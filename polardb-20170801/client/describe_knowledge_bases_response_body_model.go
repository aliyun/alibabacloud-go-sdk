// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKnowledgeBasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeKnowledgeBasesResponseBodyItems) *DescribeKnowledgeBasesResponseBody
	GetItems() []*DescribeKnowledgeBasesResponseBodyItems
	SetPageNumber(v int32) *DescribeKnowledgeBasesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeKnowledgeBasesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeKnowledgeBasesResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeKnowledgeBasesResponseBody
	GetTotalRecordCount() *int32
}

type DescribeKnowledgeBasesResponseBody struct {
	// The list of knowledge bases.
	Items []*DescribeKnowledgeBasesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **30**, **50**, and **100**.
	//
	//
	//
	// Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CDB3258F-B5DE-43C4-8935-CBA0CA******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeKnowledgeBasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBasesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBasesResponseBody) GetItems() []*DescribeKnowledgeBasesResponseBodyItems {
	return s.Items
}

func (s *DescribeKnowledgeBasesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeKnowledgeBasesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeKnowledgeBasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeKnowledgeBasesResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeKnowledgeBasesResponseBody) SetItems(v []*DescribeKnowledgeBasesResponseBodyItems) *DescribeKnowledgeBasesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeKnowledgeBasesResponseBody) SetPageNumber(v int32) *DescribeKnowledgeBasesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeKnowledgeBasesResponseBody) SetPageSize(v int32) *DescribeKnowledgeBasesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeKnowledgeBasesResponseBody) SetRequestId(v string) *DescribeKnowledgeBasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKnowledgeBasesResponseBody) SetTotalRecordCount(v int32) *DescribeKnowledgeBasesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeKnowledgeBasesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeKnowledgeBasesResponseBodyItems struct {
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
	// 2025-03-25T09:37:10Z
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
	// pkb-xxxxxx
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The type of the knowledge base. Valid values:
	//
	// - PUBLIC: public.
	//
	// - PERSONAL: personal.
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
	// 231984
	TotalSizeBytes *int64 `json:"TotalSizeBytes,omitempty" xml:"TotalSizeBytes,omitempty"`
}

func (s DescribeKnowledgeBasesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBasesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBasesResponseBodyItems) GetBindingAppCount() *int32 {
	return s.BindingAppCount
}

func (s *DescribeKnowledgeBasesResponseBodyItems) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeKnowledgeBasesResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *DescribeKnowledgeBasesResponseBodyItems) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *DescribeKnowledgeBasesResponseBodyItems) GetKnowledgeBaseType() *string {
	return s.KnowledgeBaseType
}

func (s *DescribeKnowledgeBasesResponseBodyItems) GetKnowledgeSpaceId() *string {
	return s.KnowledgeSpaceId
}

func (s *DescribeKnowledgeBasesResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *DescribeKnowledgeBasesResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeKnowledgeBasesResponseBodyItems) GetTotalDocs() *int32 {
	return s.TotalDocs
}

func (s *DescribeKnowledgeBasesResponseBodyItems) GetTotalSizeBytes() *int64 {
	return s.TotalSizeBytes
}

func (s *DescribeKnowledgeBasesResponseBodyItems) SetBindingAppCount(v int32) *DescribeKnowledgeBasesResponseBodyItems {
	s.BindingAppCount = &v
	return s
}

func (s *DescribeKnowledgeBasesResponseBodyItems) SetCreationTime(v string) *DescribeKnowledgeBasesResponseBodyItems {
	s.CreationTime = &v
	return s
}

func (s *DescribeKnowledgeBasesResponseBodyItems) SetDescription(v string) *DescribeKnowledgeBasesResponseBodyItems {
	s.Description = &v
	return s
}

func (s *DescribeKnowledgeBasesResponseBodyItems) SetKnowledgeBaseId(v string) *DescribeKnowledgeBasesResponseBodyItems {
	s.KnowledgeBaseId = &v
	return s
}

func (s *DescribeKnowledgeBasesResponseBodyItems) SetKnowledgeBaseType(v string) *DescribeKnowledgeBasesResponseBodyItems {
	s.KnowledgeBaseType = &v
	return s
}

func (s *DescribeKnowledgeBasesResponseBodyItems) SetKnowledgeSpaceId(v string) *DescribeKnowledgeBasesResponseBodyItems {
	s.KnowledgeSpaceId = &v
	return s
}

func (s *DescribeKnowledgeBasesResponseBodyItems) SetName(v string) *DescribeKnowledgeBasesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeKnowledgeBasesResponseBodyItems) SetStatus(v string) *DescribeKnowledgeBasesResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeKnowledgeBasesResponseBodyItems) SetTotalDocs(v int32) *DescribeKnowledgeBasesResponseBodyItems {
	s.TotalDocs = &v
	return s
}

func (s *DescribeKnowledgeBasesResponseBodyItems) SetTotalSizeBytes(v int64) *DescribeKnowledgeBasesResponseBodyItems {
	s.TotalSizeBytes = &v
	return s
}

func (s *DescribeKnowledgeBasesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
