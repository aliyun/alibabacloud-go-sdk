// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawAgentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgents(v []*DescribePolarClawAgentsResponseBodyAgents) *DescribePolarClawAgentsResponseBody
	GetAgents() []*DescribePolarClawAgentsResponseBodyAgents
	SetApplicationId(v string) *DescribePolarClawAgentsResponseBody
	GetApplicationId() *string
	SetCode(v int32) *DescribePolarClawAgentsResponseBody
	GetCode() *int32
	SetDefaultId(v string) *DescribePolarClawAgentsResponseBody
	GetDefaultId() *string
	SetMainKey(v string) *DescribePolarClawAgentsResponseBody
	GetMainKey() *string
	SetMessage(v string) *DescribePolarClawAgentsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribePolarClawAgentsResponseBody
	GetRequestId() *string
	SetScope(v string) *DescribePolarClawAgentsResponseBody
	GetScope() *string
}

type DescribePolarClawAgentsResponseBody struct {
	// The agent list.
	Agents []*DescribePolarClawAgentsResponseBodyAgents `json:"Agents,omitempty" xml:"Agents,omitempty" type:"Repeated"`
	// The application ID.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The default agent ID.
	//
	// example:
	//
	// main
	DefaultId *string `json:"DefaultId,omitempty" xml:"DefaultId,omitempty"`
	// The primary agent key name.
	//
	// example:
	//
	// main
	MainKey *string `json:"MainKey,omitempty" xml:"MainKey,omitempty"`
	// The response message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 24A1990B-4F6E-482B-B8CB-75C612******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The routing scope.
	//
	// example:
	//
	// per-sender
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s DescribePolarClawAgentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawAgentsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarClawAgentsResponseBody) GetAgents() []*DescribePolarClawAgentsResponseBodyAgents {
	return s.Agents
}

func (s *DescribePolarClawAgentsResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribePolarClawAgentsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribePolarClawAgentsResponseBody) GetDefaultId() *string {
	return s.DefaultId
}

func (s *DescribePolarClawAgentsResponseBody) GetMainKey() *string {
	return s.MainKey
}

func (s *DescribePolarClawAgentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePolarClawAgentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolarClawAgentsResponseBody) GetScope() *string {
	return s.Scope
}

func (s *DescribePolarClawAgentsResponseBody) SetAgents(v []*DescribePolarClawAgentsResponseBodyAgents) *DescribePolarClawAgentsResponseBody {
	s.Agents = v
	return s
}

func (s *DescribePolarClawAgentsResponseBody) SetApplicationId(v string) *DescribePolarClawAgentsResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *DescribePolarClawAgentsResponseBody) SetCode(v int32) *DescribePolarClawAgentsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePolarClawAgentsResponseBody) SetDefaultId(v string) *DescribePolarClawAgentsResponseBody {
	s.DefaultId = &v
	return s
}

func (s *DescribePolarClawAgentsResponseBody) SetMainKey(v string) *DescribePolarClawAgentsResponseBody {
	s.MainKey = &v
	return s
}

func (s *DescribePolarClawAgentsResponseBody) SetMessage(v string) *DescribePolarClawAgentsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePolarClawAgentsResponseBody) SetRequestId(v string) *DescribePolarClawAgentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarClawAgentsResponseBody) SetScope(v string) *DescribePolarClawAgentsResponseBody {
	s.Scope = &v
	return s
}

