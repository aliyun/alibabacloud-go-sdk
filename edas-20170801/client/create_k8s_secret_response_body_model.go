// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateK8sSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateK8sSecretResponseBody
	GetCode() *int32
	SetMessage(v string) *CreateK8sSecretResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateK8sSecretResponseBody
	GetRequestId() *string
}

type CreateK8sSecretResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4D9F-DR94-FD****************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateK8sSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateK8sSecretResponseBody) GoString() string {
	return s.String()
}

func (s *CreateK8sSecretResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateK8sSecretResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateK8sSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateK8sSecretResponseBody) SetCode(v int32) *CreateK8sSecretResponseBody {
	s.Code = &v
	return s
}

func (s *CreateK8sSecretResponseBody) SetMessage(v string) *CreateK8sSecretResponseBody {
	s.Message = &v
	return s
}

func (s *CreateK8sSecretResponseBody) SetRequestId(v string) *CreateK8sSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateK8sSecretResponseBody) Validate() error {
	return dara.Validate(s)
}
