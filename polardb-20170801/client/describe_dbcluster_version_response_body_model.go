// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterVersionResponseBody
	GetDBClusterId() *string
	SetDBLatestStableVersion(v string) *DescribeDBClusterVersionResponseBody
	GetDBLatestStableVersion() *string
	SetDBLatestVersion(v string) *DescribeDBClusterVersionResponseBody
	GetDBLatestVersion() *string
	SetDBMinorVersion(v string) *DescribeDBClusterVersionResponseBody
	GetDBMinorVersion() *string
	SetDBRevisionVersion(v string) *DescribeDBClusterVersionResponseBody
	GetDBRevisionVersion() *string
	SetDBRevisionVersionList(v []*DescribeDBClusterVersionResponseBodyDBRevisionVersionList) *DescribeDBClusterVersionResponseBody
	GetDBRevisionVersionList() []*DescribeDBClusterVersionResponseBodyDBRevisionVersionList
	SetDBVersion(v string) *DescribeDBClusterVersionResponseBody
	GetDBVersion() *string
	SetDBVersionStatus(v string) *DescribeDBClusterVersionResponseBody
	GetDBVersionStatus() *string
	SetIsLatestStableVersion(v string) *DescribeDBClusterVersionResponseBody
	GetIsLatestStableVersion() *string
	SetIsLatestVersion(v string) *DescribeDBClusterVersionResponseBody
	GetIsLatestVersion() *string
	SetIsProxyLatestVersion(v string) *DescribeDBClusterVersionResponseBody
	GetIsProxyLatestVersion() *string
	SetProxyLatestVersion(v string) *DescribeDBClusterVersionResponseBody
	GetProxyLatestVersion() *string
	SetProxyRevisionVersion(v string) *DescribeDBClusterVersionResponseBody
	GetProxyRevisionVersion() *string
	SetProxyRevisionVersionList(v []*DescribeDBClusterVersionResponseBodyProxyRevisionVersionList) *DescribeDBClusterVersionResponseBody
	GetProxyRevisionVersionList() []*DescribeDBClusterVersionResponseBodyProxyRevisionVersionList
	SetProxyVersionStatus(v string) *DescribeDBClusterVersionResponseBody
	GetProxyVersionStatus() *string
	SetRequestId(v string) *DescribeDBClusterVersionResponseBody
	GetRequestId() *string
}

type DescribeDBClusterVersionResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The latest stable version of PolarDB for PostgreSQL.
	//
	// example:
	//
	// 2.0.16.13.14.0
	DBLatestStableVersion *string `json:"DBLatestStableVersion,omitempty" xml:"DBLatestStableVersion,omitempty"`
	// The latest version of the database kernel engine.
	//
	// example:
	//
	// 8.0.1.1.16
	DBLatestVersion *string `json:"DBLatestVersion,omitempty" xml:"DBLatestVersion,omitempty"`
	// The minor version number of the database engine.
	//
	// 	- If `DBVersion` is **8.0**, valid values:
	//
	//     	- **8.0.2**
	//
	//     	- **8.0.1**
	//
	// 	- If `DBVersion` is **5.7**, the value is **5.7.28**.
	//
	// 	- If `DBVersion` is **5.6**, the value is **5.6.16**.
	//
	// example:
	//
	// 8.0.1
	DBMinorVersion *string `json:"DBMinorVersion,omitempty" xml:"DBMinorVersion,omitempty"`
	// The Milvus version number of the database engine.
	//
	// > For PolarDB for MySQL 5.6 clusters, only the `Milvus version` information with a release date later than August 31, 2020 is returned. Otherwise, this parameter is empty. For more information about the minor engine versions of PolarDB for MySQL clusters, see [Release notes](https://help.aliyun.com/document_detail/423884.html).
	//
	// example:
	//
	// 8.0.1.1.7
	DBRevisionVersion *string `json:"DBRevisionVersion,omitempty" xml:"DBRevisionVersion,omitempty"`
	// The list of available upgrade version information.
	DBRevisionVersionList []*DescribeDBClusterVersionResponseBodyDBRevisionVersionList `json:"DBRevisionVersionList,omitempty" xml:"DBRevisionVersionList,omitempty" type:"Repeated"`
	// The major version number of the database engine. Valid values:
	//
	// 	- **8.0**
	//
	// 	- **5.7**
	//
	// 	- **5.6**
	//
	// example:
	//
	// 8.0
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The status of the current database minor version. Valid values:
	//
	// 	- **Stable**: The current version is stable.
	//
	// 	- **Old**: The current version is outdated. Upgrade to the latest version.
	//
	// 	- **HighRisk**: The current version has critical bugs. Upgrade to the latest version immediately.
	//
	// 	- **Beta**: The current version is a beta version.
	//
	// >For more information about how to upgrade the database minor version, see [Version upgrade](https://help.aliyun.com/document_detail/158572.html).
	//
	// example:
	//
	// Stable
	DBVersionStatus *string `json:"DBVersionStatus,omitempty" xml:"DBVersionStatus,omitempty"`
	// Indicates whether the current version is the latest stable version of PolarDB for PostgreSQL.
	//
	// example:
	//
	// true
	IsLatestStableVersion *string `json:"IsLatestStableVersion,omitempty" xml:"IsLatestStableVersion,omitempty"`
	// Indicates whether the current database kernel DPI engine version is the latest database engine version. Valid values:
	//
	// -  **true**
	//
	// -  **false**
	//
	// example:
	//
	// true
	IsLatestVersion *string `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
	// Indicates whether the current PolarProxy version is the latest version. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsProxyLatestVersion *string `json:"IsProxyLatestVersion,omitempty" xml:"IsProxyLatestVersion,omitempty"`
	// The latest version of PolarProxy.
	//
	// example:
	//
	// 2.4.17
	ProxyLatestVersion *string `json:"ProxyLatestVersion,omitempty" xml:"ProxyLatestVersion,omitempty"`
	// The version of PolarProxy.
	//
	// example:
	//
	// 2.4.15
	ProxyRevisionVersion *string `json:"ProxyRevisionVersion,omitempty" xml:"ProxyRevisionVersion,omitempty"`
	// The release status of the PolarProxy version. Valid values:
	//
	// - **Stable**: The current version is stable.
	//
	// - **Old**: The current version is outdated. Upgrading to this version is not recommended.
	//
	// - **HighRisk**: The current version has critical bugs. Upgrading to this version is not recommended.
	//
	// - **Beta**: The current version is a beta version.
	ProxyRevisionVersionList []*DescribeDBClusterVersionResponseBodyProxyRevisionVersionList `json:"ProxyRevisionVersionList,omitempty" xml:"ProxyRevisionVersionList,omitempty" type:"Repeated"`
	// The version status of PolarProxy. Valid values:
	//
	// 	- **Stable**: The current version is stable.
	//
	// 	- **Old**: The current version is outdated. Upgrade to the latest version.
	//
	// 	- **HighRisk**: The current version has critical bugs. Upgrade to the latest version immediately.
	//
	// 	- **Beta**: The current version is a beta version.
	//
	// >For more information about how to upgrade the PolarProxy version, see [Version upgrade](https://help.aliyun.com/document_detail/158572.html).
	//
	// example:
	//
	// Stable
	ProxyVersionStatus *string `json:"ProxyVersionStatus,omitempty" xml:"ProxyVersionStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 47921222-0D37-4133-8C0D-017DC3******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterVersionResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterVersionResponseBody) GetDBLatestStableVersion() *string {
	return s.DBLatestStableVersion
}

func (s *DescribeDBClusterVersionResponseBody) GetDBLatestVersion() *string {
	return s.DBLatestVersion
}

func (s *DescribeDBClusterVersionResponseBody) GetDBMinorVersion() *string {
	return s.DBMinorVersion
}

func (s *DescribeDBClusterVersionResponseBody) GetDBRevisionVersion() *string {
	return s.DBRevisionVersion
}

func (s *DescribeDBClusterVersionResponseBody) GetDBRevisionVersionList() []*DescribeDBClusterVersionResponseBodyDBRevisionVersionList {
	return s.DBRevisionVersionList
}

func (s *DescribeDBClusterVersionResponseBody) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDBClusterVersionResponseBody) GetDBVersionStatus() *string {
	return s.DBVersionStatus
}

func (s *DescribeDBClusterVersionResponseBody) GetIsLatestStableVersion() *string {
	return s.IsLatestStableVersion
}

