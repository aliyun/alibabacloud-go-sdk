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
	// The latest minor version to which the instance can be updated.
	DBLatestMinorVersion *DescribeEngineVersionResponseBodyDBLatestMinorVersion `json:"DBLatestMinorVersion,omitempty" xml:"DBLatestMinorVersion,omitempty" type:"Struct"`
	// The release notes for the minor version of the instance, including the release date, minor version number, release type such as new feature, and description.
	//
	// example:
	//
	// {\\"releaseInfo\\":{\\"createTime\\":\\"2021-07-27\\",\\"level\\":1,\\"releaseVersion\\":\\"0.5.4\\",\\"releaseNote\\":\\"功能更新：增强稳定性。\\"}],\\"versionChangesLevel\\":2}
	DBVersionRelease *string `json:"DBVersionRelease,omitempty" xml:"DBVersionRelease,omitempty"`
	// Indicates whether the instance major version can be upgraded. Valid values:
	//
	// 	- **true**: The major version can be upgraded.
	//
	// 	- **false**: The major version is the latest version and cannot be upgraded.
	//
	// >  To upgrade the major version, call the [ModifyInstanceMajorVersion](https://help.aliyun.com/document_detail/473776.html) operation.
	//
	// example:
	//
	// true
	EnableUpgradeMajorVersion *bool `json:"EnableUpgradeMajorVersion,omitempty" xml:"EnableUpgradeMajorVersion,omitempty"`
	// Indicates whether the instance minor version can be updated. Valid values:
	//
	// 	- **true**: The minor version can be updated.
	//
	// 	- **false**: The minor version is the latest version and cannot be updated.
	//
	// >  To update the minor version, call the [ModifyInstanceMinorVersion](https://help.aliyun.com/document_detail/473777.html) operation.
	//
	// example:
	//
	// true
	EnableUpgradeMinorVersion *bool `json:"EnableUpgradeMinorVersion,omitempty" xml:"EnableUpgradeMinorVersion,omitempty"`
	// The database engine. Valid values: **redis*	- and **memcache**.
	//
	// example:
	//
	// redis
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// Indicates whether automatic minor version update is enabled. Valid values:
	//
	// 	- **0**: Automatic minor version update is disabled.
	//
	// 	- **1**: Automatic minor version update is enabled.
	//
	// example:
	//
	// 0
	IsAutoUpgradeOpen *string `json:"IsAutoUpgradeOpen,omitempty" xml:"IsAutoUpgradeOpen,omitempty"`
	// Indicates whether the instance uses the latest minor version. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	IsLatestVersion *bool `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
	// Indicates whether Transport Layer Security (TLS) is enabled. Valid values:
	//
	// 	- **1**: TLS is enabled.
	//
	// 	- **0**: TLS is disabled.
	//
	// example:
	//
	// 1
	IsNewSSLMode *string `json:"IsNewSSLMode,omitempty" xml:"IsNewSSLMode,omitempty"`
	// Indicates whether the NGLB mode is enabled. Valid values:
	//
	// 	- **0**: The NGLB mode is disabled.
	//
	// 	- **1**: The NGLB mode is enabled.
	//
	// example:
	//
	// 1
	IsOpenNGLB *string `json:"IsOpenNGLB,omitempty" xml:"IsOpenNGLB,omitempty"`
	// Indicates whether the instance runs a Redis version.
	//
	// example:
	//
	// 1
	IsRedisCompatibleVersion *string `json:"IsRedisCompatibleVersion,omitempty" xml:"IsRedisCompatibleVersion,omitempty"`
	// Indicates whether SSL is enabled. Valid values:
	//
	// 	- **1**: SSL is enabled.
	//
	// 	- **0**: TLS is disabled.
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
	// The latest minor version to which the proxy node can be updated.
	ProxyLatestMinorVersion *DescribeEngineVersionResponseBodyProxyLatestMinorVersion `json:"ProxyLatestMinorVersion,omitempty" xml:"ProxyLatestMinorVersion,omitempty" type:"Struct"`
	// The current minor version of the proxy node.
	//
	// >  This parameter is returned only for cluster and read/write splitting instances.
	//
	// example:
	//
	// 6.6.0
	ProxyMinorVersion *string `json:"ProxyMinorVersion,omitempty" xml:"ProxyMinorVersion,omitempty"`
	// The release notes for the minor version of proxy nodes. The release notes include the release date, minor version number, release type such as new feature, and description.
	//
	// >  This parameter is returned only for cluster and read/write splitting instances.
	//
	// example:
	//
	// {\\"releaseInfo\\":[{\\"createTime\\":\\"2021-06-08\\",\\"level\\":0,\\"releaseVersion\\":\\"6.6.2\\",\\"releaseNote\\":\\"新特性：增加对部分内部命令的支持。\\"}],\\"versionChangesLevel\\":2}
	ProxyVersionRelease *string `json:"ProxyVersionRelease,omitempty" xml:"ProxyVersionRelease,omitempty"`
	// The request ID.
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
	return dara.Validate(s)
}

