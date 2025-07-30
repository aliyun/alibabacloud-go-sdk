// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsForOrganizationalUnitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationIds(v []*string) *ListApplicationsForOrganizationalUnitRequest
	GetApplicationIds() []*string
	SetInstanceId(v string) *ListApplicationsForOrganizationalUnitRequest
	GetInstanceId() *string
	SetOrganizationalUnitId(v string) *ListApplicationsForOrganizationalUnitRequest
	GetOrganizationalUnitId() *string
	SetPageNumber(v int64) *ListApplicationsForOrganizationalUnitRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListApplicationsForOrganizationalUnitRequest
	GetPageSize() *int64
}

type ListApplicationsForOrganizationalUnitRequest struct {
	// The IDs of the applications that the EIAM organization can access. You can query a maximum of 100 application IDs at a time.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the EIAM organization.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
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

func (s ListApplicationsForOrganizationalUnitRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForOrganizationalUnitRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsForOrganizationalUnitRequest) GetApplicationIds() []*string {
	return s.ApplicationIds
}

func (s *ListApplicationsForOrganizationalUnitRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationsForOrganizationalUnitRequest) GetOrganizationalUnitId() *string {
	return s.OrganizationalUnitId
}

func (s *ListApplicationsForOrganizationalUnitRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListApplicationsForOrganizationalUnitRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListApplicationsForOrganizationalUnitRequest) SetApplicationIds(v []*string) *ListApplicationsForOrganizationalUnitRequest {
	s.ApplicationIds = v
	return s
}

func (s *ListApplicationsForOrganizationalUnitRequest) SetInstanceId(v string) *ListApplicationsForOrganizationalUnitRequest {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationsForOrganizationalUnitRequest) SetOrganizationalUnitId(v string) *ListApplicationsForOrganizationalUnitRequest {
	s.OrganizationalUnitId = &v
	return s
}

func (s *ListApplicationsForOrganizationalUnitRequest) SetPageNumber(v int64) *ListApplicationsForOrganizationalUnitRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationsForOrganizationalUnitRequest) SetPageSize(v int64) *ListApplicationsForOrganizationalUnitRequest {
	s.PageSize = &v
	return s
}

func (s *ListApplicationsForOrganizationalUnitRequest) Validate() error {
	return dara.Validate(s)
}
