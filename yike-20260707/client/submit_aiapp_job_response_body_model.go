// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIAppJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitAIAppJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitAIAppJobResponseBody
	GetRequestId() *string
}

type SubmitAIAppJobResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// 68ca759e798b40b4903b255*******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitAIAppJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIAppJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAIAppJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitAIAppJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAIAppJobResponseBody) SetJobId(v string) *SubmitAIAppJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitAIAppJobResponseBody) SetRequestId(v string) *SubmitAIAppJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAIAppJobResponseBody) Validate() error {
	return dara.Validate(s)
}
