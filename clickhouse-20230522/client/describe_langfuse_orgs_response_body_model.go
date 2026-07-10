// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLangfuseOrgsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeLangfuseOrgsResponseBodyData) *DescribeLangfuseOrgsResponseBody
	GetData() *DescribeLangfuseOrgsResponseBodyData
	SetRequestId(v string) *DescribeLangfuseOrgsResponseBody
	GetRequestId() *string
}

type DescribeLangfuseOrgsResponseBody struct {
	Data *DescribeLangfuseOrgsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD8660E9CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLangfuseOrgsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseOrgsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseOrgsResponseBody) GetData() *DescribeLangfuseOrgsResponseBodyData {
	return s.Data
}

func (s *DescribeLangfuseOrgsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLangfuseOrgsResponseBody) SetData(v *DescribeLangfuseOrgsResponseBodyData) *DescribeLangfuseOrgsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeLangfuseOrgsResponseBody) SetRequestId(v string) *DescribeLangfuseOrgsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLangfuseOrgsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLangfuseOrgsResponseBodyData struct {
	Organizations []*DescribeLangfuseOrgsResponseBodyDataOrganizations `json:"Organizations,omitempty" xml:"Organizations,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 17
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLangfuseOrgsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseOrgsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseOrgsResponseBodyData) GetOrganizations() []*DescribeLangfuseOrgsResponseBodyDataOrganizations {
	return s.Organizations
}

func (s *DescribeLangfuseOrgsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLangfuseOrgsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLangfuseOrgsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeLangfuseOrgsResponseBodyData) SetOrganizations(v []*DescribeLangfuseOrgsResponseBodyDataOrganizations) *DescribeLangfuseOrgsResponseBodyData {
	s.Organizations = v
	return s
}

func (s *DescribeLangfuseOrgsResponseBodyData) SetPageNumber(v int32) *DescribeLangfuseOrgsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *DescribeLangfuseOrgsResponseBodyData) SetPageSize(v int32) *DescribeLangfuseOrgsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeLangfuseOrgsResponseBodyData) SetTotalCount(v int64) *DescribeLangfuseOrgsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeLangfuseOrgsResponseBodyData) Validate() error {
	if s.Organizations != nil {
		for _, item := range s.Organizations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLangfuseOrgsResponseBodyDataOrganizations struct {
	// example:
	//
	// 2026-06-11T10:27:23Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// org_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// cmrbhzx930005jw2q****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// example:
	//
	// 2026-06-09T10:27:55
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s DescribeLangfuseOrgsResponseBodyDataOrganizations) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseOrgsResponseBodyDataOrganizations) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseOrgsResponseBodyDataOrganizations) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *DescribeLangfuseOrgsResponseBodyDataOrganizations) GetName() *string {
	return s.Name
}

func (s *DescribeLangfuseOrgsResponseBodyDataOrganizations) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DescribeLangfuseOrgsResponseBodyDataOrganizations) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *DescribeLangfuseOrgsResponseBodyDataOrganizations) SetCreatedAt(v string) *DescribeLangfuseOrgsResponseBodyDataOrganizations {
	s.CreatedAt = &v
	return s
}

func (s *DescribeLangfuseOrgsResponseBodyDataOrganizations) SetName(v string) *DescribeLangfuseOrgsResponseBodyDataOrganizations {
	s.Name = &v
	return s
}

func (s *DescribeLangfuseOrgsResponseBodyDataOrganizations) SetOrganizationId(v string) *DescribeLangfuseOrgsResponseBodyDataOrganizations {
	s.OrganizationId = &v
	return s
}

func (s *DescribeLangfuseOrgsResponseBodyDataOrganizations) SetUpdatedAt(v string) *DescribeLangfuseOrgsResponseBodyDataOrganizations {
	s.UpdatedAt = &v
	return s
}

func (s *DescribeLangfuseOrgsResponseBodyDataOrganizations) Validate() error {
	return dara.Validate(s)
}
