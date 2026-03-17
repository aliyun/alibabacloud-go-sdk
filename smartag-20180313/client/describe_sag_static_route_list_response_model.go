// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagStaticRouteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSagStaticRouteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSagStaticRouteListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSagStaticRouteListResponseBody) *DescribeSagStaticRouteListResponse
	GetBody() *DescribeSagStaticRouteListResponseBody
}

type DescribeSagStaticRouteListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSagStaticRouteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSagStaticRouteListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagStaticRouteListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSagStaticRouteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSagStaticRouteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSagStaticRouteListResponse) GetBody() *DescribeSagStaticRouteListResponseBody {
	return s.Body
}

func (s *DescribeSagStaticRouteListResponse) SetHeaders(v map[string]*string) *DescribeSagStaticRouteListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSagStaticRouteListResponse) SetStatusCode(v int32) *DescribeSagStaticRouteListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSagStaticRouteListResponse) SetBody(v *DescribeSagStaticRouteListResponseBody) *DescribeSagStaticRouteListResponse {
	s.Body = v
	return s
}

func (s *DescribeSagStaticRouteListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
