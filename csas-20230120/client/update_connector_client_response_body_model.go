// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConnectorClientResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnectorClient(v *UpdateConnectorClientResponseBodyConnectorClient) *UpdateConnectorClientResponseBody
	GetConnectorClient() *UpdateConnectorClientResponseBodyConnectorClient
	SetRequestId(v string) *UpdateConnectorClientResponseBody
	GetRequestId() *string
}

type UpdateConnectorClientResponseBody struct {
	// ConnectorClient。
	ConnectorClient *UpdateConnectorClientResponseBodyConnectorClient `json:"ConnectorClient,omitempty" xml:"ConnectorClient,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateConnectorClientResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectorClientResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConnectorClientResponseBody) GetConnectorClient() *UpdateConnectorClientResponseBodyConnectorClient {
	return s.ConnectorClient
}

func (s *UpdateConnectorClientResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConnectorClientResponseBody) SetConnectorClient(v *UpdateConnectorClientResponseBodyConnectorClient) *UpdateConnectorClientResponseBody {
	s.ConnectorClient = v
	return s
}

func (s *UpdateConnectorClientResponseBody) SetRequestId(v string) *UpdateConnectorClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConnectorClientResponseBody) Validate() error {
	if s.ConnectorClient != nil {
		if err := s.ConnectorClient.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateConnectorClientResponseBodyConnectorClient struct {
	// The number of CPUs of the ConnectorClient.
	//
	// example:
	//
	// 16
	CPUSize *string `json:"CPUSize,omitempty" xml:"CPUSize,omitempty"`
	// The connection status of the ConnectorClient. Valid values:
	//
	// - **Connected**: connected.
	//
	// - **Disconnected**: disconnected.
	//
	// example:
	//
	// Connected
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	// ConnectorID。
	//
	// example:
	//
	// connector-94db94e06b98****
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// The time when the connector was created.
	//
	// example:
	//
	// 2023-05-16 17:18:46
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The unique identifier of the ConnectorClient device.
	//
	// example:
	//
	// E86266A0-2859-58EF-8F5B-EDE226FF6EB2
	DevTag *string `json:"DevTag,omitempty" xml:"DevTag,omitempty"`
	// The hostname.
	//
	// example:
	//
	// connector_client_hostname
	Hosname *string `json:"Hosname,omitempty" xml:"Hosname,omitempty"`
	// The kernel version of the ConnectorClient.
	//
	// example:
	//
	// 22.3.0
	KernelVersion *string `json:"KernelVersion,omitempty" xml:"KernelVersion,omitempty"`
	// The memory size of the ConnectorClient. Unit: MB.
	//
	// example:
	//
	// 32768
	MemorySize *string `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	// The O&M status. Valid values:
	//
	// - **Running**: O&M in progress.
	//
	// - **Failed**: O&M failed.
	//
	// - (empty string): not in O&M status.
	//
	// example:
	//
	// Running
	OperationStatus *string `json:"OperationStatus,omitempty" xml:"OperationStatus,omitempty"`
	// The private IP address of the ConnectorClient.
	//
	// example:
	//
	// 192.168.0.1
	PrivateIp *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	// The program runtime. Unit: seconds.
	//
	// example:
	//
	// 100
	ProcessRunTime *int32 `json:"ProcessRunTime,omitempty" xml:"ProcessRunTime,omitempty"`
	// The public IP address of the ConnectorClient.
	//
	// example:
	//
	// 203.0.113.1
	PublicIp *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	// The version number.
	ReleaseNotes []*string `json:"ReleaseNotes,omitempty" xml:"ReleaseNotes,omitempty" type:"Repeated"`
	// The enabled status of the ConnectorClient, which can be used to force the client offline. Valid values:
	//
	// - **Enabled**: enabled.
	//
	// - **Disabled**: disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The version status of the connector. Valid values:
	//
	// - **Latest**: the current version is the latest version.
	//
	// - **NewVersionAvailable**: a newer version is available for upgrade.
	//
	// example:
	//
	// Latest
	UpgradeStatus *string `json:"UpgradeStatus,omitempty" xml:"UpgradeStatus,omitempty"`
	// The blockchain version.
	//
	// example:
	//
	// 2019-04-02
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The version to roll back to.
	//
	// example:
	//
	// 1.0.0
	VersionToRollback *string `json:"VersionToRollback,omitempty" xml:"VersionToRollback,omitempty"`
}

