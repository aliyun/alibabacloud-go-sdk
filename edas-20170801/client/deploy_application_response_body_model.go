// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *DeployApplicationResponseBody
	GetChangeOrderId() *string
	SetCode(v int32) *DeployApplicationResponseBody
	GetCode() *int32
	SetMessage(v string) *DeployApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeployApplicationResponseBody
	GetRequestId() *string
}

type DeployApplicationResponseBody struct {
	// The change process ID of the application deployment.
	//
	// example:
	//
	// adf86823-055b-48a4-a59f-fa5582******
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned for the request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// d6834ee9-5045-*************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeployApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeployApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeployApplicationResponseBody) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *DeployApplicationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeployApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeployApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployApplicationResponseBody) SetChangeOrderId(v string) *DeployApplicationResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *DeployApplicationResponseBody) SetCode(v int32) *DeployApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *DeployApplicationResponseBody) SetMessage(v string) *DeployApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *DeployApplicationResponseBody) SetRequestId(v string) *DeployApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeployApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
