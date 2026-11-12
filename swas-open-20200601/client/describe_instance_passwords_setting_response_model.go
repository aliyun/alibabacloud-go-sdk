// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancePasswordsSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstancePasswordsSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstancePasswordsSettingResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstancePasswordsSettingResponseBody) *DescribeInstancePasswordsSettingResponse
	GetBody() *DescribeInstancePasswordsSettingResponseBody
}

type DescribeInstancePasswordsSettingResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstancePasswordsSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstancePasswordsSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancePasswordsSettingResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstancePasswordsSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstancePasswordsSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstancePasswordsSettingResponse) GetBody() *DescribeInstancePasswordsSettingResponseBody {
	return s.Body
}

func (s *DescribeInstancePasswordsSettingResponse) SetHeaders(v map[string]*string) *DescribeInstancePasswordsSettingResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstancePasswordsSettingResponse) SetStatusCode(v int32) *DescribeInstancePasswordsSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstancePasswordsSettingResponse) SetBody(v *DescribeInstancePasswordsSettingResponseBody) *DescribeInstancePasswordsSettingResponse {
	s.Body = v
	return s
}

func (s *DescribeInstancePasswordsSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
