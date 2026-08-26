// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMixStreamListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeMixStreamListRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeMixStreamListRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeMixStreamListRequest
	GetEndTime() *string
	SetMixStreamId(v string) *DescribeMixStreamListRequest
	GetMixStreamId() *string
	SetOwnerId(v int64) *DescribeMixStreamListRequest
	GetOwnerId() *int64
	SetPageNo(v int32) *DescribeMixStreamListRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeMixStreamListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeMixStreamListRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeMixStreamListRequest
	GetStartTime() *string
	SetStreamName(v string) *DescribeMixStreamListRequest
	GetStreamName() *string
}

type DescribeMixStreamListRequest struct {
	// The app name.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The streaming domain.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time. The time is in the *yyyy-MM-dd*T*HH:mm:ss*Z format and in UTC.
	//
	// example:
	//
	// 2020-09-20T13:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the stream mixing task. If you create a stream mixing task by calling the [CreateMixStream](https://help.aliyun.com/document_detail/2848087.html) operation, use the MixStreamId value that is returned in the response.
	//
	// example:
	//
	// 5b2a046e-74d7-385e-d2d7-8a5b87e4****
	MixStreamId *string `json:"MixStreamId,omitempty" xml:"MixStreamId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be greater than **0*	- and cannot exceed the maximum value of the Integer data type. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of records to display on each page. Default value: **1000**.
	//
	// example:
	//
	// 1000
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time. The time is in the *yyyy-MM-dd*T*HH:mm:ss*Z format and in UTC.
	//
	// example:
	//
	// 2020-09-11T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The stream name of the stream mixing task.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeMixStreamListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMixStreamListRequest) GoString() string {
	return s.String()
}

func (s *DescribeMixStreamListRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeMixStreamListRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeMixStreamListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeMixStreamListRequest) GetMixStreamId() *string {
	return s.MixStreamId
}

func (s *DescribeMixStreamListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeMixStreamListRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeMixStreamListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMixStreamListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMixStreamListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeMixStreamListRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeMixStreamListRequest) SetAppName(v string) *DescribeMixStreamListRequest {
	s.AppName = &v
	return s
}

func (s *DescribeMixStreamListRequest) SetDomainName(v string) *DescribeMixStreamListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeMixStreamListRequest) SetEndTime(v string) *DescribeMixStreamListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeMixStreamListRequest) SetMixStreamId(v string) *DescribeMixStreamListRequest {
	s.MixStreamId = &v
	return s
}

func (s *DescribeMixStreamListRequest) SetOwnerId(v int64) *DescribeMixStreamListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMixStreamListRequest) SetPageNo(v int32) *DescribeMixStreamListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeMixStreamListRequest) SetPageSize(v int32) *DescribeMixStreamListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMixStreamListRequest) SetRegionId(v string) *DescribeMixStreamListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMixStreamListRequest) SetStartTime(v string) *DescribeMixStreamListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeMixStreamListRequest) SetStreamName(v string) *DescribeMixStreamListRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeMixStreamListRequest) Validate() error {
	return dara.Validate(s)
}
