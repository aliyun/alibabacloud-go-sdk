// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTextbookAssistantArticleDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListTextbookAssistantArticleDetailsResponseBodyData) *ListTextbookAssistantArticleDetailsResponseBody
	GetData() []*ListTextbookAssistantArticleDetailsResponseBodyData
	SetErrCode(v string) *ListTextbookAssistantArticleDetailsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListTextbookAssistantArticleDetailsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ListTextbookAssistantArticleDetailsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListTextbookAssistantArticleDetailsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTextbookAssistantArticleDetailsResponseBody
	GetSuccess() *bool
}

type ListTextbookAssistantArticleDetailsResponseBody struct {
	Data []*ListTextbookAssistantArticleDetailsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode    *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListTextbookAssistantArticleDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantArticleDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantArticleDetailsResponseBody) GetData() []*ListTextbookAssistantArticleDetailsResponseBodyData {
	return s.Data
}

func (s *ListTextbookAssistantArticleDetailsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListTextbookAssistantArticleDetailsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListTextbookAssistantArticleDetailsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTextbookAssistantArticleDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTextbookAssistantArticleDetailsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTextbookAssistantArticleDetailsResponseBody) SetData(v []*ListTextbookAssistantArticleDetailsResponseBodyData) *ListTextbookAssistantArticleDetailsResponseBody {
	s.Data = v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBody) SetErrCode(v string) *ListTextbookAssistantArticleDetailsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBody) SetErrMessage(v string) *ListTextbookAssistantArticleDetailsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBody) SetHttpStatusCode(v int32) *ListTextbookAssistantArticleDetailsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBody) SetRequestId(v string) *ListTextbookAssistantArticleDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBody) SetSuccess(v bool) *ListTextbookAssistantArticleDetailsResponseBody {
	s.Success = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTextbookAssistantArticleDetailsResponseBodyData struct {
	// example:
	//
	// 0c05700d4d9411efbe6e0c42a106bb02
	ArticleId    *string                                                            `json:"articleId,omitempty" xml:"articleId,omitempty"`
	QuestionList []*ListTextbookAssistantArticleDetailsResponseBodyDataQuestionList `json:"questionList,omitempty" xml:"questionList,omitempty" type:"Repeated"`
	SceneList    []*ListTextbookAssistantArticleDetailsResponseBodyDataSceneList    `json:"sceneList,omitempty" xml:"sceneList,omitempty" type:"Repeated"`
	SentenceList []*ListTextbookAssistantArticleDetailsResponseBodyDataSentenceList `json:"sentenceList,omitempty" xml:"sentenceList,omitempty" type:"Repeated"`
	Target       *string                                                            `json:"target,omitempty" xml:"target,omitempty"`
	Theme        *ListTextbookAssistantArticleDetailsResponseBodyDataTheme          `json:"theme,omitempty" xml:"theme,omitempty" type:"Struct"`
	Topic        *ListTextbookAssistantArticleDetailsResponseBodyDataTopic          `json:"topic,omitempty" xml:"topic,omitempty" type:"Struct"`
	WordList     []*ListTextbookAssistantArticleDetailsResponseBodyDataWordList     `json:"wordList,omitempty" xml:"wordList,omitempty" type:"Repeated"`
}

func (s ListTextbookAssistantArticleDetailsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantArticleDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyData) GetArticleId() *string {
	return s.ArticleId
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyData) GetQuestionList() []*ListTextbookAssistantArticleDetailsResponseBodyDataQuestionList {
	return s.QuestionList
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyData) GetSceneList() []*ListTextbookAssistantArticleDetailsResponseBodyDataSceneList {
	return s.SceneList
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyData) GetSentenceList() []*ListTextbookAssistantArticleDetailsResponseBodyDataSentenceList {
	return s.SentenceList
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyData) GetTarget() *string {
	return s.Target
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyData) GetTheme() *ListTextbookAssistantArticleDetailsResponseBodyDataTheme {
	return s.Theme
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyData) GetTopic() *ListTextbookAssistantArticleDetailsResponseBodyDataTopic {
	return s.Topic
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyData) GetWordList() []*ListTextbookAssistantArticleDetailsResponseBodyDataWordList {
	return s.WordList
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyData) SetArticleId(v string) *ListTextbookAssistantArticleDetailsResponseBodyData {
	s.ArticleId = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyData) SetQuestionList(v []*ListTextbookAssistantArticleDetailsResponseBodyDataQuestionList) *ListTextbookAssistantArticleDetailsResponseBodyData {
	s.QuestionList = v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyData) SetSceneList(v []*ListTextbookAssistantArticleDetailsResponseBodyDataSceneList) *ListTextbookAssistantArticleDetailsResponseBodyData {
	s.SceneList = v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyData) SetSentenceList(v []*ListTextbookAssistantArticleDetailsResponseBodyDataSentenceList) *ListTextbookAssistantArticleDetailsResponseBodyData {
	s.SentenceList = v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyData) SetTarget(v string) *ListTextbookAssistantArticleDetailsResponseBodyData {
	s.Target = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyData) SetTheme(v *ListTextbookAssistantArticleDetailsResponseBodyDataTheme) *ListTextbookAssistantArticleDetailsResponseBodyData {
	s.Theme = v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyData) SetTopic(v *ListTextbookAssistantArticleDetailsResponseBodyDataTopic) *ListTextbookAssistantArticleDetailsResponseBodyData {
	s.Topic = v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyData) SetWordList(v []*ListTextbookAssistantArticleDetailsResponseBodyDataWordList) *ListTextbookAssistantArticleDetailsResponseBodyData {
	s.WordList = v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyData) Validate() error {
	if s.QuestionList != nil {
		for _, item := range s.QuestionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SceneList != nil {
		for _, item := range s.SceneList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SentenceList != nil {
		for _, item := range s.SentenceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Theme != nil {
		if err := s.Theme.Validate(); err != nil {
			return err
		}
	}
	if s.Topic != nil {
		if err := s.Topic.Validate(); err != nil {
			return err
		}
	}
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

type ListTextbookAssistantArticleDetailsResponseBodyDataQuestionList struct {
	// example:
	//
	// I\\"m Mike Black
	Answer *string `json:"answer,omitempty" xml:"answer,omitempty"`
	// example:
	//
	// From the book, how does Mike Black introduce himself?
	Question          *string `json:"question,omitempty" xml:"question,omitempty"`
	QuestionTranslate *string `json:"questionTranslate,omitempty" xml:"questionTranslate,omitempty"`
}

func (s ListTextbookAssistantArticleDetailsResponseBodyDataQuestionList) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantArticleDetailsResponseBodyDataQuestionList) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataQuestionList) GetAnswer() *string {
	return s.Answer
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataQuestionList) GetQuestion() *string {
	return s.Question
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataQuestionList) GetQuestionTranslate() *string {
	return s.QuestionTranslate
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataQuestionList) SetAnswer(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataQuestionList {
	s.Answer = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataQuestionList) SetQuestion(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataQuestionList {
	s.Question = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataQuestionList) SetQuestionTranslate(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataQuestionList {
	s.QuestionTranslate = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataQuestionList) Validate() error {
	return dara.Validate(s)
}

type ListTextbookAssistantArticleDetailsResponseBodyDataSceneList struct {
	// example:
	//
	// In the park, you introduce yourself to John and ask his name.
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// example:
	//
	// 38cddd70509911efbe6e0c42a106bb02
	SceneId        *string   `json:"sceneId,omitempty" xml:"sceneId,omitempty"`
	SceneImageList []*string `json:"sceneImageList,omitempty" xml:"sceneImageList,omitempty" type:"Repeated"`
	SceneTranslate *string   `json:"sceneTranslate,omitempty" xml:"sceneTranslate,omitempty"`
}

func (s ListTextbookAssistantArticleDetailsResponseBodyDataSceneList) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantArticleDetailsResponseBodyDataSceneList) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataSceneList) GetScene() *string {
	return s.Scene
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataSceneList) GetSceneId() *string {
	return s.SceneId
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataSceneList) GetSceneImageList() []*string {
	return s.SceneImageList
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataSceneList) GetSceneTranslate() *string {
	return s.SceneTranslate
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataSceneList) SetScene(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataSceneList {
	s.Scene = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataSceneList) SetSceneId(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataSceneList {
	s.SceneId = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataSceneList) SetSceneImageList(v []*string) *ListTextbookAssistantArticleDetailsResponseBodyDataSceneList {
	s.SceneImageList = v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataSceneList) SetSceneTranslate(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataSceneList {
	s.SceneTranslate = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataSceneList) Validate() error {
	return dara.Validate(s)
}

type ListTextbookAssistantArticleDetailsResponseBodyDataSentenceList struct {
	SentenceAnalysis *string `json:"sentenceAnalysis,omitempty" xml:"sentenceAnalysis,omitempty"`
	// example:
	//
	// 4de677d2509811efbe6e0c42a106bb02
	SentenceId *string `json:"sentenceId,omitempty" xml:"sentenceId,omitempty"`
	// example:
	//
	// I\\"m Mike Black
	SentenceText *string `json:"sentenceText,omitempty" xml:"sentenceText,omitempty"`
}

func (s ListTextbookAssistantArticleDetailsResponseBodyDataSentenceList) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantArticleDetailsResponseBodyDataSentenceList) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataSentenceList) GetSentenceAnalysis() *string {
	return s.SentenceAnalysis
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataSentenceList) GetSentenceId() *string {
	return s.SentenceId
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataSentenceList) GetSentenceText() *string {
	return s.SentenceText
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataSentenceList) SetSentenceAnalysis(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataSentenceList {
	s.SentenceAnalysis = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataSentenceList) SetSentenceId(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataSentenceList {
	s.SentenceId = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataSentenceList) SetSentenceText(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataSentenceList {
	s.SentenceText = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataSentenceList) Validate() error {
	return dara.Validate(s)
}

type ListTextbookAssistantArticleDetailsResponseBodyDataTheme struct {
	ThemeImageList []*string `json:"themeImageList,omitempty" xml:"themeImageList,omitempty" type:"Repeated"`
	ThemeName      *string   `json:"themeName,omitempty" xml:"themeName,omitempty"`
	// example:
	//
	// Self-awareness, self-management, self-improvement
	ThemeTranslate *string `json:"themeTranslate,omitempty" xml:"themeTranslate,omitempty"`
}

func (s ListTextbookAssistantArticleDetailsResponseBodyDataTheme) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantArticleDetailsResponseBodyDataTheme) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataTheme) GetThemeImageList() []*string {
	return s.ThemeImageList
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataTheme) GetThemeName() *string {
	return s.ThemeName
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataTheme) GetThemeTranslate() *string {
	return s.ThemeTranslate
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataTheme) SetThemeImageList(v []*string) *ListTextbookAssistantArticleDetailsResponseBodyDataTheme {
	s.ThemeImageList = v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataTheme) SetThemeName(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataTheme {
	s.ThemeName = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataTheme) SetThemeTranslate(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataTheme {
	s.ThemeTranslate = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataTheme) Validate() error {
	return dara.Validate(s)
}

type ListTextbookAssistantArticleDetailsResponseBodyDataTopic struct {
	TopicImageList []*string `json:"topicImageList,omitempty" xml:"topicImageList,omitempty" type:"Repeated"`
	TopicName      *string   `json:"topicName,omitempty" xml:"topicName,omitempty"`
	// example:
	//
	// Greetings and self-introduction
	TopicTranslate *string `json:"topicTranslate,omitempty" xml:"topicTranslate,omitempty"`
}

func (s ListTextbookAssistantArticleDetailsResponseBodyDataTopic) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantArticleDetailsResponseBodyDataTopic) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataTopic) GetTopicImageList() []*string {
	return s.TopicImageList
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataTopic) GetTopicName() *string {
	return s.TopicName
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataTopic) GetTopicTranslate() *string {
	return s.TopicTranslate
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataTopic) SetTopicImageList(v []*string) *ListTextbookAssistantArticleDetailsResponseBodyDataTopic {
	s.TopicImageList = v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataTopic) SetTopicName(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataTopic {
	s.TopicName = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataTopic) SetTopicTranslate(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataTopic {
	s.TopicTranslate = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataTopic) Validate() error {
	return dara.Validate(s)
}

type ListTextbookAssistantArticleDetailsResponseBodyDataWordList struct {
	WordAnalysis *string `json:"wordAnalysis,omitempty" xml:"wordAnalysis,omitempty"`
	// example:
	//
	// a94df134ed8c11eebe6e0c42a106bb02
	WordId *string `json:"wordId,omitempty" xml:"wordId,omitempty"`
	// example:
	//
	// nice
	WordText *string `json:"wordText,omitempty" xml:"wordText,omitempty"`
}

func (s ListTextbookAssistantArticleDetailsResponseBodyDataWordList) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantArticleDetailsResponseBodyDataWordList) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataWordList) GetWordAnalysis() *string {
	return s.WordAnalysis
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataWordList) GetWordId() *string {
	return s.WordId
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataWordList) GetWordText() *string {
	return s.WordText
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataWordList) SetWordAnalysis(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataWordList {
	s.WordAnalysis = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataWordList) SetWordId(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataWordList {
	s.WordId = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataWordList) SetWordText(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataWordList {
	s.WordText = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataWordList) Validate() error {
	return dara.Validate(s)
}
