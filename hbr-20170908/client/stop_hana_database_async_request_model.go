// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopHanaDatabaseAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *StopHanaDatabaseAsyncRequest
	GetClusterId() *string
	SetDatabaseName(v string) *StopHanaDatabaseAsyncRequest
	GetDatabaseName() *string
	SetVaultId(v string) *StopHanaDatabaseAsyncRequest
	GetVaultId() *string
}

type StopHanaDatabaseAsyncRequest struct {
	// The ID of the SAP HANA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cl-000dw******45ijer
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The database name.
	//
	// This parameter is required.
	//
	// example:
	//
	// BWP
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The ID of the backup vault.
	//
	// This parameter is required.
	//
	// example:
	//
	// v-0006wkn******gzkn
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s StopHanaDatabaseAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s StopHanaDatabaseAsyncRequest) GoString() string {
	return s.String()
}

func (s *StopHanaDatabaseAsyncRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *StopHanaDatabaseAsyncRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *StopHanaDatabaseAsyncRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *StopHanaDatabaseAsyncRequest) SetClusterId(v string) *StopHanaDatabaseAsyncRequest {
	s.ClusterId = &v
	return s
}

func (s *StopHanaDatabaseAsyncRequest) SetDatabaseName(v string) *StopHanaDatabaseAsyncRequest {
	s.DatabaseName = &v
	return s
}

func (s *StopHanaDatabaseAsyncRequest) SetVaultId(v string) *StopHanaDatabaseAsyncRequest {
	s.VaultId = &v
	return s
}

func (s *StopHanaDatabaseAsyncRequest) Validate() error {
	return dara.Validate(s)
}
