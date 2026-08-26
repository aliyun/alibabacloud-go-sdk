// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataAgentThemeUploadSignatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetThemeId(v string) *GetDataAgentThemeUploadSignatureRequest
	GetThemeId() *string
}

type GetDataAgentThemeUploadSignatureRequest struct {
	// The theme UUID. By default, you do not need to specify this parameter because the backend automatically generates and returns a UUID. Specify this parameter to regenerate a signature only when the previous signature has expired.
	//
	// example:
	//
	// 0f8b2c1d-****-****-****-9a3e5f7b1c2d
	ThemeId *string `json:"ThemeId,omitempty" xml:"ThemeId,omitempty"`
}

func (s GetDataAgentThemeUploadSignatureRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataAgentThemeUploadSignatureRequest) GoString() string {
	return s.String()
}

func (s *GetDataAgentThemeUploadSignatureRequest) GetThemeId() *string {
	return s.ThemeId
}

func (s *GetDataAgentThemeUploadSignatureRequest) SetThemeId(v string) *GetDataAgentThemeUploadSignatureRequest {
	s.ThemeId = &v
	return s
}

func (s *GetDataAgentThemeUploadSignatureRequest) Validate() error {
	return dara.Validate(s)
}
