// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDnsGtmAccessModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDnsGtmAccessModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDnsGtmAccessModeResponse
	GetStatusCode() *int32
	SetBody(v *SetDnsGtmAccessModeResponseBody) *SetDnsGtmAccessModeResponse
	GetBody() *SetDnsGtmAccessModeResponseBody
}

type SetDnsGtmAccessModeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDnsGtmAccessModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDnsGtmAccessModeResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDnsGtmAccessModeResponse) GoString() string {
	return s.String()
}

func (s *SetDnsGtmAccessModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDnsGtmAccessModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDnsGtmAccessModeResponse) GetBody() *SetDnsGtmAccessModeResponseBody {
	return s.Body
}

func (s *SetDnsGtmAccessModeResponse) SetHeaders(v map[string]*string) *SetDnsGtmAccessModeResponse {
	s.Headers = v
	return s
}

func (s *SetDnsGtmAccessModeResponse) SetStatusCode(v int32) *SetDnsGtmAccessModeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDnsGtmAccessModeResponse) SetBody(v *SetDnsGtmAccessModeResponseBody) *SetDnsGtmAccessModeResponse {
	s.Body = v
	return s
}

func (s *SetDnsGtmAccessModeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
