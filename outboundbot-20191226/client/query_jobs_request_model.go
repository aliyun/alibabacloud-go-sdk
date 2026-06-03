// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactName(v string) *QueryJobsRequest
	GetContactName() *string
	SetEndTime(v int64) *QueryJobsRequest
	GetEndTime() *int64
	SetInstanceId(v string) *QueryJobsRequest
	GetInstanceId() *string
	SetJobGroupId(v string) *QueryJobsRequest
	GetJobGroupId() *string
	SetPageNumber(v int32) *QueryJobsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryJobsRequest
	GetPageSize() *int32
	SetPhoneNumber(v string) *QueryJobsRequest
	GetPhoneNumber() *string
	SetScenarioId(v string) *QueryJobsRequest
	GetScenarioId() *string
	SetStartTime(v int64) *QueryJobsRequest
	GetStartTime() *int64
	SetTimeAlignment(v string) *QueryJobsRequest
	GetTimeAlignment() *string
}

type QueryJobsRequest struct {
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// example:
	//
	// 1579077794665
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 994b8baf-7ef8-480c-b141-b7b6db77c4df
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 135****8888
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// example:
	//
	// b0f35dd1-0337-402e-9c4f-3a6c2426950a
	ScenarioId *string `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	// example:
	//
	// 1579068424883
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// start
	TimeAlignment *string `json:"TimeAlignment,omitempty" xml:"TimeAlignment,omitempty"`
}

func (s QueryJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryJobsRequest) GoString() string {
	return s.String()
}

func (s *QueryJobsRequest) GetContactName() *string {
	return s.ContactName
}

func (s *QueryJobsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryJobsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryJobsRequest) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *QueryJobsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryJobsRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *QueryJobsRequest) GetScenarioId() *string {
	return s.ScenarioId
}

func (s *QueryJobsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryJobsRequest) GetTimeAlignment() *string {
	return s.TimeAlignment
}

func (s *QueryJobsRequest) SetContactName(v string) *QueryJobsRequest {
	s.ContactName = &v
	return s
}

func (s *QueryJobsRequest) SetEndTime(v int64) *QueryJobsRequest {
	s.EndTime = &v
	return s
}

func (s *QueryJobsRequest) SetInstanceId(v string) *QueryJobsRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryJobsRequest) SetJobGroupId(v string) *QueryJobsRequest {
	s.JobGroupId = &v
	return s
}

func (s *QueryJobsRequest) SetPageNumber(v int32) *QueryJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryJobsRequest) SetPageSize(v int32) *QueryJobsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryJobsRequest) SetPhoneNumber(v string) *QueryJobsRequest {
	s.PhoneNumber = &v
	return s
}

func (s *QueryJobsRequest) SetScenarioId(v string) *QueryJobsRequest {
	s.ScenarioId = &v
	return s
}

func (s *QueryJobsRequest) SetStartTime(v int64) *QueryJobsRequest {
	s.StartTime = &v
	return s
}

func (s *QueryJobsRequest) SetTimeAlignment(v string) *QueryJobsRequest {
	s.TimeAlignment = &v
	return s
}

func (s *QueryJobsRequest) Validate() error {
	return dara.Validate(s)
}
