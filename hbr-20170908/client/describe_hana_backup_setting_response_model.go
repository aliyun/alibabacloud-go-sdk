// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHanaBackupSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHanaBackupSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHanaBackupSettingResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHanaBackupSettingResponseBody) *DescribeHanaBackupSettingResponse
	GetBody() *DescribeHanaBackupSettingResponseBody
}

type DescribeHanaBackupSettingResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHanaBackupSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHanaBackupSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaBackupSettingResponse) GoString() string {
	return s.String()
}

func (s *DescribeHanaBackupSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHanaBackupSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHanaBackupSettingResponse) GetBody() *DescribeHanaBackupSettingResponseBody {
	return s.Body
}

func (s *DescribeHanaBackupSettingResponse) SetHeaders(v map[string]*string) *DescribeHanaBackupSettingResponse {
	s.Headers = v
	return s
}

func (s *DescribeHanaBackupSettingResponse) SetStatusCode(v int32) *DescribeHanaBackupSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHanaBackupSettingResponse) SetBody(v *DescribeHanaBackupSettingResponseBody) *DescribeHanaBackupSettingResponse {
	s.Body = v
	return s
}

func (s *DescribeHanaBackupSettingResponse) Validate() error {
	return dara.Validate(s)
}
