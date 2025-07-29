// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveUserStreamMetricDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveUserStreamMetricDataRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLiveUserStreamMetricDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveUserStreamMetricDataRequest
	GetEndTime() *string
	SetPageNumber(v int64) *DescribeLiveUserStreamMetricDataRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeLiveUserStreamMetricDataRequest
	GetPageSize() *int64
	SetProtocol(v string) *DescribeLiveUserStreamMetricDataRequest
	GetProtocol() *string
	SetStartTime(v string) *DescribeLiveUserStreamMetricDataRequest
	GetStartTime() *string
	SetStreamName(v string) *DescribeLiveUserStreamMetricDataRequest
	GetStreamName() *string
}

type DescribeLiveUserStreamMetricDataRequest struct {
	// example:
	//
	// app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// test.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2015-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// flv
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// test.flv
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLiveUserStreamMetricDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUserStreamMetricDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserStreamMetricDataRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveUserStreamMetricDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveUserStreamMetricDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveUserStreamMetricDataRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeLiveUserStreamMetricDataRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeLiveUserStreamMetricDataRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeLiveUserStreamMetricDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveUserStreamMetricDataRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveUserStreamMetricDataRequest) SetAppName(v string) *DescribeLiveUserStreamMetricDataRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataRequest) SetDomainName(v string) *DescribeLiveUserStreamMetricDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataRequest) SetEndTime(v string) *DescribeLiveUserStreamMetricDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataRequest) SetPageNumber(v int64) *DescribeLiveUserStreamMetricDataRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataRequest) SetPageSize(v int64) *DescribeLiveUserStreamMetricDataRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataRequest) SetProtocol(v string) *DescribeLiveUserStreamMetricDataRequest {
	s.Protocol = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataRequest) SetStartTime(v string) *DescribeLiveUserStreamMetricDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataRequest) SetStreamName(v string) *DescribeLiveUserStreamMetricDataRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataRequest) Validate() error {
	return dara.Validate(s)
}
