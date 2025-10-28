// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCommand(v string) *AppConfig
	GetCommand() *string
	SetCommandArgs(v []*string) *AppConfig
	GetCommandArgs() []*string
	SetConfigMountDescs(v []*AppConfigConfigMountDescs) *AppConfig
	GetConfigMountDescs() []*AppConfigConfigMountDescs
	SetDeployAcrossNodes(v bool) *AppConfig
	GetDeployAcrossNodes() *bool
	SetDeployAcrossZones(v bool) *AppConfig
	GetDeployAcrossZones() *bool
	SetEmptyDirs(v []*AppConfigEmptyDirs) *AppConfig
	GetEmptyDirs() []*AppConfigEmptyDirs
	SetEnableAhas(v bool) *AppConfig
	GetEnableAhas() *bool
	SetEnvFroms(v []*AppConfigEnvFroms) *AppConfig
	GetEnvFroms() []*AppConfigEnvFroms
	SetEnvs(v []*AppConfigEnvs) *AppConfig
	GetEnvs() []*AppConfigEnvs
	SetImageConfig(v *AppConfigImageConfig) *AppConfig
	GetImageConfig() *AppConfigImageConfig
	SetIsMultilingualApp(v bool) *AppConfig
	GetIsMultilingualApp() *bool
	SetJavaStartUpConfig(v string) *AppConfig
	GetJavaStartUpConfig() *string
	SetLimitCpu(v string) *AppConfig
	GetLimitCpu() *string
	SetLimitMem(v string) *AppConfig
	GetLimitMem() *string
	SetLiveness(v string) *AppConfig
	GetLiveness() *string
	SetLocalVolumes(v []*AppConfigLocalVolumes) *AppConfig
	GetLocalVolumes() []*AppConfigLocalVolumes
	SetNasId(v string) *AppConfig
	GetNasId() *string
	SetNasMountDescs(v []*AppConfigNasMountDescs) *AppConfig
	GetNasMountDescs() []*AppConfigNasMountDescs
	SetNasStorageType(v string) *AppConfig
	GetNasStorageType() *string
	SetPackageConfig(v *AppConfigPackageConfig) *AppConfig
	GetPackageConfig() *AppConfigPackageConfig
	SetPostStart(v string) *AppConfig
	GetPostStart() *string
	SetPreStop(v string) *AppConfig
	GetPreStop() *string
	SetPvcMountDescs(v []*AppConfigPvcMountDescs) *AppConfig
	GetPvcMountDescs() []*AppConfigPvcMountDescs
	SetReadiness(v string) *AppConfig
	GetReadiness() *string
	SetReplicas(v int64) *AppConfig
	GetReplicas() *int64
	SetRequestCpu(v string) *AppConfig
	GetRequestCpu() *string
	SetRequestMem(v string) *AppConfig
	GetRequestMem() *string
	SetRuntimeClassName(v string) *AppConfig
	GetRuntimeClassName() *string
	SetSlsConfigs(v []*AppConfigSlsConfigs) *AppConfig
	GetSlsConfigs() []*AppConfigSlsConfigs
	SetWebContainerConfig(v *AppConfigWebContainerConfig) *AppConfig
	GetWebContainerConfig() *AppConfigWebContainerConfig
}

