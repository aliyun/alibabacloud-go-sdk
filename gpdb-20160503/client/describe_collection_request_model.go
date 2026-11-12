// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCollectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *DescribeCollectionRequest
	GetCollection() *string
	SetDBInstanceId(v string) *DescribeCollectionRequest
	GetDBInstanceId() *string
	SetNamespace(v string) *DescribeCollectionRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *DescribeCollectionRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *DescribeCollectionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeCollectionRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *DescribeCollectionRequest
	GetWorkspaceId() *string
}

type DescribeCollectionRequest struct {
	// The collection name.
	//
	// > You can call the [ListCollections](https://help.aliyun.com/document_detail/2401503.html) operation to query the list.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including the instance IDs.
	//
	// example:
	//
	// gp-bp152460513z****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The namespace.
	//
	// > You can call the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to query the list.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password for the namespace.
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

func (s DescribeCollectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCollectionRequest) GoString() string {
	return s.String()
}

func (s *DescribeCollectionRequest) GetCollection() *string {
	return s.Collection
}

func (s *DescribeCollectionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeCollectionRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeCollectionRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *DescribeCollectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCollectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCollectionRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DescribeCollectionRequest) SetCollection(v string) *DescribeCollectionRequest {
	s.Collection = &v
	return s
}

func (s *DescribeCollectionRequest) SetDBInstanceId(v string) *DescribeCollectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeCollectionRequest) SetNamespace(v string) *DescribeCollectionRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeCollectionRequest) SetNamespacePassword(v string) *DescribeCollectionRequest {
	s.NamespacePassword = &v
	return s
}

func (s *DescribeCollectionRequest) SetOwnerId(v int64) *DescribeCollectionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCollectionRequest) SetRegionId(v string) *DescribeCollectionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCollectionRequest) SetWorkspaceId(v string) *DescribeCollectionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeCollectionRequest) Validate() error {
	return dara.Validate(s)
}
