// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFlowShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoriesShrink(v string) *ModifyFlowShrinkRequest
	GetCategoriesShrink() *string
	SetCustSpaceId(v string) *ModifyFlowShrinkRequest
	GetCustSpaceId() *string
	SetFlowId(v string) *ModifyFlowShrinkRequest
	GetFlowId() *string
	SetFlowName(v string) *ModifyFlowShrinkRequest
	GetFlowName() *string
	SetOwnerId(v int64) *ModifyFlowShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyFlowShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyFlowShrinkRequest
	GetResourceOwnerId() *int64
}

type ModifyFlowShrinkRequest struct {
	// This parameter is required.
	CategoriesShrink *string `json:"Categories,omitempty" xml:"Categories,omitempty"`
	// example:
	//
	// 示例值示例值
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// example:
	//
	// 示例值
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	FlowName             *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyFlowShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyFlowShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyFlowShrinkRequest) GetCategoriesShrink() *string {
	return s.CategoriesShrink
}

func (s *ModifyFlowShrinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ModifyFlowShrinkRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *ModifyFlowShrinkRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *ModifyFlowShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyFlowShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyFlowShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyFlowShrinkRequest) SetCategoriesShrink(v string) *ModifyFlowShrinkRequest {
	s.CategoriesShrink = &v
	return s
}

func (s *ModifyFlowShrinkRequest) SetCustSpaceId(v string) *ModifyFlowShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ModifyFlowShrinkRequest) SetFlowId(v string) *ModifyFlowShrinkRequest {
	s.FlowId = &v
	return s
}

func (s *ModifyFlowShrinkRequest) SetFlowName(v string) *ModifyFlowShrinkRequest {
	s.FlowName = &v
	return s
}

func (s *ModifyFlowShrinkRequest) SetOwnerId(v int64) *ModifyFlowShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyFlowShrinkRequest) SetResourceOwnerAccount(v string) *ModifyFlowShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyFlowShrinkRequest) SetResourceOwnerId(v int64) *ModifyFlowShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyFlowShrinkRequest) Validate() error {
	return dara.Validate(s)
}
