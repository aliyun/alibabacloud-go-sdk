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
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the batch operation, which is in the MAP structure. The instance ID serves as a key, and the result serves as a value.
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
