// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteConsumerAuthorizationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeleteConsumerAuthorizationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeleteConsumerAuthorizationRuleResponse
	GetStatusCode() *int32
	SetBody(v *BatchDeleteConsumerAuthorizationRuleResponseBody) *BatchDeleteConsumerAuthorizationRuleResponse
	GetBody() *BatchDeleteConsumerAuthorizationRuleResponseBody
}

type BatchDeleteConsumerAuthorizationRuleResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteConsumerAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteConsumerAuthorizationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteConsumerAuthorizationRuleResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteConsumerAuthorizationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeleteConsumerAuthorizationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeleteConsumerAuthorizationRuleResponse) GetBody() *BatchDeleteConsumerAuthorizationRuleResponseBody {
	return s.Body
}

func (s *BatchDeleteConsumerAuthorizationRuleResponse) SetHeaders(v map[string]*string) *BatchDeleteConsumerAuthorizationRuleResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteConsumerAuthorizationRuleResponse) SetStatusCode(v int32) *BatchDeleteConsumerAuthorizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteConsumerAuthorizationRuleResponse) SetBody(v *BatchDeleteConsumerAuthorizationRuleResponseBody) *BatchDeleteConsumerAuthorizationRuleResponse {
	s.Body = v
	return s
}

func (s *BatchDeleteConsumerAuthorizationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
