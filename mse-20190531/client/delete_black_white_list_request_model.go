// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBlackWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteBlackWhiteListRequest
	GetAcceptLanguage() *string
	SetGatewayId(v int64) *DeleteBlackWhiteListRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *DeleteBlackWhiteListRequest
	GetGatewayUniqueId() *string
	SetId(v int64) *DeleteBlackWhiteListRequest
	GetId() *int64
}

type DeleteBlackWhiteListRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// example:
	//
	// 430
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// example:
	//
	// gw-9cdcf8e4f58144059e73ff4c5ef9****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// example:
	//
	// 120
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteBlackWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBlackWhiteListRequest) GoString() string {
	return s.String()
}

func (s *DeleteBlackWhiteListRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteBlackWhiteListRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *DeleteBlackWhiteListRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *DeleteBlackWhiteListRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteBlackWhiteListRequest) SetAcceptLanguage(v string) *DeleteBlackWhiteListRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteBlackWhiteListRequest) SetGatewayId(v int64) *DeleteBlackWhiteListRequest {
	s.GatewayId = &v
	return s
}

func (s *DeleteBlackWhiteListRequest) SetGatewayUniqueId(v string) *DeleteBlackWhiteListRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *DeleteBlackWhiteListRequest) SetId(v int64) *DeleteBlackWhiteListRequest {
	s.Id = &v
	return s
}

func (s *DeleteBlackWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
