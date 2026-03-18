// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAITeacherExpansionPracticeTaskGenerateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGrade(v string) *AITeacherExpansionPracticeTaskGenerateRequest
	GetGrade() *string
	SetKeySentences(v []*string) *AITeacherExpansionPracticeTaskGenerateRequest
	GetKeySentences() []*string
	SetKeyWords(v []*string) *AITeacherExpansionPracticeTaskGenerateRequest
	GetKeyWords() []*string
	SetLearningObject(v string) *AITeacherExpansionPracticeTaskGenerateRequest
	GetLearningObject() *string
	SetTextContent(v string) *AITeacherExpansionPracticeTaskGenerateRequest
	GetTextContent() *string
	SetTextbook(v string) *AITeacherExpansionPracticeTaskGenerateRequest
	GetTextbook() *string
	SetTopic(v string) *AITeacherExpansionPracticeTaskGenerateRequest
	GetTopic() *string
	SetUserId(v string) *AITeacherExpansionPracticeTaskGenerateRequest
	GetUserId() *string
}

type AITeacherExpansionPracticeTaskGenerateRequest struct {
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

func (s AITeacherExpansionPracticeTaskGenerateRequest) String() string {
	return dara.Prettify(s)
}

func (s AITeacherExpansionPracticeTaskGenerateRequest) GoString() string {
	return s.String()
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) GetGrade() *string {
	return s.Grade
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) GetKeySentences() []*string {
	return s.KeySentences
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) GetKeyWords() []*string {
	return s.KeyWords
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) GetLearningObject() *string {
	return s.LearningObject
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) GetTextContent() *string {
	return s.TextContent
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) GetTextbook() *string {
	return s.Textbook
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) GetTopic() *string {
	return s.Topic
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) GetUserId() *string {
	return s.UserId
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) SetGrade(v string) *AITeacherExpansionPracticeTaskGenerateRequest {
	s.Grade = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) SetKeySentences(v []*string) *AITeacherExpansionPracticeTaskGenerateRequest {
	s.KeySentences = v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) SetKeyWords(v []*string) *AITeacherExpansionPracticeTaskGenerateRequest {
	s.KeyWords = v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) SetLearningObject(v string) *AITeacherExpansionPracticeTaskGenerateRequest {
	s.LearningObject = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) SetTextContent(v string) *AITeacherExpansionPracticeTaskGenerateRequest {
	s.TextContent = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) SetTextbook(v string) *AITeacherExpansionPracticeTaskGenerateRequest {
	s.Textbook = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) SetTopic(v string) *AITeacherExpansionPracticeTaskGenerateRequest {
	s.Topic = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) SetUserId(v string) *AITeacherExpansionPracticeTaskGenerateRequest {
	s.UserId = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) Validate() error {
	return dara.Validate(s)
}
