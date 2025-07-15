// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPickOutboundNumbersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PickOutboundNumbersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PickOutboundNumbersResponse
	GetStatusCode() *int32
	SetBody(v *PickOutboundNumbersResponseBody) *PickOutboundNumbersResponse
	GetBody() *PickOutboundNumbersResponseBody
}

type PickOutboundNumbersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PickOutboundNumbersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PickOutboundNumbersResponse) String() string {
	return dara.Prettify(s)
}

func (s PickOutboundNumbersResponse) GoString() string {
	return s.String()
}

func (s *PickOutboundNumbersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PickOutboundNumbersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PickOutboundNumbersResponse) GetBody() *PickOutboundNumbersResponseBody {
	return s.Body
}

func (s *PickOutboundNumbersResponse) SetHeaders(v map[string]*string) *PickOutboundNumbersResponse {
	s.Headers = v
	return s
}

func (s *PickOutboundNumbersResponse) SetStatusCode(v int32) *PickOutboundNumbersResponse {
	s.StatusCode = &v
	return s
}

func (s *PickOutboundNumbersResponse) SetBody(v *PickOutboundNumbersResponseBody) *PickOutboundNumbersResponse {
	s.Body = v
	return s
}

func (s *PickOutboundNumbersResponse) Validate() error {
	return dara.Validate(s)
}
