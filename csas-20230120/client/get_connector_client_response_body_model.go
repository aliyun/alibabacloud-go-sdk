// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConnectorClientResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnectorClient(v *GetConnectorClientResponseBodyConnectorClient) *GetConnectorClientResponseBody
	GetConnectorClient() *GetConnectorClientResponseBodyConnectorClient
	SetRequestId(v string) *GetConnectorClientResponseBody
	GetRequestId() *string
}

type GetConnectorClientResponseBody struct {
	// ConnectorClient。
	ConnectorClient *GetConnectorClientResponseBodyConnectorClient `json:"ConnectorClient,omitempty" xml:"ConnectorClient,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetConnectorClientResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConnectorClientResponseBody) GoString() string {
	return s.String()
}

func (s *GetConnectorClientResponseBody) GetConnectorClient() *GetConnectorClientResponseBodyConnectorClient {
	return s.ConnectorClient
}

func (s *GetConnectorClientResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConnectorClientResponseBody) SetConnectorClient(v *GetConnectorClientResponseBodyConnectorClient) *GetConnectorClientResponseBody {
	s.ConnectorClient = v
	return s
}

func (s *GetConnectorClientResponseBody) SetRequestId(v string) *GetConnectorClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConnectorClientResponseBody) Validate() error {
	if s.ConnectorClient != nil {
		if err := s.ConnectorClient.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetConnectorClientResponseBodyConnectorClient struct {
	// The number of CPUs of the ConnectorClient.
	//
	// example:
	//
	// 16
	CPUSize *string `json:"CPUSize,omitempty" xml:"CPUSize,omitempty"`
	// The connection status of the ConnectorClient. Valid values:
	//
	// - **Connected**: Connected.
	//
	// - **Disconnected**: Disconnected.
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
	// The connector client version.
	//
	// example:
	//
	// 1.0.0
	ConnectorVersion *string `json:"ConnectorVersion,omitempty" xml:"ConnectorVersion,omitempty"`
	// The creation time of the ConnectorClient.
	//
	// example:
	//
	// 2023-05-16 17:18:46
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The unique device identifier of the ConnectorClient.
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
	// - (Empty string): Not in O&M status.
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
	// The release notes.
	ReleaseNotes []*string `json:"ReleaseNotes,omitempty" xml:"ReleaseNotes,omitempty" type:"Repeated"`
	// The enabled status of the ConnectorClient, which can be used to force the client offline. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The new version status of the connector. Valid values:
	//
	// - **Latest**: The current version is the latest.
	//
	// - **NewVersionAvailable**: A newer version is available for upgrade.
	//
	// example:
	//
	// latest
	UpgradeStatus *string `json:"UpgradeStatus,omitempty" xml:"UpgradeStatus,omitempty"`
	// The version to roll back to.
	//
	// example:
	//
	// 1.0.0
	VersionToRollback *string `json:"VersionToRollback,omitempty" xml:"VersionToRollback,omitempty"`
}

func (s GetConnectorClientResponseBodyConnectorClient) String() string {
	return dara.Prettify(s)
}

func (s GetConnectorClientResponseBodyConnectorClient) GoString() string {
	return s.String()
}

func (s *GetConnectorClientResponseBodyConnectorClient) GetCPUSize() *string {
	return s.CPUSize
}

func (s *GetConnectorClientResponseBodyConnectorClient) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *GetConnectorClientResponseBodyConnectorClient) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *GetConnectorClientResponseBodyConnectorClient) GetConnectorVersion() *string {
	return s.ConnectorVersion
}

func (s *GetConnectorClientResponseBodyConnectorClient) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetConnectorClientResponseBodyConnectorClient) GetDevTag() *string {
	return s.DevTag
}

func (s *GetConnectorClientResponseBodyConnectorClient) GetHosname() *string {
	return s.Hosname
}

func (s *GetConnectorClientResponseBodyConnectorClient) GetKernelVersion() *string {
	return s.KernelVersion
}

func (s *GetConnectorClientResponseBodyConnectorClient) GetMemorySize() *string {
	return s.MemorySize
}

func (s *GetConnectorClientResponseBodyConnectorClient) GetOperationStatus() *string {
	return s.OperationStatus
}

func (s *GetConnectorClientResponseBodyConnectorClient) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *GetConnectorClientResponseBodyConnectorClient) GetProcessRunTime() *int32 {
	return s.ProcessRunTime
}

func (s *GetConnectorClientResponseBodyConnectorClient) GetPublicIp() *string {
	return s.PublicIp
}

func (s *GetConnectorClientResponseBodyConnectorClient) GetReleaseNotes() []*string {
	return s.ReleaseNotes
}

func (s *GetConnectorClientResponseBodyConnectorClient) GetStatus() *string {
	return s.Status
}

func (s *GetConnectorClientResponseBodyConnectorClient) GetUpgradeStatus() *string {
	return s.UpgradeStatus
}

func (s *GetConnectorClientResponseBodyConnectorClient) GetVersionToRollback() *string {
	return s.VersionToRollback
}

func (s *GetConnectorClientResponseBodyConnectorClient) SetCPUSize(v string) *GetConnectorClientResponseBodyConnectorClient {
	s.CPUSize = &v
	return s
}

func (s *GetConnectorClientResponseBodyConnectorClient) SetConnectionStatus(v string) *GetConnectorClientResponseBodyConnectorClient {
	s.ConnectionStatus = &v
	return s
}

func (s *GetConnectorClientResponseBodyConnectorClient) SetConnectorId(v string) *GetConnectorClientResponseBodyConnectorClient {
	s.ConnectorId = &v
	return s
}

func (s *GetConnectorClientResponseBodyConnectorClient) SetConnectorVersion(v string) *GetConnectorClientResponseBodyConnectorClient {
	s.ConnectorVersion = &v
	return s
}

func (s *GetConnectorClientResponseBodyConnectorClient) SetCreateTime(v string) *GetConnectorClientResponseBodyConnectorClient {
	s.CreateTime = &v
	return s
}

func (s *GetConnectorClientResponseBodyConnectorClient) SetDevTag(v string) *GetConnectorClientResponseBodyConnectorClient {
	s.DevTag = &v
	return s
}

func (s *GetConnectorClientResponseBodyConnectorClient) SetHosname(v string) *GetConnectorClientResponseBodyConnectorClient {
	s.Hosname = &v
	return s
}

func (s *GetConnectorClientResponseBodyConnectorClient) SetKernelVersion(v string) *GetConnectorClientResponseBodyConnectorClient {
	s.KernelVersion = &v
	return s
}

func (s *GetConnectorClientResponseBodyConnectorClient) SetMemorySize(v string) *GetConnectorClientResponseBodyConnectorClient {
	s.MemorySize = &v
	return s
}

func (s *GetConnectorClientResponseBodyConnectorClient) SetOperationStatus(v string) *GetConnectorClientResponseBodyConnectorClient {
	s.OperationStatus = &v
	return s
}

func (s *GetConnectorClientResponseBodyConnectorClient) SetPrivateIp(v string) *GetConnectorClientResponseBodyConnectorClient {
	s.PrivateIp = &v
	return s
}

func (s *GetConnectorClientResponseBodyConnectorClient) SetProcessRunTime(v int32) *GetConnectorClientResponseBodyConnectorClient {
	s.ProcessRunTime = &v
	return s
}

func (s *GetConnectorClientResponseBodyConnectorClient) SetPublicIp(v string) *GetConnectorClientResponseBodyConnectorClient {
	s.PublicIp = &v
	return s
}

func (s *GetConnectorClientResponseBodyConnectorClient) SetReleaseNotes(v []*string) *GetConnectorClientResponseBodyConnectorClient {
	s.ReleaseNotes = v
	return s
}

func (s *GetConnectorClientResponseBodyConnectorClient) SetStatus(v string) *GetConnectorClientResponseBodyConnectorClient {
	s.Status = &v
	return s
}

func (s *GetConnectorClientResponseBodyConnectorClient) SetUpgradeStatus(v string) *GetConnectorClientResponseBodyConnectorClient {
	s.UpgradeStatus = &v
	return s
}

func (s *GetConnectorClientResponseBodyConnectorClient) SetVersionToRollback(v string) *GetConnectorClientResponseBodyConnectorClient {
	s.VersionToRollback = &v
	return s
}

func (s *GetConnectorClientResponseBodyConnectorClient) Validate() error {
	return dara.Validate(s)
}
