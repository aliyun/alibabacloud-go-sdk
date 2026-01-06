// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationalUnitsForApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ListOrganizationalUnitsForApplicationRequest
	GetApplicationId() *string
	SetApplicationRoleId(v string) *ListOrganizationalUnitsForApplicationRequest
	GetApplicationRoleId() *string
	SetInstanceId(v string) *ListOrganizationalUnitsForApplicationRequest
	GetInstanceId() *string
	SetOrganizationalUnitIds(v []*string) *ListOrganizationalUnitsForApplicationRequest
	GetOrganizationalUnitIds() []*string
	SetPageNumber(v int64) *ListOrganizationalUnitsForApplicationRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListOrganizationalUnitsForApplicationRequest
	GetPageSize() *int64
}

type ListOrganizationalUnitsForApplicationRequest struct {
	// The ID of the application that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// 应用角色ID。
	//
	// example:
	//
	// app_role_mkv7rgt4ds8d8v0qtzev2mxxxx
	ApplicationRoleId *string `json:"ApplicationRoleId,omitempty" xml:"ApplicationRoleId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IDs of the organizations that are allowed to access the application. You can query a maximum of 100 organization IDs at a time.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitIds []*string `json:"OrganizationalUnitIds,omitempty" xml:"OrganizationalUnitIds,omitempty" type:"Repeated"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListOrganizationalUnitsForApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationalUnitsForApplicationRequest) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsForApplicationRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListOrganizationalUnitsForApplicationRequest) GetApplicationRoleId() *string {
	return s.ApplicationRoleId
}

func (s *ListOrganizationalUnitsForApplicationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOrganizationalUnitsForApplicationRequest) GetOrganizationalUnitIds() []*string {
	return s.OrganizationalUnitIds
}

func (s *ListOrganizationalUnitsForApplicationRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListOrganizationalUnitsForApplicationRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListOrganizationalUnitsForApplicationRequest) SetApplicationId(v string) *ListOrganizationalUnitsForApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListOrganizationalUnitsForApplicationRequest) SetApplicationRoleId(v string) *ListOrganizationalUnitsForApplicationRequest {
	s.ApplicationRoleId = &v
	return s
}

func (s *ListOrganizationalUnitsForApplicationRequest) SetInstanceId(v string) *ListOrganizationalUnitsForApplicationRequest {
	s.InstanceId = &v
	return s
}

func (s *ListOrganizationalUnitsForApplicationRequest) SetOrganizationalUnitIds(v []*string) *ListOrganizationalUnitsForApplicationRequest {
	s.OrganizationalUnitIds = v
	return s
}

func (s *ListOrganizationalUnitsForApplicationRequest) SetPageNumber(v int64) *ListOrganizationalUnitsForApplicationRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOrganizationalUnitsForApplicationRequest) SetPageSize(v int64) *ListOrganizationalUnitsForApplicationRequest {
	s.PageSize = &v
	return s
}

func (s *ListOrganizationalUnitsForApplicationRequest) Validate() error {
	return dara.Validate(s)
}
