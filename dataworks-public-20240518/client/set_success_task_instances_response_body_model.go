// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSuccessTaskInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetSuccessTaskInstancesResponseBody
	GetRequestId() *string
	SetSuccessInfo(v map[string]*SuccessInfoValue) *SetSuccessTaskInstancesResponseBody
	GetSuccessInfo() map[string]*SuccessInfoValue
}

type SetSuccessTaskInstancesResponseBody struct {
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the batch operation, which is in the MAP structure. The instance ID serves as a key, and the result serves as a value.
	SuccessInfo map[string]*SuccessInfoValue `json:"SuccessInfo,omitempty" xml:"SuccessInfo,omitempty"`
}

func (s SetSuccessTaskInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetSuccessTaskInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *SetSuccessTaskInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetSuccessTaskInstancesResponseBody) GetSuccessInfo() map[string]*SuccessInfoValue {
	return s.SuccessInfo
}

func (s *SetSuccessTaskInstancesResponseBody) SetRequestId(v string) *SetSuccessTaskInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetSuccessTaskInstancesResponseBody) SetSuccessInfo(v map[string]*SuccessInfoValue) *SetSuccessTaskInstancesResponseBody {
	s.SuccessInfo = v
	return s
}

func (s *SetSuccessTaskInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
