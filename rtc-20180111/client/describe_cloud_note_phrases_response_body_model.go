// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudNotePhrasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPhrases(v []*DescribeCloudNotePhrasesResponseBodyPhrases) *DescribeCloudNotePhrasesResponseBody
	GetPhrases() []*DescribeCloudNotePhrasesResponseBodyPhrases
	SetRequestId(v string) *DescribeCloudNotePhrasesResponseBody
	GetRequestId() *string
	SetTotalNum(v int64) *DescribeCloudNotePhrasesResponseBody
	GetTotalNum() *int64
	SetTotalPage(v int64) *DescribeCloudNotePhrasesResponseBody
	GetTotalPage() *int64
}

type DescribeCloudNotePhrasesResponseBody struct {
	Phrases []*DescribeCloudNotePhrasesResponseBodyPhrases `json:"Phrases,omitempty" xml:"Phrases,omitempty" type:"Repeated"`
	// Id of the request。
	//
	// example:
	//
	// 3A26E1E3-3CBB-599E-AD68-CB78F5A42FA1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// example:
	//
	// 1
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeCloudNotePhrasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudNotePhrasesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudNotePhrasesResponseBody) GetPhrases() []*DescribeCloudNotePhrasesResponseBodyPhrases {
	return s.Phrases
}

func (s *DescribeCloudNotePhrasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudNotePhrasesResponseBody) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *DescribeCloudNotePhrasesResponseBody) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *DescribeCloudNotePhrasesResponseBody) SetPhrases(v []*DescribeCloudNotePhrasesResponseBodyPhrases) *DescribeCloudNotePhrasesResponseBody {
	s.Phrases = v
	return s
}

func (s *DescribeCloudNotePhrasesResponseBody) SetRequestId(v string) *DescribeCloudNotePhrasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudNotePhrasesResponseBody) SetTotalNum(v int64) *DescribeCloudNotePhrasesResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeCloudNotePhrasesResponseBody) SetTotalPage(v int64) *DescribeCloudNotePhrasesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeCloudNotePhrasesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudNotePhrasesResponseBodyPhrases struct {
	// example:
	//
	// 2025-03-04T06:22:15Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 水果描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1qweadca332121212
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 水果
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	WordWeights []*DescribeCloudNotePhrasesResponseBodyPhrasesWordWeights `json:"WordWeights,omitempty" xml:"WordWeights,omitempty" type:"Repeated"`
}

func (s DescribeCloudNotePhrasesResponseBodyPhrases) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudNotePhrasesResponseBodyPhrases) GoString() string {
	return s.String()
}

func (s *DescribeCloudNotePhrasesResponseBodyPhrases) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeCloudNotePhrasesResponseBodyPhrases) GetDescription() *string {
	return s.Description
}

func (s *DescribeCloudNotePhrasesResponseBodyPhrases) GetId() *string {
	return s.Id
}

func (s *DescribeCloudNotePhrasesResponseBodyPhrases) GetName() *string {
	return s.Name
}

func (s *DescribeCloudNotePhrasesResponseBodyPhrases) GetWordWeights() []*DescribeCloudNotePhrasesResponseBodyPhrasesWordWeights {
	return s.WordWeights
}

func (s *DescribeCloudNotePhrasesResponseBodyPhrases) SetCreateTime(v string) *DescribeCloudNotePhrasesResponseBodyPhrases {
	s.CreateTime = &v
	return s
}

func (s *DescribeCloudNotePhrasesResponseBodyPhrases) SetDescription(v string) *DescribeCloudNotePhrasesResponseBodyPhrases {
	s.Description = &v
	return s
}

func (s *DescribeCloudNotePhrasesResponseBodyPhrases) SetId(v string) *DescribeCloudNotePhrasesResponseBodyPhrases {
	s.Id = &v
	return s
}

func (s *DescribeCloudNotePhrasesResponseBodyPhrases) SetName(v string) *DescribeCloudNotePhrasesResponseBodyPhrases {
	s.Name = &v
	return s
}

func (s *DescribeCloudNotePhrasesResponseBodyPhrases) SetWordWeights(v []*DescribeCloudNotePhrasesResponseBodyPhrasesWordWeights) *DescribeCloudNotePhrasesResponseBodyPhrases {
	s.WordWeights = v
	return s
}

func (s *DescribeCloudNotePhrasesResponseBodyPhrases) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudNotePhrasesResponseBodyPhrasesWordWeights struct {
	// example:
	//
	// 0
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// example:
	//
	// 苹果
	Word *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s DescribeCloudNotePhrasesResponseBodyPhrasesWordWeights) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudNotePhrasesResponseBodyPhrasesWordWeights) GoString() string {
	return s.String()
}

func (s *DescribeCloudNotePhrasesResponseBodyPhrasesWordWeights) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeCloudNotePhrasesResponseBodyPhrasesWordWeights) GetWord() *string {
	return s.Word
}

func (s *DescribeCloudNotePhrasesResponseBodyPhrasesWordWeights) SetWeight(v int32) *DescribeCloudNotePhrasesResponseBodyPhrasesWordWeights {
	s.Weight = &v
	return s
}

func (s *DescribeCloudNotePhrasesResponseBodyPhrasesWordWeights) SetWord(v string) *DescribeCloudNotePhrasesResponseBodyPhrasesWordWeights {
	s.Word = &v
	return s
}

func (s *DescribeCloudNotePhrasesResponseBodyPhrasesWordWeights) Validate() error {
	return dara.Validate(s)
}
