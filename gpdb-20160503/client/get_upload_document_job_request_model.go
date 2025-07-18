// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadDocumentJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *GetUploadDocumentJobRequest
	GetCollection() *string
	SetDBInstanceId(v string) *GetUploadDocumentJobRequest
	GetDBInstanceId() *string
	SetJobId(v string) *GetUploadDocumentJobRequest
	GetJobId() *string
	SetNamespace(v string) *GetUploadDocumentJobRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *GetUploadDocumentJobRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *GetUploadDocumentJobRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetUploadDocumentJobRequest
	GetRegionId() *string
}

type GetUploadDocumentJobRequest struct {
	// The name of the document collection.
	//
	// >  You can call the [CreateDocumentCollection](https://help.aliyun.com/document_detail/2618448.html) operation to create a document collection and call the [ListDocumentCollections](https://help.aliyun.com/document_detail/2618452.html) operation to query a list of document collections.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The ID of the instance for which vector engine optimization is enabled.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the document upload job. You can call the `UploadDocumentAsync` operation to query the job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bf8f7bc4-9276-44f7-9c22-1d06edc8dfd1
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The name of the namespace. Default value: public.
	//
	// >  You can call the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation to create a namespace and call the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to query a list of namespaces.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password of the namespace.
	//
	// >  The value of this parameter is specified when you call the CreateNamespace operation.
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
}

func (s GetUploadDocumentJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUploadDocumentJobRequest) GoString() string {
	return s.String()
}

func (s *GetUploadDocumentJobRequest) GetCollection() *string {
	return s.Collection
}

func (s *GetUploadDocumentJobRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *GetUploadDocumentJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetUploadDocumentJobRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetUploadDocumentJobRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *GetUploadDocumentJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetUploadDocumentJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetUploadDocumentJobRequest) SetCollection(v string) *GetUploadDocumentJobRequest {
	s.Collection = &v
	return s
}

func (s *GetUploadDocumentJobRequest) SetDBInstanceId(v string) *GetUploadDocumentJobRequest {
	s.DBInstanceId = &v
	return s
}

func (s *GetUploadDocumentJobRequest) SetJobId(v string) *GetUploadDocumentJobRequest {
	s.JobId = &v
	return s
}

func (s *GetUploadDocumentJobRequest) SetNamespace(v string) *GetUploadDocumentJobRequest {
	s.Namespace = &v
	return s
}

func (s *GetUploadDocumentJobRequest) SetNamespacePassword(v string) *GetUploadDocumentJobRequest {
	s.NamespacePassword = &v
	return s
}

func (s *GetUploadDocumentJobRequest) SetOwnerId(v int64) *GetUploadDocumentJobRequest {
	s.OwnerId = &v
	return s
}

func (s *GetUploadDocumentJobRequest) SetRegionId(v string) *GetUploadDocumentJobRequest {
	s.RegionId = &v
	return s
}

func (s *GetUploadDocumentJobRequest) Validate() error {
	return dara.Validate(s)
}
