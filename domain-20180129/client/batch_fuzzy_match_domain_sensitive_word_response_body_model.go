// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchFuzzyMatchDomainSensitiveWordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchFuzzyMatchDomainSensitiveWordResponseBody
	GetRequestId() *string
	SetSensitiveWordMatchResultList(v *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultList) *BatchFuzzyMatchDomainSensitiveWordResponseBody
	GetSensitiveWordMatchResultList() *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultList
}

type BatchFuzzyMatchDomainSensitiveWordResponseBody struct {
	// example:
	//
	// C560A803-B975-481D-A66B-A4395EA863A1
	RequestId                    *string                                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SensitiveWordMatchResultList *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultList `json:"SensitiveWordMatchResultList,omitempty" xml:"SensitiveWordMatchResultList,omitempty" type:"Struct"`
}

func (s BatchFuzzyMatchDomainSensitiveWordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchFuzzyMatchDomainSensitiveWordResponseBody) GoString() string {
	return s.String()
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponseBody) GetSensitiveWordMatchResultList() *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultList {
	return s.SensitiveWordMatchResultList
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponseBody) SetRequestId(v string) *BatchFuzzyMatchDomainSensitiveWordResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponseBody) SetSensitiveWordMatchResultList(v *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultList) *BatchFuzzyMatchDomainSensitiveWordResponseBody {
	s.SensitiveWordMatchResultList = v
	return s
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponseBody) Validate() error {
	if s.SensitiveWordMatchResultList != nil {
		if err := s.SensitiveWordMatchResultList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultList struct {
	SensitiveWordMatchResult []*BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResult `json:"SensitiveWordMatchResult,omitempty" xml:"SensitiveWordMatchResult,omitempty" type:"Repeated"`
}

func (s BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultList) String() string {
	return dara.Prettify(s)
}

func (s BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultList) GoString() string {
	return s.String()
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultList) GetSensitiveWordMatchResult() []*BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResult {
	return s.SensitiveWordMatchResult
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultList) SetSensitiveWordMatchResult(v []*BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResult) *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultList {
	s.SensitiveWordMatchResult = v
	return s
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultList) Validate() error {
	if s.SensitiveWordMatchResult != nil {
		for _, item := range s.SensitiveWordMatchResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResult struct {
	// example:
	//
	// true
	Exist *bool `json:"Exist,omitempty" xml:"Exist,omitempty"`
	// example:
	//
	// example.com,aliyundoc.com
	Keyword             *string                                                                                                                `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	MatchedSentiveWords *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWords `json:"MatchedSentiveWords,omitempty" xml:"MatchedSentiveWords,omitempty" type:"Struct"`
}

func (s BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResult) String() string {
	return dara.Prettify(s)
}

func (s BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResult) GoString() string {
	return s.String()
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResult) GetExist() *bool {
	return s.Exist
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResult) GetKeyword() *string {
	return s.Keyword
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResult) GetMatchedSentiveWords() *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWords {
	return s.MatchedSentiveWords
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResult) SetExist(v bool) *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResult {
	s.Exist = &v
	return s
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResult) SetKeyword(v string) *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResult {
	s.Keyword = &v
	return s
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResult) SetMatchedSentiveWords(v *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWords) *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResult {
	s.MatchedSentiveWords = v
	return s
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResult) Validate() error {
	if s.MatchedSentiveWords != nil {
		if err := s.MatchedSentiveWords.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWords struct {
	MatchedSensitiveWord []*BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWordsMatchedSensitiveWord `json:"MatchedSensitiveWord,omitempty" xml:"MatchedSensitiveWord,omitempty" type:"Repeated"`
}

func (s BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWords) String() string {
	return dara.Prettify(s)
}

func (s BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWords) GoString() string {
	return s.String()
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWords) GetMatchedSensitiveWord() []*BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWordsMatchedSensitiveWord {
	return s.MatchedSensitiveWord
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWords) SetMatchedSensitiveWord(v []*BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWordsMatchedSensitiveWord) *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWords {
	s.MatchedSensitiveWord = v
	return s
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWords) Validate() error {
	if s.MatchedSensitiveWord != nil {
		for _, item := range s.MatchedSensitiveWord {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWordsMatchedSensitiveWord struct {
	// example:
	//
	// xxx
	Word *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWordsMatchedSensitiveWord) String() string {
	return dara.Prettify(s)
}

func (s BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWordsMatchedSensitiveWord) GoString() string {
	return s.String()
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWordsMatchedSensitiveWord) GetWord() *string {
	return s.Word
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWordsMatchedSensitiveWord) SetWord(v string) *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWordsMatchedSensitiveWord {
	s.Word = &v
	return s
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWordsMatchedSensitiveWord) Validate() error {
	return dara.Validate(s)
}
