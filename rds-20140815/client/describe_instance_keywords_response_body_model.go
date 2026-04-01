// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceKeywordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *DescribeInstanceKeywordsResponseBody
	GetKey() *string
	SetRequestId(v string) *DescribeInstanceKeywordsResponseBody
	GetRequestId() *string
	SetWords(v *DescribeInstanceKeywordsResponseBodyWords) *DescribeInstanceKeywordsResponseBody
	GetWords() *DescribeInstanceKeywordsResponseBodyWords
}

type DescribeInstanceKeywordsResponseBody struct {
	// The type of reserved keyword returned.
	//
	// example:
	//
	// account
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1E43AAE0-BEE8-43DA-860D-EAF2AA0724DC
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Words     *DescribeInstanceKeywordsResponseBodyWords `json:"Words,omitempty" xml:"Words,omitempty" type:"Struct"`
}

func (s DescribeInstanceKeywordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceKeywordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceKeywordsResponseBody) GetKey() *string {
	return s.Key
}

func (s *DescribeInstanceKeywordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceKeywordsResponseBody) GetWords() *DescribeInstanceKeywordsResponseBodyWords {
	return s.Words
}

func (s *DescribeInstanceKeywordsResponseBody) SetKey(v string) *DescribeInstanceKeywordsResponseBody {
	s.Key = &v
	return s
}

func (s *DescribeInstanceKeywordsResponseBody) SetRequestId(v string) *DescribeInstanceKeywordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceKeywordsResponseBody) SetWords(v *DescribeInstanceKeywordsResponseBodyWords) *DescribeInstanceKeywordsResponseBody {
	s.Words = v
	return s
}

func (s *DescribeInstanceKeywordsResponseBody) Validate() error {
	if s.Words != nil {
		if err := s.Words.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceKeywordsResponseBodyWords struct {
	Word []*string `json:"word,omitempty" xml:"word,omitempty" type:"Repeated"`
}

func (s DescribeInstanceKeywordsResponseBodyWords) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceKeywordsResponseBodyWords) GoString() string {
	return s.String()
}

func (s *DescribeInstanceKeywordsResponseBodyWords) GetWord() []*string {
	return s.Word
}

func (s *DescribeInstanceKeywordsResponseBodyWords) SetWord(v []*string) *DescribeInstanceKeywordsResponseBodyWords {
	s.Word = v
	return s
}

func (s *DescribeInstanceKeywordsResponseBodyWords) Validate() error {
	return dara.Validate(s)
}
