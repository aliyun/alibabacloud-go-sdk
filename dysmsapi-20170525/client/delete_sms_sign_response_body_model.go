// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSmsSignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteSmsSignResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteSmsSignResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSmsSignResponseBody
	GetRequestId() *string
	SetSignName(v string) *DeleteSmsSignResponseBody
	GetSignName() *string
}

type DeleteSmsSignResponseBody struct {
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- Other values indicate that the request fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The signature.
	//
	// example:
	//
	// Aliyun
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s DeleteSmsSignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSmsSignResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSmsSignResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteSmsSignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSmsSignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSmsSignResponseBody) GetSignName() *string {
	return s.SignName
}

func (s *DeleteSmsSignResponseBody) SetCode(v string) *DeleteSmsSignResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSmsSignResponseBody) SetMessage(v string) *DeleteSmsSignResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSmsSignResponseBody) SetRequestId(v string) *DeleteSmsSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSmsSignResponseBody) SetSignName(v string) *DeleteSmsSignResponseBody {
	s.SignName = &v
	return s
}

func (s *DeleteSmsSignResponseBody) Validate() error {
	return dara.Validate(s)
}
