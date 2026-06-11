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
	// An array of objects, each containing a valid API key and its associated mode.
	Credentials []*HiMarketApiKeyConfigCredentials `json:"credentials,omitempty" xml:"credentials,omitempty" type:"Repeated"`
	// The name of the parameter that holds the API key. For example, if `source` is `HEADER`, this is the request header name, such as `X-API-Key`.
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The location of the API key in the request. Valid values are `HEADER` and `QUERY`.
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
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
	// The value of the API key.
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// The operational mode for the key, such as `test` or `production`.
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
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
