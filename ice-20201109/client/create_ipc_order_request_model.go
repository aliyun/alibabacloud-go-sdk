// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpcOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCapability(v string) *CreateIpcOrderRequest
	GetCapability() *string
	SetDeviceId(v string) *CreateIpcOrderRequest
	GetDeviceId() *string
	SetPeriod(v string) *CreateIpcOrderRequest
	GetPeriod() *string
}

type CreateIpcOrderRequest struct {
	// The capability. Valid values:
	//
	// - understand: Image understanding. Supports the analysis of 300 images per day.
	//
	// - understand-reid: Image understanding with person re-identification (ReID). Supports the analysis of 300 images per day.
	//
	// - search: Search. Supports 75 searches per day.
	//
	// - understand-search: Image understanding and search. Supports the analysis of 300 images and 75 searches per day.
	//
	// - understand-reid-search: Image understanding with ReID and search. Supports the analysis of 300 images and 75 searches per day.
	//
	// example:
	//
	// understand
	Capability *string `json:"Capability,omitempty" xml:"Capability,omitempty"`
	// The device ID.
	//
	// example:
	//
	// d123
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// The subscription period. Valid values:
	//
	// - month: A monthly subscription, calculated as 30 days.
	//
	// - quarter: A quarterly subscription, calculated as 90 days.
	//
	// - year: An annual subscription, calculated as 365 days.
	//
	// example:
	//
	// month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
}

func (s CreateIpcOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIpcOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateIpcOrderRequest) GetCapability() *string {
	return s.Capability
}

func (s *CreateIpcOrderRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *CreateIpcOrderRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateIpcOrderRequest) SetCapability(v string) *CreateIpcOrderRequest {
	s.Capability = &v
	return s
}

func (s *CreateIpcOrderRequest) SetDeviceId(v string) *CreateIpcOrderRequest {
	s.DeviceId = &v
	return s
}

func (s *CreateIpcOrderRequest) SetPeriod(v string) *CreateIpcOrderRequest {
	s.Period = &v
	return s
}

func (s *CreateIpcOrderRequest) Validate() error {
	return dara.Validate(s)
}
