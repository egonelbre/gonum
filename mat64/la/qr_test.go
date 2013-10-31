// Copyright ©2013 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package la

import (
	"github.com/gonum/matrix/mat64"

	check "launchpad.net/gocheck"
)

func (s *S) TestQRD(c *check.C) {
	for _, t := range []struct {
		a *mat64.Dense
	}{
		{
			a: mustDense(mat64.NewDense(4, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})),
		},
	} {
		qf := QR(mat64.DenseCopyOf(t.a))
		r := qf.R()
		q := qf.Q()

		q.Mul(q, r)
		c.Check(t.a.EqualsApprox(q, 1e-12), check.Equals, true)
	}
}
