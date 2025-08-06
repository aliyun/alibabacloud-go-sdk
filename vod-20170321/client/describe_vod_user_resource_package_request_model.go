// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodUserResourcePackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeVodUserResourcePackageRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeVodUserResourcePackageRequest
	GetSecurityToken() *string
}

type DescribeVodUserResourcePackageRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeVodUserResourcePackageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserResourcePackageRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodUserResourcePackageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodUserResourcePackageRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeVodUserResourcePackageRequest) SetOwnerId(v int64) *DescribeVodUserResourcePackageRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodUserResourcePackageRequest) SetSecurityToken(v string) *DescribeVodUserResourcePackageRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeVodUserResourcePackageRequest) Validate() error {
	return dara.Validate(s)
}
