// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployK8sApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *DeployK8sApplicationResponseBody
	GetChangeOrderId() *string
	SetCode(v int32) *DeployK8sApplicationResponseBody
	GetCode() *int32
	SetMessage(v string) *DeployK8sApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeployK8sApplicationResponseBody
	GetRequestId() *string
}

type DeployK8sApplicationResponseBody struct {
	// The ID of the change process. You can call the GetChangeOrderInfo operation to query the change process ID. For more information, see [GetChangeOrderInfo](https://help.aliyun.com/document_detail/62072.html).
	//
	// example:
	//
	// cd65b247-****-475b-ad4b-7039040d625c
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
	// a5281053-08e4-47a5-b2ab-5c0323de*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeployK8sApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeployK8sApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeployK8sApplicationResponseBody) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *DeployK8sApplicationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeployK8sApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeployK8sApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployK8sApplicationResponseBody) SetChangeOrderId(v string) *DeployK8sApplicationResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *DeployK8sApplicationResponseBody) SetCode(v int32) *DeployK8sApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *DeployK8sApplicationResponseBody) SetMessage(v string) *DeployK8sApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *DeployK8sApplicationResponseBody) SetRequestId(v string) *DeployK8sApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeployK8sApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
