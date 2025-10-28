// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateK8sSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateK8sSecretResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateK8sSecretResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateK8sSecretResponseBody
	GetRequestId() *string
}

type UpdateK8sSecretResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D16979DC-4D42-****************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateK8sSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateK8sSecretResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateK8sSecretResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateK8sSecretResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateK8sSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateK8sSecretResponseBody) SetCode(v int32) *UpdateK8sSecretResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateK8sSecretResponseBody) SetMessage(v string) *UpdateK8sSecretResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateK8sSecretResponseBody) SetRequestId(v string) *UpdateK8sSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateK8sSecretResponseBody) Validate() error {
	return dara.Validate(s)
}
