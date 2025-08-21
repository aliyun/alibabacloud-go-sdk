// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateUserAccessTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GenerateUserAccessTokenResponseBody
	GetCode() *string
	SetData(v string) *GenerateUserAccessTokenResponseBody
	GetData() *string
	SetMessage(v string) *GenerateUserAccessTokenResponseBody
	GetMessage() *string
	SetRequestId(v string) *GenerateUserAccessTokenResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GenerateUserAccessTokenResponseBody
	GetSuccess() *bool
}

type GenerateUserAccessTokenResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// DDEXEDAFWAGASDFWAEFFAWEFAWFWEAFWAFWAEF
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Parameter.Invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E6988CE6-41CF-1103-9BEC-2B20D26C0B51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateUserAccessTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateUserAccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateUserAccessTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *GenerateUserAccessTokenResponseBody) GetData() *string {
	return s.Data
}

func (s *GenerateUserAccessTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GenerateUserAccessTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateUserAccessTokenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GenerateUserAccessTokenResponseBody) SetCode(v string) *GenerateUserAccessTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateUserAccessTokenResponseBody) SetData(v string) *GenerateUserAccessTokenResponseBody {
	s.Data = &v
	return s
}

func (s *GenerateUserAccessTokenResponseBody) SetMessage(v string) *GenerateUserAccessTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateUserAccessTokenResponseBody) SetRequestId(v string) *GenerateUserAccessTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateUserAccessTokenResponseBody) SetSuccess(v bool) *GenerateUserAccessTokenResponseBody {
	s.Success = &v
	return s
}

func (s *GenerateUserAccessTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
