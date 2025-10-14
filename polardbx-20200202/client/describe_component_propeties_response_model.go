// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComponentPropetiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeComponentPropetiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeComponentPropetiesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeComponentPropetiesResponseBody) *DescribeComponentPropetiesResponse
	GetBody() *DescribeComponentPropetiesResponseBody
}

type DescribeComponentPropetiesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeComponentPropetiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeComponentPropetiesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentPropetiesResponse) GoString() string {
	return s.String()
}

func (s *DescribeComponentPropetiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeComponentPropetiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeComponentPropetiesResponse) GetBody() *DescribeComponentPropetiesResponseBody {
	return s.Body
}

func (s *DescribeComponentPropetiesResponse) SetHeaders(v map[string]*string) *DescribeComponentPropetiesResponse {
	s.Headers = v
	return s
}

func (s *DescribeComponentPropetiesResponse) SetStatusCode(v int32) *DescribeComponentPropetiesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComponentPropetiesResponse) SetBody(v *DescribeComponentPropetiesResponseBody) *DescribeComponentPropetiesResponse {
	s.Body = v
	return s
}

func (s *DescribeComponentPropetiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
