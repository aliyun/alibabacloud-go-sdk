// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResubmitConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResubmitConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResubmitConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResubmitConfigResponseBody) *DescribeResubmitConfigResponse
	GetBody() *DescribeResubmitConfigResponseBody
}

type DescribeResubmitConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResubmitConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResubmitConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResubmitConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeResubmitConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResubmitConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResubmitConfigResponse) GetBody() *DescribeResubmitConfigResponseBody {
	return s.Body
}

func (s *DescribeResubmitConfigResponse) SetHeaders(v map[string]*string) *DescribeResubmitConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeResubmitConfigResponse) SetStatusCode(v int32) *DescribeResubmitConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResubmitConfigResponse) SetBody(v *DescribeResubmitConfigResponseBody) *DescribeResubmitConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeResubmitConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
