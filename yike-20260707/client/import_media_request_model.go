// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportMediaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v int64) *ImportMediaRequest
	GetCategoryId() *int64
	SetCoverURL(v string) *ImportMediaRequest
	GetCoverURL() *string
	SetDescription(v string) *ImportMediaRequest
	GetDescription() *string
	SetDynamicMetaData(v string) *ImportMediaRequest
	GetDynamicMetaData() *string
	SetEntityId(v string) *ImportMediaRequest
	GetEntityId() *string
	SetImportSource(v string) *ImportMediaRequest
	GetImportSource() *string
	SetInputURL(v string) *ImportMediaRequest
	GetInputURL() *string
	SetMediaTags(v string) *ImportMediaRequest
	GetMediaTags() *string
	SetMediaType(v string) *ImportMediaRequest
	GetMediaType() *string
	SetOverwrite(v bool) *ImportMediaRequest
	GetOverwrite() *bool
	SetRegisterConfig(v string) *ImportMediaRequest
	GetRegisterConfig() *string
	SetTitle(v string) *ImportMediaRequest
	GetTitle() *string
	SetUserData(v string) *ImportMediaRequest
	GetUserData() *string
}

type ImportMediaRequest struct {
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// https://outin-55c9ab3fb1b911ee817b00163e32b0a3.oss-cn-shanghai.aliyuncs.com/60425a2758a971f181385017f0e90102/covers/ice-generated/d4aee2d6c6f84769ac89f18c667699c6-cover.png
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// example:
	//
	// OK
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// {}
	DynamicMetaData *string `json:"DynamicMetaData,omitempty" xml:"DynamicMetaData,omitempty"`
	// example:
	//
	// urn:cruise:mock-saml-idp
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// example:
	//
	// url
	ImportSource *string `json:"ImportSource,omitempty" xml:"ImportSource,omitempty"`
	// example:
	//
	// https://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// example:
	//
	// 高级图生视频,AI生成
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// True
	Overwrite *bool `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	// example:
	//
	// {\\"SearchLibName\\":\\"AiSaasLib_34140718_MA\\"}
	RegisterConfig *string `json:"RegisterConfig,omitempty" xml:"RegisterConfig,omitempty"`
	// example:
	//
	// title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// {}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s ImportMediaRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportMediaRequest) GoString() string {
	return s.String()
}

func (s *ImportMediaRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *ImportMediaRequest) GetCoverURL() *string {
	return s.CoverURL
}

func (s *ImportMediaRequest) GetDescription() *string {
	return s.Description
}

func (s *ImportMediaRequest) GetDynamicMetaData() *string {
	return s.DynamicMetaData
}

func (s *ImportMediaRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *ImportMediaRequest) GetImportSource() *string {
	return s.ImportSource
}

func (s *ImportMediaRequest) GetInputURL() *string {
	return s.InputURL
}

func (s *ImportMediaRequest) GetMediaTags() *string {
	return s.MediaTags
}

func (s *ImportMediaRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *ImportMediaRequest) GetOverwrite() *bool {
	return s.Overwrite
}

func (s *ImportMediaRequest) GetRegisterConfig() *string {
	return s.RegisterConfig
}

func (s *ImportMediaRequest) GetTitle() *string {
	return s.Title
}

func (s *ImportMediaRequest) GetUserData() *string {
	return s.UserData
}

func (s *ImportMediaRequest) SetCategoryId(v int64) *ImportMediaRequest {
	s.CategoryId = &v
	return s
}

func (s *ImportMediaRequest) SetCoverURL(v string) *ImportMediaRequest {
	s.CoverURL = &v
	return s
}

func (s *ImportMediaRequest) SetDescription(v string) *ImportMediaRequest {
	s.Description = &v
	return s
}

func (s *ImportMediaRequest) SetDynamicMetaData(v string) *ImportMediaRequest {
	s.DynamicMetaData = &v
	return s
}

func (s *ImportMediaRequest) SetEntityId(v string) *ImportMediaRequest {
	s.EntityId = &v
	return s
}

func (s *ImportMediaRequest) SetImportSource(v string) *ImportMediaRequest {
	s.ImportSource = &v
	return s
}

func (s *ImportMediaRequest) SetInputURL(v string) *ImportMediaRequest {
	s.InputURL = &v
	return s
}

func (s *ImportMediaRequest) SetMediaTags(v string) *ImportMediaRequest {
	s.MediaTags = &v
	return s
}

func (s *ImportMediaRequest) SetMediaType(v string) *ImportMediaRequest {
	s.MediaType = &v
	return s
}

func (s *ImportMediaRequest) SetOverwrite(v bool) *ImportMediaRequest {
	s.Overwrite = &v
	return s
}

func (s *ImportMediaRequest) SetRegisterConfig(v string) *ImportMediaRequest {
	s.RegisterConfig = &v
	return s
}

func (s *ImportMediaRequest) SetTitle(v string) *ImportMediaRequest {
	s.Title = &v
	return s
}

func (s *ImportMediaRequest) SetUserData(v string) *ImportMediaRequest {
	s.UserData = &v
	return s
}

func (s *ImportMediaRequest) Validate() error {
	return dara.Validate(s)
}
