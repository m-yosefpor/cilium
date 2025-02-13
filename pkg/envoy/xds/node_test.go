// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

//go:build !privileged_tests
// +build !privileged_tests

package xds

import (
	. "gopkg.in/check.v1"
)

type NodeSuite struct{}

var _ = Suite(&NodeSuite{})

func (s *NodeSuite) TestIstioNodeToIP(c *C) {
	var ip string
	var err error

	ip, err = IstioNodeToIP("sidecar~10.1.1.0~v0.default~default.svc.cluster.local")
	c.Assert(err, IsNil)
	c.Check(ip, Equals, "10.1.1.0")

	_, err = IstioNodeToIP("sidecar~10.1.1.0~v0.default")
	c.Assert(err, Not(IsNil))

	_, err = IstioNodeToIP("sidecar~not-an-ip~v0.default~default.svc.cluster.local")
	c.Assert(err, Not(IsNil))
}
