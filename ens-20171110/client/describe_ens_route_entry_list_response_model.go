// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsRouteEntryListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEnsRouteEntryListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEnsRouteEntryListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEnsRouteEntryListResponseBody) *DescribeEnsRouteEntryListResponse
	GetBody() *DescribeEnsRouteEntryListResponseBody
}

type DescribeEnsRouteEntryListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEnsRouteEntryListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEnsRouteEntryListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsRouteEntryListResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnsRouteEntryListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEnsRouteEntryListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEnsRouteEntryListResponse) GetBody() *DescribeEnsRouteEntryListResponseBody {
	return s.Body
}

func (s *DescribeEnsRouteEntryListResponse) SetHeaders(v map[string]*string) *DescribeEnsRouteEntryListResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnsRouteEntryListResponse) SetStatusCode(v int32) *DescribeEnsRouteEntryListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEnsRouteEntryListResponse) SetBody(v *DescribeEnsRouteEntryListResponseBody) *DescribeEnsRouteEntryListResponse {
	s.Body = v
	return s
}

func (s *DescribeEnsRouteEntryListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
