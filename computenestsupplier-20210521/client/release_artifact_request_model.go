// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseArtifactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactId(v string) *ReleaseArtifactRequest
	GetArtifactId() *string
	SetClientToken(v string) *ReleaseArtifactRequest
	GetClientToken() *string
}

type ReleaseArtifactRequest struct {
	// The artifact ID.
	//
	// Call [ListArtifacts](https://help.aliyun.com/document_detail/469993.html) to obtain the artifact ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// artifact-9feded91880e4c78xxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// A client-generated token that ensures the idempotence of the request. The token must be unique for each request. The value of **ClientToken*	- can contain only ASCII characters and must be no more than 64 characters in length.
	//
	// example:
	//
	// 788E7CP0EN9D51P
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s ReleaseArtifactRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseArtifactRequest) GoString() string {
	return s.String()
}

func (s *ReleaseArtifactRequest) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *ReleaseArtifactRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ReleaseArtifactRequest) SetArtifactId(v string) *ReleaseArtifactRequest {
	s.ArtifactId = &v
	return s
}

func (s *ReleaseArtifactRequest) SetClientToken(v string) *ReleaseArtifactRequest {
	s.ClientToken = &v
	return s
}

func (s *ReleaseArtifactRequest) Validate() error {
	return dara.Validate(s)
}
