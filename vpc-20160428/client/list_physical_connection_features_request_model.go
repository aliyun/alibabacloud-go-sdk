// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPhysicalConnectionFeaturesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ListPhysicalConnectionFeaturesRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *ListPhysicalConnectionFeaturesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListPhysicalConnectionFeaturesRequest
	GetOwnerId() *int64
	SetPhysicalConnectionId(v string) *ListPhysicalConnectionFeaturesRequest
	GetPhysicalConnectionId() *string
	SetRegionId(v string) *ListPhysicalConnectionFeaturesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListPhysicalConnectionFeaturesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListPhysicalConnectionFeaturesRequest
	GetResourceOwnerId() *int64
}

type ListPhysicalConnectionFeaturesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **RequestId*	- as the **ClientToken**. The **RequestId*	- may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the Express Connect circuit.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp1qrb3044eqixog****
	PhysicalConnectionId *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
	// The region ID of the Express Connect circuit.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListPhysicalConnectionFeaturesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPhysicalConnectionFeaturesRequest) GoString() string {
	return s.String()
}

func (s *ListPhysicalConnectionFeaturesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListPhysicalConnectionFeaturesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListPhysicalConnectionFeaturesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListPhysicalConnectionFeaturesRequest) GetPhysicalConnectionId() *string {
	return s.PhysicalConnectionId
}

func (s *ListPhysicalConnectionFeaturesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPhysicalConnectionFeaturesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListPhysicalConnectionFeaturesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListPhysicalConnectionFeaturesRequest) SetClientToken(v string) *ListPhysicalConnectionFeaturesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListPhysicalConnectionFeaturesRequest) SetOwnerAccount(v string) *ListPhysicalConnectionFeaturesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListPhysicalConnectionFeaturesRequest) SetOwnerId(v int64) *ListPhysicalConnectionFeaturesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListPhysicalConnectionFeaturesRequest) SetPhysicalConnectionId(v string) *ListPhysicalConnectionFeaturesRequest {
	s.PhysicalConnectionId = &v
	return s
}

func (s *ListPhysicalConnectionFeaturesRequest) SetRegionId(v string) *ListPhysicalConnectionFeaturesRequest {
	s.RegionId = &v
	return s
}

func (s *ListPhysicalConnectionFeaturesRequest) SetResourceOwnerAccount(v string) *ListPhysicalConnectionFeaturesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListPhysicalConnectionFeaturesRequest) SetResourceOwnerId(v int64) *ListPhysicalConnectionFeaturesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListPhysicalConnectionFeaturesRequest) Validate() error {
	return dara.Validate(s)
}
