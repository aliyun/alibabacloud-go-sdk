// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCollectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *DeleteCollectionRequest
	GetCollection() *string
	SetDBInstanceId(v string) *DeleteCollectionRequest
	GetDBInstanceId() *string
	SetNamespace(v string) *DeleteCollectionRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *DeleteCollectionRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *DeleteCollectionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteCollectionRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *DeleteCollectionRequest
	GetWorkspaceId() *string
}

type DeleteCollectionRequest struct {
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
	// The name of the namespace.
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

func (s DeleteCollectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCollectionRequest) GoString() string {
	return s.String()
}

func (s *DeleteCollectionRequest) GetCollection() *string {
	return s.Collection
}

func (s *DeleteCollectionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteCollectionRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteCollectionRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *DeleteCollectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteCollectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCollectionRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteCollectionRequest) SetCollection(v string) *DeleteCollectionRequest {
	s.Collection = &v
	return s
}

func (s *DeleteCollectionRequest) SetDBInstanceId(v string) *DeleteCollectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteCollectionRequest) SetNamespace(v string) *DeleteCollectionRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteCollectionRequest) SetNamespacePassword(v string) *DeleteCollectionRequest {
	s.NamespacePassword = &v
	return s
}

func (s *DeleteCollectionRequest) SetOwnerId(v int64) *DeleteCollectionRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCollectionRequest) SetRegionId(v string) *DeleteCollectionRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCollectionRequest) SetWorkspaceId(v string) *DeleteCollectionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteCollectionRequest) Validate() error {
	return dara.Validate(s)
}
