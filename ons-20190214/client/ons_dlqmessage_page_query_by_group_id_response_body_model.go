// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsDLQMessagePageQueryByGroupIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMsgFoundDo(v *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo) *OnsDLQMessagePageQueryByGroupIdResponseBody
	GetMsgFoundDo() *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo
	SetRequestId(v string) *OnsDLQMessagePageQueryByGroupIdResponseBody
	GetRequestId() *string
}

type OnsDLQMessagePageQueryByGroupIdResponseBody struct {
	// The information about dead-letter messages that are queried.
	MsgFoundDo *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo `json:"MsgFoundDo,omitempty" xml:"MsgFoundDo,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// B00CD3C8-D81E-4A41-85E2-38F19252****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBody) GetMsgFoundDo() *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo {
	return s.MsgFoundDo
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBody) SetMsgFoundDo(v *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo) *OnsDLQMessagePageQueryByGroupIdResponseBody {
	s.MsgFoundDo = v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBody) SetRequestId(v string) *OnsDLQMessagePageQueryByGroupIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBody) Validate() error {
	if s.MsgFoundDo != nil {
		if err := s.MsgFoundDo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo struct {
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
	MaxPageCount *int64                                                             `json:"MaxPageCount,omitempty" xml:"MaxPageCount,omitempty"`
	MsgFoundList *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList `json:"MsgFoundList,omitempty" xml:"MsgFoundList,omitempty" type:"Struct"`
	// The ID of the query task. The first time you call this operation to query the dead-letter messages that are sent to a specified consumer group within a specified time range, this parameter is returned. You can use the task ID to query the details of dead-letter messages on other returned pages.
	//
	// example:
	//
	// 0BC1310300002A9F000021E4D7A48346
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo) String() string {
	return dara.Prettify(s)
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo) GoString() string {
	return s.String()
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo) GetMaxPageCount() *int64 {
	return s.MaxPageCount
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo) GetMsgFoundList() *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList {
	return s.MsgFoundList
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo) GetTaskId() *string {
	return s.TaskId
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo) SetCurrentPage(v int64) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo {
	s.CurrentPage = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo) SetMaxPageCount(v int64) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo {
	s.MaxPageCount = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo) SetMsgFoundList(v *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo {
	s.MsgFoundList = v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo) SetTaskId(v string) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo {
	s.TaskId = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo) Validate() error {
	if s.MsgFoundList != nil {
		if err := s.MsgFoundList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList struct {
	OnsRestMessageDo []*OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo `json:"OnsRestMessageDo,omitempty" xml:"OnsRestMessageDo,omitempty" type:"Repeated"`
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList) String() string {
	return dara.Prettify(s)
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList) GoString() string {
	return s.String()
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList) GetOnsRestMessageDo() []*OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	return s.OnsRestMessageDo
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList) SetOnsRestMessageDo(v []*OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList {
	s.OnsRestMessageDo = v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList) Validate() error {
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

type OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo struct {
	BodyCRC        *int32                                                                                         `json:"BodyCRC,omitempty" xml:"BodyCRC,omitempty"`
	BornHost       *string                                                                                        `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
	BornTimestamp  *int64                                                                                         `json:"BornTimestamp,omitempty" xml:"BornTimestamp,omitempty"`
	InstanceId     *string                                                                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId          *string                                                                                        `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	PropertyList   *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Struct"`
	ReconsumeTimes *int32                                                                                         `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	StoreHost      *string                                                                                        `json:"StoreHost,omitempty" xml:"StoreHost,omitempty"`
	StoreSize      *int32                                                                                         `json:"StoreSize,omitempty" xml:"StoreSize,omitempty"`
	StoreTimestamp *int64                                                                                         `json:"StoreTimestamp,omitempty" xml:"StoreTimestamp,omitempty"`
	Topic          *string                                                                                        `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) String() string {
	return dara.Prettify(s)
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) GoString() string {
	return s.String()
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) GetBodyCRC() *int32 {
	return s.BodyCRC
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) GetBornHost() *string {
	return s.BornHost
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) GetBornTimestamp() *int64 {
	return s.BornTimestamp
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) GetMsgId() *string {
	return s.MsgId
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) GetPropertyList() *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList {
	return s.PropertyList
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) GetReconsumeTimes() *int32 {
	return s.ReconsumeTimes
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) GetStoreHost() *string {
	return s.StoreHost
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) GetStoreSize() *int32 {
	return s.StoreSize
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) GetStoreTimestamp() *int64 {
	return s.StoreTimestamp
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) GetTopic() *string {
	return s.Topic
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetBodyCRC(v int32) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.BodyCRC = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetBornHost(v string) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.BornHost = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetBornTimestamp(v int64) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.BornTimestamp = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetInstanceId(v string) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.InstanceId = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetMsgId(v string) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.MsgId = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetPropertyList(v *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.PropertyList = v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetReconsumeTimes(v int32) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.ReconsumeTimes = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetStoreHost(v string) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.StoreHost = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetStoreSize(v int32) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.StoreSize = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetStoreTimestamp(v int64) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.StoreTimestamp = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetTopic(v string) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.Topic = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) Validate() error {
	if s.PropertyList != nil {
		if err := s.PropertyList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList struct {
	MessageProperty []*OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty `json:"MessageProperty,omitempty" xml:"MessageProperty,omitempty" type:"Repeated"`
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList) String() string {
	return dara.Prettify(s)
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList) GoString() string {
	return s.String()
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList) GetMessageProperty() []*OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty {
	return s.MessageProperty
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList) SetMessageProperty(v []*OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList {
	s.MessageProperty = v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList) Validate() error {
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

type OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty) String() string {
	return dara.Prettify(s)
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty) GoString() string {
	return s.String()
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty) GetName() *string {
	return s.Name
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty) GetValue() *string {
	return s.Value
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty) SetName(v string) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty {
	s.Name = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty) SetValue(v string) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty {
	s.Value = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty) Validate() error {
	return dara.Validate(s)
}
