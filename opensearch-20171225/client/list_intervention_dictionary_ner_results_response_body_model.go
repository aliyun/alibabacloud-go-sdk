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
	// The ID of the request.
	//
	// example:
	//
	// 8F780CA8-D4D4-2FFE-B8AC-42040822C554
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The NER result.
	//
	// For more information, see [InterventionDictionaryEntry](https://help.aliyun.com/document_detail/173606.html).
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
	// The sequence number.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// The internal name of the identified entity type. Valid values:
	//
	// 	- brand
	//
	// 	- category
	//
	// 	- material
	//
	// 	- element
	//
	// 	- style
	//
	// 	- color
	//
	// 	- function
	//
	// 	- scenario
	//
	// 	- people
	//
	// 	- season
	//
	// 	- model
	//
	// 	- region
	//
	// 	- name
	//
	// 	- adjective
	//
	// 	- category-modifier
	//
	// 	- size
	//
	// 	- quality
	//
	// 	- suit
	//
	// 	- new-release
	//
	// 	- series
	//
	// 	- marketing
	//
	// 	- entertainment
	//
	// 	- organization
	//
	// 	- movie
	//
	// 	- game
	//
	// 	- number
	//
	// 	- unit
	//
	// 	- common
	//
	// 	- new-word
	//
	// 	- proper-noun
	//
	// 	- symbol
	//
	// 	- prefix
	//
	// 	- suffix
	//
	// 	- gift
	//
	// 	- negative
	//
	// 	- agent
	//
	// example:
	//
	// category
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The description of the internal name of the identified entity type.
	//
	// example:
	//
	// category
	TagLabel *string `json:"tagLabel,omitempty" xml:"tagLabel,omitempty"`
	// The entity.
	//
	// example:
	//
	// eaa73f35-007a-4be7-88c7-37dca4a04ab7
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
