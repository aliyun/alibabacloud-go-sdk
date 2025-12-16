// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataCollctionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataCollctionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataCollctionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataCollctionResponseBody) *DescribeDataCollctionResponse
	GetBody() *DescribeDataCollctionResponseBody
}

type DescribeDataCollctionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataCollctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataCollctionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataCollctionResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataCollctionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataCollctionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataCollctionResponse) GetBody() *DescribeDataCollctionResponseBody {
	return s.Body
}

func (s *DescribeDataCollctionResponse) SetHeaders(v map[string]*string) *DescribeDataCollctionResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataCollctionResponse) SetStatusCode(v int32) *DescribeDataCollctionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataCollctionResponse) SetBody(v *DescribeDataCollctionResponseBody) *DescribeDataCollctionResponse {
	s.Body = v
	return s
}

func (s *DescribeDataCollctionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
