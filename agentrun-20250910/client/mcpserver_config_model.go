// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMCPServerConfig interface {
	dara.Model
	String() string
	GoString() string
	SetServerUrl(v string) *MCPServerConfig
	GetServerUrl() *string
	SetSsePath(v string) *MCPServerConfig
	GetSsePath() *string
	SetTransportType(v string) *MCPServerConfig
	GetTransportType() *string
}

type MCPServerConfig struct {
	ServerUrl     *string `json:"serverUrl,omitempty" xml:"serverUrl,omitempty"`
	SsePath       *string `json:"ssePath,omitempty" xml:"ssePath,omitempty"`
	TransportType *string `json:"transportType,omitempty" xml:"transportType,omitempty"`
}

func (s MCPServerConfig) String() string {
	return dara.Prettify(s)
}

func (s MCPServerConfig) GoString() string {
	return s.String()
}

func (s *MCPServerConfig) GetServerUrl() *string {
	return s.ServerUrl
}

func (s *MCPServerConfig) GetSsePath() *string {
	return s.SsePath
}

func (s *MCPServerConfig) GetTransportType() *string {
	return s.TransportType
}

func (s *MCPServerConfig) SetServerUrl(v string) *MCPServerConfig {
	s.ServerUrl = &v
	return s
}

func (s *MCPServerConfig) SetSsePath(v string) *MCPServerConfig {
	s.SsePath = &v
	return s
}

func (s *MCPServerConfig) SetTransportType(v string) *MCPServerConfig {
	s.TransportType = &v
	return s
}

func (s *MCPServerConfig) Validate() error {
	return dara.Validate(s)
}
