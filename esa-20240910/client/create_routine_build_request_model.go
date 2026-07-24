// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoutineBuildRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactUrl(v string) *CreateRoutineBuildRequest
	GetArtifactUrl() *string
	SetBranch(v string) *CreateRoutineBuildRequest
	GetBranch() *string
	SetRoutineName(v string) *CreateRoutineBuildRequest
	GetRoutineName() *string
}

type CreateRoutineBuildRequest struct {
	// The OSS object URL. This parameter is required in upload mode but is not required in git mode.
	//
	// example:
	//
	// https://bucket.oss-.aliyuncs.com/key
	ArtifactUrl *string `json:"ArtifactUrl,omitempty" xml:"ArtifactUrl,omitempty"`
	// The name of the branch to build. This parameter is not required in upload mode but is required in git mode.
	//
	// example:
	//
	// main
	Branch *string `json:"Branch,omitempty" xml:"Branch,omitempty"`
	// The ER name.
	//
	// This parameter is required.
	//
	// example:
	//
	// rwa-test
	RoutineName *string `json:"RoutineName,omitempty" xml:"RoutineName,omitempty"`
}

func (s CreateRoutineBuildRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineBuildRequest) GoString() string {
	return s.String()
}

func (s *CreateRoutineBuildRequest) GetArtifactUrl() *string {
	return s.ArtifactUrl
}

func (s *CreateRoutineBuildRequest) GetBranch() *string {
	return s.Branch
}

func (s *CreateRoutineBuildRequest) GetRoutineName() *string {
	return s.RoutineName
}

func (s *CreateRoutineBuildRequest) SetArtifactUrl(v string) *CreateRoutineBuildRequest {
	s.ArtifactUrl = &v
	return s
}

func (s *CreateRoutineBuildRequest) SetBranch(v string) *CreateRoutineBuildRequest {
	s.Branch = &v
	return s
}

func (s *CreateRoutineBuildRequest) SetRoutineName(v string) *CreateRoutineBuildRequest {
	s.RoutineName = &v
	return s
}

func (s *CreateRoutineBuildRequest) Validate() error {
	return dara.Validate(s)
}
