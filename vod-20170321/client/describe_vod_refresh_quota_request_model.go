// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodRefreshQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeVodRefreshQuotaRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeVodRefreshQuotaRequest
	GetSecurityToken() *string
}

type DescribeVodRefreshQuotaRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeVodRefreshQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodRefreshQuotaRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodRefreshQuotaRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodRefreshQuotaRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeVodRefreshQuotaRequest) SetOwnerId(v int64) *DescribeVodRefreshQuotaRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodRefreshQuotaRequest) SetSecurityToken(v string) *DescribeVodRefreshQuotaRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeVodRefreshQuotaRequest) Validate() error {
	return dara.Validate(s)
}
