// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenTrafficMirrorServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *OpenTrafficMirrorServiceRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *OpenTrafficMirrorServiceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *OpenTrafficMirrorServiceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *OpenTrafficMirrorServiceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *OpenTrafficMirrorServiceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *OpenTrafficMirrorServiceRequest
	GetResourceOwnerId() *int64
}

type OpenTrafficMirrorServiceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655442222
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

func (s OpenTrafficMirrorServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenTrafficMirrorServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenTrafficMirrorServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *OpenTrafficMirrorServiceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *OpenTrafficMirrorServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *OpenTrafficMirrorServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *OpenTrafficMirrorServiceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *OpenTrafficMirrorServiceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *OpenTrafficMirrorServiceRequest) SetClientToken(v string) *OpenTrafficMirrorServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *OpenTrafficMirrorServiceRequest) SetOwnerAccount(v string) *OpenTrafficMirrorServiceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *OpenTrafficMirrorServiceRequest) SetOwnerId(v int64) *OpenTrafficMirrorServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenTrafficMirrorServiceRequest) SetRegionId(v string) *OpenTrafficMirrorServiceRequest {
	s.RegionId = &v
	return s
}

func (s *OpenTrafficMirrorServiceRequest) SetResourceOwnerAccount(v string) *OpenTrafficMirrorServiceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OpenTrafficMirrorServiceRequest) SetResourceOwnerId(v int64) *OpenTrafficMirrorServiceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OpenTrafficMirrorServiceRequest) Validate() error {
	return dara.Validate(s)
}
