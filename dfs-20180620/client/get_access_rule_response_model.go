// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccessRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccessRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetAccessRuleResponseBody) *GetAccessRuleResponse
	GetBody() *GetAccessRuleResponseBody
}

type GetAccessRuleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccessRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccessRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccessRuleResponse) GoString() string {
	return s.String()
}

func (s *GetAccessRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccessRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccessRuleResponse) GetBody() *GetAccessRuleResponseBody {
	return s.Body
}

func (s *GetAccessRuleResponse) SetHeaders(v map[string]*string) *GetAccessRuleResponse {
	s.Headers = v
	return s
}

func (s *GetAccessRuleResponse) SetStatusCode(v int32) *GetAccessRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessRuleResponse) SetBody(v *GetAccessRuleResponseBody) *GetAccessRuleResponse {
	s.Body = v
	return s
}

func (s *GetAccessRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
