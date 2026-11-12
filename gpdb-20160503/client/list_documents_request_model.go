// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDocumentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *ListDocumentsRequest
	GetCollection() *string
	SetDBInstanceId(v string) *ListDocumentsRequest
	GetDBInstanceId() *string
	SetMaxResults(v int32) *ListDocumentsRequest
	GetMaxResults() *int32
	SetNamespace(v string) *ListDocumentsRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *ListDocumentsRequest
	GetNamespacePassword() *string
	SetNextToken(v string) *ListDocumentsRequest
	GetNextToken() *string
	SetOwnerId(v int64) *ListDocumentsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListDocumentsRequest
	GetRegionId() *string
}

type ListDocumentsRequest struct {
	// The name of the document collection.
	//
	// > The document collection is created by calling the [CreateDocumentCollection](https://help.aliyun.com/document_detail/2618448.html) operation. You can call the [ListDocumentCollections](https://help.aliyun.com/document_detail/2618452.html) operation to view the created document collections.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to view the details of all AnalyticDB for PostgreSQL instances in the target region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The maximum number of documents to return per page. Valid values: 1 to 100.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The namespace. Default value: public.
	//
	// > You can create a namespace by calling the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation and view the namespace list by calling the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password of the namespace.
	//
	// > This value is specified when you call the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpassword
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	// The paging token used to start the next paged query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListDocumentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDocumentsRequest) GoString() string {
	return s.String()
}

func (s *ListDocumentsRequest) GetCollection() *string {
	return s.Collection
}

func (s *ListDocumentsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListDocumentsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDocumentsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListDocumentsRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *ListDocumentsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDocumentsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListDocumentsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDocumentsRequest) SetCollection(v string) *ListDocumentsRequest {
	s.Collection = &v
	return s
}

func (s *ListDocumentsRequest) SetDBInstanceId(v string) *ListDocumentsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListDocumentsRequest) SetMaxResults(v int32) *ListDocumentsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDocumentsRequest) SetNamespace(v string) *ListDocumentsRequest {
	s.Namespace = &v
	return s
}

func (s *ListDocumentsRequest) SetNamespacePassword(v string) *ListDocumentsRequest {
	s.NamespacePassword = &v
	return s
}

func (s *ListDocumentsRequest) SetNextToken(v string) *ListDocumentsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDocumentsRequest) SetOwnerId(v int64) *ListDocumentsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListDocumentsRequest) SetRegionId(v string) *ListDocumentsRequest {
	s.RegionId = &v
	return s
}

func (s *ListDocumentsRequest) Validate() error {
	return dara.Validate(s)
}
