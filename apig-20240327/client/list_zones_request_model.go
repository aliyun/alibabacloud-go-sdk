// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayEdition(v string) *ListZonesRequest
	GetGatewayEdition() *string
}

type ListZonesRequest struct {
	// The target gateway edition for querying zones. Valid values:
	//
	// - Professional: standard gateway. This is the default value.
	//
	// - ServerlessV2: API multi-tenant Serverless V2.
	//
	// If this parameter is not specified, Professional is used.
	//
	// example:
	//
	// ServerlessV2
	GatewayEdition *string `json:"gatewayEdition,omitempty" xml:"gatewayEdition,omitempty"`
}

func (s ListZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListZonesRequest) GoString() string {
	return s.String()
}

func (s *ListZonesRequest) GetGatewayEdition() *string {
	return s.GatewayEdition
}

func (s *ListZonesRequest) SetGatewayEdition(v string) *ListZonesRequest {
	s.GatewayEdition = &v
	return s
}

func (s *ListZonesRequest) Validate() error {
	return dara.Validate(s)
}
