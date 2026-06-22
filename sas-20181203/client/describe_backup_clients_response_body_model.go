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
	// The anti-ransomware client information.
	Clients []*DescribeBackupClientsResponseBodyClients `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Repeated"`
	// The request ID. Alibaba Cloud generates a unique identifier for each API request. You can use the request ID to troubleshoot issues.
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
	if s.Clients != nil {
		for _, item := range s.Clients {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupClientsResponseBodyClients struct {
	// The anti-ransomware client ID.
	//
	// example:
	//
	// c-000az2f537r73dyh****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The status of the anti-ransomware client.
	//
	// Valid values:
	//
	// - **INSTALLING**: The client is being installed.
	//
	// - **ONLINE**: The client is online.
	//
	// - **UNINSTALLING**: The client is being uninstalled.
	//
	// - **NOT_INSTALLED**: The client is not installed.
	//
	// - **ACTIVATED**: The client is activated.
	//
	// - **CLIENT_CONNECTION_ERROR**: A client connection error occurred.
	//
	// example:
	//
	// ONLINE
	ClientStatus *string `json:"ClientStatus,omitempty" xml:"ClientStatus,omitempty"`
	// The version number of the anti-ransomware client.
	//
	// example:
	//
	// 2.0.0
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The instance ID of the ECS instance that has the anti-ransomware client installed.
	//
	// example:
	//
	// i-bp15hyph4aci99dv****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The UUID of the ECS instance that has the anti-ransomware client installed.
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
