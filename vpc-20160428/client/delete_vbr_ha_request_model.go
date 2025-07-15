// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVbrHaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteVbrHaRequest
	GetClientToken() *string
	SetInstanceId(v string) *DeleteVbrHaRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DeleteVbrHaRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteVbrHaRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteVbrHaRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteVbrHaRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteVbrHaRequest
	GetResourceOwnerId() *int64
}

type DeleteVbrHaRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// CBCE910E-D396-4944-8****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the VBR failover group.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbrha-sa1******
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region in which the VBR is deployed.
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

func (s DeleteVbrHaRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVbrHaRequest) GoString() string {
	return s.String()
}

func (s *DeleteVbrHaRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteVbrHaRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteVbrHaRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteVbrHaRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteVbrHaRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteVbrHaRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteVbrHaRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteVbrHaRequest) SetClientToken(v string) *DeleteVbrHaRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteVbrHaRequest) SetInstanceId(v string) *DeleteVbrHaRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteVbrHaRequest) SetOwnerAccount(v string) *DeleteVbrHaRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteVbrHaRequest) SetOwnerId(v int64) *DeleteVbrHaRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVbrHaRequest) SetRegionId(v string) *DeleteVbrHaRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVbrHaRequest) SetResourceOwnerAccount(v string) *DeleteVbrHaRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteVbrHaRequest) SetResourceOwnerId(v int64) *DeleteVbrHaRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteVbrHaRequest) Validate() error {
	return dara.Validate(s)
}
