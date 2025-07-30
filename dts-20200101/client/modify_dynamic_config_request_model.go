// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDynamicConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigList(v string) *ModifyDynamicConfigRequest
	GetConfigList() *string
	SetDtsJobId(v string) *ModifyDynamicConfigRequest
	GetDtsJobId() *string
	SetEnableLimit(v bool) *ModifyDynamicConfigRequest
	GetEnableLimit() *bool
	SetJobCode(v string) *ModifyDynamicConfigRequest
	GetJobCode() *string
	SetRegionId(v string) *ModifyDynamicConfigRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifyDynamicConfigRequest
	GetResourceGroupId() *string
}

type ModifyDynamicConfigRequest struct {
	// The specific throttling configuration.
	//
	// 	- **dts.datamove.blaster.qps.max**: The rate at which queries are made to the source database per second.
	//
	// 	- **dts.datamove.source.rps.max**: The number of rows that are fully synchronized or migrated per second.
	//
	// 	- **dts.datamove.source.bps.max**: the amount of data processed per second for full synchronization or migration. Unit: MB.
	//
	// >
	//
	// 	- If you set the **JobCode*	- parameter to **03**, you need to specify **true*	- for the **EnableLimit*	- parameter. Otherwise, the configuration cannot take effect.
	//
	// 	- If you set the **JobCode*	- parameter to **04*	- or **07**, you only need to specify the **dts.datamove.source.rps.max*	- and **dts.datamove.source.bps.max*	- parameters.
	//
	// 	- A value of \\*\\*-1\\*\\	- indicates no rate limit.
	//
	// example:
	//
	// {\\"dts.datamove.source.rps.max\\":5000,\\"dts.datamove.source.bps.max\\":10485760}
	ConfigList *string `json:"ConfigList,omitempty" xml:"ConfigList,omitempty"`
	// The ID of the data migration or synchronization task.
	//
	// >  You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ta7w132u12h****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// Specifies whether to enable throttling for data synchronization or migration. Valid values: **true*	- and **false**.
	//
	// >  Only needs to be configured when the **JobCode*	- parameter is set to **03**.
	//
	// example:
	//
	// true
	EnableLimit *bool `json:"EnableLimit,omitempty" xml:"EnableLimit,omitempty"`
	// The task type. Valid values:
	//
	// 	- **03**: a full data synchronization or full data migration task.
	//
	// 	- **04**: an incremental data migration task.
	//
	// 	- **07**: an incremental data synchronization task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 07
	JobCode *string `json:"JobCode,omitempty" xml:"JobCode,omitempty"`
	// The region ID of the DTS instance. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-aekz6zsi7ce5rpy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ModifyDynamicConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDynamicConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDynamicConfigRequest) GetConfigList() *string {
	return s.ConfigList
}

func (s *ModifyDynamicConfigRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *ModifyDynamicConfigRequest) GetEnableLimit() *bool {
	return s.EnableLimit
}

func (s *ModifyDynamicConfigRequest) GetJobCode() *string {
	return s.JobCode
}

func (s *ModifyDynamicConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDynamicConfigRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyDynamicConfigRequest) SetConfigList(v string) *ModifyDynamicConfigRequest {
	s.ConfigList = &v
	return s
}

func (s *ModifyDynamicConfigRequest) SetDtsJobId(v string) *ModifyDynamicConfigRequest {
	s.DtsJobId = &v
	return s
}

func (s *ModifyDynamicConfigRequest) SetEnableLimit(v bool) *ModifyDynamicConfigRequest {
	s.EnableLimit = &v
	return s
}

func (s *ModifyDynamicConfigRequest) SetJobCode(v string) *ModifyDynamicConfigRequest {
	s.JobCode = &v
	return s
}

func (s *ModifyDynamicConfigRequest) SetRegionId(v string) *ModifyDynamicConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDynamicConfigRequest) SetResourceGroupId(v string) *ModifyDynamicConfigRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDynamicConfigRequest) Validate() error {
	return dara.Validate(s)
}
