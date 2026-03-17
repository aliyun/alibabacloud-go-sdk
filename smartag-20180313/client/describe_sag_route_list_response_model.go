// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagRouteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSagRouteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSagRouteListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSagRouteListResponseBody) *DescribeSagRouteListResponse
	GetBody() *DescribeSagRouteListResponseBody
}

type DescribeSagRouteListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSagRouteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSagRouteListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagRouteListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSagRouteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSagRouteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSagRouteListResponse) GetBody() *DescribeSagRouteListResponseBody {
	return s.Body
}

func (s *DescribeSagRouteListResponse) SetHeaders(v map[string]*string) *DescribeSagRouteListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSagRouteListResponse) SetStatusCode(v int32) *DescribeSagRouteListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSagRouteListResponse) SetBody(v *DescribeSagRouteListResponseBody) *DescribeSagRouteListResponse {
	s.Body = v
	return s
}

func (s *DescribeSagRouteListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
