// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVocabularyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteVocabularyResponseBody
	GetCode() *string
	SetData(v string) *DeleteVocabularyResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DeleteVocabularyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteVocabularyResponseBody
	GetMessage() *string
	SetParams(v []*string) *DeleteVocabularyResponseBody
	GetParams() []*string
	SetRequestId(v string) *DeleteVocabularyResponseBody
	GetRequestId() *string
}

type DeleteVocabularyResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// a395011f-a247-400f-bc69-28796749fd52
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Instance llm-baployoyopf22m2r does not exist.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// CF6D3484-19A1-5C77-863B-AC8B5754D37C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVocabularyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVocabularyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVocabularyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteVocabularyResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteVocabularyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteVocabularyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteVocabularyResponseBody) GetParams() []*string {
	return s.Params
}

func (s *DeleteVocabularyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVocabularyResponseBody) SetCode(v string) *DeleteVocabularyResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteVocabularyResponseBody) SetData(v string) *DeleteVocabularyResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteVocabularyResponseBody) SetHttpStatusCode(v int32) *DeleteVocabularyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteVocabularyResponseBody) SetMessage(v string) *DeleteVocabularyResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVocabularyResponseBody) SetParams(v []*string) *DeleteVocabularyResponseBody {
	s.Params = v
	return s
}

func (s *DeleteVocabularyResponseBody) SetRequestId(v string) *DeleteVocabularyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVocabularyResponseBody) Validate() error {
	return dara.Validate(s)
}
