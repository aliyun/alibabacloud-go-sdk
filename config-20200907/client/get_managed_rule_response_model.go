// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetManagedRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetManagedRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetManagedRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetManagedRuleResponseBody) *GetManagedRuleResponse
	GetBody() *GetManagedRuleResponseBody
}

type GetManagedRuleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetManagedRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetManagedRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetManagedRuleResponse) GoString() string {
	return s.String()
}

func (s *GetManagedRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetManagedRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetManagedRuleResponse) GetBody() *GetManagedRuleResponseBody {
	return s.Body
}

func (s *GetManagedRuleResponse) SetHeaders(v map[string]*string) *GetManagedRuleResponse {
	s.Headers = v
	return s
}

func (s *GetManagedRuleResponse) SetStatusCode(v int32) *GetManagedRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetManagedRuleResponse) SetBody(v *GetManagedRuleResponseBody) *GetManagedRuleResponse {
	s.Body = v
	return s
}

func (s *GetManagedRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
