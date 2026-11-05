// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallBackupClientRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyVersion(v string) *UninstallBackupClientRequest
	GetPolicyVersion() *string
	SetUuid(v string) *UninstallBackupClientRequest
	GetUuid() *string
	SetUuidList(v []*string) *UninstallBackupClientRequest
	GetUuidList() []*string
}

type UninstallBackupClientRequest struct {
	// The version of the protection policy. You can call the [DescribeBackupPolicies](~~DescribeBackupPolicies~~) operation to query the version. Valid values:
	//
	// - **1.0.0**: The version of the protection policy is 1.0.0.
	//
	// - **2.0.0**: The version of the protection policy is 2.0.0.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2.0.0
	PolicyVersion *string `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
	// The UUID of the server from which you want to uninstall the anti-ransomware client.
	//
	// > UuidList.N and UUID cannot both be empty.
	//
	// example:
	//
	// D0D6E6E4-CB8C-4897-B852-46AEFDA0****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The list of UUIDs of the servers from which you want to uninstall the anti-ransomware client.
	//
	// > UuidList and UUID cannot both be empty.
	//
	// example:
	//
	// ["D0D6E6E4-CB8C-4897-B852-46AEFDA0****", "3bb30859-b3b5-4f28-868f-b0892c98****"]
	UuidList []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
}

func (s UninstallBackupClientRequest) String() string {
	return dara.Prettify(s)
}

func (s UninstallBackupClientRequest) GoString() string {
	return s.String()
}

func (s *UninstallBackupClientRequest) GetPolicyVersion() *string {
	return s.PolicyVersion
}

func (s *UninstallBackupClientRequest) GetUuid() *string {
	return s.Uuid
}

func (s *UninstallBackupClientRequest) GetUuidList() []*string {
	return s.UuidList
}

func (s *UninstallBackupClientRequest) SetPolicyVersion(v string) *UninstallBackupClientRequest {
	s.PolicyVersion = &v
	return s
}

func (s *UninstallBackupClientRequest) SetUuid(v string) *UninstallBackupClientRequest {
	s.Uuid = &v
	return s
}

func (s *UninstallBackupClientRequest) SetUuidList(v []*string) *UninstallBackupClientRequest {
	s.UuidList = v
	return s
}

func (s *UninstallBackupClientRequest) Validate() error {
	return dara.Validate(s)
}
