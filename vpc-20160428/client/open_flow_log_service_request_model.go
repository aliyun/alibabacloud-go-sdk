// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenFlowLogServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *OpenFlowLogServiceRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *OpenFlowLogServiceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *OpenFlowLogServiceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *OpenFlowLogServiceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *OpenFlowLogServiceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *OpenFlowLogServiceRequest
	GetResourceOwnerId() *int64
}

type OpenFlowLogServiceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system automatically set **ClientToken*	- to the value of **RequestId**. The value of **RequestId*	- for each API request is different.
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

func (s OpenFlowLogServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenFlowLogServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenFlowLogServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *OpenFlowLogServiceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *OpenFlowLogServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *OpenFlowLogServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *OpenFlowLogServiceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *OpenFlowLogServiceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *OpenFlowLogServiceRequest) SetClientToken(v string) *OpenFlowLogServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *OpenFlowLogServiceRequest) SetOwnerAccount(v string) *OpenFlowLogServiceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *OpenFlowLogServiceRequest) SetOwnerId(v int64) *OpenFlowLogServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenFlowLogServiceRequest) SetRegionId(v string) *OpenFlowLogServiceRequest {
	s.RegionId = &v
	return s
}

func (s *OpenFlowLogServiceRequest) SetResourceOwnerAccount(v string) *OpenFlowLogServiceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OpenFlowLogServiceRequest) SetResourceOwnerId(v int64) *OpenFlowLogServiceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OpenFlowLogServiceRequest) Validate() error {
	return dara.Validate(s)
}
