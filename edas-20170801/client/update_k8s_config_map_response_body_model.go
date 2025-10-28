// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateK8sConfigMapResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateK8sConfigMapResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateK8sConfigMapResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateK8sConfigMapResponseBody
	GetRequestId() *string
}

type UpdateK8sConfigMapResponseBody struct {
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
	// b197-40ab-9155-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateK8sConfigMapResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateK8sConfigMapResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateK8sConfigMapResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateK8sConfigMapResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateK8sConfigMapResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateK8sConfigMapResponseBody) SetCode(v int32) *UpdateK8sConfigMapResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateK8sConfigMapResponseBody) SetMessage(v string) *UpdateK8sConfigMapResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateK8sConfigMapResponseBody) SetRequestId(v string) *UpdateK8sConfigMapResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateK8sConfigMapResponseBody) Validate() error {
	return dara.Validate(s)
}
