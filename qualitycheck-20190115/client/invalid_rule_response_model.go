// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvalidRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvalidRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvalidRuleResponse
	GetStatusCode() *int32
	SetBody(v *InvalidRuleResponseBody) *InvalidRuleResponse
	GetBody() *InvalidRuleResponseBody
}

type InvalidRuleResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvalidRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvalidRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s InvalidRuleResponse) GoString() string {
	return s.String()
}

func (s *InvalidRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvalidRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvalidRuleResponse) GetBody() *InvalidRuleResponseBody {
	return s.Body
}

func (s *InvalidRuleResponse) SetHeaders(v map[string]*string) *InvalidRuleResponse {
	s.Headers = v
	return s
}

func (s *InvalidRuleResponse) SetStatusCode(v int32) *InvalidRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *InvalidRuleResponse) SetBody(v *InvalidRuleResponseBody) *InvalidRuleResponse {
	s.Body = v
	return s
}

func (s *InvalidRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
