// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVccGrantRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVccGrantRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVccGrantRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetVccGrantRuleResponseBody) *GetVccGrantRuleResponse
	GetBody() *GetVccGrantRuleResponseBody
}

type GetVccGrantRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVccGrantRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVccGrantRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVccGrantRuleResponse) GoString() string {
	return s.String()
}

func (s *GetVccGrantRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVccGrantRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVccGrantRuleResponse) GetBody() *GetVccGrantRuleResponseBody {
	return s.Body
}

func (s *GetVccGrantRuleResponse) SetHeaders(v map[string]*string) *GetVccGrantRuleResponse {
	s.Headers = v
	return s
}

func (s *GetVccGrantRuleResponse) SetStatusCode(v int32) *GetVccGrantRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVccGrantRuleResponse) SetBody(v *GetVccGrantRuleResponseBody) *GetVccGrantRuleResponse {
	s.Body = v
	return s
}

func (s *GetVccGrantRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
