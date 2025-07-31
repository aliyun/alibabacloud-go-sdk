// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveOperationTaskTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeActiveOperationTaskTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeActiveOperationTaskTypeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeActiveOperationTaskTypeResponseBody) *DescribeActiveOperationTaskTypeResponse
	GetBody() *DescribeActiveOperationTaskTypeResponseBody
}

type DescribeActiveOperationTaskTypeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeActiveOperationTaskTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeActiveOperationTaskTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationTaskTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeActiveOperationTaskTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeActiveOperationTaskTypeResponse) GetBody() *DescribeActiveOperationTaskTypeResponseBody {
	return s.Body
}

func (s *DescribeActiveOperationTaskTypeResponse) SetHeaders(v map[string]*string) *DescribeActiveOperationTaskTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeActiveOperationTaskTypeResponse) SetStatusCode(v int32) *DescribeActiveOperationTaskTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeResponse) SetBody(v *DescribeActiveOperationTaskTypeResponseBody) *DescribeActiveOperationTaskTypeResponse {
	s.Body = v
	return s
}

func (s *DescribeActiveOperationTaskTypeResponse) Validate() error {
	return dara.Validate(s)
}
