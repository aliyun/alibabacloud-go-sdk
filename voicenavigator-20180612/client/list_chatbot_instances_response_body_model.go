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
	SetPageNumber(v int32) *ListChatbotInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int64) *ListChatbotInstancesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListChatbotInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListChatbotInstancesResponseBody
	GetTotalCount() *int64
}

type ListChatbotInstancesResponseBody struct {
	Bots []*ListChatbotInstancesResponseBodyBots `json:"Bots,omitempty" xml:"Bots,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// da37319b-6c83-4268-9f19-814aed62e401
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
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

func (s *ListChatbotInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListChatbotInstancesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListChatbotInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListChatbotInstancesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListChatbotInstancesResponseBody) SetBots(v []*ListChatbotInstancesResponseBodyBots) *ListChatbotInstancesResponseBody {
	s.Bots = v
	return s
}

func (s *ListChatbotInstancesResponseBody) SetPageNumber(v int32) *ListChatbotInstancesResponseBody {
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
	// example:
	//
	// https://dss0.ali.com/70cFuHS.jpg
	Avatar *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	// example:
	//
	// 1582266750353
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ‘’
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	// example:
	//
	// zh-cn
	LanguageCode *string `json:"LanguageCode,omitempty" xml:"LanguageCode,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// UTC+8
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
