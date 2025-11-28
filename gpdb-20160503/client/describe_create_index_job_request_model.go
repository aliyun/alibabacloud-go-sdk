// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCreateIndexJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *DescribeCreateIndexJobRequest
	GetCollection() *string
	SetDBInstanceId(v string) *DescribeCreateIndexJobRequest
	GetDBInstanceId() *string
	SetJobId(v string) *DescribeCreateIndexJobRequest
	GetJobId() *string
	SetNamespace(v string) *DescribeCreateIndexJobRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *DescribeCreateIndexJobRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *DescribeCreateIndexJobRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeCreateIndexJobRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *DescribeCreateIndexJobRequest
	GetWorkspaceId() *string
}

type DescribeCreateIndexJobRequest struct {
	// The name of the collection.
	//
	// > You can call the [ListCollections](https://help.aliyun.com/document_detail/2401503.html) operation to query a list of collections.
	//
	// This parameter is required.
	//
	// example:
	//
	// testcollection
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The cluster ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the index creation job, which is returned from the `CreateIndex` operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 231460f8-75dc-405e-a669-0c5204887e91
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The namespace name.
	//
	// > You can call the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to query a list of namespaces.
	//
	// This parameter is required.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password for the namespace.
	//
	// > The value of this parameter is specified by the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpassword
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the available region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the workspace that consists of multiple AnalyticDB for PostgreSQL instances. You must specify one of the WorkspaceId and DBInstanceId parameters. If you specify both parameters, the WorkspaceId parameter takes precedence.
	//
	// Valid values:
	//
	// 	- ip
	//
	// 	- l2
	//
	// 	- cosine
	//
	// example:
	//
	// gp-ws-*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DescribeCreateIndexJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreateIndexJobRequest) GoString() string {
	return s.String()
}

func (s *DescribeCreateIndexJobRequest) GetCollection() *string {
	return s.Collection
}

func (s *DescribeCreateIndexJobRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeCreateIndexJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *DescribeCreateIndexJobRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeCreateIndexJobRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *DescribeCreateIndexJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCreateIndexJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCreateIndexJobRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DescribeCreateIndexJobRequest) SetCollection(v string) *DescribeCreateIndexJobRequest {
	s.Collection = &v
	return s
}

func (s *DescribeCreateIndexJobRequest) SetDBInstanceId(v string) *DescribeCreateIndexJobRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeCreateIndexJobRequest) SetJobId(v string) *DescribeCreateIndexJobRequest {
	s.JobId = &v
	return s
}

func (s *DescribeCreateIndexJobRequest) SetNamespace(v string) *DescribeCreateIndexJobRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeCreateIndexJobRequest) SetNamespacePassword(v string) *DescribeCreateIndexJobRequest {
	s.NamespacePassword = &v
	return s
}

func (s *DescribeCreateIndexJobRequest) SetOwnerId(v int64) *DescribeCreateIndexJobRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCreateIndexJobRequest) SetRegionId(v string) *DescribeCreateIndexJobRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCreateIndexJobRequest) SetWorkspaceId(v string) *DescribeCreateIndexJobRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeCreateIndexJobRequest) Validate() error {
	return dara.Validate(s)
}
