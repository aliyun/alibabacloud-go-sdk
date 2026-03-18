// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTextbookAssistantSceneDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListTextbookAssistantSceneDetailsResponseBodyData) *ListTextbookAssistantSceneDetailsResponseBody
	GetData() []*ListTextbookAssistantSceneDetailsResponseBodyData
	SetErrCode(v string) *ListTextbookAssistantSceneDetailsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListTextbookAssistantSceneDetailsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ListTextbookAssistantSceneDetailsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListTextbookAssistantSceneDetailsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTextbookAssistantSceneDetailsResponseBody
	GetSuccess() *bool
}

type ListTextbookAssistantSceneDetailsResponseBody struct {
	Data []*ListTextbookAssistantSceneDetailsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s ListTextbookAssistantSceneDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantSceneDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantSceneDetailsResponseBody) GetData() []*ListTextbookAssistantSceneDetailsResponseBodyData {
	return s.Data
}

func (s *ListTextbookAssistantSceneDetailsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListTextbookAssistantSceneDetailsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListTextbookAssistantSceneDetailsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTextbookAssistantSceneDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTextbookAssistantSceneDetailsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTextbookAssistantSceneDetailsResponseBody) SetData(v []*ListTextbookAssistantSceneDetailsResponseBodyData) *ListTextbookAssistantSceneDetailsResponseBody {
	s.Data = v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBody) SetErrCode(v string) *ListTextbookAssistantSceneDetailsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBody) SetErrMessage(v string) *ListTextbookAssistantSceneDetailsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBody) SetHttpStatusCode(v int32) *ListTextbookAssistantSceneDetailsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBody) SetRequestId(v string) *ListTextbookAssistantSceneDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBody) SetSuccess(v bool) *ListTextbookAssistantSceneDetailsResponseBody {
	s.Success = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBody) Validate() error {
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

type ListTextbookAssistantSceneDetailsResponseBodyData struct {
	RoleList []*ListTextbookAssistantSceneDetailsResponseBodyDataRoleList `json:"roleList,omitempty" xml:"roleList,omitempty" type:"Repeated"`
	// example:
	//
	// At school, Carl sees a photo and asks you about your family.
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// example:
	//
	// 38c41b7b509911efbe6e0c42a106bb02
	SceneId        *string                                                           `json:"sceneId,omitempty" xml:"sceneId,omitempty"`
	SceneImageList []*string                                                         `json:"sceneImageList,omitempty" xml:"sceneImageList,omitempty" type:"Repeated"`
	SceneTaskList  []*ListTextbookAssistantSceneDetailsResponseBodyDataSceneTaskList `json:"sceneTaskList,omitempty" xml:"sceneTaskList,omitempty" type:"Repeated"`
	SceneTranslate *string                                                           `json:"sceneTranslate,omitempty" xml:"sceneTranslate,omitempty"`
	SentenceList   []*ListTextbookAssistantSceneDetailsResponseBodyDataSentenceList  `json:"sentenceList,omitempty" xml:"sentenceList,omitempty" type:"Repeated"`
	// example:
	//
	// ""
	Target   *string                                                      `json:"target,omitempty" xml:"target,omitempty"`
	Theme    *ListTextbookAssistantSceneDetailsResponseBodyDataTheme      `json:"theme,omitempty" xml:"theme,omitempty" type:"Struct"`
	Topic    *ListTextbookAssistantSceneDetailsResponseBodyDataTopic      `json:"topic,omitempty" xml:"topic,omitempty" type:"Struct"`
	WordList []*ListTextbookAssistantSceneDetailsResponseBodyDataWordList `json:"wordList,omitempty" xml:"wordList,omitempty" type:"Repeated"`
}

func (s ListTextbookAssistantSceneDetailsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantSceneDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) GetRoleList() []*ListTextbookAssistantSceneDetailsResponseBodyDataRoleList {
	return s.RoleList
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) GetScene() *string {
	return s.Scene
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) GetSceneId() *string {
	return s.SceneId
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) GetSceneImageList() []*string {
	return s.SceneImageList
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) GetSceneTaskList() []*ListTextbookAssistantSceneDetailsResponseBodyDataSceneTaskList {
	return s.SceneTaskList
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) GetSceneTranslate() *string {
	return s.SceneTranslate
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) GetSentenceList() []*ListTextbookAssistantSceneDetailsResponseBodyDataSentenceList {
	return s.SentenceList
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) GetTarget() *string {
	return s.Target
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) GetTheme() *ListTextbookAssistantSceneDetailsResponseBodyDataTheme {
	return s.Theme
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) GetTopic() *ListTextbookAssistantSceneDetailsResponseBodyDataTopic {
	return s.Topic
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) GetWordList() []*ListTextbookAssistantSceneDetailsResponseBodyDataWordList {
	return s.WordList
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) SetRoleList(v []*ListTextbookAssistantSceneDetailsResponseBodyDataRoleList) *ListTextbookAssistantSceneDetailsResponseBodyData {
	s.RoleList = v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) SetScene(v string) *ListTextbookAssistantSceneDetailsResponseBodyData {
	s.Scene = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) SetSceneId(v string) *ListTextbookAssistantSceneDetailsResponseBodyData {
	s.SceneId = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) SetSceneImageList(v []*string) *ListTextbookAssistantSceneDetailsResponseBodyData {
	s.SceneImageList = v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) SetSceneTaskList(v []*ListTextbookAssistantSceneDetailsResponseBodyDataSceneTaskList) *ListTextbookAssistantSceneDetailsResponseBodyData {
	s.SceneTaskList = v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) SetSceneTranslate(v string) *ListTextbookAssistantSceneDetailsResponseBodyData {
	s.SceneTranslate = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) SetSentenceList(v []*ListTextbookAssistantSceneDetailsResponseBodyDataSentenceList) *ListTextbookAssistantSceneDetailsResponseBodyData {
	s.SentenceList = v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) SetTarget(v string) *ListTextbookAssistantSceneDetailsResponseBodyData {
	s.Target = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) SetTheme(v *ListTextbookAssistantSceneDetailsResponseBodyDataTheme) *ListTextbookAssistantSceneDetailsResponseBodyData {
	s.Theme = v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) SetTopic(v *ListTextbookAssistantSceneDetailsResponseBodyDataTopic) *ListTextbookAssistantSceneDetailsResponseBodyData {
	s.Topic = v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) SetWordList(v []*ListTextbookAssistantSceneDetailsResponseBodyDataWordList) *ListTextbookAssistantSceneDetailsResponseBodyData {
	s.WordList = v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) Validate() error {
	if s.RoleList != nil {
		for _, item := range s.RoleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SceneTaskList != nil {
		for _, item := range s.SceneTaskList {
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

type ListTextbookAssistantSceneDetailsResponseBodyDataRoleList struct {
	// example:
	//
	// Carl, a curious boy
	Introduction          *string `json:"introduction,omitempty" xml:"introduction,omitempty"`
	IntroductionTranslate *string `json:"introductionTranslate,omitempty" xml:"introductionTranslate,omitempty"`
	// example:
	//
	// Hi Noah, who is that in the photo?
	Promoting          *string `json:"promoting,omitempty" xml:"promoting,omitempty"`
	PromotingTranslate *string `json:"promotingTranslate,omitempty" xml:"promotingTranslate,omitempty"`
	// example:
	//
	// Carl
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	// example:
	//
	// Carl
	RoleNameTranslate *string `json:"roleNameTranslate,omitempty" xml:"roleNameTranslate,omitempty"`
	// example:
	//
	// 0
	RoleType *string `json:"roleType,omitempty" xml:"roleType,omitempty"`
}

func (s ListTextbookAssistantSceneDetailsResponseBodyDataRoleList) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantSceneDetailsResponseBodyDataRoleList) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList) GetIntroduction() *string {
	return s.Introduction
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList) GetIntroductionTranslate() *string {
	return s.IntroductionTranslate
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList) GetPromoting() *string {
	return s.Promoting
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList) GetPromotingTranslate() *string {
	return s.PromotingTranslate
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList) GetRoleName() *string {
	return s.RoleName
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList) GetRoleNameTranslate() *string {
	return s.RoleNameTranslate
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList) GetRoleType() *string {
	return s.RoleType
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList) SetIntroduction(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList {
	s.Introduction = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList) SetIntroductionTranslate(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList {
	s.IntroductionTranslate = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList) SetPromoting(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList {
	s.Promoting = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList) SetPromotingTranslate(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList {
	s.PromotingTranslate = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList) SetRoleName(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList {
	s.RoleName = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList) SetRoleNameTranslate(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList {
	s.RoleNameTranslate = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList) SetRoleType(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList {
	s.RoleType = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList) Validate() error {
	return dara.Validate(s)
}

type ListTextbookAssistantSceneDetailsResponseBodyDataSceneTaskList struct {
	// example:
	//
	// Say that this is your dad\\"s brother.
	SceneTask          *string `json:"sceneTask,omitempty" xml:"sceneTask,omitempty"`
	SceneTaskTranslate *string `json:"sceneTaskTranslate,omitempty" xml:"sceneTaskTranslate,omitempty"`
}

func (s ListTextbookAssistantSceneDetailsResponseBodyDataSceneTaskList) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantSceneDetailsResponseBodyDataSceneTaskList) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataSceneTaskList) GetSceneTask() *string {
	return s.SceneTask
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataSceneTaskList) GetSceneTaskTranslate() *string {
	return s.SceneTaskTranslate
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataSceneTaskList) SetSceneTask(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataSceneTaskList {
	s.SceneTask = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataSceneTaskList) SetSceneTaskTranslate(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataSceneTaskList {
	s.SceneTaskTranslate = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataSceneTaskList) Validate() error {
	return dara.Validate(s)
}

type ListTextbookAssistantSceneDetailsResponseBodyDataSentenceList struct {
	SentenceAnalysis *string `json:"sentenceAnalysis,omitempty" xml:"sentenceAnalysis,omitempty"`
	// example:
	//
	// a774c6d09c4511eebe6e0c42a106bb02
	SentenceId *string `json:"sentenceId,omitempty" xml:"sentenceId,omitempty"`
	// example:
	//
	// Is this your sister?
	SentenceText *string `json:"sentenceText,omitempty" xml:"sentenceText,omitempty"`
}

func (s ListTextbookAssistantSceneDetailsResponseBodyDataSentenceList) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantSceneDetailsResponseBodyDataSentenceList) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataSentenceList) GetSentenceAnalysis() *string {
	return s.SentenceAnalysis
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataSentenceList) GetSentenceId() *string {
	return s.SentenceId
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataSentenceList) GetSentenceText() *string {
	return s.SentenceText
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataSentenceList) SetSentenceAnalysis(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataSentenceList {
	s.SentenceAnalysis = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataSentenceList) SetSentenceId(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataSentenceList {
	s.SentenceId = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataSentenceList) SetSentenceText(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataSentenceList {
	s.SentenceText = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataSentenceList) Validate() error {
	return dara.Validate(s)
}

type ListTextbookAssistantSceneDetailsResponseBodyDataTheme struct {
	ThemeImageList []*string `json:"themeImageList,omitempty" xml:"themeImageList,omitempty" type:"Repeated"`
	ThemeName      *string   `json:"themeName,omitempty" xml:"themeName,omitempty"`
	// example:
	//
	// Family and family life
	ThemeTranslate *string `json:"themeTranslate,omitempty" xml:"themeTranslate,omitempty"`
}

func (s ListTextbookAssistantSceneDetailsResponseBodyDataTheme) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantSceneDetailsResponseBodyDataTheme) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataTheme) GetThemeImageList() []*string {
	return s.ThemeImageList
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataTheme) GetThemeName() *string {
	return s.ThemeName
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataTheme) GetThemeTranslate() *string {
	return s.ThemeTranslate
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataTheme) SetThemeImageList(v []*string) *ListTextbookAssistantSceneDetailsResponseBodyDataTheme {
	s.ThemeImageList = v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataTheme) SetThemeName(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataTheme {
	s.ThemeName = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataTheme) SetThemeTranslate(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataTheme {
	s.ThemeTranslate = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataTheme) Validate() error {
	return dara.Validate(s)
}

type ListTextbookAssistantSceneDetailsResponseBodyDataTopic struct {
	TopicImageList []*string `json:"topicImageList,omitempty" xml:"topicImageList,omitempty" type:"Repeated"`
	TopicName      *string   `json:"topicName,omitempty" xml:"topicName,omitempty"`
	// example:
	//
	// Introducing family members
	TopicTranslate *string `json:"topicTranslate,omitempty" xml:"topicTranslate,omitempty"`
}

func (s ListTextbookAssistantSceneDetailsResponseBodyDataTopic) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantSceneDetailsResponseBodyDataTopic) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataTopic) GetTopicImageList() []*string {
	return s.TopicImageList
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataTopic) GetTopicName() *string {
	return s.TopicName
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataTopic) GetTopicTranslate() *string {
	return s.TopicTranslate
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataTopic) SetTopicImageList(v []*string) *ListTextbookAssistantSceneDetailsResponseBodyDataTopic {
	s.TopicImageList = v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataTopic) SetTopicName(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataTopic {
	s.TopicName = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataTopic) SetTopicTranslate(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataTopic {
	s.TopicTranslate = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataTopic) Validate() error {
	return dara.Validate(s)
}

type ListTextbookAssistantSceneDetailsResponseBodyDataWordList struct {
	WordAnalysis *string `json:"wordAnalysis,omitempty" xml:"wordAnalysis,omitempty"`
	// example:
	//
	// a94c3337ed8c11eebe6e0c42a106bb02
	WordId *string `json:"wordId,omitempty" xml:"wordId,omitempty"`
	// example:
	//
	// family
	WordText *string `json:"wordText,omitempty" xml:"wordText,omitempty"`
}

func (s ListTextbookAssistantSceneDetailsResponseBodyDataWordList) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantSceneDetailsResponseBodyDataWordList) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataWordList) GetWordAnalysis() *string {
	return s.WordAnalysis
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataWordList) GetWordId() *string {
	return s.WordId
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataWordList) GetWordText() *string {
	return s.WordText
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataWordList) SetWordAnalysis(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataWordList {
	s.WordAnalysis = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataWordList) SetWordId(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataWordList {
	s.WordId = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataWordList) SetWordText(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataWordList {
	s.WordText = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataWordList) Validate() error {
	return dara.Validate(s)
}
