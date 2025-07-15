// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrafficMirrorServiceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetTrafficMirrorServiceStatusRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *GetTrafficMirrorServiceStatusRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetTrafficMirrorServiceStatusRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetTrafficMirrorServiceStatusRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetTrafficMirrorServiceStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetTrafficMirrorServiceStatusRequest
	GetResourceOwnerId() *int64
}

type GetTrafficMirrorServiceStatusRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **client token*	- can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the mirrored traffic belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetTrafficMirrorServiceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTrafficMirrorServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *GetTrafficMirrorServiceStatusRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetTrafficMirrorServiceStatusRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetTrafficMirrorServiceStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetTrafficMirrorServiceStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTrafficMirrorServiceStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetTrafficMirrorServiceStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetTrafficMirrorServiceStatusRequest) SetClientToken(v string) *GetTrafficMirrorServiceStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *GetTrafficMirrorServiceStatusRequest) SetOwnerAccount(v string) *GetTrafficMirrorServiceStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetTrafficMirrorServiceStatusRequest) SetOwnerId(v int64) *GetTrafficMirrorServiceStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *GetTrafficMirrorServiceStatusRequest) SetRegionId(v string) *GetTrafficMirrorServiceStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetTrafficMirrorServiceStatusRequest) SetResourceOwnerAccount(v string) *GetTrafficMirrorServiceStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetTrafficMirrorServiceStatusRequest) SetResourceOwnerId(v int64) *GetTrafficMirrorServiceStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetTrafficMirrorServiceStatusRequest) Validate() error {
	return dara.Validate(s)
}
