// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendDocumentAskQuestionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SendDocumentAskQuestionResponseBodyData) *SendDocumentAskQuestionResponseBody
	GetData() *SendDocumentAskQuestionResponseBodyData
	SetRequestId(v string) *SendDocumentAskQuestionResponseBody
	GetRequestId() *string
}

type SendDocumentAskQuestionResponseBody struct {
	// Returned data
	Data *SendDocumentAskQuestionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Request ID
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SendDocumentAskQuestionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendDocumentAskQuestionResponseBody) GoString() string {
	return s.String()
}

func (s *SendDocumentAskQuestionResponseBody) GetData() *SendDocumentAskQuestionResponseBodyData {
	return s.Data
}

func (s *SendDocumentAskQuestionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendDocumentAskQuestionResponseBody) SetData(v *SendDocumentAskQuestionResponseBodyData) *SendDocumentAskQuestionResponseBody {
	s.Data = v
	return s
}

func (s *SendDocumentAskQuestionResponseBody) SetRequestId(v string) *SendDocumentAskQuestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendDocumentAskQuestionResponseBody) Validate() error {
	return dara.Validate(s)
}

type SendDocumentAskQuestionResponseBodyData struct {
	// Q&A result
	//
	// example:
	//
	// Carbon emissions in 2023 totaled 4.681 million tons
	Answer *string `json:"answer,omitempty" xml:"answer,omitempty"`
	// Documents associated with the answer returned by the query
	Document []*string `json:"document,omitempty" xml:"document,omitempty" type:"Repeated"`
}

func (s SendDocumentAskQuestionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SendDocumentAskQuestionResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendDocumentAskQuestionResponseBodyData) GetAnswer() *string {
	return s.Answer
}

func (s *SendDocumentAskQuestionResponseBodyData) GetDocument() []*string {
	return s.Document
}

func (s *SendDocumentAskQuestionResponseBodyData) SetAnswer(v string) *SendDocumentAskQuestionResponseBodyData {
	s.Answer = &v
	return s
}

func (s *SendDocumentAskQuestionResponseBodyData) SetDocument(v []*string) *SendDocumentAskQuestionResponseBodyData {
	s.Document = v
	return s
}

func (s *SendDocumentAskQuestionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
