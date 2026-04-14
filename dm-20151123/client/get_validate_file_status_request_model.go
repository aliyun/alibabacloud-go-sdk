// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetValidateFileStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v string) *GetValidateFileStatusRequest
	GetFileId() *string
}

type GetValidateFileStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxx
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
}

func (s GetValidateFileStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetValidateFileStatusRequest) GoString() string {
	return s.String()
}

func (s *GetValidateFileStatusRequest) GetFileId() *string {
	return s.FileId
}

func (s *GetValidateFileStatusRequest) SetFileId(v string) *GetValidateFileStatusRequest {
	s.FileId = &v
	return s
}

func (s *GetValidateFileStatusRequest) Validate() error {
	return dara.Validate(s)
}
