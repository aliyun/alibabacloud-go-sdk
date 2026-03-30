// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVocabularyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateVocabularyResponseBody
	GetCode() *string
	SetData(v string) *CreateVocabularyResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateVocabularyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateVocabularyResponseBody
	GetMessage() *string
	SetParams(v []*string) *CreateVocabularyResponseBody
	GetParams() []*string
	SetRequestId(v string) *CreateVocabularyResponseBody
	GetRequestId() *string
}

type CreateVocabularyResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66b
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Instance af81a389-91f0-4157-8d82-720edd02b66a
	//
	//  does not exist.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// BA2B89BE-BF85-5333-8A52-6D9F18A100F1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVocabularyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVocabularyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVocabularyResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateVocabularyResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateVocabularyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateVocabularyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateVocabularyResponseBody) GetParams() []*string {
	return s.Params
}

func (s *CreateVocabularyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVocabularyResponseBody) SetCode(v string) *CreateVocabularyResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVocabularyResponseBody) SetData(v string) *CreateVocabularyResponseBody {
	s.Data = &v
	return s
}

func (s *CreateVocabularyResponseBody) SetHttpStatusCode(v int32) *CreateVocabularyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateVocabularyResponseBody) SetMessage(v string) *CreateVocabularyResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVocabularyResponseBody) SetParams(v []*string) *CreateVocabularyResponseBody {
	s.Params = v
	return s
}

func (s *CreateVocabularyResponseBody) SetRequestId(v string) *CreateVocabularyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVocabularyResponseBody) Validate() error {
	return dara.Validate(s)
}
