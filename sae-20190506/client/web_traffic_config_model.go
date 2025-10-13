// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebTrafficConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAuthType(v string) *WebTrafficConfig
	GetAuthType() *string
	SetDisableInternetURL(v bool) *WebTrafficConfig
	GetDisableInternetURL() *bool
	SetRevisionsTrafficWeight(v map[string]*float32) *WebTrafficConfig
	GetRevisionsTrafficWeight() map[string]*float32
	SetWebAclConfig(v *WebAclConfig) *WebTrafficConfig
	GetWebAclConfig() *WebAclConfig
}

type WebTrafficConfig struct {
	// example:
	//
	// Anonymous
	AuthType               *string             `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	DisableInternetURL     *bool               `json:"DisableInternetURL,omitempty" xml:"DisableInternetURL,omitempty"`
	RevisionsTrafficWeight map[string]*float32 `json:"RevisionsTrafficWeight,omitempty" xml:"RevisionsTrafficWeight,omitempty"`
	WebAclConfig           *WebAclConfig       `json:"WebAclConfig,omitempty" xml:"WebAclConfig,omitempty"`
}

func (s WebTrafficConfig) String() string {
	return dara.Prettify(s)
}

func (s WebTrafficConfig) GoString() string {
	return s.String()
}

func (s *WebTrafficConfig) GetAuthType() *string {
	return s.AuthType
}

func (s *WebTrafficConfig) GetDisableInternetURL() *bool {
	return s.DisableInternetURL
}

func (s *WebTrafficConfig) GetRevisionsTrafficWeight() map[string]*float32 {
	return s.RevisionsTrafficWeight
}

func (s *WebTrafficConfig) GetWebAclConfig() *WebAclConfig {
	return s.WebAclConfig
}

func (s *WebTrafficConfig) SetAuthType(v string) *WebTrafficConfig {
	s.AuthType = &v
	return s
}

func (s *WebTrafficConfig) SetDisableInternetURL(v bool) *WebTrafficConfig {
	s.DisableInternetURL = &v
	return s
}

func (s *WebTrafficConfig) SetRevisionsTrafficWeight(v map[string]*float32) *WebTrafficConfig {
	s.RevisionsTrafficWeight = v
	return s
}

func (s *WebTrafficConfig) SetWebAclConfig(v *WebAclConfig) *WebTrafficConfig {
	s.WebAclConfig = v
	return s
}

func (s *WebTrafficConfig) Validate() error {
	if s.WebAclConfig != nil {
		if err := s.WebAclConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
