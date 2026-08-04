// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGwConsumerOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateGwConsumerOrderRequest
	GetClientToken() *string
	SetExpireTime(v string) *CreateGwConsumerOrderRequest
	GetExpireTime() *string
	SetGatewayId(v string) *CreateGwConsumerOrderRequest
	GetGatewayId() *string
	SetKeyCount(v int32) *CreateGwConsumerOrderRequest
	GetKeyCount() *int32
	SetPackageSpec(v string) *CreateGwConsumerOrderRequest
	GetPackageSpec() *string
	SetRegionId(v string) *CreateGwConsumerOrderRequest
	GetRegionId() *string
}

type CreateGwConsumerOrderRequest struct {
	// The idempotency token.
	//
	// example:
	//
	// 6000170000591aed949d0f54a343f1a42***********
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The expiration time of the API key in ISO-8601 format. The value must be later than the current time.
	//
	// example:
	//
	// 2027-07-23T03:09:08Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the AI gateway instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-2ze24rr575j5b18cg
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The number of API keys to generate (the number of capacity plans to order). Valid values: 1 to 30.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6
	KeyCount *int32 `json:"KeyCount,omitempty" xml:"KeyCount,omitempty"`
	// The number of credits per API key. The value is a positive integer string.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3000
	PackageSpec *string `json:"PackageSpec,omitempty" xml:"PackageSpec,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query region information.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateGwConsumerOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGwConsumerOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateGwConsumerOrderRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateGwConsumerOrderRequest) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *CreateGwConsumerOrderRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *CreateGwConsumerOrderRequest) GetKeyCount() *int32 {
	return s.KeyCount
}

func (s *CreateGwConsumerOrderRequest) GetPackageSpec() *string {
	return s.PackageSpec
}

func (s *CreateGwConsumerOrderRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateGwConsumerOrderRequest) SetClientToken(v string) *CreateGwConsumerOrderRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateGwConsumerOrderRequest) SetExpireTime(v string) *CreateGwConsumerOrderRequest {
	s.ExpireTime = &v
	return s
}

func (s *CreateGwConsumerOrderRequest) SetGatewayId(v string) *CreateGwConsumerOrderRequest {
	s.GatewayId = &v
	return s
}

func (s *CreateGwConsumerOrderRequest) SetKeyCount(v int32) *CreateGwConsumerOrderRequest {
	s.KeyCount = &v
	return s
}

func (s *CreateGwConsumerOrderRequest) SetPackageSpec(v string) *CreateGwConsumerOrderRequest {
	s.PackageSpec = &v
	return s
}

func (s *CreateGwConsumerOrderRequest) SetRegionId(v string) *CreateGwConsumerOrderRequest {
	s.RegionId = &v
	return s
}

func (s *CreateGwConsumerOrderRequest) Validate() error {
	return dara.Validate(s)
}
