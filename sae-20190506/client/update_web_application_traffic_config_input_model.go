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
	// The authentication type. Valid values: Anonymous (default): does not require authentication. Application: requires authentication.
	//
	// example:
	//
	// Anonymous
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// Specifies whether to disable access to the default Internet domain. If you set this parameter to true, a 403 error is returned if you access the default public URL provided by the application. A value of false does not have affect the running of the function.
	//
	// example:
	//
	// true
	DisableURLInternet *bool `json:"DisableURLInternet,omitempty" xml:"DisableURLInternet,omitempty"`
	// The traffic distribution for the application versions. The sum of traffic percentages for all versions must be equal to 1.
	RevisionsTrafficWeight map[string]*float32 `json:"RevisionsTrafficWeight,omitempty" xml:"RevisionsTrafficWeight,omitempty"`
	// The configurations of the access control list (ACL) that consists of IP addresses.
	WebAclConfig *WebAclConfig `json:"WebAclConfig,omitempty" xml:"WebAclConfig,omitempty"`
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
	if s.WebAclConfig != nil {
		if err := s.WebAclConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
