// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuditHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int64) *GetAuditHistoryRequest
	GetPageNo() *int64
	SetPageSize(v int64) *GetAuditHistoryRequest
	GetPageSize() *int64
	SetSortBy(v string) *GetAuditHistoryRequest
	GetSortBy() *string
	SetVideoId(v string) *GetAuditHistoryRequest
	GetVideoId() *string
}

type GetAuditHistoryRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page. Default value: **10**. Maximum value: **100**.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sorting rule of the results. Valid values:
	//
	// 	- **CreationTime:Desc**: sorts the results based on the creation time in descending order. This is the default value.
	//
	// 	- **CreationTime:Asc**: sorts the results based on the creation time in ascending order.
	//
	// example:
	//
	// CreationTime:Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The ID of the video.
	//
	// This parameter is required.
	//
	// example:
	//
	// 93ab850b4f6f44*****6e91d24d81d4
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetAuditHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAuditHistoryRequest) GoString() string {
	return s.String()
}

func (s *GetAuditHistoryRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *GetAuditHistoryRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetAuditHistoryRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *GetAuditHistoryRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *GetAuditHistoryRequest) SetPageNo(v int64) *GetAuditHistoryRequest {
	s.PageNo = &v
	return s
}

func (s *GetAuditHistoryRequest) SetPageSize(v int64) *GetAuditHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *GetAuditHistoryRequest) SetSortBy(v string) *GetAuditHistoryRequest {
	s.SortBy = &v
	return s
}

func (s *GetAuditHistoryRequest) SetVideoId(v string) *GetAuditHistoryRequest {
	s.VideoId = &v
	return s
}

func (s *GetAuditHistoryRequest) Validate() error {
	return dara.Validate(s)
}
