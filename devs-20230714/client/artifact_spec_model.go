// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iArtifactSpec interface {
	dara.Model
	String() string
	GoString() string
	SetRuntime(v string) *ArtifactSpec
	GetRuntime() *string
	SetType(v string) *ArtifactSpec
	GetType() *string
	SetUri(v string) *ArtifactSpec
	GetUri() *string
}

type ArtifactSpec struct {
	// This parameter is required.
	//
	// example:
	//
	// custom.debian10
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FC代码包、工作流yaml
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss://cn-hangzhou/my-bucket/my.zip
	Uri *string `json:"uri,omitempty" xml:"uri,omitempty"`
}

func (s ArtifactSpec) String() string {
	return dara.Prettify(s)
}

func (s ArtifactSpec) GoString() string {
	return s.String()
}

func (s *ArtifactSpec) GetRuntime() *string {
	return s.Runtime
}

func (s *ArtifactSpec) GetType() *string {
	return s.Type
}

func (s *ArtifactSpec) GetUri() *string {
	return s.Uri
}

func (s *ArtifactSpec) SetRuntime(v string) *ArtifactSpec {
	s.Runtime = &v
	return s
}

func (s *ArtifactSpec) SetType(v string) *ArtifactSpec {
	s.Type = &v
	return s
}

func (s *ArtifactSpec) SetUri(v string) *ArtifactSpec {
	s.Uri = &v
	return s
}

func (s *ArtifactSpec) Validate() error {
	return dara.Validate(s)
}
