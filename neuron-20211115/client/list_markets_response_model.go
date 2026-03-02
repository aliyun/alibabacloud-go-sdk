// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMarketsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMarketsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMarketsResponse
	GetStatusCode() *int32
	SetBody(v *MarketListResult) *ListMarketsResponse
	GetBody() *MarketListResult
}

type ListMarketsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MarketListResult  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMarketsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMarketsResponse) GoString() string {
	return s.String()
}

func (s *ListMarketsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMarketsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMarketsResponse) GetBody() *MarketListResult {
	return s.Body
}

func (s *ListMarketsResponse) SetHeaders(v map[string]*string) *ListMarketsResponse {
	s.Headers = v
	return s
}

func (s *ListMarketsResponse) SetStatusCode(v int32) *ListMarketsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMarketsResponse) SetBody(v *MarketListResult) *ListMarketsResponse {
	s.Body = v
	return s
}

func (s *ListMarketsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
