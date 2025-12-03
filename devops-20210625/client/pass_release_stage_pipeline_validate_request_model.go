// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPassReleaseStagePipelineValidateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *PassReleaseStagePipelineValidateRequest
	GetJobId() *string
	SetOrganizationId(v string) *PassReleaseStagePipelineValidateRequest
	GetOrganizationId() *string
}

type PassReleaseStagePipelineValidateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 226241***
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// example:
	//
	// 66c0c9fffeb86b450c199***
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s PassReleaseStagePipelineValidateRequest) String() string {
	return dara.Prettify(s)
}

func (s PassReleaseStagePipelineValidateRequest) GoString() string {
	return s.String()
}

func (s *PassReleaseStagePipelineValidateRequest) GetJobId() *string {
	return s.JobId
}

func (s *PassReleaseStagePipelineValidateRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *PassReleaseStagePipelineValidateRequest) SetJobId(v string) *PassReleaseStagePipelineValidateRequest {
	s.JobId = &v
	return s
}

func (s *PassReleaseStagePipelineValidateRequest) SetOrganizationId(v string) *PassReleaseStagePipelineValidateRequest {
	s.OrganizationId = &v
	return s
}

func (s *PassReleaseStagePipelineValidateRequest) Validate() error {
	return dara.Validate(s)
}
