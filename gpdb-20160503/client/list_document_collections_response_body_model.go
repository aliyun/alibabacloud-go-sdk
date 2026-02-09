// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDocumentCollectionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *ListDocumentCollectionsResponseBody
	GetCount() *int32
	SetItems(v *ListDocumentCollectionsResponseBodyItems) *ListDocumentCollectionsResponseBody
	GetItems() *ListDocumentCollectionsResponseBodyItems
	SetMessage(v string) *ListDocumentCollectionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListDocumentCollectionsResponseBody
	GetRequestId() *string
	SetStatus(v string) *ListDocumentCollectionsResponseBody
	GetStatus() *string
}

type ListDocumentCollectionsResponseBody struct {
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	Count *int32                                    `json:"Count,omitempty" xml:"Count,omitempty"`
	Items *ListDocumentCollectionsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **success**
	//
	// 	- **fail**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDocumentCollectionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDocumentCollectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDocumentCollectionsResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListDocumentCollectionsResponseBody) GetItems() *ListDocumentCollectionsResponseBodyItems {
	return s.Items
}

func (s *ListDocumentCollectionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDocumentCollectionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDocumentCollectionsResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ListDocumentCollectionsResponseBody) SetCount(v int32) *ListDocumentCollectionsResponseBody {
	s.Count = &v
	return s
}

func (s *ListDocumentCollectionsResponseBody) SetItems(v *ListDocumentCollectionsResponseBodyItems) *ListDocumentCollectionsResponseBody {
	s.Items = v
	return s
}

func (s *ListDocumentCollectionsResponseBody) SetMessage(v string) *ListDocumentCollectionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDocumentCollectionsResponseBody) SetRequestId(v string) *ListDocumentCollectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDocumentCollectionsResponseBody) SetStatus(v string) *ListDocumentCollectionsResponseBody {
	s.Status = &v
	return s
}

func (s *ListDocumentCollectionsResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDocumentCollectionsResponseBodyItems struct {
	CollectionList []*ListDocumentCollectionsResponseBodyItemsCollectionList `json:"CollectionList,omitempty" xml:"CollectionList,omitempty" type:"Repeated"`
}

func (s ListDocumentCollectionsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListDocumentCollectionsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListDocumentCollectionsResponseBodyItems) GetCollectionList() []*ListDocumentCollectionsResponseBodyItemsCollectionList {
	return s.CollectionList
}

func (s *ListDocumentCollectionsResponseBodyItems) SetCollectionList(v []*ListDocumentCollectionsResponseBodyItemsCollectionList) *ListDocumentCollectionsResponseBodyItems {
	s.CollectionList = v
	return s
}

func (s *ListDocumentCollectionsResponseBodyItems) Validate() error {
	if s.CollectionList != nil {
		for _, item := range s.CollectionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDocumentCollectionsResponseBodyItemsCollectionList struct {
	CollectionName          *string `json:"CollectionName,omitempty" xml:"CollectionName,omitempty"`
	Dimension               *int32  `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	EmbeddingModel          *string `json:"EmbeddingModel,omitempty" xml:"EmbeddingModel,omitempty"`
	FullTextRetrievalFields *string `json:"FullTextRetrievalFields,omitempty" xml:"FullTextRetrievalFields,omitempty"`
	Metadata                *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	Metrics                 *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	Parser                  *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
	SparseRetrievalFields   *string `json:"SparseRetrievalFields,omitempty" xml:"SparseRetrievalFields,omitempty"`
	SupportSparse           *bool   `json:"SupportSparse,omitempty" xml:"SupportSparse,omitempty"`
}

func (s ListDocumentCollectionsResponseBodyItemsCollectionList) String() string {
	return dara.Prettify(s)
}

func (s ListDocumentCollectionsResponseBodyItemsCollectionList) GoString() string {
	return s.String()
}

func (s *ListDocumentCollectionsResponseBodyItemsCollectionList) GetCollectionName() *string {
	return s.CollectionName
}

func (s *ListDocumentCollectionsResponseBodyItemsCollectionList) GetDimension() *int32 {
	return s.Dimension
}

func (s *ListDocumentCollectionsResponseBodyItemsCollectionList) GetEmbeddingModel() *string {
	return s.EmbeddingModel
}

func (s *ListDocumentCollectionsResponseBodyItemsCollectionList) GetFullTextRetrievalFields() *string {
	return s.FullTextRetrievalFields
}

func (s *ListDocumentCollectionsResponseBodyItemsCollectionList) GetMetadata() *string {
	return s.Metadata
}

func (s *ListDocumentCollectionsResponseBodyItemsCollectionList) GetMetrics() *string {
	return s.Metrics
}

func (s *ListDocumentCollectionsResponseBodyItemsCollectionList) GetParser() *string {
	return s.Parser
}

func (s *ListDocumentCollectionsResponseBodyItemsCollectionList) GetSparseRetrievalFields() *string {
	return s.SparseRetrievalFields
}

func (s *ListDocumentCollectionsResponseBodyItemsCollectionList) GetSupportSparse() *bool {
	return s.SupportSparse
}

func (s *ListDocumentCollectionsResponseBodyItemsCollectionList) SetCollectionName(v string) *ListDocumentCollectionsResponseBodyItemsCollectionList {
	s.CollectionName = &v
	return s
}

func (s *ListDocumentCollectionsResponseBodyItemsCollectionList) SetDimension(v int32) *ListDocumentCollectionsResponseBodyItemsCollectionList {
	s.Dimension = &v
	return s
}

func (s *ListDocumentCollectionsResponseBodyItemsCollectionList) SetEmbeddingModel(v string) *ListDocumentCollectionsResponseBodyItemsCollectionList {
	s.EmbeddingModel = &v
	return s
}

func (s *ListDocumentCollectionsResponseBodyItemsCollectionList) SetFullTextRetrievalFields(v string) *ListDocumentCollectionsResponseBodyItemsCollectionList {
	s.FullTextRetrievalFields = &v
	return s
}

func (s *ListDocumentCollectionsResponseBodyItemsCollectionList) SetMetadata(v string) *ListDocumentCollectionsResponseBodyItemsCollectionList {
	s.Metadata = &v
	return s
}

func (s *ListDocumentCollectionsResponseBodyItemsCollectionList) SetMetrics(v string) *ListDocumentCollectionsResponseBodyItemsCollectionList {
	s.Metrics = &v
	return s
}

func (s *ListDocumentCollectionsResponseBodyItemsCollectionList) SetParser(v string) *ListDocumentCollectionsResponseBodyItemsCollectionList {
	s.Parser = &v
	return s
}

func (s *ListDocumentCollectionsResponseBodyItemsCollectionList) SetSparseRetrievalFields(v string) *ListDocumentCollectionsResponseBodyItemsCollectionList {
	s.SparseRetrievalFields = &v
	return s
}

func (s *ListDocumentCollectionsResponseBodyItemsCollectionList) SetSupportSparse(v bool) *ListDocumentCollectionsResponseBodyItemsCollectionList {
	s.SupportSparse = &v
	return s
}

func (s *ListDocumentCollectionsResponseBodyItemsCollectionList) Validate() error {
	return dara.Validate(s)
}
