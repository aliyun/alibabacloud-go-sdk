// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobPlanId(v string) *CreateJobPlanResponseBody
	GetJobPlanId() *string
	SetRequestId(v string) *CreateJobPlanResponseBody
	GetRequestId() *string
}

type CreateJobPlanResponseBody struct {
	// The job plan ID.
	//
	// example:
	//
	// jp-xxxxxx
	JobPlanId *string `json:"JobPlanId,omitempty" xml:"JobPlanId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82-9624-EC2B1779848E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateJobPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateJobPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobPlanResponseBody) GetJobPlanId() *string {
	return s.JobPlanId
}

func (s *CreateJobPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateJobPlanResponseBody) SetJobPlanId(v string) *CreateJobPlanResponseBody {
	s.JobPlanId = &v
	return s
}

func (s *CreateJobPlanResponseBody) SetRequestId(v string) *CreateJobPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateJobPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
