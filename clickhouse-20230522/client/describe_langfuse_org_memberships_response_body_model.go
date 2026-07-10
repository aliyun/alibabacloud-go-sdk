// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLangfuseOrgMembershipsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeLangfuseOrgMembershipsResponseBodyData) *DescribeLangfuseOrgMembershipsResponseBody
	GetData() *DescribeLangfuseOrgMembershipsResponseBodyData
	SetRequestId(v string) *DescribeLangfuseOrgMembershipsResponseBody
	GetRequestId() *string
}

type DescribeLangfuseOrgMembershipsResponseBody struct {
	// The returned result.
	Data *DescribeLangfuseOrgMembershipsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1F488A93-83FD-540F-9B67-0333AF64E6A0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLangfuseOrgMembershipsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseOrgMembershipsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseOrgMembershipsResponseBody) GetData() *DescribeLangfuseOrgMembershipsResponseBodyData {
	return s.Data
}

func (s *DescribeLangfuseOrgMembershipsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLangfuseOrgMembershipsResponseBody) SetData(v *DescribeLangfuseOrgMembershipsResponseBodyData) *DescribeLangfuseOrgMembershipsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeLangfuseOrgMembershipsResponseBody) SetRequestId(v string) *DescribeLangfuseOrgMembershipsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLangfuseOrgMembershipsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLangfuseOrgMembershipsResponseBodyData struct {
	// The list of user roles in the organization.
	Memberships []*DescribeLangfuseOrgMembershipsResponseBodyDataMemberships `json:"Memberships,omitempty" xml:"Memberships,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 3
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLangfuseOrgMembershipsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseOrgMembershipsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseOrgMembershipsResponseBodyData) GetMemberships() []*DescribeLangfuseOrgMembershipsResponseBodyDataMemberships {
	return s.Memberships
}

func (s *DescribeLangfuseOrgMembershipsResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeLangfuseOrgMembershipsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeLangfuseOrgMembershipsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeLangfuseOrgMembershipsResponseBodyData) SetMemberships(v []*DescribeLangfuseOrgMembershipsResponseBodyDataMemberships) *DescribeLangfuseOrgMembershipsResponseBodyData {
	s.Memberships = v
	return s
}

func (s *DescribeLangfuseOrgMembershipsResponseBodyData) SetPageNumber(v int64) *DescribeLangfuseOrgMembershipsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *DescribeLangfuseOrgMembershipsResponseBodyData) SetPageSize(v int64) *DescribeLangfuseOrgMembershipsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeLangfuseOrgMembershipsResponseBodyData) SetTotalCount(v int64) *DescribeLangfuseOrgMembershipsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeLangfuseOrgMembershipsResponseBodyData) Validate() error {
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

type DescribeLangfuseOrgMembershipsResponseBodyDataMemberships struct {
	// The time when the user was created.
	//
	// example:
	//
	// 2026-06-24T10:14:33Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
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
	// The role of the user.
	//
	// example:
	//
	// VIEWER
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s DescribeLangfuseOrgMembershipsResponseBodyDataMemberships) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseOrgMembershipsResponseBodyDataMemberships) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseOrgMembershipsResponseBodyDataMemberships) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *DescribeLangfuseOrgMembershipsResponseBodyDataMemberships) GetEmail() *string {
	return s.Email
}

func (s *DescribeLangfuseOrgMembershipsResponseBodyDataMemberships) GetName() *string {
	return s.Name
}

func (s *DescribeLangfuseOrgMembershipsResponseBodyDataMemberships) GetRole() *string {
	return s.Role
}

func (s *DescribeLangfuseOrgMembershipsResponseBodyDataMemberships) SetCreatedAt(v string) *DescribeLangfuseOrgMembershipsResponseBodyDataMemberships {
	s.CreatedAt = &v
	return s
}

func (s *DescribeLangfuseOrgMembershipsResponseBodyDataMemberships) SetEmail(v string) *DescribeLangfuseOrgMembershipsResponseBodyDataMemberships {
	s.Email = &v
	return s
}

func (s *DescribeLangfuseOrgMembershipsResponseBodyDataMemberships) SetName(v string) *DescribeLangfuseOrgMembershipsResponseBodyDataMemberships {
	s.Name = &v
	return s
}

func (s *DescribeLangfuseOrgMembershipsResponseBodyDataMemberships) SetRole(v string) *DescribeLangfuseOrgMembershipsResponseBodyDataMemberships {
	s.Role = &v
	return s
}

func (s *DescribeLangfuseOrgMembershipsResponseBodyDataMemberships) Validate() error {
	return dara.Validate(s)
}
