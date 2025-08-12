// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutCustomEventRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutCustomEventRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutCustomEventRuleResponse
	GetStatusCode() *int32
	SetBody(v *PutCustomEventRuleResponseBody) *PutCustomEventRuleResponse
	GetBody() *PutCustomEventRuleResponseBody
}

type PutCustomEventRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutCustomEventRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutCustomEventRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s PutCustomEventRuleResponse) GoString() string {
	return s.String()
}

func (s *PutCustomEventRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutCustomEventRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutCustomEventRuleResponse) GetBody() *PutCustomEventRuleResponseBody {
	return s.Body
}

func (s *PutCustomEventRuleResponse) SetHeaders(v map[string]*string) *PutCustomEventRuleResponse {
	s.Headers = v
	return s
}

func (s *PutCustomEventRuleResponse) SetStatusCode(v int32) *PutCustomEventRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *PutCustomEventRuleResponse) SetBody(v *PutCustomEventRuleResponseBody) *PutCustomEventRuleResponse {
	s.Body = v
	return s
}

func (s *PutCustomEventRuleResponse) Validate() error {
	return dara.Validate(s)
}
