// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCreateIndexJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *CancelCreateIndexJobRequest
	GetCollection() *string
	SetDBInstanceId(v string) *CancelCreateIndexJobRequest
	GetDBInstanceId() *string
	SetJobId(v string) *CancelCreateIndexJobRequest
	GetJobId() *string
	SetNamespace(v string) *CancelCreateIndexJobRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *CancelCreateIndexJobRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *CancelCreateIndexJobRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CancelCreateIndexJobRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *CancelCreateIndexJobRequest
	GetWorkspaceId() *string
}

type CancelCreateIndexJobRequest struct {
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
	// 231460f8-75dc-405e-a669-0c5204887e91
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
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

func (s CancelCreateIndexJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelCreateIndexJobRequest) GoString() string {
	return s.String()
}

func (s *CancelCreateIndexJobRequest) GetCollection() *string {
	return s.Collection
}

func (s *CancelCreateIndexJobRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CancelCreateIndexJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *CancelCreateIndexJobRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CancelCreateIndexJobRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *CancelCreateIndexJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelCreateIndexJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelCreateIndexJobRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CancelCreateIndexJobRequest) SetCollection(v string) *CancelCreateIndexJobRequest {
	s.Collection = &v
	return s
}

func (s *CancelCreateIndexJobRequest) SetDBInstanceId(v string) *CancelCreateIndexJobRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CancelCreateIndexJobRequest) SetJobId(v string) *CancelCreateIndexJobRequest {
	s.JobId = &v
	return s
}

func (s *CancelCreateIndexJobRequest) SetNamespace(v string) *CancelCreateIndexJobRequest {
	s.Namespace = &v
	return s
}

func (s *CancelCreateIndexJobRequest) SetNamespacePassword(v string) *CancelCreateIndexJobRequest {
	s.NamespacePassword = &v
	return s
}

func (s *CancelCreateIndexJobRequest) SetOwnerId(v int64) *CancelCreateIndexJobRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelCreateIndexJobRequest) SetRegionId(v string) *CancelCreateIndexJobRequest {
	s.RegionId = &v
	return s
}

func (s *CancelCreateIndexJobRequest) SetWorkspaceId(v string) *CancelCreateIndexJobRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CancelCreateIndexJobRequest) Validate() error {
	return dara.Validate(s)
}
