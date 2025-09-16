// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConfigFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *DeleteConfigFileRequest
	GetFileName() *string
	SetParentFullPath(v string) *DeleteConfigFileRequest
	GetParentFullPath() *string
}

type DeleteConfigFileRequest struct {
	// The file name.
	//
	// This parameter is required.
	//
	// example:
	//
	// /schemas/automobile_vector_schema.json
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The path of the parent directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// /
	ParentFullPath *string `json:"parentFullPath,omitempty" xml:"parentFullPath,omitempty"`
}

func (s DeleteConfigFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteConfigFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *DeleteConfigFileRequest) GetParentFullPath() *string {
	return s.ParentFullPath
}

func (s *DeleteConfigFileRequest) SetFileName(v string) *DeleteConfigFileRequest {
	s.FileName = &v
	return s
}

func (s *DeleteConfigFileRequest) SetParentFullPath(v string) *DeleteConfigFileRequest {
	s.ParentFullPath = &v
	return s
}

func (s *DeleteConfigFileRequest) Validate() error {
	return dara.Validate(s)
}
