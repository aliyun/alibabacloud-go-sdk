// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCloudRecordTextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *QueryCloudRecordTextRequest
	GetDirection() *string
	SetMaxResults(v int64) *QueryCloudRecordTextRequest
	GetMaxResults() *int64
	SetNextToken(v int64) *QueryCloudRecordTextRequest
	GetNextToken() *int64
	SetStartTime(v int64) *QueryCloudRecordTextRequest
	GetStartTime() *int64
	SetTenantContext(v *QueryCloudRecordTextRequestTenantContext) *QueryCloudRecordTextRequest
	GetTenantContext() *QueryCloudRecordTextRequestTenantContext
	SetConferenceId(v string) *QueryCloudRecordTextRequest
	GetConferenceId() *string
}

type QueryCloudRecordTextRequest struct {
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
	StartTime     *int64                                    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantContext *QueryCloudRecordTextRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 61289fxxx
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
}

func (s QueryCloudRecordTextRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCloudRecordTextRequest) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordTextRequest) GetDirection() *string {
	return s.Direction
}

func (s *QueryCloudRecordTextRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *QueryCloudRecordTextRequest) GetNextToken() *int64 {
	return s.NextToken
}

func (s *QueryCloudRecordTextRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryCloudRecordTextRequest) GetTenantContext() *QueryCloudRecordTextRequestTenantContext {
	return s.TenantContext
}

func (s *QueryCloudRecordTextRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *QueryCloudRecordTextRequest) SetDirection(v string) *QueryCloudRecordTextRequest {
	s.Direction = &v
	return s
}

func (s *QueryCloudRecordTextRequest) SetMaxResults(v int64) *QueryCloudRecordTextRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryCloudRecordTextRequest) SetNextToken(v int64) *QueryCloudRecordTextRequest {
	s.NextToken = &v
	return s
}

func (s *QueryCloudRecordTextRequest) SetStartTime(v int64) *QueryCloudRecordTextRequest {
	s.StartTime = &v
	return s
}

func (s *QueryCloudRecordTextRequest) SetTenantContext(v *QueryCloudRecordTextRequestTenantContext) *QueryCloudRecordTextRequest {
	s.TenantContext = v
	return s
}

func (s *QueryCloudRecordTextRequest) SetConferenceId(v string) *QueryCloudRecordTextRequest {
	s.ConferenceId = &v
	return s
}

func (s *QueryCloudRecordTextRequest) Validate() error {
	return dara.Validate(s)
}

type QueryCloudRecordTextRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryCloudRecordTextRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s QueryCloudRecordTextRequestTenantContext) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordTextRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryCloudRecordTextRequestTenantContext) SetTenantId(v string) *QueryCloudRecordTextRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *QueryCloudRecordTextRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
