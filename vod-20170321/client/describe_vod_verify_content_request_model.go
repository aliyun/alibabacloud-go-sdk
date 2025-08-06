// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodVerifyContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodVerifyContentRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeVodVerifyContentRequest
	GetOwnerId() *int64
}

type DescribeVodVerifyContentRequest struct {
	// The domain name for which you want to verify the ownership. You can specify only one domain name in each call.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeVodVerifyContentRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodVerifyContentRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodVerifyContentRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodVerifyContentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodVerifyContentRequest) SetDomainName(v string) *DescribeVodVerifyContentRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodVerifyContentRequest) SetOwnerId(v int64) *DescribeVodVerifyContentRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodVerifyContentRequest) Validate() error {
	return dara.Validate(s)
}
