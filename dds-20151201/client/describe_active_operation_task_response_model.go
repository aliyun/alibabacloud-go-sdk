// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveOperationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeActiveOperationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeActiveOperationTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeActiveOperationTaskResponseBody) *DescribeActiveOperationTaskResponse
	GetBody() *DescribeActiveOperationTaskResponseBody
}

type DescribeActiveOperationTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeActiveOperationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeActiveOperationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeActiveOperationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeActiveOperationTaskResponse) GetBody() *DescribeActiveOperationTaskResponseBody {
	return s.Body
}

func (s *DescribeActiveOperationTaskResponse) SetHeaders(v map[string]*string) *DescribeActiveOperationTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeActiveOperationTaskResponse) SetStatusCode(v int32) *DescribeActiveOperationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeActiveOperationTaskResponse) SetBody(v *DescribeActiveOperationTaskResponseBody) *DescribeActiveOperationTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeActiveOperationTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
