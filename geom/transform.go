// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import (
	"math"
)

type Transform struct {
	Rotation,
	ScaleX, ScaleY,
	TranslateX, TranslateY float64
}

func (t Transform) Matrix() Matrix {
	cos, sin := math.Cos(t.Rotation), math.Sin(t.Rotation)
	return TransformMatrix(
		cos*t.ScaleX,
		sin*t.ScaleY,
		-sin*t.ScaleX,
		cos*t.ScaleY,
		t.TranslateX,
		t.TranslateY,
	)
}

// Matrix is a 3x3 matrix of float64 values.
type Matrix [9]float64
type MatrixElement int

//goland:noinspection GoSnakeCaseUsage
const (
	ME_A  MatrixElement = 0
	ME_B                = 3
	ME_C                = 1
	ME_D                = 4
	ME_TX               = 2
	ME_TY               = 5
)

// TransformMatrix creates a Matrix with the following values:
//   [ a, c, tx,
//     b, d, ty,
//     0, 0, 1 ]
//
// The values A and D affect the positioning of pixels along the x and y axis
// when scaling or rotating. B and C are the values that affect the positioning
// of pixels along the x and y axis when rotating or skewing.
// TX and TY are the distances by which to translate each point along the x and
// y axis.
// The Matrix operates in 2D space so it always assumes that the (last three)
// property values u and v are 0.0, and w is 1.0.
// https://en.wikipedia.org/wiki/Transformation_matrix
// https://www.tutorialspoint.com/computer_graphics/2d_transformation.htm
func TransformMatrix(a, b, c, d, tx, ty float64) Matrix {
	return Matrix{a, c, tx, b, d, ty, 0, 0, 1}
}

func ScaleMatrix(x, y float64) Matrix {
	return Matrix{x, 0, 0, 0, y, 0, 0, 0, 1}
}

// RotationMatrix creates a Matrix which rotates the target by an angle,
// measured in radians.
// https://en.wikipedia.org/wiki/Rotation_matrix
func RotationMatrix(angle float64) Matrix {
	cos, sin := math.Cos(angle), math.Sin(angle)
	return TransformMatrix(cos, sin, -sin, cos, 0, 0)
}

// SkewMatrix creates a Matrix which progressively slides the target in a
// direction parallel to the x or y axis.
func SkewMatrix(x, y float64) Matrix {
	return Matrix{0, math.Tan(x), 0, math.Tan(y), 1, 0, 0, 0, 1}
}

// IdentityMatrix creates a Matrix with values that cause no transformation to
// the target.
func IdentityMatrix() Matrix {
	return Matrix{1, 0, 0, 0, 1, 0, 0, 0, 1}
}
