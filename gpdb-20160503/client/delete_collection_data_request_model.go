// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCollectionDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *DeleteCollectionDataRequest
	GetCollection() *string
	SetCollectionData(v string) *DeleteCollectionDataRequest
	GetCollectionData() *string
	SetCollectionDataFilter(v string) *DeleteCollectionDataRequest
	GetCollectionDataFilter() *string
	SetDBInstanceId(v string) *DeleteCollectionDataRequest
	GetDBInstanceId() *string
	SetNamespace(v string) *DeleteCollectionDataRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *DeleteCollectionDataRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *DeleteCollectionDataRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteCollectionDataRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *DeleteCollectionDataRequest
	GetWorkspaceId() *string
}

type DeleteCollectionDataRequest struct {
	// The name of the collection.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The data that you want to delete.
	//
	// example:
	//
	// {"title":["a","b"]}
	CollectionData *string `json:"CollectionData,omitempty" xml:"CollectionData,omitempty"`
	// The filter conditions for the data to be deleted.
	//
	// example:
	//
	// a < 10
	CollectionDataFilter *string `json:"CollectionDataFilter,omitempty" xml:"CollectionDataFilter,omitempty"`
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the IDs of all AnalyticDB for PostgreSQL instances within a region.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the namespace. Default value: public.
	//
	// >  You can call the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to query a list of namespaces.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password of the namespace.
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
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the most recent region list.
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

func (s DeleteCollectionDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCollectionDataRequest) GoString() string {
	return s.String()
}

func (s *DeleteCollectionDataRequest) GetCollection() *string {
	return s.Collection
}

func (s *DeleteCollectionDataRequest) GetCollectionData() *string {
	return s.CollectionData
}

func (s *DeleteCollectionDataRequest) GetCollectionDataFilter() *string {
	return s.CollectionDataFilter
}

func (s *DeleteCollectionDataRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteCollectionDataRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteCollectionDataRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *DeleteCollectionDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteCollectionDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCollectionDataRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteCollectionDataRequest) SetCollection(v string) *DeleteCollectionDataRequest {
	s.Collection = &v
	return s
}

func (s *DeleteCollectionDataRequest) SetCollectionData(v string) *DeleteCollectionDataRequest {
	s.CollectionData = &v
	return s
}

func (s *DeleteCollectionDataRequest) SetCollectionDataFilter(v string) *DeleteCollectionDataRequest {
	s.CollectionDataFilter = &v
	return s
}

func (s *DeleteCollectionDataRequest) SetDBInstanceId(v string) *DeleteCollectionDataRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteCollectionDataRequest) SetNamespace(v string) *DeleteCollectionDataRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteCollectionDataRequest) SetNamespacePassword(v string) *DeleteCollectionDataRequest {
	s.NamespacePassword = &v
	return s
}

func (s *DeleteCollectionDataRequest) SetOwnerId(v int64) *DeleteCollectionDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCollectionDataRequest) SetRegionId(v string) *DeleteCollectionDataRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCollectionDataRequest) SetWorkspaceId(v string) *DeleteCollectionDataRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteCollectionDataRequest) Validate() error {
	return dara.Validate(s)
}
