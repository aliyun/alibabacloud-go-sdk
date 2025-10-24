// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCustomGroups(v *ListCustomGroupsResponseBodyCustomGroups) *ListCustomGroupsResponseBody
	GetCustomGroups() *ListCustomGroupsResponseBodyCustomGroups
	SetPageNumber(v int32) *ListCustomGroupsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomGroupsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListCustomGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListCustomGroupsResponseBody
	GetTotalCount() *int64
}

type ListCustomGroupsResponseBody struct {
	CustomGroups *ListCustomGroupsResponseBodyCustomGroups `json:"CustomGroups,omitempty" xml:"CustomGroups,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 580e8ce3-3b80-44c5-9f3f-36ac3cc5bdd5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCustomGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomGroupsResponseBody) GetCustomGroups() *ListCustomGroupsResponseBodyCustomGroups {
	return s.CustomGroups
}

func (s *ListCustomGroupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomGroupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomGroupsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCustomGroupsResponseBody) SetCustomGroups(v *ListCustomGroupsResponseBodyCustomGroups) *ListCustomGroupsResponseBody {
	s.CustomGroups = v
	return s
}

func (s *ListCustomGroupsResponseBody) SetPageNumber(v int32) *ListCustomGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCustomGroupsResponseBody) SetPageSize(v int32) *ListCustomGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCustomGroupsResponseBody) SetRequestId(v string) *ListCustomGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomGroupsResponseBody) SetTotalCount(v int64) *ListCustomGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCustomGroupsResponseBody) Validate() error {
	if s.CustomGroups != nil {
		if err := s.CustomGroups.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCustomGroupsResponseBodyCustomGroups struct {
	CustomGroup []*ListCustomGroupsResponseBodyCustomGroupsCustomGroup `json:"CustomGroup,omitempty" xml:"CustomGroup,omitempty" type:"Repeated"`
}

func (s ListCustomGroupsResponseBodyCustomGroups) String() string {
	return dara.Prettify(s)
}

func (s ListCustomGroupsResponseBodyCustomGroups) GoString() string {
	return s.String()
}

func (s *ListCustomGroupsResponseBodyCustomGroups) GetCustomGroup() []*ListCustomGroupsResponseBodyCustomGroupsCustomGroup {
	return s.CustomGroup
}

func (s *ListCustomGroupsResponseBodyCustomGroups) SetCustomGroup(v []*ListCustomGroupsResponseBodyCustomGroupsCustomGroup) *ListCustomGroupsResponseBodyCustomGroups {
	s.CustomGroup = v
	return s
}

func (s *ListCustomGroupsResponseBodyCustomGroups) Validate() error {
	if s.CustomGroup != nil {
		for _, item := range s.CustomGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCustomGroupsResponseBodyCustomGroupsCustomGroup struct {
	CustomGroupDescription *string `json:"CustomGroupDescription,omitempty" xml:"CustomGroupDescription,omitempty"`
	// example:
	//
	// 1
	CustomGroupId   *string `json:"CustomGroupId,omitempty" xml:"CustomGroupId,omitempty"`
	CustomGroupName *string `json:"CustomGroupName,omitempty" xml:"CustomGroupName,omitempty"`
}

func (s ListCustomGroupsResponseBodyCustomGroupsCustomGroup) String() string {
	return dara.Prettify(s)
}

func (s ListCustomGroupsResponseBodyCustomGroupsCustomGroup) GoString() string {
	return s.String()
}

func (s *ListCustomGroupsResponseBodyCustomGroupsCustomGroup) GetCustomGroupDescription() *string {
	return s.CustomGroupDescription
}

func (s *ListCustomGroupsResponseBodyCustomGroupsCustomGroup) GetCustomGroupId() *string {
	return s.CustomGroupId
}

func (s *ListCustomGroupsResponseBodyCustomGroupsCustomGroup) GetCustomGroupName() *string {
	return s.CustomGroupName
}

func (s *ListCustomGroupsResponseBodyCustomGroupsCustomGroup) SetCustomGroupDescription(v string) *ListCustomGroupsResponseBodyCustomGroupsCustomGroup {
	s.CustomGroupDescription = &v
	return s
}

func (s *ListCustomGroupsResponseBodyCustomGroupsCustomGroup) SetCustomGroupId(v string) *ListCustomGroupsResponseBodyCustomGroupsCustomGroup {
	s.CustomGroupId = &v
	return s
}

func (s *ListCustomGroupsResponseBodyCustomGroupsCustomGroup) SetCustomGroupName(v string) *ListCustomGroupsResponseBodyCustomGroupsCustomGroup {
	s.CustomGroupName = &v
	return s
}

func (s *ListCustomGroupsResponseBodyCustomGroupsCustomGroup) Validate() error {
	return dara.Validate(s)
}
