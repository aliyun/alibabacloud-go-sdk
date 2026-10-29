// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComfyUserDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *DeleteComfyUserDataRequest
	GetFileName() *string
}

type DeleteComfyUserDataRequest struct {
	// The name of the file or directory to delete. If a path is passed, it is treated as a directory and must end with a forward slash (/).
	//
	// File name.
	//
	// example:
	//
	// mytest
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s DeleteComfyUserDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteComfyUserDataRequest) GoString() string {
	return s.String()
}

func (s *DeleteComfyUserDataRequest) GetFileName() *string {
	return s.FileName
}

func (s *DeleteComfyUserDataRequest) SetFileName(v string) *DeleteComfyUserDataRequest {
	s.FileName = &v
	return s
}

func (s *DeleteComfyUserDataRequest) Validate() error {
	return dara.Validate(s)
}
