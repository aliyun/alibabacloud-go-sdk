// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLockSecretNoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *LockSecretNoResponseBody
	GetCode() *string
	SetMessage(v string) *LockSecretNoResponseBody
	GetMessage() *string
	SetRequestId(v string) *LockSecretNoResponseBody
	GetRequestId() *string
}

type LockSecretNoResponseBody struct {
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

func (s LockSecretNoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LockSecretNoResponseBody) GoString() string {
	return s.String()
}

func (s *LockSecretNoResponseBody) GetCode() *string {
	return s.Code
}

func (s *LockSecretNoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *LockSecretNoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LockSecretNoResponseBody) SetCode(v string) *LockSecretNoResponseBody {
	s.Code = &v
	return s
}

func (s *LockSecretNoResponseBody) SetMessage(v string) *LockSecretNoResponseBody {
	s.Message = &v
	return s
}

func (s *LockSecretNoResponseBody) SetRequestId(v string) *LockSecretNoResponseBody {
	s.RequestId = &v
	return s
}

func (s *LockSecretNoResponseBody) Validate() error {
	return dara.Validate(s)
}
