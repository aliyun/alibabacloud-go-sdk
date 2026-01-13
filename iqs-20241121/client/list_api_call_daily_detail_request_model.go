// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiCallDailyDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiName(v string) *ListApiCallDailyDetailRequest
	GetApiName() *string
	SetEndTime(v string) *ListApiCallDailyDetailRequest
	GetEndTime() *string
	SetEngineType(v string) *ListApiCallDailyDetailRequest
	GetEngineType() *string
	SetPageNumber(v int32) *ListApiCallDailyDetailRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListApiCallDailyDetailRequest
	GetPageSize() *int32
	SetStartTime(v string) *ListApiCallDailyDetailRequest
	GetStartTime() *string
}

type ListApiCallDailyDetailRequest struct {
	// example:
	//
	// CreateVikaApp
	ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty"`
	// example:
	//
	// 20240908
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// PlaceBasic
	EngineType *string `json:"engineType,omitempty" xml:"engineType,omitempty"`
	// example:
	//
	// 0
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 20240908
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListApiCallDailyDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApiCallDailyDetailRequest) GoString() string {
	return s.String()
}

func (s *ListApiCallDailyDetailRequest) GetApiName() *string {
	return s.ApiName
}

func (s *ListApiCallDailyDetailRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListApiCallDailyDetailRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *ListApiCallDailyDetailRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListApiCallDailyDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApiCallDailyDetailRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListApiCallDailyDetailRequest) SetApiName(v string) *ListApiCallDailyDetailRequest {
	s.ApiName = &v
	return s
}

func (s *ListApiCallDailyDetailRequest) SetEndTime(v string) *ListApiCallDailyDetailRequest {
	s.EndTime = &v
	return s
}

func (s *ListApiCallDailyDetailRequest) SetEngineType(v string) *ListApiCallDailyDetailRequest {
	s.EngineType = &v
	return s
}

func (s *ListApiCallDailyDetailRequest) SetPageNumber(v int32) *ListApiCallDailyDetailRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApiCallDailyDetailRequest) SetPageSize(v int32) *ListApiCallDailyDetailRequest {
	s.PageSize = &v
	return s
}

func (s *ListApiCallDailyDetailRequest) SetStartTime(v string) *ListApiCallDailyDetailRequest {
	s.StartTime = &v
	return s
}

func (s *ListApiCallDailyDetailRequest) Validate() error {
	return dara.Validate(s)
}
