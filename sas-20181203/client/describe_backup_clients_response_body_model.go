// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupClientsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClients(v []*DescribeBackupClientsResponseBodyClients) *DescribeBackupClientsResponseBody
	GetClients() []*DescribeBackupClientsResponseBodyClients
	SetRequestId(v string) *DescribeBackupClientsResponseBody
	GetRequestId() *string
}

type DescribeBackupClientsResponseBody struct {
	// An array that consists of the information about the anti-ransomware agent.
	Clients []*DescribeBackupClientsResponseBodyClients `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// E3ED094C-9EB7-4239-962B-D0FB3D5F23C7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackupClientsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupClientsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupClientsResponseBody) GetClients() []*DescribeBackupClientsResponseBodyClients {
	return s.Clients
}

func (s *DescribeBackupClientsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupClientsResponseBody) SetClients(v []*DescribeBackupClientsResponseBodyClients) *DescribeBackupClientsResponseBody {
	s.Clients = v
	return s
}

func (s *DescribeBackupClientsResponseBody) SetRequestId(v string) *DescribeBackupClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupClientsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupClientsResponseBodyClients struct {
	// The ID of the anti-ransomware agent.
	//
	// example:
	//
	// c-000az2f537r73dyh****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The status of the anti-ransomware agent.
	//
	// Valid values:
	//
	// 	- **INSTALLING**: The agent is being installed.
	//
	// 	- **ONLINE**: The agent is online.
	//
	// 	- **UNINSTALLING**: The agent is being uninstalled.
	//
	// 	- **NOT_INSTALLED**: The agent is not installed.
	//
	// 	- **ACTIVATED**: The agent is enabled.
	//
	// 	- **CLIENT_CONNECTION_ERROR**: A connection error occurs on the agent.
	//
	// example:
	//
	// ONLINE
	ClientStatus *string `json:"ClientStatus,omitempty" xml:"ClientStatus,omitempty"`
	// The version of the anti-ransomware agent.
	//
	// example:
	//
	// 2.0.0
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The ID of the ECS instance on which the anti-ransomware agent is installed.
	//
	// example:
	//
	// i-bp15hyph4aci99dv****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The UUID of the Elastic Compute Service (ECS) instance on which the anti-ransomware agent is installed.
	//
	// example:
	//
	// 22f6550d-f294-449b-b6e6-90638fd1****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeBackupClientsResponseBodyClients) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupClientsResponseBodyClients) GoString() string {
	return s.String()
}

func (s *DescribeBackupClientsResponseBodyClients) GetClientId() *string {
	return s.ClientId
}

func (s *DescribeBackupClientsResponseBodyClients) GetClientStatus() *string {
	return s.ClientStatus
}

func (s *DescribeBackupClientsResponseBodyClients) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *DescribeBackupClientsResponseBodyClients) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeBackupClientsResponseBodyClients) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeBackupClientsResponseBodyClients) SetClientId(v string) *DescribeBackupClientsResponseBodyClients {
	s.ClientId = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetClientStatus(v string) *DescribeBackupClientsResponseBodyClients {
	s.ClientStatus = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetClientVersion(v string) *DescribeBackupClientsResponseBodyClients {
	s.ClientVersion = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetInstanceId(v string) *DescribeBackupClientsResponseBodyClients {
	s.InstanceId = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetUuid(v string) *DescribeBackupClientsResponseBodyClients {
	s.Uuid = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) Validate() error {
	return dara.Validate(s)
}
