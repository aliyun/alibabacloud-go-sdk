// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportCorpNumbersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImportCorpNumbersResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ImportCorpNumbersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ImportCorpNumbersResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImportCorpNumbersResponseBody
	GetRequestId() *string
}

type ImportCorpNumbersResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// C42981C7-93D9-55CD-B078-784F8522E0E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportCorpNumbersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportCorpNumbersResponseBody) GoString() string {
	return s.String()
}

func (s *ImportCorpNumbersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImportCorpNumbersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ImportCorpNumbersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportCorpNumbersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportCorpNumbersResponseBody) SetCode(v string) *ImportCorpNumbersResponseBody {
	s.Code = &v
	return s
}

func (s *ImportCorpNumbersResponseBody) SetHttpStatusCode(v int32) *ImportCorpNumbersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ImportCorpNumbersResponseBody) SetMessage(v string) *ImportCorpNumbersResponseBody {
	s.Message = &v
	return s
}

func (s *ImportCorpNumbersResponseBody) SetRequestId(v string) *ImportCorpNumbersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportCorpNumbersResponseBody) Validate() error {
	return dara.Validate(s)
}
