// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNluResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessageId(v string) *NluResponseBody
	GetMessageId() *string
	SetMessages(v []*NluResponseBodyMessages) *NluResponseBody
	GetMessages() []*NluResponseBodyMessages
	SetRequestId(v string) *NluResponseBody
	GetRequestId() *string
}

type NluResponseBody struct {
	// example:
	//
	// 2828708A-2C7A-1BAE-B810-87DB9DA9C661
	MessageId *string                    `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	Messages  []*NluResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// example:
	//
	// A6357C1B-1D79-1382-B259-BD9E80751B42
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s NluResponseBody) String() string {
	return dara.Prettify(s)
}

func (s NluResponseBody) GoString() string {
	return s.String()
}

func (s *NluResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *NluResponseBody) GetMessages() []*NluResponseBodyMessages {
	return s.Messages
}

func (s *NluResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *NluResponseBody) SetMessageId(v string) *NluResponseBody {
	s.MessageId = &v
	return s
}

func (s *NluResponseBody) SetMessages(v []*NluResponseBodyMessages) *NluResponseBody {
	s.Messages = v
	return s
}

func (s *NluResponseBody) SetRequestId(v string) *NluResponseBody {
	s.RequestId = &v
	return s
}

func (s *NluResponseBody) Validate() error {
	if s.Messages != nil {
		for _, item := range s.Messages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type NluResponseBodyMessages struct {
	DialogHubNluInfo *NluResponseBodyMessagesDialogHubNluInfo `json:"DialogHubNluInfo,omitempty" xml:"DialogHubNluInfo,omitempty" type:"Struct"`
	DsNluInfo        *NluResponseBodyMessagesDsNluInfo        `json:"DsNluInfo,omitempty" xml:"DsNluInfo,omitempty" type:"Struct"`
}

func (s NluResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s NluResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *NluResponseBodyMessages) GetDialogHubNluInfo() *NluResponseBodyMessagesDialogHubNluInfo {
	return s.DialogHubNluInfo
}

func (s *NluResponseBodyMessages) GetDsNluInfo() *NluResponseBodyMessagesDsNluInfo {
	return s.DsNluInfo
}

func (s *NluResponseBodyMessages) SetDialogHubNluInfo(v *NluResponseBodyMessagesDialogHubNluInfo) *NluResponseBodyMessages {
	s.DialogHubNluInfo = v
	return s
}

func (s *NluResponseBodyMessages) SetDsNluInfo(v *NluResponseBodyMessagesDsNluInfo) *NluResponseBodyMessages {
	s.DsNluInfo = v
	return s
}

func (s *NluResponseBodyMessages) Validate() error {
	if s.DialogHubNluInfo != nil {
		if err := s.DialogHubNluInfo.Validate(); err != nil {
			return err
		}
	}
	if s.DsNluInfo != nil {
		if err := s.DsNluInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type NluResponseBodyMessagesDialogHubNluInfo struct {
	GlobalDictList          []*NluResponseBodyMessagesDialogHubNluInfoGlobalDictList          `json:"GlobalDictList,omitempty" xml:"GlobalDictList,omitempty" type:"Repeated"`
	GlobalSensitiveWordList []*NluResponseBodyMessagesDialogHubNluInfoGlobalSensitiveWordList `json:"GlobalSensitiveWordList,omitempty" xml:"GlobalSensitiveWordList,omitempty" type:"Repeated"`
}

func (s NluResponseBodyMessagesDialogHubNluInfo) String() string {
	return dara.Prettify(s)
}

func (s NluResponseBodyMessagesDialogHubNluInfo) GoString() string {
	return s.String()
}

func (s *NluResponseBodyMessagesDialogHubNluInfo) GetGlobalDictList() []*NluResponseBodyMessagesDialogHubNluInfoGlobalDictList {
	return s.GlobalDictList
}

func (s *NluResponseBodyMessagesDialogHubNluInfo) GetGlobalSensitiveWordList() []*NluResponseBodyMessagesDialogHubNluInfoGlobalSensitiveWordList {
	return s.GlobalSensitiveWordList
}

func (s *NluResponseBodyMessagesDialogHubNluInfo) SetGlobalDictList(v []*NluResponseBodyMessagesDialogHubNluInfoGlobalDictList) *NluResponseBodyMessagesDialogHubNluInfo {
	s.GlobalDictList = v
	return s
}

func (s *NluResponseBodyMessagesDialogHubNluInfo) SetGlobalSensitiveWordList(v []*NluResponseBodyMessagesDialogHubNluInfoGlobalSensitiveWordList) *NluResponseBodyMessagesDialogHubNluInfo {
	s.GlobalSensitiveWordList = v
	return s
}

func (s *NluResponseBodyMessagesDialogHubNluInfo) Validate() error {
	if s.GlobalDictList != nil {
		for _, item := range s.GlobalDictList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.GlobalSensitiveWordList != nil {
		for _, item := range s.GlobalSensitiveWordList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type NluResponseBodyMessagesDialogHubNluInfoGlobalDictList struct {
	// example:
	//
	// 天气
	StandardWord *string `json:"StandardWord,omitempty" xml:"StandardWord,omitempty"`
	// example:
	//
	// 天气
	Word *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s NluResponseBodyMessagesDialogHubNluInfoGlobalDictList) String() string {
	return dara.Prettify(s)
}

func (s NluResponseBodyMessagesDialogHubNluInfoGlobalDictList) GoString() string {
	return s.String()
}

func (s *NluResponseBodyMessagesDialogHubNluInfoGlobalDictList) GetStandardWord() *string {
	return s.StandardWord
}

func (s *NluResponseBodyMessagesDialogHubNluInfoGlobalDictList) GetWord() *string {
	return s.Word
}

func (s *NluResponseBodyMessagesDialogHubNluInfoGlobalDictList) SetStandardWord(v string) *NluResponseBodyMessagesDialogHubNluInfoGlobalDictList {
	s.StandardWord = &v
	return s
}

func (s *NluResponseBodyMessagesDialogHubNluInfoGlobalDictList) SetWord(v string) *NluResponseBodyMessagesDialogHubNluInfoGlobalDictList {
	s.Word = &v
	return s
}

func (s *NluResponseBodyMessagesDialogHubNluInfoGlobalDictList) Validate() error {
	return dara.Validate(s)
}

type NluResponseBodyMessagesDialogHubNluInfoGlobalSensitiveWordList struct {
	// example:
	//
	// 天气
	StandardWord *string `json:"StandardWord,omitempty" xml:"StandardWord,omitempty"`
	// example:
	//
	// 天气
	Word *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s NluResponseBodyMessagesDialogHubNluInfoGlobalSensitiveWordList) String() string {
	return dara.Prettify(s)
}

func (s NluResponseBodyMessagesDialogHubNluInfoGlobalSensitiveWordList) GoString() string {
	return s.String()
}

func (s *NluResponseBodyMessagesDialogHubNluInfoGlobalSensitiveWordList) GetStandardWord() *string {
	return s.StandardWord
}

func (s *NluResponseBodyMessagesDialogHubNluInfoGlobalSensitiveWordList) GetWord() *string {
	return s.Word
}

func (s *NluResponseBodyMessagesDialogHubNluInfoGlobalSensitiveWordList) SetStandardWord(v string) *NluResponseBodyMessagesDialogHubNluInfoGlobalSensitiveWordList {
	s.StandardWord = &v
	return s
}

func (s *NluResponseBodyMessagesDialogHubNluInfoGlobalSensitiveWordList) SetWord(v string) *NluResponseBodyMessagesDialogHubNluInfoGlobalSensitiveWordList {
	s.Word = &v
	return s
}

func (s *NluResponseBodyMessagesDialogHubNluInfoGlobalSensitiveWordList) Validate() error {
	return dara.Validate(s)
}

type NluResponseBodyMessagesDsNluInfo struct {
	EntityList []*NluResponseBodyMessagesDsNluInfoEntityList `json:"EntityList,omitempty" xml:"EntityList,omitempty" type:"Repeated"`
	IntentList []*NluResponseBodyMessagesDsNluInfoIntentList `json:"IntentList,omitempty" xml:"IntentList,omitempty" type:"Repeated"`
}

func (s NluResponseBodyMessagesDsNluInfo) String() string {
	return dara.Prettify(s)
}

func (s NluResponseBodyMessagesDsNluInfo) GoString() string {
	return s.String()
}

func (s *NluResponseBodyMessagesDsNluInfo) GetEntityList() []*NluResponseBodyMessagesDsNluInfoEntityList {
	return s.EntityList
}

func (s *NluResponseBodyMessagesDsNluInfo) GetIntentList() []*NluResponseBodyMessagesDsNluInfoIntentList {
	return s.IntentList
}

func (s *NluResponseBodyMessagesDsNluInfo) SetEntityList(v []*NluResponseBodyMessagesDsNluInfoEntityList) *NluResponseBodyMessagesDsNluInfo {
	s.EntityList = v
	return s
}

func (s *NluResponseBodyMessagesDsNluInfo) SetIntentList(v []*NluResponseBodyMessagesDsNluInfoIntentList) *NluResponseBodyMessagesDsNluInfo {
	s.IntentList = v
	return s
}

func (s *NluResponseBodyMessagesDsNluInfo) Validate() error {
	if s.EntityList != nil {
		for _, item := range s.EntityList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.IntentList != nil {
		for _, item := range s.IntentList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type NluResponseBodyMessagesDsNluInfoEntityList struct {
	// example:
	//
	// @城市
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 北京
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 首都
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s NluResponseBodyMessagesDsNluInfoEntityList) String() string {
	return dara.Prettify(s)
}

func (s NluResponseBodyMessagesDsNluInfoEntityList) GoString() string {
	return s.String()
}

func (s *NluResponseBodyMessagesDsNluInfoEntityList) GetName() *string {
	return s.Name
}

func (s *NluResponseBodyMessagesDsNluInfoEntityList) GetOrigin() *string {
	return s.Origin
}

func (s *NluResponseBodyMessagesDsNluInfoEntityList) GetType() *string {
	return s.Type
}

func (s *NluResponseBodyMessagesDsNluInfoEntityList) GetValue() *string {
	return s.Value
}

func (s *NluResponseBodyMessagesDsNluInfoEntityList) SetName(v string) *NluResponseBodyMessagesDsNluInfoEntityList {
	s.Name = &v
	return s
}

func (s *NluResponseBodyMessagesDsNluInfoEntityList) SetOrigin(v string) *NluResponseBodyMessagesDsNluInfoEntityList {
	s.Origin = &v
	return s
}

func (s *NluResponseBodyMessagesDsNluInfoEntityList) SetType(v string) *NluResponseBodyMessagesDsNluInfoEntityList {
	s.Type = &v
	return s
}

func (s *NluResponseBodyMessagesDsNluInfoEntityList) SetValue(v string) *NluResponseBodyMessagesDsNluInfoEntityList {
	s.Value = &v
	return s
}

func (s *NluResponseBodyMessagesDsNluInfoEntityList) Validate() error {
	return dara.Validate(s)
}

type NluResponseBodyMessagesDsNluInfoIntentList struct {
	// example:
	//
	// 724387
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// example:
	//
	// classifierType=Fewshot,from=Fewshot,content=[我要查北京的天气, 帮我查北京的天气, 北京天气怎么样, 北京今天下雨吗, 北京今天多少度]
	MatchDetail *string `json:"MatchDetail,omitempty" xml:"MatchDetail,omitempty"`
	// example:
	//
	// FewShotLearning
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// example:
	//
	// 查天气意图
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0.995
	Score    *float64                                              `json:"Score,omitempty" xml:"Score,omitempty"`
	SlotList []*NluResponseBodyMessagesDsNluInfoIntentListSlotList `json:"SlotList,omitempty" xml:"SlotList,omitempty" type:"Repeated"`
}

func (s NluResponseBodyMessagesDsNluInfoIntentList) String() string {
	return dara.Prettify(s)
}

func (s NluResponseBodyMessagesDsNluInfoIntentList) GoString() string {
	return s.String()
}

func (s *NluResponseBodyMessagesDsNluInfoIntentList) GetIntentId() *int64 {
	return s.IntentId
}

func (s *NluResponseBodyMessagesDsNluInfoIntentList) GetMatchDetail() *string {
	return s.MatchDetail
}

func (s *NluResponseBodyMessagesDsNluInfoIntentList) GetMatchType() *string {
	return s.MatchType
}

func (s *NluResponseBodyMessagesDsNluInfoIntentList) GetName() *string {
	return s.Name
}

func (s *NluResponseBodyMessagesDsNluInfoIntentList) GetScore() *float64 {
	return s.Score
}

func (s *NluResponseBodyMessagesDsNluInfoIntentList) GetSlotList() []*NluResponseBodyMessagesDsNluInfoIntentListSlotList {
	return s.SlotList
}

func (s *NluResponseBodyMessagesDsNluInfoIntentList) SetIntentId(v int64) *NluResponseBodyMessagesDsNluInfoIntentList {
	s.IntentId = &v
	return s
}

func (s *NluResponseBodyMessagesDsNluInfoIntentList) SetMatchDetail(v string) *NluResponseBodyMessagesDsNluInfoIntentList {
	s.MatchDetail = &v
	return s
}

func (s *NluResponseBodyMessagesDsNluInfoIntentList) SetMatchType(v string) *NluResponseBodyMessagesDsNluInfoIntentList {
	s.MatchType = &v
	return s
}

func (s *NluResponseBodyMessagesDsNluInfoIntentList) SetName(v string) *NluResponseBodyMessagesDsNluInfoIntentList {
	s.Name = &v
	return s
}

func (s *NluResponseBodyMessagesDsNluInfoIntentList) SetScore(v float64) *NluResponseBodyMessagesDsNluInfoIntentList {
	s.Score = &v
	return s
}

func (s *NluResponseBodyMessagesDsNluInfoIntentList) SetSlotList(v []*NluResponseBodyMessagesDsNluInfoIntentListSlotList) *NluResponseBodyMessagesDsNluInfoIntentList {
	s.SlotList = v
	return s
}

func (s *NluResponseBodyMessagesDsNluInfoIntentList) Validate() error {
	if s.SlotList != nil {
		for _, item := range s.SlotList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type NluResponseBodyMessagesDsNluInfoIntentListSlotList struct {
	// example:
	//
	// @城市
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 北京
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 首都
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s NluResponseBodyMessagesDsNluInfoIntentListSlotList) String() string {
	return dara.Prettify(s)
}

func (s NluResponseBodyMessagesDsNluInfoIntentListSlotList) GoString() string {
	return s.String()
}

func (s *NluResponseBodyMessagesDsNluInfoIntentListSlotList) GetName() *string {
	return s.Name
}

func (s *NluResponseBodyMessagesDsNluInfoIntentListSlotList) GetOrigin() *string {
	return s.Origin
}

func (s *NluResponseBodyMessagesDsNluInfoIntentListSlotList) GetType() *string {
	return s.Type
}

func (s *NluResponseBodyMessagesDsNluInfoIntentListSlotList) GetValue() *string {
	return s.Value
}

func (s *NluResponseBodyMessagesDsNluInfoIntentListSlotList) SetName(v string) *NluResponseBodyMessagesDsNluInfoIntentListSlotList {
	s.Name = &v
	return s
}

func (s *NluResponseBodyMessagesDsNluInfoIntentListSlotList) SetOrigin(v string) *NluResponseBodyMessagesDsNluInfoIntentListSlotList {
	s.Origin = &v
	return s
}

func (s *NluResponseBodyMessagesDsNluInfoIntentListSlotList) SetType(v string) *NluResponseBodyMessagesDsNluInfoIntentListSlotList {
	s.Type = &v
	return s
}

func (s *NluResponseBodyMessagesDsNluInfoIntentListSlotList) SetValue(v string) *NluResponseBodyMessagesDsNluInfoIntentListSlotList {
	s.Value = &v
	return s
}

func (s *NluResponseBodyMessagesDsNluInfoIntentListSlotList) Validate() error {
	return dara.Validate(s)
}
