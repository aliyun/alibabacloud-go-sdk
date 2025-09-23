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
	DBRevisionVersionList []*DescribeDBMiniEngineVersionsResponseBodyDBRevisionVersionList `json:"DBRevisionVersionList,omitempty" xml:"DBRevisionVersionList,omitempty" type:"Repeated"`
	// Id of the request
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
	return dara.Validate(s)
}

type DescribeDBMiniEngineVersionsResponseBodyDBRevisionVersionList struct {
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
