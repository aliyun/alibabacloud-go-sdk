// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPatchOrganizationalUnitParentIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParentId(v string) *PatchOrganizationalUnitParentIdRequest
	GetParentId() *string
}

type PatchOrganizationalUnitParentIdRequest struct {
	// The ID of the parent organizational unit.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_001
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
}

func (s PatchOrganizationalUnitParentIdRequest) String() string {
	return dara.Prettify(s)
}

func (s PatchOrganizationalUnitParentIdRequest) GoString() string {
	return s.String()
}

func (s *PatchOrganizationalUnitParentIdRequest) GetParentId() *string {
	return s.ParentId
}

func (s *PatchOrganizationalUnitParentIdRequest) SetParentId(v string) *PatchOrganizationalUnitParentIdRequest {
	s.ParentId = &v
	return s
}

func (s *PatchOrganizationalUnitParentIdRequest) Validate() error {
	return dara.Validate(s)
}
