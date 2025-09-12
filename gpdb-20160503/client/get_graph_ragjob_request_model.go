// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGraphRAGJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *GetGraphRAGJobRequest
	GetCollection() *string
	SetDBInstanceId(v string) *GetGraphRAGJobRequest
	GetDBInstanceId() *string
	SetJobId(v string) *GetGraphRAGJobRequest
	GetJobId() *string
	SetNamespace(v string) *GetGraphRAGJobRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *GetGraphRAGJobRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *GetGraphRAGJobRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetGraphRAGJobRequest
	GetRegionId() *string
}

type GetGraphRAGJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// This parameter is required.
	//
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
}

func (s GetGraphRAGJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGraphRAGJobRequest) GoString() string {
	return s.String()
}

func (s *GetGraphRAGJobRequest) GetCollection() *string {
	return s.Collection
}

func (s *GetGraphRAGJobRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *GetGraphRAGJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetGraphRAGJobRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetGraphRAGJobRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *GetGraphRAGJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetGraphRAGJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetGraphRAGJobRequest) SetCollection(v string) *GetGraphRAGJobRequest {
	s.Collection = &v
	return s
}

func (s *GetGraphRAGJobRequest) SetDBInstanceId(v string) *GetGraphRAGJobRequest {
	s.DBInstanceId = &v
	return s
}

func (s *GetGraphRAGJobRequest) SetJobId(v string) *GetGraphRAGJobRequest {
	s.JobId = &v
	return s
}

func (s *GetGraphRAGJobRequest) SetNamespace(v string) *GetGraphRAGJobRequest {
	s.Namespace = &v
	return s
}

func (s *GetGraphRAGJobRequest) SetNamespacePassword(v string) *GetGraphRAGJobRequest {
	s.NamespacePassword = &v
	return s
}

func (s *GetGraphRAGJobRequest) SetOwnerId(v int64) *GetGraphRAGJobRequest {
	s.OwnerId = &v
	return s
}

func (s *GetGraphRAGJobRequest) SetRegionId(v string) *GetGraphRAGJobRequest {
	s.RegionId = &v
	return s
}

func (s *GetGraphRAGJobRequest) Validate() error {
	return dara.Validate(s)
}
