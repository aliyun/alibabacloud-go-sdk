// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePbcInvokeReviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PbcInvokeReviewCreateCmd) *CreatePbcInvokeReviewRequest
	GetBody() *PbcInvokeReviewCreateCmd
}

type CreatePbcInvokeReviewRequest struct {
	Body *PbcInvokeReviewCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePbcInvokeReviewRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePbcInvokeReviewRequest) GoString() string {
	return s.String()
}

func (s *CreatePbcInvokeReviewRequest) GetBody() *PbcInvokeReviewCreateCmd {
	return s.Body
}

func (s *CreatePbcInvokeReviewRequest) SetBody(v *PbcInvokeReviewCreateCmd) *CreatePbcInvokeReviewRequest {
	s.Body = v
	return s
}

func (s *CreatePbcInvokeReviewRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
