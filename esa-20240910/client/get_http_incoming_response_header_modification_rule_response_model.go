// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHttpIncomingResponseHeaderModificationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHttpIncomingResponseHeaderModificationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHttpIncomingResponseHeaderModificationRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetHttpIncomingResponseHeaderModificationRuleResponseBody) *GetHttpIncomingResponseHeaderModificationRuleResponse
	GetBody() *GetHttpIncomingResponseHeaderModificationRuleResponseBody
}

type GetHttpIncomingResponseHeaderModificationRuleResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHttpIncomingResponseHeaderModificationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHttpIncomingResponseHeaderModificationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHttpIncomingResponseHeaderModificationRuleResponse) GoString() string {
	return s.String()
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponse) GetBody() *GetHttpIncomingResponseHeaderModificationRuleResponseBody {
	return s.Body
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponse) SetHeaders(v map[string]*string) *GetHttpIncomingResponseHeaderModificationRuleResponse {
	s.Headers = v
	return s
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponse) SetStatusCode(v int32) *GetHttpIncomingResponseHeaderModificationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponse) SetBody(v *GetHttpIncomingResponseHeaderModificationRuleResponseBody) *GetHttpIncomingResponseHeaderModificationRuleResponse {
	s.Body = v
	return s
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
