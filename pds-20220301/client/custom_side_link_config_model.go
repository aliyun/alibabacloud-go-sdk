// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomSideLinkConfig interface {
	dara.Model
	String() string
	GoString() string
	SetIcon(v string) *CustomSideLinkConfig
	GetIcon() *string
	SetLink(v string) *CustomSideLinkConfig
	GetLink() *string
	SetText(v string) *CustomSideLinkConfig
	GetText() *string
}

type CustomSideLinkConfig struct {
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s CustomSideLinkConfig) String() string {
	return dara.Prettify(s)
}

func (s CustomSideLinkConfig) GoString() string {
	return s.String()
}

func (s *CustomSideLinkConfig) GetIcon() *string {
	return s.Icon
}

func (s *CustomSideLinkConfig) GetLink() *string {
	return s.Link
}

func (s *CustomSideLinkConfig) GetText() *string {
	return s.Text
}

func (s *CustomSideLinkConfig) SetIcon(v string) *CustomSideLinkConfig {
	s.Icon = &v
	return s
}

func (s *CustomSideLinkConfig) SetLink(v string) *CustomSideLinkConfig {
	s.Link = &v
	return s
}

func (s *CustomSideLinkConfig) SetText(v string) *CustomSideLinkConfig {
	s.Text = &v
	return s
}

func (s *CustomSideLinkConfig) Validate() error {
	return dara.Validate(s)
}
