// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConnectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnector(v *CreateConnectorResponseBodyConnector) *CreateConnectorResponseBody
	GetConnector() *CreateConnectorResponseBodyConnector
	SetRequestId(v string) *CreateConnectorResponseBody
	GetRequestId() *string
}

type CreateConnectorResponseBody struct {
	// Connector。
	Connector *CreateConnectorResponseBodyConnector `json:"Connector,omitempty" xml:"Connector,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// A1367BB2-A5D8-5E79-9403-2446757AC03C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateConnectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConnectorResponseBody) GetConnector() *CreateConnectorResponseBodyConnector {
	return s.Connector
}

func (s *CreateConnectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateConnectorResponseBody) SetConnector(v *CreateConnectorResponseBodyConnector) *CreateConnectorResponseBody {
	s.Connector = v
	return s
}

func (s *CreateConnectorResponseBody) SetRequestId(v string) *CreateConnectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConnectorResponseBody) Validate() error {
	if s.Connector != nil {
		if err := s.Connector.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateConnectorResponseBodyConnector struct {
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
	// - **Disabled**: Shutdown.
	//
	// example:
	//
	// Enabled
	SwitchStatus *string `json:"SwitchStatus,omitempty" xml:"SwitchStatus,omitempty"`
	// The connector upgrade time.
	UpgradeTime *CreateConnectorResponseBodyConnectorUpgradeTime `json:"UpgradeTime,omitempty" xml:"UpgradeTime,omitempty" type:"Struct"`
}

func (s CreateConnectorResponseBodyConnector) String() string {
	return dara.Prettify(s)
}

func (s CreateConnectorResponseBodyConnector) GoString() string {
	return s.String()
}

func (s *CreateConnectorResponseBodyConnector) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *CreateConnectorResponseBodyConnector) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateConnectorResponseBodyConnector) GetName() *string {
	return s.Name
}

func (s *CreateConnectorResponseBodyConnector) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateConnectorResponseBodyConnector) GetStatus() *string {
	return s.Status
}

func (s *CreateConnectorResponseBodyConnector) GetSwitchStatus() *string {
	return s.SwitchStatus
}

func (s *CreateConnectorResponseBodyConnector) GetUpgradeTime() *CreateConnectorResponseBodyConnectorUpgradeTime {
	return s.UpgradeTime
}

func (s *CreateConnectorResponseBodyConnector) SetConnectorId(v string) *CreateConnectorResponseBodyConnector {
	s.ConnectorId = &v
	return s
}

func (s *CreateConnectorResponseBodyConnector) SetCreateTime(v string) *CreateConnectorResponseBodyConnector {
	s.CreateTime = &v
	return s
}

func (s *CreateConnectorResponseBodyConnector) SetName(v string) *CreateConnectorResponseBodyConnector {
	s.Name = &v
	return s
}

func (s *CreateConnectorResponseBodyConnector) SetRegionId(v string) *CreateConnectorResponseBodyConnector {
	s.RegionId = &v
	return s
}

func (s *CreateConnectorResponseBodyConnector) SetStatus(v string) *CreateConnectorResponseBodyConnector {
	s.Status = &v
	return s
}

func (s *CreateConnectorResponseBodyConnector) SetSwitchStatus(v string) *CreateConnectorResponseBodyConnector {
	s.SwitchStatus = &v
	return s
}

func (s *CreateConnectorResponseBodyConnector) SetUpgradeTime(v *CreateConnectorResponseBodyConnectorUpgradeTime) *CreateConnectorResponseBodyConnector {
	s.UpgradeTime = v
	return s
}

func (s *CreateConnectorResponseBodyConnector) Validate() error {
	if s.UpgradeTime != nil {
		if err := s.UpgradeTime.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateConnectorResponseBodyConnectorUpgradeTime struct {
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

func (s CreateConnectorResponseBodyConnectorUpgradeTime) String() string {
	return dara.Prettify(s)
}

func (s CreateConnectorResponseBodyConnectorUpgradeTime) GoString() string {
	return s.String()
}

func (s *CreateConnectorResponseBodyConnectorUpgradeTime) GetEnd() *string {
	return s.End
}

func (s *CreateConnectorResponseBodyConnectorUpgradeTime) GetStart() *string {
	return s.Start
}

func (s *CreateConnectorResponseBodyConnectorUpgradeTime) SetEnd(v string) *CreateConnectorResponseBodyConnectorUpgradeTime {
	s.End = &v
	return s
}

func (s *CreateConnectorResponseBodyConnectorUpgradeTime) SetStart(v string) *CreateConnectorResponseBodyConnectorUpgradeTime {
	s.Start = &v
	return s
}

func (s *CreateConnectorResponseBodyConnectorUpgradeTime) Validate() error {
	return dara.Validate(s)
}
