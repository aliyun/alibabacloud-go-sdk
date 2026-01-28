// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizationRulesForGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAuthorizationRulesForGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAuthorizationRulesForGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListAuthorizationRulesForGroupResponseBody) *ListAuthorizationRulesForGroupResponse
	GetBody() *ListAuthorizationRulesForGroupResponseBody
}

type ListAuthorizationRulesForGroupResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthorizationRulesForGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthorizationRulesForGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationRulesForGroupResponse) GoString() string {
	return s.String()
}

func (s *ListAuthorizationRulesForGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAuthorizationRulesForGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAuthorizationRulesForGroupResponse) GetBody() *ListAuthorizationRulesForGroupResponseBody {
	return s.Body
}

func (s *ListAuthorizationRulesForGroupResponse) SetHeaders(v map[string]*string) *ListAuthorizationRulesForGroupResponse {
	s.Headers = v
	return s
}

func (s *ListAuthorizationRulesForGroupResponse) SetStatusCode(v int32) *ListAuthorizationRulesForGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthorizationRulesForGroupResponse) SetBody(v *ListAuthorizationRulesForGroupResponseBody) *ListAuthorizationRulesForGroupResponse {
	s.Body = v
	return s
}

func (s *ListAuthorizationRulesForGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