func (s *DescribeDBClusterVersionResponseBody) GetIsLatestVersion() *string {
	return s.IsLatestVersion
}

func (s *DescribeDBClusterVersionResponseBody) GetIsProxyLatestVersion() *string {
	return s.IsProxyLatestVersion
}

func (s *DescribeDBClusterVersionResponseBody) GetProxyLatestVersion() *string {
	return s.ProxyLatestVersion
}

func (s *DescribeDBClusterVersionResponseBody) GetProxyRevisionVersion() *string {
	return s.ProxyRevisionVersion
}

func (s *DescribeDBClusterVersionResponseBody) GetProxyRevisionVersionList() []*DescribeDBClusterVersionResponseBodyProxyRevisionVersionList {
	return s.ProxyRevisionVersionList
}

func (s *DescribeDBClusterVersionResponseBody) GetProxyVersionStatus() *string {
	return s.ProxyVersionStatus
}

func (s *DescribeDBClusterVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterVersionResponseBody) SetDBClusterId(v string) *DescribeDBClusterVersionResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBody) SetDBLatestStableVersion(v string) *DescribeDBClusterVersionResponseBody {
	s.DBLatestStableVersion = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBody) SetDBLatestVersion(v string) *DescribeDBClusterVersionResponseBody {
	s.DBLatestVersion = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBody) SetDBMinorVersion(v string) *DescribeDBClusterVersionResponseBody {
	s.DBMinorVersion = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBody) SetDBRevisionVersion(v string) *DescribeDBClusterVersionResponseBody {
	s.DBRevisionVersion = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBody) SetDBRevisionVersionList(v []*DescribeDBClusterVersionResponseBodyDBRevisionVersionList) *DescribeDBClusterVersionResponseBody {
	s.DBRevisionVersionList = v
	return s
}

func (s *DescribeDBClusterVersionResponseBody) SetDBVersion(v string) *DescribeDBClusterVersionResponseBody {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBody) SetDBVersionStatus(v string) *DescribeDBClusterVersionResponseBody {
	s.DBVersionStatus = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBody) SetIsLatestStableVersion(v string) *DescribeDBClusterVersionResponseBody {
	s.IsLatestStableVersion = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBody) SetIsLatestVersion(v string) *DescribeDBClusterVersionResponseBody {
	s.IsLatestVersion = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBody) SetIsProxyLatestVersion(v string) *DescribeDBClusterVersionResponseBody {
	s.IsProxyLatestVersion = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBody) SetProxyLatestVersion(v string) *DescribeDBClusterVersionResponseBody {
	s.ProxyLatestVersion = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBody) SetProxyRevisionVersion(v string) *DescribeDBClusterVersionResponseBody {
	s.ProxyRevisionVersion = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBody) SetProxyRevisionVersionList(v []*DescribeDBClusterVersionResponseBodyProxyRevisionVersionList) *DescribeDBClusterVersionResponseBody {
	s.ProxyRevisionVersionList = v
	return s
}

