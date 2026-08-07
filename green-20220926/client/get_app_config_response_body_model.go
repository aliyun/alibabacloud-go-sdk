// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetAppConfigResponseBody
	GetAppId() *string
	SetClassify(v string) *GetAppConfigResponseBody
	GetClassify() *string
	SetCustomConfig(v map[string]interface{}) *GetAppConfigResponseBody
	GetCustomConfig() map[string]interface{}
	SetDeployStatus(v string) *GetAppConfigResponseBody
	GetDeployStatus() *string
	SetGmtModified(v string) *GetAppConfigResponseBody
	GetGmtModified() *string
	SetName(v string) *GetAppConfigResponseBody
	GetName() *string
	SetOption(v map[string]interface{}) *GetAppConfigResponseBody
	GetOption() map[string]interface{}
	SetRequestId(v string) *GetAppConfigResponseBody
	GetRequestId() *string
	SetResourceType(v string) *GetAppConfigResponseBody
	GetResourceType() *string
	SetType(v string) *GetAppConfigResponseBody
	GetType() *string
	SetVersion(v int64) *GetAppConfigResponseBody
	GetVersion() *int64
}

type GetAppConfigResponseBody struct {
	// App ID。
	//
	// example:
	//
	// txt_check_pro_agent_01
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The category.
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
	// The modification time.
	//
	// example:
	//
	// 2026-06-09 10:12:50
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The application name.
	//
	// example:
	//
	// Custom text moderation
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The preset options.
	Option map[string]interface{} `json:"Option,omitempty" xml:"Option,omitempty"`
	// The ID assigned by the backend to uniquely identify a request. This ID can be used to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s GetAppConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppConfigResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *GetAppConfigResponseBody) GetClassify() *string {
	return s.Classify
}

func (s *GetAppConfigResponseBody) GetCustomConfig() map[string]interface{} {
	return s.CustomConfig
}

func (s *GetAppConfigResponseBody) GetDeployStatus() *string {
	return s.DeployStatus
}

func (s *GetAppConfigResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetAppConfigResponseBody) GetName() *string {
	return s.Name
}

func (s *GetAppConfigResponseBody) GetOption() map[string]interface{} {
	return s.Option
}

func (s *GetAppConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppConfigResponseBody) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetAppConfigResponseBody) GetType() *string {
	return s.Type
}

func (s *GetAppConfigResponseBody) GetVersion() *int64 {
	return s.Version
}

func (s *GetAppConfigResponseBody) SetAppId(v string) *GetAppConfigResponseBody {
	s.AppId = &v
	return s
}

func (s *GetAppConfigResponseBody) SetClassify(v string) *GetAppConfigResponseBody {
	s.Classify = &v
	return s
}

func (s *GetAppConfigResponseBody) SetCustomConfig(v map[string]interface{}) *GetAppConfigResponseBody {
	s.CustomConfig = v
	return s
}

func (s *GetAppConfigResponseBody) SetDeployStatus(v string) *GetAppConfigResponseBody {
	s.DeployStatus = &v
	return s
}

func (s *GetAppConfigResponseBody) SetGmtModified(v string) *GetAppConfigResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetAppConfigResponseBody) SetName(v string) *GetAppConfigResponseBody {
	s.Name = &v
	return s
}

func (s *GetAppConfigResponseBody) SetOption(v map[string]interface{}) *GetAppConfigResponseBody {
	s.Option = v
	return s
}

func (s *GetAppConfigResponseBody) SetRequestId(v string) *GetAppConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppConfigResponseBody) SetResourceType(v string) *GetAppConfigResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetAppConfigResponseBody) SetType(v string) *GetAppConfigResponseBody {
	s.Type = &v
	return s
}

func (s *GetAppConfigResponseBody) SetVersion(v int64) *GetAppConfigResponseBody {
	s.Version = &v
	return s
}

func (s *GetAppConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
