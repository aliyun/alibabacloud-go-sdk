// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCollectionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ListCollectionsRequest
	GetDBInstanceId() *string
	SetNamespace(v string) *ListCollectionsRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *ListCollectionsRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *ListCollectionsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListCollectionsRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *ListCollectionsRequest
	GetWorkspaceId() *string
}

type ListCollectionsRequest struct {
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the namespace.
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

func (s ListCollectionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCollectionsRequest) GoString() string {
	return s.String()
}

func (s *ListCollectionsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListCollectionsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListCollectionsRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *ListCollectionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListCollectionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCollectionsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListCollectionsRequest) SetDBInstanceId(v string) *ListCollectionsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListCollectionsRequest) SetNamespace(v string) *ListCollectionsRequest {
	s.Namespace = &v
	return s
}

func (s *ListCollectionsRequest) SetNamespacePassword(v string) *ListCollectionsRequest {
	s.NamespacePassword = &v
	return s
}

func (s *ListCollectionsRequest) SetOwnerId(v int64) *ListCollectionsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListCollectionsRequest) SetRegionId(v string) *ListCollectionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListCollectionsRequest) SetWorkspaceId(v string) *ListCollectionsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListCollectionsRequest) Validate() error {
	return dara.Validate(s)
}
