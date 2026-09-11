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
	// The ID of the data migration or synchronization instance. You can call the **DescribeMigrationJobs*	- or DescribeSynchronizationJobs operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtsi76118o3w92****
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The precheck items to skip. Separate multiple items with commas (,). Valid values:
	//
	// - **CHECK_SAME_OBJ**: checks whether objects with the same name exist.
	//
	// - **CHECK_SAME_USER**: checks whether accounts with different names exist.
	//
	// - **CHECK_SRC**: checks the source database version.
	//
	// - **CHECK_TOPOLOGY**: checks the topology version.
	//
	// > For the topology versions supported by DTS, see [Topology overview](https://help.aliyun.com/document_detail/124115.html).
	//
	// - **CHECK_SERVER_ID**: checks the server_id of the source database.
	//
	// This parameter is required.
	//
	// example:
	//
	// CHECK_SAME_OBJ
	PrecheckItems *string `json:"PrecheckItems,omitempty" xml:"PrecheckItems,omitempty"`
	// The ID of the region where the instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
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
