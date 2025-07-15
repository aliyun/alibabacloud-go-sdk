// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLiveDomainMultiStreamListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *QueryLiveDomainMultiStreamListRequest
	GetDomain() *string
	SetEndTime(v string) *QueryLiveDomainMultiStreamListRequest
	GetEndTime() *string
	SetOwnerId(v int64) *QueryLiveDomainMultiStreamListRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *QueryLiveDomainMultiStreamListRequest
	GetPageNumber() *int64
	SetPageSize(v int32) *QueryLiveDomainMultiStreamListRequest
	GetPageSize() *int32
	SetStartTime(v string) *QueryLiveDomainMultiStreamListRequest
	GetStartTime() *string
	SetStreamName(v string) *QueryLiveDomainMultiStreamListRequest
	GetStreamName() *string
}

type QueryLiveDomainMultiStreamListRequest struct {
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// The end time must be later than the start time. The time range specified by the StartTime and EndTime parameters cannot exceed seven days. If the two parameters are not specified, data of the last 24 hours is queried by default.
	//
	// example:
	//
	// 2024-12-02T01:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC. The time range specified by the StartTime and EndTime parameters cannot exceed seven days.
	//
	// example:
	//
	// 2024-12-01T01:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the live stream. This parameter is used for exact match.
	//
	// example:
	//
	// teststream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s QueryLiveDomainMultiStreamListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveDomainMultiStreamListRequest) GoString() string {
	return s.String()
}

func (s *QueryLiveDomainMultiStreamListRequest) GetDomain() *string {
	return s.Domain
}

func (s *QueryLiveDomainMultiStreamListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryLiveDomainMultiStreamListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryLiveDomainMultiStreamListRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *QueryLiveDomainMultiStreamListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryLiveDomainMultiStreamListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryLiveDomainMultiStreamListRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *QueryLiveDomainMultiStreamListRequest) SetDomain(v string) *QueryLiveDomainMultiStreamListRequest {
	s.Domain = &v
	return s
}

func (s *QueryLiveDomainMultiStreamListRequest) SetEndTime(v string) *QueryLiveDomainMultiStreamListRequest {
	s.EndTime = &v
	return s
}

func (s *QueryLiveDomainMultiStreamListRequest) SetOwnerId(v int64) *QueryLiveDomainMultiStreamListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryLiveDomainMultiStreamListRequest) SetPageNumber(v int64) *QueryLiveDomainMultiStreamListRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryLiveDomainMultiStreamListRequest) SetPageSize(v int32) *QueryLiveDomainMultiStreamListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryLiveDomainMultiStreamListRequest) SetStartTime(v string) *QueryLiveDomainMultiStreamListRequest {
	s.StartTime = &v
	return s
}

func (s *QueryLiveDomainMultiStreamListRequest) SetStreamName(v string) *QueryLiveDomainMultiStreamListRequest {
	s.StreamName = &v
	return s
}

func (s *QueryLiveDomainMultiStreamListRequest) Validate() error {
	return dara.Validate(s)
}
