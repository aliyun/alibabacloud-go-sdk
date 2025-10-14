// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstantSiteMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateInstantSiteMonitorResponseBody
	GetCode() *string
	SetCreateResultList(v []*CreateInstantSiteMonitorResponseBodyCreateResultList) *CreateInstantSiteMonitorResponseBody
	GetCreateResultList() []*CreateInstantSiteMonitorResponseBodyCreateResultList
	SetMessage(v string) *CreateInstantSiteMonitorResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateInstantSiteMonitorResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateInstantSiteMonitorResponseBody
	GetSuccess() *string
}

type CreateInstantSiteMonitorResponseBody struct {
	// The error code.
	//
	// > The status code 200 indicates that the call was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The results for creating the instant test task.
	CreateResultList []*CreateInstantSiteMonitorResponseBodyCreateResultList `json:"CreateResultList,omitempty" xml:"CreateResultList,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 68192f5d-0d45-4b98-9724-892813f86c71
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateInstantSiteMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstantSiteMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstantSiteMonitorResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateInstantSiteMonitorResponseBody) GetCreateResultList() []*CreateInstantSiteMonitorResponseBodyCreateResultList {
	return s.CreateResultList
}

func (s *CreateInstantSiteMonitorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateInstantSiteMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstantSiteMonitorResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateInstantSiteMonitorResponseBody) SetCode(v string) *CreateInstantSiteMonitorResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstantSiteMonitorResponseBody) SetCreateResultList(v []*CreateInstantSiteMonitorResponseBodyCreateResultList) *CreateInstantSiteMonitorResponseBody {
	s.CreateResultList = v
	return s
}

func (s *CreateInstantSiteMonitorResponseBody) SetMessage(v string) *CreateInstantSiteMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *CreateInstantSiteMonitorResponseBody) SetRequestId(v string) *CreateInstantSiteMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstantSiteMonitorResponseBody) SetSuccess(v string) *CreateInstantSiteMonitorResponseBody {
	s.Success = &v
	return s
}

func (s *CreateInstantSiteMonitorResponseBody) Validate() error {
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

type CreateInstantSiteMonitorResponseBodyCreateResultList struct {
	// The ID of the instant test task.
	//
	// example:
	//
	// 2c8dbdf9-a3ab-46a1-85a4-f094965e****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the instant test task.
	//
	// example:
	//
	// task1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateInstantSiteMonitorResponseBodyCreateResultList) String() string {
	return dara.Prettify(s)
}

func (s CreateInstantSiteMonitorResponseBodyCreateResultList) GoString() string {
	return s.String()
}

func (s *CreateInstantSiteMonitorResponseBodyCreateResultList) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateInstantSiteMonitorResponseBodyCreateResultList) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateInstantSiteMonitorResponseBodyCreateResultList) SetTaskId(v string) *CreateInstantSiteMonitorResponseBodyCreateResultList {
	s.TaskId = &v
	return s
}

func (s *CreateInstantSiteMonitorResponseBodyCreateResultList) SetTaskName(v string) *CreateInstantSiteMonitorResponseBodyCreateResultList {
	s.TaskName = &v
	return s
}

func (s *CreateInstantSiteMonitorResponseBodyCreateResultList) Validate() error {
	return dara.Validate(s)
}
