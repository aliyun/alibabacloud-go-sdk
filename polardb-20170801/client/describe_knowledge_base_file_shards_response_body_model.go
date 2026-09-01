// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKnowledgeBaseFileShardsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeKnowledgeBaseFileShardsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeKnowledgeBaseFileShardsResponseBody
	GetPageRecordCount() *int32
	SetPageSize(v int32) *DescribeKnowledgeBaseFileShardsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeKnowledgeBaseFileShardsResponseBody
	GetRequestId() *string
	SetShards(v []*DescribeKnowledgeBaseFileShardsResponseBodyShards) *DescribeKnowledgeBaseFileShardsResponseBody
	GetShards() []*DescribeKnowledgeBaseFileShardsResponseBodyShards
	SetTotalRecordCount(v int32) *DescribeKnowledgeBaseFileShardsResponseBody
	GetTotalRecordCount() *int32
}

type DescribeKnowledgeBaseFileShardsResponseBody struct {
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records on the current page.
	//
	// example:
	//
	// 10
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The number of entries per page in a paged query.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3E5CD764-xxxx-xxxx-xxxx-20E0DE84B2AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The shard information.
	Shards []*DescribeKnowledgeBaseFileShardsResponseBodyShards `json:"Shards,omitempty" xml:"Shards,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 10
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeKnowledgeBaseFileShardsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseFileShardsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseFileShardsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeKnowledgeBaseFileShardsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeKnowledgeBaseFileShardsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeKnowledgeBaseFileShardsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeKnowledgeBaseFileShardsResponseBody) GetShards() []*DescribeKnowledgeBaseFileShardsResponseBodyShards {
	return s.Shards
}

func (s *DescribeKnowledgeBaseFileShardsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeKnowledgeBaseFileShardsResponseBody) SetPageNumber(v int32) *DescribeKnowledgeBaseFileShardsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeKnowledgeBaseFileShardsResponseBody) SetPageRecordCount(v int32) *DescribeKnowledgeBaseFileShardsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeKnowledgeBaseFileShardsResponseBody) SetPageSize(v int32) *DescribeKnowledgeBaseFileShardsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeKnowledgeBaseFileShardsResponseBody) SetRequestId(v string) *DescribeKnowledgeBaseFileShardsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKnowledgeBaseFileShardsResponseBody) SetShards(v []*DescribeKnowledgeBaseFileShardsResponseBodyShards) *DescribeKnowledgeBaseFileShardsResponseBody {
	s.Shards = v
	return s
}

func (s *DescribeKnowledgeBaseFileShardsResponseBody) SetTotalRecordCount(v int32) *DescribeKnowledgeBaseFileShardsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeKnowledgeBaseFileShardsResponseBody) Validate() error {
	if s.Shards != nil {
		for _, item := range s.Shards {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeKnowledgeBaseFileShardsResponseBodyShards struct {
	// The chain of section headings to which the shard belongs.
	Headings []*string `json:"Headings,omitempty" xml:"Headings,omitempty" type:"Repeated"`
	// The list of page numbers to which the shard belongs.
	PageNumbers []*string `json:"PageNumbers,omitempty" xml:"PageNumbers,omitempty" type:"Repeated"`
	// The text content of the shard.
	//
	// example:
	//
	// ******
	ShardContent *string `json:"ShardContent,omitempty" xml:"ShardContent,omitempty"`
	// The shard index.
	//
	// example:
	//
	// 1
	ShardIndex *int32 `json:"ShardIndex,omitempty" xml:"ShardIndex,omitempty"`
}

func (s DescribeKnowledgeBaseFileShardsResponseBodyShards) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseFileShardsResponseBodyShards) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseFileShardsResponseBodyShards) GetHeadings() []*string {
	return s.Headings
}

func (s *DescribeKnowledgeBaseFileShardsResponseBodyShards) GetPageNumbers() []*string {
	return s.PageNumbers
}

func (s *DescribeKnowledgeBaseFileShardsResponseBodyShards) GetShardContent() *string {
	return s.ShardContent
}

func (s *DescribeKnowledgeBaseFileShardsResponseBodyShards) GetShardIndex() *int32 {
	return s.ShardIndex
}

func (s *DescribeKnowledgeBaseFileShardsResponseBodyShards) SetHeadings(v []*string) *DescribeKnowledgeBaseFileShardsResponseBodyShards {
	s.Headings = v
	return s
}

func (s *DescribeKnowledgeBaseFileShardsResponseBodyShards) SetPageNumbers(v []*string) *DescribeKnowledgeBaseFileShardsResponseBodyShards {
	s.PageNumbers = v
	return s
}

func (s *DescribeKnowledgeBaseFileShardsResponseBodyShards) SetShardContent(v string) *DescribeKnowledgeBaseFileShardsResponseBodyShards {
	s.ShardContent = &v
	return s
}

func (s *DescribeKnowledgeBaseFileShardsResponseBodyShards) SetShardIndex(v int32) *DescribeKnowledgeBaseFileShardsResponseBodyShards {
	s.ShardIndex = &v
	return s
}

func (s *DescribeKnowledgeBaseFileShardsResponseBodyShards) Validate() error {
	return dara.Validate(s)
}
