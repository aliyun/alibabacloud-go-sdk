// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppOperateAction interface {
	dara.Model
	String() string
	GoString() string
	SetActionKey(v string) *AppOperateAction
	GetActionKey() *string
	SetActionText(v string) *AppOperateAction
	GetActionText() *string
	SetEnable(v bool) *AppOperateAction
	GetEnable() *bool
	SetHref(v string) *AppOperateAction
	GetHref() *string
}

type AppOperateAction struct {
	// 用于唯一标识一个操作行为
	ActionKey *string `json:"ActionKey,omitempty" xml:"ActionKey,omitempty"`
	// 用于在界面中展示操作名称
	ActionText *string `json:"ActionText,omitempty" xml:"ActionText,omitempty"`
	// 标识该操作是否可用
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// 点击操作时跳转的URL地址
	Href *string `json:"Href,omitempty" xml:"Href,omitempty"`
}

func (s AppOperateAction) String() string {
	return dara.Prettify(s)
}

func (s AppOperateAction) GoString() string {
	return s.String()
}

func (s *AppOperateAction) GetActionKey() *string {
	return s.ActionKey
}

func (s *AppOperateAction) GetActionText() *string {
	return s.ActionText
}

func (s *AppOperateAction) GetEnable() *bool {
	return s.Enable
}

func (s *AppOperateAction) GetHref() *string {
	return s.Href
}

func (s *AppOperateAction) SetActionKey(v string) *AppOperateAction {
	s.ActionKey = &v
	return s
}

func (s *AppOperateAction) SetActionText(v string) *AppOperateAction {
	s.ActionText = &v
	return s
}

func (s *AppOperateAction) SetEnable(v bool) *AppOperateAction {
	s.Enable = &v
	return s
}

func (s *AppOperateAction) SetHref(v string) *AppOperateAction {
	s.Href = &v
	return s
}

func (s *AppOperateAction) Validate() error {
	return dara.Validate(s)
}
