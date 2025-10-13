// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMergeRobot interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *MergeRobot
	GetCreateTime() *string
	SetExtend(v *MergeRobotExtend) *MergeRobot
	GetExtend() *MergeRobotExtend
	SetGmtModified(v string) *MergeRobot
	GetGmtModified() *string
	SetIdentifier(v string) *MergeRobot
	GetIdentifier() *string
	SetLang(v string) *MergeRobot
	GetLang() *string
	SetName(v string) *MergeRobot
	GetName() *string
	SetSource(v string) *MergeRobot
	GetSource() *string
	SetType(v string) *MergeRobot
	GetType() *string
	SetWebhook(v string) *MergeRobot
	GetWebhook() *string
}

type MergeRobot struct {
	CreateTime  *string           `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Extend      *MergeRobotExtend `json:"extend,omitempty" xml:"extend,omitempty" type:"Struct"`
	GmtModified *string           `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Identifier  *string           `json:"identifier,omitempty" xml:"identifier,omitempty"`
	Lang        *string           `json:"lang,omitempty" xml:"lang,omitempty"`
	Name        *string           `json:"name,omitempty" xml:"name,omitempty"`
	Source      *string           `json:"source,omitempty" xml:"source,omitempty"`
	Type        *string           `json:"type,omitempty" xml:"type,omitempty"`
	Webhook     *string           `json:"webhook,omitempty" xml:"webhook,omitempty"`
}

func (s MergeRobot) String() string {
	return dara.Prettify(s)
}

func (s MergeRobot) GoString() string {
	return s.String()
}

func (s *MergeRobot) GetCreateTime() *string {
	return s.CreateTime
}

func (s *MergeRobot) GetExtend() *MergeRobotExtend {
	return s.Extend
}

func (s *MergeRobot) GetGmtModified() *string {
	return s.GmtModified
}

func (s *MergeRobot) GetIdentifier() *string {
	return s.Identifier
}

func (s *MergeRobot) GetLang() *string {
	return s.Lang
}

func (s *MergeRobot) GetName() *string {
	return s.Name
}

func (s *MergeRobot) GetSource() *string {
	return s.Source
}

func (s *MergeRobot) GetType() *string {
	return s.Type
}

func (s *MergeRobot) GetWebhook() *string {
	return s.Webhook
}

func (s *MergeRobot) SetCreateTime(v string) *MergeRobot {
	s.CreateTime = &v
	return s
}

func (s *MergeRobot) SetExtend(v *MergeRobotExtend) *MergeRobot {
	s.Extend = v
	return s
}

func (s *MergeRobot) SetGmtModified(v string) *MergeRobot {
	s.GmtModified = &v
	return s
}

func (s *MergeRobot) SetIdentifier(v string) *MergeRobot {
	s.Identifier = &v
	return s
}

func (s *MergeRobot) SetLang(v string) *MergeRobot {
	s.Lang = &v
	return s
}

func (s *MergeRobot) SetName(v string) *MergeRobot {
	s.Name = &v
	return s
}

func (s *MergeRobot) SetSource(v string) *MergeRobot {
	s.Source = &v
	return s
}

func (s *MergeRobot) SetType(v string) *MergeRobot {
	s.Type = &v
	return s
}

func (s *MergeRobot) SetWebhook(v string) *MergeRobot {
	s.Webhook = &v
	return s
}

func (s *MergeRobot) Validate() error {
	if s.Extend != nil {
		if err := s.Extend.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MergeRobotExtend struct {
	CardTemplate   *string `json:"cardTemplate,omitempty" xml:"cardTemplate,omitempty"`
	DailyNoc       *bool   `json:"dailyNoc,omitempty" xml:"dailyNoc,omitempty"`
	DailyNocTime   *string `json:"dailyNocTime,omitempty" xml:"dailyNocTime,omitempty"`
	DingSignKey    *string `json:"dingSignKey,omitempty" xml:"dingSignKey,omitempty"`
	EnableOutgoing *bool   `json:"enableOutgoing,omitempty" xml:"enableOutgoing,omitempty"`
	Token          *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s MergeRobotExtend) String() string {
	return dara.Prettify(s)
}

func (s MergeRobotExtend) GoString() string {
	return s.String()
}

func (s *MergeRobotExtend) GetCardTemplate() *string {
	return s.CardTemplate
}

func (s *MergeRobotExtend) GetDailyNoc() *bool {
	return s.DailyNoc
}

func (s *MergeRobotExtend) GetDailyNocTime() *string {
	return s.DailyNocTime
}

func (s *MergeRobotExtend) GetDingSignKey() *string {
	return s.DingSignKey
}

func (s *MergeRobotExtend) GetEnableOutgoing() *bool {
	return s.EnableOutgoing
}

func (s *MergeRobotExtend) GetToken() *string {
	return s.Token
}

func (s *MergeRobotExtend) SetCardTemplate(v string) *MergeRobotExtend {
	s.CardTemplate = &v
	return s
}

func (s *MergeRobotExtend) SetDailyNoc(v bool) *MergeRobotExtend {
	s.DailyNoc = &v
	return s
}

func (s *MergeRobotExtend) SetDailyNocTime(v string) *MergeRobotExtend {
	s.DailyNocTime = &v
	return s
}

func (s *MergeRobotExtend) SetDingSignKey(v string) *MergeRobotExtend {
	s.DingSignKey = &v
	return s
}

func (s *MergeRobotExtend) SetEnableOutgoing(v bool) *MergeRobotExtend {
	s.EnableOutgoing = &v
	return s
}

func (s *MergeRobotExtend) SetToken(v string) *MergeRobotExtend {
	s.Token = &v
	return s
}

func (s *MergeRobotExtend) Validate() error {
	return dara.Validate(s)
}
