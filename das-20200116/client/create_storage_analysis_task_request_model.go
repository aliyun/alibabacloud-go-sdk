// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStorageAnalysisTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *CreateStorageAnalysisTaskRequest
	GetDbName() *string
	SetInstanceId(v string) *CreateStorageAnalysisTaskRequest
	GetInstanceId() *string
	SetNodeId(v string) *CreateStorageAnalysisTaskRequest
	GetNodeId() *string
	SetTableName(v string) *CreateStorageAnalysisTaskRequest
	GetTableName() *string
}

type CreateStorageAnalysisTaskRequest struct {
	// The database name. If you specify a database, the operation analyzes the storage usage of the specified database.
	//
	// example:
	//
	// testdb01
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze1jdv45i7l6****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The node ID. For ApsaraDB for MongoDB instances, you can use this parameter to specify a node for storage analysis. You can call the [DescribeRoleZoneInfo](https://help.aliyun.com/document_detail/123802.html) operation to query the information about nodes of an ApsaraDB for MongoDB instance.
	//
	// 	- If you set this parameter to a value in the **InsName*	- format, such as `d-bp1872fa24d5****`, you can call this operation to analyze the hidden node that corresponds to the node ID.
	//
	// 	- If you set this parameter to a value in the `InsName#RoleId` format, such as `d-bp1872fa24d5****#299****5`, you can call this operation to analyze the specified node.
	//
	// >  If you run a storage analysis task on an ApsaraDB for MongoDB replica set instance and you do not specify this parameter, only the hidden node of the instance is analyzed by default. If you run a storage analysis task on an ApsaraDB for MongoDB sharded cluster instance, we recommend that you set this parameter to specify a node.
	//
	// example:
	//
	// 23302528
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The table name. If you specify a table in the specified database, the operation analyzes the storage usage of the specified table. If you specify a table, you must also specify the database to which the table belongs by using **DbName**.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s CreateStorageAnalysisTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStorageAnalysisTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateStorageAnalysisTaskRequest) GetDbName() *string {
	return s.DbName
}

func (s *CreateStorageAnalysisTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateStorageAnalysisTaskRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateStorageAnalysisTaskRequest) GetTableName() *string {
	return s.TableName
}

func (s *CreateStorageAnalysisTaskRequest) SetDbName(v string) *CreateStorageAnalysisTaskRequest {
	s.DbName = &v
	return s
}

func (s *CreateStorageAnalysisTaskRequest) SetInstanceId(v string) *CreateStorageAnalysisTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateStorageAnalysisTaskRequest) SetNodeId(v string) *CreateStorageAnalysisTaskRequest {
	s.NodeId = &v
	return s
}

func (s *CreateStorageAnalysisTaskRequest) SetTableName(v string) *CreateStorageAnalysisTaskRequest {
	s.TableName = &v
	return s
}

func (s *CreateStorageAnalysisTaskRequest) Validate() error {
	return dara.Validate(s)
}
