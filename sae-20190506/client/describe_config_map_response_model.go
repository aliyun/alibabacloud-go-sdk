// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConfigMapResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeConfigMapResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeConfigMapResponse
	GetStatusCode() *int32
	SetBody(v *DescribeConfigMapResponseBody) *DescribeConfigMapResponse
	GetBody() *DescribeConfigMapResponseBody
}

type DescribeConfigMapResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeConfigMapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeConfigMapResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigMapResponse) GoString() string {
	return s.String()
}

func (s *DescribeConfigMapResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeConfigMapResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeConfigMapResponse) GetBody() *DescribeConfigMapResponseBody {
	return s.Body
}

func (s *DescribeConfigMapResponse) SetHeaders(v map[string]*string) *DescribeConfigMapResponse {
	s.Headers = v
	return s
}

func (s *DescribeConfigMapResponse) SetStatusCode(v int32) *DescribeConfigMapResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConfigMapResponse) SetBody(v *DescribeConfigMapResponseBody) *DescribeConfigMapResponse {
	s.Body = v
	return s
}

func (s *DescribeConfigMapResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
