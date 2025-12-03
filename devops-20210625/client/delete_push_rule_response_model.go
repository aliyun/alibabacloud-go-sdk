// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePushRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePushRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePushRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeletePushRuleResponseBody) *DeletePushRuleResponse
	GetBody() *DeletePushRuleResponseBody
}

type DeletePushRuleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePushRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePushRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePushRuleResponse) GoString() string {
	return s.String()
}

func (s *DeletePushRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePushRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePushRuleResponse) GetBody() *DeletePushRuleResponseBody {
	return s.Body
}

func (s *DeletePushRuleResponse) SetHeaders(v map[string]*string) *DeletePushRuleResponse {
	s.Headers = v
	return s
}

func (s *DeletePushRuleResponse) SetStatusCode(v int32) *DeletePushRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePushRuleResponse) SetBody(v *DeletePushRuleResponseBody) *DeletePushRuleResponse {
	s.Body = v
	return s
}

func (s *DeletePushRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
