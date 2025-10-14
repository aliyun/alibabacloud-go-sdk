// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDnsGtmAccessStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDnsGtmAccessStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDnsGtmAccessStrategyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDnsGtmAccessStrategyResponseBody) *DeleteDnsGtmAccessStrategyResponse
	GetBody() *DeleteDnsGtmAccessStrategyResponseBody
}

type DeleteDnsGtmAccessStrategyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDnsGtmAccessStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDnsGtmAccessStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDnsGtmAccessStrategyResponse) GoString() string {
	return s.String()
}

func (s *DeleteDnsGtmAccessStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDnsGtmAccessStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDnsGtmAccessStrategyResponse) GetBody() *DeleteDnsGtmAccessStrategyResponseBody {
	return s.Body
}

func (s *DeleteDnsGtmAccessStrategyResponse) SetHeaders(v map[string]*string) *DeleteDnsGtmAccessStrategyResponse {
	s.Headers = v
	return s
}

func (s *DeleteDnsGtmAccessStrategyResponse) SetStatusCode(v int32) *DeleteDnsGtmAccessStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDnsGtmAccessStrategyResponse) SetBody(v *DeleteDnsGtmAccessStrategyResponseBody) *DeleteDnsGtmAccessStrategyResponse {
	s.Body = v
	return s
}

func (s *DeleteDnsGtmAccessStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
