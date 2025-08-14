// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPackageJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetPackageJobRequest
	GetJobId() *string
}

type GetPackageJobRequest struct {
	// The job ID. You can obtain the job ID from the response parameters of the [SubmitPackageJob](https://help.aliyun.com/document_detail/461964.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// ab4802364a2e49208c99efab82dfa8e8
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetPackageJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPackageJobRequest) GoString() string {
	return s.String()
}

func (s *GetPackageJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetPackageJobRequest) SetJobId(v string) *GetPackageJobRequest {
	s.JobId = &v
	return s
}

func (s *GetPackageJobRequest) Validate() error {
	return dara.Validate(s)
}
