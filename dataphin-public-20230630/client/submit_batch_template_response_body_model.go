// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitBatchTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitBatchTemplateResponseBody
	GetCode() *string
	SetData(v *SubmitBatchTemplateResponseBodyData) *SubmitBatchTemplateResponseBody
	GetData() *SubmitBatchTemplateResponseBodyData
	SetHttpStatusCode(v int32) *SubmitBatchTemplateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitBatchTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitBatchTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitBatchTemplateResponseBody
	GetSuccess() *bool
}

type SubmitBatchTemplateResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The submission result.
	Data *SubmitBatchTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend exception details.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitBatchTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitBatchTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitBatchTemplateResponseBody) GetData() *SubmitBatchTemplateResponseBodyData {
	return s.Data
}

func (s *SubmitBatchTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitBatchTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitBatchTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitBatchTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitBatchTemplateResponseBody) SetCode(v string) *SubmitBatchTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitBatchTemplateResponseBody) SetData(v *SubmitBatchTemplateResponseBodyData) *SubmitBatchTemplateResponseBody {
	s.Data = v
	return s
}

func (s *SubmitBatchTemplateResponseBody) SetHttpStatusCode(v int32) *SubmitBatchTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitBatchTemplateResponseBody) SetMessage(v string) *SubmitBatchTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitBatchTemplateResponseBody) SetRequestId(v string) *SubmitBatchTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitBatchTemplateResponseBody) SetSuccess(v bool) *SubmitBatchTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitBatchTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitBatchTemplateResponseBodyData struct {
	// The submission record ID.
	//
	// example:
	//
	// 1001
	SubmitId *int64 `json:"SubmitId,omitempty" xml:"SubmitId,omitempty"`
}

func (s SubmitBatchTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitBatchTemplateResponseBodyData) GetSubmitId() *int64 {
	return s.SubmitId
}

func (s *SubmitBatchTemplateResponseBodyData) SetSubmitId(v int64) *SubmitBatchTemplateResponseBodyData {
	s.SubmitId = &v
	return s
}

func (s *SubmitBatchTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
