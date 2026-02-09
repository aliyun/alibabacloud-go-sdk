// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDialogBaseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttachments(v []*DialogBaseBodyAttachments) *DialogBaseBody
	GetAttachments() []*DialogBaseBodyAttachments
	SetChannelCode(v string) *DialogBaseBody
	GetChannelCode() *string
	SetCreateTime(v int64) *DialogBaseBody
	GetCreateTime() *int64
	SetDataInfo(v *DialogBaseBodyDataInfo) *DialogBaseBody
	GetDataInfo() *DialogBaseBodyDataInfo
	SetDialogId(v int64) *DialogBaseBody
	GetDialogId() *int64
	SetFooterInfo(v *DialogBaseBodyFooterInfo) *DialogBaseBody
	GetFooterInfo() *DialogBaseBodyFooterInfo
	SetHitWords(v []*string) *DialogBaseBody
	GetHitWords() []*string
	SetModifiedTime(v int64) *DialogBaseBody
	GetModifiedTime() *int64
	SetReferInfo(v *DialogBaseBody) *DialogBaseBody
	GetReferInfo() *DialogBaseBody
	SetStage(v int32) *DialogBaseBody
	GetStage() *int32
	SetStatus(v int32) *DialogBaseBody
	GetStatus() *int32
	SetTicketId(v string) *DialogBaseBody
	GetTicketId() *string
	SetTicketStatus(v int32) *DialogBaseBody
	GetTicketStatus() *int32
	SetTimestamp(v int64) *DialogBaseBody
	GetTimestamp() *int64
	SetTip(v string) *DialogBaseBody
	GetTip() *string
	SetType(v int32) *DialogBaseBody
	GetType() *int32
	SetUserInfo(v *DialogBaseBodyUserInfo) *DialogBaseBody
	GetUserInfo() *DialogBaseBodyUserInfo
}

