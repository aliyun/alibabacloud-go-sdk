// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationalUnitsForApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrganizationalUnits(v []*ListOrganizationalUnitsForApplicationResponseBodyOrganizationalUnits) *ListOrganizationalUnitsForApplicationResponseBody
	GetOrganizationalUnits() []*ListOrganizationalUnitsForApplicationResponseBodyOrganizationalUnits
	SetRequestId(v string) *ListOrganizationalUnitsForApplicationResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListOrganizationalUnitsForApplicationResponseBody
	GetTotalCount() *int64
}

type ListOrganizationalUnitsForApplicationResponseBody struct {
	// The IDs of the organizations that are allowed to access the application.
	OrganizationalUnits []*ListOrganizationalUnitsForApplicationResponseBodyOrganizationalUnits `json:"OrganizationalUnits,omitempty" xml:"OrganizationalUnits,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the returned entries.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOrganizationalUnitsForApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationalUnitsForApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsForApplicationResponseBody) GetOrganizationalUnits() []*ListOrganizationalUnitsForApplicationResponseBodyOrganizationalUnits {
	return s.OrganizationalUnits
}

func (s *ListOrganizationalUnitsForApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOrganizationalUnitsForApplicationResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListOrganizationalUnitsForApplicationResponseBody) SetOrganizationalUnits(v []*ListOrganizationalUnitsForApplicationResponseBodyOrganizationalUnits) *ListOrganizationalUnitsForApplicationResponseBody {
	s.OrganizationalUnits = v
	return s
}

func (s *ListOrganizationalUnitsForApplicationResponseBody) SetRequestId(v string) *ListOrganizationalUnitsForApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrganizationalUnitsForApplicationResponseBody) SetTotalCount(v int64) *ListOrganizationalUnitsForApplicationResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOrganizationalUnitsForApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListOrganizationalUnitsForApplicationResponseBodyOrganizationalUnits struct {
	// The ID of the organization that is allowed to access the application.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
}

func (s ListOrganizationalUnitsForApplicationResponseBodyOrganizationalUnits) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationalUnitsForApplicationResponseBodyOrganizationalUnits) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsForApplicationResponseBodyOrganizationalUnits) GetOrganizationalUnitId() *string {
	return s.OrganizationalUnitId
}

func (s *ListOrganizationalUnitsForApplicationResponseBodyOrganizationalUnits) SetOrganizationalUnitId(v string) *ListOrganizationalUnitsForApplicationResponseBodyOrganizationalUnits {
	s.OrganizationalUnitId = &v
	return s
}

func (s *ListOrganizationalUnitsForApplicationResponseBodyOrganizationalUnits) Validate() error {
	return dara.Validate(s)
}
