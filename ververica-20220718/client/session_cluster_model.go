// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSessionCluster interface {
	dara.Model
	String() string
	GoString() string
	SetBasicResourceSetting(v *BasicResourceSetting) *SessionCluster
	GetBasicResourceSetting() *BasicResourceSetting
	SetCreatedAt(v int64) *SessionCluster
	GetCreatedAt() *int64
	SetCreator(v string) *SessionCluster
	GetCreator() *string
	SetCreatorName(v string) *SessionCluster
	GetCreatorName() *string
	SetDeploymentTargetName(v string) *SessionCluster
	GetDeploymentTargetName() *string
	SetEngineVersion(v string) *SessionCluster
	GetEngineVersion() *string
	SetFlinkConf(v map[string]interface{}) *SessionCluster
	GetFlinkConf() map[string]interface{}
	SetLabels(v map[string]interface{}) *SessionCluster
	GetLabels() map[string]interface{}
	SetLogging(v *Logging) *SessionCluster
	GetLogging() *Logging
	SetModifiedAt(v int64) *SessionCluster
	GetModifiedAt() *int64
	SetModifier(v string) *SessionCluster
	GetModifier() *string
	SetModifierName(v string) *SessionCluster
	GetModifierName() *string
	SetName(v string) *SessionCluster
	GetName() *string
	SetNamespace(v string) *SessionCluster
	GetNamespace() *string
	SetSessionClusterId(v string) *SessionCluster
	GetSessionClusterId() *string
	SetStatus(v *SessionClusterStatus) *SessionCluster
	GetStatus() *SessionClusterStatus
	SetWorkspace(v string) *SessionCluster
	GetWorkspace() *string
}

type SessionCluster struct {
	BasicResourceSetting *BasicResourceSetting `json:"basicResourceSetting,omitempty" xml:"basicResourceSetting,omitempty"`
	CreatedAt            *int64                `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 27846363877456****
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// ****@streamcompute.onaliyun.com
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// example:
	//
	// default-queue
	DeploymentTargetName *string `json:"deploymentTargetName,omitempty" xml:"deploymentTargetName,omitempty"`
	// example:
	//
	// vvr-6.0.7-flink-1.15
	EngineVersion *string `json:"engineVersion,omitempty" xml:"engineVersion,omitempty"`
	// example:
	//
	// {"taskmanager.numberOfTaskSlots":"1"}
	FlinkConf  map[string]interface{} `json:"flinkConf,omitempty" xml:"flinkConf,omitempty"`
	Labels     map[string]interface{} `json:"labels,omitempty" xml:"labels,omitempty"`
	Logging    *Logging               `json:"logging,omitempty" xml:"logging,omitempty"`
	ModifiedAt *int64                 `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
	// example:
	//
	// 27846363877456****
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// example:
	//
	// ****@streamcompute.onaliyun.com
	ModifierName *string `json:"modifierName,omitempty" xml:"modifierName,omitempty"`
	// example:
	//
	// test-sessionCluster
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// default-namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// 1f68a52c-1760-43c6-97fb-afe0674b****
	SessionClusterId *string               `json:"sessionClusterId,omitempty" xml:"sessionClusterId,omitempty"`
	Status           *SessionClusterStatus `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// edcef******b4f
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s SessionCluster) String() string {
	return dara.Prettify(s)
}

func (s SessionCluster) GoString() string {
	return s.String()
}

func (s *SessionCluster) GetBasicResourceSetting() *BasicResourceSetting {
	return s.BasicResourceSetting
}

func (s *SessionCluster) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *SessionCluster) GetCreator() *string {
	return s.Creator
}

func (s *SessionCluster) GetCreatorName() *string {
	return s.CreatorName
}

func (s *SessionCluster) GetDeploymentTargetName() *string {
	return s.DeploymentTargetName
}

func (s *SessionCluster) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *SessionCluster) GetFlinkConf() map[string]interface{} {
	return s.FlinkConf
}

func (s *SessionCluster) GetLabels() map[string]interface{} {
	return s.Labels
}

func (s *SessionCluster) GetLogging() *Logging {
	return s.Logging
}

func (s *SessionCluster) GetModifiedAt() *int64 {
	return s.ModifiedAt
}

func (s *SessionCluster) GetModifier() *string {
	return s.Modifier
}

func (s *SessionCluster) GetModifierName() *string {
	return s.ModifierName
}

func (s *SessionCluster) GetName() *string {
	return s.Name
}

func (s *SessionCluster) GetNamespace() *string {
	return s.Namespace
}

func (s *SessionCluster) GetSessionClusterId() *string {
	return s.SessionClusterId
}

func (s *SessionCluster) GetStatus() *SessionClusterStatus {
	return s.Status
}

func (s *SessionCluster) GetWorkspace() *string {
	return s.Workspace
}

func (s *SessionCluster) SetBasicResourceSetting(v *BasicResourceSetting) *SessionCluster {
	s.BasicResourceSetting = v
	return s
}

func (s *SessionCluster) SetCreatedAt(v int64) *SessionCluster {
	s.CreatedAt = &v
	return s
}

func (s *SessionCluster) SetCreator(v string) *SessionCluster {
	s.Creator = &v
	return s
}

func (s *SessionCluster) SetCreatorName(v string) *SessionCluster {
	s.CreatorName = &v
	return s
}

func (s *SessionCluster) SetDeploymentTargetName(v string) *SessionCluster {
	s.DeploymentTargetName = &v
	return s
}

func (s *SessionCluster) SetEngineVersion(v string) *SessionCluster {
	s.EngineVersion = &v
	return s
}

func (s *SessionCluster) SetFlinkConf(v map[string]interface{}) *SessionCluster {
	s.FlinkConf = v
	return s
}

func (s *SessionCluster) SetLabels(v map[string]interface{}) *SessionCluster {
	s.Labels = v
	return s
}

func (s *SessionCluster) SetLogging(v *Logging) *SessionCluster {
	s.Logging = v
	return s
}

func (s *SessionCluster) SetModifiedAt(v int64) *SessionCluster {
	s.ModifiedAt = &v
	return s
}

func (s *SessionCluster) SetModifier(v string) *SessionCluster {
	s.Modifier = &v
	return s
}

func (s *SessionCluster) SetModifierName(v string) *SessionCluster {
	s.ModifierName = &v
	return s
}

func (s *SessionCluster) SetName(v string) *SessionCluster {
	s.Name = &v
	return s
}

func (s *SessionCluster) SetNamespace(v string) *SessionCluster {
	s.Namespace = &v
	return s
}

func (s *SessionCluster) SetSessionClusterId(v string) *SessionCluster {
	s.SessionClusterId = &v
	return s
}

func (s *SessionCluster) SetStatus(v *SessionClusterStatus) *SessionCluster {
	s.Status = v
	return s
}

func (s *SessionCluster) SetWorkspace(v string) *SessionCluster {
	s.Workspace = &v
	return s
}

func (s *SessionCluster) Validate() error {
	if s.BasicResourceSetting != nil {
		if err := s.BasicResourceSetting.Validate(); err != nil {
			return err
		}
	}
	if s.Logging != nil {
		if err := s.Logging.Validate(); err != nil {
			return err
		}
	}
	if s.Status != nil {
		if err := s.Status.Validate(); err != nil {
			return err
		}
	}
	return nil
}
