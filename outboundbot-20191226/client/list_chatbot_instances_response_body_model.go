// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatbotInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBots(v []*ListChatbotInstancesResponseBodyBots) *ListChatbotInstancesResponseBody
	GetBots() []*ListChatbotInstancesResponseBodyBots
	SetCode(v string) *ListChatbotInstancesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListChatbotInstancesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListChatbotInstancesResponseBody
	GetMessage() *string
	SetPageNumber(v int64) *ListChatbotInstancesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListChatbotInstancesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListChatbotInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListChatbotInstancesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListChatbotInstancesResponseBody
	GetTotalCount() *int64
}

type ListChatbotInstancesResponseBody struct {
	// List of bot information
	Bots []*ListChatbotInstancesResponseBodyBots `json:"Bots,omitempty" xml:"Bots,omitempty" type:"Repeated"`
	// Response code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Page number
	//
	// example:
	//
	// 5
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Count
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Total count
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListChatbotInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListChatbotInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListChatbotInstancesResponseBody) GetBots() []*ListChatbotInstancesResponseBodyBots {
	return s.Bots
}

func (s *ListChatbotInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListChatbotInstancesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListChatbotInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListChatbotInstancesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListChatbotInstancesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListChatbotInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListChatbotInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListChatbotInstancesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListChatbotInstancesResponseBody) SetBots(v []*ListChatbotInstancesResponseBodyBots) *ListChatbotInstancesResponseBody {
	s.Bots = v
	return s
}

func (s *ListChatbotInstancesResponseBody) SetCode(v string) *ListChatbotInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListChatbotInstancesResponseBody) SetHttpStatusCode(v int32) *ListChatbotInstancesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListChatbotInstancesResponseBody) SetMessage(v string) *ListChatbotInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListChatbotInstancesResponseBody) SetPageNumber(v int64) *ListChatbotInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListChatbotInstancesResponseBody) SetPageSize(v int64) *ListChatbotInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListChatbotInstancesResponseBody) SetRequestId(v string) *ListChatbotInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChatbotInstancesResponseBody) SetSuccess(v bool) *ListChatbotInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *ListChatbotInstancesResponseBody) SetTotalCount(v int64) *ListChatbotInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListChatbotInstancesResponseBody) Validate() error {
	if s.Bots != nil {
		for _, item := range s.Bots {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListChatbotInstancesResponseBodyBots struct {
	// URL of the profile picture
	//
	// example:
	//
	// https://ccrm.wengine.cn/ccrm/system/common/fileDownload/noToken?fileId=975cdeaa064846e3b6004abd9ba1d7c8
	Avatar *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	// Creation Time
	//
	// example:
	//
	// 2022-01-18T02:36:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// instance ID
	//
	// example:
	//
	// e874fcf0-d2f4-4e62-9377-b6f35fe55210
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Bot remark
	//
	// example:
	//
	// 这是直播的描述
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	// Language of the bot service, such as zh-cn or en-us
	//
	// example:
	//
	// zh-cn
	LanguageCode *string `json:"LanguageCode,omitempty" xml:"LanguageCode,omitempty"`
	// Bot name
	//
	// example:
	//
	// 智能回访2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Time zone of the bot, refer to "Common - Time Zone Codes"
	//
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s ListChatbotInstancesResponseBodyBots) String() string {
	return dara.Prettify(s)
}

func (s ListChatbotInstancesResponseBodyBots) GoString() string {
	return s.String()
}

func (s *ListChatbotInstancesResponseBodyBots) GetAvatar() *string {
	return s.Avatar
}

func (s *ListChatbotInstancesResponseBodyBots) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListChatbotInstancesResponseBodyBots) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListChatbotInstancesResponseBodyBots) GetIntroduction() *string {
	return s.Introduction
}

func (s *ListChatbotInstancesResponseBodyBots) GetLanguageCode() *string {
	return s.LanguageCode
}

func (s *ListChatbotInstancesResponseBodyBots) GetName() *string {
	return s.Name
}

func (s *ListChatbotInstancesResponseBodyBots) GetTimeZone() *string {
	return s.TimeZone
}

func (s *ListChatbotInstancesResponseBodyBots) SetAvatar(v string) *ListChatbotInstancesResponseBodyBots {
	s.Avatar = &v
	return s
}

func (s *ListChatbotInstancesResponseBodyBots) SetCreateTime(v string) *ListChatbotInstancesResponseBodyBots {
	s.CreateTime = &v
	return s
}

func (s *ListChatbotInstancesResponseBodyBots) SetInstanceId(v string) *ListChatbotInstancesResponseBodyBots {
	s.InstanceId = &v
	return s
}

func (s *ListChatbotInstancesResponseBodyBots) SetIntroduction(v string) *ListChatbotInstancesResponseBodyBots {
	s.Introduction = &v
	return s
}

func (s *ListChatbotInstancesResponseBodyBots) SetLanguageCode(v string) *ListChatbotInstancesResponseBodyBots {
	s.LanguageCode = &v
	return s
}

func (s *ListChatbotInstancesResponseBodyBots) SetName(v string) *ListChatbotInstancesResponseBodyBots {
	s.Name = &v
	return s
}

func (s *ListChatbotInstancesResponseBodyBots) SetTimeZone(v string) *ListChatbotInstancesResponseBodyBots {
	s.TimeZone = &v
	return s
}

func (s *ListChatbotInstancesResponseBodyBots) Validate() error {
	return dara.Validate(s)
}
