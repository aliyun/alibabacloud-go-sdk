// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCollaborationGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCollaborationGroupsResponseBody
	GetCode() *string
	SetItems(v []*ListCollaborationGroupsResponseBodyItems) *ListCollaborationGroupsResponseBody
	GetItems() []*ListCollaborationGroupsResponseBodyItems
	SetMessage(v string) *ListCollaborationGroupsResponseBody
	GetMessage() *string
	SetPage(v int64) *ListCollaborationGroupsResponseBody
	GetPage() *int64
	SetPageSize(v int64) *ListCollaborationGroupsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListCollaborationGroupsResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListCollaborationGroupsResponseBody
	GetTotal() *int64
}

type ListCollaborationGroupsResponseBody struct {
	// The common business status code.
	//
	// example:
	//
	// ERR.Robject.Auth.TenantAccessDenied
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The list of items on the current page.
	//
	// example:
	//
	// [{"collaborationGroupId":"cg_9401","name":"Product Collaboration Space","status":"ACTIVE","currentUserRole":"OWNER"}]
	Items []*ListCollaborationGroupsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The common error message.
	//
	// example:
	//
	// You do not have permissions to access tenant 2147483647 because the user does not belong to the tenant
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The page size.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// c9cf6cb3-3f5c-4994-a8a1-c93379a6f5cf
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of filtered records.
	//
	// example:
	//
	// 39
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListCollaborationGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCollaborationGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCollaborationGroupsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCollaborationGroupsResponseBody) GetItems() []*ListCollaborationGroupsResponseBodyItems {
	return s.Items
}

func (s *ListCollaborationGroupsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCollaborationGroupsResponseBody) GetPage() *int64 {
	return s.Page
}

func (s *ListCollaborationGroupsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCollaborationGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCollaborationGroupsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListCollaborationGroupsResponseBody) SetCode(v string) *ListCollaborationGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListCollaborationGroupsResponseBody) SetItems(v []*ListCollaborationGroupsResponseBodyItems) *ListCollaborationGroupsResponseBody {
	s.Items = v
	return s
}

func (s *ListCollaborationGroupsResponseBody) SetMessage(v string) *ListCollaborationGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCollaborationGroupsResponseBody) SetPage(v int64) *ListCollaborationGroupsResponseBody {
	s.Page = &v
	return s
}

func (s *ListCollaborationGroupsResponseBody) SetPageSize(v int64) *ListCollaborationGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCollaborationGroupsResponseBody) SetRequestId(v string) *ListCollaborationGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCollaborationGroupsResponseBody) SetTotal(v int64) *ListCollaborationGroupsResponseBody {
	s.Total = &v
	return s
}

func (s *ListCollaborationGroupsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCollaborationGroupsResponseBodyItems struct {
	// The ID of the collaboration space.
	//
	// example:
	//
	// cg_9401
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The role of the current user in the space.
	//
	// example:
	//
	// ADMIN
	CurrentUserRole *string `json:"currentUserRole,omitempty" xml:"currentUserRole,omitempty"`
	// The name of the space.
	//
	// example:
	//
	// Workflow To-Do Test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the space. Currently, only ACTIVE is returned.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListCollaborationGroupsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListCollaborationGroupsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListCollaborationGroupsResponseBodyItems) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListCollaborationGroupsResponseBodyItems) GetCurrentUserRole() *string {
	return s.CurrentUserRole
}

func (s *ListCollaborationGroupsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListCollaborationGroupsResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListCollaborationGroupsResponseBodyItems) SetCollaborationGroupId(v string) *ListCollaborationGroupsResponseBodyItems {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListCollaborationGroupsResponseBodyItems) SetCurrentUserRole(v string) *ListCollaborationGroupsResponseBodyItems {
	s.CurrentUserRole = &v
	return s
}

func (s *ListCollaborationGroupsResponseBodyItems) SetName(v string) *ListCollaborationGroupsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListCollaborationGroupsResponseBodyItems) SetStatus(v string) *ListCollaborationGroupsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListCollaborationGroupsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
