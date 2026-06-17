// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridMonitorSLSGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ModifyHybridMonitorSLSGroupRequest
	GetRegionId() *string
	SetSLSGroupConfig(v []*ModifyHybridMonitorSLSGroupRequestSLSGroupConfig) *ModifyHybridMonitorSLSGroupRequest
	GetSLSGroupConfig() []*ModifyHybridMonitorSLSGroupRequestSLSGroupConfig
	SetSLSGroupDescription(v string) *ModifyHybridMonitorSLSGroupRequest
	GetSLSGroupDescription() *string
	SetSLSGroupName(v string) *ModifyHybridMonitorSLSGroupRequest
	GetSLSGroupName() *string
}

type ModifyHybridMonitorSLSGroupRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The configurations of the Logstore group.
	//
	// Valid values of N: 1 to 25.
	//
	// This parameter is required.
	SLSGroupConfig []*ModifyHybridMonitorSLSGroupRequestSLSGroupConfig `json:"SLSGroupConfig,omitempty" xml:"SLSGroupConfig,omitempty" type:"Repeated"`
	// The description of the Logstore group.
	//
	// example:
	//
	// Logstore group of Alibaba Cloud products.
	SLSGroupDescription *string `json:"SLSGroupDescription,omitempty" xml:"SLSGroupDescription,omitempty"`
	// The name of the Logstore group.
	//
	// For information about how to obtain the name of a Logstore group, see [DescribeHybridMonitorSLSGroup](https://help.aliyun.com/document_detail/429526.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// Logstore_test
	SLSGroupName *string `json:"SLSGroupName,omitempty" xml:"SLSGroupName,omitempty"`
}

func (s ModifyHybridMonitorSLSGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridMonitorSLSGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyHybridMonitorSLSGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyHybridMonitorSLSGroupRequest) GetSLSGroupConfig() []*ModifyHybridMonitorSLSGroupRequestSLSGroupConfig {
	return s.SLSGroupConfig
}

func (s *ModifyHybridMonitorSLSGroupRequest) GetSLSGroupDescription() *string {
	return s.SLSGroupDescription
}

func (s *ModifyHybridMonitorSLSGroupRequest) GetSLSGroupName() *string {
	return s.SLSGroupName
}

func (s *ModifyHybridMonitorSLSGroupRequest) SetRegionId(v string) *ModifyHybridMonitorSLSGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHybridMonitorSLSGroupRequest) SetSLSGroupConfig(v []*ModifyHybridMonitorSLSGroupRequestSLSGroupConfig) *ModifyHybridMonitorSLSGroupRequest {
	s.SLSGroupConfig = v
	return s
}

func (s *ModifyHybridMonitorSLSGroupRequest) SetSLSGroupDescription(v string) *ModifyHybridMonitorSLSGroupRequest {
	s.SLSGroupDescription = &v
	return s
}

func (s *ModifyHybridMonitorSLSGroupRequest) SetSLSGroupName(v string) *ModifyHybridMonitorSLSGroupRequest {
	s.SLSGroupName = &v
	return s
}

func (s *ModifyHybridMonitorSLSGroupRequest) Validate() error {
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

type ModifyHybridMonitorSLSGroupRequestSLSGroupConfig struct {
	// The Logstore.
	//
	// Valid values of N: 1 to 25.
	//
	// This parameter is required.
	//
	// example:
	//
	// Logstore-aliyun-all
	SLSLogstore *string `json:"SLSLogstore,omitempty" xml:"SLSLogstore,omitempty"`
	// The log project.
	//
	// Valid values of N: 1 to 25.
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyun-project
	SLSProject *string `json:"SLSProject,omitempty" xml:"SLSProject,omitempty"`
	// The region.
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
	// When you invoke this operation by using a management account, you can associate the Alibaba Cloud services of any member in the resource folder with Hybrid Cloud Monitoring through cloud native mode. You can use the resource folder to monitor the Alibaba Cloud services of multiple accounts.
	//
	// > If a member uses CloudMonitor for the first time, make sure that the member has completed authorization for the CloudMonitor service-linked role (AliyunServiceRoleForCloudMonitor). For more information, see [CloudMonitor service-linked role](https://help.aliyun.com/document_detail/170423.html).
	//
	// example:
	//
	// 120886317861****
	SLSUserId *string `json:"SLSUserId,omitempty" xml:"SLSUserId,omitempty"`
}

func (s ModifyHybridMonitorSLSGroupRequestSLSGroupConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridMonitorSLSGroupRequestSLSGroupConfig) GoString() string {
	return s.String()
}

func (s *ModifyHybridMonitorSLSGroupRequestSLSGroupConfig) GetSLSLogstore() *string {
	return s.SLSLogstore
}

func (s *ModifyHybridMonitorSLSGroupRequestSLSGroupConfig) GetSLSProject() *string {
	return s.SLSProject
}

func (s *ModifyHybridMonitorSLSGroupRequestSLSGroupConfig) GetSLSRegion() *string {
	return s.SLSRegion
}

func (s *ModifyHybridMonitorSLSGroupRequestSLSGroupConfig) GetSLSUserId() *string {
	return s.SLSUserId
}

func (s *ModifyHybridMonitorSLSGroupRequestSLSGroupConfig) SetSLSLogstore(v string) *ModifyHybridMonitorSLSGroupRequestSLSGroupConfig {
	s.SLSLogstore = &v
	return s
}

func (s *ModifyHybridMonitorSLSGroupRequestSLSGroupConfig) SetSLSProject(v string) *ModifyHybridMonitorSLSGroupRequestSLSGroupConfig {
	s.SLSProject = &v
	return s
}

func (s *ModifyHybridMonitorSLSGroupRequestSLSGroupConfig) SetSLSRegion(v string) *ModifyHybridMonitorSLSGroupRequestSLSGroupConfig {
	s.SLSRegion = &v
	return s
}

func (s *ModifyHybridMonitorSLSGroupRequestSLSGroupConfig) SetSLSUserId(v string) *ModifyHybridMonitorSLSGroupRequestSLSGroupConfig {
	s.SLSUserId = &v
	return s
}

func (s *ModifyHybridMonitorSLSGroupRequestSLSGroupConfig) Validate() error {
	return dara.Validate(s)
}
