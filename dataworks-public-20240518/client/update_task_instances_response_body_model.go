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
	// The ID of the request. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result information of the batch operation. The data structure is a map, where the key is the node instance ID and the value is the result information.
	//
	// example:
	//
	// {"10001": {"Success": false, "Message": "The instance does not exist."}}
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
