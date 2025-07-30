// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSkipPreCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsJobId(v string) *SkipPreCheckRequest
	GetDtsJobId() *string
	SetJobId(v string) *SkipPreCheckRequest
	GetJobId() *string
	SetRegionId(v string) *SkipPreCheckRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *SkipPreCheckRequest
	GetResourceGroupId() *string
	SetSkip(v bool) *SkipPreCheckRequest
	GetSkip() *bool
	SetSkipPreCheckItems(v string) *SkipPreCheckRequest
	GetSkipPreCheckItems() *string
	SetSkipPreCheckNames(v string) *SkipPreCheckRequest
	GetSkipPreCheckNames() *string
}

type SkipPreCheckRequest struct {
	// The ID of the data migration, data synchronization, or change tracking task.
	//
	// This parameter is required.
	//
	// example:
	//
	// c7412z57g8k****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The precheck task ID. You can call the **DescribePreCheckStatus*	- operation to query the task ID.
	//
	// example:
	//
	// b4my3zg929a****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The region ID of the DTS instance. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
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
	// Specifies whether to skip the precheck item. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Skip *bool `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// The shortened name of the precheck item. Valid values:
	//
	// 	- **["CHECK_SAME_OBJ"]**: object name conflict.
	//
	// 	- **["CHECK_SAME_USER"]**: username conflict.
	//
	// 	- **["CHECK_SRC"]**: source database version.
	//
	// 	- **["CHECK_TOPOLOGY"]**: topology. For more information about the topologies supported by Data Transmission Service (DTS), see [Synchronization topologies](https://help.aliyun.com/document_detail/124115.html).
	//
	// 	- **["CHECK_SERVER_ID"]**: the server ID of the source database.
	//
	// 	- **["CHECK_DEST_TABLE_EMPTY"]**: existence of objects in the destination database.
	//
	// > Separate multiple item names with commas (,). Example: **["CHECK_SRC","CHECK_SAME_OBJ"]**.
	//
	// example:
	//
	// ["CHECK_SAME_OBJ"]
	SkipPreCheckItems *string `json:"SkipPreCheckItems,omitempty" xml:"SkipPreCheckItems,omitempty"`
	// The precheck item name. This parameter corresponds to **SkipPreCheckItems**. Valid values:
	//
	// 	- **["CHECK_SAME_OBJ_DETAIL"]**: object name conflict.
	//
	// 	- **["CHECK_SAME_USER_DETAIL"]**: username conflict.
	//
	// 	- **["CHECK_SRC_DETAIL"]**: source database version.
	//
	// 	- **["CHECK_TOPOLOGY_DETAIL"]**: topology. For more information about the topologies supported by DTS, see [Synchronization topologies](https://help.aliyun.com/document_detail/124115.html).
	//
	// 	- **["CHECK_SERVER_ID_DETAIL"]**: the server ID of the source database.
	//
	// 	- **["CHECK_DEST_TABLE_EMPTY_DETAIL"]**: empty tables in the destination database.
	//
	// > Separate multiple item names with commas (,). Example: **["CHECK_SRC_DETAIL","CHECK_SAME_OBJ_DETAIL"]**.
	//
	// example:
	//
	// ["CHECK_SAME_OBJ_DETAIL"]
	SkipPreCheckNames *string `json:"SkipPreCheckNames,omitempty" xml:"SkipPreCheckNames,omitempty"`
}

func (s SkipPreCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s SkipPreCheckRequest) GoString() string {
	return s.String()
}

func (s *SkipPreCheckRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *SkipPreCheckRequest) GetJobId() *string {
	return s.JobId
}

func (s *SkipPreCheckRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SkipPreCheckRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *SkipPreCheckRequest) GetSkip() *bool {
	return s.Skip
}

func (s *SkipPreCheckRequest) GetSkipPreCheckItems() *string {
	return s.SkipPreCheckItems
}

func (s *SkipPreCheckRequest) GetSkipPreCheckNames() *string {
	return s.SkipPreCheckNames
}

func (s *SkipPreCheckRequest) SetDtsJobId(v string) *SkipPreCheckRequest {
	s.DtsJobId = &v
	return s
}

func (s *SkipPreCheckRequest) SetJobId(v string) *SkipPreCheckRequest {
	s.JobId = &v
	return s
}

func (s *SkipPreCheckRequest) SetRegionId(v string) *SkipPreCheckRequest {
	s.RegionId = &v
	return s
}

func (s *SkipPreCheckRequest) SetResourceGroupId(v string) *SkipPreCheckRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *SkipPreCheckRequest) SetSkip(v bool) *SkipPreCheckRequest {
	s.Skip = &v
	return s
}

func (s *SkipPreCheckRequest) SetSkipPreCheckItems(v string) *SkipPreCheckRequest {
	s.SkipPreCheckItems = &v
	return s
}

func (s *SkipPreCheckRequest) SetSkipPreCheckNames(v string) *SkipPreCheckRequest {
	s.SkipPreCheckNames = &v
	return s
}

func (s *SkipPreCheckRequest) Validate() error {
	return dara.Validate(s)
}
