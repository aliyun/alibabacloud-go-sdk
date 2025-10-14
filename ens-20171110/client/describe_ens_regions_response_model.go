// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEnsRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEnsRegionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEnsRegionsResponseBody) *DescribeEnsRegionsResponse
	GetBody() *DescribeEnsRegionsResponseBody
}

type DescribeEnsRegionsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEnsRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEnsRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEnsRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEnsRegionsResponse) GetBody() *DescribeEnsRegionsResponseBody {
	return s.Body
}

func (s *DescribeEnsRegionsResponse) SetHeaders(v map[string]*string) *DescribeEnsRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnsRegionsResponse) SetStatusCode(v int32) *DescribeEnsRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEnsRegionsResponse) SetBody(v *DescribeEnsRegionsResponseBody) *DescribeEnsRegionsResponse {
	s.Body = v
	return s
}

func (s *DescribeEnsRegionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
