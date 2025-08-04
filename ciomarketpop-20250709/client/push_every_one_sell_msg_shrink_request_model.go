// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushEveryOneSellMsgShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDingIdListShrink(v string) *PushEveryOneSellMsgShrinkRequest
	GetDingIdListShrink() *string
	SetPushMsg(v string) *PushEveryOneSellMsgShrinkRequest
	GetPushMsg() *string
	SetPushType(v string) *PushEveryOneSellMsgShrinkRequest
	GetPushType() *string
}

type PushEveryOneSellMsgShrinkRequest struct {
	// example:
	//
	// ["1234567"]
	DingIdListShrink *string `json:"DingIdList,omitempty" xml:"DingIdList,omitempty"`
	// example:
	//
	// 推送内容
	PushMsg *string `json:"PushMsg,omitempty" xml:"PushMsg,omitempty"`
	// example:
	//
	// 1
	PushType *string `json:"PushType,omitempty" xml:"PushType,omitempty"`
}

func (s PushEveryOneSellMsgShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PushEveryOneSellMsgShrinkRequest) GoString() string {
	return s.String()
}

func (s *PushEveryOneSellMsgShrinkRequest) GetDingIdListShrink() *string {
	return s.DingIdListShrink
}

func (s *PushEveryOneSellMsgShrinkRequest) GetPushMsg() *string {
	return s.PushMsg
}

func (s *PushEveryOneSellMsgShrinkRequest) GetPushType() *string {
	return s.PushType
}

func (s *PushEveryOneSellMsgShrinkRequest) SetDingIdListShrink(v string) *PushEveryOneSellMsgShrinkRequest {
	s.DingIdListShrink = &v
	return s
}

func (s *PushEveryOneSellMsgShrinkRequest) SetPushMsg(v string) *PushEveryOneSellMsgShrinkRequest {
	s.PushMsg = &v
	return s
}

func (s *PushEveryOneSellMsgShrinkRequest) SetPushType(v string) *PushEveryOneSellMsgShrinkRequest {
	s.PushType = &v
	return s
}

func (s *PushEveryOneSellMsgShrinkRequest) Validate() error {
	return dara.Validate(s)
}
