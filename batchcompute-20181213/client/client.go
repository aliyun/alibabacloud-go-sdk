// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type Attempt struct {
	// EndTime
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// ExitCode
	ExitCode *int32 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// Output
	Output []byte `json:"Output,omitempty" xml:"Output,omitempty"`
	// Pid
	Pid *int32 `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// Reason
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// StartTime
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// State
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// UserStages
	UserStages []*UserStage `json:"UserStages,omitempty" xml:"UserStages,omitempty" type:"Repeated"`
	// Worker
	Worker *string `json:"Worker,omitempty" xml:"Worker,omitempty"`
}

func (s Attempt) String() string {
	return tea.Prettify(s)
}

func (s Attempt) GoString() string {
	return s.String()
}

func (s *Attempt) SetEndTime(v string) *Attempt {
	s.EndTime = &v
	return s
}

func (s *Attempt) SetExitCode(v int32) *Attempt {
	s.ExitCode = &v
	return s
}

func (s *Attempt) SetOutput(v []byte) *Attempt {
	s.Output = v
	return s
}

func (s *Attempt) SetPid(v int32) *Attempt {
	s.Pid = &v
	return s
}

func (s *Attempt) SetReason(v string) *Attempt {
	s.Reason = &v
	return s
}

func (s *Attempt) SetStartTime(v string) *Attempt {
	s.StartTime = &v
	return s
}

func (s *Attempt) SetState(v string) *Attempt {
	s.State = &v
	return s
}

func (s *Attempt) SetUserStages(v []*UserStage) *Attempt {
	s.UserStages = v
	return s
}

func (s *Attempt) SetWorker(v string) *Attempt {
	s.Worker = &v
	return s
}

type AutoScaling struct {
	// Scaling
	Scaling *Scaling `json:"Scaling,omitempty" xml:"Scaling,omitempty"`
	// Trigger
	Trigger *Trigger `json:"Trigger,omitempty" xml:"Trigger,omitempty"`
}

func (s AutoScaling) String() string {
	return tea.Prettify(s)
}

func (s AutoScaling) GoString() string {
	return s.String()
}

func (s *AutoScaling) SetScaling(v *Scaling) *AutoScaling {
	s.Scaling = v
	return s
}

func (s *AutoScaling) SetTrigger(v *Trigger) *AutoScaling {
	s.Trigger = v
	return s
}

type Bootstrap struct {
	// Background
	Background *bool `json:"Background,omitempty" xml:"Background,omitempty"`
	// Command
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	// Envs
	Envs map[string]*string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// Loggings
	Loggings *Logging `json:"Loggings,omitempty" xml:"Loggings,omitempty"`
	// MountPoints
	MountPoints []*MountPoint `json:"MountPoints,omitempty" xml:"MountPoints,omitempty" type:"Repeated"`
	// PackageUri
	PackageUri *string `json:"PackageUri,omitempty" xml:"PackageUri,omitempty"`
	// RunningTimeout
	RunningTimeout *int32 `json:"RunningTimeout,omitempty" xml:"RunningTimeout,omitempty"`
	// Runtimes
	Runtimes *BootstrapRuntime `json:"Runtimes,omitempty" xml:"Runtimes,omitempty"`
	// Volumes
	Volumes []*Volume `json:"Volumes,omitempty" xml:"Volumes,omitempty" type:"Repeated"`
}

func (s Bootstrap) String() string {
	return tea.Prettify(s)
}

func (s Bootstrap) GoString() string {
	return s.String()
}

func (s *Bootstrap) SetBackground(v bool) *Bootstrap {
	s.Background = &v
	return s
}

func (s *Bootstrap) SetCommand(v []*string) *Bootstrap {
	s.Command = v
	return s
}

func (s *Bootstrap) SetEnvs(v map[string]*string) *Bootstrap {
	s.Envs = v
	return s
}

func (s *Bootstrap) SetLoggings(v *Logging) *Bootstrap {
	s.Loggings = v
	return s
}

func (s *Bootstrap) SetMountPoints(v []*MountPoint) *Bootstrap {
	s.MountPoints = v
	return s
}

func (s *Bootstrap) SetPackageUri(v string) *Bootstrap {
	s.PackageUri = &v
	return s
}

func (s *Bootstrap) SetRunningTimeout(v int32) *Bootstrap {
	s.RunningTimeout = &v
	return s
}

func (s *Bootstrap) SetRuntimes(v *BootstrapRuntime) *Bootstrap {
	s.Runtimes = v
	return s
}

func (s *Bootstrap) SetVolumes(v []*Volume) *Bootstrap {
	s.Volumes = v
	return s
}

type BootstrapRuntime struct {
	// Docker
	Docker *Docker `json:"Docker,omitempty" xml:"Docker,omitempty"`
}

func (s BootstrapRuntime) String() string {
	return tea.Prettify(s)
}

func (s BootstrapRuntime) GoString() string {
	return s.String()
}

func (s *BootstrapRuntime) SetDocker(v *Docker) *BootstrapRuntime {
	s.Docker = v
	return s
}

type ClusterDefinition struct {
	// AutoScaling
	AutoScaling []*AutoScaling `json:"AutoScaling,omitempty" xml:"AutoScaling,omitempty" type:"Repeated"`
	// Bootstrap
	Bootstrap *Bootstrap `json:"Bootstrap,omitempty" xml:"Bootstrap,omitempty"`
	// CredentialConfigs
	CredentialConfigs *CredentialConfig `json:"CredentialConfigs,omitempty" xml:"CredentialConfigs,omitempty"`
	// Docker
	Docker *Docker `json:"Docker,omitempty" xml:"Docker,omitempty"`
	// ECS
	ECS *ECS `json:"ECS,omitempty" xml:"ECS,omitempty"`
	// LivenessProbe
	LivenessProbe *Probe `json:"LivenessProbe,omitempty" xml:"LivenessProbe,omitempty"`
	// ManagedJobQueue
	ManagedJobQueue *bool `json:"ManagedJobQueue,omitempty" xml:"ManagedJobQueue,omitempty"`
	// MountPoints
	MountPoints []*MountPoint `json:"MountPoints,omitempty" xml:"MountPoints,omitempty" type:"Repeated"`
	// ProviderType
	ProviderType *string `json:"ProviderType,omitempty" xml:"ProviderType,omitempty"`
	// Resources
	Resources map[string]*string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// SLB
	SLB *SLB `json:"SLB,omitempty" xml:"SLB,omitempty"`
	// Scaling
	Scaling *Scaling `json:"Scaling,omitempty" xml:"Scaling,omitempty"`
	// StartupProbe
	StartupProbe *Probe `json:"StartupProbe,omitempty" xml:"StartupProbe,omitempty"`
	// UpgradePolicy
	UpgradePolicy *UpgradePolicy `json:"UpgradePolicy,omitempty" xml:"UpgradePolicy,omitempty"`
	// VPC
	VPC *VPC `json:"VPC,omitempty" xml:"VPC,omitempty"`
	// Volumes
	Volumes []*Volume `json:"Volumes,omitempty" xml:"Volumes,omitempty" type:"Repeated"`
	// WorkerType
	WorkerType *string `json:"WorkerType,omitempty" xml:"WorkerType,omitempty"`
}

func (s ClusterDefinition) String() string {
	return tea.Prettify(s)
}

func (s ClusterDefinition) GoString() string {
	return s.String()
}

func (s *ClusterDefinition) SetAutoScaling(v []*AutoScaling) *ClusterDefinition {
	s.AutoScaling = v
	return s
}

func (s *ClusterDefinition) SetBootstrap(v *Bootstrap) *ClusterDefinition {
	s.Bootstrap = v
	return s
}

func (s *ClusterDefinition) SetCredentialConfigs(v *CredentialConfig) *ClusterDefinition {
	s.CredentialConfigs = v
	return s
}

func (s *ClusterDefinition) SetDocker(v *Docker) *ClusterDefinition {
	s.Docker = v
	return s
}

func (s *ClusterDefinition) SetECS(v *ECS) *ClusterDefinition {
	s.ECS = v
	return s
}

func (s *ClusterDefinition) SetLivenessProbe(v *Probe) *ClusterDefinition {
	s.LivenessProbe = v
	return s
}

func (s *ClusterDefinition) SetManagedJobQueue(v bool) *ClusterDefinition {
	s.ManagedJobQueue = &v
	return s
}

func (s *ClusterDefinition) SetMountPoints(v []*MountPoint) *ClusterDefinition {
	s.MountPoints = v
	return s
}

func (s *ClusterDefinition) SetProviderType(v string) *ClusterDefinition {
	s.ProviderType = &v
	return s
}

func (s *ClusterDefinition) SetResources(v map[string]*string) *ClusterDefinition {
	s.Resources = v
	return s
}

func (s *ClusterDefinition) SetSLB(v *SLB) *ClusterDefinition {
	s.SLB = v
	return s
}

func (s *ClusterDefinition) SetScaling(v *Scaling) *ClusterDefinition {
	s.Scaling = v
	return s
}

func (s *ClusterDefinition) SetStartupProbe(v *Probe) *ClusterDefinition {
	s.StartupProbe = v
	return s
}

func (s *ClusterDefinition) SetUpgradePolicy(v *UpgradePolicy) *ClusterDefinition {
	s.UpgradePolicy = v
	return s
}

func (s *ClusterDefinition) SetVPC(v *VPC) *ClusterDefinition {
	s.VPC = v
	return s
}

func (s *ClusterDefinition) SetVolumes(v []*Volume) *ClusterDefinition {
	s.Volumes = v
	return s
}

func (s *ClusterDefinition) SetWorkerType(v string) *ClusterDefinition {
	s.WorkerType = &v
	return s
}

type Conditions struct {
	// Condition
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// Errors
	Errors []*Errors `json:"Errors,omitempty" xml:"Errors,omitempty" type:"Repeated"`
	// LastProbeTime
	LastProbeTime *string `json:"LastProbeTime,omitempty" xml:"LastProbeTime,omitempty"`
	// LastTransitionTime
	LastTransitionTime *string `json:"LastTransitionTime,omitempty" xml:"LastTransitionTime,omitempty"`
	// Status
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s Conditions) String() string {
	return tea.Prettify(s)
}

func (s Conditions) GoString() string {
	return s.String()
}

func (s *Conditions) SetCondition(v string) *Conditions {
	s.Condition = &v
	return s
}

func (s *Conditions) SetErrors(v []*Errors) *Conditions {
	s.Errors = v
	return s
}

func (s *Conditions) SetLastProbeTime(v string) *Conditions {
	s.LastProbeTime = &v
	return s
}

func (s *Conditions) SetLastTransitionTime(v string) *Conditions {
	s.LastTransitionTime = &v
	return s
}

func (s *Conditions) SetStatus(v string) *Conditions {
	s.Status = &v
	return s
}

type CredentialConfig struct {
	// Chain
	Chain []*ServiceRoleNode `json:"Chain,omitempty" xml:"Chain,omitempty" type:"Repeated"`
	// Policy
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// ServiceRole
	ServiceRole *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
}

func (s CredentialConfig) String() string {
	return tea.Prettify(s)
}

func (s CredentialConfig) GoString() string {
	return s.String()
}

func (s *CredentialConfig) SetChain(v []*ServiceRoleNode) *CredentialConfig {
	s.Chain = v
	return s
}

func (s *CredentialConfig) SetPolicy(v string) *CredentialConfig {
	s.Policy = &v
	return s
}

func (s *CredentialConfig) SetServiceRole(v string) *CredentialConfig {
	s.ServiceRole = &v
	return s
}

type Destination struct {
	OSS *OSSDescription `json:"OSS,omitempty" xml:"OSS,omitempty"`
	PDS *PDSDescription `json:"PDS,omitempty" xml:"PDS,omitempty"`
}

func (s Destination) String() string {
	return tea.Prettify(s)
}

func (s Destination) GoString() string {
	return s.String()
}

func (s *Destination) SetOSS(v *OSSDescription) *Destination {
	s.OSS = v
	return s
}

func (s *Destination) SetPDS(v *PDSDescription) *Destination {
	s.PDS = v
	return s
}

type Docker struct {
	// ExposedPorts
	ExposedPorts []*ExposedPort `json:"ExposedPorts,omitempty" xml:"ExposedPorts,omitempty" type:"Repeated"`
	// Image
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
}

func (s Docker) String() string {
	return tea.Prettify(s)
}

func (s Docker) GoString() string {
	return s.String()
}

func (s *Docker) SetExposedPorts(v []*ExposedPort) *Docker {
	s.ExposedPorts = v
	return s
}

func (s *Docker) SetImage(v string) *Docker {
	s.Image = &v
	return s
}

type ECS struct {
	// HostnamePrefix
	HostnamePrefix *string `json:"HostnamePrefix,omitempty" xml:"HostnamePrefix,omitempty"`
	// InstanceType
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// PasswordInherit
	PasswordInherit *bool `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	// ResourceType
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// SpotPriceLimit
	SpotPriceLimit *string `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// SpotStrategy
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// SystemDiskSize
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// SystemDiskType
	SystemDiskType *string `json:"SystemDiskType,omitempty" xml:"SystemDiskType,omitempty"`
	// VMImage
	VMImage *string `json:"VMImage,omitempty" xml:"VMImage,omitempty"`
}

func (s ECS) String() string {
	return tea.Prettify(s)
}

func (s ECS) GoString() string {
	return s.String()
}

func (s *ECS) SetHostnamePrefix(v string) *ECS {
	s.HostnamePrefix = &v
	return s
}

func (s *ECS) SetInstanceType(v string) *ECS {
	s.InstanceType = &v
	return s
}

func (s *ECS) SetPasswordInherit(v bool) *ECS {
	s.PasswordInherit = &v
	return s
}

func (s *ECS) SetResourceType(v string) *ECS {
	s.ResourceType = &v
	return s
}

func (s *ECS) SetSpotPriceLimit(v string) *ECS {
	s.SpotPriceLimit = &v
	return s
}

func (s *ECS) SetSpotStrategy(v string) *ECS {
	s.SpotStrategy = &v
	return s
}

func (s *ECS) SetSystemDiskSize(v int32) *ECS {
	s.SystemDiskSize = &v
	return s
}

func (s *ECS) SetSystemDiskType(v string) *ECS {
	s.SystemDiskType = &v
	return s
}

func (s *ECS) SetVMImage(v string) *ECS {
	s.VMImage = &v
	return s
}

type Errors struct {
	// Action
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Repeat
	Repeat *int32 `json:"Repeat,omitempty" xml:"Repeat,omitempty"`
}

func (s Errors) String() string {
	return tea.Prettify(s)
}

func (s Errors) GoString() string {
	return s.String()
}

func (s *Errors) SetAction(v string) *Errors {
	s.Action = &v
	return s
}

func (s *Errors) SetCode(v string) *Errors {
	s.Code = &v
	return s
}

func (s *Errors) SetMessage(v string) *Errors {
	s.Message = &v
	return s
}

func (s *Errors) SetRepeat(v int32) *Errors {
	s.Repeat = &v
	return s
}

type Exec struct {
	// Exec
	Exec *ExecAction `json:"Exec,omitempty" xml:"Exec,omitempty"`
}

func (s Exec) String() string {
	return tea.Prettify(s)
}

func (s Exec) GoString() string {
	return s.String()
}

func (s *Exec) SetExec(v *ExecAction) *Exec {
	s.Exec = v
	return s
}

type ExecAction struct {
	// Command
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
}

func (s ExecAction) String() string {
	return tea.Prettify(s)
}

func (s ExecAction) GoString() string {
	return s.String()
}

func (s *ExecAction) SetCommand(v []*string) *ExecAction {
	s.Command = v
	return s
}

type ExposedPort struct {
	// ContainerPort
	ContainerPort *int32 `json:"ContainerPort,omitempty" xml:"ContainerPort,omitempty"`
	// HostPorts
	HostPorts []*int32 `json:"HostPorts,omitempty" xml:"HostPorts,omitempty" type:"Repeated"`
	// Proto
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
}

func (s ExposedPort) String() string {
	return tea.Prettify(s)
}

func (s ExposedPort) GoString() string {
	return s.String()
}

func (s *ExposedPort) SetContainerPort(v int32) *ExposedPort {
	s.ContainerPort = &v
	return s
}

func (s *ExposedPort) SetHostPorts(v []*int32) *ExposedPort {
	s.HostPorts = v
	return s
}

func (s *ExposedPort) SetProto(v string) *ExposedPort {
	s.Proto = &v
	return s
}

type FailStrategy struct {
	// MaxRetries
	MaxRetries *int32 `json:"MaxRetries,omitempty" xml:"MaxRetries,omitempty"`
	// RunningTimeout
	RunningTimeout *int32 `json:"RunningTimeout,omitempty" xml:"RunningTimeout,omitempty"`
	// SuccessCode
	SuccessCode []*int32 `json:"SuccessCode,omitempty" xml:"SuccessCode,omitempty" type:"Repeated"`
	// WaitingTimeout
	WaitingTimeout *int32 `json:"WaitingTimeout,omitempty" xml:"WaitingTimeout,omitempty"`
}

func (s FailStrategy) String() string {
	return tea.Prettify(s)
}

func (s FailStrategy) GoString() string {
	return s.String()
}

func (s *FailStrategy) SetMaxRetries(v int32) *FailStrategy {
	s.MaxRetries = &v
	return s
}

func (s *FailStrategy) SetRunningTimeout(v int32) *FailStrategy {
	s.RunningTimeout = &v
	return s
}

func (s *FailStrategy) SetSuccessCode(v []*int32) *FailStrategy {
	s.SuccessCode = v
	return s
}

func (s *FailStrategy) SetWaitingTimeout(v int32) *FailStrategy {
	s.WaitingTimeout = &v
	return s
}

type HTTPGet struct {
	// HTTPGet
	HTTPGet *HTTPGetAction `json:"HTTPGet,omitempty" xml:"HTTPGet,omitempty"`
}

func (s HTTPGet) String() string {
	return tea.Prettify(s)
}

func (s HTTPGet) GoString() string {
	return s.String()
}

func (s *HTTPGet) SetHTTPGet(v *HTTPGetAction) *HTTPGet {
	s.HTTPGet = v
	return s
}

type HTTPGetAction struct {
	// HTTPHeaders
	HTTPHeaders []*HTTPHeader `json:"HTTPHeaders,omitempty" xml:"HTTPHeaders,omitempty" type:"Repeated"`
	// Host
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// Path
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// Port
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// Scheme
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
}

func (s HTTPGetAction) String() string {
	return tea.Prettify(s)
}

func (s HTTPGetAction) GoString() string {
	return s.String()
}

func (s *HTTPGetAction) SetHTTPHeaders(v []*HTTPHeader) *HTTPGetAction {
	s.HTTPHeaders = v
	return s
}

func (s *HTTPGetAction) SetHost(v string) *HTTPGetAction {
	s.Host = &v
	return s
}

func (s *HTTPGetAction) SetPath(v string) *HTTPGetAction {
	s.Path = &v
	return s
}

func (s *HTTPGetAction) SetPort(v int32) *HTTPGetAction {
	s.Port = &v
	return s
}

func (s *HTTPGetAction) SetScheme(v string) *HTTPGetAction {
	s.Scheme = &v
	return s
}

type HTTPHeader struct {
	// Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s HTTPHeader) String() string {
	return tea.Prettify(s)
}

func (s HTTPHeader) GoString() string {
	return s.String()
}

func (s *HTTPHeader) SetName(v string) *HTTPHeader {
	s.Name = &v
	return s
}

func (s *HTTPHeader) SetValue(v string) *HTTPHeader {
	s.Value = &v
	return s
}

type Handler struct {
	// Exec
	Exec *ExecAction `json:"Exec,omitempty" xml:"Exec,omitempty"`
	// HTTPGet
	HTTPGet *HTTPGetAction `json:"HTTPGet,omitempty" xml:"HTTPGet,omitempty"`
}

func (s Handler) String() string {
	return tea.Prettify(s)
}

func (s Handler) GoString() string {
	return s.String()
}

func (s *Handler) SetExec(v *ExecAction) *Handler {
	s.Exec = v
	return s
}

func (s *Handler) SetHTTPGet(v *HTTPGetAction) *Handler {
	s.HTTPGet = v
	return s
}

type Input struct {
	// FileMode
	FileMode *string `json:"FileMode,omitempty" xml:"FileMode,omitempty"`
	// FilePath
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	Source   *Source `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s Input) String() string {
	return tea.Prettify(s)
}

