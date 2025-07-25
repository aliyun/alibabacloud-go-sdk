// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDnsGtmAccessStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDnsGtmAccessStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDnsGtmAccessStrategyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDnsGtmAccessStrategyResponseBody) *UpdateDnsGtmAccessStrategyResponse
	GetBody() *UpdateDnsGtmAccessStrategyResponseBody
}

type UpdateDnsGtmAccessStrategyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDnsGtmAccessStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDnsGtmAccessStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDnsGtmAccessStrategyResponse) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmAccessStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDnsGtmAccessStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDnsGtmAccessStrategyResponse) GetBody() *UpdateDnsGtmAccessStrategyResponseBody {
	return s.Body
}

func (s *UpdateDnsGtmAccessStrategyResponse) SetHeaders(v map[string]*string) *UpdateDnsGtmAccessStrategyResponse {
	s.Headers = v
	return s
}

func (s *UpdateDnsGtmAccessStrategyResponse) SetStatusCode(v int32) *UpdateDnsGtmAccessStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyResponse) SetBody(v *UpdateDnsGtmAccessStrategyResponseBody) *UpdateDnsGtmAccessStrategyResponse {
	s.Body = v
	return s
}

func (s *UpdateDnsGtmAccessStrategyResponse) Validate() error {
	return dara.Validate(s)
}
