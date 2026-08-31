// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterConfig(v *UpdateComputeClusterRequestClusterConfig) *UpdateComputeClusterRequest
	GetClusterConfig() *UpdateComputeClusterRequestClusterConfig
	SetId(v int64) *UpdateComputeClusterRequest
	GetId() *int64
	SetOpTenantId(v int64) *UpdateComputeClusterRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UpdateComputeClusterRequest
	GetOpUserId() *string
}

type UpdateComputeClusterRequest struct {
	// This parameter is required.
	ClusterConfig *UpdateComputeClusterRequestClusterConfig `json:"ClusterConfig,omitempty" xml:"ClusterConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 102311
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s UpdateComputeClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeClusterRequest) GoString() string {
	return s.String()
}

func (s *UpdateComputeClusterRequest) GetClusterConfig() *UpdateComputeClusterRequestClusterConfig {
	return s.ClusterConfig
}

func (s *UpdateComputeClusterRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateComputeClusterRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateComputeClusterRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpdateComputeClusterRequest) SetClusterConfig(v *UpdateComputeClusterRequestClusterConfig) *UpdateComputeClusterRequest {
	s.ClusterConfig = v
	return s
}

func (s *UpdateComputeClusterRequest) SetId(v int64) *UpdateComputeClusterRequest {
	s.Id = &v
	return s
}

func (s *UpdateComputeClusterRequest) SetOpTenantId(v int64) *UpdateComputeClusterRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateComputeClusterRequest) SetOpUserId(v string) *UpdateComputeClusterRequest {
	s.OpUserId = &v
	return s
}

func (s *UpdateComputeClusterRequest) Validate() error {
	if s.ClusterConfig != nil {
		if err := s.ClusterConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateComputeClusterRequestClusterConfig struct {
	ClusterAdmins        []*string                                                     `json:"ClusterAdmins,omitempty" xml:"ClusterAdmins,omitempty" type:"Repeated"`
	ClusterSafetyControl *UpdateComputeClusterRequestClusterConfigClusterSafetyControl `json:"ClusterSafetyControl,omitempty" xml:"ClusterSafetyControl,omitempty" type:"Struct"`
	// This parameter is required.
	ConfigList []*UpdateComputeClusterRequestClusterConfigConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Des *string `json:"Des,omitempty" xml:"Des,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cluster_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MAX_COMPUTE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// CDH6
	TypeVersion *string `json:"TypeVersion,omitempty" xml:"TypeVersion,omitempty"`
}

func (s UpdateComputeClusterRequestClusterConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeClusterRequestClusterConfig) GoString() string {
	return s.String()
}

func (s *UpdateComputeClusterRequestClusterConfig) GetClusterAdmins() []*string {
	return s.ClusterAdmins
}

func (s *UpdateComputeClusterRequestClusterConfig) GetClusterSafetyControl() *UpdateComputeClusterRequestClusterConfigClusterSafetyControl {
	return s.ClusterSafetyControl
}

func (s *UpdateComputeClusterRequestClusterConfig) GetConfigList() []*UpdateComputeClusterRequestClusterConfigConfigList {
	return s.ConfigList
}

func (s *UpdateComputeClusterRequestClusterConfig) GetDes() *string {
	return s.Des
}

func (s *UpdateComputeClusterRequestClusterConfig) GetName() *string {
	return s.Name
}

func (s *UpdateComputeClusterRequestClusterConfig) GetType() *string {
	return s.Type
}

func (s *UpdateComputeClusterRequestClusterConfig) GetTypeVersion() *string {
	return s.TypeVersion
}

func (s *UpdateComputeClusterRequestClusterConfig) SetClusterAdmins(v []*string) *UpdateComputeClusterRequestClusterConfig {
	s.ClusterAdmins = v
	return s
}

func (s *UpdateComputeClusterRequestClusterConfig) SetClusterSafetyControl(v *UpdateComputeClusterRequestClusterConfigClusterSafetyControl) *UpdateComputeClusterRequestClusterConfig {
	s.ClusterSafetyControl = v
	return s
}

