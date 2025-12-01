// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupSetDownloadRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupGatewayId(v int64) *ModifyBackupSetDownloadRulesRequest
	GetBackupGatewayId() *int64
	SetBackupPlanId(v string) *ModifyBackupSetDownloadRulesRequest
	GetBackupPlanId() *string
	SetBackupSetDownloadDir(v string) *ModifyBackupSetDownloadRulesRequest
	GetBackupSetDownloadDir() *string
	SetBackupSetDownloadTargetType(v string) *ModifyBackupSetDownloadRulesRequest
	GetBackupSetDownloadTargetType() *string
	SetBackupSetDownloadTargetTypeLocation(v string) *ModifyBackupSetDownloadRulesRequest
	GetBackupSetDownloadTargetTypeLocation() *string
	SetClientToken(v string) *ModifyBackupSetDownloadRulesRequest
	GetClientToken() *string
	SetFullDataFormat(v string) *ModifyBackupSetDownloadRulesRequest
	GetFullDataFormat() *string
	SetIncrementDataFormat(v string) *ModifyBackupSetDownloadRulesRequest
	GetIncrementDataFormat() *string
	SetOpenAutoDownload(v bool) *ModifyBackupSetDownloadRulesRequest
	GetOpenAutoDownload() *bool
	SetOwnerId(v string) *ModifyBackupSetDownloadRulesRequest
	GetOwnerId() *string
}

type ModifyBackupSetDownloadRulesRequest struct {
	// The ID of the backup gateway that is used to download the backup set.
	//
	// example:
	//
	// 2331****
	BackupGatewayId *int64 `json:"BackupGatewayId,omitempty" xml:"BackupGatewayId,omitempty"`
	// The ID of the backup schedule.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbstooi01****
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The server directory to which the backup set is downloaded.
	//
	// example:
	//
	// test
	BackupSetDownloadDir *string `json:"BackupSetDownloadDir,omitempty" xml:"BackupSetDownloadDir,omitempty"`
	// The type of the destination server to which the backup set is downloaded.
	//
	// > Set the value to agent, which indicates a server on which a backup gateway is installed.
	//
	// example:
	//
	// agent
	BackupSetDownloadTargetType *string `json:"BackupSetDownloadTargetType,omitempty" xml:"BackupSetDownloadTargetType,omitempty"`
	// The type of the destination directory to which the backup set is downloaded. This parameter is required if the automatic download feature is enabled. Valid values:
	//
	// 	- local
	//
	// 	- nas
	//
	// 	- ftp
	//
	// 	- minio
	//
	// > Default value: local.
	//
	// example:
	//
	// local
	BackupSetDownloadTargetTypeLocation *string `json:"BackupSetDownloadTargetTypeLocation,omitempty" xml:"BackupSetDownloadTargetTypeLocation,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzx****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The format in which the full backup set is downloaded. Valid values:
	//
	// 	- Native
	//
	// 	- SQL
	//
	// 	- CSV
	//
	// 	- JSON
	//
	// > Default value: CSV.
	//
	// example:
	//
	// CSV
	FullDataFormat *string `json:"FullDataFormat,omitempty" xml:"FullDataFormat,omitempty"`
	// The format in which the incremental backup set is downloaded. Valid values:
	//
	// 	- Native
	//
	// 	- SQL
	//
	// 	- CSV
	//
	// 	- JSON
	//
	// > Default value: Native.
	//
	// example:
	//
	// Native
	IncrementDataFormat *string `json:"IncrementDataFormat,omitempty" xml:"IncrementDataFormat,omitempty"`
	// Specifies whether to enable the automatic download feature. Default value: false.
	//
	// example:
	//
	// false
	OpenAutoDownload *bool   `json:"OpenAutoDownload,omitempty" xml:"OpenAutoDownload,omitempty"`
	OwnerId          *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ModifyBackupSetDownloadRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupSetDownloadRulesRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupSetDownloadRulesRequest) GetBackupGatewayId() *int64 {
	return s.BackupGatewayId
}

func (s *ModifyBackupSetDownloadRulesRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *ModifyBackupSetDownloadRulesRequest) GetBackupSetDownloadDir() *string {
	return s.BackupSetDownloadDir
}

func (s *ModifyBackupSetDownloadRulesRequest) GetBackupSetDownloadTargetType() *string {
	return s.BackupSetDownloadTargetType
}

func (s *ModifyBackupSetDownloadRulesRequest) GetBackupSetDownloadTargetTypeLocation() *string {
	return s.BackupSetDownloadTargetTypeLocation
}

func (s *ModifyBackupSetDownloadRulesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyBackupSetDownloadRulesRequest) GetFullDataFormat() *string {
	return s.FullDataFormat
}

func (s *ModifyBackupSetDownloadRulesRequest) GetIncrementDataFormat() *string {
	return s.IncrementDataFormat
}

func (s *ModifyBackupSetDownloadRulesRequest) GetOpenAutoDownload() *bool {
	return s.OpenAutoDownload
}

func (s *ModifyBackupSetDownloadRulesRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ModifyBackupSetDownloadRulesRequest) SetBackupGatewayId(v int64) *ModifyBackupSetDownloadRulesRequest {
	s.BackupGatewayId = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesRequest) SetBackupPlanId(v string) *ModifyBackupSetDownloadRulesRequest {
	s.BackupPlanId = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesRequest) SetBackupSetDownloadDir(v string) *ModifyBackupSetDownloadRulesRequest {
	s.BackupSetDownloadDir = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesRequest) SetBackupSetDownloadTargetType(v string) *ModifyBackupSetDownloadRulesRequest {
	s.BackupSetDownloadTargetType = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesRequest) SetBackupSetDownloadTargetTypeLocation(v string) *ModifyBackupSetDownloadRulesRequest {
	s.BackupSetDownloadTargetTypeLocation = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesRequest) SetClientToken(v string) *ModifyBackupSetDownloadRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesRequest) SetFullDataFormat(v string) *ModifyBackupSetDownloadRulesRequest {
	s.FullDataFormat = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesRequest) SetIncrementDataFormat(v string) *ModifyBackupSetDownloadRulesRequest {
	s.IncrementDataFormat = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesRequest) SetOpenAutoDownload(v bool) *ModifyBackupSetDownloadRulesRequest {
	s.OpenAutoDownload = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesRequest) SetOwnerId(v string) *ModifyBackupSetDownloadRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesRequest) Validate() error {
	return dara.Validate(s)
}
