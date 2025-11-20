// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMinutesTextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHasMore(v bool) *QueryMinutesTextResponseBody
	GetHasMore() *bool
	SetNextToken(v string) *QueryMinutesTextResponseBody
	GetNextToken() *string
	SetParagraphList(v []*QueryMinutesTextResponseBodyParagraphList) *QueryMinutesTextResponseBody
	GetParagraphList() []*QueryMinutesTextResponseBodyParagraphList
	SetRequestId(v string) *QueryMinutesTextResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *QueryMinutesTextResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *QueryMinutesTextResponseBody
	GetVendorType() *string
}

type QueryMinutesTextResponseBody struct {
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// 0
	NextToken     *string                                      `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ParagraphList []*QueryMinutesTextResponseBodyParagraphList `json:"paragraphList,omitempty" xml:"paragraphList,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s QueryMinutesTextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesTextResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMinutesTextResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *QueryMinutesTextResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryMinutesTextResponseBody) GetParagraphList() []*QueryMinutesTextResponseBodyParagraphList {
	return s.ParagraphList
}

func (s *QueryMinutesTextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMinutesTextResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *QueryMinutesTextResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *QueryMinutesTextResponseBody) SetHasMore(v bool) *QueryMinutesTextResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryMinutesTextResponseBody) SetNextToken(v string) *QueryMinutesTextResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryMinutesTextResponseBody) SetParagraphList(v []*QueryMinutesTextResponseBodyParagraphList) *QueryMinutesTextResponseBody {
	s.ParagraphList = v
	return s
}

func (s *QueryMinutesTextResponseBody) SetRequestId(v string) *QueryMinutesTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMinutesTextResponseBody) SetVendorRequestId(v string) *QueryMinutesTextResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *QueryMinutesTextResponseBody) SetVendorType(v string) *QueryMinutesTextResponseBody {
	s.VendorType = &v
	return s
}

