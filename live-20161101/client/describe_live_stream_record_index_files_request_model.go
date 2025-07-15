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
	// System-defined parameter. Value: **DescribeLiveStreamRecordIndexFiles**.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// ## [](#)Usage notes
	//
	// 	- ApsaraVideo Live stores the information about an M3U8 index file for six months. You can query only the information of index files created in the previous six months.
	//
	// 	- M3U8 index files are stored in Object Storage Service (OSS) buckets. The retention period is determined by the storage configuration of the OSS buckets.
	//
	// ## [](#qps-)QPS limit
	//
	// You can call this operation up to 15 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/343507.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The name of the live stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-12-22T08:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The order in which the entries are sorted based on creation time. Valid values:
	//
	// 	- **asc*	- (default): ascending order
	//
	// 	- **desc**: descending order
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
	// The name of the application to which the live stream belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-12-21T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The main streaming domain.
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
