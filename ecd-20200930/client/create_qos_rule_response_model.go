// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQosRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateQosRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateQosRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateQosRuleResponseBody) *CreateQosRuleResponse
	GetBody() *CreateQosRuleResponseBody
}

type CreateQosRuleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQosRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQosRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateQosRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateQosRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateQosRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateQosRuleResponse) GetBody() *CreateQosRuleResponseBody {
	return s.Body
}

func (s *CreateQosRuleResponse) SetHeaders(v map[string]*string) *CreateQosRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateQosRuleResponse) SetStatusCode(v int32) *CreateQosRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQosRuleResponse) SetBody(v *CreateQosRuleResponseBody) *CreateQosRuleResponse {
	s.Body = v
	return s
}

func (s *CreateQosRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
