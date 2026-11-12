// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationAvailableVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatestVersion(v string) *DescribeApplicationAvailableVersionResponseBody
	GetLatestVersion() *string
	SetRequestId(v string) *DescribeApplicationAvailableVersionResponseBody
	GetRequestId() *string
	SetVersions(v []*DescribeApplicationAvailableVersionResponseBodyVersions) *DescribeApplicationAvailableVersionResponseBody
	GetVersions() []*DescribeApplicationAvailableVersionResponseBodyVersions
}

type DescribeApplicationAvailableVersionResponseBody struct {
	// The latest version.
	//
	// example:
	//
	// 1.0.2
	LatestVersion *string `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3E5CD764-FCCA-5C9C-838E-20E0DE84B2AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of versions.
	//
	// example:
	//
	// []
	Versions []*DescribeApplicationAvailableVersionResponseBodyVersions `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s DescribeApplicationAvailableVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAvailableVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAvailableVersionResponseBody) GetLatestVersion() *string {
	return s.LatestVersion
}

func (s *DescribeApplicationAvailableVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationAvailableVersionResponseBody) GetVersions() []*DescribeApplicationAvailableVersionResponseBodyVersions {
	return s.Versions
}

func (s *DescribeApplicationAvailableVersionResponseBody) SetLatestVersion(v string) *DescribeApplicationAvailableVersionResponseBody {
	s.LatestVersion = &v
	return s
}

func (s *DescribeApplicationAvailableVersionResponseBody) SetRequestId(v string) *DescribeApplicationAvailableVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationAvailableVersionResponseBody) SetVersions(v []*DescribeApplicationAvailableVersionResponseBodyVersions) *DescribeApplicationAvailableVersionResponseBody {
	s.Versions = v
	return s
}

func (s *DescribeApplicationAvailableVersionResponseBody) Validate() error {
	if s.Versions != nil {
		for _, item := range s.Versions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApplicationAvailableVersionResponseBodyVersions struct {
	// The release date in RFC 3339 UTC format.
	//
	// example:
	//
	// 2019-08-01T00:00:00Z
	ReleaseDate *string `json:"ReleaseDate,omitempty" xml:"ReleaseDate,omitempty"`
	// The release notes.
	//
	// example:
	//
	// release
	ReleaseNote *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	// The release type.
	//
	// example:
	//
	// Stable
	ReleaseType *string `json:"ReleaseType,omitempty" xml:"ReleaseType,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1.0.2
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The version description.
	//
	// example:
	//
	// 1.0.2
	VersionDescription *string `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
}

func (s DescribeApplicationAvailableVersionResponseBodyVersions) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAvailableVersionResponseBodyVersions) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAvailableVersionResponseBodyVersions) GetReleaseDate() *string {
	return s.ReleaseDate
}

func (s *DescribeApplicationAvailableVersionResponseBodyVersions) GetReleaseNote() *string {
	return s.ReleaseNote
}

func (s *DescribeApplicationAvailableVersionResponseBodyVersions) GetReleaseType() *string {
	return s.ReleaseType
}

func (s *DescribeApplicationAvailableVersionResponseBodyVersions) GetVersion() *string {
	return s.Version
}

func (s *DescribeApplicationAvailableVersionResponseBodyVersions) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *DescribeApplicationAvailableVersionResponseBodyVersions) SetReleaseDate(v string) *DescribeApplicationAvailableVersionResponseBodyVersions {
	s.ReleaseDate = &v
	return s
}

func (s *DescribeApplicationAvailableVersionResponseBodyVersions) SetReleaseNote(v string) *DescribeApplicationAvailableVersionResponseBodyVersions {
	s.ReleaseNote = &v
	return s
}

func (s *DescribeApplicationAvailableVersionResponseBodyVersions) SetReleaseType(v string) *DescribeApplicationAvailableVersionResponseBodyVersions {
	s.ReleaseType = &v
	return s
}

func (s *DescribeApplicationAvailableVersionResponseBodyVersions) SetVersion(v string) *DescribeApplicationAvailableVersionResponseBodyVersions {
	s.Version = &v
	return s
}

func (s *DescribeApplicationAvailableVersionResponseBodyVersions) SetVersionDescription(v string) *DescribeApplicationAvailableVersionResponseBodyVersions {
	s.VersionDescription = &v
	return s
}

func (s *DescribeApplicationAvailableVersionResponseBodyVersions) Validate() error {
	return dara.Validate(s)
}
