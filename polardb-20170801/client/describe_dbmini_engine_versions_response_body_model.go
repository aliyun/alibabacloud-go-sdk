// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBMiniEngineVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBRevisionVersionList(v []*DescribeDBMiniEngineVersionsResponseBodyDBRevisionVersionList) *DescribeDBMiniEngineVersionsResponseBody
	GetDBRevisionVersionList() []*DescribeDBMiniEngineVersionsResponseBodyDBRevisionVersionList
	SetRequestId(v string) *DescribeDBMiniEngineVersionsResponseBody
	GetRequestId() *string
}

type DescribeDBMiniEngineVersionsResponseBody struct {
	// A list of information about the versions available for an upgrade.
	DBRevisionVersionList []*DescribeDBMiniEngineVersionsResponseBodyDBRevisionVersionList `json:"DBRevisionVersionList,omitempty" xml:"DBRevisionVersionList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBMiniEngineVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBMiniEngineVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBMiniEngineVersionsResponseBody) GetDBRevisionVersionList() []*DescribeDBMiniEngineVersionsResponseBodyDBRevisionVersionList {
	return s.DBRevisionVersionList
}

func (s *DescribeDBMiniEngineVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBMiniEngineVersionsResponseBody) SetDBRevisionVersionList(v []*DescribeDBMiniEngineVersionsResponseBodyDBRevisionVersionList) *DescribeDBMiniEngineVersionsResponseBody {
	s.DBRevisionVersionList = v
	return s
}

func (s *DescribeDBMiniEngineVersionsResponseBody) SetRequestId(v string) *DescribeDBMiniEngineVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsResponseBody) Validate() error {
	if s.DBRevisionVersionList != nil {
		for _, item := range s.DBRevisionVersionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBMiniEngineVersionsResponseBodyDBRevisionVersionList struct {
	// The release notes of the version.
	//
	// example:
	//
	// ReleaseNote
	ReleaseNote *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	// The release state of the database version. Valid values:
	//
	// - **Stable**: The version is stable.
	//
	// - **Old**: The version is outdated. Upgrading to this version is not recommended.
	//
	// - **HighRisk**: The version has a critical bug. Upgrading to this version is not recommended.
	//
	// - **Beta**: The version is a beta version.
	//
	// example:
	//
	// Stable
	ReleaseType *string `json:"ReleaseType,omitempty" xml:"ReleaseType,omitempty"`
	// The code of the database engine revision version. Use this code to specify the target version for an upgrade.
	//
	// example:
	//
	// 20230707
	RevisionVersionCode *string `json:"RevisionVersionCode,omitempty" xml:"RevisionVersionCode,omitempty"`
	// The number of the database engine revision version.
	//
	// example:
	//
	// 8.0.1.1.35.1
	RevisionVersionName *string `json:"RevisionVersionName,omitempty" xml:"RevisionVersionName,omitempty"`
}

func (s DescribeDBMiniEngineVersionsResponseBodyDBRevisionVersionList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBMiniEngineVersionsResponseBodyDBRevisionVersionList) GoString() string {
	return s.String()
}

func (s *DescribeDBMiniEngineVersionsResponseBodyDBRevisionVersionList) GetReleaseNote() *string {
	return s.ReleaseNote
}

func (s *DescribeDBMiniEngineVersionsResponseBodyDBRevisionVersionList) GetReleaseType() *string {
	return s.ReleaseType
}

func (s *DescribeDBMiniEngineVersionsResponseBodyDBRevisionVersionList) GetRevisionVersionCode() *string {
	return s.RevisionVersionCode
}

func (s *DescribeDBMiniEngineVersionsResponseBodyDBRevisionVersionList) GetRevisionVersionName() *string {
	return s.RevisionVersionName
}

func (s *DescribeDBMiniEngineVersionsResponseBodyDBRevisionVersionList) SetReleaseNote(v string) *DescribeDBMiniEngineVersionsResponseBodyDBRevisionVersionList {
	s.ReleaseNote = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsResponseBodyDBRevisionVersionList) SetReleaseType(v string) *DescribeDBMiniEngineVersionsResponseBodyDBRevisionVersionList {
	s.ReleaseType = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsResponseBodyDBRevisionVersionList) SetRevisionVersionCode(v string) *DescribeDBMiniEngineVersionsResponseBodyDBRevisionVersionList {
	s.RevisionVersionCode = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsResponseBodyDBRevisionVersionList) SetRevisionVersionName(v string) *DescribeDBMiniEngineVersionsResponseBodyDBRevisionVersionList {
	s.RevisionVersionName = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsResponseBodyDBRevisionVersionList) Validate() error {
	return dara.Validate(s)
}
