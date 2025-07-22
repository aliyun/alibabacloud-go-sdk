// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCacheAnalysisJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupSetId(v string) *CreateCacheAnalysisJobRequest
	GetBackupSetId() *string
	SetInstanceId(v string) *CreateCacheAnalysisJobRequest
	GetInstanceId() *string
	SetNodeId(v string) *CreateCacheAnalysisJobRequest
	GetNodeId() *string
	SetSeparators(v string) *CreateCacheAnalysisJobRequest
	GetSeparators() *string
}

type CreateCacheAnalysisJobRequest struct {
	// The ID of the backup file. You can call the [DescribeBackups](https://help.aliyun.com/document_detail/473823.html) operation to query the ID.
	//
	// 	- If you need to specify multiple backup file IDs, separate them with commas (,). For example, you can set this parameter to `12345,67890`.
	//
	// 	- If you do not specify this parameter, the system automatically backs up the task and performs cache analysis on the backup file.
	//
	// example:
	//
	// 12345
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The ID of the ApsaraDB for Redis instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp18ff4a195d****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the data node on the instance. You can specify this parameter to query the monitoring information about the specified node.
	//
	// >  If you specify the BackupSetId parameter, the system ignores the NodeId parameter. You can call the [DescribeLogicInstanceTopology](https://help.aliyun.com/document_detail/473786.html) operation to query the node ID.
	//
	// example:
	//
	// r-x****-db-0
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The delimiters used to identify the prefixes of keys. You do not need to specify this parameter if one or more of the following default delimiters are used: `: ; , _ - + @ = | #`
	//
	// example:
	//
	// &
	Separators *string `json:"Separators,omitempty" xml:"Separators,omitempty"`
}

func (s CreateCacheAnalysisJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCacheAnalysisJobRequest) GoString() string {
	return s.String()
}

func (s *CreateCacheAnalysisJobRequest) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *CreateCacheAnalysisJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCacheAnalysisJobRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateCacheAnalysisJobRequest) GetSeparators() *string {
	return s.Separators
}

func (s *CreateCacheAnalysisJobRequest) SetBackupSetId(v string) *CreateCacheAnalysisJobRequest {
	s.BackupSetId = &v
	return s
}

func (s *CreateCacheAnalysisJobRequest) SetInstanceId(v string) *CreateCacheAnalysisJobRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCacheAnalysisJobRequest) SetNodeId(v string) *CreateCacheAnalysisJobRequest {
	s.NodeId = &v
	return s
}

func (s *CreateCacheAnalysisJobRequest) SetSeparators(v string) *CreateCacheAnalysisJobRequest {
	s.Separators = &v
	return s
}

func (s *CreateCacheAnalysisJobRequest) Validate() error {
	return dara.Validate(s)
}
