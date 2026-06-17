// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannels(v *PutContactRequestChannels) *PutContactRequest
	GetChannels() *PutContactRequestChannels
	SetContactName(v string) *PutContactRequest
	GetContactName() *string
	SetDescribe(v string) *PutContactRequest
	GetDescribe() *string
	SetLang(v string) *PutContactRequest
	GetLang() *string
}

type PutContactRequest struct {
	Channels *PutContactRequestChannels `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Struct"`
	// The name of the alert contact.
	//
	// This parameter is required.
	//
	// example:
	//
	// Alice
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// The description of the alert contact.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS_Instance
	Describe *string `json:"Describe,omitempty" xml:"Describe,omitempty"`
	// The language in which alert notifications are sent. Valid values:
	//
	// - zh-cn: simplified Chinese.
	//
	// - en: English.
	//
	// > If this parameter is not specified, the language is automatically determined based on the region of the account.
	//
	// example:
	//
	// zh-cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s PutContactRequest) String() string {
	return dara.Prettify(s)
}

func (s PutContactRequest) GoString() string {
	return s.String()
}

func (s *PutContactRequest) GetChannels() *PutContactRequestChannels {
	return s.Channels
}

func (s *PutContactRequest) GetContactName() *string {
	return s.ContactName
}

func (s *PutContactRequest) GetDescribe() *string {
	return s.Describe
}

func (s *PutContactRequest) GetLang() *string {
	return s.Lang
}

func (s *PutContactRequest) SetChannels(v *PutContactRequestChannels) *PutContactRequest {
	s.Channels = v
	return s
}

func (s *PutContactRequest) SetContactName(v string) *PutContactRequest {
	s.ContactName = &v
	return s
}

func (s *PutContactRequest) SetDescribe(v string) *PutContactRequest {
	s.Describe = &v
	return s
}

func (s *PutContactRequest) SetLang(v string) *PutContactRequest {
	s.Lang = &v
	return s
}

func (s *PutContactRequest) Validate() error {
	if s.Channels != nil {
		if err := s.Channels.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PutContactRequestChannels struct {
	// 旺旺联系人。
	//
	// <props="china">手机号码、邮箱、钉钉机器人和旺旺最少添加一种联系方式。
	//
	// <props="intl">邮箱和钉钉机器人最少添加一种联系方式。
	//
	// <props="partner">邮箱和钉钉机器人最少添加一种联系方式。
	//
	// example:
	//
	// Jim
	AliIM *string `json:"AliIM,omitempty" xml:"AliIM,omitempty"`
	// 钉钉机器人。
	//
	// <props="china">手机号码、邮箱、钉钉机器人和旺旺最少添加一种联系方式。
	//
	// <props="intl">邮箱和钉钉机器人最少添加一种联系方式。
	//
	// <props="partner">邮箱和钉钉机器人最少添加一种联系方式。
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=7d49515e8ebf21106a80a9cc4bb3d247771305d52fb15d6201234565****
	DingWebHook *string `json:"DingWebHook,omitempty" xml:"DingWebHook,omitempty"`
	// Email地址。Email会收到一个激活链接， 激活之后您才会被加入到联系人中。
	//
	// <props="china">手机号码、邮箱、钉钉机器人和旺旺最少添加一种联系方式。
	//
	// <props="intl">邮箱和钉钉机器人最少添加一种联系方式。
	//
	// <props="partner">邮箱和钉钉机器人最少添加一种联系方式。
	//
	// example:
	//
	// test@aliyun.com
	Mail *string `json:"Mail,omitempty" xml:"Mail,omitempty"`
	// 手机号码。手机号码会收到一个激活链接， 激活之后您才会被加入到联系人中。
	//
	// <props="china">手机号码、邮箱、钉钉机器人和旺旺最少添加一种联系方式。
	//
	// <props="intl">邮箱和钉钉机器人最少添加一种联系方式。
	//
	// <props="partner">邮箱和钉钉机器人最少添加一种联系方式。
	//
	// example:
	//
	// 1333333****
	SMS *string `json:"SMS,omitempty" xml:"SMS,omitempty"`
}

func (s PutContactRequestChannels) String() string {
	return dara.Prettify(s)
}

func (s PutContactRequestChannels) GoString() string {
	return s.String()
}

func (s *PutContactRequestChannels) GetAliIM() *string {
	return s.AliIM
}

func (s *PutContactRequestChannels) GetDingWebHook() *string {
	return s.DingWebHook
}

func (s *PutContactRequestChannels) GetMail() *string {
	return s.Mail
}

func (s *PutContactRequestChannels) GetSMS() *string {
	return s.SMS
}

func (s *PutContactRequestChannels) SetAliIM(v string) *PutContactRequestChannels {
	s.AliIM = &v
	return s
}

func (s *PutContactRequestChannels) SetDingWebHook(v string) *PutContactRequestChannels {
	s.DingWebHook = &v
	return s
}

func (s *PutContactRequestChannels) SetMail(v string) *PutContactRequestChannels {
	s.Mail = &v
	return s
}

func (s *PutContactRequestChannels) SetSMS(v string) *PutContactRequestChannels {
	s.SMS = &v
	return s
}

func (s *PutContactRequestChannels) Validate() error {
	return dara.Validate(s)
}
