// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLibraryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *LibraryUpdateCmd) *UpdateLibraryRequest
	GetBody() *LibraryUpdateCmd
}

type UpdateLibraryRequest struct {
	Body *LibraryUpdateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLibraryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLibraryRequest) GoString() string {
	return s.String()
}

func (s *UpdateLibraryRequest) GetBody() *LibraryUpdateCmd {
	return s.Body
}

func (s *UpdateLibraryRequest) SetBody(v *LibraryUpdateCmd) *UpdateLibraryRequest {
	s.Body = v
	return s
}

func (s *UpdateLibraryRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
