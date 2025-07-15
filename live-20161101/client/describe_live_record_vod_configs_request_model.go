// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveRecordVodConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveRecordVodConfigsRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLiveRecordVodConfigsRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeLiveRecordVodConfigsRequest
	GetOwnerId() *int64
	SetPageNum(v int64) *DescribeLiveRecordVodConfigsRequest
	GetPageNum() *int64
	SetPageSize(v int64) *DescribeLiveRecordVodConfigsRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeLiveRecordVodConfigsRequest
	GetRegionId() *string
	SetStreamName(v string) *DescribeLiveRecordVodConfigsRequest
	GetStreamName() *string
}

type DescribeLiveRecordVodConfigsRequest struct {
	// The name of the application to which the live stream belongs.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: **10**. Valid values: **5 to 100**.
	//
	// example:
	//
	// 10
	PageSize *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the live stream.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLiveRecordVodConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRecordVodConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordVodConfigsRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveRecordVodConfigsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveRecordVodConfigsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveRecordVodConfigsRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeLiveRecordVodConfigsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeLiveRecordVodConfigsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveRecordVodConfigsRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveRecordVodConfigsRequest) SetAppName(v string) *DescribeLiveRecordVodConfigsRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsRequest) SetDomainName(v string) *DescribeLiveRecordVodConfigsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsRequest) SetOwnerId(v int64) *DescribeLiveRecordVodConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsRequest) SetPageNum(v int64) *DescribeLiveRecordVodConfigsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsRequest) SetPageSize(v int64) *DescribeLiveRecordVodConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsRequest) SetRegionId(v string) *DescribeLiveRecordVodConfigsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsRequest) SetStreamName(v string) *DescribeLiveRecordVodConfigsRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsRequest) Validate() error {
	return dara.Validate(s)
}
