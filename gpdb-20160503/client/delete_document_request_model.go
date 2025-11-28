// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *DeleteDocumentRequest
	GetCollection() *string
	SetDBInstanceId(v string) *DeleteDocumentRequest
	GetDBInstanceId() *string
	SetFileName(v string) *DeleteDocumentRequest
	GetFileName() *string
	SetNamespace(v string) *DeleteDocumentRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *DeleteDocumentRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *DeleteDocumentRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteDocumentRequest
	GetRegionId() *string
}

type DeleteDocumentRequest struct {
	// Document collection name.
	//
	//
	// > Created by the [CreateDocumentCollection](https://help.aliyun.com/document_detail/2618448.html) API. You can use the [ListDocumentCollections](https://help.aliyun.com/document_detail/2618452.html) API to view the list of created document collections.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// Instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) API to view details of all AnalyticDB PostgreSQL instances in the target region, including the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// File name.
	//
	// > The name of an uploaded file. You can query the list of files using the [ListDocuments](https://help.aliyun.com/document_detail/2618453.html) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// music.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The name of the namespace. Default value: public.
	//
	// >  You can call the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation to create a namespace and call the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to query a list of namespaces.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Password for the namespace.
	//
	// > This value is specified in the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpassword
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Region ID where the instance is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDocumentRequest) GoString() string {
	return s.String()
}

func (s *DeleteDocumentRequest) GetCollection() *string {
	return s.Collection
}

func (s *DeleteDocumentRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteDocumentRequest) GetFileName() *string {
	return s.FileName
}

func (s *DeleteDocumentRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteDocumentRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *DeleteDocumentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDocumentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDocumentRequest) SetCollection(v string) *DeleteDocumentRequest {
	s.Collection = &v
	return s
}

func (s *DeleteDocumentRequest) SetDBInstanceId(v string) *DeleteDocumentRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDocumentRequest) SetFileName(v string) *DeleteDocumentRequest {
	s.FileName = &v
	return s
}

func (s *DeleteDocumentRequest) SetNamespace(v string) *DeleteDocumentRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteDocumentRequest) SetNamespacePassword(v string) *DeleteDocumentRequest {
	s.NamespacePassword = &v
	return s
}

func (s *DeleteDocumentRequest) SetOwnerId(v int64) *DeleteDocumentRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDocumentRequest) SetRegionId(v string) *DeleteDocumentRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDocumentRequest) Validate() error {
	return dara.Validate(s)
}
