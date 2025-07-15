// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppendCasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AppendCasesResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *AppendCasesResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *AppendCasesResponseBody
	GetMessage() *string
	SetRequestId(v string) *AppendCasesResponseBody
	GetRequestId() *string
}

type AppendCasesResponseBody struct {
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
	// BC976D32-AC4C-4E0F-8AA9-F4BC6C4E2B3E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AppendCasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AppendCasesResponseBody) GoString() string {
	return s.String()
}

func (s *AppendCasesResponseBody) GetCode() *string {
	return s.Code
}

func (s *AppendCasesResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *AppendCasesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AppendCasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AppendCasesResponseBody) SetCode(v string) *AppendCasesResponseBody {
	s.Code = &v
	return s
}

func (s *AppendCasesResponseBody) SetHttpStatusCode(v string) *AppendCasesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AppendCasesResponseBody) SetMessage(v string) *AppendCasesResponseBody {
	s.Message = &v
	return s
}

func (s *AppendCasesResponseBody) SetRequestId(v string) *AppendCasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AppendCasesResponseBody) Validate() error {
	return dara.Validate(s)
}
