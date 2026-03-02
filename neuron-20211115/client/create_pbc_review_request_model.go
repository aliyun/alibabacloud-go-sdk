// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePbcReviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PbcReviewCreateCmd) *CreatePbcReviewRequest
	GetBody() *PbcReviewCreateCmd
}

type CreatePbcReviewRequest struct {
	Body *PbcReviewCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePbcReviewRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePbcReviewRequest) GoString() string {
	return s.String()
}

func (s *CreatePbcReviewRequest) GetBody() *PbcReviewCreateCmd {
	return s.Body
}

func (s *CreatePbcReviewRequest) SetBody(v *PbcReviewCreateCmd) *CreatePbcReviewRequest {
	s.Body = v
	return s
}

func (s *CreatePbcReviewRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
