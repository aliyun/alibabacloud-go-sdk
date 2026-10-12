// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopWorkflowInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopWorkflowInstancesResponseBody
	GetRequestId() *string
	SetSuccessInfo(v map[string]*SuccessInfoValue) *StopWorkflowInstancesResponseBody
	GetSuccessInfo() map[string]*SuccessInfoValue
}

type StopWorkflowInstancesResponseBody struct {
	// The ID of the request. You can use this ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result information of the batch operation. The structure is a map, where the key is the workflow instance ID and the value is the result information.
	SuccessInfo map[string]*SuccessInfoValue `json:"SuccessInfo,omitempty" xml:"SuccessInfo,omitempty"`
}

func (s StopWorkflowInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopWorkflowInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *StopWorkflowInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopWorkflowInstancesResponseBody) GetSuccessInfo() map[string]*SuccessInfoValue {
	return s.SuccessInfo
}

func (s *StopWorkflowInstancesResponseBody) SetRequestId(v string) *StopWorkflowInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopWorkflowInstancesResponseBody) SetSuccessInfo(v map[string]*SuccessInfoValue) *StopWorkflowInstancesResponseBody {
	s.SuccessInfo = v
	return s
}

func (s *StopWorkflowInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
