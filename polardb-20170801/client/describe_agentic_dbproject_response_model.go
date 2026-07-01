// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAgenticDBProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAgenticDBProjectResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAgenticDBProjectResponseBody) *DescribeAgenticDBProjectResponse
	GetBody() *DescribeAgenticDBProjectResponseBody
}

type DescribeAgenticDBProjectResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAgenticDBProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAgenticDBProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBProjectResponse) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAgenticDBProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAgenticDBProjectResponse) GetBody() *DescribeAgenticDBProjectResponseBody {
	return s.Body
}

func (s *DescribeAgenticDBProjectResponse) SetHeaders(v map[string]*string) *DescribeAgenticDBProjectResponse {
	s.Headers = v
	return s
}

func (s *DescribeAgenticDBProjectResponse) SetStatusCode(v int32) *DescribeAgenticDBProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAgenticDBProjectResponse) SetBody(v *DescribeAgenticDBProjectResponseBody) *DescribeAgenticDBProjectResponse {
	s.Body = v
	return s
}

func (s *DescribeAgenticDBProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
