// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveFlowLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *ActiveFlowLogRequest
	GetCenId() *string
	SetClientToken(v string) *ActiveFlowLogRequest
	GetClientToken() *string
	SetFlowLogId(v string) *ActiveFlowLogRequest
	GetFlowLogId() *string
	SetOwnerAccount(v string) *ActiveFlowLogRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ActiveFlowLogRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ActiveFlowLogRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ActiveFlowLogRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ActiveFlowLogRequest
	GetResourceOwnerId() *int64
}

type ActiveFlowLogRequest struct {
	// The ID of the Cloud Enterprise Network (CEN) instance.
	//
	// example:
	//
	// cen-7qthudw0ll6jmc****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that the value is unique among different requests. The client token can contain only ASCII characters.
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

func (s ActiveFlowLogRequest) String() string {
	return dara.Prettify(s)
}

func (s ActiveFlowLogRequest) GoString() string {
	return s.String()
}

func (s *ActiveFlowLogRequest) GetCenId() *string {
	return s.CenId
}

func (s *ActiveFlowLogRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ActiveFlowLogRequest) GetFlowLogId() *string {
	return s.FlowLogId
}

func (s *ActiveFlowLogRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ActiveFlowLogRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ActiveFlowLogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ActiveFlowLogRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ActiveFlowLogRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ActiveFlowLogRequest) SetCenId(v string) *ActiveFlowLogRequest {
	s.CenId = &v
	return s
}

func (s *ActiveFlowLogRequest) SetClientToken(v string) *ActiveFlowLogRequest {
	s.ClientToken = &v
	return s
}

func (s *ActiveFlowLogRequest) SetFlowLogId(v string) *ActiveFlowLogRequest {
	s.FlowLogId = &v
	return s
}

func (s *ActiveFlowLogRequest) SetOwnerAccount(v string) *ActiveFlowLogRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ActiveFlowLogRequest) SetOwnerId(v int64) *ActiveFlowLogRequest {
	s.OwnerId = &v
	return s
}

func (s *ActiveFlowLogRequest) SetRegionId(v string) *ActiveFlowLogRequest {
	s.RegionId = &v
	return s
}

func (s *ActiveFlowLogRequest) SetResourceOwnerAccount(v string) *ActiveFlowLogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ActiveFlowLogRequest) SetResourceOwnerId(v int64) *ActiveFlowLogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ActiveFlowLogRequest) Validate() error {
	return dara.Validate(s)
}
