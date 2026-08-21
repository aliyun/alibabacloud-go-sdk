// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConnectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnector(v *UpdateConnectorResponseBodyConnector) *UpdateConnectorResponseBody
	GetConnector() *UpdateConnectorResponseBodyConnector
	SetRequestId(v string) *UpdateConnectorResponseBody
	GetRequestId() *string
}

type UpdateConnectorResponseBody struct {
	// Connector。
	Connector *UpdateConnectorResponseBodyConnector `json:"Connector,omitempty" xml:"Connector,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// BE4FB974-11BC-5453-9BE1-1606A73EACA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateConnectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConnectorResponseBody) GetConnector() *UpdateConnectorResponseBodyConnector {
	return s.Connector
}

func (s *UpdateConnectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConnectorResponseBody) SetConnector(v *UpdateConnectorResponseBodyConnector) *UpdateConnectorResponseBody {
	s.Connector = v
	return s
}

func (s *UpdateConnectorResponseBody) SetRequestId(v string) *UpdateConnectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConnectorResponseBody) Validate() error {
	if s.Connector != nil {
		if err := s.Connector.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateConnectorResponseBodyConnector struct {
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
	// ConnectorID。
	//
	// example:
	//
	// connector-94db94e06b98****
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// The creation time of the Connector.
	//
	// example:
	//
	// 2023-05-16 17:18:46
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The Connector name.
	//
	// example:
	//
	// connector_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The connection status of the Connector. Valid values:
	//
	// - **Online**: Online.
	//
	// - **Offline**: Offline.
	//
	// example:
	//
	// Online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	// The upgrade time of the Connector.
	UpgradeTime *UpdateConnectorResponseBodyConnectorUpgradeTime `json:"UpgradeTime,omitempty" xml:"UpgradeTime,omitempty" type:"Struct"`
	// The virtual IP address.
	//
	// example:
	//
	// 10.0.0.0/24
	VipCidr *string `json:"VipCidr,omitempty" xml:"VipCidr,omitempty"`
}

func (s UpdateConnectorResponseBodyConnector) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectorResponseBodyConnector) GoString() string {
	return s.String()
}

func (s *UpdateConnectorResponseBodyConnector) GetAccelerateStatus() *string {
	return s.AccelerateStatus
}

func (s *UpdateConnectorResponseBodyConnector) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *UpdateConnectorResponseBodyConnector) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateConnectorResponseBodyConnector) GetName() *string {
	return s.Name
}

func (s *UpdateConnectorResponseBodyConnector) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateConnectorResponseBodyConnector) GetStatus() *string {
	return s.Status
}

func (s *UpdateConnectorResponseBodyConnector) GetSwitchStatus() *string {
	return s.SwitchStatus
}

func (s *UpdateConnectorResponseBodyConnector) GetUpgradeTime() *UpdateConnectorResponseBodyConnectorUpgradeTime {
	return s.UpgradeTime
}

func (s *UpdateConnectorResponseBodyConnector) GetVipCidr() *string {
	return s.VipCidr
}

func (s *UpdateConnectorResponseBodyConnector) SetAccelerateStatus(v string) *UpdateConnectorResponseBodyConnector {
	s.AccelerateStatus = &v
	return s
}

func (s *UpdateConnectorResponseBodyConnector) SetConnectorId(v string) *UpdateConnectorResponseBodyConnector {
	s.ConnectorId = &v
	return s
}

func (s *UpdateConnectorResponseBodyConnector) SetCreateTime(v string) *UpdateConnectorResponseBodyConnector {
	s.CreateTime = &v
	return s
}

func (s *UpdateConnectorResponseBodyConnector) SetName(v string) *UpdateConnectorResponseBodyConnector {
	s.Name = &v
	return s
}

func (s *UpdateConnectorResponseBodyConnector) SetRegionId(v string) *UpdateConnectorResponseBodyConnector {
	s.RegionId = &v
	return s
}

func (s *UpdateConnectorResponseBodyConnector) SetStatus(v string) *UpdateConnectorResponseBodyConnector {
	s.Status = &v
	return s
}

func (s *UpdateConnectorResponseBodyConnector) SetSwitchStatus(v string) *UpdateConnectorResponseBodyConnector {
	s.SwitchStatus = &v
	return s
}

func (s *UpdateConnectorResponseBodyConnector) SetUpgradeTime(v *UpdateConnectorResponseBodyConnectorUpgradeTime) *UpdateConnectorResponseBodyConnector {
	s.UpgradeTime = v
	return s
}

func (s *UpdateConnectorResponseBodyConnector) SetVipCidr(v string) *UpdateConnectorResponseBodyConnector {
	s.VipCidr = &v
	return s
}

func (s *UpdateConnectorResponseBodyConnector) Validate() error {
	if s.UpgradeTime != nil {
		if err := s.UpgradeTime.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateConnectorResponseBodyConnectorUpgradeTime struct {
	// The end time.
	//
	// example:
	//
	// 23:00
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// The start time.
	//
	// example:
	//
	// 20:00
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s UpdateConnectorResponseBodyConnectorUpgradeTime) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectorResponseBodyConnectorUpgradeTime) GoString() string {
	return s.String()
}

func (s *UpdateConnectorResponseBodyConnectorUpgradeTime) GetEnd() *string {
	return s.End
}

func (s *UpdateConnectorResponseBodyConnectorUpgradeTime) GetStart() *string {
	return s.Start
}

func (s *UpdateConnectorResponseBodyConnectorUpgradeTime) SetEnd(v string) *UpdateConnectorResponseBodyConnectorUpgradeTime {
	s.End = &v
	return s
}

func (s *UpdateConnectorResponseBodyConnectorUpgradeTime) SetStart(v string) *UpdateConnectorResponseBodyConnectorUpgradeTime {
	s.Start = &v
	return s
}

func (s *UpdateConnectorResponseBodyConnectorUpgradeTime) Validate() error {
	return dara.Validate(s)
}
