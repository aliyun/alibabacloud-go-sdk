// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStackConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*ListStackConfigsResponseBodyConfigs) *ListStackConfigsResponseBody
	GetConfigs() []*ListStackConfigsResponseBodyConfigs
	SetMaxResults(v int32) *ListStackConfigsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListStackConfigsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListStackConfigsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListStackConfigsResponseBody
	GetTotalCount() *int32
}

type ListStackConfigsResponseBody struct {
	// The list of stack configurations.
	Configs []*ListStackConfigsResponseBodyConfigs `json:"configs,omitempty" xml:"configs,omitempty" type:"Repeated"`
	// The maximum number of records returned in this request.
	//
	// example:
	//
	// 24
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The position from which the current call starts reading. An empty value indicates that all data has been read.
	//
	// This parameter is required.
	//
	// example:
	//
	// jIFUaFVhy2VD6whh5GaY854dD+2BRJj42DLT6GrZysw=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9BEDBCF8-03BE-5A59-AC93-9263942B37E8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of records that match the request conditions. This parameter is optional and may not be returned by default.
	//
	// example:
	//
	// 43
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListStackConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStackConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStackConfigsResponseBody) GetConfigs() []*ListStackConfigsResponseBodyConfigs {
	return s.Configs
}

func (s *ListStackConfigsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListStackConfigsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListStackConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStackConfigsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListStackConfigsResponseBody) SetConfigs(v []*ListStackConfigsResponseBodyConfigs) *ListStackConfigsResponseBody {
	s.Configs = v
	return s
}