func (s *UpdateComputeClusterRequestClusterConfig) SetConfigList(v []*UpdateComputeClusterRequestClusterConfigConfigList) *UpdateComputeClusterRequestClusterConfig {
	s.ConfigList = v
	return s
}

func (s *UpdateComputeClusterRequestClusterConfig) SetDes(v string) *UpdateComputeClusterRequestClusterConfig {
	s.Des = &v
	return s
}

func (s *UpdateComputeClusterRequestClusterConfig) SetName(v string) *UpdateComputeClusterRequestClusterConfig {
	s.Name = &v
	return s
}

func (s *UpdateComputeClusterRequestClusterConfig) SetType(v string) *UpdateComputeClusterRequestClusterConfig {
	s.Type = &v
	return s
}

func (s *UpdateComputeClusterRequestClusterConfig) SetTypeVersion(v string) *UpdateComputeClusterRequestClusterConfig {
	s.TypeVersion = &v
	return s
}

func (s *UpdateComputeClusterRequestClusterConfig) Validate() error {
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

type UpdateComputeClusterRequestClusterConfigClusterSafetyControl struct {
	// 管控模式。CREATE_COMPUTE_SOURCE：有创建计算源权限即可使用；USER_DEFINE：仅白名单用户/用户组可用
	//
	// example:
	//
	// CREATE_COMPUTE_SOURCE
	ClusterSafetyAuthType *string   `json:"ClusterSafetyAuthType,omitempty" xml:"ClusterSafetyAuthType,omitempty"`
	UserGroupIds          []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	UserIds               []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s UpdateComputeClusterRequestClusterConfigClusterSafetyControl) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeClusterRequestClusterConfigClusterSafetyControl) GoString() string {
	return s.String()
}

func (s *UpdateComputeClusterRequestClusterConfigClusterSafetyControl) GetClusterSafetyAuthType() *string {
	return s.ClusterSafetyAuthType
}

func (s *UpdateComputeClusterRequestClusterConfigClusterSafetyControl) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *UpdateComputeClusterRequestClusterConfigClusterSafetyControl) GetUserIds() []*string {
	return s.UserIds
}

func (s *UpdateComputeClusterRequestClusterConfigClusterSafetyControl) SetClusterSafetyAuthType(v string) *UpdateComputeClusterRequestClusterConfigClusterSafetyControl {
	s.ClusterSafetyAuthType = &v
	return s
}

func (s *UpdateComputeClusterRequestClusterConfigClusterSafetyControl) SetUserGroupIds(v []*string) *UpdateComputeClusterRequestClusterConfigClusterSafetyControl {
	s.UserGroupIds = v
	return s
}

func (s *UpdateComputeClusterRequestClusterConfigClusterSafetyControl) SetUserIds(v []*string) *UpdateComputeClusterRequestClusterConfigClusterSafetyControl {
	s.UserIds = v
	return s
}

func (s *UpdateComputeClusterRequestClusterConfigClusterSafetyControl) Validate() error {
	return dara.Validate(s)
}

type UpdateComputeClusterRequestClusterConfigConfigList struct {
	// This parameter is required.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateComputeClusterRequestClusterConfigConfigList) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeClusterRequestClusterConfigConfigList) GoString() string {
	return s.String()
}

func (s *UpdateComputeClusterRequestClusterConfigConfigList) GetKey() *string {
	return s.Key
}

func (s *UpdateComputeClusterRequestClusterConfigConfigList) GetValue() *string {
	return s.Value
}

func (s *UpdateComputeClusterRequestClusterConfigConfigList) SetKey(v string) *UpdateComputeClusterRequestClusterConfigConfigList {
	s.Key = &v
	return s
}

func (s *UpdateComputeClusterRequestClusterConfigConfigList) SetValue(v string) *UpdateComputeClusterRequestClusterConfigConfigList {
	s.Value = &v
	return s
}

func (s *UpdateComputeClusterRequestClusterConfigConfigList) Validate() error {
	return dara.Validate(s)
}
