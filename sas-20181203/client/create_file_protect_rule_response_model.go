// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileProtectRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFileProtectRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFileProtectRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateFileProtectRuleResponseBody) *CreateFileProtectRuleResponse
	GetBody() *CreateFileProtectRuleResponseBody
}

type CreateFileProtectRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFileProtectRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFileProtectRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFileProtectRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateFileProtectRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFileProtectRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFileProtectRuleResponse) GetBody() *CreateFileProtectRuleResponseBody {
	return s.Body
}

func (s *CreateFileProtectRuleResponse) SetHeaders(v map[string]*string) *CreateFileProtectRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateFileProtectRuleResponse) SetStatusCode(v int32) *CreateFileProtectRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFileProtectRuleResponse) SetBody(v *CreateFileProtectRuleResponseBody) *CreateFileProtectRuleResponse {
	s.Body = v
	return s
}

func (s *CreateFileProtectRuleResponse) Validate() error {
	return dara.Validate(s)
}
