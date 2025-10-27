// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhiteListProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWhiteListProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWhiteListProcessResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWhiteListProcessResponseBody) *DescribeWhiteListProcessResponse
	GetBody() *DescribeWhiteListProcessResponseBody
}

type DescribeWhiteListProcessResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWhiteListProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWhiteListProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteListProcessResponse) GoString() string {
	return s.String()
}

func (s *DescribeWhiteListProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWhiteListProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWhiteListProcessResponse) GetBody() *DescribeWhiteListProcessResponseBody {
	return s.Body
}

func (s *DescribeWhiteListProcessResponse) SetHeaders(v map[string]*string) *DescribeWhiteListProcessResponse {
	s.Headers = v
	return s
}

func (s *DescribeWhiteListProcessResponse) SetStatusCode(v int32) *DescribeWhiteListProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWhiteListProcessResponse) SetBody(v *DescribeWhiteListProcessResponseBody) *DescribeWhiteListProcessResponse {
	s.Body = v
	return s
}

func (s *DescribeWhiteListProcessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
