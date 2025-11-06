// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateTemplateRequest
	GetAppId() *string
	SetContent(v string) *CreateTemplateRequest
	GetContent() *string
	SetDescInfo(v string) *CreateTemplateRequest
	GetDescInfo() *string
	SetIconUrls(v string) *CreateTemplateRequest
	GetIconUrls() *string
	SetImageUrls(v string) *CreateTemplateRequest
	GetImageUrls() *string
	SetJumpAction(v int32) *CreateTemplateRequest
	GetJumpAction() *int32
	SetPushStyle(v int32) *CreateTemplateRequest
	GetPushStyle() *int32
	SetShowStyle(v int64) *CreateTemplateRequest
	GetShowStyle() *int64
	SetTemplateName(v string) *CreateTemplateRequest
	GetTemplateName() *string
	SetTenantId(v string) *CreateTemplateRequest
	GetTenantId() *string
	SetTitle(v string) *CreateTemplateRequest
	GetTitle() *string
	SetUri(v string) *CreateTemplateRequest
	GetUri() *string
	SetVariables(v string) *CreateTemplateRequest
	GetVariables() *string
	SetWorkspaceId(v string) *CreateTemplateRequest
	GetWorkspaceId() *string
}

type CreateTemplateRequest struct {
	// example:
	//
	// ALIPUBE5C3F6D091419
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
	DescInfo *string `json:"DescInfo,omitempty" xml:"DescInfo,omitempty"`
	// example:
	//
	// /
	IconUrls *string `json:"IconUrls,omitempty" xml:"IconUrls,omitempty"`
	// example:
	//
	// /
	ImageUrls *string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty"`
	// example:
	//
	// 1
	JumpAction *int32 `json:"JumpAction,omitempty" xml:"JumpAction,omitempty"`
	// example:
	//
	// 0
	PushStyle *int32 `json:"PushStyle,omitempty" xml:"PushStyle,omitempty"`
	// example:
	//
	// 0
	ShowStyle    *int64  `json:"ShowStyle,omitempty" xml:"ShowStyle,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// GKDDZPFH
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Title    *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// /
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// example:
	//
	// title,content
	Variables *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
	// example:
	//
	// default
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateTemplateRequest) GetContent() *string {
	return s.Content
}

func (s *CreateTemplateRequest) GetDescInfo() *string {
	return s.DescInfo
}

func (s *CreateTemplateRequest) GetIconUrls() *string {
	return s.IconUrls
}

func (s *CreateTemplateRequest) GetImageUrls() *string {
	return s.ImageUrls
}

func (s *CreateTemplateRequest) GetJumpAction() *int32 {
	return s.JumpAction
}

func (s *CreateTemplateRequest) GetPushStyle() *int32 {
	return s.PushStyle
}

func (s *CreateTemplateRequest) GetShowStyle() *int64 {
	return s.ShowStyle
}

func (s *CreateTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateTemplateRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateTemplateRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateTemplateRequest) GetUri() *string {
	return s.Uri
}

func (s *CreateTemplateRequest) GetVariables() *string {
	return s.Variables
}

func (s *CreateTemplateRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateTemplateRequest) SetAppId(v string) *CreateTemplateRequest {
	s.AppId = &v
	return s
}

func (s *CreateTemplateRequest) SetContent(v string) *CreateTemplateRequest {
	s.Content = &v
	return s
}

func (s *CreateTemplateRequest) SetDescInfo(v string) *CreateTemplateRequest {
	s.DescInfo = &v
	return s
}

func (s *CreateTemplateRequest) SetIconUrls(v string) *CreateTemplateRequest {
	s.IconUrls = &v
	return s
}

func (s *CreateTemplateRequest) SetImageUrls(v string) *CreateTemplateRequest {
	s.ImageUrls = &v
	return s
}

func (s *CreateTemplateRequest) SetJumpAction(v int32) *CreateTemplateRequest {
	s.JumpAction = &v
	return s
}

func (s *CreateTemplateRequest) SetPushStyle(v int32) *CreateTemplateRequest {
	s.PushStyle = &v
	return s
}

func (s *CreateTemplateRequest) SetShowStyle(v int64) *CreateTemplateRequest {
	s.ShowStyle = &v
	return s
}

func (s *CreateTemplateRequest) SetTemplateName(v string) *CreateTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateTemplateRequest) SetTenantId(v string) *CreateTemplateRequest {
	s.TenantId = &v
	return s
}

func (s *CreateTemplateRequest) SetTitle(v string) *CreateTemplateRequest {
	s.Title = &v
	return s
}

func (s *CreateTemplateRequest) SetUri(v string) *CreateTemplateRequest {
	s.Uri = &v
	return s
}

func (s *CreateTemplateRequest) SetVariables(v string) *CreateTemplateRequest {
	s.Variables = &v
	return s
}

func (s *CreateTemplateRequest) SetWorkspaceId(v string) *CreateTemplateRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateTemplateRequest) Validate() error {
	return dara.Validate(s)
}
