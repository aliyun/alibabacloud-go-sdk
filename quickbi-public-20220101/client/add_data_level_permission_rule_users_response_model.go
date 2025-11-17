// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataLevelPermissionRuleUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDataLevelPermissionRuleUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDataLevelPermissionRuleUsersResponse
	GetStatusCode() *int32
	SetBody(v *AddDataLevelPermissionRuleUsersResponseBody) *AddDataLevelPermissionRuleUsersResponse
	GetBody() *AddDataLevelPermissionRuleUsersResponseBody
}

type AddDataLevelPermissionRuleUsersResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDataLevelPermissionRuleUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDataLevelPermissionRuleUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDataLevelPermissionRuleUsersResponse) GoString() string {
	return s.String()
}

func (s *AddDataLevelPermissionRuleUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDataLevelPermissionRuleUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDataLevelPermissionRuleUsersResponse) GetBody() *AddDataLevelPermissionRuleUsersResponseBody {
	return s.Body
}

func (s *AddDataLevelPermissionRuleUsersResponse) SetHeaders(v map[string]*string) *AddDataLevelPermissionRuleUsersResponse {
	s.Headers = v
	return s
}

func (s *AddDataLevelPermissionRuleUsersResponse) SetStatusCode(v int32) *AddDataLevelPermissionRuleUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDataLevelPermissionRuleUsersResponse) SetBody(v *AddDataLevelPermissionRuleUsersResponseBody) *AddDataLevelPermissionRuleUsersResponse {
	s.Body = v
	return s
}

func (s *AddDataLevelPermissionRuleUsersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
