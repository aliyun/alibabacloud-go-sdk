// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendTaskInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SuspendTaskInstancesResponseBody
	GetRequestId() *string
	SetSuccessInfo(v map[string]*SuccessInfoValue) *SuspendTaskInstancesResponseBody
	GetSuccessInfo() map[string]*SuccessInfoValue
}

type SuspendTaskInstancesResponseBody struct {
	// The ID of the request. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result information of the batch operation. The structure is a map, where the key is the task instance ID and the value is the result information.
	SuccessInfo map[string]*SuccessInfoValue `json:"SuccessInfo,omitempty" xml:"SuccessInfo,omitempty"`
}

func (s SuspendTaskInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SuspendTaskInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendTaskInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SuspendTaskInstancesResponseBody) GetSuccessInfo() map[string]*SuccessInfoValue {
	return s.SuccessInfo
}

func (s *SuspendTaskInstancesResponseBody) SetRequestId(v string) *SuspendTaskInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendTaskInstancesResponseBody) SetSuccessInfo(v map[string]*SuccessInfoValue) *SuspendTaskInstancesResponseBody {
	s.SuccessInfo = v
	return s
}

func (s *SuspendTaskInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