func (s *ListStackConfigsResponseBody) SetMaxResults(v int32) *ListStackConfigsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListStackConfigsResponseBody) SetNextToken(v string) *ListStackConfigsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListStackConfigsResponseBody) SetRequestId(v string) *ListStackConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStackConfigsResponseBody) SetTotalCount(v int32) *ListStackConfigsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListStackConfigsResponseBody) Validate() error {
	if s.Configs != nil {
		for _, item := range s.Configs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListStackConfigsResponseBodyConfigs struct {
	// The component configuration.
	ComponentConfig *ListStackConfigsResponseBodyConfigsComponentConfig `json:"componentConfig,omitempty" xml:"componentConfig,omitempty" type:"Struct"`
	// The content of the component configuration.
	//
	// example:
	//
	// format_version: IaCService/2021-08-06\\ndescription: create ALB \\nvariable:\\n  - name: region\\n    type: string\\n ...
	ComponentContent *string `json:"componentContent,omitempty" xml:"componentContent,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-08-15T16:14:06Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The deployment configuration.
	DeploymentConfig *ListStackConfigsResponseBodyConfigsDeploymentConfig `json:"deploymentConfig,omitempty" xml:"deploymentConfig,omitempty" type:"Struct"`
	// The content of the deployment configuration.
	//
	// example:
	//
	// format_version: IaCService/2021-08-06\\ndescription: create ALB\\nupstream_input:\\n  - name: stack_network\\n ...
	DeploymentContent *string `json:"deploymentContent,omitempty" xml:"deploymentContent,omitempty"`
	FailedReason      *string `json:"failedReason,omitempty" xml:"failedReason,omitempty"`
	// The status of the stack configuration.
	//
	// | Name | Description |
	//
	// |------|------|
	//
	// | Creating | Being created. |
	//
	// | Created | Created. |
	//
	// | Waiting | Waiting for deployment. |
	//
	// | Deploying | Being deployed. |
	//
	// | Deployed | Deployed. |
	//
	// | Errored | Deployment failed. |
	//
	// | Deleting | Being deleted. |
	//
	// | Deleted | Deleted. |
	//
	// | DeleteFailed | Deletion failed. |
	//
	// | DetectTriggered | Drift detection triggered. |.
	//
	// example:
	//
	// Deployed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The configuration version number, such as v1. The initial value is v1. The version number increments each time the stack is updated or refreshed and the configuration changes.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListStackConfigsResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListStackConfigsResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListStackConfigsResponseBodyConfigs) GetComponentConfig() *ListStackConfigsResponseBodyConfigsComponentConfig {
	return s.ComponentConfig
}

func (s *ListStackConfigsResponseBodyConfigs) GetComponentContent() *string {
	return s.ComponentContent
}

func (s *ListStackConfigsResponseBodyConfigs) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListStackConfigsResponseBodyConfigs) GetDeploymentConfig() *ListStackConfigsResponseBodyConfigsDeploymentConfig {
	return s.DeploymentConfig
}

func (s *ListStackConfigsResponseBodyConfigs) GetDeploymentContent() *string {
	return s.DeploymentContent
}

func (s *ListStackConfigsResponseBodyConfigs) GetFailedReason() *string {
	return s.FailedReason
}

func (s *ListStackConfigsResponseBodyConfigs) GetStatus() *string {
	return s.Status
}

func (s *ListStackConfigsResponseBodyConfigs) GetVersion() *string {
	return s.Version
}

func (s *ListStackConfigsResponseBodyConfigs) SetComponentConfig(v *ListStackConfigsResponseBodyConfigsComponentConfig) *ListStackConfigsResponseBodyConfigs {
	s.ComponentConfig = v
	return s
}

func (s *ListStackConfigsResponseBodyConfigs) SetComponentContent(v string) *ListStackConfigsResponseBodyConfigs {
	s.ComponentContent = &v
	return s
}

func (s *ListStackConfigsResponseBodyConfigs) SetCreateTime(v string) *ListStackConfigsResponseBodyConfigs {
	s.CreateTime = &v
	return s
}

func (s *ListStackConfigsResponseBodyConfigs) SetDeploymentConfig(v *ListStackConfigsResponseBodyConfigsDeploymentConfig) *ListStackConfigsResponseBodyConfigs {
	s.DeploymentConfig = v
	return s
}

func (s *ListStackConfigsResponseBodyConfigs) SetDeploymentContent(v string) *ListStackConfigsResponseBodyConfigs {
	s.DeploymentContent = &v
	return s
}

func (s *ListStackConfigsResponseBodyConfigs) SetFailedReason(v string) *ListStackConfigsResponseBodyConfigs {
	s.FailedReason = &v
	return s
}

func (s *ListStackConfigsResponseBodyConfigs) SetStatus(v string) *ListStackConfigsResponseBodyConfigs {
	s.Status = &v
	return s
}

func (s *ListStackConfigsResponseBodyConfigs) SetVersion(v string) *ListStackConfigsResponseBodyConfigs {
	s.Version = &v
	return s
}

func (s *ListStackConfigsResponseBodyConfigs) Validate() error {
	if s.ComponentConfig != nil {
		if err := s.ComponentConfig.Validate(); err != nil {
			return err
		}
	}
	if s.DeploymentConfig != nil {
		if err := s.DeploymentConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListStackConfigsResponseBodyConfigsComponentConfig struct {
	// The list of components.
	Component []*ListStackConfigsResponseBodyConfigsComponentConfigComponent `json:"component,omitempty" xml:"component,omitempty" type:"Repeated"`
	// The list of component outputs.
	Output []*ListStackConfigsResponseBodyConfigsComponentConfigOutput `json:"output,omitempty" xml:"output,omitempty" type:"Repeated"`
	// The list of component variables.
	Variable []*ListStackConfigsResponseBodyConfigsComponentConfigVariable `json:"variable,omitempty" xml:"variable,omitempty" type:"Repeated"`
}

func (s ListStackConfigsResponseBodyConfigsComponentConfig) String() string {
	return dara.Prettify(s)
}

func (s ListStackConfigsResponseBodyConfigsComponentConfig) GoString() string {
	return s.String()
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfig) GetComponent() []*ListStackConfigsResponseBodyConfigsComponentConfigComponent {
	return s.Component
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfig) GetOutput() []*ListStackConfigsResponseBodyConfigsComponentConfigOutput {
	return s.Output
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfig) GetVariable() []*ListStackConfigsResponseBodyConfigsComponentConfigVariable {
	return s.Variable
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfig) SetComponent(v []*ListStackConfigsResponseBodyConfigsComponentConfigComponent) *ListStackConfigsResponseBodyConfigsComponentConfig {
	s.Component = v
	return s
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfig) SetOutput(v []*ListStackConfigsResponseBodyConfigsComponentConfigOutput) *ListStackConfigsResponseBodyConfigsComponentConfig {
	s.Output = v
	return s
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfig) SetVariable(v []*ListStackConfigsResponseBodyConfigsComponentConfigVariable) *ListStackConfigsResponseBodyConfigsComponentConfig {
	s.Variable = v
	return s
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfig) Validate() error {
	if s.Component != nil {
		for _, item := range s.Component {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Output != nil {
		for _, item := range s.Output {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Variable != nil {
		for _, item := range s.Variable {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListStackConfigsResponseBodyConfigsComponentConfigComponent struct {
	// The component name.
	//
	// example:
	//
	// log
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListStackConfigsResponseBodyConfigsComponentConfigComponent) String() string {
	return dara.Prettify(s)
}

func (s ListStackConfigsResponseBodyConfigsComponentConfigComponent) GoString() string {
	return s.String()
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfigComponent) GetName() *string {
	return s.Name
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfigComponent) SetName(v string) *ListStackConfigsResponseBodyConfigsComponentConfigComponent {
	s.Name = &v
	return s
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfigComponent) Validate() error {
	return dara.Validate(s)
}

type ListStackConfigsResponseBodyConfigsComponentConfigOutput struct {
	// The output description.
	//
	// example:
	//
	// the name of sls project
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The output name.
	//
	// example:
	//
	// project_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The output type.
	//
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The output value.
	//
	// example:
	//
	// log-test
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListStackConfigsResponseBodyConfigsComponentConfigOutput) String() string {
	return dara.Prettify(s)
}

func (s ListStackConfigsResponseBodyConfigsComponentConfigOutput) GoString() string {
	return s.String()
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfigOutput) GetDescription() *string {
	return s.Description
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfigOutput) GetName() *string {
	return s.Name
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfigOutput) GetType() *string {
	return s.Type
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfigOutput) GetValue() *string {
	return s.Value
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfigOutput) SetDescription(v string) *ListStackConfigsResponseBodyConfigsComponentConfigOutput {
	s.Description = &v
	return s
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfigOutput) SetName(v string) *ListStackConfigsResponseBodyConfigsComponentConfigOutput {
	s.Name = &v
	return s
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfigOutput) SetType(v string) *ListStackConfigsResponseBodyConfigsComponentConfigOutput {
	s.Type = &v
	return s
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfigOutput) SetValue(v string) *ListStackConfigsResponseBodyConfigsComponentConfigOutput {
	s.Value = &v
	return s
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfigOutput) Validate() error {
	return dara.Validate(s)
}

type ListStackConfigsResponseBodyConfigsComponentConfigVariable struct {
	// The default value.
	//
	// example:
	//
	// ap-southeast-3
	Default *string `json:"default,omitempty" xml:"default,omitempty"`
	// The description.
	//
	// example:
	//
	// region of sls project
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The variable name.
	//
	// example:
	//
	// region
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Sensitive *bool   `json:"sensitive,omitempty" xml:"sensitive,omitempty"`
	// The variable type, such as:
	//
	// - string
	//
	// - list(string)
	//
	// - map(string).
	//
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListStackConfigsResponseBodyConfigsComponentConfigVariable) String() string {
	return dara.Prettify(s)
}

func (s ListStackConfigsResponseBodyConfigsComponentConfigVariable) GoString() string {
	return s.String()
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfigVariable) GetDefault() *string {
	return s.Default
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfigVariable) GetDescription() *string {
	return s.Description
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfigVariable) GetName() *string {
	return s.Name
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfigVariable) GetSensitive() *bool {
	return s.Sensitive
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfigVariable) GetType() *string {
	return s.Type
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfigVariable) SetDefault(v string) *ListStackConfigsResponseBodyConfigsComponentConfigVariable {
	s.Default = &v
	return s
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfigVariable) SetDescription(v string) *ListStackConfigsResponseBodyConfigsComponentConfigVariable {
	s.Description = &v
	return s
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfigVariable) SetName(v string) *ListStackConfigsResponseBodyConfigsComponentConfigVariable {
	s.Name = &v
	return s
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfigVariable) SetSensitive(v bool) *ListStackConfigsResponseBodyConfigsComponentConfigVariable {
	s.Sensitive = &v
	return s
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfigVariable) SetType(v string) *ListStackConfigsResponseBodyConfigsComponentConfigVariable {
	s.Type = &v
	return s
}

func (s *ListStackConfigsResponseBodyConfigsComponentConfigVariable) Validate() error {
	return dara.Validate(s)
}

type ListStackConfigsResponseBodyConfigsDeploymentConfig struct {
	// The list of deployments.
	Deployment []*ListStackConfigsResponseBodyConfigsDeploymentConfigDeployment `json:"deployment,omitempty" xml:"deployment,omitempty" type:"Repeated"`
	// The list of outputs.
	PublishOutput []*ListStackConfigsResponseBodyConfigsDeploymentConfigPublishOutput `json:"publishOutput,omitempty" xml:"publishOutput,omitempty" type:"Repeated"`
	// The list of upstream inputs.
	UpstreamInput []*ListStackConfigsResponseBodyConfigsDeploymentConfigUpstreamInput `json:"upstreamInput,omitempty" xml:"upstreamInput,omitempty" type:"Repeated"`
}

func (s ListStackConfigsResponseBodyConfigsDeploymentConfig) String() string {
	return dara.Prettify(s)
}

func (s ListStackConfigsResponseBodyConfigsDeploymentConfig) GoString() string {
	return s.String()
}

func (s *ListStackConfigsResponseBodyConfigsDeploymentConfig) GetDeployment() []*ListStackConfigsResponseBodyConfigsDeploymentConfigDeployment {
	return s.Deployment
}

func (s *ListStackConfigsResponseBodyConfigsDeploymentConfig) GetPublishOutput() []*ListStackConfigsResponseBodyConfigsDeploymentConfigPublishOutput {
	return s.PublishOutput
}

func (s *ListStackConfigsResponseBodyConfigsDeploymentConfig) GetUpstreamInput() []*ListStackConfigsResponseBodyConfigsDeploymentConfigUpstreamInput {
	return s.UpstreamInput
}

func (s *ListStackConfigsResponseBodyConfigsDeploymentConfig) SetDeployment(v []*ListStackConfigsResponseBodyConfigsDeploymentConfigDeployment) *ListStackConfigsResponseBodyConfigsDeploymentConfig {
	s.Deployment = v
	return s
}

func (s *ListStackConfigsResponseBodyConfigsDeploymentConfig) SetPublishOutput(v []*ListStackConfigsResponseBodyConfigsDeploymentConfigPublishOutput) *ListStackConfigsResponseBodyConfigsDeploymentConfig {
	s.PublishOutput = v
	return s
}

func (s *ListStackConfigsResponseBodyConfigsDeploymentConfig) SetUpstreamInput(v []*ListStackConfigsResponseBodyConfigsDeploymentConfigUpstreamInput) *ListStackConfigsResponseBodyConfigsDeploymentConfig {
	s.UpstreamInput = v
	return s
}

func (s *ListStackConfigsResponseBodyConfigsDeploymentConfig) Validate() error {
	if s.Deployment != nil {
		for _, item := range s.Deployment {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PublishOutput != nil {
		for _, item := range s.PublishOutput {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UpstreamInput != nil {
		for _, item := range s.UpstreamInput {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListStackConfigsResponseBodyConfigsDeploymentConfigDeployment struct {
	// The deployment name.
	//
	// example:
	//
	// production
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListStackConfigsResponseBodyConfigsDeploymentConfigDeployment) String() string {
	return dara.Prettify(s)
}

func (s ListStackConfigsResponseBodyConfigsDeploymentConfigDeployment) GoString() string {
	return s.String()
}

func (s *ListStackConfigsResponseBodyConfigsDeploymentConfigDeployment) GetName() *string {
	return s.Name
}

func (s *ListStackConfigsResponseBodyConfigsDeploymentConfigDeployment) SetName(v string) *ListStackConfigsResponseBodyConfigsDeploymentConfigDeployment {
	s.Name = &v
	return s
}

func (s *ListStackConfigsResponseBodyConfigsDeploymentConfigDeployment) Validate() error {
	return dara.Validate(s)
}

type ListStackConfigsResponseBodyConfigsDeploymentConfigPublishOutput struct {
	// The output description.
	//
	// example:
	//
	// the name of sls project
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The output name.
	//
	// example:
	//
	// project_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The actual output value after the stack deployment is complete.
	//
	// example:
	//
	// log-test
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// The output type, such as string or list(string).
	//
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The original definition of the output value. Currently, string or list(string) is supported. You can reference a deployment output in the format: deployment.{deploymentName}.{deploymentOutputName}.
	//
	// example:
	//
	// deployment.production.project_name
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListStackConfigsResponseBodyConfigsDeploymentConfigPublishOutput) String() string {
	return dara.Prettify(s)
}

func (s ListStackConfigsResponseBodyConfigsDeploymentConfigPublishOutput) GoString() string {
	return s.String()
}

func (s *ListStackConfigsResponseBodyConfigsDeploymentConfigPublishOutput) GetDescription() *string {
	return s.Description
}

func (s *ListStackConfigsResponseBodyConfigsDeploymentConfigPublishOutput) GetName() *string {
	return s.Name
}

func (s *ListStackConfigsResponseBodyConfigsDeploymentConfigPublishOutput) GetResult() *string {
	return s.Result
}

func (s *ListStackConfigsResponseBodyConfigsDeploymentConfigPublishOutput) GetType() *string {
	return s.Type
}

func (s *ListStackConfigsResponseBodyConfigsDeploymentConfigPublishOutput) GetValue() *string {
	return s.Value
}

func (s *ListStackConfigsResponseBodyConfigsDeploymentConfigPublishOutput) SetDescription(v string) *ListStackConfigsResponseBodyConfigsDeploymentConfigPublishOutput {
	s.Description = &v
	return s
}

func (s *ListStackConfigsResponseBodyConfigsDeploymentConfigPublishOutput) SetName(v string) *ListStackConfigsResponseBodyConfigsDeploymentConfigPublishOutput {
	s.Name = &v
	return s
}

func (s *ListStackConfigsResponseBodyConfigsDeploymentConfigPublishOutput) SetResult(v string) *ListStackConfigsResponseBodyConfigsDeploymentConfigPublishOutput {
	s.Result = &v
	return s
}

func (s *ListStackConfigsResponseBodyConfigsDeploymentConfigPublishOutput) SetType(v string) *ListStackConfigsResponseBodyConfigsDeploymentConfigPublishOutput {
	s.Type = &v
	return s
}

func (s *ListStackConfigsResponseBodyConfigsDeploymentConfigPublishOutput) SetValue(v string) *ListStackConfigsResponseBodyConfigsDeploymentConfigPublishOutput {
	s.Value = &v
	return s
}

func (s *ListStackConfigsResponseBodyConfigsDeploymentConfigPublishOutput) Validate() error {
	return dara.Validate(s)
}

type ListStackConfigsResponseBodyConfigsDeploymentConfigUpstreamInput struct {
	// The input name.
	//
	// example:
	//
	// network
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The input source. Currently, only an upstream stack can be specified. The format is {iacEndpoint}/{accountId}/{upstreamStackName}.
	//
	// example:
	//
	// IacEndpoint/156718871222312/stack_network
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s ListStackConfigsResponseBodyConfigsDeploymentConfigUpstreamInput) String() string {
	return dara.Prettify(s)
}

func (s ListStackConfigsResponseBodyConfigsDeploymentConfigUpstreamInput) GoString() string {
	return s.String()
}

func (s *ListStackConfigsResponseBodyConfigsDeploymentConfigUpstreamInput) GetName() *string {
	return s.Name
}

func (s *ListStackConfigsResponseBodyConfigsDeploymentConfigUpstreamInput) GetSource() *string {
	return s.Source
}

func (s *ListStackConfigsResponseBodyConfigsDeploymentConfigUpstreamInput) SetName(v string) *ListStackConfigsResponseBodyConfigsDeploymentConfigUpstreamInput {
	s.Name = &v
	return s
}

func (s *ListStackConfigsResponseBodyConfigsDeploymentConfigUpstreamInput) SetSource(v string) *ListStackConfigsResponseBodyConfigsDeploymentConfigUpstreamInput {
	s.Source = &v
	return s
}

func (s *ListStackConfigsResponseBodyConfigsDeploymentConfigUpstreamInput) Validate() error {
	return dara.Validate(s)
}
