// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagRouteProtocolBgpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSagRouteProtocolBgpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSagRouteProtocolBgpResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSagRouteProtocolBgpResponseBody) *DescribeSagRouteProtocolBgpResponse
	GetBody() *DescribeSagRouteProtocolBgpResponseBody
}

type DescribeSagRouteProtocolBgpResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSagRouteProtocolBgpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSagRouteProtocolBgpResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagRouteProtocolBgpResponse) GoString() string {
	return s.String()
}

func (s *DescribeSagRouteProtocolBgpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSagRouteProtocolBgpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSagRouteProtocolBgpResponse) GetBody() *DescribeSagRouteProtocolBgpResponseBody {
	return s.Body
}

func (s *DescribeSagRouteProtocolBgpResponse) SetHeaders(v map[string]*string) *DescribeSagRouteProtocolBgpResponse {
	s.Headers = v
	return s
}

func (s *DescribeSagRouteProtocolBgpResponse) SetStatusCode(v int32) *DescribeSagRouteProtocolBgpResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSagRouteProtocolBgpResponse) SetBody(v *DescribeSagRouteProtocolBgpResponseBody) *DescribeSagRouteProtocolBgpResponse {
	s.Body = v
	return s
}

func (s *DescribeSagRouteProtocolBgpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
