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
	// 29D8CB48-2384-59C4-9618-3824009635C2
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
