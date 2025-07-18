// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *DescribeIndexRequest
	GetCollection() *string
	SetDBInstanceId(v string) *DescribeIndexRequest
	GetDBInstanceId() *string
	SetIndexName(v string) *DescribeIndexRequest
	GetIndexName() *string
	SetNamespace(v string) *DescribeIndexRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *DescribeIndexRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *DescribeIndexRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeIndexRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *DescribeIndexRequest
	GetWorkspaceId() *string
}

type DescribeIndexRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// testcollection
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testindex
	IndexName *string `json:"IndexName,omitempty" xml:"IndexName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testpassword
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// gp-ws-*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DescribeIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIndexRequest) GoString() string {
	return s.String()
}

func (s *DescribeIndexRequest) GetCollection() *string {
	return s.Collection
}

func (s *DescribeIndexRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeIndexRequest) GetIndexName() *string {
	return s.IndexName
}

func (s *DescribeIndexRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeIndexRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *DescribeIndexRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeIndexRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeIndexRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DescribeIndexRequest) SetCollection(v string) *DescribeIndexRequest {
	s.Collection = &v
	return s
}

func (s *DescribeIndexRequest) SetDBInstanceId(v string) *DescribeIndexRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeIndexRequest) SetIndexName(v string) *DescribeIndexRequest {
	s.IndexName = &v
	return s
}

func (s *DescribeIndexRequest) SetNamespace(v string) *DescribeIndexRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeIndexRequest) SetNamespacePassword(v string) *DescribeIndexRequest {
	s.NamespacePassword = &v
	return s
}

func (s *DescribeIndexRequest) SetOwnerId(v int64) *DescribeIndexRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeIndexRequest) SetRegionId(v string) *DescribeIndexRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeIndexRequest) SetWorkspaceId(v string) *DescribeIndexRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeIndexRequest) Validate() error {
	return dara.Validate(s)
}
