// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSmsSignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddSmsSignResponseBody
	GetCode() *string
	SetMessage(v string) *AddSmsSignResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddSmsSignResponseBody
	GetRequestId() *string
	SetSignName(v string) *AddSmsSignResponseBody
	GetSignName() *string
}

type AddSmsSignResponseBody struct {
	// The HTTP status code.
	//
	// 	- OK is returned if the request is successful.
	//
	// 	- For other error codes, see [API error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The signature name.
	//
	// example:
	//
	// 阿里云
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s AddSmsSignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddSmsSignResponseBody) GoString() string {
	return s.String()
}

func (s *AddSmsSignResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddSmsSignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddSmsSignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddSmsSignResponseBody) GetSignName() *string {
	return s.SignName
}

func (s *AddSmsSignResponseBody) SetCode(v string) *AddSmsSignResponseBody {
	s.Code = &v
	return s
}

func (s *AddSmsSignResponseBody) SetMessage(v string) *AddSmsSignResponseBody {
	s.Message = &v
	return s
}

func (s *AddSmsSignResponseBody) SetRequestId(v string) *AddSmsSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSmsSignResponseBody) SetSignName(v string) *AddSmsSignResponseBody {
	s.SignName = &v
	return s
}

func (s *AddSmsSignResponseBody) Validate() error {
	return dara.Validate(s)
}
