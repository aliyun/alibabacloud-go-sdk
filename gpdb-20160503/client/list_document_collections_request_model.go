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
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
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
	// >  This value is specified when you call the CreateNamespace operation.
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
