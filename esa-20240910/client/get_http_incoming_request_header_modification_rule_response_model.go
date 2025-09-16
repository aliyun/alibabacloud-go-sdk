// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHttpIncomingRequestHeaderModificationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHttpIncomingRequestHeaderModificationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHttpIncomingRequestHeaderModificationRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetHttpIncomingRequestHeaderModificationRuleResponseBody) *GetHttpIncomingRequestHeaderModificationRuleResponse
	GetBody() *GetHttpIncomingRequestHeaderModificationRuleResponseBody
}

type GetHttpIncomingRequestHeaderModificationRuleResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHttpIncomingRequestHeaderModificationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHttpIncomingRequestHeaderModificationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHttpIncomingRequestHeaderModificationRuleResponse) GoString() string {
	return s.String()
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponse) GetBody() *GetHttpIncomingRequestHeaderModificationRuleResponseBody {
	return s.Body
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponse) SetHeaders(v map[string]*string) *GetHttpIncomingRequestHeaderModificationRuleResponse {
	s.Headers = v
	return s
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponse) SetStatusCode(v int32) *GetHttpIncomingRequestHeaderModificationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponse) SetBody(v *GetHttpIncomingRequestHeaderModificationRuleResponseBody) *GetHttpIncomingRequestHeaderModificationRuleResponse {
	s.Body = v
	return s
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponse) Validate() error {
	return dara.Validate(s)
}
