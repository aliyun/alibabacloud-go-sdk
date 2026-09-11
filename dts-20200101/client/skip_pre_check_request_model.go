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
	// The ID of the precheck task. You can call **DescribePreCheckStatus*	- to query the ID.
	//
	// example:
	//
	// b4my3zg929a****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The region in which the DTS instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
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
	// Specifies whether to suppress the precheck item. Valid values:
	//
	// - **true**: Suppress the precheck item.
	//
	// - **false**: Unsuppress the precheck item.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Skip *bool `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// The abbreviated names of the precheck items to suppress or unsuppress. Valid values:
	//
	// - **["CHECK_SAME_OBJ"]**: check for objects with the same name.
	//
	// - **["CHECK_SAME_USER"]**: check for accounts with different names.
	//
	// - **["CHECK_SRC"]**: source database version check.
	//
	// - **["CHECK_TOPOLOGY"]**: topology version check. For the topology versions supported by DTS, see [Topology overview](https://help.aliyun.com/document_detail/124115.html).
	//
	// - **["CHECK_SERVER_ID"]**: source database server_id check.
	//
	// - **["CHECK_DEST_TABLE_EMPTY"]**: destination database object existence check.
	//
	// - **["CHECK_SUPER_AUTH_DEST"]**: destination database super account permission check.
	//
	// > Separate multiple item names with commas (,), for example, **["CHECK_SRC","CHECK_SAME_OBJ"]**.
	//
	// example:
	//
	// ["CHECK_SAME_OBJ"]
	SkipPreCheckItems *string `json:"SkipPreCheckItems,omitempty" xml:"SkipPreCheckItems,omitempty"`
	// The full names of the precheck items to suppress or unsuppress. This parameter corresponds to the **SkipPreCheckItems*	- parameter. Valid values:
	//
	// - **["CHECK_SAME_OBJ_DETAIL"]**: check for objects with the same name.
	//
	// - **["CHECK_SAME_USER_DETAIL"]**: check for accounts with different names.
	//
	// - **["CHECK_SRC_DETAIL"]**: source database version check.
	//
	// - **["CHECK_TOPOLOGY_DETAIL"]**: topology version check. For the topology versions supported by DTS, see [Topology overview](https://help.aliyun.com/document_detail/124115.html).
	//
	// - **["CHECK_SERVER_ID_DETAIL"]**: source database server_id check.
	//
	// - **["CHECK_DEST_TABLE_EMPTY_DETAIL"]**: check whether the destination database tables are empty.
	//
	// - **["CHECK_SUPER_AUTH_DEST_DETAIL"]**: check the super account permissions of the destination database.
	//
	// > Separate multiple item names with commas (,), for example, **["CHECK_SRC_DETAIL","CHECK_SAME_OBJ_DETAIL"]**.
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
