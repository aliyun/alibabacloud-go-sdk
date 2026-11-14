// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWarmPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWarmPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWarmPoolResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWarmPoolResponseBody) *DescribeWarmPoolResponse
	GetBody() *DescribeWarmPoolResponseBody
}

type DescribeWarmPoolResponse struct {
	Headers    map[string]*string            `json:"headers" xml:"headers"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWarmPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWarmPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWarmPoolResponse) GoString() string {
	return s.String()
}

func (s *DescribeWarmPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWarmPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWarmPoolResponse) GetBody() *DescribeWarmPoolResponseBody {
	return s.Body
}

func (s *DescribeWarmPoolResponse) SetHeaders(v map[string]*string) *DescribeWarmPoolResponse {
	s.Headers = v
	return s
}

func (s *DescribeWarmPoolResponse) SetStatusCode(v int32) *DescribeWarmPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWarmPoolResponse) SetBody(v *DescribeWarmPoolResponseBody) *DescribeWarmPoolResponse {
	s.Body = v
	return s
}

func (s *DescribeWarmPoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
