// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteK8sApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *DeleteK8sApplicationResponseBody
	GetChangeOrderId() *string
	SetCode(v int32) *DeleteK8sApplicationResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteK8sApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteK8sApplicationResponseBody
	GetRequestId() *string
}

type DeleteK8sApplicationResponseBody struct {
	// The ID of the change process for this operation. If an instance on which the application is deployed is running or a Server Load Balancer (SLB) instance is bound to the application, this operation generates a change process ID and deletes the application. You can call the GetChangeOrderInfo operation to query the progress of this operation. You can determine whether the operation is successful based on the value of the Code parameter.
	//
	// example:
	//
	// 0a34531a-****-49dc-8e7f-0cbbbfa12cf0
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
	// a5281053-08e4-47a5-b2ab-5c0323******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteK8sApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteK8sApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteK8sApplicationResponseBody) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *DeleteK8sApplicationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteK8sApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteK8sApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteK8sApplicationResponseBody) SetChangeOrderId(v string) *DeleteK8sApplicationResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *DeleteK8sApplicationResponseBody) SetCode(v int32) *DeleteK8sApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteK8sApplicationResponseBody) SetMessage(v string) *DeleteK8sApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteK8sApplicationResponseBody) SetRequestId(v string) *DeleteK8sApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteK8sApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
