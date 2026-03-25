// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteSecretResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteSecretResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSecretResponseBody
	GetRequestId() *string
}

type DeleteSecretResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD4F1D33-7267-5596-8F5D-C593xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecretResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecretResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteSecretResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSecretResponseBody) SetCode(v string) *DeleteSecretResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSecretResponseBody) SetMessage(v string) *DeleteSecretResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSecretResponseBody) SetRequestId(v string) *DeleteSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSecretResponseBody) Validate() error {
	return dara.Validate(s)
}
