// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveHotspotConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParamTag(v string) *SaveHotspotConfigRequest
	GetParamTag() *string
	SetPreviewToken(v string) *SaveHotspotConfigRequest
	GetPreviewToken() *string
}

type SaveHotspotConfigRequest struct {
	// example:
	//
	// {enabledTitleTag: 0, watermarkImg: []}
	ParamTag *string `json:"ParamTag,omitempty" xml:"ParamTag,omitempty"`
	// example:
	//
	// 5dc5c2dd927e45039dadb312384b****
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
}

func (s SaveHotspotConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveHotspotConfigRequest) GoString() string {
	return s.String()
}

func (s *SaveHotspotConfigRequest) GetParamTag() *string {
	return s.ParamTag
}

func (s *SaveHotspotConfigRequest) GetPreviewToken() *string {
	return s.PreviewToken
}

func (s *SaveHotspotConfigRequest) SetParamTag(v string) *SaveHotspotConfigRequest {
	s.ParamTag = &v
	return s
}

func (s *SaveHotspotConfigRequest) SetPreviewToken(v string) *SaveHotspotConfigRequest {
	s.PreviewToken = &v
	return s
}

func (s *SaveHotspotConfigRequest) Validate() error {
	return dara.Validate(s)
}
