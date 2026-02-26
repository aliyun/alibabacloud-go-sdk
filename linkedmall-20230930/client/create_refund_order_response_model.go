// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRefundOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRefundOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRefundOrderResponse
	GetStatusCode() *int32
	SetBody(v *RefundOrderResult) *CreateRefundOrderResponse
	GetBody() *RefundOrderResult
}

type CreateRefundOrderResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefundOrderResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRefundOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRefundOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateRefundOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRefundOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRefundOrderResponse) GetBody() *RefundOrderResult {
	return s.Body
}

func (s *CreateRefundOrderResponse) SetHeaders(v map[string]*string) *CreateRefundOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateRefundOrderResponse) SetStatusCode(v int32) *CreateRefundOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRefundOrderResponse) SetBody(v *RefundOrderResult) *CreateRefundOrderResponse {
	s.Body = v
	return s
}

func (s *CreateRefundOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
