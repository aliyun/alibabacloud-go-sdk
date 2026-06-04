// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitParseDocumentLayoutTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitParseDocumentLayoutTaskResponseBody
	GetCode() *string
	SetData(v *SubmitParseDocumentLayoutTaskResponseBodyData) *SubmitParseDocumentLayoutTaskResponseBody
	GetData() *SubmitParseDocumentLayoutTaskResponseBodyData
	SetHttpStatusCode(v int32) *SubmitParseDocumentLayoutTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitParseDocumentLayoutTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitParseDocumentLayoutTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitParseDocumentLayoutTaskResponseBody
	GetSuccess() *bool
}

type SubmitParseDocumentLayoutTaskResponseBody struct {
	// example:
	//
	// successful
	Code *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitParseDocumentLayoutTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitParseDocumentLayoutTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitParseDocumentLayoutTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitParseDocumentLayoutTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitParseDocumentLayoutTaskResponseBody) GetData() *SubmitParseDocumentLayoutTaskResponseBodyData {
	return s.Data
}

func (s *SubmitParseDocumentLayoutTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitParseDocumentLayoutTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitParseDocumentLayoutTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitParseDocumentLayoutTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitParseDocumentLayoutTaskResponseBody) SetCode(v string) *SubmitParseDocumentLayoutTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitParseDocumentLayoutTaskResponseBody) SetData(v *SubmitParseDocumentLayoutTaskResponseBodyData) *SubmitParseDocumentLayoutTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitParseDocumentLayoutTaskResponseBody) SetHttpStatusCode(v int32) *SubmitParseDocumentLayoutTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitParseDocumentLayoutTaskResponseBody) SetMessage(v string) *SubmitParseDocumentLayoutTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitParseDocumentLayoutTaskResponseBody) SetRequestId(v string) *SubmitParseDocumentLayoutTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitParseDocumentLayoutTaskResponseBody) SetSuccess(v bool) *SubmitParseDocumentLayoutTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitParseDocumentLayoutTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitParseDocumentLayoutTaskResponseBodyData struct {
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitParseDocumentLayoutTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitParseDocumentLayoutTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitParseDocumentLayoutTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitParseDocumentLayoutTaskResponseBodyData) SetTaskId(v string) *SubmitParseDocumentLayoutTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitParseDocumentLayoutTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
