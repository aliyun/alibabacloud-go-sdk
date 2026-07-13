// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAtiAgentRegisterInfoMarketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchAtiAgentRegisterInfoMarketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchAtiAgentRegisterInfoMarketResponse
	GetStatusCode() *int32
	SetBody(v *SearchAtiAgentRegisterInfoMarketResponseBody) *SearchAtiAgentRegisterInfoMarketResponse
	GetBody() *SearchAtiAgentRegisterInfoMarketResponseBody
}

type SearchAtiAgentRegisterInfoMarketResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchAtiAgentRegisterInfoMarketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchAtiAgentRegisterInfoMarketResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchAtiAgentRegisterInfoMarketResponse) GoString() string {
	return s.String()
}

func (s *SearchAtiAgentRegisterInfoMarketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchAtiAgentRegisterInfoMarketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchAtiAgentRegisterInfoMarketResponse) GetBody() *SearchAtiAgentRegisterInfoMarketResponseBody {
	return s.Body
}

func (s *SearchAtiAgentRegisterInfoMarketResponse) SetHeaders(v map[string]*string) *SearchAtiAgentRegisterInfoMarketResponse {
	s.Headers = v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponse) SetStatusCode(v int32) *SearchAtiAgentRegisterInfoMarketResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponse) SetBody(v *SearchAtiAgentRegisterInfoMarketResponseBody) *SearchAtiAgentRegisterInfoMarketResponse {
	s.Body = v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
