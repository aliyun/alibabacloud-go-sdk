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
	// The collection name.
	//
	// > You can call the [ListCollections](https://help.aliyun.com/document_detail/2401503.html) operation to view the list.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in the target region, including the instance IDs.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the task for uploading vector data. The ID is returned after you call the `UpsertCollectionDataAsync` operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 231460f8-75dc-405e-a669-0c5204887e91
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The namespace. Default value: public.
	//
	// > You can call the CreateNamespace operation to create a namespace and call the ListNamespaces operation to view the list.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password for the namespace.
	//
	// > This value is specified when you call the CreateNamespace operation.
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
	// The ID of the workspace that consists of multiple database instances. This parameter and the DBInstanceId parameter cannot be left empty at the same time. If both parameters are specified, this parameter takes precedence.
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
