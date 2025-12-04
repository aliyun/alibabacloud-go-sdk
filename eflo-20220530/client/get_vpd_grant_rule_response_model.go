// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpdGrantRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVpdGrantRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVpdGrantRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetVpdGrantRuleResponseBody) *GetVpdGrantRuleResponse
	GetBody() *GetVpdGrantRuleResponseBody
}

type GetVpdGrantRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVpdGrantRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVpdGrantRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVpdGrantRuleResponse) GoString() string {
	return s.String()
}

func (s *GetVpdGrantRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVpdGrantRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVpdGrantRuleResponse) GetBody() *GetVpdGrantRuleResponseBody {
	return s.Body
}

func (s *GetVpdGrantRuleResponse) SetHeaders(v map[string]*string) *GetVpdGrantRuleResponse {
	s.Headers = v
	return s
}

func (s *GetVpdGrantRuleResponse) SetStatusCode(v int32) *GetVpdGrantRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVpdGrantRuleResponse) SetBody(v *GetVpdGrantRuleResponseBody) *GetVpdGrantRuleResponse {
	s.Body = v
	return s
}

func (s *GetVpdGrantRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
