// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportVocabularyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImportVocabularyResponseBody
	GetCode() *string
	SetData(v []*string) *ImportVocabularyResponseBody
	GetData() []*string
	SetHttpStatusCode(v int32) *ImportVocabularyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ImportVocabularyResponseBody
	GetMessage() *string
	SetParams(v []*string) *ImportVocabularyResponseBody
	GetParams() []*string
	SetRequestId(v string) *ImportVocabularyResponseBody
	GetRequestId() *string
}

type ImportVocabularyResponseBody struct {
	// example:
	//
	// OK
	Code *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Instance llm-rj6aqmctjcit4acy does not exist.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// D771A1B6-3D5F-174A-BEE1-98CE1000D337
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportVocabularyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportVocabularyResponseBody) GoString() string {
	return s.String()
}

func (s *ImportVocabularyResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImportVocabularyResponseBody) GetData() []*string {
	return s.Data
}

func (s *ImportVocabularyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ImportVocabularyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportVocabularyResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ImportVocabularyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportVocabularyResponseBody) SetCode(v string) *ImportVocabularyResponseBody {
	s.Code = &v
	return s
}

func (s *ImportVocabularyResponseBody) SetData(v []*string) *ImportVocabularyResponseBody {
	s.Data = v
	return s
}

func (s *ImportVocabularyResponseBody) SetHttpStatusCode(v int32) *ImportVocabularyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ImportVocabularyResponseBody) SetMessage(v string) *ImportVocabularyResponseBody {
	s.Message = &v
	return s
}

func (s *ImportVocabularyResponseBody) SetParams(v []*string) *ImportVocabularyResponseBody {
	s.Params = v
	return s
}

func (s *ImportVocabularyResponseBody) SetRequestId(v string) *ImportVocabularyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportVocabularyResponseBody) Validate() error {
	return dara.Validate(s)
}
