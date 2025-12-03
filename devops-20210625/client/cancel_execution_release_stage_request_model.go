// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelExecutionReleaseStageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrganizationId(v string) *CancelExecutionReleaseStageRequest
	GetOrganizationId() *string
}

type CancelExecutionReleaseStageRequest struct {
	// example:
	//
	// 66c0c9fffeb86b450c19****
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s CancelExecutionReleaseStageRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelExecutionReleaseStageRequest) GoString() string {
	return s.String()
}

func (s *CancelExecutionReleaseStageRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CancelExecutionReleaseStageRequest) SetOrganizationId(v string) *CancelExecutionReleaseStageRequest {
	s.OrganizationId = &v
	return s
}

func (s *CancelExecutionReleaseStageRequest) Validate() error {
	return dara.Validate(s)
}
