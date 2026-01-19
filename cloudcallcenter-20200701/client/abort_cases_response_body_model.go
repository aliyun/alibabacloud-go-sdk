// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbortCasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AbortCasesResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *AbortCasesResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *AbortCasesResponseBody
	GetMessage() *string
	SetRequestId(v string) *AbortCasesResponseBody
	GetRequestId() *string
}

type AbortCasesResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AbortCasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AbortCasesResponseBody) GoString() string {
	return s.String()
}

func (s *AbortCasesResponseBody) GetCode() *string {
	return s.Code
}

func (s *AbortCasesResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *AbortCasesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AbortCasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AbortCasesResponseBody) SetCode(v string) *AbortCasesResponseBody {
	s.Code = &v
	return s
}

func (s *AbortCasesResponseBody) SetHttpStatusCode(v string) *AbortCasesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AbortCasesResponseBody) SetMessage(v string) *AbortCasesResponseBody {
	s.Message = &v
	return s
}

func (s *AbortCasesResponseBody) SetRequestId(v string) *AbortCasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AbortCasesResponseBody) Validate() error {
	return dara.Validate(s)
}
