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
	// example:
	//
	// understand
	Capability *string `json:"Capability,omitempty" xml:"Capability,omitempty"`
	// example:
	//
	// d123
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
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
