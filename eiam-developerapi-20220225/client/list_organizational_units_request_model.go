// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationalUnitsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListOrganizationalUnitsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListOrganizationalUnitsRequest
	GetPageSize() *int32
	SetParentId(v string) *ListOrganizationalUnitsRequest
	GetParentId() *string
}

type ListOrganizationalUnitsRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 20. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The ID of the parent organizational unit.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
}

func (s ListOrganizationalUnitsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationalUnitsRequest) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListOrganizationalUnitsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOrganizationalUnitsRequest) GetParentId() *string {
	return s.ParentId
}

func (s *ListOrganizationalUnitsRequest) SetPageNumber(v int32) *ListOrganizationalUnitsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOrganizationalUnitsRequest) SetPageSize(v int32) *ListOrganizationalUnitsRequest {
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
