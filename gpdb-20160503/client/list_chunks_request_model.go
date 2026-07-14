// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChunksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *ListChunksRequest
	GetCollection() *string
	SetDBInstanceId(v string) *ListChunksRequest
	GetDBInstanceId() *string
	SetFileName(v string) *ListChunksRequest
	GetFileName() *string
	SetFilter(v string) *ListChunksRequest
	GetFilter() *string
	SetIncludeVector(v bool) *ListChunksRequest
	GetIncludeVector() *bool
	SetNamespace(v string) *ListChunksRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *ListChunksRequest
	GetNamespacePassword() *string
	SetPageNumber(v int64) *ListChunksRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListChunksRequest
	GetPageSize() *int64
	SetRegionId(v string) *ListChunksRequest
	GetRegionId() *string
}

type ListChunksRequest struct {
	// The name of the document collection.
	//
	// > The document collection is created by calling the [CreateDocumentCollection](https://help.aliyun.com/document_detail/2618448.html) operation. You can call the [ListDocumentCollections](https://help.aliyun.com/document_detail/2618452.html) operation to query existing document collections.
	//
	// This parameter is required.
	//
	// example:
	//
	// testcollection
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The file name.
	//
	// > The name of a previously uploaded file. You can call the [ListDocuments](https://help.aliyun.com/document_detail/2618453.html) operation to query the file list.
	//
	// example:
	//
	// mydoc.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The filter condition for the data to query, in SQL WHERE clause format. The filter is an expression that returns a Boolean value (true or false). Conditions can be simple comparison operators such as equal to (=), not equal to (<> or !=), greater than (>), less than (<), greater than or equal to (>=), and less than or equal to (<=). You can also use logical operators (AND, OR, NOT) to combine more complex expressions, as well as conditions with the IN, BETWEEN, and LIKE keywords.
	//
	// >
	//
	// > - For detailed syntax, refer to: https://www.postgresqltutorial.com/postgresql-tutorial/postgresql-where/
	//
	// example:
	//
	// title = \\"test\\" AND name like \\"test%\\"
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// Specifies whether to return vectors. Default value: false.
	//
	// > - **false**: Does not return vectors.
	//
	// > - **true**: Returns vectors.
	//
	// example:
	//
	// true
	IncludeVector *bool `json:"IncludeVector,omitempty" xml:"IncludeVector,omitempty"`
	// The namespace. Default value: public.
	//
	// > You can create a namespace by calling the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation and query the list of namespaces by calling the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password of the namespace.
	//
	// > This value is specified by the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpassword
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page. Valid values:
	//
	// - **20**
	//
	// - **50**
	//
	// - **100**
	//
	// Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListChunksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChunksRequest) GoString() string {
	return s.String()
}

func (s *ListChunksRequest) GetCollection() *string {
	return s.Collection
}

func (s *ListChunksRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListChunksRequest) GetFileName() *string {
	return s.FileName
}

func (s *ListChunksRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListChunksRequest) GetIncludeVector() *bool {
	return s.IncludeVector
}

func (s *ListChunksRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListChunksRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *ListChunksRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListChunksRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListChunksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListChunksRequest) SetCollection(v string) *ListChunksRequest {
	s.Collection = &v
	return s
}

func (s *ListChunksRequest) SetDBInstanceId(v string) *ListChunksRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListChunksRequest) SetFileName(v string) *ListChunksRequest {
	s.FileName = &v
	return s
}

func (s *ListChunksRequest) SetFilter(v string) *ListChunksRequest {
	s.Filter = &v
	return s
}

func (s *ListChunksRequest) SetIncludeVector(v bool) *ListChunksRequest {
	s.IncludeVector = &v
	return s
}

func (s *ListChunksRequest) SetNamespace(v string) *ListChunksRequest {
	s.Namespace = &v
	return s
}

func (s *ListChunksRequest) SetNamespacePassword(v string) *ListChunksRequest {
	s.NamespacePassword = &v
	return s
}

func (s *ListChunksRequest) SetPageNumber(v int64) *ListChunksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListChunksRequest) SetPageSize(v int64) *ListChunksRequest {
	s.PageSize = &v
	return s
}

func (s *ListChunksRequest) SetRegionId(v string) *ListChunksRequest {
	s.RegionId = &v
	return s
}

func (s *ListChunksRequest) Validate() error {
	return dara.Validate(s)
}
