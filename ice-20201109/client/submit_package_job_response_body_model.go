// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitPackageJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitPackageJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitPackageJobResponseBody
	GetRequestId() *string
}

type SubmitPackageJobResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// 2d705f385b704ee5b*******a36d93e0
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 31E30781-9495-5E2D-A84D-759B0A01E262
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitPackageJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitPackageJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitPackageJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitPackageJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitPackageJobResponseBody) SetJobId(v string) *SubmitPackageJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitPackageJobResponseBody) SetRequestId(v string) *SubmitPackageJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitPackageJobResponseBody) Validate() error {
	return dara.Validate(s)
}
