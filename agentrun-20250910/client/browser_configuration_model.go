// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBrowserConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetBrowserType(v string) *BrowserConfiguration
	GetBrowserType() *string
	SetEnableExtension(v []*string) *BrowserConfiguration
	GetEnableExtension() []*string
	SetHeadless(v bool) *BrowserConfiguration
	GetHeadless() *bool
	SetUserAgent(v string) *BrowserConfiguration
	GetUserAgent() *string
	SetViewPort(v *ViewPortConfiguration) *BrowserConfiguration
	GetViewPort() *ViewPortConfiguration
}

type BrowserConfiguration struct {
	BrowserType *string `json:"browserType,omitempty" xml:"browserType,omitempty"`
	// 要启用的浏览器扩展列表
	EnableExtension []*string `json:"enableExtension,omitempty" xml:"enableExtension,omitempty" type:"Repeated"`
	// 是否以无头模式运行浏览器
	Headless *bool `json:"headless,omitempty" xml:"headless,omitempty"`
	// 浏览器用户代理字符串
	UserAgent *string                `json:"userAgent,omitempty" xml:"userAgent,omitempty"`
	ViewPort  *ViewPortConfiguration `json:"viewPort,omitempty" xml:"viewPort,omitempty"`
}

func (s BrowserConfiguration) String() string {
	return dara.Prettify(s)
}

func (s BrowserConfiguration) GoString() string {
	return s.String()
}

func (s *BrowserConfiguration) GetBrowserType() *string {
	return s.BrowserType
}

func (s *BrowserConfiguration) GetEnableExtension() []*string {
	return s.EnableExtension
}

func (s *BrowserConfiguration) GetHeadless() *bool {
	return s.Headless
}

func (s *BrowserConfiguration) GetUserAgent() *string {
	return s.UserAgent
}

func (s *BrowserConfiguration) GetViewPort() *ViewPortConfiguration {
	return s.ViewPort
}

func (s *BrowserConfiguration) SetBrowserType(v string) *BrowserConfiguration {
	s.BrowserType = &v
	return s
}

func (s *BrowserConfiguration) SetEnableExtension(v []*string) *BrowserConfiguration {
	s.EnableExtension = v
	return s
}

func (s *BrowserConfiguration) SetHeadless(v bool) *BrowserConfiguration {
	s.Headless = &v
	return s
}

func (s *BrowserConfiguration) SetUserAgent(v string) *BrowserConfiguration {
	s.UserAgent = &v
	return s
}

func (s *BrowserConfiguration) SetViewPort(v *ViewPortConfiguration) *BrowserConfiguration {
	s.ViewPort = v
	return s
}

func (s *BrowserConfiguration) Validate() error {
	if s.ViewPort != nil {
		if err := s.ViewPort.Validate(); err != nil {
			return err
		}
	}
	return nil
}
