// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionType(v int32) *ModifyPolicyShrinkRequest
	GetActionType() *int32
	SetContentShrink(v string) *ModifyPolicyShrinkRequest
	GetContentShrink() *string
	SetId(v string) *ModifyPolicyShrinkRequest
	GetId() *string
	SetName(v string) *ModifyPolicyShrinkRequest
	GetName() *string
	SetPortVersion(v string) *ModifyPolicyShrinkRequest
	GetPortVersion() *string
}

type ModifyPolicyShrinkRequest struct {
	// The action type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11
	ActionType *int32 `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The policy content.
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The policy ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c52c2fa6-fdac-40c4-8753-be7c********
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The policy name.
	//
	// example:
	//
	// demo**
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version of the port-specific mitigation policy. Valid values:
	//
	// example:
	//
	// 2
	PortVersion *string `json:"PortVersion,omitempty" xml:"PortVersion,omitempty"`
}

func (s ModifyPolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolicyShrinkRequest) GetActionType() *int32 {
	return s.ActionType
}

func (s *ModifyPolicyShrinkRequest) GetContentShrink() *string {
	return s.ContentShrink
}

func (s *ModifyPolicyShrinkRequest) GetId() *string {
	return s.Id
}

func (s *ModifyPolicyShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ModifyPolicyShrinkRequest) GetPortVersion() *string {
	return s.PortVersion
}

func (s *ModifyPolicyShrinkRequest) SetActionType(v int32) *ModifyPolicyShrinkRequest {
	s.ActionType = &v
	return s
}

func (s *ModifyPolicyShrinkRequest) SetContentShrink(v string) *ModifyPolicyShrinkRequest {
	s.ContentShrink = &v
	return s
}

func (s *ModifyPolicyShrinkRequest) SetId(v string) *ModifyPolicyShrinkRequest {
	s.Id = &v
	return s
}

func (s *ModifyPolicyShrinkRequest) SetName(v string) *ModifyPolicyShrinkRequest {
	s.Name = &v
	return s
}

func (s *ModifyPolicyShrinkRequest) SetPortVersion(v string) *ModifyPolicyShrinkRequest {
	s.PortVersion = &v
	return s
}

func (s *ModifyPolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
