// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomizeFlowStrategyListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductType(v string) *ModifyCustomizeFlowStrategyListShrinkRequest
	GetProductType() *string
	SetStrategyObjectShrink(v string) *ModifyCustomizeFlowStrategyListShrinkRequest
	GetStrategyObjectShrink() *string
}

type ModifyCustomizeFlowStrategyListShrinkRequest struct {
	// Product type, currently only supports **ANT_CLOUD_AUTH*	- (Financial-grade real person), all others have been phased out.
	//
	// example:
	//
	// ANT_CLOUD_AUTH
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// Strategy list.
	StrategyObjectShrink *string `json:"StrategyObject,omitempty" xml:"StrategyObject,omitempty"`
}

func (s ModifyCustomizeFlowStrategyListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomizeFlowStrategyListShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyCustomizeFlowStrategyListShrinkRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ModifyCustomizeFlowStrategyListShrinkRequest) GetStrategyObjectShrink() *string {
	return s.StrategyObjectShrink
}

func (s *ModifyCustomizeFlowStrategyListShrinkRequest) SetProductType(v string) *ModifyCustomizeFlowStrategyListShrinkRequest {
	s.ProductType = &v
	return s
}

func (s *ModifyCustomizeFlowStrategyListShrinkRequest) SetStrategyObjectShrink(v string) *ModifyCustomizeFlowStrategyListShrinkRequest {
	s.StrategyObjectShrink = &v
	return s
}

func (s *ModifyCustomizeFlowStrategyListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
