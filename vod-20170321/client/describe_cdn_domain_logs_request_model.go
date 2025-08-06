// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDomainLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeCdnDomainLogsRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeCdnDomainLogsRequest
	GetEndTime() *string
	SetLogDay(v string) *DescribeCdnDomainLogsRequest
	GetLogDay() *string
	SetOwnerId(v string) *DescribeCdnDomainLogsRequest
	GetOwnerId() *string
	SetPageNo(v int64) *DescribeCdnDomainLogsRequest
	GetPageNo() *int64
	SetPageSize(v int64) *DescribeCdnDomainLogsRequest
	GetPageSize() *int64
	SetResourceOwnerId(v string) *DescribeCdnDomainLogsRequest
	GetResourceOwnerId() *string
	SetResourceRealOwnerId(v string) *DescribeCdnDomainLogsRequest
	GetResourceRealOwnerId() *string
	SetStartTime(v string) *DescribeCdnDomainLogsRequest
	GetStartTime() *string
}

type DescribeCdnDomainLogsRequest struct {
	// This parameter is required.
	DomainName          *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime             *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	LogDay              *string `json:"LogDay,omitempty" xml:"LogDay,omitempty"`
	OwnerId             *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo              *int64  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize            *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerId     *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId *string `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
	StartTime           *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeCdnDomainLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeCdnDomainLogsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCdnDomainLogsRequest) GetLogDay() *string {
	return s.LogDay
}

func (s *DescribeCdnDomainLogsRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeCdnDomainLogsRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DescribeCdnDomainLogsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeCdnDomainLogsRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *DescribeCdnDomainLogsRequest) GetResourceRealOwnerId() *string {
	return s.ResourceRealOwnerId
}

func (s *DescribeCdnDomainLogsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCdnDomainLogsRequest) SetDomainName(v string) *DescribeCdnDomainLogsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnDomainLogsRequest) SetEndTime(v string) *DescribeCdnDomainLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCdnDomainLogsRequest) SetLogDay(v string) *DescribeCdnDomainLogsRequest {
	s.LogDay = &v
	return s
}

func (s *DescribeCdnDomainLogsRequest) SetOwnerId(v string) *DescribeCdnDomainLogsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnDomainLogsRequest) SetPageNo(v int64) *DescribeCdnDomainLogsRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeCdnDomainLogsRequest) SetPageSize(v int64) *DescribeCdnDomainLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCdnDomainLogsRequest) SetResourceOwnerId(v string) *DescribeCdnDomainLogsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCdnDomainLogsRequest) SetResourceRealOwnerId(v string) *DescribeCdnDomainLogsRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *DescribeCdnDomainLogsRequest) SetStartTime(v string) *DescribeCdnDomainLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCdnDomainLogsRequest) Validate() error {
	return dara.Validate(s)
}
