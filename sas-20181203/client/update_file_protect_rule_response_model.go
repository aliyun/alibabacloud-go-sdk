// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileProtectRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFileProtectRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFileProtectRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFileProtectRuleResponseBody) *UpdateFileProtectRuleResponse
	GetBody() *UpdateFileProtectRuleResponseBody
}

type UpdateFileProtectRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFileProtectRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFileProtectRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileProtectRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateFileProtectRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFileProtectRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFileProtectRuleResponse) GetBody() *UpdateFileProtectRuleResponseBody {
	return s.Body
}

func (s *UpdateFileProtectRuleResponse) SetHeaders(v map[string]*string) *UpdateFileProtectRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateFileProtectRuleResponse) SetStatusCode(v int32) *UpdateFileProtectRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFileProtectRuleResponse) SetBody(v *UpdateFileProtectRuleResponseBody) *UpdateFileProtectRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateFileProtectRuleResponse) Validate() error {
	return dara.Validate(s)
}
