// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrganizationId(v string) *GetApplicationRequest
	GetOrganizationId() *string
}

type GetApplicationRequest struct {
	// example:
	//
	// 66c0c9fffeb86b450c199fcd
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s GetApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetApplicationRequest) SetOrganizationId(v string) *GetApplicationRequest {
	s.OrganizationId = &v
	return s
}

func (s *GetApplicationRequest) Validate() error {
	return dara.Validate(s)
}
