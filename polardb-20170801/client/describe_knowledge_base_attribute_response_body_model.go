// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKnowledgeBaseAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBindingAppCount(v int32) *DescribeKnowledgeBaseAttributeResponseBody
	GetBindingAppCount() *int32
	SetCreationTime(v string) *DescribeKnowledgeBaseAttributeResponseBody
	GetCreationTime() *string
	SetDescription(v string) *DescribeKnowledgeBaseAttributeResponseBody
	GetDescription() *string
	SetKnowledgeBaseId(v string) *DescribeKnowledgeBaseAttributeResponseBody
	GetKnowledgeBaseId() *string
	SetKnowledgeBaseType(v string) *DescribeKnowledgeBaseAttributeResponseBody
	GetKnowledgeBaseType() *string
	SetKnowledgeSpaceId(v string) *DescribeKnowledgeBaseAttributeResponseBody
	GetKnowledgeSpaceId() *string
	SetName(v string) *DescribeKnowledgeBaseAttributeResponseBody
	GetName() *string
	SetRequestId(v string) *DescribeKnowledgeBaseAttributeResponseBody
	GetRequestId() *string
	SetSearchMode(v string) *DescribeKnowledgeBaseAttributeResponseBody
	GetSearchMode() *string
	SetShardCount(v int32) *DescribeKnowledgeBaseAttributeResponseBody
	GetShardCount() *int32
	SetStatus(v string) *DescribeKnowledgeBaseAttributeResponseBody
	GetStatus() *string
	SetTotalDocs(v int32) *DescribeKnowledgeBaseAttributeResponseBody
	GetTotalDocs() *int32
	SetTotalSizeBytes(v int64) *DescribeKnowledgeBaseAttributeResponseBody
	GetTotalSizeBytes() *int64
}

type DescribeKnowledgeBaseAttributeResponseBody struct {
	// The number of AI applications bound to the knowledge base.
	//
	// example:
	//
	// 2
	BindingAppCount *int32 `json:"BindingAppCount,omitempty" xml:"BindingAppCount,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2020-08-14T05:58:42Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the knowledge base.
	//
	// example:
	//
	// testDesc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The unique identifier of the knowledge base.
	//
	// example:
	//
	// pkb-xxxxx
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The type of the knowledge base. Valid values:
	//
	// - PUBLIC
	//
	// - PERSONAL
	//
	// example:
	//
	// PUBLIC
	KnowledgeBaseType *string `json:"KnowledgeBaseType,omitempty" xml:"KnowledgeBaseType,omitempty"`
	// The ID of the knowledge space.
	//
	// example:
	//
	// pks-xxxxxx
	KnowledgeSpaceId *string `json:"KnowledgeSpaceId,omitempty" xml:"KnowledgeSpaceId,omitempty"`
	// The name of the knowledge base.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E56531A4-E552-40BA-9C58-137B80******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The search mode. Valid values:
	//
	// 	- balanced (default)
	//
	// 	- precise
	//
	// 	- semantic
	//
	// 	- knn
	//
	// 	- rrf
	//
	// example:
	//
	// balanced
	SearchMode *string `json:"SearchMode,omitempty" xml:"SearchMode,omitempty"`
	// The number of shards.
	//
	// example:
	//
	// 15
	ShardCount *int32 `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
	// The status of the knowledge base.
	//
	// example:
	//
	// Activation
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of documents.
	//
	// example:
	//
	// 1
	TotalDocs *int32 `json:"TotalDocs,omitempty" xml:"TotalDocs,omitempty"`
	// The total size in bytes.
	//
	// example:
	//
	// 318881
	TotalSizeBytes *int64 `json:"TotalSizeBytes,omitempty" xml:"TotalSizeBytes,omitempty"`
}

func (s DescribeKnowledgeBaseAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseAttributeResponseBody) GetBindingAppCount() *int32 {
	return s.BindingAppCount
}

func (s *DescribeKnowledgeBaseAttributeResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeKnowledgeBaseAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeKnowledgeBaseAttributeResponseBody) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *DescribeKnowledgeBaseAttributeResponseBody) GetKnowledgeBaseType() *string {
	return s.KnowledgeBaseType
}

func (s *DescribeKnowledgeBaseAttributeResponseBody) GetKnowledgeSpaceId() *string {
	return s.KnowledgeSpaceId
}

func (s *DescribeKnowledgeBaseAttributeResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeKnowledgeBaseAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeKnowledgeBaseAttributeResponseBody) GetSearchMode() *string {
	return s.SearchMode
}

func (s *DescribeKnowledgeBaseAttributeResponseBody) GetShardCount() *int32 {
	return s.ShardCount
}

func (s *DescribeKnowledgeBaseAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeKnowledgeBaseAttributeResponseBody) GetTotalDocs() *int32 {
	return s.TotalDocs
}

func (s *DescribeKnowledgeBaseAttributeResponseBody) GetTotalSizeBytes() *int64 {
	return s.TotalSizeBytes
}

func (s *DescribeKnowledgeBaseAttributeResponseBody) SetBindingAppCount(v int32) *DescribeKnowledgeBaseAttributeResponseBody {
	s.BindingAppCount = &v
	return s
}

func (s *DescribeKnowledgeBaseAttributeResponseBody) SetCreationTime(v string) *DescribeKnowledgeBaseAttributeResponseBody {
	s.CreationTime = &v
	return s
}

func (s *DescribeKnowledgeBaseAttributeResponseBody) SetDescription(v string) *DescribeKnowledgeBaseAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeKnowledgeBaseAttributeResponseBody) SetKnowledgeBaseId(v string) *DescribeKnowledgeBaseAttributeResponseBody {
	s.KnowledgeBaseId = &v
	return s
}

func (s *DescribeKnowledgeBaseAttributeResponseBody) SetKnowledgeBaseType(v string) *DescribeKnowledgeBaseAttributeResponseBody {
	s.KnowledgeBaseType = &v
	return s
}

func (s *DescribeKnowledgeBaseAttributeResponseBody) SetKnowledgeSpaceId(v string) *DescribeKnowledgeBaseAttributeResponseBody {
	s.KnowledgeSpaceId = &v
	return s
}

func (s *DescribeKnowledgeBaseAttributeResponseBody) SetName(v string) *DescribeKnowledgeBaseAttributeResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeKnowledgeBaseAttributeResponseBody) SetRequestId(v string) *DescribeKnowledgeBaseAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKnowledgeBaseAttributeResponseBody) SetSearchMode(v string) *DescribeKnowledgeBaseAttributeResponseBody {
	s.SearchMode = &v
	return s
}

func (s *DescribeKnowledgeBaseAttributeResponseBody) SetShardCount(v int32) *DescribeKnowledgeBaseAttributeResponseBody {
	s.ShardCount = &v
	return s
}

func (s *DescribeKnowledgeBaseAttributeResponseBody) SetStatus(v string) *DescribeKnowledgeBaseAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeKnowledgeBaseAttributeResponseBody) SetTotalDocs(v int32) *DescribeKnowledgeBaseAttributeResponseBody {
	s.TotalDocs = &v
	return s
}

func (s *DescribeKnowledgeBaseAttributeResponseBody) SetTotalSizeBytes(v int64) *DescribeKnowledgeBaseAttributeResponseBody {
	s.TotalSizeBytes = &v
	return s
}

func (s *DescribeKnowledgeBaseAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
