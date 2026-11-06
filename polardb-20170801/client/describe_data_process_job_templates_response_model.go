// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataProcessJobTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataProcessJobTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataProcessJobTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataProcessJobTemplatesResponseBody) *DescribeDataProcessJobTemplatesResponse
	GetBody() *DescribeDataProcessJobTemplatesResponseBody
}

type DescribeDataProcessJobTemplatesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataProcessJobTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataProcessJobTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataProcessJobTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataProcessJobTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataProcessJobTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataProcessJobTemplatesResponse) GetBody() *DescribeDataProcessJobTemplatesResponseBody {
	return s.Body
}

func (s *DescribeDataProcessJobTemplatesResponse) SetHeaders(v map[string]*string) *DescribeDataProcessJobTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataProcessJobTemplatesResponse) SetStatusCode(v int32) *DescribeDataProcessJobTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataProcessJobTemplatesResponse) SetBody(v *DescribeDataProcessJobTemplatesResponseBody) *DescribeDataProcessJobTemplatesResponse {
	s.Body = v
	return s
}

func (s *DescribeDataProcessJobTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
