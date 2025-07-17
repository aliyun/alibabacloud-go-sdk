// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHost(v string) *DeleteInstanceRequest
	GetHost() *string
	SetPort(v int32) *DeleteInstanceRequest
	GetPort() *int32
	SetSid(v string) *DeleteInstanceRequest
	GetSid() *string
	SetTid(v int64) *DeleteInstanceRequest
	GetTid() *int64
}

type DeleteInstanceRequest struct {
	// The endpoint of the database instance. You can call the [ListInstances](https://help.aliyun.com/document_detail/141936.html) or [GetInstance](https://help.aliyun.com/document_detail/141567.html) operation to obtain the endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx.mysql.rds.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The port number that is used to connect to the database instance. You can call the [ListInstances](https://help.aliyun.com/document_detail/141936.html) or [GetInstance](https://help.aliyun.com/document_detail/141567.html) operation to obtain the port number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The system ID (SID) of the database instance. You can call the [ListInstances](https://help.aliyun.com/document_detail/141936.html) or [GetInstance](https://help.aliyun.com/document_detail/141567.html) operation to obtain the SID.
	//
	// example:
	//
	// testSid
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s DeleteInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) GetHost() *string {
	return s.Host
}

func (s *DeleteInstanceRequest) GetPort() *int32 {
	return s.Port
}

func (s *DeleteInstanceRequest) GetSid() *string {
	return s.Sid
}

func (s *DeleteInstanceRequest) GetTid() *int64 {
	return s.Tid
}

func (s *DeleteInstanceRequest) SetHost(v string) *DeleteInstanceRequest {
	s.Host = &v
	return s
}

func (s *DeleteInstanceRequest) SetPort(v int32) *DeleteInstanceRequest {
	s.Port = &v
	return s
}

func (s *DeleteInstanceRequest) SetSid(v string) *DeleteInstanceRequest {
	s.Sid = &v
	return s
}

func (s *DeleteInstanceRequest) SetTid(v int64) *DeleteInstanceRequest {
	s.Tid = &v
	return s
}

func (s *DeleteInstanceRequest) Validate() error {
	return dara.Validate(s)
}
