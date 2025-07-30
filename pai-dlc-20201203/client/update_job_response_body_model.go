// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *UpdateJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *UpdateJobResponseBody
	GetRequestId() *string
}

type UpdateJobResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// dlc*************
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID, which is used for diagnostics and Q\\&A.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *UpdateJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateJobResponseBody) SetJobId(v string) *UpdateJobResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateJobResponseBody) SetRequestId(v string) *UpdateJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateJobResponseBody) Validate() error {
	return dara.Validate(s)
}
