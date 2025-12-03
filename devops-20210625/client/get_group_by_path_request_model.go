// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupByPathRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentity(v string) *GetGroupByPathRequest
	GetIdentity() *string
	SetOrganizationId(v string) *GetGroupByPathRequest
	GetOrganizationId() *string
}

type GetGroupByPathRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957/test-group
	Identity *string `json:"identity,omitempty" xml:"identity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s GetGroupByPathRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGroupByPathRequest) GoString() string {
	return s.String()
}

func (s *GetGroupByPathRequest) GetIdentity() *string {
	return s.Identity
}

func (s *GetGroupByPathRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetGroupByPathRequest) SetIdentity(v string) *GetGroupByPathRequest {
	s.Identity = &v
	return s
}

func (s *GetGroupByPathRequest) SetOrganizationId(v string) *GetGroupByPathRequest {
	s.OrganizationId = &v
	return s
}

func (s *GetGroupByPathRequest) Validate() error {
	return dara.Validate(s)
}
