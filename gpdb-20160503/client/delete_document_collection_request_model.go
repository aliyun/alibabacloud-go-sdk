// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDocumentCollectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *DeleteDocumentCollectionRequest
	GetCollection() *string
	SetDBInstanceId(v string) *DeleteDocumentCollectionRequest
	GetDBInstanceId() *string
	SetNamespace(v string) *DeleteDocumentCollectionRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *DeleteDocumentCollectionRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *DeleteDocumentCollectionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteDocumentCollectionRequest
	GetRegionId() *string
}

type DeleteDocumentCollectionRequest struct {
	// The name of the document collection to be deleted.
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
	// Namespace, default is public.
	//
	// > You can create a namespace using the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) API and view the list of namespaces using the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) API.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Password for the namespace.
	//
	// > This value is specified by the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpassword
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the instance is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDocumentCollectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDocumentCollectionRequest) GoString() string {
	return s.String()
}

func (s *DeleteDocumentCollectionRequest) GetCollection() *string {
	return s.Collection
}

func (s *DeleteDocumentCollectionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteDocumentCollectionRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteDocumentCollectionRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *DeleteDocumentCollectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDocumentCollectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDocumentCollectionRequest) SetCollection(v string) *DeleteDocumentCollectionRequest {
	s.Collection = &v
	return s
}

func (s *DeleteDocumentCollectionRequest) SetDBInstanceId(v string) *DeleteDocumentCollectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDocumentCollectionRequest) SetNamespace(v string) *DeleteDocumentCollectionRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteDocumentCollectionRequest) SetNamespacePassword(v string) *DeleteDocumentCollectionRequest {
	s.NamespacePassword = &v
	return s
}

func (s *DeleteDocumentCollectionRequest) SetOwnerId(v int64) *DeleteDocumentCollectionRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDocumentCollectionRequest) SetRegionId(v string) *DeleteDocumentCollectionRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDocumentCollectionRequest) Validate() error {
	return dara.Validate(s)
}
