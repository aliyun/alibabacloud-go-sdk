// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComputeClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterConfig(v *CreateComputeClusterRequestClusterConfig) *CreateComputeClusterRequest
	GetClusterConfig() *CreateComputeClusterRequestClusterConfig
	SetOpTenantId(v int64) *CreateComputeClusterRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *CreateComputeClusterRequest
	GetOpUserId() *string
}

type CreateComputeClusterRequest struct {
	// The cluster configuration.
	//
	// This parameter is required.
	ClusterConfig *CreateComputeClusterRequestClusterConfig `json:"ClusterConfig,omitempty" xml:"ClusterConfig,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s CreateComputeClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateComputeClusterRequest) GetClusterConfig() *CreateComputeClusterRequestClusterConfig {
	return s.ClusterConfig
}

func (s *CreateComputeClusterRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateComputeClusterRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CreateComputeClusterRequest) SetClusterConfig(v *CreateComputeClusterRequestClusterConfig) *CreateComputeClusterRequest {
	s.ClusterConfig = v
	return s
}

func (s *CreateComputeClusterRequest) SetOpTenantId(v int64) *CreateComputeClusterRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateComputeClusterRequest) SetOpUserId(v string) *CreateComputeClusterRequest {
	s.OpUserId = &v
	return s
}

func (s *CreateComputeClusterRequest) Validate() error {
	if s.ClusterConfig != nil {
		if err := s.ClusterConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateComputeClusterRequestClusterConfig struct {
	// The list of cluster administrator IDs.
	ClusterAdmins []*string `json:"ClusterAdmins,omitempty" xml:"ClusterAdmins,omitempty" type:"Repeated"`
	// The cluster security control configuration.
	ClusterSafetyControl *CreateComputeClusterRequestClusterConfigClusterSafetyControl `json:"ClusterSafetyControl,omitempty" xml:"ClusterSafetyControl,omitempty" type:"Struct"`
	// The connection configuration items.
	//
	// This parameter is required.
	ConfigList []*CreateComputeClusterRequestClusterConfigConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
	// The cluster description.
	//
	// example:
	//
	// test
	Des *string `json:"Des,omitempty" xml:"Des,omitempty"`
	// The cluster name.
	//
	// This parameter is required.
	//
	// example:
	//
	// cluster_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The cluster type.
	//
	// This parameter is required.
	//
	// example:
	//
	// MAX_COMPUTE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The cluster version.
	//
	// example:
	//
	// CDH6
	TypeVersion *string `json:"TypeVersion,omitempty" xml:"TypeVersion,omitempty"`
}

func (s CreateComputeClusterRequestClusterConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeClusterRequestClusterConfig) GoString() string {
	return s.String()
}

func (s *CreateComputeClusterRequestClusterConfig) GetClusterAdmins() []*string {
	return s.ClusterAdmins
}

func (s *CreateComputeClusterRequestClusterConfig) GetClusterSafetyControl() *CreateComputeClusterRequestClusterConfigClusterSafetyControl {
	return s.ClusterSafetyControl
}

func (s *CreateComputeClusterRequestClusterConfig) GetConfigList() []*CreateComputeClusterRequestClusterConfigConfigList {
	return s.ConfigList
}

func (s *CreateComputeClusterRequestClusterConfig) GetDes() *string {
	return s.Des
}

func (s *CreateComputeClusterRequestClusterConfig) GetName() *string {
	return s.Name
}

func (s *CreateComputeClusterRequestClusterConfig) GetType() *string {
	return s.Type
}

func (s *CreateComputeClusterRequestClusterConfig) GetTypeVersion() *string {
	return s.TypeVersion
}

func (s *CreateComputeClusterRequestClusterConfig) SetClusterAdmins(v []*string) *CreateComputeClusterRequestClusterConfig {
	s.ClusterAdmins = v
	return s
}

func (s *CreateComputeClusterRequestClusterConfig) SetClusterSafetyControl(v *CreateComputeClusterRequestClusterConfigClusterSafetyControl) *CreateComputeClusterRequestClusterConfig {
	s.ClusterSafetyControl = v
	return s
}

func (s *CreateComputeClusterRequestClusterConfig) SetConfigList(v []*CreateComputeClusterRequestClusterConfigConfigList) *CreateComputeClusterRequestClusterConfig {
	s.ConfigList = v
	return s
}

func (s *CreateComputeClusterRequestClusterConfig) SetDes(v string) *CreateComputeClusterRequestClusterConfig {
	s.Des = &v
	return s
}

func (s *CreateComputeClusterRequestClusterConfig) SetName(v string) *CreateComputeClusterRequestClusterConfig {
	s.Name = &v
	return s
}

func (s *CreateComputeClusterRequestClusterConfig) SetType(v string) *CreateComputeClusterRequestClusterConfig {
	s.Type = &v
	return s
}

func (s *CreateComputeClusterRequestClusterConfig) SetTypeVersion(v string) *CreateComputeClusterRequestClusterConfig {
	s.TypeVersion = &v
	return s
}

func (s *CreateComputeClusterRequestClusterConfig) Validate() error {
	if s.ClusterSafetyControl != nil {
		if err := s.ClusterSafetyControl.Validate(); err != nil {
			return err
		}
	}
	if s.ConfigList != nil {
		for _, item := range s.ConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateComputeClusterRequestClusterConfigClusterSafetyControl struct {
	// The control mode.
	//
	// example:
	//
	// CREATE_COMPUTE_SOURCE
	ClusterSafetyAuthType *string `json:"ClusterSafetyAuthType,omitempty" xml:"ClusterSafetyAuthType,omitempty"`
	// The list of whitelist user group IDs.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The list of whitelist user IDs.
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s CreateComputeClusterRequestClusterConfigClusterSafetyControl) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeClusterRequestClusterConfigClusterSafetyControl) GoString() string {
	return s.String()
}

func (s *CreateComputeClusterRequestClusterConfigClusterSafetyControl) GetClusterSafetyAuthType() *string {
	return s.ClusterSafetyAuthType
}

func (s *CreateComputeClusterRequestClusterConfigClusterSafetyControl) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *CreateComputeClusterRequestClusterConfigClusterSafetyControl) GetUserIds() []*string {
	return s.UserIds
}

func (s *CreateComputeClusterRequestClusterConfigClusterSafetyControl) SetClusterSafetyAuthType(v string) *CreateComputeClusterRequestClusterConfigClusterSafetyControl {
	s.ClusterSafetyAuthType = &v
	return s
}

func (s *CreateComputeClusterRequestClusterConfigClusterSafetyControl) SetUserGroupIds(v []*string) *CreateComputeClusterRequestClusterConfigClusterSafetyControl {
	s.UserGroupIds = v
	return s
}

func (s *CreateComputeClusterRequestClusterConfigClusterSafetyControl) SetUserIds(v []*string) *CreateComputeClusterRequestClusterConfigClusterSafetyControl {
	s.UserIds = v
	return s
}

func (s *CreateComputeClusterRequestClusterConfigClusterSafetyControl) Validate() error {
	return dara.Validate(s)
}

type CreateComputeClusterRequestClusterConfigConfigList struct {
	// The configuration item.
	//
	// This parameter is required.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the configuration item.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateComputeClusterRequestClusterConfigConfigList) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeClusterRequestClusterConfigConfigList) GoString() string {
	return s.String()
}

func (s *CreateComputeClusterRequestClusterConfigConfigList) GetKey() *string {
	return s.Key
}

func (s *CreateComputeClusterRequestClusterConfigConfigList) GetValue() *string {
	return s.Value
}

func (s *CreateComputeClusterRequestClusterConfigConfigList) SetKey(v string) *CreateComputeClusterRequestClusterConfigConfigList {
	s.Key = &v
	return s
}

func (s *CreateComputeClusterRequestClusterConfigConfigList) SetValue(v string) *CreateComputeClusterRequestClusterConfigConfigList {
	s.Value = &v
	return s
}

func (s *CreateComputeClusterRequestClusterConfigConfigList) Validate() error {
	return dara.Validate(s)
}