func (s Input) GoString() string {
	return s.String()
}

func (s *Input) SetFileMode(v string) *Input {
	s.FileMode = &v
	return s
}

func (s *Input) SetFilePath(v string) *Input {
	s.FilePath = &v
	return s
}

func (s *Input) SetSource(v *Source) *Input {
	s.Source = v
	return s
}

type JobDefinition struct {
	// Command
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	// CredentialConfig
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// Envs
	Envs map[string]*string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// FailStrategy
	FailStrategy *FailStrategy `json:"FailStrategy,omitempty" xml:"FailStrategy,omitempty"`
	// Inputs
	Inputs []*Input `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Repeated"`
	// Labels
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// Loggings
	Loggings *Logging `json:"Loggings,omitempty" xml:"Loggings,omitempty"`
	// MountPoints
	MountPoints []*MountPoint `json:"MountPoints,omitempty" xml:"MountPoints,omitempty" type:"Repeated"`
	// Notification
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// Outputs
	Outputs []*Output `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	// PackageUri
	PackageUri *string `json:"PackageUri,omitempty" xml:"PackageUri,omitempty"`
	// ReleaseStrategy
	ReleaseStrategy *ReleaseStrategy `json:"ReleaseStrategy,omitempty" xml:"ReleaseStrategy,omitempty"`
	// Resources
	Resources map[string]*string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// Runtimes
	Runtimes *Runtimes `json:"Runtimes,omitempty" xml:"Runtimes,omitempty"`
	// UserData
	UserData map[string]*string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// Volumes
	Volumes []*Volume `json:"Volumes,omitempty" xml:"Volumes,omitempty" type:"Repeated"`
}

func (s JobDefinition) String() string {
	return tea.Prettify(s)
}

func (s JobDefinition) GoString() string {
	return s.String()
}

func (s *JobDefinition) SetCommand(v []*string) *JobDefinition {
	s.Command = v
	return s
}

func (s *JobDefinition) SetCredentialConfig(v *CredentialConfig) *JobDefinition {
	s.CredentialConfig = v
	return s
}

func (s *JobDefinition) SetEnvs(v map[string]*string) *JobDefinition {
	s.Envs = v
	return s
}

func (s *JobDefinition) SetFailStrategy(v *FailStrategy) *JobDefinition {
	s.FailStrategy = v
	return s
}

func (s *JobDefinition) SetInputs(v []*Input) *JobDefinition {
	s.Inputs = v
	return s
}

func (s *JobDefinition) SetLabels(v map[string]*string) *JobDefinition {
	s.Labels = v
	return s
}

func (s *JobDefinition) SetLoggings(v *Logging) *JobDefinition {
	s.Loggings = v
	return s
}

func (s *JobDefinition) SetMountPoints(v []*MountPoint) *JobDefinition {
	s.MountPoints = v
	return s
}

func (s *JobDefinition) SetNotification(v *Notification) *JobDefinition {
	s.Notification = v
	return s
}

func (s *JobDefinition) SetOutputs(v []*Output) *JobDefinition {
	s.Outputs = v
	return s
}

func (s *JobDefinition) SetPackageUri(v string) *JobDefinition {
	s.PackageUri = &v
	return s
}

func (s *JobDefinition) SetReleaseStrategy(v *ReleaseStrategy) *JobDefinition {
	s.ReleaseStrategy = v
	return s
}

func (s *JobDefinition) SetResources(v map[string]*string) *JobDefinition {
	s.Resources = v
	return s
}

func (s *JobDefinition) SetRuntimes(v *Runtimes) *JobDefinition {
	s.Runtimes = v
	return s
}

func (s *JobDefinition) SetUserData(v map[string]*string) *JobDefinition {
	s.UserData = v
	return s
}

func (s *JobDefinition) SetVolumes(v []*Volume) *JobDefinition {
	s.Volumes = v
	return s
}

type JobQueueDefinition struct {
	// Labels
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// Priority
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// ProviderConfigs
	ProviderConfigs []*ProviderConfig `json:"ProviderConfigs,omitempty" xml:"ProviderConfigs,omitempty" type:"Repeated"`
	// SchedulerConfig
	SchedulerConfig *JobQueueDefinitionSchedulerConfig `json:"SchedulerConfig,omitempty" xml:"SchedulerConfig,omitempty" type:"Struct"`
}

func (s JobQueueDefinition) String() string {
	return tea.Prettify(s)
}

func (s JobQueueDefinition) GoString() string {
	return s.String()
}

func (s *JobQueueDefinition) SetLabels(v map[string]*string) *JobQueueDefinition {
	s.Labels = v
	return s
}

func (s *JobQueueDefinition) SetPriority(v int32) *JobQueueDefinition {
	s.Priority = &v
	return s
}

func (s *JobQueueDefinition) SetProviderConfigs(v []*ProviderConfig) *JobQueueDefinition {
	s.ProviderConfigs = v
	return s
}

func (s *JobQueueDefinition) SetSchedulerConfig(v *JobQueueDefinitionSchedulerConfig) *JobQueueDefinition {
	s.SchedulerConfig = v
	return s
}

type JobQueueDefinitionSchedulerConfig struct {
	// State
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s JobQueueDefinitionSchedulerConfig) String() string {
	return tea.Prettify(s)
}

func (s JobQueueDefinitionSchedulerConfig) GoString() string {
	return s.String()
}

func (s *JobQueueDefinitionSchedulerConfig) SetState(v string) *JobQueueDefinitionSchedulerConfig {
	s.State = &v
	return s
}

type JobQueueStatus struct {
	// AllocatableResources
	AllocatableResources map[string]*string `json:"AllocatableResources,omitempty" xml:"AllocatableResources,omitempty"`
	// AllocatedResources
	AllocatedResources map[string]*string `json:"AllocatedResources,omitempty" xml:"AllocatedResources,omitempty"`
	// CreateTime
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// LastUpdateTime
	LastUpdateTime *string `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
	// Managed
	Managed *bool `json:"Managed,omitempty" xml:"Managed,omitempty"`
	// ProviderStatuses
	ProviderStatuses []*ProviderStatus `json:"ProviderStatuses,omitempty" xml:"ProviderStatuses,omitempty" type:"Repeated"`
	// Reason
	Reason          *string          `json:"Reason,omitempty" xml:"Reason,omitempty"`
	SchedulerStatus *SchedulerStatus `json:"SchedulerStatus,omitempty" xml:"SchedulerStatus,omitempty"`
	// State
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s JobQueueStatus) String() string {
	return tea.Prettify(s)
}

func (s JobQueueStatus) GoString() string {
	return s.String()
}

func (s *JobQueueStatus) SetAllocatableResources(v map[string]*string) *JobQueueStatus {
	s.AllocatableResources = v
	return s
}

func (s *JobQueueStatus) SetAllocatedResources(v map[string]*string) *JobQueueStatus {
	s.AllocatedResources = v
	return s
}

func (s *JobQueueStatus) SetCreateTime(v string) *JobQueueStatus {
	s.CreateTime = &v
	return s
}

func (s *JobQueueStatus) SetLastUpdateTime(v string) *JobQueueStatus {
	s.LastUpdateTime = &v
	return s
}

func (s *JobQueueStatus) SetManaged(v bool) *JobQueueStatus {
	s.Managed = &v
	return s
}

func (s *JobQueueStatus) SetProviderStatuses(v []*ProviderStatus) *JobQueueStatus {
	s.ProviderStatuses = v
	return s
}

func (s *JobQueueStatus) SetReason(v string) *JobQueueStatus {
	s.Reason = &v
	return s
}

func (s *JobQueueStatus) SetSchedulerStatus(v *SchedulerStatus) *JobQueueStatus {
	s.SchedulerStatus = v
	return s
}

func (s *JobQueueStatus) SetState(v string) *JobQueueStatus {
	s.State = &v
	return s
}

type Logging struct {
	// Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// OSS
	OSS *OSSLogging `json:"OSS,omitempty" xml:"OSS,omitempty"`
	// Path
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// SLS
	SLS *SLSLogging `json:"SLS,omitempty" xml:"SLS,omitempty"`
}

func (s Logging) String() string {
	return tea.Prettify(s)
}

func (s Logging) GoString() string {
	return s.String()
}

func (s *Logging) SetName(v string) *Logging {
	s.Name = &v
	return s
}

func (s *Logging) SetOSS(v *OSSLogging) *Logging {
	s.OSS = v
	return s
}

func (s *Logging) SetPath(v string) *Logging {
	s.Path = &v
	return s
}

func (s *Logging) SetSLS(v *SLSLogging) *Logging {
	s.SLS = v
	return s
}

type MNSNotification struct {
	// Endpoint
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// Filters
	Filters []*string `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// Topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s MNSNotification) String() string {
	return tea.Prettify(s)
}

func (s MNSNotification) GoString() string {
	return s.String()
}

func (s *MNSNotification) SetEndpoint(v string) *MNSNotification {
	s.Endpoint = &v
	return s
}

func (s *MNSNotification) SetFilters(v []*string) *MNSNotification {
	s.Filters = v
	return s
}

func (s *MNSNotification) SetTopic(v string) *MNSNotification {
	s.Topic = &v
	return s
}

type MountPoint struct {
	// MountPath
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// ReadOnly
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// SubPath
	SubPath *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
}

func (s MountPoint) String() string {
	return tea.Prettify(s)
}

func (s MountPoint) GoString() string {
	return s.String()
}

func (s *MountPoint) SetMountPath(v string) *MountPoint {
	s.MountPath = &v
	return s
}

func (s *MountPoint) SetName(v string) *MountPoint {
	s.Name = &v
	return s
}

func (s *MountPoint) SetReadOnly(v bool) *MountPoint {
	s.ReadOnly = &v
	return s
}

func (s *MountPoint) SetSubPath(v string) *MountPoint {
	s.SubPath = &v
	return s
}

type NFSVolumeSource struct {
	// Path
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// ReadOnly
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// Server
	Server *string `json:"Server,omitempty" xml:"Server,omitempty"`
	// Version
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s NFSVolumeSource) String() string {
	return tea.Prettify(s)
}

func (s NFSVolumeSource) GoString() string {
	return s.String()
}

func (s *NFSVolumeSource) SetPath(v string) *NFSVolumeSource {
	s.Path = &v
	return s
}

func (s *NFSVolumeSource) SetReadOnly(v bool) *NFSVolumeSource {
	s.ReadOnly = &v
	return s
}

func (s *NFSVolumeSource) SetServer(v string) *NFSVolumeSource {
	s.Server = &v
	return s
}

func (s *NFSVolumeSource) SetVersion(v string) *NFSVolumeSource {
	s.Version = &v
	return s
}

type Notification struct {
	// MNS
	MNS *MNSNotification `json:"MNS,omitempty" xml:"MNS,omitempty"`
}

func (s Notification) String() string {
	return tea.Prettify(s)
}

func (s Notification) GoString() string {
	return s.String()
}

func (s *Notification) SetMNS(v *MNSNotification) *Notification {
	s.MNS = v
	return s
}

type OSSDescription struct {
	// Bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// Object
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// Prefix
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s OSSDescription) String() string {
	return tea.Prettify(s)
}

func (s OSSDescription) GoString() string {
	return s.String()
}

func (s *OSSDescription) SetBucket(v string) *OSSDescription {
	s.Bucket = &v
	return s
}

func (s *OSSDescription) SetObject(v string) *OSSDescription {
	s.Object = &v
	return s
}

func (s *OSSDescription) SetPrefix(v string) *OSSDescription {
	s.Prefix = &v
	return s
}

type OSSLogging struct {
	// Bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// Prefix
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s OSSLogging) String() string {
	return tea.Prettify(s)
}

func (s OSSLogging) GoString() string {
	return s.String()
}

func (s *OSSLogging) SetBucket(v string) *OSSLogging {
	s.Bucket = &v
	return s
}

func (s *OSSLogging) SetPrefix(v string) *OSSLogging {
	s.Prefix = &v
	return s
}

type OSSVolumeSource struct {
	// Bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// Object
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// Prefix
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// ReadOnly
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
}

func (s OSSVolumeSource) String() string {
	return tea.Prettify(s)
}

func (s OSSVolumeSource) GoString() string {
	return s.String()
}

func (s *OSSVolumeSource) SetBucket(v string) *OSSVolumeSource {
	s.Bucket = &v
	return s
}

func (s *OSSVolumeSource) SetObject(v string) *OSSVolumeSource {
	s.Object = &v
	return s
}

func (s *OSSVolumeSource) SetPrefix(v string) *OSSVolumeSource {
	s.Prefix = &v
	return s
}

func (s *OSSVolumeSource) SetReadOnly(v bool) *OSSVolumeSource {
	s.ReadOnly = &v
	return s
}

type Output struct {
	// Destination
	Destination *Destination `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// FilePattern
	FilePattern *string `json:"FilePattern,omitempty" xml:"FilePattern,omitempty"`
	// UploadConditions
	UploadConditions []*string `json:"UploadConditions,omitempty" xml:"UploadConditions,omitempty" type:"Repeated"`
	// UploadMode
	UploadMode *string `json:"UploadMode,omitempty" xml:"UploadMode,omitempty"`
}

func (s Output) String() string {
	return tea.Prettify(s)
}

func (s Output) GoString() string {
	return s.String()
}

func (s *Output) SetDestination(v *Destination) *Output {
	s.Destination = v
	return s
}

func (s *Output) SetFilePattern(v string) *Output {
	s.FilePattern = &v
	return s
}

func (s *Output) SetUploadConditions(v []*string) *Output {
	s.UploadConditions = v
	return s
}

func (s *Output) SetUploadMode(v string) *Output {
	s.UploadMode = &v
	return s
}

type PDSDescription struct {
	// Domain
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Drive
	Drive *string `json:"Drive,omitempty" xml:"Drive,omitempty"`
	// Object
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// Prefix
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s PDSDescription) String() string {
	return tea.Prettify(s)
}

func (s PDSDescription) GoString() string {
	return s.String()
}

func (s *PDSDescription) SetDomain(v string) *PDSDescription {
	s.Domain = &v
	return s
}

func (s *PDSDescription) SetDrive(v string) *PDSDescription {
	s.Drive = &v
	return s
}

func (s *PDSDescription) SetObject(v string) *PDSDescription {
	s.Object = &v
	return s
}

func (s *PDSDescription) SetPrefix(v string) *PDSDescription {
	s.Prefix = &v
	return s
}

type PDSVolumeSource struct {
	// Domain
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Drive
	Drive *string `json:"Drive,omitempty" xml:"Drive,omitempty"`
	// Object
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// Prefix
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// ReadOnly
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
}

func (s PDSVolumeSource) String() string {
	return tea.Prettify(s)
}

func (s PDSVolumeSource) GoString() string {
	return s.String()
}

func (s *PDSVolumeSource) SetDomain(v string) *PDSVolumeSource {
	s.Domain = &v
	return s
}

func (s *PDSVolumeSource) SetDrive(v string) *PDSVolumeSource {
	s.Drive = &v
	return s
}

func (s *PDSVolumeSource) SetObject(v string) *PDSVolumeSource {
	s.Object = &v
	return s
}

func (s *PDSVolumeSource) SetPrefix(v string) *PDSVolumeSource {
	s.Prefix = &v
	return s
}

func (s *PDSVolumeSource) SetReadOnly(v bool) *PDSVolumeSource {
	s.ReadOnly = &v
	return s
}

type Probe struct {
	// FailureThreshold
	FailureThreshold *int32 `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	// Handler
	Handler *Handler `json:"Handler,omitempty" xml:"Handler,omitempty"`
	// InitialDelaySeconds
	InitialDelaySeconds *int32 `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	// PeriodSeconds
	PeriodSeconds *int32 `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	// SuccessThreshold
	SuccessThreshold *int32 `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	// TimeoutSeconds
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s Probe) String() string {
	return tea.Prettify(s)
}

func (s Probe) GoString() string {
	return s.String()
}

