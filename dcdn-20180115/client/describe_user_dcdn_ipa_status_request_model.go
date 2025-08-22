// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserDcdnIpaStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeUserDcdnIpaStatusRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeUserDcdnIpaStatusRequest
	GetSecurityToken() *string
}

type DescribeUserDcdnIpaStatusRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeUserDcdnIpaStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserDcdnIpaStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserDcdnIpaStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeUserDcdnIpaStatusRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeUserDcdnIpaStatusRequest) SetOwnerId(v int64) *DescribeUserDcdnIpaStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserDcdnIpaStatusRequest) SetSecurityToken(v string) *DescribeUserDcdnIpaStatusRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeUserDcdnIpaStatusRequest) Validate() error {
	return dara.Validate(s)
}
