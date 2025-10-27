// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPushRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *QueryPushRecordsResponseBody
	GetNextToken() *string
	SetPage(v int32) *QueryPushRecordsResponseBody
	GetPage() *int32
	SetPageSize(v int32) *QueryPushRecordsResponseBody
	GetPageSize() *int32
	SetPushInfos(v *QueryPushRecordsResponseBodyPushInfos) *QueryPushRecordsResponseBody
	GetPushInfos() *QueryPushRecordsResponseBodyPushInfos
	SetRequestId(v string) *QueryPushRecordsResponseBody
	GetRequestId() *string
	SetTotal(v int32) *QueryPushRecordsResponseBody
	GetTotal() *int32
}

type QueryPushRecordsResponseBody struct {
	// example:
	//
	// i91D***********kXIh/dVBEQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 11
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 20
	PageSize  *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PushInfos *QueryPushRecordsResponseBodyPushInfos `json:"PushInfos,omitempty" xml:"PushInfos,omitempty" type:"Struct"`
	// example:
	//
	// 9B24B396-249D-55E4-8CA1-66C9B50BB734
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 193
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryPushRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPushRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPushRecordsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryPushRecordsResponseBody) GetPage() *int32 {
	return s.Page
}

func (s *QueryPushRecordsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryPushRecordsResponseBody) GetPushInfos() *QueryPushRecordsResponseBodyPushInfos {
	return s.PushInfos
}

func (s *QueryPushRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPushRecordsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *QueryPushRecordsResponseBody) SetNextToken(v string) *QueryPushRecordsResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryPushRecordsResponseBody) SetPage(v int32) *QueryPushRecordsResponseBody {
	s.Page = &v
	return s
}

func (s *QueryPushRecordsResponseBody) SetPageSize(v int32) *QueryPushRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryPushRecordsResponseBody) SetPushInfos(v *QueryPushRecordsResponseBodyPushInfos) *QueryPushRecordsResponseBody {
	s.PushInfos = v
	return s
}

func (s *QueryPushRecordsResponseBody) SetRequestId(v string) *QueryPushRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPushRecordsResponseBody) SetTotal(v int32) *QueryPushRecordsResponseBody {
	s.Total = &v
	return s
}

func (s *QueryPushRecordsResponseBody) Validate() error {
	if s.PushInfos != nil {
		if err := s.PushInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryPushRecordsResponseBodyPushInfos struct {
	PushInfo []*QueryPushRecordsResponseBodyPushInfosPushInfo `json:"PushInfo,omitempty" xml:"PushInfo,omitempty" type:"Repeated"`
}

func (s QueryPushRecordsResponseBodyPushInfos) String() string {
	return dara.Prettify(s)
}

func (s QueryPushRecordsResponseBodyPushInfos) GoString() string {
	return s.String()
}

func (s *QueryPushRecordsResponseBodyPushInfos) GetPushInfo() []*QueryPushRecordsResponseBodyPushInfosPushInfo {
	return s.PushInfo
}

func (s *QueryPushRecordsResponseBodyPushInfos) SetPushInfo(v []*QueryPushRecordsResponseBodyPushInfosPushInfo) *QueryPushRecordsResponseBodyPushInfos {
	s.PushInfo = v
	return s
}

func (s *QueryPushRecordsResponseBodyPushInfos) Validate() error {
	if s.PushInfo != nil {
		for _, item := range s.PushInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryPushRecordsResponseBodyPushInfosPushInfo struct {
	// example:
	//
	// 333526247
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// abcd
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// example:
	//
	// ANDROID
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// example:
	//
	// 510431
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// example:
	//
	// 2021-09-15T02:05:24Z
	PushTime *string `json:"PushTime,omitempty" xml:"PushTime,omitempty"`
	// example:
	//
	// NOTICE
	PushType *string `json:"PushType,omitempty" xml:"PushType,omitempty"`
	// example:
	//
	// DEVICE
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// SENT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// DEVICE
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// example:
	//
	// sssss
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s QueryPushRecordsResponseBodyPushInfosPushInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryPushRecordsResponseBodyPushInfosPushInfo) GoString() string {
	return s.String()
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) GetAppKey() *int64 {
	return s.AppKey
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) GetBody() *string {
	return s.Body
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) GetDeviceType() *string {
	return s.DeviceType
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) GetMessageId() *string {
	return s.MessageId
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) GetPushTime() *string {
	return s.PushTime
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) GetPushType() *string {
	return s.PushType
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) GetSource() *string {
	return s.Source
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) GetStatus() *string {
	return s.Status
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) GetTarget() *string {
	return s.Target
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) GetTitle() *string {
	return s.Title
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) SetAppKey(v int64) *QueryPushRecordsResponseBodyPushInfosPushInfo {
	s.AppKey = &v
	return s
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) SetBody(v string) *QueryPushRecordsResponseBodyPushInfosPushInfo {
	s.Body = &v
	return s
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) SetDeviceType(v string) *QueryPushRecordsResponseBodyPushInfosPushInfo {
	s.DeviceType = &v
	return s
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) SetMessageId(v string) *QueryPushRecordsResponseBodyPushInfosPushInfo {
	s.MessageId = &v
	return s
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) SetPushTime(v string) *QueryPushRecordsResponseBodyPushInfosPushInfo {
	s.PushTime = &v
	return s
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) SetPushType(v string) *QueryPushRecordsResponseBodyPushInfosPushInfo {
	s.PushType = &v
	return s
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) SetSource(v string) *QueryPushRecordsResponseBodyPushInfosPushInfo {
	s.Source = &v
	return s
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) SetStatus(v string) *QueryPushRecordsResponseBodyPushInfosPushInfo {
	s.Status = &v
	return s
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) SetTarget(v string) *QueryPushRecordsResponseBodyPushInfosPushInfo {
	s.Target = &v
	return s
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) SetTitle(v string) *QueryPushRecordsResponseBodyPushInfosPushInfo {
	s.Title = &v
	return s
}

func (s *QueryPushRecordsResponseBodyPushInfosPushInfo) Validate() error {
	return dara.Validate(s)
}
