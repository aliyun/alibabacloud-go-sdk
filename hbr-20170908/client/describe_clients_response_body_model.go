// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClients(v *DescribeClientsResponseBodyClients) *DescribeClientsResponseBody
	GetClients() *DescribeClientsResponseBodyClients
	SetCode(v string) *DescribeClientsResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeClientsResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeClientsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeClientsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeClientsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeClientsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *DescribeClientsResponseBody
	GetTotalCount() *int32
}

type DescribeClientsResponseBody struct {
	// The Cloud Backup clients.
	Clients *DescribeClientsResponseBodyClients `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Struct"`
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. If the request was successful, "successful" is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 99. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 17189276-465D-5EF3-8FFD-0FF51B5A41A0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeClientsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClientsResponseBody) GetClients() *DescribeClientsResponseBodyClients {
	return s.Clients
}

func (s *DescribeClientsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeClientsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeClientsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeClientsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeClientsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClientsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeClientsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeClientsResponseBody) SetClients(v *DescribeClientsResponseBodyClients) *DescribeClientsResponseBody {
	s.Clients = v
	return s
}

func (s *DescribeClientsResponseBody) SetCode(v string) *DescribeClientsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeClientsResponseBody) SetMessage(v string) *DescribeClientsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeClientsResponseBody) SetPageNumber(v int32) *DescribeClientsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeClientsResponseBody) SetPageSize(v int32) *DescribeClientsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeClientsResponseBody) SetRequestId(v string) *DescribeClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClientsResponseBody) SetSuccess(v bool) *DescribeClientsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeClientsResponseBody) SetTotalCount(v int32) *DescribeClientsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeClientsResponseBody) Validate() error {
	if s.Clients != nil {
		if err := s.Clients.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClientsResponseBodyClients struct {
	Client []*DescribeClientsResponseBodyClientsClient `json:"Client,omitempty" xml:"Client,omitempty" type:"Repeated"`
}

func (s DescribeClientsResponseBodyClients) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientsResponseBodyClients) GoString() string {
	return s.String()
}

func (s *DescribeClientsResponseBodyClients) GetClient() []*DescribeClientsResponseBodyClientsClient {
	return s.Client
}

func (s *DescribeClientsResponseBodyClients) SetClient(v []*DescribeClientsResponseBodyClientsClient) *DescribeClientsResponseBodyClients {
	s.Client = v
	return s
}

func (s *DescribeClientsResponseBodyClients) Validate() error {
	if s.Client != nil {
		for _, item := range s.Client {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClientsResponseBodyClientsClient struct {
	// The alert settings. Valid value: INHERITED, which indicates that the Cloud Backup client sends alert notifications by using the same method configured for the backup vault.
	//
	// example:
	//
	// INHERITED
	AlertSetting *string `json:"AlertSetting,omitempty" xml:"AlertSetting,omitempty"`
	// The ID of the Cloud Backup client.
	//
	// example:
	//
	// c-00062uu******2fgj
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The client name.
	//
	// example:
	//
	// client-20211224-101226
	ClientName *string `json:"ClientName,omitempty" xml:"ClientName,omitempty"`
	// The type of the Cloud Backup client. Valid value: **ECS_AGENT**, which indicates an SAP HANA backup client.
	//
	// example:
	//
	// ECS_AGENT
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// The version number of the Cloud Backup client.
	//
	// example:
	//
	// 1.11.16
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The ID of the SAP HANA instance.
	//
	// example:
	//
	// cl-0008c48frr******ncpk
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The time when the Cloud Backup client was created.
	//
	// example:
	//
	// 1554347313
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The latest heartbeat time of the Cloud Backup client. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1554347313
	HeartBeatTime *int64 `json:"HeartBeatTime,omitempty" xml:"HeartBeatTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-wz9b6wya******n8yo
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the ECS instance.
	//
	// example:
	//
	// swh-hbr
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The maximum version number of the Cloud Backup client.
	//
	// example:
	//
	// 1.11.23
	MaxVersion *string `json:"MaxVersion,omitempty" xml:"MaxVersion,omitempty"`
	// The network type. Valid values:
	//
	// 	- **CLASSIC**: the classic network
	//
	// 	- **VPC**: the virtual private cloud (VPC)
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The status of the Cloud Backup client. Valid values:
	//
	// 	- **REGISTERED**: The backup client is registered.
	//
	// 	- **ACTIVATED**: The backup client is activated.
	//
	// 	- **DEACTIVATED**: The backup client fails to be activated.
	//
	// 	- **INSTALLING**: The backup client is being installed.
	//
	// 	- **INSTALL_FAILED**: The backup client fails to be installed.
	//
	// 	- **NOT_INSTALLED**: The backup client is not installed.
	//
	// 	- **UPGRADING**: The backup client is being upgraded.
	//
	// 	- **UPGRADE_FAILED**: The backup client fails to be upgraded.
	//
	// 	- **UNINSTALLING**: The backup client is being uninstalled.
	//
	// 	- **UNINSTALL_FAILED**: The backup client fails to be uninstalled.
	//
	// 	- **STOPPED**: The backup client is out of service.
	//
	// 	- **UNKNOWN**: The backup client is disconnected.
	//
	// example:
	//
	// ACTIVATED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The status information.
	//
	// example:
	//
	// HANA_NOT_SUPPORT
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	// The time when the Cloud Backup client was updated. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1554347313
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// Indicates whether data is transmitted over HTTPS. Valid values:
	//
	// 	- true: Data is transmitted over HTTPS.
	//
	// 	- false: Data is transmitted over HTTP.
	//
	// example:
	//
	// false
	UseHttps *bool `json:"UseHttps,omitempty" xml:"UseHttps,omitempty"`
	// The ID of the backup vault.
	//
	// example:
	//
	// v-00029mx6o******n85lg
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeClientsResponseBodyClientsClient) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientsResponseBodyClientsClient) GoString() string {
	return s.String()
}

func (s *DescribeClientsResponseBodyClientsClient) GetAlertSetting() *string {
	return s.AlertSetting
}

func (s *DescribeClientsResponseBodyClientsClient) GetClientId() *string {
	return s.ClientId
}

func (s *DescribeClientsResponseBodyClientsClient) GetClientName() *string {
	return s.ClientName
}

func (s *DescribeClientsResponseBodyClientsClient) GetClientType() *string {
	return s.ClientType
}

func (s *DescribeClientsResponseBodyClientsClient) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *DescribeClientsResponseBodyClientsClient) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClientsResponseBodyClientsClient) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *DescribeClientsResponseBodyClientsClient) GetHeartBeatTime() *int64 {
	return s.HeartBeatTime
}

func (s *DescribeClientsResponseBodyClientsClient) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeClientsResponseBodyClientsClient) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeClientsResponseBodyClientsClient) GetMaxVersion() *string {
	return s.MaxVersion
}

func (s *DescribeClientsResponseBodyClientsClient) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeClientsResponseBodyClientsClient) GetStatus() *string {
	return s.Status
}

func (s *DescribeClientsResponseBodyClientsClient) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *DescribeClientsResponseBodyClientsClient) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *DescribeClientsResponseBodyClientsClient) GetUseHttps() *bool {
	return s.UseHttps
}

func (s *DescribeClientsResponseBodyClientsClient) GetVaultId() *string {
	return s.VaultId
}

func (s *DescribeClientsResponseBodyClientsClient) SetAlertSetting(v string) *DescribeClientsResponseBodyClientsClient {
	s.AlertSetting = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetClientId(v string) *DescribeClientsResponseBodyClientsClient {
	s.ClientId = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetClientName(v string) *DescribeClientsResponseBodyClientsClient {
	s.ClientName = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetClientType(v string) *DescribeClientsResponseBodyClientsClient {
	s.ClientType = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetClientVersion(v string) *DescribeClientsResponseBodyClientsClient {
	s.ClientVersion = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetClusterId(v string) *DescribeClientsResponseBodyClientsClient {
	s.ClusterId = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetCreatedTime(v int64) *DescribeClientsResponseBodyClientsClient {
	s.CreatedTime = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetHeartBeatTime(v int64) *DescribeClientsResponseBodyClientsClient {
	s.HeartBeatTime = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetInstanceId(v string) *DescribeClientsResponseBodyClientsClient {
	s.InstanceId = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetInstanceName(v string) *DescribeClientsResponseBodyClientsClient {
	s.InstanceName = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetMaxVersion(v string) *DescribeClientsResponseBodyClientsClient {
	s.MaxVersion = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetNetworkType(v string) *DescribeClientsResponseBodyClientsClient {
	s.NetworkType = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetStatus(v string) *DescribeClientsResponseBodyClientsClient {
	s.Status = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetStatusMessage(v string) *DescribeClientsResponseBodyClientsClient {
	s.StatusMessage = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetUpdatedTime(v int64) *DescribeClientsResponseBodyClientsClient {
	s.UpdatedTime = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetUseHttps(v bool) *DescribeClientsResponseBodyClientsClient {
	s.UseHttps = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetVaultId(v string) *DescribeClientsResponseBodyClientsClient {
	s.VaultId = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) Validate() error {
	return dara.Validate(s)
}
