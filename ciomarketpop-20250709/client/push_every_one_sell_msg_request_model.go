// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushEveryOneSellMsgRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDingIdList(v []*string) *PushEveryOneSellMsgRequest
	GetDingIdList() []*string
	SetPushMsg(v string) *PushEveryOneSellMsgRequest
	GetPushMsg() *string
	SetPushType(v string) *PushEveryOneSellMsgRequest
	GetPushType() *string
}

type PushEveryOneSellMsgRequest struct {
	// example:
	//
	// ["1234567"]
	DingIdList []*string `json:"DingIdList,omitempty" xml:"DingIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 推送内容
	PushMsg *string `json:"PushMsg,omitempty" xml:"PushMsg,omitempty"`
	// example:
	//
	// 1
	PushType *string `json:"PushType,omitempty" xml:"PushType,omitempty"`
}

func (s PushEveryOneSellMsgRequest) String() string {
	return dara.Prettify(s)
}

func (s PushEveryOneSellMsgRequest) GoString() string {
	return s.String()
}

func (s *PushEveryOneSellMsgRequest) GetDingIdList() []*string {
	return s.DingIdList
}

func (s *PushEveryOneSellMsgRequest) GetPushMsg() *string {
	return s.PushMsg
}

func (s *PushEveryOneSellMsgRequest) GetPushType() *string {
	return s.PushType
}

func (s *PushEveryOneSellMsgRequest) SetDingIdList(v []*string) *PushEveryOneSellMsgRequest {
	s.DingIdList = v
	return s
}

func (s *PushEveryOneSellMsgRequest) SetPushMsg(v string) *PushEveryOneSellMsgRequest {
	s.PushMsg = &v
	return s
}

func (s *PushEveryOneSellMsgRequest) SetPushType(v string) *PushEveryOneSellMsgRequest {
	s.PushType = &v
	return s
}

func (s *PushEveryOneSellMsgRequest) Validate() error {
	return dara.Validate(s)
}
