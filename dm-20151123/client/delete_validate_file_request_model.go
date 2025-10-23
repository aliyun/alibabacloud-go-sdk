// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteValidateFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v string) *DeleteValidateFileRequest
	GetFileId() *string
}

type DeleteValidateFileRequest struct {
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
}

func (s DeleteValidateFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteValidateFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteValidateFileRequest) GetFileId() *string {
	return s.FileId
}

func (s *DeleteValidateFileRequest) SetFileId(v string) *DeleteValidateFileRequest {
	s.FileId = &v
	return s
}

func (s *DeleteValidateFileRequest) Validate() error {
	return dara.Validate(s)
}
