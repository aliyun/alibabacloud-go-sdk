// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVocabularyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateVocabularyResponseBody
	GetCode() *string
	SetData(v string) *UpdateVocabularyResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateVocabularyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateVocabularyResponseBody
	GetMessage() *string
	SetParams(v []*string) *UpdateVocabularyResponseBody
	GetParams() []*string
	SetRequestId(v string) *UpdateVocabularyResponseBody
	GetRequestId() *string
}

type UpdateVocabularyResponseBody struct {
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
	// 7401D698-0AAC-5909-B68E-88C68805FFB8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVocabularyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVocabularyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVocabularyResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateVocabularyResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateVocabularyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateVocabularyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateVocabularyResponseBody) GetParams() []*string {
	return s.Params
}

func (s *UpdateVocabularyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVocabularyResponseBody) SetCode(v string) *UpdateVocabularyResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateVocabularyResponseBody) SetData(v string) *UpdateVocabularyResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateVocabularyResponseBody) SetHttpStatusCode(v int32) *UpdateVocabularyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateVocabularyResponseBody) SetMessage(v string) *UpdateVocabularyResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateVocabularyResponseBody) SetParams(v []*string) *UpdateVocabularyResponseBody {
	s.Params = v
	return s
}

func (s *UpdateVocabularyResponseBody) SetRequestId(v string) *UpdateVocabularyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVocabularyResponseBody) Validate() error {
	return dara.Validate(s)
}
