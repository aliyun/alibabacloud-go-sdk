// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileProtectRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFileProtectRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFileProtectRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetFileProtectRuleResponseBody) *GetFileProtectRuleResponse
	GetBody() *GetFileProtectRuleResponseBody
}

type GetFileProtectRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFileProtectRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileProtectRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectRuleResponse) GoString() string {
	return s.String()
}

func (s *GetFileProtectRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFileProtectRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFileProtectRuleResponse) GetBody() *GetFileProtectRuleResponseBody {
	return s.Body
}

func (s *GetFileProtectRuleResponse) SetHeaders(v map[string]*string) *GetFileProtectRuleResponse {
	s.Headers = v
	return s
}

func (s *GetFileProtectRuleResponse) SetStatusCode(v int32) *GetFileProtectRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileProtectRuleResponse) SetBody(v *GetFileProtectRuleResponseBody) *GetFileProtectRuleResponse {
	s.Body = v
	return s
}

func (s *GetFileProtectRuleResponse) Validate() error {
	return dara.Validate(s)
}