func (s *Probe) SetFailureThreshold(v int32) *Probe {
	s.FailureThreshold = &v
	return s
}

func (s *Probe) SetHandler(v *Handler) *Probe {
	s.Handler = v
	return s
}

func (s *Probe) SetInitialDelaySeconds(v int32) *Probe {
	s.InitialDelaySeconds = &v
	return s
}

func (s *Probe) SetPeriodSeconds(v int32) *Probe {
	s.PeriodSeconds = &v
	return s
}

func (s *Probe) SetSuccessThreshold(v int32) *Probe {
	s.SuccessThreshold = &v
	return s
}

func (s *Probe) SetTimeoutSeconds(v int32) *Probe {
	s.TimeoutSeconds = &v
	return s
}

type ProjectDefinition struct {
	// JobLifecycle
	JobLifecycle *int32 `json:"JobLifecycle,omitempty" xml:"JobLifecycle,omitempty"`
	// Labels
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// Role
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s ProjectDefinition) String() string {
	return tea.Prettify(s)
}

func (s ProjectDefinition) GoString() string {
	return s.String()
}

func (s *ProjectDefinition) SetJobLifecycle(v int32) *ProjectDefinition {
	s.JobLifecycle = &v
	return s
}

func (s *ProjectDefinition) SetLabels(v map[string]*string) *ProjectDefinition {
	s.Labels = v
	return s
}

func (s *ProjectDefinition) SetRole(v string) *ProjectDefinition {
	s.Role = &v
	return s
}

type ProviderConfig struct {
	// MaxWorkerCount
	MaxWorkerCount *int32 `json:"MaxWorkerCount,omitempty" xml:"MaxWorkerCount,omitempty"`
	// MinWorkerCount
	MinWorkerCount *int32 `json:"MinWorkerCount,omitempty" xml:"MinWorkerCount,omitempty"`
	// ProviderId
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	// ProviderType
	ProviderType *string `json:"ProviderType,omitempty" xml:"ProviderType,omitempty"`
	// WorkerType
	WorkerType *string `json:"WorkerType,omitempty" xml:"WorkerType,omitempty"`
}

func (s ProviderConfig) String() string {
	return tea.Prettify(s)
}

func (s ProviderConfig) GoString() string {
	return s.String()
}

func (s *ProviderConfig) SetMaxWorkerCount(v int32) *ProviderConfig {
	s.MaxWorkerCount = &v
	return s
}

func (s *ProviderConfig) SetMinWorkerCount(v int32) *ProviderConfig {
	s.MinWorkerCount = &v
	return s
}

func (s *ProviderConfig) SetProviderId(v string) *ProviderConfig {
	s.ProviderId = &v
	return s
}

func (s *ProviderConfig) SetProviderType(v string) *ProviderConfig {
	s.ProviderType = &v
	return s
}

func (s *ProviderConfig) SetWorkerType(v string) *ProviderConfig {
	s.WorkerType = &v
	return s
}

type ProviderStatus struct {
	// AllocatableResources
	AllocatableResources map[string]*string `json:"AllocatableResources,omitempty" xml:"AllocatableResources,omitempty"`
	// AllocatedResources
	AllocatedResources map[string]*string `json:"AllocatedResources,omitempty" xml:"AllocatedResources,omitempty"`
	// ProviderId
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
}

func (s ProviderStatus) String() string {
	return tea.Prettify(s)
}

func (s ProviderStatus) GoString() string {
	return s.String()
}

func (s *ProviderStatus) SetAllocatableResources(v map[string]*string) *ProviderStatus {
	s.AllocatableResources = v
	return s
}

func (s *ProviderStatus) SetAllocatedResources(v map[string]*string) *ProviderStatus {
	s.AllocatedResources = v
	return s
}

func (s *ProviderStatus) SetProviderId(v string) *ProviderStatus {
	s.ProviderId = &v
	return s
}

type ReleaseCondition struct {
	// State
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// TTLSeconds
	TTLSeconds *int64 `json:"TTLSeconds,omitempty" xml:"TTLSeconds,omitempty"`
}

func (s ReleaseCondition) String() string {
	return tea.Prettify(s)
}

func (s ReleaseCondition) GoString() string {
	return s.String()
}

func (s *ReleaseCondition) SetState(v string) *ReleaseCondition {
	s.State = &v
	return s
}

func (s *ReleaseCondition) SetTTLSeconds(v int64) *ReleaseCondition {
	s.TTLSeconds = &v
	return s
}

type ReleaseStrategy struct {
	// ReleaseConditions
	ReleaseConditions []*ReleaseCondition `json:"ReleaseConditions,omitempty" xml:"ReleaseConditions,omitempty" type:"Repeated"`
}

func (s ReleaseStrategy) String() string {
	return tea.Prettify(s)
}

func (s ReleaseStrategy) GoString() string {
	return s.String()
}

func (s *ReleaseStrategy) SetReleaseConditions(v []*ReleaseCondition) *ReleaseStrategy {
	s.ReleaseConditions = v
	return s
}

type Runtimes struct {
	// Docker
	Docker *Docker `json:"Docker,omitempty" xml:"Docker,omitempty"`
	// ECS
	ECS *ECS `json:"ECS,omitempty" xml:"ECS,omitempty"`
	// JobQueue
	JobQueue *string `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	// VPC
	VPC *VPC `json:"VPC,omitempty" xml:"VPC,omitempty"`
}

func (s Runtimes) String() string {
	return tea.Prettify(s)
}

func (s Runtimes) GoString() string {
	return s.String()
}

func (s *Runtimes) SetDocker(v *Docker) *Runtimes {
	s.Docker = v
	return s
}

func (s *Runtimes) SetECS(v *ECS) *Runtimes {
	s.ECS = v
	return s
}

func (s *Runtimes) SetJobQueue(v string) *Runtimes {
	s.JobQueue = &v
	return s
}

func (s *Runtimes) SetVPC(v *VPC) *Runtimes {
	s.VPC = v
	return s
}

type SLB struct {
	// SLBId
	SLBId *string `json:"SLBId,omitempty" xml:"SLBId,omitempty"`
}

func (s SLB) String() string {
	return tea.Prettify(s)
}

func (s SLB) GoString() string {
	return s.String()
}

func (s *SLB) SetSLBId(v string) *SLB {
	s.SLBId = &v
	return s
}

type SLSLogging struct {
	// LogtailConfigName
	LogtailConfigName *string `json:"LogtailConfigName,omitempty" xml:"LogtailConfigName,omitempty"`
	// Project
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// Store
	Store *string `json:"Store,omitempty" xml:"Store,omitempty"`
}

func (s SLSLogging) String() string {
	return tea.Prettify(s)
}

func (s SLSLogging) GoString() string {
	return s.String()
}

func (s *SLSLogging) SetLogtailConfigName(v string) *SLSLogging {
	s.LogtailConfigName = &v
	return s
}

func (s *SLSLogging) SetProject(v string) *SLSLogging {
	s.Project = &v
	return s
}

func (s *SLSLogging) SetStore(v string) *SLSLogging {
	s.Store = &v
	return s
}

type Scaling struct {
	// AdjustmentType
	AdjustmentType *string `json:"AdjustmentType,omitempty" xml:"AdjustmentType,omitempty"`
	// AdjustmentValue
	AdjustmentValue *float32 `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	// MaxWorkerCount
	MaxWorkerCount *int32 `json:"MaxWorkerCount,omitempty" xml:"MaxWorkerCount,omitempty"`
	// MetricType
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// MetricValue
	MetricValue *float32 `json:"MetricValue,omitempty" xml:"MetricValue,omitempty"`
	// MinWorkerCount
	MinWorkerCount *int32 `json:"MinWorkerCount,omitempty" xml:"MinWorkerCount,omitempty"`
	// ToleranceValue
	ToleranceValue *float32 `json:"ToleranceValue,omitempty" xml:"ToleranceValue,omitempty"`
}

func (s Scaling) String() string {
	return tea.Prettify(s)
}

func (s Scaling) GoString() string {
	return s.String()
}

func (s *Scaling) SetAdjustmentType(v string) *Scaling {
	s.AdjustmentType = &v
	return s
}

func (s *Scaling) SetAdjustmentValue(v float32) *Scaling {
	s.AdjustmentValue = &v
	return s
}

func (s *Scaling) SetMaxWorkerCount(v int32) *Scaling {
	s.MaxWorkerCount = &v
	return s
}

func (s *Scaling) SetMetricType(v string) *Scaling {
	s.MetricType = &v
	return s
}

func (s *Scaling) SetMetricValue(v float32) *Scaling {
	s.MetricValue = &v
	return s
}

func (s *Scaling) SetMinWorkerCount(v int32) *Scaling {
	s.MinWorkerCount = &v
	return s
}

func (s *Scaling) SetToleranceValue(v float32) *Scaling {
	s.ToleranceValue = &v
	return s
}

type SchedulerStatus struct {
	// CanceledJobCount
	CanceledJobCount *int64 `json:"CanceledJobCount,omitempty" xml:"CanceledJobCount,omitempty"`
	// FailedJobCount
	FailedJobCount *int64 `json:"FailedJobCount,omitempty" xml:"FailedJobCount,omitempty"`
	// RunningJobCount
	RunningJobCount *int64 `json:"RunningJobCount,omitempty" xml:"RunningJobCount,omitempty"`
	// SucceededJobCount
	SucceededJobCount *int64 `json:"SucceededJobCount,omitempty" xml:"SucceededJobCount,omitempty"`
	// WaitingJobCount
	WaitingJobCount *int64 `json:"WaitingJobCount,omitempty" xml:"WaitingJobCount,omitempty"`
}

func (s SchedulerStatus) String() string {
	return tea.Prettify(s)
}

func (s SchedulerStatus) GoString() string {
	return s.String()
}

func (s *SchedulerStatus) SetCanceledJobCount(v int64) *SchedulerStatus {
	s.CanceledJobCount = &v
	return s
}

func (s *SchedulerStatus) SetFailedJobCount(v int64) *SchedulerStatus {
	s.FailedJobCount = &v
	return s
}

func (s *SchedulerStatus) SetRunningJobCount(v int64) *SchedulerStatus {
	s.RunningJobCount = &v
	return s
}

func (s *SchedulerStatus) SetSucceededJobCount(v int64) *SchedulerStatus {
	s.SucceededJobCount = &v
	return s
}

func (s *SchedulerStatus) SetWaitingJobCount(v int64) *SchedulerStatus {
	s.WaitingJobCount = &v
	return s
}

type ServiceRoleNode struct {
	// AssumeRoleFor
	AssumeRoleFor *string `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	// Role
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// RoleType
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s ServiceRoleNode) String() string {
	return tea.Prettify(s)
}

func (s ServiceRoleNode) GoString() string {
	return s.String()
}

func (s *ServiceRoleNode) SetAssumeRoleFor(v string) *ServiceRoleNode {
	s.AssumeRoleFor = &v
	return s
}

func (s *ServiceRoleNode) SetRole(v string) *ServiceRoleNode {
	s.Role = &v
	return s
}

func (s *ServiceRoleNode) SetRoleType(v string) *ServiceRoleNode {
	s.RoleType = &v
	return s
}

type Source struct {
	OSS *OSSDescription `json:"OSS,omitempty" xml:"OSS,omitempty"`
	PDS *PDSDescription `json:"PDS,omitempty" xml:"PDS,omitempty"`
}

func (s Source) String() string {
	return tea.Prettify(s)
}

func (s Source) GoString() string {
	return s.String()
}

func (s *Source) SetOSS(v *OSSDescription) *Source {
	s.OSS = v
	return s
}

func (s *Source) SetPDS(v *PDSDescription) *Source {
	s.PDS = v
	return s
}

type Trigger struct {
	// Enabled
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// FirstLaunchTime
	FirstLaunchTime *string `json:"FirstLaunchTime,omitempty" xml:"FirstLaunchTime,omitempty"`
	// RepeatType
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// RepeatValue
	RepeatValue *string `json:"RepeatValue,omitempty" xml:"RepeatValue,omitempty"`
}

func (s Trigger) String() string {
	return tea.Prettify(s)
}

func (s Trigger) GoString() string {
	return s.String()
}

func (s *Trigger) SetEnabled(v bool) *Trigger {
	s.Enabled = &v
	return s
}

func (s *Trigger) SetFirstLaunchTime(v string) *Trigger {
	s.FirstLaunchTime = &v
	return s
}

func (s *Trigger) SetRepeatType(v string) *Trigger {
	s.RepeatType = &v
	return s
}

func (s *Trigger) SetRepeatValue(v string) *Trigger {
	s.RepeatValue = &v
	return s
}

type UpgradePolicy struct {
	// UpgradeRatio
	UpgradeRatio *float32 `json:"UpgradeRatio,omitempty" xml:"UpgradeRatio,omitempty"`
}

func (s UpgradePolicy) String() string {
	return tea.Prettify(s)
}

func (s UpgradePolicy) GoString() string {
	return s.String()
}

func (s *UpgradePolicy) SetUpgradeRatio(v float32) *UpgradePolicy {
	s.UpgradeRatio = &v
	return s
}

type UserStage struct {
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// EndTime
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// StartTime
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// State
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s UserStage) String() string {
	return tea.Prettify(s)
}

func (s UserStage) GoString() string {
	return s.String()
}

func (s *UserStage) SetDescription(v string) *UserStage {
	s.Description = &v
	return s
}

func (s *UserStage) SetEndTime(v int32) *UserStage {
	s.EndTime = &v
	return s
}

func (s *UserStage) SetStartTime(v int32) *UserStage {
	s.StartTime = &v
	return s
}

func (s *UserStage) SetState(v string) *UserStage {
	s.State = &v
	return s
}

type VPC struct {
	// SecurityGroups
	SecurityGroups []*string `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty" type:"Repeated"`
	// VPCId
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// VSwitches
	VSwitches []*string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
}

func (s VPC) String() string {
	return tea.Prettify(s)
}

func (s VPC) GoString() string {
	return s.String()
}

func (s *VPC) SetSecurityGroups(v []*string) *VPC {
	s.SecurityGroups = v
	return s
}

func (s *VPC) SetVPCId(v string) *VPC {
	s.VPCId = &v
	return s
}

func (s *VPC) SetVSwitches(v []*string) *VPC {
	s.VSwitches = v
	return s
}

type Volume struct {
	// NFS
	NFS *NFSVolumeSource `json:"NFS,omitempty" xml:"NFS,omitempty"`
	// Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// OSS
	OSS *OSSVolumeSource `json:"OSS,omitempty" xml:"OSS,omitempty"`
	// PDS
	PDS *PDSVolumeSource `json:"PDS,omitempty" xml:"PDS,omitempty"`
}

func (s Volume) String() string {
	return tea.Prettify(s)
}

func (s Volume) GoString() string {
	return s.String()
}

func (s *Volume) SetNFS(v *NFSVolumeSource) *Volume {
	s.NFS = v
	return s
}

func (s *Volume) SetName(v string) *Volume {
	s.Name = &v
	return s
}

func (s *Volume) SetOSS(v *OSSVolumeSource) *Volume {
	s.OSS = v
	return s
}

func (s *Volume) SetPDS(v *PDSVolumeSource) *Volume {
	s.PDS = v
	return s
}

type WorkerStatus struct {
	// Conditions
	Conditions []*Conditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// Container
	Container *WorkerStatusContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	// CreateTime
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// ECS
	ECS *WorkerStatusECS `json:"ECS,omitempty" xml:"ECS,omitempty" type:"Struct"`
	// JobQueue
	JobQueue *string `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	// NetworkInterfaceId
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// PoolWorkerId
	PoolWorkerId *string `json:"PoolWorkerId,omitempty" xml:"PoolWorkerId,omitempty"`
	// SecurityGroupId
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// State
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// VSwitchId
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// WorkerType
	WorkerType *int32 `json:"WorkerType,omitempty" xml:"WorkerType,omitempty"`
}

func (s WorkerStatus) String() string {
	return tea.Prettify(s)
}

func (s WorkerStatus) GoString() string {
	return s.String()
}

func (s *WorkerStatus) SetConditions(v []*Conditions) *WorkerStatus {
	s.Conditions = v
	return s
}

func (s *WorkerStatus) SetContainer(v *WorkerStatusContainer) *WorkerStatus {
	s.Container = v
	return s
}

func (s *WorkerStatus) SetCreateTime(v string) *WorkerStatus {
	s.CreateTime = &v
	return s
}

func (s *WorkerStatus) SetECS(v *WorkerStatusECS) *WorkerStatus {
	s.ECS = v
	return s
}

func (s *WorkerStatus) SetJobQueue(v string) *WorkerStatus {
	s.JobQueue = &v
	return s
}

func (s *WorkerStatus) SetNetworkInterfaceId(v string) *WorkerStatus {
	s.NetworkInterfaceId = &v
	return s
}

func (s *WorkerStatus) SetPoolWorkerId(v string) *WorkerStatus {
	s.PoolWorkerId = &v
	return s
}

func (s *WorkerStatus) SetSecurityGroupId(v string) *WorkerStatus {
	s.SecurityGroupId = &v
	return s
}

func (s *WorkerStatus) SetState(v string) *WorkerStatus {
	s.State = &v
	return s
}

func (s *WorkerStatus) SetVSwitchId(v string) *WorkerStatus {
	s.VSwitchId = &v
	return s
}

func (s *WorkerStatus) SetWorkerType(v int32) *WorkerStatus {
	s.WorkerType = &v
	return s
}

type WorkerStatusContainer struct {
	// Cpu
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// Memory
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s WorkerStatusContainer) String() string {
	return tea.Prettify(s)
}

func (s WorkerStatusContainer) GoString() string {
	return s.String()
}

func (s *WorkerStatusContainer) SetCpu(v int32) *WorkerStatusContainer {
	s.Cpu = &v
	return s
}

func (s *WorkerStatusContainer) SetMemory(v int32) *WorkerStatusContainer {
	s.Memory = &v
	return s
}

type WorkerStatusECS struct {
	// Cpu
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// Endpoint
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// Hostname
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// InstanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// InstanceType
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Memory
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// Password
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// ResourceType
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// SpotPriceLimit
	SpotPriceLimit *string `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// SpotStrategy
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// SystemDiskSize
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// SystemDiskType
	SystemDiskType *string `json:"SystemDiskType,omitempty" xml:"SystemDiskType,omitempty"`
	// VMImage
	VMImage *string `json:"VMImage,omitempty" xml:"VMImage,omitempty"`
	// ZoneId
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s WorkerStatusECS) String() string {
	return tea.Prettify(s)
}

