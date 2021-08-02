// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ClonePolicyGroupRequest struct {
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty" require:"true"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s ClonePolicyGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ClonePolicyGroupRequest) GoString() string {
	return s.String()
}

func (s *ClonePolicyGroupRequest) SetRegionId(v string) *ClonePolicyGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ClonePolicyGroupRequest) SetPolicyGroupId(v string) *ClonePolicyGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ClonePolicyGroupRequest) SetName(v string) *ClonePolicyGroupRequest {
	s.Name = &v
	return s
}

type ClonePolicyGroupResponse struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty" require:"true"`
}

func (s ClonePolicyGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ClonePolicyGroupResponse) GoString() string {
	return s.String()
}

func (s *ClonePolicyGroupResponse) SetRequestId(v string) *ClonePolicyGroupResponse {
	s.RequestId = &v
	return s
}

func (s *ClonePolicyGroupResponse) SetPolicyGroupId(v string) *ClonePolicyGroupResponse {
	s.PolicyGroupId = &v
	return s
}

type DescribeDesktopsInGroupRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty" require:"true"`
	PayType        *string `json:"PayType,omitempty" xml:"PayType,omitempty" require:"true"`
	MaxResults     *int    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeDesktopsInGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsInGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsInGroupRequest) SetRegionId(v string) *DescribeDesktopsInGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDesktopsInGroupRequest) SetDesktopGroupId(v string) *DescribeDesktopsInGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeDesktopsInGroupRequest) SetPayType(v string) *DescribeDesktopsInGroupRequest {
	s.PayType = &v
	return s
}

func (s *DescribeDesktopsInGroupRequest) SetMaxResults(v int) *DescribeDesktopsInGroupRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDesktopsInGroupRequest) SetNextToken(v string) *DescribeDesktopsInGroupRequest {
	s.NextToken = &v
	return s
}

type DescribeDesktopsInGroupResponse struct {
	RequestId                   *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PaidDesktopsCount           *int                                               `json:"PaidDesktopsCount,omitempty" xml:"PaidDesktopsCount,omitempty" require:"true"`
	PostPaidDesktopsCount       *int                                               `json:"PostPaidDesktopsCount,omitempty" xml:"PostPaidDesktopsCount,omitempty" require:"true"`
	PostPaidDesktopsTotalAmount *int                                               `json:"PostPaidDesktopsTotalAmount,omitempty" xml:"PostPaidDesktopsTotalAmount,omitempty" require:"true"`
	NextToken                   *string                                            `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	PaidDesktops                []*DescribeDesktopsInGroupResponsePaidDesktops     `json:"PaidDesktops,omitempty" xml:"PaidDesktops,omitempty" require:"true" type:"Repeated"`
	PostPaidDesktops            []*DescribeDesktopsInGroupResponsePostPaidDesktops `json:"PostPaidDesktops,omitempty" xml:"PostPaidDesktops,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDesktopsInGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsInGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsInGroupResponse) SetRequestId(v string) *DescribeDesktopsInGroupResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDesktopsInGroupResponse) SetPaidDesktopsCount(v int) *DescribeDesktopsInGroupResponse {
	s.PaidDesktopsCount = &v
	return s
}

func (s *DescribeDesktopsInGroupResponse) SetPostPaidDesktopsCount(v int) *DescribeDesktopsInGroupResponse {
	s.PostPaidDesktopsCount = &v
	return s
}

func (s *DescribeDesktopsInGroupResponse) SetPostPaidDesktopsTotalAmount(v int) *DescribeDesktopsInGroupResponse {
	s.PostPaidDesktopsTotalAmount = &v
	return s
}

func (s *DescribeDesktopsInGroupResponse) SetNextToken(v string) *DescribeDesktopsInGroupResponse {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopsInGroupResponse) SetPaidDesktops(v []*DescribeDesktopsInGroupResponsePaidDesktops) *DescribeDesktopsInGroupResponse {
	s.PaidDesktops = v
	return s
}

func (s *DescribeDesktopsInGroupResponse) SetPostPaidDesktops(v []*DescribeDesktopsInGroupResponsePostPaidDesktops) *DescribeDesktopsInGroupResponse {
	s.PostPaidDesktops = v
	return s
}

type DescribeDesktopsInGroupResponsePaidDesktops struct {
	DesktopId        *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	DesktopName      *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty" require:"true"`
	DesktopStatus    *string `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty" require:"true"`
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty" require:"true"`
	EndUserId        *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" require:"true"`
	EndUserName      *string `json:"EndUserName,omitempty" xml:"EndUserName,omitempty" require:"true"`
}

func (s DescribeDesktopsInGroupResponsePaidDesktops) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsInGroupResponsePaidDesktops) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsInGroupResponsePaidDesktops) SetDesktopId(v string) *DescribeDesktopsInGroupResponsePaidDesktops {
	s.DesktopId = &v
	return s
}

func (s *DescribeDesktopsInGroupResponsePaidDesktops) SetDesktopName(v string) *DescribeDesktopsInGroupResponsePaidDesktops {
	s.DesktopName = &v
	return s
}

func (s *DescribeDesktopsInGroupResponsePaidDesktops) SetDesktopStatus(v string) *DescribeDesktopsInGroupResponsePaidDesktops {
	s.DesktopStatus = &v
	return s
}

func (s *DescribeDesktopsInGroupResponsePaidDesktops) SetConnectionStatus(v string) *DescribeDesktopsInGroupResponsePaidDesktops {
	s.ConnectionStatus = &v
	return s
}

func (s *DescribeDesktopsInGroupResponsePaidDesktops) SetEndUserId(v string) *DescribeDesktopsInGroupResponsePaidDesktops {
	s.EndUserId = &v
	return s
}

func (s *DescribeDesktopsInGroupResponsePaidDesktops) SetEndUserName(v string) *DescribeDesktopsInGroupResponsePaidDesktops {
	s.EndUserName = &v
	return s
}

type DescribeDesktopsInGroupResponsePostPaidDesktops struct {
	DesktopId        *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	DesktopName      *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty" require:"true"`
	DesktopStatus    *string `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty" require:"true"`
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty" require:"true"`
	EndUserId        *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" require:"true"`
	EndUserName      *string `json:"EndUserName,omitempty" xml:"EndUserName,omitempty" require:"true"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	ReleaseTime      *string `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty" require:"true"`
	CreateDuration   *string `json:"CreateDuration,omitempty" xml:"CreateDuration,omitempty" require:"true"`
}

func (s DescribeDesktopsInGroupResponsePostPaidDesktops) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsInGroupResponsePostPaidDesktops) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsInGroupResponsePostPaidDesktops) SetDesktopId(v string) *DescribeDesktopsInGroupResponsePostPaidDesktops {
	s.DesktopId = &v
	return s
}

func (s *DescribeDesktopsInGroupResponsePostPaidDesktops) SetDesktopName(v string) *DescribeDesktopsInGroupResponsePostPaidDesktops {
	s.DesktopName = &v
	return s
}

func (s *DescribeDesktopsInGroupResponsePostPaidDesktops) SetDesktopStatus(v string) *DescribeDesktopsInGroupResponsePostPaidDesktops {
	s.DesktopStatus = &v
	return s
}

func (s *DescribeDesktopsInGroupResponsePostPaidDesktops) SetConnectionStatus(v string) *DescribeDesktopsInGroupResponsePostPaidDesktops {
	s.ConnectionStatus = &v
	return s
}

func (s *DescribeDesktopsInGroupResponsePostPaidDesktops) SetEndUserId(v string) *DescribeDesktopsInGroupResponsePostPaidDesktops {
	s.EndUserId = &v
	return s
}

func (s *DescribeDesktopsInGroupResponsePostPaidDesktops) SetEndUserName(v string) *DescribeDesktopsInGroupResponsePostPaidDesktops {
	s.EndUserName = &v
	return s
}

func (s *DescribeDesktopsInGroupResponsePostPaidDesktops) SetCreateTime(v string) *DescribeDesktopsInGroupResponsePostPaidDesktops {
	s.CreateTime = &v
	return s
}

func (s *DescribeDesktopsInGroupResponsePostPaidDesktops) SetReleaseTime(v string) *DescribeDesktopsInGroupResponsePostPaidDesktops {
	s.ReleaseTime = &v
	return s
}

func (s *DescribeDesktopsInGroupResponsePostPaidDesktops) SetCreateDuration(v string) *DescribeDesktopsInGroupResponsePostPaidDesktops {
	s.CreateDuration = &v
	return s
}

type DescribeUsersInGroupRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty" require:"true"`
	MaxResults     *int    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeUsersInGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersInGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeUsersInGroupRequest) SetRegionId(v string) *DescribeUsersInGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUsersInGroupRequest) SetDesktopGroupId(v string) *DescribeUsersInGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeUsersInGroupRequest) SetMaxResults(v int) *DescribeUsersInGroupRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeUsersInGroupRequest) SetNextToken(v string) *DescribeUsersInGroupRequest {
	s.NextToken = &v
	return s
}

type DescribeUsersInGroupResponse struct {
	RequestId        *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	UsersCount       *int                                    `json:"UsersCount,omitempty" xml:"UsersCount,omitempty" require:"true"`
	OnlineUsersCount *int                                    `json:"OnlineUsersCount,omitempty" xml:"OnlineUsersCount,omitempty" require:"true"`
	NextToken        *string                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	EndUsers         []*DescribeUsersInGroupResponseEndUsers `json:"EndUsers,omitempty" xml:"EndUsers,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeUsersInGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersInGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeUsersInGroupResponse) SetRequestId(v string) *DescribeUsersInGroupResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeUsersInGroupResponse) SetUsersCount(v int) *DescribeUsersInGroupResponse {
	s.UsersCount = &v
	return s
}

func (s *DescribeUsersInGroupResponse) SetOnlineUsersCount(v int) *DescribeUsersInGroupResponse {
	s.OnlineUsersCount = &v
	return s
}

func (s *DescribeUsersInGroupResponse) SetNextToken(v string) *DescribeUsersInGroupResponse {
	s.NextToken = &v
	return s
}

func (s *DescribeUsersInGroupResponse) SetEndUsers(v []*DescribeUsersInGroupResponseEndUsers) *DescribeUsersInGroupResponse {
	s.EndUsers = v
	return s
}

type DescribeUsersInGroupResponseEndUsers struct {
	EndUserId        *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" require:"true"`
	EndUserName      *string `json:"EndUserName,omitempty" xml:"EndUserName,omitempty" require:"true"`
	EndUserType      *string `json:"EndUserType,omitempty" xml:"EndUserType,omitempty" require:"true"`
	EndUserEmail     *string `json:"EndUserEmail,omitempty" xml:"EndUserEmail,omitempty" require:"true"`
	EndUserPhone     *string `json:"EndUserPhone,omitempty" xml:"EndUserPhone,omitempty" require:"true"`
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty" require:"true"`
	DesktopId        *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	DesktopName      *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty" require:"true"`
}

func (s DescribeUsersInGroupResponseEndUsers) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersInGroupResponseEndUsers) GoString() string {
	return s.String()
}

func (s *DescribeUsersInGroupResponseEndUsers) SetEndUserId(v string) *DescribeUsersInGroupResponseEndUsers {
	s.EndUserId = &v
	return s
}

func (s *DescribeUsersInGroupResponseEndUsers) SetEndUserName(v string) *DescribeUsersInGroupResponseEndUsers {
	s.EndUserName = &v
	return s
}

func (s *DescribeUsersInGroupResponseEndUsers) SetEndUserType(v string) *DescribeUsersInGroupResponseEndUsers {
	s.EndUserType = &v
	return s
}

func (s *DescribeUsersInGroupResponseEndUsers) SetEndUserEmail(v string) *DescribeUsersInGroupResponseEndUsers {
	s.EndUserEmail = &v
	return s
}

func (s *DescribeUsersInGroupResponseEndUsers) SetEndUserPhone(v string) *DescribeUsersInGroupResponseEndUsers {
	s.EndUserPhone = &v
	return s
}

func (s *DescribeUsersInGroupResponseEndUsers) SetConnectionStatus(v string) *DescribeUsersInGroupResponseEndUsers {
	s.ConnectionStatus = &v
	return s
}

func (s *DescribeUsersInGroupResponseEndUsers) SetDesktopId(v string) *DescribeUsersInGroupResponseEndUsers {
	s.DesktopId = &v
	return s
}

func (s *DescribeUsersInGroupResponseEndUsers) SetDesktopName(v string) *DescribeUsersInGroupResponseEndUsers {
	s.DesktopName = &v
	return s
}

type GetDesktopGroupDetailRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty" require:"true"`
}

func (s GetDesktopGroupDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDesktopGroupDetailRequest) GoString() string {
	return s.String()
}

func (s *GetDesktopGroupDetailRequest) SetRegionId(v string) *GetDesktopGroupDetailRequest {
	s.RegionId = &v
	return s
}

func (s *GetDesktopGroupDetailRequest) SetDesktopGroupId(v string) *GetDesktopGroupDetailRequest {
	s.DesktopGroupId = &v
	return s
}

type GetDesktopGroupDetailResponse struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Desktops  []*GetDesktopGroupDetailResponseDesktops `json:"Desktops,omitempty" xml:"Desktops,omitempty" require:"true" type:"Repeated"`
}

func (s GetDesktopGroupDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDesktopGroupDetailResponse) GoString() string {
	return s.String()
}

func (s *GetDesktopGroupDetailResponse) SetRequestId(v string) *GetDesktopGroupDetailResponse {
	s.RequestId = &v
	return s
}

func (s *GetDesktopGroupDetailResponse) SetDesktops(v []*GetDesktopGroupDetailResponseDesktops) *GetDesktopGroupDetailResponse {
	s.Desktops = v
	return s
}

type GetDesktopGroupDetailResponseDesktops struct {
	DesktopGroupId     *string  `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty" require:"true"`
	DesktopGroupName   *string  `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty" require:"true"`
	OfficeSiteId       *string  `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true"`
	OfficeSiteName     *string  `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty" require:"true"`
	OfficeSiteType     *string  `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty" require:"true"`
	PolicyGroupId      *string  `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty" require:"true"`
	PolicyGroupName    *string  `json:"PolicyGroupName,omitempty" xml:"PolicyGroupName,omitempty" require:"true"`
	OwnBundleId        *string  `json:"OwnBundleId,omitempty" xml:"OwnBundleId,omitempty" require:"true"`
	OwnBundleName      *string  `json:"OwnBundleName,omitempty" xml:"OwnBundleName,omitempty" require:"true"`
	PayType            *string  `json:"PayType,omitempty" xml:"PayType,omitempty" require:"true"`
	ExpiredTime        *string  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty" require:"true"`
	CreationTime       *string  `json:"CreationTime,omitempty" xml:"CreationTime,omitempty" require:"true"`
	DirectoryId        *string  `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" require:"true"`
	DirectoryType      *string  `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty" require:"true"`
	Cpu                *int     `json:"Cpu,omitempty" xml:"Cpu,omitempty" require:"true"`
	Memory             *int64   `json:"Memory,omitempty" xml:"Memory,omitempty" require:"true"`
	GpuCount           *float32 `json:"GpuCount,omitempty" xml:"GpuCount,omitempty" require:"true"`
	GpuSpec            *string  `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty" require:"true"`
	SystemDiskCategory *string  `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty" require:"true"`
	SystemDiskSize     *int     `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty" require:"true"`
	DataDiskCategory   *string  `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty" require:"true"`
	DataDiskSize       *string  `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty" require:"true"`
	Creator            *string  `json:"Creator,omitempty" xml:"Creator,omitempty" require:"true"`
	Comments           *string  `json:"Comments,omitempty" xml:"Comments,omitempty" require:"true"`
	KeepDuration       *int64   `json:"KeepDuration,omitempty" xml:"KeepDuration,omitempty" require:"true"`
	MinDesktopsCount   *int     `json:"MinDesktopsCount,omitempty" xml:"MinDesktopsCount,omitempty" require:"true"`
	MaxDesktopsCount   *int     `json:"MaxDesktopsCount,omitempty" xml:"MaxDesktopsCount,omitempty" require:"true"`
	ResType            *int     `json:"ResType,omitempty" xml:"ResType,omitempty" require:"true"`
	AllowAutoSetup     *int     `json:"AllowAutoSetup,omitempty" xml:"AllowAutoSetup,omitempty" require:"true"`
	AllowBufferCount   *int     `json:"AllowBufferCount,omitempty" xml:"AllowBufferCount,omitempty" require:"true"`
}

func (s GetDesktopGroupDetailResponseDesktops) String() string {
	return tea.Prettify(s)
}

func (s GetDesktopGroupDetailResponseDesktops) GoString() string {
	return s.String()
}

func (s *GetDesktopGroupDetailResponseDesktops) SetDesktopGroupId(v string) *GetDesktopGroupDetailResponseDesktops {
	s.DesktopGroupId = &v
	return s
}

func (s *GetDesktopGroupDetailResponseDesktops) SetDesktopGroupName(v string) *GetDesktopGroupDetailResponseDesktops {
	s.DesktopGroupName = &v
	return s
}

func (s *GetDesktopGroupDetailResponseDesktops) SetOfficeSiteId(v string) *GetDesktopGroupDetailResponseDesktops {
	s.OfficeSiteId = &v
	return s
}

func (s *GetDesktopGroupDetailResponseDesktops) SetOfficeSiteName(v string) *GetDesktopGroupDetailResponseDesktops {
	s.OfficeSiteName = &v
	return s
}

func (s *GetDesktopGroupDetailResponseDesktops) SetOfficeSiteType(v string) *GetDesktopGroupDetailResponseDesktops {
	s.OfficeSiteType = &v
	return s
}

func (s *GetDesktopGroupDetailResponseDesktops) SetPolicyGroupId(v string) *GetDesktopGroupDetailResponseDesktops {
	s.PolicyGroupId = &v
	return s
}

func (s *GetDesktopGroupDetailResponseDesktops) SetPolicyGroupName(v string) *GetDesktopGroupDetailResponseDesktops {
	s.PolicyGroupName = &v
	return s
}

func (s *GetDesktopGroupDetailResponseDesktops) SetOwnBundleId(v string) *GetDesktopGroupDetailResponseDesktops {
	s.OwnBundleId = &v
	return s
}

func (s *GetDesktopGroupDetailResponseDesktops) SetOwnBundleName(v string) *GetDesktopGroupDetailResponseDesktops {
	s.OwnBundleName = &v
	return s
}

func (s *GetDesktopGroupDetailResponseDesktops) SetPayType(v string) *GetDesktopGroupDetailResponseDesktops {
	s.PayType = &v
	return s
}

func (s *GetDesktopGroupDetailResponseDesktops) SetExpiredTime(v string) *GetDesktopGroupDetailResponseDesktops {
	s.ExpiredTime = &v
	return s
}

func (s *GetDesktopGroupDetailResponseDesktops) SetCreationTime(v string) *GetDesktopGroupDetailResponseDesktops {
	s.CreationTime = &v
	return s
}

func (s *GetDesktopGroupDetailResponseDesktops) SetDirectoryId(v string) *GetDesktopGroupDetailResponseDesktops {
	s.DirectoryId = &v
	return s
}

func (s *GetDesktopGroupDetailResponseDesktops) SetDirectoryType(v string) *GetDesktopGroupDetailResponseDesktops {
	s.DirectoryType = &v
	return s
}

func (s *GetDesktopGroupDetailResponseDesktops) SetCpu(v int) *GetDesktopGroupDetailResponseDesktops {
	s.Cpu = &v
	return s
}

func (s *GetDesktopGroupDetailResponseDesktops) SetMemory(v int64) *GetDesktopGroupDetailResponseDesktops {
	s.Memory = &v
	return s
}

func (s *GetDesktopGroupDetailResponseDesktops) SetGpuCount(v float32) *GetDesktopGroupDetailResponseDesktops {
	s.GpuCount = &v
	return s
}

func (s *GetDesktopGroupDetailResponseDesktops) SetGpuSpec(v string) *GetDesktopGroupDetailResponseDesktops {
	s.GpuSpec = &v
	return s
}

func (s *GetDesktopGroupDetailResponseDesktops) SetSystemDiskCategory(v string) *GetDesktopGroupDetailResponseDesktops {
	s.SystemDiskCategory = &v
	return s
}

func (s *GetDesktopGroupDetailResponseDesktops) SetSystemDiskSize(v int) *GetDesktopGroupDetailResponseDesktops {
	s.SystemDiskSize = &v
	return s
}

func (s *GetDesktopGroupDetailResponseDesktops) SetDataDiskCategory(v string) *GetDesktopGroupDetailResponseDesktops {
	s.DataDiskCategory = &v
	return s
}

func (s *GetDesktopGroupDetailResponseDesktops) SetDataDiskSize(v string) *GetDesktopGroupDetailResponseDesktops {
	s.DataDiskSize = &v
	return s
}

func (s *GetDesktopGroupDetailResponseDesktops) SetCreator(v string) *GetDesktopGroupDetailResponseDesktops {
	s.Creator = &v
	return s
}

func (s *GetDesktopGroupDetailResponseDesktops) SetComments(v string) *GetDesktopGroupDetailResponseDesktops {
	s.Comments = &v
	return s
}

func (s *GetDesktopGroupDetailResponseDesktops) SetKeepDuration(v int64) *GetDesktopGroupDetailResponseDesktops {
	s.KeepDuration = &v
	return s
}

func (s *GetDesktopGroupDetailResponseDesktops) SetMinDesktopsCount(v int) *GetDesktopGroupDetailResponseDesktops {
	s.MinDesktopsCount = &v
	return s
}

func (s *GetDesktopGroupDetailResponseDesktops) SetMaxDesktopsCount(v int) *GetDesktopGroupDetailResponseDesktops {
	s.MaxDesktopsCount = &v
	return s
}

func (s *GetDesktopGroupDetailResponseDesktops) SetResType(v int) *GetDesktopGroupDetailResponseDesktops {
	s.ResType = &v
	return s
}

func (s *GetDesktopGroupDetailResponseDesktops) SetAllowAutoSetup(v int) *GetDesktopGroupDetailResponseDesktops {
	s.AllowAutoSetup = &v
	return s
}

func (s *GetDesktopGroupDetailResponseDesktops) SetAllowBufferCount(v int) *GetDesktopGroupDetailResponseDesktops {
	s.AllowBufferCount = &v
	return s
}

type ModifyDesktopGroupRequest struct {
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopGroupId   *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty" require:"true"`
	OwnBundleId      *string `json:"OwnBundleId,omitempty" xml:"OwnBundleId,omitempty"`
	PolicyGroupId    *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	DesktopGroupName *string `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty"`
	ScaleStrategyId  *string `json:"ScaleStrategyId,omitempty" xml:"ScaleStrategyId,omitempty"`
	KeepDuration     *int64  `json:"KeepDuration,omitempty" xml:"KeepDuration,omitempty"`
	Comments         *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	MinDesktopsCount *int    `json:"MinDesktopsCount,omitempty" xml:"MinDesktopsCount,omitempty"`
	MaxDesktopsCount *int    `json:"MaxDesktopsCount,omitempty" xml:"MaxDesktopsCount,omitempty"`
	AllowAutoSetup   *int    `json:"AllowAutoSetup,omitempty" xml:"AllowAutoSetup,omitempty"`
	AllowBufferCount *int    `json:"AllowBufferCount,omitempty" xml:"AllowBufferCount,omitempty"`
}

func (s ModifyDesktopGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyDesktopGroupRequest) SetRegionId(v string) *ModifyDesktopGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetDesktopGroupId(v string) *ModifyDesktopGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetOwnBundleId(v string) *ModifyDesktopGroupRequest {
	s.OwnBundleId = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetPolicyGroupId(v string) *ModifyDesktopGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetDesktopGroupName(v string) *ModifyDesktopGroupRequest {
	s.DesktopGroupName = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetScaleStrategyId(v string) *ModifyDesktopGroupRequest {
	s.ScaleStrategyId = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetKeepDuration(v int64) *ModifyDesktopGroupRequest {
	s.KeepDuration = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetComments(v string) *ModifyDesktopGroupRequest {
	s.Comments = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetMinDesktopsCount(v int) *ModifyDesktopGroupRequest {
	s.MinDesktopsCount = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetMaxDesktopsCount(v int) *ModifyDesktopGroupRequest {
	s.MaxDesktopsCount = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetAllowAutoSetup(v int) *ModifyDesktopGroupRequest {
	s.AllowAutoSetup = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetAllowBufferCount(v int) *ModifyDesktopGroupRequest {
	s.AllowBufferCount = &v
	return s
}

type ModifyDesktopGroupResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyDesktopGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyDesktopGroupResponse) SetRequestId(v string) *ModifyDesktopGroupResponse {
	s.RequestId = &v
	return s
}

type RemoveUserFromDesktopGroupRequest struct {
	RegionId       *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopGroupId *string   `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty" require:"true"`
	EndUserIds     []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" require:"true" type:"Repeated"`
}

func (s RemoveUserFromDesktopGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromDesktopGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserFromDesktopGroupRequest) SetRegionId(v string) *RemoveUserFromDesktopGroupRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveUserFromDesktopGroupRequest) SetDesktopGroupId(v string) *RemoveUserFromDesktopGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *RemoveUserFromDesktopGroupRequest) SetEndUserIds(v []*string) *RemoveUserFromDesktopGroupRequest {
	s.EndUserIds = v
	return s
}

type RemoveUserFromDesktopGroupResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s RemoveUserFromDesktopGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromDesktopGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveUserFromDesktopGroupResponse) SetRequestId(v string) *RemoveUserFromDesktopGroupResponse {
	s.RequestId = &v
	return s
}

type DeleteScaleStrategyRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ScaleStrategyId *string `json:"ScaleStrategyId,omitempty" xml:"ScaleStrategyId,omitempty" require:"true"`
}

func (s DeleteScaleStrategyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScaleStrategyRequest) GoString() string {
	return s.String()
}

func (s *DeleteScaleStrategyRequest) SetRegionId(v string) *DeleteScaleStrategyRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteScaleStrategyRequest) SetScaleStrategyId(v string) *DeleteScaleStrategyRequest {
	s.ScaleStrategyId = &v
	return s
}

type DeleteScaleStrategyResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteScaleStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteScaleStrategyResponse) GoString() string {
	return s.String()
}

func (s *DeleteScaleStrategyResponse) SetRequestId(v string) *DeleteScaleStrategyResponse {
	s.RequestId = &v
	return s
}

type DeleteDesktopGroupRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty" require:"true"`
}

func (s DeleteDesktopGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDesktopGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteDesktopGroupRequest) SetRegionId(v string) *DeleteDesktopGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDesktopGroupRequest) SetDesktopGroupId(v string) *DeleteDesktopGroupRequest {
	s.DesktopGroupId = &v
	return s
}

type DeleteDesktopGroupResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteDesktopGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDesktopGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteDesktopGroupResponse) SetRequestId(v string) *DeleteDesktopGroupResponse {
	s.RequestId = &v
	return s
}

type DescribeDesktopGroupsRequest struct {
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	OfficeSiteId     *string   `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	DesktopGroupName *string   `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty"`
	MaxResults       *int      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken        *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	DirectoryId      *string   `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	PolicyGroupId    *string   `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	DesktopGroupId   []*string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty" type:"Repeated"`
	EndUserId        []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
	ChargeType       *string   `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ExpiredTime      *string   `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
}

func (s DescribeDesktopGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopGroupsRequest) SetRegionId(v string) *DescribeDesktopGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetOfficeSiteId(v string) *DescribeDesktopGroupsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetDesktopGroupName(v string) *DescribeDesktopGroupsRequest {
	s.DesktopGroupName = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetMaxResults(v int) *DescribeDesktopGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetNextToken(v string) *DescribeDesktopGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetDirectoryId(v string) *DescribeDesktopGroupsRequest {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetPolicyGroupId(v string) *DescribeDesktopGroupsRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetDesktopGroupId(v []*string) *DescribeDesktopGroupsRequest {
	s.DesktopGroupId = v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetEndUserId(v []*string) *DescribeDesktopGroupsRequest {
	s.EndUserId = v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetChargeType(v string) *DescribeDesktopGroupsRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetExpiredTime(v string) *DescribeDesktopGroupsRequest {
	s.ExpiredTime = &v
	return s
}

type DescribeDesktopGroupsResponse struct {
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	NextToken     *string                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	DesktopGroups []*DescribeDesktopGroupsResponseDesktopGroups `json:"DesktopGroups,omitempty" xml:"DesktopGroups,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDesktopGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopGroupsResponse) SetRequestId(v string) *DescribeDesktopGroupsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDesktopGroupsResponse) SetNextToken(v string) *DescribeDesktopGroupsResponse {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopGroupsResponse) SetDesktopGroups(v []*DescribeDesktopGroupsResponseDesktopGroups) *DescribeDesktopGroupsResponse {
	s.DesktopGroups = v
	return s
}

type DescribeDesktopGroupsResponseDesktopGroups struct {
	DesktopGroupId     *string  `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty" require:"true"`
	DesktopGroupName   *string  `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty" require:"true"`
	OfficeSiteId       *string  `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true"`
	OfficeSiteName     *string  `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty" require:"true"`
	OfficeSiteType     *string  `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty" require:"true"`
	PolicyGroupId      *string  `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty" require:"true"`
	PolicyGroupName    *string  `json:"PolicyGroupName,omitempty" xml:"PolicyGroupName,omitempty" require:"true"`
	OwnBundleId        *string  `json:"OwnBundleId,omitempty" xml:"OwnBundleId,omitempty" require:"true"`
	OwnBundleName      *string  `json:"OwnBundleName,omitempty" xml:"OwnBundleName,omitempty" require:"true"`
	EndUserCount       *int64   `json:"EndUserCount,omitempty" xml:"EndUserCount,omitempty" require:"true"`
	PayType            *string  `json:"PayType,omitempty" xml:"PayType,omitempty" require:"true"`
	ExpiredTime        *string  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty" require:"true"`
	CreateTime         *string  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	DirectoryId        *string  `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" require:"true"`
	DirectoryType      *string  `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty" require:"true"`
	Cpu                *int     `json:"Cpu,omitempty" xml:"Cpu,omitempty" require:"true"`
	Memory             *int64   `json:"Memory,omitempty" xml:"Memory,omitempty" require:"true"`
	GpuCount           *float32 `json:"GpuCount,omitempty" xml:"GpuCount,omitempty" require:"true"`
	GpuSpec            *string  `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty" require:"true"`
	SystemDiskCategory *string  `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty" require:"true"`
	SystemDiskSize     *int     `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty" require:"true"`
	DataDiskCategory   *string  `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty" require:"true"`
	DataDiskSize       *string  `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty" require:"true"`
	Creator            *string  `json:"Creator,omitempty" xml:"Creator,omitempty" require:"true"`
	Comments           *string  `json:"Comments,omitempty" xml:"Comments,omitempty" require:"true"`
	KeepDuration       *int64   `json:"KeepDuration,omitempty" xml:"KeepDuration,omitempty" require:"true"`
	MinDesktopsCount   *int     `json:"MinDesktopsCount,omitempty" xml:"MinDesktopsCount,omitempty" require:"true"`
	MaxDesktopsCount   *int     `json:"MaxDesktopsCount,omitempty" xml:"MaxDesktopsCount,omitempty" require:"true"`
}

func (s DescribeDesktopGroupsResponseDesktopGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopGroupsResponseDesktopGroups) GoString() string {
	return s.String()
}

func (s *DescribeDesktopGroupsResponseDesktopGroups) SetDesktopGroupId(v string) *DescribeDesktopGroupsResponseDesktopGroups {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeDesktopGroupsResponseDesktopGroups) SetDesktopGroupName(v string) *DescribeDesktopGroupsResponseDesktopGroups {
	s.DesktopGroupName = &v
	return s
}

func (s *DescribeDesktopGroupsResponseDesktopGroups) SetOfficeSiteId(v string) *DescribeDesktopGroupsResponseDesktopGroups {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeDesktopGroupsResponseDesktopGroups) SetOfficeSiteName(v string) *DescribeDesktopGroupsResponseDesktopGroups {
	s.OfficeSiteName = &v
	return s
}

func (s *DescribeDesktopGroupsResponseDesktopGroups) SetOfficeSiteType(v string) *DescribeDesktopGroupsResponseDesktopGroups {
	s.OfficeSiteType = &v
	return s
}

func (s *DescribeDesktopGroupsResponseDesktopGroups) SetPolicyGroupId(v string) *DescribeDesktopGroupsResponseDesktopGroups {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeDesktopGroupsResponseDesktopGroups) SetPolicyGroupName(v string) *DescribeDesktopGroupsResponseDesktopGroups {
	s.PolicyGroupName = &v
	return s
}

func (s *DescribeDesktopGroupsResponseDesktopGroups) SetOwnBundleId(v string) *DescribeDesktopGroupsResponseDesktopGroups {
	s.OwnBundleId = &v
	return s
}

func (s *DescribeDesktopGroupsResponseDesktopGroups) SetOwnBundleName(v string) *DescribeDesktopGroupsResponseDesktopGroups {
	s.OwnBundleName = &v
	return s
}

func (s *DescribeDesktopGroupsResponseDesktopGroups) SetEndUserCount(v int64) *DescribeDesktopGroupsResponseDesktopGroups {
	s.EndUserCount = &v
	return s
}

func (s *DescribeDesktopGroupsResponseDesktopGroups) SetPayType(v string) *DescribeDesktopGroupsResponseDesktopGroups {
	s.PayType = &v
	return s
}

func (s *DescribeDesktopGroupsResponseDesktopGroups) SetExpiredTime(v string) *DescribeDesktopGroupsResponseDesktopGroups {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDesktopGroupsResponseDesktopGroups) SetCreateTime(v string) *DescribeDesktopGroupsResponseDesktopGroups {
	s.CreateTime = &v
	return s
}

func (s *DescribeDesktopGroupsResponseDesktopGroups) SetDirectoryId(v string) *DescribeDesktopGroupsResponseDesktopGroups {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDesktopGroupsResponseDesktopGroups) SetDirectoryType(v string) *DescribeDesktopGroupsResponseDesktopGroups {
	s.DirectoryType = &v
	return s
}

func (s *DescribeDesktopGroupsResponseDesktopGroups) SetCpu(v int) *DescribeDesktopGroupsResponseDesktopGroups {
	s.Cpu = &v
	return s
}

func (s *DescribeDesktopGroupsResponseDesktopGroups) SetMemory(v int64) *DescribeDesktopGroupsResponseDesktopGroups {
	s.Memory = &v
	return s
}

func (s *DescribeDesktopGroupsResponseDesktopGroups) SetGpuCount(v float32) *DescribeDesktopGroupsResponseDesktopGroups {
	s.GpuCount = &v
	return s
}

func (s *DescribeDesktopGroupsResponseDesktopGroups) SetGpuSpec(v string) *DescribeDesktopGroupsResponseDesktopGroups {
	s.GpuSpec = &v
	return s
}

func (s *DescribeDesktopGroupsResponseDesktopGroups) SetSystemDiskCategory(v string) *DescribeDesktopGroupsResponseDesktopGroups {
	s.SystemDiskCategory = &v
	return s
}

func (s *DescribeDesktopGroupsResponseDesktopGroups) SetSystemDiskSize(v int) *DescribeDesktopGroupsResponseDesktopGroups {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeDesktopGroupsResponseDesktopGroups) SetDataDiskCategory(v string) *DescribeDesktopGroupsResponseDesktopGroups {
	s.DataDiskCategory = &v
	return s
}

func (s *DescribeDesktopGroupsResponseDesktopGroups) SetDataDiskSize(v string) *DescribeDesktopGroupsResponseDesktopGroups {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeDesktopGroupsResponseDesktopGroups) SetCreator(v string) *DescribeDesktopGroupsResponseDesktopGroups {
	s.Creator = &v
	return s
}

func (s *DescribeDesktopGroupsResponseDesktopGroups) SetComments(v string) *DescribeDesktopGroupsResponseDesktopGroups {
	s.Comments = &v
	return s
}

func (s *DescribeDesktopGroupsResponseDesktopGroups) SetKeepDuration(v int64) *DescribeDesktopGroupsResponseDesktopGroups {
	s.KeepDuration = &v
	return s
}

func (s *DescribeDesktopGroupsResponseDesktopGroups) SetMinDesktopsCount(v int) *DescribeDesktopGroupsResponseDesktopGroups {
	s.MinDesktopsCount = &v
	return s
}

func (s *DescribeDesktopGroupsResponseDesktopGroups) SetMaxDesktopsCount(v int) *DescribeDesktopGroupsResponseDesktopGroups {
	s.MaxDesktopsCount = &v
	return s
}

type DescribeUserConnectionRecordsRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty" require:"true"`
	EndUserId      *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" require:"true"`
	EndUserType    *string `json:"EndUserType,omitempty" xml:"EndUserType,omitempty"`
	MaxResults     *int    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeUserConnectionRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserConnectionRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserConnectionRecordsRequest) SetRegionId(v string) *DescribeUserConnectionRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserConnectionRecordsRequest) SetDesktopGroupId(v string) *DescribeUserConnectionRecordsRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeUserConnectionRecordsRequest) SetEndUserId(v string) *DescribeUserConnectionRecordsRequest {
	s.EndUserId = &v
	return s
}

func (s *DescribeUserConnectionRecordsRequest) SetEndUserType(v string) *DescribeUserConnectionRecordsRequest {
	s.EndUserType = &v
	return s
}

func (s *DescribeUserConnectionRecordsRequest) SetMaxResults(v int) *DescribeUserConnectionRecordsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeUserConnectionRecordsRequest) SetNextToken(v string) *DescribeUserConnectionRecordsRequest {
	s.NextToken = &v
	return s
}

type DescribeUserConnectionRecordsResponse struct {
	RequestId         *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	NextToken         *string                                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	ConnectionRecords []*DescribeUserConnectionRecordsResponseConnectionRecords `json:"ConnectionRecords,omitempty" xml:"ConnectionRecords,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeUserConnectionRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserConnectionRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserConnectionRecordsResponse) SetRequestId(v string) *DescribeUserConnectionRecordsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeUserConnectionRecordsResponse) SetNextToken(v string) *DescribeUserConnectionRecordsResponse {
	s.NextToken = &v
	return s
}

func (s *DescribeUserConnectionRecordsResponse) SetConnectionRecords(v []*DescribeUserConnectionRecordsResponseConnectionRecords) *DescribeUserConnectionRecordsResponse {
	s.ConnectionRecords = v
	return s
}

type DescribeUserConnectionRecordsResponseConnectionRecords struct {
	ConnectionRecordId *string `json:"ConnectionRecordId,omitempty" xml:"ConnectionRecordId,omitempty" require:"true"`
	ConnectStartTime   *string `json:"ConnectStartTime,omitempty" xml:"ConnectStartTime,omitempty" require:"true"`
	ConnectEndTime     *string `json:"ConnectEndTime,omitempty" xml:"ConnectEndTime,omitempty" require:"true"`
	ConnectDuration    *string `json:"ConnectDuration,omitempty" xml:"ConnectDuration,omitempty" require:"true"`
	DesktopId          *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	DesktopName        *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty" require:"true"`
}

func (s DescribeUserConnectionRecordsResponseConnectionRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserConnectionRecordsResponseConnectionRecords) GoString() string {
	return s.String()
}

func (s *DescribeUserConnectionRecordsResponseConnectionRecords) SetConnectionRecordId(v string) *DescribeUserConnectionRecordsResponseConnectionRecords {
	s.ConnectionRecordId = &v
	return s
}

func (s *DescribeUserConnectionRecordsResponseConnectionRecords) SetConnectStartTime(v string) *DescribeUserConnectionRecordsResponseConnectionRecords {
	s.ConnectStartTime = &v
	return s
}

func (s *DescribeUserConnectionRecordsResponseConnectionRecords) SetConnectEndTime(v string) *DescribeUserConnectionRecordsResponseConnectionRecords {
	s.ConnectEndTime = &v
	return s
}

func (s *DescribeUserConnectionRecordsResponseConnectionRecords) SetConnectDuration(v string) *DescribeUserConnectionRecordsResponseConnectionRecords {
	s.ConnectDuration = &v
	return s
}

func (s *DescribeUserConnectionRecordsResponseConnectionRecords) SetDesktopId(v string) *DescribeUserConnectionRecordsResponseConnectionRecords {
	s.DesktopId = &v
	return s
}

func (s *DescribeUserConnectionRecordsResponseConnectionRecords) SetDesktopName(v string) *DescribeUserConnectionRecordsResponseConnectionRecords {
	s.DesktopName = &v
	return s
}

type ModifyScaleStrategyRequest struct {
	RegionId                  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ScaleStrategyId           *string `json:"ScaleStrategyId,omitempty" xml:"ScaleStrategyId,omitempty" require:"true"`
	ScaleStrategyName         *string `json:"ScaleStrategyName,omitempty" xml:"ScaleStrategyName,omitempty"`
	ScaleStrategyType         *string `json:"ScaleStrategyType,omitempty" xml:"ScaleStrategyType,omitempty"`
	MinDesktopsCount          *int    `json:"MinDesktopsCount,omitempty" xml:"MinDesktopsCount,omitempty"`
	MaxDesktopsCount          *int    `json:"MaxDesktopsCount,omitempty" xml:"MaxDesktopsCount,omitempty"`
	MinAvailableDesktopsCount *int    `json:"MinAvailableDesktopsCount,omitempty" xml:"MinAvailableDesktopsCount,omitempty"`
	MaxAvailableDesktopsCount *int    `json:"MaxAvailableDesktopsCount,omitempty" xml:"MaxAvailableDesktopsCount,omitempty"`
	ScaleStep                 *int    `json:"ScaleStep,omitempty" xml:"ScaleStep,omitempty"`
}

func (s ModifyScaleStrategyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyScaleStrategyRequest) GoString() string {
	return s.String()
}

func (s *ModifyScaleStrategyRequest) SetRegionId(v string) *ModifyScaleStrategyRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyScaleStrategyRequest) SetScaleStrategyId(v string) *ModifyScaleStrategyRequest {
	s.ScaleStrategyId = &v
	return s
}

func (s *ModifyScaleStrategyRequest) SetScaleStrategyName(v string) *ModifyScaleStrategyRequest {
	s.ScaleStrategyName = &v
	return s
}

func (s *ModifyScaleStrategyRequest) SetScaleStrategyType(v string) *ModifyScaleStrategyRequest {
	s.ScaleStrategyType = &v
	return s
}

func (s *ModifyScaleStrategyRequest) SetMinDesktopsCount(v int) *ModifyScaleStrategyRequest {
	s.MinDesktopsCount = &v
	return s
}

func (s *ModifyScaleStrategyRequest) SetMaxDesktopsCount(v int) *ModifyScaleStrategyRequest {
	s.MaxDesktopsCount = &v
	return s
}

func (s *ModifyScaleStrategyRequest) SetMinAvailableDesktopsCount(v int) *ModifyScaleStrategyRequest {
	s.MinAvailableDesktopsCount = &v
	return s
}

func (s *ModifyScaleStrategyRequest) SetMaxAvailableDesktopsCount(v int) *ModifyScaleStrategyRequest {
	s.MaxAvailableDesktopsCount = &v
	return s
}

func (s *ModifyScaleStrategyRequest) SetScaleStep(v int) *ModifyScaleStrategyRequest {
	s.ScaleStep = &v
	return s
}

type ModifyScaleStrategyResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyScaleStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyScaleStrategyResponse) GoString() string {
	return s.String()
}

func (s *ModifyScaleStrategyResponse) SetRequestId(v string) *ModifyScaleStrategyResponse {
	s.RequestId = &v
	return s
}

type DescribeScaleStrategysRequest struct {
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ScaleStrategyName *string `json:"ScaleStrategyName,omitempty" xml:"ScaleStrategyName,omitempty"`
	MaxResults        *int    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken         *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeScaleStrategysRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScaleStrategysRequest) GoString() string {
	return s.String()
}

func (s *DescribeScaleStrategysRequest) SetRegionId(v string) *DescribeScaleStrategysRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScaleStrategysRequest) SetScaleStrategyName(v string) *DescribeScaleStrategysRequest {
	s.ScaleStrategyName = &v
	return s
}

func (s *DescribeScaleStrategysRequest) SetMaxResults(v int) *DescribeScaleStrategysRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeScaleStrategysRequest) SetNextToken(v string) *DescribeScaleStrategysRequest {
	s.NextToken = &v
	return s
}

type DescribeScaleStrategysResponse struct {
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	NextToken      *string                                         `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	ScaleStrategys []*DescribeScaleStrategysResponseScaleStrategys `json:"ScaleStrategys,omitempty" xml:"ScaleStrategys,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeScaleStrategysResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScaleStrategysResponse) GoString() string {
	return s.String()
}

func (s *DescribeScaleStrategysResponse) SetRequestId(v string) *DescribeScaleStrategysResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeScaleStrategysResponse) SetNextToken(v string) *DescribeScaleStrategysResponse {
	s.NextToken = &v
	return s
}

func (s *DescribeScaleStrategysResponse) SetScaleStrategys(v []*DescribeScaleStrategysResponseScaleStrategys) *DescribeScaleStrategysResponse {
	s.ScaleStrategys = v
	return s
}

type DescribeScaleStrategysResponseScaleStrategys struct {
	ScaleStrategyId           *string `json:"ScaleStrategyId,omitempty" xml:"ScaleStrategyId,omitempty" require:"true"`
	ScaleStrategyName         *string `json:"ScaleStrategyName,omitempty" xml:"ScaleStrategyName,omitempty" require:"true"`
	ScaleStrategyType         *string `json:"ScaleStrategyType,omitempty" xml:"ScaleStrategyType,omitempty" require:"true"`
	MinDesktopsCount          *int    `json:"MinDesktopsCount,omitempty" xml:"MinDesktopsCount,omitempty" require:"true"`
	MaxDesktopsCount          *int    `json:"MaxDesktopsCount,omitempty" xml:"MaxDesktopsCount,omitempty" require:"true"`
	MinAvailableDesktopsCount *int    `json:"MinAvailableDesktopsCount,omitempty" xml:"MinAvailableDesktopsCount,omitempty" require:"true"`
	MaxAvailableDesktopsCount *int    `json:"MaxAvailableDesktopsCount,omitempty" xml:"MaxAvailableDesktopsCount,omitempty" require:"true"`
	ScaleStep                 *int    `json:"ScaleStep,omitempty" xml:"ScaleStep,omitempty" require:"true"`
}

func (s DescribeScaleStrategysResponseScaleStrategys) String() string {
	return tea.Prettify(s)
}

func (s DescribeScaleStrategysResponseScaleStrategys) GoString() string {
	return s.String()
}

func (s *DescribeScaleStrategysResponseScaleStrategys) SetScaleStrategyId(v string) *DescribeScaleStrategysResponseScaleStrategys {
	s.ScaleStrategyId = &v
	return s
}

func (s *DescribeScaleStrategysResponseScaleStrategys) SetScaleStrategyName(v string) *DescribeScaleStrategysResponseScaleStrategys {
	s.ScaleStrategyName = &v
	return s
}

func (s *DescribeScaleStrategysResponseScaleStrategys) SetScaleStrategyType(v string) *DescribeScaleStrategysResponseScaleStrategys {
	s.ScaleStrategyType = &v
	return s
}

func (s *DescribeScaleStrategysResponseScaleStrategys) SetMinDesktopsCount(v int) *DescribeScaleStrategysResponseScaleStrategys {
	s.MinDesktopsCount = &v
	return s
}

func (s *DescribeScaleStrategysResponseScaleStrategys) SetMaxDesktopsCount(v int) *DescribeScaleStrategysResponseScaleStrategys {
	s.MaxDesktopsCount = &v
	return s
}

func (s *DescribeScaleStrategysResponseScaleStrategys) SetMinAvailableDesktopsCount(v int) *DescribeScaleStrategysResponseScaleStrategys {
	s.MinAvailableDesktopsCount = &v
	return s
}

func (s *DescribeScaleStrategysResponseScaleStrategys) SetMaxAvailableDesktopsCount(v int) *DescribeScaleStrategysResponseScaleStrategys {
	s.MaxAvailableDesktopsCount = &v
	return s
}

func (s *DescribeScaleStrategysResponseScaleStrategys) SetScaleStep(v int) *DescribeScaleStrategysResponseScaleStrategys {
	s.ScaleStep = &v
	return s
}

type AddUserToDesktopGroupRequest struct {
	RegionId       *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopGroupId *string   `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty" require:"true"`
	EndUserIds     []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" require:"true" type:"Repeated"`
	ClientToken    *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s AddUserToDesktopGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUserToDesktopGroupRequest) GoString() string {
	return s.String()
}

func (s *AddUserToDesktopGroupRequest) SetRegionId(v string) *AddUserToDesktopGroupRequest {
	s.RegionId = &v
	return s
}

func (s *AddUserToDesktopGroupRequest) SetDesktopGroupId(v string) *AddUserToDesktopGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *AddUserToDesktopGroupRequest) SetEndUserIds(v []*string) *AddUserToDesktopGroupRequest {
	s.EndUserIds = v
	return s
}

func (s *AddUserToDesktopGroupRequest) SetClientToken(v string) *AddUserToDesktopGroupRequest {
	s.ClientToken = &v
	return s
}

type AddUserToDesktopGroupResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AddUserToDesktopGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUserToDesktopGroupResponse) GoString() string {
	return s.String()
}

func (s *AddUserToDesktopGroupResponse) SetRequestId(v string) *AddUserToDesktopGroupResponse {
	s.RequestId = &v
	return s
}

type DescribePostPaidDesktopBillsRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty" require:"true"`
	BillStartTime  *string `json:"BillStartTime,omitempty" xml:"BillStartTime,omitempty" require:"true"`
	BillEndTime    *string `json:"BillEndTime,omitempty" xml:"BillEndTime,omitempty" require:"true"`
	MaxResults     *int    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribePostPaidDesktopBillsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePostPaidDesktopBillsRequest) GoString() string {
	return s.String()
}

func (s *DescribePostPaidDesktopBillsRequest) SetRegionId(v string) *DescribePostPaidDesktopBillsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePostPaidDesktopBillsRequest) SetDesktopGroupId(v string) *DescribePostPaidDesktopBillsRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribePostPaidDesktopBillsRequest) SetBillStartTime(v string) *DescribePostPaidDesktopBillsRequest {
	s.BillStartTime = &v
	return s
}

func (s *DescribePostPaidDesktopBillsRequest) SetBillEndTime(v string) *DescribePostPaidDesktopBillsRequest {
	s.BillEndTime = &v
	return s
}

func (s *DescribePostPaidDesktopBillsRequest) SetMaxResults(v int) *DescribePostPaidDesktopBillsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribePostPaidDesktopBillsRequest) SetNextToken(v string) *DescribePostPaidDesktopBillsRequest {
	s.NextToken = &v
	return s
}

type DescribePostPaidDesktopBillsResponse struct {
	RequestId                   *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PostPaidDesktopsCount       *int                                         `json:"PostPaidDesktopsCount,omitempty" xml:"PostPaidDesktopsCount,omitempty" require:"true"`
	PostPaidDesktopsTotalAmount *float32                                     `json:"PostPaidDesktopsTotalAmount,omitempty" xml:"PostPaidDesktopsTotalAmount,omitempty" require:"true"`
	PostPaidDesktopsBillsUrl    *string                                      `json:"PostPaidDesktopsBillsUrl,omitempty" xml:"PostPaidDesktopsBillsUrl,omitempty" require:"true"`
	NextToken                   *string                                      `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	Bills                       []*DescribePostPaidDesktopBillsResponseBills `json:"Bills,omitempty" xml:"Bills,omitempty" require:"true" type:"Repeated"`
}

func (s DescribePostPaidDesktopBillsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePostPaidDesktopBillsResponse) GoString() string {
	return s.String()
}

func (s *DescribePostPaidDesktopBillsResponse) SetRequestId(v string) *DescribePostPaidDesktopBillsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponse) SetPostPaidDesktopsCount(v int) *DescribePostPaidDesktopBillsResponse {
	s.PostPaidDesktopsCount = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponse) SetPostPaidDesktopsTotalAmount(v float32) *DescribePostPaidDesktopBillsResponse {
	s.PostPaidDesktopsTotalAmount = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponse) SetPostPaidDesktopsBillsUrl(v string) *DescribePostPaidDesktopBillsResponse {
	s.PostPaidDesktopsBillsUrl = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponse) SetNextToken(v string) *DescribePostPaidDesktopBillsResponse {
	s.NextToken = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponse) SetBills(v []*DescribePostPaidDesktopBillsResponseBills) *DescribePostPaidDesktopBillsResponse {
	s.Bills = v
	return s
}

type DescribePostPaidDesktopBillsResponseBills struct {
	BillId            *string `json:"BillId,omitempty" xml:"BillId,omitempty" require:"true"`
	BillStartTime     *string `json:"BillStartTime,omitempty" xml:"BillStartTime,omitempty" require:"true"`
	BillEndTime       *string `json:"BillEndTime,omitempty" xml:"BillEndTime,omitempty" require:"true"`
	BillType          *string `json:"BillType,omitempty" xml:"BillType,omitempty" require:"true"`
	Product           *string `json:"Product,omitempty" xml:"Product,omitempty" require:"true"`
	ProductDetail     *string `json:"ProductDetail,omitempty" xml:"ProductDetail,omitempty" require:"true"`
	ConsumeTime       *string `json:"ConsumeTime,omitempty" xml:"ConsumeTime,omitempty" require:"true"`
	ConsumeType       *string `json:"ConsumeType,omitempty" xml:"ConsumeType,omitempty" require:"true"`
	ResourceGroupId   *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty" require:"true"`
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty" require:"true"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	ChargeItem        *string `json:"ChargeItem,omitempty" xml:"ChargeItem,omitempty" require:"true"`
	Price             *string `json:"Price,omitempty" xml:"Price,omitempty" require:"true"`
	PriceUnit         *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty" require:"true"`
	Usage             *string `json:"Usage,omitempty" xml:"Usage,omitempty" require:"true"`
	UsageUnit         *string `json:"UsageUnit,omitempty" xml:"UsageUnit,omitempty" require:"true"`
	OriginalPrice     *string `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty" require:"true"`
	DiscountPrice     *string `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty" require:"true"`
	Payment           *string `json:"Payment,omitempty" xml:"Payment,omitempty" require:"true"`
	CashPayment       *string `json:"CashPayment,omitempty" xml:"CashPayment,omitempty" require:"true"`
	GoldNote          *string `json:"GoldNote,omitempty" xml:"GoldNote,omitempty" require:"true"`
}

func (s DescribePostPaidDesktopBillsResponseBills) String() string {
	return tea.Prettify(s)
}

func (s DescribePostPaidDesktopBillsResponseBills) GoString() string {
	return s.String()
}

func (s *DescribePostPaidDesktopBillsResponseBills) SetBillId(v string) *DescribePostPaidDesktopBillsResponseBills {
	s.BillId = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBills) SetBillStartTime(v string) *DescribePostPaidDesktopBillsResponseBills {
	s.BillStartTime = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBills) SetBillEndTime(v string) *DescribePostPaidDesktopBillsResponseBills {
	s.BillEndTime = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBills) SetBillType(v string) *DescribePostPaidDesktopBillsResponseBills {
	s.BillType = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBills) SetProduct(v string) *DescribePostPaidDesktopBillsResponseBills {
	s.Product = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBills) SetProductDetail(v string) *DescribePostPaidDesktopBillsResponseBills {
	s.ProductDetail = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBills) SetConsumeTime(v string) *DescribePostPaidDesktopBillsResponseBills {
	s.ConsumeTime = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBills) SetConsumeType(v string) *DescribePostPaidDesktopBillsResponseBills {
	s.ConsumeType = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBills) SetResourceGroupId(v string) *DescribePostPaidDesktopBillsResponseBills {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBills) SetResourceGroupName(v string) *DescribePostPaidDesktopBillsResponseBills {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBills) SetInstanceId(v string) *DescribePostPaidDesktopBillsResponseBills {
	s.InstanceId = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBills) SetChargeItem(v string) *DescribePostPaidDesktopBillsResponseBills {
	s.ChargeItem = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBills) SetPrice(v string) *DescribePostPaidDesktopBillsResponseBills {
	s.Price = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBills) SetPriceUnit(v string) *DescribePostPaidDesktopBillsResponseBills {
	s.PriceUnit = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBills) SetUsage(v string) *DescribePostPaidDesktopBillsResponseBills {
	s.Usage = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBills) SetUsageUnit(v string) *DescribePostPaidDesktopBillsResponseBills {
	s.UsageUnit = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBills) SetOriginalPrice(v string) *DescribePostPaidDesktopBillsResponseBills {
	s.OriginalPrice = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBills) SetDiscountPrice(v string) *DescribePostPaidDesktopBillsResponseBills {
	s.DiscountPrice = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBills) SetPayment(v string) *DescribePostPaidDesktopBillsResponseBills {
	s.Payment = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBills) SetCashPayment(v string) *DescribePostPaidDesktopBillsResponseBills {
	s.CashPayment = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBills) SetGoldNote(v string) *DescribePostPaidDesktopBillsResponseBills {
	s.GoldNote = &v
	return s
}

type CreateDesktopGroupRequest struct {
	RegionId                *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	BundleId                *string   `json:"BundleId,omitempty" xml:"BundleId,omitempty" require:"true"`
	OfficeSiteId            *string   `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true"`
	PolicyGroupId           *string   `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	DesktopGroupName        *string   `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty"`
	DirectoryId             *string   `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	ScaleStrategyId         *string   `json:"ScaleStrategyId,omitempty" xml:"ScaleStrategyId,omitempty"`
	VpcId                   *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	DefaultInitDesktopCount *int      `json:"DefaultInitDesktopCount,omitempty" xml:"DefaultInitDesktopCount,omitempty"`
	EndUserIds              []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" require:"true" type:"Repeated"`
	KeepDuration            *int64    `json:"KeepDuration,omitempty" xml:"KeepDuration,omitempty"`
	ChargeType              *string   `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Period                  *int      `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit              *string   `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	OwnType                 *int      `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	AutoPay                 *bool     `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	Comments                *string   `json:"Comments,omitempty" xml:"Comments,omitempty"`
	MinDesktopsCount        *int      `json:"MinDesktopsCount,omitempty" xml:"MinDesktopsCount,omitempty"`
	MaxDesktopsCount        *int      `json:"MaxDesktopsCount,omitempty" xml:"MaxDesktopsCount,omitempty"`
	AllowAutoSetup          *int      `json:"AllowAutoSetup,omitempty" xml:"AllowAutoSetup,omitempty"`
	AllowBufferCount        *int      `json:"AllowBufferCount,omitempty" xml:"AllowBufferCount,omitempty"`
	ClientToken             *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateDesktopGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateDesktopGroupRequest) SetRegionId(v string) *CreateDesktopGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetBundleId(v string) *CreateDesktopGroupRequest {
	s.BundleId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetOfficeSiteId(v string) *CreateDesktopGroupRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetPolicyGroupId(v string) *CreateDesktopGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetDesktopGroupName(v string) *CreateDesktopGroupRequest {
	s.DesktopGroupName = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetDirectoryId(v string) *CreateDesktopGroupRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetScaleStrategyId(v string) *CreateDesktopGroupRequest {
	s.ScaleStrategyId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetVpcId(v string) *CreateDesktopGroupRequest {
	s.VpcId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetDefaultInitDesktopCount(v int) *CreateDesktopGroupRequest {
	s.DefaultInitDesktopCount = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetEndUserIds(v []*string) *CreateDesktopGroupRequest {
	s.EndUserIds = v
	return s
}

func (s *CreateDesktopGroupRequest) SetKeepDuration(v int64) *CreateDesktopGroupRequest {
	s.KeepDuration = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetChargeType(v string) *CreateDesktopGroupRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetPeriod(v int) *CreateDesktopGroupRequest {
	s.Period = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetPeriodUnit(v string) *CreateDesktopGroupRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetOwnType(v int) *CreateDesktopGroupRequest {
	s.OwnType = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetAutoPay(v bool) *CreateDesktopGroupRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetComments(v string) *CreateDesktopGroupRequest {
	s.Comments = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetMinDesktopsCount(v int) *CreateDesktopGroupRequest {
	s.MinDesktopsCount = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetMaxDesktopsCount(v int) *CreateDesktopGroupRequest {
	s.MaxDesktopsCount = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetAllowAutoSetup(v int) *CreateDesktopGroupRequest {
	s.AllowAutoSetup = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetAllowBufferCount(v int) *CreateDesktopGroupRequest {
	s.AllowBufferCount = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetClientToken(v string) *CreateDesktopGroupRequest {
	s.ClientToken = &v
	return s
}

type CreateDesktopGroupResponse struct {
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DesktopGroupId *string   `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty" require:"true"`
	OrderIds       []*string `json:"OrderIds,omitempty" xml:"OrderIds,omitempty" require:"true" type:"Repeated"`
}

func (s CreateDesktopGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateDesktopGroupResponse) SetRequestId(v string) *CreateDesktopGroupResponse {
	s.RequestId = &v
	return s
}

func (s *CreateDesktopGroupResponse) SetDesktopGroupId(v string) *CreateDesktopGroupResponse {
	s.DesktopGroupId = &v
	return s
}

func (s *CreateDesktopGroupResponse) SetOrderIds(v []*string) *CreateDesktopGroupResponse {
	s.OrderIds = v
	return s
}

type RecreateDesktopGroupRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty" require:"true"`
	OwnBundleId    *string `json:"OwnBundleId,omitempty" xml:"OwnBundleId,omitempty"`
}

func (s RecreateDesktopGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RecreateDesktopGroupRequest) GoString() string {
	return s.String()
}

func (s *RecreateDesktopGroupRequest) SetRegionId(v string) *RecreateDesktopGroupRequest {
	s.RegionId = &v
	return s
}

func (s *RecreateDesktopGroupRequest) SetDesktopGroupId(v string) *RecreateDesktopGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *RecreateDesktopGroupRequest) SetOwnBundleId(v string) *RecreateDesktopGroupRequest {
	s.OwnBundleId = &v
	return s
}

type RecreateDesktopGroupResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s RecreateDesktopGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RecreateDesktopGroupResponse) GoString() string {
	return s.String()
}

func (s *RecreateDesktopGroupResponse) SetRequestId(v string) *RecreateDesktopGroupResponse {
	s.RequestId = &v
	return s
}

type CreateScaleStrategyRequest struct {
	RegionId                  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ScaleStrategyName         *string `json:"ScaleStrategyName,omitempty" xml:"ScaleStrategyName,omitempty"`
	ScaleStrategyType         *string `json:"ScaleStrategyType,omitempty" xml:"ScaleStrategyType,omitempty"`
	PayType                   *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	MinDesktopsCount          *int    `json:"MinDesktopsCount,omitempty" xml:"MinDesktopsCount,omitempty"`
	MaxDesktopsCount          *int    `json:"MaxDesktopsCount,omitempty" xml:"MaxDesktopsCount,omitempty"`
	MinAvailableDesktopsCount *int    `json:"MinAvailableDesktopsCount,omitempty" xml:"MinAvailableDesktopsCount,omitempty"`
	MaxAvailableDesktopsCount *int    `json:"MaxAvailableDesktopsCount,omitempty" xml:"MaxAvailableDesktopsCount,omitempty"`
	ScaleStep                 *int    `json:"ScaleStep,omitempty" xml:"ScaleStep,omitempty"`
	ClientToken               *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateScaleStrategyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScaleStrategyRequest) GoString() string {
	return s.String()
}

func (s *CreateScaleStrategyRequest) SetRegionId(v string) *CreateScaleStrategyRequest {
	s.RegionId = &v
	return s
}

func (s *CreateScaleStrategyRequest) SetScaleStrategyName(v string) *CreateScaleStrategyRequest {
	s.ScaleStrategyName = &v
	return s
}

func (s *CreateScaleStrategyRequest) SetScaleStrategyType(v string) *CreateScaleStrategyRequest {
	s.ScaleStrategyType = &v
	return s
}

func (s *CreateScaleStrategyRequest) SetPayType(v string) *CreateScaleStrategyRequest {
	s.PayType = &v
	return s
}

func (s *CreateScaleStrategyRequest) SetMinDesktopsCount(v int) *CreateScaleStrategyRequest {
	s.MinDesktopsCount = &v
	return s
}

func (s *CreateScaleStrategyRequest) SetMaxDesktopsCount(v int) *CreateScaleStrategyRequest {
	s.MaxDesktopsCount = &v
	return s
}

func (s *CreateScaleStrategyRequest) SetMinAvailableDesktopsCount(v int) *CreateScaleStrategyRequest {
	s.MinAvailableDesktopsCount = &v
	return s
}

func (s *CreateScaleStrategyRequest) SetMaxAvailableDesktopsCount(v int) *CreateScaleStrategyRequest {
	s.MaxAvailableDesktopsCount = &v
	return s
}

func (s *CreateScaleStrategyRequest) SetScaleStep(v int) *CreateScaleStrategyRequest {
	s.ScaleStep = &v
	return s
}

func (s *CreateScaleStrategyRequest) SetClientToken(v string) *CreateScaleStrategyRequest {
	s.ClientToken = &v
	return s
}

type CreateScaleStrategyResponse struct {
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ScaleStrategyId *string `json:"ScaleStrategyId,omitempty" xml:"ScaleStrategyId,omitempty" require:"true"`
}

func (s CreateScaleStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScaleStrategyResponse) GoString() string {
	return s.String()
}

func (s *CreateScaleStrategyResponse) SetRequestId(v string) *CreateScaleStrategyResponse {
	s.RequestId = &v
	return s
}

func (s *CreateScaleStrategyResponse) SetScaleStrategyId(v string) *CreateScaleStrategyResponse {
	s.ScaleStrategyId = &v
	return s
}

type ModifyUserToDesktopGroupRequest struct {
	RegionId       *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopGroupId *string   `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty" require:"true"`
	OldEndUserIds  []*string `json:"OldEndUserIds,omitempty" xml:"OldEndUserIds,omitempty" require:"true" type:"Repeated"`
	NewEndUserIds  []*string `json:"NewEndUserIds,omitempty" xml:"NewEndUserIds,omitempty" require:"true" type:"Repeated"`
}

func (s ModifyUserToDesktopGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserToDesktopGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserToDesktopGroupRequest) SetRegionId(v string) *ModifyUserToDesktopGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyUserToDesktopGroupRequest) SetDesktopGroupId(v string) *ModifyUserToDesktopGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *ModifyUserToDesktopGroupRequest) SetOldEndUserIds(v []*string) *ModifyUserToDesktopGroupRequest {
	s.OldEndUserIds = v
	return s
}

func (s *ModifyUserToDesktopGroupRequest) SetNewEndUserIds(v []*string) *ModifyUserToDesktopGroupRequest {
	s.NewEndUserIds = v
	return s
}

type ModifyUserToDesktopGroupResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyUserToDesktopGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserToDesktopGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserToDesktopGroupResponse) SetRequestId(v string) *ModifyUserToDesktopGroupResponse {
	s.RequestId = &v
	return s
}

type CreateDesktopsLiteRequest struct {
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	BundleId             *string   `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	UserAssignMode       *string   `json:"UserAssignMode,omitempty" xml:"UserAssignMode,omitempty"`
	EndUserId            []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" require:"true" type:"Repeated"`
	Amount               *int      `json:"Amount,omitempty" xml:"Amount,omitempty"`
	EnableInternetAccess *bool     `json:"EnableInternetAccess,omitempty" xml:"EnableInternetAccess,omitempty"`
	Bandwidth            *int      `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	PeriodUnit           *string   `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	Period               *int      `json:"Period,omitempty" xml:"Period,omitempty"`
}

func (s CreateDesktopsLiteRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopsLiteRequest) GoString() string {
	return s.String()
}

func (s *CreateDesktopsLiteRequest) SetRegionId(v string) *CreateDesktopsLiteRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDesktopsLiteRequest) SetBundleId(v string) *CreateDesktopsLiteRequest {
	s.BundleId = &v
	return s
}

func (s *CreateDesktopsLiteRequest) SetUserAssignMode(v string) *CreateDesktopsLiteRequest {
	s.UserAssignMode = &v
	return s
}

func (s *CreateDesktopsLiteRequest) SetEndUserId(v []*string) *CreateDesktopsLiteRequest {
	s.EndUserId = v
	return s
}

func (s *CreateDesktopsLiteRequest) SetAmount(v int) *CreateDesktopsLiteRequest {
	s.Amount = &v
	return s
}

func (s *CreateDesktopsLiteRequest) SetEnableInternetAccess(v bool) *CreateDesktopsLiteRequest {
	s.EnableInternetAccess = &v
	return s
}

func (s *CreateDesktopsLiteRequest) SetBandwidth(v int) *CreateDesktopsLiteRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateDesktopsLiteRequest) SetPeriodUnit(v string) *CreateDesktopsLiteRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateDesktopsLiteRequest) SetPeriod(v int) *CreateDesktopsLiteRequest {
	s.Period = &v
	return s
}

type CreateDesktopsLiteResponse struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	OrderId   *string   `json:"OrderId,omitempty" xml:"OrderId,omitempty" require:"true"`
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true" type:"Repeated"`
}

func (s CreateDesktopsLiteResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopsLiteResponse) GoString() string {
	return s.String()
}

func (s *CreateDesktopsLiteResponse) SetRequestId(v string) *CreateDesktopsLiteResponse {
	s.RequestId = &v
	return s
}

func (s *CreateDesktopsLiteResponse) SetOrderId(v string) *CreateDesktopsLiteResponse {
	s.OrderId = &v
	return s
}

func (s *CreateDesktopsLiteResponse) SetDesktopId(v []*string) *CreateDesktopsLiteResponse {
	s.DesktopId = v
	return s
}

type ModifyOfficeSiteAttributeRequest struct {
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	OfficeSiteId      *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true"`
	DesktopAccessType *string `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	OfficeSiteName    *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
}

func (s ModifyOfficeSiteAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyOfficeSiteAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteAttributeRequest) SetRegionId(v string) *ModifyOfficeSiteAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyOfficeSiteAttributeRequest) SetOfficeSiteId(v string) *ModifyOfficeSiteAttributeRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ModifyOfficeSiteAttributeRequest) SetDesktopAccessType(v string) *ModifyOfficeSiteAttributeRequest {
	s.DesktopAccessType = &v
	return s
}

func (s *ModifyOfficeSiteAttributeRequest) SetOfficeSiteName(v string) *ModifyOfficeSiteAttributeRequest {
	s.OfficeSiteName = &v
	return s
}

type ModifyOfficeSiteAttributeResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyOfficeSiteAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyOfficeSiteAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteAttributeResponse) SetRequestId(v string) *ModifyOfficeSiteAttributeResponse {
	s.RequestId = &v
	return s
}

type CreateServiceLinkedRoleRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s CreateServiceLinkedRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleRequest) SetRegionId(v string) *CreateServiceLinkedRoleRequest {
	s.RegionId = &v
	return s
}

type CreateServiceLinkedRoleResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s CreateServiceLinkedRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponse) SetRequestId(v string) *CreateServiceLinkedRoleResponse {
	s.RequestId = &v
	return s
}

type ModifyOfficeSiteCrossDesktopAccessRequest struct {
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	OfficeSiteId             *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true"`
	EnableCrossDesktopAccess *bool   `json:"EnableCrossDesktopAccess,omitempty" xml:"EnableCrossDesktopAccess,omitempty" require:"true"`
}

func (s ModifyOfficeSiteCrossDesktopAccessRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyOfficeSiteCrossDesktopAccessRequest) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteCrossDesktopAccessRequest) SetRegionId(v string) *ModifyOfficeSiteCrossDesktopAccessRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyOfficeSiteCrossDesktopAccessRequest) SetOfficeSiteId(v string) *ModifyOfficeSiteCrossDesktopAccessRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ModifyOfficeSiteCrossDesktopAccessRequest) SetEnableCrossDesktopAccess(v bool) *ModifyOfficeSiteCrossDesktopAccessRequest {
	s.EnableCrossDesktopAccess = &v
	return s
}

type ModifyOfficeSiteCrossDesktopAccessResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyOfficeSiteCrossDesktopAccessResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyOfficeSiteCrossDesktopAccessResponse) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteCrossDesktopAccessResponse) SetRequestId(v string) *ModifyOfficeSiteCrossDesktopAccessResponse {
	s.RequestId = &v
	return s
}

type GetDesktopUsersRequest struct {
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
}

func (s GetDesktopUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDesktopUsersRequest) GoString() string {
	return s.String()
}

func (s *GetDesktopUsersRequest) SetRegionId(v string) *GetDesktopUsersRequest {
	s.RegionId = &v
	return s
}

func (s *GetDesktopUsersRequest) SetDesktopId(v string) *GetDesktopUsersRequest {
	s.DesktopId = &v
	return s
}

type GetDesktopUsersResponse struct {
	RequestId  *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" require:"true" type:"Repeated"`
}

func (s GetDesktopUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDesktopUsersResponse) GoString() string {
	return s.String()
}

func (s *GetDesktopUsersResponse) SetRequestId(v string) *GetDesktopUsersResponse {
	s.RequestId = &v
	return s
}

func (s *GetDesktopUsersResponse) SetEndUserIds(v []*string) *GetDesktopUsersResponse {
	s.EndUserIds = v
	return s
}

type ModifyNetworkPackageEnabledRequest struct {
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	NetworkPackageId *string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty" require:"true"`
	Enabled          *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s ModifyNetworkPackageEnabledRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNetworkPackageEnabledRequest) GoString() string {
	return s.String()
}

func (s *ModifyNetworkPackageEnabledRequest) SetRegionId(v string) *ModifyNetworkPackageEnabledRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyNetworkPackageEnabledRequest) SetNetworkPackageId(v string) *ModifyNetworkPackageEnabledRequest {
	s.NetworkPackageId = &v
	return s
}

func (s *ModifyNetworkPackageEnabledRequest) SetEnabled(v bool) *ModifyNetworkPackageEnabledRequest {
	s.Enabled = &v
	return s
}

type ModifyNetworkPackageEnabledResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyNetworkPackageEnabledResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyNetworkPackageEnabledResponse) GoString() string {
	return s.String()
}

func (s *ModifyNetworkPackageEnabledResponse) SetRequestId(v string) *ModifyNetworkPackageEnabledResponse {
	s.RequestId = &v
	return s
}

type ResetNASDefaultMountTargetRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty" require:"true"`
}

func (s ResetNASDefaultMountTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetNASDefaultMountTargetRequest) GoString() string {
	return s.String()
}

func (s *ResetNASDefaultMountTargetRequest) SetRegionId(v string) *ResetNASDefaultMountTargetRequest {
	s.RegionId = &v
	return s
}

func (s *ResetNASDefaultMountTargetRequest) SetFileSystemId(v string) *ResetNASDefaultMountTargetRequest {
	s.FileSystemId = &v
	return s
}

type ResetNASDefaultMountTargetResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ResetNASDefaultMountTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetNASDefaultMountTargetResponse) GoString() string {
	return s.String()
}

func (s *ResetNASDefaultMountTargetResponse) SetRequestId(v string) *ResetNASDefaultMountTargetResponse {
	s.RequestId = &v
	return s
}

type DescribeCensRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	PageSize   *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeCensRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensRequest) GoString() string {
	return s.String()
}

func (s *DescribeCensRequest) SetRegionId(v string) *DescribeCensRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCensRequest) SetPageSize(v int) *DescribeCensRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCensRequest) SetPageNumber(v int) *DescribeCensRequest {
	s.PageNumber = &v
	return s
}

type DescribeCensResponse struct {
	RequestId  *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount *int                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageNumber *int                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                        `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	Cens       []*DescribeCensResponseCens `json:"Cens,omitempty" xml:"Cens,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeCensResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensResponse) GoString() string {
	return s.String()
}

func (s *DescribeCensResponse) SetRequestId(v string) *DescribeCensResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeCensResponse) SetTotalCount(v int) *DescribeCensResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeCensResponse) SetPageNumber(v int) *DescribeCensResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeCensResponse) SetPageSize(v int) *DescribeCensResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeCensResponse) SetCens(v []*DescribeCensResponseCens) *DescribeCensResponse {
	s.Cens = v
	return s
}

type DescribeCensResponseCens struct {
	CenId           *string                               `json:"CenId,omitempty" xml:"CenId,omitempty" require:"true"`
	Name            *string                               `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Description     *string                               `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	ProtectionLevel *string                               `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty" require:"true"`
	Status          *string                               `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	CreationTime    *string                               `json:"CreationTime,omitempty" xml:"CreationTime,omitempty" require:"true"`
	Ipv6Level       *string                               `json:"Ipv6Level,omitempty" xml:"Ipv6Level,omitempty" require:"true"`
	Tags            []*DescribeCensResponseCensTags       `json:"Tags,omitempty" xml:"Tags,omitempty" require:"true" type:"Repeated"`
	PackageIds      []*DescribeCensResponseCensPackageIds `json:"PackageIds,omitempty" xml:"PackageIds,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeCensResponseCens) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensResponseCens) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseCens) SetCenId(v string) *DescribeCensResponseCens {
	s.CenId = &v
	return s
}

func (s *DescribeCensResponseCens) SetName(v string) *DescribeCensResponseCens {
	s.Name = &v
	return s
}

func (s *DescribeCensResponseCens) SetDescription(v string) *DescribeCensResponseCens {
	s.Description = &v
	return s
}

func (s *DescribeCensResponseCens) SetProtectionLevel(v string) *DescribeCensResponseCens {
	s.ProtectionLevel = &v
	return s
}

func (s *DescribeCensResponseCens) SetStatus(v string) *DescribeCensResponseCens {
	s.Status = &v
	return s
}

func (s *DescribeCensResponseCens) SetCreationTime(v string) *DescribeCensResponseCens {
	s.CreationTime = &v
	return s
}

func (s *DescribeCensResponseCens) SetIpv6Level(v string) *DescribeCensResponseCens {
	s.Ipv6Level = &v
	return s
}

func (s *DescribeCensResponseCens) SetTags(v []*DescribeCensResponseCensTags) *DescribeCensResponseCens {
	s.Tags = v
	return s
}

func (s *DescribeCensResponseCens) SetPackageIds(v []*DescribeCensResponseCensPackageIds) *DescribeCensResponseCens {
	s.PackageIds = v
	return s
}

type DescribeCensResponseCensTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty" require:"true"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s DescribeCensResponseCensTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensResponseCensTags) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseCensTags) SetKey(v string) *DescribeCensResponseCensTags {
	s.Key = &v
	return s
}

func (s *DescribeCensResponseCensTags) SetValue(v string) *DescribeCensResponseCensTags {
	s.Value = &v
	return s
}

type DescribeCensResponseCensPackageIds struct {
	PackageId *string `json:"PackageId,omitempty" xml:"PackageId,omitempty" require:"true"`
}

func (s DescribeCensResponseCensPackageIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensResponseCensPackageIds) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseCensPackageIds) SetPackageId(v string) *DescribeCensResponseCensPackageIds {
	s.PackageId = &v
	return s
}

type CheckUserInSecurityCenterWhiteListRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s CheckUserInSecurityCenterWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckUserInSecurityCenterWhiteListRequest) GoString() string {
	return s.String()
}

func (s *CheckUserInSecurityCenterWhiteListRequest) SetRegionId(v string) *CheckUserInSecurityCenterWhiteListRequest {
	s.RegionId = &v
	return s
}

type CheckUserInSecurityCenterWhiteListResponse struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	InWhiteList *bool   `json:"InWhiteList,omitempty" xml:"InWhiteList,omitempty" require:"true"`
}

func (s CheckUserInSecurityCenterWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckUserInSecurityCenterWhiteListResponse) GoString() string {
	return s.String()
}

func (s *CheckUserInSecurityCenterWhiteListResponse) SetRequestId(v string) *CheckUserInSecurityCenterWhiteListResponse {
	s.RequestId = &v
	return s
}

func (s *CheckUserInSecurityCenterWhiteListResponse) SetInWhiteList(v bool) *CheckUserInSecurityCenterWhiteListResponse {
	s.InWhiteList = &v
	return s
}

type AddUserToSecurityCenterWhiteListRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s AddUserToSecurityCenterWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUserToSecurityCenterWhiteListRequest) GoString() string {
	return s.String()
}

func (s *AddUserToSecurityCenterWhiteListRequest) SetRegionId(v string) *AddUserToSecurityCenterWhiteListRequest {
	s.RegionId = &v
	return s
}

type AddUserToSecurityCenterWhiteListResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AddUserToSecurityCenterWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUserToSecurityCenterWhiteListResponse) GoString() string {
	return s.String()
}

func (s *AddUserToSecurityCenterWhiteListResponse) SetRequestId(v string) *AddUserToSecurityCenterWhiteListResponse {
	s.RequestId = &v
	return s
}

type DescribeVulOverviewRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s DescribeVulOverviewRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulOverviewRequest) GoString() string {
	return s.String()
}

func (s *DescribeVulOverviewRequest) SetRegionId(v string) *DescribeVulOverviewRequest {
	s.RegionId = &v
	return s
}

type DescribeVulOverviewResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	AsapCount  *int    `json:"AsapCount,omitempty" xml:"AsapCount,omitempty" require:"true"`
	LaterCount *int    `json:"LaterCount,omitempty" xml:"LaterCount,omitempty" require:"true"`
	NntfCount  *int    `json:"NntfCount,omitempty" xml:"NntfCount,omitempty" require:"true"`
}

func (s DescribeVulOverviewResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulOverviewResponse) GoString() string {
	return s.String()
}

func (s *DescribeVulOverviewResponse) SetRequestId(v string) *DescribeVulOverviewResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeVulOverviewResponse) SetAsapCount(v int) *DescribeVulOverviewResponse {
	s.AsapCount = &v
	return s
}

func (s *DescribeVulOverviewResponse) SetLaterCount(v int) *DescribeVulOverviewResponse {
	s.LaterCount = &v
	return s
}

func (s *DescribeVulOverviewResponse) SetNntfCount(v int) *DescribeVulOverviewResponse {
	s.NntfCount = &v
	return s
}

type DescribeSuspEventOverviewRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s DescribeSuspEventOverviewRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventOverviewRequest) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventOverviewRequest) SetRegionId(v string) *DescribeSuspEventOverviewRequest {
	s.RegionId = &v
	return s
}

type DescribeSuspEventOverviewResponse struct {
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SeriousCount    *int    `json:"SeriousCount,omitempty" xml:"SeriousCount,omitempty" require:"true"`
	SuspiciousCount *int    `json:"SuspiciousCount,omitempty" xml:"SuspiciousCount,omitempty" require:"true"`
	RemindCount     *int    `json:"RemindCount,omitempty" xml:"RemindCount,omitempty" require:"true"`
}

func (s DescribeSuspEventOverviewResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventOverviewResponse) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventOverviewResponse) SetRequestId(v string) *DescribeSuspEventOverviewResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeSuspEventOverviewResponse) SetSeriousCount(v int) *DescribeSuspEventOverviewResponse {
	s.SeriousCount = &v
	return s
}

func (s *DescribeSuspEventOverviewResponse) SetSuspiciousCount(v int) *DescribeSuspEventOverviewResponse {
	s.SuspiciousCount = &v
	return s
}

func (s *DescribeSuspEventOverviewResponse) SetRemindCount(v int) *DescribeSuspEventOverviewResponse {
	s.RemindCount = &v
	return s
}

type ModifyOfficeSiteMfaEnabledRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true"`
	MfaEnabled   *bool   `json:"MfaEnabled,omitempty" xml:"MfaEnabled,omitempty" require:"true"`
}

func (s ModifyOfficeSiteMfaEnabledRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyOfficeSiteMfaEnabledRequest) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteMfaEnabledRequest) SetRegionId(v string) *ModifyOfficeSiteMfaEnabledRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyOfficeSiteMfaEnabledRequest) SetOfficeSiteId(v string) *ModifyOfficeSiteMfaEnabledRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ModifyOfficeSiteMfaEnabledRequest) SetMfaEnabled(v bool) *ModifyOfficeSiteMfaEnabledRequest {
	s.MfaEnabled = &v
	return s
}

type ModifyOfficeSiteMfaEnabledResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyOfficeSiteMfaEnabledResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyOfficeSiteMfaEnabledResponse) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteMfaEnabledResponse) SetRequestId(v string) *ModifyOfficeSiteMfaEnabledResponse {
	s.RequestId = &v
	return s
}

type DeleteNASFileSystemsRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	FileSystemId []*string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteNASFileSystemsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNASFileSystemsRequest) GoString() string {
	return s.String()
}

func (s *DeleteNASFileSystemsRequest) SetRegionId(v string) *DeleteNASFileSystemsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteNASFileSystemsRequest) SetFileSystemId(v []*string) *DeleteNASFileSystemsRequest {
	s.FileSystemId = v
	return s
}

type DeleteNASFileSystemsResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteNASFileSystemsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNASFileSystemsResponse) GoString() string {
	return s.String()
}

func (s *DeleteNASFileSystemsResponse) SetRequestId(v string) *DeleteNASFileSystemsResponse {
	s.RequestId = &v
	return s
}

type ModifyNASDefaultMountTargetRequest struct {
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	FileSystemId      *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty" require:"true"`
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty" require:"true"`
}

func (s ModifyNASDefaultMountTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNASDefaultMountTargetRequest) GoString() string {
	return s.String()
}

func (s *ModifyNASDefaultMountTargetRequest) SetRegionId(v string) *ModifyNASDefaultMountTargetRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyNASDefaultMountTargetRequest) SetFileSystemId(v string) *ModifyNASDefaultMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyNASDefaultMountTargetRequest) SetMountTargetDomain(v string) *ModifyNASDefaultMountTargetRequest {
	s.MountTargetDomain = &v
	return s
}

type ModifyNASDefaultMountTargetResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyNASDefaultMountTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyNASDefaultMountTargetResponse) GoString() string {
	return s.String()
}

func (s *ModifyNASDefaultMountTargetResponse) SetRequestId(v string) *ModifyNASDefaultMountTargetResponse {
	s.RequestId = &v
	return s
}

type CreateNASFileSystemRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateNASFileSystemRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNASFileSystemRequest) GoString() string {
	return s.String()
}

func (s *CreateNASFileSystemRequest) SetRegionId(v string) *CreateNASFileSystemRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNASFileSystemRequest) SetOfficeSiteId(v string) *CreateNASFileSystemRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateNASFileSystemRequest) SetName(v string) *CreateNASFileSystemRequest {
	s.Name = &v
	return s
}

func (s *CreateNASFileSystemRequest) SetDescription(v string) *CreateNASFileSystemRequest {
	s.Description = &v
	return s
}

type CreateNASFileSystemResponse struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	OfficeSiteId      *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true"`
	FileSystemId      *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty" require:"true"`
	FileSystemName    *string `json:"FileSystemName,omitempty" xml:"FileSystemName,omitempty" require:"true"`
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty" require:"true"`
}

func (s CreateNASFileSystemResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNASFileSystemResponse) GoString() string {
	return s.String()
}

func (s *CreateNASFileSystemResponse) SetRequestId(v string) *CreateNASFileSystemResponse {
	s.RequestId = &v
	return s
}

func (s *CreateNASFileSystemResponse) SetOfficeSiteId(v string) *CreateNASFileSystemResponse {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateNASFileSystemResponse) SetFileSystemId(v string) *CreateNASFileSystemResponse {
	s.FileSystemId = &v
	return s
}

func (s *CreateNASFileSystemResponse) SetFileSystemName(v string) *CreateNASFileSystemResponse {
	s.FileSystemName = &v
	return s
}

func (s *CreateNASFileSystemResponse) SetMountTargetDomain(v string) *CreateNASFileSystemResponse {
	s.MountTargetDomain = &v
	return s
}

type DescribeNASFileSystemsRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	OfficeSiteId *string   `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	FileSystemId []*string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty" type:"Repeated"`
	MaxResults   *int      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeNASFileSystemsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNASFileSystemsRequest) GoString() string {
	return s.String()
}

func (s *DescribeNASFileSystemsRequest) SetRegionId(v string) *DescribeNASFileSystemsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeNASFileSystemsRequest) SetOfficeSiteId(v string) *DescribeNASFileSystemsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeNASFileSystemsRequest) SetFileSystemId(v []*string) *DescribeNASFileSystemsRequest {
	s.FileSystemId = v
	return s
}

func (s *DescribeNASFileSystemsRequest) SetMaxResults(v int) *DescribeNASFileSystemsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeNASFileSystemsRequest) SetNextToken(v string) *DescribeNASFileSystemsRequest {
	s.NextToken = &v
	return s
}

type DescribeNASFileSystemsResponse struct {
	RequestId   *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	NextToken   *string                                      `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	FileSystems []*DescribeNASFileSystemsResponseFileSystems `json:"FileSystems,omitempty" xml:"FileSystems,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeNASFileSystemsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNASFileSystemsResponse) GoString() string {
	return s.String()
}

func (s *DescribeNASFileSystemsResponse) SetRequestId(v string) *DescribeNASFileSystemsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeNASFileSystemsResponse) SetNextToken(v string) *DescribeNASFileSystemsResponse {
	s.NextToken = &v
	return s
}

func (s *DescribeNASFileSystemsResponse) SetFileSystems(v []*DescribeNASFileSystemsResponseFileSystems) *DescribeNASFileSystemsResponse {
	s.FileSystems = v
	return s
}

type DescribeNASFileSystemsResponseFileSystems struct {
	FileSystemId      *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty" require:"true"`
	FileSystemName    *string `json:"FileSystemName,omitempty" xml:"FileSystemName,omitempty" require:"true"`
	OfficeSiteId      *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true"`
	OfficeSiteName    *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty" require:"true"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	FileSystemType    *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty" require:"true"`
	StorageType       *string `json:"StorageType,omitempty" xml:"StorageType,omitempty" require:"true"`
	SupportAcl        *bool   `json:"SupportAcl,omitempty" xml:"SupportAcl,omitempty" require:"true"`
	Capacity          *int64  `json:"Capacity,omitempty" xml:"Capacity,omitempty" require:"true"`
	MeteredSize       *int64  `json:"MeteredSize,omitempty" xml:"MeteredSize,omitempty" require:"true"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ZoneId            *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty" require:"true"`
	FileSystemStatus  *string `json:"FileSystemStatus,omitempty" xml:"FileSystemStatus,omitempty" require:"true"`
	MountTargetStatus *string `json:"MountTargetStatus,omitempty" xml:"MountTargetStatus,omitempty" require:"true"`
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty" require:"true"`
}

func (s DescribeNASFileSystemsResponseFileSystems) String() string {
	return tea.Prettify(s)
}

func (s DescribeNASFileSystemsResponseFileSystems) GoString() string {
	return s.String()
}

func (s *DescribeNASFileSystemsResponseFileSystems) SetFileSystemId(v string) *DescribeNASFileSystemsResponseFileSystems {
	s.FileSystemId = &v
	return s
}

func (s *DescribeNASFileSystemsResponseFileSystems) SetFileSystemName(v string) *DescribeNASFileSystemsResponseFileSystems {
	s.FileSystemName = &v
	return s
}

func (s *DescribeNASFileSystemsResponseFileSystems) SetOfficeSiteId(v string) *DescribeNASFileSystemsResponseFileSystems {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeNASFileSystemsResponseFileSystems) SetOfficeSiteName(v string) *DescribeNASFileSystemsResponseFileSystems {
	s.OfficeSiteName = &v
	return s
}

func (s *DescribeNASFileSystemsResponseFileSystems) SetDescription(v string) *DescribeNASFileSystemsResponseFileSystems {
	s.Description = &v
	return s
}

func (s *DescribeNASFileSystemsResponseFileSystems) SetCreateTime(v string) *DescribeNASFileSystemsResponseFileSystems {
	s.CreateTime = &v
	return s
}

func (s *DescribeNASFileSystemsResponseFileSystems) SetFileSystemType(v string) *DescribeNASFileSystemsResponseFileSystems {
	s.FileSystemType = &v
	return s
}

func (s *DescribeNASFileSystemsResponseFileSystems) SetStorageType(v string) *DescribeNASFileSystemsResponseFileSystems {
	s.StorageType = &v
	return s
}

func (s *DescribeNASFileSystemsResponseFileSystems) SetSupportAcl(v bool) *DescribeNASFileSystemsResponseFileSystems {
	s.SupportAcl = &v
	return s
}

func (s *DescribeNASFileSystemsResponseFileSystems) SetCapacity(v int64) *DescribeNASFileSystemsResponseFileSystems {
	s.Capacity = &v
	return s
}

func (s *DescribeNASFileSystemsResponseFileSystems) SetMeteredSize(v int64) *DescribeNASFileSystemsResponseFileSystems {
	s.MeteredSize = &v
	return s
}

func (s *DescribeNASFileSystemsResponseFileSystems) SetRegionId(v string) *DescribeNASFileSystemsResponseFileSystems {
	s.RegionId = &v
	return s
}

func (s *DescribeNASFileSystemsResponseFileSystems) SetZoneId(v string) *DescribeNASFileSystemsResponseFileSystems {
	s.ZoneId = &v
	return s
}

func (s *DescribeNASFileSystemsResponseFileSystems) SetFileSystemStatus(v string) *DescribeNASFileSystemsResponseFileSystems {
	s.FileSystemStatus = &v
	return s
}

func (s *DescribeNASFileSystemsResponseFileSystems) SetMountTargetStatus(v string) *DescribeNASFileSystemsResponseFileSystems {
	s.MountTargetStatus = &v
	return s
}

func (s *DescribeNASFileSystemsResponseFileSystems) SetMountTargetDomain(v string) *DescribeNASFileSystemsResponseFileSystems {
	s.MountTargetDomain = &v
	return s
}

type StartVirusScanTaskRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	OfficeSiteId []*string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" type:"Repeated"`
	DesktopId    []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
}

func (s StartVirusScanTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s StartVirusScanTaskRequest) GoString() string {
	return s.String()
}

func (s *StartVirusScanTaskRequest) SetRegionId(v string) *StartVirusScanTaskRequest {
	s.RegionId = &v
	return s
}

func (s *StartVirusScanTaskRequest) SetOfficeSiteId(v []*string) *StartVirusScanTaskRequest {
	s.OfficeSiteId = v
	return s
}

func (s *StartVirusScanTaskRequest) SetDesktopId(v []*string) *StartVirusScanTaskRequest {
	s.DesktopId = v
	return s
}

type StartVirusScanTaskResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ScanTaskId *int64  `json:"ScanTaskId,omitempty" xml:"ScanTaskId,omitempty" require:"true"`
}

func (s StartVirusScanTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StartVirusScanTaskResponse) GoString() string {
	return s.String()
}

func (s *StartVirusScanTaskResponse) SetRequestId(v string) *StartVirusScanTaskResponse {
	s.RequestId = &v
	return s
}

func (s *StartVirusScanTaskResponse) SetScanTaskId(v int64) *StartVirusScanTaskResponse {
	s.ScanTaskId = &v
	return s
}

type ModifyADConnectorOfficeSiteRequest struct {
	RegionId            *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	OfficeSiteId        *string   `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true"`
	DomainName          *string   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainUserName      *string   `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty"`
	DomainPassword      *string   `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty"`
	DnsAddress          []*string `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" type:"Repeated"`
	OfficeSiteName      *string   `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	SubDomainDnsAddress []*string `json:"SubDomainDnsAddress,omitempty" xml:"SubDomainDnsAddress,omitempty" type:"Repeated"`
	SubDomainName       *string   `json:"SubDomainName,omitempty" xml:"SubDomainName,omitempty"`
	MfaEnabled          *bool     `json:"MfaEnabled,omitempty" xml:"MfaEnabled,omitempty"`
}

func (s ModifyADConnectorOfficeSiteRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyADConnectorOfficeSiteRequest) GoString() string {
	return s.String()
}

func (s *ModifyADConnectorOfficeSiteRequest) SetRegionId(v string) *ModifyADConnectorOfficeSiteRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetOfficeSiteId(v string) *ModifyADConnectorOfficeSiteRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetDomainName(v string) *ModifyADConnectorOfficeSiteRequest {
	s.DomainName = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetDomainUserName(v string) *ModifyADConnectorOfficeSiteRequest {
	s.DomainUserName = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetDomainPassword(v string) *ModifyADConnectorOfficeSiteRequest {
	s.DomainPassword = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetDnsAddress(v []*string) *ModifyADConnectorOfficeSiteRequest {
	s.DnsAddress = v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetOfficeSiteName(v string) *ModifyADConnectorOfficeSiteRequest {
	s.OfficeSiteName = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetSubDomainDnsAddress(v []*string) *ModifyADConnectorOfficeSiteRequest {
	s.SubDomainDnsAddress = v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetSubDomainName(v string) *ModifyADConnectorOfficeSiteRequest {
	s.SubDomainName = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetMfaEnabled(v bool) *ModifyADConnectorOfficeSiteRequest {
	s.MfaEnabled = &v
	return s
}

type ModifyADConnectorOfficeSiteResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyADConnectorOfficeSiteResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyADConnectorOfficeSiteResponse) GoString() string {
	return s.String()
}

func (s *ModifyADConnectorOfficeSiteResponse) SetRequestId(v string) *ModifyADConnectorOfficeSiteResponse {
	s.RequestId = &v
	return s
}

type DescribeFrontVulPatchListRequest struct {
	RegionId    *string                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	VulInfo     []*DescribeFrontVulPatchListRequestVulInfo `json:"VulInfo,omitempty" xml:"VulInfo,omitempty" type:"Repeated"`
	OperateType *string                                    `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	Type        *string                                    `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
}

func (s DescribeFrontVulPatchListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFrontVulPatchListRequest) GoString() string {
	return s.String()
}

func (s *DescribeFrontVulPatchListRequest) SetRegionId(v string) *DescribeFrontVulPatchListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFrontVulPatchListRequest) SetVulInfo(v []*DescribeFrontVulPatchListRequestVulInfo) *DescribeFrontVulPatchListRequest {
	s.VulInfo = v
	return s
}

func (s *DescribeFrontVulPatchListRequest) SetOperateType(v string) *DescribeFrontVulPatchListRequest {
	s.OperateType = &v
	return s
}

func (s *DescribeFrontVulPatchListRequest) SetType(v string) *DescribeFrontVulPatchListRequest {
	s.Type = &v
	return s
}

type DescribeFrontVulPatchListRequestVulInfo struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	Tag       *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s DescribeFrontVulPatchListRequestVulInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeFrontVulPatchListRequestVulInfo) GoString() string {
	return s.String()
}

func (s *DescribeFrontVulPatchListRequestVulInfo) SetName(v string) *DescribeFrontVulPatchListRequestVulInfo {
	s.Name = &v
	return s
}

func (s *DescribeFrontVulPatchListRequestVulInfo) SetDesktopId(v string) *DescribeFrontVulPatchListRequestVulInfo {
	s.DesktopId = &v
	return s
}

func (s *DescribeFrontVulPatchListRequestVulInfo) SetTag(v string) *DescribeFrontVulPatchListRequestVulInfo {
	s.Tag = &v
	return s
}

type DescribeFrontVulPatchListResponse struct {
	RequestId      *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	FrontPatchList []*DescribeFrontVulPatchListResponseFrontPatchList `json:"FrontPatchList,omitempty" xml:"FrontPatchList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeFrontVulPatchListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFrontVulPatchListResponse) GoString() string {
	return s.String()
}

func (s *DescribeFrontVulPatchListResponse) SetRequestId(v string) *DescribeFrontVulPatchListResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeFrontVulPatchListResponse) SetFrontPatchList(v []*DescribeFrontVulPatchListResponseFrontPatchList) *DescribeFrontVulPatchListResponse {
	s.FrontPatchList = v
	return s
}

type DescribeFrontVulPatchListResponseFrontPatchList struct {
	DesktopId *string                                                     `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	PatchList []*DescribeFrontVulPatchListResponseFrontPatchListPatchList `json:"PatchList,omitempty" xml:"PatchList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeFrontVulPatchListResponseFrontPatchList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFrontVulPatchListResponseFrontPatchList) GoString() string {
	return s.String()
}

func (s *DescribeFrontVulPatchListResponseFrontPatchList) SetDesktopId(v string) *DescribeFrontVulPatchListResponseFrontPatchList {
	s.DesktopId = &v
	return s
}

func (s *DescribeFrontVulPatchListResponseFrontPatchList) SetPatchList(v []*DescribeFrontVulPatchListResponseFrontPatchListPatchList) *DescribeFrontVulPatchListResponseFrontPatchList {
	s.PatchList = v
	return s
}

type DescribeFrontVulPatchListResponseFrontPatchListPatchList struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty" require:"true"`
}

func (s DescribeFrontVulPatchListResponseFrontPatchListPatchList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFrontVulPatchListResponseFrontPatchListPatchList) GoString() string {
	return s.String()
}

func (s *DescribeFrontVulPatchListResponseFrontPatchListPatchList) SetName(v string) *DescribeFrontVulPatchListResponseFrontPatchListPatchList {
	s.Name = &v
	return s
}

func (s *DescribeFrontVulPatchListResponseFrontPatchListPatchList) SetAliasName(v string) *DescribeFrontVulPatchListResponseFrontPatchListPatchList {
	s.AliasName = &v
	return s
}

type DescribeVulDetailsRequest struct {
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
}

func (s DescribeVulDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVulDetailsRequest) SetRegionId(v string) *DescribeVulDetailsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVulDetailsRequest) SetLang(v string) *DescribeVulDetailsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVulDetailsRequest) SetType(v string) *DescribeVulDetailsRequest {
	s.Type = &v
	return s
}

func (s *DescribeVulDetailsRequest) SetName(v string) *DescribeVulDetailsRequest {
	s.Name = &v
	return s
}

func (s *DescribeVulDetailsRequest) SetAliasName(v string) *DescribeVulDetailsRequest {
	s.AliasName = &v
	return s
}

type DescribeVulDetailsResponse struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Cves      []*DescribeVulDetailsResponseCves `json:"Cves,omitempty" xml:"Cves,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeVulDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVulDetailsResponse) SetRequestId(v string) *DescribeVulDetailsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeVulDetailsResponse) SetCves(v []*DescribeVulDetailsResponseCves) *DescribeVulDetailsResponse {
	s.Cves = v
	return s
}

type DescribeVulDetailsResponseCves struct {
	CveId     *string `json:"CveId,omitempty" xml:"CveId,omitempty" require:"true"`
	Title     *string `json:"Title,omitempty" xml:"Title,omitempty" require:"true"`
	CvssScore *string `json:"CvssScore,omitempty" xml:"CvssScore,omitempty" require:"true"`
	Summary   *string `json:"Summary,omitempty" xml:"Summary,omitempty" require:"true"`
}

func (s DescribeVulDetailsResponseCves) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulDetailsResponseCves) GoString() string {
	return s.String()
}

func (s *DescribeVulDetailsResponseCves) SetCveId(v string) *DescribeVulDetailsResponseCves {
	s.CveId = &v
	return s
}

func (s *DescribeVulDetailsResponseCves) SetTitle(v string) *DescribeVulDetailsResponseCves {
	s.Title = &v
	return s
}

func (s *DescribeVulDetailsResponseCves) SetCvssScore(v string) *DescribeVulDetailsResponseCves {
	s.CvssScore = &v
	return s
}

func (s *DescribeVulDetailsResponseCves) SetSummary(v string) *DescribeVulDetailsResponseCves {
	s.Summary = &v
	return s
}

type DescribeSuspEventQuaraFilesRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true"`
	CurrentPage  *int    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize     *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeSuspEventQuaraFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventQuaraFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventQuaraFilesRequest) SetRegionId(v string) *DescribeSuspEventQuaraFilesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesRequest) SetStatus(v string) *DescribeSuspEventQuaraFilesRequest {
	s.Status = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesRequest) SetOfficeSiteId(v string) *DescribeSuspEventQuaraFilesRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesRequest) SetCurrentPage(v int) *DescribeSuspEventQuaraFilesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesRequest) SetPageSize(v int) *DescribeSuspEventQuaraFilesRequest {
	s.PageSize = &v
	return s
}

type DescribeSuspEventQuaraFilesResponse struct {
	RequestId   *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageSize    *int                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	CurrentPage *int                                             `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
	TotalCount  *int                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	QuaraFiles  []*DescribeSuspEventQuaraFilesResponseQuaraFiles `json:"QuaraFiles,omitempty" xml:"QuaraFiles,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeSuspEventQuaraFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventQuaraFilesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventQuaraFilesResponse) SetRequestId(v string) *DescribeSuspEventQuaraFilesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponse) SetPageSize(v int) *DescribeSuspEventQuaraFilesResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponse) SetCurrentPage(v int) *DescribeSuspEventQuaraFilesResponse {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponse) SetTotalCount(v int) *DescribeSuspEventQuaraFilesResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponse) SetQuaraFiles(v []*DescribeSuspEventQuaraFilesResponseQuaraFiles) *DescribeSuspEventQuaraFilesResponse {
	s.QuaraFiles = v
	return s
}

type DescribeSuspEventQuaraFilesResponseQuaraFiles struct {
	EventName   *string `json:"EventName,omitempty" xml:"EventName,omitempty" require:"true"`
	EventType   *string `json:"EventType,omitempty" xml:"EventType,omitempty" require:"true"`
	Id          *int    `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
	DesktopId   *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty" require:"true"`
	Md5         *string `json:"Md5,omitempty" xml:"Md5,omitempty" require:"true"`
	ModifyTime  *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty" require:"true"`
	Path        *string `json:"Path,omitempty" xml:"Path,omitempty" require:"true"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	Tag         *string `json:"Tag,omitempty" xml:"Tag,omitempty" require:"true"`
}

func (s DescribeSuspEventQuaraFilesResponseQuaraFiles) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventQuaraFilesResponseQuaraFiles) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventQuaraFilesResponseQuaraFiles) SetEventName(v string) *DescribeSuspEventQuaraFilesResponseQuaraFiles {
	s.EventName = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseQuaraFiles) SetEventType(v string) *DescribeSuspEventQuaraFilesResponseQuaraFiles {
	s.EventType = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseQuaraFiles) SetId(v int) *DescribeSuspEventQuaraFilesResponseQuaraFiles {
	s.Id = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseQuaraFiles) SetDesktopId(v string) *DescribeSuspEventQuaraFilesResponseQuaraFiles {
	s.DesktopId = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseQuaraFiles) SetDesktopName(v string) *DescribeSuspEventQuaraFilesResponseQuaraFiles {
	s.DesktopName = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseQuaraFiles) SetMd5(v string) *DescribeSuspEventQuaraFilesResponseQuaraFiles {
	s.Md5 = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseQuaraFiles) SetModifyTime(v string) *DescribeSuspEventQuaraFilesResponseQuaraFiles {
	s.ModifyTime = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseQuaraFiles) SetPath(v string) *DescribeSuspEventQuaraFilesResponseQuaraFiles {
	s.Path = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseQuaraFiles) SetStatus(v string) *DescribeSuspEventQuaraFilesResponseQuaraFiles {
	s.Status = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseQuaraFiles) SetTag(v string) *DescribeSuspEventQuaraFilesResponseQuaraFiles {
	s.Tag = &v
	return s
}

type ModifyOperateVulRequest struct {
	RegionId    *string                           `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Type        *string                           `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	VulInfo     []*ModifyOperateVulRequestVulInfo `json:"VulInfo,omitempty" xml:"VulInfo,omitempty" type:"Repeated"`
	OperateType *string                           `json:"OperateType,omitempty" xml:"OperateType,omitempty" require:"true"`
	Reason      *string                           `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s ModifyOperateVulRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyOperateVulRequest) GoString() string {
	return s.String()
}

func (s *ModifyOperateVulRequest) SetRegionId(v string) *ModifyOperateVulRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyOperateVulRequest) SetType(v string) *ModifyOperateVulRequest {
	s.Type = &v
	return s
}

func (s *ModifyOperateVulRequest) SetVulInfo(v []*ModifyOperateVulRequestVulInfo) *ModifyOperateVulRequest {
	s.VulInfo = v
	return s
}

func (s *ModifyOperateVulRequest) SetOperateType(v string) *ModifyOperateVulRequest {
	s.OperateType = &v
	return s
}

func (s *ModifyOperateVulRequest) SetReason(v string) *ModifyOperateVulRequest {
	s.Reason = &v
	return s
}

type ModifyOperateVulRequestVulInfo struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	Tag       *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ModifyOperateVulRequestVulInfo) String() string {
	return tea.Prettify(s)
}

func (s ModifyOperateVulRequestVulInfo) GoString() string {
	return s.String()
}

func (s *ModifyOperateVulRequestVulInfo) SetName(v string) *ModifyOperateVulRequestVulInfo {
	s.Name = &v
	return s
}

func (s *ModifyOperateVulRequestVulInfo) SetDesktopId(v string) *ModifyOperateVulRequestVulInfo {
	s.DesktopId = &v
	return s
}

func (s *ModifyOperateVulRequestVulInfo) SetTag(v string) *ModifyOperateVulRequestVulInfo {
	s.Tag = &v
	return s
}

type ModifyOperateVulResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyOperateVulResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyOperateVulResponse) GoString() string {
	return s.String()
}

func (s *ModifyOperateVulResponse) SetRequestId(v string) *ModifyOperateVulResponse {
	s.RequestId = &v
	return s
}

type AttachCenRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	CenId        *string `json:"CenId,omitempty" xml:"CenId,omitempty" require:"true"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true"`
}

func (s AttachCenRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachCenRequest) GoString() string {
	return s.String()
}

func (s *AttachCenRequest) SetRegionId(v string) *AttachCenRequest {
	s.RegionId = &v
	return s
}

func (s *AttachCenRequest) SetCenId(v string) *AttachCenRequest {
	s.CenId = &v
	return s
}

func (s *AttachCenRequest) SetOfficeSiteId(v string) *AttachCenRequest {
	s.OfficeSiteId = &v
	return s
}

type AttachCenResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AttachCenResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachCenResponse) GoString() string {
	return s.String()
}

func (s *AttachCenResponse) SetRequestId(v string) *AttachCenResponse {
	s.RequestId = &v
	return s
}

type DescribeVulListRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true"`
	Necessity    *string `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	AliasName    *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	Dealed       *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	CurrentPage  *int    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize     *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeVulListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVulListRequest) SetRegionId(v string) *DescribeVulListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVulListRequest) SetLang(v string) *DescribeVulListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVulListRequest) SetType(v string) *DescribeVulListRequest {
	s.Type = &v
	return s
}

func (s *DescribeVulListRequest) SetOfficeSiteId(v string) *DescribeVulListRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeVulListRequest) SetNecessity(v string) *DescribeVulListRequest {
	s.Necessity = &v
	return s
}

func (s *DescribeVulListRequest) SetAliasName(v string) *DescribeVulListRequest {
	s.AliasName = &v
	return s
}

func (s *DescribeVulListRequest) SetDealed(v string) *DescribeVulListRequest {
	s.Dealed = &v
	return s
}

func (s *DescribeVulListRequest) SetCurrentPage(v int) *DescribeVulListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVulListRequest) SetPageSize(v int) *DescribeVulListRequest {
	s.PageSize = &v
	return s
}

type DescribeVulListResponse struct {
	RequestId   *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageSize    *int                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	CurrentPage *int                                 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
	TotalCount  *int                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	VulRecords  []*DescribeVulListResponseVulRecords `json:"VulRecords,omitempty" xml:"VulRecords,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeVulListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVulListResponse) SetRequestId(v string) *DescribeVulListResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeVulListResponse) SetPageSize(v int) *DescribeVulListResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeVulListResponse) SetCurrentPage(v int) *DescribeVulListResponse {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVulListResponse) SetTotalCount(v int) *DescribeVulListResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeVulListResponse) SetVulRecords(v []*DescribeVulListResponseVulRecords) *DescribeVulListResponse {
	s.VulRecords = v
	return s
}

type DescribeVulListResponseVulRecords struct {
	Name              *string                                             `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	AliasName         *string                                             `json:"AliasName,omitempty" xml:"AliasName,omitempty" require:"true"`
	Type              *string                                             `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	FirstTs           *int64                                              `json:"FirstTs,omitempty" xml:"FirstTs,omitempty" require:"true"`
	LastTs            *int64                                              `json:"LastTs,omitempty" xml:"LastTs,omitempty" require:"true"`
	ModifyTs          *int64                                              `json:"ModifyTs,omitempty" xml:"ModifyTs,omitempty" require:"true"`
	RepairTs          *int64                                              `json:"RepairTs,omitempty" xml:"RepairTs,omitempty" require:"true"`
	DesktopId         *string                                             `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	DesktopName       *string                                             `json:"DesktopName,omitempty" xml:"DesktopName,omitempty" require:"true"`
	Necessity         *string                                             `json:"Necessity,omitempty" xml:"Necessity,omitempty" require:"true"`
	OsVersion         *string                                             `json:"OsVersion,omitempty" xml:"OsVersion,omitempty" require:"true"`
	Related           *string                                             `json:"Related,omitempty" xml:"Related,omitempty" require:"true"`
	ResultCode        *string                                             `json:"ResultCode,omitempty" xml:"ResultCode,omitempty" require:"true"`
	ResultMessage     *string                                             `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty" require:"true"`
	Tag               *string                                             `json:"Tag,omitempty" xml:"Tag,omitempty" require:"true"`
	Online            *bool                                               `json:"Online,omitempty" xml:"Online,omitempty" require:"true"`
	Status            *int                                                `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	ExtendContentJson *DescribeVulListResponseVulRecordsExtendContentJson `json:"ExtendContentJson,omitempty" xml:"ExtendContentJson,omitempty" require:"true" type:"Struct"`
}

func (s DescribeVulListResponseVulRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulListResponseVulRecords) GoString() string {
	return s.String()
}

func (s *DescribeVulListResponseVulRecords) SetName(v string) *DescribeVulListResponseVulRecords {
	s.Name = &v
	return s
}

func (s *DescribeVulListResponseVulRecords) SetAliasName(v string) *DescribeVulListResponseVulRecords {
	s.AliasName = &v
	return s
}

func (s *DescribeVulListResponseVulRecords) SetType(v string) *DescribeVulListResponseVulRecords {
	s.Type = &v
	return s
}

func (s *DescribeVulListResponseVulRecords) SetFirstTs(v int64) *DescribeVulListResponseVulRecords {
	s.FirstTs = &v
	return s
}

func (s *DescribeVulListResponseVulRecords) SetLastTs(v int64) *DescribeVulListResponseVulRecords {
	s.LastTs = &v
	return s
}

func (s *DescribeVulListResponseVulRecords) SetModifyTs(v int64) *DescribeVulListResponseVulRecords {
	s.ModifyTs = &v
	return s
}

func (s *DescribeVulListResponseVulRecords) SetRepairTs(v int64) *DescribeVulListResponseVulRecords {
	s.RepairTs = &v
	return s
}

func (s *DescribeVulListResponseVulRecords) SetDesktopId(v string) *DescribeVulListResponseVulRecords {
	s.DesktopId = &v
	return s
}

func (s *DescribeVulListResponseVulRecords) SetDesktopName(v string) *DescribeVulListResponseVulRecords {
	s.DesktopName = &v
	return s
}

func (s *DescribeVulListResponseVulRecords) SetNecessity(v string) *DescribeVulListResponseVulRecords {
	s.Necessity = &v
	return s
}

func (s *DescribeVulListResponseVulRecords) SetOsVersion(v string) *DescribeVulListResponseVulRecords {
	s.OsVersion = &v
	return s
}

func (s *DescribeVulListResponseVulRecords) SetRelated(v string) *DescribeVulListResponseVulRecords {
	s.Related = &v
	return s
}

func (s *DescribeVulListResponseVulRecords) SetResultCode(v string) *DescribeVulListResponseVulRecords {
	s.ResultCode = &v
	return s
}

func (s *DescribeVulListResponseVulRecords) SetResultMessage(v string) *DescribeVulListResponseVulRecords {
	s.ResultMessage = &v
	return s
}

func (s *DescribeVulListResponseVulRecords) SetTag(v string) *DescribeVulListResponseVulRecords {
	s.Tag = &v
	return s
}

func (s *DescribeVulListResponseVulRecords) SetOnline(v bool) *DescribeVulListResponseVulRecords {
	s.Online = &v
	return s
}

func (s *DescribeVulListResponseVulRecords) SetStatus(v int) *DescribeVulListResponseVulRecords {
	s.Status = &v
	return s
}

func (s *DescribeVulListResponseVulRecords) SetExtendContentJson(v *DescribeVulListResponseVulRecordsExtendContentJson) *DescribeVulListResponseVulRecords {
	s.ExtendContentJson = v
	return s
}

type DescribeVulListResponseVulRecordsExtendContentJson struct {
	RpmEntityList []*DescribeVulListResponseVulRecordsExtendContentJsonRpmEntityList `json:"RpmEntityList,omitempty" xml:"RpmEntityList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeVulListResponseVulRecordsExtendContentJson) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulListResponseVulRecordsExtendContentJson) GoString() string {
	return s.String()
}

func (s *DescribeVulListResponseVulRecordsExtendContentJson) SetRpmEntityList(v []*DescribeVulListResponseVulRecordsExtendContentJsonRpmEntityList) *DescribeVulListResponseVulRecordsExtendContentJson {
	s.RpmEntityList = v
	return s
}

type DescribeVulListResponseVulRecordsExtendContentJsonRpmEntityList struct {
	FullVersion *string `json:"FullVersion,omitempty" xml:"FullVersion,omitempty" require:"true"`
	MatchDetail *string `json:"MatchDetail,omitempty" xml:"MatchDetail,omitempty" require:"true"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Path        *string `json:"Path,omitempty" xml:"Path,omitempty" require:"true"`
	UpdateCmd   *string `json:"UpdateCmd,omitempty" xml:"UpdateCmd,omitempty" require:"true"`
}

func (s DescribeVulListResponseVulRecordsExtendContentJsonRpmEntityList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulListResponseVulRecordsExtendContentJsonRpmEntityList) GoString() string {
	return s.String()
}

func (s *DescribeVulListResponseVulRecordsExtendContentJsonRpmEntityList) SetFullVersion(v string) *DescribeVulListResponseVulRecordsExtendContentJsonRpmEntityList {
	s.FullVersion = &v
	return s
}

func (s *DescribeVulListResponseVulRecordsExtendContentJsonRpmEntityList) SetMatchDetail(v string) *DescribeVulListResponseVulRecordsExtendContentJsonRpmEntityList {
	s.MatchDetail = &v
	return s
}

func (s *DescribeVulListResponseVulRecordsExtendContentJsonRpmEntityList) SetName(v string) *DescribeVulListResponseVulRecordsExtendContentJsonRpmEntityList {
	s.Name = &v
	return s
}

func (s *DescribeVulListResponseVulRecordsExtendContentJsonRpmEntityList) SetPath(v string) *DescribeVulListResponseVulRecordsExtendContentJsonRpmEntityList {
	s.Path = &v
	return s
}

func (s *DescribeVulListResponseVulRecordsExtendContentJsonRpmEntityList) SetUpdateCmd(v string) *DescribeVulListResponseVulRecordsExtendContentJsonRpmEntityList {
	s.UpdateCmd = &v
	return s
}

type DescribeOfficeSitesRequest struct {
	RegionId       *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	OfficeSiteType *string   `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	OfficeSiteId   []*string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" type:"Repeated"`
	MaxResults     *int      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeOfficeSitesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOfficeSitesRequest) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesRequest) SetRegionId(v string) *DescribeOfficeSitesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOfficeSitesRequest) SetOfficeSiteType(v string) *DescribeOfficeSitesRequest {
	s.OfficeSiteType = &v
	return s
}

func (s *DescribeOfficeSitesRequest) SetOfficeSiteId(v []*string) *DescribeOfficeSitesRequest {
	s.OfficeSiteId = v
	return s
}

func (s *DescribeOfficeSitesRequest) SetMaxResults(v int) *DescribeOfficeSitesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeOfficeSitesRequest) SetNextToken(v string) *DescribeOfficeSitesRequest {
	s.NextToken = &v
	return s
}

type DescribeOfficeSitesResponse struct {
	NextToken   *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	OfficeSites []*DescribeOfficeSitesResponseOfficeSites `json:"OfficeSites,omitempty" xml:"OfficeSites,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeOfficeSitesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOfficeSitesResponse) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesResponse) SetNextToken(v string) *DescribeOfficeSitesResponse {
	s.NextToken = &v
	return s
}

func (s *DescribeOfficeSitesResponse) SetRequestId(v string) *DescribeOfficeSitesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeOfficeSitesResponse) SetOfficeSites(v []*DescribeOfficeSitesResponseOfficeSites) *DescribeOfficeSitesResponse {
	s.OfficeSites = v
	return s
}

type DescribeOfficeSitesResponseOfficeSites struct {
	OfficeSiteId             *string                                               `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true"`
	Status                   *string                                               `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	OfficeSiteType           *string                                               `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty" require:"true"`
	CreationTime             *string                                               `json:"CreationTime,omitempty" xml:"CreationTime,omitempty" require:"true"`
	Name                     *string                                               `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	VpcId                    *string                                               `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
	CustomSecurityGroupId    *string                                               `json:"CustomSecurityGroupId,omitempty" xml:"CustomSecurityGroupId,omitempty" require:"true"`
	DnsUserName              *string                                               `json:"DnsUserName,omitempty" xml:"DnsUserName,omitempty" require:"true"`
	EnableInternetAccess     *bool                                                 `json:"EnableInternetAccess,omitempty" xml:"EnableInternetAccess,omitempty" require:"true"`
	EnableCrossDesktopAccess *bool                                                 `json:"EnableCrossDesktopAccess,omitempty" xml:"EnableCrossDesktopAccess,omitempty" require:"true"`
	EnableAdminAccess        *bool                                                 `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty" require:"true"`
	DesktopAccessType        *string                                               `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty" require:"true"`
	DesktopVpcEndpoint       *string                                               `json:"DesktopVpcEndpoint,omitempty" xml:"DesktopVpcEndpoint,omitempty" require:"true"`
	TrustPassword            *string                                               `json:"TrustPassword,omitempty" xml:"TrustPassword,omitempty" require:"true"`
	DomainName               *string                                               `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	DomainUserName           *string                                               `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty" require:"true"`
	DomainPassword           *string                                               `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty" require:"true"`
	SubDomainName            *string                                               `json:"SubDomainName,omitempty" xml:"SubDomainName,omitempty" require:"true"`
	MfaEnabled               *bool                                                 `json:"MfaEnabled,omitempty" xml:"MfaEnabled,omitempty" require:"true"`
	SsoEnabled               *bool                                                 `json:"SsoEnabled,omitempty" xml:"SsoEnabled,omitempty" require:"true"`
	Bandwidth                *int                                                  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty" require:"true"`
	CenId                    *string                                               `json:"CenId,omitempty" xml:"CenId,omitempty" require:"true"`
	NetworkPackageId         *string                                               `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty" require:"true"`
	CidrBlock                *string                                               `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty" require:"true"`
	ADConnectors             []*DescribeOfficeSitesResponseOfficeSitesADConnectors `json:"ADConnectors,omitempty" xml:"ADConnectors,omitempty" require:"true" type:"Repeated"`
	Logs                     []*DescribeOfficeSitesResponseOfficeSitesLogs         `json:"Logs,omitempty" xml:"Logs,omitempty" require:"true" type:"Repeated"`
	DnsAddress               []*string                                             `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" require:"true" type:"Repeated"`
	SubDnsAddress            []*string                                             `json:"SubDnsAddress,omitempty" xml:"SubDnsAddress,omitempty" require:"true" type:"Repeated"`
	VSwitchIds               []*string                                             `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" require:"true" type:"Repeated"`
	FileSystemIds            []*string                                             `json:"FileSystemIds,omitempty" xml:"FileSystemIds,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeOfficeSitesResponseOfficeSites) String() string {
	return tea.Prettify(s)
}

func (s DescribeOfficeSitesResponseOfficeSites) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetOfficeSiteId(v string) *DescribeOfficeSitesResponseOfficeSites {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetStatus(v string) *DescribeOfficeSitesResponseOfficeSites {
	s.Status = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetOfficeSiteType(v string) *DescribeOfficeSitesResponseOfficeSites {
	s.OfficeSiteType = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetCreationTime(v string) *DescribeOfficeSitesResponseOfficeSites {
	s.CreationTime = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetName(v string) *DescribeOfficeSitesResponseOfficeSites {
	s.Name = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetVpcId(v string) *DescribeOfficeSitesResponseOfficeSites {
	s.VpcId = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetCustomSecurityGroupId(v string) *DescribeOfficeSitesResponseOfficeSites {
	s.CustomSecurityGroupId = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetDnsUserName(v string) *DescribeOfficeSitesResponseOfficeSites {
	s.DnsUserName = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetEnableInternetAccess(v bool) *DescribeOfficeSitesResponseOfficeSites {
	s.EnableInternetAccess = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetEnableCrossDesktopAccess(v bool) *DescribeOfficeSitesResponseOfficeSites {
	s.EnableCrossDesktopAccess = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetEnableAdminAccess(v bool) *DescribeOfficeSitesResponseOfficeSites {
	s.EnableAdminAccess = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetDesktopAccessType(v string) *DescribeOfficeSitesResponseOfficeSites {
	s.DesktopAccessType = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetDesktopVpcEndpoint(v string) *DescribeOfficeSitesResponseOfficeSites {
	s.DesktopVpcEndpoint = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetTrustPassword(v string) *DescribeOfficeSitesResponseOfficeSites {
	s.TrustPassword = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetDomainName(v string) *DescribeOfficeSitesResponseOfficeSites {
	s.DomainName = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetDomainUserName(v string) *DescribeOfficeSitesResponseOfficeSites {
	s.DomainUserName = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetDomainPassword(v string) *DescribeOfficeSitesResponseOfficeSites {
	s.DomainPassword = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetSubDomainName(v string) *DescribeOfficeSitesResponseOfficeSites {
	s.SubDomainName = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetMfaEnabled(v bool) *DescribeOfficeSitesResponseOfficeSites {
	s.MfaEnabled = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetSsoEnabled(v bool) *DescribeOfficeSitesResponseOfficeSites {
	s.SsoEnabled = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetBandwidth(v int) *DescribeOfficeSitesResponseOfficeSites {
	s.Bandwidth = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetCenId(v string) *DescribeOfficeSitesResponseOfficeSites {
	s.CenId = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetNetworkPackageId(v string) *DescribeOfficeSitesResponseOfficeSites {
	s.NetworkPackageId = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetCidrBlock(v string) *DescribeOfficeSitesResponseOfficeSites {
	s.CidrBlock = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetADConnectors(v []*DescribeOfficeSitesResponseOfficeSitesADConnectors) *DescribeOfficeSitesResponseOfficeSites {
	s.ADConnectors = v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetLogs(v []*DescribeOfficeSitesResponseOfficeSitesLogs) *DescribeOfficeSitesResponseOfficeSites {
	s.Logs = v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetDnsAddress(v []*string) *DescribeOfficeSitesResponseOfficeSites {
	s.DnsAddress = v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetSubDnsAddress(v []*string) *DescribeOfficeSitesResponseOfficeSites {
	s.SubDnsAddress = v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetVSwitchIds(v []*string) *DescribeOfficeSitesResponseOfficeSites {
	s.VSwitchIds = v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSites) SetFileSystemIds(v []*string) *DescribeOfficeSitesResponseOfficeSites {
	s.FileSystemIds = v
	return s
}

type DescribeOfficeSitesResponseOfficeSitesADConnectors struct {
	ADConnectorAddress *string `json:"ADConnectorAddress,omitempty" xml:"ADConnectorAddress,omitempty" require:"true"`
	VSwitchId          *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" require:"true"`
	ConnectorStatus    *string `json:"ConnectorStatus,omitempty" xml:"ConnectorStatus,omitempty" require:"true"`
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty" require:"true"`
}

func (s DescribeOfficeSitesResponseOfficeSitesADConnectors) String() string {
	return tea.Prettify(s)
}

func (s DescribeOfficeSitesResponseOfficeSitesADConnectors) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesResponseOfficeSitesADConnectors) SetADConnectorAddress(v string) *DescribeOfficeSitesResponseOfficeSitesADConnectors {
	s.ADConnectorAddress = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSitesADConnectors) SetVSwitchId(v string) *DescribeOfficeSitesResponseOfficeSitesADConnectors {
	s.VSwitchId = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSitesADConnectors) SetConnectorStatus(v string) *DescribeOfficeSitesResponseOfficeSitesADConnectors {
	s.ConnectorStatus = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSitesADConnectors) SetNetworkInterfaceId(v string) *DescribeOfficeSitesResponseOfficeSitesADConnectors {
	s.NetworkInterfaceId = &v
	return s
}

type DescribeOfficeSitesResponseOfficeSitesLogs struct {
	Level     *string `json:"Level,omitempty" xml:"Level,omitempty" require:"true"`
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty" require:"true"`
	Step      *string `json:"Step,omitempty" xml:"Step,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
}

func (s DescribeOfficeSitesResponseOfficeSitesLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeOfficeSitesResponseOfficeSitesLogs) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesResponseOfficeSitesLogs) SetLevel(v string) *DescribeOfficeSitesResponseOfficeSitesLogs {
	s.Level = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSitesLogs) SetTimeStamp(v string) *DescribeOfficeSitesResponseOfficeSitesLogs {
	s.TimeStamp = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSitesLogs) SetStep(v string) *DescribeOfficeSitesResponseOfficeSitesLogs {
	s.Step = &v
	return s
}

func (s *DescribeOfficeSitesResponseOfficeSitesLogs) SetMessage(v string) *DescribeOfficeSitesResponseOfficeSitesLogs {
	s.Message = &v
	return s
}

type CreateSimpleOfficeSiteRequest struct {
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	CidrBlock            *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty" require:"true"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	Bandwidth            *int    `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	OfficeSiteName       *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	EnableInternetAccess *bool   `json:"EnableInternetAccess,omitempty" xml:"EnableInternetAccess,omitempty"`
	EnableAdminAccess    *bool   `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	DesktopAccessType    *string `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
}

func (s CreateSimpleOfficeSiteRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSimpleOfficeSiteRequest) GoString() string {
	return s.String()
}

func (s *CreateSimpleOfficeSiteRequest) SetRegionId(v string) *CreateSimpleOfficeSiteRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetCidrBlock(v string) *CreateSimpleOfficeSiteRequest {
	s.CidrBlock = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetCenId(v string) *CreateSimpleOfficeSiteRequest {
	s.CenId = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetBandwidth(v int) *CreateSimpleOfficeSiteRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetOfficeSiteName(v string) *CreateSimpleOfficeSiteRequest {
	s.OfficeSiteName = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetEnableInternetAccess(v bool) *CreateSimpleOfficeSiteRequest {
	s.EnableInternetAccess = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetEnableAdminAccess(v bool) *CreateSimpleOfficeSiteRequest {
	s.EnableAdminAccess = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetDesktopAccessType(v string) *CreateSimpleOfficeSiteRequest {
	s.DesktopAccessType = &v
	return s
}

type CreateSimpleOfficeSiteResponse struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true"`
}

func (s CreateSimpleOfficeSiteResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSimpleOfficeSiteResponse) GoString() string {
	return s.String()
}

func (s *CreateSimpleOfficeSiteResponse) SetRequestId(v string) *CreateSimpleOfficeSiteResponse {
	s.RequestId = &v
	return s
}

func (s *CreateSimpleOfficeSiteResponse) SetOfficeSiteId(v string) *CreateSimpleOfficeSiteResponse {
	s.OfficeSiteId = &v
	return s
}

type OperateVulsRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Type         *string   `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	VulName      []*string `json:"VulName,omitempty" xml:"VulName,omitempty" require:"true" type:"Repeated"`
	DesktopId    []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true" type:"Repeated"`
	OperateType  *string   `json:"OperateType,omitempty" xml:"OperateType,omitempty" require:"true"`
	Reason       *string   `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Precondition *int      `json:"Precondition,omitempty" xml:"Precondition,omitempty" require:"true"`
}

func (s OperateVulsRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateVulsRequest) GoString() string {
	return s.String()
}

func (s *OperateVulsRequest) SetRegionId(v string) *OperateVulsRequest {
	s.RegionId = &v
	return s
}

func (s *OperateVulsRequest) SetType(v string) *OperateVulsRequest {
	s.Type = &v
	return s
}

func (s *OperateVulsRequest) SetVulName(v []*string) *OperateVulsRequest {
	s.VulName = v
	return s
}

func (s *OperateVulsRequest) SetDesktopId(v []*string) *OperateVulsRequest {
	s.DesktopId = v
	return s
}

func (s *OperateVulsRequest) SetOperateType(v string) *OperateVulsRequest {
	s.OperateType = &v
	return s
}

func (s *OperateVulsRequest) SetReason(v string) *OperateVulsRequest {
	s.Reason = &v
	return s
}

func (s *OperateVulsRequest) SetPrecondition(v int) *OperateVulsRequest {
	s.Precondition = &v
	return s
}

type OperateVulsResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s OperateVulsResponse) String() string {
	return tea.Prettify(s)
}

func (s OperateVulsResponse) GoString() string {
	return s.String()
}

func (s *OperateVulsResponse) SetRequestId(v string) *OperateVulsResponse {
	s.RequestId = &v
	return s
}

type DescribeScanTaskProgressRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	TaskId   *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeScanTaskProgressRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScanTaskProgressRequest) GoString() string {
	return s.String()
}

func (s *DescribeScanTaskProgressRequest) SetRegionId(v string) *DescribeScanTaskProgressRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScanTaskProgressRequest) SetTaskId(v int64) *DescribeScanTaskProgressRequest {
	s.TaskId = &v
	return s
}

type DescribeScanTaskProgressResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty" require:"true"`
}

func (s DescribeScanTaskProgressResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScanTaskProgressResponse) GoString() string {
	return s.String()
}

func (s *DescribeScanTaskProgressResponse) SetRequestId(v string) *DescribeScanTaskProgressResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeScanTaskProgressResponse) SetCreateTime(v string) *DescribeScanTaskProgressResponse {
	s.CreateTime = &v
	return s
}

func (s *DescribeScanTaskProgressResponse) SetTaskStatus(v string) *DescribeScanTaskProgressResponse {
	s.TaskStatus = &v
	return s
}

type DetachCenRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true"`
}

func (s DetachCenRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachCenRequest) GoString() string {
	return s.String()
}

func (s *DetachCenRequest) SetRegionId(v string) *DetachCenRequest {
	s.RegionId = &v
	return s
}

func (s *DetachCenRequest) SetOfficeSiteId(v string) *DetachCenRequest {
	s.OfficeSiteId = &v
	return s
}

type DetachCenResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DetachCenResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachCenResponse) GoString() string {
	return s.String()
}

func (s *DetachCenResponse) SetRequestId(v string) *DetachCenResponse {
	s.RequestId = &v
	return s
}

type DescribeSecurityEventOperationStatusRequest struct {
	RegionId        *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	TaskId          *int64    `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
	SecurityEventId []*string `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeSecurityEventOperationStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationStatusRequest) SetRegionId(v string) *DescribeSecurityEventOperationStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusRequest) SetTaskId(v int64) *DescribeSecurityEventOperationStatusRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusRequest) SetSecurityEventId(v []*string) *DescribeSecurityEventOperationStatusRequest {
	s.SecurityEventId = v
	return s
}

type DescribeSecurityEventOperationStatusResponse struct {
	RequestId                      *string                                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TaskStatus                     *string                                                                       `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty" require:"true"`
	SecurityEventOperationStatuses []*DescribeSecurityEventOperationStatusResponseSecurityEventOperationStatuses `json:"SecurityEventOperationStatuses,omitempty" xml:"SecurityEventOperationStatuses,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeSecurityEventOperationStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationStatusResponse) SetRequestId(v string) *DescribeSecurityEventOperationStatusResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusResponse) SetTaskStatus(v string) *DescribeSecurityEventOperationStatusResponse {
	s.TaskStatus = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusResponse) SetSecurityEventOperationStatuses(v []*DescribeSecurityEventOperationStatusResponseSecurityEventOperationStatuses) *DescribeSecurityEventOperationStatusResponse {
	s.SecurityEventOperationStatuses = v
	return s
}

type DescribeSecurityEventOperationStatusResponseSecurityEventOperationStatuses struct {
	SecurityEventId *int64  `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty" require:"true"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	ErrorCode       *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty" require:"true"`
}

func (s DescribeSecurityEventOperationStatusResponseSecurityEventOperationStatuses) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationStatusResponseSecurityEventOperationStatuses) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationStatusResponseSecurityEventOperationStatuses) SetSecurityEventId(v int64) *DescribeSecurityEventOperationStatusResponseSecurityEventOperationStatuses {
	s.SecurityEventId = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusResponseSecurityEventOperationStatuses) SetStatus(v string) *DescribeSecurityEventOperationStatusResponseSecurityEventOperationStatuses {
	s.Status = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusResponseSecurityEventOperationStatuses) SetErrorCode(v string) *DescribeSecurityEventOperationStatusResponseSecurityEventOperationStatuses {
	s.ErrorCode = &v
	return s
}

type DescribeAlarmEventStackInfoRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopId  *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	EventName  *string `json:"EventName,omitempty" xml:"EventName,omitempty" require:"true"`
	UniqueInfo *string `json:"UniqueInfo,omitempty" xml:"UniqueInfo,omitempty" require:"true"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeAlarmEventStackInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmEventStackInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlarmEventStackInfoRequest) SetRegionId(v string) *DescribeAlarmEventStackInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlarmEventStackInfoRequest) SetDesktopId(v string) *DescribeAlarmEventStackInfoRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeAlarmEventStackInfoRequest) SetEventName(v string) *DescribeAlarmEventStackInfoRequest {
	s.EventName = &v
	return s
}

func (s *DescribeAlarmEventStackInfoRequest) SetUniqueInfo(v string) *DescribeAlarmEventStackInfoRequest {
	s.UniqueInfo = &v
	return s
}

func (s *DescribeAlarmEventStackInfoRequest) SetLang(v string) *DescribeAlarmEventStackInfoRequest {
	s.Lang = &v
	return s
}

type DescribeAlarmEventStackInfoResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	StackInfo *string `json:"StackInfo,omitempty" xml:"StackInfo,omitempty" require:"true"`
}

func (s DescribeAlarmEventStackInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmEventStackInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlarmEventStackInfoResponse) SetRequestId(v string) *DescribeAlarmEventStackInfoResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeAlarmEventStackInfoResponse) SetStackInfo(v string) *DescribeAlarmEventStackInfoResponse {
	s.StackInfo = &v
	return s
}

type ListOfficeSiteUsersRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Filter       *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults   *int    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListOfficeSiteUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeSiteUsersRequest) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteUsersRequest) SetRegionId(v string) *ListOfficeSiteUsersRequest {
	s.RegionId = &v
	return s
}

func (s *ListOfficeSiteUsersRequest) SetFilter(v string) *ListOfficeSiteUsersRequest {
	s.Filter = &v
	return s
}

func (s *ListOfficeSiteUsersRequest) SetOfficeSiteId(v string) *ListOfficeSiteUsersRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ListOfficeSiteUsersRequest) SetNextToken(v string) *ListOfficeSiteUsersRequest {
	s.NextToken = &v
	return s
}

func (s *ListOfficeSiteUsersRequest) SetMaxResults(v int) *ListOfficeSiteUsersRequest {
	s.MaxResults = &v
	return s
}

type ListOfficeSiteUsersResponse struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	NextToken *string                             `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	Users     []*ListOfficeSiteUsersResponseUsers `json:"Users,omitempty" xml:"Users,omitempty" require:"true" type:"Repeated"`
}

func (s ListOfficeSiteUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeSiteUsersResponse) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteUsersResponse) SetRequestId(v string) *ListOfficeSiteUsersResponse {
	s.RequestId = &v
	return s
}

func (s *ListOfficeSiteUsersResponse) SetNextToken(v string) *ListOfficeSiteUsersResponse {
	s.NextToken = &v
	return s
}

func (s *ListOfficeSiteUsersResponse) SetUsers(v []*ListOfficeSiteUsersResponseUsers) *ListOfficeSiteUsersResponse {
	s.Users = v
	return s
}

type ListOfficeSiteUsersResponseUsers struct {
	EndUser *string `json:"EndUser,omitempty" xml:"EndUser,omitempty" require:"true"`
}

func (s ListOfficeSiteUsersResponseUsers) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeSiteUsersResponseUsers) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteUsersResponseUsers) SetEndUser(v string) *ListOfficeSiteUsersResponseUsers {
	s.EndUser = &v
	return s
}

type DescribeSuspEventsRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	OfficeSiteId    *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true"`
	Dealed          *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	Levels          *string `json:"Levels,omitempty" xml:"Levels,omitempty"`
	ParentEventType *string `json:"ParentEventType,omitempty" xml:"ParentEventType,omitempty"`
	AlarmUniqueInfo *string `json:"AlarmUniqueInfo,omitempty" xml:"AlarmUniqueInfo,omitempty"`
	CurrentPage     *int    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize        *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeSuspEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsRequest) SetRegionId(v string) *DescribeSuspEventsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetLang(v string) *DescribeSuspEventsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetOfficeSiteId(v string) *DescribeSuspEventsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetDealed(v string) *DescribeSuspEventsRequest {
	s.Dealed = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetLevels(v string) *DescribeSuspEventsRequest {
	s.Levels = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetParentEventType(v string) *DescribeSuspEventsRequest {
	s.ParentEventType = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetAlarmUniqueInfo(v string) *DescribeSuspEventsRequest {
	s.AlarmUniqueInfo = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetCurrentPage(v int) *DescribeSuspEventsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetPageSize(v int) *DescribeSuspEventsRequest {
	s.PageSize = &v
	return s
}

type DescribeSuspEventsResponse struct {
	RequestId   *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageSize    *string                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	CurrentPage *int                                    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
	TotalCount  *int                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	SuspEvents  []*DescribeSuspEventsResponseSuspEvents `json:"SuspEvents,omitempty" xml:"SuspEvents,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeSuspEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsResponse) SetRequestId(v string) *DescribeSuspEventsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeSuspEventsResponse) SetPageSize(v string) *DescribeSuspEventsResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeSuspEventsResponse) SetCurrentPage(v int) *DescribeSuspEventsResponse {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSuspEventsResponse) SetTotalCount(v int) *DescribeSuspEventsResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeSuspEventsResponse) SetSuspEvents(v []*DescribeSuspEventsResponseSuspEvents) *DescribeSuspEventsResponse {
	s.SuspEvents = v
	return s
}

type DescribeSuspEventsResponseSuspEvents struct {
	Id                    *int64                                         `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
	DesktopId             *string                                        `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	DesktopName           *string                                        `json:"DesktopName,omitempty" xml:"DesktopName,omitempty" require:"true"`
	LastTime              *string                                        `json:"LastTime,omitempty" xml:"LastTime,omitempty" require:"true"`
	OccurrenceTime        *string                                        `json:"OccurrenceTime,omitempty" xml:"OccurrenceTime,omitempty" require:"true"`
	UniqueInfo            *string                                        `json:"UniqueInfo,omitempty" xml:"UniqueInfo,omitempty" require:"true"`
	Name                  *string                                        `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	EventSubType          *string                                        `json:"EventSubType,omitempty" xml:"EventSubType,omitempty" require:"true"`
	Level                 *string                                        `json:"Level,omitempty" xml:"Level,omitempty" require:"true"`
	EventStatus           *int                                           `json:"EventStatus,omitempty" xml:"EventStatus,omitempty" require:"true"`
	Desc                  *string                                        `json:"Desc,omitempty" xml:"Desc,omitempty" require:"true"`
	OperateMsg            *string                                        `json:"OperateMsg,omitempty" xml:"OperateMsg,omitempty" require:"true"`
	DataSource            *string                                        `json:"DataSource,omitempty" xml:"DataSource,omitempty" require:"true"`
	OperateErrorCode      *string                                        `json:"OperateErrorCode,omitempty" xml:"OperateErrorCode,omitempty" require:"true"`
	CanCancelFault        *bool                                          `json:"CanCancelFault,omitempty" xml:"CanCancelFault,omitempty" require:"true"`
	CanBeDealOnLine       *string                                        `json:"CanBeDealOnLine,omitempty" xml:"CanBeDealOnLine,omitempty" require:"true"`
	AlarmEventType        *string                                        `json:"AlarmEventType,omitempty" xml:"AlarmEventType,omitempty" require:"true"`
	AlarmEventName        *string                                        `json:"AlarmEventName,omitempty" xml:"AlarmEventName,omitempty" require:"true"`
	AlarmUniqueInfo       *string                                        `json:"AlarmUniqueInfo,omitempty" xml:"AlarmUniqueInfo,omitempty" require:"true"`
	AlarmEventNameDisplay *string                                        `json:"AlarmEventNameDisplay,omitempty" xml:"AlarmEventNameDisplay,omitempty" require:"true"`
	AlarmEventTypeDisplay *string                                        `json:"AlarmEventTypeDisplay,omitempty" xml:"AlarmEventTypeDisplay,omitempty" require:"true"`
	Details               []*DescribeSuspEventsResponseSuspEventsDetails `json:"Details,omitempty" xml:"Details,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeSuspEventsResponseSuspEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventsResponseSuspEvents) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsResponseSuspEvents) SetId(v int64) *DescribeSuspEventsResponseSuspEvents {
	s.Id = &v
	return s
}

func (s *DescribeSuspEventsResponseSuspEvents) SetDesktopId(v string) *DescribeSuspEventsResponseSuspEvents {
	s.DesktopId = &v
	return s
}

func (s *DescribeSuspEventsResponseSuspEvents) SetDesktopName(v string) *DescribeSuspEventsResponseSuspEvents {
	s.DesktopName = &v
	return s
}

func (s *DescribeSuspEventsResponseSuspEvents) SetLastTime(v string) *DescribeSuspEventsResponseSuspEvents {
	s.LastTime = &v
	return s
}

func (s *DescribeSuspEventsResponseSuspEvents) SetOccurrenceTime(v string) *DescribeSuspEventsResponseSuspEvents {
	s.OccurrenceTime = &v
	return s
}

func (s *DescribeSuspEventsResponseSuspEvents) SetUniqueInfo(v string) *DescribeSuspEventsResponseSuspEvents {
	s.UniqueInfo = &v
	return s
}

func (s *DescribeSuspEventsResponseSuspEvents) SetName(v string) *DescribeSuspEventsResponseSuspEvents {
	s.Name = &v
	return s
}

func (s *DescribeSuspEventsResponseSuspEvents) SetEventSubType(v string) *DescribeSuspEventsResponseSuspEvents {
	s.EventSubType = &v
	return s
}

func (s *DescribeSuspEventsResponseSuspEvents) SetLevel(v string) *DescribeSuspEventsResponseSuspEvents {
	s.Level = &v
	return s
}

func (s *DescribeSuspEventsResponseSuspEvents) SetEventStatus(v int) *DescribeSuspEventsResponseSuspEvents {
	s.EventStatus = &v
	return s
}

func (s *DescribeSuspEventsResponseSuspEvents) SetDesc(v string) *DescribeSuspEventsResponseSuspEvents {
	s.Desc = &v
	return s
}

func (s *DescribeSuspEventsResponseSuspEvents) SetOperateMsg(v string) *DescribeSuspEventsResponseSuspEvents {
	s.OperateMsg = &v
	return s
}

func (s *DescribeSuspEventsResponseSuspEvents) SetDataSource(v string) *DescribeSuspEventsResponseSuspEvents {
	s.DataSource = &v
	return s
}

func (s *DescribeSuspEventsResponseSuspEvents) SetOperateErrorCode(v string) *DescribeSuspEventsResponseSuspEvents {
	s.OperateErrorCode = &v
	return s
}

func (s *DescribeSuspEventsResponseSuspEvents) SetCanCancelFault(v bool) *DescribeSuspEventsResponseSuspEvents {
	s.CanCancelFault = &v
	return s
}

func (s *DescribeSuspEventsResponseSuspEvents) SetCanBeDealOnLine(v string) *DescribeSuspEventsResponseSuspEvents {
	s.CanBeDealOnLine = &v
	return s
}

func (s *DescribeSuspEventsResponseSuspEvents) SetAlarmEventType(v string) *DescribeSuspEventsResponseSuspEvents {
	s.AlarmEventType = &v
	return s
}

func (s *DescribeSuspEventsResponseSuspEvents) SetAlarmEventName(v string) *DescribeSuspEventsResponseSuspEvents {
	s.AlarmEventName = &v
	return s
}

func (s *DescribeSuspEventsResponseSuspEvents) SetAlarmUniqueInfo(v string) *DescribeSuspEventsResponseSuspEvents {
	s.AlarmUniqueInfo = &v
	return s
}

func (s *DescribeSuspEventsResponseSuspEvents) SetAlarmEventNameDisplay(v string) *DescribeSuspEventsResponseSuspEvents {
	s.AlarmEventNameDisplay = &v
	return s
}

func (s *DescribeSuspEventsResponseSuspEvents) SetAlarmEventTypeDisplay(v string) *DescribeSuspEventsResponseSuspEvents {
	s.AlarmEventTypeDisplay = &v
	return s
}

func (s *DescribeSuspEventsResponseSuspEvents) SetDetails(v []*DescribeSuspEventsResponseSuspEventsDetails) *DescribeSuspEventsResponseSuspEvents {
	s.Details = v
	return s
}

type DescribeSuspEventsResponseSuspEventsDetails struct {
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	NameDisplay  *string `json:"NameDisplay,omitempty" xml:"NameDisplay,omitempty" require:"true"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
	ValueDisplay *string `json:"ValueDisplay,omitempty" xml:"ValueDisplay,omitempty" require:"true"`
}

func (s DescribeSuspEventsResponseSuspEventsDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventsResponseSuspEventsDetails) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsResponseSuspEventsDetails) SetName(v string) *DescribeSuspEventsResponseSuspEventsDetails {
	s.Name = &v
	return s
}

func (s *DescribeSuspEventsResponseSuspEventsDetails) SetNameDisplay(v string) *DescribeSuspEventsResponseSuspEventsDetails {
	s.NameDisplay = &v
	return s
}

func (s *DescribeSuspEventsResponseSuspEventsDetails) SetType(v string) *DescribeSuspEventsResponseSuspEventsDetails {
	s.Type = &v
	return s
}

func (s *DescribeSuspEventsResponseSuspEventsDetails) SetValue(v string) *DescribeSuspEventsResponseSuspEventsDetails {
	s.Value = &v
	return s
}

func (s *DescribeSuspEventsResponseSuspEventsDetails) SetValueDisplay(v string) *DescribeSuspEventsResponseSuspEventsDetails {
	s.ValueDisplay = &v
	return s
}

type DescribeModificationPriceRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	InstanceType    *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	RootDiskSizeGib *int    `json:"RootDiskSizeGib,omitempty" xml:"RootDiskSizeGib,omitempty"`
	UserDiskSizeGib *int    `json:"UserDiskSizeGib,omitempty" xml:"UserDiskSizeGib,omitempty"`
}

func (s DescribeModificationPriceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeModificationPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeModificationPriceRequest) SetRegionId(v string) *DescribeModificationPriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeModificationPriceRequest) SetInstanceId(v string) *DescribeModificationPriceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeModificationPriceRequest) SetInstanceType(v string) *DescribeModificationPriceRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeModificationPriceRequest) SetRootDiskSizeGib(v int) *DescribeModificationPriceRequest {
	s.RootDiskSizeGib = &v
	return s
}

func (s *DescribeModificationPriceRequest) SetUserDiskSizeGib(v int) *DescribeModificationPriceRequest {
	s.UserDiskSizeGib = &v
	return s
}

type DescribeModificationPriceResponse struct {
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PriceInfo *DescribeModificationPriceResponsePriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" require:"true" type:"Struct"`
}

func (s DescribeModificationPriceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeModificationPriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeModificationPriceResponse) SetRequestId(v string) *DescribeModificationPriceResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeModificationPriceResponse) SetPriceInfo(v *DescribeModificationPriceResponsePriceInfo) *DescribeModificationPriceResponse {
	s.PriceInfo = v
	return s
}

type DescribeModificationPriceResponsePriceInfo struct {
	Rules []*DescribeModificationPriceResponsePriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" require:"true" type:"Repeated"`
	Price *DescribeModificationPriceResponsePriceInfoPrice   `json:"Price,omitempty" xml:"Price,omitempty" require:"true" type:"Struct"`
}

func (s DescribeModificationPriceResponsePriceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeModificationPriceResponsePriceInfo) GoString() string {
	return s.String()
}

func (s *DescribeModificationPriceResponsePriceInfo) SetRules(v []*DescribeModificationPriceResponsePriceInfoRules) *DescribeModificationPriceResponsePriceInfo {
	s.Rules = v
	return s
}

func (s *DescribeModificationPriceResponsePriceInfo) SetPrice(v *DescribeModificationPriceResponsePriceInfoPrice) *DescribeModificationPriceResponsePriceInfo {
	s.Price = v
	return s
}

type DescribeModificationPriceResponsePriceInfoRules struct {
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty" require:"true"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
}

func (s DescribeModificationPriceResponsePriceInfoRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeModificationPriceResponsePriceInfoRules) GoString() string {
	return s.String()
}

func (s *DescribeModificationPriceResponsePriceInfoRules) SetRuleId(v int64) *DescribeModificationPriceResponsePriceInfoRules {
	s.RuleId = &v
	return s
}

func (s *DescribeModificationPriceResponsePriceInfoRules) SetDescription(v string) *DescribeModificationPriceResponsePriceInfoRules {
	s.Description = &v
	return s
}

type DescribeModificationPriceResponsePriceInfoPrice struct {
	OriginalPrice *float32                                                     `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty" require:"true"`
	DiscountPrice *float32                                                     `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty" require:"true"`
	TradePrice    *float32                                                     `json:"TradePrice,omitempty" xml:"TradePrice,omitempty" require:"true"`
	Currency      *string                                                      `json:"Currency,omitempty" xml:"Currency,omitempty" require:"true"`
	Promotions    []*DescribeModificationPriceResponsePriceInfoPricePromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeModificationPriceResponsePriceInfoPrice) String() string {
	return tea.Prettify(s)
}

func (s DescribeModificationPriceResponsePriceInfoPrice) GoString() string {
	return s.String()
}

func (s *DescribeModificationPriceResponsePriceInfoPrice) SetOriginalPrice(v float32) *DescribeModificationPriceResponsePriceInfoPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeModificationPriceResponsePriceInfoPrice) SetDiscountPrice(v float32) *DescribeModificationPriceResponsePriceInfoPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeModificationPriceResponsePriceInfoPrice) SetTradePrice(v float32) *DescribeModificationPriceResponsePriceInfoPrice {
	s.TradePrice = &v
	return s
}

func (s *DescribeModificationPriceResponsePriceInfoPrice) SetCurrency(v string) *DescribeModificationPriceResponsePriceInfoPrice {
	s.Currency = &v
	return s
}

func (s *DescribeModificationPriceResponsePriceInfoPrice) SetPromotions(v []*DescribeModificationPriceResponsePriceInfoPricePromotions) *DescribeModificationPriceResponsePriceInfoPrice {
	s.Promotions = v
	return s
}

type DescribeModificationPriceResponsePriceInfoPricePromotions struct {
	OptionCode    *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty" require:"true"`
	PromotionId   *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty" require:"true"`
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty" require:"true"`
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty" require:"true"`
	Selected      *bool   `json:"Selected,omitempty" xml:"Selected,omitempty" require:"true"`
}

func (s DescribeModificationPriceResponsePriceInfoPricePromotions) String() string {
	return tea.Prettify(s)
}

func (s DescribeModificationPriceResponsePriceInfoPricePromotions) GoString() string {
	return s.String()
}

func (s *DescribeModificationPriceResponsePriceInfoPricePromotions) SetOptionCode(v string) *DescribeModificationPriceResponsePriceInfoPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *DescribeModificationPriceResponsePriceInfoPricePromotions) SetPromotionId(v string) *DescribeModificationPriceResponsePriceInfoPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *DescribeModificationPriceResponsePriceInfoPricePromotions) SetPromotionName(v string) *DescribeModificationPriceResponsePriceInfoPricePromotions {
	s.PromotionName = &v
	return s
}

func (s *DescribeModificationPriceResponsePriceInfoPricePromotions) SetPromotionDesc(v string) *DescribeModificationPriceResponsePriceInfoPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *DescribeModificationPriceResponsePriceInfoPricePromotions) SetSelected(v bool) *DescribeModificationPriceResponsePriceInfoPricePromotions {
	s.Selected = &v
	return s
}

type DeleteOfficeSitesRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	OfficeSiteId []*string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteOfficeSitesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteOfficeSitesRequest) GoString() string {
	return s.String()
}

func (s *DeleteOfficeSitesRequest) SetRegionId(v string) *DeleteOfficeSitesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteOfficeSitesRequest) SetOfficeSiteId(v []*string) *DeleteOfficeSitesRequest {
	s.OfficeSiteId = v
	return s
}

type DeleteOfficeSitesResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteOfficeSitesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteOfficeSitesResponse) GoString() string {
	return s.String()
}

func (s *DeleteOfficeSitesResponse) SetRequestId(v string) *DeleteOfficeSitesResponse {
	s.RequestId = &v
	return s
}

type DescribeDesktopIdsByVulNamesRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Type         *string   `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	OfficeSiteId *string   `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true"`
	Necessity    *string   `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	VulName      []*string `json:"VulName,omitempty" xml:"VulName,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDesktopIdsByVulNamesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopIdsByVulNamesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopIdsByVulNamesRequest) SetRegionId(v string) *DescribeDesktopIdsByVulNamesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDesktopIdsByVulNamesRequest) SetType(v string) *DescribeDesktopIdsByVulNamesRequest {
	s.Type = &v
	return s
}

func (s *DescribeDesktopIdsByVulNamesRequest) SetOfficeSiteId(v string) *DescribeDesktopIdsByVulNamesRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeDesktopIdsByVulNamesRequest) SetNecessity(v string) *DescribeDesktopIdsByVulNamesRequest {
	s.Necessity = &v
	return s
}

func (s *DescribeDesktopIdsByVulNamesRequest) SetVulName(v []*string) *DescribeDesktopIdsByVulNamesRequest {
	s.VulName = v
	return s
}

type DescribeDesktopIdsByVulNamesResponse struct {
	RequestId    *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DesktopItems []*DescribeDesktopIdsByVulNamesResponseDesktopItems `json:"DesktopItems,omitempty" xml:"DesktopItems,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDesktopIdsByVulNamesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopIdsByVulNamesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopIdsByVulNamesResponse) SetRequestId(v string) *DescribeDesktopIdsByVulNamesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDesktopIdsByVulNamesResponse) SetDesktopItems(v []*DescribeDesktopIdsByVulNamesResponseDesktopItems) *DescribeDesktopIdsByVulNamesResponse {
	s.DesktopItems = v
	return s
}

type DescribeDesktopIdsByVulNamesResponseDesktopItems struct {
	DesktopId   *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty" require:"true"`
}

func (s DescribeDesktopIdsByVulNamesResponseDesktopItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopIdsByVulNamesResponseDesktopItems) GoString() string {
	return s.String()
}

func (s *DescribeDesktopIdsByVulNamesResponseDesktopItems) SetDesktopId(v string) *DescribeDesktopIdsByVulNamesResponseDesktopItems {
	s.DesktopId = &v
	return s
}

func (s *DescribeDesktopIdsByVulNamesResponseDesktopItems) SetDesktopName(v string) *DescribeDesktopIdsByVulNamesResponseDesktopItems {
	s.DesktopName = &v
	return s
}

type GetOfficeSiteSsoStatusRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true"`
}

func (s GetOfficeSiteSsoStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOfficeSiteSsoStatusRequest) GoString() string {
	return s.String()
}

func (s *GetOfficeSiteSsoStatusRequest) SetRegionId(v string) *GetOfficeSiteSsoStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetOfficeSiteSsoStatusRequest) SetOfficeSiteId(v string) *GetOfficeSiteSsoStatusRequest {
	s.OfficeSiteId = &v
	return s
}

type GetOfficeSiteSsoStatusResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SsoStatus *bool   `json:"SsoStatus,omitempty" xml:"SsoStatus,omitempty" require:"true"`
}

func (s GetOfficeSiteSsoStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOfficeSiteSsoStatusResponse) GoString() string {
	return s.String()
}

func (s *GetOfficeSiteSsoStatusResponse) SetRequestId(v string) *GetOfficeSiteSsoStatusResponse {
	s.RequestId = &v
	return s
}

func (s *GetOfficeSiteSsoStatusResponse) SetSsoStatus(v bool) *GetOfficeSiteSsoStatusResponse {
	s.SsoStatus = &v
	return s
}

type DescribeSecurityEventOperationsRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	SecurityEventId *int64  `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty" require:"true"`
}

func (s DescribeSecurityEventOperationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationsRequest) SetRegionId(v string) *DescribeSecurityEventOperationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSecurityEventOperationsRequest) SetSecurityEventId(v int64) *DescribeSecurityEventOperationsRequest {
	s.SecurityEventId = &v
	return s
}

type DescribeSecurityEventOperationsResponse struct {
	RequestId               *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SecurityEventOperations []*DescribeSecurityEventOperationsResponseSecurityEventOperations `json:"SecurityEventOperations,omitempty" xml:"SecurityEventOperations,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeSecurityEventOperationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationsResponse) SetRequestId(v string) *DescribeSecurityEventOperationsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponse) SetSecurityEventOperations(v []*DescribeSecurityEventOperationsResponseSecurityEventOperations) *DescribeSecurityEventOperationsResponse {
	s.SecurityEventOperations = v
	return s
}

type DescribeSecurityEventOperationsResponseSecurityEventOperations struct {
	OperationCode   *string `json:"OperationCode,omitempty" xml:"OperationCode,omitempty" require:"true"`
	OperationParams *string `json:"OperationParams,omitempty" xml:"OperationParams,omitempty" require:"true"`
	UserCanOperate  *bool   `json:"UserCanOperate,omitempty" xml:"UserCanOperate,omitempty" require:"true"`
}

func (s DescribeSecurityEventOperationsResponseSecurityEventOperations) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationsResponseSecurityEventOperations) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationsResponseSecurityEventOperations) SetOperationCode(v string) *DescribeSecurityEventOperationsResponseSecurityEventOperations {
	s.OperationCode = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseSecurityEventOperations) SetOperationParams(v string) *DescribeSecurityEventOperationsResponseSecurityEventOperations {
	s.OperationParams = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseSecurityEventOperations) SetUserCanOperate(v bool) *DescribeSecurityEventOperationsResponseSecurityEventOperations {
	s.UserCanOperate = &v
	return s
}

type CreateNetworkPackageRequest struct {
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Bandwidth          *int    `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty" require:"true"`
	OfficeSiteId       *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true"`
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	Period             *int    `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit         *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	AutoPay            *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoRenew          *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
}

func (s CreateNetworkPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkPackageRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkPackageRequest) SetRegionId(v string) *CreateNetworkPackageRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetBandwidth(v int) *CreateNetworkPackageRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetOfficeSiteId(v string) *CreateNetworkPackageRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetInternetChargeType(v string) *CreateNetworkPackageRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetPeriod(v int) *CreateNetworkPackageRequest {
	s.Period = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetPeriodUnit(v string) *CreateNetworkPackageRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetAutoPay(v bool) *CreateNetworkPackageRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetAutoRenew(v bool) *CreateNetworkPackageRequest {
	s.AutoRenew = &v
	return s
}

type CreateNetworkPackageResponse struct {
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	NetworkPackageId *string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty" require:"true"`
	OrderId          *string `json:"OrderId,omitempty" xml:"OrderId,omitempty" require:"true"`
}

func (s CreateNetworkPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkPackageResponse) GoString() string {
	return s.String()
}

func (s *CreateNetworkPackageResponse) SetRequestId(v string) *CreateNetworkPackageResponse {
	s.RequestId = &v
	return s
}

func (s *CreateNetworkPackageResponse) SetNetworkPackageId(v string) *CreateNetworkPackageResponse {
	s.NetworkPackageId = &v
	return s
}

func (s *CreateNetworkPackageResponse) SetOrderId(v string) *CreateNetworkPackageResponse {
	s.OrderId = &v
	return s
}

type CreateADConnectorOfficeSiteRequest struct {
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	CidrBlock            *string   `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty" require:"true"`
	CenId                *string   `json:"CenId,omitempty" xml:"CenId,omitempty" require:"true"`
	Bandwidth            *int      `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	DomainName           *string   `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	DomainUserName       *string   `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty" require:"true"`
	DomainPassword       *string   `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty" require:"true"`
	DnsAddress           []*string `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" require:"true" type:"Repeated"`
	OfficeSiteName       *string   `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	EnableAdminAccess    *bool     `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	DesktopAccessType    *string   `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	EnableInternetAccess *bool     `json:"EnableInternetAccess,omitempty" xml:"EnableInternetAccess,omitempty"`
	SubDomainDnsAddress  []*string `json:"SubDomainDnsAddress,omitempty" xml:"SubDomainDnsAddress,omitempty" type:"Repeated"`
	SubDomainName        *string   `json:"SubDomainName,omitempty" xml:"SubDomainName,omitempty"`
	MfaEnabled           *bool     `json:"MfaEnabled,omitempty" xml:"MfaEnabled,omitempty"`
}

func (s CreateADConnectorOfficeSiteRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateADConnectorOfficeSiteRequest) GoString() string {
	return s.String()
}

func (s *CreateADConnectorOfficeSiteRequest) SetRegionId(v string) *CreateADConnectorOfficeSiteRequest {
	s.RegionId = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetCidrBlock(v string) *CreateADConnectorOfficeSiteRequest {
	s.CidrBlock = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetCenId(v string) *CreateADConnectorOfficeSiteRequest {
	s.CenId = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetBandwidth(v int) *CreateADConnectorOfficeSiteRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetDomainName(v string) *CreateADConnectorOfficeSiteRequest {
	s.DomainName = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetDomainUserName(v string) *CreateADConnectorOfficeSiteRequest {
	s.DomainUserName = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetDomainPassword(v string) *CreateADConnectorOfficeSiteRequest {
	s.DomainPassword = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetDnsAddress(v []*string) *CreateADConnectorOfficeSiteRequest {
	s.DnsAddress = v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetOfficeSiteName(v string) *CreateADConnectorOfficeSiteRequest {
	s.OfficeSiteName = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetEnableAdminAccess(v bool) *CreateADConnectorOfficeSiteRequest {
	s.EnableAdminAccess = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetDesktopAccessType(v string) *CreateADConnectorOfficeSiteRequest {
	s.DesktopAccessType = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetEnableInternetAccess(v bool) *CreateADConnectorOfficeSiteRequest {
	s.EnableInternetAccess = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetSubDomainDnsAddress(v []*string) *CreateADConnectorOfficeSiteRequest {
	s.SubDomainDnsAddress = v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetSubDomainName(v string) *CreateADConnectorOfficeSiteRequest {
	s.SubDomainName = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetMfaEnabled(v bool) *CreateADConnectorOfficeSiteRequest {
	s.MfaEnabled = &v
	return s
}

type CreateADConnectorOfficeSiteResponse struct {
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s CreateADConnectorOfficeSiteResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateADConnectorOfficeSiteResponse) GoString() string {
	return s.String()
}

func (s *CreateADConnectorOfficeSiteResponse) SetOfficeSiteId(v string) *CreateADConnectorOfficeSiteResponse {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateADConnectorOfficeSiteResponse) SetRequestId(v string) *CreateADConnectorOfficeSiteResponse {
	s.RequestId = &v
	return s
}

type DeleteNetworkPackagesRequest struct {
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	NetworkPackageId []*string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteNetworkPackagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkPackagesRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkPackagesRequest) SetRegionId(v string) *DeleteNetworkPackagesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteNetworkPackagesRequest) SetNetworkPackageId(v []*string) *DeleteNetworkPackagesRequest {
	s.NetworkPackageId = v
	return s
}

type DeleteNetworkPackagesResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteNetworkPackagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkPackagesResponse) GoString() string {
	return s.String()
}

func (s *DeleteNetworkPackagesResponse) SetRequestId(v string) *DeleteNetworkPackagesResponse {
	s.RequestId = &v
	return s
}

type SetOfficeSiteSsoStatusRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true"`
	EnableSso    *bool   `json:"EnableSso,omitempty" xml:"EnableSso,omitempty" require:"true"`
}

func (s SetOfficeSiteSsoStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetOfficeSiteSsoStatusRequest) GoString() string {
	return s.String()
}

func (s *SetOfficeSiteSsoStatusRequest) SetRegionId(v string) *SetOfficeSiteSsoStatusRequest {
	s.RegionId = &v
	return s
}

func (s *SetOfficeSiteSsoStatusRequest) SetOfficeSiteId(v string) *SetOfficeSiteSsoStatusRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *SetOfficeSiteSsoStatusRequest) SetEnableSso(v bool) *SetOfficeSiteSsoStatusRequest {
	s.EnableSso = &v
	return s
}

type SetOfficeSiteSsoStatusResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s SetOfficeSiteSsoStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetOfficeSiteSsoStatusResponse) GoString() string {
	return s.String()
}

func (s *SetOfficeSiteSsoStatusResponse) SetRequestId(v string) *SetOfficeSiteSsoStatusResponse {
	s.RequestId = &v
	return s
}

type HandleSecurityEventsRequest struct {
	RegionId        *string                                     `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	SecurityEvent   []*HandleSecurityEventsRequestSecurityEvent `json:"SecurityEvent,omitempty" xml:"SecurityEvent,omitempty" type:"Repeated"`
	OperationCode   *string                                     `json:"OperationCode,omitempty" xml:"OperationCode,omitempty" require:"true"`
	OperationParams *string                                     `json:"OperationParams,omitempty" xml:"OperationParams,omitempty"`
}

func (s HandleSecurityEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s HandleSecurityEventsRequest) GoString() string {
	return s.String()
}

func (s *HandleSecurityEventsRequest) SetRegionId(v string) *HandleSecurityEventsRequest {
	s.RegionId = &v
	return s
}

func (s *HandleSecurityEventsRequest) SetSecurityEvent(v []*HandleSecurityEventsRequestSecurityEvent) *HandleSecurityEventsRequest {
	s.SecurityEvent = v
	return s
}

func (s *HandleSecurityEventsRequest) SetOperationCode(v string) *HandleSecurityEventsRequest {
	s.OperationCode = &v
	return s
}

func (s *HandleSecurityEventsRequest) SetOperationParams(v string) *HandleSecurityEventsRequest {
	s.OperationParams = &v
	return s
}

type HandleSecurityEventsRequestSecurityEvent struct {
	SecurityEventId *string `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty"`
	DesktopId       *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
}

func (s HandleSecurityEventsRequestSecurityEvent) String() string {
	return tea.Prettify(s)
}

func (s HandleSecurityEventsRequestSecurityEvent) GoString() string {
	return s.String()
}

func (s *HandleSecurityEventsRequestSecurityEvent) SetSecurityEventId(v string) *HandleSecurityEventsRequestSecurityEvent {
	s.SecurityEventId = &v
	return s
}

func (s *HandleSecurityEventsRequestSecurityEvent) SetDesktopId(v string) *HandleSecurityEventsRequestSecurityEvent {
	s.DesktopId = &v
	return s
}

type HandleSecurityEventsResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TaskId    *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
}

func (s HandleSecurityEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s HandleSecurityEventsResponse) GoString() string {
	return s.String()
}

func (s *HandleSecurityEventsResponse) SetRequestId(v string) *HandleSecurityEventsResponse {
	s.RequestId = &v
	return s
}

func (s *HandleSecurityEventsResponse) SetTaskId(v int64) *HandleSecurityEventsResponse {
	s.TaskId = &v
	return s
}

type ModifyNetworkPackageRequest struct {
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	NetworkPackageId *string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty" require:"true"`
	Bandwidth        *int    `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
}

func (s ModifyNetworkPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNetworkPackageRequest) GoString() string {
	return s.String()
}

func (s *ModifyNetworkPackageRequest) SetRegionId(v string) *ModifyNetworkPackageRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyNetworkPackageRequest) SetNetworkPackageId(v string) *ModifyNetworkPackageRequest {
	s.NetworkPackageId = &v
	return s
}

func (s *ModifyNetworkPackageRequest) SetBandwidth(v int) *ModifyNetworkPackageRequest {
	s.Bandwidth = &v
	return s
}

type ModifyNetworkPackageResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyNetworkPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyNetworkPackageResponse) GoString() string {
	return s.String()
}

func (s *ModifyNetworkPackageResponse) SetRequestId(v string) *ModifyNetworkPackageResponse {
	s.RequestId = &v
	return s
}

type DescribeNetworkPackagesRequest struct {
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	NetworkPackageId []*string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty" type:"Repeated"`
	MaxResults       *int      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken        *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeNetworkPackagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkPackagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkPackagesRequest) SetRegionId(v string) *DescribeNetworkPackagesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeNetworkPackagesRequest) SetNetworkPackageId(v []*string) *DescribeNetworkPackagesRequest {
	s.NetworkPackageId = v
	return s
}

func (s *DescribeNetworkPackagesRequest) SetMaxResults(v int) *DescribeNetworkPackagesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeNetworkPackagesRequest) SetNextToken(v string) *DescribeNetworkPackagesRequest {
	s.NextToken = &v
	return s
}

type DescribeNetworkPackagesResponse struct {
	NextToken       *string                                           `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	RequestId       *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	NetworkPackages []*DescribeNetworkPackagesResponseNetworkPackages `json:"NetworkPackages,omitempty" xml:"NetworkPackages,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeNetworkPackagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkPackagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworkPackagesResponse) SetNextToken(v string) *DescribeNetworkPackagesResponse {
	s.NextToken = &v
	return s
}

func (s *DescribeNetworkPackagesResponse) SetRequestId(v string) *DescribeNetworkPackagesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeNetworkPackagesResponse) SetNetworkPackages(v []*DescribeNetworkPackagesResponseNetworkPackages) *DescribeNetworkPackagesResponse {
	s.NetworkPackages = v
	return s
}

type DescribeNetworkPackagesResponseNetworkPackages struct {
	NetworkPackageId     *string   `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty" require:"true"`
	OfficeSiteId         *string   `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true"`
	OfficeSiteName       *string   `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty" require:"true"`
	Bandwidth            *int      `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty" require:"true"`
	NetworkPackageStatus *string   `json:"NetworkPackageStatus,omitempty" xml:"NetworkPackageStatus,omitempty" require:"true"`
	CreateTime           *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	ExpiredTime          *string   `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty" require:"true"`
	InternetChargeType   *string   `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty" require:"true"`
	EipAddresses         []*string `json:"EipAddresses,omitempty" xml:"EipAddresses,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeNetworkPackagesResponseNetworkPackages) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkPackagesResponseNetworkPackages) GoString() string {
	return s.String()
}

func (s *DescribeNetworkPackagesResponseNetworkPackages) SetNetworkPackageId(v string) *DescribeNetworkPackagesResponseNetworkPackages {
	s.NetworkPackageId = &v
	return s
}

func (s *DescribeNetworkPackagesResponseNetworkPackages) SetOfficeSiteId(v string) *DescribeNetworkPackagesResponseNetworkPackages {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeNetworkPackagesResponseNetworkPackages) SetOfficeSiteName(v string) *DescribeNetworkPackagesResponseNetworkPackages {
	s.OfficeSiteName = &v
	return s
}

func (s *DescribeNetworkPackagesResponseNetworkPackages) SetBandwidth(v int) *DescribeNetworkPackagesResponseNetworkPackages {
	s.Bandwidth = &v
	return s
}

func (s *DescribeNetworkPackagesResponseNetworkPackages) SetNetworkPackageStatus(v string) *DescribeNetworkPackagesResponseNetworkPackages {
	s.NetworkPackageStatus = &v
	return s
}

func (s *DescribeNetworkPackagesResponseNetworkPackages) SetCreateTime(v string) *DescribeNetworkPackagesResponseNetworkPackages {
	s.CreateTime = &v
	return s
}

func (s *DescribeNetworkPackagesResponseNetworkPackages) SetExpiredTime(v string) *DescribeNetworkPackagesResponseNetworkPackages {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeNetworkPackagesResponseNetworkPackages) SetInternetChargeType(v string) *DescribeNetworkPackagesResponseNetworkPackages {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeNetworkPackagesResponseNetworkPackages) SetEipAddresses(v []*string) *DescribeNetworkPackagesResponseNetworkPackages {
	s.EipAddresses = v
	return s
}

type DescribeGroupedVulRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true"`
	Necessity    *string `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	Dealed       *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	CurrentPage  *int    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize     *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeGroupedVulRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupedVulRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupedVulRequest) SetRegionId(v string) *DescribeGroupedVulRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetLang(v string) *DescribeGroupedVulRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetType(v string) *DescribeGroupedVulRequest {
	s.Type = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetOfficeSiteId(v string) *DescribeGroupedVulRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetNecessity(v string) *DescribeGroupedVulRequest {
	s.Necessity = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetDealed(v string) *DescribeGroupedVulRequest {
	s.Dealed = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetCurrentPage(v int) *DescribeGroupedVulRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetPageSize(v int) *DescribeGroupedVulRequest {
	s.PageSize = &v
	return s
}

type DescribeGroupedVulResponse struct {
	RequestId       *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageSize        *int                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	CurrentPage     *int                                         `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
	TotalCount      *int                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	GroupedVulItems []*DescribeGroupedVulResponseGroupedVulItems `json:"GroupedVulItems,omitempty" xml:"GroupedVulItems,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeGroupedVulResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupedVulResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupedVulResponse) SetRequestId(v string) *DescribeGroupedVulResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupedVulResponse) SetPageSize(v int) *DescribeGroupedVulResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupedVulResponse) SetCurrentPage(v int) *DescribeGroupedVulResponse {
	s.CurrentPage = &v
	return s
}

func (s *DescribeGroupedVulResponse) SetTotalCount(v int) *DescribeGroupedVulResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeGroupedVulResponse) SetGroupedVulItems(v []*DescribeGroupedVulResponseGroupedVulItems) *DescribeGroupedVulResponse {
	s.GroupedVulItems = v
	return s
}

type DescribeGroupedVulResponseGroupedVulItems struct {
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	AliasName    *string `json:"AliasName,omitempty" xml:"AliasName,omitempty" require:"true"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	GmtLast      *string `json:"GmtLast,omitempty" xml:"GmtLast,omitempty" require:"true"`
	AsapCount    *int    `json:"AsapCount,omitempty" xml:"AsapCount,omitempty" require:"true"`
	LaterCount   *int    `json:"LaterCount,omitempty" xml:"LaterCount,omitempty" require:"true"`
	NntfCount    *int    `json:"NntfCount,omitempty" xml:"NntfCount,omitempty" require:"true"`
	HandledCount *int    `json:"HandledCount,omitempty" xml:"HandledCount,omitempty" require:"true"`
	Tags         *string `json:"Tags,omitempty" xml:"Tags,omitempty" require:"true"`
}

func (s DescribeGroupedVulResponseGroupedVulItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupedVulResponseGroupedVulItems) GoString() string {
	return s.String()
}

func (s *DescribeGroupedVulResponseGroupedVulItems) SetName(v string) *DescribeGroupedVulResponseGroupedVulItems {
	s.Name = &v
	return s
}

func (s *DescribeGroupedVulResponseGroupedVulItems) SetAliasName(v string) *DescribeGroupedVulResponseGroupedVulItems {
	s.AliasName = &v
	return s
}

func (s *DescribeGroupedVulResponseGroupedVulItems) SetType(v string) *DescribeGroupedVulResponseGroupedVulItems {
	s.Type = &v
	return s
}

func (s *DescribeGroupedVulResponseGroupedVulItems) SetGmtLast(v string) *DescribeGroupedVulResponseGroupedVulItems {
	s.GmtLast = &v
	return s
}

func (s *DescribeGroupedVulResponseGroupedVulItems) SetAsapCount(v int) *DescribeGroupedVulResponseGroupedVulItems {
	s.AsapCount = &v
	return s
}

func (s *DescribeGroupedVulResponseGroupedVulItems) SetLaterCount(v int) *DescribeGroupedVulResponseGroupedVulItems {
	s.LaterCount = &v
	return s
}

func (s *DescribeGroupedVulResponseGroupedVulItems) SetNntfCount(v int) *DescribeGroupedVulResponseGroupedVulItems {
	s.NntfCount = &v
	return s
}

func (s *DescribeGroupedVulResponseGroupedVulItems) SetHandledCount(v int) *DescribeGroupedVulResponseGroupedVulItems {
	s.HandledCount = &v
	return s
}

func (s *DescribeGroupedVulResponseGroupedVulItems) SetTags(v string) *DescribeGroupedVulResponseGroupedVulItems {
	s.Tags = &v
	return s
}

type RollbackSuspEventQuaraFileRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	QuaraFieldId *int    `json:"QuaraFieldId,omitempty" xml:"QuaraFieldId,omitempty" require:"true"`
	DesktopId    *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
}

func (s RollbackSuspEventQuaraFileRequest) String() string {
	return tea.Prettify(s)
}

func (s RollbackSuspEventQuaraFileRequest) GoString() string {
	return s.String()
}

func (s *RollbackSuspEventQuaraFileRequest) SetRegionId(v string) *RollbackSuspEventQuaraFileRequest {
	s.RegionId = &v
	return s
}

func (s *RollbackSuspEventQuaraFileRequest) SetQuaraFieldId(v int) *RollbackSuspEventQuaraFileRequest {
	s.QuaraFieldId = &v
	return s
}

func (s *RollbackSuspEventQuaraFileRequest) SetDesktopId(v string) *RollbackSuspEventQuaraFileRequest {
	s.DesktopId = &v
	return s
}

type RollbackSuspEventQuaraFileResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s RollbackSuspEventQuaraFileResponse) String() string {
	return tea.Prettify(s)
}

func (s RollbackSuspEventQuaraFileResponse) GoString() string {
	return s.String()
}

func (s *RollbackSuspEventQuaraFileResponse) SetRequestId(v string) *RollbackSuspEventQuaraFileResponse {
	s.RequestId = &v
	return s
}

type DescribePriceRequest struct {
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ResourceType       *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	InstanceType       *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	RootDiskSizeGib    *int    `json:"RootDiskSizeGib,omitempty" xml:"RootDiskSizeGib,omitempty"`
	UserDiskSizeGib    *int    `json:"UserDiskSizeGib,omitempty" xml:"UserDiskSizeGib,omitempty"`
	OsType             *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	PeriodUnit         *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	Period             *int    `json:"Period,omitempty" xml:"Period,omitempty"`
	Amount             *int    `json:"Amount,omitempty" xml:"Amount,omitempty"`
	PromotionId        *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	Bandwidth          *int    `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
}

func (s DescribePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceRequest) GoString() string {
	return s.String()
}

func (s *DescribePriceRequest) SetRegionId(v string) *DescribePriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePriceRequest) SetResourceType(v string) *DescribePriceRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribePriceRequest) SetInstanceType(v string) *DescribePriceRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribePriceRequest) SetRootDiskSizeGib(v int) *DescribePriceRequest {
	s.RootDiskSizeGib = &v
	return s
}

func (s *DescribePriceRequest) SetUserDiskSizeGib(v int) *DescribePriceRequest {
	s.UserDiskSizeGib = &v
	return s
}

func (s *DescribePriceRequest) SetOsType(v string) *DescribePriceRequest {
	s.OsType = &v
	return s
}

func (s *DescribePriceRequest) SetPeriodUnit(v string) *DescribePriceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *DescribePriceRequest) SetPeriod(v int) *DescribePriceRequest {
	s.Period = &v
	return s
}

func (s *DescribePriceRequest) SetAmount(v int) *DescribePriceRequest {
	s.Amount = &v
	return s
}

func (s *DescribePriceRequest) SetPromotionId(v string) *DescribePriceRequest {
	s.PromotionId = &v
	return s
}

func (s *DescribePriceRequest) SetInternetChargeType(v string) *DescribePriceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *DescribePriceRequest) SetBandwidth(v int) *DescribePriceRequest {
	s.Bandwidth = &v
	return s
}

type DescribePriceResponse struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PriceInfo *DescribePriceResponsePriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" require:"true" type:"Struct"`
}

func (s DescribePriceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponse) GoString() string {
	return s.String()
}

func (s *DescribePriceResponse) SetRequestId(v string) *DescribePriceResponse {
	s.RequestId = &v
	return s
}

func (s *DescribePriceResponse) SetPriceInfo(v *DescribePriceResponsePriceInfo) *DescribePriceResponse {
	s.PriceInfo = v
	return s
}

type DescribePriceResponsePriceInfo struct {
	Rules []*DescribePriceResponsePriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" require:"true" type:"Repeated"`
	Price *DescribePriceResponsePriceInfoPrice   `json:"Price,omitempty" xml:"Price,omitempty" require:"true" type:"Struct"`
}

func (s DescribePriceResponsePriceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponsePriceInfo) GoString() string {
	return s.String()
}

func (s *DescribePriceResponsePriceInfo) SetRules(v []*DescribePriceResponsePriceInfoRules) *DescribePriceResponsePriceInfo {
	s.Rules = v
	return s
}

func (s *DescribePriceResponsePriceInfo) SetPrice(v *DescribePriceResponsePriceInfoPrice) *DescribePriceResponsePriceInfo {
	s.Price = v
	return s
}

type DescribePriceResponsePriceInfoRules struct {
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty" require:"true"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
}

func (s DescribePriceResponsePriceInfoRules) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponsePriceInfoRules) GoString() string {
	return s.String()
}

func (s *DescribePriceResponsePriceInfoRules) SetRuleId(v int64) *DescribePriceResponsePriceInfoRules {
	s.RuleId = &v
	return s
}

func (s *DescribePriceResponsePriceInfoRules) SetDescription(v string) *DescribePriceResponsePriceInfoRules {
	s.Description = &v
	return s
}

type DescribePriceResponsePriceInfoPrice struct {
	OriginalPrice *float32                                         `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty" require:"true"`
	DiscountPrice *float32                                         `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty" require:"true"`
	TradePrice    *float32                                         `json:"TradePrice,omitempty" xml:"TradePrice,omitempty" require:"true"`
	Currency      *string                                          `json:"Currency,omitempty" xml:"Currency,omitempty" require:"true"`
	Promotions    []*DescribePriceResponsePriceInfoPricePromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" require:"true" type:"Repeated"`
}

func (s DescribePriceResponsePriceInfoPrice) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponsePriceInfoPrice) GoString() string {
	return s.String()
}

func (s *DescribePriceResponsePriceInfoPrice) SetOriginalPrice(v float32) *DescribePriceResponsePriceInfoPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribePriceResponsePriceInfoPrice) SetDiscountPrice(v float32) *DescribePriceResponsePriceInfoPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribePriceResponsePriceInfoPrice) SetTradePrice(v float32) *DescribePriceResponsePriceInfoPrice {
	s.TradePrice = &v
	return s
}

func (s *DescribePriceResponsePriceInfoPrice) SetCurrency(v string) *DescribePriceResponsePriceInfoPrice {
	s.Currency = &v
	return s
}

func (s *DescribePriceResponsePriceInfoPrice) SetPromotions(v []*DescribePriceResponsePriceInfoPricePromotions) *DescribePriceResponsePriceInfoPrice {
	s.Promotions = v
	return s
}

type DescribePriceResponsePriceInfoPricePromotions struct {
	OptionCode    *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty" require:"true"`
	PromotionId   *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty" require:"true"`
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty" require:"true"`
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty" require:"true"`
	Selected      *bool   `json:"Selected,omitempty" xml:"Selected,omitempty" require:"true"`
}

func (s DescribePriceResponsePriceInfoPricePromotions) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponsePriceInfoPricePromotions) GoString() string {
	return s.String()
}

func (s *DescribePriceResponsePriceInfoPricePromotions) SetOptionCode(v string) *DescribePriceResponsePriceInfoPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *DescribePriceResponsePriceInfoPricePromotions) SetPromotionId(v string) *DescribePriceResponsePriceInfoPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *DescribePriceResponsePriceInfoPricePromotions) SetPromotionName(v string) *DescribePriceResponsePriceInfoPricePromotions {
	s.PromotionName = &v
	return s
}

func (s *DescribePriceResponsePriceInfoPricePromotions) SetPromotionDesc(v string) *DescribePriceResponsePriceInfoPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *DescribePriceResponsePriceInfoPricePromotions) SetSelected(v bool) *DescribePriceResponsePriceInfoPricePromotions {
	s.Selected = &v
	return s
}

type ModifyDesktopSpecRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopId       *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	DesktopType     *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	RootDiskSizeGib *int    `json:"RootDiskSizeGib,omitempty" xml:"RootDiskSizeGib,omitempty"`
	UserDiskSizeGib *int    `json:"UserDiskSizeGib,omitempty" xml:"UserDiskSizeGib,omitempty"`
	AutoPay         *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
}

func (s ModifyDesktopSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyDesktopSpecRequest) SetRegionId(v string) *ModifyDesktopSpecRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDesktopSpecRequest) SetDesktopId(v string) *ModifyDesktopSpecRequest {
	s.DesktopId = &v
	return s
}

func (s *ModifyDesktopSpecRequest) SetDesktopType(v string) *ModifyDesktopSpecRequest {
	s.DesktopType = &v
	return s
}

func (s *ModifyDesktopSpecRequest) SetRootDiskSizeGib(v int) *ModifyDesktopSpecRequest {
	s.RootDiskSizeGib = &v
	return s
}

func (s *ModifyDesktopSpecRequest) SetUserDiskSizeGib(v int) *ModifyDesktopSpecRequest {
	s.UserDiskSizeGib = &v
	return s
}

func (s *ModifyDesktopSpecRequest) SetAutoPay(v bool) *ModifyDesktopSpecRequest {
	s.AutoPay = &v
	return s
}

type ModifyDesktopSpecResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty" require:"true"`
}

func (s ModifyDesktopSpecResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyDesktopSpecResponse) SetRequestId(v string) *ModifyDesktopSpecResponse {
	s.RequestId = &v
	return s
}

func (s *ModifyDesktopSpecResponse) SetOrderId(v string) *ModifyDesktopSpecResponse {
	s.OrderId = &v
	return s
}

type ListOfficeSiteOverviewRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ForceRefresh *bool     `json:"ForceRefresh,omitempty" xml:"ForceRefresh,omitempty"`
	OfficeSiteId []*string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" type:"Repeated"`
	MaxResults   *int      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListOfficeSiteOverviewRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeSiteOverviewRequest) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteOverviewRequest) SetRegionId(v string) *ListOfficeSiteOverviewRequest {
	s.RegionId = &v
	return s
}

func (s *ListOfficeSiteOverviewRequest) SetForceRefresh(v bool) *ListOfficeSiteOverviewRequest {
	s.ForceRefresh = &v
	return s
}

func (s *ListOfficeSiteOverviewRequest) SetOfficeSiteId(v []*string) *ListOfficeSiteOverviewRequest {
	s.OfficeSiteId = v
	return s
}

func (s *ListOfficeSiteOverviewRequest) SetMaxResults(v int) *ListOfficeSiteOverviewRequest {
	s.MaxResults = &v
	return s
}

func (s *ListOfficeSiteOverviewRequest) SetNextToken(v string) *ListOfficeSiteOverviewRequest {
	s.NextToken = &v
	return s
}

type ListOfficeSiteOverviewResponse struct {
	RequestId                 *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	NextToken                 *string                                                    `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	OfficeSiteOverviewResults []*ListOfficeSiteOverviewResponseOfficeSiteOverviewResults `json:"OfficeSiteOverviewResults,omitempty" xml:"OfficeSiteOverviewResults,omitempty" require:"true" type:"Repeated"`
}

func (s ListOfficeSiteOverviewResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeSiteOverviewResponse) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteOverviewResponse) SetRequestId(v string) *ListOfficeSiteOverviewResponse {
	s.RequestId = &v
	return s
}

func (s *ListOfficeSiteOverviewResponse) SetNextToken(v string) *ListOfficeSiteOverviewResponse {
	s.NextToken = &v
	return s
}

func (s *ListOfficeSiteOverviewResponse) SetOfficeSiteOverviewResults(v []*ListOfficeSiteOverviewResponseOfficeSiteOverviewResults) *ListOfficeSiteOverviewResponse {
	s.OfficeSiteOverviewResults = v
	return s
}

type ListOfficeSiteOverviewResponseOfficeSiteOverviewResults struct {
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	OfficeSiteId        *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true"`
	OfficeSiteName      *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty" require:"true"`
	OfficeSiteStatus    *string `json:"OfficeSiteStatus,omitempty" xml:"OfficeSiteStatus,omitempty" require:"true"`
	TotalEdsCount       *int    `json:"TotalEdsCount,omitempty" xml:"TotalEdsCount,omitempty" require:"true"`
	RunningEdsCount     *int    `json:"RunningEdsCount,omitempty" xml:"RunningEdsCount,omitempty" require:"true"`
	WillExpiredEdsCount *int    `json:"WillExpiredEdsCount,omitempty" xml:"WillExpiredEdsCount,omitempty" require:"true"`
	HasExpiredEdsCount  *int    `json:"HasExpiredEdsCount,omitempty" xml:"HasExpiredEdsCount,omitempty" require:"true"`
}

func (s ListOfficeSiteOverviewResponseOfficeSiteOverviewResults) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeSiteOverviewResponseOfficeSiteOverviewResults) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteOverviewResponseOfficeSiteOverviewResults) SetRegionId(v string) *ListOfficeSiteOverviewResponseOfficeSiteOverviewResults {
	s.RegionId = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseOfficeSiteOverviewResults) SetOfficeSiteId(v string) *ListOfficeSiteOverviewResponseOfficeSiteOverviewResults {
	s.OfficeSiteId = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseOfficeSiteOverviewResults) SetOfficeSiteName(v string) *ListOfficeSiteOverviewResponseOfficeSiteOverviewResults {
	s.OfficeSiteName = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseOfficeSiteOverviewResults) SetOfficeSiteStatus(v string) *ListOfficeSiteOverviewResponseOfficeSiteOverviewResults {
	s.OfficeSiteStatus = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseOfficeSiteOverviewResults) SetTotalEdsCount(v int) *ListOfficeSiteOverviewResponseOfficeSiteOverviewResults {
	s.TotalEdsCount = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseOfficeSiteOverviewResults) SetRunningEdsCount(v int) *ListOfficeSiteOverviewResponseOfficeSiteOverviewResults {
	s.RunningEdsCount = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseOfficeSiteOverviewResults) SetWillExpiredEdsCount(v int) *ListOfficeSiteOverviewResponseOfficeSiteOverviewResults {
	s.WillExpiredEdsCount = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseOfficeSiteOverviewResults) SetHasExpiredEdsCount(v int) *ListOfficeSiteOverviewResponseOfficeSiteOverviewResults {
	s.HasExpiredEdsCount = &v
	return s
}

type GetDirectorySsoStatusRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" require:"true"`
}

func (s GetDirectorySsoStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDirectorySsoStatusRequest) GoString() string {
	return s.String()
}

func (s *GetDirectorySsoStatusRequest) SetRegionId(v string) *GetDirectorySsoStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetDirectorySsoStatusRequest) SetDirectoryId(v string) *GetDirectorySsoStatusRequest {
	s.DirectoryId = &v
	return s
}

type GetDirectorySsoStatusResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SsoStatus *bool   `json:"SsoStatus,omitempty" xml:"SsoStatus,omitempty" require:"true"`
}

func (s GetDirectorySsoStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDirectorySsoStatusResponse) GoString() string {
	return s.String()
}

func (s *GetDirectorySsoStatusResponse) SetRequestId(v string) *GetDirectorySsoStatusResponse {
	s.RequestId = &v
	return s
}

func (s *GetDirectorySsoStatusResponse) SetSsoStatus(v bool) *GetDirectorySsoStatusResponse {
	s.SsoStatus = &v
	return s
}

type SetDirectorySsoStatusRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" require:"true"`
	EnableSso   *bool   `json:"EnableSso,omitempty" xml:"EnableSso,omitempty" require:"true"`
}

func (s SetDirectorySsoStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDirectorySsoStatusRequest) GoString() string {
	return s.String()
}

func (s *SetDirectorySsoStatusRequest) SetRegionId(v string) *SetDirectorySsoStatusRequest {
	s.RegionId = &v
	return s
}

func (s *SetDirectorySsoStatusRequest) SetDirectoryId(v string) *SetDirectorySsoStatusRequest {
	s.DirectoryId = &v
	return s
}

func (s *SetDirectorySsoStatusRequest) SetEnableSso(v bool) *SetDirectorySsoStatusRequest {
	s.EnableSso = &v
	return s
}

type SetDirectorySsoStatusResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s SetDirectorySsoStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDirectorySsoStatusResponse) GoString() string {
	return s.String()
}

func (s *SetDirectorySsoStatusResponse) SetRequestId(v string) *SetDirectorySsoStatusResponse {
	s.RequestId = &v
	return s
}

type GetSpMetadataRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DirectoryId  *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
}

func (s GetSpMetadataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSpMetadataRequest) GoString() string {
	return s.String()
}

func (s *GetSpMetadataRequest) SetRegionId(v string) *GetSpMetadataRequest {
	s.RegionId = &v
	return s
}

func (s *GetSpMetadataRequest) SetDirectoryId(v string) *GetSpMetadataRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetSpMetadataRequest) SetOfficeSiteId(v string) *GetSpMetadataRequest {
	s.OfficeSiteId = &v
	return s
}

type GetSpMetadataResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SpMetadata *string `json:"SpMetadata,omitempty" xml:"SpMetadata,omitempty" require:"true"`
}

func (s GetSpMetadataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSpMetadataResponse) GoString() string {
	return s.String()
}

func (s *GetSpMetadataResponse) SetRequestId(v string) *GetSpMetadataResponse {
	s.RequestId = &v
	return s
}

func (s *GetSpMetadataResponse) SetSpMetadata(v string) *GetSpMetadataResponse {
	s.SpMetadata = &v
	return s
}

type SetIdpMetadataRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DirectoryId  *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	IdpMetadata  *string `json:"IdpMetadata,omitempty" xml:"IdpMetadata,omitempty" require:"true"`
}

func (s SetIdpMetadataRequest) String() string {
	return tea.Prettify(s)
}

func (s SetIdpMetadataRequest) GoString() string {
	return s.String()
}

func (s *SetIdpMetadataRequest) SetRegionId(v string) *SetIdpMetadataRequest {
	s.RegionId = &v
	return s
}

func (s *SetIdpMetadataRequest) SetDirectoryId(v string) *SetIdpMetadataRequest {
	s.DirectoryId = &v
	return s
}

func (s *SetIdpMetadataRequest) SetOfficeSiteId(v string) *SetIdpMetadataRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *SetIdpMetadataRequest) SetIdpMetadata(v string) *SetIdpMetadataRequest {
	s.IdpMetadata = &v
	return s
}

type SetIdpMetadataResponse struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	IdpEntityId *string `json:"IdpEntityId,omitempty" xml:"IdpEntityId,omitempty" require:"true"`
}

func (s SetIdpMetadataResponse) String() string {
	return tea.Prettify(s)
}

func (s SetIdpMetadataResponse) GoString() string {
	return s.String()
}

func (s *SetIdpMetadataResponse) SetRequestId(v string) *SetIdpMetadataResponse {
	s.RequestId = &v
	return s
}

func (s *SetIdpMetadataResponse) SetIdpEntityId(v string) *SetIdpMetadataResponse {
	s.IdpEntityId = &v
	return s
}

type RebuildDesktopsRequest struct {
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true" type:"Repeated"`
}

func (s RebuildDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s RebuildDesktopsRequest) GoString() string {
	return s.String()
}

func (s *RebuildDesktopsRequest) SetRegionId(v string) *RebuildDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *RebuildDesktopsRequest) SetDesktopId(v []*string) *RebuildDesktopsRequest {
	s.DesktopId = v
	return s
}

type RebuildDesktopsResponse struct {
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RebuildResults []*RebuildDesktopsResponseRebuildResults `json:"RebuildResults,omitempty" xml:"RebuildResults,omitempty" require:"true" type:"Repeated"`
}

func (s RebuildDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s RebuildDesktopsResponse) GoString() string {
	return s.String()
}

func (s *RebuildDesktopsResponse) SetRequestId(v string) *RebuildDesktopsResponse {
	s.RequestId = &v
	return s
}

func (s *RebuildDesktopsResponse) SetRebuildResults(v []*RebuildDesktopsResponseRebuildResults) *RebuildDesktopsResponse {
	s.RebuildResults = v
	return s
}

type RebuildDesktopsResponseRebuildResults struct {
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
}

func (s RebuildDesktopsResponseRebuildResults) String() string {
	return tea.Prettify(s)
}

func (s RebuildDesktopsResponseRebuildResults) GoString() string {
	return s.String()
}

func (s *RebuildDesktopsResponseRebuildResults) SetDesktopId(v string) *RebuildDesktopsResponseRebuildResults {
	s.DesktopId = &v
	return s
}

func (s *RebuildDesktopsResponseRebuildResults) SetCode(v string) *RebuildDesktopsResponseRebuildResults {
	s.Code = &v
	return s
}

func (s *RebuildDesktopsResponseRebuildResults) SetMessage(v string) *RebuildDesktopsResponseRebuildResults {
	s.Message = &v
	return s
}

type ModifyBundleRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	BundleId    *string `json:"BundleId,omitempty" xml:"BundleId,omitempty" require:"true"`
	ImageId     *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	BundleName  *string `json:"BundleName,omitempty" xml:"BundleName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyBundleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBundleRequest) GoString() string {
	return s.String()
}

func (s *ModifyBundleRequest) SetRegionId(v string) *ModifyBundleRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyBundleRequest) SetBundleId(v string) *ModifyBundleRequest {
	s.BundleId = &v
	return s
}

func (s *ModifyBundleRequest) SetImageId(v string) *ModifyBundleRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyBundleRequest) SetBundleName(v string) *ModifyBundleRequest {
	s.BundleName = &v
	return s
}

func (s *ModifyBundleRequest) SetDescription(v string) *ModifyBundleRequest {
	s.Description = &v
	return s
}

type ModifyBundleResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyBundleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyBundleResponse) GoString() string {
	return s.String()
}

func (s *ModifyBundleResponse) SetRequestId(v string) *ModifyBundleResponse {
	s.RequestId = &v
	return s
}

type UnlockVirtualMFADeviceRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty" require:"true"`
}

func (s UnlockVirtualMFADeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s UnlockVirtualMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *UnlockVirtualMFADeviceRequest) SetRegionId(v string) *UnlockVirtualMFADeviceRequest {
	s.RegionId = &v
	return s
}

func (s *UnlockVirtualMFADeviceRequest) SetSerialNumber(v string) *UnlockVirtualMFADeviceRequest {
	s.SerialNumber = &v
	return s
}

type UnlockVirtualMFADeviceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UnlockVirtualMFADeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s UnlockVirtualMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *UnlockVirtualMFADeviceResponse) SetRequestId(v string) *UnlockVirtualMFADeviceResponse {
	s.RequestId = &v
	return s
}

type DescribeVirtualMFADevicesRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	MaxResults   *int      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	DirectoryId  *string   `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	OfficeSiteId *string   `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	EndUserId    []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
}

func (s DescribeVirtualMFADevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualMFADevicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeVirtualMFADevicesRequest) SetRegionId(v string) *DescribeVirtualMFADevicesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVirtualMFADevicesRequest) SetMaxResults(v int) *DescribeVirtualMFADevicesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeVirtualMFADevicesRequest) SetNextToken(v string) *DescribeVirtualMFADevicesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeVirtualMFADevicesRequest) SetDirectoryId(v string) *DescribeVirtualMFADevicesRequest {
	s.DirectoryId = &v
	return s
}

func (s *DescribeVirtualMFADevicesRequest) SetOfficeSiteId(v string) *DescribeVirtualMFADevicesRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeVirtualMFADevicesRequest) SetEndUserId(v []*string) *DescribeVirtualMFADevicesRequest {
	s.EndUserId = v
	return s
}

type DescribeVirtualMFADevicesResponse struct {
	NextToken         *string                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	RequestId         *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	VirtualMFADevices []*DescribeVirtualMFADevicesResponseVirtualMFADevices `json:"VirtualMFADevices,omitempty" xml:"VirtualMFADevices,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeVirtualMFADevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualMFADevicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeVirtualMFADevicesResponse) SetNextToken(v string) *DescribeVirtualMFADevicesResponse {
	s.NextToken = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponse) SetRequestId(v string) *DescribeVirtualMFADevicesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponse) SetVirtualMFADevices(v []*DescribeVirtualMFADevicesResponseVirtualMFADevices) *DescribeVirtualMFADevicesResponse {
	s.VirtualMFADevices = v
	return s
}

type DescribeVirtualMFADevicesResponseVirtualMFADevices struct {
	DirectoryId      *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" require:"true"`
	OfficeSiteId     *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true"`
	EndUserId        *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" require:"true"`
	GmtEnabled       *string `json:"GmtEnabled,omitempty" xml:"GmtEnabled,omitempty" require:"true"`
	GmtUnlock        *string `json:"GmtUnlock,omitempty" xml:"GmtUnlock,omitempty" require:"true"`
	ConsecutiveFails *int    `json:"ConsecutiveFails,omitempty" xml:"ConsecutiveFails,omitempty" require:"true"`
	SerialNumber     *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty" require:"true"`
	Status           *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
}

func (s DescribeVirtualMFADevicesResponseVirtualMFADevices) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualMFADevicesResponseVirtualMFADevices) GoString() string {
	return s.String()
}

func (s *DescribeVirtualMFADevicesResponseVirtualMFADevices) SetDirectoryId(v string) *DescribeVirtualMFADevicesResponseVirtualMFADevices {
	s.DirectoryId = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseVirtualMFADevices) SetOfficeSiteId(v string) *DescribeVirtualMFADevicesResponseVirtualMFADevices {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseVirtualMFADevices) SetEndUserId(v string) *DescribeVirtualMFADevicesResponseVirtualMFADevices {
	s.EndUserId = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseVirtualMFADevices) SetGmtEnabled(v string) *DescribeVirtualMFADevicesResponseVirtualMFADevices {
	s.GmtEnabled = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseVirtualMFADevices) SetGmtUnlock(v string) *DescribeVirtualMFADevicesResponseVirtualMFADevices {
	s.GmtUnlock = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseVirtualMFADevices) SetConsecutiveFails(v int) *DescribeVirtualMFADevicesResponseVirtualMFADevices {
	s.ConsecutiveFails = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseVirtualMFADevices) SetSerialNumber(v string) *DescribeVirtualMFADevicesResponseVirtualMFADevices {
	s.SerialNumber = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseVirtualMFADevices) SetStatus(v string) *DescribeVirtualMFADevicesResponseVirtualMFADevices {
	s.Status = &v
	return s
}

type LockVirtualMFADeviceRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty" require:"true"`
}

func (s LockVirtualMFADeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s LockVirtualMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *LockVirtualMFADeviceRequest) SetRegionId(v string) *LockVirtualMFADeviceRequest {
	s.RegionId = &v
	return s
}

func (s *LockVirtualMFADeviceRequest) SetSerialNumber(v string) *LockVirtualMFADeviceRequest {
	s.SerialNumber = &v
	return s
}

type LockVirtualMFADeviceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s LockVirtualMFADeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s LockVirtualMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *LockVirtualMFADeviceResponse) SetRequestId(v string) *LockVirtualMFADeviceResponse {
	s.RequestId = &v
	return s
}

type DeleteVirtualMFADeviceRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty" require:"true"`
}

func (s DeleteVirtualMFADeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteVirtualMFADeviceRequest) SetRegionId(v string) *DeleteVirtualMFADeviceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVirtualMFADeviceRequest) SetSerialNumber(v string) *DeleteVirtualMFADeviceRequest {
	s.SerialNumber = &v
	return s
}

type DeleteVirtualMFADeviceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteVirtualMFADeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *DeleteVirtualMFADeviceResponse) SetRequestId(v string) *DeleteVirtualMFADeviceResponse {
	s.RequestId = &v
	return s
}

type ModifyADConnectorDirectoryRequest struct {
	RegionId            *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DirectoryId         *string   `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" require:"true"`
	DomainName          *string   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainUserName      *string   `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty"`
	DomainPassword      *string   `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty"`
	DnsAddress          []*string `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" type:"Repeated"`
	DirectoryName       *string   `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	SubDomainDnsAddress []*string `json:"SubDomainDnsAddress,omitempty" xml:"SubDomainDnsAddress,omitempty" type:"Repeated"`
	SubDomainName       *string   `json:"SubDomainName,omitempty" xml:"SubDomainName,omitempty"`
	MfaEnabled          *bool     `json:"MfaEnabled,omitempty" xml:"MfaEnabled,omitempty"`
}

func (s ModifyADConnectorDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyADConnectorDirectoryRequest) GoString() string {
	return s.String()
}

func (s *ModifyADConnectorDirectoryRequest) SetRegionId(v string) *ModifyADConnectorDirectoryRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetDirectoryId(v string) *ModifyADConnectorDirectoryRequest {
	s.DirectoryId = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetDomainName(v string) *ModifyADConnectorDirectoryRequest {
	s.DomainName = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetDomainUserName(v string) *ModifyADConnectorDirectoryRequest {
	s.DomainUserName = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetDomainPassword(v string) *ModifyADConnectorDirectoryRequest {
	s.DomainPassword = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetDnsAddress(v []*string) *ModifyADConnectorDirectoryRequest {
	s.DnsAddress = v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetDirectoryName(v string) *ModifyADConnectorDirectoryRequest {
	s.DirectoryName = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetSubDomainDnsAddress(v []*string) *ModifyADConnectorDirectoryRequest {
	s.SubDomainDnsAddress = v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetSubDomainName(v string) *ModifyADConnectorDirectoryRequest {
	s.SubDomainName = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetMfaEnabled(v bool) *ModifyADConnectorDirectoryRequest {
	s.MfaEnabled = &v
	return s
}

type ModifyADConnectorDirectoryResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyADConnectorDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyADConnectorDirectoryResponse) GoString() string {
	return s.String()
}

func (s *ModifyADConnectorDirectoryResponse) SetRequestId(v string) *ModifyADConnectorDirectoryResponse {
	s.RequestId = &v
	return s
}

type ListTagResourcesRequest struct {
	RegionId     *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	MaxResults   *int                          `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

func (s *ListTagResourcesRequest) SetMaxResults(v int) *ListTagResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponse struct {
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	NextToken    *string                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	TagResources []*ListTagResourcesResponseTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" require:"true" type:"Repeated"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetRequestId(v string) *ListTagResourcesResponse {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponse) SetNextToken(v string) *ListTagResourcesResponse {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponse) SetTagResources(v []*ListTagResourcesResponseTagResources) *ListTagResourcesResponse {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseTagResources struct {
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty" require:"true"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty" require:"true"`
}

func (s ListTagResourcesResponseTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseTagResources) SetResourceType(v string) *ListTagResourcesResponseTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseTagResources) SetResourceId(v string) *ListTagResourcesResponseTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseTagResources) SetTagKey(v string) *ListTagResourcesResponseTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseTagResources) SetTagValue(v string) *ListTagResourcesResponseTagResources {
	s.TagValue = &v
	return s
}

type UntagResourcesRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true" type:"Repeated"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

type UntagResourcesResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetRequestId(v string) *UntagResourcesResponse {
	s.RequestId = &v
	return s
}

type TagResourcesRequest struct {
	RegionId     *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	ResourceId   []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true" type:"Repeated"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" require:"true" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetRequestId(v string) *TagResourcesResponse {
	s.RequestId = &v
	return s
}

type ModifyDesktopNameRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	NewDesktopName *string `json:"NewDesktopName,omitempty" xml:"NewDesktopName,omitempty" require:"true"`
	DesktopId      *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
}

func (s ModifyDesktopNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopNameRequest) GoString() string {
	return s.String()
}

func (s *ModifyDesktopNameRequest) SetRegionId(v string) *ModifyDesktopNameRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDesktopNameRequest) SetNewDesktopName(v string) *ModifyDesktopNameRequest {
	s.NewDesktopName = &v
	return s
}

func (s *ModifyDesktopNameRequest) SetDesktopId(v string) *ModifyDesktopNameRequest {
	s.DesktopId = &v
	return s
}

type ModifyDesktopNameResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyDesktopNameResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopNameResponse) GoString() string {
	return s.String()
}

func (s *ModifyDesktopNameResponse) SetRequestId(v string) *ModifyDesktopNameResponse {
	s.RequestId = &v
	return s
}

type RunCommandRequest struct {
	RegionId        *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Type            *string   `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	CommandContent  *string   `json:"CommandContent,omitempty" xml:"CommandContent,omitempty" require:"true"`
	Timeout         *int64    `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	DesktopId       []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true" type:"Repeated"`
	ContentEncoding *string   `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
}

func (s RunCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s RunCommandRequest) GoString() string {
	return s.String()
}

func (s *RunCommandRequest) SetRegionId(v string) *RunCommandRequest {
	s.RegionId = &v
	return s
}

func (s *RunCommandRequest) SetType(v string) *RunCommandRequest {
	s.Type = &v
	return s
}

func (s *RunCommandRequest) SetCommandContent(v string) *RunCommandRequest {
	s.CommandContent = &v
	return s
}

func (s *RunCommandRequest) SetTimeout(v int64) *RunCommandRequest {
	s.Timeout = &v
	return s
}

func (s *RunCommandRequest) SetDesktopId(v []*string) *RunCommandRequest {
	s.DesktopId = v
	return s
}

func (s *RunCommandRequest) SetContentEncoding(v string) *RunCommandRequest {
	s.ContentEncoding = &v
	return s
}

type RunCommandResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	InvokeId  *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty" require:"true"`
}

func (s RunCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s RunCommandResponse) GoString() string {
	return s.String()
}

func (s *RunCommandResponse) SetRequestId(v string) *RunCommandResponse {
	s.RequestId = &v
	return s
}

func (s *RunCommandResponse) SetInvokeId(v string) *RunCommandResponse {
	s.InvokeId = &v
	return s
}

type StopInvocationRequest struct {
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	InvokeId  *string   `json:"InvokeId,omitempty" xml:"InvokeId,omitempty" require:"true"`
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
}

func (s StopInvocationRequest) String() string {
	return tea.Prettify(s)
}

func (s StopInvocationRequest) GoString() string {
	return s.String()
}

func (s *StopInvocationRequest) SetRegionId(v string) *StopInvocationRequest {
	s.RegionId = &v
	return s
}

func (s *StopInvocationRequest) SetInvokeId(v string) *StopInvocationRequest {
	s.InvokeId = &v
	return s
}

func (s *StopInvocationRequest) SetDesktopId(v []*string) *StopInvocationRequest {
	s.DesktopId = v
	return s
}

type StopInvocationResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s StopInvocationResponse) String() string {
	return tea.Prettify(s)
}

func (s StopInvocationResponse) GoString() string {
	return s.String()
}

func (s *StopInvocationResponse) SetRequestId(v string) *StopInvocationResponse {
	s.RequestId = &v
	return s
}

type DescribeInvocationsRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	InvokeId        *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	CommandType     *string `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	InvokeStatus    *string `json:"InvokeStatus,omitempty" xml:"InvokeStatus,omitempty"`
	DesktopId       *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	IncludeOutput   *bool   `json:"IncludeOutput,omitempty" xml:"IncludeOutput,omitempty"`
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	MaxResults      *int    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeInvocationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsRequest) SetRegionId(v string) *DescribeInvocationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetInvokeId(v string) *DescribeInvocationsRequest {
	s.InvokeId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetCommandType(v string) *DescribeInvocationsRequest {
	s.CommandType = &v
	return s
}

func (s *DescribeInvocationsRequest) SetInvokeStatus(v string) *DescribeInvocationsRequest {
	s.InvokeStatus = &v
	return s
}

func (s *DescribeInvocationsRequest) SetDesktopId(v string) *DescribeInvocationsRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetIncludeOutput(v bool) *DescribeInvocationsRequest {
	s.IncludeOutput = &v
	return s
}

func (s *DescribeInvocationsRequest) SetContentEncoding(v string) *DescribeInvocationsRequest {
	s.ContentEncoding = &v
	return s
}

func (s *DescribeInvocationsRequest) SetMaxResults(v int) *DescribeInvocationsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeInvocationsRequest) SetNextToken(v string) *DescribeInvocationsRequest {
	s.NextToken = &v
	return s
}

type DescribeInvocationsResponse struct {
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	NextToken   *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	Invocations []*DescribeInvocationsResponseInvocations `json:"Invocations,omitempty" xml:"Invocations,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeInvocationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponse) SetRequestId(v string) *DescribeInvocationsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeInvocationsResponse) SetNextToken(v string) *DescribeInvocationsResponse {
	s.NextToken = &v
	return s
}

func (s *DescribeInvocationsResponse) SetInvocations(v []*DescribeInvocationsResponseInvocations) *DescribeInvocationsResponse {
	s.Invocations = v
	return s
}

type DescribeInvocationsResponseInvocations struct {
	InvokeId         *string                                                 `json:"InvokeId,omitempty" xml:"InvokeId,omitempty" require:"true"`
	CreationTime     *string                                                 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty" require:"true"`
	CommandType      *string                                                 `json:"CommandType,omitempty" xml:"CommandType,omitempty" require:"true"`
	CommandContent   *string                                                 `json:"CommandContent,omitempty" xml:"CommandContent,omitempty" require:"true"`
	InvocationStatus *string                                                 `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty" require:"true"`
	InvokeDesktops   []*DescribeInvocationsResponseInvocationsInvokeDesktops `json:"InvokeDesktops,omitempty" xml:"InvokeDesktops,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeInvocationsResponseInvocations) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponseInvocations) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseInvocations) SetInvokeId(v string) *DescribeInvocationsResponseInvocations {
	s.InvokeId = &v
	return s
}

func (s *DescribeInvocationsResponseInvocations) SetCreationTime(v string) *DescribeInvocationsResponseInvocations {
	s.CreationTime = &v
	return s
}

func (s *DescribeInvocationsResponseInvocations) SetCommandType(v string) *DescribeInvocationsResponseInvocations {
	s.CommandType = &v
	return s
}

func (s *DescribeInvocationsResponseInvocations) SetCommandContent(v string) *DescribeInvocationsResponseInvocations {
	s.CommandContent = &v
	return s
}

func (s *DescribeInvocationsResponseInvocations) SetInvocationStatus(v string) *DescribeInvocationsResponseInvocations {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeInvocationsResponseInvocations) SetInvokeDesktops(v []*DescribeInvocationsResponseInvocationsInvokeDesktops) *DescribeInvocationsResponseInvocations {
	s.InvokeDesktops = v
	return s
}

type DescribeInvocationsResponseInvocationsInvokeDesktops struct {
	DesktopId        *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	Repeats          *int    `json:"Repeats,omitempty" xml:"Repeats,omitempty" require:"true"`
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty" require:"true"`
	CreationTime     *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty" require:"true"`
	StartTime        *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	StopTime         *string `json:"StopTime,omitempty" xml:"StopTime,omitempty" require:"true"`
	FinishTime       *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty" require:"true"`
	UpdateTime       *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	Output           *string `json:"Output,omitempty" xml:"Output,omitempty" require:"true"`
	ExitCode         *int64  `json:"ExitCode,omitempty" xml:"ExitCode,omitempty" require:"true"`
	Dropped          *int    `json:"Dropped,omitempty" xml:"Dropped,omitempty" require:"true"`
	ErrorCode        *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty" require:"true"`
	ErrorInfo        *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty" require:"true"`
}

func (s DescribeInvocationsResponseInvocationsInvokeDesktops) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponseInvocationsInvokeDesktops) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseInvocationsInvokeDesktops) SetDesktopId(v string) *DescribeInvocationsResponseInvocationsInvokeDesktops {
	s.DesktopId = &v
	return s
}

func (s *DescribeInvocationsResponseInvocationsInvokeDesktops) SetRepeats(v int) *DescribeInvocationsResponseInvocationsInvokeDesktops {
	s.Repeats = &v
	return s
}

func (s *DescribeInvocationsResponseInvocationsInvokeDesktops) SetInvocationStatus(v string) *DescribeInvocationsResponseInvocationsInvokeDesktops {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeInvocationsResponseInvocationsInvokeDesktops) SetCreationTime(v string) *DescribeInvocationsResponseInvocationsInvokeDesktops {
	s.CreationTime = &v
	return s
}

func (s *DescribeInvocationsResponseInvocationsInvokeDesktops) SetStartTime(v string) *DescribeInvocationsResponseInvocationsInvokeDesktops {
	s.StartTime = &v
	return s
}

func (s *DescribeInvocationsResponseInvocationsInvokeDesktops) SetStopTime(v string) *DescribeInvocationsResponseInvocationsInvokeDesktops {
	s.StopTime = &v
	return s
}

func (s *DescribeInvocationsResponseInvocationsInvokeDesktops) SetFinishTime(v string) *DescribeInvocationsResponseInvocationsInvokeDesktops {
	s.FinishTime = &v
	return s
}

func (s *DescribeInvocationsResponseInvocationsInvokeDesktops) SetUpdateTime(v string) *DescribeInvocationsResponseInvocationsInvokeDesktops {
	s.UpdateTime = &v
	return s
}

func (s *DescribeInvocationsResponseInvocationsInvokeDesktops) SetOutput(v string) *DescribeInvocationsResponseInvocationsInvokeDesktops {
	s.Output = &v
	return s
}

func (s *DescribeInvocationsResponseInvocationsInvokeDesktops) SetExitCode(v int64) *DescribeInvocationsResponseInvocationsInvokeDesktops {
	s.ExitCode = &v
	return s
}

func (s *DescribeInvocationsResponseInvocationsInvokeDesktops) SetDropped(v int) *DescribeInvocationsResponseInvocationsInvokeDesktops {
	s.Dropped = &v
	return s
}

func (s *DescribeInvocationsResponseInvocationsInvokeDesktops) SetErrorCode(v string) *DescribeInvocationsResponseInvocationsInvokeDesktops {
	s.ErrorCode = &v
	return s
}

func (s *DescribeInvocationsResponseInvocationsInvokeDesktops) SetErrorInfo(v string) *DescribeInvocationsResponseInvocationsInvokeDesktops {
	s.ErrorInfo = &v
	return s
}

type DescribeZonesRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s DescribeZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) SetRegionId(v string) *DescribeZonesRequest {
	s.RegionId = &v
	return s
}

type DescribeZonesResponse struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Zones     []*DescribeZonesResponseZones `json:"Zones,omitempty" xml:"Zones,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponse) SetRequestId(v string) *DescribeZonesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponse) SetZones(v []*DescribeZonesResponseZones) *DescribeZonesResponse {
	s.Zones = v
	return s
}

type DescribeZonesResponseZones struct {
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty" require:"true"`
}

func (s DescribeZonesResponseZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseZones) SetZoneId(v string) *DescribeZonesResponseZones {
	s.ZoneId = &v
	return s
}

type DescribeRegionsRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Regions   []*DescribeRegionsResponseRegions `json:"Regions,omitempty" xml:"Regions,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetRequestId(v string) *DescribeRegionsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponse) SetRegions(v []*DescribeRegionsResponseRegions) *DescribeRegionsResponse {
	s.Regions = v
	return s
}

type DescribeRegionsResponseRegions struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty" require:"true"`
}

func (s DescribeRegionsResponseRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseRegions) SetRegionId(v string) *DescribeRegionsResponseRegions {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseRegions {
	s.RegionEndpoint = &v
	return s
}

type DescribeClientEventsRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	EndUserId    *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	DesktopId    *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	DesktopIp    *string `json:"DesktopIp,omitempty" xml:"DesktopIp,omitempty"`
	DirectoryId  *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	EventType    *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MaxResults   *int    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeClientEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeClientEventsRequest) SetRegionId(v string) *DescribeClientEventsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeClientEventsRequest) SetEndUserId(v string) *DescribeClientEventsRequest {
	s.EndUserId = &v
	return s
}

func (s *DescribeClientEventsRequest) SetDesktopId(v string) *DescribeClientEventsRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeClientEventsRequest) SetDesktopIp(v string) *DescribeClientEventsRequest {
	s.DesktopIp = &v
	return s
}

func (s *DescribeClientEventsRequest) SetDirectoryId(v string) *DescribeClientEventsRequest {
	s.DirectoryId = &v
	return s
}

func (s *DescribeClientEventsRequest) SetOfficeSiteId(v string) *DescribeClientEventsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeClientEventsRequest) SetEventType(v string) *DescribeClientEventsRequest {
	s.EventType = &v
	return s
}

func (s *DescribeClientEventsRequest) SetStartTime(v string) *DescribeClientEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeClientEventsRequest) SetEndTime(v string) *DescribeClientEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeClientEventsRequest) SetMaxResults(v int) *DescribeClientEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeClientEventsRequest) SetNextToken(v string) *DescribeClientEventsRequest {
	s.NextToken = &v
	return s
}

type DescribeClientEventsResponse struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	NextToken *string                               `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	Events    []*DescribeClientEventsResponseEvents `json:"Events,omitempty" xml:"Events,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeClientEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeClientEventsResponse) SetRequestId(v string) *DescribeClientEventsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeClientEventsResponse) SetNextToken(v string) *DescribeClientEventsResponse {
	s.NextToken = &v
	return s
}

func (s *DescribeClientEventsResponse) SetEvents(v []*DescribeClientEventsResponseEvents) *DescribeClientEventsResponse {
	s.Events = v
	return s
}

type DescribeClientEventsResponseEvents struct {
	EventId        *string `json:"EventId,omitempty" xml:"EventId,omitempty" require:"true"`
	EventType      *string `json:"EventType,omitempty" xml:"EventType,omitempty" require:"true"`
	EventTime      *string `json:"EventTime,omitempty" xml:"EventTime,omitempty" require:"true"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	AliUid         *string `json:"AliUid,omitempty" xml:"AliUid,omitempty" require:"true"`
	EndUserId      *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" require:"true"`
	DesktopId      *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	DesktopIp      *string `json:"DesktopIp,omitempty" xml:"DesktopIp,omitempty" require:"true"`
	DirectoryId    *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" require:"true"`
	OfficeSiteId   *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true"`
	DirectoryType  *string `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty" require:"true"`
	OfficeSiteType *string `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty" require:"true"`
	ClientOS       *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty" require:"true"`
	ClientVersion  *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty" require:"true"`
	ClientIp       *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty" require:"true"`
	BytesSend      *string `json:"BytesSend,omitempty" xml:"BytesSend,omitempty" require:"true"`
	BytesReceived  *string `json:"BytesReceived,omitempty" xml:"BytesReceived,omitempty" require:"true"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
}

func (s DescribeClientEventsResponseEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientEventsResponseEvents) GoString() string {
	return s.String()
}

func (s *DescribeClientEventsResponseEvents) SetEventId(v string) *DescribeClientEventsResponseEvents {
	s.EventId = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetEventType(v string) *DescribeClientEventsResponseEvents {
	s.EventType = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetEventTime(v string) *DescribeClientEventsResponseEvents {
	s.EventTime = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetRegionId(v string) *DescribeClientEventsResponseEvents {
	s.RegionId = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetAliUid(v string) *DescribeClientEventsResponseEvents {
	s.AliUid = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetEndUserId(v string) *DescribeClientEventsResponseEvents {
	s.EndUserId = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetDesktopId(v string) *DescribeClientEventsResponseEvents {
	s.DesktopId = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetDesktopIp(v string) *DescribeClientEventsResponseEvents {
	s.DesktopIp = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetDirectoryId(v string) *DescribeClientEventsResponseEvents {
	s.DirectoryId = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetOfficeSiteId(v string) *DescribeClientEventsResponseEvents {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetDirectoryType(v string) *DescribeClientEventsResponseEvents {
	s.DirectoryType = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetOfficeSiteType(v string) *DescribeClientEventsResponseEvents {
	s.OfficeSiteType = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetClientOS(v string) *DescribeClientEventsResponseEvents {
	s.ClientOS = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetClientVersion(v string) *DescribeClientEventsResponseEvents {
	s.ClientVersion = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetClientIp(v string) *DescribeClientEventsResponseEvents {
	s.ClientIp = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetBytesSend(v string) *DescribeClientEventsResponseEvents {
	s.BytesSend = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetBytesReceived(v string) *DescribeClientEventsResponseEvents {
	s.BytesReceived = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetStatus(v string) *DescribeClientEventsResponseEvents {
	s.Status = &v
	return s
}

type ResetSnapshotRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty" require:"true"`
}

func (s ResetSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetSnapshotRequest) GoString() string {
	return s.String()
}

func (s *ResetSnapshotRequest) SetRegionId(v string) *ResetSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *ResetSnapshotRequest) SetSnapshotId(v string) *ResetSnapshotRequest {
	s.SnapshotId = &v
	return s
}

type ResetSnapshotResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ResetSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetSnapshotResponse) GoString() string {
	return s.String()
}

func (s *ResetSnapshotResponse) SetRequestId(v string) *ResetSnapshotResponse {
	s.RequestId = &v
	return s
}

type DeleteSnapshotRequest struct {
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	SnapshotId []*string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotRequest) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotRequest) SetRegionId(v string) *DeleteSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSnapshotRequest) SetSnapshotId(v []*string) *DeleteSnapshotRequest {
	s.SnapshotId = v
	return s
}

type DeleteSnapshotResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotResponse) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotResponse) SetRequestId(v string) *DeleteSnapshotResponse {
	s.RequestId = &v
	return s
}

type CreateSnapshotRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopId      *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	SnapshotName   *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty" require:"true"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SourceDiskType *string `json:"SourceDiskType,omitempty" xml:"SourceDiskType,omitempty" require:"true"`
}

func (s CreateSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateSnapshotRequest) SetRegionId(v string) *CreateSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSnapshotRequest) SetDesktopId(v string) *CreateSnapshotRequest {
	s.DesktopId = &v
	return s
}

func (s *CreateSnapshotRequest) SetSnapshotName(v string) *CreateSnapshotRequest {
	s.SnapshotName = &v
	return s
}

func (s *CreateSnapshotRequest) SetDescription(v string) *CreateSnapshotRequest {
	s.Description = &v
	return s
}

func (s *CreateSnapshotRequest) SetSourceDiskType(v string) *CreateSnapshotRequest {
	s.SourceDiskType = &v
	return s
}

type CreateSnapshotResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty" require:"true"`
}

func (s CreateSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotResponse) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponse) SetRequestId(v string) *CreateSnapshotResponse {
	s.RequestId = &v
	return s
}

func (s *CreateSnapshotResponse) SetSnapshotId(v string) *CreateSnapshotResponse {
	s.SnapshotId = &v
	return s
}

type DescribeSnapshotsRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopId  *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	MaxResults *int    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeSnapshotsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsRequest) SetRegionId(v string) *DescribeSnapshotsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetDesktopId(v string) *DescribeSnapshotsRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetSnapshotId(v string) *DescribeSnapshotsRequest {
	s.SnapshotId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetMaxResults(v int) *DescribeSnapshotsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetNextToken(v string) *DescribeSnapshotsRequest {
	s.NextToken = &v
	return s
}

type DescribeSnapshotsResponse struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	NextToken *string                               `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	Snapshots []*DescribeSnapshotsResponseSnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeSnapshotsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponse) SetRequestId(v string) *DescribeSnapshotsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeSnapshotsResponse) SetNextToken(v string) *DescribeSnapshotsResponse {
	s.NextToken = &v
	return s
}

func (s *DescribeSnapshotsResponse) SetSnapshots(v []*DescribeSnapshotsResponseSnapshots) *DescribeSnapshotsResponse {
	s.Snapshots = v
	return s
}

type DescribeSnapshotsResponseSnapshots struct {
	SnapshotId     *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty" require:"true"`
	DesktopId      *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	SnapshotName   *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty" require:"true"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	SnapshotType   *string `json:"SnapshotType,omitempty" xml:"SnapshotType,omitempty" require:"true"`
	SourceDiskSize *string `json:"SourceDiskSize,omitempty" xml:"SourceDiskSize,omitempty" require:"true"`
	SourceDiskType *string `json:"SourceDiskType,omitempty" xml:"SourceDiskType,omitempty" require:"true"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	CreationTime   *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty" require:"true"`
	Progress       *string `json:"Progress,omitempty" xml:"Progress,omitempty" require:"true"`
	RemainTime     *int    `json:"RemainTime,omitempty" xml:"RemainTime,omitempty" require:"true"`
}

func (s DescribeSnapshotsResponseSnapshots) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsResponseSnapshots) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseSnapshots) SetSnapshotId(v string) *DescribeSnapshotsResponseSnapshots {
	s.SnapshotId = &v
	return s
}

func (s *DescribeSnapshotsResponseSnapshots) SetDesktopId(v string) *DescribeSnapshotsResponseSnapshots {
	s.DesktopId = &v
	return s
}

func (s *DescribeSnapshotsResponseSnapshots) SetSnapshotName(v string) *DescribeSnapshotsResponseSnapshots {
	s.SnapshotName = &v
	return s
}

func (s *DescribeSnapshotsResponseSnapshots) SetDescription(v string) *DescribeSnapshotsResponseSnapshots {
	s.Description = &v
	return s
}

func (s *DescribeSnapshotsResponseSnapshots) SetSnapshotType(v string) *DescribeSnapshotsResponseSnapshots {
	s.SnapshotType = &v
	return s
}

func (s *DescribeSnapshotsResponseSnapshots) SetSourceDiskSize(v string) *DescribeSnapshotsResponseSnapshots {
	s.SourceDiskSize = &v
	return s
}

func (s *DescribeSnapshotsResponseSnapshots) SetSourceDiskType(v string) *DescribeSnapshotsResponseSnapshots {
	s.SourceDiskType = &v
	return s
}

func (s *DescribeSnapshotsResponseSnapshots) SetStatus(v string) *DescribeSnapshotsResponseSnapshots {
	s.Status = &v
	return s
}

func (s *DescribeSnapshotsResponseSnapshots) SetCreationTime(v string) *DescribeSnapshotsResponseSnapshots {
	s.CreationTime = &v
	return s
}

func (s *DescribeSnapshotsResponseSnapshots) SetProgress(v string) *DescribeSnapshotsResponseSnapshots {
	s.Progress = &v
	return s
}

func (s *DescribeSnapshotsResponseSnapshots) SetRemainTime(v int) *DescribeSnapshotsResponseSnapshots {
	s.RemainTime = &v
	return s
}

type RenewDesktopsRequest struct {
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopId  []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true" type:"Repeated"`
	Period     *int      `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit *string   `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	AutoPay    *bool     `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
}

func (s RenewDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewDesktopsRequest) GoString() string {
	return s.String()
}

func (s *RenewDesktopsRequest) SetRegionId(v string) *RenewDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *RenewDesktopsRequest) SetDesktopId(v []*string) *RenewDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *RenewDesktopsRequest) SetPeriod(v int) *RenewDesktopsRequest {
	s.Period = &v
	return s
}

func (s *RenewDesktopsRequest) SetPeriodUnit(v string) *RenewDesktopsRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewDesktopsRequest) SetAutoPay(v bool) *RenewDesktopsRequest {
	s.AutoPay = &v
	return s
}

type RenewDesktopsResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty" require:"true"`
}

func (s RenewDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewDesktopsResponse) GoString() string {
	return s.String()
}

func (s *RenewDesktopsResponse) SetRequestId(v string) *RenewDesktopsResponse {
	s.RequestId = &v
	return s
}

func (s *RenewDesktopsResponse) SetOrderId(v string) *RenewDesktopsResponse {
	s.OrderId = &v
	return s
}

type StopDesktopsRequest struct {
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopId   []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true" type:"Repeated"`
	StoppedMode *string   `json:"StoppedMode,omitempty" xml:"StoppedMode,omitempty"`
}

func (s StopDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s StopDesktopsRequest) GoString() string {
	return s.String()
}

func (s *StopDesktopsRequest) SetRegionId(v string) *StopDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *StopDesktopsRequest) SetDesktopId(v []*string) *StopDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *StopDesktopsRequest) SetStoppedMode(v string) *StopDesktopsRequest {
	s.StoppedMode = &v
	return s
}

type StopDesktopsResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s StopDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s StopDesktopsResponse) GoString() string {
	return s.String()
}

func (s *StopDesktopsResponse) SetRequestId(v string) *StopDesktopsResponse {
	s.RequestId = &v
	return s
}

type StartDesktopsRequest struct {
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true" type:"Repeated"`
}

func (s StartDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s StartDesktopsRequest) GoString() string {
	return s.String()
}

func (s *StartDesktopsRequest) SetRegionId(v string) *StartDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *StartDesktopsRequest) SetDesktopId(v []*string) *StartDesktopsRequest {
	s.DesktopId = v
	return s
}

type StartDesktopsResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s StartDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s StartDesktopsResponse) GoString() string {
	return s.String()
}

func (s *StartDesktopsResponse) SetRequestId(v string) *StartDesktopsResponse {
	s.RequestId = &v
	return s
}

type ModifyPolicyGroupRequest struct {
	RegionId                    *string                                                `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	PolicyGroupId               *string                                                `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty" require:"true"`
	Name                        *string                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	Clipboard                   *string                                                `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	LocalDrive                  *string                                                `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	UsbRedirect                 *string                                                `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	VisualQuality               *string                                                `json:"VisualQuality,omitempty" xml:"VisualQuality,omitempty"`
	Html5Access                 *string                                                `json:"Html5Access,omitempty" xml:"Html5Access,omitempty"`
	Html5FileTransfer           *string                                                `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	Watermark                   *string                                                `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	WatermarkType               *string                                                `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
	WatermarkCustomText         *string                                                `json:"WatermarkCustomText,omitempty" xml:"WatermarkCustomText,omitempty"`
	WatermarkTransparency       *string                                                `json:"WatermarkTransparency,omitempty" xml:"WatermarkTransparency,omitempty"`
	PreemptLogin                *string                                                `json:"PreemptLogin,omitempty" xml:"PreemptLogin,omitempty"`
	PreemptLoginUser            []*string                                              `json:"PreemptLoginUser,omitempty" xml:"PreemptLoginUser,omitempty" type:"Repeated"`
	DomainList                  *string                                                `json:"DomainList,omitempty" xml:"DomainList,omitempty"`
	AuthorizeSecurityPolicyRule []*ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule `json:"AuthorizeSecurityPolicyRule,omitempty" xml:"AuthorizeSecurityPolicyRule,omitempty" type:"Repeated"`
	RevokeSecurityPolicyRule    []*ModifyPolicyGroupRequestRevokeSecurityPolicyRule    `json:"RevokeSecurityPolicyRule,omitempty" xml:"RevokeSecurityPolicyRule,omitempty" type:"Repeated"`
	AuthorizeAccessPolicyRule   []*ModifyPolicyGroupRequestAuthorizeAccessPolicyRule   `json:"AuthorizeAccessPolicyRule,omitempty" xml:"AuthorizeAccessPolicyRule,omitempty" type:"Repeated"`
	RevokeAccessPolicyRule      []*ModifyPolicyGroupRequestRevokeAccessPolicyRule      `json:"RevokeAccessPolicyRule,omitempty" xml:"RevokeAccessPolicyRule,omitempty" type:"Repeated"`
}

func (s ModifyPolicyGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequest) SetRegionId(v string) *ModifyPolicyGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetPolicyGroupId(v string) *ModifyPolicyGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetName(v string) *ModifyPolicyGroupRequest {
	s.Name = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetClipboard(v string) *ModifyPolicyGroupRequest {
	s.Clipboard = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetLocalDrive(v string) *ModifyPolicyGroupRequest {
	s.LocalDrive = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetUsbRedirect(v string) *ModifyPolicyGroupRequest {
	s.UsbRedirect = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetVisualQuality(v string) *ModifyPolicyGroupRequest {
	s.VisualQuality = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetHtml5Access(v string) *ModifyPolicyGroupRequest {
	s.Html5Access = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetHtml5FileTransfer(v string) *ModifyPolicyGroupRequest {
	s.Html5FileTransfer = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWatermark(v string) *ModifyPolicyGroupRequest {
	s.Watermark = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWatermarkType(v string) *ModifyPolicyGroupRequest {
	s.WatermarkType = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWatermarkCustomText(v string) *ModifyPolicyGroupRequest {
	s.WatermarkCustomText = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWatermarkTransparency(v string) *ModifyPolicyGroupRequest {
	s.WatermarkTransparency = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetPreemptLogin(v string) *ModifyPolicyGroupRequest {
	s.PreemptLogin = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetPreemptLoginUser(v []*string) *ModifyPolicyGroupRequest {
	s.PreemptLoginUser = v
	return s
}

func (s *ModifyPolicyGroupRequest) SetDomainList(v string) *ModifyPolicyGroupRequest {
	s.DomainList = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetAuthorizeSecurityPolicyRule(v []*ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) *ModifyPolicyGroupRequest {
	s.AuthorizeSecurityPolicyRule = v
	return s
}

func (s *ModifyPolicyGroupRequest) SetRevokeSecurityPolicyRule(v []*ModifyPolicyGroupRequestRevokeSecurityPolicyRule) *ModifyPolicyGroupRequest {
	s.RevokeSecurityPolicyRule = v
	return s
}

func (s *ModifyPolicyGroupRequest) SetAuthorizeAccessPolicyRule(v []*ModifyPolicyGroupRequestAuthorizeAccessPolicyRule) *ModifyPolicyGroupRequest {
	s.AuthorizeAccessPolicyRule = v
	return s
}

func (s *ModifyPolicyGroupRequest) SetRevokeAccessPolicyRule(v []*ModifyPolicyGroupRequestRevokeAccessPolicyRule) *ModifyPolicyGroupRequest {
	s.RevokeAccessPolicyRule = v
	return s
}

type ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule struct {
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	IpProtocol  *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	PortRange   *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	CidrIp      *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	Policy      *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Priority    *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) SetType(v string) *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Type = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) SetIpProtocol(v string) *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.IpProtocol = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) SetPortRange(v string) *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.PortRange = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) SetCidrIp(v string) *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.CidrIp = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) SetPolicy(v string) *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Policy = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) SetPriority(v string) *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Priority = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) SetDescription(v string) *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Description = &v
	return s
}

type ModifyPolicyGroupRequestRevokeSecurityPolicyRule struct {
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	IpProtocol  *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	PortRange   *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	CidrIp      *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	Policy      *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Priority    *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyPolicyGroupRequestRevokeSecurityPolicyRule) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupRequestRevokeSecurityPolicyRule) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) SetType(v string) *ModifyPolicyGroupRequestRevokeSecurityPolicyRule {
	s.Type = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) SetIpProtocol(v string) *ModifyPolicyGroupRequestRevokeSecurityPolicyRule {
	s.IpProtocol = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) SetPortRange(v string) *ModifyPolicyGroupRequestRevokeSecurityPolicyRule {
	s.PortRange = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) SetCidrIp(v string) *ModifyPolicyGroupRequestRevokeSecurityPolicyRule {
	s.CidrIp = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) SetPolicy(v string) *ModifyPolicyGroupRequestRevokeSecurityPolicyRule {
	s.Policy = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) SetPriority(v string) *ModifyPolicyGroupRequestRevokeSecurityPolicyRule {
	s.Priority = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) SetDescription(v string) *ModifyPolicyGroupRequestRevokeSecurityPolicyRule {
	s.Description = &v
	return s
}

type ModifyPolicyGroupRequestAuthorizeAccessPolicyRule struct {
	CidrIp      *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyPolicyGroupRequestAuthorizeAccessPolicyRule) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupRequestAuthorizeAccessPolicyRule) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequestAuthorizeAccessPolicyRule) SetCidrIp(v string) *ModifyPolicyGroupRequestAuthorizeAccessPolicyRule {
	s.CidrIp = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeAccessPolicyRule) SetDescription(v string) *ModifyPolicyGroupRequestAuthorizeAccessPolicyRule {
	s.Description = &v
	return s
}

type ModifyPolicyGroupRequestRevokeAccessPolicyRule struct {
	CidrIp      *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyPolicyGroupRequestRevokeAccessPolicyRule) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupRequestRevokeAccessPolicyRule) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequestRevokeAccessPolicyRule) SetCidrIp(v string) *ModifyPolicyGroupRequestRevokeAccessPolicyRule {
	s.CidrIp = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeAccessPolicyRule) SetDescription(v string) *ModifyPolicyGroupRequestRevokeAccessPolicyRule {
	s.Description = &v
	return s
}

type ModifyPolicyGroupResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyPolicyGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupResponse) SetRequestId(v string) *ModifyPolicyGroupResponse {
	s.RequestId = &v
	return s
}

type DescribeDesktopTypesRequest struct {
	RegionId           *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopTypeId      *string  `json:"DesktopTypeId,omitempty" xml:"DesktopTypeId,omitempty"`
	InstanceTypeFamily *string  `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	CpuCount           *int     `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	MemorySize         *int     `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	GpuCount           *float32 `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
}

func (s DescribeDesktopTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopTypesRequest) SetRegionId(v string) *DescribeDesktopTypesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetDesktopTypeId(v string) *DescribeDesktopTypesRequest {
	s.DesktopTypeId = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetInstanceTypeFamily(v string) *DescribeDesktopTypesRequest {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetCpuCount(v int) *DescribeDesktopTypesRequest {
	s.CpuCount = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetMemorySize(v int) *DescribeDesktopTypesRequest {
	s.MemorySize = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetGpuCount(v float32) *DescribeDesktopTypesRequest {
	s.GpuCount = &v
	return s
}

type DescribeDesktopTypesResponse struct {
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DesktopTypes []*DescribeDesktopTypesResponseDesktopTypes `json:"DesktopTypes,omitempty" xml:"DesktopTypes,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDesktopTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopTypesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopTypesResponse) SetRequestId(v string) *DescribeDesktopTypesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDesktopTypesResponse) SetDesktopTypes(v []*DescribeDesktopTypesResponseDesktopTypes) *DescribeDesktopTypesResponse {
	s.DesktopTypes = v
	return s
}

type DescribeDesktopTypesResponseDesktopTypes struct {
	DesktopTypeId      *string                                                  `json:"DesktopTypeId,omitempty" xml:"DesktopTypeId,omitempty" require:"true"`
	InstanceTypeFamily *string                                                  `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty" require:"true"`
	CpuCount           *string                                                  `json:"CpuCount,omitempty" xml:"CpuCount,omitempty" require:"true"`
	MemorySize         *string                                                  `json:"MemorySize,omitempty" xml:"MemorySize,omitempty" require:"true"`
	GpuCount           *float32                                                 `json:"GpuCount,omitempty" xml:"GpuCount,omitempty" require:"true"`
	GpuSpec            *string                                                  `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty" require:"true"`
	SystemDiskSize     *string                                                  `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty" require:"true"`
	DataDiskSize       *string                                                  `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty" require:"true"`
	AllowDiskSize      []*DescribeDesktopTypesResponseDesktopTypesAllowDiskSize `json:"AllowDiskSize,omitempty" xml:"AllowDiskSize,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDesktopTypesResponseDesktopTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopTypesResponseDesktopTypes) GoString() string {
	return s.String()
}

func (s *DescribeDesktopTypesResponseDesktopTypes) SetDesktopTypeId(v string) *DescribeDesktopTypesResponseDesktopTypes {
	s.DesktopTypeId = &v
	return s
}

func (s *DescribeDesktopTypesResponseDesktopTypes) SetInstanceTypeFamily(v string) *DescribeDesktopTypesResponseDesktopTypes {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeDesktopTypesResponseDesktopTypes) SetCpuCount(v string) *DescribeDesktopTypesResponseDesktopTypes {
	s.CpuCount = &v
	return s
}

func (s *DescribeDesktopTypesResponseDesktopTypes) SetMemorySize(v string) *DescribeDesktopTypesResponseDesktopTypes {
	s.MemorySize = &v
	return s
}

func (s *DescribeDesktopTypesResponseDesktopTypes) SetGpuCount(v float32) *DescribeDesktopTypesResponseDesktopTypes {
	s.GpuCount = &v
	return s
}

func (s *DescribeDesktopTypesResponseDesktopTypes) SetGpuSpec(v string) *DescribeDesktopTypesResponseDesktopTypes {
	s.GpuSpec = &v
	return s
}

func (s *DescribeDesktopTypesResponseDesktopTypes) SetSystemDiskSize(v string) *DescribeDesktopTypesResponseDesktopTypes {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeDesktopTypesResponseDesktopTypes) SetDataDiskSize(v string) *DescribeDesktopTypesResponseDesktopTypes {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeDesktopTypesResponseDesktopTypes) SetAllowDiskSize(v []*DescribeDesktopTypesResponseDesktopTypesAllowDiskSize) *DescribeDesktopTypesResponseDesktopTypes {
	s.AllowDiskSize = v
	return s
}

type DescribeDesktopTypesResponseDesktopTypesAllowDiskSize struct {
	SystemDiskSize *int `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty" require:"true"`
	DataDiskSize   *int `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty" require:"true"`
}

func (s DescribeDesktopTypesResponseDesktopTypesAllowDiskSize) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopTypesResponseDesktopTypesAllowDiskSize) GoString() string {
	return s.String()
}

func (s *DescribeDesktopTypesResponseDesktopTypesAllowDiskSize) SetSystemDiskSize(v int) *DescribeDesktopTypesResponseDesktopTypesAllowDiskSize {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeDesktopTypesResponseDesktopTypesAllowDiskSize) SetDataDiskSize(v int) *DescribeDesktopTypesResponseDesktopTypesAllowDiskSize {
	s.DataDiskSize = &v
	return s
}

type DescribeDirectoriesRequest struct {
	RegionId        *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DirectoryType   *string   `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty"`
	DirectoryStatus *string   `json:"DirectoryStatus,omitempty" xml:"DirectoryStatus,omitempty"`
	DirectoryId     []*string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" type:"Repeated"`
	MaxResults      *int      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeDirectoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesRequest) SetRegionId(v string) *DescribeDirectoriesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetDirectoryType(v string) *DescribeDirectoriesRequest {
	s.DirectoryType = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetDirectoryStatus(v string) *DescribeDirectoriesRequest {
	s.DirectoryStatus = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetDirectoryId(v []*string) *DescribeDirectoriesRequest {
	s.DirectoryId = v
	return s
}

func (s *DescribeDirectoriesRequest) SetMaxResults(v int) *DescribeDirectoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetNextToken(v string) *DescribeDirectoriesRequest {
	s.NextToken = &v
	return s
}

type DescribeDirectoriesResponse struct {
	NextToken   *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Directories []*DescribeDirectoriesResponseDirectories `json:"Directories,omitempty" xml:"Directories,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDirectoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponse) SetNextToken(v string) *DescribeDirectoriesResponse {
	s.NextToken = &v
	return s
}

func (s *DescribeDirectoriesResponse) SetRequestId(v string) *DescribeDirectoriesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDirectoriesResponse) SetDirectories(v []*DescribeDirectoriesResponseDirectories) *DescribeDirectoriesResponse {
	s.Directories = v
	return s
}

type DescribeDirectoriesResponseDirectories struct {
	DirectoryId              *string                                               `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" require:"true"`
	Status                   *string                                               `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	DirectoryType            *string                                               `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty" require:"true"`
	CreationTime             *string                                               `json:"CreationTime,omitempty" xml:"CreationTime,omitempty" require:"true"`
	Name                     *string                                               `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	VpcId                    *string                                               `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
	CustomSecurityGroupId    *string                                               `json:"CustomSecurityGroupId,omitempty" xml:"CustomSecurityGroupId,omitempty" require:"true"`
	DnsUserName              *string                                               `json:"DnsUserName,omitempty" xml:"DnsUserName,omitempty" require:"true"`
	EnableInternetAccess     *bool                                                 `json:"EnableInternetAccess,omitempty" xml:"EnableInternetAccess,omitempty" require:"true"`
	EnableCrossDesktopAccess *bool                                                 `json:"EnableCrossDesktopAccess,omitempty" xml:"EnableCrossDesktopAccess,omitempty" require:"true"`
	EnableAdminAccess        *bool                                                 `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty" require:"true"`
	DesktopAccessType        *string                                               `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty" require:"true"`
	DesktopVpcEndpoint       *string                                               `json:"DesktopVpcEndpoint,omitempty" xml:"DesktopVpcEndpoint,omitempty" require:"true"`
	TrustPassword            *string                                               `json:"TrustPassword,omitempty" xml:"TrustPassword,omitempty" require:"true"`
	DomainName               *string                                               `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	DomainUserName           *string                                               `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty" require:"true"`
	DomainPassword           *string                                               `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty" require:"true"`
	SubDomainName            *string                                               `json:"SubDomainName,omitempty" xml:"SubDomainName,omitempty" require:"true"`
	MfaEnabled               *bool                                                 `json:"MfaEnabled,omitempty" xml:"MfaEnabled,omitempty" require:"true"`
	SsoEnabled               *bool                                                 `json:"SsoEnabled,omitempty" xml:"SsoEnabled,omitempty" require:"true"`
	ADConnectors             []*DescribeDirectoriesResponseDirectoriesADConnectors `json:"ADConnectors,omitempty" xml:"ADConnectors,omitempty" require:"true" type:"Repeated"`
	Logs                     []*DescribeDirectoriesResponseDirectoriesLogs         `json:"Logs,omitempty" xml:"Logs,omitempty" require:"true" type:"Repeated"`
	DnsAddress               []*string                                             `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" require:"true" type:"Repeated"`
	SubDnsAddress            []*string                                             `json:"SubDnsAddress,omitempty" xml:"SubDnsAddress,omitempty" require:"true" type:"Repeated"`
	VSwitchIds               []*string                                             `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" require:"true" type:"Repeated"`
	FileSystemIds            []*string                                             `json:"FileSystemIds,omitempty" xml:"FileSystemIds,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDirectoriesResponseDirectories) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponseDirectories) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseDirectories) SetDirectoryId(v string) *DescribeDirectoriesResponseDirectories {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetStatus(v string) *DescribeDirectoriesResponseDirectories {
	s.Status = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetDirectoryType(v string) *DescribeDirectoriesResponseDirectories {
	s.DirectoryType = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetCreationTime(v string) *DescribeDirectoriesResponseDirectories {
	s.CreationTime = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetName(v string) *DescribeDirectoriesResponseDirectories {
	s.Name = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetVpcId(v string) *DescribeDirectoriesResponseDirectories {
	s.VpcId = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetCustomSecurityGroupId(v string) *DescribeDirectoriesResponseDirectories {
	s.CustomSecurityGroupId = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetDnsUserName(v string) *DescribeDirectoriesResponseDirectories {
	s.DnsUserName = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetEnableInternetAccess(v bool) *DescribeDirectoriesResponseDirectories {
	s.EnableInternetAccess = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetEnableCrossDesktopAccess(v bool) *DescribeDirectoriesResponseDirectories {
	s.EnableCrossDesktopAccess = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetEnableAdminAccess(v bool) *DescribeDirectoriesResponseDirectories {
	s.EnableAdminAccess = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetDesktopAccessType(v string) *DescribeDirectoriesResponseDirectories {
	s.DesktopAccessType = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetDesktopVpcEndpoint(v string) *DescribeDirectoriesResponseDirectories {
	s.DesktopVpcEndpoint = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetTrustPassword(v string) *DescribeDirectoriesResponseDirectories {
	s.TrustPassword = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetDomainName(v string) *DescribeDirectoriesResponseDirectories {
	s.DomainName = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetDomainUserName(v string) *DescribeDirectoriesResponseDirectories {
	s.DomainUserName = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetDomainPassword(v string) *DescribeDirectoriesResponseDirectories {
	s.DomainPassword = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetSubDomainName(v string) *DescribeDirectoriesResponseDirectories {
	s.SubDomainName = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetMfaEnabled(v bool) *DescribeDirectoriesResponseDirectories {
	s.MfaEnabled = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetSsoEnabled(v bool) *DescribeDirectoriesResponseDirectories {
	s.SsoEnabled = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetADConnectors(v []*DescribeDirectoriesResponseDirectoriesADConnectors) *DescribeDirectoriesResponseDirectories {
	s.ADConnectors = v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetLogs(v []*DescribeDirectoriesResponseDirectoriesLogs) *DescribeDirectoriesResponseDirectories {
	s.Logs = v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetDnsAddress(v []*string) *DescribeDirectoriesResponseDirectories {
	s.DnsAddress = v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetSubDnsAddress(v []*string) *DescribeDirectoriesResponseDirectories {
	s.SubDnsAddress = v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetVSwitchIds(v []*string) *DescribeDirectoriesResponseDirectories {
	s.VSwitchIds = v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetFileSystemIds(v []*string) *DescribeDirectoriesResponseDirectories {
	s.FileSystemIds = v
	return s
}

type DescribeDirectoriesResponseDirectoriesADConnectors struct {
	ADConnectorAddress *string `json:"ADConnectorAddress,omitempty" xml:"ADConnectorAddress,omitempty" require:"true"`
	VSwitchId          *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" require:"true"`
	ConnectorStatus    *string `json:"ConnectorStatus,omitempty" xml:"ConnectorStatus,omitempty" require:"true"`
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty" require:"true"`
}

func (s DescribeDirectoriesResponseDirectoriesADConnectors) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponseDirectoriesADConnectors) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseDirectoriesADConnectors) SetADConnectorAddress(v string) *DescribeDirectoriesResponseDirectoriesADConnectors {
	s.ADConnectorAddress = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectoriesADConnectors) SetVSwitchId(v string) *DescribeDirectoriesResponseDirectoriesADConnectors {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectoriesADConnectors) SetConnectorStatus(v string) *DescribeDirectoriesResponseDirectoriesADConnectors {
	s.ConnectorStatus = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectoriesADConnectors) SetNetworkInterfaceId(v string) *DescribeDirectoriesResponseDirectoriesADConnectors {
	s.NetworkInterfaceId = &v
	return s
}

type DescribeDirectoriesResponseDirectoriesLogs struct {
	Level     *string `json:"Level,omitempty" xml:"Level,omitempty" require:"true"`
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty" require:"true"`
	Step      *string `json:"Step,omitempty" xml:"Step,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
}

func (s DescribeDirectoriesResponseDirectoriesLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponseDirectoriesLogs) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseDirectoriesLogs) SetLevel(v string) *DescribeDirectoriesResponseDirectoriesLogs {
	s.Level = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectoriesLogs) SetTimeStamp(v string) *DescribeDirectoriesResponseDirectoriesLogs {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectoriesLogs) SetStep(v string) *DescribeDirectoriesResponseDirectoriesLogs {
	s.Step = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectoriesLogs) SetMessage(v string) *DescribeDirectoriesResponseDirectoriesLogs {
	s.Message = &v
	return s
}

type DeleteDirectoriesRequest struct {
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DirectoryId []*string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteDirectoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *DeleteDirectoriesRequest) SetRegionId(v string) *DeleteDirectoriesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDirectoriesRequest) SetDirectoryId(v []*string) *DeleteDirectoriesRequest {
	s.DirectoryId = v
	return s
}

type DeleteDirectoriesResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteDirectoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *DeleteDirectoriesResponse) SetRequestId(v string) *DeleteDirectoriesResponse {
	s.RequestId = &v
	return s
}

type ListDirectoryUsersRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Filter      *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" require:"true"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults  *int    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListDirectoryUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDirectoryUsersRequest) GoString() string {
	return s.String()
}

func (s *ListDirectoryUsersRequest) SetRegionId(v string) *ListDirectoryUsersRequest {
	s.RegionId = &v
	return s
}

func (s *ListDirectoryUsersRequest) SetFilter(v string) *ListDirectoryUsersRequest {
	s.Filter = &v
	return s
}

func (s *ListDirectoryUsersRequest) SetDirectoryId(v string) *ListDirectoryUsersRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListDirectoryUsersRequest) SetNextToken(v string) *ListDirectoryUsersRequest {
	s.NextToken = &v
	return s
}

func (s *ListDirectoryUsersRequest) SetMaxResults(v int) *ListDirectoryUsersRequest {
	s.MaxResults = &v
	return s
}

type ListDirectoryUsersResponse struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	NextToken *string                            `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	Users     []*ListDirectoryUsersResponseUsers `json:"Users,omitempty" xml:"Users,omitempty" require:"true" type:"Repeated"`
}

func (s ListDirectoryUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDirectoryUsersResponse) GoString() string {
	return s.String()
}

func (s *ListDirectoryUsersResponse) SetRequestId(v string) *ListDirectoryUsersResponse {
	s.RequestId = &v
	return s
}

func (s *ListDirectoryUsersResponse) SetNextToken(v string) *ListDirectoryUsersResponse {
	s.NextToken = &v
	return s
}

func (s *ListDirectoryUsersResponse) SetUsers(v []*ListDirectoryUsersResponseUsers) *ListDirectoryUsersResponse {
	s.Users = v
	return s
}

type ListDirectoryUsersResponseUsers struct {
	EndUser *string `json:"EndUser,omitempty" xml:"EndUser,omitempty" require:"true"`
}

func (s ListDirectoryUsersResponseUsers) String() string {
	return tea.Prettify(s)
}

func (s ListDirectoryUsersResponseUsers) GoString() string {
	return s.String()
}

func (s *ListDirectoryUsersResponseUsers) SetEndUser(v string) *ListDirectoryUsersResponseUsers {
	s.EndUser = &v
	return s
}

type CreateImageRequest struct {
	RegionId          *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopId         *string   `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	ImageName         *string   `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	Description       *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	SnapshotIds       []*string `json:"SnapshotIds,omitempty" xml:"SnapshotIds,omitempty" type:"Repeated"`
	SnapshotId        *string   `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	ImageResourceType *string   `json:"ImageResourceType,omitempty" xml:"ImageResourceType,omitempty"`
}

func (s CreateImageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageRequest) GoString() string {
	return s.String()
}

func (s *CreateImageRequest) SetRegionId(v string) *CreateImageRequest {
	s.RegionId = &v
	return s
}

func (s *CreateImageRequest) SetDesktopId(v string) *CreateImageRequest {
	s.DesktopId = &v
	return s
}

func (s *CreateImageRequest) SetImageName(v string) *CreateImageRequest {
	s.ImageName = &v
	return s
}

func (s *CreateImageRequest) SetDescription(v string) *CreateImageRequest {
	s.Description = &v
	return s
}

func (s *CreateImageRequest) SetSnapshotIds(v []*string) *CreateImageRequest {
	s.SnapshotIds = v
	return s
}

func (s *CreateImageRequest) SetSnapshotId(v string) *CreateImageRequest {
	s.SnapshotId = &v
	return s
}

func (s *CreateImageRequest) SetImageResourceType(v string) *CreateImageRequest {
	s.ImageResourceType = &v
	return s
}

type CreateImageResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
}

func (s CreateImageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateImageResponse) GoString() string {
	return s.String()
}

func (s *CreateImageResponse) SetRequestId(v string) *CreateImageResponse {
	s.RequestId = &v
	return s
}

func (s *CreateImageResponse) SetImageId(v string) *CreateImageResponse {
	s.ImageId = &v
	return s
}

type CreateRAMDirectoryRequest struct {
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DirectoryName        *string   `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	EnableInternetAccess *bool     `json:"EnableInternetAccess,omitempty" xml:"EnableInternetAccess,omitempty"`
	EnableAdminAccess    *bool     `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	DesktopAccessType    *string   `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	VSwitchId            []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" require:"true" type:"Repeated"`
}

func (s CreateRAMDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRAMDirectoryRequest) GoString() string {
	return s.String()
}

func (s *CreateRAMDirectoryRequest) SetRegionId(v string) *CreateRAMDirectoryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRAMDirectoryRequest) SetDirectoryName(v string) *CreateRAMDirectoryRequest {
	s.DirectoryName = &v
	return s
}

func (s *CreateRAMDirectoryRequest) SetEnableInternetAccess(v bool) *CreateRAMDirectoryRequest {
	s.EnableInternetAccess = &v
	return s
}

func (s *CreateRAMDirectoryRequest) SetEnableAdminAccess(v bool) *CreateRAMDirectoryRequest {
	s.EnableAdminAccess = &v
	return s
}

func (s *CreateRAMDirectoryRequest) SetDesktopAccessType(v string) *CreateRAMDirectoryRequest {
	s.DesktopAccessType = &v
	return s
}

func (s *CreateRAMDirectoryRequest) SetVSwitchId(v []*string) *CreateRAMDirectoryRequest {
	s.VSwitchId = v
	return s
}

type CreateRAMDirectoryResponse struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" require:"true"`
}

func (s CreateRAMDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRAMDirectoryResponse) GoString() string {
	return s.String()
}

func (s *CreateRAMDirectoryResponse) SetRequestId(v string) *CreateRAMDirectoryResponse {
	s.RequestId = &v
	return s
}

func (s *CreateRAMDirectoryResponse) SetDirectoryId(v string) *CreateRAMDirectoryResponse {
	s.DirectoryId = &v
	return s
}

type DeletePolicyGroupsRequest struct {
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	PolicyGroupId []*string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty" require:"true" type:"Repeated"`
}

func (s DeletePolicyGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyGroupsRequest) GoString() string {
	return s.String()
}

func (s *DeletePolicyGroupsRequest) SetRegionId(v string) *DeletePolicyGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePolicyGroupsRequest) SetPolicyGroupId(v []*string) *DeletePolicyGroupsRequest {
	s.PolicyGroupId = v
	return s
}

type DeletePolicyGroupsResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeletePolicyGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyGroupsResponse) GoString() string {
	return s.String()
}

func (s *DeletePolicyGroupsResponse) SetRequestId(v string) *DeletePolicyGroupsResponse {
	s.RequestId = &v
	return s
}

type DescribePolicyGroupsRequest struct {
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	MaxResults    *int      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PolicyGroupId []*string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty" type:"Repeated"`
}

func (s DescribePolicyGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsRequest) SetRegionId(v string) *DescribePolicyGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePolicyGroupsRequest) SetMaxResults(v int) *DescribePolicyGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribePolicyGroupsRequest) SetNextToken(v string) *DescribePolicyGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribePolicyGroupsRequest) SetPolicyGroupId(v []*string) *DescribePolicyGroupsRequest {
	s.PolicyGroupId = v
	return s
}

type DescribePolicyGroupsResponse struct {
	NextToken            *string                                             `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	RequestId            *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DescribePolicyGroups []*DescribePolicyGroupsResponseDescribePolicyGroups `json:"DescribePolicyGroups,omitempty" xml:"DescribePolicyGroups,omitempty" require:"true" type:"Repeated"`
}

func (s DescribePolicyGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponse) SetNextToken(v string) *DescribePolicyGroupsResponse {
	s.NextToken = &v
	return s
}

func (s *DescribePolicyGroupsResponse) SetRequestId(v string) *DescribePolicyGroupsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribePolicyGroupsResponse) SetDescribePolicyGroups(v []*DescribePolicyGroupsResponseDescribePolicyGroups) *DescribePolicyGroupsResponse {
	s.DescribePolicyGroups = v
	return s
}

type DescribePolicyGroupsResponseDescribePolicyGroups struct {
	PolicyGroupId                *string                                                                         `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty" require:"true"`
	PolicyGroupType              *string                                                                         `json:"PolicyGroupType,omitempty" xml:"PolicyGroupType,omitempty" require:"true"`
	Clipboard                    *string                                                                         `json:"Clipboard,omitempty" xml:"Clipboard,omitempty" require:"true"`
	LocalDrive                   *string                                                                         `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty" require:"true"`
	UsbRedirect                  *string                                                                         `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty" require:"true"`
	VisualQuality                *string                                                                         `json:"VisualQuality,omitempty" xml:"VisualQuality,omitempty" require:"true"`
	Html5Access                  *string                                                                         `json:"Html5Access,omitempty" xml:"Html5Access,omitempty" require:"true"`
	Html5FileTransfer            *string                                                                         `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty" require:"true"`
	Watermark                    *string                                                                         `json:"Watermark,omitempty" xml:"Watermark,omitempty" require:"true"`
	Name                         *string                                                                         `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	WatermarkType                *string                                                                         `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty" require:"true"`
	WatermarkCustomText          *string                                                                         `json:"WatermarkCustomText,omitempty" xml:"WatermarkCustomText,omitempty" require:"true"`
	WatermarkTransparency        *string                                                                         `json:"WatermarkTransparency,omitempty" xml:"WatermarkTransparency,omitempty" require:"true"`
	PolicyStatus                 *string                                                                         `json:"PolicyStatus,omitempty" xml:"PolicyStatus,omitempty" require:"true"`
	EdsCount                     *int                                                                            `json:"EdsCount,omitempty" xml:"EdsCount,omitempty" require:"true"`
	PreemptLogin                 *string                                                                         `json:"PreemptLogin,omitempty" xml:"PreemptLogin,omitempty" require:"true"`
	DomainList                   *string                                                                         `json:"DomainList,omitempty" xml:"DomainList,omitempty" require:"true"`
	AuthorizeSecurityPolicyRules []*DescribePolicyGroupsResponseDescribePolicyGroupsAuthorizeSecurityPolicyRules `json:"AuthorizeSecurityPolicyRules,omitempty" xml:"AuthorizeSecurityPolicyRules,omitempty" require:"true" type:"Repeated"`
	AuthorizeAccessPolicyRules   []*DescribePolicyGroupsResponseDescribePolicyGroupsAuthorizeAccessPolicyRules   `json:"AuthorizeAccessPolicyRules,omitempty" xml:"AuthorizeAccessPolicyRules,omitempty" require:"true" type:"Repeated"`
	PreemptLoginUsers            []*string                                                                       `json:"PreemptLoginUsers,omitempty" xml:"PreemptLoginUsers,omitempty" require:"true" type:"Repeated"`
}

func (s DescribePolicyGroupsResponseDescribePolicyGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGroupsResponseDescribePolicyGroups) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetPolicyGroupId(v string) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetPolicyGroupType(v string) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.PolicyGroupType = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetClipboard(v string) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.Clipboard = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetLocalDrive(v string) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.LocalDrive = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetUsbRedirect(v string) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.UsbRedirect = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetVisualQuality(v string) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.VisualQuality = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetHtml5Access(v string) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.Html5Access = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetHtml5FileTransfer(v string) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.Html5FileTransfer = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetWatermark(v string) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.Watermark = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetName(v string) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.Name = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetWatermarkType(v string) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.WatermarkType = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetWatermarkCustomText(v string) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.WatermarkCustomText = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetWatermarkTransparency(v string) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.WatermarkTransparency = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetPolicyStatus(v string) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.PolicyStatus = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetEdsCount(v int) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.EdsCount = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetPreemptLogin(v string) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.PreemptLogin = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetDomainList(v string) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.DomainList = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetAuthorizeSecurityPolicyRules(v []*DescribePolicyGroupsResponseDescribePolicyGroupsAuthorizeSecurityPolicyRules) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.AuthorizeSecurityPolicyRules = v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetAuthorizeAccessPolicyRules(v []*DescribePolicyGroupsResponseDescribePolicyGroupsAuthorizeAccessPolicyRules) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.AuthorizeAccessPolicyRules = v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetPreemptLoginUsers(v []*string) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.PreemptLoginUsers = v
	return s
}

type DescribePolicyGroupsResponseDescribePolicyGroupsAuthorizeSecurityPolicyRules struct {
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	IpProtocol  *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty" require:"true"`
	PortRange   *string `json:"PortRange,omitempty" xml:"PortRange,omitempty" require:"true"`
	CidrIp      *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty" require:"true"`
	Policy      *string `json:"Policy,omitempty" xml:"Policy,omitempty" require:"true"`
	Priority    *string `json:"Priority,omitempty" xml:"Priority,omitempty" require:"true"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
}

func (s DescribePolicyGroupsResponseDescribePolicyGroupsAuthorizeSecurityPolicyRules) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGroupsResponseDescribePolicyGroupsAuthorizeSecurityPolicyRules) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetType(v string) *DescribePolicyGroupsResponseDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.Type = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetIpProtocol(v string) *DescribePolicyGroupsResponseDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.IpProtocol = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetPortRange(v string) *DescribePolicyGroupsResponseDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.PortRange = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetCidrIp(v string) *DescribePolicyGroupsResponseDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.CidrIp = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetPolicy(v string) *DescribePolicyGroupsResponseDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.Policy = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetPriority(v string) *DescribePolicyGroupsResponseDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.Priority = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetDescription(v string) *DescribePolicyGroupsResponseDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.Description = &v
	return s
}

type DescribePolicyGroupsResponseDescribePolicyGroupsAuthorizeAccessPolicyRules struct {
	CidrIp      *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty" require:"true"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
}

func (s DescribePolicyGroupsResponseDescribePolicyGroupsAuthorizeAccessPolicyRules) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGroupsResponseDescribePolicyGroupsAuthorizeAccessPolicyRules) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroupsAuthorizeAccessPolicyRules) SetCidrIp(v string) *DescribePolicyGroupsResponseDescribePolicyGroupsAuthorizeAccessPolicyRules {
	s.CidrIp = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroupsAuthorizeAccessPolicyRules) SetDescription(v string) *DescribePolicyGroupsResponseDescribePolicyGroupsAuthorizeAccessPolicyRules {
	s.Description = &v
	return s
}

type DeleteDesktopsRequest struct {
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDesktopsRequest) GoString() string {
	return s.String()
}

func (s *DeleteDesktopsRequest) SetRegionId(v string) *DeleteDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDesktopsRequest) SetDesktopId(v []*string) *DeleteDesktopsRequest {
	s.DesktopId = v
	return s
}

type DeleteDesktopsResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDesktopsResponse) GoString() string {
	return s.String()
}

func (s *DeleteDesktopsResponse) SetRequestId(v string) *DeleteDesktopsResponse {
	s.RequestId = &v
	return s
}

type ModifyImageAttributeRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ImageId     *string `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyImageAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyImageAttributeRequest) SetRegionId(v string) *ModifyImageAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetImageId(v string) *ModifyImageAttributeRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetName(v string) *ModifyImageAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetDescription(v string) *ModifyImageAttributeRequest {
	s.Description = &v
	return s
}

type ModifyImageAttributeResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyImageAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyImageAttributeResponse) SetRequestId(v string) *ModifyImageAttributeResponse {
	s.RequestId = &v
	return s
}

type ModifyEntitlementRequest struct {
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopId *string   `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	EndUserId []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
}

func (s ModifyEntitlementRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyEntitlementRequest) GoString() string {
	return s.String()
}

func (s *ModifyEntitlementRequest) SetRegionId(v string) *ModifyEntitlementRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyEntitlementRequest) SetDesktopId(v string) *ModifyEntitlementRequest {
	s.DesktopId = &v
	return s
}

func (s *ModifyEntitlementRequest) SetEndUserId(v []*string) *ModifyEntitlementRequest {
	s.EndUserId = v
	return s
}

type ModifyEntitlementResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyEntitlementResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyEntitlementResponse) GoString() string {
	return s.String()
}

func (s *ModifyEntitlementResponse) SetRequestId(v string) *ModifyEntitlementResponse {
	s.RequestId = &v
	return s
}

type DeleteBundlesRequest struct {
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	BundleId []*string `json:"BundleId,omitempty" xml:"BundleId,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteBundlesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBundlesRequest) GoString() string {
	return s.String()
}

func (s *DeleteBundlesRequest) SetRegionId(v string) *DeleteBundlesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteBundlesRequest) SetBundleId(v []*string) *DeleteBundlesRequest {
	s.BundleId = v
	return s
}

type DeleteBundlesResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteBundlesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBundlesResponse) GoString() string {
	return s.String()
}

func (s *DeleteBundlesResponse) SetRequestId(v string) *DeleteBundlesResponse {
	s.RequestId = &v
	return s
}

type DescribeDesktopsRequest struct {
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	GroupId       *string   `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	DesktopStatus *string   `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	MaxResults    *int      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	UserName      *string   `json:"UserName,omitempty" xml:"UserName,omitempty"`
	DesktopName   *string   `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	DirectoryId   *string   `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	OfficeSiteId  *string   `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	PolicyGroupId *string   `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	DesktopId     []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	EndUserId     []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
	ChargeType    *string   `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ExpiredTime   *string   `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	ProtocolType  *string   `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
}

func (s DescribeDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsRequest) SetRegionId(v string) *DescribeDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetGroupId(v string) *DescribeDesktopsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetDesktopStatus(v string) *DescribeDesktopsRequest {
	s.DesktopStatus = &v
	return s
}

func (s *DescribeDesktopsRequest) SetMaxResults(v int) *DescribeDesktopsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDesktopsRequest) SetNextToken(v string) *DescribeDesktopsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopsRequest) SetUserName(v string) *DescribeDesktopsRequest {
	s.UserName = &v
	return s
}

func (s *DescribeDesktopsRequest) SetDesktopName(v string) *DescribeDesktopsRequest {
	s.DesktopName = &v
	return s
}

func (s *DescribeDesktopsRequest) SetDirectoryId(v string) *DescribeDesktopsRequest {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetOfficeSiteId(v string) *DescribeDesktopsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetPolicyGroupId(v string) *DescribeDesktopsRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetDesktopId(v []*string) *DescribeDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *DescribeDesktopsRequest) SetEndUserId(v []*string) *DescribeDesktopsRequest {
	s.EndUserId = v
	return s
}

func (s *DescribeDesktopsRequest) SetChargeType(v string) *DescribeDesktopsRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeDesktopsRequest) SetExpiredTime(v string) *DescribeDesktopsRequest {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDesktopsRequest) SetProtocolType(v string) *DescribeDesktopsRequest {
	s.ProtocolType = &v
	return s
}

type DescribeDesktopsResponse struct {
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount *int                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	NextToken  *string                             `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	Desktops   []*DescribeDesktopsResponseDesktops `json:"Desktops,omitempty" xml:"Desktops,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponse) SetRequestId(v string) *DescribeDesktopsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDesktopsResponse) SetTotalCount(v int) *DescribeDesktopsResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeDesktopsResponse) SetNextToken(v string) *DescribeDesktopsResponse {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopsResponse) SetDesktops(v []*DescribeDesktopsResponseDesktops) *DescribeDesktopsResponse {
	s.Desktops = v
	return s
}

type DescribeDesktopsResponseDesktops struct {
	DirectoryId        *string                                     `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" require:"true"`
	OfficeSiteId       *string                                     `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" require:"true"`
	OfficeSiteName     *string                                     `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty" require:"true"`
	DirectoryType      *string                                     `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty" require:"true"`
	OfficeSiteType     *string                                     `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty" require:"true"`
	CreationTime       *string                                     `json:"CreationTime,omitempty" xml:"CreationTime,omitempty" require:"true"`
	DesktopId          *string                                     `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	BundleId           *string                                     `json:"BundleId,omitempty" xml:"BundleId,omitempty" require:"true"`
	DesktopStatus      *string                                     `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty" require:"true"`
	DesktopName        *string                                     `json:"DesktopName,omitempty" xml:"DesktopName,omitempty" require:"true"`
	ImageId            *string                                     `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	DesktopType        *string                                     `json:"DesktopType,omitempty" xml:"DesktopType,omitempty" require:"true"`
	SystemDiskCategory *string                                     `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty" require:"true"`
	SystemDiskSize     *int                                        `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty" require:"true"`
	DataDiskCategory   *string                                     `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty" require:"true"`
	DataDiskSize       *string                                     `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty" require:"true"`
	PolicyGroupId      *string                                     `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty" require:"true"`
	PolicyGroupName    *string                                     `json:"PolicyGroupName,omitempty" xml:"PolicyGroupName,omitempty" require:"true"`
	Cpu                *int                                        `json:"Cpu,omitempty" xml:"Cpu,omitempty" require:"true"`
	Memory             *int64                                      `json:"Memory,omitempty" xml:"Memory,omitempty" require:"true"`
	NetworkInterfaceId *string                                     `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty" require:"true"`
	ExpiredTime        *string                                     `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty" require:"true"`
	ChargeType         *string                                     `json:"ChargeType,omitempty" xml:"ChargeType,omitempty" require:"true"`
	GpuCount           *float32                                    `json:"GpuCount,omitempty" xml:"GpuCount,omitempty" require:"true"`
	GpuSpec            *string                                     `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty" require:"true"`
	StartTime          *string                                     `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	ConnectionStatus   *string                                     `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty" require:"true"`
	OsType             *string                                     `json:"OsType,omitempty" xml:"OsType,omitempty" require:"true"`
	ProtocolType       *string                                     `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty" require:"true"`
	ManagementFlag     *string                                     `json:"ManagementFlag,omitempty" xml:"ManagementFlag,omitempty" require:"true"`
	NetworkInterfaceIp *string                                     `json:"NetworkInterfaceIp,omitempty" xml:"NetworkInterfaceIp,omitempty" require:"true"`
	DesktopGroupId     *string                                     `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty" require:"true"`
	Disks              []*DescribeDesktopsResponseDesktopsDisks    `json:"Disks,omitempty" xml:"Disks,omitempty" require:"true" type:"Repeated"`
	Tags               []*DescribeDesktopsResponseDesktopsTags     `json:"Tags,omitempty" xml:"Tags,omitempty" require:"true" type:"Repeated"`
	Sessions           []*DescribeDesktopsResponseDesktopsSessions `json:"Sessions,omitempty" xml:"Sessions,omitempty" require:"true" type:"Repeated"`
	EndUserIds         []*string                                   `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDesktopsResponseDesktops) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsResponseDesktops) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseDesktops) SetDirectoryId(v string) *DescribeDesktopsResponseDesktops {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetOfficeSiteId(v string) *DescribeDesktopsResponseDesktops {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetOfficeSiteName(v string) *DescribeDesktopsResponseDesktops {
	s.OfficeSiteName = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetDirectoryType(v string) *DescribeDesktopsResponseDesktops {
	s.DirectoryType = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetOfficeSiteType(v string) *DescribeDesktopsResponseDesktops {
	s.OfficeSiteType = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetCreationTime(v string) *DescribeDesktopsResponseDesktops {
	s.CreationTime = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetDesktopId(v string) *DescribeDesktopsResponseDesktops {
	s.DesktopId = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetBundleId(v string) *DescribeDesktopsResponseDesktops {
	s.BundleId = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetDesktopStatus(v string) *DescribeDesktopsResponseDesktops {
	s.DesktopStatus = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetDesktopName(v string) *DescribeDesktopsResponseDesktops {
	s.DesktopName = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetImageId(v string) *DescribeDesktopsResponseDesktops {
	s.ImageId = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetDesktopType(v string) *DescribeDesktopsResponseDesktops {
	s.DesktopType = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetSystemDiskCategory(v string) *DescribeDesktopsResponseDesktops {
	s.SystemDiskCategory = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetSystemDiskSize(v int) *DescribeDesktopsResponseDesktops {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetDataDiskCategory(v string) *DescribeDesktopsResponseDesktops {
	s.DataDiskCategory = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetDataDiskSize(v string) *DescribeDesktopsResponseDesktops {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetPolicyGroupId(v string) *DescribeDesktopsResponseDesktops {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetPolicyGroupName(v string) *DescribeDesktopsResponseDesktops {
	s.PolicyGroupName = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetCpu(v int) *DescribeDesktopsResponseDesktops {
	s.Cpu = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetMemory(v int64) *DescribeDesktopsResponseDesktops {
	s.Memory = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetNetworkInterfaceId(v string) *DescribeDesktopsResponseDesktops {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetExpiredTime(v string) *DescribeDesktopsResponseDesktops {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetChargeType(v string) *DescribeDesktopsResponseDesktops {
	s.ChargeType = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetGpuCount(v float32) *DescribeDesktopsResponseDesktops {
	s.GpuCount = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetGpuSpec(v string) *DescribeDesktopsResponseDesktops {
	s.GpuSpec = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetStartTime(v string) *DescribeDesktopsResponseDesktops {
	s.StartTime = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetConnectionStatus(v string) *DescribeDesktopsResponseDesktops {
	s.ConnectionStatus = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetOsType(v string) *DescribeDesktopsResponseDesktops {
	s.OsType = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetProtocolType(v string) *DescribeDesktopsResponseDesktops {
	s.ProtocolType = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetManagementFlag(v string) *DescribeDesktopsResponseDesktops {
	s.ManagementFlag = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetNetworkInterfaceIp(v string) *DescribeDesktopsResponseDesktops {
	s.NetworkInterfaceIp = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetDesktopGroupId(v string) *DescribeDesktopsResponseDesktops {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetDisks(v []*DescribeDesktopsResponseDesktopsDisks) *DescribeDesktopsResponseDesktops {
	s.Disks = v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetTags(v []*DescribeDesktopsResponseDesktopsTags) *DescribeDesktopsResponseDesktops {
	s.Tags = v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetSessions(v []*DescribeDesktopsResponseDesktopsSessions) *DescribeDesktopsResponseDesktops {
	s.Sessions = v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetEndUserIds(v []*string) *DescribeDesktopsResponseDesktops {
	s.EndUserIds = v
	return s
}

type DescribeDesktopsResponseDesktopsDisks struct {
	DiskId   *string `json:"DiskId,omitempty" xml:"DiskId,omitempty" require:"true"`
	DiskSize *int    `json:"DiskSize,omitempty" xml:"DiskSize,omitempty" require:"true"`
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty" require:"true"`
}

func (s DescribeDesktopsResponseDesktopsDisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsResponseDesktopsDisks) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseDesktopsDisks) SetDiskId(v string) *DescribeDesktopsResponseDesktopsDisks {
	s.DiskId = &v
	return s
}

func (s *DescribeDesktopsResponseDesktopsDisks) SetDiskSize(v int) *DescribeDesktopsResponseDesktopsDisks {
	s.DiskSize = &v
	return s
}

func (s *DescribeDesktopsResponseDesktopsDisks) SetDiskType(v string) *DescribeDesktopsResponseDesktopsDisks {
	s.DiskType = &v
	return s
}

type DescribeDesktopsResponseDesktopsTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty" require:"true"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s DescribeDesktopsResponseDesktopsTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsResponseDesktopsTags) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseDesktopsTags) SetKey(v string) *DescribeDesktopsResponseDesktopsTags {
	s.Key = &v
	return s
}

func (s *DescribeDesktopsResponseDesktopsTags) SetValue(v string) *DescribeDesktopsResponseDesktopsTags {
	s.Value = &v
	return s
}

type DescribeDesktopsResponseDesktopsSessions struct {
	EndUserId         *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" require:"true"`
	EstablishmentTime *string `json:"EstablishmentTime,omitempty" xml:"EstablishmentTime,omitempty" require:"true"`
}

func (s DescribeDesktopsResponseDesktopsSessions) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsResponseDesktopsSessions) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseDesktopsSessions) SetEndUserId(v string) *DescribeDesktopsResponseDesktopsSessions {
	s.EndUserId = &v
	return s
}

func (s *DescribeDesktopsResponseDesktopsSessions) SetEstablishmentTime(v string) *DescribeDesktopsResponseDesktopsSessions {
	s.EstablishmentTime = &v
	return s
}

type RebootDesktopsRequest struct {
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true" type:"Repeated"`
}

func (s RebootDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s RebootDesktopsRequest) GoString() string {
	return s.String()
}

func (s *RebootDesktopsRequest) SetRegionId(v string) *RebootDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *RebootDesktopsRequest) SetDesktopId(v []*string) *RebootDesktopsRequest {
	s.DesktopId = v
	return s
}

type RebootDesktopsResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s RebootDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s RebootDesktopsResponse) GoString() string {
	return s.String()
}

func (s *RebootDesktopsResponse) SetRequestId(v string) *RebootDesktopsResponse {
	s.RequestId = &v
	return s
}

type CreateBundleRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ImageId         *string `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	DesktopType     *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty" require:"true"`
	RootDiskSizeGib *int    `json:"RootDiskSizeGib,omitempty" xml:"RootDiskSizeGib,omitempty" require:"true"`
	BundleName      *string `json:"BundleName,omitempty" xml:"BundleName,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	UserDiskSizeGib []*int  `json:"UserDiskSizeGib,omitempty" xml:"UserDiskSizeGib,omitempty" require:"true" type:"Repeated"`
}

func (s CreateBundleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBundleRequest) GoString() string {
	return s.String()
}

func (s *CreateBundleRequest) SetRegionId(v string) *CreateBundleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateBundleRequest) SetImageId(v string) *CreateBundleRequest {
	s.ImageId = &v
	return s
}

func (s *CreateBundleRequest) SetDesktopType(v string) *CreateBundleRequest {
	s.DesktopType = &v
	return s
}

func (s *CreateBundleRequest) SetRootDiskSizeGib(v int) *CreateBundleRequest {
	s.RootDiskSizeGib = &v
	return s
}

func (s *CreateBundleRequest) SetBundleName(v string) *CreateBundleRequest {
	s.BundleName = &v
	return s
}

func (s *CreateBundleRequest) SetDescription(v string) *CreateBundleRequest {
	s.Description = &v
	return s
}

func (s *CreateBundleRequest) SetUserDiskSizeGib(v []*int) *CreateBundleRequest {
	s.UserDiskSizeGib = v
	return s
}

type CreateBundleResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	BundleId  *string `json:"BundleId,omitempty" xml:"BundleId,omitempty" require:"true"`
}

func (s CreateBundleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBundleResponse) GoString() string {
	return s.String()
}

func (s *CreateBundleResponse) SetRequestId(v string) *CreateBundleResponse {
	s.RequestId = &v
	return s
}

func (s *CreateBundleResponse) SetBundleId(v string) *CreateBundleResponse {
	s.BundleId = &v
	return s
}

type ModifyDesktopsPolicyGroupRequest struct {
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	PolicyGroupId *string   `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty" require:"true"`
	DesktopId     []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true" type:"Repeated"`
}

func (s ModifyDesktopsPolicyGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopsPolicyGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyDesktopsPolicyGroupRequest) SetRegionId(v string) *ModifyDesktopsPolicyGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDesktopsPolicyGroupRequest) SetPolicyGroupId(v string) *ModifyDesktopsPolicyGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ModifyDesktopsPolicyGroupRequest) SetDesktopId(v []*string) *ModifyDesktopsPolicyGroupRequest {
	s.DesktopId = v
	return s
}

type ModifyDesktopsPolicyGroupResponse struct {
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ModifyResults []*ModifyDesktopsPolicyGroupResponseModifyResults `json:"ModifyResults,omitempty" xml:"ModifyResults,omitempty" require:"true" type:"Repeated"`
}

func (s ModifyDesktopsPolicyGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopsPolicyGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyDesktopsPolicyGroupResponse) SetRequestId(v string) *ModifyDesktopsPolicyGroupResponse {
	s.RequestId = &v
	return s
}

func (s *ModifyDesktopsPolicyGroupResponse) SetModifyResults(v []*ModifyDesktopsPolicyGroupResponseModifyResults) *ModifyDesktopsPolicyGroupResponse {
	s.ModifyResults = v
	return s
}

type ModifyDesktopsPolicyGroupResponseModifyResults struct {
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
}

func (s ModifyDesktopsPolicyGroupResponseModifyResults) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopsPolicyGroupResponseModifyResults) GoString() string {
	return s.String()
}

func (s *ModifyDesktopsPolicyGroupResponseModifyResults) SetDesktopId(v string) *ModifyDesktopsPolicyGroupResponseModifyResults {
	s.DesktopId = &v
	return s
}

func (s *ModifyDesktopsPolicyGroupResponseModifyResults) SetCode(v string) *ModifyDesktopsPolicyGroupResponseModifyResults {
	s.Code = &v
	return s
}

func (s *ModifyDesktopsPolicyGroupResponseModifyResults) SetMessage(v string) *ModifyDesktopsPolicyGroupResponseModifyResults {
	s.Message = &v
	return s
}

type CreatePolicyGroupRequest struct {
	RegionId                    *string                                                `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Clipboard                   *string                                                `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	LocalDrive                  *string                                                `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	UsbRedirect                 *string                                                `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	VisualQuality               *string                                                `json:"VisualQuality,omitempty" xml:"VisualQuality,omitempty"`
	Html5Access                 *string                                                `json:"Html5Access,omitempty" xml:"Html5Access,omitempty"`
	Html5FileTransfer           *string                                                `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	Watermark                   *string                                                `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	Name                        *string                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	WatermarkType               *string                                                `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
	WatermarkCustomText         *string                                                `json:"WatermarkCustomText,omitempty" xml:"WatermarkCustomText,omitempty"`
	WatermarkTransparency       *string                                                `json:"WatermarkTransparency,omitempty" xml:"WatermarkTransparency,omitempty"`
	PreemptLogin                *string                                                `json:"PreemptLogin,omitempty" xml:"PreemptLogin,omitempty"`
	PreemptLoginUser            []*string                                              `json:"PreemptLoginUser,omitempty" xml:"PreemptLoginUser,omitempty" type:"Repeated"`
	DomainList                  *string                                                `json:"DomainList,omitempty" xml:"DomainList,omitempty"`
	AuthorizeSecurityPolicyRule []*CreatePolicyGroupRequestAuthorizeSecurityPolicyRule `json:"AuthorizeSecurityPolicyRule,omitempty" xml:"AuthorizeSecurityPolicyRule,omitempty" type:"Repeated"`
	AuthorizeAccessPolicyRule   []*CreatePolicyGroupRequestAuthorizeAccessPolicyRule   `json:"AuthorizeAccessPolicyRule,omitempty" xml:"AuthorizeAccessPolicyRule,omitempty" type:"Repeated"`
}

func (s CreatePolicyGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyGroupRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupRequest) SetRegionId(v string) *CreatePolicyGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetClipboard(v string) *CreatePolicyGroupRequest {
	s.Clipboard = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetLocalDrive(v string) *CreatePolicyGroupRequest {
	s.LocalDrive = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetUsbRedirect(v string) *CreatePolicyGroupRequest {
	s.UsbRedirect = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetVisualQuality(v string) *CreatePolicyGroupRequest {
	s.VisualQuality = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetHtml5Access(v string) *CreatePolicyGroupRequest {
	s.Html5Access = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetHtml5FileTransfer(v string) *CreatePolicyGroupRequest {
	s.Html5FileTransfer = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWatermark(v string) *CreatePolicyGroupRequest {
	s.Watermark = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetName(v string) *CreatePolicyGroupRequest {
	s.Name = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWatermarkType(v string) *CreatePolicyGroupRequest {
	s.WatermarkType = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWatermarkCustomText(v string) *CreatePolicyGroupRequest {
	s.WatermarkCustomText = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWatermarkTransparency(v string) *CreatePolicyGroupRequest {
	s.WatermarkTransparency = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetPreemptLogin(v string) *CreatePolicyGroupRequest {
	s.PreemptLogin = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetPreemptLoginUser(v []*string) *CreatePolicyGroupRequest {
	s.PreemptLoginUser = v
	return s
}

func (s *CreatePolicyGroupRequest) SetDomainList(v string) *CreatePolicyGroupRequest {
	s.DomainList = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetAuthorizeSecurityPolicyRule(v []*CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) *CreatePolicyGroupRequest {
	s.AuthorizeSecurityPolicyRule = v
	return s
}

func (s *CreatePolicyGroupRequest) SetAuthorizeAccessPolicyRule(v []*CreatePolicyGroupRequestAuthorizeAccessPolicyRule) *CreatePolicyGroupRequest {
	s.AuthorizeAccessPolicyRule = v
	return s
}

type CreatePolicyGroupRequestAuthorizeSecurityPolicyRule struct {
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	IpProtocol  *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	PortRange   *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	CidrIp      *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	Policy      *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Priority    *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) SetType(v string) *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Type = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) SetIpProtocol(v string) *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.IpProtocol = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) SetPortRange(v string) *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.PortRange = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) SetCidrIp(v string) *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.CidrIp = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) SetPolicy(v string) *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Policy = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) SetPriority(v string) *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Priority = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) SetDescription(v string) *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Description = &v
	return s
}

type CreatePolicyGroupRequestAuthorizeAccessPolicyRule struct {
	CidrIp      *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreatePolicyGroupRequestAuthorizeAccessPolicyRule) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyGroupRequestAuthorizeAccessPolicyRule) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupRequestAuthorizeAccessPolicyRule) SetCidrIp(v string) *CreatePolicyGroupRequestAuthorizeAccessPolicyRule {
	s.CidrIp = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeAccessPolicyRule) SetDescription(v string) *CreatePolicyGroupRequestAuthorizeAccessPolicyRule {
	s.Description = &v
	return s
}

type CreatePolicyGroupResponse struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty" require:"true"`
}

func (s CreatePolicyGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyGroupResponse) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupResponse) SetRequestId(v string) *CreatePolicyGroupResponse {
	s.RequestId = &v
	return s
}

func (s *CreatePolicyGroupResponse) SetPolicyGroupId(v string) *CreatePolicyGroupResponse {
	s.PolicyGroupId = &v
	return s
}

type CreateADConnectorDirectoryRequest struct {
	RegionId            *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DomainName          *string   `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	DomainUserName      *string   `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty" require:"true"`
	DomainPassword      *string   `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty" require:"true"`
	DnsAddress          []*string `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" require:"true" type:"Repeated"`
	VSwitchId           []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" require:"true" type:"Repeated"`
	DirectoryName       *string   `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	EnableAdminAccess   *bool     `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	DesktopAccessType   *string   `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	SubDomainDnsAddress []*string `json:"SubDomainDnsAddress,omitempty" xml:"SubDomainDnsAddress,omitempty" type:"Repeated"`
	SubDomainName       *string   `json:"SubDomainName,omitempty" xml:"SubDomainName,omitempty"`
	MfaEnabled          *bool     `json:"MfaEnabled,omitempty" xml:"MfaEnabled,omitempty"`
}

func (s CreateADConnectorDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateADConnectorDirectoryRequest) GoString() string {
	return s.String()
}

func (s *CreateADConnectorDirectoryRequest) SetRegionId(v string) *CreateADConnectorDirectoryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetDomainName(v string) *CreateADConnectorDirectoryRequest {
	s.DomainName = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetDomainUserName(v string) *CreateADConnectorDirectoryRequest {
	s.DomainUserName = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetDomainPassword(v string) *CreateADConnectorDirectoryRequest {
	s.DomainPassword = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetDnsAddress(v []*string) *CreateADConnectorDirectoryRequest {
	s.DnsAddress = v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetVSwitchId(v []*string) *CreateADConnectorDirectoryRequest {
	s.VSwitchId = v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetDirectoryName(v string) *CreateADConnectorDirectoryRequest {
	s.DirectoryName = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetEnableAdminAccess(v bool) *CreateADConnectorDirectoryRequest {
	s.EnableAdminAccess = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetDesktopAccessType(v string) *CreateADConnectorDirectoryRequest {
	s.DesktopAccessType = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetSubDomainDnsAddress(v []*string) *CreateADConnectorDirectoryRequest {
	s.SubDomainDnsAddress = v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetSubDomainName(v string) *CreateADConnectorDirectoryRequest {
	s.SubDomainName = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetMfaEnabled(v bool) *CreateADConnectorDirectoryRequest {
	s.MfaEnabled = &v
	return s
}

type CreateADConnectorDirectoryResponse struct {
	DirectoryId   *string                                           `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" require:"true"`
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TrustPassword *string                                           `json:"TrustPassword,omitempty" xml:"TrustPassword,omitempty" require:"true"`
	AdConnectors  []*CreateADConnectorDirectoryResponseAdConnectors `json:"AdConnectors,omitempty" xml:"AdConnectors,omitempty" require:"true" type:"Repeated"`
}

func (s CreateADConnectorDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateADConnectorDirectoryResponse) GoString() string {
	return s.String()
}

func (s *CreateADConnectorDirectoryResponse) SetDirectoryId(v string) *CreateADConnectorDirectoryResponse {
	s.DirectoryId = &v
	return s
}

func (s *CreateADConnectorDirectoryResponse) SetRequestId(v string) *CreateADConnectorDirectoryResponse {
	s.RequestId = &v
	return s
}

func (s *CreateADConnectorDirectoryResponse) SetTrustPassword(v string) *CreateADConnectorDirectoryResponse {
	s.TrustPassword = &v
	return s
}

func (s *CreateADConnectorDirectoryResponse) SetAdConnectors(v []*CreateADConnectorDirectoryResponseAdConnectors) *CreateADConnectorDirectoryResponse {
	s.AdConnectors = v
	return s
}

type CreateADConnectorDirectoryResponseAdConnectors struct {
	Address *string `json:"Address,omitempty" xml:"Address,omitempty" require:"true"`
}

func (s CreateADConnectorDirectoryResponseAdConnectors) String() string {
	return tea.Prettify(s)
}

func (s CreateADConnectorDirectoryResponseAdConnectors) GoString() string {
	return s.String()
}

func (s *CreateADConnectorDirectoryResponseAdConnectors) SetAddress(v string) *CreateADConnectorDirectoryResponseAdConnectors {
	s.Address = &v
	return s
}

type DescribeBundlesRequest struct {
	RegionId          *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	MaxResults        *int      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken         *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	BundleId          []*string `json:"BundleId,omitempty" xml:"BundleId,omitempty" type:"Repeated"`
	BundleType        *string   `json:"BundleType,omitempty" xml:"BundleType,omitempty"`
	DesktopTypeFamily *string   `json:"DesktopTypeFamily,omitempty" xml:"DesktopTypeFamily,omitempty"`
	CpuCount          *int      `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	MemorySize        *int      `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	GpuCount          *float32  `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	CheckStock        *bool     `json:"CheckStock,omitempty" xml:"CheckStock,omitempty"`
	FromDesktopGroup  *bool     `json:"FromDesktopGroup,omitempty" xml:"FromDesktopGroup,omitempty"`
	ProtocolType      *string   `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
}

func (s DescribeBundlesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBundlesRequest) GoString() string {
	return s.String()
}

func (s *DescribeBundlesRequest) SetRegionId(v string) *DescribeBundlesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBundlesRequest) SetMaxResults(v int) *DescribeBundlesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeBundlesRequest) SetNextToken(v string) *DescribeBundlesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeBundlesRequest) SetBundleId(v []*string) *DescribeBundlesRequest {
	s.BundleId = v
	return s
}

func (s *DescribeBundlesRequest) SetBundleType(v string) *DescribeBundlesRequest {
	s.BundleType = &v
	return s
}

func (s *DescribeBundlesRequest) SetDesktopTypeFamily(v string) *DescribeBundlesRequest {
	s.DesktopTypeFamily = &v
	return s
}

func (s *DescribeBundlesRequest) SetCpuCount(v int) *DescribeBundlesRequest {
	s.CpuCount = &v
	return s
}

func (s *DescribeBundlesRequest) SetMemorySize(v int) *DescribeBundlesRequest {
	s.MemorySize = &v
	return s
}

func (s *DescribeBundlesRequest) SetGpuCount(v float32) *DescribeBundlesRequest {
	s.GpuCount = &v
	return s
}

func (s *DescribeBundlesRequest) SetCheckStock(v bool) *DescribeBundlesRequest {
	s.CheckStock = &v
	return s
}

func (s *DescribeBundlesRequest) SetFromDesktopGroup(v bool) *DescribeBundlesRequest {
	s.FromDesktopGroup = &v
	return s
}

func (s *DescribeBundlesRequest) SetProtocolType(v string) *DescribeBundlesRequest {
	s.ProtocolType = &v
	return s
}

type DescribeBundlesResponse struct {
	NextToken *string                           `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Bundles   []*DescribeBundlesResponseBundles `json:"Bundles,omitempty" xml:"Bundles,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeBundlesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBundlesResponse) GoString() string {
	return s.String()
}

func (s *DescribeBundlesResponse) SetNextToken(v string) *DescribeBundlesResponse {
	s.NextToken = &v
	return s
}

func (s *DescribeBundlesResponse) SetRequestId(v string) *DescribeBundlesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeBundlesResponse) SetBundles(v []*DescribeBundlesResponseBundles) *DescribeBundlesResponse {
	s.Bundles = v
	return s
}

type DescribeBundlesResponseBundles struct {
	ImageId              *string                                             `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	BundleId             *string                                             `json:"BundleId,omitempty" xml:"BundleId,omitempty" require:"true"`
	BundleType           *string                                             `json:"BundleType,omitempty" xml:"BundleType,omitempty" require:"true"`
	BundleName           *string                                             `json:"BundleName,omitempty" xml:"BundleName,omitempty" require:"true"`
	Description          *string                                             `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	DesktopType          *string                                             `json:"DesktopType,omitempty" xml:"DesktopType,omitempty" require:"true"`
	OsType               *string                                             `json:"OsType,omitempty" xml:"OsType,omitempty" require:"true"`
	StockState           *string                                             `json:"StockState,omitempty" xml:"StockState,omitempty" require:"true"`
	ProtocolType         *string                                             `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty" require:"true"`
	Disks                []*DescribeBundlesResponseBundlesDisks              `json:"Disks,omitempty" xml:"Disks,omitempty" require:"true" type:"Repeated"`
	DesktopTypeAttribute *DescribeBundlesResponseBundlesDesktopTypeAttribute `json:"DesktopTypeAttribute,omitempty" xml:"DesktopTypeAttribute,omitempty" require:"true" type:"Struct"`
}

func (s DescribeBundlesResponseBundles) String() string {
	return tea.Prettify(s)
}

func (s DescribeBundlesResponseBundles) GoString() string {
	return s.String()
}

func (s *DescribeBundlesResponseBundles) SetImageId(v string) *DescribeBundlesResponseBundles {
	s.ImageId = &v
	return s
}

func (s *DescribeBundlesResponseBundles) SetBundleId(v string) *DescribeBundlesResponseBundles {
	s.BundleId = &v
	return s
}

func (s *DescribeBundlesResponseBundles) SetBundleType(v string) *DescribeBundlesResponseBundles {
	s.BundleType = &v
	return s
}

func (s *DescribeBundlesResponseBundles) SetBundleName(v string) *DescribeBundlesResponseBundles {
	s.BundleName = &v
	return s
}

func (s *DescribeBundlesResponseBundles) SetDescription(v string) *DescribeBundlesResponseBundles {
	s.Description = &v
	return s
}

func (s *DescribeBundlesResponseBundles) SetDesktopType(v string) *DescribeBundlesResponseBundles {
	s.DesktopType = &v
	return s
}

func (s *DescribeBundlesResponseBundles) SetOsType(v string) *DescribeBundlesResponseBundles {
	s.OsType = &v
	return s
}

func (s *DescribeBundlesResponseBundles) SetStockState(v string) *DescribeBundlesResponseBundles {
	s.StockState = &v
	return s
}

func (s *DescribeBundlesResponseBundles) SetProtocolType(v string) *DescribeBundlesResponseBundles {
	s.ProtocolType = &v
	return s
}

func (s *DescribeBundlesResponseBundles) SetDisks(v []*DescribeBundlesResponseBundlesDisks) *DescribeBundlesResponseBundles {
	s.Disks = v
	return s
}

func (s *DescribeBundlesResponseBundles) SetDesktopTypeAttribute(v *DescribeBundlesResponseBundlesDesktopTypeAttribute) *DescribeBundlesResponseBundles {
	s.DesktopTypeAttribute = v
	return s
}

type DescribeBundlesResponseBundlesDisks struct {
	DiskSize *int    `json:"DiskSize,omitempty" xml:"DiskSize,omitempty" require:"true"`
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty" require:"true"`
}

func (s DescribeBundlesResponseBundlesDisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeBundlesResponseBundlesDisks) GoString() string {
	return s.String()
}

func (s *DescribeBundlesResponseBundlesDisks) SetDiskSize(v int) *DescribeBundlesResponseBundlesDisks {
	s.DiskSize = &v
	return s
}

func (s *DescribeBundlesResponseBundlesDisks) SetDiskType(v string) *DescribeBundlesResponseBundlesDisks {
	s.DiskType = &v
	return s
}

type DescribeBundlesResponseBundlesDesktopTypeAttribute struct {
	CpuCount   *int     `json:"CpuCount,omitempty" xml:"CpuCount,omitempty" require:"true"`
	MemorySize *int     `json:"MemorySize,omitempty" xml:"MemorySize,omitempty" require:"true"`
	GpuCount   *float32 `json:"GpuCount,omitempty" xml:"GpuCount,omitempty" require:"true"`
	GpuSpec    *string  `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty" require:"true"`
}

func (s DescribeBundlesResponseBundlesDesktopTypeAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeBundlesResponseBundlesDesktopTypeAttribute) GoString() string {
	return s.String()
}

func (s *DescribeBundlesResponseBundlesDesktopTypeAttribute) SetCpuCount(v int) *DescribeBundlesResponseBundlesDesktopTypeAttribute {
	s.CpuCount = &v
	return s
}

func (s *DescribeBundlesResponseBundlesDesktopTypeAttribute) SetMemorySize(v int) *DescribeBundlesResponseBundlesDesktopTypeAttribute {
	s.MemorySize = &v
	return s
}

func (s *DescribeBundlesResponseBundlesDesktopTypeAttribute) SetGpuCount(v float32) *DescribeBundlesResponseBundlesDesktopTypeAttribute {
	s.GpuCount = &v
	return s
}

func (s *DescribeBundlesResponseBundlesDesktopTypeAttribute) SetGpuSpec(v string) *DescribeBundlesResponseBundlesDesktopTypeAttribute {
	s.GpuSpec = &v
	return s
}

type DeleteImagesRequest struct {
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ImageId  []*string `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteImagesRequest) GoString() string {
	return s.String()
}

func (s *DeleteImagesRequest) SetRegionId(v string) *DeleteImagesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteImagesRequest) SetImageId(v []*string) *DeleteImagesRequest {
	s.ImageId = v
	return s
}

type DeleteImagesResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteImagesResponse) GoString() string {
	return s.String()
}

func (s *DeleteImagesResponse) SetRequestId(v string) *DeleteImagesResponse {
	s.RequestId = &v
	return s
}

type CreateDesktopsRequest struct {
	RegionId       *string                     `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	GroupId        *string                     `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	BundleId       *string                     `json:"BundleId,omitempty" xml:"BundleId,omitempty" require:"true"`
	DesktopName    *string                     `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	UserName       *string                     `json:"UserName,omitempty" xml:"UserName,omitempty"`
	VpcId          *string                     `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Amount         *int                        `json:"Amount,omitempty" xml:"Amount,omitempty"`
	DirectoryId    *string                     `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	OfficeSiteId   *string                     `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	EndUserId      []*string                   `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" require:"true" type:"Repeated"`
	PolicyGroupId  *string                     `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty" require:"true"`
	ChargeType     *string                     `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Period         *int                        `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit     *string                     `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	AutoPay        *bool                       `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	Tag            []*CreateDesktopsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	AutoRenew      *bool                       `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	PromotionId    *string                     `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	UserAssignMode *string                     `json:"UserAssignMode,omitempty" xml:"UserAssignMode,omitempty"`
}

func (s CreateDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopsRequest) GoString() string {
	return s.String()
}

func (s *CreateDesktopsRequest) SetRegionId(v string) *CreateDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDesktopsRequest) SetGroupId(v string) *CreateDesktopsRequest {
	s.GroupId = &v
	return s
}

func (s *CreateDesktopsRequest) SetBundleId(v string) *CreateDesktopsRequest {
	s.BundleId = &v
	return s
}

func (s *CreateDesktopsRequest) SetDesktopName(v string) *CreateDesktopsRequest {
	s.DesktopName = &v
	return s
}

func (s *CreateDesktopsRequest) SetUserName(v string) *CreateDesktopsRequest {
	s.UserName = &v
	return s
}

func (s *CreateDesktopsRequest) SetVpcId(v string) *CreateDesktopsRequest {
	s.VpcId = &v
	return s
}

func (s *CreateDesktopsRequest) SetAmount(v int) *CreateDesktopsRequest {
	s.Amount = &v
	return s
}

func (s *CreateDesktopsRequest) SetDirectoryId(v string) *CreateDesktopsRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateDesktopsRequest) SetOfficeSiteId(v string) *CreateDesktopsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateDesktopsRequest) SetEndUserId(v []*string) *CreateDesktopsRequest {
	s.EndUserId = v
	return s
}

func (s *CreateDesktopsRequest) SetPolicyGroupId(v string) *CreateDesktopsRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *CreateDesktopsRequest) SetChargeType(v string) *CreateDesktopsRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateDesktopsRequest) SetPeriod(v int) *CreateDesktopsRequest {
	s.Period = &v
	return s
}

func (s *CreateDesktopsRequest) SetPeriodUnit(v string) *CreateDesktopsRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateDesktopsRequest) SetAutoPay(v bool) *CreateDesktopsRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateDesktopsRequest) SetTag(v []*CreateDesktopsRequestTag) *CreateDesktopsRequest {
	s.Tag = v
	return s
}

func (s *CreateDesktopsRequest) SetAutoRenew(v bool) *CreateDesktopsRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateDesktopsRequest) SetPromotionId(v string) *CreateDesktopsRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateDesktopsRequest) SetUserAssignMode(v string) *CreateDesktopsRequest {
	s.UserAssignMode = &v
	return s
}

type CreateDesktopsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDesktopsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopsRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDesktopsRequestTag) SetKey(v string) *CreateDesktopsRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDesktopsRequestTag) SetValue(v string) *CreateDesktopsRequestTag {
	s.Value = &v
	return s
}

type CreateDesktopsResponse struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	OrderId   *string   `json:"OrderId,omitempty" xml:"OrderId,omitempty" require:"true"`
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true" type:"Repeated"`
}

func (s CreateDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopsResponse) GoString() string {
	return s.String()
}

func (s *CreateDesktopsResponse) SetRequestId(v string) *CreateDesktopsResponse {
	s.RequestId = &v
	return s
}

func (s *CreateDesktopsResponse) SetOrderId(v string) *CreateDesktopsResponse {
	s.OrderId = &v
	return s
}

func (s *CreateDesktopsResponse) SetDesktopId(v []*string) *CreateDesktopsResponse {
	s.DesktopId = v
	return s
}

type DescribeImagesRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	MaxResults   *int      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ImageType    *string   `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	ImageStatus  *string   `json:"ImageStatus,omitempty" xml:"ImageStatus,omitempty"`
	ImageId      []*string `json:"ImageId,omitempty" xml:"ImageId,omitempty" type:"Repeated"`
	GpuCategory  *bool     `json:"GpuCategory,omitempty" xml:"GpuCategory,omitempty"`
	ProtocolType *string   `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
}

func (s DescribeImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeImagesRequest) SetRegionId(v string) *DescribeImagesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeImagesRequest) SetMaxResults(v int) *DescribeImagesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeImagesRequest) SetNextToken(v string) *DescribeImagesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeImagesRequest) SetImageType(v string) *DescribeImagesRequest {
	s.ImageType = &v
	return s
}

func (s *DescribeImagesRequest) SetImageStatus(v string) *DescribeImagesRequest {
	s.ImageStatus = &v
	return s
}

func (s *DescribeImagesRequest) SetImageId(v []*string) *DescribeImagesRequest {
	s.ImageId = v
	return s
}

func (s *DescribeImagesRequest) SetGpuCategory(v bool) *DescribeImagesRequest {
	s.GpuCategory = &v
	return s
}

func (s *DescribeImagesRequest) SetProtocolType(v string) *DescribeImagesRequest {
	s.ProtocolType = &v
	return s
}

type DescribeImagesResponse struct {
	NextToken *string                         `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Images    []*DescribeImagesResponseImages `json:"Images,omitempty" xml:"Images,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponse) SetNextToken(v string) *DescribeImagesResponse {
	s.NextToken = &v
	return s
}

func (s *DescribeImagesResponse) SetRequestId(v string) *DescribeImagesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeImagesResponse) SetImages(v []*DescribeImagesResponseImages) *DescribeImagesResponse {
	s.Images = v
	return s
}

type DescribeImagesResponseImages struct {
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty" require:"true"`
	ImageId      *string `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	ImageType    *string `json:"ImageType,omitempty" xml:"ImageType,omitempty" require:"true"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty" require:"true"`
	Size         *int    `json:"Size,omitempty" xml:"Size,omitempty" require:"true"`
	DataDiskSize *int    `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty" require:"true"`
	GpuCategory  *bool   `json:"GpuCategory,omitempty" xml:"GpuCategory,omitempty" require:"true"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	OsType       *string `json:"OsType,omitempty" xml:"OsType,omitempty" require:"true"`
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty" require:"true"`
}

func (s DescribeImagesResponseImages) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesResponseImages) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseImages) SetCreationTime(v string) *DescribeImagesResponseImages {
	s.CreationTime = &v
	return s
}

func (s *DescribeImagesResponseImages) SetImageId(v string) *DescribeImagesResponseImages {
	s.ImageId = &v
	return s
}

func (s *DescribeImagesResponseImages) SetImageType(v string) *DescribeImagesResponseImages {
	s.ImageType = &v
	return s
}

func (s *DescribeImagesResponseImages) SetName(v string) *DescribeImagesResponseImages {
	s.Name = &v
	return s
}

func (s *DescribeImagesResponseImages) SetProgress(v string) *DescribeImagesResponseImages {
	s.Progress = &v
	return s
}

func (s *DescribeImagesResponseImages) SetSize(v int) *DescribeImagesResponseImages {
	s.Size = &v
	return s
}

func (s *DescribeImagesResponseImages) SetDataDiskSize(v int) *DescribeImagesResponseImages {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeImagesResponseImages) SetGpuCategory(v bool) *DescribeImagesResponseImages {
	s.GpuCategory = &v
	return s
}

func (s *DescribeImagesResponseImages) SetStatus(v string) *DescribeImagesResponseImages {
	s.Status = &v
	return s
}

func (s *DescribeImagesResponseImages) SetDescription(v string) *DescribeImagesResponseImages {
	s.Description = &v
	return s
}

func (s *DescribeImagesResponseImages) SetOsType(v string) *DescribeImagesResponseImages {
	s.OsType = &v
	return s
}

func (s *DescribeImagesResponseImages) SetProtocolType(v string) *DescribeImagesResponseImages {
	s.ProtocolType = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ecd"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) ClonePolicyGroupWithOptions(request *ClonePolicyGroupRequest, runtime *util.RuntimeOptions) (_result *ClonePolicyGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ClonePolicyGroupResponse{}
	_body, _err := client.DoRequest(tea.String("ClonePolicyGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClonePolicyGroup(request *ClonePolicyGroupRequest) (_result *ClonePolicyGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ClonePolicyGroupResponse{}
	_body, _err := client.ClonePolicyGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDesktopsInGroupWithOptions(request *DescribeDesktopsInGroupRequest, runtime *util.RuntimeOptions) (_result *DescribeDesktopsInGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDesktopsInGroupResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDesktopsInGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDesktopsInGroup(request *DescribeDesktopsInGroupRequest) (_result *DescribeDesktopsInGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDesktopsInGroupResponse{}
	_body, _err := client.DescribeDesktopsInGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUsersInGroupWithOptions(request *DescribeUsersInGroupRequest, runtime *util.RuntimeOptions) (_result *DescribeUsersInGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeUsersInGroupResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeUsersInGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUsersInGroup(request *DescribeUsersInGroupRequest) (_result *DescribeUsersInGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUsersInGroupResponse{}
	_body, _err := client.DescribeUsersInGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDesktopGroupDetailWithOptions(request *GetDesktopGroupDetailRequest, runtime *util.RuntimeOptions) (_result *GetDesktopGroupDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetDesktopGroupDetailResponse{}
	_body, _err := client.DoRequest(tea.String("GetDesktopGroupDetail"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDesktopGroupDetail(request *GetDesktopGroupDetailRequest) (_result *GetDesktopGroupDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDesktopGroupDetailResponse{}
	_body, _err := client.GetDesktopGroupDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDesktopGroupWithOptions(request *ModifyDesktopGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyDesktopGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyDesktopGroupResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyDesktopGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDesktopGroup(request *ModifyDesktopGroupRequest) (_result *ModifyDesktopGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDesktopGroupResponse{}
	_body, _err := client.ModifyDesktopGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveUserFromDesktopGroupWithOptions(request *RemoveUserFromDesktopGroupRequest, runtime *util.RuntimeOptions) (_result *RemoveUserFromDesktopGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RemoveUserFromDesktopGroupResponse{}
	_body, _err := client.DoRequest(tea.String("RemoveUserFromDesktopGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveUserFromDesktopGroup(request *RemoveUserFromDesktopGroupRequest) (_result *RemoveUserFromDesktopGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveUserFromDesktopGroupResponse{}
	_body, _err := client.RemoveUserFromDesktopGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteScaleStrategyWithOptions(request *DeleteScaleStrategyRequest, runtime *util.RuntimeOptions) (_result *DeleteScaleStrategyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteScaleStrategyResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteScaleStrategy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteScaleStrategy(request *DeleteScaleStrategyRequest) (_result *DeleteScaleStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteScaleStrategyResponse{}
	_body, _err := client.DeleteScaleStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDesktopGroupWithOptions(request *DeleteDesktopGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteDesktopGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteDesktopGroupResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteDesktopGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDesktopGroup(request *DeleteDesktopGroupRequest) (_result *DeleteDesktopGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDesktopGroupResponse{}
	_body, _err := client.DeleteDesktopGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDesktopGroupsWithOptions(request *DescribeDesktopGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeDesktopGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDesktopGroupsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDesktopGroups"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDesktopGroups(request *DescribeDesktopGroupsRequest) (_result *DescribeDesktopGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDesktopGroupsResponse{}
	_body, _err := client.DescribeDesktopGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserConnectionRecordsWithOptions(request *DescribeUserConnectionRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeUserConnectionRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeUserConnectionRecordsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeUserConnectionRecords"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserConnectionRecords(request *DescribeUserConnectionRecordsRequest) (_result *DescribeUserConnectionRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserConnectionRecordsResponse{}
	_body, _err := client.DescribeUserConnectionRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyScaleStrategyWithOptions(request *ModifyScaleStrategyRequest, runtime *util.RuntimeOptions) (_result *ModifyScaleStrategyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyScaleStrategyResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyScaleStrategy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyScaleStrategy(request *ModifyScaleStrategyRequest) (_result *ModifyScaleStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyScaleStrategyResponse{}
	_body, _err := client.ModifyScaleStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScaleStrategysWithOptions(request *DescribeScaleStrategysRequest, runtime *util.RuntimeOptions) (_result *DescribeScaleStrategysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeScaleStrategysResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeScaleStrategys"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScaleStrategys(request *DescribeScaleStrategysRequest) (_result *DescribeScaleStrategysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScaleStrategysResponse{}
	_body, _err := client.DescribeScaleStrategysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddUserToDesktopGroupWithOptions(request *AddUserToDesktopGroupRequest, runtime *util.RuntimeOptions) (_result *AddUserToDesktopGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddUserToDesktopGroupResponse{}
	_body, _err := client.DoRequest(tea.String("AddUserToDesktopGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddUserToDesktopGroup(request *AddUserToDesktopGroupRequest) (_result *AddUserToDesktopGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUserToDesktopGroupResponse{}
	_body, _err := client.AddUserToDesktopGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePostPaidDesktopBillsWithOptions(request *DescribePostPaidDesktopBillsRequest, runtime *util.RuntimeOptions) (_result *DescribePostPaidDesktopBillsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribePostPaidDesktopBillsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribePostPaidDesktopBills"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePostPaidDesktopBills(request *DescribePostPaidDesktopBillsRequest) (_result *DescribePostPaidDesktopBillsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePostPaidDesktopBillsResponse{}
	_body, _err := client.DescribePostPaidDesktopBillsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDesktopGroupWithOptions(request *CreateDesktopGroupRequest, runtime *util.RuntimeOptions) (_result *CreateDesktopGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateDesktopGroupResponse{}
	_body, _err := client.DoRequest(tea.String("CreateDesktopGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDesktopGroup(request *CreateDesktopGroupRequest) (_result *CreateDesktopGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDesktopGroupResponse{}
	_body, _err := client.CreateDesktopGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecreateDesktopGroupWithOptions(request *RecreateDesktopGroupRequest, runtime *util.RuntimeOptions) (_result *RecreateDesktopGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RecreateDesktopGroupResponse{}
	_body, _err := client.DoRequest(tea.String("RecreateDesktopGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecreateDesktopGroup(request *RecreateDesktopGroupRequest) (_result *RecreateDesktopGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecreateDesktopGroupResponse{}
	_body, _err := client.RecreateDesktopGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateScaleStrategyWithOptions(request *CreateScaleStrategyRequest, runtime *util.RuntimeOptions) (_result *CreateScaleStrategyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateScaleStrategyResponse{}
	_body, _err := client.DoRequest(tea.String("CreateScaleStrategy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateScaleStrategy(request *CreateScaleStrategyRequest) (_result *CreateScaleStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateScaleStrategyResponse{}
	_body, _err := client.CreateScaleStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyUserToDesktopGroupWithOptions(request *ModifyUserToDesktopGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyUserToDesktopGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyUserToDesktopGroupResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyUserToDesktopGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyUserToDesktopGroup(request *ModifyUserToDesktopGroupRequest) (_result *ModifyUserToDesktopGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyUserToDesktopGroupResponse{}
	_body, _err := client.ModifyUserToDesktopGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDesktopsLiteWithOptions(request *CreateDesktopsLiteRequest, runtime *util.RuntimeOptions) (_result *CreateDesktopsLiteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateDesktopsLiteResponse{}
	_body, _err := client.DoRequest(tea.String("CreateDesktopsLite"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDesktopsLite(request *CreateDesktopsLiteRequest) (_result *CreateDesktopsLiteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDesktopsLiteResponse{}
	_body, _err := client.CreateDesktopsLiteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyOfficeSiteAttributeWithOptions(request *ModifyOfficeSiteAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyOfficeSiteAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyOfficeSiteAttributeResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyOfficeSiteAttribute"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyOfficeSiteAttribute(request *ModifyOfficeSiteAttributeRequest) (_result *ModifyOfficeSiteAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyOfficeSiteAttributeResponse{}
	_body, _err := client.ModifyOfficeSiteAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServiceLinkedRoleWithOptions(request *CreateServiceLinkedRoleRequest, runtime *util.RuntimeOptions) (_result *CreateServiceLinkedRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.DoRequest(tea.String("CreateServiceLinkedRole"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateServiceLinkedRole(request *CreateServiceLinkedRoleRequest) (_result *CreateServiceLinkedRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.CreateServiceLinkedRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyOfficeSiteCrossDesktopAccessWithOptions(request *ModifyOfficeSiteCrossDesktopAccessRequest, runtime *util.RuntimeOptions) (_result *ModifyOfficeSiteCrossDesktopAccessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyOfficeSiteCrossDesktopAccessResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyOfficeSiteCrossDesktopAccess"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyOfficeSiteCrossDesktopAccess(request *ModifyOfficeSiteCrossDesktopAccessRequest) (_result *ModifyOfficeSiteCrossDesktopAccessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyOfficeSiteCrossDesktopAccessResponse{}
	_body, _err := client.ModifyOfficeSiteCrossDesktopAccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDesktopUsersWithOptions(request *GetDesktopUsersRequest, runtime *util.RuntimeOptions) (_result *GetDesktopUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetDesktopUsersResponse{}
	_body, _err := client.DoRequest(tea.String("GetDesktopUsers"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDesktopUsers(request *GetDesktopUsersRequest) (_result *GetDesktopUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDesktopUsersResponse{}
	_body, _err := client.GetDesktopUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyNetworkPackageEnabledWithOptions(request *ModifyNetworkPackageEnabledRequest, runtime *util.RuntimeOptions) (_result *ModifyNetworkPackageEnabledResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyNetworkPackageEnabledResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyNetworkPackageEnabled"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyNetworkPackageEnabled(request *ModifyNetworkPackageEnabledRequest) (_result *ModifyNetworkPackageEnabledResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyNetworkPackageEnabledResponse{}
	_body, _err := client.ModifyNetworkPackageEnabledWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetNASDefaultMountTargetWithOptions(request *ResetNASDefaultMountTargetRequest, runtime *util.RuntimeOptions) (_result *ResetNASDefaultMountTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ResetNASDefaultMountTargetResponse{}
	_body, _err := client.DoRequest(tea.String("ResetNASDefaultMountTarget"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetNASDefaultMountTarget(request *ResetNASDefaultMountTargetRequest) (_result *ResetNASDefaultMountTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetNASDefaultMountTargetResponse{}
	_body, _err := client.ResetNASDefaultMountTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCensWithOptions(request *DescribeCensRequest, runtime *util.RuntimeOptions) (_result *DescribeCensResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeCensResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeCens"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCens(request *DescribeCensRequest) (_result *DescribeCensResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCensResponse{}
	_body, _err := client.DescribeCensWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckUserInSecurityCenterWhiteListWithOptions(request *CheckUserInSecurityCenterWhiteListRequest, runtime *util.RuntimeOptions) (_result *CheckUserInSecurityCenterWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CheckUserInSecurityCenterWhiteListResponse{}
	_body, _err := client.DoRequest(tea.String("CheckUserInSecurityCenterWhiteList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckUserInSecurityCenterWhiteList(request *CheckUserInSecurityCenterWhiteListRequest) (_result *CheckUserInSecurityCenterWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckUserInSecurityCenterWhiteListResponse{}
	_body, _err := client.CheckUserInSecurityCenterWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddUserToSecurityCenterWhiteListWithOptions(request *AddUserToSecurityCenterWhiteListRequest, runtime *util.RuntimeOptions) (_result *AddUserToSecurityCenterWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddUserToSecurityCenterWhiteListResponse{}
	_body, _err := client.DoRequest(tea.String("AddUserToSecurityCenterWhiteList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddUserToSecurityCenterWhiteList(request *AddUserToSecurityCenterWhiteListRequest) (_result *AddUserToSecurityCenterWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUserToSecurityCenterWhiteListResponse{}
	_body, _err := client.AddUserToSecurityCenterWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVulOverviewWithOptions(request *DescribeVulOverviewRequest, runtime *util.RuntimeOptions) (_result *DescribeVulOverviewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeVulOverviewResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeVulOverview"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVulOverview(request *DescribeVulOverviewRequest) (_result *DescribeVulOverviewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVulOverviewResponse{}
	_body, _err := client.DescribeVulOverviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSuspEventOverviewWithOptions(request *DescribeSuspEventOverviewRequest, runtime *util.RuntimeOptions) (_result *DescribeSuspEventOverviewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeSuspEventOverviewResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeSuspEventOverview"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSuspEventOverview(request *DescribeSuspEventOverviewRequest) (_result *DescribeSuspEventOverviewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSuspEventOverviewResponse{}
	_body, _err := client.DescribeSuspEventOverviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyOfficeSiteMfaEnabledWithOptions(request *ModifyOfficeSiteMfaEnabledRequest, runtime *util.RuntimeOptions) (_result *ModifyOfficeSiteMfaEnabledResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyOfficeSiteMfaEnabledResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyOfficeSiteMfaEnabled"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyOfficeSiteMfaEnabled(request *ModifyOfficeSiteMfaEnabledRequest) (_result *ModifyOfficeSiteMfaEnabledResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyOfficeSiteMfaEnabledResponse{}
	_body, _err := client.ModifyOfficeSiteMfaEnabledWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteNASFileSystemsWithOptions(request *DeleteNASFileSystemsRequest, runtime *util.RuntimeOptions) (_result *DeleteNASFileSystemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteNASFileSystemsResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteNASFileSystems"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteNASFileSystems(request *DeleteNASFileSystemsRequest) (_result *DeleteNASFileSystemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNASFileSystemsResponse{}
	_body, _err := client.DeleteNASFileSystemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyNASDefaultMountTargetWithOptions(request *ModifyNASDefaultMountTargetRequest, runtime *util.RuntimeOptions) (_result *ModifyNASDefaultMountTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyNASDefaultMountTargetResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyNASDefaultMountTarget"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyNASDefaultMountTarget(request *ModifyNASDefaultMountTargetRequest) (_result *ModifyNASDefaultMountTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyNASDefaultMountTargetResponse{}
	_body, _err := client.ModifyNASDefaultMountTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateNASFileSystemWithOptions(request *CreateNASFileSystemRequest, runtime *util.RuntimeOptions) (_result *CreateNASFileSystemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateNASFileSystemResponse{}
	_body, _err := client.DoRequest(tea.String("CreateNASFileSystem"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateNASFileSystem(request *CreateNASFileSystemRequest) (_result *CreateNASFileSystemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNASFileSystemResponse{}
	_body, _err := client.CreateNASFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNASFileSystemsWithOptions(request *DescribeNASFileSystemsRequest, runtime *util.RuntimeOptions) (_result *DescribeNASFileSystemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeNASFileSystemsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeNASFileSystems"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNASFileSystems(request *DescribeNASFileSystemsRequest) (_result *DescribeNASFileSystemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNASFileSystemsResponse{}
	_body, _err := client.DescribeNASFileSystemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartVirusScanTaskWithOptions(request *StartVirusScanTaskRequest, runtime *util.RuntimeOptions) (_result *StartVirusScanTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StartVirusScanTaskResponse{}
	_body, _err := client.DoRequest(tea.String("StartVirusScanTask"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartVirusScanTask(request *StartVirusScanTaskRequest) (_result *StartVirusScanTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartVirusScanTaskResponse{}
	_body, _err := client.StartVirusScanTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyADConnectorOfficeSiteWithOptions(request *ModifyADConnectorOfficeSiteRequest, runtime *util.RuntimeOptions) (_result *ModifyADConnectorOfficeSiteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyADConnectorOfficeSiteResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyADConnectorOfficeSite"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyADConnectorOfficeSite(request *ModifyADConnectorOfficeSiteRequest) (_result *ModifyADConnectorOfficeSiteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyADConnectorOfficeSiteResponse{}
	_body, _err := client.ModifyADConnectorOfficeSiteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFrontVulPatchListWithOptions(request *DescribeFrontVulPatchListRequest, runtime *util.RuntimeOptions) (_result *DescribeFrontVulPatchListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeFrontVulPatchListResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeFrontVulPatchList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFrontVulPatchList(request *DescribeFrontVulPatchListRequest) (_result *DescribeFrontVulPatchListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFrontVulPatchListResponse{}
	_body, _err := client.DescribeFrontVulPatchListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVulDetailsWithOptions(request *DescribeVulDetailsRequest, runtime *util.RuntimeOptions) (_result *DescribeVulDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeVulDetailsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeVulDetails"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVulDetails(request *DescribeVulDetailsRequest) (_result *DescribeVulDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVulDetailsResponse{}
	_body, _err := client.DescribeVulDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSuspEventQuaraFilesWithOptions(request *DescribeSuspEventQuaraFilesRequest, runtime *util.RuntimeOptions) (_result *DescribeSuspEventQuaraFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeSuspEventQuaraFilesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeSuspEventQuaraFiles"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSuspEventQuaraFiles(request *DescribeSuspEventQuaraFilesRequest) (_result *DescribeSuspEventQuaraFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSuspEventQuaraFilesResponse{}
	_body, _err := client.DescribeSuspEventQuaraFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyOperateVulWithOptions(request *ModifyOperateVulRequest, runtime *util.RuntimeOptions) (_result *ModifyOperateVulResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyOperateVulResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyOperateVul"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyOperateVul(request *ModifyOperateVulRequest) (_result *ModifyOperateVulResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyOperateVulResponse{}
	_body, _err := client.ModifyOperateVulWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachCenWithOptions(request *AttachCenRequest, runtime *util.RuntimeOptions) (_result *AttachCenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AttachCenResponse{}
	_body, _err := client.DoRequest(tea.String("AttachCen"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachCen(request *AttachCenRequest) (_result *AttachCenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachCenResponse{}
	_body, _err := client.AttachCenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVulListWithOptions(request *DescribeVulListRequest, runtime *util.RuntimeOptions) (_result *DescribeVulListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeVulListResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeVulList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVulList(request *DescribeVulListRequest) (_result *DescribeVulListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVulListResponse{}
	_body, _err := client.DescribeVulListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOfficeSitesWithOptions(request *DescribeOfficeSitesRequest, runtime *util.RuntimeOptions) (_result *DescribeOfficeSitesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeOfficeSitesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeOfficeSites"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOfficeSites(request *DescribeOfficeSitesRequest) (_result *DescribeOfficeSitesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOfficeSitesResponse{}
	_body, _err := client.DescribeOfficeSitesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSimpleOfficeSiteWithOptions(request *CreateSimpleOfficeSiteRequest, runtime *util.RuntimeOptions) (_result *CreateSimpleOfficeSiteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateSimpleOfficeSiteResponse{}
	_body, _err := client.DoRequest(tea.String("CreateSimpleOfficeSite"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSimpleOfficeSite(request *CreateSimpleOfficeSiteRequest) (_result *CreateSimpleOfficeSiteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSimpleOfficeSiteResponse{}
	_body, _err := client.CreateSimpleOfficeSiteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OperateVulsWithOptions(request *OperateVulsRequest, runtime *util.RuntimeOptions) (_result *OperateVulsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &OperateVulsResponse{}
	_body, _err := client.DoRequest(tea.String("OperateVuls"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OperateVuls(request *OperateVulsRequest) (_result *OperateVulsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OperateVulsResponse{}
	_body, _err := client.OperateVulsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScanTaskProgressWithOptions(request *DescribeScanTaskProgressRequest, runtime *util.RuntimeOptions) (_result *DescribeScanTaskProgressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeScanTaskProgressResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeScanTaskProgress"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScanTaskProgress(request *DescribeScanTaskProgressRequest) (_result *DescribeScanTaskProgressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScanTaskProgressResponse{}
	_body, _err := client.DescribeScanTaskProgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachCenWithOptions(request *DetachCenRequest, runtime *util.RuntimeOptions) (_result *DetachCenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DetachCenResponse{}
	_body, _err := client.DoRequest(tea.String("DetachCen"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachCen(request *DetachCenRequest) (_result *DetachCenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachCenResponse{}
	_body, _err := client.DetachCenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSecurityEventOperationStatusWithOptions(request *DescribeSecurityEventOperationStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeSecurityEventOperationStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeSecurityEventOperationStatusResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeSecurityEventOperationStatus"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSecurityEventOperationStatus(request *DescribeSecurityEventOperationStatusRequest) (_result *DescribeSecurityEventOperationStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSecurityEventOperationStatusResponse{}
	_body, _err := client.DescribeSecurityEventOperationStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAlarmEventStackInfoWithOptions(request *DescribeAlarmEventStackInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeAlarmEventStackInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeAlarmEventStackInfoResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeAlarmEventStackInfo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAlarmEventStackInfo(request *DescribeAlarmEventStackInfoRequest) (_result *DescribeAlarmEventStackInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlarmEventStackInfoResponse{}
	_body, _err := client.DescribeAlarmEventStackInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOfficeSiteUsersWithOptions(request *ListOfficeSiteUsersRequest, runtime *util.RuntimeOptions) (_result *ListOfficeSiteUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListOfficeSiteUsersResponse{}
	_body, _err := client.DoRequest(tea.String("ListOfficeSiteUsers"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOfficeSiteUsers(request *ListOfficeSiteUsersRequest) (_result *ListOfficeSiteUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOfficeSiteUsersResponse{}
	_body, _err := client.ListOfficeSiteUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSuspEventsWithOptions(request *DescribeSuspEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeSuspEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeSuspEventsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeSuspEvents"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSuspEvents(request *DescribeSuspEventsRequest) (_result *DescribeSuspEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSuspEventsResponse{}
	_body, _err := client.DescribeSuspEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeModificationPriceWithOptions(request *DescribeModificationPriceRequest, runtime *util.RuntimeOptions) (_result *DescribeModificationPriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeModificationPriceResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeModificationPrice"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeModificationPrice(request *DescribeModificationPriceRequest) (_result *DescribeModificationPriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeModificationPriceResponse{}
	_body, _err := client.DescribeModificationPriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteOfficeSitesWithOptions(request *DeleteOfficeSitesRequest, runtime *util.RuntimeOptions) (_result *DeleteOfficeSitesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteOfficeSitesResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteOfficeSites"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteOfficeSites(request *DeleteOfficeSitesRequest) (_result *DeleteOfficeSitesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteOfficeSitesResponse{}
	_body, _err := client.DeleteOfficeSitesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDesktopIdsByVulNamesWithOptions(request *DescribeDesktopIdsByVulNamesRequest, runtime *util.RuntimeOptions) (_result *DescribeDesktopIdsByVulNamesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDesktopIdsByVulNamesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDesktopIdsByVulNames"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDesktopIdsByVulNames(request *DescribeDesktopIdsByVulNamesRequest) (_result *DescribeDesktopIdsByVulNamesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDesktopIdsByVulNamesResponse{}
	_body, _err := client.DescribeDesktopIdsByVulNamesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOfficeSiteSsoStatusWithOptions(request *GetOfficeSiteSsoStatusRequest, runtime *util.RuntimeOptions) (_result *GetOfficeSiteSsoStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetOfficeSiteSsoStatusResponse{}
	_body, _err := client.DoRequest(tea.String("GetOfficeSiteSsoStatus"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOfficeSiteSsoStatus(request *GetOfficeSiteSsoStatusRequest) (_result *GetOfficeSiteSsoStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOfficeSiteSsoStatusResponse{}
	_body, _err := client.GetOfficeSiteSsoStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSecurityEventOperationsWithOptions(request *DescribeSecurityEventOperationsRequest, runtime *util.RuntimeOptions) (_result *DescribeSecurityEventOperationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeSecurityEventOperationsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeSecurityEventOperations"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSecurityEventOperations(request *DescribeSecurityEventOperationsRequest) (_result *DescribeSecurityEventOperationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSecurityEventOperationsResponse{}
	_body, _err := client.DescribeSecurityEventOperationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateNetworkPackageWithOptions(request *CreateNetworkPackageRequest, runtime *util.RuntimeOptions) (_result *CreateNetworkPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateNetworkPackageResponse{}
	_body, _err := client.DoRequest(tea.String("CreateNetworkPackage"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateNetworkPackage(request *CreateNetworkPackageRequest) (_result *CreateNetworkPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNetworkPackageResponse{}
	_body, _err := client.CreateNetworkPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateADConnectorOfficeSiteWithOptions(request *CreateADConnectorOfficeSiteRequest, runtime *util.RuntimeOptions) (_result *CreateADConnectorOfficeSiteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateADConnectorOfficeSiteResponse{}
	_body, _err := client.DoRequest(tea.String("CreateADConnectorOfficeSite"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateADConnectorOfficeSite(request *CreateADConnectorOfficeSiteRequest) (_result *CreateADConnectorOfficeSiteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateADConnectorOfficeSiteResponse{}
	_body, _err := client.CreateADConnectorOfficeSiteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteNetworkPackagesWithOptions(request *DeleteNetworkPackagesRequest, runtime *util.RuntimeOptions) (_result *DeleteNetworkPackagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteNetworkPackagesResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteNetworkPackages"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteNetworkPackages(request *DeleteNetworkPackagesRequest) (_result *DeleteNetworkPackagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNetworkPackagesResponse{}
	_body, _err := client.DeleteNetworkPackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetOfficeSiteSsoStatusWithOptions(request *SetOfficeSiteSsoStatusRequest, runtime *util.RuntimeOptions) (_result *SetOfficeSiteSsoStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetOfficeSiteSsoStatusResponse{}
	_body, _err := client.DoRequest(tea.String("SetOfficeSiteSsoStatus"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetOfficeSiteSsoStatus(request *SetOfficeSiteSsoStatusRequest) (_result *SetOfficeSiteSsoStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetOfficeSiteSsoStatusResponse{}
	_body, _err := client.SetOfficeSiteSsoStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HandleSecurityEventsWithOptions(request *HandleSecurityEventsRequest, runtime *util.RuntimeOptions) (_result *HandleSecurityEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &HandleSecurityEventsResponse{}
	_body, _err := client.DoRequest(tea.String("HandleSecurityEvents"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HandleSecurityEvents(request *HandleSecurityEventsRequest) (_result *HandleSecurityEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &HandleSecurityEventsResponse{}
	_body, _err := client.HandleSecurityEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyNetworkPackageWithOptions(request *ModifyNetworkPackageRequest, runtime *util.RuntimeOptions) (_result *ModifyNetworkPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyNetworkPackageResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyNetworkPackage"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyNetworkPackage(request *ModifyNetworkPackageRequest) (_result *ModifyNetworkPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyNetworkPackageResponse{}
	_body, _err := client.ModifyNetworkPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNetworkPackagesWithOptions(request *DescribeNetworkPackagesRequest, runtime *util.RuntimeOptions) (_result *DescribeNetworkPackagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeNetworkPackagesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeNetworkPackages"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNetworkPackages(request *DescribeNetworkPackagesRequest) (_result *DescribeNetworkPackagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNetworkPackagesResponse{}
	_body, _err := client.DescribeNetworkPackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGroupedVulWithOptions(request *DescribeGroupedVulRequest, runtime *util.RuntimeOptions) (_result *DescribeGroupedVulResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeGroupedVulResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeGroupedVul"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGroupedVul(request *DescribeGroupedVulRequest) (_result *DescribeGroupedVulResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGroupedVulResponse{}
	_body, _err := client.DescribeGroupedVulWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RollbackSuspEventQuaraFileWithOptions(request *RollbackSuspEventQuaraFileRequest, runtime *util.RuntimeOptions) (_result *RollbackSuspEventQuaraFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RollbackSuspEventQuaraFileResponse{}
	_body, _err := client.DoRequest(tea.String("RollbackSuspEventQuaraFile"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RollbackSuspEventQuaraFile(request *RollbackSuspEventQuaraFileRequest) (_result *RollbackSuspEventQuaraFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RollbackSuspEventQuaraFileResponse{}
	_body, _err := client.RollbackSuspEventQuaraFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePriceWithOptions(request *DescribePriceRequest, runtime *util.RuntimeOptions) (_result *DescribePriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribePriceResponse{}
	_body, _err := client.DoRequest(tea.String("DescribePrice"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePrice(request *DescribePriceRequest) (_result *DescribePriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePriceResponse{}
	_body, _err := client.DescribePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDesktopSpecWithOptions(request *ModifyDesktopSpecRequest, runtime *util.RuntimeOptions) (_result *ModifyDesktopSpecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyDesktopSpecResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyDesktopSpec"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDesktopSpec(request *ModifyDesktopSpecRequest) (_result *ModifyDesktopSpecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDesktopSpecResponse{}
	_body, _err := client.ModifyDesktopSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOfficeSiteOverviewWithOptions(request *ListOfficeSiteOverviewRequest, runtime *util.RuntimeOptions) (_result *ListOfficeSiteOverviewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListOfficeSiteOverviewResponse{}
	_body, _err := client.DoRequest(tea.String("ListOfficeSiteOverview"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOfficeSiteOverview(request *ListOfficeSiteOverviewRequest) (_result *ListOfficeSiteOverviewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOfficeSiteOverviewResponse{}
	_body, _err := client.ListOfficeSiteOverviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDirectorySsoStatusWithOptions(request *GetDirectorySsoStatusRequest, runtime *util.RuntimeOptions) (_result *GetDirectorySsoStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetDirectorySsoStatusResponse{}
	_body, _err := client.DoRequest(tea.String("GetDirectorySsoStatus"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDirectorySsoStatus(request *GetDirectorySsoStatusRequest) (_result *GetDirectorySsoStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDirectorySsoStatusResponse{}
	_body, _err := client.GetDirectorySsoStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDirectorySsoStatusWithOptions(request *SetDirectorySsoStatusRequest, runtime *util.RuntimeOptions) (_result *SetDirectorySsoStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetDirectorySsoStatusResponse{}
	_body, _err := client.DoRequest(tea.String("SetDirectorySsoStatus"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDirectorySsoStatus(request *SetDirectorySsoStatusRequest) (_result *SetDirectorySsoStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDirectorySsoStatusResponse{}
	_body, _err := client.SetDirectorySsoStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSpMetadataWithOptions(request *GetSpMetadataRequest, runtime *util.RuntimeOptions) (_result *GetSpMetadataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetSpMetadataResponse{}
	_body, _err := client.DoRequest(tea.String("GetSpMetadata"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSpMetadata(request *GetSpMetadataRequest) (_result *GetSpMetadataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSpMetadataResponse{}
	_body, _err := client.GetSpMetadataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetIdpMetadataWithOptions(request *SetIdpMetadataRequest, runtime *util.RuntimeOptions) (_result *SetIdpMetadataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetIdpMetadataResponse{}
	_body, _err := client.DoRequest(tea.String("SetIdpMetadata"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetIdpMetadata(request *SetIdpMetadataRequest) (_result *SetIdpMetadataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetIdpMetadataResponse{}
	_body, _err := client.SetIdpMetadataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RebuildDesktopsWithOptions(request *RebuildDesktopsRequest, runtime *util.RuntimeOptions) (_result *RebuildDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RebuildDesktopsResponse{}
	_body, _err := client.DoRequest(tea.String("RebuildDesktops"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RebuildDesktops(request *RebuildDesktopsRequest) (_result *RebuildDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RebuildDesktopsResponse{}
	_body, _err := client.RebuildDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyBundleWithOptions(request *ModifyBundleRequest, runtime *util.RuntimeOptions) (_result *ModifyBundleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyBundleResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyBundle"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyBundle(request *ModifyBundleRequest) (_result *ModifyBundleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyBundleResponse{}
	_body, _err := client.ModifyBundleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnlockVirtualMFADeviceWithOptions(request *UnlockVirtualMFADeviceRequest, runtime *util.RuntimeOptions) (_result *UnlockVirtualMFADeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UnlockVirtualMFADeviceResponse{}
	_body, _err := client.DoRequest(tea.String("UnlockVirtualMFADevice"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnlockVirtualMFADevice(request *UnlockVirtualMFADeviceRequest) (_result *UnlockVirtualMFADeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnlockVirtualMFADeviceResponse{}
	_body, _err := client.UnlockVirtualMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVirtualMFADevicesWithOptions(request *DescribeVirtualMFADevicesRequest, runtime *util.RuntimeOptions) (_result *DescribeVirtualMFADevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeVirtualMFADevicesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeVirtualMFADevices"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVirtualMFADevices(request *DescribeVirtualMFADevicesRequest) (_result *DescribeVirtualMFADevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVirtualMFADevicesResponse{}
	_body, _err := client.DescribeVirtualMFADevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LockVirtualMFADeviceWithOptions(request *LockVirtualMFADeviceRequest, runtime *util.RuntimeOptions) (_result *LockVirtualMFADeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &LockVirtualMFADeviceResponse{}
	_body, _err := client.DoRequest(tea.String("LockVirtualMFADevice"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LockVirtualMFADevice(request *LockVirtualMFADeviceRequest) (_result *LockVirtualMFADeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LockVirtualMFADeviceResponse{}
	_body, _err := client.LockVirtualMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVirtualMFADeviceWithOptions(request *DeleteVirtualMFADeviceRequest, runtime *util.RuntimeOptions) (_result *DeleteVirtualMFADeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteVirtualMFADeviceResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteVirtualMFADevice"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVirtualMFADevice(request *DeleteVirtualMFADeviceRequest) (_result *DeleteVirtualMFADeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVirtualMFADeviceResponse{}
	_body, _err := client.DeleteVirtualMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyADConnectorDirectoryWithOptions(request *ModifyADConnectorDirectoryRequest, runtime *util.RuntimeOptions) (_result *ModifyADConnectorDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyADConnectorDirectoryResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyADConnectorDirectory"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyADConnectorDirectory(request *ModifyADConnectorDirectoryRequest) (_result *ModifyADConnectorDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyADConnectorDirectoryResponse{}
	_body, _err := client.ModifyADConnectorDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRequest(tea.String("ListTagResources"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRequest(tea.String("UntagResources"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRequest(tea.String("TagResources"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDesktopNameWithOptions(request *ModifyDesktopNameRequest, runtime *util.RuntimeOptions) (_result *ModifyDesktopNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyDesktopNameResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyDesktopName"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDesktopName(request *ModifyDesktopNameRequest) (_result *ModifyDesktopNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDesktopNameResponse{}
	_body, _err := client.ModifyDesktopNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunCommandWithOptions(request *RunCommandRequest, runtime *util.RuntimeOptions) (_result *RunCommandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RunCommandResponse{}
	_body, _err := client.DoRequest(tea.String("RunCommand"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunCommand(request *RunCommandRequest) (_result *RunCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunCommandResponse{}
	_body, _err := client.RunCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopInvocationWithOptions(request *StopInvocationRequest, runtime *util.RuntimeOptions) (_result *StopInvocationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StopInvocationResponse{}
	_body, _err := client.DoRequest(tea.String("StopInvocation"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopInvocation(request *StopInvocationRequest) (_result *StopInvocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopInvocationResponse{}
	_body, _err := client.StopInvocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInvocationsWithOptions(request *DescribeInvocationsRequest, runtime *util.RuntimeOptions) (_result *DescribeInvocationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeInvocationsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeInvocations"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInvocations(request *DescribeInvocationsRequest) (_result *DescribeInvocationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInvocationsResponse{}
	_body, _err := client.DescribeInvocationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *util.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeZones"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeZones(request *DescribeZonesRequest) (_result *DescribeZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DescribeZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeRegions"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClientEventsWithOptions(request *DescribeClientEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeClientEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClientEventsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeClientEvents"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClientEvents(request *DescribeClientEventsRequest) (_result *DescribeClientEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClientEventsResponse{}
	_body, _err := client.DescribeClientEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetSnapshotWithOptions(request *ResetSnapshotRequest, runtime *util.RuntimeOptions) (_result *ResetSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ResetSnapshotResponse{}
	_body, _err := client.DoRequest(tea.String("ResetSnapshot"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetSnapshot(request *ResetSnapshotRequest) (_result *ResetSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetSnapshotResponse{}
	_body, _err := client.ResetSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSnapshotWithOptions(request *DeleteSnapshotRequest, runtime *util.RuntimeOptions) (_result *DeleteSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteSnapshot"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSnapshot(request *DeleteSnapshotRequest) (_result *DeleteSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.DeleteSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSnapshotWithOptions(request *CreateSnapshotRequest, runtime *util.RuntimeOptions) (_result *CreateSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.DoRequest(tea.String("CreateSnapshot"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSnapshot(request *CreateSnapshotRequest) (_result *CreateSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CreateSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSnapshotsWithOptions(request *DescribeSnapshotsRequest, runtime *util.RuntimeOptions) (_result *DescribeSnapshotsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeSnapshotsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeSnapshots"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (_result *DescribeSnapshotsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSnapshotsResponse{}
	_body, _err := client.DescribeSnapshotsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenewDesktopsWithOptions(request *RenewDesktopsRequest, runtime *util.RuntimeOptions) (_result *RenewDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RenewDesktopsResponse{}
	_body, _err := client.DoRequest(tea.String("RenewDesktops"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenewDesktops(request *RenewDesktopsRequest) (_result *RenewDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewDesktopsResponse{}
	_body, _err := client.RenewDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopDesktopsWithOptions(request *StopDesktopsRequest, runtime *util.RuntimeOptions) (_result *StopDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StopDesktopsResponse{}
	_body, _err := client.DoRequest(tea.String("StopDesktops"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopDesktops(request *StopDesktopsRequest) (_result *StopDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopDesktopsResponse{}
	_body, _err := client.StopDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartDesktopsWithOptions(request *StartDesktopsRequest, runtime *util.RuntimeOptions) (_result *StartDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StartDesktopsResponse{}
	_body, _err := client.DoRequest(tea.String("StartDesktops"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartDesktops(request *StartDesktopsRequest) (_result *StartDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartDesktopsResponse{}
	_body, _err := client.StartDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyPolicyGroupWithOptions(request *ModifyPolicyGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyPolicyGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyPolicyGroupResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyPolicyGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyPolicyGroup(request *ModifyPolicyGroupRequest) (_result *ModifyPolicyGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPolicyGroupResponse{}
	_body, _err := client.ModifyPolicyGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDesktopTypesWithOptions(request *DescribeDesktopTypesRequest, runtime *util.RuntimeOptions) (_result *DescribeDesktopTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDesktopTypesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDesktopTypes"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDesktopTypes(request *DescribeDesktopTypesRequest) (_result *DescribeDesktopTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDesktopTypesResponse{}
	_body, _err := client.DescribeDesktopTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDirectoriesWithOptions(request *DescribeDirectoriesRequest, runtime *util.RuntimeOptions) (_result *DescribeDirectoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDirectoriesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDirectories"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDirectories(request *DescribeDirectoriesRequest) (_result *DescribeDirectoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDirectoriesResponse{}
	_body, _err := client.DescribeDirectoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDirectoriesWithOptions(request *DeleteDirectoriesRequest, runtime *util.RuntimeOptions) (_result *DeleteDirectoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteDirectoriesResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteDirectories"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDirectories(request *DeleteDirectoriesRequest) (_result *DeleteDirectoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDirectoriesResponse{}
	_body, _err := client.DeleteDirectoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDirectoryUsersWithOptions(request *ListDirectoryUsersRequest, runtime *util.RuntimeOptions) (_result *ListDirectoryUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListDirectoryUsersResponse{}
	_body, _err := client.DoRequest(tea.String("ListDirectoryUsers"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDirectoryUsers(request *ListDirectoryUsersRequest) (_result *ListDirectoryUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDirectoryUsersResponse{}
	_body, _err := client.ListDirectoryUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateImageWithOptions(request *CreateImageRequest, runtime *util.RuntimeOptions) (_result *CreateImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateImageResponse{}
	_body, _err := client.DoRequest(tea.String("CreateImage"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateImage(request *CreateImageRequest) (_result *CreateImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateImageResponse{}
	_body, _err := client.CreateImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRAMDirectoryWithOptions(request *CreateRAMDirectoryRequest, runtime *util.RuntimeOptions) (_result *CreateRAMDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateRAMDirectoryResponse{}
	_body, _err := client.DoRequest(tea.String("CreateRAMDirectory"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRAMDirectory(request *CreateRAMDirectoryRequest) (_result *CreateRAMDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRAMDirectoryResponse{}
	_body, _err := client.CreateRAMDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePolicyGroupsWithOptions(request *DeletePolicyGroupsRequest, runtime *util.RuntimeOptions) (_result *DeletePolicyGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeletePolicyGroupsResponse{}
	_body, _err := client.DoRequest(tea.String("DeletePolicyGroups"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePolicyGroups(request *DeletePolicyGroupsRequest) (_result *DeletePolicyGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePolicyGroupsResponse{}
	_body, _err := client.DeletePolicyGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePolicyGroupsWithOptions(request *DescribePolicyGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribePolicyGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribePolicyGroupsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribePolicyGroups"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePolicyGroups(request *DescribePolicyGroupsRequest) (_result *DescribePolicyGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePolicyGroupsResponse{}
	_body, _err := client.DescribePolicyGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDesktopsWithOptions(request *DeleteDesktopsRequest, runtime *util.RuntimeOptions) (_result *DeleteDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteDesktopsResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteDesktops"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDesktops(request *DeleteDesktopsRequest) (_result *DeleteDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDesktopsResponse{}
	_body, _err := client.DeleteDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyImageAttributeWithOptions(request *ModifyImageAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyImageAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyImageAttributeResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyImageAttribute"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyImageAttribute(request *ModifyImageAttributeRequest) (_result *ModifyImageAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyImageAttributeResponse{}
	_body, _err := client.ModifyImageAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyEntitlementWithOptions(request *ModifyEntitlementRequest, runtime *util.RuntimeOptions) (_result *ModifyEntitlementResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyEntitlementResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyEntitlement"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyEntitlement(request *ModifyEntitlementRequest) (_result *ModifyEntitlementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyEntitlementResponse{}
	_body, _err := client.ModifyEntitlementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBundlesWithOptions(request *DeleteBundlesRequest, runtime *util.RuntimeOptions) (_result *DeleteBundlesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteBundlesResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteBundles"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBundles(request *DeleteBundlesRequest) (_result *DeleteBundlesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBundlesResponse{}
	_body, _err := client.DeleteBundlesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDesktopsWithOptions(request *DescribeDesktopsRequest, runtime *util.RuntimeOptions) (_result *DescribeDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDesktopsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDesktops"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDesktops(request *DescribeDesktopsRequest) (_result *DescribeDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDesktopsResponse{}
	_body, _err := client.DescribeDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RebootDesktopsWithOptions(request *RebootDesktopsRequest, runtime *util.RuntimeOptions) (_result *RebootDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RebootDesktopsResponse{}
	_body, _err := client.DoRequest(tea.String("RebootDesktops"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RebootDesktops(request *RebootDesktopsRequest) (_result *RebootDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RebootDesktopsResponse{}
	_body, _err := client.RebootDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBundleWithOptions(request *CreateBundleRequest, runtime *util.RuntimeOptions) (_result *CreateBundleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateBundleResponse{}
	_body, _err := client.DoRequest(tea.String("CreateBundle"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBundle(request *CreateBundleRequest) (_result *CreateBundleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBundleResponse{}
	_body, _err := client.CreateBundleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDesktopsPolicyGroupWithOptions(request *ModifyDesktopsPolicyGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyDesktopsPolicyGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyDesktopsPolicyGroupResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyDesktopsPolicyGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDesktopsPolicyGroup(request *ModifyDesktopsPolicyGroupRequest) (_result *ModifyDesktopsPolicyGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDesktopsPolicyGroupResponse{}
	_body, _err := client.ModifyDesktopsPolicyGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePolicyGroupWithOptions(request *CreatePolicyGroupRequest, runtime *util.RuntimeOptions) (_result *CreatePolicyGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreatePolicyGroupResponse{}
	_body, _err := client.DoRequest(tea.String("CreatePolicyGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePolicyGroup(request *CreatePolicyGroupRequest) (_result *CreatePolicyGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePolicyGroupResponse{}
	_body, _err := client.CreatePolicyGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateADConnectorDirectoryWithOptions(request *CreateADConnectorDirectoryRequest, runtime *util.RuntimeOptions) (_result *CreateADConnectorDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateADConnectorDirectoryResponse{}
	_body, _err := client.DoRequest(tea.String("CreateADConnectorDirectory"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateADConnectorDirectory(request *CreateADConnectorDirectoryRequest) (_result *CreateADConnectorDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateADConnectorDirectoryResponse{}
	_body, _err := client.CreateADConnectorDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBundlesWithOptions(request *DescribeBundlesRequest, runtime *util.RuntimeOptions) (_result *DescribeBundlesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeBundlesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeBundles"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBundles(request *DescribeBundlesRequest) (_result *DescribeBundlesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBundlesResponse{}
	_body, _err := client.DescribeBundlesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteImagesWithOptions(request *DeleteImagesRequest, runtime *util.RuntimeOptions) (_result *DeleteImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteImagesResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteImages"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteImages(request *DeleteImagesRequest) (_result *DeleteImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteImagesResponse{}
	_body, _err := client.DeleteImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDesktopsWithOptions(request *CreateDesktopsRequest, runtime *util.RuntimeOptions) (_result *CreateDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateDesktopsResponse{}
	_body, _err := client.DoRequest(tea.String("CreateDesktops"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDesktops(request *CreateDesktopsRequest) (_result *CreateDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDesktopsResponse{}
	_body, _err := client.CreateDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImagesWithOptions(request *DescribeImagesRequest, runtime *util.RuntimeOptions) (_result *DescribeImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeImagesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeImages"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImages(request *DescribeImagesRequest) (_result *DescribeImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImagesResponse{}
	_body, _err := client.DescribeImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
