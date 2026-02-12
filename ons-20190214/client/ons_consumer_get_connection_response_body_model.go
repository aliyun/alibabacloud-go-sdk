// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsConsumerGetConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OnsConsumerGetConnectionResponseBodyData) *OnsConsumerGetConnectionResponseBody
	GetData() *OnsConsumerGetConnectionResponseBodyData
	SetRequestId(v string) *OnsConsumerGetConnectionResponseBody
	GetRequestId() *string
}

type OnsConsumerGetConnectionResponseBody struct {
	// The data returned.
	Data *OnsConsumerGetConnectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// DE4140C7-F42D-473D-A5FF-B1E31692****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsConsumerGetConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerGetConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *OnsConsumerGetConnectionResponseBody) GetData() *OnsConsumerGetConnectionResponseBodyData {
	return s.Data
}

func (s *OnsConsumerGetConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsConsumerGetConnectionResponseBody) SetData(v *OnsConsumerGetConnectionResponseBodyData) *OnsConsumerGetConnectionResponseBody {
	s.Data = v
	return s
}

func (s *OnsConsumerGetConnectionResponseBody) SetRequestId(v string) *OnsConsumerGetConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsConsumerGetConnectionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsConsumerGetConnectionResponseBodyData struct {
	ConnectionList *OnsConsumerGetConnectionResponseBodyDataConnectionList `json:"ConnectionList,omitempty" xml:"ConnectionList,omitempty" type:"Struct"`
	MessageModel   *string                                                 `json:"MessageModel,omitempty" xml:"MessageModel,omitempty"`
}

func (s OnsConsumerGetConnectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerGetConnectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsConsumerGetConnectionResponseBodyData) GetConnectionList() *OnsConsumerGetConnectionResponseBodyDataConnectionList {
	return s.ConnectionList
}

func (s *OnsConsumerGetConnectionResponseBodyData) GetMessageModel() *string {
	return s.MessageModel
}

func (s *OnsConsumerGetConnectionResponseBodyData) SetConnectionList(v *OnsConsumerGetConnectionResponseBodyDataConnectionList) *OnsConsumerGetConnectionResponseBodyData {
	s.ConnectionList = v
	return s
}

func (s *OnsConsumerGetConnectionResponseBodyData) SetMessageModel(v string) *OnsConsumerGetConnectionResponseBodyData {
	s.MessageModel = &v
	return s
}

func (s *OnsConsumerGetConnectionResponseBodyData) Validate() error {
	if s.ConnectionList != nil {
		if err := s.ConnectionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsConsumerGetConnectionResponseBodyDataConnectionList struct {
	ConnectionDo []*OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo `json:"ConnectionDo,omitempty" xml:"ConnectionDo,omitempty" type:"Repeated"`
}

func (s OnsConsumerGetConnectionResponseBodyDataConnectionList) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerGetConnectionResponseBodyDataConnectionList) GoString() string {
	return s.String()
}

func (s *OnsConsumerGetConnectionResponseBodyDataConnectionList) GetConnectionDo() []*OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo {
	return s.ConnectionDo
}

func (s *OnsConsumerGetConnectionResponseBodyDataConnectionList) SetConnectionDo(v []*OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo) *OnsConsumerGetConnectionResponseBodyDataConnectionList {
	s.ConnectionDo = v
	return s
}

func (s *OnsConsumerGetConnectionResponseBodyDataConnectionList) Validate() error {
	if s.ConnectionDo != nil {
		for _, item := range s.ConnectionDo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo struct {
	ClientAddr *string `json:"ClientAddr,omitempty" xml:"ClientAddr,omitempty"`
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	Language   *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Version    *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo) GoString() string {
	return s.String()
}

func (s *OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo) GetClientAddr() *string {
	return s.ClientAddr
}

func (s *OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo) GetClientId() *string {
	return s.ClientId
}

func (s *OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo) GetLanguage() *string {
	return s.Language
}

func (s *OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo) GetVersion() *string {
	return s.Version
}

func (s *OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo) SetClientAddr(v string) *OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo {
	s.ClientAddr = &v
	return s
}

func (s *OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo) SetClientId(v string) *OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo {
	s.ClientId = &v
	return s
}

func (s *OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo) SetLanguage(v string) *OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo {
	s.Language = &v
	return s
}

func (s *OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo) SetVersion(v string) *OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo {
	s.Version = &v
	return s
}

func (s *OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo) Validate() error {
	return dara.Validate(s)
}
