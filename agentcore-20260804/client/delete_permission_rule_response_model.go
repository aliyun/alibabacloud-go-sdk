// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePermissionRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePermissionRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePermissionRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeletePermissionRuleResponseBody) *DeletePermissionRuleResponse
	GetBody() *DeletePermissionRuleResponseBody
}

type DeletePermissionRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePermissionRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePermissionRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePermissionRuleResponse) GoString() string {
	return s.String()
}

func (s *DeletePermissionRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePermissionRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePermissionRuleResponse) GetBody() *DeletePermissionRuleResponseBody {
	return s.Body
}

func (s *DeletePermissionRuleResponse) SetHeaders(v map[string]*string) *DeletePermissionRuleResponse {
	s.Headers = v
	return s
}

func (s *DeletePermissionRuleResponse) SetStatusCode(v int32) *DeletePermissionRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePermissionRuleResponse) SetBody(v *DeletePermissionRuleResponseBody) *DeletePermissionRuleResponse {
	s.Body = v
	return s
}

func (s *DeletePermissionRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
