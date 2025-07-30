// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRuleToSchemeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRuleToSchemeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRuleToSchemeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRuleToSchemeResponseBody) *UpdateRuleToSchemeResponse
	GetBody() *UpdateRuleToSchemeResponseBody
}

type UpdateRuleToSchemeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRuleToSchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRuleToSchemeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleToSchemeResponse) GoString() string {
	return s.String()
}

func (s *UpdateRuleToSchemeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRuleToSchemeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRuleToSchemeResponse) GetBody() *UpdateRuleToSchemeResponseBody {
	return s.Body
}

func (s *UpdateRuleToSchemeResponse) SetHeaders(v map[string]*string) *UpdateRuleToSchemeResponse {
	s.Headers = v
	return s
}

func (s *UpdateRuleToSchemeResponse) SetStatusCode(v int32) *UpdateRuleToSchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRuleToSchemeResponse) SetBody(v *UpdateRuleToSchemeResponseBody) *UpdateRuleToSchemeResponse {
	s.Body = v
	return s
}

func (s *UpdateRuleToSchemeResponse) Validate() error {
	return dara.Validate(s)
}
