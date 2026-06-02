// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMeetingFlashMinutesTextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHasNext(v bool) *MeetingFlashMinutesTextResponseBody
	GetHasNext() *bool
	SetNextToken(v string) *MeetingFlashMinutesTextResponseBody
	GetNextToken() *string
	SetParagraphList(v []*MeetingFlashMinutesTextResponseBodyParagraphList) *MeetingFlashMinutesTextResponseBody
	GetParagraphList() []*MeetingFlashMinutesTextResponseBodyParagraphList
	SetRequestId(v string) *MeetingFlashMinutesTextResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *MeetingFlashMinutesTextResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *MeetingFlashMinutesTextResponseBody
	GetVendorType() *string
}

type MeetingFlashMinutesTextResponseBody struct {
	HasNext *bool `json:"hasNext,omitempty" xml:"hasNext,omitempty"`
	// example:
	//
	// 1778490366045000_62XXX
	NextToken     *string                                             `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ParagraphList []*MeetingFlashMinutesTextResponseBodyParagraphList `json:"paragraphList,omitempty" xml:"paragraphList,omitempty" type:"Repeated"`
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

func (s MeetingFlashMinutesTextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MeetingFlashMinutesTextResponseBody) GoString() string {
	return s.String()
}

func (s *MeetingFlashMinutesTextResponseBody) GetHasNext() *bool {
	return s.HasNext
}

func (s *MeetingFlashMinutesTextResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *MeetingFlashMinutesTextResponseBody) GetParagraphList() []*MeetingFlashMinutesTextResponseBodyParagraphList {
	return s.ParagraphList
}

func (s *MeetingFlashMinutesTextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MeetingFlashMinutesTextResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *MeetingFlashMinutesTextResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *MeetingFlashMinutesTextResponseBody) SetHasNext(v bool) *MeetingFlashMinutesTextResponseBody {
	s.HasNext = &v
	return s
}

func (s *MeetingFlashMinutesTextResponseBody) SetNextToken(v string) *MeetingFlashMinutesTextResponseBody {
	s.NextToken = &v
	return s
}

func (s *MeetingFlashMinutesTextResponseBody) SetParagraphList(v []*MeetingFlashMinutesTextResponseBodyParagraphList) *MeetingFlashMinutesTextResponseBody {
	s.ParagraphList = v
	return s
}

func (s *MeetingFlashMinutesTextResponseBody) SetRequestId(v string) *MeetingFlashMinutesTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *MeetingFlashMinutesTextResponseBody) SetVendorRequestId(v string) *MeetingFlashMinutesTextResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *MeetingFlashMinutesTextResponseBody) SetVendorType(v string) *MeetingFlashMinutesTextResponseBody {
	s.VendorType = &v
	return s
}

func (s *MeetingFlashMinutesTextResponseBody) Validate() error {
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

type MeetingFlashMinutesTextResponseBodyParagraphList struct {
	// example:
	//
	// 7920000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 小钉
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// example:
	//
	// 这里是小钉
	Paragraph      *string                                                         `json:"paragraph,omitempty" xml:"paragraph,omitempty"`
	SentenceList   []*MeetingFlashMinutesTextResponseBodyParagraphListSentenceList `json:"sentenceList,omitempty" xml:"sentenceList,omitempty" type:"Repeated"`
	SpeakerDisplay *MeetingFlashMinutesTextResponseBodyParagraphListSpeakerDisplay `json:"speakerDisplay,omitempty" xml:"speakerDisplay,omitempty" type:"Struct"`
	// example:
	//
	// 7910000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 012345
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s MeetingFlashMinutesTextResponseBodyParagraphList) String() string {
	return dara.Prettify(s)
}

func (s MeetingFlashMinutesTextResponseBodyParagraphList) GoString() string {
	return s.String()
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphList) GetNickName() *string {
	return s.NickName
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphList) GetParagraph() *string {
	return s.Paragraph
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphList) GetSentenceList() []*MeetingFlashMinutesTextResponseBodyParagraphListSentenceList {
	return s.SentenceList
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphList) GetSpeakerDisplay() *MeetingFlashMinutesTextResponseBodyParagraphListSpeakerDisplay {
	return s.SpeakerDisplay
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphList) GetUserId() *string {
	return s.UserId
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphList) SetEndTime(v int64) *MeetingFlashMinutesTextResponseBodyParagraphList {
	s.EndTime = &v
	return s
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphList) SetNickName(v string) *MeetingFlashMinutesTextResponseBodyParagraphList {
	s.NickName = &v
	return s
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphList) SetParagraph(v string) *MeetingFlashMinutesTextResponseBodyParagraphList {
	s.Paragraph = &v
	return s
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphList) SetSentenceList(v []*MeetingFlashMinutesTextResponseBodyParagraphListSentenceList) *MeetingFlashMinutesTextResponseBodyParagraphList {
	s.SentenceList = v
	return s
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphList) SetSpeakerDisplay(v *MeetingFlashMinutesTextResponseBodyParagraphListSpeakerDisplay) *MeetingFlashMinutesTextResponseBodyParagraphList {
	s.SpeakerDisplay = v
	return s
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphList) SetStartTime(v int64) *MeetingFlashMinutesTextResponseBodyParagraphList {
	s.StartTime = &v
	return s
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphList) SetUserId(v string) *MeetingFlashMinutesTextResponseBodyParagraphList {
	s.UserId = &v
	return s
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphList) Validate() error {
	if s.SentenceList != nil {
		for _, item := range s.SentenceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SpeakerDisplay != nil {
		if err := s.SpeakerDisplay.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MeetingFlashMinutesTextResponseBodyParagraphListSentenceList struct {
	// example:
	//
	// 7920000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 这里是小钉
	Sentence *string `json:"sentence,omitempty" xml:"sentence,omitempty"`
	// example:
	//
	// 7910000
	StartTime *int64                                                                  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	WordList  []*MeetingFlashMinutesTextResponseBodyParagraphListSentenceListWordList `json:"wordList,omitempty" xml:"wordList,omitempty" type:"Repeated"`
}

func (s MeetingFlashMinutesTextResponseBodyParagraphListSentenceList) String() string {
	return dara.Prettify(s)
}

func (s MeetingFlashMinutesTextResponseBodyParagraphListSentenceList) GoString() string {
	return s.String()
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphListSentenceList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphListSentenceList) GetSentence() *string {
	return s.Sentence
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphListSentenceList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphListSentenceList) GetWordList() []*MeetingFlashMinutesTextResponseBodyParagraphListSentenceListWordList {
	return s.WordList
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphListSentenceList) SetEndTime(v int64) *MeetingFlashMinutesTextResponseBodyParagraphListSentenceList {
	s.EndTime = &v
	return s
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphListSentenceList) SetSentence(v string) *MeetingFlashMinutesTextResponseBodyParagraphListSentenceList {
	s.Sentence = &v
	return s
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphListSentenceList) SetStartTime(v int64) *MeetingFlashMinutesTextResponseBodyParagraphListSentenceList {
	s.StartTime = &v
	return s
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphListSentenceList) SetWordList(v []*MeetingFlashMinutesTextResponseBodyParagraphListSentenceListWordList) *MeetingFlashMinutesTextResponseBodyParagraphListSentenceList {
	s.WordList = v
	return s
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphListSentenceList) Validate() error {
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

type MeetingFlashMinutesTextResponseBodyParagraphListSentenceListWordList struct {
	// example:
	//
	// 7920000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 7910000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 单词
	Word *string `json:"word,omitempty" xml:"word,omitempty"`
}

func (s MeetingFlashMinutesTextResponseBodyParagraphListSentenceListWordList) String() string {
	return dara.Prettify(s)
}

func (s MeetingFlashMinutesTextResponseBodyParagraphListSentenceListWordList) GoString() string {
	return s.String()
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphListSentenceListWordList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphListSentenceListWordList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphListSentenceListWordList) GetWord() *string {
	return s.Word
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphListSentenceListWordList) SetEndTime(v int64) *MeetingFlashMinutesTextResponseBodyParagraphListSentenceListWordList {
	s.EndTime = &v
	return s
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphListSentenceListWordList) SetStartTime(v int64) *MeetingFlashMinutesTextResponseBodyParagraphListSentenceListWordList {
	s.StartTime = &v
	return s
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphListSentenceListWordList) SetWord(v string) *MeetingFlashMinutesTextResponseBodyParagraphListSentenceListWordList {
	s.Word = &v
	return s
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphListSentenceListWordList) Validate() error {
	return dara.Validate(s)
}

type MeetingFlashMinutesTextResponseBodyParagraphListSpeakerDisplay struct {
	// example:
	//
	// https://xxx
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// example:
	//
	// 小钉
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
}

func (s MeetingFlashMinutesTextResponseBodyParagraphListSpeakerDisplay) String() string {
	return dara.Prettify(s)
}

func (s MeetingFlashMinutesTextResponseBodyParagraphListSpeakerDisplay) GoString() string {
	return s.String()
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphListSpeakerDisplay) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphListSpeakerDisplay) GetNickName() *string {
	return s.NickName
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphListSpeakerDisplay) SetAvatarUrl(v string) *MeetingFlashMinutesTextResponseBodyParagraphListSpeakerDisplay {
	s.AvatarUrl = &v
	return s
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphListSpeakerDisplay) SetNickName(v string) *MeetingFlashMinutesTextResponseBodyParagraphListSpeakerDisplay {
	s.NickName = &v
	return s
}

func (s *MeetingFlashMinutesTextResponseBodyParagraphListSpeakerDisplay) Validate() error {
	return dara.Validate(s)
}
