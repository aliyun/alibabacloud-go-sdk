// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeListModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeListModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeListModelResponse
	GetStatusCode() *int32
	SetBody(v *DescribeListModelResponseBody) *DescribeListModelResponse
	GetBody() *DescribeListModelResponseBody
}

type DescribeListModelResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeListModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeListModelResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeListModelResponse) GoString() string {
	return s.String()
}

func (s *DescribeListModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeListModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeListModelResponse) GetBody() *DescribeListModelResponseBody {
	return s.Body
}

func (s *DescribeListModelResponse) SetHeaders(v map[string]*string) *DescribeListModelResponse {
	s.Headers = v
	return s
}

func (s *DescribeListModelResponse) SetStatusCode(v int32) *DescribeListModelResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeListModelResponse) SetBody(v *DescribeListModelResponseBody) *DescribeListModelResponse {
	s.Body = v
	return s
}

func (s *DescribeListModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