func (s UpdateConnectorClientResponseBodyConnectorClient) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectorClientResponseBodyConnectorClient) GoString() string {
	return s.String()
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) GetCPUSize() *string {
	return s.CPUSize
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) GetDevTag() *string {
	return s.DevTag
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) GetHosname() *string {
	return s.Hosname
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) GetKernelVersion() *string {
	return s.KernelVersion
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) GetMemorySize() *string {
	return s.MemorySize
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) GetOperationStatus() *string {
	return s.OperationStatus
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) GetProcessRunTime() *int32 {
	return s.ProcessRunTime
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) GetPublicIp() *string {
	return s.PublicIp
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) GetReleaseNotes() []*string {
	return s.ReleaseNotes
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) GetStatus() *string {
	return s.Status
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) GetUpgradeStatus() *string {
	return s.UpgradeStatus
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) GetVersion() *string {
	return s.Version
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) GetVersionToRollback() *string {
	return s.VersionToRollback
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) SetCPUSize(v string) *UpdateConnectorClientResponseBodyConnectorClient {
	s.CPUSize = &v
	return s
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) SetConnectionStatus(v string) *UpdateConnectorClientResponseBodyConnectorClient {
	s.ConnectionStatus = &v
	return s
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) SetConnectorId(v string) *UpdateConnectorClientResponseBodyConnectorClient {
	s.ConnectorId = &v
	return s
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) SetCreateTime(v string) *UpdateConnectorClientResponseBodyConnectorClient {
	s.CreateTime = &v
	return s
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) SetDevTag(v string) *UpdateConnectorClientResponseBodyConnectorClient {
	s.DevTag = &v
	return s
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) SetHosname(v string) *UpdateConnectorClientResponseBodyConnectorClient {
	s.Hosname = &v
	return s
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) SetKernelVersion(v string) *UpdateConnectorClientResponseBodyConnectorClient {
	s.KernelVersion = &v
	return s
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) SetMemorySize(v string) *UpdateConnectorClientResponseBodyConnectorClient {
	s.MemorySize = &v
	return s
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) SetOperationStatus(v string) *UpdateConnectorClientResponseBodyConnectorClient {
	s.OperationStatus = &v
	return s
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) SetPrivateIp(v string) *UpdateConnectorClientResponseBodyConnectorClient {
	s.PrivateIp = &v
	return s
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) SetProcessRunTime(v int32) *UpdateConnectorClientResponseBodyConnectorClient {
	s.ProcessRunTime = &v
	return s
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) SetPublicIp(v string) *UpdateConnectorClientResponseBodyConnectorClient {
	s.PublicIp = &v
	return s
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) SetReleaseNotes(v []*string) *UpdateConnectorClientResponseBodyConnectorClient {
	s.ReleaseNotes = v
	return s
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) SetStatus(v string) *UpdateConnectorClientResponseBodyConnectorClient {
	s.Status = &v
	return s
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) SetUpgradeStatus(v string) *UpdateConnectorClientResponseBodyConnectorClient {
	s.UpgradeStatus = &v
	return s
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) SetVersion(v string) *UpdateConnectorClientResponseBodyConnectorClient {
	s.Version = &v
	return s
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) SetVersionToRollback(v string) *UpdateConnectorClientResponseBodyConnectorClient {
	s.VersionToRollback = &v
	return s
}

func (s *UpdateConnectorClientResponseBodyConnectorClient) Validate() error {
	return dara.Validate(s)
}
