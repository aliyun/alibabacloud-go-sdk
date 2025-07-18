// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUpsertCollectionDataJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *GetUpsertCollectionDataJobRequest
	GetCollection() *string
	SetDBInstanceId(v string) *GetUpsertCollectionDataJobRequest
	GetDBInstanceId() *string
	SetJobId(v string) *GetUpsertCollectionDataJobRequest
	GetJobId() *string
	SetNamespace(v string) *GetUpsertCollectionDataJobRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *GetUpsertCollectionDataJobRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *GetUpsertCollectionDataJobRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetUpsertCollectionDataJobRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *GetUpsertCollectionDataJobRequest
	GetWorkspaceId() *string
}

type GetUpsertCollectionDataJobRequest struct {
	// The name of the collection.
	//
	// >  You can call the [ListCollections](https://help.aliyun.com/document_detail/2401503.html) operation to query a list of collections.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the vector data upload job. You can call the `UpsertCollectionDataAsync` operation to query the job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 231460f8-75dc-405e-a669-0c5204887e91
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The name of the namespace. Default value: public.
	//
	// >  You can call the CreateNamespace operation to create a namespace and call the ListNamespaces operation to query a list of namespaces.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password of the namespace.
	//
	// >  The value of this parameter is specified when you call the CreateNamespace operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpassword
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the workspace that consists of multiple AnalyticDB for PostgreSQL instances. You must specify one of the WorkspaceId and DBInstanceId parameters. If you specify both parameters, the WorkspaceId parameter takes effect.
	//
	// example:
	//
	// gp-ws-*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetUpsertCollectionDataJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUpsertCollectionDataJobRequest) GoString() string {
	return s.String()
}

func (s *GetUpsertCollectionDataJobRequest) GetCollection() *string {
	return s.Collection
}

func (s *GetUpsertCollectionDataJobRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *GetUpsertCollectionDataJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetUpsertCollectionDataJobRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetUpsertCollectionDataJobRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *GetUpsertCollectionDataJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetUpsertCollectionDataJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetUpsertCollectionDataJobRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetUpsertCollectionDataJobRequest) SetCollection(v string) *GetUpsertCollectionDataJobRequest {
	s.Collection = &v
	return s
}

func (s *GetUpsertCollectionDataJobRequest) SetDBInstanceId(v string) *GetUpsertCollectionDataJobRequest {
	s.DBInstanceId = &v
	return s
}

func (s *GetUpsertCollectionDataJobRequest) SetJobId(v string) *GetUpsertCollectionDataJobRequest {
	s.JobId = &v
	return s
}

func (s *GetUpsertCollectionDataJobRequest) SetNamespace(v string) *GetUpsertCollectionDataJobRequest {
	s.Namespace = &v
	return s
}

func (s *GetUpsertCollectionDataJobRequest) SetNamespacePassword(v string) *GetUpsertCollectionDataJobRequest {
	s.NamespacePassword = &v
	return s
}

func (s *GetUpsertCollectionDataJobRequest) SetOwnerId(v int64) *GetUpsertCollectionDataJobRequest {
	s.OwnerId = &v
	return s
}

func (s *GetUpsertCollectionDataJobRequest) SetRegionId(v string) *GetUpsertCollectionDataJobRequest {
	s.RegionId = &v
	return s
}

func (s *GetUpsertCollectionDataJobRequest) SetWorkspaceId(v string) *GetUpsertCollectionDataJobRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetUpsertCollectionDataJobRequest) Validate() error {
	return dara.Validate(s)
}
