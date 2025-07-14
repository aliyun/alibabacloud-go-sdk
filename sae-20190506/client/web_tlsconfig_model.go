// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebTLSConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCipherSuites(v []*string) *WebTLSConfig
	GetCipherSuites() []*string
	SetMaxVersion(v string) *WebTLSConfig
	GetMaxVersion() *string
	SetMinVersion(v string) *WebTLSConfig
	GetMinVersion() *string
}

type WebTLSConfig struct {
	CipherSuites []*string `json:"CipherSuites,omitempty" xml:"CipherSuites,omitempty" type:"Repeated"`
	MaxVersion   *string   `json:"MaxVersion,omitempty" xml:"MaxVersion,omitempty"`
	MinVersion   *string   `json:"MinVersion,omitempty" xml:"MinVersion,omitempty"`
}

func (s WebTLSConfig) String() string {
	return dara.Prettify(s)
}

func (s WebTLSConfig) GoString() string {
	return s.String()
}

func (s *WebTLSConfig) GetCipherSuites() []*string {
	return s.CipherSuites
}

func (s *WebTLSConfig) GetMaxVersion() *string {
	return s.MaxVersion
}

func (s *WebTLSConfig) GetMinVersion() *string {
	return s.MinVersion
}

func (s *WebTLSConfig) SetCipherSuites(v []*string) *WebTLSConfig {
	s.CipherSuites = v
	return s
}

func (s *WebTLSConfig) SetMaxVersion(v string) *WebTLSConfig {
	s.MaxVersion = &v
	return s
}

func (s *WebTLSConfig) SetMinVersion(v string) *WebTLSConfig {
	s.MinVersion = &v
	return s
}

func (s *WebTLSConfig) Validate() error {
	return dara.Validate(s)
}
