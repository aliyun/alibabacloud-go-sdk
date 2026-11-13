// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartWorkflowInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartWorkflowInstancesResponseBody
	GetRequestId() *string
	SetSuccessInfo(v map[string]*SuccessInfoValue) *StartWorkflowInstancesResponseBody
	GetSuccessInfo() map[string]*SuccessInfoValue
}

type StartWorkflowInstancesResponseBody struct {
	// The ID of the request. You can use this ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The results of the batch operation. The data structure is a map where the key is the workflow instance ID and the value is the result information.
	//
	// example:
	//
	// {"10001": {"Success": false, "Message": "The instance does not exist."}}
	SuccessInfo map[string]*SuccessInfoValue `json:"SuccessInfo,omitempty" xml:"SuccessInfo,omitempty"`
}

func (s StartWorkflowInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartWorkflowInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *StartWorkflowInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartWorkflowInstancesResponseBody) GetSuccessInfo() map[string]*SuccessInfoValue {
	return s.SuccessInfo
}

func (s *StartWorkflowInstancesResponseBody) SetRequestId(v string) *StartWorkflowInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartWorkflowInstancesResponseBody) SetSuccessInfo(v map[string]*SuccessInfoValue) *StartWorkflowInstancesResponseBody {
	s.SuccessInfo = v
	return s
}

func (s *StartWorkflowInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
