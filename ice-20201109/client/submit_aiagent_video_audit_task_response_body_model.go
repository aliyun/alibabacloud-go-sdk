// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIAgentVideoAuditTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitAIAgentVideoAuditTaskResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitAIAgentVideoAuditTaskResponseBody
	GetRequestId() *string
}

type SubmitAIAgentVideoAuditTaskResponseBody struct {
	// example:
	//
	// **********fb04483915d4f2**********
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// **********-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitAIAgentVideoAuditTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIAgentVideoAuditTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAIAgentVideoAuditTaskResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitAIAgentVideoAuditTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAIAgentVideoAuditTaskResponseBody) SetJobId(v string) *SubmitAIAgentVideoAuditTaskResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitAIAgentVideoAuditTaskResponseBody) SetRequestId(v string) *SubmitAIAgentVideoAuditTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAIAgentVideoAuditTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
