// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSmartClipTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitSmartClipTaskResponseBody
	GetCode() *string
	SetData(v *SubmitSmartClipTaskResponseBodyData) *SubmitSmartClipTaskResponseBody
	GetData() *SubmitSmartClipTaskResponseBodyData
	SetHttpStatusCode(v int32) *SubmitSmartClipTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitSmartClipTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitSmartClipTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitSmartClipTaskResponseBody
	GetSuccess() *bool
}

type SubmitSmartClipTaskResponseBody struct {
	// example:
	//
	// NoData
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitSmartClipTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s SubmitSmartClipTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmartClipTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSmartClipTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitSmartClipTaskResponseBody) GetData() *SubmitSmartClipTaskResponseBodyData {
	return s.Data
}

func (s *SubmitSmartClipTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitSmartClipTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitSmartClipTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitSmartClipTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitSmartClipTaskResponseBody) SetCode(v string) *SubmitSmartClipTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitSmartClipTaskResponseBody) SetData(v *SubmitSmartClipTaskResponseBodyData) *SubmitSmartClipTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitSmartClipTaskResponseBody) SetHttpStatusCode(v int32) *SubmitSmartClipTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitSmartClipTaskResponseBody) SetMessage(v string) *SubmitSmartClipTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitSmartClipTaskResponseBody) SetRequestId(v string) *SubmitSmartClipTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSmartClipTaskResponseBody) SetSuccess(v bool) *SubmitSmartClipTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitSmartClipTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitSmartClipTaskResponseBodyData struct {
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitSmartClipTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmartClipTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitSmartClipTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitSmartClipTaskResponseBodyData) SetTaskId(v string) *SubmitSmartClipTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitSmartClipTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
