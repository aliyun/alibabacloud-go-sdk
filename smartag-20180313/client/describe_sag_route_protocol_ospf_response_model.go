// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagRouteProtocolOspfResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSagRouteProtocolOspfResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSagRouteProtocolOspfResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSagRouteProtocolOspfResponseBody) *DescribeSagRouteProtocolOspfResponse
	GetBody() *DescribeSagRouteProtocolOspfResponseBody
}

type DescribeSagRouteProtocolOspfResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSagRouteProtocolOspfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSagRouteProtocolOspfResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagRouteProtocolOspfResponse) GoString() string {
	return s.String()
}

func (s *DescribeSagRouteProtocolOspfResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSagRouteProtocolOspfResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSagRouteProtocolOspfResponse) GetBody() *DescribeSagRouteProtocolOspfResponseBody {
	return s.Body
}

func (s *DescribeSagRouteProtocolOspfResponse) SetHeaders(v map[string]*string) *DescribeSagRouteProtocolOspfResponse {
	s.Headers = v
	return s
}

func (s *DescribeSagRouteProtocolOspfResponse) SetStatusCode(v int32) *DescribeSagRouteProtocolOspfResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSagRouteProtocolOspfResponse) SetBody(v *DescribeSagRouteProtocolOspfResponseBody) *DescribeSagRouteProtocolOspfResponse {
	s.Body = v
	return s
}

func (s *DescribeSagRouteProtocolOspfResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
