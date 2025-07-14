// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iArmsConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAgentVersion(v string) *ArmsConfig
	GetAgentVersion() *string
	SetAppId(v string) *ArmsConfig
	GetAppId() *string
	SetLicenseKey(v string) *ArmsConfig
	GetLicenseKey() *string
}

type ArmsConfig struct {
	AgentVersion *string `json:"agentVersion,omitempty" xml:"agentVersion,omitempty"`
	AppId        *string `json:"appId,omitempty" xml:"appId,omitempty"`
	LicenseKey   *string `json:"licenseKey,omitempty" xml:"licenseKey,omitempty"`
}

func (s ArmsConfig) String() string {
	return dara.Prettify(s)
}

func (s ArmsConfig) GoString() string {
	return s.String()
}

func (s *ArmsConfig) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *ArmsConfig) GetAppId() *string {
	return s.AppId
}

func (s *ArmsConfig) GetLicenseKey() *string {
	return s.LicenseKey
}

func (s *ArmsConfig) SetAgentVersion(v string) *ArmsConfig {
	s.AgentVersion = &v
	return s
}

func (s *ArmsConfig) SetAppId(v string) *ArmsConfig {
	s.AppId = &v
	return s
}

func (s *ArmsConfig) SetLicenseKey(v string) *ArmsConfig {
	s.LicenseKey = &v
	return s
}

func (s *ArmsConfig) Validate() error {
	return dara.Validate(s)
}
