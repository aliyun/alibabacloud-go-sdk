// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *GetFileRequest
	GetFileName() *string
}

type GetFileRequest struct {
	// The name of the file in full path
	//
	// This parameter is required.
	//
	// example:
	//
	// /schemas/automobile_vector_schema.json
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s GetFileRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileRequest) GoString() string {
	return s.String()
}

func (s *GetFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *GetFileRequest) SetFileName(v string) *GetFileRequest {
	s.FileName = &v
	return s
}

func (s *GetFileRequest) Validate() error {
	return dara.Validate(s)
}
