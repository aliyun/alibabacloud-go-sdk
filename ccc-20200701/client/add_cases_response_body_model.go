// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddCasesResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *AddCasesResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *AddCasesResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddCasesResponseBody
	GetRequestId() *string
}

type AddCasesResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 41298961-CAD7-5270-9378-FFD69F153144
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddCasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCasesResponseBody) GoString() string {
	return s.String()
}

func (s *AddCasesResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddCasesResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *AddCasesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddCasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCasesResponseBody) SetCode(v string) *AddCasesResponseBody {
	s.Code = &v
	return s
}

func (s *AddCasesResponseBody) SetHttpStatusCode(v string) *AddCasesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddCasesResponseBody) SetMessage(v string) *AddCasesResponseBody {
	s.Message = &v
	return s
}

func (s *AddCasesResponseBody) SetRequestId(v string) *AddCasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCasesResponseBody) Validate() error {
	return dara.Validate(s)
}
