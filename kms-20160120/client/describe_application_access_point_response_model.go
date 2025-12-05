// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationAccessPointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationAccessPointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationAccessPointResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationAccessPointResponseBody) *DescribeApplicationAccessPointResponse
	GetBody() *DescribeApplicationAccessPointResponseBody
}

type DescribeApplicationAccessPointResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationAccessPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationAccessPointResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAccessPointResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAccessPointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationAccessPointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationAccessPointResponse) GetBody() *DescribeApplicationAccessPointResponseBody {
	return s.Body
}

func (s *DescribeApplicationAccessPointResponse) SetHeaders(v map[string]*string) *DescribeApplicationAccessPointResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationAccessPointResponse) SetStatusCode(v int32) *DescribeApplicationAccessPointResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationAccessPointResponse) SetBody(v *DescribeApplicationAccessPointResponseBody) *DescribeApplicationAccessPointResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationAccessPointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
