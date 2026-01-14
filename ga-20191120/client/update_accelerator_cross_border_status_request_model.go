// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAcceleratorCrossBorderStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *UpdateAcceleratorCrossBorderStatusRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *UpdateAcceleratorCrossBorderStatusRequest
	GetClientToken() *string
	SetCrossBorderStatus(v bool) *UpdateAcceleratorCrossBorderStatusRequest
	GetCrossBorderStatus() *bool
	SetRegionId(v string) *UpdateAcceleratorCrossBorderStatusRequest
	GetRegionId() *string
}

type UpdateAcceleratorCrossBorderStatusRequest struct {
	// The ID of the GA instance.
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
	// Specifies whether to enable the cross-border acceleration feature for the GA instance. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	CrossBorderStatus *bool `json:"CrossBorderStatus,omitempty" xml:"CrossBorderStatus,omitempty"`
	// The region ID of the GA instance. Set the value to cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateAcceleratorCrossBorderStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAcceleratorCrossBorderStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateAcceleratorCrossBorderStatusRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *UpdateAcceleratorCrossBorderStatusRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAcceleratorCrossBorderStatusRequest) GetCrossBorderStatus() *bool {
	return s.CrossBorderStatus
}

func (s *UpdateAcceleratorCrossBorderStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateAcceleratorCrossBorderStatusRequest) SetAcceleratorId(v string) *UpdateAcceleratorCrossBorderStatusRequest {
	s.AcceleratorId = &v
	return s
}

func (s *UpdateAcceleratorCrossBorderStatusRequest) SetClientToken(v string) *UpdateAcceleratorCrossBorderStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAcceleratorCrossBorderStatusRequest) SetCrossBorderStatus(v bool) *UpdateAcceleratorCrossBorderStatusRequest {
	s.CrossBorderStatus = &v
	return s
}

func (s *UpdateAcceleratorCrossBorderStatusRequest) SetRegionId(v string) *UpdateAcceleratorCrossBorderStatusRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAcceleratorCrossBorderStatusRequest) Validate() error {
	return dara.Validate(s)
}
