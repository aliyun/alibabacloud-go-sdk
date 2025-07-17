// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVersion(v string) *UpgradeGatewayRequest
	GetVersion() *string
}

type UpgradeGatewayRequest struct {
	// Gateway version.
	//
	// example:
	//
	// 2.0.2
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UpgradeGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeGatewayRequest) GoString() string {
	return s.String()
}

func (s *UpgradeGatewayRequest) GetVersion() *string {
	return s.Version
}

func (s *UpgradeGatewayRequest) SetVersion(v string) *UpgradeGatewayRequest {
	s.Version = &v
	return s
}

func (s *UpgradeGatewayRequest) Validate() error {
	return dara.Validate(s)
}
