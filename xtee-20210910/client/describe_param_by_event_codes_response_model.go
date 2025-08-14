// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParamByEventCodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeParamByEventCodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeParamByEventCodesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeParamByEventCodesResponseBody) *DescribeParamByEventCodesResponse
	GetBody() *DescribeParamByEventCodesResponseBody
}

type DescribeParamByEventCodesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeParamByEventCodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeParamByEventCodesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeParamByEventCodesResponse) GoString() string {
	return s.String()
}

func (s *DescribeParamByEventCodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeParamByEventCodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeParamByEventCodesResponse) GetBody() *DescribeParamByEventCodesResponseBody {
	return s.Body
}

func (s *DescribeParamByEventCodesResponse) SetHeaders(v map[string]*string) *DescribeParamByEventCodesResponse {
	s.Headers = v
	return s
}

func (s *DescribeParamByEventCodesResponse) SetStatusCode(v int32) *DescribeParamByEventCodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeParamByEventCodesResponse) SetBody(v *DescribeParamByEventCodesResponseBody) *DescribeParamByEventCodesResponse {
	s.Body = v
	return s
}

func (s *DescribeParamByEventCodesResponse) Validate() error {
	return dara.Validate(s)
}
