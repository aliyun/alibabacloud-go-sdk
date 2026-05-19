// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileProtectClientRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFileProtectClientRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFileProtectClientRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateFileProtectClientRuleResponseBody) *CreateFileProtectClientRuleResponse
	GetBody() *CreateFileProtectClientRuleResponseBody
}

type CreateFileProtectClientRuleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFileProtectClientRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFileProtectClientRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFileProtectClientRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateFileProtectClientRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFileProtectClientRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFileProtectClientRuleResponse) GetBody() *CreateFileProtectClientRuleResponseBody {
	return s.Body
}

func (s *CreateFileProtectClientRuleResponse) SetHeaders(v map[string]*string) *CreateFileProtectClientRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateFileProtectClientRuleResponse) SetStatusCode(v int32) *CreateFileProtectClientRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFileProtectClientRuleResponse) SetBody(v *CreateFileProtectClientRuleResponseBody) *CreateFileProtectClientRuleResponse {
	s.Body = v
	return s
}

func (s *CreateFileProtectClientRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
