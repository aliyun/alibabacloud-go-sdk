// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoDetectShotTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitVideoDetectShotTaskResponseBody
	GetCode() *string
	SetData(v *SubmitVideoDetectShotTaskResponseBodyData) *SubmitVideoDetectShotTaskResponseBody
	GetData() *SubmitVideoDetectShotTaskResponseBodyData
	SetHttpStatusCode(v int32) *SubmitVideoDetectShotTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitVideoDetectShotTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitVideoDetectShotTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitVideoDetectShotTaskResponseBody
	GetSuccess() *bool
}

type SubmitVideoDetectShotTaskResponseBody struct {
	// example:
	//
	// xx
	Code *string                                    `json:"code,omitempty" xml:"code,omitempty"`
	Data *SubmitVideoDetectShotTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 117F5ABE-CF02-5502-9A3F-E56BC9081A64
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SubmitVideoDetectShotTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoDetectShotTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitVideoDetectShotTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitVideoDetectShotTaskResponseBody) GetData() *SubmitVideoDetectShotTaskResponseBodyData {
	return s.Data
}

func (s *SubmitVideoDetectShotTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitVideoDetectShotTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitVideoDetectShotTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitVideoDetectShotTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitVideoDetectShotTaskResponseBody) SetCode(v string) *SubmitVideoDetectShotTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitVideoDetectShotTaskResponseBody) SetData(v *SubmitVideoDetectShotTaskResponseBodyData) *SubmitVideoDetectShotTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitVideoDetectShotTaskResponseBody) SetHttpStatusCode(v int32) *SubmitVideoDetectShotTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitVideoDetectShotTaskResponseBody) SetMessage(v string) *SubmitVideoDetectShotTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitVideoDetectShotTaskResponseBody) SetRequestId(v string) *SubmitVideoDetectShotTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitVideoDetectShotTaskResponseBody) SetSuccess(v bool) *SubmitVideoDetectShotTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitVideoDetectShotTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitVideoDetectShotTaskResponseBodyData struct {
	// example:
	//
	// 22ec888712de45b39b97983f8d166831
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SubmitVideoDetectShotTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoDetectShotTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitVideoDetectShotTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitVideoDetectShotTaskResponseBodyData) SetTaskId(v string) *SubmitVideoDetectShotTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitVideoDetectShotTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
