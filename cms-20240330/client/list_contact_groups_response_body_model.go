// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContactGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroups(v []*ListContactGroupsResponseBodyContactGroups) *ListContactGroupsResponseBody
	GetContactGroups() []*ListContactGroupsResponseBodyContactGroups
	SetPageNumber(v int64) *ListContactGroupsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListContactGroupsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListContactGroupsResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListContactGroupsResponseBody
	GetTotal() *int64
}

type ListContactGroupsResponseBody struct {
	ContactGroups []*ListContactGroupsResponseBodyContactGroups `json:"contactGroups,omitempty" xml:"contactGroups,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 15
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListContactGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListContactGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListContactGroupsResponseBody) GetContactGroups() []*ListContactGroupsResponseBodyContactGroups {
	return s.ContactGroups
}

func (s *ListContactGroupsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListContactGroupsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListContactGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListContactGroupsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListContactGroupsResponseBody) SetContactGroups(v []*ListContactGroupsResponseBodyContactGroups) *ListContactGroupsResponseBody {
	s.ContactGroups = v
	return s
}

func (s *ListContactGroupsResponseBody) SetPageNumber(v int64) *ListContactGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListContactGroupsResponseBody) SetPageSize(v int64) *ListContactGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListContactGroupsResponseBody) SetRequestId(v string) *ListContactGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContactGroupsResponseBody) SetTotal(v int64) *ListContactGroupsResponseBody {
	s.Total = &v
	return s
}

func (s *ListContactGroupsResponseBody) Validate() error {
	if s.ContactGroups != nil {
		for _, item := range s.ContactGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListContactGroupsResponseBodyContactGroups struct {
	// example:
	//
	// test
	ContactGroupId *string   `json:"contactGroupId,omitempty" xml:"contactGroupId,omitempty"`
	ContactIds     []*string `json:"contactIds,omitempty" xml:"contactIds,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListContactGroupsResponseBodyContactGroups) String() string {
	return dara.Prettify(s)
}

func (s ListContactGroupsResponseBodyContactGroups) GoString() string {
	return s.String()
}

func (s *ListContactGroupsResponseBodyContactGroups) GetContactGroupId() *string {
	return s.ContactGroupId
}

func (s *ListContactGroupsResponseBodyContactGroups) GetContactIds() []*string {
	return s.ContactIds
}

func (s *ListContactGroupsResponseBodyContactGroups) GetName() *string {
	return s.Name
}

func (s *ListContactGroupsResponseBodyContactGroups) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListContactGroupsResponseBodyContactGroups) SetContactGroupId(v string) *ListContactGroupsResponseBodyContactGroups {
	s.ContactGroupId = &v
	return s
}

func (s *ListContactGroupsResponseBodyContactGroups) SetContactIds(v []*string) *ListContactGroupsResponseBodyContactGroups {
	s.ContactIds = v
	return s
}

func (s *ListContactGroupsResponseBodyContactGroups) SetName(v string) *ListContactGroupsResponseBodyContactGroups {
	s.Name = &v
	return s
}

func (s *ListContactGroupsResponseBodyContactGroups) SetWorkspace(v string) *ListContactGroupsResponseBodyContactGroups {
	s.Workspace = &v
	return s
}

func (s *ListContactGroupsResponseBodyContactGroups) Validate() error {
	return dara.Validate(s)
}
