// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInterventionDictionaryNerResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListInterventionDictionaryNerResultsResponseBody
	GetRequestId() *string
	SetResult(v []*ListInterventionDictionaryNerResultsResponseBodyResult) *ListInterventionDictionaryNerResultsResponseBody
	GetResult() []*ListInterventionDictionaryNerResultsResponseBodyResult
}

type ListInterventionDictionaryNerResultsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8F780CA8-D4D4-2FFE-B8AC-42040822C554
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The NER results.
	//
	// For more information, see [Named Entity Recognition (NER)](https://help.aliyun.com/document_detail/173606.html).
	Result []*ListInterventionDictionaryNerResultsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListInterventionDictionaryNerResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInterventionDictionaryNerResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryNerResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInterventionDictionaryNerResultsResponseBody) GetResult() []*ListInterventionDictionaryNerResultsResponseBodyResult {
	return s.Result
}

func (s *ListInterventionDictionaryNerResultsResponseBody) SetRequestId(v string) *ListInterventionDictionaryNerResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInterventionDictionaryNerResultsResponseBody) SetResult(v []*ListInterventionDictionaryNerResultsResponseBodyResult) *ListInterventionDictionaryNerResultsResponseBody {
	s.Result = v
	return s
}

func (s *ListInterventionDictionaryNerResultsResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInterventionDictionaryNerResultsResponseBodyResult struct {
	// The ordinal number.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// The tag of the detected entity.
	//
	// - brand: Brand
	//
	// - category: Category
	//
	// - material: Material
	//
	// - element: Element
	//
	// - style: Style
	//
	// - color: Color
	//
	// - function: Function
	//
	// - scenario: Scenario
	//
	// - people: People
	//
	// - season: Season
	//
	// - model: Model
	//
	// - region: Region
	//
	// - name: Name
	//
	// - adjective: Adjective
	//
	// - category-modifier: Category modifier
	//
	// - size: Size
	//
	// - quality: Quality
	//
	// - suit: Suit
	//
	// - new-release: New release
	//
	// - series: Series
	//
	// - marketing: Marketing
	//
	// - entertainment: Entertainment
	//
	// - organization: Organization
	//
	// - movie: Movie
	//
	// - game: Game
	//
	// - number: Number
	//
	// - unit: Unit
	//
	// - common: Common word
	//
	// - new-word: New word
	//
	// - proper-noun: Proper noun
	//
	// - symbol: Symbol
	//
	// - prefix: Prefix
	//
	// - suffix: Suffix
	//
	// - gift: Gift
	//
	// - negative: Negative
	//
	// - agent: Agent
	//
	// example:
	//
	// category
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The label of the tag.
	//
	// example:
	//
	// 品类
	TagLabel *string `json:"tagLabel,omitempty" xml:"tagLabel,omitempty"`
	// The detected entity.
	//
	// example:
	//
	// milk
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s ListInterventionDictionaryNerResultsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListInterventionDictionaryNerResultsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryNerResultsResponseBodyResult) GetOrder() *int32 {
	return s.Order
}

func (s *ListInterventionDictionaryNerResultsResponseBodyResult) GetTag() *string {
	return s.Tag
}

func (s *ListInterventionDictionaryNerResultsResponseBodyResult) GetTagLabel() *string {
	return s.TagLabel
}

func (s *ListInterventionDictionaryNerResultsResponseBodyResult) GetToken() *string {
	return s.Token
}

func (s *ListInterventionDictionaryNerResultsResponseBodyResult) SetOrder(v int32) *ListInterventionDictionaryNerResultsResponseBodyResult {
	s.Order = &v
	return s
}

func (s *ListInterventionDictionaryNerResultsResponseBodyResult) SetTag(v string) *ListInterventionDictionaryNerResultsResponseBodyResult {
	s.Tag = &v
	return s
}

func (s *ListInterventionDictionaryNerResultsResponseBodyResult) SetTagLabel(v string) *ListInterventionDictionaryNerResultsResponseBodyResult {
	s.TagLabel = &v
	return s
}

func (s *ListInterventionDictionaryNerResultsResponseBodyResult) SetToken(v string) *ListInterventionDictionaryNerResultsResponseBodyResult {
	s.Token = &v
	return s
}

func (s *ListInterventionDictionaryNerResultsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
