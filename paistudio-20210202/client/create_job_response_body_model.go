// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *CreateJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *CreateJobResponseBody
	GetRequestId() *string
}

type CreateJobResponseBody struct {
	// example:
	//
	// job-jcs4mmyangk7r2zdqz
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 85082123-2A3F-52E7-B27A-F04700B82FFB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *CreateJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateJobResponseBody) SetJobId(v string) *CreateJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateJobResponseBody) SetRequestId(v string) *CreateJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateJobResponseBody) Validate() error {
	return dara.Validate(s)
}