type AppConfig struct {
	Command            *string                      `json:"Command,omitempty" xml:"Command,omitempty"`
	CommandArgs        []*string                    `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty" type:"Repeated"`
	ConfigMountDescs   []*AppConfigConfigMountDescs `json:"ConfigMountDescs,omitempty" xml:"ConfigMountDescs,omitempty" type:"Repeated"`
	DeployAcrossNodes  *bool                        `json:"DeployAcrossNodes,omitempty" xml:"DeployAcrossNodes,omitempty"`
	DeployAcrossZones  *bool                        `json:"DeployAcrossZones,omitempty" xml:"DeployAcrossZones,omitempty"`
	EmptyDirs          []*AppConfigEmptyDirs        `json:"EmptyDirs,omitempty" xml:"EmptyDirs,omitempty" type:"Repeated"`
	EnableAhas         *bool                        `json:"EnableAhas,omitempty" xml:"EnableAhas,omitempty"`
	EnvFroms           []*AppConfigEnvFroms         `json:"EnvFroms,omitempty" xml:"EnvFroms,omitempty" type:"Repeated"`
	Envs               []*AppConfigEnvs             `json:"Envs,omitempty" xml:"Envs,omitempty" type:"Repeated"`
	ImageConfig        *AppConfigImageConfig        `json:"ImageConfig,omitempty" xml:"ImageConfig,omitempty" type:"Struct"`
	IsMultilingualApp  *bool                        `json:"IsMultilingualApp,omitempty" xml:"IsMultilingualApp,omitempty"`
	JavaStartUpConfig  *string                      `json:"JavaStartUpConfig,omitempty" xml:"JavaStartUpConfig,omitempty"`
	LimitCpu           *string                      `json:"LimitCpu,omitempty" xml:"LimitCpu,omitempty"`
	LimitMem           *string                      `json:"LimitMem,omitempty" xml:"LimitMem,omitempty"`
	Liveness           *string                      `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	LocalVolumes       []*AppConfigLocalVolumes     `json:"LocalVolumes,omitempty" xml:"LocalVolumes,omitempty" type:"Repeated"`
	NasId              *string                      `json:"NasId,omitempty" xml:"NasId,omitempty"`
	NasMountDescs      []*AppConfigNasMountDescs    `json:"NasMountDescs,omitempty" xml:"NasMountDescs,omitempty" type:"Repeated"`
	NasStorageType     *string                      `json:"NasStorageType,omitempty" xml:"NasStorageType,omitempty"`
	PackageConfig      *AppConfigPackageConfig      `json:"PackageConfig,omitempty" xml:"PackageConfig,omitempty" type:"Struct"`
	PostStart          *string                      `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	PreStop            *string                      `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	PvcMountDescs      []*AppConfigPvcMountDescs    `json:"PvcMountDescs,omitempty" xml:"PvcMountDescs,omitempty" type:"Repeated"`
	Readiness          *string                      `json:"Readiness,omitempty" xml:"Readiness,omitempty"`
	Replicas           *int64                       `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	RequestCpu         *string                      `json:"RequestCpu,omitempty" xml:"RequestCpu,omitempty"`
	RequestMem         *string                      `json:"RequestMem,omitempty" xml:"RequestMem,omitempty"`
	RuntimeClassName   *string                      `json:"RuntimeClassName,omitempty" xml:"RuntimeClassName,omitempty"`
	SlsConfigs         []*AppConfigSlsConfigs       `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty" type:"Repeated"`
	WebContainerConfig *AppConfigWebContainerConfig `json:"WebContainerConfig,omitempty" xml:"WebContainerConfig,omitempty" type:"Struct"`
}

func (s AppConfig) String() string {
	return dara.Prettify(s)
}

func (s AppConfig) GoString() string {
	return s.String()
}

func (s *AppConfig) GetCommand() *string {
	return s.Command
}

func (s *AppConfig) GetCommandArgs() []*string {
	return s.CommandArgs
}

func (s *AppConfig) GetConfigMountDescs() []*AppConfigConfigMountDescs {
	return s.ConfigMountDescs
}

func (s *AppConfig) GetDeployAcrossNodes() *bool {
	return s.DeployAcrossNodes
}

func (s *AppConfig) GetDeployAcrossZones() *bool {
	return s.DeployAcrossZones
}

func (s *AppConfig) GetEmptyDirs() []*AppConfigEmptyDirs {
	return s.EmptyDirs
}

func (s *AppConfig) GetEnableAhas() *bool {
	return s.EnableAhas
}

func (s *AppConfig) GetEnvFroms() []*AppConfigEnvFroms {
	return s.EnvFroms
}

func (s *AppConfig) GetEnvs() []*AppConfigEnvs {
	return s.Envs
}

func (s *AppConfig) GetImageConfig() *AppConfigImageConfig {
	return s.ImageConfig
}

func (s *AppConfig) GetIsMultilingualApp() *bool {
	return s.IsMultilingualApp
}

func (s *AppConfig) GetJavaStartUpConfig() *string {
	return s.JavaStartUpConfig
}

func (s *AppConfig) GetLimitCpu() *string {
	return s.LimitCpu
}

func (s *AppConfig) GetLimitMem() *string {
	return s.LimitMem
}

func (s *AppConfig) GetLiveness() *string {
	return s.Liveness
}

func (s *AppConfig) GetLocalVolumes() []*AppConfigLocalVolumes {
	return s.LocalVolumes
}

func (s *AppConfig) GetNasId() *string {
	return s.NasId
}

func (s *AppConfig) GetNasMountDescs() []*AppConfigNasMountDescs {
	return s.NasMountDescs
}

func (s *AppConfig) GetNasStorageType() *string {
	return s.NasStorageType
}

func (s *AppConfig) GetPackageConfig() *AppConfigPackageConfig {
	return s.PackageConfig
}

func (s *AppConfig) GetPostStart() *string {
	return s.PostStart
}

func (s *AppConfig) GetPreStop() *string {
	return s.PreStop
}

