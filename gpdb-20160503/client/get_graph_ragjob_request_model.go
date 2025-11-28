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
	// The name of the document collection.
	//
	// > You can call the [CreateDocumentCollection](https://help.aliyun.com/document_detail/2618448.html) operation to create a document collection and call the [ListDocumentCollections](https://help.aliyun.com/document_detail/2618452.html) operation to query a list of document collections.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The cluster ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 231460f8-75dc-405e-a669-0c5204887e91
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The name of the namespace. Default value: public.
	//
	// > You can call the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to query a list of namespaces.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password of the namespace.
	//
	// > The value of this parameter is specified when you call the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpassword
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the most recent region list.
	//
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
