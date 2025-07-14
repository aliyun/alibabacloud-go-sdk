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
	SetImage(v string) *Version
	GetImage() *string
	SetLastModifiedTime(v string) *Version
	GetLastModifiedTime() *string
	SetRequestId(v string) *Version
	GetRequestId() *string
	SetVersionId(v string) *Version
	GetVersionId() *string
	SetWeight(v float64) *Version
	GetWeight() *float64
}

type Version struct {
	CreatedTime      *string  `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	Description      *string  `json:"description,omitempty" xml:"description,omitempty"`
	Image            *string  `json:"image,omitempty" xml:"image,omitempty"`
	LastModifiedTime *string  `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	RequestId        *string  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	VersionId        *string  `json:"versionId,omitempty" xml:"versionId,omitempty"`
	Weight           *float64 `json:"weight,omitempty" xml:"weight,omitempty"`
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

func (s *Version) GetImage() *string {
	return s.Image
}

func (s *Version) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *Version) GetRequestId() *string {
	return s.RequestId
}

func (s *Version) GetVersionId() *string {
	return s.VersionId
}

func (s *Version) GetWeight() *float64 {
	return s.Weight
}

func (s *Version) SetCreatedTime(v string) *Version {
	s.CreatedTime = &v
	return s
}

func (s *Version) SetDescription(v string) *Version {
	s.Description = &v
	return s
}

func (s *Version) SetImage(v string) *Version {
	s.Image = &v
	return s
}

func (s *Version) SetLastModifiedTime(v string) *Version {
	s.LastModifiedTime = &v
	return s
}

func (s *Version) SetRequestId(v string) *Version {
	s.RequestId = &v
	return s
}

func (s *Version) SetVersionId(v string) *Version {
	s.VersionId = &v
	return s
}

func (s *Version) SetWeight(v float64) *Version {
	s.Weight = &v
	return s
}

func (s *Version) Validate() error {
	return dara.Validate(s)
}
