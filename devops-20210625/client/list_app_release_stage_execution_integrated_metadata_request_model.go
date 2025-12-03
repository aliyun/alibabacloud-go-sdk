// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppReleaseStageExecutionIntegratedMetadataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrganizationId(v string) *ListAppReleaseStageExecutionIntegratedMetadataRequest
	GetOrganizationId() *string
}

type ListAppReleaseStageExecutionIntegratedMetadataRequest struct {
	// example:
	//
	// 66c0c9fffeb86b450c199fcd
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s ListAppReleaseStageExecutionIntegratedMetadataRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppReleaseStageExecutionIntegratedMetadataRequest) GoString() string {
	return s.String()
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataRequest) SetOrganizationId(v string) *ListAppReleaseStageExecutionIntegratedMetadataRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataRequest) Validate() error {
	return dara.Validate(s)
}
