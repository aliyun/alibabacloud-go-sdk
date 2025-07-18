// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *CreateIndexRequest
	GetCollection() *string
	SetDBInstanceId(v string) *CreateIndexRequest
	GetDBInstanceId() *string
	SetIndexConfig(v string) *CreateIndexRequest
	GetIndexConfig() *string
	SetIndexField(v string) *CreateIndexRequest
	GetIndexField() *string
	SetIndexName(v string) *CreateIndexRequest
	GetIndexName() *string
	SetNamespace(v string) *CreateIndexRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *CreateIndexRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *CreateIndexRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateIndexRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *CreateIndexRequest
	GetWorkspaceId() *string
}

type CreateIndexRequest struct {
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
	IndexConfig  *string `json:"IndexConfig,omitempty" xml:"IndexConfig,omitempty"`
	// example:
	//
	// title
	IndexField *string `json:"IndexField,omitempty" xml:"IndexField,omitempty"`
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

func (s CreateIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIndexRequest) GoString() string {
	return s.String()
}

func (s *CreateIndexRequest) GetCollection() *string {
	return s.Collection
}

func (s *CreateIndexRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateIndexRequest) GetIndexConfig() *string {
	return s.IndexConfig
}

func (s *CreateIndexRequest) GetIndexField() *string {
	return s.IndexField
}

func (s *CreateIndexRequest) GetIndexName() *string {
	return s.IndexName
}

func (s *CreateIndexRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateIndexRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *CreateIndexRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateIndexRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateIndexRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateIndexRequest) SetCollection(v string) *CreateIndexRequest {
	s.Collection = &v
	return s
}

func (s *CreateIndexRequest) SetDBInstanceId(v string) *CreateIndexRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateIndexRequest) SetIndexConfig(v string) *CreateIndexRequest {
	s.IndexConfig = &v
	return s
}

func (s *CreateIndexRequest) SetIndexField(v string) *CreateIndexRequest {
	s.IndexField = &v
	return s
}

func (s *CreateIndexRequest) SetIndexName(v string) *CreateIndexRequest {
	s.IndexName = &v
	return s
}

func (s *CreateIndexRequest) SetNamespace(v string) *CreateIndexRequest {
	s.Namespace = &v
	return s
}

func (s *CreateIndexRequest) SetNamespacePassword(v string) *CreateIndexRequest {
	s.NamespacePassword = &v
	return s
}

func (s *CreateIndexRequest) SetOwnerId(v int64) *CreateIndexRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateIndexRequest) SetRegionId(v string) *CreateIndexRequest {
	s.RegionId = &v
	return s
}

func (s *CreateIndexRequest) SetWorkspaceId(v string) *CreateIndexRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateIndexRequest) Validate() error {
	return dara.Validate(s)
}
