// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBlackWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetBlackWhiteListRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *GetBlackWhiteListRequest
	GetGatewayUniqueId() *string
	SetIsWhite(v bool) *GetBlackWhiteListRequest
	GetIsWhite() *bool
	SetResourceType(v string) *GetBlackWhiteListRequest
	GetResourceType() *string
	SetType(v string) *GetBlackWhiteListRequest
	GetType() *string
}

type GetBlackWhiteListRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-32d5c9769c1842b1a2cc3426c59e****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// Specifies whether to enable the whitelist.
	//
	// example:
	//
	// true
	IsWhite *bool `json:"IsWhite,omitempty" xml:"IsWhite,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// GATEWAY
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The type.
	//
	// example:
	//
	// IP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetBlackWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBlackWhiteListRequest) GoString() string {
	return s.String()
}

func (s *GetBlackWhiteListRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetBlackWhiteListRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GetBlackWhiteListRequest) GetIsWhite() *bool {
	return s.IsWhite
}

func (s *GetBlackWhiteListRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetBlackWhiteListRequest) GetType() *string {
	return s.Type
}

func (s *GetBlackWhiteListRequest) SetAcceptLanguage(v string) *GetBlackWhiteListRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetBlackWhiteListRequest) SetGatewayUniqueId(v string) *GetBlackWhiteListRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *GetBlackWhiteListRequest) SetIsWhite(v bool) *GetBlackWhiteListRequest {
	s.IsWhite = &v
	return s
}

func (s *GetBlackWhiteListRequest) SetResourceType(v string) *GetBlackWhiteListRequest {
	s.ResourceType = &v
	return s
}

func (s *GetBlackWhiteListRequest) SetType(v string) *GetBlackWhiteListRequest {
	s.Type = &v
	return s
}

func (s *GetBlackWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
