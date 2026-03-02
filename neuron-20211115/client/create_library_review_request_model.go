// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLibraryReviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *LibraryReviewCreateCmd) *CreateLibraryReviewRequest
	GetBody() *LibraryReviewCreateCmd
}

type CreateLibraryReviewRequest struct {
	Body *LibraryReviewCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLibraryReviewRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLibraryReviewRequest) GoString() string {
	return s.String()
}

func (s *CreateLibraryReviewRequest) GetBody() *LibraryReviewCreateCmd {
	return s.Body
}

func (s *CreateLibraryReviewRequest) SetBody(v *LibraryReviewCreateCmd) *CreateLibraryReviewRequest {
	s.Body = v
	return s
}

func (s *CreateLibraryReviewRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
