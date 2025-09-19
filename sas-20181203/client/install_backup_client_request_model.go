// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallBackupClientRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyVersion(v string) *InstallBackupClientRequest
	GetPolicyVersion() *string
	SetUuid(v string) *InstallBackupClientRequest
	GetUuid() *string
	SetUuidList(v []*string) *InstallBackupClientRequest
	GetUuidList() []*string
}

type InstallBackupClientRequest struct {
	// The version of the anti-ransomware policy. Valid values:
	//
	// 	- **1.0.0**
	//
	// 	- **2.0.0**
	//
	// This parameter is required.
	//
	// example:
	//
	// 2.0.0
	PolicyVersion *string `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
	// The UUID of the server on which you want to install the anti-ransomware agent.
	//
	// > You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of servers. You must specify at least one of the UuidList and Uuid parameters.
	//
	// example:
	//
	// inet-617eddab-7df4-4a51-b217-a3f59194****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The UUIDs of servers on which you want to install the anti-ransomware agent.
	//
	// >  You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of servers.
	//
	// example:
	//
	// ["3bb30859-b3b5-4f28-868f-b0892c98****", "3bb30859-b3b5-4f28-868f-b0892c98****"]
	UuidList []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
}

func (s InstallBackupClientRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallBackupClientRequest) GoString() string {
	return s.String()
}

func (s *InstallBackupClientRequest) GetPolicyVersion() *string {
	return s.PolicyVersion
}

func (s *InstallBackupClientRequest) GetUuid() *string {
	return s.Uuid
}

func (s *InstallBackupClientRequest) GetUuidList() []*string {
	return s.UuidList
}

func (s *InstallBackupClientRequest) SetPolicyVersion(v string) *InstallBackupClientRequest {
	s.PolicyVersion = &v
	return s
}

func (s *InstallBackupClientRequest) SetUuid(v string) *InstallBackupClientRequest {
	s.Uuid = &v
	return s
}

func (s *InstallBackupClientRequest) SetUuidList(v []*string) *InstallBackupClientRequest {
	s.UuidList = v
	return s
}

func (s *InstallBackupClientRequest) Validate() error {
	return dara.Validate(s)
}
