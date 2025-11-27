// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInstanceTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRCInstanceTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRCInstanceTypesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRCInstanceTypesResponseBody) *DescribeRCInstanceTypesResponse
	GetBody() *DescribeRCInstanceTypesResponseBody
}

type DescribeRCInstanceTypesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRCInstanceTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRCInstanceTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceTypesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRCInstanceTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRCInstanceTypesResponse) GetBody() *DescribeRCInstanceTypesResponseBody {
	return s.Body
}

func (s *DescribeRCInstanceTypesResponse) SetHeaders(v map[string]*string) *DescribeRCInstanceTypesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRCInstanceTypesResponse) SetStatusCode(v int32) *DescribeRCInstanceTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRCInstanceTypesResponse) SetBody(v *DescribeRCInstanceTypesResponseBody) *DescribeRCInstanceTypesResponse {
	s.Body = v
	return s
}

func (s *DescribeRCInstanceTypesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
