// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNodePoolAmountShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *ModifyNodePoolAmountShrinkRequest
	GetAppInstanceGroupId() *string
	SetNodePoolShrink(v string) *ModifyNodePoolAmountShrinkRequest
	GetNodePoolShrink() *string
	SetProductType(v string) *ModifyNodePoolAmountShrinkRequest
	GetProductType() *string
}

type ModifyNodePoolAmountShrinkRequest struct {
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
	NodePoolShrink *string `json:"NodePool,omitempty" xml:"NodePool,omitempty"`
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

func (s ModifyNodePoolAmountShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodePoolAmountShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAmountShrinkRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ModifyNodePoolAmountShrinkRequest) GetNodePoolShrink() *string {
	return s.NodePoolShrink
}

func (s *ModifyNodePoolAmountShrinkRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ModifyNodePoolAmountShrinkRequest) SetAppInstanceGroupId(v string) *ModifyNodePoolAmountShrinkRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ModifyNodePoolAmountShrinkRequest) SetNodePoolShrink(v string) *ModifyNodePoolAmountShrinkRequest {
	s.NodePoolShrink = &v
	return s
}

func (s *ModifyNodePoolAmountShrinkRequest) SetProductType(v string) *ModifyNodePoolAmountShrinkRequest {
	s.ProductType = &v
	return s
}

func (s *ModifyNodePoolAmountShrinkRequest) Validate() error {
	return dara.Validate(s)
}
