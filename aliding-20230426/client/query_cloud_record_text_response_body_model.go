// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCloudRecordTextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHasMore(v bool) *QueryCloudRecordTextResponseBody
	GetHasMore() *bool
	SetParagraphList(v []*QueryCloudRecordTextResponseBodyParagraphList) *QueryCloudRecordTextResponseBody
	GetParagraphList() []*QueryCloudRecordTextResponseBodyParagraphList
	SetRequestId(v string) *QueryCloudRecordTextResponseBody
	GetRequestId() *string
}

type QueryCloudRecordTextResponseBody struct {
	// example:
	//
	// true
	HasMore       *bool                                            `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	ParagraphList []*QueryCloudRecordTextResponseBodyParagraphList `json:"paragraphList,omitempty" xml:"paragraphList,omitempty" type:"Repeated"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryCloudRecordTextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCloudRecordTextResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordTextResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *QueryCloudRecordTextResponseBody) GetParagraphList() []*QueryCloudRecordTextResponseBodyParagraphList {
	return s.ParagraphList
}

func (s *QueryCloudRecordTextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCloudRecordTextResponseBody) SetHasMore(v bool) *QueryCloudRecordTextResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryCloudRecordTextResponseBody) SetParagraphList(v []*QueryCloudRecordTextResponseBodyParagraphList) *QueryCloudRecordTextResponseBody {
	s.ParagraphList = v
	return s
}

func (s *QueryCloudRecordTextResponseBody) SetRequestId(v string) *QueryCloudRecordTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCloudRecordTextResponseBody) Validate() error {
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

type QueryCloudRecordTextResponseBodyParagraphList struct {
	// example:
	//
	// 7940
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1631172045153000
	NextTtoken *int64 `json:"NextTtoken,omitempty" xml:"NextTtoken,omitempty"`
	// example:
	//
	// 小钉
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// example:
	//
	// 嘿！你好，这里是小钉
	Paragraph *string `json:"Paragraph,omitempty" xml:"Paragraph,omitempty"`
	// example:
	//
	// 44444
	RecordId     *int64                                                       `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	SentenceList []*QueryCloudRecordTextResponseBodyParagraphListSentenceList `json:"SentenceList,omitempty" xml:"SentenceList,omitempty" type:"Repeated"`
	// example:
	//
	// 7940
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryCloudRecordTextResponseBodyParagraphList) String() string {
	return dara.Prettify(s)
}

func (s QueryCloudRecordTextResponseBodyParagraphList) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) GetNextTtoken() *int64 {
	return s.NextTtoken
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) GetNickName() *string {
	return s.NickName
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) GetParagraph() *string {
	return s.Paragraph
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) GetRecordId() *int64 {
	return s.RecordId
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) GetSentenceList() []*QueryCloudRecordTextResponseBodyParagraphListSentenceList {
	return s.SentenceList
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) GetStatus() *int64 {
	return s.Status
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) GetUserId() *string {
	return s.UserId
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) SetEndTime(v int64) *QueryCloudRecordTextResponseBodyParagraphList {
	s.EndTime = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) SetNextTtoken(v int64) *QueryCloudRecordTextResponseBodyParagraphList {
	s.NextTtoken = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) SetNickName(v string) *QueryCloudRecordTextResponseBodyParagraphList {
	s.NickName = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) SetParagraph(v string) *QueryCloudRecordTextResponseBodyParagraphList {
	s.Paragraph = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) SetRecordId(v int64) *QueryCloudRecordTextResponseBodyParagraphList {
	s.RecordId = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) SetSentenceList(v []*QueryCloudRecordTextResponseBodyParagraphListSentenceList) *QueryCloudRecordTextResponseBodyParagraphList {
	s.SentenceList = v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) SetStartTime(v int64) *QueryCloudRecordTextResponseBodyParagraphList {
	s.StartTime = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) SetStatus(v int64) *QueryCloudRecordTextResponseBodyParagraphList {
	s.Status = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) SetUserId(v string) *QueryCloudRecordTextResponseBodyParagraphList {
	s.UserId = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) Validate() error {
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

type QueryCloudRecordTextResponseBodyParagraphListSentenceList struct {
	// example:
	//
	// 7940
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 这里是小钉
	Sentence *string `json:"Sentence,omitempty" xml:"Sentence,omitempty"`
	// example:
	//
	// 7940
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 012345
	UserId   *string                                                              `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WordList []*QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList `json:"WordList,omitempty" xml:"WordList,omitempty" type:"Repeated"`
}

func (s QueryCloudRecordTextResponseBodyParagraphListSentenceList) String() string {
	return dara.Prettify(s)
}

func (s QueryCloudRecordTextResponseBodyParagraphListSentenceList) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceList) GetSentence() *string {
	return s.Sentence
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceList) GetUserId() *string {
	return s.UserId
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceList) GetWordList() []*QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList {
	return s.WordList
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceList) SetEndTime(v int64) *QueryCloudRecordTextResponseBodyParagraphListSentenceList {
	s.EndTime = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceList) SetSentence(v string) *QueryCloudRecordTextResponseBodyParagraphListSentenceList {
	s.Sentence = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceList) SetStartTime(v int64) *QueryCloudRecordTextResponseBodyParagraphListSentenceList {
	s.StartTime = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceList) SetUserId(v string) *QueryCloudRecordTextResponseBodyParagraphListSentenceList {
	s.UserId = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceList) SetWordList(v []*QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList) *QueryCloudRecordTextResponseBodyParagraphListSentenceList {
	s.WordList = v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceList) Validate() error {
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

type QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList struct {
	// example:
	//
	// 7940
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 7940
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 这里
	Word *string `json:"Word,omitempty" xml:"Word,omitempty"`
	// example:
	//
	// 1631172050535000#0
	WordId *string `json:"WordId,omitempty" xml:"WordId,omitempty"`
}

func (s QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList) String() string {
	return dara.Prettify(s)
}

func (s QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList) GetWord() *string {
	return s.Word
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList) GetWordId() *string {
	return s.WordId
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList) SetEndTime(v int64) *QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList {
	s.EndTime = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList) SetStartTime(v int64) *QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList {
	s.StartTime = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList) SetWord(v string) *QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList {
	s.Word = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList) SetWordId(v string) *QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList {
	s.WordId = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList) Validate() error {
	return dara.Validate(s)
}
