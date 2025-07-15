// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitExportTermsTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitExportTermsTaskResponseBody
	GetCode() *string
	SetData(v *SubmitExportTermsTaskResponseBodyData) *SubmitExportTermsTaskResponseBody
	GetData() *SubmitExportTermsTaskResponseBodyData
	SetHttpStatusCode(v int32) *SubmitExportTermsTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitExportTermsTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitExportTermsTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitExportTermsTaskResponseBody
	GetSuccess() *bool
}

type SubmitExportTermsTaskResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitExportTermsTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F2F366D6-E9FE-1006-BB70-2C650896AAB5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitExportTermsTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitExportTermsTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitExportTermsTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitExportTermsTaskResponseBody) GetData() *SubmitExportTermsTaskResponseBodyData {
	return s.Data
}

func (s *SubmitExportTermsTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitExportTermsTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitExportTermsTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitExportTermsTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitExportTermsTaskResponseBody) SetCode(v string) *SubmitExportTermsTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitExportTermsTaskResponseBody) SetData(v *SubmitExportTermsTaskResponseBodyData) *SubmitExportTermsTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitExportTermsTaskResponseBody) SetHttpStatusCode(v int32) *SubmitExportTermsTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitExportTermsTaskResponseBody) SetMessage(v string) *SubmitExportTermsTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitExportTermsTaskResponseBody) SetRequestId(v string) *SubmitExportTermsTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitExportTermsTaskResponseBody) SetSuccess(v bool) *SubmitExportTermsTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitExportTermsTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitExportTermsTaskResponseBodyData struct {
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitExportTermsTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitExportTermsTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitExportTermsTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitExportTermsTaskResponseBodyData) SetTaskId(v string) *SubmitExportTermsTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitExportTermsTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
