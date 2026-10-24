// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateTargetConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSelectionKey(v string) *OperateTargetConfigShrinkRequest
	GetSelectionKey() *string
	SetStatus(v string) *OperateTargetConfigShrinkRequest
	GetStatus() *string
	SetSwitchIdListShrink(v string) *OperateTargetConfigShrinkRequest
	GetSwitchIdListShrink() *string
}

type OperateTargetConfigShrinkRequest struct {
	// The key of the rule switch configuration associated with the policy. It must be the same as the key used in the ListSystemConfigRules query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8b37f46a-72d9-4c8c-bada-2f374a9c3410
	SelectionKey *string `json:"SelectionKey,omitempty" xml:"SelectionKey,omitempty"`
	// The status of the rule switch. Valid values:
	//
	// - **add**: Enabled.
	//
	// - **del**: Disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// add
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of system rule switch identifiers whose status is to be updated. The elements are obtained from the SwitchId returned by ListSystemConfigRules. A maximum of 100 items are supported. The backend updates the items one by one based on the list. The request is rejected if the list is not provided.
	SwitchIdListShrink *string `json:"SwitchIdList,omitempty" xml:"SwitchIdList,omitempty"`
}

func (s OperateTargetConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateTargetConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *OperateTargetConfigShrinkRequest) GetSelectionKey() *string {
	return s.SelectionKey
}

func (s *OperateTargetConfigShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *OperateTargetConfigShrinkRequest) GetSwitchIdListShrink() *string {
	return s.SwitchIdListShrink
}

func (s *OperateTargetConfigShrinkRequest) SetSelectionKey(v string) *OperateTargetConfigShrinkRequest {
	s.SelectionKey = &v
	return s
}

func (s *OperateTargetConfigShrinkRequest) SetStatus(v string) *OperateTargetConfigShrinkRequest {
	s.Status = &v
	return s
}

func (s *OperateTargetConfigShrinkRequest) SetSwitchIdListShrink(v string) *OperateTargetConfigShrinkRequest {
	s.SwitchIdListShrink = &v
	return s
}

func (s *OperateTargetConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
