// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowlogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *DeleteFlowlogRequest
	GetCenId() *string
	SetClientToken(v string) *DeleteFlowlogRequest
	GetClientToken() *string
	SetFlowLogId(v string) *DeleteFlowlogRequest
	GetFlowLogId() *string
	SetOwnerAccount(v string) *DeleteFlowlogRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteFlowlogRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteFlowlogRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteFlowlogRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteFlowlogRequest
	GetResourceOwnerId() *int64
}

type DeleteFlowlogRequest struct {
	// The ID of the Cloud Enterprise Network (CEN) instance.
	//
	// example:
	//
	// cen-7qthudw0ll6jmc****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, ClientToken is set to the value of RequestId. The value of RequestId may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the flow log.
	//
	// This parameter is required.
	//
	// example:
	//
	// flowlog-m5evbtbpt****
	FlowLogId    *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the flow log is deployed.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
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

func (s DeleteFlowlogRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowlogRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowlogRequest) GetCenId() *string {
	return s.CenId
}

func (s *DeleteFlowlogRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteFlowlogRequest) GetFlowLogId() *string {
	return s.FlowLogId
}

func (s *DeleteFlowlogRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteFlowlogRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteFlowlogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteFlowlogRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteFlowlogRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteFlowlogRequest) SetCenId(v string) *DeleteFlowlogRequest {
	s.CenId = &v
	return s
}

func (s *DeleteFlowlogRequest) SetClientToken(v string) *DeleteFlowlogRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteFlowlogRequest) SetFlowLogId(v string) *DeleteFlowlogRequest {
	s.FlowLogId = &v
	return s
}

func (s *DeleteFlowlogRequest) SetOwnerAccount(v string) *DeleteFlowlogRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteFlowlogRequest) SetOwnerId(v int64) *DeleteFlowlogRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteFlowlogRequest) SetRegionId(v string) *DeleteFlowlogRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteFlowlogRequest) SetResourceOwnerAccount(v string) *DeleteFlowlogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteFlowlogRequest) SetResourceOwnerId(v int64) *DeleteFlowlogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteFlowlogRequest) Validate() error {
	return dara.Validate(s)
}
