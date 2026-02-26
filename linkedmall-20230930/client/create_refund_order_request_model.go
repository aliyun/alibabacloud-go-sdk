// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRefundOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *RefundOrderCmd) *CreateRefundOrderRequest
	GetBody() *RefundOrderCmd
}

type CreateRefundOrderRequest struct {
	// This parameter is required.
	Body *RefundOrderCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRefundOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRefundOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateRefundOrderRequest) GetBody() *RefundOrderCmd {
	return s.Body
}

func (s *CreateRefundOrderRequest) SetBody(v *RefundOrderCmd) *CreateRefundOrderRequest {
	s.Body = v
	return s
}

func (s *CreateRefundOrderRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
