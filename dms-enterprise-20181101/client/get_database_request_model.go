// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHost(v string) *GetDatabaseRequest
	GetHost() *string
	SetPort(v int32) *GetDatabaseRequest
	GetPort() *int32
	SetRealLoginUserUid(v string) *GetDatabaseRequest
	GetRealLoginUserUid() *string
	SetSchemaName(v string) *GetDatabaseRequest
	GetSchemaName() *string
	SetSid(v string) *GetDatabaseRequest
	GetSid() *string
	SetTid(v int64) *GetDatabaseRequest
	GetTid() *int64
}

type GetDatabaseRequest struct {
	// The endpoint that is used to connect to the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.XX.XX
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The port that is used to connect to the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port             *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	RealLoginUserUid *string `json:"RealLoginUserUid,omitempty" xml:"RealLoginUserUid,omitempty"`
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The system identifier (SID) of the database.
	//
	// >  The SID uniquely identifies an Oracle database. After a database is created, a SID is generated for the database.
	//
	// example:
	//
	// test_sid
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to query the ID of the tenant.
	//
	// example:
	//
	// -1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseRequest) GoString() string {
	return s.String()
}

func (s *GetDatabaseRequest) GetHost() *string {
	return s.Host
}

func (s *GetDatabaseRequest) GetPort() *int32 {
	return s.Port
}

func (s *GetDatabaseRequest) GetRealLoginUserUid() *string {
	return s.RealLoginUserUid
}

func (s *GetDatabaseRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetDatabaseRequest) GetSid() *string {
	return s.Sid
}

func (s *GetDatabaseRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetDatabaseRequest) SetHost(v string) *GetDatabaseRequest {
	s.Host = &v
	return s
}

func (s *GetDatabaseRequest) SetPort(v int32) *GetDatabaseRequest {
	s.Port = &v
	return s
}

func (s *GetDatabaseRequest) SetRealLoginUserUid(v string) *GetDatabaseRequest {
	s.RealLoginUserUid = &v
	return s
}

func (s *GetDatabaseRequest) SetSchemaName(v string) *GetDatabaseRequest {
	s.SchemaName = &v
	return s
}

func (s *GetDatabaseRequest) SetSid(v string) *GetDatabaseRequest {
	s.Sid = &v
	return s
}

func (s *GetDatabaseRequest) SetTid(v int64) *GetDatabaseRequest {
	s.Tid = &v
	return s
}

func (s *GetDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
