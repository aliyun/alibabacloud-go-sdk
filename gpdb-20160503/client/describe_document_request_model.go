// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *DescribeDocumentRequest
	GetCollection() *string
	SetDBInstanceId(v string) *DescribeDocumentRequest
	GetDBInstanceId() *string
	SetFileName(v string) *DescribeDocumentRequest
	GetFileName() *string
	SetNamespace(v string) *DescribeDocumentRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *DescribeDocumentRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *DescribeDocumentRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDocumentRequest
	GetRegionId() *string
}

type DescribeDocumentRequest struct {
	// The name of the document collection.
	//
	// > You can call the [CreateDocumentCollection](https://help.aliyun.com/document_detail/2618448.html) operation to create a document collection. and call the [ListDocumentCollections](https://help.aliyun.com/document_detail/2618452.html) operation to query a list of document collections.
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
	// The name of the document.
	//
	// > You can call the [ListDocuments](https://help.aliyun.com/document_detail/2618453.html) operation to query a list of documents.
	//
	// This parameter is required.
	//
	// example:
	//
	// music.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The name of the namespace. Default value: public.
	//
	// > You can call the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation to create a namespace and call the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to query a list of namespaces.
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
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocumentRequest) GoString() string {
	return s.String()
}

func (s *DescribeDocumentRequest) GetCollection() *string {
	return s.Collection
}

func (s *DescribeDocumentRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDocumentRequest) GetFileName() *string {
	return s.FileName
}

func (s *DescribeDocumentRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeDocumentRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *DescribeDocumentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDocumentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDocumentRequest) SetCollection(v string) *DescribeDocumentRequest {
	s.Collection = &v
	return s
}

func (s *DescribeDocumentRequest) SetDBInstanceId(v string) *DescribeDocumentRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDocumentRequest) SetFileName(v string) *DescribeDocumentRequest {
	s.FileName = &v
	return s
}

func (s *DescribeDocumentRequest) SetNamespace(v string) *DescribeDocumentRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeDocumentRequest) SetNamespacePassword(v string) *DescribeDocumentRequest {
	s.NamespacePassword = &v
	return s
}

func (s *DescribeDocumentRequest) SetOwnerId(v int64) *DescribeDocumentRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDocumentRequest) SetRegionId(v string) *DescribeDocumentRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDocumentRequest) Validate() error {
	return dara.Validate(s)
}
