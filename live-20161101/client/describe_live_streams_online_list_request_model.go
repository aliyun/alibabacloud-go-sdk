// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamsOnlineListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveStreamsOnlineListRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLiveStreamsOnlineListRequest
	GetDomainName() *string
	SetOnlyStream(v string) *DescribeLiveStreamsOnlineListRequest
	GetOnlyStream() *string
	SetOwnerId(v int64) *DescribeLiveStreamsOnlineListRequest
	GetOwnerId() *int64
	SetPageNum(v int32) *DescribeLiveStreamsOnlineListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeLiveStreamsOnlineListRequest
	GetPageSize() *int32
	SetQueryType(v string) *DescribeLiveStreamsOnlineListRequest
	GetQueryType() *string
	SetRegionId(v string) *DescribeLiveStreamsOnlineListRequest
	GetRegionId() *string
	SetStreamName(v string) *DescribeLiveStreamsOnlineListRequest
	GetStreamName() *string
	SetStreamType(v string) *DescribeLiveStreamsOnlineListRequest
	GetStreamType() *string
}

type DescribeLiveStreamsOnlineListRequest struct {
	// The name of the application to which the live stream belongs. You can view the application name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
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
	// Specifies whether to return only specific parameters. Valid values:
	//
	// 	- **yes**: returns only the DomainName, AppName, StreamName, and PublishTime parameters.
	//
	// 	- **no**: returns all parameters. This is the default value.
	//
	// example:
	//
	// no
	OnlyStream *string `json:"OnlyStream,omitempty" xml:"OnlyStream,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Valid values: 1 to 3000. Default value: 2000.
	//
	// example:
	//
	// 1500
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The mode in which stream names are matched. Valid values:
	//
	// 	- **fuzzy*	- (default): fuzzy match
	//
	// 	- **strict**: exact match
	//
	// example:
	//
	// fuzzy
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the live stream. You can specify only one live stream. You can view the stream name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The type of the streams to query. Valid values:
	//
	// 	- **all*	- (default): all streams
	//
	// 	- **raw**: source streams
	//
	// 	- **trans**: transcoded streams
	//
	// example:
	//
	// all
	StreamType *string `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s DescribeLiveStreamsOnlineListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsOnlineListRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsOnlineListRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveStreamsOnlineListRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamsOnlineListRequest) GetOnlyStream() *string {
	return s.OnlyStream
}

func (s *DescribeLiveStreamsOnlineListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveStreamsOnlineListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeLiveStreamsOnlineListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveStreamsOnlineListRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *DescribeLiveStreamsOnlineListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveStreamsOnlineListRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveStreamsOnlineListRequest) GetStreamType() *string {
	return s.StreamType
}

func (s *DescribeLiveStreamsOnlineListRequest) SetAppName(v string) *DescribeLiveStreamsOnlineListRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListRequest) SetDomainName(v string) *DescribeLiveStreamsOnlineListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListRequest) SetOnlyStream(v string) *DescribeLiveStreamsOnlineListRequest {
	s.OnlyStream = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListRequest) SetOwnerId(v int64) *DescribeLiveStreamsOnlineListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListRequest) SetPageNum(v int32) *DescribeLiveStreamsOnlineListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListRequest) SetPageSize(v int32) *DescribeLiveStreamsOnlineListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListRequest) SetQueryType(v string) *DescribeLiveStreamsOnlineListRequest {
	s.QueryType = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListRequest) SetRegionId(v string) *DescribeLiveStreamsOnlineListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListRequest) SetStreamName(v string) *DescribeLiveStreamsOnlineListRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListRequest) SetStreamType(v string) *DescribeLiveStreamsOnlineListRequest {
	s.StreamType = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListRequest) Validate() error {
	return dara.Validate(s)
}
