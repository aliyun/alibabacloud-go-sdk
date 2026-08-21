// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConnectorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerateStatus(v string) *UpdateConnectorRequest
	GetAccelerateStatus() *string
	SetConnectorId(v string) *UpdateConnectorRequest
	GetConnectorId() *string
	SetName(v string) *UpdateConnectorRequest
	GetName() *string
	SetSwitchStatus(v string) *UpdateConnectorRequest
	GetSwitchStatus() *string
	SetVipCidr(v string) *UpdateConnectorRequest
	GetVipCidr() *string
}

type UpdateConnectorRequest struct {
	// Specifies whether to enable Global Accelerator. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Enabled
	AccelerateStatus *string `json:"AccelerateStatus,omitempty" xml:"AccelerateStatus,omitempty"`
	// The Connector ID. You can obtain the value by calling [ListConnectors](~~ListConnectors~~).
	//
	// This parameter is required.
	//
	// example:
	//
	// connector-94db94e06b98****
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// The Connector name. The name must be 1 to 128 characters in length and can contain Chinese characters, letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// connector_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The instance status of the Connector. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Shutdown.
	//
	// example:
	//
	// Enabled
	SwitchStatus *string `json:"SwitchStatus,omitempty" xml:"SwitchStatus,omitempty"`
	// The CIDR block of the virtual IP address.
	//
	// example:
	//
	// 10.0.0.0/24
	VipCidr *string `json:"VipCidr,omitempty" xml:"VipCidr,omitempty"`
}

func (s UpdateConnectorRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectorRequest) GoString() string {
	return s.String()
}

func (s *UpdateConnectorRequest) GetAccelerateStatus() *string {
	return s.AccelerateStatus
}

func (s *UpdateConnectorRequest) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *UpdateConnectorRequest) GetName() *string {
	return s.Name
}

func (s *UpdateConnectorRequest) GetSwitchStatus() *string {
	return s.SwitchStatus
}

func (s *UpdateConnectorRequest) GetVipCidr() *string {
	return s.VipCidr
}

func (s *UpdateConnectorRequest) SetAccelerateStatus(v string) *UpdateConnectorRequest {
	s.AccelerateStatus = &v
	return s
}

func (s *UpdateConnectorRequest) SetConnectorId(v string) *UpdateConnectorRequest {
	s.ConnectorId = &v
	return s
}

func (s *UpdateConnectorRequest) SetName(v string) *UpdateConnectorRequest {
	s.Name = &v
	return s
}

func (s *UpdateConnectorRequest) SetSwitchStatus(v string) *UpdateConnectorRequest {
	s.SwitchStatus = &v
	return s
}

func (s *UpdateConnectorRequest) SetVipCidr(v string) *UpdateConnectorRequest {
	s.VipCidr = &v
	return s
}

func (s *UpdateConnectorRequest) Validate() error {
	return dara.Validate(s)
}