func (s WorkerStatusECS) GoString() string {
	return s.String()
}

func (s *WorkerStatusECS) SetCpu(v int32) *WorkerStatusECS {
	s.Cpu = &v
	return s
}

func (s *WorkerStatusECS) SetEndpoint(v string) *WorkerStatusECS {
	s.Endpoint = &v
	return s
}

func (s *WorkerStatusECS) SetHostname(v string) *WorkerStatusECS {
	s.Hostname = &v
	return s
}

func (s *WorkerStatusECS) SetInstanceId(v string) *WorkerStatusECS {
	s.InstanceId = &v
	return s
}

func (s *WorkerStatusECS) SetInstanceType(v string) *WorkerStatusECS {
	s.InstanceType = &v
	return s
}

func (s *WorkerStatusECS) SetMemory(v int32) *WorkerStatusECS {
	s.Memory = &v
	return s
}

func (s *WorkerStatusECS) SetPassword(v string) *WorkerStatusECS {
	s.Password = &v
	return s
}

func (s *WorkerStatusECS) SetResourceType(v string) *WorkerStatusECS {
	s.ResourceType = &v
	return s
}

func (s *WorkerStatusECS) SetSpotPriceLimit(v string) *WorkerStatusECS {
	s.SpotPriceLimit = &v
	return s
}

func (s *WorkerStatusECS) SetSpotStrategy(v string) *WorkerStatusECS {
	s.SpotStrategy = &v
	return s
}

func (s *WorkerStatusECS) SetSystemDiskSize(v int32) *WorkerStatusECS {
	s.SystemDiskSize = &v
	return s
}

func (s *WorkerStatusECS) SetSystemDiskType(v string) *WorkerStatusECS {
	s.SystemDiskType = &v
	return s
}

func (s *WorkerStatusECS) SetVMImage(v string) *WorkerStatusECS {
	s.VMImage = &v
	return s
}

func (s *WorkerStatusECS) SetZoneId(v string) *WorkerStatusECS {
	s.ZoneId = &v
	return s
}

type CancelJobRequest struct {
	ExitCode *string `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	JobId    *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Reason   *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s CancelJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelJobRequest) GoString() string {
	return s.String()
}

func (s *CancelJobRequest) SetExitCode(v string) *CancelJobRequest {
	s.ExitCode = &v
	return s
}

func (s *CancelJobRequest) SetJobId(v string) *CancelJobRequest {
	s.JobId = &v
	return s
}

func (s *CancelJobRequest) SetProject(v string) *CancelJobRequest {
	s.Project = &v
	return s
}

func (s *CancelJobRequest) SetReason(v string) *CancelJobRequest {
	s.Reason = &v
	return s
}

type CancelJobResponseBody struct {
	HostId    *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelJobResponseBody) GoString() string {
	return s.String()
}

func (s *CancelJobResponseBody) SetHostId(v string) *CancelJobResponseBody {
	s.HostId = &v
	return s
}

func (s *CancelJobResponseBody) SetRequestId(v string) *CancelJobResponseBody {
	s.RequestId = &v
	return s
}

type CancelJobResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelJobResponse) GoString() string {
	return s.String()
}

func (s *CancelJobResponse) SetHeaders(v map[string]*string) *CancelJobResponse {
	s.Headers = v
	return s
}

func (s *CancelJobResponse) SetBody(v *CancelJobResponseBody) *CancelJobResponse {
	s.Body = v
	return s
}

type CreateClusterRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Definition
	Definition *ClusterDefinition `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Project
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s CreateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) SetClientToken(v string) *CreateClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateClusterRequest) SetDefinition(v *ClusterDefinition) *CreateClusterRequest {
	s.Definition = v
	return s
}

func (s *CreateClusterRequest) SetDescription(v string) *CreateClusterRequest {
	s.Description = &v
	return s
}

func (s *CreateClusterRequest) SetName(v string) *CreateClusterRequest {
	s.Name = &v
	return s
}

func (s *CreateClusterRequest) SetProject(v string) *CreateClusterRequest {
	s.Project = &v
	return s
}

type CreateClusterShrinkRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Definition
	DefinitionShrink *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Project
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s CreateClusterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterShrinkRequest) SetClientToken(v string) *CreateClusterShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetDefinitionShrink(v string) *CreateClusterShrinkRequest {
	s.DefinitionShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetDescription(v string) *CreateClusterShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetName(v string) *CreateClusterShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetProject(v string) *CreateClusterShrinkRequest {
	s.Project = &v
	return s
}

type CreateClusterResponseBody struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	HostId    *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) SetClusterId(v string) *CreateClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterResponseBody) SetHostId(v string) *CreateClusterResponseBody {
	s.HostId = &v
	return s
}

func (s *CreateClusterResponseBody) SetRequestId(v string) *CreateClusterResponseBody {
	s.RequestId = &v
	return s
}

type CreateClusterResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterResponse) SetHeaders(v map[string]*string) *CreateClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateClusterResponse) SetBody(v *CreateClusterResponseBody) *CreateClusterResponse {
	s.Body = v
	return s
}

type CreateJobRequest struct {
	ClientToken *string        `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Definition  *JobDefinition `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description *string        `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string        `json:"Name,omitempty" xml:"Name,omitempty"`
	Project     *string        `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s CreateJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequest) GoString() string {
	return s.String()
}

func (s *CreateJobRequest) SetClientToken(v string) *CreateJobRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateJobRequest) SetDefinition(v *JobDefinition) *CreateJobRequest {
	s.Definition = v
	return s
}

func (s *CreateJobRequest) SetDescription(v string) *CreateJobRequest {
	s.Description = &v
	return s
}

func (s *CreateJobRequest) SetName(v string) *CreateJobRequest {
	s.Name = &v
	return s
}

func (s *CreateJobRequest) SetProject(v string) *CreateJobRequest {
	s.Project = &v
	return s
}

type CreateJobShrinkRequest struct {
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DefinitionShrink *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Project          *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s CreateJobShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateJobShrinkRequest) SetClientToken(v string) *CreateJobShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateJobShrinkRequest) SetDefinitionShrink(v string) *CreateJobShrinkRequest {
	s.DefinitionShrink = &v
	return s
}

func (s *CreateJobShrinkRequest) SetDescription(v string) *CreateJobShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateJobShrinkRequest) SetName(v string) *CreateJobShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateJobShrinkRequest) SetProject(v string) *CreateJobShrinkRequest {
	s.Project = &v
	return s
}

type CreateJobResponseBody struct {
	HostId    *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobResponseBody) SetHostId(v string) *CreateJobResponseBody {
	s.HostId = &v
	return s
}

func (s *CreateJobResponseBody) SetJobId(v string) *CreateJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateJobResponseBody) SetRequestId(v string) *CreateJobResponseBody {
	s.RequestId = &v
	return s
}

type CreateJobResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateJobResponse) GoString() string {
	return s.String()
}

func (s *CreateJobResponse) SetHeaders(v map[string]*string) *CreateJobResponse {
	s.Headers = v
	return s
}

func (s *CreateJobResponse) SetBody(v *CreateJobResponseBody) *CreateJobResponse {
	s.Body = v
	return s
}

type CreateJobQueueRequest struct {
	ClientToken *string             `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Definition  *JobQueueDefinition `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description *string             `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string             `json:"Name,omitempty" xml:"Name,omitempty"`
	Project     *string             `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s CreateJobQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateJobQueueRequest) GoString() string {
	return s.String()
}

func (s *CreateJobQueueRequest) SetClientToken(v string) *CreateJobQueueRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateJobQueueRequest) SetDefinition(v *JobQueueDefinition) *CreateJobQueueRequest {
	s.Definition = v
	return s
}

func (s *CreateJobQueueRequest) SetDescription(v string) *CreateJobQueueRequest {
	s.Description = &v
	return s
}

func (s *CreateJobQueueRequest) SetName(v string) *CreateJobQueueRequest {
	s.Name = &v
	return s
}

func (s *CreateJobQueueRequest) SetProject(v string) *CreateJobQueueRequest {
	s.Project = &v
	return s
}

type CreateJobQueueShrinkRequest struct {
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DefinitionShrink *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Project          *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s CreateJobQueueShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateJobQueueShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateJobQueueShrinkRequest) SetClientToken(v string) *CreateJobQueueShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateJobQueueShrinkRequest) SetDefinitionShrink(v string) *CreateJobQueueShrinkRequest {
	s.DefinitionShrink = &v
	return s
}

func (s *CreateJobQueueShrinkRequest) SetDescription(v string) *CreateJobQueueShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateJobQueueShrinkRequest) SetName(v string) *CreateJobQueueShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateJobQueueShrinkRequest) SetProject(v string) *CreateJobQueueShrinkRequest {
	s.Project = &v
	return s
}

type CreateJobQueueResponseBody struct {
	HostId    *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateJobQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateJobQueueResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobQueueResponseBody) SetHostId(v string) *CreateJobQueueResponseBody {
	s.HostId = &v
	return s
}

func (s *CreateJobQueueResponseBody) SetName(v string) *CreateJobQueueResponseBody {
	s.Name = &v
	return s
}

func (s *CreateJobQueueResponseBody) SetRequestId(v string) *CreateJobQueueResponseBody {
	s.RequestId = &v
	return s
}

type CreateJobQueueResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateJobQueueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateJobQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateJobQueueResponse) GoString() string {
	return s.String()
}

func (s *CreateJobQueueResponse) SetHeaders(v map[string]*string) *CreateJobQueueResponse {
	s.Headers = v
	return s
}

func (s *CreateJobQueueResponse) SetBody(v *CreateJobQueueResponseBody) *CreateJobQueueResponse {
	s.Body = v
	return s
}

type CreateProjectRequest struct {
	ClientToken *string            `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Definition  *ProjectDefinition `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description *string            `json:"Description,omitempty" xml:"Description,omitempty"`
	Project     *string            `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetClientToken(v string) *CreateProjectRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateProjectRequest) SetDefinition(v *ProjectDefinition) *CreateProjectRequest {
	s.Definition = v
	return s
}

func (s *CreateProjectRequest) SetDescription(v string) *CreateProjectRequest {
	s.Description = &v
	return s
}

func (s *CreateProjectRequest) SetProject(v string) *CreateProjectRequest {
	s.Project = &v
	return s
}

type CreateProjectShrinkRequest struct {
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DefinitionShrink *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Project          *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s CreateProjectShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectShrinkRequest) SetClientToken(v string) *CreateProjectShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetDefinitionShrink(v string) *CreateProjectShrinkRequest {
	s.DefinitionShrink = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetDescription(v string) *CreateProjectShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetProject(v string) *CreateProjectShrinkRequest {
	s.Project = &v
	return s
}

type CreateProjectResponseBody struct {
	HostId    *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	Project   *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) SetHostId(v string) *CreateProjectResponseBody {
	s.HostId = &v
	return s
}

func (s *CreateProjectResponseBody) SetProject(v string) *CreateProjectResponseBody {
	s.Project = &v
	return s
}

func (s *CreateProjectResponseBody) SetRequestId(v string) *CreateProjectResponseBody {
	s.RequestId = &v
	return s
}

type CreateProjectResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectResponse) SetHeaders(v map[string]*string) *CreateProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectResponse) SetBody(v *CreateProjectResponseBody) *CreateProjectResponse {
	s.Body = v
	return s
}

type DeleteClusterRequest struct {
	// ClusterId
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Project
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s DeleteClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterRequest) SetClusterId(v string) *DeleteClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteClusterRequest) SetProject(v string) *DeleteClusterRequest {
	s.Project = &v
	return s
}

type DeleteClusterResponseBody struct {
	HostId    *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponseBody) SetHostId(v string) *DeleteClusterResponseBody {
	s.HostId = &v
	return s
}

func (s *DeleteClusterResponseBody) SetRequestId(v string) *DeleteClusterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteClusterResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponse) SetHeaders(v map[string]*string) *DeleteClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteClusterResponse) SetBody(v *DeleteClusterResponseBody) *DeleteClusterResponse {
	s.Body = v
	return s
}

type DeleteJobRequest struct {
	JobId   *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s DeleteJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobRequest) SetJobId(v string) *DeleteJobRequest {
	s.JobId = &v
	return s
}

func (s *DeleteJobRequest) SetProject(v string) *DeleteJobRequest {
	s.Project = &v
	return s
}

type DeleteJobResponseBody struct {
	HostId    *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobResponseBody) SetHostId(v string) *DeleteJobResponseBody {
	s.HostId = &v
	return s
}

func (s *DeleteJobResponseBody) SetRequestId(v string) *DeleteJobResponseBody {
	s.RequestId = &v
	return s
}

type DeleteJobResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteJobResponse) SetHeaders(v map[string]*string) *DeleteJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteJobResponse) SetBody(v *DeleteJobResponseBody) *DeleteJobResponse {
	s.Body = v
	return s
}

type DeleteJobQueueRequest struct {
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s DeleteJobQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobQueueRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobQueueRequest) SetName(v string) *DeleteJobQueueRequest {
	s.Name = &v
	return s
}

func (s *DeleteJobQueueRequest) SetProject(v string) *DeleteJobQueueRequest {
	s.Project = &v
	return s
}

type DeleteJobQueueResponseBody struct {
	HostId    *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteJobQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobQueueResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobQueueResponseBody) SetHostId(v string) *DeleteJobQueueResponseBody {
	s.HostId = &v
	return s
}

func (s *DeleteJobQueueResponseBody) SetRequestId(v string) *DeleteJobQueueResponseBody {
	s.RequestId = &v
	return s
}

type DeleteJobQueueResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteJobQueueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteJobQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobQueueResponse) GoString() string {
	return s.String()
}

func (s *DeleteJobQueueResponse) SetHeaders(v map[string]*string) *DeleteJobQueueResponse {
	s.Headers = v
	return s
}

func (s *DeleteJobQueueResponse) SetBody(v *DeleteJobQueueResponseBody) *DeleteJobQueueResponse {
	s.Body = v
	return s
}

type DeleteProjectRequest struct {
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s DeleteProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectRequest) SetProject(v string) *DeleteProjectRequest {
	s.Project = &v
	return s
}

type DeleteProjectResponseBody struct {
	HostId    *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponseBody) SetHostId(v string) *DeleteProjectResponseBody {
	s.HostId = &v
	return s
}

func (s *DeleteProjectResponseBody) SetRequestId(v string) *DeleteProjectResponseBody {
	s.RequestId = &v
	return s
}

type DeleteProjectResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponse) SetHeaders(v map[string]*string) *DeleteProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteProjectResponse) SetBody(v *DeleteProjectResponseBody) *DeleteProjectResponse {
	s.Body = v
	return s
}

type GetClusterRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Project   *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s GetClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRequest) GoString() string {
	return s.String()
}

func (s *GetClusterRequest) SetClusterId(v string) *GetClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *GetClusterRequest) SetProject(v string) *GetClusterRequest {
	s.Project = &v
	return s
}

type GetClusterResponseBody struct {
	ClusterId   *string                       `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Definition  *ClusterDefinition            `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description *string                       `json:"Description,omitempty" xml:"Description,omitempty"`
	HostId      *string                       `json:"HostId,omitempty" xml:"HostId,omitempty"`
	Name        *string                       `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId     *string                       `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Project     *string                       `json:"Project,omitempty" xml:"Project,omitempty"`
	RequestId   *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status      *GetClusterResponseBodyStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
}

func (s GetClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBody) SetClusterId(v string) *GetClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *GetClusterResponseBody) SetDefinition(v *ClusterDefinition) *GetClusterResponseBody {
	s.Definition = v
	return s
}

func (s *GetClusterResponseBody) SetDescription(v string) *GetClusterResponseBody {
	s.Description = &v
	return s
}

func (s *GetClusterResponseBody) SetHostId(v string) *GetClusterResponseBody {
	s.HostId = &v
	return s
}

func (s *GetClusterResponseBody) SetName(v string) *GetClusterResponseBody {
	s.Name = &v
	return s
}

func (s *GetClusterResponseBody) SetOwnerId(v string) *GetClusterResponseBody {
	s.OwnerId = &v
	return s
}

func (s *GetClusterResponseBody) SetProject(v string) *GetClusterResponseBody {
	s.Project = &v
	return s
}

func (s *GetClusterResponseBody) SetRequestId(v string) *GetClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterResponseBody) SetStatus(v *GetClusterResponseBodyStatus) *GetClusterResponseBody {
	s.Status = v
	return s
}

type GetClusterResponseBodyStatus struct {
	AllocatableResources map[string]*string `json:"AllocatableResources,omitempty" xml:"AllocatableResources,omitempty"`
	AllocatedResources   map[string]*string `json:"AllocatedResources,omitempty" xml:"AllocatedResources,omitempty"`
	Conditions           []*Conditions      `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	CreateTime           *string            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CurrentWorkerCount   *int32             `json:"CurrentWorkerCount,omitempty" xml:"CurrentWorkerCount,omitempty"`
	DesiredWorkerCount   *int32             `json:"DesiredWorkerCount,omitempty" xml:"DesiredWorkerCount,omitempty"`
	State                *string            `json:"State,omitempty" xml:"State,omitempty"`
}

func (s GetClusterResponseBodyStatus) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBodyStatus) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyStatus) SetAllocatableResources(v map[string]*string) *GetClusterResponseBodyStatus {
	s.AllocatableResources = v
	return s
}

func (s *GetClusterResponseBodyStatus) SetAllocatedResources(v map[string]*string) *GetClusterResponseBodyStatus {
	s.AllocatedResources = v
	return s
}

func (s *GetClusterResponseBodyStatus) SetConditions(v []*Conditions) *GetClusterResponseBodyStatus {
	s.Conditions = v
	return s
}

