// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTLSConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCipherSuites(v []*string) *TLSConfig
	GetCipherSuites() []*string
	SetMaxVersion(v string) *TLSConfig
	GetMaxVersion() *string
	SetMinVersion(v string) *TLSConfig
	GetMinVersion() *string
}

type TLSConfig struct {
	CipherSuites []*string `json:"cipherSuites,omitempty" xml:"cipherSuites,omitempty" type:"Repeated"`
	MaxVersion   *string   `json:"maxVersion,omitempty" xml:"maxVersion,omitempty"`
	MinVersion   *string   `json:"minVersion,omitempty" xml:"minVersion,omitempty"`
}

func (s TLSConfig) String() string {
	return dara.Prettify(s)
}

func (s TLSConfig) GoString() string {
	return s.String()
}

func (s *TLSConfig) GetCipherSuites() []*string {
	return s.CipherSuites
}

func (s *TLSConfig) GetMaxVersion() *string {
	return s.MaxVersion
}

func (s *TLSConfig) GetMinVersion() *string {
	return s.MinVersion
}

func (s *TLSConfig) SetCipherSuites(v []*string) *TLSConfig {
	s.CipherSuites = v
	return s
}

func (s *TLSConfig) SetMaxVersion(v string) *TLSConfig {
	s.MaxVersion = &v
	return s
}

func (s *TLSConfig) SetMinVersion(v string) *TLSConfig {
	s.MinVersion = &v
	return s
}

func (s *TLSConfig) Validate() error {
	return dara.Validate(s)
}
