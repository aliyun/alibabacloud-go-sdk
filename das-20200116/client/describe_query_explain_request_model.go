// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQueryExplainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *DescribeQueryExplainRequest
	GetDbName() *string
	SetInstanceId(v string) *DescribeQueryExplainRequest
	GetInstanceId() *string
	SetNodeId(v string) *DescribeQueryExplainRequest
	GetNodeId() *string
	SetSchema(v string) *DescribeQueryExplainRequest
	GetSchema() *string
	SetSql(v string) *DescribeQueryExplainRequest
	GetSql() *string
}

type DescribeQueryExplainRequest struct {
	// example:
	//
	// dbtest01
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// pi-bp1v203xzzh0a****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// select 	- from test where name = \\"mockUser\\"
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
}

func (s DescribeQueryExplainRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeQueryExplainRequest) GoString() string {
	return s.String()
}

func (s *DescribeQueryExplainRequest) GetDbName() *string {
	return s.DbName
}

func (s *DescribeQueryExplainRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeQueryExplainRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeQueryExplainRequest) GetSchema() *string {
	return s.Schema
}

func (s *DescribeQueryExplainRequest) GetSql() *string {
	return s.Sql
}

func (s *DescribeQueryExplainRequest) SetDbName(v string) *DescribeQueryExplainRequest {
	s.DbName = &v
	return s
}

func (s *DescribeQueryExplainRequest) SetInstanceId(v string) *DescribeQueryExplainRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeQueryExplainRequest) SetNodeId(v string) *DescribeQueryExplainRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeQueryExplainRequest) SetSchema(v string) *DescribeQueryExplainRequest {
	s.Schema = &v
	return s
}

func (s *DescribeQueryExplainRequest) SetSql(v string) *DescribeQueryExplainRequest {
	s.Sql = &v
	return s
}

func (s *DescribeQueryExplainRequest) Validate() error {
	return dara.Validate(s)
}