type DialogBaseBody struct {
	Attachments  []*DialogBaseBodyAttachments `json:"Attachments,omitempty" xml:"Attachments,omitempty" type:"Repeated"`
	ChannelCode  *string                      `json:"ChannelCode,omitempty" xml:"ChannelCode,omitempty"`
	CreateTime   *int64                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DataInfo     *DialogBaseBodyDataInfo      `json:"DataInfo,omitempty" xml:"DataInfo,omitempty" type:"Struct"`
	DialogId     *int64                       `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	FooterInfo   *DialogBaseBodyFooterInfo    `json:"FooterInfo,omitempty" xml:"FooterInfo,omitempty" type:"Struct"`
	HitWords     []*string                    `json:"HitWords,omitempty" xml:"HitWords,omitempty" type:"Repeated"`
	ModifiedTime *int64                       `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ReferInfo    *DialogBaseBody              `json:"ReferInfo,omitempty" xml:"ReferInfo,omitempty"`
	Stage        *int32                       `json:"Stage,omitempty" xml:"Stage,omitempty"`
	Status       *int32                       `json:"Status,omitempty" xml:"Status,omitempty"`
	TicketId     *string                      `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	TicketStatus *int32                       `json:"TicketStatus,omitempty" xml:"TicketStatus,omitempty"`
	Timestamp    *int64                       `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Tip          *string                      `json:"Tip,omitempty" xml:"Tip,omitempty"`
	Type         *int32                       `json:"Type,omitempty" xml:"Type,omitempty"`
	UserInfo     *DialogBaseBodyUserInfo      `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s DialogBaseBody) String() string {
	return dara.Prettify(s)
}

func (s DialogBaseBody) GoString() string {
	return s.String()
}

func (s *DialogBaseBody) GetAttachments() []*DialogBaseBodyAttachments {
	return s.Attachments
}

func (s *DialogBaseBody) GetChannelCode() *string {
	return s.ChannelCode
}

func (s *DialogBaseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DialogBaseBody) GetDataInfo() *DialogBaseBodyDataInfo {
	return s.DataInfo
}

func (s *DialogBaseBody) GetDialogId() *int64 {
	return s.DialogId
}

func (s *DialogBaseBody) GetFooterInfo() *DialogBaseBodyFooterInfo {
	return s.FooterInfo
}

func (s *DialogBaseBody) GetHitWords() []*string {
	return s.HitWords
}

func (s *DialogBaseBody) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *DialogBaseBody) GetReferInfo() *DialogBaseBody {
	return s.ReferInfo
}

func (s *DialogBaseBody) GetStage() *int32 {
	return s.Stage
}

func (s *DialogBaseBody) GetStatus() *int32 {
	return s.Status
}

func (s *DialogBaseBody) GetTicketId() *string {
	return s.TicketId
}

func (s *DialogBaseBody) GetTicketStatus() *int32 {
	return s.TicketStatus
}

func (s *DialogBaseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DialogBaseBody) GetTip() *string {
	return s.Tip
}

func (s *DialogBaseBody) GetType() *int32 {
	return s.Type
}

func (s *DialogBaseBody) GetUserInfo() *DialogBaseBodyUserInfo {
	return s.UserInfo
}

func (s *DialogBaseBody) SetAttachments(v []*DialogBaseBodyAttachments) *DialogBaseBody {
	s.Attachments = v
	return s
}

func (s *DialogBaseBody) SetChannelCode(v string) *DialogBaseBody {
	s.ChannelCode = &v
	return s
}

func (s *DialogBaseBody) SetCreateTime(v int64) *DialogBaseBody {
	s.CreateTime = &v
	return s
}

func (s *DialogBaseBody) SetDataInfo(v *DialogBaseBodyDataInfo) *DialogBaseBody {
	s.DataInfo = v
	return s
}

func (s *DialogBaseBody) SetDialogId(v int64) *DialogBaseBody {
	s.DialogId = &v
	return s
}

func (s *DialogBaseBody) SetFooterInfo(v *DialogBaseBodyFooterInfo) *DialogBaseBody {
	s.FooterInfo = v
	return s
}

func (s *DialogBaseBody) SetHitWords(v []*string) *DialogBaseBody {
	s.HitWords = v
	return s
}

func (s *DialogBaseBody) SetModifiedTime(v int64) *DialogBaseBody {
	s.ModifiedTime = &v
	return s
}

func (s *DialogBaseBody) SetReferInfo(v *DialogBaseBody) *DialogBaseBody {
	s.ReferInfo = v
	return s
}

func (s *DialogBaseBody) SetStage(v int32) *DialogBaseBody {
	s.Stage = &v
	return s
}

func (s *DialogBaseBody) SetStatus(v int32) *DialogBaseBody {
	s.Status = &v
	return s
}

func (s *DialogBaseBody) SetTicketId(v string) *DialogBaseBody {
	s.TicketId = &v
	return s
}

func (s *DialogBaseBody) SetTicketStatus(v int32) *DialogBaseBody {
	s.TicketStatus = &v
	return s
}

func (s *DialogBaseBody) SetTimestamp(v int64) *DialogBaseBody {
	s.Timestamp = &v
	return s
}

func (s *DialogBaseBody) SetTip(v string) *DialogBaseBody {
	s.Tip = &v
	return s
}

func (s *DialogBaseBody) SetType(v int32) *DialogBaseBody {
	s.Type = &v
	return s
}

func (s *DialogBaseBody) SetUserInfo(v *DialogBaseBodyUserInfo) *DialogBaseBody {
	s.UserInfo = v
	return s
}

func (s *DialogBaseBody) Validate() error {
	if s.Attachments != nil {
		for _, item := range s.Attachments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DataInfo != nil {
		if err := s.DataInfo.Validate(); err != nil {
			return err
		}
	}
	if s.FooterInfo != nil {
		if err := s.FooterInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ReferInfo != nil {
		if err := s.ReferInfo.Validate(); err != nil {
			return err
		}
	}
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DialogBaseBodyAttachments struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url  *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DialogBaseBodyAttachments) String() string {
	return dara.Prettify(s)
}

func (s DialogBaseBodyAttachments) GoString() string {
	return s.String()
}

func (s *DialogBaseBodyAttachments) GetName() *string {
	return s.Name
}

func (s *DialogBaseBodyAttachments) GetSize() *string {
	return s.Size
}

func (s *DialogBaseBodyAttachments) GetType() *string {
	return s.Type
}

func (s *DialogBaseBodyAttachments) GetUrl() *string {
	return s.Url
}

func (s *DialogBaseBodyAttachments) SetName(v string) *DialogBaseBodyAttachments {
	s.Name = &v
	return s
}

func (s *DialogBaseBodyAttachments) SetSize(v string) *DialogBaseBodyAttachments {
	s.Size = &v
	return s
}

func (s *DialogBaseBodyAttachments) SetType(v string) *DialogBaseBodyAttachments {
	s.Type = &v
	return s
}

func (s *DialogBaseBodyAttachments) SetUrl(v string) *DialogBaseBodyAttachments {
	s.Url = &v
	return s
}

func (s *DialogBaseBodyAttachments) Validate() error {
	return dara.Validate(s)
}

type DialogBaseBodyDataInfo struct {
	BizId               *string                  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType             *int32                   `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Component           []map[string]interface{} `json:"Component,omitempty" xml:"Component,omitempty" type:"Repeated"`
	Container           interface{}              `json:"Container,omitempty" xml:"Container,omitempty"`
	Content             *string                  `json:"Content,omitempty" xml:"Content,omitempty"`
	ContentDesensitized *string                  `json:"ContentDesensitized,omitempty" xml:"ContentDesensitized,omitempty"`
	Editable            *int32                   `json:"Editable,omitempty" xml:"Editable,omitempty"`
	Props               []map[string]interface{} `json:"Props,omitempty" xml:"Props,omitempty" type:"Repeated"`
	Schema              *string                  `json:"Schema,omitempty" xml:"Schema,omitempty"`
	SchemaId            *int64                   `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
	ServiceChannel      *string                  `json:"ServiceChannel,omitempty" xml:"ServiceChannel,omitempty"`
	Title               *string                  `json:"Title,omitempty" xml:"Title,omitempty"`
	Values              map[string]interface{}   `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DialogBaseBodyDataInfo) String() string {
	return dara.Prettify(s)
}

