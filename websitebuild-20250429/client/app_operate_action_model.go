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
	// The key used to uniquely identify an action.
	//
	// example:
	//
	// BeginDialogue
	ActionKey *string `json:"ActionKey,omitempty" xml:"ActionKey,omitempty"`
	// The text used to display the action name in the user interface.
	//
	// example:
	//
	// Renew
	ActionText *string `json:"ActionText,omitempty" xml:"ActionText,omitempty"`
	// Indicates whether the action is available.
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The URL to redirect to when the action is clicked.
	//
	// example:
	//
	// https://common-buy.aliyun.com/?orderType=RENEW&commodityCode=domain_website_public_cn&instanceId=12313213
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
