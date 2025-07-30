// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShieldPrecheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsInstanceId(v string) *ShieldPrecheckRequest
	GetDtsInstanceId() *string
	SetPrecheckItems(v string) *ShieldPrecheckRequest
	GetPrecheckItems() *string
	SetRegionId(v string) *ShieldPrecheckRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ShieldPrecheckRequest
	GetResourceGroupId() *string
}

type ShieldPrecheckRequest struct {
	// The ID of the change tracking instance. You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the instance ID.
	//
	// >  You must specify at least one of the **DtsInstanceId*	- and **DtsJobId*	- parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtsi76118o3w92****
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The precheck items that you want to ignore. Separate multiple items with commas (,). Valid values:
	//
	//
	//
	// 	- **CHECK_SAME_OBJ**: schema name conflict
	//
	//
	//
	// 	- **CHECK_SAME_USER**: multiple usernames for one instance
	//
	//
	//
	// 	- **CHECK_SRC**: source database version
	//
	//
	//
	// 	- **CHECK_TOPOLOGY**: topology
	//
	//
	//
	// > For more information about the topologies supported by DTS, see [Synchronization topologies](https://help.aliyun.com/document_detail/124115.html).
	//
	//
	//
	// 	- **CHECK_SERVER_ID**: value of server_id in the source database
	//
	// This parameter is required.
	//
	// example:
	//
	// CHECK_SAME_OBJ
	PrecheckItems *string `json:"PrecheckItems,omitempty" xml:"PrecheckItems,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ShieldPrecheckRequest) String() string {
	return dara.Prettify(s)
}

func (s ShieldPrecheckRequest) GoString() string {
	return s.String()
}

func (s *ShieldPrecheckRequest) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *ShieldPrecheckRequest) GetPrecheckItems() *string {
	return s.PrecheckItems
}

func (s *ShieldPrecheckRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ShieldPrecheckRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ShieldPrecheckRequest) SetDtsInstanceId(v string) *ShieldPrecheckRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *ShieldPrecheckRequest) SetPrecheckItems(v string) *ShieldPrecheckRequest {
	s.PrecheckItems = &v
	return s
}

func (s *ShieldPrecheckRequest) SetRegionId(v string) *ShieldPrecheckRequest {
	s.RegionId = &v
	return s
}

func (s *ShieldPrecheckRequest) SetResourceGroupId(v string) *ShieldPrecheckRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ShieldPrecheckRequest) Validate() error {
	return dara.Validate(s)
}
