// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeModelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeModelsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeModelsResponseBody) *DescribeModelsResponse
	GetBody() *DescribeModelsResponseBody
}

type DescribeModelsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeModelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeModelsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelsResponse) GoString() string {
	return s.String()
}

func (s *DescribeModelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeModelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeModelsResponse) GetBody() *DescribeModelsResponseBody {
	return s.Body
}

func (s *DescribeModelsResponse) SetHeaders(v map[string]*string) *DescribeModelsResponse {
	s.Headers = v
	return s
}

func (s *DescribeModelsResponse) SetStatusCode(v int32) *DescribeModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeModelsResponse) SetBody(v *DescribeModelsResponseBody) *DescribeModelsResponse {
	s.Body = v
	return s
}

func (s *DescribeModelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
