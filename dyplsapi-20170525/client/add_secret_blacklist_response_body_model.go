// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSecretBlacklistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddSecretBlacklistResponseBody
	GetCode() *string
	SetMessage(v string) *AddSecretBlacklistResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddSecretBlacklistResponseBody
	GetRequestId() *string
}

type AddSecretBlacklistResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other status codes indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
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
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddSecretBlacklistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddSecretBlacklistResponseBody) GoString() string {
	return s.String()
}

func (s *AddSecretBlacklistResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddSecretBlacklistResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddSecretBlacklistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddSecretBlacklistResponseBody) SetCode(v string) *AddSecretBlacklistResponseBody {
	s.Code = &v
	return s
}

func (s *AddSecretBlacklistResponseBody) SetMessage(v string) *AddSecretBlacklistResponseBody {
	s.Message = &v
	return s
}

func (s *AddSecretBlacklistResponseBody) SetRequestId(v string) *AddSecretBlacklistResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSecretBlacklistResponseBody) Validate() error {
	return dara.Validate(s)
}
