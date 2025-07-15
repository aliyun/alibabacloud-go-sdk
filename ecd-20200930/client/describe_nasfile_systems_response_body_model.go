// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNASFileSystemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystems(v []*DescribeNASFileSystemsResponseBodyFileSystems) *DescribeNASFileSystemsResponseBody
	GetFileSystems() []*DescribeNASFileSystemsResponseBodyFileSystems
	SetNextToken(v string) *DescribeNASFileSystemsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeNASFileSystemsResponseBody
	GetRequestId() *string
}

type DescribeNASFileSystemsResponseBody struct {
	// The NAS file systems.
	FileSystems []*DescribeNASFileSystemsResponseBodyFileSystems `json:"FileSystems,omitempty" xml:"FileSystems,omitempty" type:"Repeated"`
	// The token that determines the start point of the next query. This parameter is empty if no additional results exist.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 269BDB16-2CD8-4865-84BD-11C40BC21DB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNASFileSystemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNASFileSystemsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNASFileSystemsResponseBody) GetFileSystems() []*DescribeNASFileSystemsResponseBodyFileSystems {
	return s.FileSystems
}

func (s *DescribeNASFileSystemsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeNASFileSystemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNASFileSystemsResponseBody) SetFileSystems(v []*DescribeNASFileSystemsResponseBodyFileSystems) *DescribeNASFileSystemsResponseBody {
	s.FileSystems = v
	return s
}

