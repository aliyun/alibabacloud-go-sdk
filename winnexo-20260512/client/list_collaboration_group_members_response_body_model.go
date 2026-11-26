// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCollaborationGroupMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCollaborationGroupMembersResponseBody
	GetCode() *string
	SetItems(v []*ListCollaborationGroupMembersResponseBodyItems) *ListCollaborationGroupMembersResponseBody
	GetItems() []*ListCollaborationGroupMembersResponseBodyItems
	SetMessage(v string) *ListCollaborationGroupMembersResponseBody
	GetMessage() *string
	SetPage(v int64) *ListCollaborationGroupMembersResponseBody
	GetPage() *int64
	SetPageSize(v int64) *ListCollaborationGroupMembersResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListCollaborationGroupMembersResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListCollaborationGroupMembersResponseBody
	GetTotal() *int64
}

type ListCollaborationGroupMembersResponseBody struct {
	// The common business status code.
	//
	// example:
	//
	// ERR.Robject.Auth.TenantAccessDenied
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The list of entries on the current page.
	//
	// example:
	//
	// [{"userId":"5","displayName":"Zhang San","role":"OWNER","isActive":true}]
	Items []*ListCollaborationGroupMembersResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The common error description.
	//
	// example:
	//
	// No permission to access tenant 2147483647. The user does not belong to this tenant
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
	// The request tracking ID.
	//
	// example:
	//
	// c9cf6cb3-3f5c-4994-a8a1-c93379a6f5cf
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of records after filtering.
	//
	// example:
	//
	// 10
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListCollaborationGroupMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCollaborationGroupMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListCollaborationGroupMembersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCollaborationGroupMembersResponseBody) GetItems() []*ListCollaborationGroupMembersResponseBodyItems {
	return s.Items
}

func (s *ListCollaborationGroupMembersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCollaborationGroupMembersResponseBody) GetPage() *int64 {
	return s.Page
}

func (s *ListCollaborationGroupMembersResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCollaborationGroupMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCollaborationGroupMembersResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListCollaborationGroupMembersResponseBody) SetCode(v string) *ListCollaborationGroupMembersResponseBody {
	s.Code = &v
	return s
}

func (s *ListCollaborationGroupMembersResponseBody) SetItems(v []*ListCollaborationGroupMembersResponseBodyItems) *ListCollaborationGroupMembersResponseBody {
	s.Items = v
	return s
}

func (s *ListCollaborationGroupMembersResponseBody) SetMessage(v string) *ListCollaborationGroupMembersResponseBody {
	s.Message = &v
	return s
}

func (s *ListCollaborationGroupMembersResponseBody) SetPage(v int64) *ListCollaborationGroupMembersResponseBody {
	s.Page = &v
	return s
}

func (s *ListCollaborationGroupMembersResponseBody) SetPageSize(v int64) *ListCollaborationGroupMembersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCollaborationGroupMembersResponseBody) SetRequestId(v string) *ListCollaborationGroupMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCollaborationGroupMembersResponseBody) SetTotal(v int64) *ListCollaborationGroupMembersResponseBody {
	s.Total = &v
	return s
}

func (s *ListCollaborationGroupMembersResponseBody) Validate() error {
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

type ListCollaborationGroupMembersResponseBodyItems struct {
	// The display name. This field is empty if not set.
	//
	// example:
	//
	// TEST001
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The enabling status of the tenant account. Only active members are returned, so this value is always true. It does not represent the employment status.
	//
	// example:
	//
	// true
	IsActive *bool `json:"isActive,omitempty" xml:"isActive,omitempty"`
	// The role of the workspace member.
	//
	// example:
	//
	// ADMIN
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// The platform user ID.
	//
	// example:
	//
	// 5
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListCollaborationGroupMembersResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListCollaborationGroupMembersResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListCollaborationGroupMembersResponseBodyItems) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListCollaborationGroupMembersResponseBodyItems) GetIsActive() *bool {
	return s.IsActive
}

func (s *ListCollaborationGroupMembersResponseBodyItems) GetRole() *string {
	return s.Role
}

func (s *ListCollaborationGroupMembersResponseBodyItems) GetUserId() *string {
	return s.UserId
}

func (s *ListCollaborationGroupMembersResponseBodyItems) SetDisplayName(v string) *ListCollaborationGroupMembersResponseBodyItems {
	s.DisplayName = &v
	return s
}

func (s *ListCollaborationGroupMembersResponseBodyItems) SetIsActive(v bool) *ListCollaborationGroupMembersResponseBodyItems {
	s.IsActive = &v
	return s
}

func (s *ListCollaborationGroupMembersResponseBodyItems) SetRole(v string) *ListCollaborationGroupMembersResponseBodyItems {
	s.Role = &v
	return s
}

func (s *ListCollaborationGroupMembersResponseBodyItems) SetUserId(v string) *ListCollaborationGroupMembersResponseBodyItems {
	s.UserId = &v
	return s
}

func (s *ListCollaborationGroupMembersResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
