// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListAppConfigsResponseBodyData) *ListAppConfigsResponseBody
	GetData() []*ListAppConfigsResponseBodyData
	SetRequestId(v string) *ListAppConfigsResponseBody
	GetRequestId() *string
}

type ListAppConfigsResponseBody struct {
	// The returned data.
	Data []*ListAppConfigsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID assigned by the backend to uniquely identify a request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAppConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppConfigsResponseBody) GetData() []*ListAppConfigsResponseBodyData {
	return s.Data
}

func (s *ListAppConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppConfigsResponseBody) SetData(v []*ListAppConfigsResponseBodyData) *ListAppConfigsResponseBody {
	s.Data = v
	return s
}

func (s *ListAppConfigsResponseBody) SetRequestId(v string) *ListAppConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppConfigsResponseBody) Validate() error {
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

type ListAppConfigsResponseBodyData struct {
	// App ID。
	//
	// example:
	//
	// txt_check_pro_agent_01
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The classification.
	//
	// example:
	//
	// guard-scene
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// The configuration details.
	//
	// example:
	//
	// {"agentItemConfigs": "[{\\"agentId\\":\\"ag.abcxxx\\",\\"enable\\":true,\\"name\\":\\"Agent 1\\"}]"}
	CustomConfig map[string]interface{} `json:"CustomConfig,omitempty" xml:"CustomConfig,omitempty"`
	// The publish status.
	//
	// example:
	//
	// editing
	DeployStatus *string `json:"DeployStatus,omitempty" xml:"DeployStatus,omitempty"`
	// The last modification time.
	//
	// example:
	//
	// 2026-06-15 10:17:49
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The name.
	//
	// example:
	//
	// Custom text moderation
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The preset options.
	Option map[string]interface{} `json:"Option,omitempty" xml:"Option,omitempty"`
	// The resource type.
	//
	// example:
	//
	// agent_text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The type.
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

func (s ListAppConfigsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAppConfigsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAppConfigsResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *ListAppConfigsResponseBodyData) GetClassify() *string {
	return s.Classify
}

func (s *ListAppConfigsResponseBodyData) GetCustomConfig() map[string]interface{} {
	return s.CustomConfig
}

func (s *ListAppConfigsResponseBodyData) GetDeployStatus() *string {
	return s.DeployStatus
}

func (s *ListAppConfigsResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListAppConfigsResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListAppConfigsResponseBodyData) GetOption() map[string]interface{} {
	return s.Option
}

func (s *ListAppConfigsResponseBodyData) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListAppConfigsResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ListAppConfigsResponseBodyData) GetVersion() *int64 {
	return s.Version
}

func (s *ListAppConfigsResponseBodyData) SetAppId(v string) *ListAppConfigsResponseBodyData {
	s.AppId = &v
	return s
}

func (s *ListAppConfigsResponseBodyData) SetClassify(v string) *ListAppConfigsResponseBodyData {
	s.Classify = &v
	return s
}

func (s *ListAppConfigsResponseBodyData) SetCustomConfig(v map[string]interface{}) *ListAppConfigsResponseBodyData {
	s.CustomConfig = v
	return s
}

func (s *ListAppConfigsResponseBodyData) SetDeployStatus(v string) *ListAppConfigsResponseBodyData {
	s.DeployStatus = &v
	return s
}

func (s *ListAppConfigsResponseBodyData) SetGmtModified(v string) *ListAppConfigsResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListAppConfigsResponseBodyData) SetName(v string) *ListAppConfigsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListAppConfigsResponseBodyData) SetOption(v map[string]interface{}) *ListAppConfigsResponseBodyData {
	s.Option = v
	return s
}

func (s *ListAppConfigsResponseBodyData) SetResourceType(v string) *ListAppConfigsResponseBodyData {
	s.ResourceType = &v
	return s
}

func (s *ListAppConfigsResponseBodyData) SetType(v string) *ListAppConfigsResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListAppConfigsResponseBodyData) SetVersion(v int64) *ListAppConfigsResponseBodyData {
	s.Version = &v
	return s
}

func (s *ListAppConfigsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
