// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWorkitemAttachmentCreateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileKey(v string) *WorkitemAttachmentCreateRequest
	GetFileKey() *string
	SetOriginalFilename(v string) *WorkitemAttachmentCreateRequest
	GetOriginalFilename() *string
}

type WorkitemAttachmentCreateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1106473328927498/1106473328927498/5ec071g0e5ij85fche8cddchje.xlsx
	FileKey *string `json:"fileKey,omitempty" xml:"fileKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx.tgz
	OriginalFilename *string `json:"originalFilename,omitempty" xml:"originalFilename,omitempty"`
}

func (s WorkitemAttachmentCreateRequest) String() string {
	return dara.Prettify(s)
}

func (s WorkitemAttachmentCreateRequest) GoString() string {
	return s.String()
}

func (s *WorkitemAttachmentCreateRequest) GetFileKey() *string {
	return s.FileKey
}

func (s *WorkitemAttachmentCreateRequest) GetOriginalFilename() *string {
	return s.OriginalFilename
}

func (s *WorkitemAttachmentCreateRequest) SetFileKey(v string) *WorkitemAttachmentCreateRequest {
	s.FileKey = &v
	return s
}

func (s *WorkitemAttachmentCreateRequest) SetOriginalFilename(v string) *WorkitemAttachmentCreateRequest {
	s.OriginalFilename = &v
	return s
}

func (s *WorkitemAttachmentCreateRequest) Validate() error {
	return dara.Validate(s)
}
