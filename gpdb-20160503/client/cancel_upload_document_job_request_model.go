// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelUploadDocumentJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *CancelUploadDocumentJobRequest
	GetCollection() *string
	SetDBInstanceId(v string) *CancelUploadDocumentJobRequest
	GetDBInstanceId() *string
	SetJobId(v string) *CancelUploadDocumentJobRequest
	GetJobId() *string
	SetNamespace(v string) *CancelUploadDocumentJobRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *CancelUploadDocumentJobRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *CancelUploadDocumentJobRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CancelUploadDocumentJobRequest
	GetRegionId() *string
}

type CancelUploadDocumentJobRequest struct {
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
	// 231460f8-75dc-405e-a669-0c5204887e91
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The name of the namespace. Default value: public.
	//
	// >  You can call the CreateNamespace operation to create a namespace and call the ListNamespaces operation to query a list of namespaces.
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

func (s CancelUploadDocumentJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelUploadDocumentJobRequest) GoString() string {
	return s.String()
}

func (s *CancelUploadDocumentJobRequest) GetCollection() *string {
	return s.Collection
}

func (s *CancelUploadDocumentJobRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CancelUploadDocumentJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *CancelUploadDocumentJobRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CancelUploadDocumentJobRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *CancelUploadDocumentJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelUploadDocumentJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelUploadDocumentJobRequest) SetCollection(v string) *CancelUploadDocumentJobRequest {
	s.Collection = &v
	return s
}

func (s *CancelUploadDocumentJobRequest) SetDBInstanceId(v string) *CancelUploadDocumentJobRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CancelUploadDocumentJobRequest) SetJobId(v string) *CancelUploadDocumentJobRequest {
	s.JobId = &v
	return s
}

func (s *CancelUploadDocumentJobRequest) SetNamespace(v string) *CancelUploadDocumentJobRequest {
	s.Namespace = &v
	return s
}

func (s *CancelUploadDocumentJobRequest) SetNamespacePassword(v string) *CancelUploadDocumentJobRequest {
	s.NamespacePassword = &v
	return s
}

func (s *CancelUploadDocumentJobRequest) SetOwnerId(v int64) *CancelUploadDocumentJobRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelUploadDocumentJobRequest) SetRegionId(v string) *CancelUploadDocumentJobRequest {
	s.RegionId = &v
	return s
}

func (s *CancelUploadDocumentJobRequest) Validate() error {
	return dara.Validate(s)
}
