// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetUserPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ResetUserPasswordResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ResetUserPasswordResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ResetUserPasswordResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResetUserPasswordResponseBody
	GetRequestId() *string
}

type ResetUserPasswordResponseBody struct {
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
	// BA03159C-E808-4FF1-B27E-A61B6E888D7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetUserPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetUserPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResetUserPasswordResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ResetUserPasswordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResetUserPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetUserPasswordResponseBody) SetCode(v string) *ResetUserPasswordResponseBody {
	s.Code = &v
	return s
}

func (s *ResetUserPasswordResponseBody) SetHttpStatusCode(v int32) *ResetUserPasswordResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ResetUserPasswordResponseBody) SetMessage(v string) *ResetUserPasswordResponseBody {
	s.Message = &v
	return s
}

func (s *ResetUserPasswordResponseBody) SetRequestId(v string) *ResetUserPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetUserPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
