// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertCollectionDataAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *UpsertCollectionDataAsyncRequest
	GetCollection() *string
	SetDBInstanceId(v string) *UpsertCollectionDataAsyncRequest
	GetDBInstanceId() *string
	SetFileUrl(v string) *UpsertCollectionDataAsyncRequest
	GetFileUrl() *string
	SetNamespace(v string) *UpsertCollectionDataAsyncRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *UpsertCollectionDataAsyncRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *UpsertCollectionDataAsyncRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpsertCollectionDataAsyncRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *UpsertCollectionDataAsyncRequest
	GetWorkspaceId() *string
}

type UpsertCollectionDataAsyncRequest struct {
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
	// The Internet-accessible vector data file URL.
	//
	// >
	//
	// 	- The file content must be in the JSONL format. Each line consists of a list of JSON data, which specifies a set of vector data.
	//
	// 	- Data format of each line: `{String Id; Map<String, Object> Metadata; List<Double> Vector}`. Example: `{"Id":"myid", "Metadata": {"my_meta_key": "my_meta_value"}, "Vector": [1.234, -0.123]}`.
	//
	// 	- We recommend that you use SDKs to call this operation. SDKs encapsulate the UpsertCollectionDataAsyncAdvance method to upload on-premises files as data sources.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://xx/vectors.jsonl
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
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
	// The ID of the Workspace composed of multiple database instances. This parameter and the DBInstanceId parameter cannot both be empty. When both are specified, this parameter takes precedence.
	//
	// example:
	//
	// gp-ws-*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpsertCollectionDataAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s UpsertCollectionDataAsyncRequest) GoString() string {
	return s.String()
}

func (s *UpsertCollectionDataAsyncRequest) GetCollection() *string {
	return s.Collection
}

func (s *UpsertCollectionDataAsyncRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpsertCollectionDataAsyncRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *UpsertCollectionDataAsyncRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpsertCollectionDataAsyncRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *UpsertCollectionDataAsyncRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpsertCollectionDataAsyncRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpsertCollectionDataAsyncRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpsertCollectionDataAsyncRequest) SetCollection(v string) *UpsertCollectionDataAsyncRequest {
	s.Collection = &v
	return s
}

func (s *UpsertCollectionDataAsyncRequest) SetDBInstanceId(v string) *UpsertCollectionDataAsyncRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpsertCollectionDataAsyncRequest) SetFileUrl(v string) *UpsertCollectionDataAsyncRequest {
	s.FileUrl = &v
	return s
}

func (s *UpsertCollectionDataAsyncRequest) SetNamespace(v string) *UpsertCollectionDataAsyncRequest {
	s.Namespace = &v
	return s
}

func (s *UpsertCollectionDataAsyncRequest) SetNamespacePassword(v string) *UpsertCollectionDataAsyncRequest {
	s.NamespacePassword = &v
	return s
}

func (s *UpsertCollectionDataAsyncRequest) SetOwnerId(v int64) *UpsertCollectionDataAsyncRequest {
	s.OwnerId = &v
	return s
}

func (s *UpsertCollectionDataAsyncRequest) SetRegionId(v string) *UpsertCollectionDataAsyncRequest {
	s.RegionId = &v
	return s
}

func (s *UpsertCollectionDataAsyncRequest) SetWorkspaceId(v string) *UpsertCollectionDataAsyncRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpsertCollectionDataAsyncRequest) Validate() error {
	return dara.Validate(s)
}