func (s *DescribeNASFileSystemsResponseBody) SetNextToken(v string) *DescribeNASFileSystemsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBody) SetRequestId(v string) *DescribeNASFileSystemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeNASFileSystemsResponseBodyFileSystems struct {
	// >  This parameter is not publicly available.
	AllowOperateUserDrive *bool `json:"AllowOperateUserDrive,omitempty" xml:"AllowOperateUserDrive,omitempty"`
	// The application delivery groups that are associated with the UPM-supported NAS file systems.
	AppInstanceGroups []*DescribeNASFileSystemsResponseBodyFileSystemsAppInstanceGroups `json:"AppInstanceGroups,omitempty" xml:"AppInstanceGroups,omitempty" type:"Repeated"`
	// The total capacity of the NAS file system. Unit: GiB.
	//
	// 	- The Capacity type has 10 PiB of storage, which is equal to 10,485,760 GiB.
	//
	// 	- The Performance type has 1 PiB of storage, which is equal to 1,048,576 GiB.
	//
	// example:
	//
	// 10485760
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The time when the NAS file system was created.
	//
	// example:
	//
	// 2021-05-10T11:39Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the NAS file system.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The cloud computer shares that are associated with the UPM-supported NAS file systems.
	DesktopGroups []*DescribeNASFileSystemsResponseBodyFileSystemsDesktopGroups `json:"DesktopGroups,omitempty" xml:"DesktopGroups,omitempty" type:"Repeated"`
	// Indicates whether disk encryption is enabled.
	//
	// example:
	//
	// false
	EncryptionEnabled *bool `json:"EncryptionEnabled,omitempty" xml:"EncryptionEnabled,omitempty"`
	// The ID of the NAS file system.
	//
	// example:
	//
	// 04f314****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The name of the NAS file system.
	//
	// example:
	//
	// testNAS
	FileSystemName *string `json:"FileSystemName,omitempty" xml:"FileSystemName,omitempty"`
	// The status of the NAS file system. The possible values include:
	//
	// 	- Pending: The NAS file system is being created.
	//
	// 	- Running: The NAS file system is running.
	//
	// 	- Stopped: The NAS file system is stopped.
	//
	// 	- Deleting: The NAS file system is being deleted.
	//
	// 	- Deleted: The NAS file system is deleted.
	//
	// 	- Invalid: The NAS file system is invalid.
	//
	// example:
	//
	// Running
	FileSystemStatus *string `json:"FileSystemStatus,omitempty" xml:"FileSystemStatus,omitempty"`
	// The type of the NAS file system. The only valid value is `standard`.
	//
	// example:
	//
	// standard
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The used capacity of the NAS file system. Unit: bytes.
	//
	// example:
	//
	// 0
	MeteredSize *int64 `json:"MeteredSize,omitempty" xml:"MeteredSize,omitempty"`
	// The domain name of the mount target.
	//
	// example:
	//
	// 04f314****-at***.cn-hangzhou.nas.aliyuncs.com
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	// The status of the mount target. The possible values include:
	//
	// 	- Pending: The mount target is being created.
	//
	// 	- Active: The mount target is enabled.
	//
	// 	- Inactive: The mount target is disabled.
	//
	// 	- Deleting: The mount target is being deleted.
	//
	// 	- Invalid: The mount target is invalid.
	//
	// example:
	//
	// Active
	MountTargetStatus *string `json:"MountTargetStatus,omitempty" xml:"MountTargetStatus,omitempty"`
	// The ID of the office network.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The name of the office network.
	//
	// example:
	//
	// test
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	// The office networks.
	OfficeSites []*DescribeNASFileSystemsResponseBodyFileSystemsOfficeSites `json:"OfficeSites,omitempty" xml:"OfficeSites,omitempty" type:"Repeated"`
	// Indicates whether the User Profile Management (UPM) feature is supported.
	//
	// example:
	//
	// false
	ProfileCompatible *bool `json:"ProfileCompatible,omitempty" xml:"ProfileCompatible,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The storage type of the NAS file system.
	//
	// Valid values:
	//
	// 	- Upm: the UPM-supported NAS file system.
	//
	// 	- ShareNas: the shared NAS file system.
	//
	// example:
	//
	// Upm
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The storage type of the NAS file system. Valid values:
	//
	// 	- Capacity
	//
	// 	- Performance
	//
	// example:
	//
	// Capacity
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// Indicates whether the Server Message Block (SMB) access control list (ACL) feature was enabled.
	//
	// example:
	//
	// false
	SupportAcl *bool `json:"SupportAcl,omitempty" xml:"SupportAcl,omitempty"`
	// The ID of the zone where the NAS file system resides.
	//
	// example:
	//
	// cn-hangzhou-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeNASFileSystemsResponseBodyFileSystems) String() string {
	return dara.Prettify(s)
}

func (s DescribeNASFileSystemsResponseBodyFileSystems) GoString() string {
	return s.String()
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) GetAllowOperateUserDrive() *bool {
	return s.AllowOperateUserDrive
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) GetAppInstanceGroups() []*DescribeNASFileSystemsResponseBodyFileSystemsAppInstanceGroups {
	return s.AppInstanceGroups
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) GetCapacity() *int64 {
	return s.Capacity
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) GetDescription() *string {
	return s.Description
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) GetDesktopGroups() []*DescribeNASFileSystemsResponseBodyFileSystemsDesktopGroups {
	return s.DesktopGroups
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) GetEncryptionEnabled() *bool {
	return s.EncryptionEnabled
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) GetFileSystemName() *string {
	return s.FileSystemName
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) GetFileSystemStatus() *string {
	return s.FileSystemStatus
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) GetMeteredSize() *int64 {
	return s.MeteredSize
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) GetMountTargetDomain() *string {
	return s.MountTargetDomain
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) GetMountTargetStatus() *string {
	return s.MountTargetStatus
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) GetOfficeSiteName() *string {
	return s.OfficeSiteName
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) GetOfficeSites() []*DescribeNASFileSystemsResponseBodyFileSystemsOfficeSites {
	return s.OfficeSites
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) GetProfileCompatible() *bool {
	return s.ProfileCompatible
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) GetScene() *string {
	return s.Scene
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) GetSupportAcl() *bool {
	return s.SupportAcl
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetAllowOperateUserDrive(v bool) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.AllowOperateUserDrive = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetAppInstanceGroups(v []*DescribeNASFileSystemsResponseBodyFileSystemsAppInstanceGroups) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.AppInstanceGroups = v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetCapacity(v int64) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.Capacity = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetCreateTime(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.CreateTime = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetDescription(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.Description = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetDesktopGroups(v []*DescribeNASFileSystemsResponseBodyFileSystemsDesktopGroups) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.DesktopGroups = v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetEncryptionEnabled(v bool) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.EncryptionEnabled = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetFileSystemId(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.FileSystemId = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetFileSystemName(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.FileSystemName = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetFileSystemStatus(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.FileSystemStatus = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetFileSystemType(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.FileSystemType = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetMeteredSize(v int64) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.MeteredSize = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetMountTargetDomain(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.MountTargetDomain = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetMountTargetStatus(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.MountTargetStatus = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetOfficeSiteId(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetOfficeSiteName(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.OfficeSiteName = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetOfficeSites(v []*DescribeNASFileSystemsResponseBodyFileSystemsOfficeSites) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.OfficeSites = v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetProfileCompatible(v bool) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.ProfileCompatible = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetRegionId(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.RegionId = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetScene(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.Scene = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetStorageType(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.StorageType = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetSupportAcl(v bool) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.SupportAcl = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetZoneId(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.ZoneId = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) Validate() error {
	return dara.Validate(s)
}

type DescribeNASFileSystemsResponseBodyFileSystemsAppInstanceGroups struct {
	// The ID of the delivery group.
	//
	// example:
	//
	// aig-0bz55ibznu9p7****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The name of the delivery group.
	//
	// example:
	//
	// DemoDeliveryGroup
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
}

func (s DescribeNASFileSystemsResponseBodyFileSystemsAppInstanceGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeNASFileSystemsResponseBodyFileSystemsAppInstanceGroups) GoString() string {
	return s.String()
}

func (s *DescribeNASFileSystemsResponseBodyFileSystemsAppInstanceGroups) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *DescribeNASFileSystemsResponseBodyFileSystemsAppInstanceGroups) GetAppInstanceGroupName() *string {
	return s.AppInstanceGroupName
}

func (s *DescribeNASFileSystemsResponseBodyFileSystemsAppInstanceGroups) SetAppInstanceGroupId(v string) *DescribeNASFileSystemsResponseBodyFileSystemsAppInstanceGroups {
	s.AppInstanceGroupId = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystemsAppInstanceGroups) SetAppInstanceGroupName(v string) *DescribeNASFileSystemsResponseBodyFileSystemsAppInstanceGroups {
	s.AppInstanceGroupName = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystemsAppInstanceGroups) Validate() error {
	return dara.Validate(s)
}

type DescribeNASFileSystemsResponseBodyFileSystemsDesktopGroups struct {
	// The ID of the cloud computer share.
	//
	// example:
	//
	// dg-9eeyf15b25nyl****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The name of the cloud computer share.
	//
	// example:
	//
	// test_dg
	DesktopGroupName *string `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty"`
}

