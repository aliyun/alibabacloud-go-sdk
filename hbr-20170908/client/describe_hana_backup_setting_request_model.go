// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHanaBackupSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeHanaBackupSettingRequest
	GetClusterId() *string
	SetDatabaseName(v string) *DescribeHanaBackupSettingRequest
	GetDatabaseName() *string
	SetVaultId(v string) *DescribeHanaBackupSettingRequest
	GetVaultId() *string
}

type DescribeHanaBackupSettingRequest struct {
	// The ID of the SAP HANA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cl-00068btz******oku
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// SYSTEMDB
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The ID of the backup vault.
	//
	// example:
	//
	// v-000css******za5uyqr
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeHanaBackupSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaBackupSettingRequest) GoString() string {
	return s.String()
}

func (s *DescribeHanaBackupSettingRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeHanaBackupSettingRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeHanaBackupSettingRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *DescribeHanaBackupSettingRequest) SetClusterId(v string) *DescribeHanaBackupSettingRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeHanaBackupSettingRequest) SetDatabaseName(v string) *DescribeHanaBackupSettingRequest {
	s.DatabaseName = &v
	return s
}

func (s *DescribeHanaBackupSettingRequest) SetVaultId(v string) *DescribeHanaBackupSettingRequest {
	s.VaultId = &v
	return s
}

func (s *DescribeHanaBackupSettingRequest) Validate() error {
	return dara.Validate(s)
}
