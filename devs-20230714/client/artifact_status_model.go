// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iArtifactStatus interface {
	dara.Model
	String() string
	GoString() string
	SetArn(v string) *ArtifactStatus
	GetArn() *string
	SetChecksum(v string) *ArtifactStatus
	GetChecksum() *string
	SetSize(v int64) *ArtifactStatus
	GetSize() *int64
}

type ArtifactStatus struct {
	// example:
	//
	// acs:devs:cn-hangzhou:123456:artifacts/my-first-artifact
	Arn *string `json:"arn,omitempty" xml:"arn,omitempty"`
	// example:
	//
	// 2825179536350****
	Checksum *string `json:"checksum,omitempty" xml:"checksum,omitempty"`
	// example:
	//
	// 1024
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ArtifactStatus) String() string {
	return dara.Prettify(s)
}

func (s ArtifactStatus) GoString() string {
	return s.String()
}

func (s *ArtifactStatus) GetArn() *string {
	return s.Arn
}

func (s *ArtifactStatus) GetChecksum() *string {
	return s.Checksum
}

func (s *ArtifactStatus) GetSize() *int64 {
	return s.Size
}

func (s *ArtifactStatus) SetArn(v string) *ArtifactStatus {
	s.Arn = &v
	return s
}

func (s *ArtifactStatus) SetChecksum(v string) *ArtifactStatus {
	s.Checksum = &v
	return s
}

func (s *ArtifactStatus) SetSize(v int64) *ArtifactStatus {
	s.Size = &v
	return s
}

func (s *ArtifactStatus) Validate() error {
	return dara.Validate(s)
}
