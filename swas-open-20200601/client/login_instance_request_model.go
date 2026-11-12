// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoginInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *LoginInstanceRequest
	GetInstanceId() *string
	SetPassword(v string) *LoginInstanceRequest
	GetPassword() *string
	SetPort(v int32) *LoginInstanceRequest
	GetPort() *int32
	SetRegionId(v string) *LoginInstanceRequest
	GetRegionId() *string
	SetUsername(v string) *LoginInstanceRequest
	GetUsername() *string
}

type LoginInstanceRequest struct {
	// The instance ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9ae7106e68eb4402b0dcbd48a9de****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The password associated with the username of the simple application server.
	//
	// - Linux server: No password is required.
	//
	// - Windows server: Enter the configured password. If you have not set a password after purchasing the simple application server, see [Reset the password](https://help.aliyun.com/document_detail/60055.html).
	//
	// example:
	//
	// Test****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The port number for remote logon to the instance.
	//
	// - Linux server: The default port is 22.
	//
	// - Windows server: The default port is 3389.
	//
	// > If you want to connect to the server by using a custom port, you must first change the default remote port of the server. For more information, see [Settings for custom remote connection ports](https://help.aliyun.com/document_detail/2807402.html).
	//
	// example:
	//
	// 3389
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region ID of the specified Simple Application Server instance. You can call [ListRegions](https://help.aliyun.com/document_detail/189315.html) to query the list of supported Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The username for the simple application server.
	//
	// - You do not need to enter a username for a Linux server.
	//
	// - The default username for a Windows server is `administrator`.
	//
	// example:
	//
	// administrator
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s LoginInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s LoginInstanceRequest) GoString() string {
	return s.String()
}

func (s *LoginInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *LoginInstanceRequest) GetPassword() *string {
	return s.Password
}

func (s *LoginInstanceRequest) GetPort() *int32 {
	return s.Port
}

func (s *LoginInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *LoginInstanceRequest) GetUsername() *string {
	return s.Username
}

func (s *LoginInstanceRequest) SetInstanceId(v string) *LoginInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *LoginInstanceRequest) SetPassword(v string) *LoginInstanceRequest {
	s.Password = &v
	return s
}

func (s *LoginInstanceRequest) SetPort(v int32) *LoginInstanceRequest {
	s.Port = &v
	return s
}

func (s *LoginInstanceRequest) SetRegionId(v string) *LoginInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *LoginInstanceRequest) SetUsername(v string) *LoginInstanceRequest {
	s.Username = &v
	return s
}

func (s *LoginInstanceRequest) Validate() error {
	return dara.Validate(s)
}
