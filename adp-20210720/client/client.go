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

type ProductComponentRelationDetail struct {
	// appVersion
	AppVersion *string `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	// category
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// class
	Class *string `json:"class,omitempty" xml:"class,omitempty"`
	// componentName
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// componentOrchestrationValues
	ComponentOrchestrationValues *string `json:"componentOrchestrationValues,omitempty" xml:"componentOrchestrationValues,omitempty"`
	// componentUID
	ComponentUID *string `json:"componentUID,omitempty" xml:"componentUID,omitempty"`
	// componentVersionUID
	ComponentVersionUID *string `json:"componentVersionUID,omitempty" xml:"componentVersionUID,omitempty"`
	// createdAt
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// documents
	Documents *string `json:"documents,omitempty" xml:"documents,omitempty"`
	// enable
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// imagesMapping
	ImagesMapping *string `json:"imagesMapping,omitempty" xml:"imagesMapping,omitempty"`
	// namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// orchestrationType
	OrchestrationType *string `json:"orchestrationType,omitempty" xml:"orchestrationType,omitempty"`
	// parentComponent
	ParentComponent *bool `json:"parentComponent,omitempty" xml:"parentComponent,omitempty"`
	// parentComponentVersionRelationUID
	ParentComponentVersionRelationUID *string `json:"parentComponentVersionRelationUID,omitempty" xml:"parentComponentVersionRelationUID,omitempty"`
	// parentComponentVersionUID
	ParentComponentVersionUID *string `json:"parentComponentVersionUID,omitempty" xml:"parentComponentVersionUID,omitempty"`
	// priority
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
	// productVersionUID
	ProductVersionUID *string `json:"productVersionUID,omitempty" xml:"productVersionUID,omitempty"`
	// provider
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// public
	Public *bool `json:"public,omitempty" xml:"public,omitempty"`
	// readme
	Readme *string `json:"readme,omitempty" xml:"readme,omitempty"`
	// relationUID
	RelationUID *string `json:"relationUID,omitempty" xml:"relationUID,omitempty"`
	// releaseName
	ReleaseName *string `json:"releaseName,omitempty" xml:"releaseName,omitempty"`
	// resources
	Resources *string `json:"resources,omitempty" xml:"resources,omitempty"`
	// sequence
	Sequence *int32 `json:"sequence,omitempty" xml:"sequence,omitempty"`
	// singleton
	Singleton *bool `json:"singleton,omitempty" xml:"singleton,omitempty"`
	// source
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// version
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ProductComponentRelationDetail) String() string {
	return tea.Prettify(s)
}

func (s ProductComponentRelationDetail) GoString() string {
	return s.String()
}

func (s *ProductComponentRelationDetail) SetAppVersion(v string) *ProductComponentRelationDetail {
	s.AppVersion = &v
	return s
}

func (s *ProductComponentRelationDetail) SetCategory(v string) *ProductComponentRelationDetail {
	s.Category = &v
	return s
}

func (s *ProductComponentRelationDetail) SetClass(v string) *ProductComponentRelationDetail {
	s.Class = &v
	return s
}

func (s *ProductComponentRelationDetail) SetComponentName(v string) *ProductComponentRelationDetail {
	s.ComponentName = &v
	return s
}

func (s *ProductComponentRelationDetail) SetComponentOrchestrationValues(v string) *ProductComponentRelationDetail {
	s.ComponentOrchestrationValues = &v
	return s
}

func (s *ProductComponentRelationDetail) SetComponentUID(v string) *ProductComponentRelationDetail {
	s.ComponentUID = &v
	return s
}

func (s *ProductComponentRelationDetail) SetComponentVersionUID(v string) *ProductComponentRelationDetail {
	s.ComponentVersionUID = &v
	return s
}

func (s *ProductComponentRelationDetail) SetCreatedAt(v string) *ProductComponentRelationDetail {
	s.CreatedAt = &v
	return s
}

func (s *ProductComponentRelationDetail) SetDescription(v string) *ProductComponentRelationDetail {
	s.Description = &v
	return s
}

func (s *ProductComponentRelationDetail) SetDocuments(v string) *ProductComponentRelationDetail {
	s.Documents = &v
	return s
}

func (s *ProductComponentRelationDetail) SetEnable(v bool) *ProductComponentRelationDetail {
	s.Enable = &v
	return s
}

func (s *ProductComponentRelationDetail) SetImagesMapping(v string) *ProductComponentRelationDetail {
	s.ImagesMapping = &v
	return s
}

func (s *ProductComponentRelationDetail) SetNamespace(v string) *ProductComponentRelationDetail {
	s.Namespace = &v
	return s
}

func (s *ProductComponentRelationDetail) SetOrchestrationType(v string) *ProductComponentRelationDetail {
	s.OrchestrationType = &v
	return s
}

func (s *ProductComponentRelationDetail) SetParentComponent(v bool) *ProductComponentRelationDetail {
	s.ParentComponent = &v
	return s
}

func (s *ProductComponentRelationDetail) SetParentComponentVersionRelationUID(v string) *ProductComponentRelationDetail {
	s.ParentComponentVersionRelationUID = &v
	return s
}

func (s *ProductComponentRelationDetail) SetParentComponentVersionUID(v string) *ProductComponentRelationDetail {
	s.ParentComponentVersionUID = &v
	return s
}

func (s *ProductComponentRelationDetail) SetPriority(v int32) *ProductComponentRelationDetail {
	s.Priority = &v
	return s
}

func (s *ProductComponentRelationDetail) SetProductVersionUID(v string) *ProductComponentRelationDetail {
	s.ProductVersionUID = &v
	return s
}

func (s *ProductComponentRelationDetail) SetProvider(v string) *ProductComponentRelationDetail {
	s.Provider = &v
	return s
}

func (s *ProductComponentRelationDetail) SetPublic(v bool) *ProductComponentRelationDetail {
	s.Public = &v
	return s
}

func (s *ProductComponentRelationDetail) SetReadme(v string) *ProductComponentRelationDetail {
	s.Readme = &v
	return s
}

func (s *ProductComponentRelationDetail) SetRelationUID(v string) *ProductComponentRelationDetail {
	s.RelationUID = &v
	return s
}

func (s *ProductComponentRelationDetail) SetReleaseName(v string) *ProductComponentRelationDetail {
	s.ReleaseName = &v
	return s
}

func (s *ProductComponentRelationDetail) SetResources(v string) *ProductComponentRelationDetail {
	s.Resources = &v
	return s
}

func (s *ProductComponentRelationDetail) SetSequence(v int32) *ProductComponentRelationDetail {
	s.Sequence = &v
	return s
}

func (s *ProductComponentRelationDetail) SetSingleton(v bool) *ProductComponentRelationDetail {
	s.Singleton = &v
	return s
}

func (s *ProductComponentRelationDetail) SetSource(v string) *ProductComponentRelationDetail {
	s.Source = &v
	return s
}

func (s *ProductComponentRelationDetail) SetVersion(v string) *ProductComponentRelationDetail {
	s.Version = &v
	return s
}

type Platform struct {
	// architecture
	Architecture *string `json:"architecture,omitempty" xml:"architecture,omitempty"`
	// os
	Os *string `json:"os,omitempty" xml:"os,omitempty"`
}

func (s Platform) String() string {
	return tea.Prettify(s)
}

func (s Platform) GoString() string {
	return s.String()
}

func (s *Platform) SetArchitecture(v string) *Platform {
	s.Architecture = &v
	return s
}

func (s *Platform) SetOs(v string) *Platform {
	s.Os = &v
	return s
}

type FoundationVersion struct {
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// documents
	Documents *string `json:"documents,omitempty" xml:"documents,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// status
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// uid
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
	// version
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// platforms
	Platforms []*Platform `json:"platforms,omitempty" xml:"platforms,omitempty" type:"Repeated"`
}

func (s FoundationVersion) String() string {
	return tea.Prettify(s)
}

func (s FoundationVersion) GoString() string {
	return s.String()
}

func (s *FoundationVersion) SetDescription(v string) *FoundationVersion {
	s.Description = &v
	return s
}

func (s *FoundationVersion) SetDocuments(v string) *FoundationVersion {
	s.Documents = &v
	return s
}

func (s *FoundationVersion) SetName(v string) *FoundationVersion {
	s.Name = &v
	return s
}

func (s *FoundationVersion) SetStatus(v string) *FoundationVersion {
	s.Status = &v
	return s
}

func (s *FoundationVersion) SetUid(v string) *FoundationVersion {
	s.Uid = &v
	return s
}

func (s *FoundationVersion) SetVersion(v string) *FoundationVersion {
	s.Version = &v
	return s
}

func (s *FoundationVersion) SetPlatforms(v []*Platform) *FoundationVersion {
	s.Platforms = v
	return s
}

type GetInstanceInfoResponse struct {
	// identifier
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// hostName
	HostName *string `json:"hostName,omitempty" xml:"hostName,omitempty"`
	// os
	Os *string `json:"os,omitempty" xml:"os,omitempty"`
	// osVersion
	OsVersion *string `json:"osVersion,omitempty" xml:"osVersion,omitempty"`
	// arch
	Arch *string `json:"arch,omitempty" xml:"arch,omitempty"`
	// kernel
	Kernel *string `json:"kernel,omitempty" xml:"kernel,omitempty"`
	// macAddress
	MacAddress *string `json:"macAddress,omitempty" xml:"macAddress,omitempty"`
	// cpu
	Cpu *string `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// memory
	Memory *string `json:"memory,omitempty" xml:"memory,omitempty"`
	// systemDisk
	SystemDisk []*Disk `json:"systemDisk,omitempty" xml:"systemDisk,omitempty" type:"Repeated"`
	// dataDisk
	DataDisk []*Disk `json:"dataDisk,omitempty" xml:"dataDisk,omitempty" type:"Repeated"`
	// privateIP
	PrivateIP *string `json:"privateIP,omitempty" xml:"privateIP,omitempty"`
	// publicIP
	PublicIP *string `json:"publicIP,omitempty" xml:"publicIP,omitempty"`
	// internetBandwidth
	InternetBandwidth *int32 `json:"internetBandwidth,omitempty" xml:"internetBandwidth,omitempty"`
	// networkCards
	NetworkCards []*GetInstanceInfoResponseNetworkCards `json:"networkCards,omitempty" xml:"networkCards,omitempty" type:"Repeated"`
	// imageID
	ImageID *string `json:"imageID,omitempty" xml:"imageID,omitempty"`
	// uid
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
	// instanceType
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// systemInfo
	SystemInfo *string `json:"systemInfo,omitempty" xml:"systemInfo,omitempty"`
	// rootPassword
	RootPassword *string `json:"rootPassword,omitempty" xml:"rootPassword,omitempty"`
	// labels
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// taints
	Taints []*GetInstanceInfoResponseTaints `json:"taints,omitempty" xml:"taints,omitempty" type:"Repeated"`
	// annotations
	Annotations map[string]*string `json:"annotations,omitempty" xml:"annotations,omitempty"`
}

func (s GetInstanceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceInfoResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceInfoResponse) SetIdentifier(v string) *GetInstanceInfoResponse {
	s.Identifier = &v
	return s
}

func (s *GetInstanceInfoResponse) SetHostName(v string) *GetInstanceInfoResponse {
	s.HostName = &v
	return s
}

func (s *GetInstanceInfoResponse) SetOs(v string) *GetInstanceInfoResponse {
	s.Os = &v
	return s
}

func (s *GetInstanceInfoResponse) SetOsVersion(v string) *GetInstanceInfoResponse {
	s.OsVersion = &v
	return s
}

func (s *GetInstanceInfoResponse) SetArch(v string) *GetInstanceInfoResponse {
	s.Arch = &v
	return s
}

func (s *GetInstanceInfoResponse) SetKernel(v string) *GetInstanceInfoResponse {
	s.Kernel = &v
	return s
}

func (s *GetInstanceInfoResponse) SetMacAddress(v string) *GetInstanceInfoResponse {
	s.MacAddress = &v
	return s
}

func (s *GetInstanceInfoResponse) SetCpu(v string) *GetInstanceInfoResponse {
	s.Cpu = &v
	return s
}

func (s *GetInstanceInfoResponse) SetMemory(v string) *GetInstanceInfoResponse {
	s.Memory = &v
	return s
}

func (s *GetInstanceInfoResponse) SetSystemDisk(v []*Disk) *GetInstanceInfoResponse {
	s.SystemDisk = v
	return s
}

func (s *GetInstanceInfoResponse) SetDataDisk(v []*Disk) *GetInstanceInfoResponse {
	s.DataDisk = v
	return s
}

func (s *GetInstanceInfoResponse) SetPrivateIP(v string) *GetInstanceInfoResponse {
	s.PrivateIP = &v
	return s
}

func (s *GetInstanceInfoResponse) SetPublicIP(v string) *GetInstanceInfoResponse {
	s.PublicIP = &v
	return s
}

func (s *GetInstanceInfoResponse) SetInternetBandwidth(v int32) *GetInstanceInfoResponse {
	s.InternetBandwidth = &v
	return s
}

func (s *GetInstanceInfoResponse) SetNetworkCards(v []*GetInstanceInfoResponseNetworkCards) *GetInstanceInfoResponse {
	s.NetworkCards = v
	return s
}

func (s *GetInstanceInfoResponse) SetImageID(v string) *GetInstanceInfoResponse {
	s.ImageID = &v
	return s
}

func (s *GetInstanceInfoResponse) SetUid(v string) *GetInstanceInfoResponse {
	s.Uid = &v
	return s
}

func (s *GetInstanceInfoResponse) SetInstanceType(v string) *GetInstanceInfoResponse {
	s.InstanceType = &v
	return s
}

func (s *GetInstanceInfoResponse) SetSystemInfo(v string) *GetInstanceInfoResponse {
	s.SystemInfo = &v
	return s
}

func (s *GetInstanceInfoResponse) SetRootPassword(v string) *GetInstanceInfoResponse {
	s.RootPassword = &v
	return s
}

func (s *GetInstanceInfoResponse) SetLabels(v map[string]*string) *GetInstanceInfoResponse {
	s.Labels = v
	return s
}

func (s *GetInstanceInfoResponse) SetTaints(v []*GetInstanceInfoResponseTaints) *GetInstanceInfoResponse {
	s.Taints = v
	return s
}

func (s *GetInstanceInfoResponse) SetAnnotations(v map[string]*string) *GetInstanceInfoResponse {
	s.Annotations = v
	return s
}

type GetInstanceInfoResponseNetworkCards struct {
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// ip
	Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
}

func (s GetInstanceInfoResponseNetworkCards) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceInfoResponseNetworkCards) GoString() string {
	return s.String()
}

func (s *GetInstanceInfoResponseNetworkCards) SetName(v string) *GetInstanceInfoResponseNetworkCards {
	s.Name = &v
	return s
}

func (s *GetInstanceInfoResponseNetworkCards) SetIp(v string) *GetInstanceInfoResponseNetworkCards {
	s.Ip = &v
	return s
}

type GetInstanceInfoResponseTaints struct {
	// key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// effect
	Effect *string `json:"effect,omitempty" xml:"effect,omitempty"`
}

func (s GetInstanceInfoResponseTaints) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceInfoResponseTaints) GoString() string {
	return s.String()
}

func (s *GetInstanceInfoResponseTaints) SetKey(v string) *GetInstanceInfoResponseTaints {
	s.Key = &v
	return s
}

func (s *GetInstanceInfoResponseTaints) SetValue(v string) *GetInstanceInfoResponseTaints {
	s.Value = &v
	return s
}

func (s *GetInstanceInfoResponseTaints) SetEffect(v string) *GetInstanceInfoResponseTaints {
	s.Effect = &v
	return s
}

type InstanceInfo struct {
	// identifier
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// hostName
	HostName *string `json:"hostName,omitempty" xml:"hostName,omitempty"`
	// os
	Os *string `json:"os,omitempty" xml:"os,omitempty"`
	// osVersion
	OsVersion *string `json:"osVersion,omitempty" xml:"osVersion,omitempty"`
	// arch
	Arch *string `json:"arch,omitempty" xml:"arch,omitempty"`
	// kernel
	Kernel *string `json:"kernel,omitempty" xml:"kernel,omitempty"`
	// macAddress
	MacAddress *string `json:"macAddress,omitempty" xml:"macAddress,omitempty"`
	// cpu
	Cpu *string `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// memory
	Memory *string `json:"memory,omitempty" xml:"memory,omitempty"`
	// systemDisk
	SystemDisk []*Disk `json:"systemDisk,omitempty" xml:"systemDisk,omitempty" type:"Repeated"`
	// dataDisk
	DataDisk []*Disk `json:"dataDisk,omitempty" xml:"dataDisk,omitempty" type:"Repeated"`
	// privateIP
	PrivateIP *string `json:"privateIP,omitempty" xml:"privateIP,omitempty"`
	// publicIP
	PublicIP *string `json:"publicIP,omitempty" xml:"publicIP,omitempty"`
	// internetBandwidth
	InternetBandwidth *int32 `json:"internetBandwidth,omitempty" xml:"internetBandwidth,omitempty"`
	// networkCards
	NetworkCards []*InstanceInfoNetworkCards `json:"networkCards,omitempty" xml:"networkCards,omitempty" type:"Repeated"`
	// imageID
	ImageID *string `json:"imageID,omitempty" xml:"imageID,omitempty"`
	// instanceType
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// systemInfo
	SystemInfo *string `json:"systemInfo,omitempty" xml:"systemInfo,omitempty"`
	// rootPassword
	RootPassword *string `json:"rootPassword,omitempty" xml:"rootPassword,omitempty"`
	// labels
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// taints
	Taints []*InstanceInfoTaints `json:"taints,omitempty" xml:"taints,omitempty" type:"Repeated"`
	// annotations
	Annotations map[string]*string `json:"annotations,omitempty" xml:"annotations,omitempty"`
}

func (s InstanceInfo) String() string {
	return tea.Prettify(s)
}

func (s InstanceInfo) GoString() string {
	return s.String()
}

func (s *InstanceInfo) SetIdentifier(v string) *InstanceInfo {
	s.Identifier = &v
	return s
}

func (s *InstanceInfo) SetHostName(v string) *InstanceInfo {
	s.HostName = &v
	return s
}

func (s *InstanceInfo) SetOs(v string) *InstanceInfo {
	s.Os = &v
	return s
}

func (s *InstanceInfo) SetOsVersion(v string) *InstanceInfo {
	s.OsVersion = &v
	return s
}

func (s *InstanceInfo) SetArch(v string) *InstanceInfo {
	s.Arch = &v
	return s
}

func (s *InstanceInfo) SetKernel(v string) *InstanceInfo {
	s.Kernel = &v
	return s
}

func (s *InstanceInfo) SetMacAddress(v string) *InstanceInfo {
	s.MacAddress = &v
	return s
}

func (s *InstanceInfo) SetCpu(v string) *InstanceInfo {
	s.Cpu = &v
	return s
}

func (s *InstanceInfo) SetMemory(v string) *InstanceInfo {
	s.Memory = &v
	return s
}

func (s *InstanceInfo) SetSystemDisk(v []*Disk) *InstanceInfo {
	s.SystemDisk = v
	return s
}

func (s *InstanceInfo) SetDataDisk(v []*Disk) *InstanceInfo {
	s.DataDisk = v
	return s
}

func (s *InstanceInfo) SetPrivateIP(v string) *InstanceInfo {
	s.PrivateIP = &v
	return s
}

func (s *InstanceInfo) SetPublicIP(v string) *InstanceInfo {
	s.PublicIP = &v
	return s
}

func (s *InstanceInfo) SetInternetBandwidth(v int32) *InstanceInfo {
	s.InternetBandwidth = &v
	return s
}

func (s *InstanceInfo) SetNetworkCards(v []*InstanceInfoNetworkCards) *InstanceInfo {
	s.NetworkCards = v
	return s
}

func (s *InstanceInfo) SetImageID(v string) *InstanceInfo {
	s.ImageID = &v
	return s
}

func (s *InstanceInfo) SetInstanceType(v string) *InstanceInfo {
	s.InstanceType = &v
	return s
}

func (s *InstanceInfo) SetSystemInfo(v string) *InstanceInfo {
	s.SystemInfo = &v
	return s
}

func (s *InstanceInfo) SetRootPassword(v string) *InstanceInfo {
	s.RootPassword = &v
	return s
}

func (s *InstanceInfo) SetLabels(v map[string]*string) *InstanceInfo {
	s.Labels = v
	return s
}

func (s *InstanceInfo) SetTaints(v []*InstanceInfoTaints) *InstanceInfo {
	s.Taints = v
	return s
}

func (s *InstanceInfo) SetAnnotations(v map[string]*string) *InstanceInfo {
	s.Annotations = v
	return s
}

type InstanceInfoNetworkCards struct {
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// ip
	Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
}

func (s InstanceInfoNetworkCards) String() string {
	return tea.Prettify(s)
}

func (s InstanceInfoNetworkCards) GoString() string {
	return s.String()
}

func (s *InstanceInfoNetworkCards) SetName(v string) *InstanceInfoNetworkCards {
	s.Name = &v
	return s
}

func (s *InstanceInfoNetworkCards) SetIp(v string) *InstanceInfoNetworkCards {
	s.Ip = &v
	return s
}

type InstanceInfoTaints struct {
	// key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// effect
	Effect *string `json:"effect,omitempty" xml:"effect,omitempty"`
}

func (s InstanceInfoTaints) String() string {
	return tea.Prettify(s)
}

func (s InstanceInfoTaints) GoString() string {
	return s.String()
}

func (s *InstanceInfoTaints) SetKey(v string) *InstanceInfoTaints {
	s.Key = &v
	return s
}

func (s *InstanceInfoTaints) SetValue(v string) *InstanceInfoTaints {
	s.Value = &v
	return s
}

func (s *InstanceInfoTaints) SetEffect(v string) *InstanceInfoTaints {
	s.Effect = &v
	return s
}

type Disk struct {
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// capacity
	Capacity *int32 `json:"capacity,omitempty" xml:"capacity,omitempty"`
	// remain
	Remain *int32 `json:"remain,omitempty" xml:"remain,omitempty"`
	// fsType
	FsType *string `json:"fsType,omitempty" xml:"fsType,omitempty"`
	// mountPoint
	MountPoint *string `json:"mountPoint,omitempty" xml:"mountPoint,omitempty"`
	// type
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s Disk) String() string {
	return tea.Prettify(s)
}

func (s Disk) GoString() string {
	return s.String()
}

func (s *Disk) SetName(v string) *Disk {
	s.Name = &v
	return s
}

func (s *Disk) SetCapacity(v int32) *Disk {
	s.Capacity = &v
	return s
}

func (s *Disk) SetRemain(v int32) *Disk {
	s.Remain = &v
	return s
}

func (s *Disk) SetFsType(v string) *Disk {
	s.FsType = &v
	return s
}

func (s *Disk) SetMountPoint(v string) *Disk {
	s.MountPoint = &v
	return s
}

func (s *Disk) SetType(v string) *Disk {
	s.Type = &v
	return s
}

