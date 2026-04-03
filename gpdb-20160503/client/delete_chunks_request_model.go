// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChunksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChunkIds(v []*string) *DeleteChunksRequest
	GetChunkIds() []*string
	SetCollection(v string) *DeleteChunksRequest
	GetCollection() *string
	SetDBInstanceId(v string) *DeleteChunksRequest
	GetDBInstanceId() *string
	SetNamespace(v string) *DeleteChunksRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *DeleteChunksRequest
	GetNamespacePassword() *string
	SetRegionId(v string) *DeleteChunksRequest
	GetRegionId() *string
}

type DeleteChunksRequest struct {
	// This parameter is required.
	ChunkIds []*string `json:"ChunkIds,omitempty" xml:"ChunkIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// testcollection
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteChunksRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteChunksRequest) GoString() string {
	return s.String()
}

func (s *DeleteChunksRequest) GetChunkIds() []*string {
	return s.ChunkIds
}

func (s *DeleteChunksRequest) GetCollection() *string {
	return s.Collection
}

func (s *DeleteChunksRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteChunksRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteChunksRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *DeleteChunksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteChunksRequest) SetChunkIds(v []*string) *DeleteChunksRequest {
	s.ChunkIds = v
	return s
}

func (s *DeleteChunksRequest) SetCollection(v string) *DeleteChunksRequest {
	s.Collection = &v
	return s
}

func (s *DeleteChunksRequest) SetDBInstanceId(v string) *DeleteChunksRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteChunksRequest) SetNamespace(v string) *DeleteChunksRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteChunksRequest) SetNamespacePassword(v string) *DeleteChunksRequest {
	s.NamespacePassword = &v
	return s
}

func (s *DeleteChunksRequest) SetRegionId(v string) *DeleteChunksRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteChunksRequest) Validate() error {
	return dara.Validate(s)
}
