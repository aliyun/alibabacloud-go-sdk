// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsMessagePageQueryByTopicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMsgFoundDo(v *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo) *OnsMessagePageQueryByTopicResponseBody
	GetMsgFoundDo() *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo
	SetRequestId(v string) *OnsMessagePageQueryByTopicResponseBody
	GetRequestId() *string
}

type OnsMessagePageQueryByTopicResponseBody struct {
	// The information about the message that is queried.
	MsgFoundDo *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo `json:"MsgFoundDo,omitempty" xml:"MsgFoundDo,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 5DC2A47E-2B31-4722-96C8-FA59C9*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsMessagePageQueryByTopicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsMessagePageQueryByTopicResponseBody) GoString() string {
	return s.String()
}

func (s *OnsMessagePageQueryByTopicResponseBody) GetMsgFoundDo() *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo {
	return s.MsgFoundDo
}

func (s *OnsMessagePageQueryByTopicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsMessagePageQueryByTopicResponseBody) SetMsgFoundDo(v *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo) *OnsMessagePageQueryByTopicResponseBody {
	s.MsgFoundDo = v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBody) SetRequestId(v string) *OnsMessagePageQueryByTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBody) Validate() error {
	if s.MsgFoundDo != nil {
		if err := s.MsgFoundDo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsMessagePageQueryByTopicResponseBodyMsgFoundDo struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The total number of returned pages.
	//
	// example:
	//
	// 400
	MaxPageCount *int64                                                        `json:"MaxPageCount,omitempty" xml:"MaxPageCount,omitempty"`
	MsgFoundList *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList `json:"MsgFoundList,omitempty" xml:"MsgFoundList,omitempty" type:"Struct"`
	// The ID of the query task. The first time you call this operation to query the messages that are sent to a specified consumer group within a specified time range, this parameter is returned. You can use the task ID to query the details of messages on other returned pages.
	//
	// example:
	//
	// 0BC1310300002A9F000021E4D7A48346
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s OnsMessagePageQueryByTopicResponseBodyMsgFoundDo) String() string {
	return dara.Prettify(s)
}

func (s OnsMessagePageQueryByTopicResponseBodyMsgFoundDo) GoString() string {
	return s.String()
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo) GetMaxPageCount() *int64 {
	return s.MaxPageCount
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo) GetMsgFoundList() *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList {
	return s.MsgFoundList
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo) GetTaskId() *string {
	return s.TaskId
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo) SetCurrentPage(v int64) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo {
	s.CurrentPage = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo) SetMaxPageCount(v int64) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo {
	s.MaxPageCount = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo) SetMsgFoundList(v *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo {
	s.MsgFoundList = v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo) SetTaskId(v string) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo {
	s.TaskId = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo) Validate() error {
	if s.MsgFoundList != nil {
		if err := s.MsgFoundList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList struct {
	OnsRestMessageDo []*OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo `json:"OnsRestMessageDo,omitempty" xml:"OnsRestMessageDo,omitempty" type:"Repeated"`
}

func (s OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList) String() string {
	return dara.Prettify(s)
}

func (s OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList) GoString() string {
	return s.String()
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList) GetOnsRestMessageDo() []*OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	return s.OnsRestMessageDo
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList) SetOnsRestMessageDo(v []*OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList {
	s.OnsRestMessageDo = v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList) Validate() error {
	if s.OnsRestMessageDo != nil {
		for _, item := range s.OnsRestMessageDo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo struct {
	BodyCRC        *int32                                                                                    `json:"BodyCRC,omitempty" xml:"BodyCRC,omitempty"`
	BornHost       *string                                                                                   `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
	BornTimestamp  *int64                                                                                    `json:"BornTimestamp,omitempty" xml:"BornTimestamp,omitempty"`
	InstanceId     *string                                                                                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId          *string                                                                                   `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	PropertyList   *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Struct"`
	ReconsumeTimes *int32                                                                                    `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	StoreHost      *string                                                                                   `json:"StoreHost,omitempty" xml:"StoreHost,omitempty"`
	StoreSize      *int32                                                                                    `json:"StoreSize,omitempty" xml:"StoreSize,omitempty"`
	StoreTimestamp *int64                                                                                    `json:"StoreTimestamp,omitempty" xml:"StoreTimestamp,omitempty"`
	Topic          *string                                                                                   `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) String() string {
	return dara.Prettify(s)
}

func (s OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) GoString() string {
	return s.String()
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) GetBodyCRC() *int32 {
	return s.BodyCRC
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) GetBornHost() *string {
	return s.BornHost
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) GetBornTimestamp() *int64 {
	return s.BornTimestamp
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) GetMsgId() *string {
	return s.MsgId
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) GetPropertyList() *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList {
	return s.PropertyList
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) GetReconsumeTimes() *int32 {
	return s.ReconsumeTimes
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) GetStoreHost() *string {
	return s.StoreHost
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) GetStoreSize() *int32 {
	return s.StoreSize
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) GetStoreTimestamp() *int64 {
	return s.StoreTimestamp
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) GetTopic() *string {
	return s.Topic
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetBodyCRC(v int32) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.BodyCRC = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetBornHost(v string) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.BornHost = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetBornTimestamp(v int64) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.BornTimestamp = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetInstanceId(v string) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.InstanceId = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetMsgId(v string) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.MsgId = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetPropertyList(v *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.PropertyList = v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetReconsumeTimes(v int32) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.ReconsumeTimes = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetStoreHost(v string) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.StoreHost = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetStoreSize(v int32) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.StoreSize = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetStoreTimestamp(v int64) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.StoreTimestamp = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetTopic(v string) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.Topic = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) Validate() error {
	if s.PropertyList != nil {
		if err := s.PropertyList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList struct {
	MessageProperty []*OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty `json:"MessageProperty,omitempty" xml:"MessageProperty,omitempty" type:"Repeated"`
}

func (s OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList) String() string {
	return dara.Prettify(s)
}

func (s OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList) GoString() string {
	return s.String()
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList) GetMessageProperty() []*OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty {
	return s.MessageProperty
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList) SetMessageProperty(v []*OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList {
	s.MessageProperty = v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList) Validate() error {
	if s.MessageProperty != nil {
		for _, item := range s.MessageProperty {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty) String() string {
	return dara.Prettify(s)
}

func (s OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty) GoString() string {
	return s.String()
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty) GetName() *string {
	return s.Name
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty) GetValue() *string {
	return s.Value
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty) SetName(v string) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty {
	s.Name = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty) SetValue(v string) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty {
	s.Value = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty) Validate() error {
	return dara.Validate(s)
}
