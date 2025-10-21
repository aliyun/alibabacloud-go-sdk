// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFolderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Folder) *CreateFolderRequest
	GetBody() *Folder
}

type CreateFolderRequest struct {
	Body *Folder `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFolderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFolderRequest) GoString() string {
	return s.String()
}

func (s *CreateFolderRequest) GetBody() *Folder {
	return s.Body
}

func (s *CreateFolderRequest) SetBody(v *Folder) *CreateFolderRequest {
	s.Body = v
	return s
}

func (s *CreateFolderRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
