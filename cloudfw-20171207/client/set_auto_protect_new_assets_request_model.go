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
	SetSourceIp(v string) *SetAutoProtectNewAssetsRequest
	GetSourceIp() *string
}

type SetAutoProtectNewAssetsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	AutoProtect *bool `json:"AutoProtect,omitempty" xml:"AutoProtect,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 60.182.79.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
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

func (s *SetAutoProtectNewAssetsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *SetAutoProtectNewAssetsRequest) SetAutoProtect(v bool) *SetAutoProtectNewAssetsRequest {
	s.AutoProtect = &v
	return s
}

func (s *SetAutoProtectNewAssetsRequest) SetLang(v string) *SetAutoProtectNewAssetsRequest {
	s.Lang = &v
	return s
}

func (s *SetAutoProtectNewAssetsRequest) SetSourceIp(v string) *SetAutoProtectNewAssetsRequest {
	s.SourceIp = &v
	return s
}

func (s *SetAutoProtectNewAssetsRequest) Validate() error {
	return dara.Validate(s)
}
