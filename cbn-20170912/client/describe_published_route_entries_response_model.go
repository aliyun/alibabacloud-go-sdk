// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePublishedRouteEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePublishedRouteEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePublishedRouteEntriesResponse
	GetStatusCode() *int32
	SetBody(v *DescribePublishedRouteEntriesResponseBody) *DescribePublishedRouteEntriesResponse
	GetBody() *DescribePublishedRouteEntriesResponseBody
}

type DescribePublishedRouteEntriesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePublishedRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePublishedRouteEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePublishedRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *DescribePublishedRouteEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePublishedRouteEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePublishedRouteEntriesResponse) GetBody() *DescribePublishedRouteEntriesResponseBody {
	return s.Body
}

func (s *DescribePublishedRouteEntriesResponse) SetHeaders(v map[string]*string) *DescribePublishedRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *DescribePublishedRouteEntriesResponse) SetStatusCode(v int32) *DescribePublishedRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponse) SetBody(v *DescribePublishedRouteEntriesResponseBody) *DescribePublishedRouteEntriesResponse {
	s.Body = v
	return s
}

func (s *DescribePublishedRouteEntriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
