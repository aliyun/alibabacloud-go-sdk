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
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The queried document collections.
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type ListDocumentCollectionsResponseBodyItemsCollectionList struct {
	// The name of the document collection.
	//
	// example:
	//
	// document
	CollectionName *string `json:"CollectionName,omitempty" xml:"CollectionName,omitempty"`
	// The number of vector dimensions.
	//
	// example:
	//
	// 1536
	Dimension *int32 `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The name of the vector algorithm.
	//
	// example:
	//
	// text-embeddding-v1
	EmbeddingModel *string `json:"EmbeddingModel,omitempty" xml:"EmbeddingModel,omitempty"`
	// The fields that are used for full-text search. Multiple fields are separated by commas (,).
	//
	// example:
	//
	// title
	FullTextRetrievalFields *string `json:"FullTextRetrievalFields,omitempty" xml:"FullTextRetrievalFields,omitempty"`
	// The metadata.
	//
	// example:
	//
	// {\\"page\\": \\"int\\", \\"title\\": \\"text\\"}
	Metadata *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The vector similarity algorithm.
	//
	// example:
	//
	// cosine
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// The analyzer that is used for full-text search.
	//
	// example:
	//
	// zh_cn
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
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

func (s *ListDocumentCollectionsResponseBodyItemsCollectionList) Validate() error {
	return dara.Validate(s)
}
