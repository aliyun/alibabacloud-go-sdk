// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCheckWarningsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCheckWarningsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCheckWarningsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCheckWarningsResponseBody) *DescribeCheckWarningsResponse
	GetBody() *DescribeCheckWarningsResponseBody
}

type DescribeCheckWarningsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCheckWarningsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCheckWarningsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckWarningsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCheckWarningsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCheckWarningsResponse) GetBody() *DescribeCheckWarningsResponseBody {
	return s.Body
}

func (s *DescribeCheckWarningsResponse) SetHeaders(v map[string]*string) *DescribeCheckWarningsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCheckWarningsResponse) SetStatusCode(v int32) *DescribeCheckWarningsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCheckWarningsResponse) SetBody(v *DescribeCheckWarningsResponseBody) *DescribeCheckWarningsResponse {
	s.Body = v
	return s
}

func (s *DescribeCheckWarningsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
