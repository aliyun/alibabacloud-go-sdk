// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssV2ResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOssV2ResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOssV2ResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOssV2ResultResponseBody) *DescribeOssV2ResultResponse
	GetBody() *DescribeOssV2ResultResponseBody
}

type DescribeOssV2ResultResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOssV2ResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOssV2ResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssV2ResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeOssV2ResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOssV2ResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOssV2ResultResponse) GetBody() *DescribeOssV2ResultResponseBody {
	return s.Body
}

func (s *DescribeOssV2ResultResponse) SetHeaders(v map[string]*string) *DescribeOssV2ResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeOssV2ResultResponse) SetStatusCode(v int32) *DescribeOssV2ResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOssV2ResultResponse) SetBody(v *DescribeOssV2ResultResponseBody) *DescribeOssV2ResultResponse {
	s.Body = v
	return s
}

func (s *DescribeOssV2ResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
