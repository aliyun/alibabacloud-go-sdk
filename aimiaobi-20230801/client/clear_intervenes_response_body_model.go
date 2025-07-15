// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearIntervenesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ClearIntervenesResponseBody
	GetCode() *string
	SetData(v *ClearIntervenesResponseBodyData) *ClearIntervenesResponseBody
	GetData() *ClearIntervenesResponseBodyData
	SetHttpStatusCode(v int32) *ClearIntervenesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ClearIntervenesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClearIntervenesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ClearIntervenesResponseBody
	GetSuccess() *bool
}

type ClearIntervenesResponseBody struct {
	// example:
	//
	// 0
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClearIntervenesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ClearIntervenesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClearIntervenesResponseBody) GoString() string {
	return s.String()
}

func (s *ClearIntervenesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClearIntervenesResponseBody) GetData() *ClearIntervenesResponseBodyData {
	return s.Data
}

func (s *ClearIntervenesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ClearIntervenesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClearIntervenesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClearIntervenesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ClearIntervenesResponseBody) SetCode(v string) *ClearIntervenesResponseBody {
	s.Code = &v
	return s
}

func (s *ClearIntervenesResponseBody) SetData(v *ClearIntervenesResponseBodyData) *ClearIntervenesResponseBody {
	s.Data = v
	return s
}

func (s *ClearIntervenesResponseBody) SetHttpStatusCode(v int32) *ClearIntervenesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ClearIntervenesResponseBody) SetMessage(v string) *ClearIntervenesResponseBody {
	s.Message = &v
	return s
}

func (s *ClearIntervenesResponseBody) SetRequestId(v string) *ClearIntervenesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClearIntervenesResponseBody) SetSuccess(v bool) *ClearIntervenesResponseBody {
	s.Success = &v
	return s
}

func (s *ClearIntervenesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ClearIntervenesResponseBodyData struct {
	Code       *int32    `json:"Code,omitempty" xml:"Code,omitempty"`
	FailIdList []*string `json:"FailIdList,omitempty" xml:"FailIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ClearIntervenesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClearIntervenesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClearIntervenesResponseBodyData) GetCode() *int32 {
	return s.Code
}

func (s *ClearIntervenesResponseBodyData) GetFailIdList() []*string {
	return s.FailIdList
}

func (s *ClearIntervenesResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ClearIntervenesResponseBodyData) SetCode(v int32) *ClearIntervenesResponseBodyData {
	s.Code = &v
	return s
}

func (s *ClearIntervenesResponseBodyData) SetFailIdList(v []*string) *ClearIntervenesResponseBodyData {
	s.FailIdList = v
	return s
}

func (s *ClearIntervenesResponseBodyData) SetTaskId(v string) *ClearIntervenesResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ClearIntervenesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
