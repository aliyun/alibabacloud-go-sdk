// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeManualTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeManualTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeManualTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeManualTaskResponseBody) *DescribeManualTaskResponse
	GetBody() *DescribeManualTaskResponseBody
}

type DescribeManualTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeManualTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeManualTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeManualTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeManualTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeManualTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeManualTaskResponse) GetBody() *DescribeManualTaskResponseBody {
	return s.Body
}

func (s *DescribeManualTaskResponse) SetHeaders(v map[string]*string) *DescribeManualTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeManualTaskResponse) SetStatusCode(v int32) *DescribeManualTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeManualTaskResponse) SetBody(v *DescribeManualTaskResponseBody) *DescribeManualTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeManualTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
