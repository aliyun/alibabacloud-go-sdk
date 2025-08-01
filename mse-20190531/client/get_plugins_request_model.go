// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPluginsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetPluginsRequest
	GetAcceptLanguage() *string
	SetCategory(v int32) *GetPluginsRequest
	GetCategory() *int32
	SetEnableOnly(v bool) *GetPluginsRequest
	GetEnableOnly() *bool
	SetGatewayUniqueId(v string) *GetPluginsRequest
	GetGatewayUniqueId() *string
	SetName(v string) *GetPluginsRequest
	GetName() *string
}

type GetPluginsRequest struct {
	// The language of the response. Valid values:
	//
	// zh: Chinese en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The type of the plug-in. Valid values:
	//
	// 	- 0: custom
	//
	// 	- 1: permission authorization
	//
	// 	- 2: security protection
	//
	// 	- 3: transmission protocol
	//
	// 	- 4: traffic control
	//
	// 	- 5: traffic observation
	//
	// example:
	//
	// 1
	Category *int32 `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to enable the plug-in.
	//
	// example:
	//
	// true
	EnableOnly *bool `json:"EnableOnly,omitempty" xml:"EnableOnly,omitempty"`
	// The ID of the gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-0adf3ad751284cc69fcf9669fba*****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The name of the plug-in.
	//
	// example:
	//
	// key-auth
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPluginsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPluginsRequest) GoString() string {
	return s.String()
}

func (s *GetPluginsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetPluginsRequest) GetCategory() *int32 {
	return s.Category
}

func (s *GetPluginsRequest) GetEnableOnly() *bool {
	return s.EnableOnly
}

func (s *GetPluginsRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GetPluginsRequest) GetName() *string {
	return s.Name
}

func (s *GetPluginsRequest) SetAcceptLanguage(v string) *GetPluginsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetPluginsRequest) SetCategory(v int32) *GetPluginsRequest {
	s.Category = &v
	return s
}

func (s *GetPluginsRequest) SetEnableOnly(v bool) *GetPluginsRequest {
	s.EnableOnly = &v
	return s
}

func (s *GetPluginsRequest) SetGatewayUniqueId(v string) *GetPluginsRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *GetPluginsRequest) SetName(v string) *GetPluginsRequest {
	s.Name = &v
	return s
}

func (s *GetPluginsRequest) Validate() error {
	return dara.Validate(s)
}
