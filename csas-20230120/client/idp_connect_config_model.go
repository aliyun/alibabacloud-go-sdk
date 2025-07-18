// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIdpConnectConfig interface {
	dara.Model
	String() string
	GoString() string
	SetConnectorId(v string) *IdpConnectConfig
	GetConnectorId() *string
	SetUseConnector(v bool) *IdpConnectConfig
	GetUseConnector() *bool
}

type IdpConnectConfig struct {
	ConnectorId  *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	UseConnector *bool   `json:"UseConnector,omitempty" xml:"UseConnector,omitempty"`
}

func (s IdpConnectConfig) String() string {
	return dara.Prettify(s)
}

func (s IdpConnectConfig) GoString() string {
	return s.String()
}

func (s *IdpConnectConfig) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *IdpConnectConfig) GetUseConnector() *bool {
	return s.UseConnector
}

func (s *IdpConnectConfig) SetConnectorId(v string) *IdpConnectConfig {
	s.ConnectorId = &v
	return s
}

func (s *IdpConnectConfig) SetUseConnector(v bool) *IdpConnectConfig {
	s.UseConnector = &v
	return s
}

func (s *IdpConnectConfig) Validate() error {
	return dara.Validate(s)
}
