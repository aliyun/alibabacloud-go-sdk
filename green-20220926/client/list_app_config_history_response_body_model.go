// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppConfigHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHistory(v []*ListAppConfigHistoryResponseBodyHistory) *ListAppConfigHistoryResponseBody
	GetHistory() []*ListAppConfigHistoryResponseBodyHistory
	SetPublish(v []*ListAppConfigHistoryResponseBodyPublish) *ListAppConfigHistoryResponseBody
	GetPublish() []*ListAppConfigHistoryResponseBodyPublish
	SetRequestId(v string) *ListAppConfigHistoryResponseBody
	GetRequestId() *string
}

type ListAppConfigHistoryResponseBody struct {
	// The historical versions.
	History []*ListAppConfigHistoryResponseBodyHistory `json:"History,omitempty" xml:"History,omitempty" type:"Repeated"`
	// The published versions.
	Publish []*ListAppConfigHistoryResponseBodyPublish `json:"Publish,omitempty" xml:"Publish,omitempty" type:"Repeated"`
	// The ID assigned by the backend to uniquely identify a request. It can be used to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAppConfigHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppConfigHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppConfigHistoryResponseBody) GetHistory() []*ListAppConfigHistoryResponseBodyHistory {
	return s.History
}

func (s *ListAppConfigHistoryResponseBody) GetPublish() []*ListAppConfigHistoryResponseBodyPublish {
	return s.Publish
}

func (s *ListAppConfigHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppConfigHistoryResponseBody) SetHistory(v []*ListAppConfigHistoryResponseBodyHistory) *ListAppConfigHistoryResponseBody {
	s.History = v
	return s
}

func (s *ListAppConfigHistoryResponseBody) SetPublish(v []*ListAppConfigHistoryResponseBodyPublish) *ListAppConfigHistoryResponseBody {
	s.Publish = v
	return s
}