type ComponentVersion struct {
	// appVersion
	AppVersion *string `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	// componentName
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// componentUID
	ComponentUID *string `json:"componentUID,omitempty" xml:"componentUID,omitempty"`
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// documents
	Documents *string `json:"documents,omitempty" xml:"documents,omitempty"`
	// imagesMapping
	ImagesMapping *string `json:"imagesMapping,omitempty" xml:"imagesMapping,omitempty"`
	// namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// orchestrationType
	OrchestrationType *string `json:"orchestrationType,omitempty" xml:"orchestrationType,omitempty"`
	// orchestrationValues
	OrchestrationValues *string `json:"orchestrationValues,omitempty" xml:"orchestrationValues,omitempty"`
	// packageURL
	PackageURL *string `json:"packageURL,omitempty" xml:"packageURL,omitempty"`
	// parentComponent
	ParentComponent *bool `json:"parentComponent,omitempty" xml:"parentComponent,omitempty"`
	// platforms
	Platforms []*Platform `json:"platforms,omitempty" xml:"platforms,omitempty" type:"Repeated"`
	// readme
	Readme *string `json:"readme,omitempty" xml:"readme,omitempty"`
	// resources
	Resources *string `json:"resources,omitempty" xml:"resources,omitempty"`
	// source
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// uid
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
	// version
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ComponentVersion) String() string {
	return tea.Prettify(s)
}

func (s ComponentVersion) GoString() string {
	return s.String()
}

func (s *ComponentVersion) SetAppVersion(v string) *ComponentVersion {
	s.AppVersion = &v
	return s
}

func (s *ComponentVersion) SetComponentName(v string) *ComponentVersion {
	s.ComponentName = &v
	return s
}

func (s *ComponentVersion) SetComponentUID(v string) *ComponentVersion {
	s.ComponentUID = &v
	return s
}

func (s *ComponentVersion) SetDescription(v string) *ComponentVersion {
	s.Description = &v
	return s
}

func (s *ComponentVersion) SetDocuments(v string) *ComponentVersion {
	s.Documents = &v
	return s
}

func (s *ComponentVersion) SetImagesMapping(v string) *ComponentVersion {
	s.ImagesMapping = &v
	return s
}

func (s *ComponentVersion) SetNamespace(v string) *ComponentVersion {
	s.Namespace = &v
	return s
}

func (s *ComponentVersion) SetOrchestrationType(v string) *ComponentVersion {
	s.OrchestrationType = &v
	return s
}

func (s *ComponentVersion) SetOrchestrationValues(v string) *ComponentVersion {
	s.OrchestrationValues = &v
	return s
}

func (s *ComponentVersion) SetPackageURL(v string) *ComponentVersion {
	s.PackageURL = &v
	return s
}

func (s *ComponentVersion) SetParentComponent(v bool) *ComponentVersion {
	s.ParentComponent = &v
	return s
}

func (s *ComponentVersion) SetPlatforms(v []*Platform) *ComponentVersion {
	s.Platforms = v
	return s
}

func (s *ComponentVersion) SetReadme(v string) *ComponentVersion {
	s.Readme = &v
	return s
}

func (s *ComponentVersion) SetResources(v string) *ComponentVersion {
	s.Resources = &v
	return s
}

func (s *ComponentVersion) SetSource(v string) *ComponentVersion {
	s.Source = &v
	return s
}

func (s *ComponentVersion) SetUid(v string) *ComponentVersion {
	s.Uid = &v
	return s
}

func (s *ComponentVersion) SetVersion(v string) *ComponentVersion {
	s.Version = &v
	return s
}

type GetEnvironmentResponseBody struct {
	Data *GetEnvironmentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Code *string                         `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string                         `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s GetEnvironmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *GetEnvironmentResponseBody) SetData(v *GetEnvironmentResponseBodyData) *GetEnvironmentResponseBody {
	s.Data = v
	return s
}

func (s *GetEnvironmentResponseBody) SetCode(v string) *GetEnvironmentResponseBody {
	s.Code = &v
	return s
}

func (s *GetEnvironmentResponseBody) SetMsg(v string) *GetEnvironmentResponseBody {
	s.Msg = &v
	return s
}

type GetEnvironmentResponseBodyData struct {
	ClusterId      *string                                 `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	CreatedAt      *string                                 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description    *string                                 `json:"description,omitempty" xml:"description,omitempty"`
	InstanceList   []*InstanceInfo                         `json:"instanceList,omitempty" xml:"instanceList,omitempty" type:"Repeated"`
	Location       *string                                 `json:"location,omitempty" xml:"location,omitempty"`
	Name           *string                                 `json:"name,omitempty" xml:"name,omitempty"`
	Platform       *GetEnvironmentResponseBodyDataPlatform `json:"platform,omitempty" xml:"platform,omitempty" type:"Struct"`
	ProductName    *string                                 `json:"productName,omitempty" xml:"productName,omitempty"`
	ProductVersion *string                                 `json:"productVersion,omitempty" xml:"productVersion,omitempty"`
	Uid            *string                                 `json:"uid,omitempty" xml:"uid,omitempty"`
	VendorConfig   *string                                 `json:"vendorConfig,omitempty" xml:"vendorConfig,omitempty"`
	VendorType     *string                                 `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
	InstanceStatus *string                                 `json:"instanceStatus,omitempty" xml:"instanceStatus,omitempty"`
}

func (s GetEnvironmentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEnvironmentResponseBodyData) SetClusterId(v string) *GetEnvironmentResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetCreatedAt(v string) *GetEnvironmentResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetDescription(v string) *GetEnvironmentResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetInstanceList(v []*InstanceInfo) *GetEnvironmentResponseBodyData {
	s.InstanceList = v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetLocation(v string) *GetEnvironmentResponseBodyData {
	s.Location = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetName(v string) *GetEnvironmentResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetPlatform(v *GetEnvironmentResponseBodyDataPlatform) *GetEnvironmentResponseBodyData {
	s.Platform = v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetProductName(v string) *GetEnvironmentResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetProductVersion(v string) *GetEnvironmentResponseBodyData {
	s.ProductVersion = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetUid(v string) *GetEnvironmentResponseBodyData {
	s.Uid = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetVendorConfig(v string) *GetEnvironmentResponseBodyData {
	s.VendorConfig = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetVendorType(v string) *GetEnvironmentResponseBodyData {
	s.VendorType = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetInstanceStatus(v string) *GetEnvironmentResponseBodyData {
	s.InstanceStatus = &v
	return s
}

type GetEnvironmentResponseBodyDataPlatform struct {
	Architecture *string `json:"architecture,omitempty" xml:"architecture,omitempty"`
	Os           *string `json:"os,omitempty" xml:"os,omitempty"`
}

func (s GetEnvironmentResponseBodyDataPlatform) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentResponseBodyDataPlatform) GoString() string {
	return s.String()
}

func (s *GetEnvironmentResponseBodyDataPlatform) SetArchitecture(v string) *GetEnvironmentResponseBodyDataPlatform {
	s.Architecture = &v
	return s
}

func (s *GetEnvironmentResponseBodyDataPlatform) SetOs(v string) *GetEnvironmentResponseBodyDataPlatform {
	s.Os = &v
	return s
}

type GetEnvironmentResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEnvironmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEnvironmentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *GetEnvironmentResponse) SetHeaders(v map[string]*string) *GetEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *GetEnvironmentResponse) SetBody(v *GetEnvironmentResponseBody) *GetEnvironmentResponse {
	s.Body = v
	return s
}

type GetProductVersionPackageRequest struct {
	PackageUID           *string `json:"packageUID,omitempty" xml:"packageUID,omitempty"`
	Platform             *string `json:"platform,omitempty" xml:"platform,omitempty"`
	PackageType          *string `json:"packageType,omitempty" xml:"packageType,omitempty"`
	PackageContentType   *string `json:"packageContentType,omitempty" xml:"packageContentType,omitempty"`
	OldProductVersionUID *string `json:"oldProductVersionUID,omitempty" xml:"oldProductVersionUID,omitempty"`
	WithURL              *bool   `json:"withURL,omitempty" xml:"withURL,omitempty"`
}

func (s GetProductVersionPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProductVersionPackageRequest) GoString() string {
	return s.String()
}

func (s *GetProductVersionPackageRequest) SetPackageUID(v string) *GetProductVersionPackageRequest {
	s.PackageUID = &v
	return s
}

func (s *GetProductVersionPackageRequest) SetPlatform(v string) *GetProductVersionPackageRequest {
	s.Platform = &v
	return s
}

func (s *GetProductVersionPackageRequest) SetPackageType(v string) *GetProductVersionPackageRequest {
	s.PackageType = &v
	return s
}

func (s *GetProductVersionPackageRequest) SetPackageContentType(v string) *GetProductVersionPackageRequest {
	s.PackageContentType = &v
	return s
}

func (s *GetProductVersionPackageRequest) SetOldProductVersionUID(v string) *GetProductVersionPackageRequest {
	s.OldProductVersionUID = &v
	return s
}

func (s *GetProductVersionPackageRequest) SetWithURL(v bool) *GetProductVersionPackageRequest {
	s.WithURL = &v
	return s
}

type GetProductVersionPackageResponseBody struct {
	Data *GetProductVersionPackageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Code *string                                   `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string                                   `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s GetProductVersionPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProductVersionPackageResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductVersionPackageResponseBody) SetData(v *GetProductVersionPackageResponseBodyData) *GetProductVersionPackageResponseBody {
	s.Data = v
	return s
}

func (s *GetProductVersionPackageResponseBody) SetCode(v string) *GetProductVersionPackageResponseBody {
	s.Code = &v
	return s
}

func (s *GetProductVersionPackageResponseBody) SetMsg(v string) *GetProductVersionPackageResponseBody {
	s.Msg = &v
	return s
}

type GetProductVersionPackageResponseBodyData struct {
	PackageUID         *string   `json:"packageUID,omitempty" xml:"packageUID,omitempty"`
	PackageStatus      *string   `json:"packageStatus,omitempty" xml:"packageStatus,omitempty"`
	PackageURL         *string   `json:"packageURL,omitempty" xml:"packageURL,omitempty"`
	PackageSize        *string   `json:"packageSize,omitempty" xml:"packageSize,omitempty"`
	PackageType        *string   `json:"packageType,omitempty" xml:"packageType,omitempty"`
	PackageContentType *string   `json:"packageContentType,omitempty" xml:"packageContentType,omitempty"`
	Platform           *Platform `json:"platform,omitempty" xml:"platform,omitempty"`
}

func (s GetProductVersionPackageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetProductVersionPackageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProductVersionPackageResponseBodyData) SetPackageUID(v string) *GetProductVersionPackageResponseBodyData {
	s.PackageUID = &v
	return s
}

func (s *GetProductVersionPackageResponseBodyData) SetPackageStatus(v string) *GetProductVersionPackageResponseBodyData {
	s.PackageStatus = &v
	return s
}

func (s *GetProductVersionPackageResponseBodyData) SetPackageURL(v string) *GetProductVersionPackageResponseBodyData {
	s.PackageURL = &v
	return s
}

func (s *GetProductVersionPackageResponseBodyData) SetPackageSize(v string) *GetProductVersionPackageResponseBodyData {
	s.PackageSize = &v
	return s
}

func (s *GetProductVersionPackageResponseBodyData) SetPackageType(v string) *GetProductVersionPackageResponseBodyData {
	s.PackageType = &v
	return s
}

func (s *GetProductVersionPackageResponseBodyData) SetPackageContentType(v string) *GetProductVersionPackageResponseBodyData {
	s.PackageContentType = &v
	return s
}

func (s *GetProductVersionPackageResponseBodyData) SetPlatform(v *Platform) *GetProductVersionPackageResponseBodyData {
	s.Platform = v
	return s
}

type GetProductVersionPackageResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProductVersionPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProductVersionPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProductVersionPackageResponse) GoString() string {
	return s.String()
}

func (s *GetProductVersionPackageResponse) SetHeaders(v map[string]*string) *GetProductVersionPackageResponse {
	s.Headers = v
	return s
}

func (s *GetProductVersionPackageResponse) SetBody(v *GetProductVersionPackageResponseBody) *GetProductVersionPackageResponse {
	s.Body = v
	return s
}

type ListProductDeploymentsRequest struct {
	EnvironmentUID    *string `json:"environmentUID,omitempty" xml:"environmentUID,omitempty"`
	ProductVersionUID *string `json:"productVersionUID,omitempty" xml:"productVersionUID,omitempty"`
	PageNum           *int32  `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize          *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListProductDeploymentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductDeploymentsRequest) GoString() string {
	return s.String()
}

func (s *ListProductDeploymentsRequest) SetEnvironmentUID(v string) *ListProductDeploymentsRequest {
	s.EnvironmentUID = &v
	return s
}

func (s *ListProductDeploymentsRequest) SetProductVersionUID(v string) *ListProductDeploymentsRequest {
	s.ProductVersionUID = &v
	return s
}

func (s *ListProductDeploymentsRequest) SetPageNum(v int32) *ListProductDeploymentsRequest {
	s.PageNum = &v
	return s
}

func (s *ListProductDeploymentsRequest) SetPageSize(v int32) *ListProductDeploymentsRequest {
	s.PageSize = &v
	return s
}

type ListProductDeploymentsResponseBody struct {
	Data *ListProductDeploymentsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Code *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string                                 `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s ListProductDeploymentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductDeploymentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductDeploymentsResponseBody) SetData(v *ListProductDeploymentsResponseBodyData) *ListProductDeploymentsResponseBody {
	s.Data = v
	return s
}

func (s *ListProductDeploymentsResponseBody) SetCode(v string) *ListProductDeploymentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListProductDeploymentsResponseBody) SetMsg(v string) *ListProductDeploymentsResponseBody {
	s.Msg = &v
	return s
}

type ListProductDeploymentsResponseBodyData struct {
	List     []*ListProductDeploymentsResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNum  *int32                                        `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32                                        `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int32                                        `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListProductDeploymentsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListProductDeploymentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProductDeploymentsResponseBodyData) SetList(v []*ListProductDeploymentsResponseBodyDataList) *ListProductDeploymentsResponseBodyData {
	s.List = v
	return s
}

func (s *ListProductDeploymentsResponseBodyData) SetPageNum(v int32) *ListProductDeploymentsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListProductDeploymentsResponseBodyData) SetPageSize(v int32) *ListProductDeploymentsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListProductDeploymentsResponseBodyData) SetTotal(v int32) *ListProductDeploymentsResponseBodyData {
	s.Total = &v
	return s
}

type ListProductDeploymentsResponseBodyDataList struct {
	EnvUID             *string `json:"envUID,omitempty" xml:"envUID,omitempty"`
	Status             *string `json:"status,omitempty" xml:"status,omitempty"`
	Uid                *string `json:"uid,omitempty" xml:"uid,omitempty"`
	ProductVersion     *string `json:"productVersion,omitempty" xml:"productVersion,omitempty"`
	OldProductVersion  *string `json:"oldProductVersion,omitempty" xml:"oldProductVersion,omitempty"`
	PackageInfoUID     *string `json:"packageInfoUID,omitempty" xml:"packageInfoUID,omitempty"`
	EnvParams          *string `json:"envParams,omitempty" xml:"envParams,omitempty"`
	PackageType        *string `json:"packageType,omitempty" xml:"packageType,omitempty"`
	PackageContentType *string `json:"packageContentType,omitempty" xml:"packageContentType,omitempty"`
	ProductName        *string `json:"productName,omitempty" xml:"productName,omitempty"`
}

func (s ListProductDeploymentsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListProductDeploymentsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListProductDeploymentsResponseBodyDataList) SetEnvUID(v string) *ListProductDeploymentsResponseBodyDataList {
	s.EnvUID = &v
	return s
}

func (s *ListProductDeploymentsResponseBodyDataList) SetStatus(v string) *ListProductDeploymentsResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *ListProductDeploymentsResponseBodyDataList) SetUid(v string) *ListProductDeploymentsResponseBodyDataList {
	s.Uid = &v
	return s
}

func (s *ListProductDeploymentsResponseBodyDataList) SetProductVersion(v string) *ListProductDeploymentsResponseBodyDataList {
	s.ProductVersion = &v
	return s
}

func (s *ListProductDeploymentsResponseBodyDataList) SetOldProductVersion(v string) *ListProductDeploymentsResponseBodyDataList {
	s.OldProductVersion = &v
	return s
}

func (s *ListProductDeploymentsResponseBodyDataList) SetPackageInfoUID(v string) *ListProductDeploymentsResponseBodyDataList {
	s.PackageInfoUID = &v
	return s
}

func (s *ListProductDeploymentsResponseBodyDataList) SetEnvParams(v string) *ListProductDeploymentsResponseBodyDataList {
	s.EnvParams = &v
	return s
}

func (s *ListProductDeploymentsResponseBodyDataList) SetPackageType(v string) *ListProductDeploymentsResponseBodyDataList {
	s.PackageType = &v
	return s
}

func (s *ListProductDeploymentsResponseBodyDataList) SetPackageContentType(v string) *ListProductDeploymentsResponseBodyDataList {
	s.PackageContentType = &v
	return s
}

func (s *ListProductDeploymentsResponseBodyDataList) SetProductName(v string) *ListProductDeploymentsResponseBodyDataList {
	s.ProductName = &v
	return s
}

type ListProductDeploymentsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProductDeploymentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProductDeploymentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProductDeploymentsResponse) GoString() string {
	return s.String()
}

func (s *ListProductDeploymentsResponse) SetHeaders(v map[string]*string) *ListProductDeploymentsResponse {
	s.Headers = v
	return s
}

func (s *ListProductDeploymentsResponse) SetBody(v *ListProductDeploymentsResponseBody) *ListProductDeploymentsResponse {
	s.Body = v
	return s
}

type ListEnvironmentNodesRequest struct {
	PageNum  *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListEnvironmentNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentNodesRequest) GoString() string {
	return s.String()
}

func (s *ListEnvironmentNodesRequest) SetPageNum(v int32) *ListEnvironmentNodesRequest {
	s.PageNum = &v
	return s
}

func (s *ListEnvironmentNodesRequest) SetPageSize(v int32) *ListEnvironmentNodesRequest {
	s.PageSize = &v
	return s
}

type ListEnvironmentNodesResponseBody struct {
	Data *ListEnvironmentNodesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Code *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string                               `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s ListEnvironmentNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnvironmentNodesResponseBody) SetData(v *ListEnvironmentNodesResponseBodyData) *ListEnvironmentNodesResponseBody {
	s.Data = v
	return s
}

func (s *ListEnvironmentNodesResponseBody) SetCode(v string) *ListEnvironmentNodesResponseBody {
	s.Code = &v
	return s
}

func (s *ListEnvironmentNodesResponseBody) SetMsg(v string) *ListEnvironmentNodesResponseBody {
	s.Msg = &v
	return s
}

type ListEnvironmentNodesResponseBodyData struct {
	List     []*GetInstanceInfoResponse `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNum  *int64                     `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int64                     `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int64                     `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListEnvironmentNodesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentNodesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnvironmentNodesResponseBodyData) SetList(v []*GetInstanceInfoResponse) *ListEnvironmentNodesResponseBodyData {
	s.List = v
	return s
}

func (s *ListEnvironmentNodesResponseBodyData) SetPageNum(v int64) *ListEnvironmentNodesResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListEnvironmentNodesResponseBodyData) SetPageSize(v int64) *ListEnvironmentNodesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListEnvironmentNodesResponseBodyData) SetTotal(v int64) *ListEnvironmentNodesResponseBodyData {
	s.Total = &v
	return s
}

type ListEnvironmentNodesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEnvironmentNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEnvironmentNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentNodesResponse) GoString() string {
	return s.String()
}

func (s *ListEnvironmentNodesResponse) SetHeaders(v map[string]*string) *ListEnvironmentNodesResponse {
	s.Headers = v
	return s
}

func (s *ListEnvironmentNodesResponse) SetBody(v *ListEnvironmentNodesResponseBody) *ListEnvironmentNodesResponse {
	s.Body = v
	return s
}

type AddProductComponentVersionRequest struct {
	ReleaseName *string `json:"releaseName,omitempty" xml:"releaseName,omitempty"`
}

func (s AddProductComponentVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s AddProductComponentVersionRequest) GoString() string {
	return s.String()
}

func (s *AddProductComponentVersionRequest) SetReleaseName(v string) *AddProductComponentVersionRequest {
	s.ReleaseName = &v
	return s
}

type AddProductComponentVersionResponseBody struct {
	Code *string                                     `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string                                     `json:"msg,omitempty" xml:"msg,omitempty"`
	Data *AddProductComponentVersionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s AddProductComponentVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddProductComponentVersionResponseBody) GoString() string {
	return s.String()
}

func (s *AddProductComponentVersionResponseBody) SetCode(v string) *AddProductComponentVersionResponseBody {
	s.Code = &v
	return s
}

func (s *AddProductComponentVersionResponseBody) SetMsg(v string) *AddProductComponentVersionResponseBody {
	s.Msg = &v
	return s
}

func (s *AddProductComponentVersionResponseBody) SetData(v *AddProductComponentVersionResponseBodyData) *AddProductComponentVersionResponseBody {
	s.Data = v
	return s
}

type AddProductComponentVersionResponseBodyData struct {
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s AddProductComponentVersionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddProductComponentVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddProductComponentVersionResponseBodyData) SetUid(v string) *AddProductComponentVersionResponseBodyData {
	s.Uid = &v
	return s
}

type AddProductComponentVersionResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddProductComponentVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddProductComponentVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s AddProductComponentVersionResponse) GoString() string {
	return s.String()
}

func (s *AddProductComponentVersionResponse) SetHeaders(v map[string]*string) *AddProductComponentVersionResponse {
	s.Headers = v
	return s
}

func (s *AddProductComponentVersionResponse) SetBody(v *AddProductComponentVersionResponseBody) *AddProductComponentVersionResponse {
	s.Body = v
	return s
}

type UpdateEnvironmentNodeRequest struct {
	Labels       map[string]interface{}                `json:"labels,omitempty" xml:"labels,omitempty"`
	RootPassword *string                               `json:"rootPassword,omitempty" xml:"rootPassword,omitempty"`
	Taints       []*UpdateEnvironmentNodeRequestTaints `json:"taints,omitempty" xml:"taints,omitempty" type:"Repeated"`
	// etcd数据盘
	EtcdDisk *string `json:"etcdDisk,omitempty" xml:"etcdDisk,omitempty"`
	// k8s管控数据盘
	TridentSystemDisk *string `json:"tridentSystemDisk,omitempty" xml:"tridentSystemDisk,omitempty"`
	// k8s管控数据盘大小
	TridentSystemSizeDisk *int32 `json:"tridentSystemSizeDisk,omitempty" xml:"tridentSystemSizeDisk,omitempty"`
	// 保留业务分区
	ApplicationDisk *string `json:"applicationDisk,omitempty" xml:"applicationDisk,omitempty"`
}

func (s UpdateEnvironmentNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentNodeRequest) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentNodeRequest) SetLabels(v map[string]interface{}) *UpdateEnvironmentNodeRequest {
	s.Labels = v
	return s
}

func (s *UpdateEnvironmentNodeRequest) SetRootPassword(v string) *UpdateEnvironmentNodeRequest {
	s.RootPassword = &v
	return s
}

func (s *UpdateEnvironmentNodeRequest) SetTaints(v []*UpdateEnvironmentNodeRequestTaints) *UpdateEnvironmentNodeRequest {
	s.Taints = v
	return s
}

func (s *UpdateEnvironmentNodeRequest) SetEtcdDisk(v string) *UpdateEnvironmentNodeRequest {
	s.EtcdDisk = &v
	return s
}

func (s *UpdateEnvironmentNodeRequest) SetTridentSystemDisk(v string) *UpdateEnvironmentNodeRequest {
	s.TridentSystemDisk = &v
	return s
}

func (s *UpdateEnvironmentNodeRequest) SetTridentSystemSizeDisk(v int32) *UpdateEnvironmentNodeRequest {
	s.TridentSystemSizeDisk = &v
	return s
}

func (s *UpdateEnvironmentNodeRequest) SetApplicationDisk(v string) *UpdateEnvironmentNodeRequest {
	s.ApplicationDisk = &v
	return s
}

type UpdateEnvironmentNodeRequestTaints struct {
	Effect *string `json:"effect,omitempty" xml:"effect,omitempty"`
	Key    *string `json:"key,omitempty" xml:"key,omitempty"`
	Value  *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateEnvironmentNodeRequestTaints) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentNodeRequestTaints) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentNodeRequestTaints) SetEffect(v string) *UpdateEnvironmentNodeRequestTaints {
	s.Effect = &v
	return s
}

func (s *UpdateEnvironmentNodeRequestTaints) SetKey(v string) *UpdateEnvironmentNodeRequestTaints {
	s.Key = &v
	return s
}

func (s *UpdateEnvironmentNodeRequestTaints) SetValue(v string) *UpdateEnvironmentNodeRequestTaints {
	s.Value = &v
	return s
}

type UpdateEnvironmentNodeResponseBody struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s UpdateEnvironmentNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentNodeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentNodeResponseBody) SetCode(v string) *UpdateEnvironmentNodeResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEnvironmentNodeResponseBody) SetMsg(v string) *UpdateEnvironmentNodeResponseBody {
	s.Msg = &v
	return s
}

type UpdateEnvironmentNodeResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateEnvironmentNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateEnvironmentNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentNodeResponse) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentNodeResponse) SetHeaders(v map[string]*string) *UpdateEnvironmentNodeResponse {
	s.Headers = v
	return s
}

func (s *UpdateEnvironmentNodeResponse) SetBody(v *UpdateEnvironmentNodeResponseBody) *UpdateEnvironmentNodeResponse {
	s.Body = v
	return s
}

type DeleteProductVersionResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Msg       *string `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s DeleteProductVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProductVersionResponseBody) SetRequestId(v string) *DeleteProductVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProductVersionResponseBody) SetCode(v string) *DeleteProductVersionResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteProductVersionResponseBody) SetMsg(v string) *DeleteProductVersionResponseBody {
	s.Msg = &v
	return s
}

type DeleteProductVersionResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteProductVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProductVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteProductVersionResponse) SetHeaders(v map[string]*string) *DeleteProductVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteProductVersionResponse) SetBody(v *DeleteProductVersionResponseBody) *DeleteProductVersionResponse {
	s.Body = v
	return s
}

type CreateProductHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ClientToken   *string            `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateProductHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateProductHeaders) GoString() string {
	return s.String()
}

func (s *CreateProductHeaders) SetCommonHeaders(v map[string]*string) *CreateProductHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateProductHeaders) SetClientToken(v string) *CreateProductHeaders {
	s.ClientToken = &v
	return s
}

type CreateProductRequest struct {
	Description          *string `json:"description,omitempty" xml:"description,omitempty"`
	FoundationVersionUID *string `json:"foundationVersionUID,omitempty" xml:"foundationVersionUID,omitempty"`
	ProductName          *string `json:"productName,omitempty" xml:"productName,omitempty"`
}

func (s CreateProductRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProductRequest) GoString() string {
	return s.String()
}

func (s *CreateProductRequest) SetDescription(v string) *CreateProductRequest {
	s.Description = &v
	return s
}

func (s *CreateProductRequest) SetFoundationVersionUID(v string) *CreateProductRequest {
	s.FoundationVersionUID = &v
	return s
}

func (s *CreateProductRequest) SetProductName(v string) *CreateProductRequest {
	s.ProductName = &v
	return s
}

type CreateProductResponseBody struct {
	Data *CreateProductResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Code *string                        `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string                        `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s CreateProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProductResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProductResponseBody) SetData(v *CreateProductResponseBodyData) *CreateProductResponseBody {
	s.Data = v
	return s
}

func (s *CreateProductResponseBody) SetCode(v string) *CreateProductResponseBody {
	s.Code = &v
	return s
}

func (s *CreateProductResponseBody) SetMsg(v string) *CreateProductResponseBody {
	s.Msg = &v
	return s
}

type CreateProductResponseBodyData struct {
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s CreateProductResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateProductResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateProductResponseBodyData) SetUid(v string) *CreateProductResponseBodyData {
	s.Uid = &v
	return s
}

type CreateProductResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateProductResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProductResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProductResponse) GoString() string {
	return s.String()
}

func (s *CreateProductResponse) SetHeaders(v map[string]*string) *CreateProductResponse {
	s.Headers = v
	return s
}

func (s *CreateProductResponse) SetBody(v *CreateProductResponseBody) *CreateProductResponse {
	s.Body = v
	return s
}

type ValidateEnvironmentTunnelRequest struct {
	// 通道类型
	TunnelType *string `json:"tunnelType,omitempty" xml:"tunnelType,omitempty"`
	// 通道配置
	TunnelConfig *ValidateEnvironmentTunnelRequestTunnelConfig `json:"tunnelConfig,omitempty" xml:"tunnelConfig,omitempty" type:"Struct"`
}

func (s ValidateEnvironmentTunnelRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateEnvironmentTunnelRequest) GoString() string {
	return s.String()
}

func (s *ValidateEnvironmentTunnelRequest) SetTunnelType(v string) *ValidateEnvironmentTunnelRequest {
	s.TunnelType = &v
	return s
}

func (s *ValidateEnvironmentTunnelRequest) SetTunnelConfig(v *ValidateEnvironmentTunnelRequestTunnelConfig) *ValidateEnvironmentTunnelRequest {
	s.TunnelConfig = v
	return s
}

type ValidateEnvironmentTunnelRequestTunnelConfig struct {
	// 跳板机hostname
	Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty"`
	// 跳板机ssh端口号
	SshPort *int32 `json:"sshPort,omitempty" xml:"sshPort,omitempty"`
	// 跳板机用户名
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
	// 跳板机密码
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// 直连vpcId
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// 直连地域id
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ValidateEnvironmentTunnelRequestTunnelConfig) String() string {
	return tea.Prettify(s)
}

func (s ValidateEnvironmentTunnelRequestTunnelConfig) GoString() string {
	return s.String()
}

