// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFolderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Folder) *UpdateFolderRequest
	GetBody() *Folder
}

type UpdateFolderRequest struct {
	// This parameter is required.
	Body *Folder `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFolderRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFolderRequest) GoString() string {
	return s.String()
}

func (s *UpdateFolderRequest) GetBody() *Folder {
	return s.Body
}

func (s *UpdateFolderRequest) SetBody(v *Folder) *UpdateFolderRequest {
	s.Body = v
	return s
}

func (s *UpdateFolderRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
