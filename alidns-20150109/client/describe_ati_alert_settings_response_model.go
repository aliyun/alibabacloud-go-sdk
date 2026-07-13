// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAtiAlertSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAtiAlertSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAtiAlertSettingsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAtiAlertSettingsResponseBody) *DescribeAtiAlertSettingsResponse
	GetBody() *DescribeAtiAlertSettingsResponseBody
}

type DescribeAtiAlertSettingsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAtiAlertSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAtiAlertSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAtiAlertSettingsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAtiAlertSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAtiAlertSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAtiAlertSettingsResponse) GetBody() *DescribeAtiAlertSettingsResponseBody {
	return s.Body
}

func (s *DescribeAtiAlertSettingsResponse) SetHeaders(v map[string]*string) *DescribeAtiAlertSettingsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAtiAlertSettingsResponse) SetStatusCode(v int32) *DescribeAtiAlertSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAtiAlertSettingsResponse) SetBody(v *DescribeAtiAlertSettingsResponseBody) *DescribeAtiAlertSettingsResponse {
	s.Body = v
	return s
}

func (s *DescribeAtiAlertSettingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
