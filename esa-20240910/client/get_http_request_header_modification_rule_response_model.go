// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHttpRequestHeaderModificationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHttpRequestHeaderModificationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHttpRequestHeaderModificationRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetHttpRequestHeaderModificationRuleResponseBody) *GetHttpRequestHeaderModificationRuleResponse
	GetBody() *GetHttpRequestHeaderModificationRuleResponseBody
}

type GetHttpRequestHeaderModificationRuleResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHttpRequestHeaderModificationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHttpRequestHeaderModificationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHttpRequestHeaderModificationRuleResponse) GoString() string {
	return s.String()
}

func (s *GetHttpRequestHeaderModificationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHttpRequestHeaderModificationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHttpRequestHeaderModificationRuleResponse) GetBody() *GetHttpRequestHeaderModificationRuleResponseBody {
	return s.Body
}

func (s *GetHttpRequestHeaderModificationRuleResponse) SetHeaders(v map[string]*string) *GetHttpRequestHeaderModificationRuleResponse {
	s.Headers = v
	return s
}

func (s *GetHttpRequestHeaderModificationRuleResponse) SetStatusCode(v int32) *GetHttpRequestHeaderModificationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHttpRequestHeaderModificationRuleResponse) SetBody(v *GetHttpRequestHeaderModificationRuleResponseBody) *GetHttpRequestHeaderModificationRuleResponse {
	s.Body = v
	return s
}

func (s *GetHttpRequestHeaderModificationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
