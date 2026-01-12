// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSourceFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v string) *DeleteSourceFileRequest
	GetFileId() *string
	SetJwtToken(v string) *DeleteSourceFileRequest
	GetJwtToken() *string
}

type DeleteSourceFileRequest struct {
	// This parameter is required.
	FileId   *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
}

func (s DeleteSourceFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSourceFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteSourceFileRequest) GetFileId() *string {
	return s.FileId
}

func (s *DeleteSourceFileRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *DeleteSourceFileRequest) SetFileId(v string) *DeleteSourceFileRequest {
	s.FileId = &v
	return s
}

func (s *DeleteSourceFileRequest) SetJwtToken(v string) *DeleteSourceFileRequest {
	s.JwtToken = &v
	return s
}

func (s *DeleteSourceFileRequest) Validate() error {
	return dara.Validate(s)
}