func (s *GetClusterResponseBodyStatus) SetCreateTime(v string) *GetClusterResponseBodyStatus {
	s.CreateTime = &v
	return s
}

func (s *GetClusterResponseBodyStatus) SetCurrentWorkerCount(v int32) *GetClusterResponseBodyStatus {
	s.CurrentWorkerCount = &v
	return s
}

func (s *GetClusterResponseBodyStatus) SetDesiredWorkerCount(v int32) *GetClusterResponseBodyStatus {
	s.DesiredWorkerCount = &v
	return s
}

func (s *GetClusterResponseBodyStatus) SetState(v string) *GetClusterResponseBodyStatus {
	s.State = &v
	return s
}

type GetClusterResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponse) GoString() string {
	return s.String()
}

func (s *GetClusterResponse) SetHeaders(v map[string]*string) *GetClusterResponse {
	s.Headers = v
	return s
}

func (s *GetClusterResponse) SetBody(v *GetClusterResponseBody) *GetClusterResponse {
	s.Body = v
	return s
}

type GetJobRequest struct {
	JobId   *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s GetJobRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobRequest) GoString() string {
	return s.String()
}

func (s *GetJobRequest) SetJobId(v string) *GetJobRequest {
	s.JobId = &v
	return s
}

func (s *GetJobRequest) SetProject(v string) *GetJobRequest {
	s.Project = &v
	return s
}

type GetJobResponseBody struct {
	Definition  *JobDefinition            `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description *string                   `json:"Description,omitempty" xml:"Description,omitempty"`
	HostId      *string                   `json:"HostId,omitempty" xml:"HostId,omitempty"`
	JobId       *string                   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Name        *string                   `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId     *string                   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Project     *string                   `json:"Project,omitempty" xml:"Project,omitempty"`
	RequestId   *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status      *GetJobResponseBodyStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
}

func (s GetJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobResponseBody) SetDefinition(v *JobDefinition) *GetJobResponseBody {
	s.Definition = v
	return s
}

func (s *GetJobResponseBody) SetDescription(v string) *GetJobResponseBody {
	s.Description = &v
	return s
}

func (s *GetJobResponseBody) SetHostId(v string) *GetJobResponseBody {
	s.HostId = &v
	return s
}

func (s *GetJobResponseBody) SetJobId(v string) *GetJobResponseBody {
	s.JobId = &v
	return s
}

func (s *GetJobResponseBody) SetName(v string) *GetJobResponseBody {
	s.Name = &v
	return s
}

func (s *GetJobResponseBody) SetOwnerId(v string) *GetJobResponseBody {
	s.OwnerId = &v
	return s
}

func (s *GetJobResponseBody) SetProject(v string) *GetJobResponseBody {
	s.Project = &v
	return s
}

func (s *GetJobResponseBody) SetRequestId(v string) *GetJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobResponseBody) SetStatus(v *GetJobResponseBodyStatus) *GetJobResponseBody {
	s.Status = v
	return s
}

type GetJobResponseBodyStatus struct {
	Attempts   []*Attempt `json:"Attempts,omitempty" xml:"Attempts,omitempty" type:"Repeated"`
	CreateTime *string    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EndTime    *string    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExitCode   *int32     `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	Pid        *int32     `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Reason     *string    `json:"Reason,omitempty" xml:"Reason,omitempty"`
	StartTime  *string    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State      *string    `json:"State,omitempty" xml:"State,omitempty"`
	Worker     *string    `json:"Worker,omitempty" xml:"Worker,omitempty"`
}

func (s GetJobResponseBodyStatus) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyStatus) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyStatus) SetAttempts(v []*Attempt) *GetJobResponseBodyStatus {
	s.Attempts = v
	return s
}

func (s *GetJobResponseBodyStatus) SetCreateTime(v string) *GetJobResponseBodyStatus {
	s.CreateTime = &v
	return s
}

func (s *GetJobResponseBodyStatus) SetEndTime(v string) *GetJobResponseBodyStatus {
	s.EndTime = &v
	return s
}

func (s *GetJobResponseBodyStatus) SetExitCode(v int32) *GetJobResponseBodyStatus {
	s.ExitCode = &v
	return s
}

func (s *GetJobResponseBodyStatus) SetPid(v int32) *GetJobResponseBodyStatus {
	s.Pid = &v
	return s
}

func (s *GetJobResponseBodyStatus) SetReason(v string) *GetJobResponseBodyStatus {
	s.Reason = &v
	return s
}

func (s *GetJobResponseBodyStatus) SetStartTime(v string) *GetJobResponseBodyStatus {
	s.StartTime = &v
	return s
}

func (s *GetJobResponseBodyStatus) SetState(v string) *GetJobResponseBodyStatus {
	s.State = &v
	return s
}

func (s *GetJobResponseBodyStatus) SetWorker(v string) *GetJobResponseBodyStatus {
	s.Worker = &v
	return s
}

type GetJobResponse struct {
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponse) GoString() string {
	return s.String()
}

func (s *GetJobResponse) SetHeaders(v map[string]*string) *GetJobResponse {
	s.Headers = v
	return s
}

func (s *GetJobResponse) SetBody(v *GetJobResponseBody) *GetJobResponse {
	s.Body = v
	return s
}

type GetJobQueueRequest struct {
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s GetJobQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobQueueRequest) GoString() string {
	return s.String()
}

func (s *GetJobQueueRequest) SetName(v string) *GetJobQueueRequest {
	s.Name = &v
	return s
}

func (s *GetJobQueueRequest) SetProject(v string) *GetJobQueueRequest {
	s.Project = &v
	return s
}

type GetJobQueueResponseBody struct {
	Definition  *JobQueueDefinition `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description *string             `json:"Description,omitempty" xml:"Description,omitempty"`
	HostId      *string             `json:"HostId,omitempty" xml:"HostId,omitempty"`
	Name        *string             `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId     *string             `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Project     *string             `json:"Project,omitempty" xml:"Project,omitempty"`
	RequestId   *string             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status      *JobQueueStatus     `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetJobQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobQueueResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobQueueResponseBody) SetDefinition(v *JobQueueDefinition) *GetJobQueueResponseBody {
	s.Definition = v
	return s
}

func (s *GetJobQueueResponseBody) SetDescription(v string) *GetJobQueueResponseBody {
	s.Description = &v
	return s
}

func (s *GetJobQueueResponseBody) SetHostId(v string) *GetJobQueueResponseBody {
	s.HostId = &v
	return s
}

func (s *GetJobQueueResponseBody) SetName(v string) *GetJobQueueResponseBody {
	s.Name = &v
	return s
}

func (s *GetJobQueueResponseBody) SetOwnerId(v string) *GetJobQueueResponseBody {
	s.OwnerId = &v
	return s
}

func (s *GetJobQueueResponseBody) SetProject(v string) *GetJobQueueResponseBody {
	s.Project = &v
	return s
}

func (s *GetJobQueueResponseBody) SetRequestId(v string) *GetJobQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobQueueResponseBody) SetStatus(v *JobQueueStatus) *GetJobQueueResponseBody {
	s.Status = v
	return s
}

type GetJobQueueResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetJobQueueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJobQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobQueueResponse) GoString() string {
	return s.String()
}

func (s *GetJobQueueResponse) SetHeaders(v map[string]*string) *GetJobQueueResponse {
	s.Headers = v
	return s
}

func (s *GetJobQueueResponse) SetBody(v *GetJobQueueResponseBody) *GetJobQueueResponse {
	s.Body = v
	return s
}

type GetProjectRequest struct {
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s GetProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProjectRequest) GoString() string {
	return s.String()
}

func (s *GetProjectRequest) SetProject(v string) *GetProjectRequest {
	s.Project = &v
	return s
}

type GetProjectResponseBody struct {
	Definition  *ProjectDefinition            `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description *string                       `json:"Description,omitempty" xml:"Description,omitempty"`
	HostId      *string                       `json:"HostId,omitempty" xml:"HostId,omitempty"`
	Project     *string                       `json:"Project,omitempty" xml:"Project,omitempty"`
	RequestId   *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status      *GetProjectResponseBodyStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
}

func (s GetProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBody) SetDefinition(v *ProjectDefinition) *GetProjectResponseBody {
	s.Definition = v
	return s
}

func (s *GetProjectResponseBody) SetDescription(v string) *GetProjectResponseBody {
	s.Description = &v
	return s
}

func (s *GetProjectResponseBody) SetHostId(v string) *GetProjectResponseBody {
	s.HostId = &v
	return s
}

func (s *GetProjectResponseBody) SetProject(v string) *GetProjectResponseBody {
	s.Project = &v
	return s
}

func (s *GetProjectResponseBody) SetRequestId(v string) *GetProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectResponseBody) SetStatus(v *GetProjectResponseBodyStatus) *GetProjectResponseBody {
	s.Status = v
	return s
}

type GetProjectResponseBodyStatus struct {
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
}

func (s GetProjectResponseBodyStatus) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyStatus) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyStatus) SetCreateTime(v string) *GetProjectResponseBodyStatus {
	s.CreateTime = &v
	return s
}

func (s *GetProjectResponseBodyStatus) SetLastModifiedTime(v string) *GetProjectResponseBodyStatus {
	s.LastModifiedTime = &v
	return s
}

type GetProjectResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponse) GoString() string {
	return s.String()
}

func (s *GetProjectResponse) SetHeaders(v map[string]*string) *GetProjectResponse {
	s.Headers = v
	return s
}

func (s *GetProjectResponse) SetBody(v *GetProjectResponseBody) *GetProjectResponse {
	s.Body = v
	return s
}

type GetWorkerRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Project   *string `json:"Project,omitempty" xml:"Project,omitempty"`
	WorkerId  *string `json:"WorkerId,omitempty" xml:"WorkerId,omitempty"`
}

func (s GetWorkerRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWorkerRequest) GoString() string {
	return s.String()
}

func (s *GetWorkerRequest) SetClusterId(v string) *GetWorkerRequest {
	s.ClusterId = &v
	return s
}

func (s *GetWorkerRequest) SetProject(v string) *GetWorkerRequest {
	s.Project = &v
	return s
}

func (s *GetWorkerRequest) SetWorkerId(v string) *GetWorkerRequest {
	s.WorkerId = &v
	return s
}

type GetWorkerResponseBody struct {
	HostId    *string                      `json:"HostId,omitempty" xml:"HostId,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *GetWorkerResponseBodyStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
	WorkerId  *string                      `json:"WorkerId,omitempty" xml:"WorkerId,omitempty"`
}

func (s GetWorkerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkerResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkerResponseBody) SetHostId(v string) *GetWorkerResponseBody {
	s.HostId = &v
	return s
}

func (s *GetWorkerResponseBody) SetRequestId(v string) *GetWorkerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkerResponseBody) SetStatus(v *GetWorkerResponseBodyStatus) *GetWorkerResponseBody {
	s.Status = v
	return s
}

func (s *GetWorkerResponseBody) SetWorkerId(v string) *GetWorkerResponseBody {
	s.WorkerId = &v
	return s
}

