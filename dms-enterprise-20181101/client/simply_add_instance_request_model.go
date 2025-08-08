// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSimplyAddInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabasePassword(v string) *SimplyAddInstanceRequest
	GetDatabasePassword() *string
	SetDatabaseUser(v string) *SimplyAddInstanceRequest
	GetDatabaseUser() *string
	SetHost(v string) *SimplyAddInstanceRequest
	GetHost() *string
	SetInstanceId(v string) *SimplyAddInstanceRequest
	GetInstanceId() *string
	SetInstanceRegion(v string) *SimplyAddInstanceRequest
	GetInstanceRegion() *string
	SetPort(v int32) *SimplyAddInstanceRequest
	GetPort() *int32
	SetRealLoginUserUid(v string) *SimplyAddInstanceRequest
	GetRealLoginUserUid() *string
}

type SimplyAddInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test***
	DatabasePassword *string `json:"DatabasePassword,omitempty" xml:"DatabasePassword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abc
	DatabaseUser *string `json:"DatabaseUser,omitempty" xml:"DatabaseUser,omitempty"`
	// example:
	//
	// 192.XXX.0.56
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// rm****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	InstanceRegion *string `json:"InstanceRegion,omitempty" xml:"InstanceRegion,omitempty"`
	// example:
	//
	// 5432
	Port             *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	RealLoginUserUid *string `json:"RealLoginUserUid,omitempty" xml:"RealLoginUserUid,omitempty"`
}

func (s SimplyAddInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SimplyAddInstanceRequest) GoString() string {
	return s.String()
}

func (s *SimplyAddInstanceRequest) GetDatabasePassword() *string {
	return s.DatabasePassword
}

func (s *SimplyAddInstanceRequest) GetDatabaseUser() *string {
	return s.DatabaseUser
}

func (s *SimplyAddInstanceRequest) GetHost() *string {
	return s.Host
}

func (s *SimplyAddInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SimplyAddInstanceRequest) GetInstanceRegion() *string {
	return s.InstanceRegion
}

func (s *SimplyAddInstanceRequest) GetPort() *int32 {
	return s.Port
}

func (s *SimplyAddInstanceRequest) GetRealLoginUserUid() *string {
	return s.RealLoginUserUid
}

func (s *SimplyAddInstanceRequest) SetDatabasePassword(v string) *SimplyAddInstanceRequest {
	s.DatabasePassword = &v
	return s
}

func (s *SimplyAddInstanceRequest) SetDatabaseUser(v string) *SimplyAddInstanceRequest {
	s.DatabaseUser = &v
	return s
}

func (s *SimplyAddInstanceRequest) SetHost(v string) *SimplyAddInstanceRequest {
	s.Host = &v
	return s
}

func (s *SimplyAddInstanceRequest) SetInstanceId(v string) *SimplyAddInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *SimplyAddInstanceRequest) SetInstanceRegion(v string) *SimplyAddInstanceRequest {
	s.InstanceRegion = &v
	return s
}

func (s *SimplyAddInstanceRequest) SetPort(v int32) *SimplyAddInstanceRequest {
	s.Port = &v
	return s
}

func (s *SimplyAddInstanceRequest) SetRealLoginUserUid(v string) *SimplyAddInstanceRequest {
	s.RealLoginUserUid = &v
	return s
}

func (s *SimplyAddInstanceRequest) Validate() error {
	return dara.Validate(s)
}
