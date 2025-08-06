// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserVodStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeUserVodStatusRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeUserVodStatusRequest
	GetSecurityToken() *string
}

type DescribeUserVodStatusRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeUserVodStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserVodStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserVodStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeUserVodStatusRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeUserVodStatusRequest) SetOwnerId(v int64) *DescribeUserVodStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserVodStatusRequest) SetSecurityToken(v string) *DescribeUserVodStatusRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeUserVodStatusRequest) Validate() error {
	return dara.Validate(s)
}
