// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileProtectClientRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFileProtectClientRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFileProtectClientRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetFileProtectClientRuleResponseBody) *GetFileProtectClientRuleResponse
	GetBody() *GetFileProtectClientRuleResponseBody
}

type GetFileProtectClientRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFileProtectClientRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileProtectClientRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectClientRuleResponse) GoString() string {
	return s.String()
}

func (s *GetFileProtectClientRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFileProtectClientRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFileProtectClientRuleResponse) GetBody() *GetFileProtectClientRuleResponseBody {
	return s.Body
}

func (s *GetFileProtectClientRuleResponse) SetHeaders(v map[string]*string) *GetFileProtectClientRuleResponse {
	s.Headers = v
	return s
}

func (s *GetFileProtectClientRuleResponse) SetStatusCode(v int32) *GetFileProtectClientRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileProtectClientRuleResponse) SetBody(v *GetFileProtectClientRuleResponseBody) *GetFileProtectClientRuleResponse {
	s.Body = v
	return s
}

func (s *GetFileProtectClientRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
