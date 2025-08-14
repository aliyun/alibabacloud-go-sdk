// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChannelAlertsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ListChannelAlertsRequest
	GetCategory() *string
	SetGmtEnd(v string) *ListChannelAlertsRequest
	GetGmtEnd() *string
	SetGmtStart(v string) *ListChannelAlertsRequest
	GetGmtStart() *string
	SetPageNo(v int32) *ListChannelAlertsRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListChannelAlertsRequest
	GetPageSize() *int32
	SetResourceArn(v string) *ListChannelAlertsRequest
	GetResourceArn() *string
	SetSortByModifiedTime(v string) *ListChannelAlertsRequest
	GetSortByModifiedTime() *string
}

type ListChannelAlertsRequest struct {
	// The alert type.
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The end of the time range to query.
	//
	// example:
	//
	// 2024-11-21T16:10:45Z
	GmtEnd *string `json:"GmtEnd,omitempty" xml:"GmtEnd,omitempty"`
	// The beginning of the time range to query.
	//
	// example:
	//
	// 2024-11-21T16:10:45Z
	GmtStart *string `json:"GmtStart,omitempty" xml:"GmtStart,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ARN of the channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:ims:mediaweaver:<regionId>:<userId>:channel/myChannel
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	// The sorting order by modification time. Valid values: asc and desc.
	//
	// example:
	//
	// desc
	SortByModifiedTime *string `json:"SortByModifiedTime,omitempty" xml:"SortByModifiedTime,omitempty"`
}

func (s ListChannelAlertsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChannelAlertsRequest) GoString() string {
	return s.String()
}

func (s *ListChannelAlertsRequest) GetCategory() *string {
	return s.Category
}

func (s *ListChannelAlertsRequest) GetGmtEnd() *string {
	return s.GmtEnd
}

func (s *ListChannelAlertsRequest) GetGmtStart() *string {
	return s.GmtStart
}

func (s *ListChannelAlertsRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListChannelAlertsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListChannelAlertsRequest) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *ListChannelAlertsRequest) GetSortByModifiedTime() *string {
	return s.SortByModifiedTime
}

func (s *ListChannelAlertsRequest) SetCategory(v string) *ListChannelAlertsRequest {
	s.Category = &v
	return s
}

func (s *ListChannelAlertsRequest) SetGmtEnd(v string) *ListChannelAlertsRequest {
	s.GmtEnd = &v
	return s
}

func (s *ListChannelAlertsRequest) SetGmtStart(v string) *ListChannelAlertsRequest {
	s.GmtStart = &v
	return s
}

func (s *ListChannelAlertsRequest) SetPageNo(v int32) *ListChannelAlertsRequest {
	s.PageNo = &v
	return s
}

func (s *ListChannelAlertsRequest) SetPageSize(v int32) *ListChannelAlertsRequest {
	s.PageSize = &v
	return s
}

func (s *ListChannelAlertsRequest) SetResourceArn(v string) *ListChannelAlertsRequest {
	s.ResourceArn = &v
	return s
}

func (s *ListChannelAlertsRequest) SetSortByModifiedTime(v string) *ListChannelAlertsRequest {
	s.SortByModifiedTime = &v
	return s
}

func (s *ListChannelAlertsRequest) Validate() error {
	return dara.Validate(s)
}
