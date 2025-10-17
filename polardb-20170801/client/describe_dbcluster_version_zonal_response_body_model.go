// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterVersionZonalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterVersionZonalResponseBody
	GetDBClusterId() *string
	SetDBLatestVersion(v string) *DescribeDBClusterVersionZonalResponseBody
	GetDBLatestVersion() *string
	SetDBMinorVersion(v string) *DescribeDBClusterVersionZonalResponseBody
	GetDBMinorVersion() *string
	SetDBRevisionVersion(v string) *DescribeDBClusterVersionZonalResponseBody
	GetDBRevisionVersion() *string
	SetDBRevisionVersionList(v []*DescribeDBClusterVersionZonalResponseBodyDBRevisionVersionList) *DescribeDBClusterVersionZonalResponseBody
	GetDBRevisionVersionList() []*DescribeDBClusterVersionZonalResponseBodyDBRevisionVersionList
	SetDBVersion(v string) *DescribeDBClusterVersionZonalResponseBody
	GetDBVersion() *string
	SetDBVersionStatus(v string) *DescribeDBClusterVersionZonalResponseBody
	GetDBVersionStatus() *string
	SetIsLatestVersion(v string) *DescribeDBClusterVersionZonalResponseBody
	GetIsLatestVersion() *string
	SetIsProxyLatestVersion(v string) *DescribeDBClusterVersionZonalResponseBody
	GetIsProxyLatestVersion() *string
	SetProxyLatestVersion(v string) *DescribeDBClusterVersionZonalResponseBody
	GetProxyLatestVersion() *string
	SetProxyRevisionVersion(v string) *DescribeDBClusterVersionZonalResponseBody
	GetProxyRevisionVersion() *string
	SetProxyRevisionVersionList(v []*DescribeDBClusterVersionZonalResponseBodyProxyRevisionVersionList) *DescribeDBClusterVersionZonalResponseBody
	GetProxyRevisionVersionList() []*DescribeDBClusterVersionZonalResponseBodyProxyRevisionVersionList
	SetProxyVersionStatus(v string) *DescribeDBClusterVersionZonalResponseBody
	GetProxyVersionStatus() *string
	SetRequestId(v string) *DescribeDBClusterVersionZonalResponseBody
	GetRequestId() *string
}

