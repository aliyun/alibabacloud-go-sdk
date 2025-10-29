// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRerunWorkflowInstancesResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetRerunWorkflowInstancesResultResponseBody
	GetRequestId() *string
	SetResult(v *GetRerunWorkflowInstancesResultResponseBodyResult) *GetRerunWorkflowInstancesResultResponseBody
	GetResult() *GetRerunWorkflowInstancesResultResponseBodyResult
}

type GetRerunWorkflowInstancesResultResponseBody struct {
	// The request ID, used for log tracing and troubleshooting.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the workflow instance rerun.
	Result *GetRerunWorkflowInstancesResultResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetRerunWorkflowInstancesResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRerunWorkflowInstancesResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetRerunWorkflowInstancesResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRerunWorkflowInstancesResultResponseBody) GetResult() *GetRerunWorkflowInstancesResultResponseBodyResult {
	return s.Result
}

func (s *GetRerunWorkflowInstancesResultResponseBody) SetRequestId(v string) *GetRerunWorkflowInstancesResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRerunWorkflowInstancesResultResponseBody) SetResult(v *GetRerunWorkflowInstancesResultResponseBodyResult) *GetRerunWorkflowInstancesResultResponseBody {
	s.Result = v
	return s
}

func (s *GetRerunWorkflowInstancesResultResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRerunWorkflowInstancesResultResponseBodyResult struct {
	// The failure message. Returned if the rerun fails.
	//
	// example:
	//
	// Invalid Param xxx
	FailureMessage *string `json:"FailureMessage,omitempty" xml:"FailureMessage,omitempty"`
	// The status. NotRun Success Failure
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetRerunWorkflowInstancesResultResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetRerunWorkflowInstancesResultResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRerunWorkflowInstancesResultResponseBodyResult) GetFailureMessage() *string {
	return s.FailureMessage
}

func (s *GetRerunWorkflowInstancesResultResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetRerunWorkflowInstancesResultResponseBodyResult) SetFailureMessage(v string) *GetRerunWorkflowInstancesResultResponseBodyResult {
	s.FailureMessage = &v
	return s
}

func (s *GetRerunWorkflowInstancesResultResponseBodyResult) SetStatus(v string) *GetRerunWorkflowInstancesResultResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetRerunWorkflowInstancesResultResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
