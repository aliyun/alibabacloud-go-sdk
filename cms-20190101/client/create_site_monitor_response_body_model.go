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
	Code             *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
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
	if s.CreateResultList != nil {
		if err := s.CreateResultList.Validate(); err != nil {
			return err
		}
	}
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.CreateResultList != nil {
		for _, item := range s.CreateResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateSiteMonitorResponseBodyCreateResultListCreateResultList struct {
	TaskId   *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	if s.AttachAlertResult != nil {
		if err := s.AttachAlertResult.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.Contact != nil {
		for _, item := range s.Contact {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateSiteMonitorResponseBodyDataAttachAlertResultContact struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleId    *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
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
