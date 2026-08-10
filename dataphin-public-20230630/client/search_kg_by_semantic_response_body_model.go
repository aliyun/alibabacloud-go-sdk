// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchKgBySemanticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SearchKgBySemanticResponseBody
	GetCode() *string
	SetData(v *SearchKgBySemanticResponseBodyData) *SearchKgBySemanticResponseBody
	GetData() *SearchKgBySemanticResponseBodyData
	SetHttpStatusCode(v int32) *SearchKgBySemanticResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SearchKgBySemanticResponseBody
	GetMessage() *string
	SetRequestId(v string) *SearchKgBySemanticResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SearchKgBySemanticResponseBody
	GetSuccess() *bool
}

type SearchKgBySemanticResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The search results.
	Data *SearchKgBySemanticResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend exception details.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchKgBySemanticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchKgBySemanticResponseBody) GoString() string {
	return s.String()
}

func (s *SearchKgBySemanticResponseBody) GetCode() *string {
	return s.Code
}

func (s *SearchKgBySemanticResponseBody) GetData() *SearchKgBySemanticResponseBodyData {
	return s.Data
}

func (s *SearchKgBySemanticResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SearchKgBySemanticResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SearchKgBySemanticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchKgBySemanticResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchKgBySemanticResponseBody) SetCode(v string) *SearchKgBySemanticResponseBody {
	s.Code = &v
	return s
}

func (s *SearchKgBySemanticResponseBody) SetData(v *SearchKgBySemanticResponseBodyData) *SearchKgBySemanticResponseBody {
	s.Data = v
	return s
}

func (s *SearchKgBySemanticResponseBody) SetHttpStatusCode(v int32) *SearchKgBySemanticResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SearchKgBySemanticResponseBody) SetMessage(v string) *SearchKgBySemanticResponseBody {
	s.Message = &v
	return s
}

func (s *SearchKgBySemanticResponseBody) SetRequestId(v string) *SearchKgBySemanticResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchKgBySemanticResponseBody) SetSuccess(v bool) *SearchKgBySemanticResponseBody {
	s.Success = &v
	return s
}

func (s *SearchKgBySemanticResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchKgBySemanticResponseBodyData struct {
	// The list of search results.
	SearchResults []*SearchKgBySemanticResponseBodyDataSearchResults `json:"SearchResults,omitempty" xml:"SearchResults,omitempty" type:"Repeated"`
	// The total number of results.
	//
	// example:
	//
	// 1640276
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchKgBySemanticResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SearchKgBySemanticResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchKgBySemanticResponseBodyData) GetSearchResults() []*SearchKgBySemanticResponseBodyDataSearchResults {
	return s.SearchResults
}

func (s *SearchKgBySemanticResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *SearchKgBySemanticResponseBodyData) SetSearchResults(v []*SearchKgBySemanticResponseBodyDataSearchResults) *SearchKgBySemanticResponseBodyData {
	s.SearchResults = v
	return s
}

func (s *SearchKgBySemanticResponseBodyData) SetTotalCount(v int32) *SearchKgBySemanticResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *SearchKgBySemanticResponseBodyData) Validate() error {
	if s.SearchResults != nil {
		for _, item := range s.SearchResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchKgBySemanticResponseBodyDataSearchResults struct {
	// The ID of the matched entity record.
	//
	// example:
	//
	// itemId
	ItemId *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	// The entity type code.
	//
	// example:
	//
	// EntityTypeCode1
	ItemTypeCode *string `json:"ItemTypeCode,omitempty" xml:"ItemTypeCode,omitempty"`
	// The property code that matched the semantic search.
	//
	// example:
	//
	// name
	MatchedPropertyCode *string `json:"MatchedPropertyCode,omitempty" xml:"MatchedPropertyCode,omitempty"`
	// The actual value of the matched property.
	//
	// example:
	//
	// 张三
	MatchedPropertyValue *string `json:"MatchedPropertyValue,omitempty" xml:"MatchedPropertyValue,omitempty"`
	// The similarity score ranging from 0.0 to 1.0, based on cosine similarity.
	//
	// example:
	//
	// 0.88
	SimilarityScore *float32 `json:"SimilarityScore,omitempty" xml:"SimilarityScore,omitempty"`
}

func (s SearchKgBySemanticResponseBodyDataSearchResults) String() string {
	return dara.Prettify(s)
}

func (s SearchKgBySemanticResponseBodyDataSearchResults) GoString() string {
	return s.String()
}

func (s *SearchKgBySemanticResponseBodyDataSearchResults) GetItemId() *string {
	return s.ItemId
}

func (s *SearchKgBySemanticResponseBodyDataSearchResults) GetItemTypeCode() *string {
	return s.ItemTypeCode
}

func (s *SearchKgBySemanticResponseBodyDataSearchResults) GetMatchedPropertyCode() *string {
	return s.MatchedPropertyCode
}

func (s *SearchKgBySemanticResponseBodyDataSearchResults) GetMatchedPropertyValue() *string {
	return s.MatchedPropertyValue
}

func (s *SearchKgBySemanticResponseBodyDataSearchResults) GetSimilarityScore() *float32 {
	return s.SimilarityScore
}

func (s *SearchKgBySemanticResponseBodyDataSearchResults) SetItemId(v string) *SearchKgBySemanticResponseBodyDataSearchResults {
	s.ItemId = &v
	return s
}

func (s *SearchKgBySemanticResponseBodyDataSearchResults) SetItemTypeCode(v string) *SearchKgBySemanticResponseBodyDataSearchResults {
	s.ItemTypeCode = &v
	return s
}

func (s *SearchKgBySemanticResponseBodyDataSearchResults) SetMatchedPropertyCode(v string) *SearchKgBySemanticResponseBodyDataSearchResults {
	s.MatchedPropertyCode = &v
	return s
}

func (s *SearchKgBySemanticResponseBodyDataSearchResults) SetMatchedPropertyValue(v string) *SearchKgBySemanticResponseBodyDataSearchResults {
	s.MatchedPropertyValue = &v
	return s
}

func (s *SearchKgBySemanticResponseBodyDataSearchResults) SetSimilarityScore(v float32) *SearchKgBySemanticResponseBodyDataSearchResults {
	s.SimilarityScore = &v
	return s
}

func (s *SearchKgBySemanticResponseBodyDataSearchResults) Validate() error {
	return dara.Validate(s)
}