func (s *ValidateEnvironmentTunnelRequestTunnelConfig) SetHostname(v string) *ValidateEnvironmentTunnelRequestTunnelConfig {
	s.Hostname = &v
	return s
}

func (s *ValidateEnvironmentTunnelRequestTunnelConfig) SetSshPort(v int32) *ValidateEnvironmentTunnelRequestTunnelConfig {
	s.SshPort = &v
	return s
}

func (s *ValidateEnvironmentTunnelRequestTunnelConfig) SetUsername(v string) *ValidateEnvironmentTunnelRequestTunnelConfig {
	s.Username = &v
	return s
}

func (s *ValidateEnvironmentTunnelRequestTunnelConfig) SetPassword(v string) *ValidateEnvironmentTunnelRequestTunnelConfig {
	s.Password = &v
	return s
}

func (s *ValidateEnvironmentTunnelRequestTunnelConfig) SetVpcId(v string) *ValidateEnvironmentTunnelRequestTunnelConfig {
	s.VpcId = &v
	return s
}

func (s *ValidateEnvironmentTunnelRequestTunnelConfig) SetRegionId(v string) *ValidateEnvironmentTunnelRequestTunnelConfig {
	s.RegionId = &v
	return s
}

type ValidateEnvironmentTunnelResponseBody struct {
	// 错误码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s ValidateEnvironmentTunnelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateEnvironmentTunnelResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateEnvironmentTunnelResponseBody) SetCode(v string) *ValidateEnvironmentTunnelResponseBody {
	s.Code = &v
	return s
}

func (s *ValidateEnvironmentTunnelResponseBody) SetMsg(v string) *ValidateEnvironmentTunnelResponseBody {
	s.Msg = &v
	return s
}

type ValidateEnvironmentTunnelResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ValidateEnvironmentTunnelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateEnvironmentTunnelResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateEnvironmentTunnelResponse) GoString() string {
	return s.String()
}

func (s *ValidateEnvironmentTunnelResponse) SetHeaders(v map[string]*string) *ValidateEnvironmentTunnelResponse {
	s.Headers = v
	return s
}

func (s *ValidateEnvironmentTunnelResponse) SetBody(v *ValidateEnvironmentTunnelResponseBody) *ValidateEnvironmentTunnelResponse {
	s.Body = v
	return s
}

type CreateProductDeploymentRequest struct {
	// 部署包uid
	PackageUID *string `json:"packageUID,omitempty" xml:"packageUID,omitempty"`
	// 环境uid
	EnvironmentUID *string `json:"environmentUID,omitempty" xml:"environmentUID,omitempty"`
	// 产品版本uid
	ProductVersionUID *string `json:"productVersionUID,omitempty" xml:"productVersionUID,omitempty"`
	// 旧产品版本uid
	OldProductVersionUID *string `json:"oldProductVersionUID,omitempty" xml:"oldProductVersionUID,omitempty"`
}

func (s CreateProductDeploymentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProductDeploymentRequest) GoString() string {
	return s.String()
}

func (s *CreateProductDeploymentRequest) SetPackageUID(v string) *CreateProductDeploymentRequest {
	s.PackageUID = &v
	return s
}

func (s *CreateProductDeploymentRequest) SetEnvironmentUID(v string) *CreateProductDeploymentRequest {
	s.EnvironmentUID = &v
	return s
}

func (s *CreateProductDeploymentRequest) SetProductVersionUID(v string) *CreateProductDeploymentRequest {
	s.ProductVersionUID = &v
	return s
}

func (s *CreateProductDeploymentRequest) SetOldProductVersionUID(v string) *CreateProductDeploymentRequest {
	s.OldProductVersionUID = &v
	return s
}

type CreateProductDeploymentResponseBody struct {
	// 错误码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 数据
	Data *CreateProductDeploymentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s CreateProductDeploymentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProductDeploymentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProductDeploymentResponseBody) SetCode(v string) *CreateProductDeploymentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateProductDeploymentResponseBody) SetMsg(v string) *CreateProductDeploymentResponseBody {
	s.Msg = &v
	return s
}

func (s *CreateProductDeploymentResponseBody) SetData(v *CreateProductDeploymentResponseBodyData) *CreateProductDeploymentResponseBody {
	s.Data = v
	return s
}

type CreateProductDeploymentResponseBodyData struct {
	// 部署uid
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s CreateProductDeploymentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateProductDeploymentResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateProductDeploymentResponseBodyData) SetUid(v string) *CreateProductDeploymentResponseBodyData {
	s.Uid = &v
	return s
}

type CreateProductDeploymentResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateProductDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProductDeploymentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProductDeploymentResponse) GoString() string {
	return s.String()
}

func (s *CreateProductDeploymentResponse) SetHeaders(v map[string]*string) *CreateProductDeploymentResponse {
	s.Headers = v
	return s
}

func (s *CreateProductDeploymentResponse) SetBody(v *CreateProductDeploymentResponseBody) *CreateProductDeploymentResponse {
	s.Body = v
	return s
}

type DeleteProductComponentVersionResponseBody struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s DeleteProductComponentVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductComponentVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProductComponentVersionResponseBody) SetCode(v string) *DeleteProductComponentVersionResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteProductComponentVersionResponseBody) SetMsg(v string) *DeleteProductComponentVersionResponseBody {
	s.Msg = &v
	return s
}

type DeleteProductComponentVersionResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteProductComponentVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProductComponentVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductComponentVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteProductComponentVersionResponse) SetHeaders(v map[string]*string) *DeleteProductComponentVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteProductComponentVersionResponse) SetBody(v *DeleteProductComponentVersionResponseBody) *DeleteProductComponentVersionResponse {
	s.Body = v
	return s
}

type DeleteEnvironmentResponseBody struct {
	Msg  *string `json:"msg,omitempty" xml:"msg,omitempty"`
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s DeleteEnvironmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEnvironmentResponseBody) SetMsg(v string) *DeleteEnvironmentResponseBody {
	s.Msg = &v
	return s
}

func (s *DeleteEnvironmentResponseBody) SetCode(v string) *DeleteEnvironmentResponseBody {
	s.Code = &v
	return s
}

type DeleteEnvironmentResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteEnvironmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEnvironmentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteEnvironmentResponse) SetHeaders(v map[string]*string) *DeleteEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteEnvironmentResponse) SetBody(v *DeleteEnvironmentResponseBody) *DeleteEnvironmentResponse {
	s.Body = v
	return s
}

type CreateProductVersionResponseBody struct {
	Data *CreateProductVersionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Code *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string                               `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s CreateProductVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProductVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProductVersionResponseBody) SetData(v *CreateProductVersionResponseBodyData) *CreateProductVersionResponseBody {
	s.Data = v
	return s
}

func (s *CreateProductVersionResponseBody) SetCode(v string) *CreateProductVersionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateProductVersionResponseBody) SetMsg(v string) *CreateProductVersionResponseBody {
	s.Msg = &v
	return s
}

type CreateProductVersionResponseBodyData struct {
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s CreateProductVersionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateProductVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateProductVersionResponseBodyData) SetUid(v string) *CreateProductVersionResponseBodyData {
	s.Uid = &v
	return s
}

type CreateProductVersionResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateProductVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProductVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProductVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateProductVersionResponse) SetHeaders(v map[string]*string) *CreateProductVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateProductVersionResponse) SetBody(v *CreateProductVersionResponseBody) *CreateProductVersionResponse {
	s.Body = v
	return s
}

type PutEnvironmentTunnelRequest struct {
	// 通道类型
	TunnelType *string `json:"tunnelType,omitempty" xml:"tunnelType,omitempty"`
	// 通道配置
	TunnelConfig *PutEnvironmentTunnelRequestTunnelConfig `json:"tunnelConfig,omitempty" xml:"tunnelConfig,omitempty" type:"Struct"`
}

func (s PutEnvironmentTunnelRequest) String() string {
	return tea.Prettify(s)
}

func (s PutEnvironmentTunnelRequest) GoString() string {
	return s.String()
}

func (s *PutEnvironmentTunnelRequest) SetTunnelType(v string) *PutEnvironmentTunnelRequest {
	s.TunnelType = &v
	return s
}

func (s *PutEnvironmentTunnelRequest) SetTunnelConfig(v *PutEnvironmentTunnelRequestTunnelConfig) *PutEnvironmentTunnelRequest {
	s.TunnelConfig = v
	return s
}

type PutEnvironmentTunnelRequestTunnelConfig struct {
	// 跳板机ssh端口号
	SshPort *int32 `json:"sshPort,omitempty" xml:"sshPort,omitempty"`
	// 跳板机hostname
	Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty"`
	// 跳板机用户名
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
	// 跳板机密码
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// 直连vpcId
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// 直连地域id
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s PutEnvironmentTunnelRequestTunnelConfig) String() string {
	return tea.Prettify(s)
}

func (s PutEnvironmentTunnelRequestTunnelConfig) GoString() string {
	return s.String()
}

func (s *PutEnvironmentTunnelRequestTunnelConfig) SetSshPort(v int32) *PutEnvironmentTunnelRequestTunnelConfig {
	s.SshPort = &v
	return s
}

func (s *PutEnvironmentTunnelRequestTunnelConfig) SetHostname(v string) *PutEnvironmentTunnelRequestTunnelConfig {
	s.Hostname = &v
	return s
}

func (s *PutEnvironmentTunnelRequestTunnelConfig) SetUsername(v string) *PutEnvironmentTunnelRequestTunnelConfig {
	s.Username = &v
	return s
}

func (s *PutEnvironmentTunnelRequestTunnelConfig) SetPassword(v string) *PutEnvironmentTunnelRequestTunnelConfig {
	s.Password = &v
	return s
}

func (s *PutEnvironmentTunnelRequestTunnelConfig) SetVpcId(v string) *PutEnvironmentTunnelRequestTunnelConfig {
	s.VpcId = &v
	return s
}

func (s *PutEnvironmentTunnelRequestTunnelConfig) SetRegionId(v string) *PutEnvironmentTunnelRequestTunnelConfig {
	s.RegionId = &v
	return s
}

type PutEnvironmentTunnelResponseBody struct {
	// 错误码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 数据
	Data *PutEnvironmentTunnelResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s PutEnvironmentTunnelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutEnvironmentTunnelResponseBody) GoString() string {
	return s.String()
}

func (s *PutEnvironmentTunnelResponseBody) SetCode(v string) *PutEnvironmentTunnelResponseBody {
	s.Code = &v
	return s
}

func (s *PutEnvironmentTunnelResponseBody) SetMsg(v string) *PutEnvironmentTunnelResponseBody {
	s.Msg = &v
	return s
}

func (s *PutEnvironmentTunnelResponseBody) SetData(v *PutEnvironmentTunnelResponseBodyData) *PutEnvironmentTunnelResponseBody {
	s.Data = v
	return s
}

type PutEnvironmentTunnelResponseBodyData struct {
	// 通道id，可空
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s PutEnvironmentTunnelResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PutEnvironmentTunnelResponseBodyData) GoString() string {
	return s.String()
}

func (s *PutEnvironmentTunnelResponseBodyData) SetUid(v string) *PutEnvironmentTunnelResponseBodyData {
	s.Uid = &v
	return s
}

type PutEnvironmentTunnelResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutEnvironmentTunnelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutEnvironmentTunnelResponse) String() string {
	return tea.Prettify(s)
}

func (s PutEnvironmentTunnelResponse) GoString() string {
	return s.String()
}

func (s *PutEnvironmentTunnelResponse) SetHeaders(v map[string]*string) *PutEnvironmentTunnelResponse {
	s.Headers = v
	return s
}

func (s *PutEnvironmentTunnelResponse) SetBody(v *PutEnvironmentTunnelResponseBody) *PutEnvironmentTunnelResponse {
	s.Body = v
	return s
}

type ListProductEnvironmentsRequest struct {
	EnvType *string `json:"envType,omitempty" xml:"envType,omitempty"`
	// 兼容版本产品版本id
	CompatibleProductVersionUID *string                                    `json:"compatibleProductVersionUID,omitempty" xml:"compatibleProductVersionUID,omitempty"`
	Platforms                   []*ListProductEnvironmentsRequestPlatforms `json:"platforms,omitempty" xml:"platforms,omitempty" type:"Repeated"`
	ProductVersionUID           *string                                    `json:"productVersionUID,omitempty" xml:"productVersionUID,omitempty"`
}

func (s ListProductEnvironmentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductEnvironmentsRequest) GoString() string {
	return s.String()
}

func (s *ListProductEnvironmentsRequest) SetEnvType(v string) *ListProductEnvironmentsRequest {
	s.EnvType = &v
	return s
}

func (s *ListProductEnvironmentsRequest) SetCompatibleProductVersionUID(v string) *ListProductEnvironmentsRequest {
	s.CompatibleProductVersionUID = &v
	return s
}

func (s *ListProductEnvironmentsRequest) SetPlatforms(v []*ListProductEnvironmentsRequestPlatforms) *ListProductEnvironmentsRequest {
	s.Platforms = v
	return s
}

func (s *ListProductEnvironmentsRequest) SetProductVersionUID(v string) *ListProductEnvironmentsRequest {
	s.ProductVersionUID = &v
	return s
}

type ListProductEnvironmentsRequestPlatforms struct {
	Architecture *string `json:"architecture,omitempty" xml:"architecture,omitempty"`
	Os           *string `json:"os,omitempty" xml:"os,omitempty"`
}

func (s ListProductEnvironmentsRequestPlatforms) String() string {
	return tea.Prettify(s)
}

func (s ListProductEnvironmentsRequestPlatforms) GoString() string {
	return s.String()
}

func (s *ListProductEnvironmentsRequestPlatforms) SetArchitecture(v string) *ListProductEnvironmentsRequestPlatforms {
	s.Architecture = &v
	return s
}

func (s *ListProductEnvironmentsRequestPlatforms) SetOs(v string) *ListProductEnvironmentsRequestPlatforms {
	s.Os = &v
	return s
}

type ListProductEnvironmentsShrinkRequest struct {
	EnvType *string `json:"envType,omitempty" xml:"envType,omitempty"`
	// 兼容版本产品版本id
	CompatibleProductVersionUID *string `json:"compatibleProductVersionUID,omitempty" xml:"compatibleProductVersionUID,omitempty"`
	PlatformsShrink             *string `json:"platforms,omitempty" xml:"platforms,omitempty"`
	ProductVersionUID           *string `json:"productVersionUID,omitempty" xml:"productVersionUID,omitempty"`
}

func (s ListProductEnvironmentsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductEnvironmentsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListProductEnvironmentsShrinkRequest) SetEnvType(v string) *ListProductEnvironmentsShrinkRequest {
	s.EnvType = &v
	return s
}

func (s *ListProductEnvironmentsShrinkRequest) SetCompatibleProductVersionUID(v string) *ListProductEnvironmentsShrinkRequest {
	s.CompatibleProductVersionUID = &v
	return s
}

func (s *ListProductEnvironmentsShrinkRequest) SetPlatformsShrink(v string) *ListProductEnvironmentsShrinkRequest {
	s.PlatformsShrink = &v
	return s
}

func (s *ListProductEnvironmentsShrinkRequest) SetProductVersionUID(v string) *ListProductEnvironmentsShrinkRequest {
	s.ProductVersionUID = &v
	return s
}

type ListProductEnvironmentsResponseBody struct {
	Code *string                                    `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string                                    `json:"msg,omitempty" xml:"msg,omitempty"`
	Data []*ListProductEnvironmentsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s ListProductEnvironmentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductEnvironmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductEnvironmentsResponseBody) SetCode(v string) *ListProductEnvironmentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListProductEnvironmentsResponseBody) SetMsg(v string) *ListProductEnvironmentsResponseBody {
	s.Msg = &v
	return s
}

func (s *ListProductEnvironmentsResponseBody) SetData(v []*ListProductEnvironmentsResponseBodyData) *ListProductEnvironmentsResponseBody {
	s.Data = v
	return s
}

type ListProductEnvironmentsResponseBodyData struct {
	Uid                  *string `json:"uid,omitempty" xml:"uid,omitempty"`
	ProductUID           *string `json:"productUID,omitempty" xml:"productUID,omitempty"`
	ProductName          *string `json:"productName,omitempty" xml:"productName,omitempty"`
	ProductVersion       *string `json:"productVersion,omitempty" xml:"productVersion,omitempty"`
	ProductVersionUID    *string `json:"productVersionUID,omitempty" xml:"productVersionUID,omitempty"`
	Provider             *string `json:"provider,omitempty" xml:"provider,omitempty"`
	EnvUID               *string `json:"envUID,omitempty" xml:"envUID,omitempty"`
	EnvType              *string `json:"envType,omitempty" xml:"envType,omitempty"`
	EnvName              *string `json:"envName,omitempty" xml:"envName,omitempty"`
	OldProductVersion    *string `json:"oldProductVersion,omitempty" xml:"oldProductVersion,omitempty"`
	OldProductVersionUID *string `json:"oldProductVersionUID,omitempty" xml:"oldProductVersionUID,omitempty"`
	VendorType           *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
	InstanceStatus       *string `json:"instanceStatus,omitempty" xml:"instanceStatus,omitempty"`
	PlatformStatus       *string `json:"platformStatus,omitempty" xml:"platformStatus,omitempty"`
	VendorConfig         *string `json:"vendorConfig,omitempty" xml:"vendorConfig,omitempty"`
}

func (s ListProductEnvironmentsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListProductEnvironmentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProductEnvironmentsResponseBodyData) SetUid(v string) *ListProductEnvironmentsResponseBodyData {
	s.Uid = &v
	return s
}

func (s *ListProductEnvironmentsResponseBodyData) SetProductUID(v string) *ListProductEnvironmentsResponseBodyData {
	s.ProductUID = &v
	return s
}

func (s *ListProductEnvironmentsResponseBodyData) SetProductName(v string) *ListProductEnvironmentsResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *ListProductEnvironmentsResponseBodyData) SetProductVersion(v string) *ListProductEnvironmentsResponseBodyData {
	s.ProductVersion = &v
	return s
}

func (s *ListProductEnvironmentsResponseBodyData) SetProductVersionUID(v string) *ListProductEnvironmentsResponseBodyData {
	s.ProductVersionUID = &v
	return s
}

func (s *ListProductEnvironmentsResponseBodyData) SetProvider(v string) *ListProductEnvironmentsResponseBodyData {
	s.Provider = &v
	return s
}

func (s *ListProductEnvironmentsResponseBodyData) SetEnvUID(v string) *ListProductEnvironmentsResponseBodyData {
	s.EnvUID = &v
	return s
}

func (s *ListProductEnvironmentsResponseBodyData) SetEnvType(v string) *ListProductEnvironmentsResponseBodyData {
	s.EnvType = &v
	return s
}

func (s *ListProductEnvironmentsResponseBodyData) SetEnvName(v string) *ListProductEnvironmentsResponseBodyData {
	s.EnvName = &v
	return s
}

func (s *ListProductEnvironmentsResponseBodyData) SetOldProductVersion(v string) *ListProductEnvironmentsResponseBodyData {
	s.OldProductVersion = &v
	return s
}

func (s *ListProductEnvironmentsResponseBodyData) SetOldProductVersionUID(v string) *ListProductEnvironmentsResponseBodyData {
	s.OldProductVersionUID = &v
	return s
}

func (s *ListProductEnvironmentsResponseBodyData) SetVendorType(v string) *ListProductEnvironmentsResponseBodyData {
	s.VendorType = &v
	return s
}

func (s *ListProductEnvironmentsResponseBodyData) SetInstanceStatus(v string) *ListProductEnvironmentsResponseBodyData {
	s.InstanceStatus = &v
	return s
}

func (s *ListProductEnvironmentsResponseBodyData) SetPlatformStatus(v string) *ListProductEnvironmentsResponseBodyData {
	s.PlatformStatus = &v
	return s
}

func (s *ListProductEnvironmentsResponseBodyData) SetVendorConfig(v string) *ListProductEnvironmentsResponseBodyData {
	s.VendorConfig = &v
	return s
}

type ListProductEnvironmentsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProductEnvironmentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProductEnvironmentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProductEnvironmentsResponse) GoString() string {
	return s.String()
}

func (s *ListProductEnvironmentsResponse) SetHeaders(v map[string]*string) *ListProductEnvironmentsResponse {
	s.Headers = v
	return s
}

func (s *ListProductEnvironmentsResponse) SetBody(v *ListProductEnvironmentsResponseBody) *ListProductEnvironmentsResponse {
	s.Body = v
	return s
}

type DeleteEnvironmentNodeResponseBody struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s DeleteEnvironmentNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEnvironmentNodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEnvironmentNodeResponseBody) SetCode(v string) *DeleteEnvironmentNodeResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEnvironmentNodeResponseBody) SetMsg(v string) *DeleteEnvironmentNodeResponseBody {
	s.Msg = &v
	return s
}

type DeleteEnvironmentNodeResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteEnvironmentNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEnvironmentNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEnvironmentNodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteEnvironmentNodeResponse) SetHeaders(v map[string]*string) *DeleteEnvironmentNodeResponse {
	s.Headers = v
	return s
}

func (s *DeleteEnvironmentNodeResponse) SetBody(v *DeleteEnvironmentNodeResponseBody) *DeleteEnvironmentNodeResponse {
	s.Body = v
	return s
}

type UpdateProductComponentVersionRequest struct {
	ComponentOrchestrationValues *string `json:"componentOrchestrationValues,omitempty" xml:"componentOrchestrationValues,omitempty"`
	Enable                       *bool   `json:"enable,omitempty" xml:"enable,omitempty"`
	ReleaseName                  *string `json:"releaseName,omitempty" xml:"releaseName,omitempty"`
	// 如果该参数不为空，则代表更新组件版本
	NewComponentVersionUID *string `json:"newComponentVersionUID,omitempty" xml:"newComponentVersionUID,omitempty"`
}

func (s UpdateProductComponentVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductComponentVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateProductComponentVersionRequest) SetComponentOrchestrationValues(v string) *UpdateProductComponentVersionRequest {
	s.ComponentOrchestrationValues = &v
	return s
}

func (s *UpdateProductComponentVersionRequest) SetEnable(v bool) *UpdateProductComponentVersionRequest {
	s.Enable = &v
	return s
}

func (s *UpdateProductComponentVersionRequest) SetReleaseName(v string) *UpdateProductComponentVersionRequest {
	s.ReleaseName = &v
	return s
}

func (s *UpdateProductComponentVersionRequest) SetNewComponentVersionUID(v string) *UpdateProductComponentVersionRequest {
	s.NewComponentVersionUID = &v
	return s
}

type UpdateProductComponentVersionResponseBody struct {
	Msg  *string                                        `json:"msg,omitempty" xml:"msg,omitempty"`
	Code *string                                        `json:"code,omitempty" xml:"code,omitempty"`
	Data *UpdateProductComponentVersionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s UpdateProductComponentVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductComponentVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProductComponentVersionResponseBody) SetMsg(v string) *UpdateProductComponentVersionResponseBody {
	s.Msg = &v
	return s
}

func (s *UpdateProductComponentVersionResponseBody) SetCode(v string) *UpdateProductComponentVersionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateProductComponentVersionResponseBody) SetData(v *UpdateProductComponentVersionResponseBodyData) *UpdateProductComponentVersionResponseBody {
	s.Data = v
	return s
}

type UpdateProductComponentVersionResponseBodyData struct {
	RelationUID *string `json:"relationUID,omitempty" xml:"relationUID,omitempty"`
}

func (s UpdateProductComponentVersionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductComponentVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateProductComponentVersionResponseBodyData) SetRelationUID(v string) *UpdateProductComponentVersionResponseBodyData {
	s.RelationUID = &v
	return s
}

type UpdateProductComponentVersionResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateProductComponentVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProductComponentVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductComponentVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateProductComponentVersionResponse) SetHeaders(v map[string]*string) *UpdateProductComponentVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateProductComponentVersionResponse) SetBody(v *UpdateProductComponentVersionResponseBody) *UpdateProductComponentVersionResponse {
	s.Body = v
	return s
}

type DeleteProductResponseBody struct {
	Msg  *string `json:"msg,omitempty" xml:"msg,omitempty"`
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s DeleteProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProductResponseBody) SetMsg(v string) *DeleteProductResponseBody {
	s.Msg = &v
	return s
}

func (s *DeleteProductResponseBody) SetCode(v string) *DeleteProductResponseBody {
	s.Code = &v
	return s
}

type DeleteProductResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteProductResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProductResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductResponse) GoString() string {
	return s.String()
}

func (s *DeleteProductResponse) SetHeaders(v map[string]*string) *DeleteProductResponse {
	s.Headers = v
	return s
}

func (s *DeleteProductResponse) SetBody(v *DeleteProductResponseBody) *DeleteProductResponse {
	s.Body = v
	return s
}

type AddProductVersionConfigRequest struct {
	// 配置信息key
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 配置信息value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 配置说明
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 组件uid
	ComponentVersionUID *string `json:"componentVersionUID,omitempty" xml:"componentVersionUID,omitempty"`
	// 父组件uid
	ParentComponentVersionUID *string `json:"parentComponentVersionUID,omitempty" xml:"parentComponentVersionUID,omitempty"`
}

func (s AddProductVersionConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s AddProductVersionConfigRequest) GoString() string {
	return s.String()
}

func (s *AddProductVersionConfigRequest) SetName(v string) *AddProductVersionConfigRequest {
	s.Name = &v
	return s
}

func (s *AddProductVersionConfigRequest) SetValue(v string) *AddProductVersionConfigRequest {
	s.Value = &v
	return s
}

func (s *AddProductVersionConfigRequest) SetDescription(v string) *AddProductVersionConfigRequest {
	s.Description = &v
	return s
}

func (s *AddProductVersionConfigRequest) SetComponentVersionUID(v string) *AddProductVersionConfigRequest {
	s.ComponentVersionUID = &v
	return s
}

func (s *AddProductVersionConfigRequest) SetParentComponentVersionUID(v string) *AddProductVersionConfigRequest {
	s.ParentComponentVersionUID = &v
	return s
}

type AddProductVersionConfigResponseBody struct {
	// Id of the request
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Code      *string                                  `json:"code,omitempty" xml:"code,omitempty"`
	Msg       *string                                  `json:"msg,omitempty" xml:"msg,omitempty"`
	Data      *AddProductVersionConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s AddProductVersionConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddProductVersionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *AddProductVersionConfigResponseBody) SetRequestId(v string) *AddProductVersionConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddProductVersionConfigResponseBody) SetCode(v string) *AddProductVersionConfigResponseBody {
	s.Code = &v
	return s
}

func (s *AddProductVersionConfigResponseBody) SetMsg(v string) *AddProductVersionConfigResponseBody {
	s.Msg = &v
	return s
}

func (s *AddProductVersionConfigResponseBody) SetData(v *AddProductVersionConfigResponseBodyData) *AddProductVersionConfigResponseBody {
	s.Data = v
	return s
}

type AddProductVersionConfigResponseBodyData struct {
	// product version config uid
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s AddProductVersionConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddProductVersionConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddProductVersionConfigResponseBodyData) SetUid(v string) *AddProductVersionConfigResponseBodyData {
	s.Uid = &v
	return s
}

type AddProductVersionConfigResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddProductVersionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddProductVersionConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s AddProductVersionConfigResponse) GoString() string {
	return s.String()
}

func (s *AddProductVersionConfigResponse) SetHeaders(v map[string]*string) *AddProductVersionConfigResponse {
	s.Headers = v
	return s
}

func (s *AddProductVersionConfigResponse) SetBody(v *AddProductVersionConfigResponseBody) *AddProductVersionConfigResponse {
	s.Body = v
	return s
}

type InitEnvironmentResourceRequest struct {
	AccessKeyID     *string `json:"accessKeyID,omitempty" xml:"accessKeyID,omitempty"`
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	SecurityToken   *string `json:"securityToken,omitempty" xml:"securityToken,omitempty"`
}

func (s InitEnvironmentResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s InitEnvironmentResourceRequest) GoString() string {
	return s.String()
}

func (s *InitEnvironmentResourceRequest) SetAccessKeyID(v string) *InitEnvironmentResourceRequest {
	s.AccessKeyID = &v
	return s
}

func (s *InitEnvironmentResourceRequest) SetAccessKeySecret(v string) *InitEnvironmentResourceRequest {
	s.AccessKeySecret = &v
	return s
}

func (s *InitEnvironmentResourceRequest) SetSecurityToken(v string) *InitEnvironmentResourceRequest {
	s.SecurityToken = &v
	return s
}

type InitEnvironmentResourceResponseBody struct {
	Data *InitEnvironmentResourceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Code *string                                  `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string                                  `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s InitEnvironmentResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitEnvironmentResourceResponseBody) GoString() string {
	return s.String()
}

func (s *InitEnvironmentResourceResponseBody) SetData(v *InitEnvironmentResourceResponseBodyData) *InitEnvironmentResourceResponseBody {
	s.Data = v
	return s
}

func (s *InitEnvironmentResourceResponseBody) SetCode(v string) *InitEnvironmentResourceResponseBody {
	s.Code = &v
	return s
}

func (s *InitEnvironmentResourceResponseBody) SetMsg(v string) *InitEnvironmentResourceResponseBody {
	s.Msg = &v
	return s
}

type InitEnvironmentResourceResponseBodyData struct {
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s InitEnvironmentResourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s InitEnvironmentResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *InitEnvironmentResourceResponseBodyData) SetStatus(v string) *InitEnvironmentResourceResponseBodyData {
	s.Status = &v
	return s
}

type InitEnvironmentResourceResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InitEnvironmentResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InitEnvironmentResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s InitEnvironmentResourceResponse) GoString() string {
	return s.String()
}

func (s *InitEnvironmentResourceResponse) SetHeaders(v map[string]*string) *InitEnvironmentResourceResponse {
	s.Headers = v
	return s
}

func (s *InitEnvironmentResourceResponse) SetBody(v *InitEnvironmentResourceResponseBody) *InitEnvironmentResourceResponse {
	s.Body = v
	return s
}

type DeleteProductInstanceConfigRequest struct {
	EnvironmentUID    *string `json:"environmentUID,omitempty" xml:"environmentUID,omitempty"`
	ProductVersionUID *string `json:"productVersionUID,omitempty" xml:"productVersionUID,omitempty"`
}

func (s DeleteProductInstanceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductInstanceConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteProductInstanceConfigRequest) SetEnvironmentUID(v string) *DeleteProductInstanceConfigRequest {
	s.EnvironmentUID = &v
	return s
}

func (s *DeleteProductInstanceConfigRequest) SetProductVersionUID(v string) *DeleteProductInstanceConfigRequest {
	s.ProductVersionUID = &v
	return s
}

type DeleteProductInstanceConfigResponseBody struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s DeleteProductInstanceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductInstanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProductInstanceConfigResponseBody) SetCode(v string) *DeleteProductInstanceConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteProductInstanceConfigResponseBody) SetMsg(v string) *DeleteProductInstanceConfigResponseBody {
	s.Msg = &v
	return s
}

type DeleteProductInstanceConfigResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteProductInstanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProductInstanceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductInstanceConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteProductInstanceConfigResponse) SetHeaders(v map[string]*string) *DeleteProductInstanceConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteProductInstanceConfigResponse) SetBody(v *DeleteProductInstanceConfigResponseBody) *DeleteProductInstanceConfigResponse {
	s.Body = v
	return s
}

type CreateEnvironmentHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ClientToken   *string            `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateEnvironmentHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentHeaders) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentHeaders) SetCommonHeaders(v map[string]*string) *CreateEnvironmentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateEnvironmentHeaders) SetClientToken(v string) *CreateEnvironmentHeaders {
	s.ClientToken = &v
	return s
}

