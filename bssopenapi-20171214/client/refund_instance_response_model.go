// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefundInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefundInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RefundInstanceResponseBody) *RefundInstanceResponse
	GetBody() *RefundInstanceResponseBody
}

type RefundInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefundInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefundInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RefundInstanceResponse) GoString() string {
	return s.String()
}

func (s *RefundInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefundInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefundInstanceResponse) GetBody() *RefundInstanceResponseBody {
	return s.Body
}

func (s *RefundInstanceResponse) SetHeaders(v map[string]*string) *RefundInstanceResponse {
	s.Headers = v
	return s
}

func (s *RefundInstanceResponse) SetStatusCode(v int32) *RefundInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RefundInstanceResponse) SetBody(v *RefundInstanceResponseBody) *RefundInstanceResponse {
	s.Body = v
	return s
}

func (s *RefundInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
