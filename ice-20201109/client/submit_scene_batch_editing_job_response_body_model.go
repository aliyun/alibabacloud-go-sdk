// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSceneBatchEditingJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitSceneBatchEditingJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitSceneBatchEditingJobResponseBody
	GetRequestId() *string
}

type SubmitSceneBatchEditingJobResponseBody struct {
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// ****C702-41BE-467E-AF2E-883D4517****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitSceneBatchEditingJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitSceneBatchEditingJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSceneBatchEditingJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitSceneBatchEditingJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitSceneBatchEditingJobResponseBody) SetJobId(v string) *SubmitSceneBatchEditingJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitSceneBatchEditingJobResponseBody) SetRequestId(v string) *SubmitSceneBatchEditingJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSceneBatchEditingJobResponseBody) Validate() error {
	return dara.Validate(s)
}
