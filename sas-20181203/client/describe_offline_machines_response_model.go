// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOfflineMachinesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOfflineMachinesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOfflineMachinesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOfflineMachinesResponseBody) *DescribeOfflineMachinesResponse
	GetBody() *DescribeOfflineMachinesResponseBody
}

type DescribeOfflineMachinesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOfflineMachinesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOfflineMachinesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOfflineMachinesResponse) GoString() string {
	return s.String()
}

func (s *DescribeOfflineMachinesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOfflineMachinesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOfflineMachinesResponse) GetBody() *DescribeOfflineMachinesResponseBody {
	return s.Body
}

func (s *DescribeOfflineMachinesResponse) SetHeaders(v map[string]*string) *DescribeOfflineMachinesResponse {
	s.Headers = v
	return s
}

func (s *DescribeOfflineMachinesResponse) SetStatusCode(v int32) *DescribeOfflineMachinesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOfflineMachinesResponse) SetBody(v *DescribeOfflineMachinesResponseBody) *DescribeOfflineMachinesResponse {
	s.Body = v
	return s
}

func (s *DescribeOfflineMachinesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