type CreateEnvironmentRequest struct {
	Description       *string                           `json:"description,omitempty" xml:"description,omitempty"`
	Location          *string                           `json:"location,omitempty" xml:"location,omitempty"`
	Name              *string                           `json:"name,omitempty" xml:"name,omitempty"`
	Platform          *CreateEnvironmentRequestPlatform `json:"platform,omitempty" xml:"platform,omitempty" type:"Struct"`
	VendorType        *string                           `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
	ProductVersionUID *string                           `json:"productVersionUID,omitempty" xml:"productVersionUID,omitempty"`
	// 环境类型
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateEnvironmentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentRequest) SetDescription(v string) *CreateEnvironmentRequest {
	s.Description = &v
	return s
}

func (s *CreateEnvironmentRequest) SetLocation(v string) *CreateEnvironmentRequest {
	s.Location = &v
	return s
}

func (s *CreateEnvironmentRequest) SetName(v string) *CreateEnvironmentRequest {
	s.Name = &v
	return s
}

func (s *CreateEnvironmentRequest) SetPlatform(v *CreateEnvironmentRequestPlatform) *CreateEnvironmentRequest {
	s.Platform = v
	return s
}

func (s *CreateEnvironmentRequest) SetVendorType(v string) *CreateEnvironmentRequest {
	s.VendorType = &v
	return s
}

func (s *CreateEnvironmentRequest) SetProductVersionUID(v string) *CreateEnvironmentRequest {
	s.ProductVersionUID = &v
	return s
}

func (s *CreateEnvironmentRequest) SetType(v string) *CreateEnvironmentRequest {
	s.Type = &v
	return s
}

type CreateEnvironmentRequestPlatform struct {
	Architecture *string `json:"architecture,omitempty" xml:"architecture,omitempty"`
	Os           *string `json:"os,omitempty" xml:"os,omitempty"`
}

func (s CreateEnvironmentRequestPlatform) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentRequestPlatform) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentRequestPlatform) SetArchitecture(v string) *CreateEnvironmentRequestPlatform {
	s.Architecture = &v
	return s
}

func (s *CreateEnvironmentRequestPlatform) SetOs(v string) *CreateEnvironmentRequestPlatform {
	s.Os = &v
	return s
}

type CreateEnvironmentResponseBody struct {
	Data *CreateEnvironmentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Code *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string                            `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s CreateEnvironmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentResponseBody) SetData(v *CreateEnvironmentResponseBodyData) *CreateEnvironmentResponseBody {
	s.Data = v
	return s
}

func (s *CreateEnvironmentResponseBody) SetCode(v string) *CreateEnvironmentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEnvironmentResponseBody) SetMsg(v string) *CreateEnvironmentResponseBody {
	s.Msg = &v
	return s
}

type CreateEnvironmentResponseBodyData struct {
	EnvironmentUID *string `json:"environmentUID,omitempty" xml:"environmentUID,omitempty"`
	VendorConfig   *string `json:"vendorConfig,omitempty" xml:"vendorConfig,omitempty"`
}

func (s CreateEnvironmentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentResponseBodyData) SetEnvironmentUID(v string) *CreateEnvironmentResponseBodyData {
	s.EnvironmentUID = &v
	return s
}

func (s *CreateEnvironmentResponseBodyData) SetVendorConfig(v string) *CreateEnvironmentResponseBodyData {
	s.VendorConfig = &v
	return s
}

type CreateEnvironmentResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateEnvironmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEnvironmentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentResponse) SetHeaders(v map[string]*string) *CreateEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *CreateEnvironmentResponse) SetBody(v *CreateEnvironmentResponseBody) *CreateEnvironmentResponse {
	s.Body = v
	return s
}

type ListProductVersionConfigsRequest struct {
	PageNum    *string `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize   *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ConfigType *string `json:"configType,omitempty" xml:"configType,omitempty"`
}

func (s ListProductVersionConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductVersionConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListProductVersionConfigsRequest) SetPageNum(v string) *ListProductVersionConfigsRequest {
	s.PageNum = &v
	return s
}

func (s *ListProductVersionConfigsRequest) SetPageSize(v string) *ListProductVersionConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *ListProductVersionConfigsRequest) SetConfigType(v string) *ListProductVersionConfigsRequest {
	s.ConfigType = &v
	return s
}

type ListProductVersionConfigsResponseBody struct {
	Code *string                                    `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string                                    `json:"msg,omitempty" xml:"msg,omitempty"`
	Data *ListProductVersionConfigsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s ListProductVersionConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductVersionConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductVersionConfigsResponseBody) SetCode(v string) *ListProductVersionConfigsResponseBody {
	s.Code = &v
	return s
}

func (s *ListProductVersionConfigsResponseBody) SetMsg(v string) *ListProductVersionConfigsResponseBody {
	s.Msg = &v
	return s
}

func (s *ListProductVersionConfigsResponseBody) SetData(v *ListProductVersionConfigsResponseBodyData) *ListProductVersionConfigsResponseBody {
	s.Data = v
	return s
}

type ListProductVersionConfigsResponseBodyData struct {
	List     []*ListProductVersionConfigsResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNum  *int32                                           `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32                                           `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int32                                           `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListProductVersionConfigsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListProductVersionConfigsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProductVersionConfigsResponseBodyData) SetList(v []*ListProductVersionConfigsResponseBodyDataList) *ListProductVersionConfigsResponseBodyData {
	s.List = v
	return s
}

func (s *ListProductVersionConfigsResponseBodyData) SetPageNum(v int32) *ListProductVersionConfigsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListProductVersionConfigsResponseBodyData) SetPageSize(v int32) *ListProductVersionConfigsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListProductVersionConfigsResponseBodyData) SetTotal(v int32) *ListProductVersionConfigsResponseBodyData {
	s.Total = &v
	return s
}

type ListProductVersionConfigsResponseBodyDataList struct {
	ProductVersionUID          *string `json:"productVersionUID,omitempty" xml:"productVersionUID,omitempty"`
	ComponentVersionUID        *string `json:"componentVersionUID,omitempty" xml:"componentVersionUID,omitempty"`
	Name                       *string `json:"name,omitempty" xml:"name,omitempty"`
	Value                      *string `json:"value,omitempty" xml:"value,omitempty"`
	Description                *string `json:"description,omitempty" xml:"description,omitempty"`
	ParentComponentVersionUID  *string `json:"parentComponentVersionUID,omitempty" xml:"parentComponentVersionUID,omitempty"`
	ComponentName              *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	ParentComponentName        *string `json:"parentComponentName,omitempty" xml:"parentComponentName,omitempty"`
	ComponentReleaseName       *string `json:"componentReleaseName,omitempty" xml:"componentReleaseName,omitempty"`
	ParentComponentReleaseName *string `json:"parentComponentReleaseName,omitempty" xml:"parentComponentReleaseName,omitempty"`
	Uid                        *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s ListProductVersionConfigsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListProductVersionConfigsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListProductVersionConfigsResponseBodyDataList) SetProductVersionUID(v string) *ListProductVersionConfigsResponseBodyDataList {
	s.ProductVersionUID = &v
	return s
}

func (s *ListProductVersionConfigsResponseBodyDataList) SetComponentVersionUID(v string) *ListProductVersionConfigsResponseBodyDataList {
	s.ComponentVersionUID = &v
	return s
}

func (s *ListProductVersionConfigsResponseBodyDataList) SetName(v string) *ListProductVersionConfigsResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListProductVersionConfigsResponseBodyDataList) SetValue(v string) *ListProductVersionConfigsResponseBodyDataList {
	s.Value = &v
	return s
}

func (s *ListProductVersionConfigsResponseBodyDataList) SetDescription(v string) *ListProductVersionConfigsResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListProductVersionConfigsResponseBodyDataList) SetParentComponentVersionUID(v string) *ListProductVersionConfigsResponseBodyDataList {
	s.ParentComponentVersionUID = &v
	return s
}

func (s *ListProductVersionConfigsResponseBodyDataList) SetComponentName(v string) *ListProductVersionConfigsResponseBodyDataList {
	s.ComponentName = &v
	return s
}

func (s *ListProductVersionConfigsResponseBodyDataList) SetParentComponentName(v string) *ListProductVersionConfigsResponseBodyDataList {
	s.ParentComponentName = &v
	return s
}

func (s *ListProductVersionConfigsResponseBodyDataList) SetComponentReleaseName(v string) *ListProductVersionConfigsResponseBodyDataList {
	s.ComponentReleaseName = &v
	return s
}

func (s *ListProductVersionConfigsResponseBodyDataList) SetParentComponentReleaseName(v string) *ListProductVersionConfigsResponseBodyDataList {
	s.ParentComponentReleaseName = &v
	return s
}

func (s *ListProductVersionConfigsResponseBodyDataList) SetUid(v string) *ListProductVersionConfigsResponseBodyDataList {
	s.Uid = &v
	return s
}

type ListProductVersionConfigsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProductVersionConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProductVersionConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProductVersionConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListProductVersionConfigsResponse) SetHeaders(v map[string]*string) *ListProductVersionConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListProductVersionConfigsResponse) SetBody(v *ListProductVersionConfigsResponseBody) *ListProductVersionConfigsResponse {
	s.Body = v
	return s
}

type AddEnvironmentNodesRequest struct {
	Labels map[string]interface{} `json:"labels,omitempty" xml:"labels,omitempty"`
	// master privateId
	MasterPrivateIPs []*string                           `json:"masterPrivateIPs,omitempty" xml:"masterPrivateIPs,omitempty" type:"Repeated"`
	RootPassword     *string                             `json:"rootPassword,omitempty" xml:"rootPassword,omitempty"`
	Taints           []*AddEnvironmentNodesRequestTaints `json:"taints,omitempty" xml:"taints,omitempty" type:"Repeated"`
	// work privateIp
	WorkerPrivateIPs []*string `json:"workerPrivateIPs,omitempty" xml:"workerPrivateIPs,omitempty" type:"Repeated"`
	// etcd数据盘
	EtcdDisk *string `json:"etcdDisk,omitempty" xml:"etcdDisk,omitempty"`
	// k8s管控数据盘
	TridentSystemDisk *string `json:"tridentSystemDisk,omitempty" xml:"tridentSystemDisk,omitempty"`
	// k8s管控数据盘大小
	TridentSystemSizeDisk *int32 `json:"tridentSystemSizeDisk,omitempty" xml:"tridentSystemSizeDisk,omitempty"`
	// 保留业务分区
	ApplicationDisk *string                                 `json:"applicationDisk,omitempty" xml:"applicationDisk,omitempty"`
	Cpu             *int32                                  `json:"cpu,omitempty" xml:"cpu,omitempty"`
	DataDisk        []*AddEnvironmentNodesRequestDataDisk   `json:"dataDisk,omitempty" xml:"dataDisk,omitempty" type:"Repeated"`
	HostName        *string                                 `json:"hostName,omitempty" xml:"hostName,omitempty"`
	Memory          *int32                                  `json:"memory,omitempty" xml:"memory,omitempty"`
	Os              *string                                 `json:"os,omitempty" xml:"os,omitempty"`
	SystemDisk      []*AddEnvironmentNodesRequestSystemDisk `json:"systemDisk,omitempty" xml:"systemDisk,omitempty" type:"Repeated"`
}

func (s AddEnvironmentNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s AddEnvironmentNodesRequest) GoString() string {
	return s.String()
}

func (s *AddEnvironmentNodesRequest) SetLabels(v map[string]interface{}) *AddEnvironmentNodesRequest {
	s.Labels = v
	return s
}

func (s *AddEnvironmentNodesRequest) SetMasterPrivateIPs(v []*string) *AddEnvironmentNodesRequest {
	s.MasterPrivateIPs = v
	return s
}

func (s *AddEnvironmentNodesRequest) SetRootPassword(v string) *AddEnvironmentNodesRequest {
	s.RootPassword = &v
	return s
}

func (s *AddEnvironmentNodesRequest) SetTaints(v []*AddEnvironmentNodesRequestTaints) *AddEnvironmentNodesRequest {
	s.Taints = v
	return s
}

func (s *AddEnvironmentNodesRequest) SetWorkerPrivateIPs(v []*string) *AddEnvironmentNodesRequest {
	s.WorkerPrivateIPs = v
	return s
}

func (s *AddEnvironmentNodesRequest) SetEtcdDisk(v string) *AddEnvironmentNodesRequest {
	s.EtcdDisk = &v
	return s
}

func (s *AddEnvironmentNodesRequest) SetTridentSystemDisk(v string) *AddEnvironmentNodesRequest {
	s.TridentSystemDisk = &v
	return s
}

func (s *AddEnvironmentNodesRequest) SetTridentSystemSizeDisk(v int32) *AddEnvironmentNodesRequest {
	s.TridentSystemSizeDisk = &v
	return s
}

func (s *AddEnvironmentNodesRequest) SetApplicationDisk(v string) *AddEnvironmentNodesRequest {
	s.ApplicationDisk = &v
	return s
}

func (s *AddEnvironmentNodesRequest) SetCpu(v int32) *AddEnvironmentNodesRequest {
	s.Cpu = &v
	return s
}

func (s *AddEnvironmentNodesRequest) SetDataDisk(v []*AddEnvironmentNodesRequestDataDisk) *AddEnvironmentNodesRequest {
	s.DataDisk = v
	return s
}

func (s *AddEnvironmentNodesRequest) SetHostName(v string) *AddEnvironmentNodesRequest {
	s.HostName = &v
	return s
}

func (s *AddEnvironmentNodesRequest) SetMemory(v int32) *AddEnvironmentNodesRequest {
	s.Memory = &v
	return s
}

func (s *AddEnvironmentNodesRequest) SetOs(v string) *AddEnvironmentNodesRequest {
	s.Os = &v
	return s
}

func (s *AddEnvironmentNodesRequest) SetSystemDisk(v []*AddEnvironmentNodesRequestSystemDisk) *AddEnvironmentNodesRequest {
	s.SystemDisk = v
	return s
}

type AddEnvironmentNodesRequestTaints struct {
	Effect *string `json:"effect,omitempty" xml:"effect,omitempty"`
	Key    *string `json:"key,omitempty" xml:"key,omitempty"`
	Value  *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AddEnvironmentNodesRequestTaints) String() string {
	return tea.Prettify(s)
}

func (s AddEnvironmentNodesRequestTaints) GoString() string {
	return s.String()
}

func (s *AddEnvironmentNodesRequestTaints) SetEffect(v string) *AddEnvironmentNodesRequestTaints {
	s.Effect = &v
	return s
}

func (s *AddEnvironmentNodesRequestTaints) SetKey(v string) *AddEnvironmentNodesRequestTaints {
	s.Key = &v
	return s
}

func (s *AddEnvironmentNodesRequestTaints) SetValue(v string) *AddEnvironmentNodesRequestTaints {
	s.Value = &v
	return s
}

type AddEnvironmentNodesRequestDataDisk struct {
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Required *int32  `json:"required,omitempty" xml:"required,omitempty"`
}

func (s AddEnvironmentNodesRequestDataDisk) String() string {
	return tea.Prettify(s)
}

func (s AddEnvironmentNodesRequestDataDisk) GoString() string {
	return s.String()
}

func (s *AddEnvironmentNodesRequestDataDisk) SetName(v string) *AddEnvironmentNodesRequestDataDisk {
	s.Name = &v
	return s
}

func (s *AddEnvironmentNodesRequestDataDisk) SetRequired(v int32) *AddEnvironmentNodesRequestDataDisk {
	s.Required = &v
	return s
}

type AddEnvironmentNodesRequestSystemDisk struct {
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Required *int32  `json:"required,omitempty" xml:"required,omitempty"`
}

func (s AddEnvironmentNodesRequestSystemDisk) String() string {
	return tea.Prettify(s)
}

func (s AddEnvironmentNodesRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *AddEnvironmentNodesRequestSystemDisk) SetName(v string) *AddEnvironmentNodesRequestSystemDisk {
	s.Name = &v
	return s
}

func (s *AddEnvironmentNodesRequestSystemDisk) SetRequired(v int32) *AddEnvironmentNodesRequestSystemDisk {
	s.Required = &v
	return s
}

type AddEnvironmentNodesResponseBody struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s AddEnvironmentNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddEnvironmentNodesResponseBody) GoString() string {
	return s.String()
}

func (s *AddEnvironmentNodesResponseBody) SetCode(v string) *AddEnvironmentNodesResponseBody {
	s.Code = &v
	return s
}

func (s *AddEnvironmentNodesResponseBody) SetMsg(v string) *AddEnvironmentNodesResponseBody {
	s.Msg = &v
	return s
}

type AddEnvironmentNodesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddEnvironmentNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddEnvironmentNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s AddEnvironmentNodesResponse) GoString() string {
	return s.String()
}

func (s *AddEnvironmentNodesResponse) SetHeaders(v map[string]*string) *AddEnvironmentNodesResponse {
	s.Headers = v
	return s
}

func (s *AddEnvironmentNodesResponse) SetBody(v *AddEnvironmentNodesResponseBody) *AddEnvironmentNodesResponse {
	s.Body = v
	return s
}

type CreateProductVersionPackageHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ClientToken   *string            `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateProductVersionPackageHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateProductVersionPackageHeaders) GoString() string {
	return s.String()
}

func (s *CreateProductVersionPackageHeaders) SetCommonHeaders(v map[string]*string) *CreateProductVersionPackageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateProductVersionPackageHeaders) SetClientToken(v string) *CreateProductVersionPackageHeaders {
	s.ClientToken = &v
	return s
}

type CreateProductVersionPackageRequest struct {
	// ENUM:["full","upgrade"]
	PackageType *string `json:"packageType,omitempty" xml:"packageType,omitempty"`
	// ENUM:["all","base"."application"]
	PackageContentType   *string `json:"packageContentType,omitempty" xml:"packageContentType,omitempty"`
	OldProductVersionUID *string `json:"oldProductVersionUID,omitempty" xml:"oldProductVersionUID,omitempty"`
	Platform             *string `json:"platform,omitempty" xml:"platform,omitempty"`
}

func (s CreateProductVersionPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProductVersionPackageRequest) GoString() string {
	return s.String()
}

func (s *CreateProductVersionPackageRequest) SetPackageType(v string) *CreateProductVersionPackageRequest {
	s.PackageType = &v
	return s
}

func (s *CreateProductVersionPackageRequest) SetPackageContentType(v string) *CreateProductVersionPackageRequest {
	s.PackageContentType = &v
	return s
}

func (s *CreateProductVersionPackageRequest) SetOldProductVersionUID(v string) *CreateProductVersionPackageRequest {
	s.OldProductVersionUID = &v
	return s
}

func (s *CreateProductVersionPackageRequest) SetPlatform(v string) *CreateProductVersionPackageRequest {
	s.Platform = &v
	return s
}

type CreateProductVersionPackageResponseBody struct {
	Code *string                                      `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string                                      `json:"msg,omitempty" xml:"msg,omitempty"`
	Data *CreateProductVersionPackageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s CreateProductVersionPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProductVersionPackageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProductVersionPackageResponseBody) SetCode(v string) *CreateProductVersionPackageResponseBody {
	s.Code = &v
	return s
}

func (s *CreateProductVersionPackageResponseBody) SetMsg(v string) *CreateProductVersionPackageResponseBody {
	s.Msg = &v
	return s
}

func (s *CreateProductVersionPackageResponseBody) SetData(v *CreateProductVersionPackageResponseBodyData) *CreateProductVersionPackageResponseBody {
	s.Data = v
	return s
}

type CreateProductVersionPackageResponseBodyData struct {
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s CreateProductVersionPackageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateProductVersionPackageResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateProductVersionPackageResponseBodyData) SetUid(v string) *CreateProductVersionPackageResponseBodyData {
	s.Uid = &v
	return s
}

type CreateProductVersionPackageResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateProductVersionPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProductVersionPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProductVersionPackageResponse) GoString() string {
	return s.String()
}

func (s *CreateProductVersionPackageResponse) SetHeaders(v map[string]*string) *CreateProductVersionPackageResponse {
	s.Headers = v
	return s
}

func (s *CreateProductVersionPackageResponse) SetBody(v *CreateProductVersionPackageResponseBody) *CreateProductVersionPackageResponse {
	s.Body = v
	return s
}

type ListProductInstanceConfigsRequest struct {
	PageNum  *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 组件和全局类型字段
	ParamType         *string `json:"paramType,omitempty" xml:"paramType,omitempty"`
	ProductVersionUID *string `json:"productVersionUID,omitempty" xml:"productVersionUID,omitempty"`
	EnvironmentUID    *string `json:"environmentUID,omitempty" xml:"environmentUID,omitempty"`
}

func (s ListProductInstanceConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductInstanceConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListProductInstanceConfigsRequest) SetPageNum(v int32) *ListProductInstanceConfigsRequest {
	s.PageNum = &v
	return s
}

func (s *ListProductInstanceConfigsRequest) SetPageSize(v int32) *ListProductInstanceConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *ListProductInstanceConfigsRequest) SetParamType(v string) *ListProductInstanceConfigsRequest {
	s.ParamType = &v
	return s
}

func (s *ListProductInstanceConfigsRequest) SetProductVersionUID(v string) *ListProductInstanceConfigsRequest {
	s.ProductVersionUID = &v
	return s
}

func (s *ListProductInstanceConfigsRequest) SetEnvironmentUID(v string) *ListProductInstanceConfigsRequest {
	s.EnvironmentUID = &v
	return s
}

type ListProductInstanceConfigsResponseBody struct {
	Data *ListProductInstanceConfigsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Code *string                                     `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string                                     `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s ListProductInstanceConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductInstanceConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductInstanceConfigsResponseBody) SetData(v *ListProductInstanceConfigsResponseBodyData) *ListProductInstanceConfigsResponseBody {
	s.Data = v
	return s
}

func (s *ListProductInstanceConfigsResponseBody) SetCode(v string) *ListProductInstanceConfigsResponseBody {
	s.Code = &v
	return s
}

func (s *ListProductInstanceConfigsResponseBody) SetMsg(v string) *ListProductInstanceConfigsResponseBody {
	s.Msg = &v
	return s
}

type ListProductInstanceConfigsResponseBodyData struct {
	List     []*ListProductInstanceConfigsResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNum  *int32                                            `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32                                            `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int32                                            `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListProductInstanceConfigsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListProductInstanceConfigsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProductInstanceConfigsResponseBodyData) SetList(v []*ListProductInstanceConfigsResponseBodyDataList) *ListProductInstanceConfigsResponseBodyData {
	s.List = v
	return s
}

