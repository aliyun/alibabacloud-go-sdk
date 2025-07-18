// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelUpsertCollectionDataJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *CancelUpsertCollectionDataJobRequest
	GetCollection() *string
	SetDBInstanceId(v string) *CancelUpsertCollectionDataJobRequest
	GetDBInstanceId() *string
	SetJobId(v string) *CancelUpsertCollectionDataJobRequest
	GetJobId() *string
	SetNamespace(v string) *CancelUpsertCollectionDataJobRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *CancelUpsertCollectionDataJobRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *CancelUpsertCollectionDataJobRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CancelUpsertCollectionDataJobRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *CancelUpsertCollectionDataJobRequest
	GetWorkspaceId() *string
}

type CancelUpsertCollectionDataJobRequest struct {
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
	// bf8f7bc4-9276-44f7-9c22-1d06edc8dfd1
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

func (s CancelUpsertCollectionDataJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelUpsertCollectionDataJobRequest) GoString() string {
	return s.String()
}

func (s *CancelUpsertCollectionDataJobRequest) GetCollection() *string {
	return s.Collection
}

func (s *CancelUpsertCollectionDataJobRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CancelUpsertCollectionDataJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *CancelUpsertCollectionDataJobRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CancelUpsertCollectionDataJobRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *CancelUpsertCollectionDataJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelUpsertCollectionDataJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelUpsertCollectionDataJobRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CancelUpsertCollectionDataJobRequest) SetCollection(v string) *CancelUpsertCollectionDataJobRequest {
	s.Collection = &v
	return s
}

func (s *CancelUpsertCollectionDataJobRequest) SetDBInstanceId(v string) *CancelUpsertCollectionDataJobRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CancelUpsertCollectionDataJobRequest) SetJobId(v string) *CancelUpsertCollectionDataJobRequest {
	s.JobId = &v
	return s
}

func (s *CancelUpsertCollectionDataJobRequest) SetNamespace(v string) *CancelUpsertCollectionDataJobRequest {
	s.Namespace = &v
	return s
}

func (s *CancelUpsertCollectionDataJobRequest) SetNamespacePassword(v string) *CancelUpsertCollectionDataJobRequest {
	s.NamespacePassword = &v
	return s
}

func (s *CancelUpsertCollectionDataJobRequest) SetOwnerId(v int64) *CancelUpsertCollectionDataJobRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelUpsertCollectionDataJobRequest) SetRegionId(v string) *CancelUpsertCollectionDataJobRequest {
	s.RegionId = &v
	return s
}

func (s *CancelUpsertCollectionDataJobRequest) SetWorkspaceId(v string) *CancelUpsertCollectionDataJobRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CancelUpsertCollectionDataJobRequest) Validate() error {
	return dara.Validate(s)
}
