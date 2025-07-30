// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationalUnitParentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParents(v []*ListOrganizationalUnitParentsResponseBodyParents) *ListOrganizationalUnitParentsResponseBody
	GetParents() []*ListOrganizationalUnitParentsResponseBodyParents
	SetRequestId(v string) *ListOrganizationalUnitParentsResponseBody
	GetRequestId() *string
}

type ListOrganizationalUnitParentsResponseBody struct {
	// The parent organizations.
	Parents []*ListOrganizationalUnitParentsResponseBodyParents `json:"Parents,omitempty" xml:"Parents,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListOrganizationalUnitParentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationalUnitParentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitParentsResponseBody) GetParents() []*ListOrganizationalUnitParentsResponseBodyParents {
	return s.Parents
}

func (s *ListOrganizationalUnitParentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOrganizationalUnitParentsResponseBody) SetParents(v []*ListOrganizationalUnitParentsResponseBodyParents) *ListOrganizationalUnitParentsResponseBody {
	s.Parents = v
	return s
}

func (s *ListOrganizationalUnitParentsResponseBody) SetRequestId(v string) *ListOrganizationalUnitParentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrganizationalUnitParentsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListOrganizationalUnitParentsResponseBodyParents struct {
	// The organization ID.
	//
	// example:
	//
	// ou_4lag76zc2km5ssg5vsmm2xxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
	// The parent organization ID.
	//
	// example:
	//
	// ou_x3beoyepv2ls5iwuge3xhjxxxx
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s ListOrganizationalUnitParentsResponseBodyParents) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationalUnitParentsResponseBodyParents) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitParentsResponseBodyParents) GetOrganizationalUnitId() *string {
	return s.OrganizationalUnitId
}

func (s *ListOrganizationalUnitParentsResponseBodyParents) GetParentId() *string {
	return s.ParentId
}

func (s *ListOrganizationalUnitParentsResponseBodyParents) SetOrganizationalUnitId(v string) *ListOrganizationalUnitParentsResponseBodyParents {
	s.OrganizationalUnitId = &v
	return s
}

func (s *ListOrganizationalUnitParentsResponseBodyParents) SetParentId(v string) *ListOrganizationalUnitParentsResponseBodyParents {
	s.ParentId = &v
	return s
}

func (s *ListOrganizationalUnitParentsResponseBodyParents) Validate() error {
	return dara.Validate(s)
}
