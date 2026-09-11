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
	// The rate limit configurations.
	//
	// - **dts.datamove.blaster.qps.max**: the queries per second (QPS) for querying the source database.
	//
	// - **dts.datamove.source.rps.max**: the records per second (RPS) for full data synchronization or migration.
	//
	// - **dts.datamove.source.bps.max**: the data volume per second for full data synchronization or migration. Unit: bytes per second.
	//
	// > - If **JobCode*	- is set to **03**, you must set **EnableLimit*	- to **true*	- for the three parameters to take effect.
	//
	// - If **JobCode*	- is set to **04*	- or **07**, you only need to configure **dts.datamove.source.rps.max*	- and **dts.datamove.source.bps.max**.
	//
	// - A value of **-1*	- indicates that no rate limit is applied.
	//
	// example:
	//
	// {
	//
	//       "dts.datamove.source.rps.max": 5000,
	//
	//       "dts.datamove.source.bps.max": 10485760
	//
	// }
	ConfigList *string `json:"ConfigList,omitempty" xml:"ConfigList,omitempty"`
	// The ID of the data synchronization or migration task.
	//
	// > You can call [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) to obtain the task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ta7w132u12h****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// Specifies whether to limit the migration rate of the full data synchronization or migration task. Valid values: **true*	- and **false**.
	//
	// > This parameter is required only when **JobCode*	- is set to **03**.
	//
	// example:
	//
	// true
	EnableLimit *bool `json:"EnableLimit,omitempty" xml:"EnableLimit,omitempty"`
	// The task code. Valid values:
	//
	// - **03**: full data synchronization or migration task.
	//
	// - **04**: incremental data migration task.
	//
	// - **07**: incremental data synchronization task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 07
	JobCode *string `json:"JobCode,omitempty" xml:"JobCode,omitempty"`
	// The ID of the region where the DTS instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
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
