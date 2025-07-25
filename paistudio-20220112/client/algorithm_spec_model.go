// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlgorithmSpec interface {
	dara.Model
	String() string
	GoString() string
	SetCodeDir(v *Location) *AlgorithmSpec
	GetCodeDir() *Location
	SetCommand(v []*string) *AlgorithmSpec
	GetCommand() []*string
	SetComputeResource(v *AlgorithmSpecComputeResource) *AlgorithmSpec
	GetComputeResource() *AlgorithmSpecComputeResource
	SetCustomization(v *AlgorithmSpecCustomization) *AlgorithmSpec
	GetCustomization() *AlgorithmSpecCustomization
	SetHyperParameters(v []*HyperParameterDefinition) *AlgorithmSpec
	GetHyperParameters() []*HyperParameterDefinition
	SetImage(v string) *AlgorithmSpec
	GetImage() *string
	SetInputChannels(v []*Channel) *AlgorithmSpec
	GetInputChannels() []*Channel
	SetJobType(v string) *AlgorithmSpec
	GetJobType() *string
	SetMetricDefinitions(v []*MetricDefinition) *AlgorithmSpec
	GetMetricDefinitions() []*MetricDefinition
	SetOutputChannels(v []*Channel) *AlgorithmSpec
	GetOutputChannels() []*Channel
	SetProgressDefinitions(v *AlgorithmSpecProgressDefinitions) *AlgorithmSpec
	GetProgressDefinitions() *AlgorithmSpecProgressDefinitions
	SetResourceRequirements(v []*ConditionExpression) *AlgorithmSpec
	GetResourceRequirements() []*ConditionExpression
	SetSupportedInstanceTypes(v []*string) *AlgorithmSpec
	GetSupportedInstanceTypes() []*string
	SetSupportsDistributedTraining(v bool) *AlgorithmSpec
	GetSupportsDistributedTraining() *bool
}

type AlgorithmSpec struct {
	CodeDir *Location `json:"CodeDir,omitempty" xml:"CodeDir,omitempty"`
	// This parameter is required.
	Command         []*string                     `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	ComputeResource *AlgorithmSpecComputeResource `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty" type:"Struct"`
	Customization   *AlgorithmSpecCustomization   `json:"Customization,omitempty" xml:"Customization,omitempty" type:"Struct"`
	HyperParameters []*HyperParameterDefinition   `json:"HyperParameters,omitempty" xml:"HyperParameters,omitempty" type:"Repeated"`
	// This parameter is required.
	Image         *string    `json:"Image,omitempty" xml:"Image,omitempty"`
	InputChannels []*Channel `json:"InputChannels,omitempty" xml:"InputChannels,omitempty" type:"Repeated"`
	// This parameter is required.
	JobType                     *string                           `json:"JobType,omitempty" xml:"JobType,omitempty"`
	MetricDefinitions           []*MetricDefinition               `json:"MetricDefinitions,omitempty" xml:"MetricDefinitions,omitempty" type:"Repeated"`
	OutputChannels              []*Channel                        `json:"OutputChannels,omitempty" xml:"OutputChannels,omitempty" type:"Repeated"`
	ProgressDefinitions         *AlgorithmSpecProgressDefinitions `json:"ProgressDefinitions,omitempty" xml:"ProgressDefinitions,omitempty" type:"Struct"`
	ResourceRequirements        []*ConditionExpression            `json:"ResourceRequirements,omitempty" xml:"ResourceRequirements,omitempty" type:"Repeated"`
	SupportedInstanceTypes      []*string                         `json:"SupportedInstanceTypes,omitempty" xml:"SupportedInstanceTypes,omitempty" type:"Repeated"`
	SupportsDistributedTraining *bool                             `json:"SupportsDistributedTraining,omitempty" xml:"SupportsDistributedTraining,omitempty"`
}

func (s AlgorithmSpec) String() string {
	return dara.Prettify(s)
}

func (s AlgorithmSpec) GoString() string {
	return s.String()
}

func (s *AlgorithmSpec) GetCodeDir() *Location {
	return s.CodeDir
}

func (s *AlgorithmSpec) GetCommand() []*string {
	return s.Command
}

func (s *AlgorithmSpec) GetComputeResource() *AlgorithmSpecComputeResource {
	return s.ComputeResource
}

func (s *AlgorithmSpec) GetCustomization() *AlgorithmSpecCustomization {
	return s.Customization
}

func (s *AlgorithmSpec) GetHyperParameters() []*HyperParameterDefinition {
	return s.HyperParameters
}

func (s *AlgorithmSpec) GetImage() *string {
	return s.Image
}

func (s *AlgorithmSpec) GetInputChannels() []*Channel {
	return s.InputChannels
}

