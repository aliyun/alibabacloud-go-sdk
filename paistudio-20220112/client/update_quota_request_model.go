// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterSpec(v *UpdateQuotaRequestClusterSpec) *UpdateQuotaRequest
	GetClusterSpec() *UpdateQuotaRequestClusterSpec
	SetDescription(v string) *UpdateQuotaRequest
	GetDescription() *string
	SetEnableGracefulUpgrade(v bool) *UpdateQuotaRequest
	GetEnableGracefulUpgrade() *bool
	SetGracefulUpgradeConfig(v *UpdateQuotaRequestGracefulUpgradeConfig) *UpdateQuotaRequest
	GetGracefulUpgradeConfig() *UpdateQuotaRequestGracefulUpgradeConfig
	SetLabels(v []*Label) *UpdateQuotaRequest
	GetLabels() []*Label
	SetPropagateDefaultGPUDriver(v bool) *UpdateQuotaRequest
	GetPropagateDefaultGPUDriver() *bool
	SetQueueStrategy(v string) *UpdateQuotaRequest
	GetQueueStrategy() *string
	SetQuotaConfig(v *QuotaConfig) *UpdateQuotaRequest
	GetQuotaConfig() *QuotaConfig
	SetQuotaName(v string) *UpdateQuotaRequest
	GetQuotaName() *string
}

type UpdateQuotaRequest struct {
	// The configurations of the compute cluster bound to the quota (RNQ Ray cluster). This parameter is only available for RNQ quotas and will be rejected for common quotas. It can only be modified when no operations are in progress on the quota. Unpassed sub-fields retain the current effective configurations.
	ClusterSpec *UpdateQuotaRequestClusterSpec `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty" type:"Struct"`
	// The description of the resource quota.
	//
	// example:
	//
	// this is a test quota
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to execute the current ClusterSpec change as a rolling upgrade (draining and migrating nodes one by one to the new cluster) instead of the default stop-and-rebuild method.
	//
	// example:
	//
	// true
	EnableGracefulUpgrade *bool `json:"EnableGracefulUpgrade,omitempty" xml:"EnableGracefulUpgrade,omitempty"`
	// The tuning parameters for the migration process. This parameter is only allowed when EnableGracefulUpgrade is set to true.
	GracefulUpgradeConfig *UpdateQuotaRequestGracefulUpgradeConfig `json:"GracefulUpgradeConfig,omitempty" xml:"GracefulUpgradeConfig,omitempty" type:"Struct"`
	// The list of user-defined labels.
	//
	// This is a full update.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// Specifies whether to pass the default GPU driver configuration to child quotas.
	PropagateDefaultGPUDriver *bool `json:"PropagateDefaultGPUDriver,omitempty" xml:"PropagateDefaultGPUDriver,omitempty"`
	// The queuing strategy for tasks on the quota.
	//
	// example:
	//
	// PaiStrategyIntelligent
	QueueStrategy *string `json:"QueueStrategy,omitempty" xml:"QueueStrategy,omitempty"`
	// The resource quota configuration.
	//
	// if can be null:
	// true
	QuotaConfig *QuotaConfig `json:"QuotaConfig,omitempty" xml:"QuotaConfig,omitempty"`
	// The name of the resource quota.
	//
	// example:
	//
	// test
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
}

func (s UpdateQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateQuotaRequest) GoString() string {
	return s.String()
}

func (s *UpdateQuotaRequest) GetClusterSpec() *UpdateQuotaRequestClusterSpec {
	return s.ClusterSpec
}

func (s *UpdateQuotaRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateQuotaRequest) GetEnableGracefulUpgrade() *bool {
	return s.EnableGracefulUpgrade
}

func (s *UpdateQuotaRequest) GetGracefulUpgradeConfig() *UpdateQuotaRequestGracefulUpgradeConfig {
	return s.GracefulUpgradeConfig
}

func (s *UpdateQuotaRequest) GetLabels() []*Label {
	return s.Labels
}

func (s *UpdateQuotaRequest) GetPropagateDefaultGPUDriver() *bool {
	return s.PropagateDefaultGPUDriver
}

func (s *UpdateQuotaRequest) GetQueueStrategy() *string {
	return s.QueueStrategy
}

func (s *UpdateQuotaRequest) GetQuotaConfig() *QuotaConfig {
	return s.QuotaConfig
}

func (s *UpdateQuotaRequest) GetQuotaName() *string {
	return s.QuotaName
}

func (s *UpdateQuotaRequest) SetClusterSpec(v *UpdateQuotaRequestClusterSpec) *UpdateQuotaRequest {
	s.ClusterSpec = v
	return s
}

func (s *UpdateQuotaRequest) SetDescription(v string) *UpdateQuotaRequest {
	s.Description = &v
	return s
}

func (s *UpdateQuotaRequest) SetEnableGracefulUpgrade(v bool) *UpdateQuotaRequest {
	s.EnableGracefulUpgrade = &v
	return s
}

func (s *UpdateQuotaRequest) SetGracefulUpgradeConfig(v *UpdateQuotaRequestGracefulUpgradeConfig) *UpdateQuotaRequest {
	s.GracefulUpgradeConfig = v
	return s
}

func (s *UpdateQuotaRequest) SetLabels(v []*Label) *UpdateQuotaRequest {
	s.Labels = v
	return s
}

func (s *UpdateQuotaRequest) SetPropagateDefaultGPUDriver(v bool) *UpdateQuotaRequest {
	s.PropagateDefaultGPUDriver = &v
	return s
}

