// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateFlowLogRequest
	GetClientToken() *string
	SetDescription(v string) *CreateFlowLogRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateFlowLogRequest
	GetDryRun() *bool
	SetEcrId(v string) *CreateFlowLogRequest
	GetEcrId() *string
	SetFlowLogName(v string) *CreateFlowLogRequest
	GetFlowLogName() *string
	SetInstanceId(v string) *CreateFlowLogRequest
	GetInstanceId() *string
	SetInstanceType(v string) *CreateFlowLogRequest
	GetInstanceType() *string
	SetInterval(v int32) *CreateFlowLogRequest
	GetInterval() *int32
	SetLogStoreName(v string) *CreateFlowLogRequest
	GetLogStoreName() *string
	SetProjectName(v string) *CreateFlowLogRequest
	GetProjectName() *string
	SetResourceGroupId(v string) *CreateFlowLogRequest
	GetResourceGroupId() *string
	SetSamplingRate(v string) *CreateFlowLogRequest
	GetSamplingRate() *string
	SetTag(v []*CreateFlowLogRequestTag) *CreateFlowLogRequest
	GetTag() []*CreateFlowLogRequestTag
	SetTargetSlsRegionId(v string) *CreateFlowLogRequest
	GetTargetSlsRegionId() *string
	SetVersion(v string) *CreateFlowLogRequest
	GetVersion() *string
}

type CreateFlowLogRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate a parameter value on your client to ensure that the value is unique across different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the flow log.
	//
	// > The description can be empty or 0 to 256 characters in length. It cannot start with http:// or https://.
	//
	// example:
	//
	// myFlowlog
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether the request is only a dry run. Valid values:
	//
	// - **true**: Performs a dry run without modifying EPS resources.
	//
	// - **false*	- (default): Sends a normal request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The instance ID of the Express Connect Router (ECR).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The name of the flow log.
	//
	// > The name can be empty or 0 to 128 characters in length. It cannot start with http:// or https://.
	//
	// example:
	//
	// myFlowlog
	FlowLogName *string `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	// The instance ID of the VBR.
	//
	// >The VBR instance must be created by calling the CreateVirtualBorderRouter operation of the Vpc/2016-04-28 API version. The prerequisite is that a physical connection exists.
	//
	// >The VBR instance must be mounted as a child instance to the Express Connect Router (ECR) corresponding to EcrId by calling AttachExpressConnectRouterChildInstance, and the mount status must be ACTIVE before you can create a flow log. You can call DescribeExpressConnectRouterChildInstance to query the mount status of the child instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the network instance. Valid values:
	//
	// - **VBR**: virtual border router (VBR) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// VBR
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The duration of the capture window for traffic captured by the flow log. Unit: seconds. Valid values:
	//
	// - **60*	-
	//
	// - **600**
	//
	// Default value: **600**.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the Logstore used to store catch traffic.
	//
	// - If you have already created a Logstore in the current region, you can enter the name of the existing Logstore.
	//
	// - If you have not created a Logstore in the current region, you can specify a custom name for the Logstore, and the system will automatically create the Logstore. The naming rules for the Logstore name are as follows:
	//
	// - The Logstore name must be unique within the same project.
	//
	// - It can contain only lowercase letters, digits, hyphens (-), and underscores (_).
	//
	// - It must start and end with a lowercase letter or a digit.
	//
	// - It must be 3 to 63 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// flowlog-logstore
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The name of the Simple Log Service project used to store catch traffic.
	//
	// - If you have already created a project in the current region, you can enter the name of the existing project.
	//
	// - If you have not created a project in the current region, you can specify a custom name for the project, and the system will automatically create the project.
	//
	//
	// The project name must be globally unique within an Alibaba Cloud region and cannot be modified after creation. The naming rules are as follows:
	//
	//
	//
	// - The project name must be globally unique.
	//
	// - It can contain only lowercase
	//
	// - letters, digits, and hyphens (-).
	//
	// - It must start and end with a lowercase letter or a digit.
	//
	// - It must be 3 to 63 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// flowlog-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The ID of the resource group to which the flow log instance belongs.
	//
	// example:
	//
	// rg-aek2aq7f4va****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The sampling ratio. Valid values:
	//
	// - **1:4096**
	//
	// - **1:2048**
	//
	// - **1:1024**
	//
	// Default value: **1:4096**.
	//
	// example:
	//
	// 1:4096
	SamplingRate *string `json:"SamplingRate,omitempty" xml:"SamplingRate,omitempty"`
	// The tags.
	//
	// You can add up to 20 tags.
	Tag []*CreateFlowLogRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The region where the project to which logs are delivered resides. You only need to specify this parameter for cross-region log delivery. You do not need to specify this parameter for same-region log delivery.
	//
	// example:
	//
	// cn-hangzhou
	TargetSlsRegionId *string `json:"TargetSlsRegionId,omitempty" xml:"TargetSlsRegionId,omitempty"`
	Version           *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateFlowLogRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowLogRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowLogRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateFlowLogRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateFlowLogRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateFlowLogRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *CreateFlowLogRequest) GetFlowLogName() *string {
	return s.FlowLogName
}

func (s *CreateFlowLogRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateFlowLogRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateFlowLogRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *CreateFlowLogRequest) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *CreateFlowLogRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateFlowLogRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateFlowLogRequest) GetSamplingRate() *string {
	return s.SamplingRate
}

func (s *CreateFlowLogRequest) GetTag() []*CreateFlowLogRequestTag {
	return s.Tag
}

func (s *CreateFlowLogRequest) GetTargetSlsRegionId() *string {
	return s.TargetSlsRegionId
}

func (s *CreateFlowLogRequest) GetVersion() *string {
	return s.Version
}

func (s *CreateFlowLogRequest) SetClientToken(v string) *CreateFlowLogRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFlowLogRequest) SetDescription(v string) *CreateFlowLogRequest {
	s.Description = &v
	return s
}

func (s *CreateFlowLogRequest) SetDryRun(v bool) *CreateFlowLogRequest {
	s.DryRun = &v
	return s
}

func (s *CreateFlowLogRequest) SetEcrId(v string) *CreateFlowLogRequest {
	s.EcrId = &v
	return s
}

func (s *CreateFlowLogRequest) SetFlowLogName(v string) *CreateFlowLogRequest {
	s.FlowLogName = &v
	return s
}

func (s *CreateFlowLogRequest) SetInstanceId(v string) *CreateFlowLogRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateFlowLogRequest) SetInstanceType(v string) *CreateFlowLogRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateFlowLogRequest) SetInterval(v int32) *CreateFlowLogRequest {
	s.Interval = &v
	return s
}

func (s *CreateFlowLogRequest) SetLogStoreName(v string) *CreateFlowLogRequest {
	s.LogStoreName = &v
	return s
}

func (s *CreateFlowLogRequest) SetProjectName(v string) *CreateFlowLogRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateFlowLogRequest) SetResourceGroupId(v string) *CreateFlowLogRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateFlowLogRequest) SetSamplingRate(v string) *CreateFlowLogRequest {
	s.SamplingRate = &v
	return s
}

func (s *CreateFlowLogRequest) SetTag(v []*CreateFlowLogRequestTag) *CreateFlowLogRequest {
	s.Tag = v
	return s
}

func (s *CreateFlowLogRequest) SetTargetSlsRegionId(v string) *CreateFlowLogRequest {
	s.TargetSlsRegionId = &v
	return s
}

func (s *CreateFlowLogRequest) SetVersion(v string) *CreateFlowLogRequest {
	s.Version = &v
	return s
}

func (s *CreateFlowLogRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateFlowLogRequestTag struct {
	// The tag key of the resource. If you specify this parameter, it cannot be an empty string. It can be up to 128 characters in length, and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. Valid values of N: 1 to 20. If you specify this parameter, it can be an empty string. It can be up to 128 characters in length, and cannot start with acs:. It cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateFlowLogRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowLogRequestTag) GoString() string {
	return s.String()
}

func (s *CreateFlowLogRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateFlowLogRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateFlowLogRequestTag) SetKey(v string) *CreateFlowLogRequestTag {
	s.Key = &v
	return s
}

func (s *CreateFlowLogRequestTag) SetValue(v string) *CreateFlowLogRequestTag {
	s.Value = &v
	return s
}

func (s *CreateFlowLogRequestTag) Validate() error {
	return dara.Validate(s)
}
