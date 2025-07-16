// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCloudRecordTextShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *QueryCloudRecordTextShrinkRequest
	GetDirection() *string
	SetMaxResults(v int64) *QueryCloudRecordTextShrinkRequest
	GetMaxResults() *int64
	SetNextToken(v int64) *QueryCloudRecordTextShrinkRequest
	GetNextToken() *int64
	SetStartTime(v int64) *QueryCloudRecordTextShrinkRequest
	GetStartTime() *int64
	SetTenantContextShrink(v string) *QueryCloudRecordTextShrinkRequest
	GetTenantContextShrink() *string
	SetConferenceId(v string) *QueryCloudRecordTextShrinkRequest
	GetConferenceId() *string
}

type QueryCloudRecordTextShrinkRequest struct {
	// example:
	//
	// 0
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// example:
	//
	// 20000
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 1631172045153000
	NextToken *int64 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 7940
	StartTime           *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 61289fxxx
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
}

func (s QueryCloudRecordTextShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCloudRecordTextShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordTextShrinkRequest) GetDirection() *string {
	return s.Direction
}

func (s *QueryCloudRecordTextShrinkRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *QueryCloudRecordTextShrinkRequest) GetNextToken() *int64 {
	return s.NextToken
}

func (s *QueryCloudRecordTextShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryCloudRecordTextShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *QueryCloudRecordTextShrinkRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *QueryCloudRecordTextShrinkRequest) SetDirection(v string) *QueryCloudRecordTextShrinkRequest {
	s.Direction = &v
	return s
}

func (s *QueryCloudRecordTextShrinkRequest) SetMaxResults(v int64) *QueryCloudRecordTextShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryCloudRecordTextShrinkRequest) SetNextToken(v int64) *QueryCloudRecordTextShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *QueryCloudRecordTextShrinkRequest) SetStartTime(v int64) *QueryCloudRecordTextShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *QueryCloudRecordTextShrinkRequest) SetTenantContextShrink(v string) *QueryCloudRecordTextShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *QueryCloudRecordTextShrinkRequest) SetConferenceId(v string) *QueryCloudRecordTextShrinkRequest {
	s.ConferenceId = &v
	return s
}

func (s *QueryCloudRecordTextShrinkRequest) Validate() error {
	return dara.Validate(s)
}