func (s *UpdateQuotaRequest) SetQueueStrategy(v string) *UpdateQuotaRequest {
	s.QueueStrategy = &v
	return s
}

func (s *UpdateQuotaRequest) SetQuotaConfig(v *QuotaConfig) *UpdateQuotaRequest {
	s.QuotaConfig = v
	return s
}

func (s *UpdateQuotaRequest) SetQuotaName(v string) *UpdateQuotaRequest {
	s.QuotaName = &v
	return s
}

func (s *UpdateQuotaRequest) Validate() error {
	if s.ClusterSpec != nil {
		if err := s.ClusterSpec.Validate(); err != nil {
			return err
		}
	}
	if s.GracefulUpgradeConfig != nil {
		if err := s.GracefulUpgradeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.QuotaConfig != nil {
		if err := s.QuotaConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateQuotaRequestClusterSpec struct {
	// The cluster type. Currently, only RayCluster is supported.
	//
	// example:
	//
	// RayCluster
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The list of data sources mounted to the cluster. This parameter uses the overall replacement semantics: if it is not passed, the current state is maintained; if an empty array is passed, the list is cleared. Each element is a JSON object string that contains DataSourceId or Uri (one of them is required and they are mutually exclusive) and an optional MountPath.
	DataSources []*string `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// The container image used by the Ray cluster. If this parameter is explicitly passed, the system switches to this image. If it is left empty and the data sources remain unchanged, and GetQuota returns UpgradeAvailable=true, this update is considered an upgrade to the platform default image. To keep the current image, explicitly pass the current value.
	//
	// example:
	//
	// dsw-registry-vpc.ap-southeast-1.cr.aliyuncs.com/pai-common/ray-runtime
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
}

func (s UpdateQuotaRequestClusterSpec) String() string {
	return dara.Prettify(s)
}

func (s UpdateQuotaRequestClusterSpec) GoString() string {
	return s.String()
}

func (s *UpdateQuotaRequestClusterSpec) GetClusterType() *string {
	return s.ClusterType
}

func (s *UpdateQuotaRequestClusterSpec) GetDataSources() []*string {
	return s.DataSources
}

func (s *UpdateQuotaRequestClusterSpec) GetImage() *string {
	return s.Image
}

func (s *UpdateQuotaRequestClusterSpec) SetClusterType(v string) *UpdateQuotaRequestClusterSpec {
	s.ClusterType = &v
	return s
}

func (s *UpdateQuotaRequestClusterSpec) SetDataSources(v []*string) *UpdateQuotaRequestClusterSpec {
	s.DataSources = v
	return s
}

func (s *UpdateQuotaRequestClusterSpec) SetImage(v string) *UpdateQuotaRequestClusterSpec {
	s.Image = &v
	return s
}

func (s *UpdateQuotaRequestClusterSpec) Validate() error {
	return dara.Validate(s)
}

type UpdateQuotaRequestGracefulUpgradeConfig struct {
	// The timeout limit for Ray draining of a single node. After the timeout, the node is still directly scaled down, which does not cause the upgrade to fail.
	//
	// example:
	//
	// 3600
	DrainTimeoutSeconds *int64 `json:"DrainTimeoutSeconds,omitempty" xml:"DrainTimeoutSeconds,omitempty"`
	// The migration window width, which is the maximum number of workers in the old cluster that are simultaneously in an unavailable state (draining or drained and pending removal), that is, the number of nodes migrated in parallel.
	//
	// example:
	//
	// 25%
	MaxUnavailable *string `json:"MaxUnavailable,omitempty" xml:"MaxUnavailable,omitempty"`
	// The number of workers in the old cluster that are kept in a schedulable state while waiting for the remaining Ray tasks in the old cluster to complete.
	//
	// example:
	//
	// 5
	TaskDrainReserveWorkers *int64 `json:"TaskDrainReserveWorkers,omitempty" xml:"TaskDrainReserveWorkers,omitempty"`
}

func (s UpdateQuotaRequestGracefulUpgradeConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateQuotaRequestGracefulUpgradeConfig) GoString() string {
	return s.String()
}

func (s *UpdateQuotaRequestGracefulUpgradeConfig) GetDrainTimeoutSeconds() *int64 {
	return s.DrainTimeoutSeconds
}

func (s *UpdateQuotaRequestGracefulUpgradeConfig) GetMaxUnavailable() *string {
	return s.MaxUnavailable
}

func (s *UpdateQuotaRequestGracefulUpgradeConfig) GetTaskDrainReserveWorkers() *int64 {
	return s.TaskDrainReserveWorkers
}

func (s *UpdateQuotaRequestGracefulUpgradeConfig) SetDrainTimeoutSeconds(v int64) *UpdateQuotaRequestGracefulUpgradeConfig {
	s.DrainTimeoutSeconds = &v
	return s
}

func (s *UpdateQuotaRequestGracefulUpgradeConfig) SetMaxUnavailable(v string) *UpdateQuotaRequestGracefulUpgradeConfig {
	s.MaxUnavailable = &v
	return s
}

func (s *UpdateQuotaRequestGracefulUpgradeConfig) SetTaskDrainReserveWorkers(v int64) *UpdateQuotaRequestGracefulUpgradeConfig {
	s.TaskDrainReserveWorkers = &v
	return s
}

func (s *UpdateQuotaRequestGracefulUpgradeConfig) Validate() error {
	return dara.Validate(s)
}
