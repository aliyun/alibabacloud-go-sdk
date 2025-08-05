// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHanaInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteHanaInstanceRequest
	GetClusterId() *string
	SetResourceGroupId(v string) *DeleteHanaInstanceRequest
	GetResourceGroupId() *string
	SetSid(v string) *DeleteHanaInstanceRequest
	GetSid() *string
	SetVaultId(v string) *DeleteHanaInstanceRequest
	GetVaultId() *string
}

type DeleteHanaInstanceRequest struct {
	// The ID of the SAP HANA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cl-000g4z09******9cfc
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmwutpyat2kwy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The security identifier (SID) of the SAP HANA database. You must specify a valid SID. The SID must be three characters in length and start with a letter. For more information, see [How to find sid user and instance number of HANA db?](https://answers.sap.com/questions/555192/how-to-find-sid-user-and-instance-number-of-hana-d.html?)
	//
	// This parameter is required.
	//
	// example:
	//
	// HXE
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// The ID of the backup vault.
	//
	// This parameter is required.
	//
	// example:
	//
	// v-0008n2******ax3
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DeleteHanaInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHanaInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteHanaInstanceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteHanaInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteHanaInstanceRequest) GetSid() *string {
	return s.Sid
}

func (s *DeleteHanaInstanceRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *DeleteHanaInstanceRequest) SetClusterId(v string) *DeleteHanaInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteHanaInstanceRequest) SetResourceGroupId(v string) *DeleteHanaInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteHanaInstanceRequest) SetSid(v string) *DeleteHanaInstanceRequest {
	s.Sid = &v
	return s
}

func (s *DeleteHanaInstanceRequest) SetVaultId(v string) *DeleteHanaInstanceRequest {
	s.VaultId = &v
	return s
}

func (s *DeleteHanaInstanceRequest) Validate() error {
	return dara.Validate(s)
}
