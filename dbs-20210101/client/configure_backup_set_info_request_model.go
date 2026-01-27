// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureBackupSetInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupMethod(v string) *ConfigureBackupSetInfoRequest
	GetBackupMethod() *string
	SetBackupMode(v string) *ConfigureBackupSetInfoRequest
	GetBackupMode() *string
	SetBackupSetId(v string) *ConfigureBackupSetInfoRequest
	GetBackupSetId() *string
	SetBackupSetName(v string) *ConfigureBackupSetInfoRequest
	GetBackupSetName() *string
	SetBackupType(v string) *ConfigureBackupSetInfoRequest
	GetBackupType() *string
	SetCheckSum(v string) *ConfigureBackupSetInfoRequest
	GetCheckSum() *string
	SetClientToken(v string) *ConfigureBackupSetInfoRequest
	GetClientToken() *string
	SetDataSourceId(v string) *ConfigureBackupSetInfoRequest
	GetDataSourceId() *string
	SetExtraMeta(v string) *ConfigureBackupSetInfoRequest
	GetExtraMeta() *string
	SetRegionCode(v string) *ConfigureBackupSetInfoRequest
	GetRegionCode() *string
	SetRegionId(v string) *ConfigureBackupSetInfoRequest
	GetRegionId() *string
	SetUploadStatus(v string) *ConfigureBackupSetInfoRequest
	GetUploadStatus() *string
}

type ConfigureBackupSetInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Physical
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Automated
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// example:
	//
	// ee-d84wwm3c****
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// This parameter is required.
	BackupSetName *string `json:"BackupSetName,omitempty" xml:"BackupSetName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FullBackup
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// example:
	//
	// ****
	CheckSum *string `json:"CheckSum,omitempty" xml:"CheckSum,omitempty"`
	// example:
	//
	// dbs
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ds-7iz7uw****
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	ExtraMeta    *string `json:"ExtraMeta,omitempty" xml:"ExtraMeta,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// WaitingUpload
	UploadStatus *string `json:"UploadStatus,omitempty" xml:"UploadStatus,omitempty"`
}

func (s ConfigureBackupSetInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigureBackupSetInfoRequest) GoString() string {
	return s.String()
}

func (s *ConfigureBackupSetInfoRequest) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *ConfigureBackupSetInfoRequest) GetBackupMode() *string {
	return s.BackupMode
}

func (s *ConfigureBackupSetInfoRequest) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *ConfigureBackupSetInfoRequest) GetBackupSetName() *string {
	return s.BackupSetName
}

func (s *ConfigureBackupSetInfoRequest) GetBackupType() *string {
	return s.BackupType
}

func (s *ConfigureBackupSetInfoRequest) GetCheckSum() *string {
	return s.CheckSum
}

func (s *ConfigureBackupSetInfoRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ConfigureBackupSetInfoRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *ConfigureBackupSetInfoRequest) GetExtraMeta() *string {
	return s.ExtraMeta
}

func (s *ConfigureBackupSetInfoRequest) GetRegionCode() *string {
	return s.RegionCode
}

func (s *ConfigureBackupSetInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConfigureBackupSetInfoRequest) GetUploadStatus() *string {
	return s.UploadStatus
}

func (s *ConfigureBackupSetInfoRequest) SetBackupMethod(v string) *ConfigureBackupSetInfoRequest {
	s.BackupMethod = &v
	return s
}

func (s *ConfigureBackupSetInfoRequest) SetBackupMode(v string) *ConfigureBackupSetInfoRequest {
	s.BackupMode = &v
	return s
}

func (s *ConfigureBackupSetInfoRequest) SetBackupSetId(v string) *ConfigureBackupSetInfoRequest {
	s.BackupSetId = &v
	return s
}

func (s *ConfigureBackupSetInfoRequest) SetBackupSetName(v string) *ConfigureBackupSetInfoRequest {
	s.BackupSetName = &v
	return s
}

func (s *ConfigureBackupSetInfoRequest) SetBackupType(v string) *ConfigureBackupSetInfoRequest {
	s.BackupType = &v
	return s
}

func (s *ConfigureBackupSetInfoRequest) SetCheckSum(v string) *ConfigureBackupSetInfoRequest {
	s.CheckSum = &v
	return s
}

func (s *ConfigureBackupSetInfoRequest) SetClientToken(v string) *ConfigureBackupSetInfoRequest {
	s.ClientToken = &v
	return s
}

func (s *ConfigureBackupSetInfoRequest) SetDataSourceId(v string) *ConfigureBackupSetInfoRequest {
	s.DataSourceId = &v
	return s
}

func (s *ConfigureBackupSetInfoRequest) SetExtraMeta(v string) *ConfigureBackupSetInfoRequest {
	s.ExtraMeta = &v
	return s
}

func (s *ConfigureBackupSetInfoRequest) SetRegionCode(v string) *ConfigureBackupSetInfoRequest {
	s.RegionCode = &v
	return s
}

func (s *ConfigureBackupSetInfoRequest) SetRegionId(v string) *ConfigureBackupSetInfoRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigureBackupSetInfoRequest) SetUploadStatus(v string) *ConfigureBackupSetInfoRequest {
	s.UploadStatus = &v
	return s
}

func (s *ConfigureBackupSetInfoRequest) Validate() error {
	return dara.Validate(s)
}
