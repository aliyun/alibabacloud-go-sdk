// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTerminalSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTerminalSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTerminalSettingsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTerminalSettingsResponseBody) *DescribeTerminalSettingsResponse
	GetBody() *DescribeTerminalSettingsResponseBody
}

type DescribeTerminalSettingsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTerminalSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTerminalSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTerminalSettingsResponse) GoString() string {
	return s.String()
}

func (s *DescribeTerminalSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTerminalSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTerminalSettingsResponse) GetBody() *DescribeTerminalSettingsResponseBody {
	return s.Body
}

func (s *DescribeTerminalSettingsResponse) SetHeaders(v map[string]*string) *DescribeTerminalSettingsResponse {
	s.Headers = v
	return s
}

func (s *DescribeTerminalSettingsResponse) SetStatusCode(v int32) *DescribeTerminalSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTerminalSettingsResponse) SetBody(v *DescribeTerminalSettingsResponseBody) *DescribeTerminalSettingsResponse {
	s.Body = v
	return s
}

func (s *DescribeTerminalSettingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
