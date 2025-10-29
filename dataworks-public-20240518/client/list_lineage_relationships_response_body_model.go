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
	// The pagination result.
	PagingInfo *ListLineageRelationshipsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID. Used for locating and troubleshooting issues.
	//
	// example:
	//
	// SDFSDFSDF-SDFSDF-SDFDSF-SDFSDF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
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
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLineageRelationshipsResponseBodyPagingInfo struct {
	// The list of data tables.
	LineageRelationships []*LineageRelationship `json:"LineageRelationships,omitempty" xml:"LineageRelationships,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total count.
	//
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
	if s.LineageRelationships != nil {
		for _, item := range s.LineageRelationships {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
