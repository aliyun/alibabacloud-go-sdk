// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManageTerraformStateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *ManageTerraformStateResponseBody
	GetJobId() *string
	SetRequestId(v string) *ManageTerraformStateResponseBody
	GetRequestId() *string
}

type ManageTerraformStateResponseBody struct {
	// example:
	//
	// job-5fd38c9xxxxx
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0A80B4F1-4E8C-515A-B3C1-0E1A9B85B6A7
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ManageTerraformStateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ManageTerraformStateResponseBody) GoString() string {
	return s.String()
}

func (s *ManageTerraformStateResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *ManageTerraformStateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ManageTerraformStateResponseBody) SetJobId(v string) *ManageTerraformStateResponseBody {
	s.JobId = &v
	return s
}

func (s *ManageTerraformStateResponseBody) SetRequestId(v string) *ManageTerraformStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ManageTerraformStateResponseBody) Validate() error {
	return dara.Validate(s)
}
