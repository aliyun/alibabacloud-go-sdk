// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWxTrustedDomainConfig interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *WxTrustedDomainConfig
	GetContent() *string
	SetName(v string) *WxTrustedDomainConfig
	GetName() *string
	SetShow(v bool) *WxTrustedDomainConfig
	GetShow() *bool
}

type WxTrustedDomainConfig struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	Show    *bool   `json:"show,omitempty" xml:"show,omitempty"`
}

func (s WxTrustedDomainConfig) String() string {
	return dara.Prettify(s)
}

func (s WxTrustedDomainConfig) GoString() string {
	return s.String()
}

func (s *WxTrustedDomainConfig) GetContent() *string {
	return s.Content
}

func (s *WxTrustedDomainConfig) GetName() *string {
	return s.Name
}

func (s *WxTrustedDomainConfig) GetShow() *bool {
	return s.Show
}

func (s *WxTrustedDomainConfig) SetContent(v string) *WxTrustedDomainConfig {
	s.Content = &v
	return s
}

func (s *WxTrustedDomainConfig) SetName(v string) *WxTrustedDomainConfig {
	s.Name = &v
	return s
}

func (s *WxTrustedDomainConfig) SetShow(v bool) *WxTrustedDomainConfig {
	s.Show = &v
	return s
}

func (s *WxTrustedDomainConfig) Validate() error {
	return dara.Validate(s)
}
