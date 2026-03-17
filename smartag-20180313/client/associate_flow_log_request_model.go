// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateFlowLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowLogId(v string) *AssociateFlowLogRequest
	GetFlowLogId() *string
	SetOwnerAccount(v string) *AssociateFlowLogRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AssociateFlowLogRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AssociateFlowLogRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AssociateFlowLogRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AssociateFlowLogRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *AssociateFlowLogRequest
	GetSmartAGId() *string
}

type AssociateFlowLogRequest struct {
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
	// The ID of the SAG instance to be associated.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-39u6j9a49i03wk****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s AssociateFlowLogRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateFlowLogRequest) GoString() string {
	return s.String()
}

func (s *AssociateFlowLogRequest) GetFlowLogId() *string {
	return s.FlowLogId
}

func (s *AssociateFlowLogRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AssociateFlowLogRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AssociateFlowLogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssociateFlowLogRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AssociateFlowLogRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AssociateFlowLogRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *AssociateFlowLogRequest) SetFlowLogId(v string) *AssociateFlowLogRequest {
	s.FlowLogId = &v
	return s
}

func (s *AssociateFlowLogRequest) SetOwnerAccount(v string) *AssociateFlowLogRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AssociateFlowLogRequest) SetOwnerId(v int64) *AssociateFlowLogRequest {
	s.OwnerId = &v
	return s
}

func (s *AssociateFlowLogRequest) SetRegionId(v string) *AssociateFlowLogRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateFlowLogRequest) SetResourceOwnerAccount(v string) *AssociateFlowLogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AssociateFlowLogRequest) SetResourceOwnerId(v int64) *AssociateFlowLogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AssociateFlowLogRequest) SetSmartAGId(v string) *AssociateFlowLogRequest {
	s.SmartAGId = &v
	return s
}

func (s *AssociateFlowLogRequest) Validate() error {
	return dara.Validate(s)
}
