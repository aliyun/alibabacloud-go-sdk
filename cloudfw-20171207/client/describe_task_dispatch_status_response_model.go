// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskDispatchStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTaskDispatchStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTaskDispatchStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTaskDispatchStatusResponseBody) *DescribeTaskDispatchStatusResponse
	GetBody() *DescribeTaskDispatchStatusResponseBody
}

type DescribeTaskDispatchStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTaskDispatchStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTaskDispatchStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskDispatchStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeTaskDispatchStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTaskDispatchStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTaskDispatchStatusResponse) GetBody() *DescribeTaskDispatchStatusResponseBody {
	return s.Body
}

func (s *DescribeTaskDispatchStatusResponse) SetHeaders(v map[string]*string) *DescribeTaskDispatchStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeTaskDispatchStatusResponse) SetStatusCode(v int32) *DescribeTaskDispatchStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTaskDispatchStatusResponse) SetBody(v *DescribeTaskDispatchStatusResponseBody) *DescribeTaskDispatchStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeTaskDispatchStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
