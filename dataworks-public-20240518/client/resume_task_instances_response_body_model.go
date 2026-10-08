// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeTaskInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResumeTaskInstancesResponseBody
	GetRequestId() *string
	SetSuccessInfo(v map[string]*SuccessInfoValue) *ResumeTaskInstancesResponseBody
	GetSuccessInfo() map[string]*SuccessInfoValue
}

type ResumeTaskInstancesResponseBody struct {
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the batch operation, which is in the MAP structure. The instance ID serves as a key, and the result serves as a value.
	SuccessInfo map[string]*SuccessInfoValue `json:"SuccessInfo,omitempty" xml:"SuccessInfo,omitempty"`
}

func (s ResumeTaskInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeTaskInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeTaskInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeTaskInstancesResponseBody) GetSuccessInfo() map[string]*SuccessInfoValue {
	return s.SuccessInfo
}

func (s *ResumeTaskInstancesResponseBody) SetRequestId(v string) *ResumeTaskInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeTaskInstancesResponseBody) SetSuccessInfo(v map[string]*SuccessInfoValue) *ResumeTaskInstancesResponseBody {
	s.SuccessInfo = v
	return s
}

func (s *ResumeTaskInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
