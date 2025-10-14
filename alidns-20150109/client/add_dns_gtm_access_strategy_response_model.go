// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDnsGtmAccessStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDnsGtmAccessStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDnsGtmAccessStrategyResponse
	GetStatusCode() *int32
	SetBody(v *AddDnsGtmAccessStrategyResponseBody) *AddDnsGtmAccessStrategyResponse
	GetBody() *AddDnsGtmAccessStrategyResponseBody
}

type AddDnsGtmAccessStrategyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDnsGtmAccessStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDnsGtmAccessStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDnsGtmAccessStrategyResponse) GoString() string {
	return s.String()
}

func (s *AddDnsGtmAccessStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDnsGtmAccessStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDnsGtmAccessStrategyResponse) GetBody() *AddDnsGtmAccessStrategyResponseBody {
	return s.Body
}

func (s *AddDnsGtmAccessStrategyResponse) SetHeaders(v map[string]*string) *AddDnsGtmAccessStrategyResponse {
	s.Headers = v
	return s
}

func (s *AddDnsGtmAccessStrategyResponse) SetStatusCode(v int32) *AddDnsGtmAccessStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDnsGtmAccessStrategyResponse) SetBody(v *AddDnsGtmAccessStrategyResponseBody) *AddDnsGtmAccessStrategyResponse {
	s.Body = v
	return s
}

func (s *AddDnsGtmAccessStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
