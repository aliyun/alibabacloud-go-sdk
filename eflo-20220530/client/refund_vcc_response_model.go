// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundVccResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefundVccResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefundVccResponse
	GetStatusCode() *int32
	SetBody(v *RefundVccResponseBody) *RefundVccResponse
	GetBody() *RefundVccResponseBody
}

type RefundVccResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefundVccResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefundVccResponse) String() string {
	return dara.Prettify(s)
}

func (s RefundVccResponse) GoString() string {
	return s.String()
}

func (s *RefundVccResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefundVccResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefundVccResponse) GetBody() *RefundVccResponseBody {
	return s.Body
}

func (s *RefundVccResponse) SetHeaders(v map[string]*string) *RefundVccResponse {
	s.Headers = v
	return s
}

func (s *RefundVccResponse) SetStatusCode(v int32) *RefundVccResponse {
	s.StatusCode = &v
	return s
}

func (s *RefundVccResponse) SetBody(v *RefundVccResponseBody) *RefundVccResponse {
	s.Body = v
	return s
}

func (s *RefundVccResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
