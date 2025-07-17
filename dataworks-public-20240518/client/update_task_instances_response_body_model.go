// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTaskInstancesResponseBody
	GetRequestId() *string
	SetSuccessInfo(v map[string]*SuccessInfoValue) *UpdateTaskInstancesResponseBody
	GetSuccessInfo() map[string]*SuccessInfoValue
}

type UpdateTaskInstancesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the batch operation, which is in the MAP structure. The instance ID serves as a key, and the result serves as a value.
	SuccessInfo map[string]*SuccessInfoValue `json:"SuccessInfo,omitempty" xml:"SuccessInfo,omitempty"`
}

func (s UpdateTaskInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTaskInstancesResponseBody) GetSuccessInfo() map[string]*SuccessInfoValue {
	return s.SuccessInfo
}

func (s *UpdateTaskInstancesResponseBody) SetRequestId(v string) *UpdateTaskInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTaskInstancesResponseBody) SetSuccessInfo(v map[string]*SuccessInfoValue) *UpdateTaskInstancesResponseBody {
	s.SuccessInfo = v
	return s
}

func (s *UpdateTaskInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
