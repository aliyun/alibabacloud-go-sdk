// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortAutoCcStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePortAutoCcStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePortAutoCcStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribePortAutoCcStatusResponseBody) *DescribePortAutoCcStatusResponse
	GetBody() *DescribePortAutoCcStatusResponseBody
}

type DescribePortAutoCcStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePortAutoCcStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePortAutoCcStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePortAutoCcStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribePortAutoCcStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePortAutoCcStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePortAutoCcStatusResponse) GetBody() *DescribePortAutoCcStatusResponseBody {
	return s.Body
}

func (s *DescribePortAutoCcStatusResponse) SetHeaders(v map[string]*string) *DescribePortAutoCcStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribePortAutoCcStatusResponse) SetStatusCode(v int32) *DescribePortAutoCcStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePortAutoCcStatusResponse) SetBody(v *DescribePortAutoCcStatusResponseBody) *DescribePortAutoCcStatusResponse {
	s.Body = v
	return s
}

func (s *DescribePortAutoCcStatusResponse) Validate() error {
	return dara.Validate(s)
}
