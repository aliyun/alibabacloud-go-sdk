// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResponseRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateResponseRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateResponseRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateResponseRuleResponseBody) *CreateResponseRuleResponse
	GetBody() *CreateResponseRuleResponseBody
}

type CreateResponseRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResponseRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResponseRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateResponseRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateResponseRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateResponseRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateResponseRuleResponse) GetBody() *CreateResponseRuleResponseBody {
	return s.Body
}

func (s *CreateResponseRuleResponse) SetHeaders(v map[string]*string) *CreateResponseRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateResponseRuleResponse) SetStatusCode(v int32) *CreateResponseRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResponseRuleResponse) SetBody(v *CreateResponseRuleResponseBody) *CreateResponseRuleResponse {
	s.Body = v
	return s
}

func (s *CreateResponseRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
