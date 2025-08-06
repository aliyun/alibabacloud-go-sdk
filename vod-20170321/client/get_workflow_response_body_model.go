// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetWorkflowResponseBody
	GetRequestId() *string
	SetWorkflowInfo(v *GetWorkflowResponseBodyWorkflowInfo) *GetWorkflowResponseBody
	GetWorkflowInfo() *GetWorkflowResponseBodyWorkflowInfo
}

type GetWorkflowResponseBody struct {
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WorkflowInfo *GetWorkflowResponseBodyWorkflowInfo `json:"WorkflowInfo,omitempty" xml:"WorkflowInfo,omitempty" type:"Struct"`
}

func (s GetWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkflowResponseBody) GetWorkflowInfo() *GetWorkflowResponseBodyWorkflowInfo {
	return s.WorkflowInfo
}

func (s *GetWorkflowResponseBody) SetRequestId(v string) *GetWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkflowResponseBody) SetWorkflowInfo(v *GetWorkflowResponseBodyWorkflowInfo) *GetWorkflowResponseBody {
	s.WorkflowInfo = v
	return s
}

func (s *GetWorkflowResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetWorkflowResponseBodyWorkflowInfo struct {
	ActionList     *string `json:"ActionList,omitempty" xml:"ActionList,omitempty"`
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CallbackConfig *string `json:"CallbackConfig,omitempty" xml:"CallbackConfig,omitempty"`
	CreationTime   *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ModifyTime     *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	State          *string `json:"State,omitempty" xml:"State,omitempty"`
	WorkflowId     *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s GetWorkflowResponseBodyWorkflowInfo) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowResponseBodyWorkflowInfo) GoString() string {
	return s.String()
}

func (s *GetWorkflowResponseBodyWorkflowInfo) GetActionList() *string {
	return s.ActionList
}

func (s *GetWorkflowResponseBodyWorkflowInfo) GetAppId() *string {
	return s.AppId
}

func (s *GetWorkflowResponseBodyWorkflowInfo) GetCallbackConfig() *string {
	return s.CallbackConfig
}

func (s *GetWorkflowResponseBodyWorkflowInfo) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetWorkflowResponseBodyWorkflowInfo) GetDescription() *string {
	return s.Description
}

func (s *GetWorkflowResponseBodyWorkflowInfo) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetWorkflowResponseBodyWorkflowInfo) GetName() *string {
	return s.Name
}

func (s *GetWorkflowResponseBodyWorkflowInfo) GetState() *string {
	return s.State
}

func (s *GetWorkflowResponseBodyWorkflowInfo) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *GetWorkflowResponseBodyWorkflowInfo) SetActionList(v string) *GetWorkflowResponseBodyWorkflowInfo {
	s.ActionList = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowInfo) SetAppId(v string) *GetWorkflowResponseBodyWorkflowInfo {
	s.AppId = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowInfo) SetCallbackConfig(v string) *GetWorkflowResponseBodyWorkflowInfo {
	s.CallbackConfig = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowInfo) SetCreationTime(v string) *GetWorkflowResponseBodyWorkflowInfo {
	s.CreationTime = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowInfo) SetDescription(v string) *GetWorkflowResponseBodyWorkflowInfo {
	s.Description = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowInfo) SetModifyTime(v string) *GetWorkflowResponseBodyWorkflowInfo {
	s.ModifyTime = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowInfo) SetName(v string) *GetWorkflowResponseBodyWorkflowInfo {
	s.Name = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowInfo) SetState(v string) *GetWorkflowResponseBodyWorkflowInfo {
	s.State = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowInfo) SetWorkflowId(v string) *GetWorkflowResponseBodyWorkflowInfo {
	s.WorkflowId = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowInfo) Validate() error {
	return dara.Validate(s)
}
