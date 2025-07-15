// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDocClusterTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitDocClusterTaskResponseBody
	GetCode() *string
	SetData(v *SubmitDocClusterTaskResponseBodyData) *SubmitDocClusterTaskResponseBody
	GetData() *SubmitDocClusterTaskResponseBodyData
	SetHttpStatusCode(v int32) *SubmitDocClusterTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitDocClusterTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitDocClusterTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitDocClusterTaskResponseBody
	GetSuccess() *bool
}

type SubmitDocClusterTaskResponseBody struct {
	// example:
	//
	// NoData
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitDocClusterTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s SubmitDocClusterTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocClusterTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDocClusterTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitDocClusterTaskResponseBody) GetData() *SubmitDocClusterTaskResponseBodyData {
	return s.Data
}

func (s *SubmitDocClusterTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitDocClusterTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitDocClusterTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitDocClusterTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitDocClusterTaskResponseBody) SetCode(v string) *SubmitDocClusterTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitDocClusterTaskResponseBody) SetData(v *SubmitDocClusterTaskResponseBodyData) *SubmitDocClusterTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitDocClusterTaskResponseBody) SetHttpStatusCode(v int32) *SubmitDocClusterTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitDocClusterTaskResponseBody) SetMessage(v string) *SubmitDocClusterTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitDocClusterTaskResponseBody) SetRequestId(v string) *SubmitDocClusterTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitDocClusterTaskResponseBody) SetSuccess(v bool) *SubmitDocClusterTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitDocClusterTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitDocClusterTaskResponseBodyData struct {
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitDocClusterTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocClusterTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitDocClusterTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitDocClusterTaskResponseBodyData) SetTaskId(v string) *SubmitDocClusterTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitDocClusterTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
