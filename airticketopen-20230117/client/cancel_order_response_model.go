// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelOrderResponse
	GetStatusCode() *int32
	SetBody(v *CancelOrderResponseBody) *CancelOrderResponse
	GetBody() *CancelOrderResponseBody
}

type CancelOrderResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelOrderResponse) GoString() string {
	return s.String()
}

func (s *CancelOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelOrderResponse) GetBody() *CancelOrderResponseBody {
	return s.Body
}

func (s *CancelOrderResponse) SetHeaders(v map[string]*string) *CancelOrderResponse {
	s.Headers = v
	return s
}

func (s *CancelOrderResponse) SetStatusCode(v int32) *CancelOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelOrderResponse) SetBody(v *CancelOrderResponseBody) *CancelOrderResponse {
	s.Body = v
	return s
}

func (s *CancelOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
