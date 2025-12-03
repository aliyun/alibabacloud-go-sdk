// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefuseReleaseStagePipelineValidateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *RefuseReleaseStagePipelineValidateRequest
	GetJobId() *string
	SetOrganizationId(v string) *RefuseReleaseStagePipelineValidateRequest
	GetOrganizationId() *string
}

type RefuseReleaseStagePipelineValidateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 226241***
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// example:
	//
	// 66c0c9fffeb86b450c19****
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s RefuseReleaseStagePipelineValidateRequest) String() string {
	return dara.Prettify(s)
}

func (s RefuseReleaseStagePipelineValidateRequest) GoString() string {
	return s.String()
}

func (s *RefuseReleaseStagePipelineValidateRequest) GetJobId() *string {
	return s.JobId
}

func (s *RefuseReleaseStagePipelineValidateRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *RefuseReleaseStagePipelineValidateRequest) SetJobId(v string) *RefuseReleaseStagePipelineValidateRequest {
	s.JobId = &v
	return s
}

func (s *RefuseReleaseStagePipelineValidateRequest) SetOrganizationId(v string) *RefuseReleaseStagePipelineValidateRequest {
	s.OrganizationId = &v
	return s
}

func (s *RefuseReleaseStagePipelineValidateRequest) Validate() error {
	return dara.Validate(s)
}
