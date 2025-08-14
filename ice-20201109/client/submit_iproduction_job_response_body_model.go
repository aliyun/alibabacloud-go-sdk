// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitIProductionJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitIProductionJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitIProductionJobResponseBody
	GetRequestId() *string
}

type SubmitIProductionJobResponseBody struct {
	// The ID of the intelligent production job.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C1849434-FC47-5DC1-92B6-F7EAAFE3851E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitIProductionJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitIProductionJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitIProductionJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitIProductionJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitIProductionJobResponseBody) SetJobId(v string) *SubmitIProductionJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitIProductionJobResponseBody) SetRequestId(v string) *SubmitIProductionJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitIProductionJobResponseBody) Validate() error {
	return dara.Validate(s)
}