func (s *AlgorithmSpec) GetJobType() *string {
	return s.JobType
}

func (s *AlgorithmSpec) GetMetricDefinitions() []*MetricDefinition {
	return s.MetricDefinitions
}

func (s *AlgorithmSpec) GetOutputChannels() []*Channel {
	return s.OutputChannels
}

func (s *AlgorithmSpec) GetProgressDefinitions() *AlgorithmSpecProgressDefinitions {
	return s.ProgressDefinitions
}

func (s *AlgorithmSpec) GetResourceRequirements() []*ConditionExpression {
	return s.ResourceRequirements
}

func (s *AlgorithmSpec) GetSupportedInstanceTypes() []*string {
	return s.SupportedInstanceTypes
}

func (s *AlgorithmSpec) GetSupportsDistributedTraining() *bool {
	return s.SupportsDistributedTraining
}

func (s *AlgorithmSpec) SetCodeDir(v *Location) *AlgorithmSpec {
	s.CodeDir = v
	return s
}

func (s *AlgorithmSpec) SetCommand(v []*string) *AlgorithmSpec {
	s.Command = v
	return s
}

func (s *AlgorithmSpec) SetComputeResource(v *AlgorithmSpecComputeResource) *AlgorithmSpec {
	s.ComputeResource = v
	return s
}

func (s *AlgorithmSpec) SetCustomization(v *AlgorithmSpecCustomization) *AlgorithmSpec {
	s.Customization = v
	return s
}

func (s *AlgorithmSpec) SetHyperParameters(v []*HyperParameterDefinition) *AlgorithmSpec {
	s.HyperParameters = v
	return s
}

func (s *AlgorithmSpec) SetImage(v string) *AlgorithmSpec {
	s.Image = &v
	return s
}

func (s *AlgorithmSpec) SetInputChannels(v []*Channel) *AlgorithmSpec {
	s.InputChannels = v
	return s
}

func (s *AlgorithmSpec) SetJobType(v string) *AlgorithmSpec {
	s.JobType = &v
	return s
}

func (s *AlgorithmSpec) SetMetricDefinitions(v []*MetricDefinition) *AlgorithmSpec {
	s.MetricDefinitions = v
	return s
}

func (s *AlgorithmSpec) SetOutputChannels(v []*Channel) *AlgorithmSpec {
	s.OutputChannels = v
	return s
}

func (s *AlgorithmSpec) SetProgressDefinitions(v *AlgorithmSpecProgressDefinitions) *AlgorithmSpec {
	s.ProgressDefinitions = v
	return s
}

func (s *AlgorithmSpec) SetResourceRequirements(v []*ConditionExpression) *AlgorithmSpec {
	s.ResourceRequirements = v
	return s
}

func (s *AlgorithmSpec) SetSupportedInstanceTypes(v []*string) *AlgorithmSpec {
	s.SupportedInstanceTypes = v
	return s
}

func (s *AlgorithmSpec) SetSupportsDistributedTraining(v bool) *AlgorithmSpec {
	s.SupportsDistributedTraining = &v
	return s
}

func (s *AlgorithmSpec) Validate() error {
	return dara.Validate(s)
}

type AlgorithmSpecComputeResource struct {
	// This parameter is required.
	Policy *AlgorithmSpecComputeResourcePolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
}

func (s AlgorithmSpecComputeResource) String() string {
	return dara.Prettify(s)
}

func (s AlgorithmSpecComputeResource) GoString() string {
	return s.String()
}

func (s *AlgorithmSpecComputeResource) GetPolicy() *AlgorithmSpecComputeResourcePolicy {
	return s.Policy
}

func (s *AlgorithmSpecComputeResource) SetPolicy(v *AlgorithmSpecComputeResourcePolicy) *AlgorithmSpecComputeResource {
	s.Policy = v
	return s
}

func (s *AlgorithmSpecComputeResource) Validate() error {
	return dara.Validate(s)
}

type AlgorithmSpecComputeResourcePolicy struct {
	// This parameter is required.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// This parameter is required.
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s AlgorithmSpecComputeResourcePolicy) String() string {
	return dara.Prettify(s)
}

func (s AlgorithmSpecComputeResourcePolicy) GoString() string {
	return s.String()
}

func (s *AlgorithmSpecComputeResourcePolicy) GetValue() *string {
	return s.Value
}

func (s *AlgorithmSpecComputeResourcePolicy) GetVersion() *string {
	return s.Version
}

func (s *AlgorithmSpecComputeResourcePolicy) SetValue(v string) *AlgorithmSpecComputeResourcePolicy {
	s.Value = &v
	return s
}

func (s *AlgorithmSpecComputeResourcePolicy) SetVersion(v string) *AlgorithmSpecComputeResourcePolicy {
	s.Version = &v
	return s
}

