// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePluginConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeletePluginConfigRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *DeletePluginConfigRequest
	GetGatewayUniqueId() *string
	SetPluginConfigId(v int64) *DeletePluginConfigRequest
	GetPluginConfigId() *int64
}

type DeletePluginConfigRequest struct {
	// The language in which you want to display the results. Valid values: zh and en. zh indicates Chinese, which is the default value. en indicates English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The unique ID of the gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-e2d226bba4b2445c9e29fa7f8216****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The plug-in configuration ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200
	PluginConfigId *int64 `json:"PluginConfigId,omitempty" xml:"PluginConfigId,omitempty"`
}

func (s DeletePluginConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePluginConfigRequest) GoString() string {
	return s.String()
}

func (s *DeletePluginConfigRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeletePluginConfigRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *DeletePluginConfigRequest) GetPluginConfigId() *int64 {
	return s.PluginConfigId
}

func (s *DeletePluginConfigRequest) SetAcceptLanguage(v string) *DeletePluginConfigRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeletePluginConfigRequest) SetGatewayUniqueId(v string) *DeletePluginConfigRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *DeletePluginConfigRequest) SetPluginConfigId(v int64) *DeletePluginConfigRequest {
	s.PluginConfigId = &v
	return s
}

func (s *DeletePluginConfigRequest) Validate() error {
	return dara.Validate(s)
}
