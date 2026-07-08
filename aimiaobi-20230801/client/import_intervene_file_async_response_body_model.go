// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportInterveneFileAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImportInterveneFileAsyncResponseBody
	GetCode() *string
	SetData(v *ImportInterveneFileAsyncResponseBodyData) *ImportInterveneFileAsyncResponseBody
	GetData() *ImportInterveneFileAsyncResponseBodyData
	SetHttpStatusCode(v int32) *ImportInterveneFileAsyncResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ImportInterveneFileAsyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImportInterveneFileAsyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImportInterveneFileAsyncResponseBody
	GetSuccess() *bool
}

type ImportInterveneFileAsyncResponseBody struct {
	// Status code
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Business data
	Data *ImportInterveneFileAsyncResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Error message
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Unique request identifier
	//
	// example:
	//
	// 94512A33-8EC1-5452-A793-5C91F18ED2F0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded. true means success. false means failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImportInterveneFileAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportInterveneFileAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *ImportInterveneFileAsyncResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImportInterveneFileAsyncResponseBody) GetData() *ImportInterveneFileAsyncResponseBodyData {
	return s.Data
}

func (s *ImportInterveneFileAsyncResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ImportInterveneFileAsyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportInterveneFileAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportInterveneFileAsyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImportInterveneFileAsyncResponseBody) SetCode(v string) *ImportInterveneFileAsyncResponseBody {
	s.Code = &v
	return s
}

func (s *ImportInterveneFileAsyncResponseBody) SetData(v *ImportInterveneFileAsyncResponseBodyData) *ImportInterveneFileAsyncResponseBody {
	s.Data = v
	return s
}

func (s *ImportInterveneFileAsyncResponseBody) SetHttpStatusCode(v int32) *ImportInterveneFileAsyncResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ImportInterveneFileAsyncResponseBody) SetMessage(v string) *ImportInterveneFileAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *ImportInterveneFileAsyncResponseBody) SetRequestId(v string) *ImportInterveneFileAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportInterveneFileAsyncResponseBody) SetSuccess(v bool) *ImportInterveneFileAsyncResponseBody {
	s.Success = &v
	return s
}

func (s *ImportInterveneFileAsyncResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImportInterveneFileAsyncResponseBodyData struct {
	// Status code returned by the intervention service
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// List of failed index IDs
	FailIdList []*string `json:"FailIdList,omitempty" xml:"FailIdList,omitempty" type:"Repeated"`
	// Task ID
	//
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ImportInterveneFileAsyncResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImportInterveneFileAsyncResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImportInterveneFileAsyncResponseBodyData) GetCode() *int32 {
	return s.Code
}

func (s *ImportInterveneFileAsyncResponseBodyData) GetFailIdList() []*string {
	return s.FailIdList
}

func (s *ImportInterveneFileAsyncResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ImportInterveneFileAsyncResponseBodyData) SetCode(v int32) *ImportInterveneFileAsyncResponseBodyData {
	s.Code = &v
	return s
}

func (s *ImportInterveneFileAsyncResponseBodyData) SetFailIdList(v []*string) *ImportInterveneFileAsyncResponseBodyData {
	s.FailIdList = v
	return s
}

func (s *ImportInterveneFileAsyncResponseBodyData) SetTaskId(v string) *ImportInterveneFileAsyncResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ImportInterveneFileAsyncResponseBodyData) Validate() error {
	return dara.Validate(s)
}
