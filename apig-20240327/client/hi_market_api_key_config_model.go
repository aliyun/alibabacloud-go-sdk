// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHiMarketApiKeyConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCredentials(v []*HiMarketApiKeyConfigCredentials) *HiMarketApiKeyConfig
	GetCredentials() []*HiMarketApiKeyConfigCredentials
	SetKey(v string) *HiMarketApiKeyConfig
	GetKey() *string
	SetSource(v string) *HiMarketApiKeyConfig
	GetSource() *string
}

type HiMarketApiKeyConfig struct {
	Credentials []*HiMarketApiKeyConfigCredentials `json:"credentials,omitempty" xml:"credentials,omitempty" type:"Repeated"`
	Key         *string                            `json:"key,omitempty" xml:"key,omitempty"`
	Source      *string                            `json:"source,omitempty" xml:"source,omitempty"`
}

func (s HiMarketApiKeyConfig) String() string {
	return dara.Prettify(s)
}

func (s HiMarketApiKeyConfig) GoString() string {
	return s.String()
}

func (s *HiMarketApiKeyConfig) GetCredentials() []*HiMarketApiKeyConfigCredentials {
	return s.Credentials
}

func (s *HiMarketApiKeyConfig) GetKey() *string {
	return s.Key
}

func (s *HiMarketApiKeyConfig) GetSource() *string {
	return s.Source
}

func (s *HiMarketApiKeyConfig) SetCredentials(v []*HiMarketApiKeyConfigCredentials) *HiMarketApiKeyConfig {
	s.Credentials = v
	return s
}

func (s *HiMarketApiKeyConfig) SetKey(v string) *HiMarketApiKeyConfig {
	s.Key = &v
	return s
}

func (s *HiMarketApiKeyConfig) SetSource(v string) *HiMarketApiKeyConfig {
	s.Source = &v
	return s
}

func (s *HiMarketApiKeyConfig) Validate() error {
	if s.Credentials != nil {
		for _, item := range s.Credentials {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type HiMarketApiKeyConfigCredentials struct {
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	Mode   *string `json:"mode,omitempty" xml:"mode,omitempty"`
}

func (s HiMarketApiKeyConfigCredentials) String() string {
	return dara.Prettify(s)
}

func (s HiMarketApiKeyConfigCredentials) GoString() string {
	return s.String()
}

func (s *HiMarketApiKeyConfigCredentials) GetApiKey() *string {
	return s.ApiKey
}

func (s *HiMarketApiKeyConfigCredentials) GetMode() *string {
	return s.Mode
}

func (s *HiMarketApiKeyConfigCredentials) SetApiKey(v string) *HiMarketApiKeyConfigCredentials {
	s.ApiKey = &v
	return s
}

func (s *HiMarketApiKeyConfigCredentials) SetMode(v string) *HiMarketApiKeyConfigCredentials {
	s.Mode = &v
	return s
}

func (s *HiMarketApiKeyConfigCredentials) Validate() error {
	return dara.Validate(s)
}
