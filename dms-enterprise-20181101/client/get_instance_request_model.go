// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHost(v string) *GetInstanceRequest
	GetHost() *string
	SetPort(v int32) *GetInstanceRequest
	GetPort() *int32
	SetRealLoginUserUid(v string) *GetInstanceRequest
	GetRealLoginUserUid() *string
	SetSid(v string) *GetInstanceRequest
	GetSid() *string
	SetTid(v int64) *GetInstanceRequest
	GetTid() *int64
}

type GetInstanceRequest struct {
	// The endpoint of the database instance. You can call the [ListInstances](https://help.aliyun.com/document_detail/141936.html) operation to obtain the endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.17.XXX.XXX
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The port number that is used to connect to the database instance. You can call the [ListInstances](https://help.aliyun.com/document_detail/141936.html) operation to obtain the port number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5432
	Port             *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	RealLoginUserUid *string `json:"RealLoginUserUid,omitempty" xml:"RealLoginUserUid,omitempty"`
	// The system ID (SID) of the database instance. You can call the [ListInstances](https://help.aliyun.com/document_detail/141936.html) operation to obtain the SID.
	//
	// example:
	//
	// test
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 2***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceRequest) GetHost() *string {
	return s.Host
}

func (s *GetInstanceRequest) GetPort() *int32 {
	return s.Port
}

func (s *GetInstanceRequest) GetRealLoginUserUid() *string {
	return s.RealLoginUserUid
}

func (s *GetInstanceRequest) GetSid() *string {
	return s.Sid
}

func (s *GetInstanceRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetInstanceRequest) SetHost(v string) *GetInstanceRequest {
	s.Host = &v
	return s
}

func (s *GetInstanceRequest) SetPort(v int32) *GetInstanceRequest {
	s.Port = &v
	return s
}

func (s *GetInstanceRequest) SetRealLoginUserUid(v string) *GetInstanceRequest {
	s.RealLoginUserUid = &v
	return s
}

func (s *GetInstanceRequest) SetSid(v string) *GetInstanceRequest {
	s.Sid = &v
	return s
}

func (s *GetInstanceRequest) SetTid(v int64) *GetInstanceRequest {
	s.Tid = &v
	return s
}

func (s *GetInstanceRequest) Validate() error {
	return dara.Validate(s)
}
