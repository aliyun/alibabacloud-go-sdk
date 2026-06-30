// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawAgentToolsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *DescribePolarClawAgentToolsResponseBody
	GetAgentId() *string
	SetApplicationId(v string) *DescribePolarClawAgentToolsResponseBody
	GetApplicationId() *string
	SetCode(v int32) *DescribePolarClawAgentToolsResponseBody
	GetCode() *int32
	SetCurrentConfig(v *DescribePolarClawAgentToolsResponseBodyCurrentConfig) *DescribePolarClawAgentToolsResponseBody
	GetCurrentConfig() *DescribePolarClawAgentToolsResponseBodyCurrentConfig
	SetGroups(v []*DescribePolarClawAgentToolsResponseBodyGroups) *DescribePolarClawAgentToolsResponseBody
	GetGroups() []*DescribePolarClawAgentToolsResponseBodyGroups
	SetMessage(v string) *DescribePolarClawAgentToolsResponseBody
	GetMessage() *string
	SetProfiles(v []*DescribePolarClawAgentToolsResponseBodyProfiles) *DescribePolarClawAgentToolsResponseBody
	GetProfiles() []*DescribePolarClawAgentToolsResponseBodyProfiles
	SetRequestId(v string) *DescribePolarClawAgentToolsResponseBody
	GetRequestId() *string
}

