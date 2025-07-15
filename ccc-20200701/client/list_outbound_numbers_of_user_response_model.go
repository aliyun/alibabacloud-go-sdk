// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOutboundNumbersOfUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOutboundNumbersOfUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOutboundNumbersOfUserResponse
	GetStatusCode() *int32
	SetBody(v *ListOutboundNumbersOfUserResponseBody) *ListOutboundNumbersOfUserResponse
	GetBody() *ListOutboundNumbersOfUserResponseBody
}

type ListOutboundNumbersOfUserResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOutboundNumbersOfUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOutboundNumbersOfUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOutboundNumbersOfUserResponse) GoString() string {
	return s.String()
}

func (s *ListOutboundNumbersOfUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOutboundNumbersOfUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOutboundNumbersOfUserResponse) GetBody() *ListOutboundNumbersOfUserResponseBody {
	return s.Body
}

func (s *ListOutboundNumbersOfUserResponse) SetHeaders(v map[string]*string) *ListOutboundNumbersOfUserResponse {
	s.Headers = v
	return s
}

func (s *ListOutboundNumbersOfUserResponse) SetStatusCode(v int32) *ListOutboundNumbersOfUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOutboundNumbersOfUserResponse) SetBody(v *ListOutboundNumbersOfUserResponseBody) *ListOutboundNumbersOfUserResponse {
	s.Body = v
	return s
}

func (s *ListOutboundNumbersOfUserResponse) Validate() error {
	return dara.Validate(s)
}
