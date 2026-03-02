// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLibraryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *LibraryCreateCmd) *CreateLibraryRequest
	GetBody() *LibraryCreateCmd
}

type CreateLibraryRequest struct {
	Body *LibraryCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLibraryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLibraryRequest) GoString() string {
	return s.String()
}

func (s *CreateLibraryRequest) GetBody() *LibraryCreateCmd {
	return s.Body
}

func (s *CreateLibraryRequest) SetBody(v *LibraryCreateCmd) *CreateLibraryRequest {
	s.Body = v
	return s
}

func (s *CreateLibraryRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
