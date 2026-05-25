// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNumberHLRResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNumberHLRResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNumberHLRResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNumberHLRResponseBody) *DescribeNumberHLRResponse
	GetBody() *DescribeNumberHLRResponseBody
}

type DescribeNumberHLRResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNumberHLRResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNumberHLRResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNumberHLRResponse) GoString() string {
	return s.String()
}

func (s *DescribeNumberHLRResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNumberHLRResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNumberHLRResponse) GetBody() *DescribeNumberHLRResponseBody {
	return s.Body
}

func (s *DescribeNumberHLRResponse) SetHeaders(v map[string]*string) *DescribeNumberHLRResponse {
	s.Headers = v
	return s
}

func (s *DescribeNumberHLRResponse) SetStatusCode(v int32) *DescribeNumberHLRResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNumberHLRResponse) SetBody(v *DescribeNumberHLRResponseBody) *DescribeNumberHLRResponse {
	s.Body = v
	return s
}

func (s *DescribeNumberHLRResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
