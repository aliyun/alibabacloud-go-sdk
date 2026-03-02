// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokePbcReviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *RevokeReviewCmd) *RevokePbcReviewRequest
	GetBody() *RevokeReviewCmd
}

type RevokePbcReviewRequest struct {
	Body *RevokeReviewCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokePbcReviewRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokePbcReviewRequest) GoString() string {
	return s.String()
}

func (s *RevokePbcReviewRequest) GetBody() *RevokeReviewCmd {
	return s.Body
}

func (s *RevokePbcReviewRequest) SetBody(v *RevokeReviewCmd) *RevokePbcReviewRequest {
	s.Body = v
	return s
}

func (s *RevokePbcReviewRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
