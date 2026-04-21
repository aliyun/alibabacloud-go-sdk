// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHiMarketHmacConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCredentials(v []*HiMarketHmacConfigCredentials) *HiMarketHmacConfig
	GetCredentials() []*HiMarketHmacConfigCredentials
}

type HiMarketHmacConfig struct {
	Credentials []*HiMarketHmacConfigCredentials `json:"credentials,omitempty" xml:"credentials,omitempty" type:"Repeated"`
}

func (s HiMarketHmacConfig) String() string {
	return dara.Prettify(s)
}

func (s HiMarketHmacConfig) GoString() string {
	return s.String()
}

func (s *HiMarketHmacConfig) GetCredentials() []*HiMarketHmacConfigCredentials {
	return s.Credentials
}

func (s *HiMarketHmacConfig) SetCredentials(v []*HiMarketHmacConfigCredentials) *HiMarketHmacConfig {
	s.Credentials = v
	return s
}

func (s *HiMarketHmacConfig) Validate() error {
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

type HiMarketHmacConfigCredentials struct {
	Ak   *string `json:"ak,omitempty" xml:"ak,omitempty"`
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	Sk   *string `json:"sk,omitempty" xml:"sk,omitempty"`
}

func (s HiMarketHmacConfigCredentials) String() string {
	return dara.Prettify(s)
}

func (s HiMarketHmacConfigCredentials) GoString() string {
	return s.String()
}

func (s *HiMarketHmacConfigCredentials) GetAk() *string {
	return s.Ak
}

func (s *HiMarketHmacConfigCredentials) GetMode() *string {
	return s.Mode
}

func (s *HiMarketHmacConfigCredentials) GetSk() *string {
	return s.Sk
}

func (s *HiMarketHmacConfigCredentials) SetAk(v string) *HiMarketHmacConfigCredentials {
	s.Ak = &v
	return s
}

func (s *HiMarketHmacConfigCredentials) SetMode(v string) *HiMarketHmacConfigCredentials {
	s.Mode = &v
	return s
}

func (s *HiMarketHmacConfigCredentials) SetSk(v string) *HiMarketHmacConfigCredentials {
	s.Sk = &v
	return s
}

func (s *HiMarketHmacConfigCredentials) Validate() error {
	return dara.Validate(s)
}
