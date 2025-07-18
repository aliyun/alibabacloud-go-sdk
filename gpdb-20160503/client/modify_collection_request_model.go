// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCollectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *ModifyCollectionRequest
	GetCollection() *string
	SetDBInstanceId(v string) *ModifyCollectionRequest
	GetDBInstanceId() *string
	SetMetadata(v string) *ModifyCollectionRequest
	GetMetadata() *string
	SetNamespace(v string) *ModifyCollectionRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *ModifyCollectionRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *ModifyCollectionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyCollectionRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *ModifyCollectionRequest
	GetWorkspaceId() *string
}

type ModifyCollectionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"operations":[
	//
	// {"operator":"add","newMetaType":"int","newMetaName":"ext1"},
	//
	// {"operator":"replace","oldMetaName":"ext2","newMetaName":"ext3"},
	//
	// {"operator":"replace","newMetaType":"bigint","oldMetaName":"ext4"},
	//
	// {"operator":"replace","newMetaType":"int","oldMetaName":"ext5","newMetaName":"ext6"}
	//
	// ]}
	Metadata *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
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

func (s ModifyCollectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCollectionRequest) GoString() string {
	return s.String()
}

func (s *ModifyCollectionRequest) GetCollection() *string {
	return s.Collection
}

func (s *ModifyCollectionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyCollectionRequest) GetMetadata() *string {
	return s.Metadata
}

func (s *ModifyCollectionRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ModifyCollectionRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *ModifyCollectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyCollectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCollectionRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ModifyCollectionRequest) SetCollection(v string) *ModifyCollectionRequest {
	s.Collection = &v
	return s
}

func (s *ModifyCollectionRequest) SetDBInstanceId(v string) *ModifyCollectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyCollectionRequest) SetMetadata(v string) *ModifyCollectionRequest {
	s.Metadata = &v
	return s
}

func (s *ModifyCollectionRequest) SetNamespace(v string) *ModifyCollectionRequest {
	s.Namespace = &v
	return s
}

func (s *ModifyCollectionRequest) SetNamespacePassword(v string) *ModifyCollectionRequest {
	s.NamespacePassword = &v
	return s
}

func (s *ModifyCollectionRequest) SetOwnerId(v int64) *ModifyCollectionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCollectionRequest) SetRegionId(v string) *ModifyCollectionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCollectionRequest) SetWorkspaceId(v string) *ModifyCollectionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ModifyCollectionRequest) Validate() error {
	return dara.Validate(s)
}
