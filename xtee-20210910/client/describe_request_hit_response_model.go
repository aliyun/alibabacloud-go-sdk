// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRequestHitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRequestHitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRequestHitResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRequestHitResponseBody) *DescribeRequestHitResponse
	GetBody() *DescribeRequestHitResponseBody
}

type DescribeRequestHitResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRequestHitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRequestHitResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRequestHitResponse) GoString() string {
	return s.String()
}

func (s *DescribeRequestHitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRequestHitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRequestHitResponse) GetBody() *DescribeRequestHitResponseBody {
	return s.Body
}

func (s *DescribeRequestHitResponse) SetHeaders(v map[string]*string) *DescribeRequestHitResponse {
	s.Headers = v
	return s
}

func (s *DescribeRequestHitResponse) SetStatusCode(v int32) *DescribeRequestHitResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRequestHitResponse) SetBody(v *DescribeRequestHitResponseBody) *DescribeRequestHitResponse {
	s.Body = v
	return s
}

func (s *DescribeRequestHitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
