// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHiMarketConsumersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHiMarketConsumersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHiMarketConsumersResponse
	GetStatusCode() *int32
	SetBody(v *ListHiMarketConsumersResponseBody) *ListHiMarketConsumersResponse
	GetBody() *ListHiMarketConsumersResponseBody
}

type ListHiMarketConsumersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHiMarketConsumersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHiMarketConsumersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHiMarketConsumersResponse) GoString() string {
	return s.String()
}

func (s *ListHiMarketConsumersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHiMarketConsumersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHiMarketConsumersResponse) GetBody() *ListHiMarketConsumersResponseBody {
	return s.Body
}

func (s *ListHiMarketConsumersResponse) SetHeaders(v map[string]*string) *ListHiMarketConsumersResponse {
	s.Headers = v
	return s
}

func (s *ListHiMarketConsumersResponse) SetStatusCode(v int32) *ListHiMarketConsumersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHiMarketConsumersResponse) SetBody(v *ListHiMarketConsumersResponseBody) *ListHiMarketConsumersResponse {
	s.Body = v
	return s
}

func (s *ListHiMarketConsumersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
