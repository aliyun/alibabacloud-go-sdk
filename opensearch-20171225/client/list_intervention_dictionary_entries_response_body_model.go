// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInterventionDictionaryEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListInterventionDictionaryEntriesResponseBody
	GetRequestId() *string
	SetResult(v []*ListInterventionDictionaryEntriesResponseBodyResult) *ListInterventionDictionaryEntriesResponseBody
	GetResult() []*ListInterventionDictionaryEntriesResponseBodyResult
	SetTotalCount(v int32) *ListInterventionDictionaryEntriesResponseBody
	GetTotalCount() *int32
}

type ListInterventionDictionaryEntriesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 516A02B7-2167-8D92-12D0-B639A2A0F3C5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about intervention entries.
	//
	// For more information, see [InterventionDictionaryEntry](https://help.aliyun.com/document_detail/173606.html).
	Result []*ListInterventionDictionaryEntriesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 8
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListInterventionDictionaryEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInterventionDictionaryEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInterventionDictionaryEntriesResponseBody) GetResult() []*ListInterventionDictionaryEntriesResponseBodyResult {
	return s.Result
}

func (s *ListInterventionDictionaryEntriesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListInterventionDictionaryEntriesResponseBody) SetRequestId(v string) *ListInterventionDictionaryEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBody) SetResult(v []*ListInterventionDictionaryEntriesResponseBodyResult) *ListInterventionDictionaryEntriesResponseBody {
	s.Result = v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBody) SetTotalCount(v int32) *ListInterventionDictionaryEntriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBody) Validate() error {
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

type ListInterventionDictionaryEntriesResponseBodyResult struct {
	// The command. Valid values:
	//
	// 	- add
	//
	// 	- delete
	//
	// example:
	//
	// add
	Cmd *string `json:"cmd,omitempty" xml:"cmd,omitempty"`
	// The timestamp when the intervention entry was created.
	//
	// example:
	//
	// 1536690285
	Created *int64 `json:"created,omitempty" xml:"created,omitempty"`
	// The content of an intervention entry for category prediction. The field value consists of key-value pairs. The key in a key-value pair indicates the ID of the category. The value in a key-value pair indicates the relevance to the category. A value of 0 indicates irrelevant. A value of 1 indicates slightly relevant. A value of 2 indicates relevant. Example: {"2":1, "100":0}
	//
	// example:
	//
	// {                 "100": "0",                 "200": "2"             }
	Relevance map[string]interface{} `json:"relevance,omitempty" xml:"relevance,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// 	- ACTIVE: The intervention entry takes effect.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The content of the intervention entry for term weight analysis.
	Tokens []*ListInterventionDictionaryEntriesResponseBodyResultTokens `json:"tokens,omitempty" xml:"tokens,omitempty" type:"Repeated"`
	// The timestamp when the intervention entry was last updated.
	//
	// example:
	//
	// 1537348987
	Updated *int64 `json:"updated,omitempty" xml:"updated,omitempty"`
	// The intervention entry.
	//
	// example:
	//
	// \\u8fc7\\u513f
	Word *string `json:"word,omitempty" xml:"word,omitempty"`
}

func (s ListInterventionDictionaryEntriesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListInterventionDictionaryEntriesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryEntriesResponseBodyResult) GetCmd() *string {
	return s.Cmd
}

func (s *ListInterventionDictionaryEntriesResponseBodyResult) GetCreated() *int64 {
	return s.Created
}

func (s *ListInterventionDictionaryEntriesResponseBodyResult) GetRelevance() map[string]interface{} {
	return s.Relevance
}

func (s *ListInterventionDictionaryEntriesResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListInterventionDictionaryEntriesResponseBodyResult) GetTokens() []*ListInterventionDictionaryEntriesResponseBodyResultTokens {
	return s.Tokens
}

func (s *ListInterventionDictionaryEntriesResponseBodyResult) GetUpdated() *int64 {
	return s.Updated
}

func (s *ListInterventionDictionaryEntriesResponseBodyResult) GetWord() *string {
	return s.Word
}

func (s *ListInterventionDictionaryEntriesResponseBodyResult) SetCmd(v string) *ListInterventionDictionaryEntriesResponseBodyResult {
	s.Cmd = &v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBodyResult) SetCreated(v int64) *ListInterventionDictionaryEntriesResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBodyResult) SetRelevance(v map[string]interface{}) *ListInterventionDictionaryEntriesResponseBodyResult {
	s.Relevance = v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBodyResult) SetStatus(v string) *ListInterventionDictionaryEntriesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBodyResult) SetTokens(v []*ListInterventionDictionaryEntriesResponseBodyResultTokens) *ListInterventionDictionaryEntriesResponseBodyResult {
	s.Tokens = v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBodyResult) SetUpdated(v int64) *ListInterventionDictionaryEntriesResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBodyResult) SetWord(v string) *ListInterventionDictionaryEntriesResponseBodyResult {
	s.Word = &v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBodyResult) Validate() error {
	if s.Tokens != nil {
		for _, item := range s.Tokens {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInterventionDictionaryEntriesResponseBodyResultTokens struct {
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
	// category
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s ListInterventionDictionaryEntriesResponseBodyResultTokens) String() string {
	return dara.Prettify(s)
}

func (s ListInterventionDictionaryEntriesResponseBodyResultTokens) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryEntriesResponseBodyResultTokens) GetOrder() *int32 {
	return s.Order
}

func (s *ListInterventionDictionaryEntriesResponseBodyResultTokens) GetTag() *string {
	return s.Tag
}

func (s *ListInterventionDictionaryEntriesResponseBodyResultTokens) GetTagLabel() *string {
	return s.TagLabel
}

func (s *ListInterventionDictionaryEntriesResponseBodyResultTokens) GetToken() *string {
	return s.Token
}

func (s *ListInterventionDictionaryEntriesResponseBodyResultTokens) SetOrder(v int32) *ListInterventionDictionaryEntriesResponseBodyResultTokens {
	s.Order = &v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBodyResultTokens) SetTag(v string) *ListInterventionDictionaryEntriesResponseBodyResultTokens {
	s.Tag = &v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBodyResultTokens) SetTagLabel(v string) *ListInterventionDictionaryEntriesResponseBodyResultTokens {
	s.TagLabel = &v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBodyResultTokens) SetToken(v string) *ListInterventionDictionaryEntriesResponseBodyResultTokens {
	s.Token = &v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBodyResultTokens) Validate() error {
	return dara.Validate(s)
}
