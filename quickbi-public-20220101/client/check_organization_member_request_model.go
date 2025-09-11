// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckOrganizationMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *CheckOrganizationMemberRequest
	GetUserId() *string
}

type CheckOrganizationMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// adfssd-sdf****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CheckOrganizationMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckOrganizationMemberRequest) GoString() string {
	return s.String()
}

func (s *CheckOrganizationMemberRequest) GetUserId() *string {
	return s.UserId
}

func (s *CheckOrganizationMemberRequest) SetUserId(v string) *CheckOrganizationMemberRequest {
	s.UserId = &v
	return s
}

func (s *CheckOrganizationMemberRequest) Validate() error {
	return dara.Validate(s)
}
