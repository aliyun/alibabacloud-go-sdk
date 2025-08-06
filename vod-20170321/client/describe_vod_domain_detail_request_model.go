// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainDetailRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeVodDomainDetailRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeVodDomainDetailRequest
	GetSecurityToken() *string
}

type DescribeVodDomainDetailRequest struct {
	// The domain name for CDN.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeVodDomainDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainDetailRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainDetailRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeVodDomainDetailRequest) SetDomainName(v string) *DescribeVodDomainDetailRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainDetailRequest) SetOwnerId(v int64) *DescribeVodDomainDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainDetailRequest) SetSecurityToken(v string) *DescribeVodDomainDetailRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeVodDomainDetailRequest) Validate() error {
	return dara.Validate(s)
}
