// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdsResourceSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRdsResourceSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRdsResourceSettingsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRdsResourceSettingsResponseBody) *DescribeRdsResourceSettingsResponse
	GetBody() *DescribeRdsResourceSettingsResponseBody
}

type DescribeRdsResourceSettingsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRdsResourceSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRdsResourceSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsResourceSettingsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRdsResourceSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRdsResourceSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRdsResourceSettingsResponse) GetBody() *DescribeRdsResourceSettingsResponseBody {
	return s.Body
}

func (s *DescribeRdsResourceSettingsResponse) SetHeaders(v map[string]*string) *DescribeRdsResourceSettingsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRdsResourceSettingsResponse) SetStatusCode(v int32) *DescribeRdsResourceSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRdsResourceSettingsResponse) SetBody(v *DescribeRdsResourceSettingsResponseBody) *DescribeRdsResourceSettingsResponse {
	s.Body = v
	return s
}

func (s *DescribeRdsResourceSettingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
