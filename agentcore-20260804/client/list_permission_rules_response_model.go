// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPermissionRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPermissionRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListPermissionRulesResponseBody) *ListPermissionRulesResponse
	GetBody() *ListPermissionRulesResponseBody
}

type ListPermissionRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPermissionRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPermissionRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionRulesResponse) GoString() string {
	return s.String()
}

func (s *ListPermissionRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPermissionRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPermissionRulesResponse) GetBody() *ListPermissionRulesResponseBody {
	return s.Body
}

func (s *ListPermissionRulesResponse) SetHeaders(v map[string]*string) *ListPermissionRulesResponse {
	s.Headers = v
	return s
}

func (s *ListPermissionRulesResponse) SetStatusCode(v int32) *ListPermissionRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPermissionRulesResponse) SetBody(v *ListPermissionRulesResponseBody) *ListPermissionRulesResponse {
	s.Body = v
	return s
}

func (s *ListPermissionRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