func (s DialogBaseBodyDataInfo) GoString() string {
	return s.String()
}

func (s *DialogBaseBodyDataInfo) GetBizId() *string {
	return s.BizId
}

func (s *DialogBaseBodyDataInfo) GetBizType() *int32 {
	return s.BizType
}

func (s *DialogBaseBodyDataInfo) GetComponent() []map[string]interface{} {
	return s.Component
}

func (s *DialogBaseBodyDataInfo) GetContainer() interface{} {
	return s.Container
}

func (s *DialogBaseBodyDataInfo) GetContent() *string {
	return s.Content
}

func (s *DialogBaseBodyDataInfo) GetContentDesensitized() *string {
	return s.ContentDesensitized
}

func (s *DialogBaseBodyDataInfo) GetEditable() *int32 {
	return s.Editable
}

func (s *DialogBaseBodyDataInfo) GetProps() []map[string]interface{} {
	return s.Props
}

func (s *DialogBaseBodyDataInfo) GetSchema() *string {
	return s.Schema
}

func (s *DialogBaseBodyDataInfo) GetSchemaId() *int64 {
	return s.SchemaId
}

func (s *DialogBaseBodyDataInfo) GetServiceChannel() *string {
	return s.ServiceChannel
}

func (s *DialogBaseBodyDataInfo) GetTitle() *string {
	return s.Title
}

func (s *DialogBaseBodyDataInfo) GetValues() map[string]interface{} {
	return s.Values
}

func (s *DialogBaseBodyDataInfo) SetBizId(v string) *DialogBaseBodyDataInfo {
	s.BizId = &v
	return s
}

func (s *DialogBaseBodyDataInfo) SetBizType(v int32) *DialogBaseBodyDataInfo {
	s.BizType = &v
	return s
}

func (s *DialogBaseBodyDataInfo) SetComponent(v []map[string]interface{}) *DialogBaseBodyDataInfo {
	s.Component = v
	return s
}

