// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomResponseCodeRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomResponseCodeRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomResponseCodeRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomResponseCodeRuleResponseBody) *DeleteCustomResponseCodeRuleResponse
	GetBody() *DeleteCustomResponseCodeRuleResponseBody
}

type DeleteCustomResponseCodeRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomResponseCodeRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomResponseCodeRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomResponseCodeRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomResponseCodeRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomResponseCodeRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomResponseCodeRuleResponse) GetBody() *DeleteCustomResponseCodeRuleResponseBody {
	return s.Body
}

func (s *DeleteCustomResponseCodeRuleResponse) SetHeaders(v map[string]*string) *DeleteCustomResponseCodeRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomResponseCodeRuleResponse) SetStatusCode(v int32) *DeleteCustomResponseCodeRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomResponseCodeRuleResponse) SetBody(v *DeleteCustomResponseCodeRuleResponseBody) *DeleteCustomResponseCodeRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomResponseCodeRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
