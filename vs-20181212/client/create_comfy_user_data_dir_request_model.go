// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComfyUserDataDirRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPath(v string) *CreateComfyUserDataDirRequest
	GetPath() *string
}

type CreateComfyUserDataDirRequest struct {
	// example:
	//
	// /userdir1
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s CreateComfyUserDataDirRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateComfyUserDataDirRequest) GoString() string {
	return s.String()
}

func (s *CreateComfyUserDataDirRequest) GetPath() *string {
	return s.Path
}

func (s *CreateComfyUserDataDirRequest) SetPath(v string) *CreateComfyUserDataDirRequest {
	s.Path = &v
	return s
}

func (s *CreateComfyUserDataDirRequest) Validate() error {
	return dara.Validate(s)
}
