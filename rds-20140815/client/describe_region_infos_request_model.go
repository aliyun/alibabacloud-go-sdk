// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeRegionInfosRequest
	GetClientToken() *string
	SetOwnerId(v int64) *DescribeRegionInfosRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeRegionInfosRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeRegionInfosRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRegionInfosRequest
	GetResourceOwnerId() *int64
}

type DescribeRegionInfosRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the generated token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz*****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeRegionInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionInfosRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionInfosRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeRegionInfosRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRegionInfosRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionInfosRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRegionInfosRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRegionInfosRequest) SetClientToken(v string) *DescribeRegionInfosRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeRegionInfosRequest) SetOwnerId(v int64) *DescribeRegionInfosRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRegionInfosRequest) SetRegionId(v string) *DescribeRegionInfosRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionInfosRequest) SetResourceOwnerAccount(v string) *DescribeRegionInfosRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRegionInfosRequest) SetResourceOwnerId(v int64) *DescribeRegionInfosRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRegionInfosRequest) Validate() error {
	return dara.Validate(s)
}