func (s *AppConfig) GetPvcMountDescs() []*AppConfigPvcMountDescs {
	return s.PvcMountDescs
}

func (s *AppConfig) GetReadiness() *string {
	return s.Readiness
}

func (s *AppConfig) GetReplicas() *int64 {
	return s.Replicas
}

func (s *AppConfig) GetRequestCpu() *string {
	return s.RequestCpu
}

func (s *AppConfig) GetRequestMem() *string {
	return s.RequestMem
}

func (s *AppConfig) GetRuntimeClassName() *string {
	return s.RuntimeClassName
}

func (s *AppConfig) GetSlsConfigs() []*AppConfigSlsConfigs {
	return s.SlsConfigs
}

func (s *AppConfig) GetWebContainerConfig() *AppConfigWebContainerConfig {
	return s.WebContainerConfig
}

func (s *AppConfig) SetCommand(v string) *AppConfig {
	s.Command = &v
	return s
}

func (s *AppConfig) SetCommandArgs(v []*string) *AppConfig {
	s.CommandArgs = v
	return s
}

func (s *AppConfig) SetConfigMountDescs(v []*AppConfigConfigMountDescs) *AppConfig {
	s.ConfigMountDescs = v
	return s
}

func (s *AppConfig) SetDeployAcrossNodes(v bool) *AppConfig {
	s.DeployAcrossNodes = &v
	return s
}

func (s *AppConfig) SetDeployAcrossZones(v bool) *AppConfig {
	s.DeployAcrossZones = &v
	return s
}

func (s *AppConfig) SetEmptyDirs(v []*AppConfigEmptyDirs) *AppConfig {
	s.EmptyDirs = v
	return s
}

func (s *AppConfig) SetEnableAhas(v bool) *AppConfig {
	s.EnableAhas = &v
	return s
}

func (s *AppConfig) SetEnvFroms(v []*AppConfigEnvFroms) *AppConfig {
	s.EnvFroms = v
	return s
}

func (s *AppConfig) SetEnvs(v []*AppConfigEnvs) *AppConfig {
	s.Envs = v
	return s
}

func (s *AppConfig) SetImageConfig(v *AppConfigImageConfig) *AppConfig {
	s.ImageConfig = v
	return s
}

func (s *AppConfig) SetIsMultilingualApp(v bool) *AppConfig {
	s.IsMultilingualApp = &v
	return s
}

func (s *AppConfig) SetJavaStartUpConfig(v string) *AppConfig {
	s.JavaStartUpConfig = &v
	return s
}

func (s *AppConfig) SetLimitCpu(v string) *AppConfig {
	s.LimitCpu = &v
	return s
}

func (s *AppConfig) SetLimitMem(v string) *AppConfig {
	s.LimitMem = &v
	return s
}

func (s *AppConfig) SetLiveness(v string) *AppConfig {
	s.Liveness = &v
	return s
}

func (s *AppConfig) SetLocalVolumes(v []*AppConfigLocalVolumes) *AppConfig {
	s.LocalVolumes = v
	return s
}

func (s *AppConfig) SetNasId(v string) *AppConfig {
	s.NasId = &v
	return s
}

func (s *AppConfig) SetNasMountDescs(v []*AppConfigNasMountDescs) *AppConfig {
	s.NasMountDescs = v
	return s
}

func (s *AppConfig) SetNasStorageType(v string) *AppConfig {
	s.NasStorageType = &v
	return s
}

func (s *AppConfig) SetPackageConfig(v *AppConfigPackageConfig) *AppConfig {
	s.PackageConfig = v
	return s
}

func (s *AppConfig) SetPostStart(v string) *AppConfig {
	s.PostStart = &v
	return s
}

func (s *AppConfig) SetPreStop(v string) *AppConfig {
	s.PreStop = &v
	return s
}

func (s *AppConfig) SetPvcMountDescs(v []*AppConfigPvcMountDescs) *AppConfig {
	s.PvcMountDescs = v
	return s
}

func (s *AppConfig) SetReadiness(v string) *AppConfig {
	s.Readiness = &v
	return s
}

func (s *AppConfig) SetReplicas(v int64) *AppConfig {
	s.Replicas = &v
	return s
}

func (s *AppConfig) SetRequestCpu(v string) *AppConfig {
	s.RequestCpu = &v
	return s
}

func (s *AppConfig) SetRequestMem(v string) *AppConfig {
	s.RequestMem = &v
	return s
}

func (s *AppConfig) SetRuntimeClassName(v string) *AppConfig {
	s.RuntimeClassName = &v
	return s
}

