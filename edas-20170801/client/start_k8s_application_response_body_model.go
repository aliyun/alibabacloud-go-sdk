// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartK8sApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *StartK8sApplicationResponseBody
	GetChangeOrderId() *string
	SetCode(v int32) *StartK8sApplicationResponseBody
	GetCode() *int32
	SetMessage(v string) *StartK8sApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartK8sApplicationResponseBody
	GetRequestId() *string
}

type StartK8sApplicationResponseBody struct {
	// The ID of the change process.
	//
	// example:
	//
	// *********d237-4827-a4f4-ed2ae98de18d
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
	// 03FD1520-0FD6-436A-****-265318D7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartK8sApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartK8sApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *StartK8sApplicationResponseBody) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *StartK8sApplicationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *StartK8sApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartK8sApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartK8sApplicationResponseBody) SetChangeOrderId(v string) *StartK8sApplicationResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *StartK8sApplicationResponseBody) SetCode(v int32) *StartK8sApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *StartK8sApplicationResponseBody) SetMessage(v string) *StartK8sApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *StartK8sApplicationResponseBody) SetRequestId(v string) *StartK8sApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartK8sApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