func (s *ListAppConfigHistoryResponseBody) SetRequestId(v string) *ListAppConfigHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppConfigHistoryResponseBody) Validate() error {
	if s.History != nil {
		for _, item := range s.History {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Publish != nil {
		for _, item := range s.Publish {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAppConfigHistoryResponseBodyHistory struct {
	// App ID。
	//
	// example:
	//
	// txt_check_agent_01
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// **[Deprecated]*	- The categorization.
	//
	// example:
	//
	// guard-scene
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// **[Deprecated]*	- The configuration details.
	//
	// example:
	//
	// {"agentItemConfigs": "[{\\"agentId\\":\\"ag.abcxxx\\",\\"enable\\":true,\\"name\\":\\"Agent 1\\"}]"}
	CustomConfig map[string]interface{} `json:"CustomConfig,omitempty" xml:"CustomConfig,omitempty"`
	// **[Deprecated]*	- The publish status.
	//
	// example:
	//
	// editing
	DeployStatus *string `json:"DeployStatus,omitempty" xml:"DeployStatus,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2026-05-29 10:05:27
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// **[Deprecated]*	- The name.
	//
	// example:
	//
	// Custom text moderation
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// **[Deprecated]*	- The preset options.
	Option map[string]interface{} `json:"Option,omitempty" xml:"Option,omitempty"`
	// **[Deprecated]*	- The resource type.
	//
	// example:
	//
	// agent_text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// **[Deprecated]*	- The type.
	//
	// example:
	//
	// plus
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1785888163
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListAppConfigHistoryResponseBodyHistory) String() string {
	return dara.Prettify(s)
}

func (s ListAppConfigHistoryResponseBodyHistory) GoString() string {
	return s.String()
}

func (s *ListAppConfigHistoryResponseBodyHistory) GetAppId() *string {
	return s.AppId
}

func (s *ListAppConfigHistoryResponseBodyHistory) GetClassify() *string {
	return s.Classify
}

func (s *ListAppConfigHistoryResponseBodyHistory) GetCustomConfig() map[string]interface{} {
	return s.CustomConfig
}

func (s *ListAppConfigHistoryResponseBodyHistory) GetDeployStatus() *string {
	return s.DeployStatus
}

func (s *ListAppConfigHistoryResponseBodyHistory) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListAppConfigHistoryResponseBodyHistory) GetName() *string {
	return s.Name
}

func (s *ListAppConfigHistoryResponseBodyHistory) GetOption() map[string]interface{} {
	return s.Option
}

func (s *ListAppConfigHistoryResponseBodyHistory) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListAppConfigHistoryResponseBodyHistory) GetType() *string {
	return s.Type
}

func (s *ListAppConfigHistoryResponseBodyHistory) GetVersion() *int64 {
	return s.Version
}

func (s *ListAppConfigHistoryResponseBodyHistory) SetAppId(v string) *ListAppConfigHistoryResponseBodyHistory {
	s.AppId = &v
	return s
}

func (s *ListAppConfigHistoryResponseBodyHistory) SetClassify(v string) *ListAppConfigHistoryResponseBodyHistory {
	s.Classify = &v
	return s
}

func (s *ListAppConfigHistoryResponseBodyHistory) SetCustomConfig(v map[string]interface{}) *ListAppConfigHistoryResponseBodyHistory {
	s.CustomConfig = v
	return s
}

func (s *ListAppConfigHistoryResponseBodyHistory) SetDeployStatus(v string) *ListAppConfigHistoryResponseBodyHistory {
	s.DeployStatus = &v
	return s
}

func (s *ListAppConfigHistoryResponseBodyHistory) SetGmtModified(v string) *ListAppConfigHistoryResponseBodyHistory {
	s.GmtModified = &v
	return s
}

func (s *ListAppConfigHistoryResponseBodyHistory) SetName(v string) *ListAppConfigHistoryResponseBodyHistory {
	s.Name = &v
	return s
}

func (s *ListAppConfigHistoryResponseBodyHistory) SetOption(v map[string]interface{}) *ListAppConfigHistoryResponseBodyHistory {
	s.Option = v
	return s
}

func (s *ListAppConfigHistoryResponseBodyHistory) SetResourceType(v string) *ListAppConfigHistoryResponseBodyHistory {
	s.ResourceType = &v
	return s
}

func (s *ListAppConfigHistoryResponseBodyHistory) SetType(v string) *ListAppConfigHistoryResponseBodyHistory {
	s.Type = &v
	return s
}

func (s *ListAppConfigHistoryResponseBodyHistory) SetVersion(v int64) *ListAppConfigHistoryResponseBodyHistory {
	s.Version = &v
	return s
}

func (s *ListAppConfigHistoryResponseBodyHistory) Validate() error {
	return dara.Validate(s)
}

type ListAppConfigHistoryResponseBodyPublish struct {
	// App ID。
	//
	// example:
	//
	// txt_check_agent_01
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// **[Deprecated]*	- The categorization.
	//
	// example:
	//
	// guard-scene
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// **[Deprecated]*	- The configuration details.
	//
	// example:
	//
	// {"agentItemConfigs": "[{\\"agentId\\":\\"ag.abcxxx\\",\\"enable\\":true,\\"name\\":\\"Agent 1\\"}]"}
	CustomConfig map[string]interface{} `json:"CustomConfig,omitempty" xml:"CustomConfig,omitempty"`
	// **[Deprecated]*	- The publish status.
	//
	// example:
	//
	// editing
	DeployStatus *string `json:"DeployStatus,omitempty" xml:"DeployStatus,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2026-06-25 09:52:12
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// **[Deprecated]*	- The name.
	//
	// example:
	//
	// Custom text moderation
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// **[Deprecated]*	- The preset options.
	Option map[string]interface{} `json:"Option,omitempty" xml:"Option,omitempty"`
	// **[Deprecated]*	- The resource type.
	//
	// example:
	//
	// agent_text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// **[Deprecated]*	- The type.
	//
	// example:
	//
	// plus
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1785898163
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListAppConfigHistoryResponseBodyPublish) String() string {
	return dara.Prettify(s)
}

func (s ListAppConfigHistoryResponseBodyPublish) GoString() string {
	return s.String()
}

func (s *ListAppConfigHistoryResponseBodyPublish) GetAppId() *string {
	return s.AppId
}

func (s *ListAppConfigHistoryResponseBodyPublish) GetClassify() *string {
	return s.Classify
}

func (s *ListAppConfigHistoryResponseBodyPublish) GetCustomConfig() map[string]interface{} {
	return s.CustomConfig
}

func (s *ListAppConfigHistoryResponseBodyPublish) GetDeployStatus() *string {
	return s.DeployStatus
}

func (s *ListAppConfigHistoryResponseBodyPublish) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListAppConfigHistoryResponseBodyPublish) GetName() *string {
	return s.Name
}

func (s *ListAppConfigHistoryResponseBodyPublish) GetOption() map[string]interface{} {
	return s.Option
}

func (s *ListAppConfigHistoryResponseBodyPublish) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListAppConfigHistoryResponseBodyPublish) GetType() *string {
	return s.Type
}

func (s *ListAppConfigHistoryResponseBodyPublish) GetVersion() *int64 {
	return s.Version
}

func (s *ListAppConfigHistoryResponseBodyPublish) SetAppId(v string) *ListAppConfigHistoryResponseBodyPublish {
	s.AppId = &v
	return s
}

func (s *ListAppConfigHistoryResponseBodyPublish) SetClassify(v string) *ListAppConfigHistoryResponseBodyPublish {
	s.Classify = &v
	return s
}

func (s *ListAppConfigHistoryResponseBodyPublish) SetCustomConfig(v map[string]interface{}) *ListAppConfigHistoryResponseBodyPublish {
	s.CustomConfig = v
	return s
}

func (s *ListAppConfigHistoryResponseBodyPublish) SetDeployStatus(v string) *ListAppConfigHistoryResponseBodyPublish {
	s.DeployStatus = &v
	return s
}

func (s *ListAppConfigHistoryResponseBodyPublish) SetGmtModified(v string) *ListAppConfigHistoryResponseBodyPublish {
	s.GmtModified = &v
	return s
}

func (s *ListAppConfigHistoryResponseBodyPublish) SetName(v string) *ListAppConfigHistoryResponseBodyPublish {
	s.Name = &v
	return s
}

func (s *ListAppConfigHistoryResponseBodyPublish) SetOption(v map[string]interface{}) *ListAppConfigHistoryResponseBodyPublish {
	s.Option = v
	return s
}

func (s *ListAppConfigHistoryResponseBodyPublish) SetResourceType(v string) *ListAppConfigHistoryResponseBodyPublish {
	s.ResourceType = &v
	return s
}

func (s *ListAppConfigHistoryResponseBodyPublish) SetType(v string) *ListAppConfigHistoryResponseBodyPublish {
	s.Type = &v
	return s
}

func (s *ListAppConfigHistoryResponseBodyPublish) SetVersion(v int64) *ListAppConfigHistoryResponseBodyPublish {
	s.Version = &v
	return s
}

func (s *ListAppConfigHistoryResponseBodyPublish) Validate() error {
	return dara.Validate(s)
}
