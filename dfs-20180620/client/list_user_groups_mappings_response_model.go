// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupsMappingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserGroupsMappingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserGroupsMappingsResponse
	GetStatusCode() *int32
	SetBody(v *ListUserGroupsMappingsResponseBody) *ListUserGroupsMappingsResponse
	GetBody() *ListUserGroupsMappingsResponseBody
}

type ListUserGroupsMappingsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserGroupsMappingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserGroupsMappingsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsMappingsResponse) GoString() string {
	return s.String()
}

func (s *ListUserGroupsMappingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserGroupsMappingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserGroupsMappingsResponse) GetBody() *ListUserGroupsMappingsResponseBody {
	return s.Body
}

func (s *ListUserGroupsMappingsResponse) SetHeaders(v map[string]*string) *ListUserGroupsMappingsResponse {
	s.Headers = v
	return s
}

func (s *ListUserGroupsMappingsResponse) SetStatusCode(v int32) *ListUserGroupsMappingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserGroupsMappingsResponse) SetBody(v *ListUserGroupsMappingsResponseBody) *ListUserGroupsMappingsResponse {
	s.Body = v
	return s
}

func (s *ListUserGroupsMappingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
