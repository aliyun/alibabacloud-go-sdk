// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeManualTaskInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeManualTaskInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeManualTaskInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeManualTaskInstanceResponseBody) *DescribeManualTaskInstanceResponse
	GetBody() *DescribeManualTaskInstanceResponseBody
}

type DescribeManualTaskInstanceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeManualTaskInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeManualTaskInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeManualTaskInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeManualTaskInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeManualTaskInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeManualTaskInstanceResponse) GetBody() *DescribeManualTaskInstanceResponseBody {
	return s.Body
}

func (s *DescribeManualTaskInstanceResponse) SetHeaders(v map[string]*string) *DescribeManualTaskInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeManualTaskInstanceResponse) SetStatusCode(v int32) *DescribeManualTaskInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeManualTaskInstanceResponse) SetBody(v *DescribeManualTaskInstanceResponseBody) *DescribeManualTaskInstanceResponse {
	s.Body = v
	return s
}

func (s *DescribeManualTaskInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
