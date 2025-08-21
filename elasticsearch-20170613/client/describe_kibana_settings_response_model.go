// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKibanaSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeKibanaSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeKibanaSettingsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeKibanaSettingsResponseBody) *DescribeKibanaSettingsResponse
	GetBody() *DescribeKibanaSettingsResponseBody
}

type DescribeKibanaSettingsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeKibanaSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeKibanaSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeKibanaSettingsResponse) GoString() string {
	return s.String()
}

func (s *DescribeKibanaSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeKibanaSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeKibanaSettingsResponse) GetBody() *DescribeKibanaSettingsResponseBody {
	return s.Body
}

func (s *DescribeKibanaSettingsResponse) SetHeaders(v map[string]*string) *DescribeKibanaSettingsResponse {
	s.Headers = v
	return s
}

func (s *DescribeKibanaSettingsResponse) SetStatusCode(v int32) *DescribeKibanaSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKibanaSettingsResponse) SetBody(v *DescribeKibanaSettingsResponseBody) *DescribeKibanaSettingsResponse {
	s.Body = v
	return s
}

func (s *DescribeKibanaSettingsResponse) Validate() error {
	return dara.Validate(s)
}
