// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineArtifactUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *GetPipelineArtifactUrlRequest
	GetFileName() *string
	SetFilePath(v string) *GetPipelineArtifactUrlRequest
	GetFilePath() *string
}

type GetPipelineArtifactUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test.tgz
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /test/test/test.tgz
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
}

func (s GetPipelineArtifactUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineArtifactUrlRequest) GoString() string {
	return s.String()
}

func (s *GetPipelineArtifactUrlRequest) GetFileName() *string {
	return s.FileName
}

func (s *GetPipelineArtifactUrlRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *GetPipelineArtifactUrlRequest) SetFileName(v string) *GetPipelineArtifactUrlRequest {
	s.FileName = &v
	return s
}

func (s *GetPipelineArtifactUrlRequest) SetFilePath(v string) *GetPipelineArtifactUrlRequest {
	s.FilePath = &v
	return s
}

func (s *GetPipelineArtifactUrlRequest) Validate() error {
	return dara.Validate(s)
}