type GetWorkerResponseBodyStatus struct {
	ActiveTime         *string                               `json:"ActiveTime,omitempty" xml:"ActiveTime,omitempty"`
	AllocateTime       *string                               `json:"AllocateTime,omitempty" xml:"AllocateTime,omitempty"`
	Conditions         []*Conditions                         `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	Container          *GetWorkerResponseBodyStatusContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	CreateTime         *string                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ECS                *GetWorkerResponseBodyStatusECS       `json:"ECS,omitempty" xml:"ECS,omitempty" type:"Struct"`
	JobQueue           *string                               `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	NetworkInterfaceId *string                               `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	PoolWorkerId       *string                               `json:"PoolWorkerId,omitempty" xml:"PoolWorkerId,omitempty"`
	SecurityGroupId    *string                               `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	State              *string                               `json:"State,omitempty" xml:"State,omitempty"`
	VSwitchId          *string                               `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	WorkerType         *int32                                `json:"WorkerType,omitempty" xml:"WorkerType,omitempty"`
}

func (s GetWorkerResponseBodyStatus) String() string {
	return tea.Prettify(s)
}

func (s GetWorkerResponseBodyStatus) GoString() string {
	return s.String()
}

func (s *GetWorkerResponseBodyStatus) SetActiveTime(v string) *GetWorkerResponseBodyStatus {
	s.ActiveTime = &v
	return s
}

func (s *GetWorkerResponseBodyStatus) SetAllocateTime(v string) *GetWorkerResponseBodyStatus {
	s.AllocateTime = &v
	return s
}

func (s *GetWorkerResponseBodyStatus) SetConditions(v []*Conditions) *GetWorkerResponseBodyStatus {
	s.Conditions = v
	return s
}

func (s *GetWorkerResponseBodyStatus) SetContainer(v *GetWorkerResponseBodyStatusContainer) *GetWorkerResponseBodyStatus {
	s.Container = v
	return s
}

func (s *GetWorkerResponseBodyStatus) SetCreateTime(v string) *GetWorkerResponseBodyStatus {
	s.CreateTime = &v
	return s
}

func (s *GetWorkerResponseBodyStatus) SetECS(v *GetWorkerResponseBodyStatusECS) *GetWorkerResponseBodyStatus {
	s.ECS = v
	return s
}

func (s *GetWorkerResponseBodyStatus) SetJobQueue(v string) *GetWorkerResponseBodyStatus {
	s.JobQueue = &v
	return s
}

func (s *GetWorkerResponseBodyStatus) SetNetworkInterfaceId(v string) *GetWorkerResponseBodyStatus {
	s.NetworkInterfaceId = &v
	return s
}

func (s *GetWorkerResponseBodyStatus) SetPoolWorkerId(v string) *GetWorkerResponseBodyStatus {
	s.PoolWorkerId = &v
	return s
}

func (s *GetWorkerResponseBodyStatus) SetSecurityGroupId(v string) *GetWorkerResponseBodyStatus {
	s.SecurityGroupId = &v
	return s
}

func (s *GetWorkerResponseBodyStatus) SetState(v string) *GetWorkerResponseBodyStatus {
	s.State = &v
	return s
}

func (s *GetWorkerResponseBodyStatus) SetVSwitchId(v string) *GetWorkerResponseBodyStatus {
	s.VSwitchId = &v
	return s
}

func (s *GetWorkerResponseBodyStatus) SetWorkerType(v int32) *GetWorkerResponseBodyStatus {
	s.WorkerType = &v
	return s
}

type GetWorkerResponseBodyStatusContainer struct {
	Cpu    *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s GetWorkerResponseBodyStatusContainer) String() string {
	return tea.Prettify(s)
}

func (s GetWorkerResponseBodyStatusContainer) GoString() string {
	return s.String()
}

func (s *GetWorkerResponseBodyStatusContainer) SetCpu(v int32) *GetWorkerResponseBodyStatusContainer {
	s.Cpu = &v
	return s
}

func (s *GetWorkerResponseBodyStatusContainer) SetMemory(v int32) *GetWorkerResponseBodyStatusContainer {
	s.Memory = &v
	return s
}

type GetWorkerResponseBodyStatusECS struct {
	Cpu            *int32  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Endpoint       *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Hostname       *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType   *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Memory         *int32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Password       *string `json:"Password,omitempty" xml:"Password,omitempty"`
	ResourceType   *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SpotPriceLimit *string `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	SpotStrategy   *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	SystemDiskSize *int32  `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	SystemDiskType *string `json:"SystemDiskType,omitempty" xml:"SystemDiskType,omitempty"`
	VMImage        *string `json:"VMImage,omitempty" xml:"VMImage,omitempty"`
	ZoneId         *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetWorkerResponseBodyStatusECS) String() string {
	return tea.Prettify(s)
}

func (s GetWorkerResponseBodyStatusECS) GoString() string {
	return s.String()
}

func (s *GetWorkerResponseBodyStatusECS) SetCpu(v int32) *GetWorkerResponseBodyStatusECS {
	s.Cpu = &v
	return s
}

func (s *GetWorkerResponseBodyStatusECS) SetEndpoint(v string) *GetWorkerResponseBodyStatusECS {
	s.Endpoint = &v
	return s
}

func (s *GetWorkerResponseBodyStatusECS) SetHostname(v string) *GetWorkerResponseBodyStatusECS {
	s.Hostname = &v
	return s
}

func (s *GetWorkerResponseBodyStatusECS) SetInstanceId(v string) *GetWorkerResponseBodyStatusECS {
	s.InstanceId = &v
	return s
}

func (s *GetWorkerResponseBodyStatusECS) SetInstanceType(v string) *GetWorkerResponseBodyStatusECS {
	s.InstanceType = &v
	return s
}

func (s *GetWorkerResponseBodyStatusECS) SetMemory(v int32) *GetWorkerResponseBodyStatusECS {
	s.Memory = &v
	return s
}

func (s *GetWorkerResponseBodyStatusECS) SetPassword(v string) *GetWorkerResponseBodyStatusECS {
	s.Password = &v
	return s
}

func (s *GetWorkerResponseBodyStatusECS) SetResourceType(v string) *GetWorkerResponseBodyStatusECS {
	s.ResourceType = &v
	return s
}

func (s *GetWorkerResponseBodyStatusECS) SetSpotPriceLimit(v string) *GetWorkerResponseBodyStatusECS {
	s.SpotPriceLimit = &v
	return s
}

func (s *GetWorkerResponseBodyStatusECS) SetSpotStrategy(v string) *GetWorkerResponseBodyStatusECS {
	s.SpotStrategy = &v
	return s
}

func (s *GetWorkerResponseBodyStatusECS) SetSystemDiskSize(v int32) *GetWorkerResponseBodyStatusECS {
	s.SystemDiskSize = &v
	return s
}

func (s *GetWorkerResponseBodyStatusECS) SetSystemDiskType(v string) *GetWorkerResponseBodyStatusECS {
	s.SystemDiskType = &v
	return s
}

func (s *GetWorkerResponseBodyStatusECS) SetVMImage(v string) *GetWorkerResponseBodyStatusECS {
	s.VMImage = &v
	return s
}

func (s *GetWorkerResponseBodyStatusECS) SetZoneId(v string) *GetWorkerResponseBodyStatusECS {
	s.ZoneId = &v
	return s
}

type GetWorkerResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWorkerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWorkerResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWorkerResponse) GoString() string {
	return s.String()
}

func (s *GetWorkerResponse) SetHeaders(v map[string]*string) *GetWorkerResponse {
	s.Headers = v
	return s
}

func (s *GetWorkerResponse) SetBody(v *GetWorkerResponseBody) *GetWorkerResponse {
	s.Body = v
	return s
}

type KillWorkerRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Project   *string `json:"Project,omitempty" xml:"Project,omitempty"`
	WorkerId  *string `json:"WorkerId,omitempty" xml:"WorkerId,omitempty"`
}

func (s KillWorkerRequest) String() string {
	return tea.Prettify(s)
}

func (s KillWorkerRequest) GoString() string {
	return s.String()
}

func (s *KillWorkerRequest) SetClusterId(v string) *KillWorkerRequest {
	s.ClusterId = &v
	return s
}

func (s *KillWorkerRequest) SetProject(v string) *KillWorkerRequest {
	s.Project = &v
	return s
}

func (s *KillWorkerRequest) SetWorkerId(v string) *KillWorkerRequest {
	s.WorkerId = &v
	return s
}

type KillWorkerResponseBody struct {
	HostId    *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s KillWorkerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s KillWorkerResponseBody) GoString() string {
	return s.String()
}

func (s *KillWorkerResponseBody) SetHostId(v string) *KillWorkerResponseBody {
	s.HostId = &v
	return s
}

func (s *KillWorkerResponseBody) SetRequestId(v string) *KillWorkerResponseBody {
	s.RequestId = &v
	return s
}

type KillWorkerResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *KillWorkerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s KillWorkerResponse) String() string {
	return tea.Prettify(s)
}

func (s KillWorkerResponse) GoString() string {
	return s.String()
}

func (s *KillWorkerResponse) SetHeaders(v map[string]*string) *KillWorkerResponse {
	s.Headers = v
	return s
}

func (s *KillWorkerResponse) SetBody(v *KillWorkerResponseBody) *KillWorkerResponse {
	s.Body = v
	return s
}

type ListClustersRequest struct {
	Filter     *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Project    *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s ListClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClustersRequest) GoString() string {
	return s.String()
}

func (s *ListClustersRequest) SetFilter(v string) *ListClustersRequest {
	s.Filter = &v
	return s
}

func (s *ListClustersRequest) SetMaxResults(v int32) *ListClustersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListClustersRequest) SetNextToken(v string) *ListClustersRequest {
	s.NextToken = &v
	return s
}

func (s *ListClustersRequest) SetProject(v string) *ListClustersRequest {
	s.Project = &v
	return s
}

type ListClustersResponseBody struct {
	Clusters   []*ListClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	HostId     *string                             `json:"HostId,omitempty" xml:"HostId,omitempty"`
	NextToken  *string                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBody) SetClusters(v []*ListClustersResponseBodyClusters) *ListClustersResponseBody {
	s.Clusters = v
	return s
}

func (s *ListClustersResponseBody) SetHostId(v string) *ListClustersResponseBody {
	s.HostId = &v
	return s
}

func (s *ListClustersResponseBody) SetNextToken(v string) *ListClustersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListClustersResponseBody) SetRequestId(v string) *ListClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClustersResponseBody) SetTotalCount(v int32) *ListClustersResponseBody {
	s.TotalCount = &v
	return s
}

type ListClustersResponseBodyClusters struct {
	ClusterId   *string                                 `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Definition  *ClusterDefinition                      `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description *string                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId     *string                                 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Project     *string                                 `json:"Project,omitempty" xml:"Project,omitempty"`
	Status      *ListClustersResponseBodyClustersStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
}

func (s ListClustersResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClusters) SetClusterId(v string) *ListClustersResponseBodyClusters {
	s.ClusterId = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetDefinition(v *ClusterDefinition) *ListClustersResponseBodyClusters {
	s.Definition = v
	return s
}

func (s *ListClustersResponseBodyClusters) SetDescription(v string) *ListClustersResponseBodyClusters {
	s.Description = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetName(v string) *ListClustersResponseBodyClusters {
	s.Name = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetOwnerId(v string) *ListClustersResponseBodyClusters {
	s.OwnerId = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetProject(v string) *ListClustersResponseBodyClusters {
	s.Project = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetStatus(v *ListClustersResponseBodyClustersStatus) *ListClustersResponseBodyClusters {
	s.Status = v
	return s
}

type ListClustersResponseBodyClustersStatus struct {
	AllocatableResources map[string]*string `json:"AllocatableResources,omitempty" xml:"AllocatableResources,omitempty"`
	AllocatedResources   map[string]*string `json:"AllocatedResources,omitempty" xml:"AllocatedResources,omitempty"`
	Conditions           []*Conditions      `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	CreateTime           *string            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CurrentWorkerCount   *int32             `json:"CurrentWorkerCount,omitempty" xml:"CurrentWorkerCount,omitempty"`
	DesiredWorkerCount   *int32             `json:"DesiredWorkerCount,omitempty" xml:"DesiredWorkerCount,omitempty"`
	State                *string            `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListClustersResponseBodyClustersStatus) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersStatus) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersStatus) SetAllocatableResources(v map[string]*string) *ListClustersResponseBodyClustersStatus {
	s.AllocatableResources = v
	return s
}

func (s *ListClustersResponseBodyClustersStatus) SetAllocatedResources(v map[string]*string) *ListClustersResponseBodyClustersStatus {
	s.AllocatedResources = v
	return s
}

func (s *ListClustersResponseBodyClustersStatus) SetConditions(v []*Conditions) *ListClustersResponseBodyClustersStatus {
	s.Conditions = v
	return s
}

func (s *ListClustersResponseBodyClustersStatus) SetCreateTime(v string) *ListClustersResponseBodyClustersStatus {
	s.CreateTime = &v
	return s
}

func (s *ListClustersResponseBodyClustersStatus) SetCurrentWorkerCount(v int32) *ListClustersResponseBodyClustersStatus {
	s.CurrentWorkerCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersStatus) SetDesiredWorkerCount(v int32) *ListClustersResponseBodyClustersStatus {
	s.DesiredWorkerCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersStatus) SetState(v string) *ListClustersResponseBodyClustersStatus {
	s.State = &v
	return s
}

type ListClustersResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponse) GoString() string {
	return s.String()
}

func (s *ListClustersResponse) SetHeaders(v map[string]*string) *ListClustersResponse {
	s.Headers = v
	return s
}

func (s *ListClustersResponse) SetBody(v *ListClustersResponseBody) *ListClustersResponse {
	s.Body = v
	return s
}

type ListJobQueuesRequest struct {
	LabelSelector *string `json:"LabelSelector,omitempty" xml:"LabelSelector,omitempty"`
	MaxResults    *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OrderBy       *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	OrderReverse  *bool   `json:"OrderReverse,omitempty" xml:"OrderReverse,omitempty"`
	Project       *string `json:"Project,omitempty" xml:"Project,omitempty"`
	State         *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListJobQueuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobQueuesRequest) GoString() string {
	return s.String()
}

func (s *ListJobQueuesRequest) SetLabelSelector(v string) *ListJobQueuesRequest {
	s.LabelSelector = &v
	return s
}

func (s *ListJobQueuesRequest) SetMaxResults(v int32) *ListJobQueuesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListJobQueuesRequest) SetNextToken(v string) *ListJobQueuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListJobQueuesRequest) SetOrderBy(v string) *ListJobQueuesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListJobQueuesRequest) SetOrderReverse(v bool) *ListJobQueuesRequest {
	s.OrderReverse = &v
	return s
}

func (s *ListJobQueuesRequest) SetProject(v string) *ListJobQueuesRequest {
	s.Project = &v
	return s
}

func (s *ListJobQueuesRequest) SetState(v string) *ListJobQueuesRequest {
	s.State = &v
	return s
}

type ListJobQueuesResponseBody struct {
	HostId    *string                               `json:"HostId,omitempty" xml:"HostId,omitempty"`
	JobQueues []*ListJobQueuesResponseBodyJobQueues `json:"JobQueues,omitempty" xml:"JobQueues,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListJobQueuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobQueuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobQueuesResponseBody) SetHostId(v string) *ListJobQueuesResponseBody {
	s.HostId = &v
	return s
}

func (s *ListJobQueuesResponseBody) SetJobQueues(v []*ListJobQueuesResponseBodyJobQueues) *ListJobQueuesResponseBody {
	s.JobQueues = v
	return s
}

func (s *ListJobQueuesResponseBody) SetRequestId(v string) *ListJobQueuesResponseBody {
	s.RequestId = &v
	return s
}

type ListJobQueuesResponseBodyJobQueues struct {
	Definition  *JobQueueDefinition `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description *string             `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string             `json:"Name,omitempty" xml:"Name,omitempty"`
	NextToken   *string             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerId     *string             `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Project     *string             `json:"Project,omitempty" xml:"Project,omitempty"`
	Status      *JobQueueStatus     `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalCount  *int32              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListJobQueuesResponseBodyJobQueues) String() string {
	return tea.Prettify(s)
}

func (s ListJobQueuesResponseBodyJobQueues) GoString() string {
	return s.String()
}

func (s *ListJobQueuesResponseBodyJobQueues) SetDefinition(v *JobQueueDefinition) *ListJobQueuesResponseBodyJobQueues {
	s.Definition = v
	return s
}

func (s *ListJobQueuesResponseBodyJobQueues) SetDescription(v string) *ListJobQueuesResponseBodyJobQueues {
	s.Description = &v
	return s
}

func (s *ListJobQueuesResponseBodyJobQueues) SetName(v string) *ListJobQueuesResponseBodyJobQueues {
	s.Name = &v
	return s
}

func (s *ListJobQueuesResponseBodyJobQueues) SetNextToken(v string) *ListJobQueuesResponseBodyJobQueues {
	s.NextToken = &v
	return s
}

func (s *ListJobQueuesResponseBodyJobQueues) SetOwnerId(v string) *ListJobQueuesResponseBodyJobQueues {
	s.OwnerId = &v
	return s
}

func (s *ListJobQueuesResponseBodyJobQueues) SetProject(v string) *ListJobQueuesResponseBodyJobQueues {
	s.Project = &v
	return s
}

func (s *ListJobQueuesResponseBodyJobQueues) SetStatus(v *JobQueueStatus) *ListJobQueuesResponseBodyJobQueues {
	s.Status = v
	return s
}

func (s *ListJobQueuesResponseBodyJobQueues) SetTotalCount(v int32) *ListJobQueuesResponseBodyJobQueues {
	s.TotalCount = &v
	return s
}

type ListJobQueuesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListJobQueuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListJobQueuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobQueuesResponse) GoString() string {
	return s.String()
}

func (s *ListJobQueuesResponse) SetHeaders(v map[string]*string) *ListJobQueuesResponse {
	s.Headers = v
	return s
}

func (s *ListJobQueuesResponse) SetBody(v *ListJobQueuesResponseBody) *ListJobQueuesResponse {
	s.Body = v
	return s
}

type ListJobsRequest struct {
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	LabelSelector *string `json:"LabelSelector,omitempty" xml:"LabelSelector,omitempty"`
	MaxResults    *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NextToken     *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OrderBy       *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	OrderReverse  *bool   `json:"OrderReverse,omitempty" xml:"OrderReverse,omitempty"`
	Project       *string `json:"Project,omitempty" xml:"Project,omitempty"`
	State         *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) SetClusterId(v string) *ListJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListJobsRequest) SetLabelSelector(v string) *ListJobsRequest {
	s.LabelSelector = &v
	return s
}

func (s *ListJobsRequest) SetMaxResults(v int32) *ListJobsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListJobsRequest) SetName(v string) *ListJobsRequest {
	s.Name = &v
	return s
}

func (s *ListJobsRequest) SetNextToken(v string) *ListJobsRequest {
	s.NextToken = &v
	return s
}

func (s *ListJobsRequest) SetOrderBy(v string) *ListJobsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListJobsRequest) SetOrderReverse(v bool) *ListJobsRequest {
	s.OrderReverse = &v
	return s
}

func (s *ListJobsRequest) SetProject(v string) *ListJobsRequest {
	s.Project = &v
	return s
}

func (s *ListJobsRequest) SetState(v string) *ListJobsRequest {
	s.State = &v
	return s
}

type ListJobsResponseBody struct {
	HostId     *string                     `json:"HostId,omitempty" xml:"HostId,omitempty"`
	Jobs       []*ListJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	NextToken  *string                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBody) SetHostId(v string) *ListJobsResponseBody {
	s.HostId = &v
	return s
}

func (s *ListJobsResponseBody) SetJobs(v []*ListJobsResponseBodyJobs) *ListJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListJobsResponseBody) SetNextToken(v string) *ListJobsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListJobsResponseBody) SetRequestId(v string) *ListJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobsResponseBody) SetTotalCount(v int32) *ListJobsResponseBody {
	s.TotalCount = &v
	return s
}

type ListJobsResponseBodyJobs struct {
	Definition  *JobDefinition                  `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description *string                         `json:"Description,omitempty" xml:"Description,omitempty"`
	JobId       *string                         `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Name        *string                         `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId     *string                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Project     *string                         `json:"Project,omitempty" xml:"Project,omitempty"`
	Status      *ListJobsResponseBodyJobsStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
}

func (s ListJobsResponseBodyJobs) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobs) SetDefinition(v *JobDefinition) *ListJobsResponseBodyJobs {
	s.Definition = v
	return s
}

func (s *ListJobsResponseBodyJobs) SetDescription(v string) *ListJobsResponseBodyJobs {
	s.Description = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetJobId(v string) *ListJobsResponseBodyJobs {
	s.JobId = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetName(v string) *ListJobsResponseBodyJobs {
	s.Name = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetOwnerId(v string) *ListJobsResponseBodyJobs {
	s.OwnerId = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetProject(v string) *ListJobsResponseBodyJobs {
	s.Project = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetStatus(v *ListJobsResponseBodyJobsStatus) *ListJobsResponseBodyJobs {
	s.Status = v
	return s
}

type ListJobsResponseBodyJobsStatus struct {
	Attempts   []*Attempt `json:"Attempts,omitempty" xml:"Attempts,omitempty" type:"Repeated"`
	CreateTime *string    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EndTime    *string    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExitCode   *int32     `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	Pid        *int32     `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Reason     *string    `json:"Reason,omitempty" xml:"Reason,omitempty"`
	StartTime  *string    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State      *string    `json:"State,omitempty" xml:"State,omitempty"`
	Worker     *string    `json:"Worker,omitempty" xml:"Worker,omitempty"`
}

func (s ListJobsResponseBodyJobsStatus) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyJobsStatus) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobsStatus) SetAttempts(v []*Attempt) *ListJobsResponseBodyJobsStatus {
	s.Attempts = v
	return s
}

func (s *ListJobsResponseBodyJobsStatus) SetCreateTime(v string) *ListJobsResponseBodyJobsStatus {
	s.CreateTime = &v
	return s
}

func (s *ListJobsResponseBodyJobsStatus) SetEndTime(v string) *ListJobsResponseBodyJobsStatus {
	s.EndTime = &v
	return s
}

func (s *ListJobsResponseBodyJobsStatus) SetExitCode(v int32) *ListJobsResponseBodyJobsStatus {
	s.ExitCode = &v
	return s
}

func (s *ListJobsResponseBodyJobsStatus) SetPid(v int32) *ListJobsResponseBodyJobsStatus {
	s.Pid = &v
	return s
}

func (s *ListJobsResponseBodyJobsStatus) SetReason(v string) *ListJobsResponseBodyJobsStatus {
	s.Reason = &v
	return s
}

func (s *ListJobsResponseBodyJobsStatus) SetStartTime(v string) *ListJobsResponseBodyJobsStatus {
	s.StartTime = &v
	return s
}

func (s *ListJobsResponseBodyJobsStatus) SetState(v string) *ListJobsResponseBodyJobsStatus {
	s.State = &v
	return s
}

func (s *ListJobsResponseBodyJobsStatus) SetWorker(v string) *ListJobsResponseBodyJobsStatus {
	s.Worker = &v
	return s
}

type ListJobsResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponse) GoString() string {
	return s.String()
}

func (s *ListJobsResponse) SetHeaders(v map[string]*string) *ListJobsResponse {
	s.Headers = v
	return s
}

func (s *ListJobsResponse) SetBody(v *ListJobsResponseBody) *ListJobsResponse {
	s.Body = v
	return s
}

type ListProjectsRequest struct {
	LabelSelector *string `json:"LabelSelector,omitempty" xml:"LabelSelector,omitempty"`
	MaxResults    *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OrderBy       *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	OrderReverse  *bool   `json:"OrderReverse,omitempty" xml:"OrderReverse,omitempty"`
}

func (s ListProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) SetLabelSelector(v string) *ListProjectsRequest {
	s.LabelSelector = &v
	return s
}

func (s *ListProjectsRequest) SetMaxResults(v int32) *ListProjectsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProjectsRequest) SetNextToken(v string) *ListProjectsRequest {
	s.NextToken = &v
	return s
}

func (s *ListProjectsRequest) SetOrderBy(v string) *ListProjectsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListProjectsRequest) SetOrderReverse(v bool) *ListProjectsRequest {
	s.OrderReverse = &v
	return s
}

type ListProjectsResponseBody struct {
	HostId     *string                             `json:"HostId,omitempty" xml:"HostId,omitempty"`
	MaxResults *int32                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Projects   []*ListProjectsResponseBodyProjects `json:"Projects,omitempty" xml:"Projects,omitempty" type:"Repeated"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) SetHostId(v string) *ListProjectsResponseBody {
	s.HostId = &v
	return s
}

func (s *ListProjectsResponseBody) SetMaxResults(v int32) *ListProjectsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListProjectsResponseBody) SetNextToken(v string) *ListProjectsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProjectsResponseBody) SetProjects(v []*ListProjectsResponseBodyProjects) *ListProjectsResponseBody {
	s.Projects = v
	return s
}

func (s *ListProjectsResponseBody) SetRequestId(v string) *ListProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectsResponseBody) SetTotalCount(v int32) *ListProjectsResponseBody {
	s.TotalCount = &v
	return s
}

