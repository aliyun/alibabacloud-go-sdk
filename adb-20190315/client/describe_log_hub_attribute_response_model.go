// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogHubAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLogHubAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLogHubAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLogHubAttributeResponseBody) *DescribeLogHubAttributeResponse
	GetBody() *DescribeLogHubAttributeResponseBody
}

type DescribeLogHubAttributeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLogHubAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLogHubAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogHubAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogHubAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLogHubAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLogHubAttributeResponse) GetBody() *DescribeLogHubAttributeResponseBody {
	return s.Body
}

func (s *DescribeLogHubAttributeResponse) SetHeaders(v map[string]*string) *DescribeLogHubAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogHubAttributeResponse) SetStatusCode(v int32) *DescribeLogHubAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogHubAttributeResponse) SetBody(v *DescribeLogHubAttributeResponseBody) *DescribeLogHubAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeLogHubAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
