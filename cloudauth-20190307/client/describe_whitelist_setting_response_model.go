// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhitelistSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWhitelistSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWhitelistSettingResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWhitelistSettingResponseBody) *DescribeWhitelistSettingResponse
	GetBody() *DescribeWhitelistSettingResponseBody
}

type DescribeWhitelistSettingResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWhitelistSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWhitelistSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhitelistSettingResponse) GoString() string {
	return s.String()
}

func (s *DescribeWhitelistSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWhitelistSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWhitelistSettingResponse) GetBody() *DescribeWhitelistSettingResponseBody {
	return s.Body
}

func (s *DescribeWhitelistSettingResponse) SetHeaders(v map[string]*string) *DescribeWhitelistSettingResponse {
	s.Headers = v
	return s
}

func (s *DescribeWhitelistSettingResponse) SetStatusCode(v int32) *DescribeWhitelistSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWhitelistSettingResponse) SetBody(v *DescribeWhitelistSettingResponseBody) *DescribeWhitelistSettingResponse {
	s.Body = v
	return s
}

func (s *DescribeWhitelistSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
