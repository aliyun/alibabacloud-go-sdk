// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConnectorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnectors(v []*ListConnectorsResponseBodyConnectors) *ListConnectorsResponseBody
	GetConnectors() []*ListConnectorsResponseBodyConnectors
	SetRequestId(v string) *ListConnectorsResponseBody
	GetRequestId() *string
	SetTotalNum(v int32) *ListConnectorsResponseBody
	GetTotalNum() *int32
}

type ListConnectorsResponseBody struct {
	// The list of connectors.
	Connectors []*ListConnectorsResponseBodyConnectors `json:"Connectors,omitempty" xml:"Connectors,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of connectors.
	//
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListConnectorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConnectorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConnectorsResponseBody) GetConnectors() []*ListConnectorsResponseBodyConnectors {
	return s.Connectors
}

func (s *ListConnectorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConnectorsResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *ListConnectorsResponseBody) SetConnectors(v []*ListConnectorsResponseBodyConnectors) *ListConnectorsResponseBody {
	s.Connectors = v
	return s
}

func (s *ListConnectorsResponseBody) SetRequestId(v string) *ListConnectorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConnectorsResponseBody) SetTotalNum(v int32) *ListConnectorsResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListConnectorsResponseBody) Validate() error {
	if s.Connectors != nil {
		for _, item := range s.Connectors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListConnectorsResponseBodyConnectors struct {
	// Specifies whether global acceleration is enabled. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Enabled
	AccelerateStatus *string `json:"AccelerateStatus,omitempty" xml:"AccelerateStatus,omitempty"`
	// The internal access applications associated with the connector.
	Applications []*ListConnectorsResponseBodyConnectorsApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// The IP address of the cluster.
	//
	// example:
	//
	// 1.1.1.1
	ClusterIP *string `json:"ClusterIP,omitempty" xml:"ClusterIP,omitempty"`
	// The port of the cluster.
	//
	// example:
	//
	// 8000
	ClusterPort *string `json:"ClusterPort,omitempty" xml:"ClusterPort,omitempty"`
	// The deployed connector clients.
	ConnectorClients []*ListConnectorsResponseBodyConnectorsConnectorClients `json:"ConnectorClients,omitempty" xml:"ConnectorClients,omitempty" type:"Repeated"`
	// The connector ID.
	//
	// example:
	//
	// connector-94db94e06b98****
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// The time when the connector was created. The time is in the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2022-09-27 18:10:25
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The name of the connector.
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
	// The connection status of the connector. Valid values:
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
	// The upgrade time of the connector.
	UpgradeTime *ListConnectorsResponseBodyConnectorsUpgradeTime `json:"UpgradeTime,omitempty" xml:"UpgradeTime,omitempty" type:"Struct"`
	VipCidr     *string                                          `json:"VipCidr,omitempty" xml:"VipCidr,omitempty"`
}

func (s ListConnectorsResponseBodyConnectors) String() string {
	return dara.Prettify(s)
}

func (s ListConnectorsResponseBodyConnectors) GoString() string {
	return s.String()
}

func (s *ListConnectorsResponseBodyConnectors) GetAccelerateStatus() *string {
	return s.AccelerateStatus
}

func (s *ListConnectorsResponseBodyConnectors) GetApplications() []*ListConnectorsResponseBodyConnectorsApplications {
	return s.Applications
}

func (s *ListConnectorsResponseBodyConnectors) GetClusterIP() *string {
	return s.ClusterIP
}

func (s *ListConnectorsResponseBodyConnectors) GetClusterPort() *string {
	return s.ClusterPort
}

func (s *ListConnectorsResponseBodyConnectors) GetConnectorClients() []*ListConnectorsResponseBodyConnectorsConnectorClients {
	return s.ConnectorClients
}

func (s *ListConnectorsResponseBodyConnectors) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *ListConnectorsResponseBodyConnectors) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListConnectorsResponseBodyConnectors) GetName() *string {
	return s.Name
}

func (s *ListConnectorsResponseBodyConnectors) GetRegionId() *string {
	return s.RegionId
}

func (s *ListConnectorsResponseBodyConnectors) GetStatus() *string {
	return s.Status
}

func (s *ListConnectorsResponseBodyConnectors) GetSwitchStatus() *string {
	return s.SwitchStatus
}

func (s *ListConnectorsResponseBodyConnectors) GetUpgradeTime() *ListConnectorsResponseBodyConnectorsUpgradeTime {
	return s.UpgradeTime
}

func (s *ListConnectorsResponseBodyConnectors) GetVipCidr() *string {
	return s.VipCidr
}

func (s *ListConnectorsResponseBodyConnectors) SetAccelerateStatus(v string) *ListConnectorsResponseBodyConnectors {
	s.AccelerateStatus = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetApplications(v []*ListConnectorsResponseBodyConnectorsApplications) *ListConnectorsResponseBodyConnectors {
	s.Applications = v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetClusterIP(v string) *ListConnectorsResponseBodyConnectors {
	s.ClusterIP = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetClusterPort(v string) *ListConnectorsResponseBodyConnectors {
	s.ClusterPort = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetConnectorClients(v []*ListConnectorsResponseBodyConnectorsConnectorClients) *ListConnectorsResponseBodyConnectors {
	s.ConnectorClients = v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetConnectorId(v string) *ListConnectorsResponseBodyConnectors {
	s.ConnectorId = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetCreateTime(v string) *ListConnectorsResponseBodyConnectors {
	s.CreateTime = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetName(v string) *ListConnectorsResponseBodyConnectors {
	s.Name = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetRegionId(v string) *ListConnectorsResponseBodyConnectors {
	s.RegionId = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetStatus(v string) *ListConnectorsResponseBodyConnectors {
	s.Status = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetSwitchStatus(v string) *ListConnectorsResponseBodyConnectors {
	s.SwitchStatus = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetUpgradeTime(v *ListConnectorsResponseBodyConnectorsUpgradeTime) *ListConnectorsResponseBodyConnectors {
	s.UpgradeTime = v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetVipCidr(v string) *ListConnectorsResponseBodyConnectors {
	s.VipCidr = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) Validate() error {
	if s.Applications != nil {
		for _, item := range s.Applications {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ConnectorClients != nil {
		for _, item := range s.ConnectorClients {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UpgradeTime != nil {
		if err := s.UpgradeTime.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListConnectorsResponseBodyConnectorsApplications struct {
	// The ID of the internal access application.
	//
	// example:
	//
	// pa-application-e12860ef6c48****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The name of the internal access application.
	//
	// example:
	//
	// application_name
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
}

func (s ListConnectorsResponseBodyConnectorsApplications) String() string {
	return dara.Prettify(s)
}

func (s ListConnectorsResponseBodyConnectorsApplications) GoString() string {
	return s.String()
}

func (s *ListConnectorsResponseBodyConnectorsApplications) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListConnectorsResponseBodyConnectorsApplications) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ListConnectorsResponseBodyConnectorsApplications) SetApplicationId(v string) *ListConnectorsResponseBodyConnectorsApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectorsApplications) SetApplicationName(v string) *ListConnectorsResponseBodyConnectorsApplications {
	s.ApplicationName = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectorsApplications) Validate() error {
	return dara.Validate(s)
}

type ListConnectorsResponseBodyConnectorsConnectorClients struct {
	// The connection status between the connector client and the connector server.
	//
	// example:
	//
	// Disconnected
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	// The unique device identifier of the connector client.
	//
	// example:
	//
	// C50A2386-F851-4F11-920B-DF7148DA0C22
	DevTag *string `json:"DevTag,omitempty" xml:"DevTag,omitempty"`
	// The hostname of the connector client.
	//
	// example:
	//
	// connector_client
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The public IP address of the connector client.
	//
	// example:
	//
	// 192.0.2.1
	PublicIp *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
}

func (s ListConnectorsResponseBodyConnectorsConnectorClients) String() string {
	return dara.Prettify(s)
}

func (s ListConnectorsResponseBodyConnectorsConnectorClients) GoString() string {
	return s.String()
}

func (s *ListConnectorsResponseBodyConnectorsConnectorClients) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *ListConnectorsResponseBodyConnectorsConnectorClients) GetDevTag() *string {
	return s.DevTag
}

func (s *ListConnectorsResponseBodyConnectorsConnectorClients) GetHostname() *string {
	return s.Hostname
}

func (s *ListConnectorsResponseBodyConnectorsConnectorClients) GetPublicIp() *string {
	return s.PublicIp
}

func (s *ListConnectorsResponseBodyConnectorsConnectorClients) SetConnectionStatus(v string) *ListConnectorsResponseBodyConnectorsConnectorClients {
	s.ConnectionStatus = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectorsConnectorClients) SetDevTag(v string) *ListConnectorsResponseBodyConnectorsConnectorClients {
	s.DevTag = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectorsConnectorClients) SetHostname(v string) *ListConnectorsResponseBodyConnectorsConnectorClients {
	s.Hostname = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectorsConnectorClients) SetPublicIp(v string) *ListConnectorsResponseBodyConnectorsConnectorClients {
	s.PublicIp = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectorsConnectorClients) Validate() error {
	return dara.Validate(s)
}

type ListConnectorsResponseBodyConnectorsUpgradeTime struct {
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

func (s ListConnectorsResponseBodyConnectorsUpgradeTime) String() string {
	return dara.Prettify(s)
}

func (s ListConnectorsResponseBodyConnectorsUpgradeTime) GoString() string {
	return s.String()
}

func (s *ListConnectorsResponseBodyConnectorsUpgradeTime) GetEnd() *string {
	return s.End
}

func (s *ListConnectorsResponseBodyConnectorsUpgradeTime) GetStart() *string {
	return s.Start
}

func (s *ListConnectorsResponseBodyConnectorsUpgradeTime) SetEnd(v string) *ListConnectorsResponseBodyConnectorsUpgradeTime {
	s.End = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectorsUpgradeTime) SetStart(v string) *ListConnectorsResponseBodyConnectorsUpgradeTime {
	s.Start = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectorsUpgradeTime) Validate() error {
	return dara.Validate(s)
}
