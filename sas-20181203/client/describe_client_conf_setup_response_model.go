// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientConfSetupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClientConfSetupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClientConfSetupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClientConfSetupResponseBody) *DescribeClientConfSetupResponse
	GetBody() *DescribeClientConfSetupResponseBody
}

type DescribeClientConfSetupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClientConfSetupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClientConfSetupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientConfSetupResponse) GoString() string {
	return s.String()
}

func (s *DescribeClientConfSetupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClientConfSetupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClientConfSetupResponse) GetBody() *DescribeClientConfSetupResponseBody {
	return s.Body
}

func (s *DescribeClientConfSetupResponse) SetHeaders(v map[string]*string) *DescribeClientConfSetupResponse {
	s.Headers = v
	return s
}

func (s *DescribeClientConfSetupResponse) SetStatusCode(v int32) *DescribeClientConfSetupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClientConfSetupResponse) SetBody(v *DescribeClientConfSetupResponseBody) *DescribeClientConfSetupResponse {
	s.Body = v
	return s
}

func (s *DescribeClientConfSetupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
