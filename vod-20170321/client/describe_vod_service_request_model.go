// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeVodServiceRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeVodServiceRequest
	GetSecurityToken() *string
}

type DescribeVodServiceRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeVodServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodServiceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeVodServiceRequest) SetOwnerId(v int64) *DescribeVodServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodServiceRequest) SetSecurityToken(v string) *DescribeVodServiceRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeVodServiceRequest) Validate() error {
	return dara.Validate(s)
}
