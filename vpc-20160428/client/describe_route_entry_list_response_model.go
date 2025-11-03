// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouteEntryListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRouteEntryListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRouteEntryListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRouteEntryListResponseBody) *DescribeRouteEntryListResponse
	GetBody() *DescribeRouteEntryListResponseBody
}

type DescribeRouteEntryListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRouteEntryListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRouteEntryListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteEntryListResponse) GoString() string {
	return s.String()
}

func (s *DescribeRouteEntryListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRouteEntryListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRouteEntryListResponse) GetBody() *DescribeRouteEntryListResponseBody {
	return s.Body
}

func (s *DescribeRouteEntryListResponse) SetHeaders(v map[string]*string) *DescribeRouteEntryListResponse {
	s.Headers = v
	return s
}

func (s *DescribeRouteEntryListResponse) SetStatusCode(v int32) *DescribeRouteEntryListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRouteEntryListResponse) SetBody(v *DescribeRouteEntryListResponseBody) *DescribeRouteEntryListResponse {
	s.Body = v
	return s
}

func (s *DescribeRouteEntryListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
