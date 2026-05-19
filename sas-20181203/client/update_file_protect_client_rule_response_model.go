// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileProtectClientRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFileProtectClientRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFileProtectClientRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFileProtectClientRuleResponseBody) *UpdateFileProtectClientRuleResponse
	GetBody() *UpdateFileProtectClientRuleResponseBody
}

type UpdateFileProtectClientRuleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFileProtectClientRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFileProtectClientRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileProtectClientRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateFileProtectClientRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFileProtectClientRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFileProtectClientRuleResponse) GetBody() *UpdateFileProtectClientRuleResponseBody {
	return s.Body
}

func (s *UpdateFileProtectClientRuleResponse) SetHeaders(v map[string]*string) *UpdateFileProtectClientRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateFileProtectClientRuleResponse) SetStatusCode(v int32) *UpdateFileProtectClientRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFileProtectClientRuleResponse) SetBody(v *UpdateFileProtectClientRuleResponseBody) *UpdateFileProtectClientRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateFileProtectClientRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
