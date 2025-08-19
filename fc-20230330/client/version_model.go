// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVersion interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *Version
	GetCreatedTime() *string
	SetDescription(v string) *Version
	GetDescription() *string
	SetLastModifiedTime(v string) *Version
	GetLastModifiedTime() *string
	SetVersionId(v string) *Version
	GetVersionId() *string
}

type Version struct {
	// example:
	//
	// 2006-01-02T15:04:05Z07:00
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// my version
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 2006-01-02T15:04:05Z07:00
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// example:
	//
	// 1
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s Version) String() string {
	return dara.Prettify(s)
}

func (s Version) GoString() string {
	return s.String()
}

func (s *Version) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *Version) GetDescription() *string {
	return s.Description
}

func (s *Version) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *Version) GetVersionId() *string {
	return s.VersionId
}

func (s *Version) SetCreatedTime(v string) *Version {
	s.CreatedTime = &v
	return s
}

func (s *Version) SetDescription(v string) *Version {
	s.Description = &v
	return s
}

func (s *Version) SetLastModifiedTime(v string) *Version {
	s.LastModifiedTime = &v
	return s
}

func (s *Version) SetVersionId(v string) *Version {
	s.VersionId = &v
	return s
}

func (s *Version) Validate() error {
	return dara.Validate(s)
}
