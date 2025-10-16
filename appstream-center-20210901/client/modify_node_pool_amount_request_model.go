// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNodePoolAmountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *ModifyNodePoolAmountRequest
	GetAppInstanceGroupId() *string
	SetNodePool(v *ModifyNodePoolAmountRequestNodePool) *ModifyNodePoolAmountRequest
	GetNodePool() *ModifyNodePoolAmountRequestNodePool
	SetProductType(v string) *ModifyNodePoolAmountRequest
	GetProductType() *string
}

type ModifyNodePoolAmountRequest struct {
	// The ID of the delivery group.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The parameters related to the configuration change of the node pool.
	//
	// This parameter is required.
	NodePool *ModifyNodePoolAmountRequestNodePool `json:"NodePool,omitempty" xml:"NodePool,omitempty" type:"Struct"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ModifyNodePoolAmountRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodePoolAmountRequest) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAmountRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ModifyNodePoolAmountRequest) GetNodePool() *ModifyNodePoolAmountRequestNodePool {
	return s.NodePool
}

func (s *ModifyNodePoolAmountRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ModifyNodePoolAmountRequest) SetAppInstanceGroupId(v string) *ModifyNodePoolAmountRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ModifyNodePoolAmountRequest) SetNodePool(v *ModifyNodePoolAmountRequestNodePool) *ModifyNodePoolAmountRequest {
	s.NodePool = v
	return s
}

func (s *ModifyNodePoolAmountRequest) SetProductType(v string) *ModifyNodePoolAmountRequest {
	s.ProductType = &v
	return s
}

func (s *ModifyNodePoolAmountRequest) Validate() error {
	if s.NodePool != nil {
		if err := s.NodePool.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyNodePoolAmountRequestNodePool struct {
	// The total number of subscription nodes after the change.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	NodeAmount *int32 `json:"NodeAmount,omitempty" xml:"NodeAmount,omitempty"`
	// The change mode of subscription nodes.
	//
	// Valid value:
	//
	// 	- EXPAND_FROM_POST_PAID_EXPLICIT: changes from specified pay-as-you-go nodes
	//
	// example:
	//
	// EXPAND_FROM_POST_PAID_EXPLICIT
	PrePaidNodeAmountModifyMode *string `json:"PrePaidNodeAmountModifyMode,omitempty" xml:"PrePaidNodeAmountModifyMode,omitempty"`
	// The nodes for which you want to change the billing method.
	PrePaidNodeAmountModifyNodeIds []*string `json:"PrePaidNodeAmountModifyNodeIds,omitempty" xml:"PrePaidNodeAmountModifyNodeIds,omitempty" type:"Repeated"`
}

func (s ModifyNodePoolAmountRequestNodePool) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodePoolAmountRequestNodePool) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAmountRequestNodePool) GetNodeAmount() *int32 {
	return s.NodeAmount
}

func (s *ModifyNodePoolAmountRequestNodePool) GetPrePaidNodeAmountModifyMode() *string {
	return s.PrePaidNodeAmountModifyMode
}

func (s *ModifyNodePoolAmountRequestNodePool) GetPrePaidNodeAmountModifyNodeIds() []*string {
	return s.PrePaidNodeAmountModifyNodeIds
}

func (s *ModifyNodePoolAmountRequestNodePool) SetNodeAmount(v int32) *ModifyNodePoolAmountRequestNodePool {
	s.NodeAmount = &v
	return s
}

func (s *ModifyNodePoolAmountRequestNodePool) SetPrePaidNodeAmountModifyMode(v string) *ModifyNodePoolAmountRequestNodePool {
	s.PrePaidNodeAmountModifyMode = &v
	return s
}

func (s *ModifyNodePoolAmountRequestNodePool) SetPrePaidNodeAmountModifyNodeIds(v []*string) *ModifyNodePoolAmountRequestNodePool {
	s.PrePaidNodeAmountModifyNodeIds = v
	return s
}

func (s *ModifyNodePoolAmountRequestNodePool) Validate() error {
	return dara.Validate(s)
}
