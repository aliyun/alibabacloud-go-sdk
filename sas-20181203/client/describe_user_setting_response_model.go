// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserSettingResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserSettingResponseBody) *DescribeUserSettingResponse
	GetBody() *DescribeUserSettingResponseBody
}

type DescribeUserSettingResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserSettingResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserSettingResponse) GetBody() *DescribeUserSettingResponseBody {
	return s.Body
}

func (s *DescribeUserSettingResponse) SetHeaders(v map[string]*string) *DescribeUserSettingResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserSettingResponse) SetStatusCode(v int32) *DescribeUserSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserSettingResponse) SetBody(v *DescribeUserSettingResponseBody) *DescribeUserSettingResponse {
	s.Body = v
	return s
}

func (s *DescribeUserSettingResponse) Validate() error {
	return dara.Validate(s)
}
