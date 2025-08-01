// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPluginConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetPluginConfigRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *GetPluginConfigRequest
	GetGatewayUniqueId() *string
	SetPluginId(v int64) *GetPluginConfigRequest
	GetPluginId() *int64
}

type GetPluginConfigRequest struct {
	// The language of the response. Valid values:
	//
	// zh: Chinese en: English
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
	// gw-ubuwqygbq4783gqb2y3f87q****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the gateway plug-in.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	PluginId *int64 `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
}

func (s GetPluginConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPluginConfigRequest) GoString() string {
	return s.String()
}

func (s *GetPluginConfigRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetPluginConfigRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GetPluginConfigRequest) GetPluginId() *int64 {
	return s.PluginId
}

func (s *GetPluginConfigRequest) SetAcceptLanguage(v string) *GetPluginConfigRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetPluginConfigRequest) SetGatewayUniqueId(v string) *GetPluginConfigRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *GetPluginConfigRequest) SetPluginId(v int64) *GetPluginConfigRequest {
	s.PluginId = &v
	return s
}

func (s *GetPluginConfigRequest) Validate() error {
	return dara.Validate(s)
}
