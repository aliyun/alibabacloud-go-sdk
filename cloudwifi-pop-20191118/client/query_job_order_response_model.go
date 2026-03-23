// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryJobOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryJobOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryJobOrderResponse
	GetStatusCode() *int32
	SetBody(v *QueryJobOrderResponseBody) *QueryJobOrderResponse
	GetBody() *QueryJobOrderResponseBody
}

type QueryJobOrderResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryJobOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryJobOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryJobOrderResponse) GoString() string {
	return s.String()
}

func (s *QueryJobOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryJobOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryJobOrderResponse) GetBody() *QueryJobOrderResponseBody {
	return s.Body
}

func (s *QueryJobOrderResponse) SetHeaders(v map[string]*string) *QueryJobOrderResponse {
	s.Headers = v
	return s
}

func (s *QueryJobOrderResponse) SetStatusCode(v int32) *QueryJobOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryJobOrderResponse) SetBody(v *QueryJobOrderResponseBody) *QueryJobOrderResponse {
	s.Body = v
	return s
}

func (s *QueryJobOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
