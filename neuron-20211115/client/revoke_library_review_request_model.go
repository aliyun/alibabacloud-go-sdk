// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeLibraryReviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *RevokeLibraryReviewCmd) *RevokeLibraryReviewRequest
	GetBody() *RevokeLibraryReviewCmd
}

type RevokeLibraryReviewRequest struct {
	Body *RevokeLibraryReviewCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeLibraryReviewRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeLibraryReviewRequest) GoString() string {
	return s.String()
}

func (s *RevokeLibraryReviewRequest) GetBody() *RevokeLibraryReviewCmd {
	return s.Body
}

func (s *RevokeLibraryReviewRequest) SetBody(v *RevokeLibraryReviewCmd) *RevokeLibraryReviewRequest {
	s.Body = v
	return s
}

func (s *RevokeLibraryReviewRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
