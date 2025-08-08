// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iArtifactCode interface {
	dara.Model
	String() string
	GoString() string
	SetChecksum(v string) *ArtifactCode
	GetChecksum() *string
	SetUrl(v string) *ArtifactCode
	GetUrl() *string
}

type ArtifactCode struct {
	Checksum *string `json:"checksum,omitempty" xml:"checksum,omitempty"`
	Url      *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ArtifactCode) String() string {
	return dara.Prettify(s)
}

func (s ArtifactCode) GoString() string {
	return s.String()
}

func (s *ArtifactCode) GetChecksum() *string {
	return s.Checksum
}

func (s *ArtifactCode) GetUrl() *string {
	return s.Url
}

func (s *ArtifactCode) SetChecksum(v string) *ArtifactCode {
	s.Checksum = &v
	return s
}

func (s *ArtifactCode) SetUrl(v string) *ArtifactCode {
	s.Url = &v
	return s
}

func (s *ArtifactCode) Validate() error {
	return dara.Validate(s)
}