func (s *DescribePolarClawAgentsResponseBody) Validate() error {
	if s.Agents != nil {
		for _, item := range s.Agents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolarClawAgentsResponseBodyAgents struct {
	// example:
	//
	// true
	Default *bool                                             `json:"Default,omitempty" xml:"Default,omitempty"`
	Files   []*DescribePolarClawAgentsResponseBodyAgentsFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// Agent ID
	//
	// example:
	//
	// main
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The identity information.
	Identity *DescribePolarClawAgentsResponseBodyAgentsIdentity `json:"Identity,omitempty" xml:"Identity,omitempty" type:"Struct"`
	Model    *DescribePolarClawAgentsResponseBodyAgentsModel    `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// The display name of the agent.
	//
	// example:
	//
	// main
	Name   *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Skills []*string `json:"Skills,omitempty" xml:"Skills,omitempty" type:"Repeated"`
	// example:
	//
	// /home/node/.openclaw/workspace-work
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s DescribePolarClawAgentsResponseBodyAgents) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawAgentsResponseBodyAgents) GoString() string {
	return s.String()
}

func (s *DescribePolarClawAgentsResponseBodyAgents) GetDefault() *bool {
	return s.Default
}

func (s *DescribePolarClawAgentsResponseBodyAgents) GetFiles() []*DescribePolarClawAgentsResponseBodyAgentsFiles {
	return s.Files
}

func (s *DescribePolarClawAgentsResponseBodyAgents) GetId() *string {
	return s.Id
}

func (s *DescribePolarClawAgentsResponseBodyAgents) GetIdentity() *DescribePolarClawAgentsResponseBodyAgentsIdentity {
	return s.Identity
}

func (s *DescribePolarClawAgentsResponseBodyAgents) GetModel() *DescribePolarClawAgentsResponseBodyAgentsModel {
	return s.Model
}

func (s *DescribePolarClawAgentsResponseBodyAgents) GetName() *string {
	return s.Name
}

func (s *DescribePolarClawAgentsResponseBodyAgents) GetSkills() []*string {
	return s.Skills
}

func (s *DescribePolarClawAgentsResponseBodyAgents) GetWorkspace() *string {
	return s.Workspace
}

func (s *DescribePolarClawAgentsResponseBodyAgents) SetDefault(v bool) *DescribePolarClawAgentsResponseBodyAgents {
	s.Default = &v
	return s
}

func (s *DescribePolarClawAgentsResponseBodyAgents) SetFiles(v []*DescribePolarClawAgentsResponseBodyAgentsFiles) *DescribePolarClawAgentsResponseBodyAgents {
	s.Files = v
	return s
}

func (s *DescribePolarClawAgentsResponseBodyAgents) SetId(v string) *DescribePolarClawAgentsResponseBodyAgents {
	s.Id = &v
	return s
}

func (s *DescribePolarClawAgentsResponseBodyAgents) SetIdentity(v *DescribePolarClawAgentsResponseBodyAgentsIdentity) *DescribePolarClawAgentsResponseBodyAgents {
	s.Identity = v
	return s
}

func (s *DescribePolarClawAgentsResponseBodyAgents) SetModel(v *DescribePolarClawAgentsResponseBodyAgentsModel) *DescribePolarClawAgentsResponseBodyAgents {
	s.Model = v
	return s
}

func (s *DescribePolarClawAgentsResponseBodyAgents) SetName(v string) *DescribePolarClawAgentsResponseBodyAgents {
	s.Name = &v
	return s
}

func (s *DescribePolarClawAgentsResponseBodyAgents) SetSkills(v []*string) *DescribePolarClawAgentsResponseBodyAgents {
	s.Skills = v
	return s
}

func (s *DescribePolarClawAgentsResponseBodyAgents) SetWorkspace(v string) *DescribePolarClawAgentsResponseBodyAgents {
	s.Workspace = &v
	return s
}

func (s *DescribePolarClawAgentsResponseBodyAgents) Validate() error {
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Identity != nil {
		if err := s.Identity.Validate(); err != nil {
			return err
		}
	}
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePolarClawAgentsResponseBodyAgentsFiles struct {
	// example:
	//
	// false
	Missing *bool `json:"Missing,omitempty" xml:"Missing,omitempty"`
	// example:
	//
	// SOUL.md
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// /home/node/.openclaw/workspace-work/SOUL.md
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// 1024
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// 1716000000000
	UpdatedAtMs *int64 `json:"UpdatedAtMs,omitempty" xml:"UpdatedAtMs,omitempty"`
}

func (s DescribePolarClawAgentsResponseBodyAgentsFiles) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawAgentsResponseBodyAgentsFiles) GoString() string {
	return s.String()
}

func (s *DescribePolarClawAgentsResponseBodyAgentsFiles) GetMissing() *bool {
	return s.Missing
}

func (s *DescribePolarClawAgentsResponseBodyAgentsFiles) GetName() *string {
	return s.Name
}

func (s *DescribePolarClawAgentsResponseBodyAgentsFiles) GetPath() *string {
	return s.Path
}

func (s *DescribePolarClawAgentsResponseBodyAgentsFiles) GetSize() *int64 {
	return s.Size
}

func (s *DescribePolarClawAgentsResponseBodyAgentsFiles) GetUpdatedAtMs() *int64 {
	return s.UpdatedAtMs
}

func (s *DescribePolarClawAgentsResponseBodyAgentsFiles) SetMissing(v bool) *DescribePolarClawAgentsResponseBodyAgentsFiles {
	s.Missing = &v
	return s
}

func (s *DescribePolarClawAgentsResponseBodyAgentsFiles) SetName(v string) *DescribePolarClawAgentsResponseBodyAgentsFiles {
	s.Name = &v
	return s
}

func (s *DescribePolarClawAgentsResponseBodyAgentsFiles) SetPath(v string) *DescribePolarClawAgentsResponseBodyAgentsFiles {
	s.Path = &v
	return s
}

func (s *DescribePolarClawAgentsResponseBodyAgentsFiles) SetSize(v int64) *DescribePolarClawAgentsResponseBodyAgentsFiles {
	s.Size = &v
	return s
}

func (s *DescribePolarClawAgentsResponseBodyAgentsFiles) SetUpdatedAtMs(v int64) *DescribePolarClawAgentsResponseBodyAgentsFiles {
	s.UpdatedAtMs = &v
	return s
}

func (s *DescribePolarClawAgentsResponseBodyAgentsFiles) Validate() error {
	return dara.Validate(s)
}

type DescribePolarClawAgentsResponseBodyAgentsIdentity struct {
	// The avatar path or content.
	//
	// example:
	//
	// test
	Avatar *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	// The avatar URL.
	//
	// example:
	//
	// test
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	// The emoji identifier in Unicode encoding format such as U+1F99E, or a direct emoji character.
	//
	// example:
	//
	// U+1F99E
	Emoji *string `json:"Emoji,omitempty" xml:"Emoji,omitempty"`
	// The identity name.
	//
	// example:
	//
	// PolarClaw
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The theme.
	//
	// example:
	//
	// space lobster
	Theme *string `json:"Theme,omitempty" xml:"Theme,omitempty"`
}

func (s DescribePolarClawAgentsResponseBodyAgentsIdentity) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawAgentsResponseBodyAgentsIdentity) GoString() string {
	return s.String()
}

func (s *DescribePolarClawAgentsResponseBodyAgentsIdentity) GetAvatar() *string {
	return s.Avatar
}

func (s *DescribePolarClawAgentsResponseBodyAgentsIdentity) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *DescribePolarClawAgentsResponseBodyAgentsIdentity) GetEmoji() *string {
	return s.Emoji
}

func (s *DescribePolarClawAgentsResponseBodyAgentsIdentity) GetName() *string {
	return s.Name
}

func (s *DescribePolarClawAgentsResponseBodyAgentsIdentity) GetTheme() *string {
	return s.Theme
}

func (s *DescribePolarClawAgentsResponseBodyAgentsIdentity) SetAvatar(v string) *DescribePolarClawAgentsResponseBodyAgentsIdentity {
	s.Avatar = &v
	return s
}

func (s *DescribePolarClawAgentsResponseBodyAgentsIdentity) SetAvatarUrl(v string) *DescribePolarClawAgentsResponseBodyAgentsIdentity {
	s.AvatarUrl = &v
	return s
}

func (s *DescribePolarClawAgentsResponseBodyAgentsIdentity) SetEmoji(v string) *DescribePolarClawAgentsResponseBodyAgentsIdentity {
	s.Emoji = &v
	return s
}

func (s *DescribePolarClawAgentsResponseBodyAgentsIdentity) SetName(v string) *DescribePolarClawAgentsResponseBodyAgentsIdentity {
	s.Name = &v
	return s
}

func (s *DescribePolarClawAgentsResponseBodyAgentsIdentity) SetTheme(v string) *DescribePolarClawAgentsResponseBodyAgentsIdentity {
	s.Theme = &v
	return s
}

func (s *DescribePolarClawAgentsResponseBodyAgentsIdentity) Validate() error {
	return dara.Validate(s)
}

type DescribePolarClawAgentsResponseBodyAgentsModel struct {
	Fallbacks []*string `json:"Fallbacks,omitempty" xml:"Fallbacks,omitempty" type:"Repeated"`
	// example:
	//
	// claude-sonnet-4-5
	Primary *string `json:"Primary,omitempty" xml:"Primary,omitempty"`
}

func (s DescribePolarClawAgentsResponseBodyAgentsModel) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawAgentsResponseBodyAgentsModel) GoString() string {
	return s.String()
}

func (s *DescribePolarClawAgentsResponseBodyAgentsModel) GetFallbacks() []*string {
	return s.Fallbacks
}

func (s *DescribePolarClawAgentsResponseBodyAgentsModel) GetPrimary() *string {
	return s.Primary
}

func (s *DescribePolarClawAgentsResponseBodyAgentsModel) SetFallbacks(v []*string) *DescribePolarClawAgentsResponseBodyAgentsModel {
	s.Fallbacks = v
	return s
}

func (s *DescribePolarClawAgentsResponseBodyAgentsModel) SetPrimary(v string) *DescribePolarClawAgentsResponseBodyAgentsModel {
	s.Primary = &v
	return s
}

func (s *DescribePolarClawAgentsResponseBodyAgentsModel) Validate() error {
	return dara.Validate(s)
}
