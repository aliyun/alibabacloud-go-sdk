// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetValidateFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v string) *GetValidateFileRequest
	GetFileId() *string
}

type GetValidateFileRequest struct {
	// example:
	//
	// xxxx-xxxx-xxxx-xxxx
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
}

func (s GetValidateFileRequest) String() string {
	return dara.Prettify(s)
}

func (s GetValidateFileRequest) GoString() string {
	return s.String()
}

func (s *GetValidateFileRequest) GetFileId() *string {
	return s.FileId
}

func (s *GetValidateFileRequest) SetFileId(v string) *GetValidateFileRequest {
	s.FileId = &v
	return s
}

func (s *GetValidateFileRequest) Validate() error {
	return dara.Validate(s)
}
