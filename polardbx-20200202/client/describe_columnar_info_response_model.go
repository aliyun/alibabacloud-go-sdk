// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeColumnarInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeColumnarInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeColumnarInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeColumnarInfoResponseBody) *DescribeColumnarInfoResponse
	GetBody() *DescribeColumnarInfoResponseBody
}

type DescribeColumnarInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeColumnarInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeColumnarInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeColumnarInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeColumnarInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeColumnarInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeColumnarInfoResponse) GetBody() *DescribeColumnarInfoResponseBody {
	return s.Body
}

func (s *DescribeColumnarInfoResponse) SetHeaders(v map[string]*string) *DescribeColumnarInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeColumnarInfoResponse) SetStatusCode(v int32) *DescribeColumnarInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeColumnarInfoResponse) SetBody(v *DescribeColumnarInfoResponseBody) *DescribeColumnarInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeColumnarInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
