// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDocumentSummaryResponseBody
	GetCode() *string
	SetData(v *GetDocumentSummaryResponseBodyData) *GetDocumentSummaryResponseBody
	GetData() *GetDocumentSummaryResponseBodyData
	SetHttpStatusCode(v int32) *GetDocumentSummaryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDocumentSummaryResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDocumentSummaryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDocumentSummaryResponseBody
	GetSuccess() *bool
}

type GetDocumentSummaryResponseBody struct {
	// Interface return code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data query result.
	Data *GetDocumentSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message for the returned result.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7903F2DE-D9EE-5D16-8A08-E9223E54B281
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful. Values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDocumentSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentSummaryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDocumentSummaryResponseBody) GetData() *GetDocumentSummaryResponseBodyData {
	return s.Data
}

func (s *GetDocumentSummaryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDocumentSummaryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDocumentSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDocumentSummaryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDocumentSummaryResponseBody) SetCode(v string) *GetDocumentSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetDocumentSummaryResponseBody) SetData(v *GetDocumentSummaryResponseBodyData) *GetDocumentSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetDocumentSummaryResponseBody) SetHttpStatusCode(v int32) *GetDocumentSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDocumentSummaryResponseBody) SetMessage(v string) *GetDocumentSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocumentSummaryResponseBody) SetRequestId(v string) *GetDocumentSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentSummaryResponseBody) SetSuccess(v bool) *GetDocumentSummaryResponseBody {
	s.Success = &v
	return s
}

func (s *GetDocumentSummaryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDocumentSummaryResponseBodyData struct {
	// Number of documents.
	//
	// example:
	//
	// 10
	DocumentCount *int64 `json:"DocumentCount,omitempty" xml:"DocumentCount,omitempty"`
	// Number of services or days.
	//
	// example:
	//
	// 10
	Frequency *int64 `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
}

func (s GetDocumentSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDocumentSummaryResponseBodyData) GetDocumentCount() *int64 {
	return s.DocumentCount
}

func (s *GetDocumentSummaryResponseBodyData) GetFrequency() *int64 {
	return s.Frequency
}

func (s *GetDocumentSummaryResponseBodyData) SetDocumentCount(v int64) *GetDocumentSummaryResponseBodyData {
	s.DocumentCount = &v
	return s
}

func (s *GetDocumentSummaryResponseBodyData) SetFrequency(v int64) *GetDocumentSummaryResponseBodyData {
	s.Frequency = &v
	return s
}

func (s *GetDocumentSummaryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
