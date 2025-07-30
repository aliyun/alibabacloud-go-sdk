// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationalUnitsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListOrganizationalUnitsRequest
	GetInstanceId() *string
	SetOrganizationalUnitIds(v []*string) *ListOrganizationalUnitsRequest
	GetOrganizationalUnitIds() []*string
	SetOrganizationalUnitName(v string) *ListOrganizationalUnitsRequest
	GetOrganizationalUnitName() *string
	SetOrganizationalUnitNameStartsWith(v string) *ListOrganizationalUnitsRequest
	GetOrganizationalUnitNameStartsWith() *string
	SetPageNumber(v int64) *ListOrganizationalUnitsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListOrganizationalUnitsRequest
	GetPageSize() *int64
	SetParentId(v string) *ListOrganizationalUnitsRequest
	GetParentId() *string
}

type ListOrganizationalUnitsRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IDs of organizational units.
	//
	// example:
	//
	// [ou_wovwffm62xifdziem7an7xxxxx]
	OrganizationalUnitIds []*string `json:"OrganizationalUnitIds,omitempty" xml:"OrganizationalUnitIds,omitempty" type:"Repeated"`
	// The name of the organizational unit.
	//
	// example:
	//
	// name_001
	OrganizationalUnitName *string `json:"OrganizationalUnitName,omitempty" xml:"OrganizationalUnitName,omitempty"`
	// Organization name, matching left
	//
	// example:
	//
	// name
	OrganizationalUnitNameStartsWith *string `json:"OrganizationalUnitNameStartsWith,omitempty" xml:"OrganizationalUnitNameStartsWith,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the parent organizational unit.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s ListOrganizationalUnitsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationalUnitsRequest) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOrganizationalUnitsRequest) GetOrganizationalUnitIds() []*string {
	return s.OrganizationalUnitIds
}

func (s *ListOrganizationalUnitsRequest) GetOrganizationalUnitName() *string {
	return s.OrganizationalUnitName
}

func (s *ListOrganizationalUnitsRequest) GetOrganizationalUnitNameStartsWith() *string {
	return s.OrganizationalUnitNameStartsWith
}

func (s *ListOrganizationalUnitsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListOrganizationalUnitsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListOrganizationalUnitsRequest) GetParentId() *string {
	return s.ParentId
}

func (s *ListOrganizationalUnitsRequest) SetInstanceId(v string) *ListOrganizationalUnitsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListOrganizationalUnitsRequest) SetOrganizationalUnitIds(v []*string) *ListOrganizationalUnitsRequest {
	s.OrganizationalUnitIds = v
	return s
}

func (s *ListOrganizationalUnitsRequest) SetOrganizationalUnitName(v string) *ListOrganizationalUnitsRequest {
	s.OrganizationalUnitName = &v
	return s
}

func (s *ListOrganizationalUnitsRequest) SetOrganizationalUnitNameStartsWith(v string) *ListOrganizationalUnitsRequest {
	s.OrganizationalUnitNameStartsWith = &v
	return s
}

func (s *ListOrganizationalUnitsRequest) SetPageNumber(v int64) *ListOrganizationalUnitsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOrganizationalUnitsRequest) SetPageSize(v int64) *ListOrganizationalUnitsRequest {
	s.PageSize = &v
	return s
}

func (s *ListOrganizationalUnitsRequest) SetParentId(v string) *ListOrganizationalUnitsRequest {
	s.ParentId = &v
	return s
}

func (s *ListOrganizationalUnitsRequest) Validate() error {
	return dara.Validate(s)
}
