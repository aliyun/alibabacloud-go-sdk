// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchEditingProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *SearchEditingProjectRequest
	GetEndTime() *string
	SetOwnerAccount(v string) *SearchEditingProjectRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *SearchEditingProjectRequest
	GetOwnerId() *string
	SetPageNo(v int32) *SearchEditingProjectRequest
	GetPageNo() *int32
	SetPageSize(v int32) *SearchEditingProjectRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *SearchEditingProjectRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *SearchEditingProjectRequest
	GetResourceOwnerId() *string
	SetSortBy(v string) *SearchEditingProjectRequest
	GetSortBy() *string
	SetStartTime(v string) *SearchEditingProjectRequest
	GetStartTime() *string
	SetStatus(v string) *SearchEditingProjectRequest
	GetStatus() *string
	SetTitle(v string) *SearchEditingProjectRequest
	GetTitle() *string
}

type SearchEditingProjectRequest struct {
	// The end of the time range to query. The query is performed based on the time range during which the required online editing projects were created. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2017-01-11T13:00:00Z
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page. Default value: **10**. Maximum value: **100**.
	//
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The sorting rule of results. Valid values:
	//
	// 	- **CreationTime:Desc**: sorts the results based on the creation time in descending order. This is the default value.
	//
	// 	- **CreationTime:Asc**: sorts the results based on the creation time in ascending order.
	//
	// example:
	//
	// CreationTime:Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The beginning of the time range to query. The query is performed based on the time range during which the required online editing projects were created. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the online editing project. Separate multiple states with commas (,). By default, all online editing projects are queried. Valid values:
	//
	// 	- **Normal**: indicates that the online editing project is in draft.
	//
	// 	- **Producing**: indicates that the video is being produced.
	//
	// 	- **Produced**: indicates that the video was produced.
	//
	// 	- **ProduceFailed**: indicates that the video failed to be produced.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The title of the online editing project.
	//
	// example:
	//
	// test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s SearchEditingProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchEditingProjectRequest) GoString() string {
	return s.String()
}

func (s *SearchEditingProjectRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *SearchEditingProjectRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SearchEditingProjectRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SearchEditingProjectRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *SearchEditingProjectRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchEditingProjectRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SearchEditingProjectRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *SearchEditingProjectRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *SearchEditingProjectRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *SearchEditingProjectRequest) GetStatus() *string {
	return s.Status
}

func (s *SearchEditingProjectRequest) GetTitle() *string {
	return s.Title
}

func (s *SearchEditingProjectRequest) SetEndTime(v string) *SearchEditingProjectRequest {
	s.EndTime = &v
	return s
}

func (s *SearchEditingProjectRequest) SetOwnerAccount(v string) *SearchEditingProjectRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SearchEditingProjectRequest) SetOwnerId(v string) *SearchEditingProjectRequest {
	s.OwnerId = &v
	return s
}

func (s *SearchEditingProjectRequest) SetPageNo(v int32) *SearchEditingProjectRequest {
	s.PageNo = &v
	return s
}

func (s *SearchEditingProjectRequest) SetPageSize(v int32) *SearchEditingProjectRequest {
	s.PageSize = &v
	return s
}

func (s *SearchEditingProjectRequest) SetResourceOwnerAccount(v string) *SearchEditingProjectRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SearchEditingProjectRequest) SetResourceOwnerId(v string) *SearchEditingProjectRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SearchEditingProjectRequest) SetSortBy(v string) *SearchEditingProjectRequest {
	s.SortBy = &v
	return s
}

func (s *SearchEditingProjectRequest) SetStartTime(v string) *SearchEditingProjectRequest {
	s.StartTime = &v
	return s
}

func (s *SearchEditingProjectRequest) SetStatus(v string) *SearchEditingProjectRequest {
	s.Status = &v
	return s
}

func (s *SearchEditingProjectRequest) SetTitle(v string) *SearchEditingProjectRequest {
	s.Title = &v
	return s
}

func (s *SearchEditingProjectRequest) Validate() error {
	return dara.Validate(s)
}
