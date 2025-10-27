// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupMachineStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupMachineStatus(v *DescribeBackupMachineStatusResponseBodyBackupMachineStatus) *DescribeBackupMachineStatusResponseBody
	GetBackupMachineStatus() *DescribeBackupMachineStatusResponseBodyBackupMachineStatus
	SetRequestId(v string) *DescribeBackupMachineStatusResponseBody
	GetRequestId() *string
}

type DescribeBackupMachineStatusResponseBody struct {
	// The backup status of the server.
	BackupMachineStatus *DescribeBackupMachineStatusResponseBodyBackupMachineStatus `json:"BackupMachineStatus,omitempty" xml:"BackupMachineStatus,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 09969D2C-4FAD-429E-BFBF-9A60DEF8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackupMachineStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupMachineStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupMachineStatusResponseBody) GetBackupMachineStatus() *DescribeBackupMachineStatusResponseBodyBackupMachineStatus {
	return s.BackupMachineStatus
}

func (s *DescribeBackupMachineStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupMachineStatusResponseBody) SetBackupMachineStatus(v *DescribeBackupMachineStatusResponseBodyBackupMachineStatus) *DescribeBackupMachineStatusResponseBody {
	s.BackupMachineStatus = v
	return s
}

func (s *DescribeBackupMachineStatusResponseBody) SetRequestId(v string) *DescribeBackupMachineStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupMachineStatusResponseBody) Validate() error {
	if s.BackupMachineStatus != nil {
		if err := s.BackupMachineStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBackupMachineStatusResponseBodyBackupMachineStatus struct {
	// The ID of the anti-ransomware agent.
	//
	// example:
	//
	// c-000dbefaw9f7gnbw****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The status of the anti-ransomware agent. Valid values:
	//
	// 	- **ONLINE**: normal
	//
	// 	- **CLIENT_CONNECTION_ERROR**: abnormal
	//
	// 	- **UNINSTALLING**: being uninstalled
	//
	// 	- **UNINSTALL_FAILED**: failed to be uninstalled
	//
	// 	- **UPGRADING**: being upgraded
	//
	// 	- **UPGRADE_FAILED**: failed to be upgraded
	//
	// example:
	//
	// ONLINE
	ClientStatus *string `json:"ClientStatus,omitempty" xml:"ClientStatus,omitempty"`
	// The version of the anti-ransomware agent.
	//
	// example:
	//
	// 2.11.0
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The error code returned.
	//
	// example:
	//
	// CLIENT_CONNECTION_ERROR
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// An array that consists of the error information reported by the backup server.
	ErrorList []*DescribeBackupMachineStatusResponseBodyBackupMachineStatusErrorList `json:"ErrorList,omitempty" xml:"ErrorList,omitempty" type:"Repeated"`
	// The ID of the server.
	//
	// example:
	//
	// i-2zeaqkb80vloxjcj****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region in which the server resides.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of backup versions.
	//
	// example:
	//
	// 7
	SavedBackupCount *int32 `json:"SavedBackupCount,omitempty" xml:"SavedBackupCount,omitempty"`
	// The status of the anti-ransomware service. Valid values:
	//
	// 	- **SERVICE_EXCEPTION**: Service exception
	//
	// 	- **RESTORING**: Restoring
	//
	// 	- **BACKING_UP**: Backup in process
	//
	// example:
	//
	// RESTORING
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// The status of the anti-ransomware agent. Valid values:
	//
	// 	- **NOT_INSTALLED**: not installed
	//
	// 	- **CLIENT_CONNECTION_ERROR**: abnormal
	//
	// 	- **ACTIVATED**: normal
	//
	// example:
	//
	// ACTIVATED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// eb2c782e-64f2-4590-a86c-d90164df****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The ID of the backup vault in which the backup data is stored.
	//
	// example:
	//
	// v-0005i2qh5fcr6seo****
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeBackupMachineStatusResponseBodyBackupMachineStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupMachineStatusResponseBodyBackupMachineStatus) GoString() string {
	return s.String()
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatus) GetClientId() *string {
	return s.ClientId
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatus) GetClientStatus() *string {
	return s.ClientStatus
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatus) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatus) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatus) GetErrorList() []*DescribeBackupMachineStatusResponseBodyBackupMachineStatusErrorList {
	return s.ErrorList
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatus) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatus) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatus) GetSavedBackupCount() *int32 {
	return s.SavedBackupCount
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatus) GetServiceStatus() *string {
	return s.ServiceStatus
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatus) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatus) GetVaultId() *string {
	return s.VaultId
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatus) SetClientId(v string) *DescribeBackupMachineStatusResponseBodyBackupMachineStatus {
	s.ClientId = &v
	return s
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatus) SetClientStatus(v string) *DescribeBackupMachineStatusResponseBodyBackupMachineStatus {
	s.ClientStatus = &v
	return s
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatus) SetClientVersion(v string) *DescribeBackupMachineStatusResponseBodyBackupMachineStatus {
	s.ClientVersion = &v
	return s
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatus) SetErrorCode(v string) *DescribeBackupMachineStatusResponseBodyBackupMachineStatus {
	s.ErrorCode = &v
	return s
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatus) SetErrorList(v []*DescribeBackupMachineStatusResponseBodyBackupMachineStatusErrorList) *DescribeBackupMachineStatusResponseBodyBackupMachineStatus {
	s.ErrorList = v
	return s
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatus) SetInstanceId(v string) *DescribeBackupMachineStatusResponseBodyBackupMachineStatus {
	s.InstanceId = &v
	return s
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatus) SetRegionId(v string) *DescribeBackupMachineStatusResponseBodyBackupMachineStatus {
	s.RegionId = &v
	return s
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatus) SetSavedBackupCount(v int32) *DescribeBackupMachineStatusResponseBodyBackupMachineStatus {
	s.SavedBackupCount = &v
	return s
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatus) SetServiceStatus(v string) *DescribeBackupMachineStatusResponseBodyBackupMachineStatus {
	s.ServiceStatus = &v
	return s
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatus) SetStatus(v string) *DescribeBackupMachineStatusResponseBodyBackupMachineStatus {
	s.Status = &v
	return s
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatus) SetUuid(v string) *DescribeBackupMachineStatusResponseBodyBackupMachineStatus {
	s.Uuid = &v
	return s
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatus) SetVaultId(v string) *DescribeBackupMachineStatusResponseBodyBackupMachineStatus {
	s.VaultId = &v
	return s
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatus) Validate() error {
	if s.ErrorList != nil {
		for _, item := range s.ErrorList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupMachineStatusResponseBodyBackupMachineStatusErrorList struct {
	// The error code.
	//
	// example:
	//
	// TARGET_NOT_EXIST
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// FAILED
	ErrorStatus *string `json:"ErrorStatus,omitempty" xml:"ErrorStatus,omitempty"`
}

func (s DescribeBackupMachineStatusResponseBodyBackupMachineStatusErrorList) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupMachineStatusResponseBodyBackupMachineStatusErrorList) GoString() string {
	return s.String()
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatusErrorList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatusErrorList) GetErrorStatus() *string {
	return s.ErrorStatus
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatusErrorList) SetErrorCode(v string) *DescribeBackupMachineStatusResponseBodyBackupMachineStatusErrorList {
	s.ErrorCode = &v
	return s
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatusErrorList) SetErrorStatus(v string) *DescribeBackupMachineStatusResponseBodyBackupMachineStatusErrorList {
	s.ErrorStatus = &v
	return s
}

func (s *DescribeBackupMachineStatusResponseBodyBackupMachineStatusErrorList) Validate() error {
	return dara.Validate(s)
}
