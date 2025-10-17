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
	// The ID of cluster.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The latest version of the database engine.
	//
	// example:
	//
	// 8.0.1.1.16
	DBLatestVersion *string `json:"DBLatestVersion,omitempty" xml:"DBLatestVersion,omitempty"`
	// The minor version of the database engine.
	//
	// - If DBVersion is 8.0, the valid values of this parameter are:
	//
	//   - 8.0.2
	//
	//   - 8.0.1
	//
	// - If DBVersion is 5.7, set the value of this parameter to 5.7.28.
	//
	// - If DBVersion is 5.6, the value of this parameter is 5.6.16.
	//
	// example:
	//
	// 8.0.1
	DBMinorVersion *string `json:"DBMinorVersion,omitempty" xml:"DBMinorVersion,omitempty"`
	// The revision version of the database engine.
	//
	// >For a cluster of the PolarDB for MySQL 5.6, the DBRevisionVersion parameter returns the revision version information only if the Revision Version is released later than August 31, 2020. Otherwise, this parameter returns an empty value.
	//
	// example:
	//
	// 8.0.1.1.7
	DBRevisionVersion *string `json:"DBRevisionVersion,omitempty" xml:"DBRevisionVersion,omitempty"`
	// The versions to which the cluster can be upgraded.
	DBRevisionVersionList []*DescribeDBClusterVersionResponseBodyDBRevisionVersionList `json:"DBRevisionVersionList,omitempty" xml:"DBRevisionVersionList,omitempty" type:"Repeated"`
	// The version of the database engine. Valid values:
	//
	// - 5.6
	//
	// - 5.7
	//
	// - 8.0
	//
	// example:
	//
	// 8.0
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The status of the minor version. Valid values:
	//
	// 	- **Stable**: The minor version is stable.
	//
	// 	- **Old**: The minor version is outdated. We recommend that you upgrade the cluster to the latest version.
	//
	// 	- **HighRisk**: The minor version has critical defects. We recommend that you immediately update the cluster to the latest minor version.
	//
	// >  For more information about how to update the minor version, see [Minor version update](https://help.aliyun.com/document_detail/158572.html).
	//
	// example:
	//
	// Stable
	DBVersionStatus *string `json:"DBVersionStatus,omitempty" xml:"DBVersionStatus,omitempty"`
	// Indicates whether the kernel is of the latest version. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// true
	IsLatestVersion *string `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
	// Indicates whether PolarProxy uses the latest version. Valid values:
	//
	// - true
	//
	// - false
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
	// The revision version of the database engine.
	//
	// example:
	//
	// 2.4.15
	ProxyRevisionVersion *string `json:"ProxyRevisionVersion,omitempty" xml:"ProxyRevisionVersion,omitempty"`
	// The release status of the PolarProxy version. Valid values:
	//
	// 	- **Stable**: The PolarProxy revision version is stable.
	//
	// 	- **Old**: The PolarProxy revision version is outdated. We recommend that you do not update the PolarProxy to this revision version.
	//
	// 	- **HighRisk**: The PolarProxy revision version has critical defects. We recommend that you do not update the PolarProxy to this revision version.
	//
	// 	- **Beta**: The PolarProxy revision version is a Beta version.
	ProxyRevisionVersionList []*DescribeDBClusterVersionResponseBodyProxyRevisionVersionList `json:"ProxyRevisionVersionList,omitempty" xml:"ProxyRevisionVersionList,omitempty" type:"Repeated"`
	// The status of PolarProxy. Valid values:
	//
	// - Stable: The minor version is stable.
	//
	// - Old: The minor version is outdated. We recommend that you upgrade the cluster to the latest version.
	//
	// - HighRisk: The minor version has critical defects. We recommend that you immediately upgrade the cluster to the latest version.
	//
	// - Beta: The minor version is a beta version.
	//
	// example:
	//
	// Stable
	ProxyVersionStatus *string `json:"ProxyVersionStatus,omitempty" xml:"ProxyVersionStatus,omitempty"`
	// The ID of the request.
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
	// The release notes for the database engine revision version.
	//
	// example:
	//
	// ReleaseNote
	ReleaseNote *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	// The release status of the database engine revision version. Valid values:
	//
	// 	- **Stable**: The database engine revision version is stable.
	//
	// 	- **Old**: The database engine revision version is outdated. We recommend that you do not update the database engine to this revision version.
	//
	// 	- **HighRisk**: The database engine revision version has critical defects. We recommend that you do not update the database engine to this revision version.
	//
	// 	- **Beta**: The database engine revision version is a Beta version.
	//
	// example:
	//
	// Stable
	ReleaseType *string `json:"ReleaseType,omitempty" xml:"ReleaseType,omitempty"`
	// The code of the database engine revision version. You can use the code to specify the database engine revision version.
	//
	// example:
	//
	// 20230707
	RevisionVersionCode *string `json:"RevisionVersionCode,omitempty" xml:"RevisionVersionCode,omitempty"`
	// The database engine revision version number.
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
	// The release notes for the PolarProxy revision version.
	//
	// example:
	//
	// ReleaseNote
	ReleaseNote *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	// The release type. Valid values:
	//
	// 	- **LTS**: a long-term version
	//
	// 	- **BETA**: a preview version
	//
	// example:
	//
	// LTS
	ReleaseType *string `json:"ReleaseType,omitempty" xml:"ReleaseType,omitempty"`
	// The PolarProxy revision version code. You can use this code to specify the PolarProxy revision version.
	//
	// example:
	//
	// 20230707
	RevisionVersionCode *string `json:"RevisionVersionCode,omitempty" xml:"RevisionVersionCode,omitempty"`
	// The PolarProxy revision version number.
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
