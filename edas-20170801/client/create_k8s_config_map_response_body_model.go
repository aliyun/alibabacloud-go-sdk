// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateK8sConfigMapResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateK8sConfigMapResponseBody
	GetCode() *int32
	SetMessage(v string) *CreateK8sConfigMapResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateK8sConfigMapResponseBody
	GetRequestId() *string
}

type CreateK8sConfigMapResponseBody struct {
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
	// D16979DC-4D42-**************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateK8sConfigMapResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateK8sConfigMapResponseBody) GoString() string {
	return s.String()
}

func (s *CreateK8sConfigMapResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateK8sConfigMapResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateK8sConfigMapResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateK8sConfigMapResponseBody) SetCode(v int32) *CreateK8sConfigMapResponseBody {
	s.Code = &v
	return s
}

func (s *CreateK8sConfigMapResponseBody) SetMessage(v string) *CreateK8sConfigMapResponseBody {
	s.Message = &v
	return s
}

func (s *CreateK8sConfigMapResponseBody) SetRequestId(v string) *CreateK8sConfigMapResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateK8sConfigMapResponseBody) Validate() error {
	return dara.Validate(s)
}
