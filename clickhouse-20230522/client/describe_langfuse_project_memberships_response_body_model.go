// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLangfuseProjectMembershipsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeLangfuseProjectMembershipsResponseBodyData) *DescribeLangfuseProjectMembershipsResponseBody
	GetData() *DescribeLangfuseProjectMembershipsResponseBodyData
	SetRequestId(v string) *DescribeLangfuseProjectMembershipsResponseBody
	GetRequestId() *string
}

type DescribeLangfuseProjectMembershipsResponseBody struct {
	// The returned data.
	Data *DescribeLangfuseProjectMembershipsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// A82758F8-E793-5610-BE11-0E46664305C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLangfuseProjectMembershipsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseProjectMembershipsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseProjectMembershipsResponseBody) GetData() *DescribeLangfuseProjectMembershipsResponseBodyData {
	return s.Data
}

func (s *DescribeLangfuseProjectMembershipsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLangfuseProjectMembershipsResponseBody) SetData(v *DescribeLangfuseProjectMembershipsResponseBodyData) *DescribeLangfuseProjectMembershipsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeLangfuseProjectMembershipsResponseBody) SetRequestId(v string) *DescribeLangfuseProjectMembershipsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLangfuseProjectMembershipsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLangfuseProjectMembershipsResponseBodyData struct {
	// The list of user roles in the Langfuse project.
	Memberships []*DescribeLangfuseProjectMembershipsResponseBodyDataMemberships `json:"Memberships,omitempty" xml:"Memberships,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 4
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLangfuseProjectMembershipsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseProjectMembershipsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseProjectMembershipsResponseBodyData) GetMemberships() []*DescribeLangfuseProjectMembershipsResponseBodyDataMemberships {
	return s.Memberships
}

func (s *DescribeLangfuseProjectMembershipsResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeLangfuseProjectMembershipsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeLangfuseProjectMembershipsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeLangfuseProjectMembershipsResponseBodyData) SetMemberships(v []*DescribeLangfuseProjectMembershipsResponseBodyDataMemberships) *DescribeLangfuseProjectMembershipsResponseBodyData {
	s.Memberships = v
	return s
}

func (s *DescribeLangfuseProjectMembershipsResponseBodyData) SetPageNumber(v int64) *DescribeLangfuseProjectMembershipsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *DescribeLangfuseProjectMembershipsResponseBodyData) SetPageSize(v int64) *DescribeLangfuseProjectMembershipsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeLangfuseProjectMembershipsResponseBodyData) SetTotalCount(v int64) *DescribeLangfuseProjectMembershipsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeLangfuseProjectMembershipsResponseBodyData) Validate() error {
	if s.Memberships != nil {
		for _, item := range s.Memberships {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLangfuseProjectMembershipsResponseBodyDataMemberships struct {
	// The email address of the user.
	//
	// example:
	//
	// john@company.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The username.
	//
	// example:
	//
	// john
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The role of the user in the Langfuse organization.
	//
	// example:
	//
	// ADMIN
	OrgRole *string `json:"OrgRole,omitempty" xml:"OrgRole,omitempty"`
	// The role of the user in the Langfuse project.
	//
	// example:
	//
	// VIEWER
	ProjectRole *string `json:"ProjectRole,omitempty" xml:"ProjectRole,omitempty"`
}

func (s DescribeLangfuseProjectMembershipsResponseBodyDataMemberships) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseProjectMembershipsResponseBodyDataMemberships) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseProjectMembershipsResponseBodyDataMemberships) GetEmail() *string {
	return s.Email
}

func (s *DescribeLangfuseProjectMembershipsResponseBodyDataMemberships) GetName() *string {
	return s.Name
}

func (s *DescribeLangfuseProjectMembershipsResponseBodyDataMemberships) GetOrgRole() *string {
	return s.OrgRole
}

func (s *DescribeLangfuseProjectMembershipsResponseBodyDataMemberships) GetProjectRole() *string {
	return s.ProjectRole
}

func (s *DescribeLangfuseProjectMembershipsResponseBodyDataMemberships) SetEmail(v string) *DescribeLangfuseProjectMembershipsResponseBodyDataMemberships {
	s.Email = &v
	return s
}

func (s *DescribeLangfuseProjectMembershipsResponseBodyDataMemberships) SetName(v string) *DescribeLangfuseProjectMembershipsResponseBodyDataMemberships {
	s.Name = &v
	return s
}

func (s *DescribeLangfuseProjectMembershipsResponseBodyDataMemberships) SetOrgRole(v string) *DescribeLangfuseProjectMembershipsResponseBodyDataMemberships {
	s.OrgRole = &v
	return s
}

func (s *DescribeLangfuseProjectMembershipsResponseBodyDataMemberships) SetProjectRole(v string) *DescribeLangfuseProjectMembershipsResponseBodyDataMemberships {
	s.ProjectRole = &v
	return s
}

func (s *DescribeLangfuseProjectMembershipsResponseBodyDataMemberships) Validate() error {
	return dara.Validate(s)
}
