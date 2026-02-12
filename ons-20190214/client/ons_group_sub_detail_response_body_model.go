// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsGroupSubDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OnsGroupSubDetailResponseBodyData) *OnsGroupSubDetailResponseBody
	GetData() *OnsGroupSubDetailResponseBodyData
	SetRequestId(v string) *OnsGroupSubDetailResponseBody
	GetRequestId() *string
}

type OnsGroupSubDetailResponseBody struct {
	// The data returned.
	Data *OnsGroupSubDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 3364E875-013B-442A-BC3C-C1A84DC6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsGroupSubDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsGroupSubDetailResponseBody) GoString() string {
	return s.String()
}

func (s *OnsGroupSubDetailResponseBody) GetData() *OnsGroupSubDetailResponseBodyData {
	return s.Data
}

func (s *OnsGroupSubDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsGroupSubDetailResponseBody) SetData(v *OnsGroupSubDetailResponseBodyData) *OnsGroupSubDetailResponseBody {
	s.Data = v
	return s
}

func (s *OnsGroupSubDetailResponseBody) SetRequestId(v string) *OnsGroupSubDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsGroupSubDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsGroupSubDetailResponseBodyData struct {
	// The ID of the consumer group that you want to query.
	//
	// example:
	//
	// GID_test_group_id
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The consumption mode. Valid values:
	//
	// 	- **CLUSTERING**: the clustering consumption mode
	//
	// 	- **BROADCASTING**: the broadcasting consumption mode
	//
	// For more information about consumption modes, see [Clustering consumption and broadcasting consumption](https://help.aliyun.com/document_detail/43163.html).
	//
	// example:
	//
	// CLUSTERING
	MessageModel *string `json:"MessageModel,omitempty" xml:"MessageModel,omitempty"`
	// Indicates whether consumers in the group are online.
	//
	// example:
	//
	// true
	Online               *bool                                                  `json:"Online,omitempty" xml:"Online,omitempty"`
	SubscriptionDataList *OnsGroupSubDetailResponseBodyDataSubscriptionDataList `json:"SubscriptionDataList,omitempty" xml:"SubscriptionDataList,omitempty" type:"Struct"`
}

func (s OnsGroupSubDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OnsGroupSubDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsGroupSubDetailResponseBodyData) GetGroupId() *string {
	return s.GroupId
}

func (s *OnsGroupSubDetailResponseBodyData) GetMessageModel() *string {
	return s.MessageModel
}

func (s *OnsGroupSubDetailResponseBodyData) GetOnline() *bool {
	return s.Online
}

func (s *OnsGroupSubDetailResponseBodyData) GetSubscriptionDataList() *OnsGroupSubDetailResponseBodyDataSubscriptionDataList {
	return s.SubscriptionDataList
}

func (s *OnsGroupSubDetailResponseBodyData) SetGroupId(v string) *OnsGroupSubDetailResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *OnsGroupSubDetailResponseBodyData) SetMessageModel(v string) *OnsGroupSubDetailResponseBodyData {
	s.MessageModel = &v
	return s
}

func (s *OnsGroupSubDetailResponseBodyData) SetOnline(v bool) *OnsGroupSubDetailResponseBodyData {
	s.Online = &v
	return s
}

func (s *OnsGroupSubDetailResponseBodyData) SetSubscriptionDataList(v *OnsGroupSubDetailResponseBodyDataSubscriptionDataList) *OnsGroupSubDetailResponseBodyData {
	s.SubscriptionDataList = v
	return s
}

func (s *OnsGroupSubDetailResponseBodyData) Validate() error {
	if s.SubscriptionDataList != nil {
		if err := s.SubscriptionDataList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsGroupSubDetailResponseBodyDataSubscriptionDataList struct {
	SubscriptionDataList []*OnsGroupSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList `json:"SubscriptionDataList,omitempty" xml:"SubscriptionDataList,omitempty" type:"Repeated"`
}

func (s OnsGroupSubDetailResponseBodyDataSubscriptionDataList) String() string {
	return dara.Prettify(s)
}

func (s OnsGroupSubDetailResponseBodyDataSubscriptionDataList) GoString() string {
	return s.String()
}

func (s *OnsGroupSubDetailResponseBodyDataSubscriptionDataList) GetSubscriptionDataList() []*OnsGroupSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList {
	return s.SubscriptionDataList
}

func (s *OnsGroupSubDetailResponseBodyDataSubscriptionDataList) SetSubscriptionDataList(v []*OnsGroupSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) *OnsGroupSubDetailResponseBodyDataSubscriptionDataList {
	s.SubscriptionDataList = v
	return s
}

func (s *OnsGroupSubDetailResponseBodyDataSubscriptionDataList) Validate() error {
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

type OnsGroupSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList struct {
	SubString *string `json:"SubString,omitempty" xml:"SubString,omitempty"`
	Topic     *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsGroupSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) String() string {
	return dara.Prettify(s)
}

func (s OnsGroupSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) GoString() string {
	return s.String()
}

func (s *OnsGroupSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) GetSubString() *string {
	return s.SubString
}

func (s *OnsGroupSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) GetTopic() *string {
	return s.Topic
}

func (s *OnsGroupSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) SetSubString(v string) *OnsGroupSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList {
	s.SubString = &v
	return s
}

func (s *OnsGroupSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) SetTopic(v string) *OnsGroupSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList {
	s.Topic = &v
	return s
}

func (s *OnsGroupSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) Validate() error {
	return dara.Validate(s)
}
