// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodUserQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeVodUserQuotaRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeVodUserQuotaRequest
	GetSecurityToken() *string
}

type DescribeVodUserQuotaRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeVodUserQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserQuotaRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodUserQuotaRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodUserQuotaRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeVodUserQuotaRequest) SetOwnerId(v int64) *DescribeVodUserQuotaRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodUserQuotaRequest) SetSecurityToken(v string) *DescribeVodUserQuotaRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeVodUserQuotaRequest) Validate() error {
	return dara.Validate(s)
}
