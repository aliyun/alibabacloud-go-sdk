// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRemediationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRemediationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRemediationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRemediationResponseBody) *DescribeRemediationResponse
	GetBody() *DescribeRemediationResponseBody
}

type DescribeRemediationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRemediationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRemediationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRemediationResponse) GoString() string {
	return s.String()
}

func (s *DescribeRemediationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRemediationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRemediationResponse) GetBody() *DescribeRemediationResponseBody {
	return s.Body
}

func (s *DescribeRemediationResponse) SetHeaders(v map[string]*string) *DescribeRemediationResponse {
	s.Headers = v
	return s
}

func (s *DescribeRemediationResponse) SetStatusCode(v int32) *DescribeRemediationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRemediationResponse) SetBody(v *DescribeRemediationResponseBody) *DescribeRemediationResponse {
	s.Body = v
	return s
}

func (s *DescribeRemediationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
