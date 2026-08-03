// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyRefundResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyRefundResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyRefundResponse
	GetStatusCode() *int32
	SetBody(v *ApplyRefundResponseBody) *ApplyRefundResponse
	GetBody() *ApplyRefundResponseBody
}

type ApplyRefundResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyRefundResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyRefundResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyRefundResponse) GoString() string {
	return s.String()
}

func (s *ApplyRefundResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyRefundResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyRefundResponse) GetBody() *ApplyRefundResponseBody {
	return s.Body
}

func (s *ApplyRefundResponse) SetHeaders(v map[string]*string) *ApplyRefundResponse {
	s.Headers = v
	return s
}

func (s *ApplyRefundResponse) SetStatusCode(v int32) *ApplyRefundResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyRefundResponse) SetBody(v *ApplyRefundResponseBody) *ApplyRefundResponse {
	s.Body = v
	return s
}

func (s *ApplyRefundResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
