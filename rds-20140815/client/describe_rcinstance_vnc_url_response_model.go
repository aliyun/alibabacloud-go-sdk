// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInstanceVncUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRCInstanceVncUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRCInstanceVncUrlResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRCInstanceVncUrlResponseBody) *DescribeRCInstanceVncUrlResponse
	GetBody() *DescribeRCInstanceVncUrlResponseBody
}

type DescribeRCInstanceVncUrlResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRCInstanceVncUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRCInstanceVncUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceVncUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceVncUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRCInstanceVncUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRCInstanceVncUrlResponse) GetBody() *DescribeRCInstanceVncUrlResponseBody {
	return s.Body
}

func (s *DescribeRCInstanceVncUrlResponse) SetHeaders(v map[string]*string) *DescribeRCInstanceVncUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeRCInstanceVncUrlResponse) SetStatusCode(v int32) *DescribeRCInstanceVncUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRCInstanceVncUrlResponse) SetBody(v *DescribeRCInstanceVncUrlResponseBody) *DescribeRCInstanceVncUrlResponse {
	s.Body = v
	return s
}

func (s *DescribeRCInstanceVncUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
