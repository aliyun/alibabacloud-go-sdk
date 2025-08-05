// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHanaRetentionSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHanaRetentionSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHanaRetentionSettingResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHanaRetentionSettingResponseBody) *DescribeHanaRetentionSettingResponse
	GetBody() *DescribeHanaRetentionSettingResponseBody
}

type DescribeHanaRetentionSettingResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHanaRetentionSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHanaRetentionSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaRetentionSettingResponse) GoString() string {
	return s.String()
}

func (s *DescribeHanaRetentionSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHanaRetentionSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHanaRetentionSettingResponse) GetBody() *DescribeHanaRetentionSettingResponseBody {
	return s.Body
}

func (s *DescribeHanaRetentionSettingResponse) SetHeaders(v map[string]*string) *DescribeHanaRetentionSettingResponse {
	s.Headers = v
	return s
}

func (s *DescribeHanaRetentionSettingResponse) SetStatusCode(v int32) *DescribeHanaRetentionSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHanaRetentionSettingResponse) SetBody(v *DescribeHanaRetentionSettingResponseBody) *DescribeHanaRetentionSettingResponse {
	s.Body = v
	return s
}

func (s *DescribeHanaRetentionSettingResponse) Validate() error {
	return dara.Validate(s)
}
