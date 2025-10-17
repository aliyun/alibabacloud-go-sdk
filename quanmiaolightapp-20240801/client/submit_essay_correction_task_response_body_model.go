// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitEssayCorrectionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitEssayCorrectionTaskResponseBody
	GetCode() *string
	SetData(v *SubmitEssayCorrectionTaskResponseBodyData) *SubmitEssayCorrectionTaskResponseBody
	GetData() *SubmitEssayCorrectionTaskResponseBodyData
	SetHttpStatusCode(v int32) *SubmitEssayCorrectionTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitEssayCorrectionTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitEssayCorrectionTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitEssayCorrectionTaskResponseBody
	GetSuccess() *bool
}

type SubmitEssayCorrectionTaskResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                    `json:"code,omitempty" xml:"code,omitempty"`
	Data *SubmitEssayCorrectionTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SubmitEssayCorrectionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitEssayCorrectionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitEssayCorrectionTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitEssayCorrectionTaskResponseBody) GetData() *SubmitEssayCorrectionTaskResponseBodyData {
	return s.Data
}

func (s *SubmitEssayCorrectionTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitEssayCorrectionTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitEssayCorrectionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitEssayCorrectionTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitEssayCorrectionTaskResponseBody) SetCode(v string) *SubmitEssayCorrectionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitEssayCorrectionTaskResponseBody) SetData(v *SubmitEssayCorrectionTaskResponseBodyData) *SubmitEssayCorrectionTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitEssayCorrectionTaskResponseBody) SetHttpStatusCode(v int32) *SubmitEssayCorrectionTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitEssayCorrectionTaskResponseBody) SetMessage(v string) *SubmitEssayCorrectionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitEssayCorrectionTaskResponseBody) SetRequestId(v string) *SubmitEssayCorrectionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitEssayCorrectionTaskResponseBody) SetSuccess(v bool) *SubmitEssayCorrectionTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitEssayCorrectionTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitEssayCorrectionTaskResponseBodyData struct {
	// example:
	//
	// 3feb69ed02d9b1a17d0f1a942675d300
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SubmitEssayCorrectionTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitEssayCorrectionTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitEssayCorrectionTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitEssayCorrectionTaskResponseBodyData) SetTaskId(v string) *SubmitEssayCorrectionTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitEssayCorrectionTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
