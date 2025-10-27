// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExposedCheckWarningResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExposedCheckWarningResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExposedCheckWarningResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExposedCheckWarningResponseBody) *DescribeExposedCheckWarningResponse
	GetBody() *DescribeExposedCheckWarningResponseBody
}

type DescribeExposedCheckWarningResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExposedCheckWarningResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExposedCheckWarningResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExposedCheckWarningResponse) GoString() string {
	return s.String()
}

func (s *DescribeExposedCheckWarningResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExposedCheckWarningResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExposedCheckWarningResponse) GetBody() *DescribeExposedCheckWarningResponseBody {
	return s.Body
}

func (s *DescribeExposedCheckWarningResponse) SetHeaders(v map[string]*string) *DescribeExposedCheckWarningResponse {
	s.Headers = v
	return s
}

func (s *DescribeExposedCheckWarningResponse) SetStatusCode(v int32) *DescribeExposedCheckWarningResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExposedCheckWarningResponse) SetBody(v *DescribeExposedCheckWarningResponseBody) *DescribeExposedCheckWarningResponse {
	s.Body = v
	return s
}

func (s *DescribeExposedCheckWarningResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
