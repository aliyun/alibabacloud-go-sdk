// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpIncomingRequestHeaderModificationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHttpIncomingRequestHeaderModificationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHttpIncomingRequestHeaderModificationRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateHttpIncomingRequestHeaderModificationRuleResponseBody) *CreateHttpIncomingRequestHeaderModificationRuleResponse
	GetBody() *CreateHttpIncomingRequestHeaderModificationRuleResponseBody
}

type CreateHttpIncomingRequestHeaderModificationRuleResponse struct {
	Headers    map[string]*string                                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHttpIncomingRequestHeaderModificationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHttpIncomingRequestHeaderModificationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpIncomingRequestHeaderModificationRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleResponse) GetBody() *CreateHttpIncomingRequestHeaderModificationRuleResponseBody {
	return s.Body
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleResponse) SetHeaders(v map[string]*string) *CreateHttpIncomingRequestHeaderModificationRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleResponse) SetStatusCode(v int32) *CreateHttpIncomingRequestHeaderModificationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleResponse) SetBody(v *CreateHttpIncomingRequestHeaderModificationRuleResponseBody) *CreateHttpIncomingRequestHeaderModificationRuleResponse {
	s.Body = v
	return s
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