func (s *ListProductInstanceConfigsResponseBodyData) SetPageNum(v int32) *ListProductInstanceConfigsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListProductInstanceConfigsResponseBodyData) SetPageSize(v int32) *ListProductInstanceConfigsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListProductInstanceConfigsResponseBodyData) SetTotal(v int32) *ListProductInstanceConfigsResponseBodyData {
	s.Total = &v
	return s
}

type ListProductInstanceConfigsResponseBodyDataList struct {
	CreatedAt                  *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description                *string `json:"description,omitempty" xml:"description,omitempty"`
	Name                       *string `json:"name,omitempty" xml:"name,omitempty"`
	ProductName                *string `json:"productName,omitempty" xml:"productName,omitempty"`
	ProductVersion             *string `json:"productVersion,omitempty" xml:"productVersion,omitempty"`
	ProductVersionUID          *string `json:"productVersionUID,omitempty" xml:"productVersionUID,omitempty"`
	Uid                        *string `json:"uid,omitempty" xml:"uid,omitempty"`
	VendorType                 *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
	Value                      *string `json:"value,omitempty" xml:"value,omitempty"`
	ComponentUID               *string `json:"componentUID,omitempty" xml:"componentUID,omitempty"`
	ComponentVersionUID        *string `json:"componentVersionUID,omitempty" xml:"componentVersionUID,omitempty"`
	ComponentName              *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	ParentComponentName        *string `json:"parentComponentName,omitempty" xml:"parentComponentName,omitempty"`
	ComponentReleaseName       *string `json:"componentReleaseName,omitempty" xml:"componentReleaseName,omitempty"`
	ParentComponentReleaseName *string `json:"parentComponentReleaseName,omitempty" xml:"parentComponentReleaseName,omitempty"`
	ParentComponentVersionUID  *string `json:"parentComponentVersionUID,omitempty" xml:"parentComponentVersionUID,omitempty"`
	EnvUID                     *string `json:"envUID,omitempty" xml:"envUID,omitempty"`
}

func (s ListProductInstanceConfigsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListProductInstanceConfigsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListProductInstanceConfigsResponseBodyDataList) SetCreatedAt(v string) *ListProductInstanceConfigsResponseBodyDataList {
	s.CreatedAt = &v
	return s
}

func (s *ListProductInstanceConfigsResponseBodyDataList) SetDescription(v string) *ListProductInstanceConfigsResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListProductInstanceConfigsResponseBodyDataList) SetName(v string) *ListProductInstanceConfigsResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListProductInstanceConfigsResponseBodyDataList) SetProductName(v string) *ListProductInstanceConfigsResponseBodyDataList {
	s.ProductName = &v
	return s
}

func (s *ListProductInstanceConfigsResponseBodyDataList) SetProductVersion(v string) *ListProductInstanceConfigsResponseBodyDataList {
	s.ProductVersion = &v
	return s
}

func (s *ListProductInstanceConfigsResponseBodyDataList) SetProductVersionUID(v string) *ListProductInstanceConfigsResponseBodyDataList {
	s.ProductVersionUID = &v
	return s
}

func (s *ListProductInstanceConfigsResponseBodyDataList) SetUid(v string) *ListProductInstanceConfigsResponseBodyDataList {
	s.Uid = &v
	return s
}

func (s *ListProductInstanceConfigsResponseBodyDataList) SetVendorType(v string) *ListProductInstanceConfigsResponseBodyDataList {
	s.VendorType = &v
	return s
}

func (s *ListProductInstanceConfigsResponseBodyDataList) SetValue(v string) *ListProductInstanceConfigsResponseBodyDataList {
	s.Value = &v
	return s
}

func (s *ListProductInstanceConfigsResponseBodyDataList) SetComponentUID(v string) *ListProductInstanceConfigsResponseBodyDataList {
	s.ComponentUID = &v
	return s
}

func (s *ListProductInstanceConfigsResponseBodyDataList) SetComponentVersionUID(v string) *ListProductInstanceConfigsResponseBodyDataList {
	s.ComponentVersionUID = &v
	return s
}

func (s *ListProductInstanceConfigsResponseBodyDataList) SetComponentName(v string) *ListProductInstanceConfigsResponseBodyDataList {
	s.ComponentName = &v
	return s
}

func (s *ListProductInstanceConfigsResponseBodyDataList) SetParentComponentName(v string) *ListProductInstanceConfigsResponseBodyDataList {
	s.ParentComponentName = &v
	return s
}

func (s *ListProductInstanceConfigsResponseBodyDataList) SetComponentReleaseName(v string) *ListProductInstanceConfigsResponseBodyDataList {
	s.ComponentReleaseName = &v
	return s
}

func (s *ListProductInstanceConfigsResponseBodyDataList) SetParentComponentReleaseName(v string) *ListProductInstanceConfigsResponseBodyDataList {
	s.ParentComponentReleaseName = &v
	return s
}

func (s *ListProductInstanceConfigsResponseBodyDataList) SetParentComponentVersionUID(v string) *ListProductInstanceConfigsResponseBodyDataList {
	s.ParentComponentVersionUID = &v
	return s
}

func (s *ListProductInstanceConfigsResponseBodyDataList) SetEnvUID(v string) *ListProductInstanceConfigsResponseBodyDataList {
	s.EnvUID = &v
	return s
}

type ListProductInstanceConfigsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProductInstanceConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProductInstanceConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProductInstanceConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListProductInstanceConfigsResponse) SetHeaders(v map[string]*string) *ListProductInstanceConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListProductInstanceConfigsResponse) SetBody(v *ListProductInstanceConfigsResponseBody) *ListProductInstanceConfigsResponse {
	s.Body = v
	return s
}

type UpdateProductVersionRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 更新版本号，只允许从latest 更新到其他版本号，更新之后代表已发布
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UpdateProductVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateProductVersionRequest) SetDescription(v string) *UpdateProductVersionRequest {
	s.Description = &v
	return s
}

func (s *UpdateProductVersionRequest) SetVersion(v string) *UpdateProductVersionRequest {
	s.Version = &v
	return s
}

type UpdateProductVersionResponseBody struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s UpdateProductVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProductVersionResponseBody) SetCode(v string) *UpdateProductVersionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateProductVersionResponseBody) SetMsg(v string) *UpdateProductVersionResponseBody {
	s.Msg = &v
	return s
}

type UpdateProductVersionResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateProductVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProductVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateProductVersionResponse) SetHeaders(v map[string]*string) *UpdateProductVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateProductVersionResponse) SetBody(v *UpdateProductVersionResponseBody) *UpdateProductVersionResponse {
	s.Body = v
	return s
}

type GetProductComponentVersionResponseBody struct {
	Data []*GetProductComponentVersionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Code *string                                       `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string                                       `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s GetProductComponentVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProductComponentVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductComponentVersionResponseBody) SetData(v []*GetProductComponentVersionResponseBodyData) *GetProductComponentVersionResponseBody {
	s.Data = v
	return s
}

func (s *GetProductComponentVersionResponseBody) SetCode(v string) *GetProductComponentVersionResponseBody {
	s.Code = &v
	return s
}

func (s *GetProductComponentVersionResponseBody) SetMsg(v string) *GetProductComponentVersionResponseBody {
	s.Msg = &v
	return s
}

type GetProductComponentVersionResponseBodyData struct {
	RelationUID                       *string `json:"relationUID,omitempty" xml:"relationUID,omitempty"`
	ProductVersionUID                 *string `json:"productVersionUID,omitempty" xml:"productVersionUID,omitempty"`
	ParentComponentVersionUID         *string `json:"parentComponentVersionUID,omitempty" xml:"parentComponentVersionUID,omitempty"`
	ParentComponentVersionRelationUID *string `json:"parentComponentVersionRelationUID,omitempty" xml:"parentComponentVersionRelationUID,omitempty"`
	OrchestrationValues               *string `json:"orchestrationValues,omitempty" xml:"orchestrationValues,omitempty"`
	Values                            *string `json:"values,omitempty" xml:"values,omitempty"`
	ReleaseName                       *string `json:"releaseName,omitempty" xml:"releaseName,omitempty"`
	Sequence                          *int32  `json:"sequence,omitempty" xml:"sequence,omitempty"`
	ComponentVersionUID               *string `json:"componentVersionUID,omitempty" xml:"componentVersionUID,omitempty"`
	ComponentUID                      *string `json:"componentUID,omitempty" xml:"componentUID,omitempty"`
	ComponentName                     *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	ComponentDescription              *string `json:"componentDescription,omitempty" xml:"componentDescription,omitempty"`
	Version                           *string `json:"version,omitempty" xml:"version,omitempty"`
	AppVersion                        *string `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	ComponentVersionDescription       *string `json:"componentVersionDescription,omitempty" xml:"componentVersionDescription,omitempty"`
	ParentComponent                   *bool   `json:"parentComponent,omitempty" xml:"parentComponent,omitempty"`
	Resources                         *string `json:"resources,omitempty" xml:"resources,omitempty"`
	Namespace                         *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	Enable                            *bool   `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s GetProductComponentVersionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetProductComponentVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProductComponentVersionResponseBodyData) SetRelationUID(v string) *GetProductComponentVersionResponseBodyData {
	s.RelationUID = &v
	return s
}

func (s *GetProductComponentVersionResponseBodyData) SetProductVersionUID(v string) *GetProductComponentVersionResponseBodyData {
	s.ProductVersionUID = &v
	return s
}

func (s *GetProductComponentVersionResponseBodyData) SetParentComponentVersionUID(v string) *GetProductComponentVersionResponseBodyData {
	s.ParentComponentVersionUID = &v
	return s
}

func (s *GetProductComponentVersionResponseBodyData) SetParentComponentVersionRelationUID(v string) *GetProductComponentVersionResponseBodyData {
	s.ParentComponentVersionRelationUID = &v
	return s
}

func (s *GetProductComponentVersionResponseBodyData) SetOrchestrationValues(v string) *GetProductComponentVersionResponseBodyData {
	s.OrchestrationValues = &v
	return s
}

func (s *GetProductComponentVersionResponseBodyData) SetValues(v string) *GetProductComponentVersionResponseBodyData {
	s.Values = &v
	return s
}

func (s *GetProductComponentVersionResponseBodyData) SetReleaseName(v string) *GetProductComponentVersionResponseBodyData {
	s.ReleaseName = &v
	return s
}

func (s *GetProductComponentVersionResponseBodyData) SetSequence(v int32) *GetProductComponentVersionResponseBodyData {
	s.Sequence = &v
	return s
}

func (s *GetProductComponentVersionResponseBodyData) SetComponentVersionUID(v string) *GetProductComponentVersionResponseBodyData {
	s.ComponentVersionUID = &v
	return s
}

func (s *GetProductComponentVersionResponseBodyData) SetComponentUID(v string) *GetProductComponentVersionResponseBodyData {
	s.ComponentUID = &v
	return s
}

func (s *GetProductComponentVersionResponseBodyData) SetComponentName(v string) *GetProductComponentVersionResponseBodyData {
	s.ComponentName = &v
	return s
}

func (s *GetProductComponentVersionResponseBodyData) SetComponentDescription(v string) *GetProductComponentVersionResponseBodyData {
	s.ComponentDescription = &v
	return s
}

func (s *GetProductComponentVersionResponseBodyData) SetVersion(v string) *GetProductComponentVersionResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetProductComponentVersionResponseBodyData) SetAppVersion(v string) *GetProductComponentVersionResponseBodyData {
	s.AppVersion = &v
	return s
}

func (s *GetProductComponentVersionResponseBodyData) SetComponentVersionDescription(v string) *GetProductComponentVersionResponseBodyData {
	s.ComponentVersionDescription = &v
	return s
}

func (s *GetProductComponentVersionResponseBodyData) SetParentComponent(v bool) *GetProductComponentVersionResponseBodyData {
	s.ParentComponent = &v
	return s
}

func (s *GetProductComponentVersionResponseBodyData) SetResources(v string) *GetProductComponentVersionResponseBodyData {
	s.Resources = &v
	return s
}

func (s *GetProductComponentVersionResponseBodyData) SetNamespace(v string) *GetProductComponentVersionResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *GetProductComponentVersionResponseBodyData) SetEnable(v bool) *GetProductComponentVersionResponseBodyData {
	s.Enable = &v
	return s
}

type GetProductComponentVersionResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProductComponentVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProductComponentVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProductComponentVersionResponse) GoString() string {
	return s.String()
}

func (s *GetProductComponentVersionResponse) SetHeaders(v map[string]*string) *GetProductComponentVersionResponse {
	s.Headers = v
	return s
}

func (s *GetProductComponentVersionResponse) SetBody(v *GetProductComponentVersionResponseBody) *GetProductComponentVersionResponse {
	s.Body = v
	return s
}

type ListEnvironmentsRequest struct {
	PageNum        *int32  `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	Fuzzy          *string `json:"fuzzy,omitempty" xml:"fuzzy,omitempty"`
	InstanceStatus *string `json:"instanceStatus,omitempty" xml:"instanceStatus,omitempty"`
	VendorType     *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
	ClusterUID     *string `json:"clusterUID,omitempty" xml:"clusterUID,omitempty"`
}

func (s ListEnvironmentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsRequest) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsRequest) SetPageNum(v int32) *ListEnvironmentsRequest {
	s.PageNum = &v
	return s
}

func (s *ListEnvironmentsRequest) SetPageSize(v int32) *ListEnvironmentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEnvironmentsRequest) SetName(v string) *ListEnvironmentsRequest {
	s.Name = &v
	return s
}

func (s *ListEnvironmentsRequest) SetFuzzy(v string) *ListEnvironmentsRequest {
	s.Fuzzy = &v
	return s
}

func (s *ListEnvironmentsRequest) SetInstanceStatus(v string) *ListEnvironmentsRequest {
	s.InstanceStatus = &v
	return s
}

func (s *ListEnvironmentsRequest) SetVendorType(v string) *ListEnvironmentsRequest {
	s.VendorType = &v
	return s
}

func (s *ListEnvironmentsRequest) SetClusterUID(v string) *ListEnvironmentsRequest {
	s.ClusterUID = &v
	return s
}

type ListEnvironmentsResponseBody struct {
	Data *ListEnvironmentsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Code *string                           `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string                           `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s ListEnvironmentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponseBody) SetData(v *ListEnvironmentsResponseBodyData) *ListEnvironmentsResponseBody {
	s.Data = v
	return s
}

func (s *ListEnvironmentsResponseBody) SetCode(v string) *ListEnvironmentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListEnvironmentsResponseBody) SetMsg(v string) *ListEnvironmentsResponseBody {
	s.Msg = &v
	return s
}

type ListEnvironmentsResponseBodyData struct {
	List     []*ListEnvironmentsResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNum  *int32                                  `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32                                  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int32                                  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListEnvironmentsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponseBodyData) SetList(v []*ListEnvironmentsResponseBodyDataList) *ListEnvironmentsResponseBodyData {
	s.List = v
	return s
}

func (s *ListEnvironmentsResponseBodyData) SetPageNum(v int32) *ListEnvironmentsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListEnvironmentsResponseBodyData) SetPageSize(v int32) *ListEnvironmentsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListEnvironmentsResponseBodyData) SetTotal(v int32) *ListEnvironmentsResponseBodyData {
	s.Total = &v
	return s
}

type ListEnvironmentsResponseBodyDataList struct {
	CreatedAt         *string                                       `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description       *string                                       `json:"description,omitempty" xml:"description,omitempty"`
	Id                *int32                                        `json:"id,omitempty" xml:"id,omitempty"`
	Location          *string                                       `json:"location,omitempty" xml:"location,omitempty"`
	Name              *string                                       `json:"name,omitempty" xml:"name,omitempty"`
	Platform          *ListEnvironmentsResponseBodyDataListPlatform `json:"platform,omitempty" xml:"platform,omitempty" type:"Struct"`
	ProductName       *string                                       `json:"productName,omitempty" xml:"productName,omitempty"`
	ProductVersion    *string                                       `json:"productVersion,omitempty" xml:"productVersion,omitempty"`
	ProductVersionUID *string                                       `json:"productVersionUID,omitempty" xml:"productVersionUID,omitempty"`
	Uid               *string                                       `json:"uid,omitempty" xml:"uid,omitempty"`
	VendorType        *string                                       `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s ListEnvironmentsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponseBodyDataList) SetCreatedAt(v string) *ListEnvironmentsResponseBodyDataList {
	s.CreatedAt = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataList) SetDescription(v string) *ListEnvironmentsResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataList) SetId(v int32) *ListEnvironmentsResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataList) SetLocation(v string) *ListEnvironmentsResponseBodyDataList {
	s.Location = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataList) SetName(v string) *ListEnvironmentsResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataList) SetPlatform(v *ListEnvironmentsResponseBodyDataListPlatform) *ListEnvironmentsResponseBodyDataList {
	s.Platform = v
	return s
}

func (s *ListEnvironmentsResponseBodyDataList) SetProductName(v string) *ListEnvironmentsResponseBodyDataList {
	s.ProductName = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataList) SetProductVersion(v string) *ListEnvironmentsResponseBodyDataList {
	s.ProductVersion = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataList) SetProductVersionUID(v string) *ListEnvironmentsResponseBodyDataList {
	s.ProductVersionUID = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataList) SetUid(v string) *ListEnvironmentsResponseBodyDataList {
	s.Uid = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataList) SetVendorType(v string) *ListEnvironmentsResponseBodyDataList {
	s.VendorType = &v
	return s
}

type ListEnvironmentsResponseBodyDataListPlatform struct {
	Architecture *string `json:"architecture,omitempty" xml:"architecture,omitempty"`
	Os           *string `json:"os,omitempty" xml:"os,omitempty"`
}

func (s ListEnvironmentsResponseBodyDataListPlatform) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsResponseBodyDataListPlatform) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponseBodyDataListPlatform) SetArchitecture(v string) *ListEnvironmentsResponseBodyDataListPlatform {
	s.Architecture = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataListPlatform) SetOs(v string) *ListEnvironmentsResponseBodyDataListPlatform {
	s.Os = &v
	return s
}

type ListEnvironmentsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEnvironmentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEnvironmentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsResponse) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponse) SetHeaders(v map[string]*string) *ListEnvironmentsResponse {
	s.Headers = v
	return s
}

func (s *ListEnvironmentsResponse) SetBody(v *ListEnvironmentsResponseBody) *ListEnvironmentsResponse {
	s.Body = v
	return s
}

type UpdateProductRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateProductRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductRequest) GoString() string {
	return s.String()
}

func (s *UpdateProductRequest) SetDescription(v string) *UpdateProductRequest {
	s.Description = &v
	return s
}

type UpdateProductResponseBody struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s UpdateProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProductResponseBody) SetCode(v string) *UpdateProductResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateProductResponseBody) SetMsg(v string) *UpdateProductResponseBody {
	s.Msg = &v
	return s
}

type UpdateProductResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateProductResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProductResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductResponse) GoString() string {
	return s.String()
}

func (s *UpdateProductResponse) SetHeaders(v map[string]*string) *UpdateProductResponse {
	s.Headers = v
	return s
}

func (s *UpdateProductResponse) SetBody(v *UpdateProductResponseBody) *UpdateProductResponse {
	s.Body = v
	return s
}

type DeleteProductVersionConfigResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Msg       *string `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s DeleteProductVersionConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductVersionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProductVersionConfigResponseBody) SetRequestId(v string) *DeleteProductVersionConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProductVersionConfigResponseBody) SetCode(v string) *DeleteProductVersionConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteProductVersionConfigResponseBody) SetMsg(v string) *DeleteProductVersionConfigResponseBody {
	s.Msg = &v
	return s
}

type DeleteProductVersionConfigResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteProductVersionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProductVersionConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductVersionConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteProductVersionConfigResponse) SetHeaders(v map[string]*string) *DeleteProductVersionConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteProductVersionConfigResponse) SetBody(v *DeleteProductVersionConfigResponseBody) *DeleteProductVersionConfigResponse {
	s.Body = v
	return s
}

type ListProductComponentVersionsRequest struct {
	PageNum    *string `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize   *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	SortKey    *string `json:"sortKey,omitempty" xml:"sortKey,omitempty"`
	SortDirect *string `json:"sortDirect,omitempty" xml:"sortDirect,omitempty"`
	// 组件类型，不填写则为全部
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
}

func (s ListProductComponentVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductComponentVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListProductComponentVersionsRequest) SetPageNum(v string) *ListProductComponentVersionsRequest {
	s.PageNum = &v
	return s
}

func (s *ListProductComponentVersionsRequest) SetPageSize(v string) *ListProductComponentVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListProductComponentVersionsRequest) SetSortKey(v string) *ListProductComponentVersionsRequest {
	s.SortKey = &v
	return s
}

func (s *ListProductComponentVersionsRequest) SetSortDirect(v string) *ListProductComponentVersionsRequest {
	s.SortDirect = &v
	return s
}

func (s *ListProductComponentVersionsRequest) SetCategory(v string) *ListProductComponentVersionsRequest {
	s.Category = &v
	return s
}

type ListProductComponentVersionsResponseBody struct {
	Data *ListProductComponentVersionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Code *string                                       `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string                                       `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s ListProductComponentVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductComponentVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductComponentVersionsResponseBody) SetData(v *ListProductComponentVersionsResponseBodyData) *ListProductComponentVersionsResponseBody {
	s.Data = v
	return s
}

func (s *ListProductComponentVersionsResponseBody) SetCode(v string) *ListProductComponentVersionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListProductComponentVersionsResponseBody) SetMsg(v string) *ListProductComponentVersionsResponseBody {
	s.Msg = &v
	return s
}

type ListProductComponentVersionsResponseBodyData struct {
	List []*ProductComponentRelationDetail `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s ListProductComponentVersionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListProductComponentVersionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProductComponentVersionsResponseBodyData) SetList(v []*ProductComponentRelationDetail) *ListProductComponentVersionsResponseBodyData {
	s.List = v
	return s
}

type ListProductComponentVersionsResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProductComponentVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProductComponentVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProductComponentVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListProductComponentVersionsResponse) SetHeaders(v map[string]*string) *ListProductComponentVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListProductComponentVersionsResponse) SetBody(v *ListProductComponentVersionsResponseBody) *ListProductComponentVersionsResponse {
	s.Body = v
	return s
}

type GetEnvironmentNodeResponseBody struct {
	Data *GetEnvironmentNodeResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Code *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string                             `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s GetEnvironmentNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentNodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetEnvironmentNodeResponseBody) SetData(v *GetEnvironmentNodeResponseBodyData) *GetEnvironmentNodeResponseBody {
	s.Data = v
	return s
}

func (s *GetEnvironmentNodeResponseBody) SetCode(v string) *GetEnvironmentNodeResponseBody {
	s.Code = &v
	return s
}

func (s *GetEnvironmentNodeResponseBody) SetMsg(v string) *GetEnvironmentNodeResponseBody {
	s.Msg = &v
	return s
}

