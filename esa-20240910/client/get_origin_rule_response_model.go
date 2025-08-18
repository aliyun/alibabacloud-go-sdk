// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOriginRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOriginRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOriginRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetOriginRuleResponseBody) *GetOriginRuleResponse
	GetBody() *GetOriginRuleResponseBody
}

type GetOriginRuleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOriginRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOriginRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOriginRuleResponse) GoString() string {
	return s.String()
}

func (s *GetOriginRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOriginRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOriginRuleResponse) GetBody() *GetOriginRuleResponseBody {
	return s.Body
}

func (s *GetOriginRuleResponse) SetHeaders(v map[string]*string) *GetOriginRuleResponse {
	s.Headers = v
	return s
}

func (s *GetOriginRuleResponse) SetStatusCode(v int32) *GetOriginRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOriginRuleResponse) SetBody(v *GetOriginRuleResponseBody) *GetOriginRuleResponse {
	s.Body = v
	return s
}

func (s *GetOriginRuleResponse) Validate() error {
	return dara.Validate(s)
}
