// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateForkReviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ForkReviewCreateCmd) *CreateForkReviewRequest
	GetBody() *ForkReviewCreateCmd
}

type CreateForkReviewRequest struct {
	Body *ForkReviewCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateForkReviewRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateForkReviewRequest) GoString() string {
	return s.String()
}

func (s *CreateForkReviewRequest) GetBody() *ForkReviewCreateCmd {
	return s.Body
}

func (s *CreateForkReviewRequest) SetBody(v *ForkReviewCreateCmd) *CreateForkReviewRequest {
	s.Body = v
	return s
}

func (s *CreateForkReviewRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
