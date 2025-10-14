// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveOperationTaskCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeActiveOperationTaskCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeActiveOperationTaskCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeActiveOperationTaskCountResponseBody) *DescribeActiveOperationTaskCountResponse
	GetBody() *DescribeActiveOperationTaskCountResponseBody
}

type DescribeActiveOperationTaskCountResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeActiveOperationTaskCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeActiveOperationTaskCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationTaskCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeActiveOperationTaskCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeActiveOperationTaskCountResponse) GetBody() *DescribeActiveOperationTaskCountResponseBody {
	return s.Body
}

func (s *DescribeActiveOperationTaskCountResponse) SetHeaders(v map[string]*string) *DescribeActiveOperationTaskCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeActiveOperationTaskCountResponse) SetStatusCode(v int32) *DescribeActiveOperationTaskCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeActiveOperationTaskCountResponse) SetBody(v *DescribeActiveOperationTaskCountResponseBody) *DescribeActiveOperationTaskCountResponse {
	s.Body = v
	return s
}

func (s *DescribeActiveOperationTaskCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
