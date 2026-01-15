// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateModelFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilePath(v string) *ValidateModelFileRequest
	GetFilePath() *string
}

type ValidateModelFileRequest struct {
	// File path.
	//
	// example:
	//
	// saf/xxxxxxx/xxx.pmml
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
}

func (s ValidateModelFileRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateModelFileRequest) GoString() string {
	return s.String()
}

func (s *ValidateModelFileRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *ValidateModelFileRequest) SetFilePath(v string) *ValidateModelFileRequest {
	s.FilePath = &v
	return s
}

func (s *ValidateModelFileRequest) Validate() error {
	return dara.Validate(s)
}
