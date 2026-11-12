// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertCollectionDataShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *UpsertCollectionDataShrinkRequest
	GetCollection() *string
	SetDBInstanceId(v string) *UpsertCollectionDataShrinkRequest
	GetDBInstanceId() *string
	SetNamespace(v string) *UpsertCollectionDataShrinkRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *UpsertCollectionDataShrinkRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *UpsertCollectionDataShrinkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpsertCollectionDataShrinkRequest
	GetRegionId() *string
	SetRowsShrink(v string) *UpsertCollectionDataShrinkRequest
	GetRowsShrink() *string
	SetWorkspaceId(v string) *UpsertCollectionDataShrinkRequest
	GetWorkspaceId() *string
}

type UpsertCollectionDataShrinkRequest struct {
	// The name of the collection.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the namespace. Default value: public.
	//
	// >  You can call the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation to create a namespace and call the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to query a list of namespaces.
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
	// The vector data that you want to upload.
	RowsShrink *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// The ID of the workspace that consists of multiple AnalyticDB for PostgreSQL instances. You must specify one of the WorkspaceId and DBInstanceId parameters. If you specify both parameters, the WorkspaceId parameter takes effect.
	//
	// example:
	//
	// gp-ws-*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpsertCollectionDataShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpsertCollectionDataShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpsertCollectionDataShrinkRequest) GetCollection() *string {
	return s.Collection
}

func (s *UpsertCollectionDataShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpsertCollectionDataShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpsertCollectionDataShrinkRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *UpsertCollectionDataShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpsertCollectionDataShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpsertCollectionDataShrinkRequest) GetRowsShrink() *string {
	return s.RowsShrink
}

func (s *UpsertCollectionDataShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpsertCollectionDataShrinkRequest) SetCollection(v string) *UpsertCollectionDataShrinkRequest {
	s.Collection = &v
	return s
}

func (s *UpsertCollectionDataShrinkRequest) SetDBInstanceId(v string) *UpsertCollectionDataShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpsertCollectionDataShrinkRequest) SetNamespace(v string) *UpsertCollectionDataShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *UpsertCollectionDataShrinkRequest) SetNamespacePassword(v string) *UpsertCollectionDataShrinkRequest {
	s.NamespacePassword = &v
	return s
}

func (s *UpsertCollectionDataShrinkRequest) SetOwnerId(v int64) *UpsertCollectionDataShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpsertCollectionDataShrinkRequest) SetRegionId(v string) *UpsertCollectionDataShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpsertCollectionDataShrinkRequest) SetRowsShrink(v string) *UpsertCollectionDataShrinkRequest {
	s.RowsShrink = &v
	return s
}

func (s *UpsertCollectionDataShrinkRequest) SetWorkspaceId(v string) *UpsertCollectionDataShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpsertCollectionDataShrinkRequest) Validate() error {
	return dara.Validate(s)
}
