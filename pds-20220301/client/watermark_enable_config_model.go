// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWatermarkEnableConfig interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayAccessUserName(v bool) *WatermarkEnableConfig
	GetDisplayAccessUserName() *bool
	SetDisplayCustomText(v string) *WatermarkEnableConfig
	GetDisplayCustomText() *string
	SetDisplayShareLinkCreatorName(v bool) *WatermarkEnableConfig
	GetDisplayShareLinkCreatorName() *bool
	SetEnableDocPreview(v bool) *WatermarkEnableConfig
	GetEnableDocPreview() *bool
}

type WatermarkEnableConfig struct {
	DisplayAccessUserName       *bool   `json:"display_access_user_name,omitempty" xml:"display_access_user_name,omitempty"`
	DisplayCustomText           *string `json:"display_custom_text,omitempty" xml:"display_custom_text,omitempty"`
	DisplayShareLinkCreatorName *bool   `json:"display_shareLink_creator_name,omitempty" xml:"display_shareLink_creator_name,omitempty"`
	EnableDocPreview            *bool   `json:"enable_doc_preview,omitempty" xml:"enable_doc_preview,omitempty"`
}

func (s WatermarkEnableConfig) String() string {
	return dara.Prettify(s)
}

func (s WatermarkEnableConfig) GoString() string {
	return s.String()
}

func (s *WatermarkEnableConfig) GetDisplayAccessUserName() *bool {
	return s.DisplayAccessUserName
}

func (s *WatermarkEnableConfig) GetDisplayCustomText() *string {
	return s.DisplayCustomText
}

func (s *WatermarkEnableConfig) GetDisplayShareLinkCreatorName() *bool {
	return s.DisplayShareLinkCreatorName
}

func (s *WatermarkEnableConfig) GetEnableDocPreview() *bool {
	return s.EnableDocPreview
}

func (s *WatermarkEnableConfig) SetDisplayAccessUserName(v bool) *WatermarkEnableConfig {
	s.DisplayAccessUserName = &v
	return s
}

func (s *WatermarkEnableConfig) SetDisplayCustomText(v string) *WatermarkEnableConfig {
	s.DisplayCustomText = &v
	return s
}

func (s *WatermarkEnableConfig) SetDisplayShareLinkCreatorName(v bool) *WatermarkEnableConfig {
	s.DisplayShareLinkCreatorName = &v
	return s
}

func (s *WatermarkEnableConfig) SetEnableDocPreview(v bool) *WatermarkEnableConfig {
	s.EnableDocPreview = &v
	return s
}

func (s *WatermarkEnableConfig) Validate() error {
	return dara.Validate(s)
}
