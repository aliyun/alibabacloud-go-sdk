// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiPathFieldConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *GetApiPathFieldConfigRequest
	GetGatewayId() *string
	SetHttpApiId(v string) *GetApiPathFieldConfigRequest
	GetHttpApiId() *string
}

type GetApiPathFieldConfigRequest struct {
	// The ID of the cloud-native API gateway.
	//
	// example:
	//
	// gw-123456
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The ID of the HTTP API.
	//
	// example:
	//
	// api-123456
	HttpApiId *string `json:"httpApiId,omitempty" xml:"httpApiId,omitempty"`
}

func (s GetApiPathFieldConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApiPathFieldConfigRequest) GoString() string {
	return s.String()
}

func (s *GetApiPathFieldConfigRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetApiPathFieldConfigRequest) GetHttpApiId() *string {
	return s.HttpApiId
}

func (s *GetApiPathFieldConfigRequest) SetGatewayId(v string) *GetApiPathFieldConfigRequest {
	s.GatewayId = &v
	return s
}

func (s *GetApiPathFieldConfigRequest) SetHttpApiId(v string) *GetApiPathFieldConfigRequest {
	s.HttpApiId = &v
	return s
}

func (s *GetApiPathFieldConfigRequest) Validate() error {
	return dara.Validate(s)
}
