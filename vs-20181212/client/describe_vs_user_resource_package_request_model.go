// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsUserResourcePackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeVsUserResourcePackageRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeVsUserResourcePackageRequest
	GetSecurityToken() *string
}

type DescribeVsUserResourcePackageRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeVsUserResourcePackageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsUserResourcePackageRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsUserResourcePackageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVsUserResourcePackageRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeVsUserResourcePackageRequest) SetOwnerId(v int64) *DescribeVsUserResourcePackageRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsUserResourcePackageRequest) SetSecurityToken(v string) *DescribeVsUserResourcePackageRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeVsUserResourcePackageRequest) Validate() error {
	return dara.Validate(s)
}
