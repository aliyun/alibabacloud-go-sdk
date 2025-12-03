// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReleaseStagePipelineRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrganizationId(v string) *GetReleaseStagePipelineRunRequest
	GetOrganizationId() *string
}

type GetReleaseStagePipelineRunRequest struct {
	// example:
	//
	// 66c0c9fffeb86b450c199fcd
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s GetReleaseStagePipelineRunRequest) String() string {
	return dara.Prettify(s)
}

func (s GetReleaseStagePipelineRunRequest) GoString() string {
	return s.String()
}

func (s *GetReleaseStagePipelineRunRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetReleaseStagePipelineRunRequest) SetOrganizationId(v string) *GetReleaseStagePipelineRunRequest {
	s.OrganizationId = &v
	return s
}

func (s *GetReleaseStagePipelineRunRequest) Validate() error {
	return dara.Validate(s)
}
