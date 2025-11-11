// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitImportTermsTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitImportTermsTaskResponseBody
	GetCode() *string
	SetData(v *SubmitImportTermsTaskResponseBodyData) *SubmitImportTermsTaskResponseBody
	GetData() *SubmitImportTermsTaskResponseBodyData
	SetHttpStatusCode(v int32) *SubmitImportTermsTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitImportTermsTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitImportTermsTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitImportTermsTaskResponseBody
	GetSuccess() *bool
}

type SubmitImportTermsTaskResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitImportTermsTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s SubmitImportTermsTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitImportTermsTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitImportTermsTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitImportTermsTaskResponseBody) GetData() *SubmitImportTermsTaskResponseBodyData {
	return s.Data
}

func (s *SubmitImportTermsTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitImportTermsTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitImportTermsTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitImportTermsTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitImportTermsTaskResponseBody) SetCode(v string) *SubmitImportTermsTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitImportTermsTaskResponseBody) SetData(v *SubmitImportTermsTaskResponseBodyData) *SubmitImportTermsTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitImportTermsTaskResponseBody) SetHttpStatusCode(v int32) *SubmitImportTermsTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitImportTermsTaskResponseBody) SetMessage(v string) *SubmitImportTermsTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitImportTermsTaskResponseBody) SetRequestId(v string) *SubmitImportTermsTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitImportTermsTaskResponseBody) SetSuccess(v bool) *SubmitImportTermsTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitImportTermsTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitImportTermsTaskResponseBodyData struct {
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitImportTermsTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitImportTermsTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitImportTermsTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitImportTermsTaskResponseBodyData) SetTaskId(v string) *SubmitImportTermsTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitImportTermsTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