type GetEnvironmentNodeResponseBodyData struct {
	Uid        *string `json:"uid,omitempty" xml:"uid,omitempty"`
	CreatedAt  *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	ClusterUID *string `json:"clusterUID,omitempty" xml:"clusterUID,omitempty"`
	// identifier
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// hostName
	HostName *string `json:"hostName,omitempty" xml:"hostName,omitempty"`
	// os
	Os *string `json:"os,omitempty" xml:"os,omitempty"`
	// osVersion
	OsVersion *string `json:"osVersion,omitempty" xml:"osVersion,omitempty"`
	// arch
	Arch *string `json:"arch,omitempty" xml:"arch,omitempty"`
	// kernel
	Kernel *string `json:"kernel,omitempty" xml:"kernel,omitempty"`
	// macAddress
	MacAddress *string `json:"macAddress,omitempty" xml:"macAddress,omitempty"`
	// cpu
	Cpu *string `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// memory
	Memory *string `json:"memory,omitempty" xml:"memory,omitempty"`
	// systemDisk
	SystemDisk []*Disk `json:"systemDisk,omitempty" xml:"systemDisk,omitempty" type:"Repeated"`
	// dataDisk
	DataDisk []*Disk `json:"dataDisk,omitempty" xml:"dataDisk,omitempty" type:"Repeated"`
	// privateIP
	PrivateIP *string `json:"privateIP,omitempty" xml:"privateIP,omitempty"`
	// publicIP
	PublicIP *string `json:"publicIP,omitempty" xml:"publicIP,omitempty"`
	// internetBandwidth
	InternetBandwidth *int32 `json:"internetBandwidth,omitempty" xml:"internetBandwidth,omitempty"`
	// networkCards
	NetworkCards []*GetEnvironmentNodeResponseBodyDataNetworkCards `json:"networkCards,omitempty" xml:"networkCards,omitempty" type:"Repeated"`
	// imageID
	ImageID *string `json:"imageID,omitempty" xml:"imageID,omitempty"`
	// instanceType
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// systemInfo
	SystemInfo *string `json:"systemInfo,omitempty" xml:"systemInfo,omitempty"`
	// rootPassword
	RootPassword *string `json:"rootPassword,omitempty" xml:"rootPassword,omitempty"`
	// labels
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// taints
	Taints []*GetEnvironmentNodeResponseBodyDataTaints `json:"taints,omitempty" xml:"taints,omitempty" type:"Repeated"`
}

func (s GetEnvironmentNodeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentNodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEnvironmentNodeResponseBodyData) SetUid(v string) *GetEnvironmentNodeResponseBodyData {
	s.Uid = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetCreatedAt(v string) *GetEnvironmentNodeResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetClusterUID(v string) *GetEnvironmentNodeResponseBodyData {
	s.ClusterUID = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetIdentifier(v string) *GetEnvironmentNodeResponseBodyData {
	s.Identifier = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetHostName(v string) *GetEnvironmentNodeResponseBodyData {
	s.HostName = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetOs(v string) *GetEnvironmentNodeResponseBodyData {
	s.Os = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetOsVersion(v string) *GetEnvironmentNodeResponseBodyData {
	s.OsVersion = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetArch(v string) *GetEnvironmentNodeResponseBodyData {
	s.Arch = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetKernel(v string) *GetEnvironmentNodeResponseBodyData {
	s.Kernel = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetMacAddress(v string) *GetEnvironmentNodeResponseBodyData {
	s.MacAddress = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetCpu(v string) *GetEnvironmentNodeResponseBodyData {
	s.Cpu = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetMemory(v string) *GetEnvironmentNodeResponseBodyData {
	s.Memory = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetSystemDisk(v []*Disk) *GetEnvironmentNodeResponseBodyData {
	s.SystemDisk = v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetDataDisk(v []*Disk) *GetEnvironmentNodeResponseBodyData {
	s.DataDisk = v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetPrivateIP(v string) *GetEnvironmentNodeResponseBodyData {
	s.PrivateIP = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetPublicIP(v string) *GetEnvironmentNodeResponseBodyData {
	s.PublicIP = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetInternetBandwidth(v int32) *GetEnvironmentNodeResponseBodyData {
	s.InternetBandwidth = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetNetworkCards(v []*GetEnvironmentNodeResponseBodyDataNetworkCards) *GetEnvironmentNodeResponseBodyData {
	s.NetworkCards = v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetImageID(v string) *GetEnvironmentNodeResponseBodyData {
	s.ImageID = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetInstanceType(v string) *GetEnvironmentNodeResponseBodyData {
	s.InstanceType = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetSystemInfo(v string) *GetEnvironmentNodeResponseBodyData {
	s.SystemInfo = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetRootPassword(v string) *GetEnvironmentNodeResponseBodyData {
	s.RootPassword = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetLabels(v map[string]*string) *GetEnvironmentNodeResponseBodyData {
	s.Labels = v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetTaints(v []*GetEnvironmentNodeResponseBodyDataTaints) *GetEnvironmentNodeResponseBodyData {
	s.Taints = v
	return s
}

type GetEnvironmentNodeResponseBodyDataNetworkCards struct {
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// ip
	Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
}

func (s GetEnvironmentNodeResponseBodyDataNetworkCards) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentNodeResponseBodyDataNetworkCards) GoString() string {
	return s.String()
}

func (s *GetEnvironmentNodeResponseBodyDataNetworkCards) SetName(v string) *GetEnvironmentNodeResponseBodyDataNetworkCards {
	s.Name = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyDataNetworkCards) SetIp(v string) *GetEnvironmentNodeResponseBodyDataNetworkCards {
	s.Ip = &v
	return s
}

type GetEnvironmentNodeResponseBodyDataTaints struct {
	// key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// effect
	Effect *string `json:"effect,omitempty" xml:"effect,omitempty"`
}

func (s GetEnvironmentNodeResponseBodyDataTaints) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentNodeResponseBodyDataTaints) GoString() string {
	return s.String()
}

func (s *GetEnvironmentNodeResponseBodyDataTaints) SetKey(v string) *GetEnvironmentNodeResponseBodyDataTaints {
	s.Key = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyDataTaints) SetValue(v string) *GetEnvironmentNodeResponseBodyDataTaints {
	s.Value = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyDataTaints) SetEffect(v string) *GetEnvironmentNodeResponseBodyDataTaints {
	s.Effect = &v
	return s
}

type GetEnvironmentNodeResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEnvironmentNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEnvironmentNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentNodeResponse) GoString() string {
	return s.String()
}

func (s *GetEnvironmentNodeResponse) SetHeaders(v map[string]*string) *GetEnvironmentNodeResponse {
	s.Headers = v
	return s
}

func (s *GetEnvironmentNodeResponse) SetBody(v *GetEnvironmentNodeResponseBody) *GetEnvironmentNodeResponse {
	s.Body = v
	return s
}

type ListProductsRequest struct {
	PageNum  *int32  `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Fuzzy    *string `json:"fuzzy,omitempty" xml:"fuzzy,omitempty"`
}

func (s ListProductsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductsRequest) GoString() string {
	return s.String()
}

func (s *ListProductsRequest) SetPageNum(v int32) *ListProductsRequest {
	s.PageNum = &v
	return s
}

func (s *ListProductsRequest) SetPageSize(v int32) *ListProductsRequest {
	s.PageSize = &v
	return s
}

func (s *ListProductsRequest) SetName(v string) *ListProductsRequest {
	s.Name = &v
	return s
}

func (s *ListProductsRequest) SetFuzzy(v string) *ListProductsRequest {
	s.Fuzzy = &v
	return s
}

type ListProductsResponseBody struct {
	Data *ListProductsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Code *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string                       `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s ListProductsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBody) SetData(v *ListProductsResponseBodyData) *ListProductsResponseBody {
	s.Data = v
	return s
}

func (s *ListProductsResponseBody) SetCode(v string) *ListProductsResponseBody {
	s.Code = &v
	return s
}

func (s *ListProductsResponseBody) SetMsg(v string) *ListProductsResponseBody {
	s.Msg = &v
	return s
}

type ListProductsResponseBodyData struct {
	List     []*ListProductsResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNum  *int32                              `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32                              `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int32                              `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListProductsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyData) SetList(v []*ListProductsResponseBodyDataList) *ListProductsResponseBodyData {
	s.List = v
	return s
}

func (s *ListProductsResponseBodyData) SetPageNum(v int32) *ListProductsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListProductsResponseBodyData) SetPageSize(v int32) *ListProductsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListProductsResponseBodyData) SetTotal(v int32) *ListProductsResponseBodyData {
	s.Total = &v
	return s
}

type ListProductsResponseBodyDataList struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Uid         *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s ListProductsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyDataList) SetDescription(v string) *ListProductsResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListProductsResponseBodyDataList) SetName(v string) *ListProductsResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListProductsResponseBodyDataList) SetUid(v string) *ListProductsResponseBodyDataList {
	s.Uid = &v
	return s
}

type ListProductsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProductsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProductsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponse) GoString() string {
	return s.String()
}

func (s *ListProductsResponse) SetHeaders(v map[string]*string) *ListProductsResponse {
	s.Headers = v
	return s
}

func (s *ListProductsResponse) SetBody(v *ListProductsResponseBody) *ListProductsResponse {
	s.Body = v
	return s
}

type GenerateProductInstanceDeploymentConfigRequest struct {
	EnvironmentUID    *string `json:"environmentUID,omitempty" xml:"environmentUID,omitempty"`
	ProductVersionUID *string `json:"productVersionUID,omitempty" xml:"productVersionUID,omitempty"`
}

func (s GenerateProductInstanceDeploymentConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateProductInstanceDeploymentConfigRequest) GoString() string {
	return s.String()
}

func (s *GenerateProductInstanceDeploymentConfigRequest) SetEnvironmentUID(v string) *GenerateProductInstanceDeploymentConfigRequest {
	s.EnvironmentUID = &v
	return s
}

func (s *GenerateProductInstanceDeploymentConfigRequest) SetProductVersionUID(v string) *GenerateProductInstanceDeploymentConfigRequest {
	s.ProductVersionUID = &v
	return s
}

type GenerateProductInstanceDeploymentConfigResponseBody struct {
	Data *GenerateProductInstanceDeploymentConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Code *string                                                  `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string                                                  `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s GenerateProductInstanceDeploymentConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateProductInstanceDeploymentConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateProductInstanceDeploymentConfigResponseBody) SetData(v *GenerateProductInstanceDeploymentConfigResponseBodyData) *GenerateProductInstanceDeploymentConfigResponseBody {
	s.Data = v
	return s
}

func (s *GenerateProductInstanceDeploymentConfigResponseBody) SetCode(v string) *GenerateProductInstanceDeploymentConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateProductInstanceDeploymentConfigResponseBody) SetMsg(v string) *GenerateProductInstanceDeploymentConfigResponseBody {
	s.Msg = &v
	return s
}

type GenerateProductInstanceDeploymentConfigResponseBodyData struct {
	PackageConfig *string `json:"packageConfig,omitempty" xml:"packageConfig,omitempty"`
}

func (s GenerateProductInstanceDeploymentConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateProductInstanceDeploymentConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateProductInstanceDeploymentConfigResponseBodyData) SetPackageConfig(v string) *GenerateProductInstanceDeploymentConfigResponseBodyData {
	s.PackageConfig = &v
	return s
}

type GenerateProductInstanceDeploymentConfigResponse struct {
	Headers map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GenerateProductInstanceDeploymentConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateProductInstanceDeploymentConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateProductInstanceDeploymentConfigResponse) GoString() string {
	return s.String()
}

func (s *GenerateProductInstanceDeploymentConfigResponse) SetHeaders(v map[string]*string) *GenerateProductInstanceDeploymentConfigResponse {
	s.Headers = v
	return s
}

func (s *GenerateProductInstanceDeploymentConfigResponse) SetBody(v *GenerateProductInstanceDeploymentConfigResponseBody) *GenerateProductInstanceDeploymentConfigResponse {
	s.Body = v
	return s
}

type GetProductDeploymentRequest struct {
	EnvironmentUID    *string `json:"environmentUID,omitempty" xml:"environmentUID,omitempty"`
	ProductVersionUID *string `json:"productVersionUID,omitempty" xml:"productVersionUID,omitempty"`
	WithParamConfig   *bool   `json:"withParamConfig,omitempty" xml:"withParamConfig,omitempty"`
}

func (s GetProductDeploymentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProductDeploymentRequest) GoString() string {
	return s.String()
}

func (s *GetProductDeploymentRequest) SetEnvironmentUID(v string) *GetProductDeploymentRequest {
	s.EnvironmentUID = &v
	return s
}

func (s *GetProductDeploymentRequest) SetProductVersionUID(v string) *GetProductDeploymentRequest {
	s.ProductVersionUID = &v
	return s
}

func (s *GetProductDeploymentRequest) SetWithParamConfig(v bool) *GetProductDeploymentRequest {
	s.WithParamConfig = &v
	return s
}

type GetProductDeploymentResponseBody struct {
	Code *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string                               `json:"msg,omitempty" xml:"msg,omitempty"`
	Data *GetProductDeploymentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s GetProductDeploymentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProductDeploymentResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductDeploymentResponseBody) SetCode(v string) *GetProductDeploymentResponseBody {
	s.Code = &v
	return s
}

func (s *GetProductDeploymentResponseBody) SetMsg(v string) *GetProductDeploymentResponseBody {
	s.Msg = &v
	return s
}

func (s *GetProductDeploymentResponseBody) SetData(v *GetProductDeploymentResponseBodyData) *GetProductDeploymentResponseBody {
	s.Data = v
	return s
}

type GetProductDeploymentResponseBodyData struct {
	EnvUID    *string `json:"envUID,omitempty" xml:"envUID,omitempty"`
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
	Uid       *string `json:"uid,omitempty" xml:"uid,omitempty"`
	EnvParams *string `json:"envParams,omitempty" xml:"envParams,omitempty"`
}

func (s GetProductDeploymentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetProductDeploymentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProductDeploymentResponseBodyData) SetEnvUID(v string) *GetProductDeploymentResponseBodyData {
	s.EnvUID = &v
	return s
}

func (s *GetProductDeploymentResponseBodyData) SetStatus(v string) *GetProductDeploymentResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetProductDeploymentResponseBodyData) SetUid(v string) *GetProductDeploymentResponseBodyData {
	s.Uid = &v
	return s
}

func (s *GetProductDeploymentResponseBodyData) SetEnvParams(v string) *GetProductDeploymentResponseBodyData {
	s.EnvParams = &v
	return s
}

type GetProductDeploymentResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProductDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProductDeploymentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProductDeploymentResponse) GoString() string {
	return s.String()
}

func (s *GetProductDeploymentResponse) SetHeaders(v map[string]*string) *GetProductDeploymentResponse {
	s.Headers = v
	return s
}

func (s *GetProductDeploymentResponse) SetBody(v *GetProductDeploymentResponseBody) *GetProductDeploymentResponse {
	s.Body = v
	return s
}

type ListFoundationReferenceComponentsRequest struct {
	// 产品版本uid
	ProductVersionUID *string `json:"productVersionUID,omitempty" xml:"productVersionUID,omitempty"`
	// foundation version uid
	FoundationVersionUID *string `json:"foundationVersionUID,omitempty" xml:"foundationVersionUID,omitempty"`
	OnlyEnabled          *bool   `json:"onlyEnabled,omitempty" xml:"onlyEnabled,omitempty"`
}

func (s ListFoundationReferenceComponentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFoundationReferenceComponentsRequest) GoString() string {
	return s.String()
}

func (s *ListFoundationReferenceComponentsRequest) SetProductVersionUID(v string) *ListFoundationReferenceComponentsRequest {
	s.ProductVersionUID = &v
	return s
}

func (s *ListFoundationReferenceComponentsRequest) SetFoundationVersionUID(v string) *ListFoundationReferenceComponentsRequest {
	s.FoundationVersionUID = &v
	return s
}

func (s *ListFoundationReferenceComponentsRequest) SetOnlyEnabled(v bool) *ListFoundationReferenceComponentsRequest {
	s.OnlyEnabled = &v
	return s
}

type ListFoundationReferenceComponentsResponseBody struct {
	Data []*ProductComponentRelationDetail `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Code *string                           `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string                           `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s ListFoundationReferenceComponentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFoundationReferenceComponentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFoundationReferenceComponentsResponseBody) SetData(v []*ProductComponentRelationDetail) *ListFoundationReferenceComponentsResponseBody {
	s.Data = v
	return s
}

func (s *ListFoundationReferenceComponentsResponseBody) SetCode(v string) *ListFoundationReferenceComponentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListFoundationReferenceComponentsResponseBody) SetMsg(v string) *ListFoundationReferenceComponentsResponseBody {
	s.Msg = &v
	return s
}

type ListFoundationReferenceComponentsResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFoundationReferenceComponentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFoundationReferenceComponentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFoundationReferenceComponentsResponse) GoString() string {
	return s.String()
}

func (s *ListFoundationReferenceComponentsResponse) SetHeaders(v map[string]*string) *ListFoundationReferenceComponentsResponse {
	s.Headers = v
	return s
}

func (s *ListFoundationReferenceComponentsResponse) SetBody(v *ListFoundationReferenceComponentsResponseBody) *ListFoundationReferenceComponentsResponse {
	s.Body = v
	return s
}

type GetProductVersionDifferencesRequest struct {
	// 上一个产品版本id
	PreVersionUID *string `json:"preVersionUID,omitempty" xml:"preVersionUID,omitempty"`
}

func (s GetProductVersionDifferencesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProductVersionDifferencesRequest) GoString() string {
	return s.String()
}

func (s *GetProductVersionDifferencesRequest) SetPreVersionUID(v string) *GetProductVersionDifferencesRequest {
	s.PreVersionUID = &v
	return s
}

type GetProductVersionDifferencesResponseBody struct {
	// Id of the request
	RequestId *string                                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Data      []*GetProductVersionDifferencesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Code      *string                                         `json:"code,omitempty" xml:"code,omitempty"`
	Msg       *string                                         `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s GetProductVersionDifferencesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProductVersionDifferencesResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductVersionDifferencesResponseBody) SetRequestId(v string) *GetProductVersionDifferencesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProductVersionDifferencesResponseBody) SetData(v []*GetProductVersionDifferencesResponseBodyData) *GetProductVersionDifferencesResponseBody {
	s.Data = v
	return s
}

func (s *GetProductVersionDifferencesResponseBody) SetCode(v string) *GetProductVersionDifferencesResponseBody {
	s.Code = &v
	return s
}

func (s *GetProductVersionDifferencesResponseBody) SetMsg(v string) *GetProductVersionDifferencesResponseBody {
	s.Msg = &v
	return s
}

type GetProductVersionDifferencesResponseBodyData struct {
	// 组件名称
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// 组件实例名称
	ReleaseName *string `json:"releaseName,omitempty" xml:"releaseName,omitempty"`
	// 组件当前的版本号
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// 组件之前的版本号
	PreVersion *string `json:"preVersion,omitempty" xml:"preVersion,omitempty"`
	// 变更类型，ENUM 类型
	Difference  *string `json:"difference,omitempty" xml:"difference,omitempty"`
	UpgradeFlag *bool   `json:"upgradeFlag,omitempty" xml:"upgradeFlag,omitempty"`
	// 变更描述信息
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetProductVersionDifferencesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetProductVersionDifferencesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProductVersionDifferencesResponseBodyData) SetComponentName(v string) *GetProductVersionDifferencesResponseBodyData {
	s.ComponentName = &v
	return s
}

func (s *GetProductVersionDifferencesResponseBodyData) SetReleaseName(v string) *GetProductVersionDifferencesResponseBodyData {
	s.ReleaseName = &v
	return s
}

func (s *GetProductVersionDifferencesResponseBodyData) SetVersion(v string) *GetProductVersionDifferencesResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetProductVersionDifferencesResponseBodyData) SetPreVersion(v string) *GetProductVersionDifferencesResponseBodyData {
	s.PreVersion = &v
	return s
}

func (s *GetProductVersionDifferencesResponseBodyData) SetDifference(v string) *GetProductVersionDifferencesResponseBodyData {
	s.Difference = &v
	return s
}

func (s *GetProductVersionDifferencesResponseBodyData) SetUpgradeFlag(v bool) *GetProductVersionDifferencesResponseBodyData {
	s.UpgradeFlag = &v
	return s
}

func (s *GetProductVersionDifferencesResponseBodyData) SetMessage(v string) *GetProductVersionDifferencesResponseBodyData {
	s.Message = &v
	return s
}

type GetProductVersionDifferencesResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProductVersionDifferencesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProductVersionDifferencesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProductVersionDifferencesResponse) GoString() string {
	return s.String()
}

func (s *GetProductVersionDifferencesResponse) SetHeaders(v map[string]*string) *GetProductVersionDifferencesResponse {
	s.Headers = v
	return s
}

func (s *GetProductVersionDifferencesResponse) SetBody(v *GetProductVersionDifferencesResponseBody) *GetProductVersionDifferencesResponse {
	s.Body = v
	return s
}

type UpdateEnvironmentRequest struct {
	Description  *string `json:"description,omitempty" xml:"description,omitempty"`
	Location     *string `json:"location,omitempty" xml:"location,omitempty"`
	VendorConfig *string `json:"vendorConfig,omitempty" xml:"vendorConfig,omitempty"`
}

func (s UpdateEnvironmentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentRequest) SetDescription(v string) *UpdateEnvironmentRequest {
	s.Description = &v
	return s
}

func (s *UpdateEnvironmentRequest) SetLocation(v string) *UpdateEnvironmentRequest {
	s.Location = &v
	return s
}

func (s *UpdateEnvironmentRequest) SetVendorConfig(v string) *UpdateEnvironmentRequest {
	s.VendorConfig = &v
	return s
}

type UpdateEnvironmentResponseBody struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s UpdateEnvironmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentResponseBody) SetCode(v string) *UpdateEnvironmentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEnvironmentResponseBody) SetMsg(v string) *UpdateEnvironmentResponseBody {
	s.Msg = &v
	return s
}

type UpdateEnvironmentResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateEnvironmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateEnvironmentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentResponse) SetHeaders(v map[string]*string) *UpdateEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *UpdateEnvironmentResponse) SetBody(v *UpdateEnvironmentResponseBody) *UpdateEnvironmentResponse {
	s.Body = v
	return s
}

type ListEnvironmentTunnelsResponseBody struct {
	// 错误码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Msg  *string                                 `json:"msg,omitempty" xml:"msg,omitempty"`
	Data *ListEnvironmentTunnelsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s ListEnvironmentTunnelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentTunnelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnvironmentTunnelsResponseBody) SetCode(v string) *ListEnvironmentTunnelsResponseBody {
	s.Code = &v
	return s
}

func (s *ListEnvironmentTunnelsResponseBody) SetMsg(v string) *ListEnvironmentTunnelsResponseBody {
	s.Msg = &v
	return s
}

func (s *ListEnvironmentTunnelsResponseBody) SetData(v *ListEnvironmentTunnelsResponseBodyData) *ListEnvironmentTunnelsResponseBody {
	s.Data = v
	return s
}

type ListEnvironmentTunnelsResponseBodyData struct {
	List []*ListEnvironmentTunnelsResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s ListEnvironmentTunnelsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentTunnelsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnvironmentTunnelsResponseBodyData) SetList(v []*ListEnvironmentTunnelsResponseBodyDataList) *ListEnvironmentTunnelsResponseBodyData {
	s.List = v
	return s
}

type ListEnvironmentTunnelsResponseBodyDataList struct {
	// 通道类型
	TunnelType   *string                                                 `json:"tunnelType,omitempty" xml:"tunnelType,omitempty"`
	TunnelConfig *ListEnvironmentTunnelsResponseBodyDataListTunnelConfig `json:"tunnelConfig,omitempty" xml:"tunnelConfig,omitempty" type:"Struct"`
}

func (s ListEnvironmentTunnelsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentTunnelsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListEnvironmentTunnelsResponseBodyDataList) SetTunnelType(v string) *ListEnvironmentTunnelsResponseBodyDataList {
	s.TunnelType = &v
	return s
}

func (s *ListEnvironmentTunnelsResponseBodyDataList) SetTunnelConfig(v *ListEnvironmentTunnelsResponseBodyDataListTunnelConfig) *ListEnvironmentTunnelsResponseBodyDataList {
	s.TunnelConfig = v
	return s
}

type ListEnvironmentTunnelsResponseBodyDataListTunnelConfig struct {
	// 跳板机hostname
	Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty"`
	// 跳板机ssh端口号
	SshPort *int32 `json:"sshPort,omitempty" xml:"sshPort,omitempty"`
	// 跳板机用户名
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
	// 跳板机密码
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// 直连vpcId
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// 直连regionId
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListEnvironmentTunnelsResponseBodyDataListTunnelConfig) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentTunnelsResponseBodyDataListTunnelConfig) GoString() string {
	return s.String()
}

func (s *ListEnvironmentTunnelsResponseBodyDataListTunnelConfig) SetHostname(v string) *ListEnvironmentTunnelsResponseBodyDataListTunnelConfig {
	s.Hostname = &v
	return s
}

func (s *ListEnvironmentTunnelsResponseBodyDataListTunnelConfig) SetSshPort(v int32) *ListEnvironmentTunnelsResponseBodyDataListTunnelConfig {
	s.SshPort = &v
	return s
}

func (s *ListEnvironmentTunnelsResponseBodyDataListTunnelConfig) SetUsername(v string) *ListEnvironmentTunnelsResponseBodyDataListTunnelConfig {
	s.Username = &v
	return s
}

func (s *ListEnvironmentTunnelsResponseBodyDataListTunnelConfig) SetPassword(v string) *ListEnvironmentTunnelsResponseBodyDataListTunnelConfig {
	s.Password = &v
	return s
}

func (s *ListEnvironmentTunnelsResponseBodyDataListTunnelConfig) SetVpcId(v string) *ListEnvironmentTunnelsResponseBodyDataListTunnelConfig {
	s.VpcId = &v
	return s
}

func (s *ListEnvironmentTunnelsResponseBodyDataListTunnelConfig) SetRegionId(v string) *ListEnvironmentTunnelsResponseBodyDataListTunnelConfig {
	s.RegionId = &v
	return s
}

type ListEnvironmentTunnelsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEnvironmentTunnelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEnvironmentTunnelsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentTunnelsResponse) GoString() string {
	return s.String()
}

func (s *ListEnvironmentTunnelsResponse) SetHeaders(v map[string]*string) *ListEnvironmentTunnelsResponse {
	s.Headers = v
	return s
}

func (s *ListEnvironmentTunnelsResponse) SetBody(v *ListEnvironmentTunnelsResponseBody) *ListEnvironmentTunnelsResponse {
	s.Body = v
	return s
}

type ListProductVersionsRequest struct {
	// 过滤参数，产品uid
	ProductUID *string `json:"productUID,omitempty" xml:"productUID,omitempty"`
	// 过滤参数，产品名称
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// 过滤参数，产品版本号
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// 过滤参数，是否已发布
	Released *bool `json:"released,omitempty" xml:"released,omitempty"`
	// 过滤参数，支持的环境类型
	SupportedFoundationTypes []*string `json:"supportedFoundationTypes,omitempty" xml:"supportedFoundationTypes,omitempty" type:"Repeated"`
	// 支持的platform，为空则不过滤
	Platforms []*ListProductVersionsRequestPlatforms `json:"platforms,omitempty" xml:"platforms,omitempty" type:"Repeated"`
	PageNum   *string                                `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize  *string                                `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListProductVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListProductVersionsRequest) SetProductUID(v string) *ListProductVersionsRequest {
	s.ProductUID = &v
	return s
}

func (s *ListProductVersionsRequest) SetProductName(v string) *ListProductVersionsRequest {
	s.ProductName = &v
	return s
}

func (s *ListProductVersionsRequest) SetVersion(v string) *ListProductVersionsRequest {
	s.Version = &v
	return s
}

func (s *ListProductVersionsRequest) SetReleased(v bool) *ListProductVersionsRequest {
	s.Released = &v
	return s
}

func (s *ListProductVersionsRequest) SetSupportedFoundationTypes(v []*string) *ListProductVersionsRequest {
	s.SupportedFoundationTypes = v
	return s
}

func (s *ListProductVersionsRequest) SetPlatforms(v []*ListProductVersionsRequestPlatforms) *ListProductVersionsRequest {
	s.Platforms = v
	return s
}

func (s *ListProductVersionsRequest) SetPageNum(v string) *ListProductVersionsRequest {
	s.PageNum = &v
	return s
}

func (s *ListProductVersionsRequest) SetPageSize(v string) *ListProductVersionsRequest {
	s.PageSize = &v
	return s
}