func (s *DescribeDBClusterVersionResponseBody) SetProxyVersionStatus(v string) *DescribeDBClusterVersionResponseBody {
	s.ProxyVersionStatus = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBody) SetRequestId(v string) *DescribeDBClusterVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBody) Validate() error {
	if s.DBRevisionVersionList != nil {
		for _, item := range s.DBRevisionVersionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ProxyRevisionVersionList != nil {
		for _, item := range s.ProxyRevisionVersionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterVersionResponseBodyDBRevisionVersionList struct {
	// The release notes for the version.
	//
	// example:
	//
	// ReleaseNote
	ReleaseNote *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	// The release status of the database version. Valid values:
	//
	// 	- **Stable**: The current version is stable.
	//
	// 	- **Old**: The current version is outdated. Upgrading to this version is not recommended.
	//
	// 	- **HighRisk**: The current version has critical bugs. Upgrading to this version is not recommended.
	//
	// 	- **Beta**: The current version is a beta version.
	//
	// example:
	//
	// Stable
	ReleaseType *string `json:"ReleaseType,omitempty" xml:"ReleaseType,omitempty"`
	// The revision version code of the database engine, which is used to specify the target version for the upgrade.
	//
	// example:
	//
	// 20230707
	RevisionVersionCode *string `json:"RevisionVersionCode,omitempty" xml:"RevisionVersionCode,omitempty"`
	// The revision version number of the database engine.
	//
	// example:
	//
	// 8.0.1.1.35.1
	RevisionVersionName *string `json:"RevisionVersionName,omitempty" xml:"RevisionVersionName,omitempty"`
}

func (s DescribeDBClusterVersionResponseBodyDBRevisionVersionList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterVersionResponseBodyDBRevisionVersionList) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterVersionResponseBodyDBRevisionVersionList) GetReleaseNote() *string {
	return s.ReleaseNote
}

func (s *DescribeDBClusterVersionResponseBodyDBRevisionVersionList) GetReleaseType() *string {
	return s.ReleaseType
}

func (s *DescribeDBClusterVersionResponseBodyDBRevisionVersionList) GetRevisionVersionCode() *string {
	return s.RevisionVersionCode
}

func (s *DescribeDBClusterVersionResponseBodyDBRevisionVersionList) GetRevisionVersionName() *string {
	return s.RevisionVersionName
}

func (s *DescribeDBClusterVersionResponseBodyDBRevisionVersionList) SetReleaseNote(v string) *DescribeDBClusterVersionResponseBodyDBRevisionVersionList {
	s.ReleaseNote = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBodyDBRevisionVersionList) SetReleaseType(v string) *DescribeDBClusterVersionResponseBodyDBRevisionVersionList {
	s.ReleaseType = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBodyDBRevisionVersionList) SetRevisionVersionCode(v string) *DescribeDBClusterVersionResponseBodyDBRevisionVersionList {
	s.RevisionVersionCode = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBodyDBRevisionVersionList) SetRevisionVersionName(v string) *DescribeDBClusterVersionResponseBodyDBRevisionVersionList {
	s.RevisionVersionName = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBodyDBRevisionVersionList) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterVersionResponseBodyProxyRevisionVersionList struct {
	// The release notes for the version.
	//
	// example:
	//
	// ReleaseNote
	ReleaseNote *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	// The release type. Valid values:
	//
	// 	- **LTS**: Long-term support version.
	//
	// 	- **BETA**: Preview version.
	//
	// example:
	//
	// LTS
	ReleaseType *string `json:"ReleaseType,omitempty" xml:"ReleaseType,omitempty"`
	// The revision version code of the PolarProxy engine, which is used to specify the target version for the upgrade.
	//
	// example:
	//
	// 20230707
	RevisionVersionCode *string `json:"RevisionVersionCode,omitempty" xml:"RevisionVersionCode,omitempty"`
	// The revision version number of the PolarProxy engine.
	//
	// example:
	//
	// 2.8.24
	RevisionVersionName *string `json:"RevisionVersionName,omitempty" xml:"RevisionVersionName,omitempty"`
}

func (s DescribeDBClusterVersionResponseBodyProxyRevisionVersionList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterVersionResponseBodyProxyRevisionVersionList) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterVersionResponseBodyProxyRevisionVersionList) GetReleaseNote() *string {
	return s.ReleaseNote
}

func (s *DescribeDBClusterVersionResponseBodyProxyRevisionVersionList) GetReleaseType() *string {
	return s.ReleaseType
}

func (s *DescribeDBClusterVersionResponseBodyProxyRevisionVersionList) GetRevisionVersionCode() *string {
	return s.RevisionVersionCode
}

func (s *DescribeDBClusterVersionResponseBodyProxyRevisionVersionList) GetRevisionVersionName() *string {
	return s.RevisionVersionName
}

func (s *DescribeDBClusterVersionResponseBodyProxyRevisionVersionList) SetReleaseNote(v string) *DescribeDBClusterVersionResponseBodyProxyRevisionVersionList {
	s.ReleaseNote = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBodyProxyRevisionVersionList) SetReleaseType(v string) *DescribeDBClusterVersionResponseBodyProxyRevisionVersionList {
	s.ReleaseType = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBodyProxyRevisionVersionList) SetRevisionVersionCode(v string) *DescribeDBClusterVersionResponseBodyProxyRevisionVersionList {
	s.RevisionVersionCode = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBodyProxyRevisionVersionList) SetRevisionVersionName(v string) *DescribeDBClusterVersionResponseBodyProxyRevisionVersionList {
	s.RevisionVersionName = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBodyProxyRevisionVersionList) Validate() error {
	return dara.Validate(s)
}
