// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpIncomingResponseHeaderModificationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHttpIncomingResponseHeaderModificationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHttpIncomingResponseHeaderModificationRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateHttpIncomingResponseHeaderModificationRuleResponseBody) *CreateHttpIncomingResponseHeaderModificationRuleResponse
	GetBody() *CreateHttpIncomingResponseHeaderModificationRuleResponseBody
}

type CreateHttpIncomingResponseHeaderModificationRuleResponse struct {
	Headers    map[string]*string                                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHttpIncomingResponseHeaderModificationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHttpIncomingResponseHeaderModificationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpIncomingResponseHeaderModificationRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleResponse) GetBody() *CreateHttpIncomingResponseHeaderModificationRuleResponseBody {
	return s.Body
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleResponse) SetHeaders(v map[string]*string) *CreateHttpIncomingResponseHeaderModificationRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleResponse) SetStatusCode(v int32) *CreateHttpIncomingResponseHeaderModificationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleResponse) SetBody(v *CreateHttpIncomingResponseHeaderModificationRuleResponseBody) *CreateHttpIncomingResponseHeaderModificationRuleResponse {
	s.Body = v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
