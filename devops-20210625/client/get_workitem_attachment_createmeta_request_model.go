// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkitemAttachmentCreatemetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *GetWorkitemAttachmentCreatemetaRequest
	GetFileName() *string
}

type GetWorkitemAttachmentCreatemetaRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// application.jar
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s GetWorkitemAttachmentCreatemetaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkitemAttachmentCreatemetaRequest) GoString() string {
	return s.String()
}

func (s *GetWorkitemAttachmentCreatemetaRequest) GetFileName() *string {
	return s.FileName
}

func (s *GetWorkitemAttachmentCreatemetaRequest) SetFileName(v string) *GetWorkitemAttachmentCreatemetaRequest {
	s.FileName = &v
	return s
}

func (s *GetWorkitemAttachmentCreatemetaRequest) Validate() error {
	return dara.Validate(s)
}
