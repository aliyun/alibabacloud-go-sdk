// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePatternTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePatternTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePatternTypesResponse
	GetStatusCode() *int32
	SetBody(v *DescribePatternTypesResponseBody) *DescribePatternTypesResponse
	GetBody() *DescribePatternTypesResponseBody
}

type DescribePatternTypesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePatternTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePatternTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePatternTypesResponse) GoString() string {
	return s.String()
}

func (s *DescribePatternTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePatternTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePatternTypesResponse) GetBody() *DescribePatternTypesResponseBody {
	return s.Body
}

func (s *DescribePatternTypesResponse) SetHeaders(v map[string]*string) *DescribePatternTypesResponse {
	s.Headers = v
	return s
}

func (s *DescribePatternTypesResponse) SetStatusCode(v int32) *DescribePatternTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePatternTypesResponse) SetBody(v *DescribePatternTypesResponseBody) *DescribePatternTypesResponse {
	s.Body = v
	return s
}

func (s *DescribePatternTypesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
