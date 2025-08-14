// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExcuteNumResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExcuteNumResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExcuteNumResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExcuteNumResponseBody) *DescribeExcuteNumResponse
	GetBody() *DescribeExcuteNumResponseBody
}

type DescribeExcuteNumResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExcuteNumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExcuteNumResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExcuteNumResponse) GoString() string {
	return s.String()
}

func (s *DescribeExcuteNumResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExcuteNumResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExcuteNumResponse) GetBody() *DescribeExcuteNumResponseBody {
	return s.Body
}

func (s *DescribeExcuteNumResponse) SetHeaders(v map[string]*string) *DescribeExcuteNumResponse {
	s.Headers = v
	return s
}

func (s *DescribeExcuteNumResponse) SetStatusCode(v int32) *DescribeExcuteNumResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExcuteNumResponse) SetBody(v *DescribeExcuteNumResponseBody) *DescribeExcuteNumResponse {
	s.Body = v
	return s
}

func (s *DescribeExcuteNumResponse) Validate() error {
	return dara.Validate(s)
}
