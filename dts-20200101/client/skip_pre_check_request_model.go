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
	// The ID of the precheck task. You can call the **DescribePreCheckStatus*	- operation to query the task ID.
	//
	// example:
	//
	// b4my3zg929a****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The region where the DTS instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specifies whether to skip the precheck item. Valid values:
	//
	// - **true**: Skips the precheck item.
	//
	// - **false**: Does not skip the precheck item.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Skip *bool `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// The abbreviations of the precheck items to skip or unskip. Valid values:
	//
	// - **["CHECK_SAME_OBJ"]**: Checks for the existence of objects with the same name.
	//
	// - **["CHECK_SAME_USER"]**: Checks for the existence of accounts with different names.
	//
	// - **["CHECK_SRC"]**: Checks the version of the source database.
	//
	// - **["CHECK_TOPOLOGY"]**: Checks the topology version. For more information about the topology versions supported by DTS, see [Topology overview](https://help.aliyun.com/document_detail/124115.html).
	//
	// - **["CHECK_SERVER_ID"]**: Checks the server_id of the source database.
	//
	// - **["CHECK_DEST_TABLE_EMPTY"]**: Checks for the existence of objects in the destination database.
	//
	// - **["CHECK_SUPER_AUTH_DEST"]**: Checks the permissions of the super account in the destination database.
	//
	// > Separate multiple item names with commas (,), for example, **["CHECK_SRC","CHECK_SAME_OBJ"]**.
	//
	// example:
	//
	// ["CHECK_SAME_OBJ"]
	SkipPreCheckItems *string `json:"SkipPreCheckItems,omitempty" xml:"SkipPreCheckItems,omitempty"`
	// The names of the precheck items to skip or unskip. This parameter corresponds to the **SkipPreCheckItems*	- parameter. Valid values:
	//
	// - **["CHECK_SAME_OBJ_DETAIL"]**: Checks for the existence of objects with the same name.
	//
	// - **["CHECK_SAME_USER_DETAIL"]**: Checks for the existence of accounts with different names.
	//
	// - **["CHECK_SRC_DETAIL"]**: Checks the version of the source database.
	//
	// - **["CHECK_TOPOLOGY_DETAIL"]**: Checks the topology version. For more information about the topology versions supported by DTS, see [Topology overview](https://help.aliyun.com/document_detail/124115.html).
	//
	// - **["CHECK_SERVER_ID_DETAIL"]**: Checks the server_id of the source database.
	//
	// - **["CHECK_DEST_TABLE_EMPTY_DETAIL"]**: Checks whether the tables in the destination database are empty.
	//
	// - **["CHECK_SUPER_AUTH_DEST_DETAIL"]**: Checks the permissions of the super account in the destination database.
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
