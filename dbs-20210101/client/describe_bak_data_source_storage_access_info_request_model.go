// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBakDataSourceStorageAccessInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupSetId(v string) *DescribeBakDataSourceStorageAccessInfoRequest
	GetBackupSetId() *string
	SetBackupType(v string) *DescribeBakDataSourceStorageAccessInfoRequest
	GetBackupType() *string
	SetClientToken(v string) *DescribeBakDataSourceStorageAccessInfoRequest
	GetClientToken() *string
	SetDataSourceId(v string) *DescribeBakDataSourceStorageAccessInfoRequest
	GetDataSourceId() *string
	SetDurationSeconds(v int64) *DescribeBakDataSourceStorageAccessInfoRequest
	GetDurationSeconds() *int64
	SetRegionCode(v string) *DescribeBakDataSourceStorageAccessInfoRequest
	GetRegionCode() *string
	SetRegionId(v string) *DescribeBakDataSourceStorageAccessInfoRequest
	GetRegionId() *string
	SetStorageType(v string) *DescribeBakDataSourceStorageAccessInfoRequest
	GetStorageType() *string
}

type DescribeBakDataSourceStorageAccessInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ee-d84wwm3cazdbjli1m1*****
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FullBackup
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// example:
	//
	// dbs
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ds-7iz7uwzgcgumznkd02xn*****
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// 3600
	DurationSeconds *int64 `json:"DurationSeconds,omitempty" xml:"DurationSeconds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeBakDataSourceStorageAccessInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBakDataSourceStorageAccessInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeBakDataSourceStorageAccessInfoRequest) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *DescribeBakDataSourceStorageAccessInfoRequest) GetBackupType() *string {
	return s.BackupType
}

func (s *DescribeBakDataSourceStorageAccessInfoRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeBakDataSourceStorageAccessInfoRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *DescribeBakDataSourceStorageAccessInfoRequest) GetDurationSeconds() *int64 {
	return s.DurationSeconds
}

func (s *DescribeBakDataSourceStorageAccessInfoRequest) GetRegionCode() *string {
	return s.RegionCode
}

func (s *DescribeBakDataSourceStorageAccessInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBakDataSourceStorageAccessInfoRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeBakDataSourceStorageAccessInfoRequest) SetBackupSetId(v string) *DescribeBakDataSourceStorageAccessInfoRequest {
	s.BackupSetId = &v
	return s
}

func (s *DescribeBakDataSourceStorageAccessInfoRequest) SetBackupType(v string) *DescribeBakDataSourceStorageAccessInfoRequest {
	s.BackupType = &v
	return s
}

func (s *DescribeBakDataSourceStorageAccessInfoRequest) SetClientToken(v string) *DescribeBakDataSourceStorageAccessInfoRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeBakDataSourceStorageAccessInfoRequest) SetDataSourceId(v string) *DescribeBakDataSourceStorageAccessInfoRequest {
	s.DataSourceId = &v
	return s
}

func (s *DescribeBakDataSourceStorageAccessInfoRequest) SetDurationSeconds(v int64) *DescribeBakDataSourceStorageAccessInfoRequest {
	s.DurationSeconds = &v
	return s
}

func (s *DescribeBakDataSourceStorageAccessInfoRequest) SetRegionCode(v string) *DescribeBakDataSourceStorageAccessInfoRequest {
	s.RegionCode = &v
	return s
}

func (s *DescribeBakDataSourceStorageAccessInfoRequest) SetRegionId(v string) *DescribeBakDataSourceStorageAccessInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBakDataSourceStorageAccessInfoRequest) SetStorageType(v string) *DescribeBakDataSourceStorageAccessInfoRequest {
	s.StorageType = &v
	return s
}

func (s *DescribeBakDataSourceStorageAccessInfoRequest) Validate() error {
	return dara.Validate(s)
}
