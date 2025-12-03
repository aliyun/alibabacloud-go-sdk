// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccountId(v string) *UpdateApplicationRequest
	GetOwnerAccountId() *string
	SetOrganizationId(v string) *UpdateApplicationRequest
	GetOrganizationId() *string
}

type UpdateApplicationRequest struct {
	// example:
	//
	// 1332695887xxxxxx
	OwnerAccountId *string `json:"ownerAccountId,omitempty" xml:"ownerAccountId,omitempty"`
	// example:
	//
	// 66c0c9fffeb86b450c199fcd
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s UpdateApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationRequest) GetOwnerAccountId() *string {
	return s.OwnerAccountId
}

func (s *UpdateApplicationRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *UpdateApplicationRequest) SetOwnerAccountId(v string) *UpdateApplicationRequest {
	s.OwnerAccountId = &v
	return s
}

func (s *UpdateApplicationRequest) SetOrganizationId(v string) *UpdateApplicationRequest {
	s.OrganizationId = &v
	return s
}

func (s *UpdateApplicationRequest) Validate() error {
	return dara.Validate(s)
}
