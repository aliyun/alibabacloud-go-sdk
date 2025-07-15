// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceClusterInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceClusterName(v string) *DescribeInstanceClusterInfoRequest
	GetInstanceClusterName() *string
	SetSecurityToken(v string) *DescribeInstanceClusterInfoRequest
	GetSecurityToken() *string
}

type DescribeInstanceClusterInfoRequest struct {
	// The name of the dedicated instance cluster.
	//
	// example:
	//
	// testvpc
	InstanceClusterName *string `json:"InstanceClusterName,omitempty" xml:"InstanceClusterName,omitempty"`
	SecurityToken       *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeInstanceClusterInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceClusterInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceClusterInfoRequest) GetInstanceClusterName() *string {
	return s.InstanceClusterName
}

func (s *DescribeInstanceClusterInfoRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeInstanceClusterInfoRequest) SetInstanceClusterName(v string) *DescribeInstanceClusterInfoRequest {
	s.InstanceClusterName = &v
	return s
}

func (s *DescribeInstanceClusterInfoRequest) SetSecurityToken(v string) *DescribeInstanceClusterInfoRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeInstanceClusterInfoRequest) Validate() error {
	return dara.Validate(s)
}
