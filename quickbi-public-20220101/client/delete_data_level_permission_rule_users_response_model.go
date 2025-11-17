// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataLevelPermissionRuleUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataLevelPermissionRuleUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataLevelPermissionRuleUsersResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataLevelPermissionRuleUsersResponseBody) *DeleteDataLevelPermissionRuleUsersResponse
	GetBody() *DeleteDataLevelPermissionRuleUsersResponseBody
}

type DeleteDataLevelPermissionRuleUsersResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataLevelPermissionRuleUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataLevelPermissionRuleUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataLevelPermissionRuleUsersResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataLevelPermissionRuleUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataLevelPermissionRuleUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataLevelPermissionRuleUsersResponse) GetBody() *DeleteDataLevelPermissionRuleUsersResponseBody {
	return s.Body
}

func (s *DeleteDataLevelPermissionRuleUsersResponse) SetHeaders(v map[string]*string) *DeleteDataLevelPermissionRuleUsersResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataLevelPermissionRuleUsersResponse) SetStatusCode(v int32) *DeleteDataLevelPermissionRuleUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataLevelPermissionRuleUsersResponse) SetBody(v *DeleteDataLevelPermissionRuleUsersResponseBody) *DeleteDataLevelPermissionRuleUsersResponse {
	s.Body = v
	return s
}

func (s *DeleteDataLevelPermissionRuleUsersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
