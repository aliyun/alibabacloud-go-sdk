// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCreateIndexJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCreateIndexJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCreateIndexJobResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCreateIndexJobResponseBody) *DescribeCreateIndexJobResponse
	GetBody() *DescribeCreateIndexJobResponseBody
}

type DescribeCreateIndexJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCreateIndexJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCreateIndexJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreateIndexJobResponse) GoString() string {
	return s.String()
}

func (s *DescribeCreateIndexJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCreateIndexJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCreateIndexJobResponse) GetBody() *DescribeCreateIndexJobResponseBody {
	return s.Body
}

func (s *DescribeCreateIndexJobResponse) SetHeaders(v map[string]*string) *DescribeCreateIndexJobResponse {
	s.Headers = v
	return s
}

func (s *DescribeCreateIndexJobResponse) SetStatusCode(v int32) *DescribeCreateIndexJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCreateIndexJobResponse) SetBody(v *DescribeCreateIndexJobResponseBody) *DescribeCreateIndexJobResponse {
	s.Body = v
	return s
}

func (s *DescribeCreateIndexJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
