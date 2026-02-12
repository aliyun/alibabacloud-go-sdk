// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTopicSubDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OnsTopicSubDetailResponseBodyData) *OnsTopicSubDetailResponseBody
	GetData() *OnsTopicSubDetailResponseBodyData
	SetRequestId(v string) *OnsTopicSubDetailResponseBody
	GetRequestId() *string
}

type OnsTopicSubDetailResponseBody struct {
	// The data returned.
	Data *OnsTopicSubDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 87B6207F-2908-42B5-A134-84956DCA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsTopicSubDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsTopicSubDetailResponseBody) GoString() string {
	return s.String()
}

func (s *OnsTopicSubDetailResponseBody) GetData() *OnsTopicSubDetailResponseBodyData {
	return s.Data
}

func (s *OnsTopicSubDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsTopicSubDetailResponseBody) SetData(v *OnsTopicSubDetailResponseBodyData) *OnsTopicSubDetailResponseBody {
	s.Data = v
	return s
}

func (s *OnsTopicSubDetailResponseBody) SetRequestId(v string) *OnsTopicSubDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsTopicSubDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsTopicSubDetailResponseBodyData struct {
	SubscriptionDataList *OnsTopicSubDetailResponseBodyDataSubscriptionDataList `json:"SubscriptionDataList,omitempty" xml:"SubscriptionDataList,omitempty" type:"Struct"`
	// The topic name.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsTopicSubDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OnsTopicSubDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsTopicSubDetailResponseBodyData) GetSubscriptionDataList() *OnsTopicSubDetailResponseBodyDataSubscriptionDataList {
	return s.SubscriptionDataList
}

func (s *OnsTopicSubDetailResponseBodyData) GetTopic() *string {
	return s.Topic
}

func (s *OnsTopicSubDetailResponseBodyData) SetSubscriptionDataList(v *OnsTopicSubDetailResponseBodyDataSubscriptionDataList) *OnsTopicSubDetailResponseBodyData {
	s.SubscriptionDataList = v
	return s
}

func (s *OnsTopicSubDetailResponseBodyData) SetTopic(v string) *OnsTopicSubDetailResponseBodyData {
	s.Topic = &v
	return s
}

func (s *OnsTopicSubDetailResponseBodyData) Validate() error {
	if s.SubscriptionDataList != nil {
		if err := s.SubscriptionDataList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsTopicSubDetailResponseBodyDataSubscriptionDataList struct {
	SubscriptionDataList []*OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList `json:"SubscriptionDataList,omitempty" xml:"SubscriptionDataList,omitempty" type:"Repeated"`
}

func (s OnsTopicSubDetailResponseBodyDataSubscriptionDataList) String() string {
	return dara.Prettify(s)
}

func (s OnsTopicSubDetailResponseBodyDataSubscriptionDataList) GoString() string {
	return s.String()
}

func (s *OnsTopicSubDetailResponseBodyDataSubscriptionDataList) GetSubscriptionDataList() []*OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList {
	return s.SubscriptionDataList
}

func (s *OnsTopicSubDetailResponseBodyDataSubscriptionDataList) SetSubscriptionDataList(v []*OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) *OnsTopicSubDetailResponseBodyDataSubscriptionDataList {
	s.SubscriptionDataList = v
	return s
}

func (s *OnsTopicSubDetailResponseBodyDataSubscriptionDataList) Validate() error {
	if s.SubscriptionDataList != nil {
		for _, item := range s.SubscriptionDataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList struct {
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	MessageModel *string `json:"MessageModel,omitempty" xml:"MessageModel,omitempty"`
	Online       *string `json:"Online,omitempty" xml:"Online,omitempty"`
	SubString    *string `json:"SubString,omitempty" xml:"SubString,omitempty"`
}

func (s OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) String() string {
	return dara.Prettify(s)
}

func (s OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) GoString() string {
	return s.String()
}

func (s *OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) GetGroupId() *string {
	return s.GroupId
}

func (s *OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) GetMessageModel() *string {
	return s.MessageModel
}

func (s *OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) GetOnline() *string {
	return s.Online
}

func (s *OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) GetSubString() *string {
	return s.SubString
}

func (s *OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) SetGroupId(v string) *OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList {
	s.GroupId = &v
	return s
}

func (s *OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) SetMessageModel(v string) *OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList {
	s.MessageModel = &v
	return s
}

func (s *OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) SetOnline(v string) *OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList {
	s.Online = &v
	return s
}

func (s *OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) SetSubString(v string) *OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList {
	s.SubString = &v
	return s
}

func (s *OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) Validate() error {
	return dara.Validate(s)
}
