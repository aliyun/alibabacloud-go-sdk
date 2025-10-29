// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStreamLocationBlockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStreamLocationBlockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStreamLocationBlockResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStreamLocationBlockResponseBody) *DescribeStreamLocationBlockResponse
	GetBody() *DescribeStreamLocationBlockResponseBody
}

type DescribeStreamLocationBlockResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStreamLocationBlockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStreamLocationBlockResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamLocationBlockResponse) GoString() string {
	return s.String()
}

func (s *DescribeStreamLocationBlockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStreamLocationBlockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStreamLocationBlockResponse) GetBody() *DescribeStreamLocationBlockResponseBody {
	return s.Body
}

func (s *DescribeStreamLocationBlockResponse) SetHeaders(v map[string]*string) *DescribeStreamLocationBlockResponse {
	s.Headers = v
	return s
}

func (s *DescribeStreamLocationBlockResponse) SetStatusCode(v int32) *DescribeStreamLocationBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStreamLocationBlockResponse) SetBody(v *DescribeStreamLocationBlockResponseBody) *DescribeStreamLocationBlockResponse {
	s.Body = v
	return s
}

func (s *DescribeStreamLocationBlockResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
