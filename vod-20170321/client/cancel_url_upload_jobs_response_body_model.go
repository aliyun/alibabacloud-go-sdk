// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelUrlUploadJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCanceledJobs(v []*string) *CancelUrlUploadJobsResponseBody
	GetCanceledJobs() []*string
	SetNonExists(v []*string) *CancelUrlUploadJobsResponseBody
	GetNonExists() []*string
	SetRequestId(v string) *CancelUrlUploadJobsResponseBody
	GetRequestId() *string
}

type CancelUrlUploadJobsResponseBody struct {
	// The IDs of canceled jobs.
	CanceledJobs []*string `json:"CanceledJobs,omitempty" xml:"CanceledJobs,omitempty" type:"Repeated"`
	// The jobs that do not exist.
	NonExists []*string `json:"NonExists,omitempty" xml:"NonExists,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4D5C-3C3D-D7393642****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelUrlUploadJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelUrlUploadJobsResponseBody) GoString() string {
	return s.String()
}

func (s *CancelUrlUploadJobsResponseBody) GetCanceledJobs() []*string {
	return s.CanceledJobs
}

func (s *CancelUrlUploadJobsResponseBody) GetNonExists() []*string {
	return s.NonExists
}

func (s *CancelUrlUploadJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelUrlUploadJobsResponseBody) SetCanceledJobs(v []*string) *CancelUrlUploadJobsResponseBody {
	s.CanceledJobs = v
	return s
}

func (s *CancelUrlUploadJobsResponseBody) SetNonExists(v []*string) *CancelUrlUploadJobsResponseBody {
	s.NonExists = v
	return s
}

func (s *CancelUrlUploadJobsResponseBody) SetRequestId(v string) *CancelUrlUploadJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelUrlUploadJobsResponseBody) Validate() error {
	return dara.Validate(s)
}
