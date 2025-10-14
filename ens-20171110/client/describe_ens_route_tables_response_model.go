// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsRouteTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEnsRouteTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEnsRouteTablesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEnsRouteTablesResponseBody) *DescribeEnsRouteTablesResponse
	GetBody() *DescribeEnsRouteTablesResponseBody
}

type DescribeEnsRouteTablesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEnsRouteTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEnsRouteTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsRouteTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnsRouteTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEnsRouteTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEnsRouteTablesResponse) GetBody() *DescribeEnsRouteTablesResponseBody {
	return s.Body
}

func (s *DescribeEnsRouteTablesResponse) SetHeaders(v map[string]*string) *DescribeEnsRouteTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnsRouteTablesResponse) SetStatusCode(v int32) *DescribeEnsRouteTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEnsRouteTablesResponse) SetBody(v *DescribeEnsRouteTablesResponseBody) *DescribeEnsRouteTablesResponse {
	s.Body = v
	return s
}

func (s *DescribeEnsRouteTablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
