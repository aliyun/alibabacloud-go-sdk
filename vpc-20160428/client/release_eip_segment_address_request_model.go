// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseEipSegmentAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ReleaseEipSegmentAddressRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *ReleaseEipSegmentAddressRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ReleaseEipSegmentAddressRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ReleaseEipSegmentAddressRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ReleaseEipSegmentAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReleaseEipSegmentAddressRequest
	GetResourceOwnerId() *int64
	SetSegmentInstanceId(v string) *ReleaseEipSegmentAddressRequest
	GetSegmentInstanceId() *string
}

type ReleaseEipSegmentAddressRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001sdfg
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the contiguous EIPs. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the contiguous EIP group to be released.
	//
	// The system releases all EIPs in the group.
	//
	// This parameter is required.
	//
	// example:
	//
	// eipsg-2zett8ba055tbsxme****
	SegmentInstanceId *string `json:"SegmentInstanceId,omitempty" xml:"SegmentInstanceId,omitempty"`
}

func (s ReleaseEipSegmentAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseEipSegmentAddressRequest) GoString() string {
	return s.String()
}

func (s *ReleaseEipSegmentAddressRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ReleaseEipSegmentAddressRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ReleaseEipSegmentAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReleaseEipSegmentAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReleaseEipSegmentAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReleaseEipSegmentAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReleaseEipSegmentAddressRequest) GetSegmentInstanceId() *string {
	return s.SegmentInstanceId
}

func (s *ReleaseEipSegmentAddressRequest) SetClientToken(v string) *ReleaseEipSegmentAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *ReleaseEipSegmentAddressRequest) SetOwnerAccount(v string) *ReleaseEipSegmentAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReleaseEipSegmentAddressRequest) SetOwnerId(v int64) *ReleaseEipSegmentAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseEipSegmentAddressRequest) SetRegionId(v string) *ReleaseEipSegmentAddressRequest {
	s.RegionId = &v
	return s
}

func (s *ReleaseEipSegmentAddressRequest) SetResourceOwnerAccount(v string) *ReleaseEipSegmentAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleaseEipSegmentAddressRequest) SetResourceOwnerId(v int64) *ReleaseEipSegmentAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleaseEipSegmentAddressRequest) SetSegmentInstanceId(v string) *ReleaseEipSegmentAddressRequest {
	s.SegmentInstanceId = &v
	return s
}

func (s *ReleaseEipSegmentAddressRequest) Validate() error {
	return dara.Validate(s)
}