type ListProductVersionsRequestPlatforms struct {
	Architecture *string `json:"architecture,omitempty" xml:"architecture,omitempty"`
	Os           *string `json:"os,omitempty" xml:"os,omitempty"`
}

func (s ListProductVersionsRequestPlatforms) String() string {
	return tea.Prettify(s)
}

func (s ListProductVersionsRequestPlatforms) GoString() string {
	return s.String()
}

func (s *ListProductVersionsRequestPlatforms) SetArchitecture(v string) *ListProductVersionsRequestPlatforms {
	s.Architecture = &v
	return s
}

func (s *ListProductVersionsRequestPlatforms) SetOs(v string) *ListProductVersionsRequestPlatforms {
	s.Os = &v
	return s
}

type ListProductVersionsShrinkRequest struct {
	// 过滤参数，产品uid
	ProductUID *string `json:"productUID,omitempty" xml:"productUID,omitempty"`
	// 过滤参数，产品名称
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// 过滤参数，产品版本号
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// 过滤参数，是否已发布
	Released *bool `json:"released,omitempty" xml:"released,omitempty"`
	// 过滤参数，支持的环境类型
	SupportedFoundationTypesShrink *string `json:"supportedFoundationTypes,omitempty" xml:"supportedFoundationTypes,omitempty"`
	// 支持的platform，为空则不过滤
	PlatformsShrink *string `json:"platforms,omitempty" xml:"platforms,omitempty"`
	PageNum         *string `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize        *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListProductVersionsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductVersionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListProductVersionsShrinkRequest) SetProductUID(v string) *ListProductVersionsShrinkRequest {
	s.ProductUID = &v
	return s
}

func (s *ListProductVersionsShrinkRequest) SetProductName(v string) *ListProductVersionsShrinkRequest {
	s.ProductName = &v
	return s
}

func (s *ListProductVersionsShrinkRequest) SetVersion(v string) *ListProductVersionsShrinkRequest {
	s.Version = &v
	return s
}

func (s *ListProductVersionsShrinkRequest) SetReleased(v bool) *ListProductVersionsShrinkRequest {
	s.Released = &v
	return s
}

func (s *ListProductVersionsShrinkRequest) SetSupportedFoundationTypesShrink(v string) *ListProductVersionsShrinkRequest {
	s.SupportedFoundationTypesShrink = &v
	return s
}

func (s *ListProductVersionsShrinkRequest) SetPlatformsShrink(v string) *ListProductVersionsShrinkRequest {
	s.PlatformsShrink = &v
	return s
}

func (s *ListProductVersionsShrinkRequest) SetPageNum(v string) *ListProductVersionsShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *ListProductVersionsShrinkRequest) SetPageSize(v string) *ListProductVersionsShrinkRequest {
	s.PageSize = &v
	return s
}

type ListProductVersionsResponseBody struct {
	Data *ListProductVersionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Code *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string                              `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s ListProductVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductVersionsResponseBody) SetData(v *ListProductVersionsResponseBodyData) *ListProductVersionsResponseBody {
	s.Data = v
	return s
}

func (s *ListProductVersionsResponseBody) SetCode(v string) *ListProductVersionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListProductVersionsResponseBody) SetMsg(v string) *ListProductVersionsResponseBody {
	s.Msg = &v
	return s
}

type ListProductVersionsResponseBodyData struct {
	List     []*ListProductVersionsResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNum  *int32                                     `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32                                     `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int32                                     `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListProductVersionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListProductVersionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProductVersionsResponseBodyData) SetList(v []*ListProductVersionsResponseBodyDataList) *ListProductVersionsResponseBodyData {
	s.List = v
	return s
}

func (s *ListProductVersionsResponseBodyData) SetPageNum(v int32) *ListProductVersionsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListProductVersionsResponseBodyData) SetPageSize(v int32) *ListProductVersionsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListProductVersionsResponseBodyData) SetTotal(v int32) *ListProductVersionsResponseBodyData {
	s.Total = &v
	return s
}

type ListProductVersionsResponseBodyDataList struct {
	Annotations *ListProductVersionsResponseBodyDataListAnnotations `json:"annotations,omitempty" xml:"annotations,omitempty" type:"Struct"`
	Description *string                                             `json:"description,omitempty" xml:"description,omitempty"`
	PackageURL  *string                                             `json:"packageURL,omitempty" xml:"packageURL,omitempty"`
	ProductName *string                                             `json:"productName,omitempty" xml:"productName,omitempty"`
	ProductUID  *string                                             `json:"productUID,omitempty" xml:"productUID,omitempty"`
	Provider    *string                                             `json:"provider,omitempty" xml:"provider,omitempty"`
	Uid         *string                                             `json:"uid,omitempty" xml:"uid,omitempty"`
	Version     *string                                             `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListProductVersionsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListProductVersionsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListProductVersionsResponseBodyDataList) SetAnnotations(v *ListProductVersionsResponseBodyDataListAnnotations) *ListProductVersionsResponseBodyDataList {
	s.Annotations = v
	return s
}

func (s *ListProductVersionsResponseBodyDataList) SetDescription(v string) *ListProductVersionsResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListProductVersionsResponseBodyDataList) SetPackageURL(v string) *ListProductVersionsResponseBodyDataList {
	s.PackageURL = &v
	return s
}

func (s *ListProductVersionsResponseBodyDataList) SetProductName(v string) *ListProductVersionsResponseBodyDataList {
	s.ProductName = &v
	return s
}

func (s *ListProductVersionsResponseBodyDataList) SetProductUID(v string) *ListProductVersionsResponseBodyDataList {
	s.ProductUID = &v
	return s
}

func (s *ListProductVersionsResponseBodyDataList) SetProvider(v string) *ListProductVersionsResponseBodyDataList {
	s.Provider = &v
	return s
}

func (s *ListProductVersionsResponseBodyDataList) SetUid(v string) *ListProductVersionsResponseBodyDataList {
	s.Uid = &v
	return s
}

func (s *ListProductVersionsResponseBodyDataList) SetVersion(v string) *ListProductVersionsResponseBodyDataList {
	s.Version = &v
	return s
}

type ListProductVersionsResponseBodyDataListAnnotations struct {
	AdditionalProp1 *string `json:"additionalProp1,omitempty" xml:"additionalProp1,omitempty"`
	AdditionalProp2 *string `json:"additionalProp2,omitempty" xml:"additionalProp2,omitempty"`
	AdditionalProp3 *string `json:"additionalProp3,omitempty" xml:"additionalProp3,omitempty"`
}

func (s ListProductVersionsResponseBodyDataListAnnotations) String() string {
	return tea.Prettify(s)
}

func (s ListProductVersionsResponseBodyDataListAnnotations) GoString() string {
	return s.String()
}

func (s *ListProductVersionsResponseBodyDataListAnnotations) SetAdditionalProp1(v string) *ListProductVersionsResponseBodyDataListAnnotations {
	s.AdditionalProp1 = &v
	return s
}

func (s *ListProductVersionsResponseBodyDataListAnnotations) SetAdditionalProp2(v string) *ListProductVersionsResponseBodyDataListAnnotations {
	s.AdditionalProp2 = &v
	return s
}

func (s *ListProductVersionsResponseBodyDataListAnnotations) SetAdditionalProp3(v string) *ListProductVersionsResponseBodyDataListAnnotations {
	s.AdditionalProp3 = &v
	return s
}

type ListProductVersionsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProductVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProductVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProductVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListProductVersionsResponse) SetHeaders(v map[string]*string) *ListProductVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListProductVersionsResponse) SetBody(v *ListProductVersionsResponseBody) *ListProductVersionsResponse {
	s.Body = v
	return s
}

type PutProductInstanceConfigRequest struct {
	EnvironmentUID            *string   `json:"environmentUID,omitempty" xml:"environmentUID,omitempty"`
	ProductVersionUID         *string   `json:"productVersionUID,omitempty" xml:"productVersionUID,omitempty"`
	ComponentUID              *string   `json:"componentUID,omitempty" xml:"componentUID,omitempty"`
	ComponentVersionUID       *string   `json:"componentVersionUID,omitempty" xml:"componentVersionUID,omitempty"`
	Name                      *string   `json:"name,omitempty" xml:"name,omitempty"`
	ConfigUID                 *string   `json:"configUID,omitempty" xml:"configUID,omitempty"`
	ReleaseName               *string   `json:"releaseName,omitempty" xml:"releaseName,omitempty"`
	Scope                     []*string `json:"scope,omitempty" xml:"scope,omitempty" type:"Repeated"`
	Value                     *string   `json:"value,omitempty" xml:"value,omitempty"`
	Description               *string   `json:"description,omitempty" xml:"description,omitempty"`
	ParentComponentName       *string   `json:"parentComponentName,omitempty" xml:"parentComponentName,omitempty"`
	ParentComponentVersionUID *string   `json:"parentComponentVersionUID,omitempty" xml:"parentComponentVersionUID,omitempty"`
}

func (s PutProductInstanceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s PutProductInstanceConfigRequest) GoString() string {
	return s.String()
}

func (s *PutProductInstanceConfigRequest) SetEnvironmentUID(v string) *PutProductInstanceConfigRequest {
	s.EnvironmentUID = &v
	return s
}

func (s *PutProductInstanceConfigRequest) SetProductVersionUID(v string) *PutProductInstanceConfigRequest {
	s.ProductVersionUID = &v
	return s
}

func (s *PutProductInstanceConfigRequest) SetComponentUID(v string) *PutProductInstanceConfigRequest {
	s.ComponentUID = &v
	return s
}

func (s *PutProductInstanceConfigRequest) SetComponentVersionUID(v string) *PutProductInstanceConfigRequest {
	s.ComponentVersionUID = &v
	return s
}

func (s *PutProductInstanceConfigRequest) SetName(v string) *PutProductInstanceConfigRequest {
	s.Name = &v
	return s
}

func (s *PutProductInstanceConfigRequest) SetConfigUID(v string) *PutProductInstanceConfigRequest {
	s.ConfigUID = &v
	return s
}

func (s *PutProductInstanceConfigRequest) SetReleaseName(v string) *PutProductInstanceConfigRequest {
	s.ReleaseName = &v
	return s
}

func (s *PutProductInstanceConfigRequest) SetScope(v []*string) *PutProductInstanceConfigRequest {
	s.Scope = v
	return s
}

func (s *PutProductInstanceConfigRequest) SetValue(v string) *PutProductInstanceConfigRequest {
	s.Value = &v
	return s
}

func (s *PutProductInstanceConfigRequest) SetDescription(v string) *PutProductInstanceConfigRequest {
	s.Description = &v
	return s
}

func (s *PutProductInstanceConfigRequest) SetParentComponentName(v string) *PutProductInstanceConfigRequest {
	s.ParentComponentName = &v
	return s
}

func (s *PutProductInstanceConfigRequest) SetParentComponentVersionUID(v string) *PutProductInstanceConfigRequest {
	s.ParentComponentVersionUID = &v
	return s
}

type PutProductInstanceConfigResponseBody struct {
	Code *string                                   `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string                                   `json:"msg,omitempty" xml:"msg,omitempty"`
	Data *PutProductInstanceConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s PutProductInstanceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutProductInstanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *PutProductInstanceConfigResponseBody) SetCode(v string) *PutProductInstanceConfigResponseBody {
	s.Code = &v
	return s
}

func (s *PutProductInstanceConfigResponseBody) SetMsg(v string) *PutProductInstanceConfigResponseBody {
	s.Msg = &v
	return s
}

func (s *PutProductInstanceConfigResponseBody) SetData(v *PutProductInstanceConfigResponseBodyData) *PutProductInstanceConfigResponseBody {
	s.Data = v
	return s
}

type PutProductInstanceConfigResponseBodyData struct {
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s PutProductInstanceConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PutProductInstanceConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *PutProductInstanceConfigResponseBodyData) SetUid(v string) *PutProductInstanceConfigResponseBodyData {
	s.Uid = &v
	return s
}

type PutProductInstanceConfigResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutProductInstanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutProductInstanceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s PutProductInstanceConfigResponse) GoString() string {
	return s.String()
}

func (s *PutProductInstanceConfigResponse) SetHeaders(v map[string]*string) *PutProductInstanceConfigResponse {
	s.Headers = v
	return s
}

func (s *PutProductInstanceConfigResponse) SetBody(v *PutProductInstanceConfigResponseBody) *PutProductInstanceConfigResponse {
	s.Body = v
	return s
}

type UpdateEnvironmentProductVersionRequest struct {
	ProductVersionUID *string `json:"productVersionUID,omitempty" xml:"productVersionUID,omitempty"`
}

func (s UpdateEnvironmentProductVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentProductVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentProductVersionRequest) SetProductVersionUID(v string) *UpdateEnvironmentProductVersionRequest {
	s.ProductVersionUID = &v
	return s
}

type UpdateEnvironmentProductVersionResponseBody struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s UpdateEnvironmentProductVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentProductVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentProductVersionResponseBody) SetCode(v string) *UpdateEnvironmentProductVersionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEnvironmentProductVersionResponseBody) SetMsg(v string) *UpdateEnvironmentProductVersionResponseBody {
	s.Msg = &v
	return s
}

type UpdateEnvironmentProductVersionResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateEnvironmentProductVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateEnvironmentProductVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentProductVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentProductVersionResponse) SetHeaders(v map[string]*string) *UpdateEnvironmentProductVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateEnvironmentProductVersionResponse) SetBody(v *UpdateEnvironmentProductVersionResponseBody) *UpdateEnvironmentProductVersionResponse {
	s.Body = v
	return s
}

type GetProductVersionResponseBody struct {
	Data *GetProductVersionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Code *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string                            `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s GetProductVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProductVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductVersionResponseBody) SetData(v *GetProductVersionResponseBodyData) *GetProductVersionResponseBody {
	s.Data = v
	return s
}

func (s *GetProductVersionResponseBody) SetCode(v string) *GetProductVersionResponseBody {
	s.Code = &v
	return s
}

func (s *GetProductVersionResponseBody) SetMsg(v string) *GetProductVersionResponseBody {
	s.Msg = &v
	return s
}

type GetProductVersionResponseBodyData struct {
	Description          *string     `json:"description,omitempty" xml:"description,omitempty"`
	Provider             *string     `json:"provider,omitempty" xml:"provider,omitempty"`
	Uid                  *string     `json:"uid,omitempty" xml:"uid,omitempty"`
	ProductUID           *string     `json:"productUID,omitempty" xml:"productUID,omitempty"`
	ProductName          *string     `json:"productName,omitempty" xml:"productName,omitempty"`
	Version              *string     `json:"version,omitempty" xml:"version,omitempty"`
	FoundationVersionUID *string     `json:"foundationVersionUID,omitempty" xml:"foundationVersionUID,omitempty"`
	PackageURL           *string     `json:"packageURL,omitempty" xml:"packageURL,omitempty"`
	Platforms            []*Platform `json:"platforms,omitempty" xml:"platforms,omitempty" type:"Repeated"`
}

func (s GetProductVersionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetProductVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProductVersionResponseBodyData) SetDescription(v string) *GetProductVersionResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetProductVersionResponseBodyData) SetProvider(v string) *GetProductVersionResponseBodyData {
	s.Provider = &v
	return s
}

func (s *GetProductVersionResponseBodyData) SetUid(v string) *GetProductVersionResponseBodyData {
	s.Uid = &v
	return s
}

func (s *GetProductVersionResponseBodyData) SetProductUID(v string) *GetProductVersionResponseBodyData {
	s.ProductUID = &v
	return s
}

func (s *GetProductVersionResponseBodyData) SetProductName(v string) *GetProductVersionResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *GetProductVersionResponseBodyData) SetVersion(v string) *GetProductVersionResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetProductVersionResponseBodyData) SetFoundationVersionUID(v string) *GetProductVersionResponseBodyData {
	s.FoundationVersionUID = &v
	return s
}

func (s *GetProductVersionResponseBodyData) SetPackageURL(v string) *GetProductVersionResponseBodyData {
	s.PackageURL = &v
	return s
}

func (s *GetProductVersionResponseBodyData) SetPlatforms(v []*Platform) *GetProductVersionResponseBodyData {
	s.Platforms = v
	return s
}

type GetProductVersionResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProductVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProductVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProductVersionResponse) GoString() string {
	return s.String()
}

func (s *GetProductVersionResponse) SetHeaders(v map[string]*string) *GetProductVersionResponse {
	s.Headers = v
	return s
}

func (s *GetProductVersionResponse) SetBody(v *GetProductVersionResponseBody) *GetProductVersionResponse {
	s.Body = v
	return s
}

type UpdateProductFoundationVersionRequest struct {
	// 底座版本uid
	FoundationVersionUID *string `json:"foundationVersionUID,omitempty" xml:"foundationVersionUID,omitempty"`
}

func (s UpdateProductFoundationVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductFoundationVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateProductFoundationVersionRequest) SetFoundationVersionUID(v string) *UpdateProductFoundationVersionRequest {
	s.FoundationVersionUID = &v
	return s
}

type UpdateProductFoundationVersionResponseBody struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s UpdateProductFoundationVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductFoundationVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProductFoundationVersionResponseBody) SetCode(v string) *UpdateProductFoundationVersionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateProductFoundationVersionResponseBody) SetMsg(v string) *UpdateProductFoundationVersionResponseBody {
	s.Msg = &v
	return s
}

type UpdateProductFoundationVersionResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateProductFoundationVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProductFoundationVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductFoundationVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateProductFoundationVersionResponse) SetHeaders(v map[string]*string) *UpdateProductFoundationVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateProductFoundationVersionResponse) SetBody(v *UpdateProductFoundationVersionResponseBody) *UpdateProductFoundationVersionResponse {
	s.Body = v
	return s
}

type GetProductResponseBody struct {
	Data *GetProductResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Code *string                     `json:"code,omitempty" xml:"code,omitempty"`
	Msg  *string                     `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s GetProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProductResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductResponseBody) SetData(v *GetProductResponseBodyData) *GetProductResponseBody {
	s.Data = v
	return s
}

func (s *GetProductResponseBody) SetCode(v string) *GetProductResponseBody {
	s.Code = &v
	return s
}

func (s *GetProductResponseBody) SetMsg(v string) *GetProductResponseBody {
	s.Msg = &v
	return s
}

type GetProductResponseBodyData struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Uid         *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s GetProductResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetProductResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProductResponseBodyData) SetDescription(v string) *GetProductResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetProductResponseBodyData) SetName(v string) *GetProductResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetProductResponseBodyData) SetUid(v string) *GetProductResponseBodyData {
	s.Uid = &v
	return s
}

type GetProductResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProductResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProductResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProductResponse) GoString() string {
	return s.String()
}

func (s *GetProductResponse) SetHeaders(v map[string]*string) *GetProductResponse {
	s.Headers = v
	return s
}

func (s *GetProductResponse) SetBody(v *GetProductResponseBody) *GetProductResponse {
	s.Body = v
	return s
}

type UpdateProductVersionConfigRequest struct {
	// 配置信息key
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 配置信息value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 配置说明
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 子组件versinid
	ComponentVersionUID *string `json:"componentVersionUID,omitempty" xml:"componentVersionUID,omitempty"`
	// 父组件versionid
	ParentComponentVersionUID *string `json:"parentComponentVersionUID,omitempty" xml:"parentComponentVersionUID,omitempty"`
}

func (s UpdateProductVersionConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductVersionConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateProductVersionConfigRequest) SetName(v string) *UpdateProductVersionConfigRequest {
	s.Name = &v
	return s
}

func (s *UpdateProductVersionConfigRequest) SetValue(v string) *UpdateProductVersionConfigRequest {
	s.Value = &v
	return s
}

func (s *UpdateProductVersionConfigRequest) SetDescription(v string) *UpdateProductVersionConfigRequest {
	s.Description = &v
	return s
}

func (s *UpdateProductVersionConfigRequest) SetComponentVersionUID(v string) *UpdateProductVersionConfigRequest {
	s.ComponentVersionUID = &v
	return s
}

func (s *UpdateProductVersionConfigRequest) SetParentComponentVersionUID(v string) *UpdateProductVersionConfigRequest {
	s.ParentComponentVersionUID = &v
	return s
}

type UpdateProductVersionConfigResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Msg       *string `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s UpdateProductVersionConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductVersionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProductVersionConfigResponseBody) SetRequestId(v string) *UpdateProductVersionConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProductVersionConfigResponseBody) SetCode(v string) *UpdateProductVersionConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateProductVersionConfigResponseBody) SetMsg(v string) *UpdateProductVersionConfigResponseBody {
	s.Msg = &v
	return s
}

type UpdateProductVersionConfigResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateProductVersionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProductVersionConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductVersionConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateProductVersionConfigResponse) SetHeaders(v map[string]*string) *UpdateProductVersionConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateProductVersionConfigResponse) SetBody(v *UpdateProductVersionConfigResponseBody) *UpdateProductVersionConfigResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("adp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GetEnvironment(uid *string) (_result *GetEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEnvironmentResponse{}
	_body, _err := client.GetEnvironmentWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEnvironmentWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEnvironmentResponse, _err error) {
	uid = openapiutil.GetEncodeParam(uid)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetEnvironment"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/environments/" + tea.StringValue(uid)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEnvironmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProductVersionPackage(uid *string, request *GetProductVersionPackageRequest) (_result *GetProductVersionPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProductVersionPackageResponse{}
	_body, _err := client.GetProductVersionPackageWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProductVersionPackageWithOptions(uid *string, request *GetProductVersionPackageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProductVersionPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	uid = openapiutil.GetEncodeParam(uid)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PackageUID)) {
		query["packageUID"] = request.PackageUID
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		query["platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.PackageType)) {
		query["packageType"] = request.PackageType
	}

	if !tea.BoolValue(util.IsUnset(request.PackageContentType)) {
		query["packageContentType"] = request.PackageContentType
	}

	if !tea.BoolValue(util.IsUnset(request.OldProductVersionUID)) {
		query["oldProductVersionUID"] = request.OldProductVersionUID
	}

	if !tea.BoolValue(util.IsUnset(request.WithURL)) {
		query["withURL"] = request.WithURL
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProductVersionPackage"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/hosting/product-versions/" + tea.StringValue(uid) + "/packages"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProductVersionPackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProductDeployments(request *ListProductDeploymentsRequest) (_result *ListProductDeploymentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProductDeploymentsResponse{}
	_body, _err := client.ListProductDeploymentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProductDeploymentsWithOptions(request *ListProductDeploymentsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProductDeploymentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvironmentUID)) {
		query["environmentUID"] = request.EnvironmentUID
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersionUID)) {
		query["productVersionUID"] = request.ProductVersionUID
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProductDeployments"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/product-instances/deployments"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProductDeploymentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEnvironmentNodes(uid *string, request *ListEnvironmentNodesRequest) (_result *ListEnvironmentNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEnvironmentNodesResponse{}
	_body, _err := client.ListEnvironmentNodesWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEnvironmentNodesWithOptions(uid *string, request *ListEnvironmentNodesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEnvironmentNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	uid = openapiutil.GetEncodeParam(uid)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEnvironmentNodes"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/environments/" + tea.StringValue(uid) + "/nodes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEnvironmentNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddProductComponentVersion(uid *string, componentVersionUID *string, request *AddProductComponentVersionRequest) (_result *AddProductComponentVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddProductComponentVersionResponse{}
	_body, _err := client.AddProductComponentVersionWithOptions(uid, componentVersionUID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddProductComponentVersionWithOptions(uid *string, componentVersionUID *string, request *AddProductComponentVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddProductComponentVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	uid = openapiutil.GetEncodeParam(uid)
	componentVersionUID = openapiutil.GetEncodeParam(componentVersionUID)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ReleaseName)) {
		body["releaseName"] = request.ReleaseName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddProductComponentVersion"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/integration/api/v2/product-versions/" + tea.StringValue(uid) + "/component-versions/{componentVersionUID}"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddProductComponentVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEnvironmentNode(uid *string, nodeUID *string, request *UpdateEnvironmentNodeRequest) (_result *UpdateEnvironmentNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateEnvironmentNodeResponse{}
	_body, _err := client.UpdateEnvironmentNodeWithOptions(uid, nodeUID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEnvironmentNodeWithOptions(uid *string, nodeUID *string, request *UpdateEnvironmentNodeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateEnvironmentNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	uid = openapiutil.GetEncodeParam(uid)
	nodeUID = openapiutil.GetEncodeParam(nodeUID)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.RootPassword)) {
		body["rootPassword"] = request.RootPassword
	}

	if !tea.BoolValue(util.IsUnset(request.Taints)) {
		body["taints"] = request.Taints
	}

	if !tea.BoolValue(util.IsUnset(request.EtcdDisk)) {
		body["etcdDisk"] = request.EtcdDisk
	}

	if !tea.BoolValue(util.IsUnset(request.TridentSystemDisk)) {
		body["tridentSystemDisk"] = request.TridentSystemDisk
	}

	if !tea.BoolValue(util.IsUnset(request.TridentSystemSizeDisk)) {
		body["tridentSystemSizeDisk"] = request.TridentSystemSizeDisk
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationDisk)) {
		body["applicationDisk"] = request.ApplicationDisk
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEnvironmentNode"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/environments/" + tea.StringValue(uid) + "/nodes/{nodeUID}"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEnvironmentNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProductVersion(uid *string) (_result *DeleteProductVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteProductVersionResponse{}
	_body, _err := client.DeleteProductVersionWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProductVersionWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteProductVersionResponse, _err error) {
	uid = openapiutil.GetEncodeParam(uid)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProductVersion"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/products/" + tea.StringValue(uid) + "/versions"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProductVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProduct(request *CreateProductRequest) (_result *CreateProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateProductHeaders{}
	_result = &CreateProductResponse{}
	_body, _err := client.CreateProductWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProductWithOptions(request *CreateProductRequest, headers *CreateProductHeaders, runtime *util.RuntimeOptions) (_result *CreateProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FoundationVersionUID)) {
		body["foundationVersionUID"] = request.FoundationVersionUID
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		body["productName"] = request.ProductName
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ClientToken)) {
		realHeaders["ClientToken"] = headers.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProduct"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/integration/api/v2/products"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateEnvironmentTunnel(uid *string, request *ValidateEnvironmentTunnelRequest) (_result *ValidateEnvironmentTunnelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ValidateEnvironmentTunnelResponse{}
	_body, _err := client.ValidateEnvironmentTunnelWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateEnvironmentTunnelWithOptions(uid *string, request *ValidateEnvironmentTunnelRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ValidateEnvironmentTunnelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	uid = openapiutil.GetEncodeParam(uid)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TunnelType)) {
		body["tunnelType"] = request.TunnelType
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TunnelConfig))) {
		body["tunnelConfig"] = request.TunnelConfig
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ValidateEnvironmentTunnel"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/environments/" + tea.StringValue(uid) + "/tunnels/validation"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ValidateEnvironmentTunnelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProductDeployment(request *CreateProductDeploymentRequest) (_result *CreateProductDeploymentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateProductDeploymentResponse{}
	_body, _err := client.CreateProductDeploymentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProductDeploymentWithOptions(request *CreateProductDeploymentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateProductDeploymentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PackageUID)) {
		body["packageUID"] = request.PackageUID
	}

	if !tea.BoolValue(util.IsUnset(request.EnvironmentUID)) {
		body["environmentUID"] = request.EnvironmentUID
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersionUID)) {
		body["productVersionUID"] = request.ProductVersionUID
	}

	if !tea.BoolValue(util.IsUnset(request.OldProductVersionUID)) {
		body["oldProductVersionUID"] = request.OldProductVersionUID
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProductDeployment"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/product-instances/deployments"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProductDeploymentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProductComponentVersion(uid *string, relationUID *string) (_result *DeleteProductComponentVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteProductComponentVersionResponse{}
	_body, _err := client.DeleteProductComponentVersionWithOptions(uid, relationUID, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProductComponentVersionWithOptions(uid *string, relationUID *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteProductComponentVersionResponse, _err error) {
	uid = openapiutil.GetEncodeParam(uid)
	relationUID = openapiutil.GetEncodeParam(relationUID)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProductComponentVersion"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/product-versions/" + tea.StringValue(uid) + "/relations/{relationUID}"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProductComponentVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEnvironment(uid *string) (_result *DeleteEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteEnvironmentResponse{}
	_body, _err := client.DeleteEnvironmentWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEnvironmentWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteEnvironmentResponse, _err error) {
	uid = openapiutil.GetEncodeParam(uid)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEnvironment"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/environments/" + tea.StringValue(uid)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEnvironmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProductVersion(uid *string) (_result *CreateProductVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateProductVersionResponse{}
	_body, _err := client.CreateProductVersionWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProductVersionWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateProductVersionResponse, _err error) {
	uid = openapiutil.GetEncodeParam(uid)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProductVersion"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/integration/api/v2/products/" + tea.StringValue(uid) + "/versions"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProductVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutEnvironmentTunnel(uid *string, request *PutEnvironmentTunnelRequest) (_result *PutEnvironmentTunnelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutEnvironmentTunnelResponse{}
	_body, _err := client.PutEnvironmentTunnelWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutEnvironmentTunnelWithOptions(uid *string, request *PutEnvironmentTunnelRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutEnvironmentTunnelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	uid = openapiutil.GetEncodeParam(uid)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TunnelType)) {
		body["tunnelType"] = request.TunnelType
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TunnelConfig))) {
		body["tunnelConfig"] = request.TunnelConfig
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutEnvironmentTunnel"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/environments/" + tea.StringValue(uid) + "/tunnels"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PutEnvironmentTunnelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProductEnvironments(uid *string, request *ListProductEnvironmentsRequest) (_result *ListProductEnvironmentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProductEnvironmentsResponse{}
	_body, _err := client.ListProductEnvironmentsWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProductEnvironmentsWithOptions(uid *string, tmpReq *ListProductEnvironmentsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProductEnvironmentsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	uid = openapiutil.GetEncodeParam(uid)
	request := &ListProductEnvironmentsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Platforms)) {
		request.PlatformsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Platforms, tea.String("platforms"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvType)) {
		query["envType"] = request.EnvType
	}

	if !tea.BoolValue(util.IsUnset(request.CompatibleProductVersionUID)) {
		query["compatibleProductVersionUID"] = request.CompatibleProductVersionUID
	}

	if !tea.BoolValue(util.IsUnset(request.PlatformsShrink)) {
		query["platforms"] = request.PlatformsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersionUID)) {
		query["productVersionUID"] = request.ProductVersionUID
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProductEnvironments"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/hosting/products/" + tea.StringValue(uid) + "/environments"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProductEnvironmentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEnvironmentNode(uid *string, nodeUID *string) (_result *DeleteEnvironmentNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteEnvironmentNodeResponse{}
	_body, _err := client.DeleteEnvironmentNodeWithOptions(uid, nodeUID, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEnvironmentNodeWithOptions(uid *string, nodeUID *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteEnvironmentNodeResponse, _err error) {
	uid = openapiutil.GetEncodeParam(uid)
	nodeUID = openapiutil.GetEncodeParam(nodeUID)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEnvironmentNode"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/environments/" + tea.StringValue(uid) + "/nodes/{nodeUID}"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEnvironmentNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProductComponentVersion(uid *string, relationUID *string, request *UpdateProductComponentVersionRequest) (_result *UpdateProductComponentVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProductComponentVersionResponse{}
	_body, _err := client.UpdateProductComponentVersionWithOptions(uid, relationUID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProductComponentVersionWithOptions(uid *string, relationUID *string, request *UpdateProductComponentVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateProductComponentVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	uid = openapiutil.GetEncodeParam(uid)
	relationUID = openapiutil.GetEncodeParam(relationUID)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ComponentOrchestrationValues)) {
		body["componentOrchestrationValues"] = request.ComponentOrchestrationValues
	}

	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		body["enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.ReleaseName)) {
		body["releaseName"] = request.ReleaseName
	}

	if !tea.BoolValue(util.IsUnset(request.NewComponentVersionUID)) {
		body["newComponentVersionUID"] = request.NewComponentVersionUID
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProductComponentVersion"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/product-versions/" + tea.StringValue(uid) + "/relations/{relationUID}"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProductComponentVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProduct(uid *string) (_result *DeleteProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteProductResponse{}
	_body, _err := client.DeleteProductWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProductWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteProductResponse, _err error) {
	uid = openapiutil.GetEncodeParam(uid)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProduct"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/integration/api/v2/products/" + tea.StringValue(uid)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddProductVersionConfig(uid *string, request *AddProductVersionConfigRequest) (_result *AddProductVersionConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddProductVersionConfigResponse{}
	_body, _err := client.AddProductVersionConfigWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddProductVersionConfigWithOptions(uid *string, request *AddProductVersionConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddProductVersionConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	uid = openapiutil.GetEncodeParam(uid)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		body["value"] = request.Value
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ComponentVersionUID)) {
		body["componentVersionUID"] = request.ComponentVersionUID
	}

	if !tea.BoolValue(util.IsUnset(request.ParentComponentVersionUID)) {
		body["parentComponentVersionUID"] = request.ParentComponentVersionUID
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddProductVersionConfig"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/product-versions/" + tea.StringValue(uid) + "/configs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddProductVersionConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitEnvironmentResource(uid *string, request *InitEnvironmentResourceRequest) (_result *InitEnvironmentResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InitEnvironmentResourceResponse{}
	_body, _err := client.InitEnvironmentResourceWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitEnvironmentResourceWithOptions(uid *string, request *InitEnvironmentResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *InitEnvironmentResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	uid = openapiutil.GetEncodeParam(uid)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKeyID)) {
		body["accessKeyID"] = request.AccessKeyID
	}

	if !tea.BoolValue(util.IsUnset(request.AccessKeySecret)) {
		body["accessKeySecret"] = request.AccessKeySecret
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		body["securityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InitEnvironmentResource"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/environments/" + tea.StringValue(uid) + "/resources"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &InitEnvironmentResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProductInstanceConfig(configUID *string, request *DeleteProductInstanceConfigRequest) (_result *DeleteProductInstanceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteProductInstanceConfigResponse{}
	_body, _err := client.DeleteProductInstanceConfigWithOptions(configUID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProductInstanceConfigWithOptions(configUID *string, request *DeleteProductInstanceConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteProductInstanceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	configUID = openapiutil.GetEncodeParam(configUID)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvironmentUID)) {
		query["environmentUID"] = request.EnvironmentUID
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersionUID)) {
		query["productVersionUID"] = request.ProductVersionUID
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProductInstanceConfig"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/product-instances/configs/" + tea.StringValue(configUID)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProductInstanceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEnvironment(request *CreateEnvironmentRequest) (_result *CreateEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateEnvironmentHeaders{}
	_result = &CreateEnvironmentResponse{}
	_body, _err := client.CreateEnvironmentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEnvironmentWithOptions(request *CreateEnvironmentRequest, headers *CreateEnvironmentHeaders, runtime *util.RuntimeOptions) (_result *CreateEnvironmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Platform))) {
		body["platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.VendorType)) {
		body["vendorType"] = request.VendorType
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersionUID)) {
		body["productVersionUID"] = request.ProductVersionUID
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ClientToken)) {
		realHeaders["ClientToken"] = headers.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEnvironment"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/environments"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEnvironmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProductVersionConfigs(uid *string, request *ListProductVersionConfigsRequest) (_result *ListProductVersionConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProductVersionConfigsResponse{}
	_body, _err := client.ListProductVersionConfigsWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProductVersionConfigsWithOptions(uid *string, request *ListProductVersionConfigsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProductVersionConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	uid = openapiutil.GetEncodeParam(uid)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigType)) {
		query["configType"] = request.ConfigType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProductVersionConfigs"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/product-versions/" + tea.StringValue(uid) + "/configs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProductVersionConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddEnvironmentNodes(uid *string, request *AddEnvironmentNodesRequest) (_result *AddEnvironmentNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddEnvironmentNodesResponse{}
	_body, _err := client.AddEnvironmentNodesWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddEnvironmentNodesWithOptions(uid *string, request *AddEnvironmentNodesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddEnvironmentNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	uid = openapiutil.GetEncodeParam(uid)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.MasterPrivateIPs)) {
		body["masterPrivateIPs"] = request.MasterPrivateIPs
	}

	if !tea.BoolValue(util.IsUnset(request.RootPassword)) {
		body["rootPassword"] = request.RootPassword
	}

	if !tea.BoolValue(util.IsUnset(request.Taints)) {
		body["taints"] = request.Taints
	}

	if !tea.BoolValue(util.IsUnset(request.WorkerPrivateIPs)) {
		body["workerPrivateIPs"] = request.WorkerPrivateIPs
	}

	if !tea.BoolValue(util.IsUnset(request.EtcdDisk)) {
		body["etcdDisk"] = request.EtcdDisk
	}

	if !tea.BoolValue(util.IsUnset(request.TridentSystemDisk)) {
		body["tridentSystemDisk"] = request.TridentSystemDisk
	}

	if !tea.BoolValue(util.IsUnset(request.TridentSystemSizeDisk)) {
		body["tridentSystemSizeDisk"] = request.TridentSystemSizeDisk
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationDisk)) {
		body["applicationDisk"] = request.ApplicationDisk
	}

	if !tea.BoolValue(util.IsUnset(request.Cpu)) {
		body["cpu"] = request.Cpu
	}

	if !tea.BoolValue(util.IsUnset(request.DataDisk)) {
		body["dataDisk"] = request.DataDisk
	}

	if !tea.BoolValue(util.IsUnset(request.HostName)) {
		body["hostName"] = request.HostName
	}

	if !tea.BoolValue(util.IsUnset(request.Memory)) {
		body["memory"] = request.Memory
	}

	if !tea.BoolValue(util.IsUnset(request.Os)) {
		body["os"] = request.Os
	}

	if !tea.BoolValue(util.IsUnset(request.SystemDisk)) {
		body["systemDisk"] = request.SystemDisk
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddEnvironmentNodes"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/environments/" + tea.StringValue(uid) + "/nodes"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddEnvironmentNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProductVersionPackage(uid *string, request *CreateProductVersionPackageRequest) (_result *CreateProductVersionPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateProductVersionPackageHeaders{}
	_result = &CreateProductVersionPackageResponse{}
	_body, _err := client.CreateProductVersionPackageWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProductVersionPackageWithOptions(uid *string, request *CreateProductVersionPackageRequest, headers *CreateProductVersionPackageHeaders, runtime *util.RuntimeOptions) (_result *CreateProductVersionPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	uid = openapiutil.GetEncodeParam(uid)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PackageType)) {
		query["packageType"] = request.PackageType
	}

	if !tea.BoolValue(util.IsUnset(request.PackageContentType)) {
		query["packageContentType"] = request.PackageContentType
	}

	if !tea.BoolValue(util.IsUnset(request.OldProductVersionUID)) {
		query["oldProductVersionUID"] = request.OldProductVersionUID
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		query["platform"] = request.Platform
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ClientToken)) {
		realHeaders["ClientToken"] = headers.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProductVersionPackage"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/hosting/product-versions/" + tea.StringValue(uid) + "/packages"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProductVersionPackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProductInstanceConfigs(request *ListProductInstanceConfigsRequest) (_result *ListProductInstanceConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProductInstanceConfigsResponse{}
	_body, _err := client.ListProductInstanceConfigsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProductInstanceConfigsWithOptions(request *ListProductInstanceConfigsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProductInstanceConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ParamType)) {
		query["paramType"] = request.ParamType
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersionUID)) {
		query["productVersionUID"] = request.ProductVersionUID
	}

	if !tea.BoolValue(util.IsUnset(request.EnvironmentUID)) {
		query["environmentUID"] = request.EnvironmentUID
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProductInstanceConfigs"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/product-instances/configs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProductInstanceConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProductVersion(uid *string, request *UpdateProductVersionRequest) (_result *UpdateProductVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProductVersionResponse{}
	_body, _err := client.UpdateProductVersionWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProductVersionWithOptions(uid *string, request *UpdateProductVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateProductVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	uid = openapiutil.GetEncodeParam(uid)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		body["version"] = request.Version
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProductVersion"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/product-versions/" + tea.StringValue(uid)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProductVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProductComponentVersion(relationUID *string, uid *string) (_result *GetProductComponentVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProductComponentVersionResponse{}
	_body, _err := client.GetProductComponentVersionWithOptions(relationUID, uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProductComponentVersionWithOptions(relationUID *string, uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProductComponentVersionResponse, _err error) {
	relationUID = openapiutil.GetEncodeParam(relationUID)
	uid = openapiutil.GetEncodeParam(uid)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetProductComponentVersion"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/integration/api/v2/product-versions/" + tea.StringValue(uid) + "/relations/{relationUID}"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProductComponentVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEnvironments(request *ListEnvironmentsRequest) (_result *ListEnvironmentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEnvironmentsResponse{}
	_body, _err := client.ListEnvironmentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEnvironmentsWithOptions(request *ListEnvironmentsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEnvironmentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Fuzzy)) {
		query["fuzzy"] = request.Fuzzy
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceStatus)) {
		query["instanceStatus"] = request.InstanceStatus
	}

	if !tea.BoolValue(util.IsUnset(request.VendorType)) {
		query["vendorType"] = request.VendorType
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterUID)) {
		query["clusterUID"] = request.ClusterUID
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEnvironments"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/environments"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEnvironmentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProduct(uid *string, request *UpdateProductRequest) (_result *UpdateProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProductResponse{}
	_body, _err := client.UpdateProductWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProductWithOptions(uid *string, request *UpdateProductRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	uid = openapiutil.GetEncodeParam(uid)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProduct"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/products/" + tea.StringValue(uid)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProductVersionConfig(uid *string, configUID *string) (_result *DeleteProductVersionConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteProductVersionConfigResponse{}
	_body, _err := client.DeleteProductVersionConfigWithOptions(uid, configUID, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProductVersionConfigWithOptions(uid *string, configUID *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteProductVersionConfigResponse, _err error) {
	uid = openapiutil.GetEncodeParam(uid)
	configUID = openapiutil.GetEncodeParam(configUID)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProductVersionConfig"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/product-versions/" + tea.StringValue(uid) + "/configs/{configUID}"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProductVersionConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProductComponentVersions(uid *string, request *ListProductComponentVersionsRequest) (_result *ListProductComponentVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProductComponentVersionsResponse{}
	_body, _err := client.ListProductComponentVersionsWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProductComponentVersionsWithOptions(uid *string, request *ListProductComponentVersionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProductComponentVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	uid = openapiutil.GetEncodeParam(uid)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortKey)) {
		query["sortKey"] = request.SortKey
	}

	if !tea.BoolValue(util.IsUnset(request.SortDirect)) {
		query["sortDirect"] = request.SortDirect
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["category"] = request.Category
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProductComponentVersions"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/product-versions/" + tea.StringValue(uid) + "/component-versions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProductComponentVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEnvironmentNode(uid *string, nodeUID *string) (_result *GetEnvironmentNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEnvironmentNodeResponse{}
	_body, _err := client.GetEnvironmentNodeWithOptions(uid, nodeUID, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEnvironmentNodeWithOptions(uid *string, nodeUID *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEnvironmentNodeResponse, _err error) {
	uid = openapiutil.GetEncodeParam(uid)
	nodeUID = openapiutil.GetEncodeParam(nodeUID)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetEnvironmentNode"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/environments/" + tea.StringValue(uid) + "/nodes/{nodeUID}"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEnvironmentNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProducts(request *ListProductsRequest) (_result *ListProductsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProductsResponse{}
	_body, _err := client.ListProductsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProductsWithOptions(request *ListProductsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProductsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Fuzzy)) {
		query["fuzzy"] = request.Fuzzy
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProducts"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/products"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProductsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateProductInstanceDeploymentConfig(request *GenerateProductInstanceDeploymentConfigRequest) (_result *GenerateProductInstanceDeploymentConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateProductInstanceDeploymentConfigResponse{}
	_body, _err := client.GenerateProductInstanceDeploymentConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateProductInstanceDeploymentConfigWithOptions(request *GenerateProductInstanceDeploymentConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GenerateProductInstanceDeploymentConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvironmentUID)) {
		body["environmentUID"] = request.EnvironmentUID
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersionUID)) {
		body["productVersionUID"] = request.ProductVersionUID
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateProductInstanceDeploymentConfig"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/product-instances/package-configs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateProductInstanceDeploymentConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProductDeployment(deploymentUID *string, request *GetProductDeploymentRequest) (_result *GetProductDeploymentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProductDeploymentResponse{}
	_body, _err := client.GetProductDeploymentWithOptions(deploymentUID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProductDeploymentWithOptions(deploymentUID *string, request *GetProductDeploymentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProductDeploymentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	deploymentUID = openapiutil.GetEncodeParam(deploymentUID)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvironmentUID)) {
		query["environmentUID"] = request.EnvironmentUID
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersionUID)) {
		query["productVersionUID"] = request.ProductVersionUID
	}

	if !tea.BoolValue(util.IsUnset(request.WithParamConfig)) {
		query["withParamConfig"] = request.WithParamConfig
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProductDeployment"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/product-instances/deployments/" + tea.StringValue(deploymentUID)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProductDeploymentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFoundationReferenceComponents(request *ListFoundationReferenceComponentsRequest) (_result *ListFoundationReferenceComponentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFoundationReferenceComponentsResponse{}
	_body, _err := client.ListFoundationReferenceComponentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFoundationReferenceComponentsWithOptions(request *ListFoundationReferenceComponentsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFoundationReferenceComponentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductVersionUID)) {
		query["productVersionUID"] = request.ProductVersionUID
	}

	if !tea.BoolValue(util.IsUnset(request.FoundationVersionUID)) {
		query["foundationVersionUID"] = request.FoundationVersionUID
	}

	if !tea.BoolValue(util.IsUnset(request.OnlyEnabled)) {
		query["onlyEnabled"] = request.OnlyEnabled
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFoundationReferenceComponents"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/foundation-references/component-versions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFoundationReferenceComponentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProductVersionDifferences(uid *string, versionUID *string, request *GetProductVersionDifferencesRequest) (_result *GetProductVersionDifferencesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProductVersionDifferencesResponse{}
	_body, _err := client.GetProductVersionDifferencesWithOptions(uid, versionUID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProductVersionDifferencesWithOptions(uid *string, versionUID *string, request *GetProductVersionDifferencesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProductVersionDifferencesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	uid = openapiutil.GetEncodeParam(uid)
	versionUID = openapiutil.GetEncodeParam(versionUID)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PreVersionUID)) {
		query["preVersionUID"] = request.PreVersionUID
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProductVersionDifferences"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/integration/api/v2/products/" + tea.StringValue(uid) + "/versions/{versionUID}/differences"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProductVersionDifferencesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEnvironment(uid *string, request *UpdateEnvironmentRequest) (_result *UpdateEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateEnvironmentResponse{}
	_body, _err := client.UpdateEnvironmentWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEnvironmentWithOptions(uid *string, request *UpdateEnvironmentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateEnvironmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	uid = openapiutil.GetEncodeParam(uid)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(request.VendorConfig)) {
		body["vendorConfig"] = request.VendorConfig
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEnvironment"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/environments/" + tea.StringValue(uid)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEnvironmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEnvironmentTunnels(uid *string) (_result *ListEnvironmentTunnelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEnvironmentTunnelsResponse{}
	_body, _err := client.ListEnvironmentTunnelsWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEnvironmentTunnelsWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEnvironmentTunnelsResponse, _err error) {
	uid = openapiutil.GetEncodeParam(uid)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListEnvironmentTunnels"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/environments/" + tea.StringValue(uid) + "/tunnels"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEnvironmentTunnelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProductVersions(request *ListProductVersionsRequest) (_result *ListProductVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProductVersionsResponse{}
	_body, _err := client.ListProductVersionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProductVersionsWithOptions(tmpReq *ListProductVersionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProductVersionsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListProductVersionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SupportedFoundationTypes)) {
		request.SupportedFoundationTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SupportedFoundationTypes, tea.String("supportedFoundationTypes"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Platforms)) {
		request.PlatformsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Platforms, tea.String("platforms"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductUID)) {
		query["productUID"] = request.ProductUID
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		query["productName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		query["version"] = request.Version
	}

	if !tea.BoolValue(util.IsUnset(request.Released)) {
		query["released"] = request.Released
	}

	if !tea.BoolValue(util.IsUnset(request.SupportedFoundationTypesShrink)) {
		query["supportedFoundationTypes"] = request.SupportedFoundationTypesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PlatformsShrink)) {
		query["platforms"] = request.PlatformsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProductVersions"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/product-versions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProductVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutProductInstanceConfig(request *PutProductInstanceConfigRequest) (_result *PutProductInstanceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutProductInstanceConfigResponse{}
	_body, _err := client.PutProductInstanceConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutProductInstanceConfigWithOptions(request *PutProductInstanceConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutProductInstanceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvironmentUID)) {
		body["environmentUID"] = request.EnvironmentUID
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersionUID)) {
		body["productVersionUID"] = request.ProductVersionUID
	}

	if !tea.BoolValue(util.IsUnset(request.ComponentUID)) {
		body["componentUID"] = request.ComponentUID
	}

	if !tea.BoolValue(util.IsUnset(request.ComponentVersionUID)) {
		body["componentVersionUID"] = request.ComponentVersionUID
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigUID)) {
		body["configUID"] = request.ConfigUID
	}

	if !tea.BoolValue(util.IsUnset(request.ReleaseName)) {
		body["releaseName"] = request.ReleaseName
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		body["scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		body["value"] = request.Value
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ParentComponentName)) {
		body["parentComponentName"] = request.ParentComponentName
	}

	if !tea.BoolValue(util.IsUnset(request.ParentComponentVersionUID)) {
		body["parentComponentVersionUID"] = request.ParentComponentVersionUID
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutProductInstanceConfig"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/product-instances/configs"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PutProductInstanceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEnvironmentProductVersion(uid *string, request *UpdateEnvironmentProductVersionRequest) (_result *UpdateEnvironmentProductVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateEnvironmentProductVersionResponse{}
	_body, _err := client.UpdateEnvironmentProductVersionWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEnvironmentProductVersionWithOptions(uid *string, request *UpdateEnvironmentProductVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateEnvironmentProductVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	uid = openapiutil.GetEncodeParam(uid)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductVersionUID)) {
		body["productVersionUID"] = request.ProductVersionUID
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEnvironmentProductVersion"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/environments/" + tea.StringValue(uid) + "/product-versions"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEnvironmentProductVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProductVersion(uid *string) (_result *GetProductVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProductVersionResponse{}
	_body, _err := client.GetProductVersionWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProductVersionWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProductVersionResponse, _err error) {
	uid = openapiutil.GetEncodeParam(uid)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetProductVersion"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/product-versions/" + tea.StringValue(uid)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProductVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProductFoundationVersion(uid *string, request *UpdateProductFoundationVersionRequest) (_result *UpdateProductFoundationVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProductFoundationVersionResponse{}
	_body, _err := client.UpdateProductFoundationVersionWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProductFoundationVersionWithOptions(uid *string, request *UpdateProductFoundationVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateProductFoundationVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	uid = openapiutil.GetEncodeParam(uid)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FoundationVersionUID)) {
		body["foundationVersionUID"] = request.FoundationVersionUID
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProductFoundationVersion"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/product-versions/" + tea.StringValue(uid) + "/foundation"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProductFoundationVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProduct(uid *string) (_result *GetProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProductResponse{}
	_body, _err := client.GetProductWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProductWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProductResponse, _err error) {
	uid = openapiutil.GetEncodeParam(uid)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetProduct"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/products/" + tea.StringValue(uid)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProductVersionConfig(uid *string, configUID *string, request *UpdateProductVersionConfigRequest) (_result *UpdateProductVersionConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProductVersionConfigResponse{}
	_body, _err := client.UpdateProductVersionConfigWithOptions(uid, configUID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProductVersionConfigWithOptions(uid *string, configUID *string, request *UpdateProductVersionConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateProductVersionConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	uid = openapiutil.GetEncodeParam(uid)
	configUID = openapiutil.GetEncodeParam(configUID)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		body["value"] = request.Value
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ComponentVersionUID)) {
		body["componentVersionUID"] = request.ComponentVersionUID
	}

	if !tea.BoolValue(util.IsUnset(request.ParentComponentVersionUID)) {
		body["parentComponentVersionUID"] = request.ParentComponentVersionUID
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProductVersionConfig"),
		Version:     tea.String("2021-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/product-versions/" + tea.StringValue(uid) + "/configs/{configUID}"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProductVersionConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
