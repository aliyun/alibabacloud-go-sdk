// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartPolarClawGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *RestartPolarClawGatewayRequest
	GetApplicationId() *string
	SetMode(v string) *RestartPolarClawGatewayRequest
	GetMode() *string
}

type RestartPolarClawGatewayRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-xxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// in-process
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s RestartPolarClawGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartPolarClawGatewayRequest) GoString() string {
	return s.String()
}

func (s *RestartPolarClawGatewayRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *RestartPolarClawGatewayRequest) GetMode() *string {
	return s.Mode
}

func (s *RestartPolarClawGatewayRequest) SetApplicationId(v string) *RestartPolarClawGatewayRequest {
	s.ApplicationId = &v
	return s
}

func (s *RestartPolarClawGatewayRequest) SetMode(v string) *RestartPolarClawGatewayRequest {
	s.Mode = &v
	return s
}

func (s *RestartPolarClawGatewayRequest) Validate() error {
	return dara.Validate(s)
}
