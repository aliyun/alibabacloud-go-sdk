// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAITeacherSyncPracticeTaskGenerateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGrade(v string) *AITeacherSyncPracticeTaskGenerateRequest
	GetGrade() *string
	SetKeySentences(v []*string) *AITeacherSyncPracticeTaskGenerateRequest
	GetKeySentences() []*string
	SetKeyWords(v []*string) *AITeacherSyncPracticeTaskGenerateRequest
	GetKeyWords() []*string
	SetLearningObject(v string) *AITeacherSyncPracticeTaskGenerateRequest
	GetLearningObject() *string
	SetTextContent(v string) *AITeacherSyncPracticeTaskGenerateRequest
	GetTextContent() *string
	SetTextbook(v string) *AITeacherSyncPracticeTaskGenerateRequest
	GetTextbook() *string
	SetTopic(v string) *AITeacherSyncPracticeTaskGenerateRequest
	GetTopic() *string
	SetUserId(v string) *AITeacherSyncPracticeTaskGenerateRequest
	GetUserId() *string
}

type AITeacherSyncPracticeTaskGenerateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 13
	Grade        *string   `json:"grade,omitempty" xml:"grade,omitempty"`
	KeySentences []*string `json:"keySentences,omitempty" xml:"keySentences,omitempty" type:"Repeated"`
	KeyWords     []*string `json:"keyWords,omitempty" xml:"keyWords,omitempty" type:"Repeated"`
	// example:
	//
	// Understanding unique professions such as dog walkers, hotel test sleepers, and food tasters, including their job responsibilities and the benefits or challenges associated with each role.
	LearningObject *string `json:"learningObject,omitempty" xml:"learningObject,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Dog walker Dog walking, as a profession, originated in the US. Some may think that it\\"s a perfect job, because dog walkers won\\"t be imprisoned in an office. But it\\"s actually manual labour. At their busiest, dog walkers may have more than ten dogs to take care of in a day. Hotel test sleeper A hotel test sleeper, as the name suggests, has to write expert reviews about the facilities, locations, prices, dining and other services of hotels, in order to provide evaluations and guides for travelers. Hotel test sleepers don\\"t need to punch in for work and they get about ten thousand yuan as income every month. What a comfortable job! Food taster In ancient times, a food taster was a person who tasted foods (or drinks) to be served to someone else, to confirm that it was safe to eat. But now, those working as food tasters just get to taste various new foods and drinks aimed at specific regions across the world. They then give their opinions on these products to the companies and suggest improvements.
	TextContent *string `json:"textContent,omitempty" xml:"textContent,omitempty"`
	Textbook    *string `json:"textbook,omitempty" xml:"textbook,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// talk about your dream job.
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6440xxxxxxxxxx5fafc98c421
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s AITeacherSyncPracticeTaskGenerateRequest) String() string {
	return dara.Prettify(s)
}

func (s AITeacherSyncPracticeTaskGenerateRequest) GoString() string {
	return s.String()
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) GetGrade() *string {
	return s.Grade
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) GetKeySentences() []*string {
	return s.KeySentences
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) GetKeyWords() []*string {
	return s.KeyWords
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) GetLearningObject() *string {
	return s.LearningObject
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) GetTextContent() *string {
	return s.TextContent
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) GetTextbook() *string {
	return s.Textbook
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) GetTopic() *string {
	return s.Topic
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) GetUserId() *string {
	return s.UserId
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) SetGrade(v string) *AITeacherSyncPracticeTaskGenerateRequest {
	s.Grade = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) SetKeySentences(v []*string) *AITeacherSyncPracticeTaskGenerateRequest {
	s.KeySentences = v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) SetKeyWords(v []*string) *AITeacherSyncPracticeTaskGenerateRequest {
	s.KeyWords = v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) SetLearningObject(v string) *AITeacherSyncPracticeTaskGenerateRequest {
	s.LearningObject = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) SetTextContent(v string) *AITeacherSyncPracticeTaskGenerateRequest {
	s.TextContent = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) SetTextbook(v string) *AITeacherSyncPracticeTaskGenerateRequest {
	s.Textbook = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) SetTopic(v string) *AITeacherSyncPracticeTaskGenerateRequest {
	s.Topic = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) SetUserId(v string) *AITeacherSyncPracticeTaskGenerateRequest {
	s.UserId = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) Validate() error {
	return dara.Validate(s)
}
