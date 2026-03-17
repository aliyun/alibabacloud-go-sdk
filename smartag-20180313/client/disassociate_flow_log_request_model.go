// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateFlowLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowLogId(v string) *DisassociateFlowLogRequest
	GetFlowLogId() *string
	SetOwnerAccount(v string) *DisassociateFlowLogRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DisassociateFlowLogRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DisassociateFlowLogRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DisassociateFlowLogRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DisassociateFlowLogRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DisassociateFlowLogRequest
	GetSmartAGId() *string
}

type DisassociateFlowLogRequest struct {
	// The instance ID of the flow log.
	//
	// This parameter is required.
	//
	// example:
	//
	// fl-l934tsa5504yuc****
	FlowLogId    *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the flow log belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-0v3pmd7qpnvx5f****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s DisassociateFlowLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DisassociateFlowLogRequest) GoString() string {
	return s.String()
}

func (s *DisassociateFlowLogRequest) GetFlowLogId() *string {
	return s.FlowLogId
}

func (s *DisassociateFlowLogRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DisassociateFlowLogRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DisassociateFlowLogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisassociateFlowLogRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DisassociateFlowLogRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DisassociateFlowLogRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DisassociateFlowLogRequest) SetFlowLogId(v string) *DisassociateFlowLogRequest {
	s.FlowLogId = &v
	return s
}

func (s *DisassociateFlowLogRequest) SetOwnerAccount(v string) *DisassociateFlowLogRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DisassociateFlowLogRequest) SetOwnerId(v int64) *DisassociateFlowLogRequest {
	s.OwnerId = &v
	return s
}

func (s *DisassociateFlowLogRequest) SetRegionId(v string) *DisassociateFlowLogRequest {
	s.RegionId = &v
	return s
}

func (s *DisassociateFlowLogRequest) SetResourceOwnerAccount(v string) *DisassociateFlowLogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DisassociateFlowLogRequest) SetResourceOwnerId(v int64) *DisassociateFlowLogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DisassociateFlowLogRequest) SetSmartAGId(v string) *DisassociateFlowLogRequest {
	s.SmartAGId = &v
	return s
}

func (s *DisassociateFlowLogRequest) Validate() error {
	return dara.Validate(s)
}
