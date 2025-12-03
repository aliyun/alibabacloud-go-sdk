// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllReleaseWorkflowsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrganizationId(v string) *ListAllReleaseWorkflowsRequest
	GetOrganizationId() *string
}

type ListAllReleaseWorkflowsRequest struct {
	// example:
	//
	// 66c0c9fffeb86b450c199fcd
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s ListAllReleaseWorkflowsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAllReleaseWorkflowsRequest) GoString() string {
	return s.String()
}

func (s *ListAllReleaseWorkflowsRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListAllReleaseWorkflowsRequest) SetOrganizationId(v string) *ListAllReleaseWorkflowsRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListAllReleaseWorkflowsRequest) Validate() error {
	return dara.Validate(s)
}
