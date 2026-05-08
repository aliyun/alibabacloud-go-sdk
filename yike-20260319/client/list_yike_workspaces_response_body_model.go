// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListYikeWorkspacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListYikeWorkspacesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListYikeWorkspacesResponseBody
	GetTotalCount() *int32
	SetWorkspaceList(v []*ListYikeWorkspacesResponseBodyWorkspaceList) *ListYikeWorkspacesResponseBody
	GetWorkspaceList() []*ListYikeWorkspacesResponseBodyWorkspaceList
}

type ListYikeWorkspacesResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 4E84BE44-58A7-****-****-FBEBEA16EF94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 54
	TotalCount    *int32                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	WorkspaceList []*ListYikeWorkspacesResponseBodyWorkspaceList `json:"WorkspaceList,omitempty" xml:"WorkspaceList,omitempty" type:"Repeated"`
}

func (s ListYikeWorkspacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListYikeWorkspacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListYikeWorkspacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListYikeWorkspacesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListYikeWorkspacesResponseBody) GetWorkspaceList() []*ListYikeWorkspacesResponseBodyWorkspaceList {
	return s.WorkspaceList
}

func (s *ListYikeWorkspacesResponseBody) SetRequestId(v string) *ListYikeWorkspacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListYikeWorkspacesResponseBody) SetTotalCount(v int32) *ListYikeWorkspacesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListYikeWorkspacesResponseBody) SetWorkspaceList(v []*ListYikeWorkspacesResponseBodyWorkspaceList) *ListYikeWorkspacesResponseBody {
	s.WorkspaceList = v
	return s
}

func (s *ListYikeWorkspacesResponseBody) Validate() error {
	if s.WorkspaceList != nil {
		for _, item := range s.WorkspaceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListYikeWorkspacesResponseBodyWorkspaceList struct {
	// example:
	//
	// ABCD12
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 2026-04-29T10:22:44Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// pd_1234***
	DefaultProductionId *string `json:"DefaultProductionId,omitempty" xml:"DefaultProductionId,omitempty"`
	// example:
	//
	// Online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// workspace title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 1
	UserCount *string `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
	// example:
	//
	// ws_1243****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListYikeWorkspacesResponseBodyWorkspaceList) String() string {
	return dara.Prettify(s)
}

func (s ListYikeWorkspacesResponseBodyWorkspaceList) GoString() string {
	return s.String()
}

func (s *ListYikeWorkspacesResponseBodyWorkspaceList) GetCode() *string {
	return s.Code
}

func (s *ListYikeWorkspacesResponseBodyWorkspaceList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListYikeWorkspacesResponseBodyWorkspaceList) GetDefaultProductionId() *string {
	return s.DefaultProductionId
}

func (s *ListYikeWorkspacesResponseBodyWorkspaceList) GetStatus() *string {
	return s.Status
}

func (s *ListYikeWorkspacesResponseBodyWorkspaceList) GetTitle() *string {
	return s.Title
}

func (s *ListYikeWorkspacesResponseBodyWorkspaceList) GetUserCount() *string {
	return s.UserCount
}

func (s *ListYikeWorkspacesResponseBodyWorkspaceList) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListYikeWorkspacesResponseBodyWorkspaceList) SetCode(v string) *ListYikeWorkspacesResponseBodyWorkspaceList {
	s.Code = &v
	return s
}

func (s *ListYikeWorkspacesResponseBodyWorkspaceList) SetCreateTime(v string) *ListYikeWorkspacesResponseBodyWorkspaceList {
	s.CreateTime = &v
	return s
}

func (s *ListYikeWorkspacesResponseBodyWorkspaceList) SetDefaultProductionId(v string) *ListYikeWorkspacesResponseBodyWorkspaceList {
	s.DefaultProductionId = &v
	return s
}

func (s *ListYikeWorkspacesResponseBodyWorkspaceList) SetStatus(v string) *ListYikeWorkspacesResponseBodyWorkspaceList {
	s.Status = &v
	return s
}

func (s *ListYikeWorkspacesResponseBodyWorkspaceList) SetTitle(v string) *ListYikeWorkspacesResponseBodyWorkspaceList {
	s.Title = &v
	return s
}

func (s *ListYikeWorkspacesResponseBodyWorkspaceList) SetUserCount(v string) *ListYikeWorkspacesResponseBodyWorkspaceList {
	s.UserCount = &v
	return s
}

func (s *ListYikeWorkspacesResponseBodyWorkspaceList) SetWorkspaceId(v string) *ListYikeWorkspacesResponseBodyWorkspaceList {
	s.WorkspaceId = &v
	return s
}

func (s *ListYikeWorkspacesResponseBodyWorkspaceList) Validate() error {
	return dara.Validate(s)
}
