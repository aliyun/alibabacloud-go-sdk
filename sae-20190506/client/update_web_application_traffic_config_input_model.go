// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWebApplicationTrafficConfigInput interface {
	dara.Model
	String() string
	GoString() string
	SetAuthType(v string) *UpdateWebApplicationTrafficConfigInput
	GetAuthType() *string
	SetDisableURLInternet(v bool) *UpdateWebApplicationTrafficConfigInput
	GetDisableURLInternet() *bool
	SetRevisionsTrafficWeight(v map[string]*float32) *UpdateWebApplicationTrafficConfigInput
	GetRevisionsTrafficWeight() map[string]*float32
	SetWebAclConfig(v *WebAclConfig) *UpdateWebApplicationTrafficConfigInput
	GetWebAclConfig() *WebAclConfig
}

type UpdateWebApplicationTrafficConfigInput struct {
	// example:
	//
	// Anonymous
	AuthType               *string             `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	DisableURLInternet     *bool               `json:"DisableURLInternet,omitempty" xml:"DisableURLInternet,omitempty"`
	RevisionsTrafficWeight map[string]*float32 `json:"RevisionsTrafficWeight,omitempty" xml:"RevisionsTrafficWeight,omitempty"`
	WebAclConfig           *WebAclConfig       `json:"WebAclConfig,omitempty" xml:"WebAclConfig,omitempty"`
}

func (s UpdateWebApplicationTrafficConfigInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebApplicationTrafficConfigInput) GoString() string {
	return s.String()
}

func (s *UpdateWebApplicationTrafficConfigInput) GetAuthType() *string {
	return s.AuthType
}

func (s *UpdateWebApplicationTrafficConfigInput) GetDisableURLInternet() *bool {
	return s.DisableURLInternet
}

func (s *UpdateWebApplicationTrafficConfigInput) GetRevisionsTrafficWeight() map[string]*float32 {
	return s.RevisionsTrafficWeight
}

func (s *UpdateWebApplicationTrafficConfigInput) GetWebAclConfig() *WebAclConfig {
	return s.WebAclConfig
}

func (s *UpdateWebApplicationTrafficConfigInput) SetAuthType(v string) *UpdateWebApplicationTrafficConfigInput {
	s.AuthType = &v
	return s
}

func (s *UpdateWebApplicationTrafficConfigInput) SetDisableURLInternet(v bool) *UpdateWebApplicationTrafficConfigInput {
	s.DisableURLInternet = &v
	return s
}

func (s *UpdateWebApplicationTrafficConfigInput) SetRevisionsTrafficWeight(v map[string]*float32) *UpdateWebApplicationTrafficConfigInput {
	s.RevisionsTrafficWeight = v
	return s
}

func (s *UpdateWebApplicationTrafficConfigInput) SetWebAclConfig(v *WebAclConfig) *UpdateWebApplicationTrafficConfigInput {
	s.WebAclConfig = v
	return s
}

func (s *UpdateWebApplicationTrafficConfigInput) Validate() error {
	return dara.Validate(s)
}
