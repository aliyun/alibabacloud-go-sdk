// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConnectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnector(v *GetConnectorResponseBodyConnector) *GetConnectorResponseBody
	GetConnector() *GetConnectorResponseBodyConnector
	SetRequestId(v string) *GetConnectorResponseBody
	GetRequestId() *string
}

type GetConnectorResponseBody struct {
	// Connector。
	Connector *GetConnectorResponseBodyConnector `json:"Connector,omitempty" xml:"Connector,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 748CFDC7-1EB6-5B8B-9405-DA76ED5BB60D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetConnectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *GetConnectorResponseBody) GetConnector() *GetConnectorResponseBodyConnector {
	return s.Connector
}

func (s *GetConnectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConnectorResponseBody) SetConnector(v *GetConnectorResponseBodyConnector) *GetConnectorResponseBody {
	s.Connector = v
	return s
}

func (s *GetConnectorResponseBody) SetRequestId(v string) *GetConnectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConnectorResponseBody) Validate() error {
	if s.Connector != nil {
		if err := s.Connector.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetConnectorResponseBodyConnector struct {
	// ConnectorID。
	//
	// example:
	//
	// connector-94db94e06b98****
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// The connector creation time.
	//
	// example:
	//
	// 2023-05-16 17:18:46
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The connector name.
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
	// The connector connection status. Valid values:
	//
	// - **Online**: Online.
	//
	// - **Offline**: Offline.
	//
	// example:
	//
	// Online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The connector instance status. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Enabled
	SwitchStatus *string `json:"SwitchStatus,omitempty" xml:"SwitchStatus,omitempty"`
	// The connector upgrade time.
	UpgradeTime *GetConnectorResponseBodyConnectorUpgradeTime `json:"UpgradeTime,omitempty" xml:"UpgradeTime,omitempty" type:"Struct"`
	// The virtual IP address range.
	//
	// example:
	//
	// 10.0.0.0/24
	VipCidr *string `json:"VipCidr,omitempty" xml:"VipCidr,omitempty"`
}

func (s GetConnectorResponseBodyConnector) String() string {
	return dara.Prettify(s)
}

func (s GetConnectorResponseBodyConnector) GoString() string {
	return s.String()
}

func (s *GetConnectorResponseBodyConnector) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *GetConnectorResponseBodyConnector) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetConnectorResponseBodyConnector) GetName() *string {
	return s.Name
}

func (s *GetConnectorResponseBodyConnector) GetRegionId() *string {
	return s.RegionId
}

func (s *GetConnectorResponseBodyConnector) GetStatus() *string {
	return s.Status
}

func (s *GetConnectorResponseBodyConnector) GetSwitchStatus() *string {
	return s.SwitchStatus
}

func (s *GetConnectorResponseBodyConnector) GetUpgradeTime() *GetConnectorResponseBodyConnectorUpgradeTime {
	return s.UpgradeTime
}

func (s *GetConnectorResponseBodyConnector) GetVipCidr() *string {
	return s.VipCidr
}

func (s *GetConnectorResponseBodyConnector) SetConnectorId(v string) *GetConnectorResponseBodyConnector {
	s.ConnectorId = &v
	return s
}

func (s *GetConnectorResponseBodyConnector) SetCreateTime(v string) *GetConnectorResponseBodyConnector {
	s.CreateTime = &v
	return s
}

func (s *GetConnectorResponseBodyConnector) SetName(v string) *GetConnectorResponseBodyConnector {
	s.Name = &v
	return s
}

func (s *GetConnectorResponseBodyConnector) SetRegionId(v string) *GetConnectorResponseBodyConnector {
	s.RegionId = &v
	return s
}

func (s *GetConnectorResponseBodyConnector) SetStatus(v string) *GetConnectorResponseBodyConnector {
	s.Status = &v
	return s
}

func (s *GetConnectorResponseBodyConnector) SetSwitchStatus(v string) *GetConnectorResponseBodyConnector {
	s.SwitchStatus = &v
	return s
}

func (s *GetConnectorResponseBodyConnector) SetUpgradeTime(v *GetConnectorResponseBodyConnectorUpgradeTime) *GetConnectorResponseBodyConnector {
	s.UpgradeTime = v
	return s
}

func (s *GetConnectorResponseBodyConnector) SetVipCidr(v string) *GetConnectorResponseBodyConnector {
	s.VipCidr = &v
	return s
}

func (s *GetConnectorResponseBodyConnector) Validate() error {
	if s.UpgradeTime != nil {
		if err := s.UpgradeTime.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetConnectorResponseBodyConnectorUpgradeTime struct {
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

func (s GetConnectorResponseBodyConnectorUpgradeTime) String() string {
	return dara.Prettify(s)
}

func (s GetConnectorResponseBodyConnectorUpgradeTime) GoString() string {
	return s.String()
}

func (s *GetConnectorResponseBodyConnectorUpgradeTime) GetEnd() *string {
	return s.End
}

func (s *GetConnectorResponseBodyConnectorUpgradeTime) GetStart() *string {
	return s.Start
}

func (s *GetConnectorResponseBodyConnectorUpgradeTime) SetEnd(v string) *GetConnectorResponseBodyConnectorUpgradeTime {
	s.End = &v
	return s
}

func (s *GetConnectorResponseBodyConnectorUpgradeTime) SetStart(v string) *GetConnectorResponseBodyConnectorUpgradeTime {
	s.Start = &v
	return s
}

func (s *GetConnectorResponseBodyConnectorUpgradeTime) Validate() error {
	return dara.Validate(s)
}
