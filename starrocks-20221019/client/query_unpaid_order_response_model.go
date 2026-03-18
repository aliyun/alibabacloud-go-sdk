// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUnpaidOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryUnpaidOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryUnpaidOrderResponse
	GetStatusCode() *int32
	SetBody(v *QueryUnpaidOrderResponseBody) *QueryUnpaidOrderResponse
	GetBody() *QueryUnpaidOrderResponseBody
}

type QueryUnpaidOrderResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUnpaidOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUnpaidOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryUnpaidOrderResponse) GoString() string {
	return s.String()
}

func (s *QueryUnpaidOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryUnpaidOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryUnpaidOrderResponse) GetBody() *QueryUnpaidOrderResponseBody {
	return s.Body
}

func (s *QueryUnpaidOrderResponse) SetHeaders(v map[string]*string) *QueryUnpaidOrderResponse {
	s.Headers = v
	return s
}

func (s *QueryUnpaidOrderResponse) SetStatusCode(v int32) *QueryUnpaidOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUnpaidOrderResponse) SetBody(v *QueryUnpaidOrderResponseBody) *QueryUnpaidOrderResponse {
	s.Body = v
	return s
}

func (s *QueryUnpaidOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
