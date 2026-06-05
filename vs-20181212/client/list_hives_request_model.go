// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHivesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListHivesRequest
	GetEndTime() *string
	SetHiveId(v string) *ListHivesRequest
	GetHiveId() *string
	SetName(v string) *ListHivesRequest
	GetName() *string
	SetPageNumber(v int32) *ListHivesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHivesRequest
	GetPageSize() *int32
	SetStartTime(v string) *ListHivesRequest
	GetStartTime() *string
}

type ListHivesRequest struct {
	// example:
	//
	// 2025-05-14T15:20:37+08:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// g-xxxx
	HiveId *string `json:"HiveId,omitempty" xml:"HiveId,omitempty"`
	// example:
	//
	// test001
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2025-05-14T15:20:37+08:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListHivesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHivesRequest) GoString() string {
	return s.String()
}

func (s *ListHivesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListHivesRequest) GetHiveId() *string {
	return s.HiveId
}

func (s *ListHivesRequest) GetName() *string {
	return s.Name
}

func (s *ListHivesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHivesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHivesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListHivesRequest) SetEndTime(v string) *ListHivesRequest {
	s.EndTime = &v
	return s
}

func (s *ListHivesRequest) SetHiveId(v string) *ListHivesRequest {
	s.HiveId = &v
	return s
}

func (s *ListHivesRequest) SetName(v string) *ListHivesRequest {
	s.Name = &v
	return s
}

func (s *ListHivesRequest) SetPageNumber(v int32) *ListHivesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHivesRequest) SetPageSize(v int32) *ListHivesRequest {
	s.PageSize = &v
	return s
}

func (s *ListHivesRequest) SetStartTime(v string) *ListHivesRequest {
	s.StartTime = &v
	return s
}

func (s *ListHivesRequest) Validate() error {
	return dara.Validate(s)
}
