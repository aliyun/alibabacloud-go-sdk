// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamsNotifyRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNotifyRecordsInfo(v *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfo) *DescribeLiveStreamsNotifyRecordsResponseBody
	GetNotifyRecordsInfo() *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfo
	SetPageNum(v int32) *DescribeLiveStreamsNotifyRecordsResponseBody
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeLiveStreamsNotifyRecordsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeLiveStreamsNotifyRecordsResponseBody
	GetRequestId() *string
	SetTotalNum(v int32) *DescribeLiveStreamsNotifyRecordsResponseBody
	GetTotalNum() *int32
	SetTotalPage(v int32) *DescribeLiveStreamsNotifyRecordsResponseBody
	GetTotalPage() *int32
}

type DescribeLiveStreamsNotifyRecordsResponseBody struct {
	// The stream ingest callback records.
	NotifyRecordsInfo *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfo `json:"NotifyRecordsInfo,omitempty" xml:"NotifyRecordsInfo,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F675E4B4-125D-1533-901B-11A724644285
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries that meet the specified conditions.
	//
	// example:
	//
	// 20
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 20
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeLiveStreamsNotifyRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsNotifyRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBody) GetNotifyRecordsInfo() *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfo {
	return s.NotifyRecordsInfo
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBody) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBody) SetNotifyRecordsInfo(v *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfo) *DescribeLiveStreamsNotifyRecordsResponseBody {
	s.NotifyRecordsInfo = v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBody) SetPageNum(v int32) *DescribeLiveStreamsNotifyRecordsResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBody) SetPageSize(v int32) *DescribeLiveStreamsNotifyRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBody) SetRequestId(v string) *DescribeLiveStreamsNotifyRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBody) SetTotalNum(v int32) *DescribeLiveStreamsNotifyRecordsResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBody) SetTotalPage(v int32) *DescribeLiveStreamsNotifyRecordsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfo struct {
	LiveStreamNotifyRecordsInfo []*DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo `json:"LiveStreamNotifyRecordsInfo,omitempty" xml:"LiveStreamNotifyRecordsInfo,omitempty" type:"Repeated"`
}

func (s DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfo) GetLiveStreamNotifyRecordsInfo() []*DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo {
	return s.LiveStreamNotifyRecordsInfo
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfo) SetLiveStreamNotifyRecordsInfo(v []*DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo) *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfo {
	s.LiveStreamNotifyRecordsInfo = v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo struct {
	// The name of the application to which the live stream belongs.
	//
	// example:
	//
	// app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The result of the request. If success is returned, the request is successful. If an error message is returned, the request failed.
	//
	// example:
	//
	// success
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ingest domain.
	//
	// example:
	//
	// push.example1.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The callback content.
	//
	// example:
	//
	// {\\"action\\":\\"publish_done\\",\\"app\\":\\"push.example1.com\\"}
	NotifyContent  *string `json:"NotifyContent,omitempty" xml:"NotifyContent,omitempty"`
	NotifyResponse *string `json:"NotifyResponse,omitempty" xml:"NotifyResponse,omitempty"`
	// The callback result. Valid values:
	//
	// 	- success
	//
	// 	- failed
	//
	// example:
	//
	// success
	NotifyResult *string `json:"NotifyResult,omitempty" xml:"NotifyResult,omitempty"`
	// The time when the callback was invoked. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-10-19T19:09:28Z
	NotifyTime *string `json:"NotifyTime,omitempty" xml:"NotifyTime,omitempty"`
	// The event. Valid values:
	//
	// 	- publish: The stream ingest starts.
	//
	// 	- publish_done: The stream ingest is interrupted.
	//
	// example:
	//
	// publish_done
	NotifyType *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	// The callback URL.
	//
	// example:
	//
	// http://xx.xx.xx.xx/callbacks
	NotifyUrl *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	// The name of the live stream.
	//
	// example:
	//
	// stream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo) GetDescription() *string {
	return s.Description
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo) GetNotifyContent() *string {
	return s.NotifyContent
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo) GetNotifyResponse() *string {
	return s.NotifyResponse
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo) GetNotifyResult() *string {
	return s.NotifyResult
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo) GetNotifyTime() *string {
	return s.NotifyTime
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo) GetNotifyType() *string {
	return s.NotifyType
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo) GetNotifyUrl() *string {
	return s.NotifyUrl
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo) SetAppName(v string) *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo) SetDescription(v string) *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo {
	s.Description = &v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo) SetDomainName(v string) *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo) SetNotifyContent(v string) *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo {
	s.NotifyContent = &v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo) SetNotifyResponse(v string) *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo {
	s.NotifyResponse = &v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo) SetNotifyResult(v string) *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo {
	s.NotifyResult = &v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo) SetNotifyTime(v string) *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo {
	s.NotifyTime = &v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo) SetNotifyType(v string) *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo {
	s.NotifyType = &v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo) SetNotifyUrl(v string) *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo {
	s.NotifyUrl = &v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo) SetStreamName(v string) *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsResponseBodyNotifyRecordsInfoLiveStreamNotifyRecordsInfo) Validate() error {
	return dara.Validate(s)
}
