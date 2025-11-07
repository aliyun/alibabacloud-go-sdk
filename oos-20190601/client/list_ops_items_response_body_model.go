// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOpsItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListOpsItemsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListOpsItemsResponseBody
	GetNextToken() *string
	SetOpsItems(v []*ListOpsItemsResponseBodyOpsItems) *ListOpsItemsResponseBody
	GetOpsItems() []*ListOpsItemsResponseBodyOpsItems
	SetRequestId(v string) *ListOpsItemsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListOpsItemsResponseBody
	GetTotalCount() *int32
}

type ListOpsItemsResponseBody struct {
	// The number of entries returned on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// gAAAAABfTgv5ewUWmNdJ3g7JVLvX70sPH90GZOVGC6KPDUL0FIIb
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of O\\&M items.
	OpsItems []*ListOpsItemsResponseBodyOpsItems `json:"OpsItems,omitempty" xml:"OpsItems,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 113DD533-389C-5F83-9C69-F64D5BAB10B2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOpsItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOpsItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOpsItemsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListOpsItemsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOpsItemsResponseBody) GetOpsItems() []*ListOpsItemsResponseBodyOpsItems {
	return s.OpsItems
}

func (s *ListOpsItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOpsItemsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListOpsItemsResponseBody) SetMaxResults(v int32) *ListOpsItemsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListOpsItemsResponseBody) SetNextToken(v string) *ListOpsItemsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListOpsItemsResponseBody) SetOpsItems(v []*ListOpsItemsResponseBodyOpsItems) *ListOpsItemsResponseBody {
	s.OpsItems = v
	return s
}

func (s *ListOpsItemsResponseBody) SetRequestId(v string) *ListOpsItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOpsItemsResponseBody) SetTotalCount(v int32) *ListOpsItemsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOpsItemsResponseBody) Validate() error {
	if s.OpsItems != nil {
		for _, item := range s.OpsItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOpsItemsResponseBodyOpsItems struct {
	// The category.
	//
	// example:
	//
	// Security
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time when the O\\&M item was created.
	//
	// example:
	//
	// 2023-07-09T10:01Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The ID of the O\\&M item.
	//
	// example:
	//
	// oi-d52b08695e2b46ae8413
	OpsItemId *string `json:"OpsItemId,omitempty" xml:"OpsItemId,omitempty"`
	// The priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The Alibaba Resource Names (ARNs) of the associated resources.
	Resources []*string `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The severity level.
	//
	// example:
	//
	// Medium
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The source business.
	//
	// example:
	//
	// /aliyun/ecs
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The status of the O\\&M item.
	//
	// example:
	//
	// Open
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1":"v1"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the O\\&M item.
	//
	// example:
	//
	// Test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The time when the O\\&M item was updated.
	//
	// example:
	//
	// 2023-07-09T10:01Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListOpsItemsResponseBodyOpsItems) String() string {
	return dara.Prettify(s)
}

func (s ListOpsItemsResponseBodyOpsItems) GoString() string {
	return s.String()
}

func (s *ListOpsItemsResponseBodyOpsItems) GetCategory() *string {
	return s.Category
}

func (s *ListOpsItemsResponseBodyOpsItems) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListOpsItemsResponseBodyOpsItems) GetOpsItemId() *string {
	return s.OpsItemId
}

func (s *ListOpsItemsResponseBodyOpsItems) GetPriority() *int32 {
	return s.Priority
}

func (s *ListOpsItemsResponseBodyOpsItems) GetResources() []*string {
	return s.Resources
}

func (s *ListOpsItemsResponseBodyOpsItems) GetSeverity() *string {
	return s.Severity
}

func (s *ListOpsItemsResponseBodyOpsItems) GetSource() *string {
	return s.Source
}

func (s *ListOpsItemsResponseBodyOpsItems) GetStatus() *string {
	return s.Status
}

func (s *ListOpsItemsResponseBodyOpsItems) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *ListOpsItemsResponseBodyOpsItems) GetTitle() *string {
	return s.Title
}

func (s *ListOpsItemsResponseBodyOpsItems) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *ListOpsItemsResponseBodyOpsItems) SetCategory(v string) *ListOpsItemsResponseBodyOpsItems {
	s.Category = &v
	return s
}

func (s *ListOpsItemsResponseBodyOpsItems) SetCreateDate(v string) *ListOpsItemsResponseBodyOpsItems {
	s.CreateDate = &v
	return s
}

func (s *ListOpsItemsResponseBodyOpsItems) SetOpsItemId(v string) *ListOpsItemsResponseBodyOpsItems {
	s.OpsItemId = &v
	return s
}

func (s *ListOpsItemsResponseBodyOpsItems) SetPriority(v int32) *ListOpsItemsResponseBodyOpsItems {
	s.Priority = &v
	return s
}

func (s *ListOpsItemsResponseBodyOpsItems) SetResources(v []*string) *ListOpsItemsResponseBodyOpsItems {
	s.Resources = v
	return s
}

func (s *ListOpsItemsResponseBodyOpsItems) SetSeverity(v string) *ListOpsItemsResponseBodyOpsItems {
	s.Severity = &v
	return s
}

func (s *ListOpsItemsResponseBodyOpsItems) SetSource(v string) *ListOpsItemsResponseBodyOpsItems {
	s.Source = &v
	return s
}

func (s *ListOpsItemsResponseBodyOpsItems) SetStatus(v string) *ListOpsItemsResponseBodyOpsItems {
	s.Status = &v
	return s
}

func (s *ListOpsItemsResponseBodyOpsItems) SetTags(v map[string]interface{}) *ListOpsItemsResponseBodyOpsItems {
	s.Tags = v
	return s
}

func (s *ListOpsItemsResponseBodyOpsItems) SetTitle(v string) *ListOpsItemsResponseBodyOpsItems {
	s.Title = &v
	return s
}

func (s *ListOpsItemsResponseBodyOpsItems) SetUpdateDate(v string) *ListOpsItemsResponseBodyOpsItems {
	s.UpdateDate = &v
	return s
}

func (s *ListOpsItemsResponseBodyOpsItems) Validate() error {
	return dara.Validate(s)
}
