// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterMemberPage(v *ListClusterMembersResponseBodyClusterMemberPage) *ListClusterMembersResponseBody
	GetClusterMemberPage() *ListClusterMembersResponseBodyClusterMemberPage
	SetCode(v int32) *ListClusterMembersResponseBody
	GetCode() *int32
	SetMessage(v string) *ListClusterMembersResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListClusterMembersResponseBody
	GetRequestId() *string
}

type ListClusterMembersResponseBody struct {
	// The information about the ECS instances in the cluster.
	ClusterMemberPage *ListClusterMembersResponseBodyClusterMemberPage `json:"ClusterMemberPage,omitempty" xml:"ClusterMemberPage,omitempty" type:"Struct"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// b197-40ab-9155-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClusterMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClusterMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterMembersResponseBody) GetClusterMemberPage() *ListClusterMembersResponseBodyClusterMemberPage {
	return s.ClusterMemberPage
}

func (s *ListClusterMembersResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListClusterMembersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListClusterMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClusterMembersResponseBody) SetClusterMemberPage(v *ListClusterMembersResponseBodyClusterMemberPage) *ListClusterMembersResponseBody {
	s.ClusterMemberPage = v
	return s
}

func (s *ListClusterMembersResponseBody) SetCode(v int32) *ListClusterMembersResponseBody {
	s.Code = &v
	return s
}

func (s *ListClusterMembersResponseBody) SetMessage(v string) *ListClusterMembersResponseBody {
	s.Message = &v
	return s
}

func (s *ListClusterMembersResponseBody) SetRequestId(v string) *ListClusterMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterMembersResponseBody) Validate() error {
	if s.ClusterMemberPage != nil {
		if err := s.ClusterMemberPage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListClusterMembersResponseBodyClusterMemberPage struct {
	ClusterMemberList *ListClusterMembersResponseBodyClusterMemberPageClusterMemberList `json:"ClusterMemberList,omitempty" xml:"ClusterMemberList,omitempty" type:"Struct"`
	// The page number of the returned page. If this parameter is not returned, the first page is returned.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of ECS instances returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of pages returned when all ECS instances are returned based on the specified PageSize parameter.
	//
	// example:
	//
	// 5
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListClusterMembersResponseBodyClusterMemberPage) String() string {
	return dara.Prettify(s)
}

func (s ListClusterMembersResponseBodyClusterMemberPage) GoString() string {
	return s.String()
}

func (s *ListClusterMembersResponseBodyClusterMemberPage) GetClusterMemberList() *ListClusterMembersResponseBodyClusterMemberPageClusterMemberList {
	return s.ClusterMemberList
}

func (s *ListClusterMembersResponseBodyClusterMemberPage) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListClusterMembersResponseBodyClusterMemberPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListClusterMembersResponseBodyClusterMemberPage) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListClusterMembersResponseBodyClusterMemberPage) SetClusterMemberList(v *ListClusterMembersResponseBodyClusterMemberPageClusterMemberList) *ListClusterMembersResponseBodyClusterMemberPage {
	s.ClusterMemberList = v
	return s
}

func (s *ListClusterMembersResponseBodyClusterMemberPage) SetCurrentPage(v int32) *ListClusterMembersResponseBodyClusterMemberPage {
	s.CurrentPage = &v
	return s
}

func (s *ListClusterMembersResponseBodyClusterMemberPage) SetPageSize(v int32) *ListClusterMembersResponseBodyClusterMemberPage {
	s.PageSize = &v
	return s
}

func (s *ListClusterMembersResponseBodyClusterMemberPage) SetTotalSize(v int32) *ListClusterMembersResponseBodyClusterMemberPage {
	s.TotalSize = &v
	return s
}

func (s *ListClusterMembersResponseBodyClusterMemberPage) Validate() error {
	if s.ClusterMemberList != nil {
		if err := s.ClusterMemberList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListClusterMembersResponseBodyClusterMemberPageClusterMemberList struct {
	ClusterMember []*ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember `json:"ClusterMember,omitempty" xml:"ClusterMember,omitempty" type:"Repeated"`
}

func (s ListClusterMembersResponseBodyClusterMemberPageClusterMemberList) String() string {
	return dara.Prettify(s)
}

func (s ListClusterMembersResponseBodyClusterMemberPageClusterMemberList) GoString() string {
	return s.String()
}

func (s *ListClusterMembersResponseBodyClusterMemberPageClusterMemberList) GetClusterMember() []*ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember {
	return s.ClusterMember
}

func (s *ListClusterMembersResponseBodyClusterMemberPageClusterMemberList) SetClusterMember(v []*ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) *ListClusterMembersResponseBodyClusterMemberPageClusterMemberList {
	s.ClusterMember = v
	return s
}

func (s *ListClusterMembersResponseBodyClusterMemberPageClusterMemberList) Validate() error {
	if s.ClusterMember != nil {
		for _, item := range s.ClusterMember {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterMemberId *string `json:"ClusterMemberId,omitempty" xml:"ClusterMemberId,omitempty"`
	CreateTime      *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EcsId           *string `json:"EcsId,omitempty" xml:"EcsId,omitempty"`
	EcuId           *string `json:"EcuId,omitempty" xml:"EcuId,omitempty"`
	PrivateIp       *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	Status          *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime      *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) String() string {
	return dara.Prettify(s)
}

func (s ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) GoString() string {
	return s.String()
}

func (s *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) GetClusterMemberId() *string {
	return s.ClusterMemberId
}

func (s *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) GetEcsId() *string {
	return s.EcsId
}

func (s *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) GetEcuId() *string {
	return s.EcuId
}

func (s *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) GetStatus() *int32 {
	return s.Status
}

func (s *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) SetClusterId(v string) *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember {
	s.ClusterId = &v
	return s
}

func (s *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) SetClusterMemberId(v string) *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember {
	s.ClusterMemberId = &v
	return s
}

func (s *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) SetCreateTime(v int64) *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember {
	s.CreateTime = &v
	return s
}

func (s *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) SetEcsId(v string) *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember {
	s.EcsId = &v
	return s
}

func (s *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) SetEcuId(v string) *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember {
	s.EcuId = &v
	return s
}

func (s *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) SetPrivateIp(v string) *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember {
	s.PrivateIp = &v
	return s
}

func (s *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) SetStatus(v int32) *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember {
	s.Status = &v
	return s
}

func (s *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) SetUpdateTime(v int64) *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember {
	s.UpdateTime = &v
	return s
}

func (s *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) Validate() error {
	return dara.Validate(s)
}