func (s *DialogBaseBodyDataInfo) SetContainer(v interface{}) *DialogBaseBodyDataInfo {
	s.Container = v
	return s
}

func (s *DialogBaseBodyDataInfo) SetContent(v string) *DialogBaseBodyDataInfo {
	s.Content = &v
	return s
}

func (s *DialogBaseBodyDataInfo) SetContentDesensitized(v string) *DialogBaseBodyDataInfo {
	s.ContentDesensitized = &v
	return s
}

func (s *DialogBaseBodyDataInfo) SetEditable(v int32) *DialogBaseBodyDataInfo {
	s.Editable = &v
	return s
}

func (s *DialogBaseBodyDataInfo) SetProps(v []map[string]interface{}) *DialogBaseBodyDataInfo {
	s.Props = v
	return s
}

func (s *DialogBaseBodyDataInfo) SetSchema(v string) *DialogBaseBodyDataInfo {
	s.Schema = &v
	return s
}

func (s *DialogBaseBodyDataInfo) SetSchemaId(v int64) *DialogBaseBodyDataInfo {
	s.SchemaId = &v
	return s
}

func (s *DialogBaseBodyDataInfo) SetServiceChannel(v string) *DialogBaseBodyDataInfo {
	s.ServiceChannel = &v
	return s
}

func (s *DialogBaseBodyDataInfo) SetTitle(v string) *DialogBaseBodyDataInfo {
	s.Title = &v
	return s
}

func (s *DialogBaseBodyDataInfo) SetValues(v map[string]interface{}) *DialogBaseBodyDataInfo {
	s.Values = v
	return s
}

func (s *DialogBaseBodyDataInfo) Validate() error {
	return dara.Validate(s)
}

type DialogBaseBodyFooterInfo struct {
	Ext    map[string]interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Schema *string                `json:"Schema,omitempty" xml:"Schema,omitempty"`
}

func (s DialogBaseBodyFooterInfo) String() string {
	return dara.Prettify(s)
}

func (s DialogBaseBodyFooterInfo) GoString() string {
	return s.String()
}

func (s *DialogBaseBodyFooterInfo) GetExt() map[string]interface{} {
	return s.Ext
}

func (s *DialogBaseBodyFooterInfo) GetSchema() *string {
	return s.Schema
}

func (s *DialogBaseBodyFooterInfo) SetExt(v map[string]interface{}) *DialogBaseBodyFooterInfo {
	s.Ext = v
	return s
}

func (s *DialogBaseBodyFooterInfo) SetSchema(v string) *DialogBaseBodyFooterInfo {
	s.Schema = &v
	return s
}

func (s *DialogBaseBodyFooterInfo) Validate() error {
	return dara.Validate(s)
}

type DialogBaseBodyUserInfo struct {
	Avatar   *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	Role     *int32  `json:"Role,omitempty" xml:"Role,omitempty"`
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DialogBaseBodyUserInfo) String() string {
	return dara.Prettify(s)
}

func (s DialogBaseBodyUserInfo) GoString() string {
	return s.String()
}

func (s *DialogBaseBodyUserInfo) GetAvatar() *string {
	return s.Avatar
}

func (s *DialogBaseBodyUserInfo) GetRole() *int32 {
	return s.Role
}

func (s *DialogBaseBodyUserInfo) GetUserId() *string {
	return s.UserId
}

func (s *DialogBaseBodyUserInfo) GetUserName() *string {
	return s.UserName
}

func (s *DialogBaseBodyUserInfo) SetAvatar(v string) *DialogBaseBodyUserInfo {
	s.Avatar = &v
	return s
}

func (s *DialogBaseBodyUserInfo) SetRole(v int32) *DialogBaseBodyUserInfo {
	s.Role = &v
	return s
}

func (s *DialogBaseBodyUserInfo) SetUserId(v string) *DialogBaseBodyUserInfo {
	s.UserId = &v
	return s
}

func (s *DialogBaseBodyUserInfo) SetUserName(v string) *DialogBaseBodyUserInfo {
	s.UserName = &v
	return s
}

func (s *DialogBaseBodyUserInfo) Validate() error {
	return dara.Validate(s)
}
