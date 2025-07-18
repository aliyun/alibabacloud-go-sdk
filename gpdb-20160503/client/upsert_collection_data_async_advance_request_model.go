// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iUpsertCollectionDataAsyncAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *UpsertCollectionDataAsyncAdvanceRequest
	GetCollection() *string
	SetDBInstanceId(v string) *UpsertCollectionDataAsyncAdvanceRequest
	GetDBInstanceId() *string
	SetFileUrlObject(v io.Reader) *UpsertCollectionDataAsyncAdvanceRequest
	GetFileUrlObject() io.Reader
	SetNamespace(v string) *UpsertCollectionDataAsyncAdvanceRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *UpsertCollectionDataAsyncAdvanceRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *UpsertCollectionDataAsyncAdvanceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpsertCollectionDataAsyncAdvanceRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *UpsertCollectionDataAsyncAdvanceRequest
	GetWorkspaceId() *string
}

type UpsertCollectionDataAsyncAdvanceRequest struct {
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
	FileUrlObject io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
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

func (s UpsertCollectionDataAsyncAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpsertCollectionDataAsyncAdvanceRequest) GoString() string {
	return s.String()
}

func (s *UpsertCollectionDataAsyncAdvanceRequest) GetCollection() *string {
	return s.Collection
}

func (s *UpsertCollectionDataAsyncAdvanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpsertCollectionDataAsyncAdvanceRequest) GetFileUrlObject() io.Reader {
	return s.FileUrlObject
}

func (s *UpsertCollectionDataAsyncAdvanceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpsertCollectionDataAsyncAdvanceRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *UpsertCollectionDataAsyncAdvanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpsertCollectionDataAsyncAdvanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpsertCollectionDataAsyncAdvanceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpsertCollectionDataAsyncAdvanceRequest) SetCollection(v string) *UpsertCollectionDataAsyncAdvanceRequest {
	s.Collection = &v
	return s
}

func (s *UpsertCollectionDataAsyncAdvanceRequest) SetDBInstanceId(v string) *UpsertCollectionDataAsyncAdvanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpsertCollectionDataAsyncAdvanceRequest) SetFileUrlObject(v io.Reader) *UpsertCollectionDataAsyncAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *UpsertCollectionDataAsyncAdvanceRequest) SetNamespace(v string) *UpsertCollectionDataAsyncAdvanceRequest {
	s.Namespace = &v
	return s
}

func (s *UpsertCollectionDataAsyncAdvanceRequest) SetNamespacePassword(v string) *UpsertCollectionDataAsyncAdvanceRequest {
	s.NamespacePassword = &v
	return s
}

func (s *UpsertCollectionDataAsyncAdvanceRequest) SetOwnerId(v int64) *UpsertCollectionDataAsyncAdvanceRequest {
	s.OwnerId = &v
	return s
}

func (s *UpsertCollectionDataAsyncAdvanceRequest) SetRegionId(v string) *UpsertCollectionDataAsyncAdvanceRequest {
	s.RegionId = &v
	return s
}

func (s *UpsertCollectionDataAsyncAdvanceRequest) SetWorkspaceId(v string) *UpsertCollectionDataAsyncAdvanceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpsertCollectionDataAsyncAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
