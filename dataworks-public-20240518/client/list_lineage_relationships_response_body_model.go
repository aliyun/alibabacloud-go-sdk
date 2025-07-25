// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLineageRelationshipsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListLineageRelationshipsResponseBodyPagingInfo) *ListLineageRelationshipsResponseBody
	GetPagingInfo() *ListLineageRelationshipsResponseBodyPagingInfo
	SetRequestId(v string) *ListLineageRelationshipsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListLineageRelationshipsResponseBody
	GetSuccess() *bool
}

type ListLineageRelationshipsResponseBody struct {
	PagingInfo *ListLineageRelationshipsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// SDFSDFSDF-SDFSDF-SDFDSF-SDFSDF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListLineageRelationshipsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLineageRelationshipsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLineageRelationshipsResponseBody) GetPagingInfo() *ListLineageRelationshipsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListLineageRelationshipsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLineageRelationshipsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListLineageRelationshipsResponseBody) SetPagingInfo(v *ListLineageRelationshipsResponseBodyPagingInfo) *ListLineageRelationshipsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListLineageRelationshipsResponseBody) SetRequestId(v string) *ListLineageRelationshipsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLineageRelationshipsResponseBody) SetSuccess(v bool) *ListLineageRelationshipsResponseBody {
	s.Success = &v
	return s
}

func (s *ListLineageRelationshipsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLineageRelationshipsResponseBodyPagingInfo struct {
	LineageRelationships []*LineageRelationship `json:"LineageRelationships,omitempty" xml:"LineageRelationships,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 123
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLineageRelationshipsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListLineageRelationshipsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListLineageRelationshipsResponseBodyPagingInfo) GetLineageRelationships() []*LineageRelationship {
	return s.LineageRelationships
}

func (s *ListLineageRelationshipsResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLineageRelationshipsResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLineageRelationshipsResponseBodyPagingInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListLineageRelationshipsResponseBodyPagingInfo) SetLineageRelationships(v []*LineageRelationship) *ListLineageRelationshipsResponseBodyPagingInfo {
	s.LineageRelationships = v
	return s
}

func (s *ListLineageRelationshipsResponseBodyPagingInfo) SetPageNumber(v int32) *ListLineageRelationshipsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListLineageRelationshipsResponseBodyPagingInfo) SetPageSize(v int32) *ListLineageRelationshipsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListLineageRelationshipsResponseBodyPagingInfo) SetTotalCount(v int64) *ListLineageRelationshipsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListLineageRelationshipsResponseBodyPagingInfo) Validate() error {
	return dara.Validate(s)
}
