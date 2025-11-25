// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebAccessLogDispatchStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebAccessLogDispatchStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebAccessLogDispatchStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWebAccessLogDispatchStatusResponseBody) *DescribeWebAccessLogDispatchStatusResponse
	GetBody() *DescribeWebAccessLogDispatchStatusResponseBody
}

type DescribeWebAccessLogDispatchStatusResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebAccessLogDispatchStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebAccessLogDispatchStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebAccessLogDispatchStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebAccessLogDispatchStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebAccessLogDispatchStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebAccessLogDispatchStatusResponse) GetBody() *DescribeWebAccessLogDispatchStatusResponseBody {
	return s.Body
}

func (s *DescribeWebAccessLogDispatchStatusResponse) SetHeaders(v map[string]*string) *DescribeWebAccessLogDispatchStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebAccessLogDispatchStatusResponse) SetStatusCode(v int32) *DescribeWebAccessLogDispatchStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebAccessLogDispatchStatusResponse) SetBody(v *DescribeWebAccessLogDispatchStatusResponseBody) *DescribeWebAccessLogDispatchStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeWebAccessLogDispatchStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
