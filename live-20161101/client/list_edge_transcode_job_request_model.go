// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeTranscodeJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListEdgeTranscodeJobRequest
	GetClusterId() *string
	SetKeyword(v string) *ListEdgeTranscodeJobRequest
	GetKeyword() *string
	SetOwnerId(v int64) *ListEdgeTranscodeJobRequest
	GetOwnerId() *int64
	SetPageNo(v int32) *ListEdgeTranscodeJobRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListEdgeTranscodeJobRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListEdgeTranscodeJobRequest
	GetRegionId() *string
	SetSortBy(v string) *ListEdgeTranscodeJobRequest
	GetSortBy() *string
	SetStatus(v int32) *ListEdgeTranscodeJobRequest
	GetStatus() *int32
	SetType(v string) *ListEdgeTranscodeJobRequest
	GetType() *string
}

type ListEdgeTranscodeJobRequest struct {
	// The data center ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cluster-1
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The search keyword. Valid values:
	//
	// - Task ID. Exact match is supported.
	//
	// - Task name. Fuzzy match is supported.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The collation based on CreateTime. Default value: desc. Valid values:
	//
	// - desc: descending sorting.
	//
	// - asc: ascending sorting.
	//
	// example:
	//
	// desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The edge transcoding task status. Valid values:
	//
	// - 0: not started.
	//
	// - 1: running.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The edge transcoding type. Valid values:
	//
	// - common: default transcoding (standard + Narrowband HD 1.0).
	//
	// - nbhd-2: Narrowband HD 2.0.
	//
	// - ultra-hd: ultra-high definition.
	//
	// > If this parameter is not specified, transcoding templates for which the user has the corresponding transcoding type permissions are displayed.
	//
	// example:
	//
	// common
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListEdgeTranscodeJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeTranscodeJobRequest) GoString() string {
	return s.String()
}

func (s *ListEdgeTranscodeJobRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListEdgeTranscodeJobRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListEdgeTranscodeJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListEdgeTranscodeJobRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListEdgeTranscodeJobRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEdgeTranscodeJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEdgeTranscodeJobRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListEdgeTranscodeJobRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListEdgeTranscodeJobRequest) GetType() *string {
	return s.Type
}

func (s *ListEdgeTranscodeJobRequest) SetClusterId(v string) *ListEdgeTranscodeJobRequest {
	s.ClusterId = &v
	return s
}

func (s *ListEdgeTranscodeJobRequest) SetKeyword(v string) *ListEdgeTranscodeJobRequest {
	s.Keyword = &v
	return s
}

func (s *ListEdgeTranscodeJobRequest) SetOwnerId(v int64) *ListEdgeTranscodeJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ListEdgeTranscodeJobRequest) SetPageNo(v int32) *ListEdgeTranscodeJobRequest {
	s.PageNo = &v
	return s
}

func (s *ListEdgeTranscodeJobRequest) SetPageSize(v int32) *ListEdgeTranscodeJobRequest {
	s.PageSize = &v
	return s
}

func (s *ListEdgeTranscodeJobRequest) SetRegionId(v string) *ListEdgeTranscodeJobRequest {
	s.RegionId = &v
	return s
}

func (s *ListEdgeTranscodeJobRequest) SetSortBy(v string) *ListEdgeTranscodeJobRequest {
	s.SortBy = &v
	return s
}

func (s *ListEdgeTranscodeJobRequest) SetStatus(v int32) *ListEdgeTranscodeJobRequest {
	s.Status = &v
	return s
}

func (s *ListEdgeTranscodeJobRequest) SetType(v string) *ListEdgeTranscodeJobRequest {
	s.Type = &v
	return s
}

func (s *ListEdgeTranscodeJobRequest) Validate() error {
	return dara.Validate(s)
}
