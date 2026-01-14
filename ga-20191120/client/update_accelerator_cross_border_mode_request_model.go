// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAcceleratorCrossBorderModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *UpdateAcceleratorCrossBorderModeRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *UpdateAcceleratorCrossBorderModeRequest
	GetClientToken() *string
	SetCrossBorderMode(v string) *UpdateAcceleratorCrossBorderModeRequest
	GetCrossBorderMode() *string
	SetRegionId(v string) *UpdateAcceleratorCrossBorderModeRequest
	GetRegionId() *string
}

type UpdateAcceleratorCrossBorderModeRequest struct {
	// The GA instance ID.
	//
	// > The bandwidth metering method of the GA instance must be pay-by-data-transfer.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The type of transmission network of the GA instance. Valid values:
	//
	// 	- **bgpPro**: BGP (Multi-ISP) Pro. BGP (Multi-ISP) Pro lines are used for cross-border acceleration. You do not need to complete real-name verification.
	//
	// 	- **private**: cross-border Express Connect circuit. Cross-border Express Connect circuits provide better acceleration performance but require real-name verification.
	//
	// This parameter is required.
	//
	// example:
	//
	// bgpPro
	CrossBorderMode *string `json:"CrossBorderMode,omitempty" xml:"CrossBorderMode,omitempty"`
	// The region ID of the GA instance. Set the value to **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateAcceleratorCrossBorderModeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAcceleratorCrossBorderModeRequest) GoString() string {
	return s.String()
}

func (s *UpdateAcceleratorCrossBorderModeRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *UpdateAcceleratorCrossBorderModeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAcceleratorCrossBorderModeRequest) GetCrossBorderMode() *string {
	return s.CrossBorderMode
}

func (s *UpdateAcceleratorCrossBorderModeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateAcceleratorCrossBorderModeRequest) SetAcceleratorId(v string) *UpdateAcceleratorCrossBorderModeRequest {
	s.AcceleratorId = &v
	return s
}

func (s *UpdateAcceleratorCrossBorderModeRequest) SetClientToken(v string) *UpdateAcceleratorCrossBorderModeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAcceleratorCrossBorderModeRequest) SetCrossBorderMode(v string) *UpdateAcceleratorCrossBorderModeRequest {
	s.CrossBorderMode = &v
	return s
}

func (s *UpdateAcceleratorCrossBorderModeRequest) SetRegionId(v string) *UpdateAcceleratorCrossBorderModeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAcceleratorCrossBorderModeRequest) Validate() error {
	return dara.Validate(s)
}
