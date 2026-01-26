// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnvironmentFeatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeEnvironmentFeatureResponseBody
	GetCode() *int32
	SetData(v *DescribeEnvironmentFeatureResponseBodyData) *DescribeEnvironmentFeatureResponseBody
	GetData() *DescribeEnvironmentFeatureResponseBodyData
	SetMessage(v string) *DescribeEnvironmentFeatureResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeEnvironmentFeatureResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeEnvironmentFeatureResponseBody
	GetSuccess() *bool
}

type DescribeEnvironmentFeatureResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The struct returned.
	Data *DescribeEnvironmentFeatureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 01FF8DD9-A09C-47A1-895A-B6E321BE77B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeEnvironmentFeatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnvironmentFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnvironmentFeatureResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeEnvironmentFeatureResponseBody) GetData() *DescribeEnvironmentFeatureResponseBodyData {
	return s.Data
}

func (s *DescribeEnvironmentFeatureResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeEnvironmentFeatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEnvironmentFeatureResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeEnvironmentFeatureResponseBody) SetCode(v int32) *DescribeEnvironmentFeatureResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBody) SetData(v *DescribeEnvironmentFeatureResponseBodyData) *DescribeEnvironmentFeatureResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBody) SetMessage(v string) *DescribeEnvironmentFeatureResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBody) SetRequestId(v string) *DescribeEnvironmentFeatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBody) SetSuccess(v bool) *DescribeEnvironmentFeatureResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEnvironmentFeatureResponseBodyData struct {
	// The installation information about the feature.
	Feature *DescribeEnvironmentFeatureResponseBodyDataFeature `json:"Feature,omitempty" xml:"Feature,omitempty" type:"Struct"`
	// The status of the feature.
	FeatureStatus *DescribeEnvironmentFeatureResponseBodyDataFeatureStatus `json:"FeatureStatus,omitempty" xml:"FeatureStatus,omitempty" type:"Struct"`
	// The feature configurations.
	//
	// example:
	//
	// {}
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
}

func (s DescribeEnvironmentFeatureResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnvironmentFeatureResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEnvironmentFeatureResponseBodyData) GetFeature() *DescribeEnvironmentFeatureResponseBodyDataFeature {
	return s.Feature
}

func (s *DescribeEnvironmentFeatureResponseBodyData) GetFeatureStatus() *DescribeEnvironmentFeatureResponseBodyDataFeatureStatus {
	return s.FeatureStatus
}

func (s *DescribeEnvironmentFeatureResponseBodyData) GetConfig() *string {
	return s.Config
}

func (s *DescribeEnvironmentFeatureResponseBodyData) SetFeature(v *DescribeEnvironmentFeatureResponseBodyDataFeature) *DescribeEnvironmentFeatureResponseBodyData {
	s.Feature = v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBodyData) SetFeatureStatus(v *DescribeEnvironmentFeatureResponseBodyDataFeatureStatus) *DescribeEnvironmentFeatureResponseBodyData {
	s.FeatureStatus = v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBodyData) SetConfig(v string) *DescribeEnvironmentFeatureResponseBodyData {
	s.Config = &v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBodyData) Validate() error {
	if s.Feature != nil {
		if err := s.Feature.Validate(); err != nil {
			return err
		}
	}
	if s.FeatureStatus != nil {
		if err := s.FeatureStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEnvironmentFeatureResponseBodyDataFeature struct {
	// The alias of the feature.
	//
	// example:
	//
	// Prometheus agent.
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The configuration of the feature.
	Config map[string]*string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The description of the feature.
	//
	// example:
	//
	// Collect Metric data using the Prometheus collection specification.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The environment ID.
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
	// The language.
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
	Managed *bool `json:"Managed,omitempty" xml:"Managed,omitempty"`
	// The name of the feature.
	//
	// example:
	//
	// metric-agent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The installation status of the agent.
	//
	// 	- Installing: The agent is being installed.
	//
	// 	- Success: The agent is installed.
	//
	// 	- Failed: The agent failed to be installed.
	//
	// 	- UnInstall: The agent is uninstalled or has not been installed.
	//
	// 	- Uninstalling: The agent is being uninstalled.
	//
	// 	- UnInstallFailed: The agent failed to be uninstalled.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1.1.17
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeEnvironmentFeatureResponseBodyDataFeature) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnvironmentFeatureResponseBodyDataFeature) GoString() string {
	return s.String()
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeature) GetAlias() *string {
	return s.Alias
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeature) GetConfig() map[string]*string {
	return s.Config
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeature) GetDescription() *string {
	return s.Description
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeature) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeature) GetIcon() *string {
	return s.Icon
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeature) GetLanguage() *string {
	return s.Language
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeature) GetLatestVersion() *string {
	return s.LatestVersion
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeature) GetManaged() *bool {
	return s.Managed
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeature) GetName() *string {
	return s.Name
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeature) GetStatus() *string {
	return s.Status
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeature) GetVersion() *string {
	return s.Version
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeature) SetAlias(v string) *DescribeEnvironmentFeatureResponseBodyDataFeature {
	s.Alias = &v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeature) SetConfig(v map[string]*string) *DescribeEnvironmentFeatureResponseBodyDataFeature {
	s.Config = v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeature) SetDescription(v string) *DescribeEnvironmentFeatureResponseBodyDataFeature {
	s.Description = &v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeature) SetEnvironmentId(v string) *DescribeEnvironmentFeatureResponseBodyDataFeature {
	s.EnvironmentId = &v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeature) SetIcon(v string) *DescribeEnvironmentFeatureResponseBodyDataFeature {
	s.Icon = &v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeature) SetLanguage(v string) *DescribeEnvironmentFeatureResponseBodyDataFeature {
	s.Language = &v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeature) SetLatestVersion(v string) *DescribeEnvironmentFeatureResponseBodyDataFeature {
	s.LatestVersion = &v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeature) SetManaged(v bool) *DescribeEnvironmentFeatureResponseBodyDataFeature {
	s.Managed = &v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeature) SetName(v string) *DescribeEnvironmentFeatureResponseBodyDataFeature {
	s.Name = &v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeature) SetStatus(v string) *DescribeEnvironmentFeatureResponseBodyDataFeature {
	s.Status = &v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeature) SetVersion(v string) *DescribeEnvironmentFeatureResponseBodyDataFeature {
	s.Version = &v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeature) Validate() error {
	return dara.Validate(s)
}

