// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNotificationTemplate interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *NotificationTemplate
	GetCreateTime() *string
	SetDescription(v string) *NotificationTemplate
	GetDescription() *string
	SetEnContent(v string) *NotificationTemplate
	GetEnContent() *string
	SetEnItemContent(v string) *NotificationTemplate
	GetEnItemContent() *string
	SetEnTitle(v string) *NotificationTemplate
	GetEnTitle() *string
	SetName(v string) *NotificationTemplate
	GetName() *string
	SetType(v string) *NotificationTemplate
	GetType() *string
	SetUpdateTime(v string) *NotificationTemplate
	GetUpdateTime() *string
	SetUserId(v string) *NotificationTemplate
	GetUserId() *string
	SetUuid(v string) *NotificationTemplate
	GetUuid() *string
	SetWraperType(v string) *NotificationTemplate
	GetWraperType() *string
	SetZhContent(v string) *NotificationTemplate
	GetZhContent() *string
	SetZhItemContent(v string) *NotificationTemplate
	GetZhItemContent() *string
	SetZhTitle(v string) *NotificationTemplate
	GetZhTitle() *string
}

type NotificationTemplate struct {
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Alarm $.alertName
	EnContent *string `json:"EnContent,omitempty" xml:"EnContent,omitempty"`
	// example:
	//
	// Alarm $.alertName
	EnItemContent *string `json:"EnItemContent,omitempty" xml:"EnItemContent,omitempty"`
	// example:
	//
	// Alarm $.alertName
	EnTitle *string `json:"EnTitle,omitempty" xml:"EnTitle,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// DATA,  SMS,  ONCALL,  MAIL,  DING,  WEIXIN,  FEISHU,  SLACK
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Uuid       *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// example:
	//
	// TEXT, MARKDOWN,CARD
	WraperType *string `json:"WraperType,omitempty" xml:"WraperType,omitempty"`
	// example:
	//
	// 报警 $.alertName
	ZhContent *string `json:"ZhContent,omitempty" xml:"ZhContent,omitempty"`
	// example:
	//
	// 报警 $.alertName
	ZhItemContent *string `json:"ZhItemContent,omitempty" xml:"ZhItemContent,omitempty"`
	// example:
	//
	// 报警通知 $.alertName
	ZhTitle *string `json:"ZhTitle,omitempty" xml:"ZhTitle,omitempty"`
}

func (s NotificationTemplate) String() string {
	return dara.Prettify(s)
}

func (s NotificationTemplate) GoString() string {
	return s.String()
}

func (s *NotificationTemplate) GetCreateTime() *string {
	return s.CreateTime
}

func (s *NotificationTemplate) GetDescription() *string {
	return s.Description
}

func (s *NotificationTemplate) GetEnContent() *string {
	return s.EnContent
}

func (s *NotificationTemplate) GetEnItemContent() *string {
	return s.EnItemContent
}

func (s *NotificationTemplate) GetEnTitle() *string {
	return s.EnTitle
}

func (s *NotificationTemplate) GetName() *string {
	return s.Name
}

func (s *NotificationTemplate) GetType() *string {
	return s.Type
}

func (s *NotificationTemplate) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *NotificationTemplate) GetUserId() *string {
	return s.UserId
}

func (s *NotificationTemplate) GetUuid() *string {
	return s.Uuid
}

func (s *NotificationTemplate) GetWraperType() *string {
	return s.WraperType
}

func (s *NotificationTemplate) GetZhContent() *string {
	return s.ZhContent
}

func (s *NotificationTemplate) GetZhItemContent() *string {
	return s.ZhItemContent
}

func (s *NotificationTemplate) GetZhTitle() *string {
	return s.ZhTitle
}

func (s *NotificationTemplate) SetCreateTime(v string) *NotificationTemplate {
	s.CreateTime = &v
	return s
}

func (s *NotificationTemplate) SetDescription(v string) *NotificationTemplate {
	s.Description = &v
	return s
}

func (s *NotificationTemplate) SetEnContent(v string) *NotificationTemplate {
	s.EnContent = &v
	return s
}

func (s *NotificationTemplate) SetEnItemContent(v string) *NotificationTemplate {
	s.EnItemContent = &v
	return s
}

func (s *NotificationTemplate) SetEnTitle(v string) *NotificationTemplate {
	s.EnTitle = &v
	return s
}

func (s *NotificationTemplate) SetName(v string) *NotificationTemplate {
	s.Name = &v
	return s
}

func (s *NotificationTemplate) SetType(v string) *NotificationTemplate {
	s.Type = &v
	return s
}

func (s *NotificationTemplate) SetUpdateTime(v string) *NotificationTemplate {
	s.UpdateTime = &v
	return s
}

func (s *NotificationTemplate) SetUserId(v string) *NotificationTemplate {
	s.UserId = &v
	return s
}

func (s *NotificationTemplate) SetUuid(v string) *NotificationTemplate {
	s.Uuid = &v
	return s
}

func (s *NotificationTemplate) SetWraperType(v string) *NotificationTemplate {
	s.WraperType = &v
	return s
}

func (s *NotificationTemplate) SetZhContent(v string) *NotificationTemplate {
	s.ZhContent = &v
	return s
}

func (s *NotificationTemplate) SetZhItemContent(v string) *NotificationTemplate {
	s.ZhItemContent = &v
	return s
}

func (s *NotificationTemplate) SetZhTitle(v string) *NotificationTemplate {
	s.ZhTitle = &v
	return s
}

func (s *NotificationTemplate) Validate() error {
	return dara.Validate(s)
}
