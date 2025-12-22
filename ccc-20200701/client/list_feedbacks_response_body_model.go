// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeedbacksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListFeedbacksResponseBody
	GetCode() *string
	SetData(v []*ListFeedbacksResponseBodyData) *ListFeedbacksResponseBody
	GetData() []*ListFeedbacksResponseBodyData
	SetHttpStatusCode(v int32) *ListFeedbacksResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListFeedbacksResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListFeedbacksResponseBody
	GetRequestId() *string
}

type ListFeedbacksResponseBody struct {
	// example:
	//
	// OK
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListFeedbacksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 26E54929-CA86-1035-9B42-0C8F291BB027
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFeedbacksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFeedbacksResponseBody) GoString() string {
	return s.String()
}

func (s *ListFeedbacksResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListFeedbacksResponseBody) GetData() []*ListFeedbacksResponseBodyData {
	return s.Data
}

func (s *ListFeedbacksResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListFeedbacksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListFeedbacksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFeedbacksResponseBody) SetCode(v string) *ListFeedbacksResponseBody {
	s.Code = &v
	return s
}

func (s *ListFeedbacksResponseBody) SetData(v []*ListFeedbacksResponseBodyData) *ListFeedbacksResponseBody {
	s.Data = v
	return s
}

func (s *ListFeedbacksResponseBody) SetHttpStatusCode(v int32) *ListFeedbacksResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListFeedbacksResponseBody) SetMessage(v string) *ListFeedbacksResponseBody {
	s.Message = &v
	return s
}

func (s *ListFeedbacksResponseBody) SetRequestId(v string) *ListFeedbacksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFeedbacksResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFeedbacksResponseBodyData struct {
	// example:
	//
	// job-25920271311543****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 20251216-8B9B7B02-16FE-54BE-942A-F59DE0656032
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// PostCallAnalyzer:solution
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// -1
	UserRating *int32 `json:"UserRating,omitempty" xml:"UserRating,omitempty"`
	// example:
	//
	// xxxxxxx
	UserResponse *string `json:"UserResponse,omitempty" xml:"UserResponse,omitempty"`
}

func (s ListFeedbacksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListFeedbacksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFeedbacksResponseBodyData) GetContactId() *string {
	return s.ContactId
}

func (s *ListFeedbacksResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListFeedbacksResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ListFeedbacksResponseBodyData) GetTaskName() *string {
	return s.TaskName
}

func (s *ListFeedbacksResponseBodyData) GetUserRating() *int32 {
	return s.UserRating
}

func (s *ListFeedbacksResponseBodyData) GetUserResponse() *string {
	return s.UserResponse
}

func (s *ListFeedbacksResponseBodyData) SetContactId(v string) *ListFeedbacksResponseBodyData {
	s.ContactId = &v
	return s
}

func (s *ListFeedbacksResponseBodyData) SetInstanceId(v string) *ListFeedbacksResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListFeedbacksResponseBodyData) SetTaskId(v string) *ListFeedbacksResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ListFeedbacksResponseBodyData) SetTaskName(v string) *ListFeedbacksResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *ListFeedbacksResponseBodyData) SetUserRating(v int32) *ListFeedbacksResponseBodyData {
	s.UserRating = &v
	return s
}

func (s *ListFeedbacksResponseBodyData) SetUserResponse(v string) *ListFeedbacksResponseBodyData {
	s.UserResponse = &v
	return s
}

func (s *ListFeedbacksResponseBodyData) Validate() error {
	return dara.Validate(s)
}
