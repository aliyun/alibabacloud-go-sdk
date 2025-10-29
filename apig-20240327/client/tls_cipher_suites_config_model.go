// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTlsCipherSuitesConfig interface {
	dara.Model
	String() string
	GoString() string
	SetConfigType(v string) *TlsCipherSuitesConfig
	GetConfigType() *string
	SetTlsCipherSuite(v []*TlsCipherSuitesConfigTlsCipherSuite) *TlsCipherSuitesConfig
	GetTlsCipherSuite() []*TlsCipherSuitesConfigTlsCipherSuite
}

type TlsCipherSuitesConfig struct {
	ConfigType     *string                                `json:"configType,omitempty" xml:"configType,omitempty"`
	TlsCipherSuite []*TlsCipherSuitesConfigTlsCipherSuite `json:"tlsCipherSuite,omitempty" xml:"tlsCipherSuite,omitempty" type:"Repeated"`
}

func (s TlsCipherSuitesConfig) String() string {
	return dara.Prettify(s)
}

func (s TlsCipherSuitesConfig) GoString() string {
	return s.String()
}

func (s *TlsCipherSuitesConfig) GetConfigType() *string {
	return s.ConfigType
}

func (s *TlsCipherSuitesConfig) GetTlsCipherSuite() []*TlsCipherSuitesConfigTlsCipherSuite {
	return s.TlsCipherSuite
}

func (s *TlsCipherSuitesConfig) SetConfigType(v string) *TlsCipherSuitesConfig {
	s.ConfigType = &v
	return s
}

func (s *TlsCipherSuitesConfig) SetTlsCipherSuite(v []*TlsCipherSuitesConfigTlsCipherSuite) *TlsCipherSuitesConfig {
	s.TlsCipherSuite = v
	return s
}

func (s *TlsCipherSuitesConfig) Validate() error {
	if s.TlsCipherSuite != nil {
		for _, item := range s.TlsCipherSuite {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TlsCipherSuitesConfigTlsCipherSuite struct {
	Name            *string   `json:"name,omitempty" xml:"name,omitempty"`
	SupportVersions []*string `json:"supportVersions,omitempty" xml:"supportVersions,omitempty" type:"Repeated"`
}

func (s TlsCipherSuitesConfigTlsCipherSuite) String() string {
	return dara.Prettify(s)
}

func (s TlsCipherSuitesConfigTlsCipherSuite) GoString() string {
	return s.String()
}

func (s *TlsCipherSuitesConfigTlsCipherSuite) GetName() *string {
	return s.Name
}

func (s *TlsCipherSuitesConfigTlsCipherSuite) GetSupportVersions() []*string {
	return s.SupportVersions
}

func (s *TlsCipherSuitesConfigTlsCipherSuite) SetName(v string) *TlsCipherSuitesConfigTlsCipherSuite {
	s.Name = &v
	return s
}

func (s *TlsCipherSuitesConfigTlsCipherSuite) SetSupportVersions(v []*string) *TlsCipherSuitesConfigTlsCipherSuite {
	s.SupportVersions = v
	return s
}

func (s *TlsCipherSuitesConfigTlsCipherSuite) Validate() error {
	return dara.Validate(s)
}
