// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientUserDefineRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClientUserDefineRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClientUserDefineRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetClientUserDefineRuleResponseBody) *GetClientUserDefineRuleResponse
	GetBody() *GetClientUserDefineRuleResponseBody
}

type GetClientUserDefineRuleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClientUserDefineRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClientUserDefineRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClientUserDefineRuleResponse) GoString() string {
	return s.String()
}

func (s *GetClientUserDefineRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClientUserDefineRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClientUserDefineRuleResponse) GetBody() *GetClientUserDefineRuleResponseBody {
	return s.Body
}

func (s *GetClientUserDefineRuleResponse) SetHeaders(v map[string]*string) *GetClientUserDefineRuleResponse {
	s.Headers = v
	return s
}

func (s *GetClientUserDefineRuleResponse) SetStatusCode(v int32) *GetClientUserDefineRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClientUserDefineRuleResponse) SetBody(v *GetClientUserDefineRuleResponseBody) *GetClientUserDefineRuleResponse {
	s.Body = v
	return s
}

func (s *GetClientUserDefineRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
