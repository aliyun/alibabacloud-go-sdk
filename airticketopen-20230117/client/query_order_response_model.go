// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryOrderResponse
	GetStatusCode() *int32
	SetBody(v *QueryOrderResponseBody) *QueryOrderResponse
	GetBody() *QueryOrderResponseBody
}

type QueryOrderResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryOrderResponse) GoString() string {
	return s.String()
}

func (s *QueryOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryOrderResponse) GetBody() *QueryOrderResponseBody {
	return s.Body
}

func (s *QueryOrderResponse) SetHeaders(v map[string]*string) *QueryOrderResponse {
	s.Headers = v
	return s
}

func (s *QueryOrderResponse) SetStatusCode(v int32) *QueryOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrderResponse) SetBody(v *QueryOrderResponseBody) *QueryOrderResponse {
	s.Body = v
	return s
}

func (s *QueryOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
