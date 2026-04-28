// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReceivedMsg interface {
	dara.Model
	String() string
	GoString() string
	SetHasRead(v bool) *ReceivedMsg
	GetHasRead() *bool
	SetMsgCategory(v string) *ReceivedMsg
	GetMsgCategory() *string
	SetMsgContent(v *ReceivedMsgMsgContent) *ReceivedMsg
	GetMsgContent() *ReceivedMsgMsgContent
	SetMsgId(v string) *ReceivedMsg
	GetMsgId() *string
	SetMsgSubCategory(v string) *ReceivedMsg
	GetMsgSubCategory() *string
	SetMsgType(v string) *ReceivedMsg
	GetMsgType() *string
	SetPublishAt(v int64) *ReceivedMsg
	GetPublishAt() *int64
	SetReadAt(v int64) *ReceivedMsg
	GetReadAt() *int64
}

type ReceivedMsg struct {
	// example:
	//
	// false
	HasRead *bool `json:"has_read,omitempty" xml:"has_read,omitempty"`
	// example:
	//
	// system
	MsgCategory *string                `json:"msg_category,omitempty" xml:"msg_category,omitempty"`
	MsgContent  *ReceivedMsgMsgContent `json:"msg_content,omitempty" xml:"msg_content,omitempty" type:"Struct"`
	// example:
	//
	// 50d6f2aaa16525c7d053998e48fac265962f585f
	MsgId *string `json:"msg_id,omitempty" xml:"msg_id,omitempty"`
	// example:
	//
	// change_user_setting
	MsgSubCategory *string `json:"msg_sub_category,omitempty" xml:"msg_sub_category,omitempty"`
	// example:
	//
	// edit_user
	MsgType *string `json:"msg_type,omitempty" xml:"msg_type,omitempty"`
	// example:
	//
	// 1716363191123
	PublishAt *int64 `json:"publish_at,omitempty" xml:"publish_at,omitempty"`
	// example:
	//
	// 1716363191123
	ReadAt *int64 `json:"read_at,omitempty" xml:"read_at,omitempty"`
}

func (s ReceivedMsg) String() string {
	return dara.Prettify(s)
}

func (s ReceivedMsg) GoString() string {
	return s.String()
}

func (s *ReceivedMsg) GetHasRead() *bool {
	return s.HasRead
}

func (s *ReceivedMsg) GetMsgCategory() *string {
	return s.MsgCategory
}

func (s *ReceivedMsg) GetMsgContent() *ReceivedMsgMsgContent {
	return s.MsgContent
}

func (s *ReceivedMsg) GetMsgId() *string {
	return s.MsgId
}

func (s *ReceivedMsg) GetMsgSubCategory() *string {
	return s.MsgSubCategory
}

func (s *ReceivedMsg) GetMsgType() *string {
	return s.MsgType
}

func (s *ReceivedMsg) GetPublishAt() *int64 {
	return s.PublishAt
}

func (s *ReceivedMsg) GetReadAt() *int64 {
	return s.ReadAt
}

func (s *ReceivedMsg) SetHasRead(v bool) *ReceivedMsg {
	s.HasRead = &v
	return s
}

func (s *ReceivedMsg) SetMsgCategory(v string) *ReceivedMsg {
	s.MsgCategory = &v
	return s
}

func (s *ReceivedMsg) SetMsgContent(v *ReceivedMsgMsgContent) *ReceivedMsg {
	s.MsgContent = v
	return s
}

func (s *ReceivedMsg) SetMsgId(v string) *ReceivedMsg {
	s.MsgId = &v
	return s
}

func (s *ReceivedMsg) SetMsgSubCategory(v string) *ReceivedMsg {
	s.MsgSubCategory = &v
	return s
}

func (s *ReceivedMsg) SetMsgType(v string) *ReceivedMsg {
	s.MsgType = &v
	return s
}

func (s *ReceivedMsg) SetPublishAt(v int64) *ReceivedMsg {
	s.PublishAt = &v
	return s
}

func (s *ReceivedMsg) SetReadAt(v int64) *ReceivedMsg {
	s.ReadAt = &v
	return s
}

func (s *ReceivedMsg) Validate() error {
	if s.MsgContent != nil {
		if err := s.MsgContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReceivedMsgMsgContent struct {
	MsgData map[string]interface{} `json:"msg_data,omitempty" xml:"msg_data,omitempty"`
}

func (s ReceivedMsgMsgContent) String() string {
	return dara.Prettify(s)
}

func (s ReceivedMsgMsgContent) GoString() string {
	return s.String()
}

func (s *ReceivedMsgMsgContent) GetMsgData() map[string]interface{} {
	return s.MsgData
}

func (s *ReceivedMsgMsgContent) SetMsgData(v map[string]interface{}) *ReceivedMsgMsgContent {
	s.MsgData = v
	return s
}

func (s *ReceivedMsgMsgContent) Validate() error {
	return dara.Validate(s)
}