func (s DescribeNASFileSystemsResponseBodyFileSystemsDesktopGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeNASFileSystemsResponseBodyFileSystemsDesktopGroups) GoString() string {
	return s.String()
}

func (s *DescribeNASFileSystemsResponseBodyFileSystemsDesktopGroups) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *DescribeNASFileSystemsResponseBodyFileSystemsDesktopGroups) GetDesktopGroupName() *string {
	return s.DesktopGroupName
}

func (s *DescribeNASFileSystemsResponseBodyFileSystemsDesktopGroups) SetDesktopGroupId(v string) *DescribeNASFileSystemsResponseBodyFileSystemsDesktopGroups {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystemsDesktopGroups) SetDesktopGroupName(v string) *DescribeNASFileSystemsResponseBodyFileSystemsDesktopGroups {
	s.DesktopGroupName = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystemsDesktopGroups) Validate() error {
	return dara.Validate(s)
}

type DescribeNASFileSystemsResponseBodyFileSystemsOfficeSites struct {
	// The ID of the office network.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The name of the office network.
	//
	// example:
	//
	// DemoOfficeNetwork
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
}

func (s DescribeNASFileSystemsResponseBodyFileSystemsOfficeSites) String() string {
	return dara.Prettify(s)
}

func (s DescribeNASFileSystemsResponseBodyFileSystemsOfficeSites) GoString() string {
	return s.String()
}

func (s *DescribeNASFileSystemsResponseBodyFileSystemsOfficeSites) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeNASFileSystemsResponseBodyFileSystemsOfficeSites) GetOfficeSiteName() *string {
	return s.OfficeSiteName
}

func (s *DescribeNASFileSystemsResponseBodyFileSystemsOfficeSites) SetOfficeSiteId(v string) *DescribeNASFileSystemsResponseBodyFileSystemsOfficeSites {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystemsOfficeSites) SetOfficeSiteName(v string) *DescribeNASFileSystemsResponseBodyFileSystemsOfficeSites {
	s.OfficeSiteName = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystemsOfficeSites) Validate() error {
	return dara.Validate(s)
}
