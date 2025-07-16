// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRefreshQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeRefreshQuotaRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeRefreshQuotaRequest
	GetSecurityToken() *string
}

type DescribeRefreshQuotaRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeRefreshQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRefreshQuotaRequest) GoString() string {
	return s.String()
}

func (s *DescribeRefreshQuotaRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRefreshQuotaRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeRefreshQuotaRequest) SetOwnerId(v int64) *DescribeRefreshQuotaRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRefreshQuotaRequest) SetSecurityToken(v string) *DescribeRefreshQuotaRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeRefreshQuotaRequest) Validate() error {
	return dara.Validate(s)
}
