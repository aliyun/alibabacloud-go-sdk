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
	Agents []*DescribePolarClawAgentsResponseBodyAgents `json:"Agents,omitempty" xml:"Agents,omitempty" type:"Repeated"`
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// main
	DefaultId *string `json:"DefaultId,omitempty" xml:"DefaultId,omitempty"`
	// example:
	//
	// main
	MainKey *string `json:"MainKey,omitempty" xml:"MainKey,omitempty"`
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
	// Agent ID
	//
	// example:
	//
	// main
	Id       *string                                            `json:"Id,omitempty" xml:"Id,omitempty"`
	Identity *DescribePolarClawAgentsResponseBodyAgentsIdentity `json:"Identity,omitempty" xml:"Identity,omitempty" type:"Struct"`
	// example:
	//
	// main
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribePolarClawAgentsResponseBodyAgents) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawAgentsResponseBodyAgents) GoString() string {
	return s.String()
}

func (s *DescribePolarClawAgentsResponseBodyAgents) GetId() *string {
	return s.Id
}

func (s *DescribePolarClawAgentsResponseBodyAgents) GetIdentity() *DescribePolarClawAgentsResponseBodyAgentsIdentity {
	return s.Identity
}

func (s *DescribePolarClawAgentsResponseBodyAgents) GetName() *string {
	return s.Name
}

func (s *DescribePolarClawAgentsResponseBodyAgents) SetId(v string) *DescribePolarClawAgentsResponseBodyAgents {
	s.Id = &v
	return s
}

func (s *DescribePolarClawAgentsResponseBodyAgents) SetIdentity(v *DescribePolarClawAgentsResponseBodyAgentsIdentity) *DescribePolarClawAgentsResponseBodyAgents {
	s.Identity = v
	return s
}

func (s *DescribePolarClawAgentsResponseBodyAgents) SetName(v string) *DescribePolarClawAgentsResponseBodyAgents {
	s.Name = &v
	return s
}

func (s *DescribePolarClawAgentsResponseBodyAgents) Validate() error {
	if s.Identity != nil {
		if err := s.Identity.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePolarClawAgentsResponseBodyAgentsIdentity struct {
	// example:
	//
	// test
	Avatar *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	// example:
	//
	// test
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	// example:
	//
	// U+1F99E
	Emoji *string `json:"Emoji,omitempty" xml:"Emoji,omitempty"`
	// example:
	//
	// PolarClaw
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
