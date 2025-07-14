// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTomcatConfig interface {
	dara.Model
	String() string
	GoString() string
	SetContextPath(v string) *TomcatConfig
	GetContextPath() *string
	SetMaxThreads(v int32) *TomcatConfig
	GetMaxThreads() *int32
	SetPort(v int32) *TomcatConfig
	GetPort() *int32
	SetUriEncoding(v string) *TomcatConfig
	GetUriEncoding() *string
	SetUseBodyEncodingForUri(v bool) *TomcatConfig
	GetUseBodyEncodingForUri() *bool
	SetVersion(v string) *TomcatConfig
	GetVersion() *string
}

type TomcatConfig struct {
	ContextPath           *string `json:"ContextPath,omitempty" xml:"ContextPath,omitempty"`
	MaxThreads            *int32  `json:"MaxThreads,omitempty" xml:"MaxThreads,omitempty"`
	Port                  *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	UriEncoding           *string `json:"UriEncoding,omitempty" xml:"UriEncoding,omitempty"`
	UseBodyEncodingForUri *bool   `json:"UseBodyEncodingForUri,omitempty" xml:"UseBodyEncodingForUri,omitempty"`
	Version               *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s TomcatConfig) String() string {
	return dara.Prettify(s)
}

func (s TomcatConfig) GoString() string {
	return s.String()
}

func (s *TomcatConfig) GetContextPath() *string {
	return s.ContextPath
}

func (s *TomcatConfig) GetMaxThreads() *int32 {
	return s.MaxThreads
}

func (s *TomcatConfig) GetPort() *int32 {
	return s.Port
}

func (s *TomcatConfig) GetUriEncoding() *string {
	return s.UriEncoding
}

func (s *TomcatConfig) GetUseBodyEncodingForUri() *bool {
	return s.UseBodyEncodingForUri
}

func (s *TomcatConfig) GetVersion() *string {
	return s.Version
}

func (s *TomcatConfig) SetContextPath(v string) *TomcatConfig {
	s.ContextPath = &v
	return s
}

func (s *TomcatConfig) SetMaxThreads(v int32) *TomcatConfig {
	s.MaxThreads = &v
	return s
}

func (s *TomcatConfig) SetPort(v int32) *TomcatConfig {
	s.Port = &v
	return s
}

func (s *TomcatConfig) SetUriEncoding(v string) *TomcatConfig {
	s.UriEncoding = &v
	return s
}

func (s *TomcatConfig) SetUseBodyEncodingForUri(v bool) *TomcatConfig {
	s.UseBodyEncodingForUri = &v
	return s
}

func (s *TomcatConfig) SetVersion(v string) *TomcatConfig {
	s.Version = &v
	return s
}

func (s *TomcatConfig) Validate() error {
	return dara.Validate(s)
}
