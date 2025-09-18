// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStreamTagListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetStreamTagListRequest
	GetEndTime() *string
	SetMediaId(v string) *GetStreamTagListRequest
	GetMediaId() *string
	SetNamespace(v string) *GetStreamTagListRequest
	GetNamespace() *string
	SetNextToken(v string) *GetStreamTagListRequest
	GetNextToken() *string
	SetPageNo(v int32) *GetStreamTagListRequest
	GetPageNo() *int32
	SetPageSize(v int32) *GetStreamTagListRequest
	GetPageSize() *int32
	SetSearchLibName(v string) *GetStreamTagListRequest
	GetSearchLibName() *string
	SetSortBy(v string) *GetStreamTagListRequest
	GetSortBy() *string
	SetStartTime(v string) *GetStreamTagListRequest
	GetStartTime() *string
}

type GetStreamTagListRequest struct {
	// example:
	//
	// 2017-12-22T08:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// ******b48fb04483915d4f2cd8******
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// name-1
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// ****73f33c91-d59383e8280b****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// Stream_xxx
	SearchLibName *string `json:"SearchLibName,omitempty" xml:"SearchLibName,omitempty"`
	// example:
	//
	// StartTime:Asc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// 2025-04-23T02:26:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetStreamTagListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStreamTagListRequest) GoString() string {
	return s.String()
}

func (s *GetStreamTagListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetStreamTagListRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *GetStreamTagListRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetStreamTagListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetStreamTagListRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetStreamTagListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetStreamTagListRequest) GetSearchLibName() *string {
	return s.SearchLibName
}

func (s *GetStreamTagListRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *GetStreamTagListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetStreamTagListRequest) SetEndTime(v string) *GetStreamTagListRequest {
	s.EndTime = &v
	return s
}

func (s *GetStreamTagListRequest) SetMediaId(v string) *GetStreamTagListRequest {
	s.MediaId = &v
	return s
}

func (s *GetStreamTagListRequest) SetNamespace(v string) *GetStreamTagListRequest {
	s.Namespace = &v
	return s
}

func (s *GetStreamTagListRequest) SetNextToken(v string) *GetStreamTagListRequest {
	s.NextToken = &v
	return s
}

func (s *GetStreamTagListRequest) SetPageNo(v int32) *GetStreamTagListRequest {
	s.PageNo = &v
	return s
}

func (s *GetStreamTagListRequest) SetPageSize(v int32) *GetStreamTagListRequest {
	s.PageSize = &v
	return s
}

func (s *GetStreamTagListRequest) SetSearchLibName(v string) *GetStreamTagListRequest {
	s.SearchLibName = &v
	return s
}

func (s *GetStreamTagListRequest) SetSortBy(v string) *GetStreamTagListRequest {
	s.SortBy = &v
	return s
}

func (s *GetStreamTagListRequest) SetStartTime(v string) *GetStreamTagListRequest {
	s.StartTime = &v
	return s
}

func (s *GetStreamTagListRequest) Validate() error {
	return dara.Validate(s)
}
