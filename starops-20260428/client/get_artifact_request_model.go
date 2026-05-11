// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArtifactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactPath(v string) *GetArtifactRequest
	GetArtifactPath() *string
}

type GetArtifactRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// missions/mission-xxx/artifacts/2026-05/05-01/xxxx.md
	ArtifactPath *string `json:"artifactPath,omitempty" xml:"artifactPath,omitempty"`
}

func (s GetArtifactRequest) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactRequest) GoString() string {
	return s.String()
}

func (s *GetArtifactRequest) GetArtifactPath() *string {
	return s.ArtifactPath
}

func (s *GetArtifactRequest) SetArtifactPath(v string) *GetArtifactRequest {
	s.ArtifactPath = &v
	return s
}

func (s *GetArtifactRequest) Validate() error {
	return dara.Validate(s)
}
