// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRequestDiagnosisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabase(v string) *CreateRequestDiagnosisRequest
	GetDatabase() *string
	SetInstanceId(v string) *CreateRequestDiagnosisRequest
	GetInstanceId() *string
	SetNodeId(v string) *CreateRequestDiagnosisRequest
	GetNodeId() *string
	SetSql(v string) *CreateRequestDiagnosisRequest
	GetSql() *string
}

type CreateRequestDiagnosisRequest struct {
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// das
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-0iwhhl8gx0ld6****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The node ID.
	//
	// >  This parameter must be specified for PolarDB for MySQL, PolarDB for PostgreSQL (Compatible with Oracle), and ApsaraDB for MongoDB instances.
	//
	// example:
	//
	// 202****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The SQL statement that you want to diagnose.
	//
	// This parameter is required.
	//
	// example:
	//
	// select 	- from test where name = \\"mockUser\\"
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
}

func (s CreateRequestDiagnosisRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRequestDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *CreateRequestDiagnosisRequest) GetDatabase() *string {
	return s.Database
}

func (s *CreateRequestDiagnosisRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateRequestDiagnosisRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateRequestDiagnosisRequest) GetSql() *string {
	return s.Sql
}

func (s *CreateRequestDiagnosisRequest) SetDatabase(v string) *CreateRequestDiagnosisRequest {
	s.Database = &v
	return s
}

func (s *CreateRequestDiagnosisRequest) SetInstanceId(v string) *CreateRequestDiagnosisRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRequestDiagnosisRequest) SetNodeId(v string) *CreateRequestDiagnosisRequest {
	s.NodeId = &v
	return s
}

func (s *CreateRequestDiagnosisRequest) SetSql(v string) *CreateRequestDiagnosisRequest {
	s.Sql = &v
	return s
}

func (s *CreateRequestDiagnosisRequest) Validate() error {
	return dara.Validate(s)
}
