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
}

type CreateFlowLogRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the flow log.
	//
	// > The description can be empty or 1 to 256 characters in length. It cannot start with http:// or https://.
	//
	// example:
	//
	// myFlowlog
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the ECR.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The name of the flow log.
	//
	// > The name can be empty or 1 to 128 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// myFlowlog
	FlowLogName *string `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	// The VBR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of network instance. Valid values:
	//
	// 	- **VBR**
	//
	// This parameter is required.
	//
	// example:
	//
	// VBR
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The time window for collecting log data. Unit: seconds. Valid values:
	//
	// - **60**
	//
	// - **600**
	//
	// Default value: **600**.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The Logstore that stores the captured traffic data.
	//
	// 	- If a Logstore is already created in the selected region, enter the name of the Logstore.
	//
	// 	- If no Logstores are created in the selected region, enter a name and the system automatically creates a Logstore. The name of the Logstore. The name must meet the following requirements:
	//
	// 	- The name must be unique in a project.
	//
	// 	- It can contain only lowercase letters, digits, hyphens (-), and underscores (_).
	//
	// 	- The name must start and end with a lowercase letter or a digit.
	//
	// 	- The name must be 3 to 63 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// flowlog-logstore
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The project that stores the captured traffic data.
	//
	// 	- If a project is already created in the selected region, enter the name of the project.
	//
	// 	- If no projects are created in the selected region, enter a name and the system automatically creates a project.
	//
	// The project name must be unique in a region. You cannot change the name after the project is created. The name must meet the following requirements:
	//
	// 	- The name must be globally unique.
	//
	// 	- The name can contain only lowercase letters,
	//
	// 	- digits, and hyphens (-).
	//
	// 	- The name must start and end with a lowercase letter or a digit.
	//
	// 	- The name must be 3 to 63 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// flowlog-project
	ProjectName     *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The sampling proportion. Valid values:
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
	SamplingRate *string                    `json:"SamplingRate,omitempty" xml:"SamplingRate,omitempty"`
	Tag          []*CreateFlowLogRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *CreateFlowLogRequest) Validate() error {
	return dara.Validate(s)
}

type CreateFlowLogRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
