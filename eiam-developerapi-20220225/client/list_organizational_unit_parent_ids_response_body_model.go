// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationalUnitParentIdsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParentIds(v []*string) *ListOrganizationalUnitParentIdsResponseBody
	GetParentIds() []*string
}

type ListOrganizationalUnitParentIdsResponseBody struct {
	// The IDs of the parent organizational units. The IDs of the organizational unit are ordered based on their levels from high to low. Only the IDs of the organizational units within the authorization scope are displayed.
	//
	// example:
	//
	// [ou_xxx001]
	ParentIds []*string `json:"parentIds,omitempty" xml:"parentIds,omitempty" type:"Repeated"`
}

func (s ListOrganizationalUnitParentIdsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationalUnitParentIdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitParentIdsResponseBody) GetParentIds() []*string {
	return s.ParentIds
}

func (s *ListOrganizationalUnitParentIdsResponseBody) SetParentIds(v []*string) *ListOrganizationalUnitParentIdsResponseBody {
	s.ParentIds = v
	return s
}

func (s *ListOrganizationalUnitParentIdsResponseBody) Validate() error {
	return dara.Validate(s)
}
