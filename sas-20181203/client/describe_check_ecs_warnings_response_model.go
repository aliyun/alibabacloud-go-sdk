// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCheckEcsWarningsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCheckEcsWarningsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCheckEcsWarningsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCheckEcsWarningsResponseBody) *DescribeCheckEcsWarningsResponse
	GetBody() *DescribeCheckEcsWarningsResponseBody
}

type DescribeCheckEcsWarningsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCheckEcsWarningsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCheckEcsWarningsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckEcsWarningsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCheckEcsWarningsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCheckEcsWarningsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCheckEcsWarningsResponse) GetBody() *DescribeCheckEcsWarningsResponseBody {
	return s.Body
}

func (s *DescribeCheckEcsWarningsResponse) SetHeaders(v map[string]*string) *DescribeCheckEcsWarningsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCheckEcsWarningsResponse) SetStatusCode(v int32) *DescribeCheckEcsWarningsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCheckEcsWarningsResponse) SetBody(v *DescribeCheckEcsWarningsResponseBody) *DescribeCheckEcsWarningsResponse {
	s.Body = v
	return s
}

func (s *DescribeCheckEcsWarningsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
