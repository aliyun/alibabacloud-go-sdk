// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArtifactDownloadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactPath(v string) *GetArtifactDownloadUrlRequest
	GetArtifactPath() *string
}

type GetArtifactDownloadUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// reports/summary.pdf
	ArtifactPath *string `json:"artifactPath,omitempty" xml:"artifactPath,omitempty"`
}

func (s GetArtifactDownloadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetArtifactDownloadUrlRequest) GetArtifactPath() *string {
	return s.ArtifactPath
}

func (s *GetArtifactDownloadUrlRequest) SetArtifactPath(v string) *GetArtifactDownloadUrlRequest {
	s.ArtifactPath = &v
	return s
}

func (s *GetArtifactDownloadUrlRequest) Validate() error {
	return dara.Validate(s)
}
