// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagGlobalRouteProtocolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSagGlobalRouteProtocolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSagGlobalRouteProtocolResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSagGlobalRouteProtocolResponseBody) *DescribeSagGlobalRouteProtocolResponse
	GetBody() *DescribeSagGlobalRouteProtocolResponseBody
}

type DescribeSagGlobalRouteProtocolResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSagGlobalRouteProtocolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSagGlobalRouteProtocolResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagGlobalRouteProtocolResponse) GoString() string {
	return s.String()
}

func (s *DescribeSagGlobalRouteProtocolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSagGlobalRouteProtocolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSagGlobalRouteProtocolResponse) GetBody() *DescribeSagGlobalRouteProtocolResponseBody {
	return s.Body
}

func (s *DescribeSagGlobalRouteProtocolResponse) SetHeaders(v map[string]*string) *DescribeSagGlobalRouteProtocolResponse {
	s.Headers = v
	return s
}

func (s *DescribeSagGlobalRouteProtocolResponse) SetStatusCode(v int32) *DescribeSagGlobalRouteProtocolResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSagGlobalRouteProtocolResponse) SetBody(v *DescribeSagGlobalRouteProtocolResponseBody) *DescribeSagGlobalRouteProtocolResponse {
	s.Body = v
	return s
}

func (s *DescribeSagGlobalRouteProtocolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
