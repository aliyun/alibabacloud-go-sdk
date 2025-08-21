// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsVerifyContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVsVerifyContentRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeVsVerifyContentRequest
	GetOwnerId() *int64
}

type DescribeVsVerifyContentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeVsVerifyContentRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsVerifyContentRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsVerifyContentRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsVerifyContentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVsVerifyContentRequest) SetDomainName(v string) *DescribeVsVerifyContentRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsVerifyContentRequest) SetOwnerId(v int64) *DescribeVsVerifyContentRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsVerifyContentRequest) Validate() error {
	return dara.Validate(s)
}