type DescribeEngineVersionResponseBodyDBLatestMinorVersion struct {
	// The version update level. Valid values:
	//
	// 	- **0**: regular
	//
	// 	- **1**: recommended
	//
	// 	- **2**: critical
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
	// The version update path from the current minor version to the latest minor version of the instance, which is consistent with the version documentation. For more detailed information, you can directly refer to the release notes.
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
	return dara.Validate(s)
}

type DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionRelease struct {
	// The information about the minor versions.
	ReleaseInfo *DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfo `json:"ReleaseInfo,omitempty" xml:"ReleaseInfo,omitempty" type:"Struct"`
	// The version update level, which indicates how strongly the update is recommended. Valid values:
	//
	// 	- 0: regular
	//
	// 	- 1: recommended
	//
	// 	- 2: critical
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type DescribeEngineVersionResponseBodyDBLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList struct {
	// The creation time of the instance.
	//
	// example:
	//
	// 2022-11-21T13:28Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The importance level.
	//
	// example:
	//
	// 0
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The release notes.
	//
	// example:
	//
	// netbank1022
	ReleaseNote *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	// The description of the minor versions to which the instance can be updated.
	//
	// example:
	//
	// ReleaseNoteEn
	ReleaseNoteEn *string `json:"ReleaseNoteEn,omitempty" xml:"ReleaseNoteEn,omitempty"`
	// The release version of EMR.
	//
	// example:
	//
	// EMR-5.7.0
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
	// The version update level. Valid values:
	//
	// 	- **0**: regular
	//
	// 	- **1**: recommended
	//
	// 	- **2**: critical
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
	// The version update path from the current minor version to the latest minor version of the instance, which is consistent with the version documentation. For more detailed information, you can directly refer to the release notes.
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
	return dara.Validate(s)
}

type DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionRelease struct {
	// The information about the minor versions.
	ReleaseInfo *DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfo `json:"ReleaseInfo,omitempty" xml:"ReleaseInfo,omitempty" type:"Struct"`
	// The version update level, which indicates how strongly the update is recommended. Valid values:
	//
	// 	- 0: regular
	//
	// 	- 1: recommended
	//
	// 	- 2: critical
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type DescribeEngineVersionResponseBodyProxyLatestMinorVersionVersionReleaseReleaseInfoReleaseInfoList struct {
	// The time when the version was released.
	//
	// example:
	//
	// 2022-08-23T14:26:20Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The version update level. Valid values:
	//
	// 	- **0**: regular
	//
	// 	- **1**: recommended
	//
	// 	- **2**: critical
	//
	// example:
	//
	// 0
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The release notes.
	//
	// example:
	//
	// x x x x
	ReleaseNote *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	// The description of the minor versions to which the instance can be updated.
	//
	// example:
	//
	// ReleaseNoteEn
	ReleaseNoteEn *string `json:"ReleaseNoteEn,omitempty" xml:"ReleaseNoteEn,omitempty"`
	// The release version of EMR.
	//
	// example:
	//
	// EMR-5.9.1
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
