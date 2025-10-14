// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpIncomingResponseHeaderModificationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHttpIncomingResponseHeaderModificationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHttpIncomingResponseHeaderModificationRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHttpIncomingResponseHeaderModificationRuleResponseBody) *DeleteHttpIncomingResponseHeaderModificationRuleResponse
	GetBody() *DeleteHttpIncomingResponseHeaderModificationRuleResponseBody
}

type DeleteHttpIncomingResponseHeaderModificationRuleResponse struct {
	Headers    map[string]*string                                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHttpIncomingResponseHeaderModificationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHttpIncomingResponseHeaderModificationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpIncomingResponseHeaderModificationRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteHttpIncomingResponseHeaderModificationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHttpIncomingResponseHeaderModificationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHttpIncomingResponseHeaderModificationRuleResponse) GetBody() *DeleteHttpIncomingResponseHeaderModificationRuleResponseBody {
	return s.Body
}

func (s *DeleteHttpIncomingResponseHeaderModificationRuleResponse) SetHeaders(v map[string]*string) *DeleteHttpIncomingResponseHeaderModificationRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteHttpIncomingResponseHeaderModificationRuleResponse) SetStatusCode(v int32) *DeleteHttpIncomingResponseHeaderModificationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHttpIncomingResponseHeaderModificationRuleResponse) SetBody(v *DeleteHttpIncomingResponseHeaderModificationRuleResponseBody) *DeleteHttpIncomingResponseHeaderModificationRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteHttpIncomingResponseHeaderModificationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
