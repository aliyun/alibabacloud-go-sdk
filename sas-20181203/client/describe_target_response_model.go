// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTargetResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTargetResponseBody) *DescribeTargetResponse
	GetBody() *DescribeTargetResponseBody
}

type DescribeTargetResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTargetResponse) GoString() string {
	return s.String()
}

func (s *DescribeTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTargetResponse) GetBody() *DescribeTargetResponseBody {
	return s.Body
}

func (s *DescribeTargetResponse) SetHeaders(v map[string]*string) *DescribeTargetResponse {
	s.Headers = v
	return s
}

func (s *DescribeTargetResponse) SetStatusCode(v int32) *DescribeTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTargetResponse) SetBody(v *DescribeTargetResponseBody) *DescribeTargetResponse {
	s.Body = v
	return s
}

func (s *DescribeTargetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
