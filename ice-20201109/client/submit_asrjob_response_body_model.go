// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitASRJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitASRJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitASRJobResponseBody
	GetRequestId() *string
	SetState(v string) *SubmitASRJobResponseBody
	GetState() *string
}

type SubmitASRJobResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The job state. Valid values:
	//
	// 	- Created
	//
	// 	- Executing
	//
	// 	- Finished
	//
	// 	- Failed
	//
	// example:
	//
	// Finished
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s SubmitASRJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitASRJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitASRJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitASRJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitASRJobResponseBody) GetState() *string {
	return s.State
}

func (s *SubmitASRJobResponseBody) SetJobId(v string) *SubmitASRJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitASRJobResponseBody) SetRequestId(v string) *SubmitASRJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitASRJobResponseBody) SetState(v string) *SubmitASRJobResponseBody {
	s.State = &v
	return s
}

func (s *SubmitASRJobResponseBody) Validate() error {
	return dara.Validate(s)
}
