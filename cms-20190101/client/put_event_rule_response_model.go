// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutEventRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutEventRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutEventRuleResponse
	GetStatusCode() *int32
	SetBody(v *PutEventRuleResponseBody) *PutEventRuleResponse
	GetBody() *PutEventRuleResponseBody
}

type PutEventRuleResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutEventRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutEventRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s PutEventRuleResponse) GoString() string {
	return s.String()
}

func (s *PutEventRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutEventRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutEventRuleResponse) GetBody() *PutEventRuleResponseBody {
	return s.Body
}

func (s *PutEventRuleResponse) SetHeaders(v map[string]*string) *PutEventRuleResponse {
	s.Headers = v
	return s
}

func (s *PutEventRuleResponse) SetStatusCode(v int32) *PutEventRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *PutEventRuleResponse) SetBody(v *PutEventRuleResponseBody) *PutEventRuleResponse {
	s.Body = v
	return s
}

func (s *PutEventRuleResponse) Validate() error {
	return dara.Validate(s)
}
