// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*ListAgentsResponseBodyList) *ListAgentsResponseBody
	GetList() []*ListAgentsResponseBodyList
	SetRequestId(v string) *ListAgentsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAgentsResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListAgentsResponseBody
	GetTotal() *int32
}

type ListAgentsResponseBody struct {
	List      []*ListAgentsResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	RequestId *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                         `json:"success,omitempty" xml:"success,omitempty"`
	Total     *int32                        `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAgentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentsResponseBody) GetList() []*ListAgentsResponseBodyList {
	return s.List
}

func (s *ListAgentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAgentsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListAgentsResponseBody) SetList(v []*ListAgentsResponseBodyList) *ListAgentsResponseBody {
	s.List = v
	return s
}

func (s *ListAgentsResponseBody) SetRequestId(v string) *ListAgentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentsResponseBody) SetSuccess(v bool) *ListAgentsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAgentsResponseBody) SetTotal(v int32) *ListAgentsResponseBody {
	s.Total = &v
	return s
}

func (s *ListAgentsResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAgentsResponseBodyList struct {
	AgentIconUrl          *string `json:"agentIconUrl,omitempty" xml:"agentIconUrl,omitempty"`
	AgentId               *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	AgentName             *string `json:"agentName,omitempty" xml:"agentName,omitempty"`
	AgentScene            *string `json:"agentScene,omitempty" xml:"agentScene,omitempty"`
	CharacterAgeStage     *string `json:"characterAgeStage,omitempty" xml:"characterAgeStage,omitempty"`
	CharacterGender       *string `json:"characterGender,omitempty" xml:"characterGender,omitempty"`
	CharacterName         *string `json:"characterName,omitempty" xml:"characterName,omitempty"`
	CharactersDescription *string `json:"charactersDescription,omitempty" xml:"charactersDescription,omitempty"`
	EnableInteraction     *int32  `json:"enableInteraction,omitempty" xml:"enableInteraction,omitempty"`
	Industry              *string `json:"industry,omitempty" xml:"industry,omitempty"`
	OnlineSearch          *bool   `json:"onlineSearch,omitempty" xml:"onlineSearch,omitempty"`
	Owner                 *string `json:"owner,omitempty" xml:"owner,omitempty"`
	ReferenceUrl          *string `json:"referenceUrl,omitempty" xml:"referenceUrl,omitempty"`
	Status                *int32  `json:"status,omitempty" xml:"status,omitempty"`
	TextStyle             *string `json:"textStyle,omitempty" xml:"textStyle,omitempty"`
	Viewer                *string `json:"viewer,omitempty" xml:"viewer,omitempty"`
}

func (s ListAgentsResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListAgentsResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListAgentsResponseBodyList) GetAgentIconUrl() *string {
	return s.AgentIconUrl
}

func (s *ListAgentsResponseBodyList) GetAgentId() *string {
	return s.AgentId
}

func (s *ListAgentsResponseBodyList) GetAgentName() *string {
	return s.AgentName
}

func (s *ListAgentsResponseBodyList) GetAgentScene() *string {
	return s.AgentScene
}

func (s *ListAgentsResponseBodyList) GetCharacterAgeStage() *string {
	return s.CharacterAgeStage
}

func (s *ListAgentsResponseBodyList) GetCharacterGender() *string {
	return s.CharacterGender
}

func (s *ListAgentsResponseBodyList) GetCharacterName() *string {
	return s.CharacterName
}

func (s *ListAgentsResponseBodyList) GetCharactersDescription() *string {
	return s.CharactersDescription
}

func (s *ListAgentsResponseBodyList) GetEnableInteraction() *int32 {
	return s.EnableInteraction
}

func (s *ListAgentsResponseBodyList) GetIndustry() *string {
	return s.Industry
}

func (s *ListAgentsResponseBodyList) GetOnlineSearch() *bool {
	return s.OnlineSearch
}

func (s *ListAgentsResponseBodyList) GetOwner() *string {
	return s.Owner
}

func (s *ListAgentsResponseBodyList) GetReferenceUrl() *string {
	return s.ReferenceUrl
}

func (s *ListAgentsResponseBodyList) GetStatus() *int32 {
	return s.Status
}

func (s *ListAgentsResponseBodyList) GetTextStyle() *string {
	return s.TextStyle
}

func (s *ListAgentsResponseBodyList) GetViewer() *string {
	return s.Viewer
}

func (s *ListAgentsResponseBodyList) SetAgentIconUrl(v string) *ListAgentsResponseBodyList {
	s.AgentIconUrl = &v
	return s
}

func (s *ListAgentsResponseBodyList) SetAgentId(v string) *ListAgentsResponseBodyList {
	s.AgentId = &v
	return s
}

func (s *ListAgentsResponseBodyList) SetAgentName(v string) *ListAgentsResponseBodyList {
	s.AgentName = &v
	return s
}

func (s *ListAgentsResponseBodyList) SetAgentScene(v string) *ListAgentsResponseBodyList {
	s.AgentScene = &v
	return s
}

func (s *ListAgentsResponseBodyList) SetCharacterAgeStage(v string) *ListAgentsResponseBodyList {
	s.CharacterAgeStage = &v
	return s
}

func (s *ListAgentsResponseBodyList) SetCharacterGender(v string) *ListAgentsResponseBodyList {
	s.CharacterGender = &v
	return s
}

func (s *ListAgentsResponseBodyList) SetCharacterName(v string) *ListAgentsResponseBodyList {
	s.CharacterName = &v
	return s
}

func (s *ListAgentsResponseBodyList) SetCharactersDescription(v string) *ListAgentsResponseBodyList {
	s.CharactersDescription = &v
	return s
}

func (s *ListAgentsResponseBodyList) SetEnableInteraction(v int32) *ListAgentsResponseBodyList {
	s.EnableInteraction = &v
	return s
}

func (s *ListAgentsResponseBodyList) SetIndustry(v string) *ListAgentsResponseBodyList {
	s.Industry = &v
	return s
}

func (s *ListAgentsResponseBodyList) SetOnlineSearch(v bool) *ListAgentsResponseBodyList {
	s.OnlineSearch = &v
	return s
}

func (s *ListAgentsResponseBodyList) SetOwner(v string) *ListAgentsResponseBodyList {
	s.Owner = &v
	return s
}

func (s *ListAgentsResponseBodyList) SetReferenceUrl(v string) *ListAgentsResponseBodyList {
	s.ReferenceUrl = &v
	return s
}

func (s *ListAgentsResponseBodyList) SetStatus(v int32) *ListAgentsResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListAgentsResponseBodyList) SetTextStyle(v string) *ListAgentsResponseBodyList {
	s.TextStyle = &v
	return s
}

func (s *ListAgentsResponseBodyList) SetViewer(v string) *ListAgentsResponseBodyList {
	s.Viewer = &v
	return s
}

func (s *ListAgentsResponseBodyList) Validate() error {
	return dara.Validate(s)
}
