// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateTargetConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSelectionKey(v string) *OperateTargetConfigRequest
	GetSelectionKey() *string
	SetStatus(v string) *OperateTargetConfigRequest
	GetStatus() *string
	SetSwitchIdList(v []*string) *OperateTargetConfigRequest
	GetSwitchIdList() []*string
}

type OperateTargetConfigRequest struct {
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
	SwitchIdList []*string `json:"SwitchIdList,omitempty" xml:"SwitchIdList,omitempty" type:"Repeated"`
}

func (s OperateTargetConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateTargetConfigRequest) GoString() string {
	return s.String()
}

func (s *OperateTargetConfigRequest) GetSelectionKey() *string {
	return s.SelectionKey
}

func (s *OperateTargetConfigRequest) GetStatus() *string {
	return s.Status
}

func (s *OperateTargetConfigRequest) GetSwitchIdList() []*string {
	return s.SwitchIdList
}

func (s *OperateTargetConfigRequest) SetSelectionKey(v string) *OperateTargetConfigRequest {
	s.SelectionKey = &v
	return s
}

func (s *OperateTargetConfigRequest) SetStatus(v string) *OperateTargetConfigRequest {
	s.Status = &v
	return s
}

func (s *OperateTargetConfigRequest) SetSwitchIdList(v []*string) *OperateTargetConfigRequest {
	s.SwitchIdList = v
	return s
}

func (s *OperateTargetConfigRequest) Validate() error {
	return dara.Validate(s)
}
