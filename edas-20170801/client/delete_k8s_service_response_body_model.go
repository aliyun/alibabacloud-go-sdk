// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteK8sServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *DeleteK8sServiceResponseBody
	GetChangeOrderId() *string
	SetCode(v int32) *DeleteK8sServiceResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteK8sServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteK8sServiceResponseBody
	GetRequestId() *string
}

type DeleteK8sServiceResponseBody struct {
	// The ID of the change process.
	//
	// example:
	//
	// b4b37bde-a125-43fc-****-741f7f4a9ae3
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
	// 9041389c-*****-459c-8253-724bca7f51f0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteK8sServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteK8sServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteK8sServiceResponseBody) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *DeleteK8sServiceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteK8sServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteK8sServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteK8sServiceResponseBody) SetChangeOrderId(v string) *DeleteK8sServiceResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *DeleteK8sServiceResponseBody) SetCode(v int32) *DeleteK8sServiceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteK8sServiceResponseBody) SetMessage(v string) *DeleteK8sServiceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteK8sServiceResponseBody) SetRequestId(v string) *DeleteK8sServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteK8sServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
