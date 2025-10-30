// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockSecretNoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UnlockSecretNoResponseBody
	GetCode() *string
	SetMessage(v string) *UnlockSecretNoResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnlockSecretNoResponseBody
	GetRequestId() *string
}

type UnlockSecretNoResponseBody struct {
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
	// 2D1AEB96-96D0-454E-B0DC-AE2A8DF08020
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnlockSecretNoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnlockSecretNoResponseBody) GoString() string {
	return s.String()
}

func (s *UnlockSecretNoResponseBody) GetCode() *string {
	return s.Code
}

func (s *UnlockSecretNoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnlockSecretNoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnlockSecretNoResponseBody) SetCode(v string) *UnlockSecretNoResponseBody {
	s.Code = &v
	return s
}

func (s *UnlockSecretNoResponseBody) SetMessage(v string) *UnlockSecretNoResponseBody {
	s.Message = &v
	return s
}

func (s *UnlockSecretNoResponseBody) SetRequestId(v string) *UnlockSecretNoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnlockSecretNoResponseBody) Validate() error {
	return dara.Validate(s)
}
