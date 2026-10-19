// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEngineVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBLatestMinorVersion(v *DescribeEngineVersionResponseBodyDBLatestMinorVersion) *DescribeEngineVersionResponseBody
	GetDBLatestMinorVersion() *DescribeEngineVersionResponseBodyDBLatestMinorVersion
	SetDBVersionRelease(v string) *DescribeEngineVersionResponseBody
	GetDBVersionRelease() *string
	SetEnableUpgradeMajorVersion(v bool) *DescribeEngineVersionResponseBody
	GetEnableUpgradeMajorVersion() *bool
	SetEnableUpgradeMinorVersion(v bool) *DescribeEngineVersionResponseBody
	GetEnableUpgradeMinorVersion() *bool
	SetEngine(v string) *DescribeEngineVersionResponseBody
	GetEngine() *string
	SetIsAutoUpgradeOpen(v string) *DescribeEngineVersionResponseBody
	GetIsAutoUpgradeOpen() *string
	SetIsLatestVersion(v bool) *DescribeEngineVersionResponseBody
	GetIsLatestVersion() *bool
	SetIsNewSSLMode(v string) *DescribeEngineVersionResponseBody
	GetIsNewSSLMode() *string
	SetIsOpenNGLB(v string) *DescribeEngineVersionResponseBody
	GetIsOpenNGLB() *string
	SetIsRedisCompatibleVersion(v string) *DescribeEngineVersionResponseBody
	GetIsRedisCompatibleVersion() *string
	SetIsSSLEnable(v string) *DescribeEngineVersionResponseBody
	GetIsSSLEnable() *string
	SetMajorVersion(v string) *DescribeEngineVersionResponseBody
	GetMajorVersion() *string
	SetMinorVersion(v string) *DescribeEngineVersionResponseBody
	GetMinorVersion() *string
	SetProxyLatestMinorVersion(v *DescribeEngineVersionResponseBodyProxyLatestMinorVersion) *DescribeEngineVersionResponseBody
	GetProxyLatestMinorVersion() *DescribeEngineVersionResponseBodyProxyLatestMinorVersion
	SetProxyMinorVersion(v string) *DescribeEngineVersionResponseBody
	GetProxyMinorVersion() *string
	SetProxyVersionRelease(v string) *DescribeEngineVersionResponseBody
	GetProxyVersionRelease() *string
	SetRequestId(v string) *DescribeEngineVersionResponseBody
	GetRequestId() *string
}

