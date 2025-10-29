// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveRecordNotifyRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCallbackList(v []*DescribeLiveRecordNotifyRecordsResponseBodyCallbackList) *DescribeLiveRecordNotifyRecordsResponseBody
	GetCallbackList() []*DescribeLiveRecordNotifyRecordsResponseBodyCallbackList
	SetCode(v int32) *DescribeLiveRecordNotifyRecordsResponseBody
	GetCode() *int32
	SetMsg(v string) *DescribeLiveRecordNotifyRecordsResponseBody
	GetMsg() *string
	SetPageNum(v int32) *DescribeLiveRecordNotifyRecordsResponseBody
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeLiveRecordNotifyRecordsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeLiveRecordNotifyRecordsResponseBody
	GetRequestId() *string
	SetTotalNum(v int32) *DescribeLiveRecordNotifyRecordsResponseBody
	GetTotalNum() *int32
	SetTotalPage(v int32) *DescribeLiveRecordNotifyRecordsResponseBody
	GetTotalPage() *int32
}

type DescribeLiveRecordNotifyRecordsResponseBody struct {
	// The callback records.
	CallbackList []*DescribeLiveRecordNotifyRecordsResponseBodyCallbackList `json:"CallbackList,omitempty" xml:"CallbackList,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// ok
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
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
	// 180FA0D2-1A02-5158-A36B-115DBF7B218D
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

func (s DescribeLiveRecordNotifyRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRecordNotifyRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordNotifyRecordsResponseBody) GetCallbackList() []*DescribeLiveRecordNotifyRecordsResponseBodyCallbackList {
	return s.CallbackList
}

func (s *DescribeLiveRecordNotifyRecordsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeLiveRecordNotifyRecordsResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *DescribeLiveRecordNotifyRecordsResponseBody) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeLiveRecordNotifyRecordsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveRecordNotifyRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveRecordNotifyRecordsResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *DescribeLiveRecordNotifyRecordsResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeLiveRecordNotifyRecordsResponseBody) SetCallbackList(v []*DescribeLiveRecordNotifyRecordsResponseBodyCallbackList) *DescribeLiveRecordNotifyRecordsResponseBody {
	s.CallbackList = v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsResponseBody) SetCode(v int32) *DescribeLiveRecordNotifyRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsResponseBody) SetMsg(v string) *DescribeLiveRecordNotifyRecordsResponseBody {
	s.Msg = &v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsResponseBody) SetPageNum(v int32) *DescribeLiveRecordNotifyRecordsResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsResponseBody) SetPageSize(v int32) *DescribeLiveRecordNotifyRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsResponseBody) SetRequestId(v string) *DescribeLiveRecordNotifyRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsResponseBody) SetTotalNum(v int32) *DescribeLiveRecordNotifyRecordsResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsResponseBody) SetTotalPage(v int32) *DescribeLiveRecordNotifyRecordsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsResponseBody) Validate() error {
	if s.CallbackList != nil {
		for _, item := range s.CallbackList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveRecordNotifyRecordsResponseBodyCallbackList struct {
	// The name of the application to which the live stream belongs.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The description of the result. A value of success indicates that the request is successful. If the request fails, an error message is returned.
	//
	// example:
	//
	// success
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The main streaming domain.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The callback content.
	//
	// example:
	//
	// July 26,  16:14{"domain":"al.xxxx.com","stream":"livestream01","pull _stream_result":true,"cdn":"al"}
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
	// The time when the callback was returned. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-10-19T19:09:28Z
	NotifyTime *string `json:"NotifyTime,omitempty" xml:"NotifyTime,omitempty"`
	// The callback type. Valid values:
	//
	// 	- file_created: The recording file is created.
	//
	// 	- record_error: A recording error occurs.
	//
	// 	- record_started: Recording is started.
	//
	// 	- record_paused: Recording is paused.
	//
	// 	- record_resumed: Recording is resumed.
	//
	// 	- record_force_transcode_fail: The recording task fails to trigger transcoding.
	//
	// 	- transformat_error: An error occurs when the live stream is parsed.
	//
	// example:
	//
	// record_started
	NotifyType *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	// The recording callback URL.
	//
	// example:
	//
	// http://learn.aliyundoc.com/examplecallback.action
	NotifyUrl *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	// The name of the live stream.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLiveRecordNotifyRecordsResponseBodyCallbackList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRecordNotifyRecordsResponseBodyCallbackList) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList) GetDescription() *string {
	return s.Description
}

func (s *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList) GetNotifyContent() *string {
	return s.NotifyContent
}

func (s *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList) GetNotifyResponse() *string {
	return s.NotifyResponse
}

func (s *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList) GetNotifyResult() *string {
	return s.NotifyResult
}

func (s *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList) GetNotifyTime() *string {
	return s.NotifyTime
}

func (s *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList) GetNotifyType() *string {
	return s.NotifyType
}

func (s *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList) GetNotifyUrl() *string {
	return s.NotifyUrl
}

func (s *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList) SetAppName(v string) *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList {
	s.AppName = &v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList) SetDescription(v string) *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList {
	s.Description = &v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList) SetDomainName(v string) *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList) SetNotifyContent(v string) *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList {
	s.NotifyContent = &v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList) SetNotifyResponse(v string) *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList {
	s.NotifyResponse = &v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList) SetNotifyResult(v string) *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList {
	s.NotifyResult = &v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList) SetNotifyTime(v string) *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList {
	s.NotifyTime = &v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList) SetNotifyType(v string) *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList {
	s.NotifyType = &v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList) SetNotifyUrl(v string) *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList {
	s.NotifyUrl = &v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList) SetStreamName(v string) *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsResponseBodyCallbackList) Validate() error {
	return dara.Validate(s)
}
