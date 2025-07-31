// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeRegionsRequest
	GetAcceptLanguage() *string
	SetOwnerAccount(v string) *DescribeRegionsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeRegionsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeRegionsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeRegionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRegionsRequest
	GetResourceOwnerId() *int64
}

type DescribeRegionsRequest struct {
	// The language of the values of the returned **RegionName*	- and **ZoneName*	- parameters. Valid values:
	//
	// 	- **zh*	- (default)
	//
	// 	- **en**
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// >  If you do not specify this parameter, all supported regions are queried.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeRegionsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeRegionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRegionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRegionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeRegionsRequest) SetOwnerAccount(v string) *DescribeRegionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRegionsRequest) SetOwnerId(v int64) *DescribeRegionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerAccount(v string) *DescribeRegionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerId(v int64) *DescribeRegionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRegionsRequest) Validate() error {
	return dara.Validate(s)
}
