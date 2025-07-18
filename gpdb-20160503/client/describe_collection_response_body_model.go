// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCollectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeCollectionResponseBody
	GetDBInstanceId() *string
	SetDimension(v int32) *DescribeCollectionResponseBody
	GetDimension() *int32
	SetFullTextRetrievalFields(v string) *DescribeCollectionResponseBody
	GetFullTextRetrievalFields() *string
	SetMessage(v string) *DescribeCollectionResponseBody
	GetMessage() *string
	SetMetadata(v map[string]*string) *DescribeCollectionResponseBody
	GetMetadata() map[string]*string
	SetMetrics(v string) *DescribeCollectionResponseBody
	GetMetrics() *string
	SetNamespace(v string) *DescribeCollectionResponseBody
	GetNamespace() *string
	SetParser(v string) *DescribeCollectionResponseBody
	GetParser() *string
	SetRegionId(v string) *DescribeCollectionResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeCollectionResponseBody
	GetRequestId() *string
	SetSparseVectorMetrics(v string) *DescribeCollectionResponseBody
	GetSparseVectorMetrics() *string
	SetStatus(v string) *DescribeCollectionResponseBody
	GetStatus() *string
	SetSupportSparse(v bool) *DescribeCollectionResponseBody
	GetSupportSparse() *bool
}

type DescribeCollectionResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The number of vector dimensions.
	//
	// example:
	//
	// 1024
	Dimension *int32 `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The fields that are used for full-text search. Multiple fields are separated by commas (,).
	//
	// example:
	//
	// title,content
	FullTextRetrievalFields *string `json:"FullTextRetrievalFields,omitempty" xml:"FullTextRetrievalFields,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The metadata of vector data, which is a JSON string in the MAP format. The key specifies the field name, and the value specifies the data type.
	//
	// **
	//
	// **Warning*	- Reserved fields such as id, vector, and to_tsvector cannot be used.
	Metadata map[string]*string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The distance metrics.
	//
	// example:
	//
	// cosine
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The analyzer that is used for full-text search.
	//
	// example:
	//
	// zh_cn
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SparseVectorMetrics *string `json:"SparseVectorMetrics,omitempty" xml:"SparseVectorMetrics,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **success**
	//
	// 	- **fail**
	//
	// example:
	//
	// success
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SupportSparse *bool   `json:"SupportSparse,omitempty" xml:"SupportSparse,omitempty"`
}

func (s DescribeCollectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCollectionResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeCollectionResponseBody) GetDimension() *int32 {
	return s.Dimension
}

func (s *DescribeCollectionResponseBody) GetFullTextRetrievalFields() *string {
	return s.FullTextRetrievalFields
}

func (s *DescribeCollectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCollectionResponseBody) GetMetadata() map[string]*string {
	return s.Metadata
}

func (s *DescribeCollectionResponseBody) GetMetrics() *string {
	return s.Metrics
}

func (s *DescribeCollectionResponseBody) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeCollectionResponseBody) GetParser() *string {
	return s.Parser
}

func (s *DescribeCollectionResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCollectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCollectionResponseBody) GetSparseVectorMetrics() *string {
	return s.SparseVectorMetrics
}

func (s *DescribeCollectionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeCollectionResponseBody) GetSupportSparse() *bool {
	return s.SupportSparse
}

func (s *DescribeCollectionResponseBody) SetDBInstanceId(v string) *DescribeCollectionResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeCollectionResponseBody) SetDimension(v int32) *DescribeCollectionResponseBody {
	s.Dimension = &v
	return s
}

func (s *DescribeCollectionResponseBody) SetFullTextRetrievalFields(v string) *DescribeCollectionResponseBody {
	s.FullTextRetrievalFields = &v
	return s
}

func (s *DescribeCollectionResponseBody) SetMessage(v string) *DescribeCollectionResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCollectionResponseBody) SetMetadata(v map[string]*string) *DescribeCollectionResponseBody {
	s.Metadata = v
	return s
}

func (s *DescribeCollectionResponseBody) SetMetrics(v string) *DescribeCollectionResponseBody {
	s.Metrics = &v
	return s
}

func (s *DescribeCollectionResponseBody) SetNamespace(v string) *DescribeCollectionResponseBody {
	s.Namespace = &v
	return s
}

func (s *DescribeCollectionResponseBody) SetParser(v string) *DescribeCollectionResponseBody {
	s.Parser = &v
	return s
}

func (s *DescribeCollectionResponseBody) SetRegionId(v string) *DescribeCollectionResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeCollectionResponseBody) SetRequestId(v string) *DescribeCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCollectionResponseBody) SetSparseVectorMetrics(v string) *DescribeCollectionResponseBody {
	s.SparseVectorMetrics = &v
	return s
}

func (s *DescribeCollectionResponseBody) SetStatus(v string) *DescribeCollectionResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeCollectionResponseBody) SetSupportSparse(v bool) *DescribeCollectionResponseBody {
	s.SupportSparse = &v
	return s
}

func (s *DescribeCollectionResponseBody) Validate() error {
	return dara.Validate(s)
}
