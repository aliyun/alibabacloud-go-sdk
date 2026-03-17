// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagPortRouteProtocolListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSagPortRouteProtocolListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSagPortRouteProtocolListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSagPortRouteProtocolListResponseBody) *DescribeSagPortRouteProtocolListResponse
	GetBody() *DescribeSagPortRouteProtocolListResponseBody
}

type DescribeSagPortRouteProtocolListResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSagPortRouteProtocolListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSagPortRouteProtocolListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagPortRouteProtocolListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSagPortRouteProtocolListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSagPortRouteProtocolListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSagPortRouteProtocolListResponse) GetBody() *DescribeSagPortRouteProtocolListResponseBody {
	return s.Body
}

func (s *DescribeSagPortRouteProtocolListResponse) SetHeaders(v map[string]*string) *DescribeSagPortRouteProtocolListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSagPortRouteProtocolListResponse) SetStatusCode(v int32) *DescribeSagPortRouteProtocolListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSagPortRouteProtocolListResponse) SetBody(v *DescribeSagPortRouteProtocolListResponseBody) *DescribeSagPortRouteProtocolListResponse {
	s.Body = v
	return s
}

func (s *DescribeSagPortRouteProtocolListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
