// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpIncomingRequestHeaderModificationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHttpIncomingRequestHeaderModificationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHttpIncomingRequestHeaderModificationRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHttpIncomingRequestHeaderModificationRuleResponseBody) *DeleteHttpIncomingRequestHeaderModificationRuleResponse
	GetBody() *DeleteHttpIncomingRequestHeaderModificationRuleResponseBody
}

type DeleteHttpIncomingRequestHeaderModificationRuleResponse struct {
	Headers    map[string]*string                                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHttpIncomingRequestHeaderModificationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHttpIncomingRequestHeaderModificationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpIncomingRequestHeaderModificationRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteHttpIncomingRequestHeaderModificationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHttpIncomingRequestHeaderModificationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHttpIncomingRequestHeaderModificationRuleResponse) GetBody() *DeleteHttpIncomingRequestHeaderModificationRuleResponseBody {
	return s.Body
}

func (s *DeleteHttpIncomingRequestHeaderModificationRuleResponse) SetHeaders(v map[string]*string) *DeleteHttpIncomingRequestHeaderModificationRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteHttpIncomingRequestHeaderModificationRuleResponse) SetStatusCode(v int32) *DeleteHttpIncomingRequestHeaderModificationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHttpIncomingRequestHeaderModificationRuleResponse) SetBody(v *DeleteHttpIncomingRequestHeaderModificationRuleResponseBody) *DeleteHttpIncomingRequestHeaderModificationRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteHttpIncomingRequestHeaderModificationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
