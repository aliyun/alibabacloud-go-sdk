// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsMessageGetByKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OnsMessageGetByKeyResponseBodyData) *OnsMessageGetByKeyResponseBody
	GetData() *OnsMessageGetByKeyResponseBodyData
	SetRequestId(v string) *OnsMessageGetByKeyResponseBody
	GetRequestId() *string
}

type OnsMessageGetByKeyResponseBody struct {
	Data *OnsMessageGetByKeyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// A07E3902-B92E-44A6-B6C5-6AA111111****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsMessageGetByKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsMessageGetByKeyResponseBody) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByKeyResponseBody) GetData() *OnsMessageGetByKeyResponseBodyData {
	return s.Data
}

func (s *OnsMessageGetByKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsMessageGetByKeyResponseBody) SetData(v *OnsMessageGetByKeyResponseBodyData) *OnsMessageGetByKeyResponseBody {
	s.Data = v
	return s
}

func (s *OnsMessageGetByKeyResponseBody) SetRequestId(v string) *OnsMessageGetByKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsMessageGetByKeyResponseBodyData struct {
	OnsRestMessageDo []*OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo `json:"OnsRestMessageDo,omitempty" xml:"OnsRestMessageDo,omitempty" type:"Repeated"`
}

func (s OnsMessageGetByKeyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OnsMessageGetByKeyResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByKeyResponseBodyData) GetOnsRestMessageDo() []*OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo {
	return s.OnsRestMessageDo
}

func (s *OnsMessageGetByKeyResponseBodyData) SetOnsRestMessageDo(v []*OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) *OnsMessageGetByKeyResponseBodyData {
	s.OnsRestMessageDo = v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyData) Validate() error {
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

type OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo struct {
	BodyCRC        *int32                                                          `json:"BodyCRC,omitempty" xml:"BodyCRC,omitempty"`
	BornHost       *string                                                         `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
	BornTimestamp  *int64                                                          `json:"BornTimestamp,omitempty" xml:"BornTimestamp,omitempty"`
	InstanceId     *string                                                         `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId          *string                                                         `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	PropertyList   *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Struct"`
	ReconsumeTimes *int32                                                          `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	StoreHost      *string                                                         `json:"StoreHost,omitempty" xml:"StoreHost,omitempty"`
	StoreSize      *int32                                                          `json:"StoreSize,omitempty" xml:"StoreSize,omitempty"`
	StoreTimestamp *int64                                                          `json:"StoreTimestamp,omitempty" xml:"StoreTimestamp,omitempty"`
	Topic          *string                                                         `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) String() string {
	return dara.Prettify(s)
}

func (s OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) GetBodyCRC() *int32 {
	return s.BodyCRC
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) GetBornHost() *string {
	return s.BornHost
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) GetBornTimestamp() *int64 {
	return s.BornTimestamp
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) GetMsgId() *string {
	return s.MsgId
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) GetPropertyList() *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyList {
	return s.PropertyList
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) GetReconsumeTimes() *int32 {
	return s.ReconsumeTimes
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) GetStoreHost() *string {
	return s.StoreHost
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) GetStoreSize() *int32 {
	return s.StoreSize
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) GetStoreTimestamp() *int64 {
	return s.StoreTimestamp
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) GetTopic() *string {
	return s.Topic
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) SetBodyCRC(v int32) *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo {
	s.BodyCRC = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) SetBornHost(v string) *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo {
	s.BornHost = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) SetBornTimestamp(v int64) *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo {
	s.BornTimestamp = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) SetInstanceId(v string) *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo {
	s.InstanceId = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) SetMsgId(v string) *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo {
	s.MsgId = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) SetPropertyList(v *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyList) *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo {
	s.PropertyList = v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) SetReconsumeTimes(v int32) *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo {
	s.ReconsumeTimes = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) SetStoreHost(v string) *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo {
	s.StoreHost = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) SetStoreSize(v int32) *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo {
	s.StoreSize = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) SetStoreTimestamp(v int64) *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo {
	s.StoreTimestamp = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) SetTopic(v string) *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo {
	s.Topic = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) Validate() error {
	if s.PropertyList != nil {
		if err := s.PropertyList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyList struct {
	MessageProperty []*OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyListMessageProperty `json:"MessageProperty,omitempty" xml:"MessageProperty,omitempty" type:"Repeated"`
}

func (s OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyList) String() string {
	return dara.Prettify(s)
}

func (s OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyList) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyList) GetMessageProperty() []*OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyListMessageProperty {
	return s.MessageProperty
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyList) SetMessageProperty(v []*OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyListMessageProperty) *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyList {
	s.MessageProperty = v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyList) Validate() error {
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

type OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyListMessageProperty struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyListMessageProperty) String() string {
	return dara.Prettify(s)
}

func (s OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyListMessageProperty) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyListMessageProperty) GetName() *string {
	return s.Name
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyListMessageProperty) GetValue() *string {
	return s.Value
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyListMessageProperty) SetName(v string) *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyListMessageProperty {
	s.Name = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyListMessageProperty) SetValue(v string) *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyListMessageProperty {
	s.Value = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyListMessageProperty) Validate() error {
	return dara.Validate(s)
}
