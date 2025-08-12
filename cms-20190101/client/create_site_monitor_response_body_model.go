// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSiteMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSiteMonitorResponseBody
	GetCode() *string
	SetCreateResultList(v *CreateSiteMonitorResponseBodyCreateResultList) *CreateSiteMonitorResponseBody
	GetCreateResultList() *CreateSiteMonitorResponseBodyCreateResultList
	SetData(v *CreateSiteMonitorResponseBodyData) *CreateSiteMonitorResponseBody
	GetData() *CreateSiteMonitorResponseBodyData
	SetMessage(v string) *CreateSiteMonitorResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSiteMonitorResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateSiteMonitorResponseBody
	GetSuccess() *string
}

type CreateSiteMonitorResponseBody struct {
	// The HTTP status code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	//
	// If a site monitoring task is created, the result is returned. If a site monitoring task fails to be created, no result is returned.
	CreateResultList *CreateSiteMonitorResponseBodyCreateResultList `json:"CreateResultList,omitempty" xml:"CreateResultList,omitempty" type:"Struct"`
	// The result of the site monitoring task.
	Data *CreateSiteMonitorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 68192f5d-0d45-4b98-9724-892813f86c71
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSiteMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSiteMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSiteMonitorResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSiteMonitorResponseBody) GetCreateResultList() *CreateSiteMonitorResponseBodyCreateResultList {
	return s.CreateResultList
}

func (s *CreateSiteMonitorResponseBody) GetData() *CreateSiteMonitorResponseBodyData {
	return s.Data
}

func (s *CreateSiteMonitorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSiteMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSiteMonitorResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateSiteMonitorResponseBody) SetCode(v string) *CreateSiteMonitorResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSiteMonitorResponseBody) SetCreateResultList(v *CreateSiteMonitorResponseBodyCreateResultList) *CreateSiteMonitorResponseBody {
	s.CreateResultList = v
	return s
}

func (s *CreateSiteMonitorResponseBody) SetData(v *CreateSiteMonitorResponseBodyData) *CreateSiteMonitorResponseBody {
	s.Data = v
	return s
}

func (s *CreateSiteMonitorResponseBody) SetMessage(v string) *CreateSiteMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSiteMonitorResponseBody) SetRequestId(v string) *CreateSiteMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSiteMonitorResponseBody) SetSuccess(v string) *CreateSiteMonitorResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSiteMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateSiteMonitorResponseBodyCreateResultList struct {
	CreateResultList []*CreateSiteMonitorResponseBodyCreateResultListCreateResultList `json:"CreateResultList,omitempty" xml:"CreateResultList,omitempty" type:"Repeated"`
}

func (s CreateSiteMonitorResponseBodyCreateResultList) String() string {
	return dara.Prettify(s)
}

func (s CreateSiteMonitorResponseBodyCreateResultList) GoString() string {
	return s.String()
}

func (s *CreateSiteMonitorResponseBodyCreateResultList) GetCreateResultList() []*CreateSiteMonitorResponseBodyCreateResultListCreateResultList {
	return s.CreateResultList
}

func (s *CreateSiteMonitorResponseBodyCreateResultList) SetCreateResultList(v []*CreateSiteMonitorResponseBodyCreateResultListCreateResultList) *CreateSiteMonitorResponseBodyCreateResultList {
	s.CreateResultList = v
	return s
}

func (s *CreateSiteMonitorResponseBodyCreateResultList) Validate() error {
	return dara.Validate(s)
}

type CreateSiteMonitorResponseBodyCreateResultListCreateResultList struct {
	// The ID of the site monitoring task.
	//
	// example:
	//
	// 2c8dbdf9-a3ab-46a1-85a4-f094965e****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the site monitoring task.
	//
	// example:
	//
	// HanZhou_ECS1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateSiteMonitorResponseBodyCreateResultListCreateResultList) String() string {
	return dara.Prettify(s)
}

func (s CreateSiteMonitorResponseBodyCreateResultListCreateResultList) GoString() string {
	return s.String()
}

func (s *CreateSiteMonitorResponseBodyCreateResultListCreateResultList) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateSiteMonitorResponseBodyCreateResultListCreateResultList) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateSiteMonitorResponseBodyCreateResultListCreateResultList) SetTaskId(v string) *CreateSiteMonitorResponseBodyCreateResultListCreateResultList {
	s.TaskId = &v
	return s
}

func (s *CreateSiteMonitorResponseBodyCreateResultListCreateResultList) SetTaskName(v string) *CreateSiteMonitorResponseBodyCreateResultListCreateResultList {
	s.TaskName = &v
	return s
}

