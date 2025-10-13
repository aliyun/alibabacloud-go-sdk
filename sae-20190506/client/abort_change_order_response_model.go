// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbortChangeOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AbortChangeOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AbortChangeOrderResponse
	GetStatusCode() *int32
	SetBody(v *AbortChangeOrderResponseBody) *AbortChangeOrderResponse
	GetBody() *AbortChangeOrderResponseBody
}

type AbortChangeOrderResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AbortChangeOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AbortChangeOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s AbortChangeOrderResponse) GoString() string {
	return s.String()
}

func (s *AbortChangeOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AbortChangeOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AbortChangeOrderResponse) GetBody() *AbortChangeOrderResponseBody {
	return s.Body
}

func (s *AbortChangeOrderResponse) SetHeaders(v map[string]*string) *AbortChangeOrderResponse {
	s.Headers = v
	return s
}

func (s *AbortChangeOrderResponse) SetStatusCode(v int32) *AbortChangeOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *AbortChangeOrderResponse) SetBody(v *AbortChangeOrderResponseBody) *AbortChangeOrderResponse {
	s.Body = v
	return s
}

func (s *AbortChangeOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
