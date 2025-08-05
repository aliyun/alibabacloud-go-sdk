// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartHanaDatabaseAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *StartHanaDatabaseAsyncRequest
	GetClusterId() *string
	SetDatabaseName(v string) *StartHanaDatabaseAsyncRequest
	GetDatabaseName() *string
	SetVaultId(v string) *StartHanaDatabaseAsyncRequest
	GetVaultId() *string
}

type StartHanaDatabaseAsyncRequest struct {
	// The ID of the SAP HANA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cl-000a9ipe******sme
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The database name.
	//
	// This parameter is required.
	//
	// example:
	//
	// HXE
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The ID of the backup vault.
	//
	// This parameter is required.
	//
	// example:
	//
	// v-0000hrh******vhr3i
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s StartHanaDatabaseAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s StartHanaDatabaseAsyncRequest) GoString() string {
	return s.String()
}

func (s *StartHanaDatabaseAsyncRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *StartHanaDatabaseAsyncRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *StartHanaDatabaseAsyncRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *StartHanaDatabaseAsyncRequest) SetClusterId(v string) *StartHanaDatabaseAsyncRequest {
	s.ClusterId = &v
	return s
}

func (s *StartHanaDatabaseAsyncRequest) SetDatabaseName(v string) *StartHanaDatabaseAsyncRequest {
	s.DatabaseName = &v
	return s
}

func (s *StartHanaDatabaseAsyncRequest) SetVaultId(v string) *StartHanaDatabaseAsyncRequest {
	s.VaultId = &v
	return s
}

func (s *StartHanaDatabaseAsyncRequest) Validate() error {
	return dara.Validate(s)
}
