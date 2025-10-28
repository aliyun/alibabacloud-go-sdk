// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateK8sResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateK8sResourceResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateK8sResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateK8sResourceResponseBody
	GetRequestId() *string
}

type UpdateK8sResourceResponseBody struct {
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
	// 7638276F-****-****-884F-54CC0BC84A8D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateK8sResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateK8sResourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateK8sResourceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateK8sResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateK8sResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateK8sResourceResponseBody) SetCode(v int32) *UpdateK8sResourceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateK8sResourceResponseBody) SetMessage(v string) *UpdateK8sResourceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateK8sResourceResponseBody) SetRequestId(v string) *UpdateK8sResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateK8sResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
