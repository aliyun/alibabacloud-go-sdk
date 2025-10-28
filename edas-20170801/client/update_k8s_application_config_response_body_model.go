// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateK8sApplicationConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *UpdateK8sApplicationConfigResponseBody
	GetChangeOrderId() *string
	SetCode(v int32) *UpdateK8sApplicationConfigResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateK8sApplicationConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateK8sApplicationConfigResponseBody
	GetRequestId() *string
}

type UpdateK8sApplicationConfigResponseBody struct {
	// The ID of the change process.
	//
	// example:
	//
	// 8806d1c6-****-48eb-9373-6bdef3007466
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
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
	// 4823-bhjf-23u4-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateK8sApplicationConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateK8sApplicationConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateK8sApplicationConfigResponseBody) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *UpdateK8sApplicationConfigResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateK8sApplicationConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateK8sApplicationConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateK8sApplicationConfigResponseBody) SetChangeOrderId(v string) *UpdateK8sApplicationConfigResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *UpdateK8sApplicationConfigResponseBody) SetCode(v int32) *UpdateK8sApplicationConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateK8sApplicationConfigResponseBody) SetMessage(v string) *UpdateK8sApplicationConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateK8sApplicationConfigResponseBody) SetRequestId(v string) *UpdateK8sApplicationConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateK8sApplicationConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
