// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIndicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *ListIndicesRequest
	GetCollection() *string
	SetDBInstanceId(v string) *ListIndicesRequest
	GetDBInstanceId() *string
	SetNamespace(v string) *ListIndicesRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *ListIndicesRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *ListIndicesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListIndicesRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *ListIndicesRequest
	GetWorkspaceId() *string
}

type ListIndicesRequest struct {
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

func (s ListIndicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIndicesRequest) GoString() string {
	return s.String()
}

func (s *ListIndicesRequest) GetCollection() *string {
	return s.Collection
}

func (s *ListIndicesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListIndicesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListIndicesRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *ListIndicesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListIndicesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListIndicesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListIndicesRequest) SetCollection(v string) *ListIndicesRequest {
	s.Collection = &v
	return s
}

func (s *ListIndicesRequest) SetDBInstanceId(v string) *ListIndicesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListIndicesRequest) SetNamespace(v string) *ListIndicesRequest {
	s.Namespace = &v
	return s
}

func (s *ListIndicesRequest) SetNamespacePassword(v string) *ListIndicesRequest {
	s.NamespacePassword = &v
	return s
}

func (s *ListIndicesRequest) SetOwnerId(v int64) *ListIndicesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListIndicesRequest) SetRegionId(v string) *ListIndicesRequest {
	s.RegionId = &v
	return s
}

func (s *ListIndicesRequest) SetWorkspaceId(v string) *ListIndicesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListIndicesRequest) Validate() error {
	return dara.Validate(s)
}