func (s *QueryMinutesTextResponseBody) Validate() error {
	if s.ParagraphList != nil {
		for _, item := range s.ParagraphList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMinutesTextResponseBodyParagraphList struct {
	// example:
	//
	// 7910000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 小钉
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// example:
	//
	// 123
	Paragraph *string `json:"Paragraph,omitempty" xml:"Paragraph,omitempty"`
	// example:
	//
	// 123
	ParagraphId *int64 `json:"ParagraphId,omitempty" xml:"ParagraphId,omitempty"`
	// example:
	//
	// 44444
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// example:
	//
	// []
	SentenceList []*QueryMinutesTextResponseBodyParagraphListSentenceList `json:"SentenceList,omitempty" xml:"SentenceList,omitempty" type:"Repeated"`
	// example:
	//
	// 7910000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryMinutesTextResponseBodyParagraphList) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesTextResponseBodyParagraphList) GoString() string {
	return s.String()
}

func (s *QueryMinutesTextResponseBodyParagraphList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryMinutesTextResponseBodyParagraphList) GetNickName() *string {
	return s.NickName
}

func (s *QueryMinutesTextResponseBodyParagraphList) GetParagraph() *string {
	return s.Paragraph
}

func (s *QueryMinutesTextResponseBodyParagraphList) GetParagraphId() *int64 {
	return s.ParagraphId
}

func (s *QueryMinutesTextResponseBodyParagraphList) GetRecordId() *int64 {
	return s.RecordId
}

func (s *QueryMinutesTextResponseBodyParagraphList) GetSentenceList() []*QueryMinutesTextResponseBodyParagraphListSentenceList {
	return s.SentenceList
}

func (s *QueryMinutesTextResponseBodyParagraphList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryMinutesTextResponseBodyParagraphList) GetUserId() *string {
	return s.UserId
}

func (s *QueryMinutesTextResponseBodyParagraphList) SetEndTime(v int64) *QueryMinutesTextResponseBodyParagraphList {
	s.EndTime = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphList) SetNickName(v string) *QueryMinutesTextResponseBodyParagraphList {
	s.NickName = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphList) SetParagraph(v string) *QueryMinutesTextResponseBodyParagraphList {
	s.Paragraph = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphList) SetParagraphId(v int64) *QueryMinutesTextResponseBodyParagraphList {
	s.ParagraphId = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphList) SetRecordId(v int64) *QueryMinutesTextResponseBodyParagraphList {
	s.RecordId = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphList) SetSentenceList(v []*QueryMinutesTextResponseBodyParagraphListSentenceList) *QueryMinutesTextResponseBodyParagraphList {
	s.SentenceList = v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphList) SetStartTime(v int64) *QueryMinutesTextResponseBodyParagraphList {
	s.StartTime = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphList) SetUserId(v string) *QueryMinutesTextResponseBodyParagraphList {
	s.UserId = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphList) Validate() error {
	if s.SentenceList != nil {
		for _, item := range s.SentenceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMinutesTextResponseBodyParagraphListSentenceList struct {
	// example:
	//
	// 7910000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 这里是小钉
	Sentence *string `json:"Sentence,omitempty" xml:"Sentence,omitempty"`
	// example:
	//
	// 7910000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// []
	WordList []*QueryMinutesTextResponseBodyParagraphListSentenceListWordList `json:"WordList,omitempty" xml:"WordList,omitempty" type:"Repeated"`
}

func (s QueryMinutesTextResponseBodyParagraphListSentenceList) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesTextResponseBodyParagraphListSentenceList) GoString() string {
	return s.String()
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceList) GetSentence() *string {
	return s.Sentence
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceList) GetUserId() *string {
	return s.UserId
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceList) GetWordList() []*QueryMinutesTextResponseBodyParagraphListSentenceListWordList {
	return s.WordList
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceList) SetEndTime(v int64) *QueryMinutesTextResponseBodyParagraphListSentenceList {
	s.EndTime = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceList) SetSentence(v string) *QueryMinutesTextResponseBodyParagraphListSentenceList {
	s.Sentence = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceList) SetStartTime(v int64) *QueryMinutesTextResponseBodyParagraphListSentenceList {
	s.StartTime = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceList) SetUserId(v string) *QueryMinutesTextResponseBodyParagraphListSentenceList {
	s.UserId = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceList) SetWordList(v []*QueryMinutesTextResponseBodyParagraphListSentenceListWordList) *QueryMinutesTextResponseBodyParagraphListSentenceList {
	s.WordList = v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceList) Validate() error {
	if s.WordList != nil {
		for _, item := range s.WordList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMinutesTextResponseBodyParagraphListSentenceListWordList struct {
	// example:
	//
	// 7910000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 7910000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 单词
	Word *string `json:"Word,omitempty" xml:"Word,omitempty"`
	// example:
	//
	// 172
	WordId *string `json:"WordId,omitempty" xml:"WordId,omitempty"`
}

func (s QueryMinutesTextResponseBodyParagraphListSentenceListWordList) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesTextResponseBodyParagraphListSentenceListWordList) GoString() string {
	return s.String()
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceListWordList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceListWordList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceListWordList) GetWord() *string {
	return s.Word
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceListWordList) GetWordId() *string {
	return s.WordId
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceListWordList) SetEndTime(v int64) *QueryMinutesTextResponseBodyParagraphListSentenceListWordList {
	s.EndTime = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceListWordList) SetStartTime(v int64) *QueryMinutesTextResponseBodyParagraphListSentenceListWordList {
	s.StartTime = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceListWordList) SetWord(v string) *QueryMinutesTextResponseBodyParagraphListSentenceListWordList {
	s.Word = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceListWordList) SetWordId(v string) *QueryMinutesTextResponseBodyParagraphListSentenceListWordList {
	s.WordId = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceListWordList) Validate() error {
	return dara.Validate(s)
}
