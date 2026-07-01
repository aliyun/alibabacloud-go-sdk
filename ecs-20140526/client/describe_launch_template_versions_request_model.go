// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLaunchTemplateVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultVersion(v bool) *DescribeLaunchTemplateVersionsRequest
	GetDefaultVersion() *bool
	SetDetailFlag(v bool) *DescribeLaunchTemplateVersionsRequest
	GetDetailFlag() *bool
	SetLaunchTemplateId(v string) *DescribeLaunchTemplateVersionsRequest
	GetLaunchTemplateId() *string
	SetLaunchTemplateName(v string) *DescribeLaunchTemplateVersionsRequest
	GetLaunchTemplateName() *string
	SetLaunchTemplateVersion(v []*int64) *DescribeLaunchTemplateVersionsRequest
	GetLaunchTemplateVersion() []*int64
	SetMaxVersion(v int64) *DescribeLaunchTemplateVersionsRequest
	GetMaxVersion() *int64
	SetMinVersion(v int64) *DescribeLaunchTemplateVersionsRequest
	GetMinVersion() *int64
	SetOwnerAccount(v string) *DescribeLaunchTemplateVersionsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeLaunchTemplateVersionsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeLaunchTemplateVersionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLaunchTemplateVersionsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeLaunchTemplateVersionsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeLaunchTemplateVersionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeLaunchTemplateVersionsRequest
	GetResourceOwnerId() *int64
}

type DescribeLaunchTemplateVersionsRequest struct {
	// Specifies whether to query the default version.
	//
	// example:
	//
	// true
	DefaultVersion *bool `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// Specifies whether to query detailed template configuration information. Valid values:
	//
	// - true: Queries detailed template configuration information. In addition to basic template information, detailed configuration such as image ID and system disk size is returned.
	//
	// - false: Queries only basic template information, such as template ID, template name, and default version.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	DetailFlag *bool `json:"DetailFlag,omitempty" xml:"DetailFlag,omitempty"`
	// The launch template ID.
	//
	// You must specify `LaunchTemplateId` or `LaunchTemplateName` to determine the template.
	//
	// example:
	//
	// lt-bp168lnahrdwl39p****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The launch template name.
	//
	// You must specify `LaunchTemplateId` or `LaunchTemplateName` to determine the template.
	//
	// example:
	//
	// testLaunchTemplateName
	LaunchTemplateName *string `json:"LaunchTemplateName,omitempty" xml:"LaunchTemplateName,omitempty"`
	// One or more launch template version numbers.
	//
	// example:
	//
	// 1
	LaunchTemplateVersion []*int64 `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty" type:"Repeated"`
	// The maximum version number used to filter query results. Used together with `MinVersion` to query version information within the range between the minimum and maximum version numbers.
	//
	// example:
	//
	// 10
	MaxVersion *int64 `json:"MaxVersion,omitempty" xml:"MaxVersion,omitempty"`
	// The minimum version number used to filter query results. Used together with `MaxVersion` to query version information within the range between the minimum and maximum version numbers.
	//
	// example:
	//
	// 1
	MinVersion   *int64  `json:"MinVersion,omitempty" xml:"MinVersion,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the launch template list.
	//
	// Minimum value: 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paged query. Settings this parameter for paging.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the launch template.
	//
	// You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeLaunchTemplateVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLaunchTemplateVersionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLaunchTemplateVersionsRequest) GetDefaultVersion() *bool {
	return s.DefaultVersion
}

func (s *DescribeLaunchTemplateVersionsRequest) GetDetailFlag() *bool {
	return s.DetailFlag
}

func (s *DescribeLaunchTemplateVersionsRequest) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *DescribeLaunchTemplateVersionsRequest) GetLaunchTemplateName() *string {
	return s.LaunchTemplateName
}

func (s *DescribeLaunchTemplateVersionsRequest) GetLaunchTemplateVersion() []*int64 {
	return s.LaunchTemplateVersion
}

func (s *DescribeLaunchTemplateVersionsRequest) GetMaxVersion() *int64 {
	return s.MaxVersion
}

func (s *DescribeLaunchTemplateVersionsRequest) GetMinVersion() *int64 {
	return s.MinVersion
}

func (s *DescribeLaunchTemplateVersionsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeLaunchTemplateVersionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLaunchTemplateVersionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLaunchTemplateVersionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLaunchTemplateVersionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLaunchTemplateVersionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLaunchTemplateVersionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeLaunchTemplateVersionsRequest) SetDefaultVersion(v bool) *DescribeLaunchTemplateVersionsRequest {
	s.DefaultVersion = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsRequest) SetDetailFlag(v bool) *DescribeLaunchTemplateVersionsRequest {
	s.DetailFlag = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsRequest) SetLaunchTemplateId(v string) *DescribeLaunchTemplateVersionsRequest {
	s.LaunchTemplateId = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsRequest) SetLaunchTemplateName(v string) *DescribeLaunchTemplateVersionsRequest {
	s.LaunchTemplateName = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsRequest) SetLaunchTemplateVersion(v []*int64) *DescribeLaunchTemplateVersionsRequest {
	s.LaunchTemplateVersion = v
	return s
}

func (s *DescribeLaunchTemplateVersionsRequest) SetMaxVersion(v int64) *DescribeLaunchTemplateVersionsRequest {
	s.MaxVersion = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsRequest) SetMinVersion(v int64) *DescribeLaunchTemplateVersionsRequest {
	s.MinVersion = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsRequest) SetOwnerAccount(v string) *DescribeLaunchTemplateVersionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsRequest) SetOwnerId(v int64) *DescribeLaunchTemplateVersionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsRequest) SetPageNumber(v int32) *DescribeLaunchTemplateVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsRequest) SetPageSize(v int32) *DescribeLaunchTemplateVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsRequest) SetRegionId(v string) *DescribeLaunchTemplateVersionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsRequest) SetResourceOwnerAccount(v string) *DescribeLaunchTemplateVersionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsRequest) SetResourceOwnerId(v int64) *DescribeLaunchTemplateVersionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsRequest) Validate() error {
	return dara.Validate(s)
}
