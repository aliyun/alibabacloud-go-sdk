// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFuzzyMatchDomainSensitiveWordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExist(v bool) *FuzzyMatchDomainSensitiveWordResponseBody
	GetExist() *bool
	SetKeyword(v string) *FuzzyMatchDomainSensitiveWordResponseBody
	GetKeyword() *string
	SetMatchedSentiveWords(v *FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWords) *FuzzyMatchDomainSensitiveWordResponseBody
	GetMatchedSentiveWords() *FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWords
	SetRequestId(v string) *FuzzyMatchDomainSensitiveWordResponseBody
	GetRequestId() *string
}

type FuzzyMatchDomainSensitiveWordResponseBody struct {
	// example:
	//
	// true
	Exist *bool `json:"Exist,omitempty" xml:"Exist,omitempty"`
	// example:
	//
	// xxx**.cn
	Keyword             *string                                                       `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	MatchedSentiveWords *FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWords `json:"MatchedSentiveWords,omitempty" xml:"MatchedSentiveWords,omitempty" type:"Struct"`
	// example:
	//
	// D15F91FD-0B34-4E48-8CBF-EFA5D2A31586
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FuzzyMatchDomainSensitiveWordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FuzzyMatchDomainSensitiveWordResponseBody) GoString() string {
	return s.String()
}

func (s *FuzzyMatchDomainSensitiveWordResponseBody) GetExist() *bool {
	return s.Exist
}

func (s *FuzzyMatchDomainSensitiveWordResponseBody) GetKeyword() *string {
	return s.Keyword
}

func (s *FuzzyMatchDomainSensitiveWordResponseBody) GetMatchedSentiveWords() *FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWords {
	return s.MatchedSentiveWords
}

func (s *FuzzyMatchDomainSensitiveWordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FuzzyMatchDomainSensitiveWordResponseBody) SetExist(v bool) *FuzzyMatchDomainSensitiveWordResponseBody {
	s.Exist = &v
	return s
}

func (s *FuzzyMatchDomainSensitiveWordResponseBody) SetKeyword(v string) *FuzzyMatchDomainSensitiveWordResponseBody {
	s.Keyword = &v
	return s
}

func (s *FuzzyMatchDomainSensitiveWordResponseBody) SetMatchedSentiveWords(v *FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWords) *FuzzyMatchDomainSensitiveWordResponseBody {
	s.MatchedSentiveWords = v
	return s
}

func (s *FuzzyMatchDomainSensitiveWordResponseBody) SetRequestId(v string) *FuzzyMatchDomainSensitiveWordResponseBody {
	s.RequestId = &v
	return s
}

func (s *FuzzyMatchDomainSensitiveWordResponseBody) Validate() error {
	if s.MatchedSentiveWords != nil {
		if err := s.MatchedSentiveWords.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWords struct {
	MatchedSensitiveWord []*FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWordsMatchedSensitiveWord `json:"MatchedSensitiveWord,omitempty" xml:"MatchedSensitiveWord,omitempty" type:"Repeated"`
}

func (s FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWords) String() string {
	return dara.Prettify(s)
}

func (s FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWords) GoString() string {
	return s.String()
}

func (s *FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWords) GetMatchedSensitiveWord() []*FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWordsMatchedSensitiveWord {
	return s.MatchedSensitiveWord
}

func (s *FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWords) SetMatchedSensitiveWord(v []*FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWordsMatchedSensitiveWord) *FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWords {
	s.MatchedSensitiveWord = v
	return s
}

func (s *FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWords) Validate() error {
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

type FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWordsMatchedSensitiveWord struct {
	// example:
	//
	// xxx
	Word *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWordsMatchedSensitiveWord) String() string {
	return dara.Prettify(s)
}

func (s FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWordsMatchedSensitiveWord) GoString() string {
	return s.String()
}

func (s *FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWordsMatchedSensitiveWord) GetWord() *string {
	return s.Word
}

func (s *FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWordsMatchedSensitiveWord) SetWord(v string) *FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWordsMatchedSensitiveWord {
	s.Word = &v
	return s
}

func (s *FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWordsMatchedSensitiveWord) Validate() error {
	return dara.Validate(s)
}