type ListProjectsResponseBodyProjects struct {
	Definition  *ProjectDefinition                      `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description *string                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	Project     *string                                 `json:"Project,omitempty" xml:"Project,omitempty"`
	Status      *ListProjectsResponseBodyProjectsStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
}

func (s ListProjectsResponseBodyProjects) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyProjects) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyProjects) SetDefinition(v *ProjectDefinition) *ListProjectsResponseBodyProjects {
	s.Definition = v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetDescription(v string) *ListProjectsResponseBodyProjects {
	s.Description = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetProject(v string) *ListProjectsResponseBodyProjects {
	s.Project = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetStatus(v *ListProjectsResponseBodyProjectsStatus) *ListProjectsResponseBodyProjects {
	s.Status = v
	return s
}

type ListProjectsResponseBodyProjectsStatus struct {
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
}

func (s ListProjectsResponseBodyProjectsStatus) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyProjectsStatus) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyProjectsStatus) SetCreateTime(v string) *ListProjectsResponseBodyProjectsStatus {
	s.CreateTime = &v
	return s
}

func (s *ListProjectsResponseBodyProjectsStatus) SetLastModifiedTime(v string) *ListProjectsResponseBodyProjectsStatus {
	s.LastModifiedTime = &v
	return s
}

type ListProjectsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponse) GoString() string {
	return s.String()
}

func (s *ListProjectsResponse) SetHeaders(v map[string]*string) *ListProjectsResponse {
	s.Headers = v
	return s
}

func (s *ListProjectsResponse) SetBody(v *ListProjectsResponseBody) *ListProjectsResponse {
	s.Body = v
	return s
}

type ListRegionsRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s ListRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsRequest) GoString() string {
	return s.String()
}

func (s *ListRegionsRequest) SetAcceptLanguage(v string) *ListRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

type ListRegionsResponseBody struct {
	HostId    *string                           `json:"HostId,omitempty" xml:"HostId,omitempty"`
	Regions   []*ListRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) SetHostId(v string) *ListRegionsResponseBody {
	s.HostId = &v
	return s
}

func (s *ListRegionsResponseBody) SetRegions(v []*ListRegionsResponseBodyRegions) *ListRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

type ListRegionsResponseBodyRegions struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyRegions) SetLocalName(v string) *ListRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *ListRegionsResponseBodyRegions) SetRegionId(v string) *ListRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type ListRegionsResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListRegionsResponse) SetHeaders(v map[string]*string) *ListRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListRegionsResponse) SetBody(v *ListRegionsResponseBody) *ListRegionsResponse {
	s.Body = v
	return s
}

type ListWorkersRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Project    *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s ListWorkersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkersRequest) GoString() string {
	return s.String()
}

func (s *ListWorkersRequest) SetClusterId(v string) *ListWorkersRequest {
	s.ClusterId = &v
	return s
}

func (s *ListWorkersRequest) SetMaxResults(v int32) *ListWorkersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkersRequest) SetNextToken(v string) *ListWorkersRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkersRequest) SetProject(v string) *ListWorkersRequest {
	s.Project = &v
	return s
}

type ListWorkersResponseBody struct {
	HostId     *string                           `json:"HostId,omitempty" xml:"HostId,omitempty"`
	NextToken  *string                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Workers    []*ListWorkersResponseBodyWorkers `json:"Workers,omitempty" xml:"Workers,omitempty" type:"Repeated"`
}

func (s ListWorkersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWorkersResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkersResponseBody) SetHostId(v string) *ListWorkersResponseBody {
	s.HostId = &v
	return s
}

func (s *ListWorkersResponseBody) SetNextToken(v string) *ListWorkersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkersResponseBody) SetRequestId(v string) *ListWorkersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkersResponseBody) SetTotalCount(v int32) *ListWorkersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWorkersResponseBody) SetWorkers(v []*ListWorkersResponseBodyWorkers) *ListWorkersResponseBody {
	s.Workers = v
	return s
}

type ListWorkersResponseBodyWorkers struct {
	Status   *ListWorkersResponseBodyWorkersStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
	WorkerId *string                               `json:"WorkerId,omitempty" xml:"WorkerId,omitempty"`
}

func (s ListWorkersResponseBodyWorkers) String() string {
	return tea.Prettify(s)
}

func (s ListWorkersResponseBodyWorkers) GoString() string {
	return s.String()
}

func (s *ListWorkersResponseBodyWorkers) SetStatus(v *ListWorkersResponseBodyWorkersStatus) *ListWorkersResponseBodyWorkers {
	s.Status = v
	return s
}

func (s *ListWorkersResponseBodyWorkers) SetWorkerId(v string) *ListWorkersResponseBodyWorkers {
	s.WorkerId = &v
	return s
}

type ListWorkersResponseBodyWorkersStatus struct {
	ActiveTime         *string                                        `json:"ActiveTime,omitempty" xml:"ActiveTime,omitempty"`
	AllocateTime       *string                                        `json:"AllocateTime,omitempty" xml:"AllocateTime,omitempty"`
	Conditions         []*Conditions                                  `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	Container          *ListWorkersResponseBodyWorkersStatusContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	CreateTime         *string                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ECS                *ListWorkersResponseBodyWorkersStatusECS       `json:"ECS,omitempty" xml:"ECS,omitempty" type:"Struct"`
	JobQueue           *string                                        `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	NetworkInterfaceId *string                                        `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	PoolWorkerId       *string                                        `json:"PoolWorkerId,omitempty" xml:"PoolWorkerId,omitempty"`
	SecurityGroupId    *string                                        `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	State              *string                                        `json:"State,omitempty" xml:"State,omitempty"`
	VSwitchId          *string                                        `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	WorkerType         *int32                                         `json:"WorkerType,omitempty" xml:"WorkerType,omitempty"`
}

func (s ListWorkersResponseBodyWorkersStatus) String() string {
	return tea.Prettify(s)
}

func (s ListWorkersResponseBodyWorkersStatus) GoString() string {
	return s.String()
}

func (s *ListWorkersResponseBodyWorkersStatus) SetActiveTime(v string) *ListWorkersResponseBodyWorkersStatus {
	s.ActiveTime = &v
	return s
}

func (s *ListWorkersResponseBodyWorkersStatus) SetAllocateTime(v string) *ListWorkersResponseBodyWorkersStatus {
	s.AllocateTime = &v
	return s
}

func (s *ListWorkersResponseBodyWorkersStatus) SetConditions(v []*Conditions) *ListWorkersResponseBodyWorkersStatus {
	s.Conditions = v
	return s
}

func (s *ListWorkersResponseBodyWorkersStatus) SetContainer(v *ListWorkersResponseBodyWorkersStatusContainer) *ListWorkersResponseBodyWorkersStatus {
	s.Container = v
	return s
}

func (s *ListWorkersResponseBodyWorkersStatus) SetCreateTime(v string) *ListWorkersResponseBodyWorkersStatus {
	s.CreateTime = &v
	return s
}

func (s *ListWorkersResponseBodyWorkersStatus) SetECS(v *ListWorkersResponseBodyWorkersStatusECS) *ListWorkersResponseBodyWorkersStatus {
	s.ECS = v
	return s
}

func (s *ListWorkersResponseBodyWorkersStatus) SetJobQueue(v string) *ListWorkersResponseBodyWorkersStatus {
	s.JobQueue = &v
	return s
}

func (s *ListWorkersResponseBodyWorkersStatus) SetNetworkInterfaceId(v string) *ListWorkersResponseBodyWorkersStatus {
	s.NetworkInterfaceId = &v
	return s
}

func (s *ListWorkersResponseBodyWorkersStatus) SetPoolWorkerId(v string) *ListWorkersResponseBodyWorkersStatus {
	s.PoolWorkerId = &v
	return s
}

func (s *ListWorkersResponseBodyWorkersStatus) SetSecurityGroupId(v string) *ListWorkersResponseBodyWorkersStatus {
	s.SecurityGroupId = &v
	return s
}

func (s *ListWorkersResponseBodyWorkersStatus) SetState(v string) *ListWorkersResponseBodyWorkersStatus {
	s.State = &v
	return s
}

func (s *ListWorkersResponseBodyWorkersStatus) SetVSwitchId(v string) *ListWorkersResponseBodyWorkersStatus {
	s.VSwitchId = &v
	return s
}

func (s *ListWorkersResponseBodyWorkersStatus) SetWorkerType(v int32) *ListWorkersResponseBodyWorkersStatus {
	s.WorkerType = &v
	return s
}

type ListWorkersResponseBodyWorkersStatusContainer struct {
	Cpu    *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListWorkersResponseBodyWorkersStatusContainer) String() string {
	return tea.Prettify(s)
}

func (s ListWorkersResponseBodyWorkersStatusContainer) GoString() string {
	return s.String()
}

func (s *ListWorkersResponseBodyWorkersStatusContainer) SetCpu(v int32) *ListWorkersResponseBodyWorkersStatusContainer {
	s.Cpu = &v
	return s
}

func (s *ListWorkersResponseBodyWorkersStatusContainer) SetMemory(v int32) *ListWorkersResponseBodyWorkersStatusContainer {
	s.Memory = &v
	return s
}

type ListWorkersResponseBodyWorkersStatusECS struct {
	Cpu            *int32  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Endpoint       *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Hostname       *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType   *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Memory         *int32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Password       *string `json:"Password,omitempty" xml:"Password,omitempty"`
	ResourceType   *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SpotPriceLimit *string `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	SpotStrategy   *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	SystemDiskSize *int32  `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	SystemDiskType *string `json:"SystemDiskType,omitempty" xml:"SystemDiskType,omitempty"`
	VMImage        *string `json:"VMImage,omitempty" xml:"VMImage,omitempty"`
	ZoneId         *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListWorkersResponseBodyWorkersStatusECS) String() string {
	return tea.Prettify(s)
}

func (s ListWorkersResponseBodyWorkersStatusECS) GoString() string {
	return s.String()
}

func (s *ListWorkersResponseBodyWorkersStatusECS) SetCpu(v int32) *ListWorkersResponseBodyWorkersStatusECS {
	s.Cpu = &v
	return s
}

func (s *ListWorkersResponseBodyWorkersStatusECS) SetEndpoint(v string) *ListWorkersResponseBodyWorkersStatusECS {
	s.Endpoint = &v
	return s
}

func (s *ListWorkersResponseBodyWorkersStatusECS) SetHostname(v string) *ListWorkersResponseBodyWorkersStatusECS {
	s.Hostname = &v
	return s
}

func (s *ListWorkersResponseBodyWorkersStatusECS) SetInstanceId(v string) *ListWorkersResponseBodyWorkersStatusECS {
	s.InstanceId = &v
	return s
}

func (s *ListWorkersResponseBodyWorkersStatusECS) SetInstanceType(v string) *ListWorkersResponseBodyWorkersStatusECS {
	s.InstanceType = &v
	return s
}

func (s *ListWorkersResponseBodyWorkersStatusECS) SetMemory(v int32) *ListWorkersResponseBodyWorkersStatusECS {
	s.Memory = &v
	return s
}

func (s *ListWorkersResponseBodyWorkersStatusECS) SetPassword(v string) *ListWorkersResponseBodyWorkersStatusECS {
	s.Password = &v
	return s
}

func (s *ListWorkersResponseBodyWorkersStatusECS) SetResourceType(v string) *ListWorkersResponseBodyWorkersStatusECS {
	s.ResourceType = &v
	return s
}

func (s *ListWorkersResponseBodyWorkersStatusECS) SetSpotPriceLimit(v string) *ListWorkersResponseBodyWorkersStatusECS {
	s.SpotPriceLimit = &v
	return s
}

func (s *ListWorkersResponseBodyWorkersStatusECS) SetSpotStrategy(v string) *ListWorkersResponseBodyWorkersStatusECS {
	s.SpotStrategy = &v
	return s
}

func (s *ListWorkersResponseBodyWorkersStatusECS) SetSystemDiskSize(v int32) *ListWorkersResponseBodyWorkersStatusECS {
	s.SystemDiskSize = &v
	return s
}

func (s *ListWorkersResponseBodyWorkersStatusECS) SetSystemDiskType(v string) *ListWorkersResponseBodyWorkersStatusECS {
	s.SystemDiskType = &v
	return s
}

func (s *ListWorkersResponseBodyWorkersStatusECS) SetVMImage(v string) *ListWorkersResponseBodyWorkersStatusECS {
	s.VMImage = &v
	return s
}

func (s *ListWorkersResponseBodyWorkersStatusECS) SetZoneId(v string) *ListWorkersResponseBodyWorkersStatusECS {
	s.ZoneId = &v
	return s
}

type ListWorkersResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListWorkersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListWorkersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWorkersResponse) GoString() string {
	return s.String()
}

func (s *ListWorkersResponse) SetHeaders(v map[string]*string) *ListWorkersResponse {
	s.Headers = v
	return s
}

func (s *ListWorkersResponse) SetBody(v *ListWorkersResponseBody) *ListWorkersResponse {
	s.Body = v
	return s
}

type OpenBatchComputeServiceResponseBody struct {
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenBatchComputeServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenBatchComputeServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenBatchComputeServiceResponseBody) SetOrderId(v string) *OpenBatchComputeServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenBatchComputeServiceResponseBody) SetRequestId(v string) *OpenBatchComputeServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenBatchComputeServiceResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenBatchComputeServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenBatchComputeServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenBatchComputeServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenBatchComputeServiceResponse) SetHeaders(v map[string]*string) *OpenBatchComputeServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenBatchComputeServiceResponse) SetBody(v *OpenBatchComputeServiceResponseBody) *OpenBatchComputeServiceResponse {
	s.Body = v
	return s
}

type PollCmdRequest struct {
	Queue       *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
	WaitSeconds *string `json:"WaitSeconds,omitempty" xml:"WaitSeconds,omitempty"`
	WorkerId    *string `json:"WorkerId,omitempty" xml:"WorkerId,omitempty"`
}

func (s PollCmdRequest) String() string {
	return tea.Prettify(s)
}

func (s PollCmdRequest) GoString() string {
	return s.String()
}

func (s *PollCmdRequest) SetQueue(v string) *PollCmdRequest {
	s.Queue = &v
	return s
}

func (s *PollCmdRequest) SetWaitSeconds(v string) *PollCmdRequest {
	s.WaitSeconds = &v
	return s
}

func (s *PollCmdRequest) SetWorkerId(v string) *PollCmdRequest {
	s.WorkerId = &v
	return s
}

type PollCmdResponseBody struct {
	HostId    *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PollCmdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PollCmdResponseBody) GoString() string {
	return s.String()
}

func (s *PollCmdResponseBody) SetHostId(v string) *PollCmdResponseBody {
	s.HostId = &v
	return s
}

func (s *PollCmdResponseBody) SetMessage(v string) *PollCmdResponseBody {
	s.Message = &v
	return s
}

func (s *PollCmdResponseBody) SetRequestId(v string) *PollCmdResponseBody {
	s.RequestId = &v
	return s
}

type PollCmdResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PollCmdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PollCmdResponse) String() string {
	return tea.Prettify(s)
}

func (s PollCmdResponse) GoString() string {
	return s.String()
}

func (s *PollCmdResponse) SetHeaders(v map[string]*string) *PollCmdResponse {
	s.Headers = v
	return s
}

func (s *PollCmdResponse) SetBody(v *PollCmdResponseBody) *PollCmdResponse {
	s.Body = v
	return s
}

type RecreateWorkerRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Project   *string `json:"Project,omitempty" xml:"Project,omitempty"`
	WorkerId  *string `json:"WorkerId,omitempty" xml:"WorkerId,omitempty"`
}

func (s RecreateWorkerRequest) String() string {
	return tea.Prettify(s)
}

func (s RecreateWorkerRequest) GoString() string {
	return s.String()
}

func (s *RecreateWorkerRequest) SetClusterId(v string) *RecreateWorkerRequest {
	s.ClusterId = &v
	return s
}

func (s *RecreateWorkerRequest) SetProject(v string) *RecreateWorkerRequest {
	s.Project = &v
	return s
}

func (s *RecreateWorkerRequest) SetWorkerId(v string) *RecreateWorkerRequest {
	s.WorkerId = &v
	return s
}

type RecreateWorkerResponseBody struct {
	HostId    *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecreateWorkerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecreateWorkerResponseBody) GoString() string {
	return s.String()
}

func (s *RecreateWorkerResponseBody) SetHostId(v string) *RecreateWorkerResponseBody {
	s.HostId = &v
	return s
}

func (s *RecreateWorkerResponseBody) SetRequestId(v string) *RecreateWorkerResponseBody {
	s.RequestId = &v
	return s
}

type RecreateWorkerResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecreateWorkerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecreateWorkerResponse) String() string {
	return tea.Prettify(s)
}

func (s RecreateWorkerResponse) GoString() string {
	return s.String()
}

func (s *RecreateWorkerResponse) SetHeaders(v map[string]*string) *RecreateWorkerResponse {
	s.Headers = v
	return s
}

func (s *RecreateWorkerResponse) SetBody(v *RecreateWorkerResponseBody) *RecreateWorkerResponse {
	s.Body = v
	return s
}

type RunJobRequest struct {
	ClientToken *string        `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Definition  *JobDefinition `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description *string        `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string        `json:"Name,omitempty" xml:"Name,omitempty"`
	Project     *string        `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s RunJobRequest) String() string {
	return tea.Prettify(s)
}

func (s RunJobRequest) GoString() string {
	return s.String()
}

func (s *RunJobRequest) SetClientToken(v string) *RunJobRequest {
	s.ClientToken = &v
	return s
}

func (s *RunJobRequest) SetDefinition(v *JobDefinition) *RunJobRequest {
	s.Definition = v
	return s
}

func (s *RunJobRequest) SetDescription(v string) *RunJobRequest {
	s.Description = &v
	return s
}

func (s *RunJobRequest) SetName(v string) *RunJobRequest {
	s.Name = &v
	return s
}

func (s *RunJobRequest) SetProject(v string) *RunJobRequest {
	s.Project = &v
	return s
}