type DescribeEnvironmentFeatureResponseBodyDataFeatureStatus struct {
	// The ID of the resource.
	//
	// example:
	//
	// c013823b55e4b4d6bb6b6f28682bd38a7
	BindResourceId *string `json:"BindResourceId,omitempty" xml:"BindResourceId,omitempty"`
	// The containers of the feature.
	FeatureContainers []*DescribeEnvironmentFeatureResponseBodyDataFeatureStatusFeatureContainers `json:"FeatureContainers,omitempty" xml:"FeatureContainers,omitempty" type:"Repeated"`
	// The IP address of the pod.
	Ips []*string `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
	// The Kubernetes resource name of the feature.
	//
	// example:
	//
	// arms-prometheus-ack-arms-prometheus
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace.
	//
	// example:
	//
	// arms-prom
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-bp1c9fcexoalq9po6cp8
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The status of the agent. Valid values:
	//
	// 	- Success: The agent is running.
	//
	// 	- Failed: The agent failed to run.
	//
	// 	- Not Found: The agent is not installed.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-bp1qt6ict0dbxgv4wer8l
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeEnvironmentFeatureResponseBodyDataFeatureStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnvironmentFeatureResponseBodyDataFeatureStatus) GoString() string {
	return s.String()
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeatureStatus) GetBindResourceId() *string {
	return s.BindResourceId
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeatureStatus) GetFeatureContainers() []*DescribeEnvironmentFeatureResponseBodyDataFeatureStatusFeatureContainers {
	return s.FeatureContainers
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeatureStatus) GetIps() []*string {
	return s.Ips
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeatureStatus) GetName() *string {
	return s.Name
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeatureStatus) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeatureStatus) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeatureStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeatureStatus) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeatureStatus) SetBindResourceId(v string) *DescribeEnvironmentFeatureResponseBodyDataFeatureStatus {
	s.BindResourceId = &v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeatureStatus) SetFeatureContainers(v []*DescribeEnvironmentFeatureResponseBodyDataFeatureStatusFeatureContainers) *DescribeEnvironmentFeatureResponseBodyDataFeatureStatus {
	s.FeatureContainers = v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeatureStatus) SetIps(v []*string) *DescribeEnvironmentFeatureResponseBodyDataFeatureStatus {
	s.Ips = v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeatureStatus) SetName(v string) *DescribeEnvironmentFeatureResponseBodyDataFeatureStatus {
	s.Name = &v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeatureStatus) SetNamespace(v string) *DescribeEnvironmentFeatureResponseBodyDataFeatureStatus {
	s.Namespace = &v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeatureStatus) SetSecurityGroupId(v string) *DescribeEnvironmentFeatureResponseBodyDataFeatureStatus {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeatureStatus) SetStatus(v string) *DescribeEnvironmentFeatureResponseBodyDataFeatureStatus {
	s.Status = &v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeatureStatus) SetVSwitchId(v string) *DescribeEnvironmentFeatureResponseBodyDataFeatureStatus {
	s.VSwitchId = &v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeatureStatus) Validate() error {
	if s.FeatureContainers != nil {
		for _, item := range s.FeatureContainers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEnvironmentFeatureResponseBodyDataFeatureStatusFeatureContainers struct {
	// The container parameters.
	Args []*string `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	// The container image.
	//
	// example:
	//
	// registry-cn-hangzhou-vpc.ack.aliyuncs.com/acs/arms-prometheus-agent:v4.0.0
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The container name.
	//
	// example:
	//
	// arms-prometheus-operator
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeEnvironmentFeatureResponseBodyDataFeatureStatusFeatureContainers) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnvironmentFeatureResponseBodyDataFeatureStatusFeatureContainers) GoString() string {
	return s.String()
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeatureStatusFeatureContainers) GetArgs() []*string {
	return s.Args
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeatureStatusFeatureContainers) GetImage() *string {
	return s.Image
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeatureStatusFeatureContainers) GetName() *string {
	return s.Name
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeatureStatusFeatureContainers) SetArgs(v []*string) *DescribeEnvironmentFeatureResponseBodyDataFeatureStatusFeatureContainers {
	s.Args = v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeatureStatusFeatureContainers) SetImage(v string) *DescribeEnvironmentFeatureResponseBodyDataFeatureStatusFeatureContainers {
	s.Image = &v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeatureStatusFeatureContainers) SetName(v string) *DescribeEnvironmentFeatureResponseBodyDataFeatureStatusFeatureContainers {
	s.Name = &v
	return s
}

func (s *DescribeEnvironmentFeatureResponseBodyDataFeatureStatusFeatureContainers) Validate() error {
	return dara.Validate(s)
}
