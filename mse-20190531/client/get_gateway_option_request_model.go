// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayOptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetGatewayOptionRequest
	GetAcceptLanguage() *string
	SetGatewayId(v int64) *GetGatewayOptionRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *GetGatewayOptionRequest
	GetGatewayUniqueId() *string
}

type GetGatewayOptionRequest struct {
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
	// The ID of the gateway.
	//
	// example:
	//
	// 429
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-c9bc5afd61014165bd58f621b491*****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
}

func (s GetGatewayOptionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayOptionRequest) GoString() string {
	return s.String()
}

func (s *GetGatewayOptionRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetGatewayOptionRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *GetGatewayOptionRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GetGatewayOptionRequest) SetAcceptLanguage(v string) *GetGatewayOptionRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetGatewayOptionRequest) SetGatewayId(v int64) *GetGatewayOptionRequest {
	s.GatewayId = &v
	return s
}

func (s *GetGatewayOptionRequest) SetGatewayUniqueId(v string) *GetGatewayOptionRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *GetGatewayOptionRequest) Validate() error {
	return dara.Validate(s)
}
