// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTaskGroupResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *CreateTaskGroupResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *CreateTaskGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTaskGroupResponseBody
	GetRequestId() *string
	SetResultObject(v *CreateTaskGroupResponseBodyResultObject) *CreateTaskGroupResponseBody
	GetResultObject() *CreateTaskGroupResponseBodyResultObject
}

type CreateTaskGroupResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *CreateTaskGroupResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s CreateTaskGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTaskGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTaskGroupResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *CreateTaskGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTaskGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTaskGroupResponseBody) GetResultObject() *CreateTaskGroupResponseBodyResultObject {
	return s.ResultObject
}

func (s *CreateTaskGroupResponseBody) SetCode(v string) *CreateTaskGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTaskGroupResponseBody) SetHttpStatusCode(v string) *CreateTaskGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateTaskGroupResponseBody) SetMessage(v string) *CreateTaskGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTaskGroupResponseBody) SetRequestId(v string) *CreateTaskGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTaskGroupResponseBody) SetResultObject(v *CreateTaskGroupResponseBodyResultObject) *CreateTaskGroupResponseBody {
	s.ResultObject = v
	return s
}

func (s *CreateTaskGroupResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTaskGroupResponseBodyResultObject struct {
	// example:
	//
	// 1750645267000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 345298
	CreatorUserId *int32 `json:"CreatorUserId,omitempty" xml:"CreatorUserId,omitempty"`
	// example:
	//
	// RUNNING
	GroupStatus *string   `json:"GroupStatus,omitempty" xml:"GroupStatus,omitempty"`
	SampleNames []*string `json:"SampleNames,omitempty" xml:"SampleNames,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	SubTaskCount *int32 `json:"SubTaskCount,omitempty" xml:"SubTaskCount,omitempty"`
	// example:
	//
	// FINANCE
	Tab *string `json:"Tab,omitempty" xml:"Tab,omitempty"`
	// example:
	//
	// g-0jlcreertd0p471l6f72
	TaskGroupId *int32 `json:"TaskGroupId,omitempty" xml:"TaskGroupId,omitempty"`
	// example:
	//
	// GroupTest
	TaskGroupName *string `json:"TaskGroupName,omitempty" xml:"TaskGroupName,omitempty"`
}

func (s CreateTaskGroupResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskGroupResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *CreateTaskGroupResponseBodyResultObject) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CreateTaskGroupResponseBodyResultObject) GetCreatorUserId() *int32 {
	return s.CreatorUserId
}

func (s *CreateTaskGroupResponseBodyResultObject) GetGroupStatus() *string {
	return s.GroupStatus
}

func (s *CreateTaskGroupResponseBodyResultObject) GetSampleNames() []*string {
	return s.SampleNames
}

func (s *CreateTaskGroupResponseBodyResultObject) GetSubTaskCount() *int32 {
	return s.SubTaskCount
}

func (s *CreateTaskGroupResponseBodyResultObject) GetTab() *string {
	return s.Tab
}

func (s *CreateTaskGroupResponseBodyResultObject) GetTaskGroupId() *int32 {
	return s.TaskGroupId
}

func (s *CreateTaskGroupResponseBodyResultObject) GetTaskGroupName() *string {
	return s.TaskGroupName
}

func (s *CreateTaskGroupResponseBodyResultObject) SetCreateTime(v int64) *CreateTaskGroupResponseBodyResultObject {
	s.CreateTime = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObject) SetCreatorUserId(v int32) *CreateTaskGroupResponseBodyResultObject {
	s.CreatorUserId = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObject) SetGroupStatus(v string) *CreateTaskGroupResponseBodyResultObject {
	s.GroupStatus = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObject) SetSampleNames(v []*string) *CreateTaskGroupResponseBodyResultObject {
	s.SampleNames = v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObject) SetSubTaskCount(v int32) *CreateTaskGroupResponseBodyResultObject {
	s.SubTaskCount = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObject) SetTab(v string) *CreateTaskGroupResponseBodyResultObject {
	s.Tab = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObject) SetTaskGroupId(v int32) *CreateTaskGroupResponseBodyResultObject {
	s.TaskGroupId = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObject) SetTaskGroupName(v string) *CreateTaskGroupResponseBodyResultObject {
	s.TaskGroupName = &v
	return s
}

func (s *CreateTaskGroupResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
