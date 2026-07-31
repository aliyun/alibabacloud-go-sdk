// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssStsTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileType(v string) *GetOssStsTokenRequest
	GetFileType() *string
}

type GetOssStsTokenRequest struct {
	// The file type.
	//
	// example:
	//
	// SKILL
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
}

func (s GetOssStsTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOssStsTokenRequest) GoString() string {
	return s.String()
}

func (s *GetOssStsTokenRequest) GetFileType() *string {
	return s.FileType
}

func (s *GetOssStsTokenRequest) SetFileType(v string) *GetOssStsTokenRequest {
	s.FileType = &v
	return s
}

func (s *GetOssStsTokenRequest) Validate() error {
	return dara.Validate(s)
}
