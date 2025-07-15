// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFlowLogServiceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetFlowLogServiceStatusRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *GetFlowLogServiceStatusRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetFlowLogServiceStatusRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetFlowLogServiceStatusRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetFlowLogServiceStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetFlowLogServiceStatusRequest
	GetResourceOwnerId() *int64
}

type GetFlowLogServiceStatusRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system automatically uses **RequestId*	- as **ClientToken**. The value of **RequestId*	- in each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the flow log.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
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

func (s GetFlowLogServiceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFlowLogServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *GetFlowLogServiceStatusRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetFlowLogServiceStatusRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetFlowLogServiceStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetFlowLogServiceStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetFlowLogServiceStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetFlowLogServiceStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetFlowLogServiceStatusRequest) SetClientToken(v string) *GetFlowLogServiceStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *GetFlowLogServiceStatusRequest) SetOwnerAccount(v string) *GetFlowLogServiceStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetFlowLogServiceStatusRequest) SetOwnerId(v int64) *GetFlowLogServiceStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *GetFlowLogServiceStatusRequest) SetRegionId(v string) *GetFlowLogServiceStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetFlowLogServiceStatusRequest) SetResourceOwnerAccount(v string) *GetFlowLogServiceStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetFlowLogServiceStatusRequest) SetResourceOwnerId(v int64) *GetFlowLogServiceStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetFlowLogServiceStatusRequest) Validate() error {
	return dara.Validate(s)
}
