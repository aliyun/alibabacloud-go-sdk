// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMCPServiceConfig interface {
	dara.Model
	String() string
	GoString() string
	SetServiceId(v string) *MCPServiceConfig
	GetServiceId() *string
}

type MCPServiceConfig struct {
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
}

func (s MCPServiceConfig) String() string {
	return dara.Prettify(s)
}

func (s MCPServiceConfig) GoString() string {
	return s.String()
}

func (s *MCPServiceConfig) GetServiceId() *string {
	return s.ServiceId
}

func (s *MCPServiceConfig) SetServiceId(v string) *MCPServiceConfig {
	s.ServiceId = &v
	return s
}

func (s *MCPServiceConfig) Validate() error {
	return dara.Validate(s)
}
