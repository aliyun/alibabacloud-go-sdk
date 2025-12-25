// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScenePreviewResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDraft(v bool) *GetScenePreviewResourceRequest
	GetDraft() *bool
	SetPreviewToken(v string) *GetScenePreviewResourceRequest
	GetPreviewToken() *string
}

type GetScenePreviewResourceRequest struct {
	// example:
	//
	// false
	Draft *bool `json:"Draft,omitempty" xml:"Draft,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2deb941b3e1****
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
}

func (s GetScenePreviewResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScenePreviewResourceRequest) GoString() string {
	return s.String()
}

func (s *GetScenePreviewResourceRequest) GetDraft() *bool {
	return s.Draft
}

func (s *GetScenePreviewResourceRequest) GetPreviewToken() *string {
	return s.PreviewToken
}

func (s *GetScenePreviewResourceRequest) SetDraft(v bool) *GetScenePreviewResourceRequest {
	s.Draft = &v
	return s
}

func (s *GetScenePreviewResourceRequest) SetPreviewToken(v string) *GetScenePreviewResourceRequest {
	s.PreviewToken = &v
	return s
}

func (s *GetScenePreviewResourceRequest) Validate() error {
	return dara.Validate(s)
}