func (s *CreateSiteMonitorResponseBodyCreateResultListCreateResultList) Validate() error {
	return dara.Validate(s)
}

type CreateSiteMonitorResponseBodyData struct {
	// The result that is returned after you associate the existing alert rule with the site monitoring task.
	AttachAlertResult *CreateSiteMonitorResponseBodyDataAttachAlertResult `json:"AttachAlertResult,omitempty" xml:"AttachAlertResult,omitempty" type:"Struct"`
}

func (s CreateSiteMonitorResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateSiteMonitorResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSiteMonitorResponseBodyData) GetAttachAlertResult() *CreateSiteMonitorResponseBodyDataAttachAlertResult {
	return s.AttachAlertResult
}

func (s *CreateSiteMonitorResponseBodyData) SetAttachAlertResult(v *CreateSiteMonitorResponseBodyDataAttachAlertResult) *CreateSiteMonitorResponseBodyData {
	s.AttachAlertResult = v
	return s
}

func (s *CreateSiteMonitorResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type CreateSiteMonitorResponseBodyDataAttachAlertResult struct {
	Contact []*CreateSiteMonitorResponseBodyDataAttachAlertResultContact `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Repeated"`
}

func (s CreateSiteMonitorResponseBodyDataAttachAlertResult) String() string {
	return dara.Prettify(s)
}

func (s CreateSiteMonitorResponseBodyDataAttachAlertResult) GoString() string {
	return s.String()
}

func (s *CreateSiteMonitorResponseBodyDataAttachAlertResult) GetContact() []*CreateSiteMonitorResponseBodyDataAttachAlertResultContact {
	return s.Contact
}

func (s *CreateSiteMonitorResponseBodyDataAttachAlertResult) SetContact(v []*CreateSiteMonitorResponseBodyDataAttachAlertResultContact) *CreateSiteMonitorResponseBodyDataAttachAlertResult {
	s.Contact = v
	return s
}

func (s *CreateSiteMonitorResponseBodyDataAttachAlertResult) Validate() error {
	return dara.Validate(s)
}

type CreateSiteMonitorResponseBodyDataAttachAlertResultContact struct {
	// The status code that is returned after you associate the existing alert rule with the site monitoring task.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned after you associate the existing alert rule with the site monitoring task.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request that was sent to associate the existing alert rule with the site monitoring task.
	//
	// example:
	//
	// 5dd33455-4f65-4b0c-9200-33d66f3f340b
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the alert rule.
	//
	// example:
	//
	// SystemDefault_acs_ecs_dashboard_InternetOutRate_Percent
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// Indicates whether the existing alert rule was associated with the site monitoring task. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSiteMonitorResponseBodyDataAttachAlertResultContact) String() string {
	return dara.Prettify(s)
}

func (s CreateSiteMonitorResponseBodyDataAttachAlertResultContact) GoString() string {
	return s.String()
}

func (s *CreateSiteMonitorResponseBodyDataAttachAlertResultContact) GetCode() *string {
	return s.Code
}

func (s *CreateSiteMonitorResponseBodyDataAttachAlertResultContact) GetMessage() *string {
	return s.Message
}

func (s *CreateSiteMonitorResponseBodyDataAttachAlertResultContact) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSiteMonitorResponseBodyDataAttachAlertResultContact) GetRuleId() *string {
	return s.RuleId
}

func (s *CreateSiteMonitorResponseBodyDataAttachAlertResultContact) GetSuccess() *string {
	return s.Success
}

func (s *CreateSiteMonitorResponseBodyDataAttachAlertResultContact) SetCode(v string) *CreateSiteMonitorResponseBodyDataAttachAlertResultContact {
	s.Code = &v
	return s
}

func (s *CreateSiteMonitorResponseBodyDataAttachAlertResultContact) SetMessage(v string) *CreateSiteMonitorResponseBodyDataAttachAlertResultContact {
	s.Message = &v
	return s
}

func (s *CreateSiteMonitorResponseBodyDataAttachAlertResultContact) SetRequestId(v string) *CreateSiteMonitorResponseBodyDataAttachAlertResultContact {
	s.RequestId = &v
	return s
}

func (s *CreateSiteMonitorResponseBodyDataAttachAlertResultContact) SetRuleId(v string) *CreateSiteMonitorResponseBodyDataAttachAlertResultContact {
	s.RuleId = &v
	return s
}

func (s *CreateSiteMonitorResponseBodyDataAttachAlertResultContact) SetSuccess(v string) *CreateSiteMonitorResponseBodyDataAttachAlertResultContact {
	s.Success = &v
	return s
}

func (s *CreateSiteMonitorResponseBodyDataAttachAlertResultContact) Validate() error {
	return dara.Validate(s)
}
