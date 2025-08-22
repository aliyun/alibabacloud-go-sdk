// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserLogserviceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeUserLogserviceStatusRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeUserLogserviceStatusRequest
	GetSecurityToken() *string
}

type DescribeUserLogserviceStatusRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeUserLogserviceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserLogserviceStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserLogserviceStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeUserLogserviceStatusRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeUserLogserviceStatusRequest) SetOwnerId(v int64) *DescribeUserLogserviceStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserLogserviceStatusRequest) SetSecurityToken(v string) *DescribeUserLogserviceStatusRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeUserLogserviceStatusRequest) Validate() error {
	return dara.Validate(s)
}
