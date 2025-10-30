// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecretBlacklistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteSecretBlacklistResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteSecretBlacklistResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSecretBlacklistResponseBody
	GetRequestId() *string
}

type DeleteSecretBlacklistResponseBody struct {
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

func (s DeleteSecretBlacklistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecretBlacklistResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecretBlacklistResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteSecretBlacklistResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSecretBlacklistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSecretBlacklistResponseBody) SetCode(v string) *DeleteSecretBlacklistResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSecretBlacklistResponseBody) SetMessage(v string) *DeleteSecretBlacklistResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSecretBlacklistResponseBody) SetRequestId(v string) *DeleteSecretBlacklistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSecretBlacklistResponseBody) Validate() error {
	return dara.Validate(s)
}
