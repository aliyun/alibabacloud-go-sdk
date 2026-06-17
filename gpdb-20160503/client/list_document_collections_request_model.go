// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDocumentCollectionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ListDocumentCollectionsRequest
	GetDBInstanceId() *string
	SetNamespace(v string) *ListDocumentCollectionsRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *ListDocumentCollectionsRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *ListDocumentCollectionsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListDocumentCollectionsRequest
	GetRegionId() *string
}

type ListDocumentCollectionsRequest struct {
	// The ID of the instance.
	//
	// > To view details of all AnalyticDB for PostgreSQL instances in a region, including their IDs, call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The namespace. Default value: public.
	//
	// > To create a namespace, call the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation. To list namespaces, call the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password for the namespace.
	//
	// > You set this password when you call the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpassword
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListDocumentCollectionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDocumentCollectionsRequest) GoString() string {
	return s.String()
}

func (s *ListDocumentCollectionsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListDocumentCollectionsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListDocumentCollectionsRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *ListDocumentCollectionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListDocumentCollectionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDocumentCollectionsRequest) SetDBInstanceId(v string) *ListDocumentCollectionsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListDocumentCollectionsRequest) SetNamespace(v string) *ListDocumentCollectionsRequest {
	s.Namespace = &v
	return s
}

func (s *ListDocumentCollectionsRequest) SetNamespacePassword(v string) *ListDocumentCollectionsRequest {
	s.NamespacePassword = &v
	return s
}

func (s *ListDocumentCollectionsRequest) SetOwnerId(v int64) *ListDocumentCollectionsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListDocumentCollectionsRequest) SetRegionId(v string) *ListDocumentCollectionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListDocumentCollectionsRequest) Validate() error {
	return dara.Validate(s)
}
