// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOutboundCallNumbersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOutboundCallNumbersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOutboundCallNumbersResponse
	GetStatusCode() *int32
	SetBody(v *ListOutboundCallNumbersResponseBody) *ListOutboundCallNumbersResponse
	GetBody() *ListOutboundCallNumbersResponseBody
}

type ListOutboundCallNumbersResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOutboundCallNumbersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOutboundCallNumbersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOutboundCallNumbersResponse) GoString() string {
	return s.String()
}

func (s *ListOutboundCallNumbersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOutboundCallNumbersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOutboundCallNumbersResponse) GetBody() *ListOutboundCallNumbersResponseBody {
	return s.Body
}

func (s *ListOutboundCallNumbersResponse) SetHeaders(v map[string]*string) *ListOutboundCallNumbersResponse {
	s.Headers = v
	return s
}

func (s *ListOutboundCallNumbersResponse) SetStatusCode(v int32) *ListOutboundCallNumbersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOutboundCallNumbersResponse) SetBody(v *ListOutboundCallNumbersResponseBody) *ListOutboundCallNumbersResponse {
	s.Body = v
	return s
}

func (s *ListOutboundCallNumbersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