type RunJobShrinkRequest struct {
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DefinitionShrink *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Project          *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s RunJobShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RunJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunJobShrinkRequest) SetClientToken(v string) *RunJobShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *RunJobShrinkRequest) SetDefinitionShrink(v string) *RunJobShrinkRequest {
	s.DefinitionShrink = &v
	return s
}

func (s *RunJobShrinkRequest) SetDescription(v string) *RunJobShrinkRequest {
	s.Description = &v
	return s
}

func (s *RunJobShrinkRequest) SetName(v string) *RunJobShrinkRequest {
	s.Name = &v
	return s
}

func (s *RunJobShrinkRequest) SetProject(v string) *RunJobShrinkRequest {
	s.Project = &v
	return s
}

type RunJobResponseBody struct {
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExitCode   *int32  `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	HostId     *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	Pid        *int32  `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Reason     *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State      *string `json:"State,omitempty" xml:"State,omitempty"`
	Worker     *string `json:"Worker,omitempty" xml:"Worker,omitempty"`
}

func (s RunJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunJobResponseBody) GoString() string {
	return s.String()
}

func (s *RunJobResponseBody) SetCreateTime(v string) *RunJobResponseBody {
	s.CreateTime = &v
	return s
}

func (s *RunJobResponseBody) SetEndTime(v string) *RunJobResponseBody {
	s.EndTime = &v
	return s
}

func (s *RunJobResponseBody) SetExitCode(v int32) *RunJobResponseBody {
	s.ExitCode = &v
	return s
}

func (s *RunJobResponseBody) SetHostId(v string) *RunJobResponseBody {
	s.HostId = &v
	return s
}

func (s *RunJobResponseBody) SetPid(v int32) *RunJobResponseBody {
	s.Pid = &v
	return s
}

func (s *RunJobResponseBody) SetReason(v string) *RunJobResponseBody {
	s.Reason = &v
	return s
}

func (s *RunJobResponseBody) SetRequestId(v string) *RunJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunJobResponseBody) SetStartTime(v string) *RunJobResponseBody {
	s.StartTime = &v
	return s
}

func (s *RunJobResponseBody) SetState(v string) *RunJobResponseBody {
	s.State = &v
	return s
}

func (s *RunJobResponseBody) SetWorker(v string) *RunJobResponseBody {
	s.Worker = &v
	return s
}

type RunJobResponse struct {
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RunJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RunJobResponse) String() string {
	return tea.Prettify(s)
}

func (s RunJobResponse) GoString() string {
	return s.String()
}

func (s *RunJobResponse) SetHeaders(v map[string]*string) *RunJobResponse {
	s.Headers = v
	return s
}

func (s *RunJobResponse) SetBody(v *RunJobResponseBody) *RunJobResponse {
	s.Body = v
	return s
}

type UpdateClusterRequest struct {
	ClusterId  *string            `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Definition *ClusterDefinition `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Project    *string            `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s UpdateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterRequest) GoString() string {
	return s.String()
}

func (s *UpdateClusterRequest) SetClusterId(v string) *UpdateClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateClusterRequest) SetDefinition(v *ClusterDefinition) *UpdateClusterRequest {
	s.Definition = v
	return s
}

func (s *UpdateClusterRequest) SetProject(v string) *UpdateClusterRequest {
	s.Project = &v
	return s
}

type UpdateClusterShrinkRequest struct {
	ClusterId        *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DefinitionShrink *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Project          *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s UpdateClusterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateClusterShrinkRequest) SetClusterId(v string) *UpdateClusterShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateClusterShrinkRequest) SetDefinitionShrink(v string) *UpdateClusterShrinkRequest {
	s.DefinitionShrink = &v
	return s
}

func (s *UpdateClusterShrinkRequest) SetProject(v string) *UpdateClusterShrinkRequest {
	s.Project = &v
	return s
}

type UpdateClusterResponseBody struct {
	HostId    *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClusterResponseBody) SetHostId(v string) *UpdateClusterResponseBody {
	s.HostId = &v
	return s
}

func (s *UpdateClusterResponseBody) SetRequestId(v string) *UpdateClusterResponseBody {
	s.RequestId = &v
	return s
}

type UpdateClusterResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterResponse) GoString() string {
	return s.String()
}

func (s *UpdateClusterResponse) SetHeaders(v map[string]*string) *UpdateClusterResponse {
	s.Headers = v
	return s
}

func (s *UpdateClusterResponse) SetBody(v *UpdateClusterResponseBody) *UpdateClusterResponse {
	s.Body = v
	return s
}

type UpdateJobQueueRequest struct {
	Definition *JobQueueDefinition `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Name       *string             `json:"Name,omitempty" xml:"Name,omitempty"`
	Project    *string             `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s UpdateJobQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobQueueRequest) GoString() string {
	return s.String()
}

func (s *UpdateJobQueueRequest) SetDefinition(v *JobQueueDefinition) *UpdateJobQueueRequest {
	s.Definition = v
	return s
}

func (s *UpdateJobQueueRequest) SetName(v string) *UpdateJobQueueRequest {
	s.Name = &v
	return s
}

func (s *UpdateJobQueueRequest) SetProject(v string) *UpdateJobQueueRequest {
	s.Project = &v
	return s
}

type UpdateJobQueueShrinkRequest struct {
	DefinitionShrink *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Project          *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s UpdateJobQueueShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobQueueShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateJobQueueShrinkRequest) SetDefinitionShrink(v string) *UpdateJobQueueShrinkRequest {
	s.DefinitionShrink = &v
	return s
}

func (s *UpdateJobQueueShrinkRequest) SetName(v string) *UpdateJobQueueShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateJobQueueShrinkRequest) SetProject(v string) *UpdateJobQueueShrinkRequest {
	s.Project = &v
	return s
}

type UpdateJobQueueResponseBody struct {
	HostId    *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateJobQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobQueueResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateJobQueueResponseBody) SetHostId(v string) *UpdateJobQueueResponseBody {
	s.HostId = &v
	return s
}

func (s *UpdateJobQueueResponseBody) SetRequestId(v string) *UpdateJobQueueResponseBody {
	s.RequestId = &v
	return s
}

type UpdateJobQueueResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateJobQueueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateJobQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobQueueResponse) GoString() string {
	return s.String()
}

func (s *UpdateJobQueueResponse) SetHeaders(v map[string]*string) *UpdateJobQueueResponse {
	s.Headers = v
	return s
}

func (s *UpdateJobQueueResponse) SetBody(v *UpdateJobQueueResponseBody) *UpdateJobQueueResponse {
	s.Body = v
	return s
}

type UpdateProjectRequest struct {
	Definition  *ProjectDefinition `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description *string            `json:"Description,omitempty" xml:"Description,omitempty"`
	Project     *string            `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s UpdateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequest) SetDefinition(v *ProjectDefinition) *UpdateProjectRequest {
	s.Definition = v
	return s
}

func (s *UpdateProjectRequest) SetDescription(v string) *UpdateProjectRequest {
	s.Description = &v
	return s
}

func (s *UpdateProjectRequest) SetProject(v string) *UpdateProjectRequest {
	s.Project = &v
	return s
}

type UpdateProjectShrinkRequest struct {
	DefinitionShrink *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Project          *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s UpdateProjectShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectShrinkRequest) SetDefinitionShrink(v string) *UpdateProjectShrinkRequest {
	s.DefinitionShrink = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetDescription(v string) *UpdateProjectShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetProject(v string) *UpdateProjectShrinkRequest {
	s.Project = &v
	return s
}

type UpdateProjectResponseBody struct {
	HostId    *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponseBody) SetHostId(v string) *UpdateProjectResponseBody {
	s.HostId = &v
	return s
}

func (s *UpdateProjectResponseBody) SetRequestId(v string) *UpdateProjectResponseBody {
	s.RequestId = &v
	return s
}

type UpdateProjectResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponse) SetHeaders(v map[string]*string) *UpdateProjectResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectResponse) SetBody(v *UpdateProjectResponseBody) *UpdateProjectResponse {
	s.Body = v
	return s
}

type UpdateWorkerStatusRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	WorkerId  *string `json:"WorkerId,omitempty" xml:"WorkerId,omitempty"`
}

func (s UpdateWorkerStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkerStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkerStatusRequest) SetClusterId(v string) *UpdateWorkerStatusRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateWorkerStatusRequest) SetStatus(v string) *UpdateWorkerStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateWorkerStatusRequest) SetWorkerId(v string) *UpdateWorkerStatusRequest {
	s.WorkerId = &v
	return s
}

type UpdateWorkerStatusResponseBody struct {
	HostId    *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWorkerStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkerStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkerStatusResponseBody) SetHostId(v string) *UpdateWorkerStatusResponseBody {
	s.HostId = &v
	return s
}

func (s *UpdateWorkerStatusResponseBody) SetRequestId(v string) *UpdateWorkerStatusResponseBody {
	s.RequestId = &v
	return s
}

type UpdateWorkerStatusResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateWorkerStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateWorkerStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkerStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkerStatusResponse) SetHeaders(v map[string]*string) *UpdateWorkerStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkerStatusResponse) SetBody(v *UpdateWorkerStatusResponseBody) *UpdateWorkerStatusResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("batchcompute"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelJobWithOptions(request *CancelJobRequest, runtime *util.RuntimeOptions) (_result *CancelJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CancelJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CancelJob"), tea.String("2018-12-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelJob(request *CancelJobRequest) (_result *CancelJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelJobResponse{}
	_body, _err := client.CancelJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateClusterWithOptions(tmpReq *CreateClusterRequest, runtime *util.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Definition))) {
		request.DefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Definition), tea.String("Definition"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCluster"), tea.String("2018-12-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCluster(request *CreateClusterRequest) (_result *CreateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClusterResponse{}
	_body, _err := client.CreateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateJobWithOptions(tmpReq *CreateJobRequest, runtime *util.RuntimeOptions) (_result *CreateJobResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Definition))) {
		request.DefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Definition), tea.String("Definition"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateJob"), tea.String("2018-12-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateJob(request *CreateJobRequest) (_result *CreateJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateJobResponse{}
	_body, _err := client.CreateJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateJobQueueWithOptions(tmpReq *CreateJobQueueRequest, runtime *util.RuntimeOptions) (_result *CreateJobQueueResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateJobQueueShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Definition))) {
		request.DefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Definition), tea.String("Definition"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateJobQueueResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateJobQueue"), tea.String("2018-12-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateJobQueue(request *CreateJobQueueRequest) (_result *CreateJobQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateJobQueueResponse{}
	_body, _err := client.CreateJobQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProjectWithOptions(tmpReq *CreateProjectRequest, runtime *util.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateProjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Definition))) {
		request.DefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Definition), tea.String("Definition"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateProject"), tea.String("2018-12-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProject(request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteClusterWithOptions(request *DeleteClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteCluster"), tea.String("2018-12-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCluster(request *DeleteClusterRequest) (_result *DeleteClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DeleteClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteJobWithOptions(request *DeleteJobRequest, runtime *util.RuntimeOptions) (_result *DeleteJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteJob"), tea.String("2018-12-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteJob(request *DeleteJobRequest) (_result *DeleteJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteJobResponse{}
	_body, _err := client.DeleteJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteJobQueueWithOptions(request *DeleteJobQueueRequest, runtime *util.RuntimeOptions) (_result *DeleteJobQueueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteJobQueueResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteJobQueue"), tea.String("2018-12-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteJobQueue(request *DeleteJobQueueRequest) (_result *DeleteJobQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteJobQueueResponse{}
	_body, _err := client.DeleteJobQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProjectWithOptions(request *DeleteProjectRequest, runtime *util.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteProject"), tea.String("2018-12-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProject(request *DeleteProjectRequest) (_result *DeleteProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProjectResponse{}
	_body, _err := client.DeleteProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetClusterWithOptions(request *GetClusterRequest, runtime *util.RuntimeOptions) (_result *GetClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetCluster"), tea.String("2018-12-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCluster(request *GetClusterRequest) (_result *GetClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetClusterResponse{}
	_body, _err := client.GetClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobWithOptions(request *GetJobRequest, runtime *util.RuntimeOptions) (_result *GetJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetJob"), tea.String("2018-12-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJob(request *GetJobRequest) (_result *GetJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJobResponse{}
	_body, _err := client.GetJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobQueueWithOptions(request *GetJobQueueRequest, runtime *util.RuntimeOptions) (_result *GetJobQueueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetJobQueueResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetJobQueue"), tea.String("2018-12-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobQueue(request *GetJobQueueRequest) (_result *GetJobQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJobQueueResponse{}
	_body, _err := client.GetJobQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProjectWithOptions(request *GetProjectRequest, runtime *util.RuntimeOptions) (_result *GetProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetProject"), tea.String("2018-12-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProject(request *GetProjectRequest) (_result *GetProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProjectResponse{}
	_body, _err := client.GetProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWorkerWithOptions(request *GetWorkerRequest, runtime *util.RuntimeOptions) (_result *GetWorkerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetWorkerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetWorker"), tea.String("2018-12-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWorker(request *GetWorkerRequest) (_result *GetWorkerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWorkerResponse{}
	_body, _err := client.GetWorkerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) KillWorkerWithOptions(request *KillWorkerRequest, runtime *util.RuntimeOptions) (_result *KillWorkerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &KillWorkerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("KillWorker"), tea.String("2018-12-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) KillWorker(request *KillWorkerRequest) (_result *KillWorkerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &KillWorkerResponse{}
	_body, _err := client.KillWorkerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClustersWithOptions(request *ListClustersRequest, runtime *util.RuntimeOptions) (_result *ListClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListClustersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListClusters"), tea.String("2018-12-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusters(request *ListClustersRequest) (_result *ListClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClustersResponse{}
	_body, _err := client.ListClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListJobQueuesWithOptions(request *ListJobQueuesRequest, runtime *util.RuntimeOptions) (_result *ListJobQueuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListJobQueuesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListJobQueues"), tea.String("2018-12-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListJobQueues(request *ListJobQueuesRequest) (_result *ListJobQueuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListJobQueuesResponse{}
	_body, _err := client.ListJobQueuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListJobsWithOptions(request *ListJobsRequest, runtime *util.RuntimeOptions) (_result *ListJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListJobsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListJobs"), tea.String("2018-12-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListJobs(request *ListJobsRequest) (_result *ListJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListJobsResponse{}
	_body, _err := client.ListJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProjectsWithOptions(request *ListProjectsRequest, runtime *util.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListProjectsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListProjects"), tea.String("2018-12-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProjects(request *ListProjectsRequest) (_result *ListProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProjectsResponse{}
	_body, _err := client.ListProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRegionsWithOptions(request *ListRegionsRequest, runtime *util.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRegions"), tea.String("2018-12-13"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRegions(request *ListRegionsRequest) (_result *ListRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListWorkersWithOptions(request *ListWorkersRequest, runtime *util.RuntimeOptions) (_result *ListWorkersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListWorkersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListWorkers"), tea.String("2018-12-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListWorkers(request *ListWorkersRequest) (_result *ListWorkersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWorkersResponse{}
	_body, _err := client.ListWorkersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenBatchComputeServiceWithOptions(runtime *util.RuntimeOptions) (_result *OpenBatchComputeServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &OpenBatchComputeServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OpenBatchComputeService"), tea.String("2018-12-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenBatchComputeService() (_result *OpenBatchComputeServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenBatchComputeServiceResponse{}
	_body, _err := client.OpenBatchComputeServiceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PollCmdWithOptions(request *PollCmdRequest, runtime *util.RuntimeOptions) (_result *PollCmdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PollCmdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PollCmd"), tea.String("2018-12-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PollCmd(request *PollCmdRequest) (_result *PollCmdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PollCmdResponse{}
	_body, _err := client.PollCmdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecreateWorkerWithOptions(request *RecreateWorkerRequest, runtime *util.RuntimeOptions) (_result *RecreateWorkerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecreateWorkerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecreateWorker"), tea.String("2018-12-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecreateWorker(request *RecreateWorkerRequest) (_result *RecreateWorkerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecreateWorkerResponse{}
	_body, _err := client.RecreateWorkerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunJobWithOptions(tmpReq *RunJobRequest, runtime *util.RuntimeOptions) (_result *RunJobResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RunJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Definition))) {
		request.DefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Definition), tea.String("Definition"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RunJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RunJob"), tea.String("2018-12-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunJob(request *RunJobRequest) (_result *RunJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunJobResponse{}
	_body, _err := client.RunJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateClusterWithOptions(tmpReq *UpdateClusterRequest, runtime *util.RuntimeOptions) (_result *UpdateClusterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Definition))) {
		request.DefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Definition), tea.String("Definition"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateCluster"), tea.String("2018-12-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCluster(request *UpdateClusterRequest) (_result *UpdateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateClusterResponse{}
	_body, _err := client.UpdateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateJobQueueWithOptions(tmpReq *UpdateJobQueueRequest, runtime *util.RuntimeOptions) (_result *UpdateJobQueueResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateJobQueueShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Definition))) {
		request.DefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Definition), tea.String("Definition"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateJobQueueResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateJobQueue"), tea.String("2018-12-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateJobQueue(request *UpdateJobQueueRequest) (_result *UpdateJobQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateJobQueueResponse{}
	_body, _err := client.UpdateJobQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProjectWithOptions(tmpReq *UpdateProjectRequest, runtime *util.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateProjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Definition))) {
		request.DefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Definition), tea.String("Definition"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateProject"), tea.String("2018-12-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProject(request *UpdateProjectRequest) (_result *UpdateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateProjectResponse{}
	_body, _err := client.UpdateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWorkerStatusWithOptions(request *UpdateWorkerStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateWorkerStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateWorkerStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateWorkerStatus"), tea.String("2018-12-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWorkerStatus(request *UpdateWorkerStatusRequest) (_result *UpdateWorkerStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWorkerStatusResponse{}
	_body, _err := client.UpdateWorkerStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
