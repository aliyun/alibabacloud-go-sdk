// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScenegroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupUrl(v string) *GetScenegroupResponseBody
	GetGroupUrl() *string
	SetIcon(v string) *GetScenegroupResponseBody
	GetIcon() *string
	SetManagementOptions(v *GetScenegroupResponseBodyManagementOptions) *GetScenegroupResponseBody
	GetManagementOptions() *GetScenegroupResponseBodyManagementOptions
	SetOpenConversationId(v string) *GetScenegroupResponseBody
	GetOpenConversationId() *string
	SetOwnerStaffId(v string) *GetScenegroupResponseBody
	GetOwnerStaffId() *string
	SetRequestId(v string) *GetScenegroupResponseBody
	GetRequestId() *string
	SetSubAdminStaffIds(v []*string) *GetScenegroupResponseBody
	GetSubAdminStaffIds() []*string
	SetTemplateId(v string) *GetScenegroupResponseBody
	GetTemplateId() *string
	SetTitle(v string) *GetScenegroupResponseBody
	GetTitle() *string
	SetVendorRequestId(v string) *GetScenegroupResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetScenegroupResponseBody
	GetVendorType() *string
}

type GetScenegroupResponseBody struct {
	GroupUrl           *string                                     `json:"groupUrl,omitempty" xml:"groupUrl,omitempty"`
	Icon               *string                                     `json:"icon,omitempty" xml:"icon,omitempty"`
	ManagementOptions  *GetScenegroupResponseBodyManagementOptions `json:"managementOptions,omitempty" xml:"managementOptions,omitempty" type:"Struct"`
	OpenConversationId *string                                     `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	OwnerStaffId       *string                                     `json:"ownerStaffId,omitempty" xml:"ownerStaffId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId        *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SubAdminStaffIds []*string `json:"subAdminStaffIds,omitempty" xml:"subAdminStaffIds,omitempty" type:"Repeated"`
	TemplateId       *string   `json:"templateId,omitempty" xml:"templateId,omitempty"`
	Title            *string   `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetScenegroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScenegroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetScenegroupResponseBody) GetGroupUrl() *string {
	return s.GroupUrl
}

func (s *GetScenegroupResponseBody) GetIcon() *string {
	return s.Icon
}

func (s *GetScenegroupResponseBody) GetManagementOptions() *GetScenegroupResponseBodyManagementOptions {
	return s.ManagementOptions
}

func (s *GetScenegroupResponseBody) GetOpenConversationId() *string {
	return s.OpenConversationId
}

func (s *GetScenegroupResponseBody) GetOwnerStaffId() *string {
	return s.OwnerStaffId
}

func (s *GetScenegroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScenegroupResponseBody) GetSubAdminStaffIds() []*string {
	return s.SubAdminStaffIds
}

func (s *GetScenegroupResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetScenegroupResponseBody) GetTitle() *string {
	return s.Title
}

func (s *GetScenegroupResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetScenegroupResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetScenegroupResponseBody) SetGroupUrl(v string) *GetScenegroupResponseBody {
	s.GroupUrl = &v
	return s
}

func (s *GetScenegroupResponseBody) SetIcon(v string) *GetScenegroupResponseBody {
	s.Icon = &v
	return s
}

func (s *GetScenegroupResponseBody) SetManagementOptions(v *GetScenegroupResponseBodyManagementOptions) *GetScenegroupResponseBody {
	s.ManagementOptions = v
	return s
}

func (s *GetScenegroupResponseBody) SetOpenConversationId(v string) *GetScenegroupResponseBody {
	s.OpenConversationId = &v
	return s
}

func (s *GetScenegroupResponseBody) SetOwnerStaffId(v string) *GetScenegroupResponseBody {
	s.OwnerStaffId = &v
	return s
}

func (s *GetScenegroupResponseBody) SetRequestId(v string) *GetScenegroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScenegroupResponseBody) SetSubAdminStaffIds(v []*string) *GetScenegroupResponseBody {
	s.SubAdminStaffIds = v
	return s
}

func (s *GetScenegroupResponseBody) SetTemplateId(v string) *GetScenegroupResponseBody {
	s.TemplateId = &v
	return s
}

func (s *GetScenegroupResponseBody) SetTitle(v string) *GetScenegroupResponseBody {
	s.Title = &v
	return s
}

func (s *GetScenegroupResponseBody) SetVendorRequestId(v string) *GetScenegroupResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetScenegroupResponseBody) SetVendorType(v string) *GetScenegroupResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetScenegroupResponseBody) Validate() error {
	if s.ManagementOptions != nil {
		if err := s.ManagementOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScenegroupResponseBodyManagementOptions struct {
	ChatBannedType      *string `json:"ChatBannedType,omitempty" xml:"ChatBannedType,omitempty"`
	ManagementType      *string `json:"ManagementType,omitempty" xml:"ManagementType,omitempty"`
	MentionAllAuthority *string `json:"MentionAllAuthority,omitempty" xml:"MentionAllAuthority,omitempty"`
	Searchable          *string `json:"Searchable,omitempty" xml:"Searchable,omitempty"`
	ShowHistoryType     *string `json:"ShowHistoryType,omitempty" xml:"ShowHistoryType,omitempty"`
	ValidationType      *string `json:"ValidationType,omitempty" xml:"ValidationType,omitempty"`
}

func (s GetScenegroupResponseBodyManagementOptions) String() string {
	return dara.Prettify(s)
}

func (s GetScenegroupResponseBodyManagementOptions) GoString() string {
	return s.String()
}

func (s *GetScenegroupResponseBodyManagementOptions) GetChatBannedType() *string {
	return s.ChatBannedType
}

func (s *GetScenegroupResponseBodyManagementOptions) GetManagementType() *string {
	return s.ManagementType
}

func (s *GetScenegroupResponseBodyManagementOptions) GetMentionAllAuthority() *string {
	return s.MentionAllAuthority
}

func (s *GetScenegroupResponseBodyManagementOptions) GetSearchable() *string {
	return s.Searchable
}

func (s *GetScenegroupResponseBodyManagementOptions) GetShowHistoryType() *string {
	return s.ShowHistoryType
}

func (s *GetScenegroupResponseBodyManagementOptions) GetValidationType() *string {
	return s.ValidationType
}

func (s *GetScenegroupResponseBodyManagementOptions) SetChatBannedType(v string) *GetScenegroupResponseBodyManagementOptions {
	s.ChatBannedType = &v
	return s
}

func (s *GetScenegroupResponseBodyManagementOptions) SetManagementType(v string) *GetScenegroupResponseBodyManagementOptions {
	s.ManagementType = &v
	return s
}

func (s *GetScenegroupResponseBodyManagementOptions) SetMentionAllAuthority(v string) *GetScenegroupResponseBodyManagementOptions {
	s.MentionAllAuthority = &v
	return s
}

func (s *GetScenegroupResponseBodyManagementOptions) SetSearchable(v string) *GetScenegroupResponseBodyManagementOptions {
	s.Searchable = &v
	return s
}

func (s *GetScenegroupResponseBodyManagementOptions) SetShowHistoryType(v string) *GetScenegroupResponseBodyManagementOptions {
	s.ShowHistoryType = &v
	return s
}

func (s *GetScenegroupResponseBodyManagementOptions) SetValidationType(v string) *GetScenegroupResponseBodyManagementOptions {
	s.ValidationType = &v
	return s
}

func (s *GetScenegroupResponseBodyManagementOptions) Validate() error {
	return dara.Validate(s)
}
