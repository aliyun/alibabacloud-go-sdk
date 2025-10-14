// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundApplyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefundApplyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefundApplyResponse
	GetStatusCode() *int32
	SetBody(v *RefundApplyResponseBody) *RefundApplyResponse
	GetBody() *RefundApplyResponseBody
}

type RefundApplyResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefundApplyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefundApplyResponse) String() string {
	return dara.Prettify(s)
}

func (s RefundApplyResponse) GoString() string {
	return s.String()
}

func (s *RefundApplyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefundApplyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefundApplyResponse) GetBody() *RefundApplyResponseBody {
	return s.Body
}

func (s *RefundApplyResponse) SetHeaders(v map[string]*string) *RefundApplyResponse {
	s.Headers = v
	return s
}

func (s *RefundApplyResponse) SetStatusCode(v int32) *RefundApplyResponse {
	s.StatusCode = &v
	return s
}

func (s *RefundApplyResponse) SetBody(v *RefundApplyResponseBody) *RefundApplyResponse {
	s.Body = v
	return s
}

func (s *RefundApplyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
