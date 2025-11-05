// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAccountInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SetAccountInfoResponseBody
	GetCode() *string
	SetMessage(v string) *SetAccountInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetAccountInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetAccountInfoResponseBody
	GetSuccess() *bool
}

type SetAccountInfoResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetAccountInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetAccountInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SetAccountInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *SetAccountInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetAccountInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetAccountInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetAccountInfoResponseBody) SetCode(v string) *SetAccountInfoResponseBody {
	s.Code = &v
	return s
}

func (s *SetAccountInfoResponseBody) SetMessage(v string) *SetAccountInfoResponseBody {
	s.Message = &v
	return s
}

func (s *SetAccountInfoResponseBody) SetRequestId(v string) *SetAccountInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAccountInfoResponseBody) SetSuccess(v bool) *SetAccountInfoResponseBody {
	s.Success = &v
	return s
}

func (s *SetAccountInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
