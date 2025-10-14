// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDetectionRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDetectionRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDetectionRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDetectionRuleResponseBody) *UpdateDetectionRuleResponse
	GetBody() *UpdateDetectionRuleResponseBody
}

type UpdateDetectionRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDetectionRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDetectionRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDetectionRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateDetectionRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDetectionRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDetectionRuleResponse) GetBody() *UpdateDetectionRuleResponseBody {
	return s.Body
}

func (s *UpdateDetectionRuleResponse) SetHeaders(v map[string]*string) *UpdateDetectionRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateDetectionRuleResponse) SetStatusCode(v int32) *UpdateDetectionRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDetectionRuleResponse) SetBody(v *UpdateDetectionRuleResponseBody) *UpdateDetectionRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateDetectionRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
