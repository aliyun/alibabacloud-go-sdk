// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigInstanceSecurityGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigInstanceSecurityGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigInstanceSecurityGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ConfigInstanceSecurityGroupsResponseBody) *ConfigInstanceSecurityGroupsResponse
	GetBody() *ConfigInstanceSecurityGroupsResponseBody
}

type ConfigInstanceSecurityGroupsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigInstanceSecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigInstanceSecurityGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigInstanceSecurityGroupsResponse) GoString() string {
	return s.String()
}

func (s *ConfigInstanceSecurityGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigInstanceSecurityGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigInstanceSecurityGroupsResponse) GetBody() *ConfigInstanceSecurityGroupsResponseBody {
	return s.Body
}

func (s *ConfigInstanceSecurityGroupsResponse) SetHeaders(v map[string]*string) *ConfigInstanceSecurityGroupsResponse {
	s.Headers = v
	return s
}

func (s *ConfigInstanceSecurityGroupsResponse) SetStatusCode(v int32) *ConfigInstanceSecurityGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigInstanceSecurityGroupsResponse) SetBody(v *ConfigInstanceSecurityGroupsResponseBody) *ConfigInstanceSecurityGroupsResponse {
	s.Body = v
	return s
}

func (s *ConfigInstanceSecurityGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
