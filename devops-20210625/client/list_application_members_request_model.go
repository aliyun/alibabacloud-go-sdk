// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrganizationId(v string) *ListApplicationMembersRequest
	GetOrganizationId() *string
}

type ListApplicationMembersRequest struct {
	// example:
	//
	// 66c0c9fffeb86b450c199fcd
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s ListApplicationMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationMembersRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationMembersRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListApplicationMembersRequest) SetOrganizationId(v string) *ListApplicationMembersRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListApplicationMembersRequest) Validate() error {
	return dara.Validate(s)
}
