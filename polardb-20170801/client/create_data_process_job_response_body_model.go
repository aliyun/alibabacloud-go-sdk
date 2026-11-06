// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataProcessJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *CreateDataProcessJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *CreateDataProcessJobResponseBody
	GetRequestId() *string
}

type CreateDataProcessJobResponseBody struct {
	// The ID of the created data processing task.
	//
	// example:
	//
	// job-123456
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataProcessJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataProcessJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataProcessJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *CreateDataProcessJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataProcessJobResponseBody) SetJobId(v string) *CreateDataProcessJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateDataProcessJobResponseBody) SetRequestId(v string) *CreateDataProcessJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataProcessJobResponseBody) Validate() error {
	return dara.Validate(s)
}