type DescribeEngineVersionResponseBody struct {
	// The information about the latest minor version to which the instance can be upgraded.
	DBLatestMinorVersion *DescribeEngineVersionResponseBodyDBLatestMinorVersion `json:"DBLatestMinorVersion,omitempty" xml:"DBLatestMinorVersion,omitempty" type:"Struct"`
	// The update notes for the minor version of the instance, including the release time, minor version number, type (such as new feature), and detailed description.
	//
	// example:
	//
	// {\\"releaseInfo\\":{\\"createTime\\":\\"2021-07-27\\",\\"level\\":1,\\"releaseVersion\\":\\"0.5.4\\",\\"releaseNote\\":\\"Feature update: Enhanced stability.\\"}],\\"versionChangesLevel\\":2}
	DBVersionRelease *string `json:"DBVersionRelease,omitempty" xml:"DBVersionRelease,omitempty"`
	// Indicates whether the major version of the instance can be upgraded. Valid values:
	//
	// 	- **true**: The major version can be upgraded.
	//
	// 	- **false**: The major version cannot be upgraded, usually because it is already the latest version.
	//
	// > To upgrade the major version, call the [ModifyInstanceMajorVersion](https://help.aliyun.com/document_detail/473776.html) operation.
	//
	// example:
	//
	// true
	EnableUpgradeMajorVersion *bool `json:"EnableUpgradeMajorVersion,omitempty" xml:"EnableUpgradeMajorVersion,omitempty"`
	// Indicates whether the minor version of the instance can be upgraded. Valid values:
	//
	// 	- **true**: The minor version can be upgraded.
	//
	// 	- **false**: The minor version is already the latest and cannot be upgraded.
	//
	// > To upgrade the minor version, call the [ModifyInstanceMinorVersion](https://help.aliyun.com/document_detail/473777.html) operation.
	//
	// example:
	//
	// true
	EnableUpgradeMinorVersion *bool `json:"EnableUpgradeMinorVersion,omitempty" xml:"EnableUpgradeMinorVersion,omitempty"`
	// The database engine type. Valid values: **redis*	- and **memcache**.
	//
	// example:
	//
	// redis
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// Indicates whether automatic minor version upgrade is enabled. Valid values:
	//
	// - **0**: Disabled.
	//
	// - **1**: Enabled.
	//
	// example:
	//
	// 0
	IsAutoUpgradeOpen *string `json:"IsAutoUpgradeOpen,omitempty" xml:"IsAutoUpgradeOpen,omitempty"`
	// Indicates whether the minor version of the instance is the latest. Valid values:
	//
	// 	- **true**: Yes.
	//
	// 	- **false**: No.
	//
	// example:
	//
	// false
	IsLatestVersion *bool `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
	// Indicates whether the new TLS version is supported. Valid values:
	//
	// 	- **1**: Enabled.
	//
	// 	- **0**: Disabled.
	//
	// example:
	//
	// 1
	IsNewSSLMode *string `json:"IsNewSSLMode,omitempty" xml:"IsNewSSLMode,omitempty"`
	// Indicates whether NGLB mode is enabled. Valid values:
	//
	// 	- **0**: Disabled.
	//
	// 	- **1**: Enabled.
	//
	// example:
	//
	// 1
	IsOpenNGLB *string `json:"IsOpenNGLB,omitempty" xml:"IsOpenNGLB,omitempty"`
	// Indicates whether the instance is a Redis-compatible version.
	//
	// example:
	//
	// 1
	IsRedisCompatibleVersion *string `json:"IsRedisCompatibleVersion,omitempty" xml:"IsRedisCompatibleVersion,omitempty"`
	// Indicates whether SSL is enabled. Valid values:
	//
	// 	- **1**: Enabled.
	//
	// 	- **0**: Disabled.
	//
	// example:
	//
	// 1
	IsSSLEnable *string `json:"IsSSLEnable,omitempty" xml:"IsSSLEnable,omitempty"`
	// The major version of the instance.
	//
	// example:
	//
	// 5.0
	MajorVersion *string `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	// The current minor version of the instance.
	//
	// example:
	//
	// redis-5.0_0.5.0
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	// The information about the latest minor version to which the proxy node can be upgraded.
	ProxyLatestMinorVersion *DescribeEngineVersionResponseBodyProxyLatestMinorVersion `json:"ProxyLatestMinorVersion,omitempty" xml:"ProxyLatestMinorVersion,omitempty" type:"Struct"`
	// The current minor version of the proxy node.
	//
	// > This parameter is returned only for instances of the cluster or read/write splitting architecture.
	//
	// example:
	//
	// 6.6.0
	ProxyMinorVersion *string `json:"ProxyMinorVersion,omitempty" xml:"ProxyMinorVersion,omitempty"`
	// The update notes for the minor version of the proxy node, including the release time, minor version number, type (such as new feature), and detailed description.
	//
	// > This parameter is returned only for instances of the cluster or read/write splitting architecture.
	//
	// example:
	//
	// {\\"releaseInfo\\":[{\\"createTime\\":\\"2021-06-08\\",\\"level\\":0,\\"releaseVersion\\":\\"6.6.2\\",\\"releaseNote\\":\\"New feature: Added support for some internal commands.\\"}],\\"versionChangesLevel\\":2}
	ProxyVersionRelease *string `json:"ProxyVersionRelease,omitempty" xml:"ProxyVersionRelease,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A52974D1-9D57-4805-86CC-92E6EDE8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEngineVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEngineVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEngineVersionResponseBody) GetDBLatestMinorVersion() *DescribeEngineVersionResponseBodyDBLatestMinorVersion {
	return s.DBLatestMinorVersion
}

func (s *DescribeEngineVersionResponseBody) GetDBVersionRelease() *string {
	return s.DBVersionRelease
}

func (s *DescribeEngineVersionResponseBody) GetEnableUpgradeMajorVersion() *bool {
	return s.EnableUpgradeMajorVersion
}

func (s *DescribeEngineVersionResponseBody) GetEnableUpgradeMinorVersion() *bool {
	return s.EnableUpgradeMinorVersion
}

func (s *DescribeEngineVersionResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeEngineVersionResponseBody) GetIsAutoUpgradeOpen() *string {
	return s.IsAutoUpgradeOpen
}

func (s *DescribeEngineVersionResponseBody) GetIsLatestVersion() *bool {
	return s.IsLatestVersion
}

func (s *DescribeEngineVersionResponseBody) GetIsNewSSLMode() *string {
	return s.IsNewSSLMode
}

func (s *DescribeEngineVersionResponseBody) GetIsOpenNGLB() *string {
	return s.IsOpenNGLB
}

func (s *DescribeEngineVersionResponseBody) GetIsRedisCompatibleVersion() *string {
	return s.IsRedisCompatibleVersion
}

func (s *DescribeEngineVersionResponseBody) GetIsSSLEnable() *string {
	return s.IsSSLEnable
}

func (s *DescribeEngineVersionResponseBody) GetMajorVersion() *string {
	return s.MajorVersion
}

func (s *DescribeEngineVersionResponseBody) GetMinorVersion() *string {
	return s.MinorVersion
}

func (s *DescribeEngineVersionResponseBody) GetProxyLatestMinorVersion() *DescribeEngineVersionResponseBodyProxyLatestMinorVersion {
	return s.ProxyLatestMinorVersion
}

func (s *DescribeEngineVersionResponseBody) GetProxyMinorVersion() *string {
	return s.ProxyMinorVersion
}

func (s *DescribeEngineVersionResponseBody) GetProxyVersionRelease() *string {
	return s.ProxyVersionRelease
}

func (s *DescribeEngineVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEngineVersionResponseBody) SetDBLatestMinorVersion(v *DescribeEngineVersionResponseBodyDBLatestMinorVersion) *DescribeEngineVersionResponseBody {
	s.DBLatestMinorVersion = v
	return s
}

func (s *DescribeEngineVersionResponseBody) SetDBVersionRelease(v string) *DescribeEngineVersionResponseBody {
	s.DBVersionRelease = &v
	return s
}

func (s *DescribeEngineVersionResponseBody) SetEnableUpgradeMajorVersion(v bool) *DescribeEngineVersionResponseBody {
	s.EnableUpgradeMajorVersion = &v
	return s
}

func (s *DescribeEngineVersionResponseBody) SetEnableUpgradeMinorVersion(v bool) *DescribeEngineVersionResponseBody {
	s.EnableUpgradeMinorVersion = &v
	return s
}

func (s *DescribeEngineVersionResponseBody) SetEngine(v string) *DescribeEngineVersionResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeEngineVersionResponseBody) SetIsAutoUpgradeOpen(v string) *DescribeEngineVersionResponseBody {
	s.IsAutoUpgradeOpen = &v
	return s
}

func (s *DescribeEngineVersionResponseBody) SetIsLatestVersion(v bool) *DescribeEngineVersionResponseBody {
	s.IsLatestVersion = &v
	return s
}

func (s *DescribeEngineVersionResponseBody) SetIsNewSSLMode(v string) *DescribeEngineVersionResponseBody {
	s.IsNewSSLMode = &v
	return s
}

func (s *DescribeEngineVersionResponseBody) SetIsOpenNGLB(v string) *DescribeEngineVersionResponseBody {
	s.IsOpenNGLB = &v
	return s
}

func (s *DescribeEngineVersionResponseBody) SetIsRedisCompatibleVersion(v string) *DescribeEngineVersionResponseBody {
	s.IsRedisCompatibleVersion = &v
	return s
}

func (s *DescribeEngineVersionResponseBody) SetIsSSLEnable(v string) *DescribeEngineVersionResponseBody {
	s.IsSSLEnable = &v
	return s
}

func (s *DescribeEngineVersionResponseBody) SetMajorVersion(v string) *DescribeEngineVersionResponseBody {
	s.MajorVersion = &v
	return s
}

func (s *DescribeEngineVersionResponseBody) SetMinorVersion(v string) *DescribeEngineVersionResponseBody {
	s.MinorVersion = &v
	return s
}

func (s *DescribeEngineVersionResponseBody) SetProxyLatestMinorVersion(v *DescribeEngineVersionResponseBodyProxyLatestMinorVersion) *DescribeEngineVersionResponseBody {
	s.ProxyLatestMinorVersion = v
	return s
}

func (s *DescribeEngineVersionResponseBody) SetProxyMinorVersion(v string) *DescribeEngineVersionResponseBody {
	s.ProxyMinorVersion = &v
	return s
}

func (s *DescribeEngineVersionResponseBody) SetProxyVersionRelease(v string) *DescribeEngineVersionResponseBody {
	s.ProxyVersionRelease = &v
	return s
}

func (s *DescribeEngineVersionResponseBody) SetRequestId(v string) *DescribeEngineVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEngineVersionResponseBody) Validate() error {
	if s.DBLatestMinorVersion != nil {
		if err := s.DBLatestMinorVersion.Validate(); err != nil {
			return err
		}
	}
	if s.ProxyLatestMinorVersion != nil {
		if err := s.ProxyLatestMinorVersion.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEngineVersionResponseBodyDBLatestMinorVersion struct {
	// The importance level of the version change. Valid values:
	//
	// 	- **0**: Normal.
	//
	// 	- **1**: Important.
	//
	// 	- **2**: Critical.
	//
	// example:
	//
	// 0
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The version number.
	//
	// example:
	//
	// 7.0.1.4
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	// The version evolution path from the current minor version of the instance to the latest minor version. This is consistent with the version documentation. You can refer to the version release notes for more detailed information.
	VersionRelease *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionRelease `json:"VersionRelease,omitempty" xml:"VersionRelease,omitempty" type:"Struct"`
}

func (s DescribeEngineVersionResponseBodyDBLatestMinorVersion) String() string {
	return dara.Prettify(s)
}

func (s DescribeEngineVersionResponseBodyDBLatestMinorVersion) GoString() string {
	return s.String()
}

func (s *DescribeEngineVersionResponseBodyDBLatestMinorVersion) GetLevel() *string {
	return s.Level
}

func (s *DescribeEngineVersionResponseBodyDBLatestMinorVersion) GetMinorVersion() *string {
	return s.MinorVersion
}

func (s *DescribeEngineVersionResponseBodyDBLatestMinorVersion) GetVersionRelease() *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionRelease {
	return s.VersionRelease
}

func (s *DescribeEngineVersionResponseBodyDBLatestMinorVersion) SetLevel(v string) *DescribeEngineVersionResponseBodyDBLatestMinorVersion {
	s.Level = &v
	return s
}

func (s *DescribeEngineVersionResponseBodyDBLatestMinorVersion) SetMinorVersion(v string) *DescribeEngineVersionResponseBodyDBLatestMinorVersion {
	s.MinorVersion = &v
	return s
}

func (s *DescribeEngineVersionResponseBodyDBLatestMinorVersion) SetVersionRelease(v *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionRelease) *DescribeEngineVersionResponseBodyDBLatestMinorVersion {
	s.VersionRelease = v
	return s
}

func (s *DescribeEngineVersionResponseBodyDBLatestMinorVersion) Validate() error {
	if s.VersionRelease != nil {
		if err := s.VersionRelease.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionRelease struct {
	ReleaseInfo *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfo `json:"ReleaseInfo,omitempty" xml:"ReleaseInfo,omitempty" type:"Struct"`
	// The importance level of the version upgrade (recommended upgrade level). Valid values:
	//
	// 	- 0: Normal.
	//
	// 	- 1: Important.
	//
	// 	- 2: Critical.
	//
	// example:
	//
	// 0
	VersionChangesLevel *string `json:"VersionChangesLevel,omitempty" xml:"VersionChangesLevel,omitempty"`
}

func (s DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionRelease) String() string {
	return dara.Prettify(s)
}

func (s DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionRelease) GoString() string {
	return s.String()
}

func (s *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionRelease) GetReleaseInfo() *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfo {
	return s.ReleaseInfo
}

func (s *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionRelease) GetVersionChangesLevel() *string {
	return s.VersionChangesLevel
}

func (s *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionRelease) SetReleaseInfo(v *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfo) *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionRelease {
	s.ReleaseInfo = v
	return s
}

func (s *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionRelease) SetVersionChangesLevel(v string) *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionRelease {
	s.VersionChangesLevel = &v
	return s
}

func (s *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionRelease) Validate() error {
	if s.ReleaseInfo != nil {
		if err := s.ReleaseInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfo struct {
	ReleaseInfoList []*DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList `json:"ReleaseInfoList,omitempty" xml:"ReleaseInfoList,omitempty" type:"Repeated"`
}

func (s DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfo) GoString() string {
	return s.String()
}

func (s *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfo) GetReleaseInfoList() []*DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList {
	return s.ReleaseInfoList
}

func (s *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfo) SetReleaseInfoList(v []*DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList) *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfo {
	s.ReleaseInfoList = v
	return s
}

func (s *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfo) Validate() error {
	if s.ReleaseInfoList != nil {
		for _, item := range s.ReleaseInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList struct {
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Level          *string `json:"Level,omitempty" xml:"Level,omitempty"`
	ReleaseNote    *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	ReleaseNoteEn  *string `json:"ReleaseNoteEn,omitempty" xml:"ReleaseNoteEn,omitempty"`
	ReleaseVersion *string `json:"ReleaseVersion,omitempty" xml:"ReleaseVersion,omitempty"`
}

func (s DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList) GoString() string {
	return s.String()
}

func (s *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList) GetLevel() *string {
	return s.Level
}

func (s *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList) GetReleaseNote() *string {
	return s.ReleaseNote
}

func (s *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList) GetReleaseNoteEn() *string {
	return s.ReleaseNoteEn
}

func (s *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList) SetCreateTime(v string) *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList {
	s.CreateTime = &v
	return s
}

func (s *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList) SetLevel(v string) *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList {
	s.Level = &v
	return s
}

func (s *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList) SetReleaseNote(v string) *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList {
	s.ReleaseNote = &v
	return s
}

func (s *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList) SetReleaseNoteEn(v string) *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList {
	s.ReleaseNoteEn = &v
	return s
}

func (s *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList) SetReleaseVersion(v string) *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList {
	s.ReleaseVersion = &v
	return s
}

func (s *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList) Validate() error {
	return dara.Validate(s)
}

type DescribeEngineVersionResponseBodyProxyLatestMinorVersion struct {
	// The importance level of the version change. Valid values:
	//
	// 	- **0**: Normal.
	//
	// 	- **1**: Important.
	//
	// 	- **2**: Critical.
	//
	// example:
	//
	// 0
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The version number.
	//
	// example:
	//
	// 7.0.6
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	// The version evolution path from the current minor version of the instance to the latest minor version. This is consistent with the version documentation. You can refer to the version release notes for more detailed information.
	VersionRelease *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionRelease `json:"VersionRelease,omitempty" xml:"VersionRelease,omitempty" type:"Struct"`
}

func (s DescribeEngineVersionResponseBodyProxyLatestMinorVersion) String() string {
	return dara.Prettify(s)
}

func (s DescribeEngineVersionResponseBodyProxyLatestMinorVersion) GoString() string {
	return s.String()
}

func (s *DescribeEngineVersionResponseBodyProxyLatestMinorVersion) GetLevel() *string {
	return s.Level
}

func (s *DescribeEngineVersionResponseBodyProxyLatestMinorVersion) GetMinorVersion() *string {
	return s.MinorVersion
}

func (s *DescribeEngineVersionResponseBodyProxyLatestMinorVersion) GetVersionRelease() *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionRelease {
	return s.VersionRelease
}

func (s *DescribeEngineVersionResponseBodyProxyLatestMinorVersion) SetLevel(v string) *DescribeEngineVersionResponseBodyProxyLatestMinorVersion {
	s.Level = &v
	return s
}

func (s *DescribeEngineVersionResponseBodyProxyLatestMinorVersion) SetMinorVersion(v string) *DescribeEngineVersionResponseBodyProxyLatestMinorVersion {
	s.MinorVersion = &v
	return s
}

func (s *DescribeEngineVersionResponseBodyProxyLatestMinorVersion) SetVersionRelease(v *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionRelease) *DescribeEngineVersionResponseBodyProxyLatestMinorVersion {
	s.VersionRelease = v
	return s
}

func (s *DescribeEngineVersionResponseBodyProxyLatestMinorVersion) Validate() error {
	if s.VersionRelease != nil {
		if err := s.VersionRelease.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionRelease struct {
	ReleaseInfo *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfo `json:"ReleaseInfo,omitempty" xml:"ReleaseInfo,omitempty" type:"Struct"`
	// The importance level of the version upgrade (recommended upgrade level). Valid values:
	//
	// 	- 0: Normal.
	//
	// 	- 1: Important.
	//
	// 	- 2: Critical.
	//
	// example:
	//
	// 0
	VersionChangesLevel *string `json:"VersionChangesLevel,omitempty" xml:"VersionChangesLevel,omitempty"`
}

func (s DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionRelease) String() string {
	return dara.Prettify(s)
}

func (s DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionRelease) GoString() string {
	return s.String()
}

func (s *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionRelease) GetReleaseInfo() *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfo {
	return s.ReleaseInfo
}

func (s *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionRelease) GetVersionChangesLevel() *string {
	return s.VersionChangesLevel
}

func (s *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionRelease) SetReleaseInfo(v *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfo) *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionRelease {
	s.ReleaseInfo = v
	return s
}

func (s *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionRelease) SetVersionChangesLevel(v string) *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionRelease {
	s.VersionChangesLevel = &v
	return s
}

func (s *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionRelease) Validate() error {
	if s.ReleaseInfo != nil {
		if err := s.ReleaseInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfo struct {
	ReleaseInfoList []*DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList `json:"ReleaseInfoList,omitempty" xml:"ReleaseInfoList,omitempty" type:"Repeated"`
}

func (s DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfo) GoString() string {
	return s.String()
}

func (s *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfo) GetReleaseInfoList() []*DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList {
	return s.ReleaseInfoList
}

func (s *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfo) SetReleaseInfoList(v []*DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList) *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfo {
	s.ReleaseInfoList = v
	return s
}

func (s *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfo) Validate() error {
	if s.ReleaseInfoList != nil {
		for _, item := range s.ReleaseInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList struct {
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Level          *string `json:"Level,omitempty" xml:"Level,omitempty"`
	ReleaseNote    *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	ReleaseNoteEn  *string `json:"ReleaseNoteEn,omitempty" xml:"ReleaseNoteEn,omitempty"`
	ReleaseVersion *string `json:"ReleaseVersion,omitempty" xml:"ReleaseVersion,omitempty"`
}

func (s DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList) GoString() string {
	return s.String()
}

func (s *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList) GetLevel() *string {
	return s.Level
}

func (s *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList) GetReleaseNote() *string {
	return s.ReleaseNote
}

func (s *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList) GetReleaseNoteEn() *string {
	return s.ReleaseNoteEn
}

func (s *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList) SetCreateTime(v string) *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList {
	s.CreateTime = &v
	return s
}

func (s *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList) SetLevel(v string) *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList {
	s.Level = &v
	return s
}

func (s *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList) SetReleaseNote(v string) *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList {
	s.ReleaseNote = &v
	return s
}

func (s *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList) SetReleaseNoteEn(v string) *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList {
	s.ReleaseNoteEn = &v
	return s
}

func (s *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList) SetReleaseVersion(v string) *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList {
	s.ReleaseVersion = &v
	return s
}

func (s *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList) Validate() error {
	return dara.Validate(s)
}
