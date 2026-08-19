// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTaskInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopTaskInstancesResponseBody
	GetRequestId() *string
	SetSuccessInfo(v map[string]*SuccessInfoValue) *StopTaskInstancesResponseBody
	GetSuccessInfo() map[string]*SuccessInfoValue
}

type StopTaskInstancesResponseBody struct {
	// The request ID, which is used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result information of the batch operation. The structure is a map in which instance IDs are used as keys and result information is used as values.
	SuccessInfo map[string]*SuccessInfoValue `json:"SuccessInfo,omitempty" xml:"SuccessInfo,omitempty"`
}

func (s StopTaskInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopTaskInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *StopTaskInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopTaskInstancesResponseBody) GetSuccessInfo() map[string]*SuccessInfoValue {
	return s.SuccessInfo
}

func (s *StopTaskInstancesResponseBody) SetRequestId(v string) *StopTaskInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopTaskInstancesResponseBody) SetSuccessInfo(v map[string]*SuccessInfoValue) *StopTaskInstancesResponseBody {
	s.SuccessInfo = v
	return s
}

func (s *StopTaskInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
