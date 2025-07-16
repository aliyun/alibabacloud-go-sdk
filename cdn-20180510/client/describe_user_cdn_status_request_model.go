// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserCdnStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeUserCdnStatusRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeUserCdnStatusRequest
	GetSecurityToken() *string
}

type DescribeUserCdnStatusRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeUserCdnStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserCdnStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserCdnStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeUserCdnStatusRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeUserCdnStatusRequest) SetOwnerId(v int64) *DescribeUserCdnStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserCdnStatusRequest) SetSecurityToken(v string) *DescribeUserCdnStatusRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeUserCdnStatusRequest) Validate() error {
	return dara.Validate(s)
}
