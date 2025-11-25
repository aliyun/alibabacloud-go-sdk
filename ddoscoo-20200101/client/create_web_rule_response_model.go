// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWebRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWebRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWebRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateWebRuleResponseBody) *CreateWebRuleResponse
	GetBody() *CreateWebRuleResponseBody
}

type CreateWebRuleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWebRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWebRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWebRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateWebRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWebRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWebRuleResponse) GetBody() *CreateWebRuleResponseBody {
	return s.Body
}

func (s *CreateWebRuleResponse) SetHeaders(v map[string]*string) *CreateWebRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateWebRuleResponse) SetStatusCode(v int32) *CreateWebRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWebRuleResponse) SetBody(v *CreateWebRuleResponseBody) *CreateWebRuleResponse {
	s.Body = v
	return s
}

func (s *CreateWebRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
