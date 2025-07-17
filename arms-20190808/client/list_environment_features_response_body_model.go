// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentFeaturesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListEnvironmentFeaturesResponseBody
	GetCode() *int32
	SetData(v []*ListEnvironmentFeaturesResponseBodyData) *ListEnvironmentFeaturesResponseBody
	GetData() []*ListEnvironmentFeaturesResponseBodyData
	SetMessage(v string) *ListEnvironmentFeaturesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListEnvironmentFeaturesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListEnvironmentFeaturesResponseBody
	GetSuccess() *bool
}

type ListEnvironmentFeaturesResponseBody struct {
	// Status Code. Description 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned struct.
	Data []*ListEnvironmentFeaturesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 40B10E04-81E8-4643-970D-F1B38F2E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the alert rule was deleted. Valid values:
	//
	// 	- `true`: The alert rule was deleted.
	//
	// 	- `false`: The alert rule failed to be deleted.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListEnvironmentFeaturesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentFeaturesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnvironmentFeaturesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListEnvironmentFeaturesResponseBody) GetData() []*ListEnvironmentFeaturesResponseBodyData {
	return s.Data
}

func (s *ListEnvironmentFeaturesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEnvironmentFeaturesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEnvironmentFeaturesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListEnvironmentFeaturesResponseBody) SetCode(v int32) *ListEnvironmentFeaturesResponseBody {
	s.Code = &v
	return s
}

func (s *ListEnvironmentFeaturesResponseBody) SetData(v []*ListEnvironmentFeaturesResponseBodyData) *ListEnvironmentFeaturesResponseBody {
	s.Data = v
	return s
}

func (s *ListEnvironmentFeaturesResponseBody) SetMessage(v string) *ListEnvironmentFeaturesResponseBody {
	s.Message = &v
	return s
}

func (s *ListEnvironmentFeaturesResponseBody) SetRequestId(v string) *ListEnvironmentFeaturesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnvironmentFeaturesResponseBody) SetSuccess(v bool) *ListEnvironmentFeaturesResponseBody {
	s.Success = &v
	return s
}

func (s *ListEnvironmentFeaturesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListEnvironmentFeaturesResponseBodyData struct {
	// The alias of the feature.
	//
	// example:
	//
	// Prometheus Agent
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The feature configuration.
	Config map[string]*string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The description of the feature.
	//
	// example:
	//
	// Collect Metric data using the Prometheus collection specification.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the environment instance.
	//
	// example:
	//
	// env-xxxxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The URL of the icon.
	//
	// example:
	//
	// http://xxx
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// The language. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The latest version number.
	//
	// example:
	//
	// 1.1.17
	LatestVersion *string `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	// Indicates whether the component is fully managed.
	//
	// example:
	//
	// true
	Managed *bool `json:"Managed,omitempty" xml:"Managed,omitempty"`
	// The name of the feature.
	//
	// example:
	//
	// metric-agent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the feature. Valid values:
	//
	// 	- Installing: The agent is being installed.
	//
	// 	- Success: The agent is installed.
	//
	// 	- Failed: The agent failed to be installed.
	//
	// 	- UnInstall: The agent is uninstalled.
	//
	// 	- Uninstalling: The agent is being uninstalled.
	//
	// 	- UnInstallFailed: The agent failed to be uninstalled.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The version of the feature.
	//
	// example:
	//
	// 1.1.17
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListEnvironmentFeaturesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentFeaturesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnvironmentFeaturesResponseBodyData) GetAlias() *string {
	return s.Alias
}

func (s *ListEnvironmentFeaturesResponseBodyData) GetConfig() map[string]*string {
	return s.Config
}

func (s *ListEnvironmentFeaturesResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListEnvironmentFeaturesResponseBodyData) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ListEnvironmentFeaturesResponseBodyData) GetIcon() *string {
	return s.Icon
}

func (s *ListEnvironmentFeaturesResponseBodyData) GetLanguage() *string {
	return s.Language
}

func (s *ListEnvironmentFeaturesResponseBodyData) GetLatestVersion() *string {
	return s.LatestVersion
}

func (s *ListEnvironmentFeaturesResponseBodyData) GetManaged() *bool {
	return s.Managed
}

func (s *ListEnvironmentFeaturesResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListEnvironmentFeaturesResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListEnvironmentFeaturesResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *ListEnvironmentFeaturesResponseBodyData) SetAlias(v string) *ListEnvironmentFeaturesResponseBodyData {
	s.Alias = &v
	return s
}

func (s *ListEnvironmentFeaturesResponseBodyData) SetConfig(v map[string]*string) *ListEnvironmentFeaturesResponseBodyData {
	s.Config = v
	return s
}

func (s *ListEnvironmentFeaturesResponseBodyData) SetDescription(v string) *ListEnvironmentFeaturesResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListEnvironmentFeaturesResponseBodyData) SetEnvironmentId(v string) *ListEnvironmentFeaturesResponseBodyData {
	s.EnvironmentId = &v
	return s
}

func (s *ListEnvironmentFeaturesResponseBodyData) SetIcon(v string) *ListEnvironmentFeaturesResponseBodyData {
	s.Icon = &v
	return s
}

func (s *ListEnvironmentFeaturesResponseBodyData) SetLanguage(v string) *ListEnvironmentFeaturesResponseBodyData {
	s.Language = &v
	return s
}

func (s *ListEnvironmentFeaturesResponseBodyData) SetLatestVersion(v string) *ListEnvironmentFeaturesResponseBodyData {
	s.LatestVersion = &v
	return s
}

func (s *ListEnvironmentFeaturesResponseBodyData) SetManaged(v bool) *ListEnvironmentFeaturesResponseBodyData {
	s.Managed = &v
	return s
}

func (s *ListEnvironmentFeaturesResponseBodyData) SetName(v string) *ListEnvironmentFeaturesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListEnvironmentFeaturesResponseBodyData) SetStatus(v string) *ListEnvironmentFeaturesResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListEnvironmentFeaturesResponseBodyData) SetVersion(v string) *ListEnvironmentFeaturesResponseBodyData {
	s.Version = &v
	return s
}

func (s *ListEnvironmentFeaturesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