type DescribePolarClawAgentToolsResponseBody struct {
	// Agent ID
	//
	// example:
	//
	// main
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The application ID.
	//
	// example:
	//
	// pa-xxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The current tool configuration.
	CurrentConfig *DescribePolarClawAgentToolsResponseBodyCurrentConfig `json:"CurrentConfig,omitempty" xml:"CurrentConfig,omitempty" type:"Struct"`
	// The list of tool groups.
	Groups []*DescribePolarClawAgentToolsResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of available profiles.
	Profiles []*DescribePolarClawAgentToolsResponseBodyProfiles `json:"Profiles,omitempty" xml:"Profiles,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// F45FFACC-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePolarClawAgentToolsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawAgentToolsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarClawAgentToolsResponseBody) GetAgentId() *string {
	return s.AgentId
}

func (s *DescribePolarClawAgentToolsResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribePolarClawAgentToolsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribePolarClawAgentToolsResponseBody) GetCurrentConfig() *DescribePolarClawAgentToolsResponseBodyCurrentConfig {
	return s.CurrentConfig
}

func (s *DescribePolarClawAgentToolsResponseBody) GetGroups() []*DescribePolarClawAgentToolsResponseBodyGroups {
	return s.Groups
}

func (s *DescribePolarClawAgentToolsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePolarClawAgentToolsResponseBody) GetProfiles() []*DescribePolarClawAgentToolsResponseBodyProfiles {
	return s.Profiles
}

func (s *DescribePolarClawAgentToolsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolarClawAgentToolsResponseBody) SetAgentId(v string) *DescribePolarClawAgentToolsResponseBody {
	s.AgentId = &v
	return s
}

func (s *DescribePolarClawAgentToolsResponseBody) SetApplicationId(v string) *DescribePolarClawAgentToolsResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *DescribePolarClawAgentToolsResponseBody) SetCode(v int32) *DescribePolarClawAgentToolsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePolarClawAgentToolsResponseBody) SetCurrentConfig(v *DescribePolarClawAgentToolsResponseBodyCurrentConfig) *DescribePolarClawAgentToolsResponseBody {
	s.CurrentConfig = v
	return s
}

func (s *DescribePolarClawAgentToolsResponseBody) SetGroups(v []*DescribePolarClawAgentToolsResponseBodyGroups) *DescribePolarClawAgentToolsResponseBody {
	s.Groups = v
	return s
}

func (s *DescribePolarClawAgentToolsResponseBody) SetMessage(v string) *DescribePolarClawAgentToolsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePolarClawAgentToolsResponseBody) SetProfiles(v []*DescribePolarClawAgentToolsResponseBodyProfiles) *DescribePolarClawAgentToolsResponseBody {
	s.Profiles = v
	return s
}

func (s *DescribePolarClawAgentToolsResponseBody) SetRequestId(v string) *DescribePolarClawAgentToolsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarClawAgentToolsResponseBody) Validate() error {
	if s.CurrentConfig != nil {
		if err := s.CurrentConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Groups != nil {
		for _, item := range s.Groups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Profiles != nil {
		for _, item := range s.Profiles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolarClawAgentToolsResponseBodyCurrentConfig struct {
	// The list of explicitly allowed tools.
	Allow []*string `json:"Allow,omitempty" xml:"Allow,omitempty" type:"Repeated"`
	// The list of additionally allowed tools.
	AlsoAllow []*string `json:"AlsoAllow,omitempty" xml:"AlsoAllow,omitempty" type:"Repeated"`
	// The list of denied tools.
	Deny []*string `json:"Deny,omitempty" xml:"Deny,omitempty" type:"Repeated"`
	// The tool profile.
	//
	// example:
	//
	// full
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
}

func (s DescribePolarClawAgentToolsResponseBodyCurrentConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawAgentToolsResponseBodyCurrentConfig) GoString() string {
	return s.String()
}

func (s *DescribePolarClawAgentToolsResponseBodyCurrentConfig) GetAllow() []*string {
	return s.Allow
}

func (s *DescribePolarClawAgentToolsResponseBodyCurrentConfig) GetAlsoAllow() []*string {
	return s.AlsoAllow
}

func (s *DescribePolarClawAgentToolsResponseBodyCurrentConfig) GetDeny() []*string {
	return s.Deny
}

func (s *DescribePolarClawAgentToolsResponseBodyCurrentConfig) GetProfile() *string {
	return s.Profile
}

func (s *DescribePolarClawAgentToolsResponseBodyCurrentConfig) SetAllow(v []*string) *DescribePolarClawAgentToolsResponseBodyCurrentConfig {
	s.Allow = v
	return s
}

func (s *DescribePolarClawAgentToolsResponseBodyCurrentConfig) SetAlsoAllow(v []*string) *DescribePolarClawAgentToolsResponseBodyCurrentConfig {
	s.AlsoAllow = v
	return s
}

func (s *DescribePolarClawAgentToolsResponseBodyCurrentConfig) SetDeny(v []*string) *DescribePolarClawAgentToolsResponseBodyCurrentConfig {
	s.Deny = v
	return s
}

func (s *DescribePolarClawAgentToolsResponseBodyCurrentConfig) SetProfile(v string) *DescribePolarClawAgentToolsResponseBodyCurrentConfig {
	s.Profile = &v
	return s
}

func (s *DescribePolarClawAgentToolsResponseBodyCurrentConfig) Validate() error {
	return dara.Validate(s)
}

type DescribePolarClawAgentToolsResponseBodyGroups struct {
	// The group identifier.
	//
	// example:
	//
	// file
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The group name.
	//
	// example:
	//
	// 文件操作
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The source, which is core or a plugin ID.
	//
	// example:
	//
	// core
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The list of tools.
	Tools []*DescribePolarClawAgentToolsResponseBodyGroupsTools `json:"Tools,omitempty" xml:"Tools,omitempty" type:"Repeated"`
}

func (s DescribePolarClawAgentToolsResponseBodyGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawAgentToolsResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *DescribePolarClawAgentToolsResponseBodyGroups) GetId() *string {
	return s.Id
}

func (s *DescribePolarClawAgentToolsResponseBodyGroups) GetLabel() *string {
	return s.Label
}

func (s *DescribePolarClawAgentToolsResponseBodyGroups) GetSource() *string {
	return s.Source
}

func (s *DescribePolarClawAgentToolsResponseBodyGroups) GetTools() []*DescribePolarClawAgentToolsResponseBodyGroupsTools {
	return s.Tools
}

func (s *DescribePolarClawAgentToolsResponseBodyGroups) SetId(v string) *DescribePolarClawAgentToolsResponseBodyGroups {
	s.Id = &v
	return s
}

func (s *DescribePolarClawAgentToolsResponseBodyGroups) SetLabel(v string) *DescribePolarClawAgentToolsResponseBodyGroups {
	s.Label = &v
	return s
}

func (s *DescribePolarClawAgentToolsResponseBodyGroups) SetSource(v string) *DescribePolarClawAgentToolsResponseBodyGroups {
	s.Source = &v
	return s
}

func (s *DescribePolarClawAgentToolsResponseBodyGroups) SetTools(v []*DescribePolarClawAgentToolsResponseBodyGroupsTools) *DescribePolarClawAgentToolsResponseBodyGroups {
	s.Tools = v
	return s
}

func (s *DescribePolarClawAgentToolsResponseBodyGroups) Validate() error {
	if s.Tools != nil {
		for _, item := range s.Tools {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolarClawAgentToolsResponseBodyGroupsTools struct {
	// The list of profiles that include this tool by default.
	DefaultProfiles []*string `json:"DefaultProfiles,omitempty" xml:"DefaultProfiles,omitempty" type:"Repeated"`
	// The tool description.
	//
	// example:
	//
	// 读取指定文件内容
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The tool identifier.
	//
	// example:
	//
	// read
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The tool name.
	//
	// example:
	//
	// 读取文件
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The source.
	//
	// example:
	//
	// core
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribePolarClawAgentToolsResponseBodyGroupsTools) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawAgentToolsResponseBodyGroupsTools) GoString() string {
	return s.String()
}

func (s *DescribePolarClawAgentToolsResponseBodyGroupsTools) GetDefaultProfiles() []*string {
	return s.DefaultProfiles
}

func (s *DescribePolarClawAgentToolsResponseBodyGroupsTools) GetDescription() *string {
	return s.Description
}

func (s *DescribePolarClawAgentToolsResponseBodyGroupsTools) GetId() *string {
	return s.Id
}

func (s *DescribePolarClawAgentToolsResponseBodyGroupsTools) GetLabel() *string {
	return s.Label
}

func (s *DescribePolarClawAgentToolsResponseBodyGroupsTools) GetSource() *string {
	return s.Source
}

func (s *DescribePolarClawAgentToolsResponseBodyGroupsTools) SetDefaultProfiles(v []*string) *DescribePolarClawAgentToolsResponseBodyGroupsTools {
	s.DefaultProfiles = v
	return s
}

func (s *DescribePolarClawAgentToolsResponseBodyGroupsTools) SetDescription(v string) *DescribePolarClawAgentToolsResponseBodyGroupsTools {
	s.Description = &v
	return s
}

func (s *DescribePolarClawAgentToolsResponseBodyGroupsTools) SetId(v string) *DescribePolarClawAgentToolsResponseBodyGroupsTools {
	s.Id = &v
	return s
}

func (s *DescribePolarClawAgentToolsResponseBodyGroupsTools) SetLabel(v string) *DescribePolarClawAgentToolsResponseBodyGroupsTools {
	s.Label = &v
	return s
}

func (s *DescribePolarClawAgentToolsResponseBodyGroupsTools) SetSource(v string) *DescribePolarClawAgentToolsResponseBodyGroupsTools {
	s.Source = &v
	return s
}

func (s *DescribePolarClawAgentToolsResponseBodyGroupsTools) Validate() error {
	return dara.Validate(s)
}

type DescribePolarClawAgentToolsResponseBodyProfiles struct {
	// The profile identifier.
	//
	// example:
	//
	// full
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The display name.
	//
	// example:
	//
	// 全部工具
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s DescribePolarClawAgentToolsResponseBodyProfiles) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawAgentToolsResponseBodyProfiles) GoString() string {
	return s.String()
}

func (s *DescribePolarClawAgentToolsResponseBodyProfiles) GetId() *string {
	return s.Id
}

func (s *DescribePolarClawAgentToolsResponseBodyProfiles) GetLabel() *string {
	return s.Label
}

func (s *DescribePolarClawAgentToolsResponseBodyProfiles) SetId(v string) *DescribePolarClawAgentToolsResponseBodyProfiles {
	s.Id = &v
	return s
}

func (s *DescribePolarClawAgentToolsResponseBodyProfiles) SetLabel(v string) *DescribePolarClawAgentToolsResponseBodyProfiles {
	s.Label = &v
	return s
}

func (s *DescribePolarClawAgentToolsResponseBodyProfiles) Validate() error {
	return dara.Validate(s)
}
