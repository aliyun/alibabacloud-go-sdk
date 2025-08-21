// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVsDomainDetailRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeVsDomainDetailRequest
	GetOwnerId() *int64
}

type DescribeVsDomainDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeVsDomainDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainDetailRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsDomainDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVsDomainDetailRequest) SetDomainName(v string) *DescribeVsDomainDetailRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainDetailRequest) SetOwnerId(v int64) *DescribeVsDomainDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsDomainDetailRequest) Validate() error {
	return dara.Validate(s)
}
