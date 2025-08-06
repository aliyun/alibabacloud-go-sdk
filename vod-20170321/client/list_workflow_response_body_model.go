// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListWorkflowResponseBody
	GetRequestId() *string
	SetWorkflowInfoList(v []*ListWorkflowResponseBodyWorkflowInfoList) *ListWorkflowResponseBody
	GetWorkflowInfoList() []*ListWorkflowResponseBodyWorkflowInfoList
}

type ListWorkflowResponseBody struct {
	RequestId        *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WorkflowInfoList []*ListWorkflowResponseBodyWorkflowInfoList `json:"WorkflowInfoList,omitempty" xml:"WorkflowInfoList,omitempty" type:"Repeated"`
}

func (s ListWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkflowResponseBody) GetWorkflowInfoList() []*ListWorkflowResponseBodyWorkflowInfoList {
	return s.WorkflowInfoList
}

func (s *ListWorkflowResponseBody) SetRequestId(v string) *ListWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkflowResponseBody) SetWorkflowInfoList(v []*ListWorkflowResponseBodyWorkflowInfoList) *ListWorkflowResponseBody {
	s.WorkflowInfoList = v
	return s
}

func (s *ListWorkflowResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListWorkflowResponseBodyWorkflowInfoList struct {
	ActionList   *string `json:"ActionList,omitempty" xml:"ActionList,omitempty"`
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ModifyTime   *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
	WorkflowId   *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s ListWorkflowResponseBodyWorkflowInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowResponseBodyWorkflowInfoList) GoString() string {
	return s.String()
}

func (s *ListWorkflowResponseBodyWorkflowInfoList) GetActionList() *string {
	return s.ActionList
}

func (s *ListWorkflowResponseBodyWorkflowInfoList) GetAppId() *string {
	return s.AppId
}

func (s *ListWorkflowResponseBodyWorkflowInfoList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListWorkflowResponseBodyWorkflowInfoList) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListWorkflowResponseBodyWorkflowInfoList) GetName() *string {
	return s.Name
}

func (s *ListWorkflowResponseBodyWorkflowInfoList) GetState() *string {
	return s.State
}

func (s *ListWorkflowResponseBodyWorkflowInfoList) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *ListWorkflowResponseBodyWorkflowInfoList) SetActionList(v string) *ListWorkflowResponseBodyWorkflowInfoList {
	s.ActionList = &v
	return s
}

func (s *ListWorkflowResponseBodyWorkflowInfoList) SetAppId(v string) *ListWorkflowResponseBodyWorkflowInfoList {
	s.AppId = &v
	return s
}

func (s *ListWorkflowResponseBodyWorkflowInfoList) SetCreationTime(v string) *ListWorkflowResponseBodyWorkflowInfoList {
	s.CreationTime = &v
	return s
}

func (s *ListWorkflowResponseBodyWorkflowInfoList) SetModifyTime(v string) *ListWorkflowResponseBodyWorkflowInfoList {
	s.ModifyTime = &v
	return s
}

func (s *ListWorkflowResponseBodyWorkflowInfoList) SetName(v string) *ListWorkflowResponseBodyWorkflowInfoList {
	s.Name = &v
	return s
}

func (s *ListWorkflowResponseBodyWorkflowInfoList) SetState(v string) *ListWorkflowResponseBodyWorkflowInfoList {
	s.State = &v
	return s
}

func (s *ListWorkflowResponseBodyWorkflowInfoList) SetWorkflowId(v string) *ListWorkflowResponseBodyWorkflowInfoList {
	s.WorkflowId = &v
	return s
}

func (s *ListWorkflowResponseBodyWorkflowInfoList) Validate() error {
	return dara.Validate(s)
}
