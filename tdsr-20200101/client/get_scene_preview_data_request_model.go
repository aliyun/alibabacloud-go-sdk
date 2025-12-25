// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScenePreviewDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *GetScenePreviewDataRequest
	GetDomain() *string
	SetEnabled(v bool) *GetScenePreviewDataRequest
	GetEnabled() *bool
	SetPreviewToken(v string) *GetScenePreviewDataRequest
	GetPreviewToken() *string
	SetShowTag(v bool) *GetScenePreviewDataRequest
	GetShowTag() *bool
}

type GetScenePreviewDataRequest struct {
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// true/false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2735913e96da44ea8c86f8e777c8****
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
	// example:
	//
	// true/false
	ShowTag *bool `json:"ShowTag,omitempty" xml:"ShowTag,omitempty"`
}

func (s GetScenePreviewDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScenePreviewDataRequest) GoString() string {
	return s.String()
}

func (s *GetScenePreviewDataRequest) GetDomain() *string {
	return s.Domain
}

func (s *GetScenePreviewDataRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetScenePreviewDataRequest) GetPreviewToken() *string {
	return s.PreviewToken
}

func (s *GetScenePreviewDataRequest) GetShowTag() *bool {
	return s.ShowTag
}

func (s *GetScenePreviewDataRequest) SetDomain(v string) *GetScenePreviewDataRequest {
	s.Domain = &v
	return s
}

func (s *GetScenePreviewDataRequest) SetEnabled(v bool) *GetScenePreviewDataRequest {
	s.Enabled = &v
	return s
}

func (s *GetScenePreviewDataRequest) SetPreviewToken(v string) *GetScenePreviewDataRequest {
	s.PreviewToken = &v
	return s
}

func (s *GetScenePreviewDataRequest) SetShowTag(v bool) *GetScenePreviewDataRequest {
	s.ShowTag = &v
	return s
}

func (s *GetScenePreviewDataRequest) Validate() error {
	return dara.Validate(s)
}