type DescribeDBClusterVersionZonalResponseBody struct {
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// 8.0.1.1.16
	DBLatestVersion *string `json:"DBLatestVersion,omitempty" xml:"DBLatestVersion,omitempty"`
	// example:
	//
	// 8.0.1
	DBMinorVersion *string `json:"DBMinorVersion,omitempty" xml:"DBMinorVersion,omitempty"`
	// example:
	//
	// 8.0.1.1.7
	DBRevisionVersion     *string                                                           `json:"DBRevisionVersion,omitempty" xml:"DBRevisionVersion,omitempty"`
	DBRevisionVersionList []*DescribeDBClusterVersionZonalResponseBodyDBRevisionVersionList `json:"DBRevisionVersionList,omitempty" xml:"DBRevisionVersionList,omitempty" type:"Repeated"`
	// example:
	//
	// 5.6
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// example:
	//
	// Stable
	DBVersionStatus *string `json:"DBVersionStatus,omitempty" xml:"DBVersionStatus,omitempty"`
	// example:
	//
	// true
	IsLatestVersion *string `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
	// example:
	//
	// true
	IsProxyLatestVersion *string `json:"IsProxyLatestVersion,omitempty" xml:"IsProxyLatestVersion,omitempty"`
	// example:
	//
	// 2.4.17
	ProxyLatestVersion *string `json:"ProxyLatestVersion,omitempty" xml:"ProxyLatestVersion,omitempty"`
	// example:
	//
	// 2.4.15
	ProxyRevisionVersion     *string                                                              `json:"ProxyRevisionVersion,omitempty" xml:"ProxyRevisionVersion,omitempty"`
	ProxyRevisionVersionList []*DescribeDBClusterVersionZonalResponseBodyProxyRevisionVersionList `json:"ProxyRevisionVersionList,omitempty" xml:"ProxyRevisionVersionList,omitempty" type:"Repeated"`
	// example:
	//
	// Stable
	ProxyVersionStatus *string `json:"ProxyVersionStatus,omitempty" xml:"ProxyVersionStatus,omitempty"`
	// example:
	//
	// 47921222-0D37-4133-8C0D-017DC3******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterVersionZonalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterVersionZonalResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterVersionZonalResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterVersionZonalResponseBody) GetDBLatestVersion() *string {
	return s.DBLatestVersion
}

func (s *DescribeDBClusterVersionZonalResponseBody) GetDBMinorVersion() *string {
	return s.DBMinorVersion
}

func (s *DescribeDBClusterVersionZonalResponseBody) GetDBRevisionVersion() *string {
	return s.DBRevisionVersion
}

func (s *DescribeDBClusterVersionZonalResponseBody) GetDBRevisionVersionList() []*DescribeDBClusterVersionZonalResponseBodyDBRevisionVersionList {
	return s.DBRevisionVersionList
}

func (s *DescribeDBClusterVersionZonalResponseBody) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDBClusterVersionZonalResponseBody) GetDBVersionStatus() *string {
	return s.DBVersionStatus
}

func (s *DescribeDBClusterVersionZonalResponseBody) GetIsLatestVersion() *string {
	return s.IsLatestVersion
}

func (s *DescribeDBClusterVersionZonalResponseBody) GetIsProxyLatestVersion() *string {
	return s.IsProxyLatestVersion
}

func (s *DescribeDBClusterVersionZonalResponseBody) GetProxyLatestVersion() *string {
	return s.ProxyLatestVersion
}

func (s *DescribeDBClusterVersionZonalResponseBody) GetProxyRevisionVersion() *string {
	return s.ProxyRevisionVersion
}

func (s *DescribeDBClusterVersionZonalResponseBody) GetProxyRevisionVersionList() []*DescribeDBClusterVersionZonalResponseBodyProxyRevisionVersionList {
	return s.ProxyRevisionVersionList
}

func (s *DescribeDBClusterVersionZonalResponseBody) GetProxyVersionStatus() *string {
	return s.ProxyVersionStatus
}

func (s *DescribeDBClusterVersionZonalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterVersionZonalResponseBody) SetDBClusterId(v string) *DescribeDBClusterVersionZonalResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterVersionZonalResponseBody) SetDBLatestVersion(v string) *DescribeDBClusterVersionZonalResponseBody {
	s.DBLatestVersion = &v
	return s
}

func (s *DescribeDBClusterVersionZonalResponseBody) SetDBMinorVersion(v string) *DescribeDBClusterVersionZonalResponseBody {
	s.DBMinorVersion = &v
	return s
}

func (s *DescribeDBClusterVersionZonalResponseBody) SetDBRevisionVersion(v string) *DescribeDBClusterVersionZonalResponseBody {
	s.DBRevisionVersion = &v
	return s
}

func (s *DescribeDBClusterVersionZonalResponseBody) SetDBRevisionVersionList(v []*DescribeDBClusterVersionZonalResponseBodyDBRevisionVersionList) *DescribeDBClusterVersionZonalResponseBody {
	s.DBRevisionVersionList = v
	return s
}

func (s *DescribeDBClusterVersionZonalResponseBody) SetDBVersion(v string) *DescribeDBClusterVersionZonalResponseBody {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClusterVersionZonalResponseBody) SetDBVersionStatus(v string) *DescribeDBClusterVersionZonalResponseBody {
	s.DBVersionStatus = &v
	return s
}

func (s *DescribeDBClusterVersionZonalResponseBody) SetIsLatestVersion(v string) *DescribeDBClusterVersionZonalResponseBody {
	s.IsLatestVersion = &v
	return s
}

func (s *DescribeDBClusterVersionZonalResponseBody) SetIsProxyLatestVersion(v string) *DescribeDBClusterVersionZonalResponseBody {
	s.IsProxyLatestVersion = &v
	return s
}

func (s *DescribeDBClusterVersionZonalResponseBody) SetProxyLatestVersion(v string) *DescribeDBClusterVersionZonalResponseBody {
	s.ProxyLatestVersion = &v
	return s
}

func (s *DescribeDBClusterVersionZonalResponseBody) SetProxyRevisionVersion(v string) *DescribeDBClusterVersionZonalResponseBody {
	s.ProxyRevisionVersion = &v
	return s
}

func (s *DescribeDBClusterVersionZonalResponseBody) SetProxyRevisionVersionList(v []*DescribeDBClusterVersionZonalResponseBodyProxyRevisionVersionList) *DescribeDBClusterVersionZonalResponseBody {
	s.ProxyRevisionVersionList = v
	return s
}

func (s *DescribeDBClusterVersionZonalResponseBody) SetProxyVersionStatus(v string) *DescribeDBClusterVersionZonalResponseBody {
	s.ProxyVersionStatus = &v
	return s
}

func (s *DescribeDBClusterVersionZonalResponseBody) SetRequestId(v string) *DescribeDBClusterVersionZonalResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterVersionZonalResponseBody) Validate() error {
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

type DescribeDBClusterVersionZonalResponseBodyDBRevisionVersionList struct {
	// example:
	//
	// ReleaseNote
	ReleaseNote *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	// example:
	//
	// Stable
	ReleaseType *string `json:"ReleaseType,omitempty" xml:"ReleaseType,omitempty"`
	// example:
	//
	// 20230707
	RevisionVersionCode *string `json:"RevisionVersionCode,omitempty" xml:"RevisionVersionCode,omitempty"`
	// example:
	//
	// 8.0.1.1.35.1
	RevisionVersionName *string `json:"RevisionVersionName,omitempty" xml:"RevisionVersionName,omitempty"`
}

func (s DescribeDBClusterVersionZonalResponseBodyDBRevisionVersionList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterVersionZonalResponseBodyDBRevisionVersionList) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterVersionZonalResponseBodyDBRevisionVersionList) GetReleaseNote() *string {
	return s.ReleaseNote
}

func (s *DescribeDBClusterVersionZonalResponseBodyDBRevisionVersionList) GetReleaseType() *string {
	return s.ReleaseType
}

func (s *DescribeDBClusterVersionZonalResponseBodyDBRevisionVersionList) GetRevisionVersionCode() *string {
	return s.RevisionVersionCode
}

func (s *DescribeDBClusterVersionZonalResponseBodyDBRevisionVersionList) GetRevisionVersionName() *string {
	return s.RevisionVersionName
}

func (s *DescribeDBClusterVersionZonalResponseBodyDBRevisionVersionList) SetReleaseNote(v string) *DescribeDBClusterVersionZonalResponseBodyDBRevisionVersionList {
	s.ReleaseNote = &v
	return s
}

func (s *DescribeDBClusterVersionZonalResponseBodyDBRevisionVersionList) SetReleaseType(v string) *DescribeDBClusterVersionZonalResponseBodyDBRevisionVersionList {
	s.ReleaseType = &v
	return s
}

func (s *DescribeDBClusterVersionZonalResponseBodyDBRevisionVersionList) SetRevisionVersionCode(v string) *DescribeDBClusterVersionZonalResponseBodyDBRevisionVersionList {
	s.RevisionVersionCode = &v
	return s
}

func (s *DescribeDBClusterVersionZonalResponseBodyDBRevisionVersionList) SetRevisionVersionName(v string) *DescribeDBClusterVersionZonalResponseBodyDBRevisionVersionList {
	s.RevisionVersionName = &v
	return s
}

func (s *DescribeDBClusterVersionZonalResponseBodyDBRevisionVersionList) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterVersionZonalResponseBodyProxyRevisionVersionList struct {
	// example:
	//
	// ReleaseNote
	ReleaseNote *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	// example:
	//
	// LTS
	ReleaseType *string `json:"ReleaseType,omitempty" xml:"ReleaseType,omitempty"`
	// example:
	//
	// 20230707
	RevisionVersionCode *string `json:"RevisionVersionCode,omitempty" xml:"RevisionVersionCode,omitempty"`
	// example:
	//
	// 2.8.24
	RevisionVersionName *string `json:"RevisionVersionName,omitempty" xml:"RevisionVersionName,omitempty"`
}

func (s DescribeDBClusterVersionZonalResponseBodyProxyRevisionVersionList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterVersionZonalResponseBodyProxyRevisionVersionList) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterVersionZonalResponseBodyProxyRevisionVersionList) GetReleaseNote() *string {
	return s.ReleaseNote
}

func (s *DescribeDBClusterVersionZonalResponseBodyProxyRevisionVersionList) GetReleaseType() *string {
	return s.ReleaseType
}

func (s *DescribeDBClusterVersionZonalResponseBodyProxyRevisionVersionList) GetRevisionVersionCode() *string {
	return s.RevisionVersionCode
}

func (s *DescribeDBClusterVersionZonalResponseBodyProxyRevisionVersionList) GetRevisionVersionName() *string {
	return s.RevisionVersionName
}

func (s *DescribeDBClusterVersionZonalResponseBodyProxyRevisionVersionList) SetReleaseNote(v string) *DescribeDBClusterVersionZonalResponseBodyProxyRevisionVersionList {
	s.ReleaseNote = &v
	return s
}

func (s *DescribeDBClusterVersionZonalResponseBodyProxyRevisionVersionList) SetReleaseType(v string) *DescribeDBClusterVersionZonalResponseBodyProxyRevisionVersionList {
	s.ReleaseType = &v
	return s
}

func (s *DescribeDBClusterVersionZonalResponseBodyProxyRevisionVersionList) SetRevisionVersionCode(v string) *DescribeDBClusterVersionZonalResponseBodyProxyRevisionVersionList {
	s.RevisionVersionCode = &v
	return s
}

func (s *DescribeDBClusterVersionZonalResponseBodyProxyRevisionVersionList) SetRevisionVersionName(v string) *DescribeDBClusterVersionZonalResponseBodyProxyRevisionVersionList {
	s.RevisionVersionName = &v
	return s
}

func (s *DescribeDBClusterVersionZonalResponseBodyProxyRevisionVersionList) Validate() error {
	return dara.Validate(s)
}
