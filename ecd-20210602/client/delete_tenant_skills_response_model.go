// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTenantSkillsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTenantSkillsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTenantSkillsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTenantSkillsResponseBody) *DeleteTenantSkillsResponse
	GetBody() *DeleteTenantSkillsResponseBody
}

type DeleteTenantSkillsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTenantSkillsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTenantSkillsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTenantSkillsResponse) GoString() string {
	return s.String()
}

func (s *DeleteTenantSkillsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTenantSkillsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTenantSkillsResponse) GetBody() *DeleteTenantSkillsResponseBody {
	return s.Body
}

func (s *DeleteTenantSkillsResponse) SetHeaders(v map[string]*string) *DeleteTenantSkillsResponse {
	s.Headers = v
	return s
}

func (s *DeleteTenantSkillsResponse) SetStatusCode(v int32) *DeleteTenantSkillsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTenantSkillsResponse) SetBody(v *DeleteTenantSkillsResponseBody) *DeleteTenantSkillsResponse {
	s.Body = v
	return s
}

func (s *DeleteTenantSkillsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