func (s *AlgorithmSpecComputeResourcePolicy) Validate() error {
	return dara.Validate(s)
}

type AlgorithmSpecCustomization struct {
	CodeDir *bool `json:"CodeDir,omitempty" xml:"CodeDir,omitempty"`
}

func (s AlgorithmSpecCustomization) String() string {
	return dara.Prettify(s)
}

func (s AlgorithmSpecCustomization) GoString() string {
	return s.String()
}

func (s *AlgorithmSpecCustomization) GetCodeDir() *bool {
	return s.CodeDir
}

func (s *AlgorithmSpecCustomization) SetCodeDir(v bool) *AlgorithmSpecCustomization {
	s.CodeDir = &v
	return s
}

func (s *AlgorithmSpecCustomization) Validate() error {
	return dara.Validate(s)
}

type AlgorithmSpecProgressDefinitions struct {
	OverallProgress *AlgorithmSpecProgressDefinitionsOverallProgress `json:"OverallProgress,omitempty" xml:"OverallProgress,omitempty" type:"Struct"`
	RemainingTime   *AlgorithmSpecProgressDefinitionsRemainingTime   `json:"RemainingTime,omitempty" xml:"RemainingTime,omitempty" type:"Struct"`
}

func (s AlgorithmSpecProgressDefinitions) String() string {
	return dara.Prettify(s)
}

func (s AlgorithmSpecProgressDefinitions) GoString() string {
	return s.String()
}

func (s *AlgorithmSpecProgressDefinitions) GetOverallProgress() *AlgorithmSpecProgressDefinitionsOverallProgress {
	return s.OverallProgress
}

func (s *AlgorithmSpecProgressDefinitions) GetRemainingTime() *AlgorithmSpecProgressDefinitionsRemainingTime {
	return s.RemainingTime
}

func (s *AlgorithmSpecProgressDefinitions) SetOverallProgress(v *AlgorithmSpecProgressDefinitionsOverallProgress) *AlgorithmSpecProgressDefinitions {
	s.OverallProgress = v
	return s
}

func (s *AlgorithmSpecProgressDefinitions) SetRemainingTime(v *AlgorithmSpecProgressDefinitionsRemainingTime) *AlgorithmSpecProgressDefinitions {
	s.RemainingTime = v
	return s
}

func (s *AlgorithmSpecProgressDefinitions) Validate() error {
	return dara.Validate(s)
}

type AlgorithmSpecProgressDefinitionsOverallProgress struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Regex       *string `json:"Regex,omitempty" xml:"Regex,omitempty"`
}

func (s AlgorithmSpecProgressDefinitionsOverallProgress) String() string {
	return dara.Prettify(s)
}

func (s AlgorithmSpecProgressDefinitionsOverallProgress) GoString() string {
	return s.String()
}

func (s *AlgorithmSpecProgressDefinitionsOverallProgress) GetDescription() *string {
	return s.Description
}

func (s *AlgorithmSpecProgressDefinitionsOverallProgress) GetRegex() *string {
	return s.Regex
}

func (s *AlgorithmSpecProgressDefinitionsOverallProgress) SetDescription(v string) *AlgorithmSpecProgressDefinitionsOverallProgress {
	s.Description = &v
	return s
}

func (s *AlgorithmSpecProgressDefinitionsOverallProgress) SetRegex(v string) *AlgorithmSpecProgressDefinitionsOverallProgress {
	s.Regex = &v
	return s
}

func (s *AlgorithmSpecProgressDefinitionsOverallProgress) Validate() error {
	return dara.Validate(s)
}

type AlgorithmSpecProgressDefinitionsRemainingTime struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Regex       *string `json:"Regex,omitempty" xml:"Regex,omitempty"`
}

func (s AlgorithmSpecProgressDefinitionsRemainingTime) String() string {
	return dara.Prettify(s)
}

func (s AlgorithmSpecProgressDefinitionsRemainingTime) GoString() string {
	return s.String()
}

func (s *AlgorithmSpecProgressDefinitionsRemainingTime) GetDescription() *string {
	return s.Description
}

func (s *AlgorithmSpecProgressDefinitionsRemainingTime) GetRegex() *string {
	return s.Regex
}

func (s *AlgorithmSpecProgressDefinitionsRemainingTime) SetDescription(v string) *AlgorithmSpecProgressDefinitionsRemainingTime {
	s.Description = &v
	return s
}

func (s *AlgorithmSpecProgressDefinitionsRemainingTime) SetRegex(v string) *AlgorithmSpecProgressDefinitionsRemainingTime {
	s.Regex = &v
	return s
}

func (s *AlgorithmSpecProgressDefinitionsRemainingTime) Validate() error {
	return dara.Validate(s)
}