func (s *AppConfig) SetSlsConfigs(v []*AppConfigSlsConfigs) *AppConfig {
	s.SlsConfigs = v
	return s
}

func (s *AppConfig) SetWebContainerConfig(v *AppConfigWebContainerConfig) *AppConfig {
	s.WebContainerConfig = v
	return s
}

func (s *AppConfig) Validate() error {
	if s.ConfigMountDescs != nil {
		for _, item := range s.ConfigMountDescs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EmptyDirs != nil {
		for _, item := range s.EmptyDirs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EnvFroms != nil {
		for _, item := range s.EnvFroms {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Envs != nil {
		for _, item := range s.Envs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ImageConfig != nil {
		if err := s.ImageConfig.Validate(); err != nil {
			return err
		}
	}
	if s.LocalVolumes != nil {
		for _, item := range s.LocalVolumes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NasMountDescs != nil {
		for _, item := range s.NasMountDescs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PackageConfig != nil {
		if err := s.PackageConfig.Validate(); err != nil {
			return err
		}
	}
	if s.PvcMountDescs != nil {
		for _, item := range s.PvcMountDescs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SlsConfigs != nil {
		for _, item := range s.SlsConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WebContainerConfig != nil {
		if err := s.WebContainerConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AppConfigConfigMountDescs struct {
	MountItems []*AppConfigConfigMountDescsMountItems `json:"MountItems,omitempty" xml:"MountItems,omitempty" type:"Repeated"`
	MountPath  *string                                `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	Name       *string                                `json:"Name,omitempty" xml:"Name,omitempty"`
	Type       *string                                `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AppConfigConfigMountDescs) String() string {
	return dara.Prettify(s)
}

func (s AppConfigConfigMountDescs) GoString() string {
	return s.String()
}

func (s *AppConfigConfigMountDescs) GetMountItems() []*AppConfigConfigMountDescsMountItems {
	return s.MountItems
}

func (s *AppConfigConfigMountDescs) GetMountPath() *string {
	return s.MountPath
}

func (s *AppConfigConfigMountDescs) GetName() *string {
	return s.Name
}

func (s *AppConfigConfigMountDescs) GetType() *string {
	return s.Type
}

func (s *AppConfigConfigMountDescs) SetMountItems(v []*AppConfigConfigMountDescsMountItems) *AppConfigConfigMountDescs {
	s.MountItems = v
	return s
}

func (s *AppConfigConfigMountDescs) SetMountPath(v string) *AppConfigConfigMountDescs {
	s.MountPath = &v
	return s
}

func (s *AppConfigConfigMountDescs) SetName(v string) *AppConfigConfigMountDescs {
	s.Name = &v
	return s
}

func (s *AppConfigConfigMountDescs) SetType(v string) *AppConfigConfigMountDescs {
	s.Type = &v
	return s
}

func (s *AppConfigConfigMountDescs) Validate() error {
	if s.MountItems != nil {
		for _, item := range s.MountItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AppConfigConfigMountDescsMountItems struct {
	Key  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s AppConfigConfigMountDescsMountItems) String() string {
	return dara.Prettify(s)
}

func (s AppConfigConfigMountDescsMountItems) GoString() string {
	return s.String()
}

func (s *AppConfigConfigMountDescsMountItems) GetKey() *string {
	return s.Key
}

func (s *AppConfigConfigMountDescsMountItems) GetPath() *string {
	return s.Path
}

func (s *AppConfigConfigMountDescsMountItems) SetKey(v string) *AppConfigConfigMountDescsMountItems {
	s.Key = &v
	return s
}

func (s *AppConfigConfigMountDescsMountItems) SetPath(v string) *AppConfigConfigMountDescsMountItems {
	s.Path = &v
	return s
}

func (s *AppConfigConfigMountDescsMountItems) Validate() error {
	return dara.Validate(s)
}

type AppConfigEmptyDirs struct {
	MountPath   *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ReadOnly    *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	SubPathExpr *string `json:"SubPathExpr,omitempty" xml:"SubPathExpr,omitempty"`
}

func (s AppConfigEmptyDirs) String() string {
	return dara.Prettify(s)
}

func (s AppConfigEmptyDirs) GoString() string {
	return s.String()
}

func (s *AppConfigEmptyDirs) GetMountPath() *string {
	return s.MountPath
}

func (s *AppConfigEmptyDirs) GetName() *string {
	return s.Name
}

func (s *AppConfigEmptyDirs) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *AppConfigEmptyDirs) GetSubPathExpr() *string {
	return s.SubPathExpr
}

func (s *AppConfigEmptyDirs) SetMountPath(v string) *AppConfigEmptyDirs {
	s.MountPath = &v
	return s
}

func (s *AppConfigEmptyDirs) SetName(v string) *AppConfigEmptyDirs {
	s.Name = &v
	return s
}

func (s *AppConfigEmptyDirs) SetReadOnly(v bool) *AppConfigEmptyDirs {
	s.ReadOnly = &v
	return s
}

func (s *AppConfigEmptyDirs) SetSubPathExpr(v string) *AppConfigEmptyDirs {
	s.SubPathExpr = &v
	return s
}

func (s *AppConfigEmptyDirs) Validate() error {
	return dara.Validate(s)
}

type AppConfigEnvFroms struct {
	ConfigMapRef *string `json:"ConfigMapRef,omitempty" xml:"ConfigMapRef,omitempty"`
	SecretRef    *string `json:"SecretRef,omitempty" xml:"SecretRef,omitempty"`
}

func (s AppConfigEnvFroms) String() string {
	return dara.Prettify(s)
}

func (s AppConfigEnvFroms) GoString() string {
	return s.String()
}

func (s *AppConfigEnvFroms) GetConfigMapRef() *string {
	return s.ConfigMapRef
}

func (s *AppConfigEnvFroms) GetSecretRef() *string {
	return s.SecretRef
}

func (s *AppConfigEnvFroms) SetConfigMapRef(v string) *AppConfigEnvFroms {
	s.ConfigMapRef = &v
	return s
}

func (s *AppConfigEnvFroms) SetSecretRef(v string) *AppConfigEnvFroms {
	s.SecretRef = &v
	return s
}

func (s *AppConfigEnvFroms) Validate() error {
	return dara.Validate(s)
}

type AppConfigEnvs struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
	ValueFrom *string `json:"ValueFrom,omitempty" xml:"ValueFrom,omitempty"`
}

func (s AppConfigEnvs) String() string {
	return dara.Prettify(s)
}

func (s AppConfigEnvs) GoString() string {
	return s.String()
}

func (s *AppConfigEnvs) GetName() *string {
	return s.Name
}

func (s *AppConfigEnvs) GetValue() *string {
	return s.Value
}

func (s *AppConfigEnvs) GetValueFrom() *string {
	return s.ValueFrom
}

func (s *AppConfigEnvs) SetName(v string) *AppConfigEnvs {
	s.Name = &v
	return s
}

func (s *AppConfigEnvs) SetValue(v string) *AppConfigEnvs {
	s.Value = &v
	return s
}

func (s *AppConfigEnvs) SetValueFrom(v string) *AppConfigEnvs {
	s.ValueFrom = &v
	return s
}

func (s *AppConfigEnvs) Validate() error {
	return dara.Validate(s)
}

type AppConfigImageConfig struct {
	ContainerRegistryId *string `json:"ContainerRegistryId,omitempty" xml:"ContainerRegistryId,omitempty"`
	CrInstanceId        *string `json:"CrInstanceId,omitempty" xml:"CrInstanceId,omitempty"`
	CrRegionId          *string `json:"CrRegionId,omitempty" xml:"CrRegionId,omitempty"`
	ImageUrl            *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s AppConfigImageConfig) String() string {
	return dara.Prettify(s)
}

func (s AppConfigImageConfig) GoString() string {
	return s.String()
}

func (s *AppConfigImageConfig) GetContainerRegistryId() *string {
	return s.ContainerRegistryId
}

func (s *AppConfigImageConfig) GetCrInstanceId() *string {
	return s.CrInstanceId
}

func (s *AppConfigImageConfig) GetCrRegionId() *string {
	return s.CrRegionId
}

func (s *AppConfigImageConfig) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *AppConfigImageConfig) SetContainerRegistryId(v string) *AppConfigImageConfig {
	s.ContainerRegistryId = &v
	return s
}

func (s *AppConfigImageConfig) SetCrInstanceId(v string) *AppConfigImageConfig {
	s.CrInstanceId = &v
	return s
}

func (s *AppConfigImageConfig) SetCrRegionId(v string) *AppConfigImageConfig {
	s.CrRegionId = &v
	return s
}

func (s *AppConfigImageConfig) SetImageUrl(v string) *AppConfigImageConfig {
	s.ImageUrl = &v
	return s
}

func (s *AppConfigImageConfig) Validate() error {
	return dara.Validate(s)
}

type AppConfigLocalVolumes struct {
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NodePath  *string `json:"NodePath,omitempty" xml:"NodePath,omitempty"`
	OpsAuth   *int64  `json:"OpsAuth,omitempty" xml:"OpsAuth,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AppConfigLocalVolumes) String() string {
	return dara.Prettify(s)
}

func (s AppConfigLocalVolumes) GoString() string {
	return s.String()
}

func (s *AppConfigLocalVolumes) GetMountPath() *string {
	return s.MountPath
}

func (s *AppConfigLocalVolumes) GetName() *string {
	return s.Name
}

func (s *AppConfigLocalVolumes) GetNodePath() *string {
	return s.NodePath
}

func (s *AppConfigLocalVolumes) GetOpsAuth() *int64 {
	return s.OpsAuth
}

func (s *AppConfigLocalVolumes) GetType() *string {
	return s.Type
}

func (s *AppConfigLocalVolumes) SetMountPath(v string) *AppConfigLocalVolumes {
	s.MountPath = &v
	return s
}

func (s *AppConfigLocalVolumes) SetName(v string) *AppConfigLocalVolumes {
	s.Name = &v
	return s
}

func (s *AppConfigLocalVolumes) SetNodePath(v string) *AppConfigLocalVolumes {
	s.NodePath = &v
	return s
}

func (s *AppConfigLocalVolumes) SetOpsAuth(v int64) *AppConfigLocalVolumes {
	s.OpsAuth = &v
	return s
}

func (s *AppConfigLocalVolumes) SetType(v string) *AppConfigLocalVolumes {
	s.Type = &v
	return s
}

func (s *AppConfigLocalVolumes) Validate() error {
	return dara.Validate(s)
}

type AppConfigNasMountDescs struct {
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	NasPath   *string `json:"NasPath,omitempty" xml:"NasPath,omitempty"`
}

func (s AppConfigNasMountDescs) String() string {
	return dara.Prettify(s)
}

func (s AppConfigNasMountDescs) GoString() string {
	return s.String()
}

func (s *AppConfigNasMountDescs) GetMountPath() *string {
	return s.MountPath
}

func (s *AppConfigNasMountDescs) GetNasPath() *string {
	return s.NasPath
}

func (s *AppConfigNasMountDescs) SetMountPath(v string) *AppConfigNasMountDescs {
	s.MountPath = &v
	return s
}

func (s *AppConfigNasMountDescs) SetNasPath(v string) *AppConfigNasMountDescs {
	s.NasPath = &v
	return s
}

func (s *AppConfigNasMountDescs) Validate() error {
	return dara.Validate(s)
}

type AppConfigPackageConfig struct {
	EdasContainerVersion *string `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
	Jdk                  *string `json:"Jdk,omitempty" xml:"Jdk,omitempty"`
	PackageType          *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	PackageUrl           *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	PackageVersion       *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	Timezone             *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	UriEncoding          *string `json:"UriEncoding,omitempty" xml:"UriEncoding,omitempty"`
	UseBodyEncoding      *bool   `json:"UseBodyEncoding,omitempty" xml:"UseBodyEncoding,omitempty"`
	WebContainer         *string `json:"WebContainer,omitempty" xml:"WebContainer,omitempty"`
}

func (s AppConfigPackageConfig) String() string {
	return dara.Prettify(s)
}

func (s AppConfigPackageConfig) GoString() string {
	return s.String()
}

func (s *AppConfigPackageConfig) GetEdasContainerVersion() *string {
	return s.EdasContainerVersion
}

func (s *AppConfigPackageConfig) GetJdk() *string {
	return s.Jdk
}

func (s *AppConfigPackageConfig) GetPackageType() *string {
	return s.PackageType
}

func (s *AppConfigPackageConfig) GetPackageUrl() *string {
	return s.PackageUrl
}

func (s *AppConfigPackageConfig) GetPackageVersion() *string {
	return s.PackageVersion
}

func (s *AppConfigPackageConfig) GetTimezone() *string {
	return s.Timezone
}

func (s *AppConfigPackageConfig) GetUriEncoding() *string {
	return s.UriEncoding
}

func (s *AppConfigPackageConfig) GetUseBodyEncoding() *bool {
	return s.UseBodyEncoding
}

func (s *AppConfigPackageConfig) GetWebContainer() *string {
	return s.WebContainer
}

func (s *AppConfigPackageConfig) SetEdasContainerVersion(v string) *AppConfigPackageConfig {
	s.EdasContainerVersion = &v
	return s
}

func (s *AppConfigPackageConfig) SetJdk(v string) *AppConfigPackageConfig {
	s.Jdk = &v
	return s
}

func (s *AppConfigPackageConfig) SetPackageType(v string) *AppConfigPackageConfig {
	s.PackageType = &v
	return s
}

func (s *AppConfigPackageConfig) SetPackageUrl(v string) *AppConfigPackageConfig {
	s.PackageUrl = &v
	return s
}

func (s *AppConfigPackageConfig) SetPackageVersion(v string) *AppConfigPackageConfig {
	s.PackageVersion = &v
	return s
}

func (s *AppConfigPackageConfig) SetTimezone(v string) *AppConfigPackageConfig {
	s.Timezone = &v
	return s
}

func (s *AppConfigPackageConfig) SetUriEncoding(v string) *AppConfigPackageConfig {
	s.UriEncoding = &v
	return s
}

func (s *AppConfigPackageConfig) SetUseBodyEncoding(v bool) *AppConfigPackageConfig {
	s.UseBodyEncoding = &v
	return s
}

func (s *AppConfigPackageConfig) SetWebContainer(v string) *AppConfigPackageConfig {
	s.WebContainer = &v
	return s
}

func (s *AppConfigPackageConfig) Validate() error {
	return dara.Validate(s)
}

type AppConfigPvcMountDescs struct {
	MountPaths []*AppConfigPvcMountDescsMountPaths `json:"MountPaths,omitempty" xml:"MountPaths,omitempty" type:"Repeated"`
	PvcName    *string                             `json:"PvcName,omitempty" xml:"PvcName,omitempty"`
}

func (s AppConfigPvcMountDescs) String() string {
	return dara.Prettify(s)
}

func (s AppConfigPvcMountDescs) GoString() string {
	return s.String()
}

func (s *AppConfigPvcMountDescs) GetMountPaths() []*AppConfigPvcMountDescsMountPaths {
	return s.MountPaths
}

func (s *AppConfigPvcMountDescs) GetPvcName() *string {
	return s.PvcName
}

func (s *AppConfigPvcMountDescs) SetMountPaths(v []*AppConfigPvcMountDescsMountPaths) *AppConfigPvcMountDescs {
	s.MountPaths = v
	return s
}

func (s *AppConfigPvcMountDescs) SetPvcName(v string) *AppConfigPvcMountDescs {
	s.PvcName = &v
	return s
}

func (s *AppConfigPvcMountDescs) Validate() error {
	if s.MountPaths != nil {
		for _, item := range s.MountPaths {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AppConfigPvcMountDescsMountPaths struct {
	MountPath   *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	ReadOnly    *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	SubPathExpr *string `json:"SubPathExpr,omitempty" xml:"SubPathExpr,omitempty"`
}

func (s AppConfigPvcMountDescsMountPaths) String() string {
	return dara.Prettify(s)
}

func (s AppConfigPvcMountDescsMountPaths) GoString() string {
	return s.String()
}

func (s *AppConfigPvcMountDescsMountPaths) GetMountPath() *string {
	return s.MountPath
}

func (s *AppConfigPvcMountDescsMountPaths) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *AppConfigPvcMountDescsMountPaths) GetSubPathExpr() *string {
	return s.SubPathExpr
}

func (s *AppConfigPvcMountDescsMountPaths) SetMountPath(v string) *AppConfigPvcMountDescsMountPaths {
	s.MountPath = &v
	return s
}

func (s *AppConfigPvcMountDescsMountPaths) SetReadOnly(v bool) *AppConfigPvcMountDescsMountPaths {
	s.ReadOnly = &v
	return s
}

func (s *AppConfigPvcMountDescsMountPaths) SetSubPathExpr(v string) *AppConfigPvcMountDescsMountPaths {
	s.SubPathExpr = &v
	return s
}

func (s *AppConfigPvcMountDescsMountPaths) Validate() error {
	return dara.Validate(s)
}

type AppConfigSlsConfigs struct {
	LogDir   *string `json:"LogDir,omitempty" xml:"LogDir,omitempty"`
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AppConfigSlsConfigs) String() string {
	return dara.Prettify(s)
}

func (s AppConfigSlsConfigs) GoString() string {
	return s.String()
}

func (s *AppConfigSlsConfigs) GetLogDir() *string {
	return s.LogDir
}

func (s *AppConfigSlsConfigs) GetLogstore() *string {
	return s.Logstore
}

func (s *AppConfigSlsConfigs) GetProject() *string {
	return s.Project
}

func (s *AppConfigSlsConfigs) GetType() *string {
	return s.Type
}

func (s *AppConfigSlsConfigs) SetLogDir(v string) *AppConfigSlsConfigs {
	s.LogDir = &v
	return s
}

func (s *AppConfigSlsConfigs) SetLogstore(v string) *AppConfigSlsConfigs {
	s.Logstore = &v
	return s
}

func (s *AppConfigSlsConfigs) SetProject(v string) *AppConfigSlsConfigs {
	s.Project = &v
	return s
}

func (s *AppConfigSlsConfigs) SetType(v string) *AppConfigSlsConfigs {
	s.Type = &v
	return s
}

func (s *AppConfigSlsConfigs) Validate() error {
	return dara.Validate(s)
}

type AppConfigWebContainerConfig struct {
	ConnectorType        *string `json:"ConnectorType,omitempty" xml:"ConnectorType,omitempty"`
	ContextInputType     *string `json:"ContextInputType,omitempty" xml:"ContextInputType,omitempty"`
	ContextPath          *string `json:"ContextPath,omitempty" xml:"ContextPath,omitempty"`
	HttpPort             *int64  `json:"HttpPort,omitempty" xml:"HttpPort,omitempty"`
	MaxThreads           *int64  `json:"MaxThreads,omitempty" xml:"MaxThreads,omitempty"`
	ServerXml            *string `json:"ServerXml,omitempty" xml:"ServerXml,omitempty"`
	UriEncoding          *string `json:"UriEncoding,omitempty" xml:"UriEncoding,omitempty"`
	UseAdvancedServerXml *bool   `json:"UseAdvancedServerXml,omitempty" xml:"UseAdvancedServerXml,omitempty"`
	UseBodyEncoding      *bool   `json:"UseBodyEncoding,omitempty" xml:"UseBodyEncoding,omitempty"`
	UseDefaultConfig     *bool   `json:"UseDefaultConfig,omitempty" xml:"UseDefaultConfig,omitempty"`
}

func (s AppConfigWebContainerConfig) String() string {
	return dara.Prettify(s)
}

func (s AppConfigWebContainerConfig) GoString() string {
	return s.String()
}

func (s *AppConfigWebContainerConfig) GetConnectorType() *string {
	return s.ConnectorType
}

func (s *AppConfigWebContainerConfig) GetContextInputType() *string {
	return s.ContextInputType
}

func (s *AppConfigWebContainerConfig) GetContextPath() *string {
	return s.ContextPath
}

func (s *AppConfigWebContainerConfig) GetHttpPort() *int64 {
	return s.HttpPort
}

func (s *AppConfigWebContainerConfig) GetMaxThreads() *int64 {
	return s.MaxThreads
}

func (s *AppConfigWebContainerConfig) GetServerXml() *string {
	return s.ServerXml
}

func (s *AppConfigWebContainerConfig) GetUriEncoding() *string {
	return s.UriEncoding
}

func (s *AppConfigWebContainerConfig) GetUseAdvancedServerXml() *bool {
	return s.UseAdvancedServerXml
}

func (s *AppConfigWebContainerConfig) GetUseBodyEncoding() *bool {
	return s.UseBodyEncoding
}

func (s *AppConfigWebContainerConfig) GetUseDefaultConfig() *bool {
	return s.UseDefaultConfig
}

func (s *AppConfigWebContainerConfig) SetConnectorType(v string) *AppConfigWebContainerConfig {
	s.ConnectorType = &v
	return s
}

func (s *AppConfigWebContainerConfig) SetContextInputType(v string) *AppConfigWebContainerConfig {
	s.ContextInputType = &v
	return s
}

func (s *AppConfigWebContainerConfig) SetContextPath(v string) *AppConfigWebContainerConfig {
	s.ContextPath = &v
	return s
}

func (s *AppConfigWebContainerConfig) SetHttpPort(v int64) *AppConfigWebContainerConfig {
	s.HttpPort = &v
	return s
}

func (s *AppConfigWebContainerConfig) SetMaxThreads(v int64) *AppConfigWebContainerConfig {
	s.MaxThreads = &v
	return s
}

func (s *AppConfigWebContainerConfig) SetServerXml(v string) *AppConfigWebContainerConfig {
	s.ServerXml = &v
	return s
}

func (s *AppConfigWebContainerConfig) SetUriEncoding(v string) *AppConfigWebContainerConfig {
	s.UriEncoding = &v
	return s
}

func (s *AppConfigWebContainerConfig) SetUseAdvancedServerXml(v bool) *AppConfigWebContainerConfig {
	s.UseAdvancedServerXml = &v
	return s
}

func (s *AppConfigWebContainerConfig) SetUseBodyEncoding(v bool) *AppConfigWebContainerConfig {
	s.UseBodyEncoding = &v
	return s
}

func (s *AppConfigWebContainerConfig) SetUseDefaultConfig(v bool) *AppConfigWebContainerConfig {
	s.UseDefaultConfig = &v
	return s
}

func (s *AppConfigWebContainerConfig) Validate() error {
	return dara.Validate(s)
}
