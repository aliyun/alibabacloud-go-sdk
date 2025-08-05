// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHanaRetentionSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeHanaRetentionSettingRequest
	GetClusterId() *string
	SetDatabaseName(v string) *DescribeHanaRetentionSettingRequest
	GetDatabaseName() *string
	SetVaultId(v string) *DescribeHanaRetentionSettingRequest
	GetVaultId() *string
}

type DescribeHanaRetentionSettingRequest struct {
	// The ID of the SAP HANA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cl-0002ys1i******wwtf
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The database name.
	//
	// example:
	//
	// C4P
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The ID of the backup vault.
	//
	// This parameter is required.
	//
	// example:
	//
	// v-000ii8t******ntrt2
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeHanaRetentionSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaRetentionSettingRequest) GoString() string {
	return s.String()
}

func (s *DescribeHanaRetentionSettingRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeHanaRetentionSettingRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeHanaRetentionSettingRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *DescribeHanaRetentionSettingRequest) SetClusterId(v string) *DescribeHanaRetentionSettingRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeHanaRetentionSettingRequest) SetDatabaseName(v string) *DescribeHanaRetentionSettingRequest {
	s.DatabaseName = &v
	return s
}

func (s *DescribeHanaRetentionSettingRequest) SetVaultId(v string) *DescribeHanaRetentionSettingRequest {
	s.VaultId = &v
	return s
}

func (s *DescribeHanaRetentionSettingRequest) Validate() error {
	return dara.Validate(s)
}
