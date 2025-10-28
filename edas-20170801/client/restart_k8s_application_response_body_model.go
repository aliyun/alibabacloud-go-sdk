// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartK8sApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *RestartK8sApplicationResponseBody
	GetChangeOrderId() *string
	SetCode(v int32) *RestartK8sApplicationResponseBody
	GetCode() *int32
	SetMessage(v string) *RestartK8sApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *RestartK8sApplicationResponseBody
	GetRequestId() *string
}

type RestartK8sApplicationResponseBody struct {
	// The ID of the change process.
	//
	// example:
	//
	// *********-ed2ae98de18d
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

func (s RestartK8sApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartK8sApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *RestartK8sApplicationResponseBody) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *RestartK8sApplicationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *RestartK8sApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RestartK8sApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartK8sApplicationResponseBody) SetChangeOrderId(v string) *RestartK8sApplicationResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *RestartK8sApplicationResponseBody) SetCode(v int32) *RestartK8sApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *RestartK8sApplicationResponseBody) SetMessage(v string) *RestartK8sApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *RestartK8sApplicationResponseBody) SetRequestId(v string) *RestartK8sApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartK8sApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
