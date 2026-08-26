// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamRecordIndexFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveStreamRecordIndexFilesRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLiveStreamRecordIndexFilesRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveStreamRecordIndexFilesRequest
	GetEndTime() *string
	SetOrder(v string) *DescribeLiveStreamRecordIndexFilesRequest
	GetOrder() *string
	SetOwnerId(v int64) *DescribeLiveStreamRecordIndexFilesRequest
	GetOwnerId() *int64
	SetPageNum(v int32) *DescribeLiveStreamRecordIndexFilesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeLiveStreamRecordIndexFilesRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeLiveStreamRecordIndexFilesRequest
	GetSecurityToken() *string
	SetStartTime(v string) *DescribeLiveStreamRecordIndexFilesRequest
	GetStartTime() *string
	SetStreamName(v string) *DescribeLiveStreamRecordIndexFilesRequest
	GetStreamName() *string
}

type DescribeLiveStreamRecordIndexFilesRequest struct {
	// The name of the application to which the live stream belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The streaming domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time. The interval between EndTime and StartTime cannot exceed 4 days. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-12-22T08:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The order in which entries are sorted by creation time. Valid values:
	//
	// - **asc*	- (default): ascending order.
	//
	// - **desc**: descending order.
	//
	// example:
	//
	// asc
	Order   *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Valid values: **5 to 30**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The start time. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-12-21T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The stream name.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLiveStreamRecordIndexFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamRecordIndexFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) SetAppName(v string) *DescribeLiveStreamRecordIndexFilesRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) SetDomainName(v string) *DescribeLiveStreamRecordIndexFilesRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) SetEndTime(v string) *DescribeLiveStreamRecordIndexFilesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) SetOrder(v string) *DescribeLiveStreamRecordIndexFilesRequest {
	s.Order = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) SetOwnerId(v int64) *DescribeLiveStreamRecordIndexFilesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) SetPageNum(v int32) *DescribeLiveStreamRecordIndexFilesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) SetPageSize(v int32) *DescribeLiveStreamRecordIndexFilesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) SetSecurityToken(v string) *DescribeLiveStreamRecordIndexFilesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) SetStartTime(v string) *DescribeLiveStreamRecordIndexFilesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) SetStreamName(v string) *DescribeLiveStreamRecordIndexFilesRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) Validate() error {
	return dara.Validate(s)
}
