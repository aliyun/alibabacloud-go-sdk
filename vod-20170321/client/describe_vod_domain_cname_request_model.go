// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainCnameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainCnameRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeVodDomainCnameRequest
	GetOwnerId() *int64
}

type DescribeVodDomainCnameRequest struct {
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeVodDomainCnameRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainCnameRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainCnameRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainCnameRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainCnameRequest) SetDomainName(v string) *DescribeVodDomainCnameRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainCnameRequest) SetOwnerId(v int64) *DescribeVodDomainCnameRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainCnameRequest) Validate() error {
	return dara.Validate(s)
}
