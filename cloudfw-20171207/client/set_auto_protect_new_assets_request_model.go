// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAutoProtectNewAssetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoProtect(v bool) *SetAutoProtectNewAssetsRequest
	GetAutoProtect() *bool
	SetLang(v string) *SetAutoProtectNewAssetsRequest
	GetLang() *string
}

type SetAutoProtectNewAssetsRequest struct {
	// Specifies whether to enable automatic traffic redirection. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	AutoProtect *bool `json:"AutoProtect,omitempty" xml:"AutoProtect,omitempty"`
	// The language type. Valid values:
	//
	// - **zh**: Chinese. Default value.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s SetAutoProtectNewAssetsRequest) String() string {
	return dara.Prettify(s)
}

func (s SetAutoProtectNewAssetsRequest) GoString() string {
	return s.String()
}

func (s *SetAutoProtectNewAssetsRequest) GetAutoProtect() *bool {
	return s.AutoProtect
}

func (s *SetAutoProtectNewAssetsRequest) GetLang() *string {
	return s.Lang
}

func (s *SetAutoProtectNewAssetsRequest) SetAutoProtect(v bool) *SetAutoProtectNewAssetsRequest {
	s.AutoProtect = &v
	return s
}

func (s *SetAutoProtectNewAssetsRequest) SetLang(v string) *SetAutoProtectNewAssetsRequest {
	s.Lang = &v
	return s
}

func (s *SetAutoProtectNewAssetsRequest) Validate() error {
	return dara.Validate(s)
}
