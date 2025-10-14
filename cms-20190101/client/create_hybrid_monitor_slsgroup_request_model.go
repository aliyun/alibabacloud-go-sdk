// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHybridMonitorSLSGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CreateHybridMonitorSLSGroupRequest
	GetRegionId() *string
	SetSLSGroupConfig(v []*CreateHybridMonitorSLSGroupRequestSLSGroupConfig) *CreateHybridMonitorSLSGroupRequest
	GetSLSGroupConfig() []*CreateHybridMonitorSLSGroupRequestSLSGroupConfig
	SetSLSGroupDescription(v string) *CreateHybridMonitorSLSGroupRequest
	GetSLSGroupDescription() *string
	SetSLSGroupName(v string) *CreateHybridMonitorSLSGroupRequest
	GetSLSGroupName() *string
}

type CreateHybridMonitorSLSGroupRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The configurations of the Logstore group.
	//
	// Valid values of N: 1 to 25.
	//
	// This parameter is required.
	SLSGroupConfig []*CreateHybridMonitorSLSGroupRequestSLSGroupConfig `json:"SLSGroupConfig,omitempty" xml:"SLSGroupConfig,omitempty" type:"Repeated"`
	// The description of the Logstore group.
	SLSGroupDescription *string `json:"SLSGroupDescription,omitempty" xml:"SLSGroupDescription,omitempty"`
	// The name of the Logstore group.
	//
	// The name must be 2 to 32 characters in length and can contain uppercase letters, lowercase letters, digits, and underscores (_). The name must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// Logstore_test
	SLSGroupName *string `json:"SLSGroupName,omitempty" xml:"SLSGroupName,omitempty"`
}

func (s CreateHybridMonitorSLSGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHybridMonitorSLSGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateHybridMonitorSLSGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateHybridMonitorSLSGroupRequest) GetSLSGroupConfig() []*CreateHybridMonitorSLSGroupRequestSLSGroupConfig {
	return s.SLSGroupConfig
}

func (s *CreateHybridMonitorSLSGroupRequest) GetSLSGroupDescription() *string {
	return s.SLSGroupDescription
}

func (s *CreateHybridMonitorSLSGroupRequest) GetSLSGroupName() *string {
	return s.SLSGroupName
}

func (s *CreateHybridMonitorSLSGroupRequest) SetRegionId(v string) *CreateHybridMonitorSLSGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateHybridMonitorSLSGroupRequest) SetSLSGroupConfig(v []*CreateHybridMonitorSLSGroupRequestSLSGroupConfig) *CreateHybridMonitorSLSGroupRequest {
	s.SLSGroupConfig = v
	return s
}

func (s *CreateHybridMonitorSLSGroupRequest) SetSLSGroupDescription(v string) *CreateHybridMonitorSLSGroupRequest {
	s.SLSGroupDescription = &v
	return s
}

func (s *CreateHybridMonitorSLSGroupRequest) SetSLSGroupName(v string) *CreateHybridMonitorSLSGroupRequest {
	s.SLSGroupName = &v
	return s
}

func (s *CreateHybridMonitorSLSGroupRequest) Validate() error {
	if s.SLSGroupConfig != nil {
		for _, item := range s.SLSGroupConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateHybridMonitorSLSGroupRequestSLSGroupConfig struct {
	// The Logstore.
	//
	// Valid values of N: 1 to 25.
	//
	// This parameter is required.
	//
	// example:
	//
	// Logstore-ECS
	SLSLogstore *string `json:"SLSLogstore,omitempty" xml:"SLSLogstore,omitempty"`
	// The Simple Log Service project.
	//
	// Valid values of N: 1 to 25.
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyun-project
	SLSProject *string `json:"SLSProject,omitempty" xml:"SLSProject,omitempty"`
	// The region ID.
	//
	// Valid values of N: 1 to 25.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	SLSRegion *string `json:"SLSRegion,omitempty" xml:"SLSRegion,omitempty"`
	// The member ID.
	//
	// Valid values of N: 1 to 25.
	//
	// If you call this operation by using the management account of a resource directory, you can connect the Alibaba Cloud services that are activated for all members in the resource directory to Hybrid Cloud Monitoring. You can use the resource directory to monitor Alibaba Cloud services across enterprise accounts.
	//
	// > If a member uses CloudMonitor for the first time, you must make sure that the service-linked role AliyunServiceRoleForCloudMonitor is attached to the member. For more information, see [Manage the service-linked role for CloudMonitor](https://help.aliyun.com/document_detail/170423.html).
	//
	// example:
	//
	// 120886317861****
	SLSUserId *string `json:"SLSUserId,omitempty" xml:"SLSUserId,omitempty"`
}

func (s CreateHybridMonitorSLSGroupRequestSLSGroupConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateHybridMonitorSLSGroupRequestSLSGroupConfig) GoString() string {
	return s.String()
}

func (s *CreateHybridMonitorSLSGroupRequestSLSGroupConfig) GetSLSLogstore() *string {
	return s.SLSLogstore
}

func (s *CreateHybridMonitorSLSGroupRequestSLSGroupConfig) GetSLSProject() *string {
	return s.SLSProject
}

func (s *CreateHybridMonitorSLSGroupRequestSLSGroupConfig) GetSLSRegion() *string {
	return s.SLSRegion
}

func (s *CreateHybridMonitorSLSGroupRequestSLSGroupConfig) GetSLSUserId() *string {
	return s.SLSUserId
}

func (s *CreateHybridMonitorSLSGroupRequestSLSGroupConfig) SetSLSLogstore(v string) *CreateHybridMonitorSLSGroupRequestSLSGroupConfig {
	s.SLSLogstore = &v
	return s
}

func (s *CreateHybridMonitorSLSGroupRequestSLSGroupConfig) SetSLSProject(v string) *CreateHybridMonitorSLSGroupRequestSLSGroupConfig {
	s.SLSProject = &v
	return s
}

func (s *CreateHybridMonitorSLSGroupRequestSLSGroupConfig) SetSLSRegion(v string) *CreateHybridMonitorSLSGroupRequestSLSGroupConfig {
	s.SLSRegion = &v
	return s
}

func (s *CreateHybridMonitorSLSGroupRequestSLSGroupConfig) SetSLSUserId(v string) *CreateHybridMonitorSLSGroupRequestSLSGroupConfig {
	s.SLSUserId = &v
	return s
}

func (s *CreateHybridMonitorSLSGroupRequestSLSGroupConfig) Validate() error {
	return dara.Validate(s)
}
