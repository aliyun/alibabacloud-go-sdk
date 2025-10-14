// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDetectionRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDetectionRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDetectionRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateDetectionRuleResponseBody) *CreateDetectionRuleResponse
	GetBody() *CreateDetectionRuleResponseBody
}

type CreateDetectionRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDetectionRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDetectionRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDetectionRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateDetectionRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDetectionRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDetectionRuleResponse) GetBody() *CreateDetectionRuleResponseBody {
	return s.Body
}

func (s *CreateDetectionRuleResponse) SetHeaders(v map[string]*string) *CreateDetectionRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateDetectionRuleResponse) SetStatusCode(v int32) *CreateDetectionRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDetectionRuleResponse) SetBody(v *CreateDetectionRuleResponseBody) *CreateDetectionRuleResponse {
	s.Body = v
	return s
}

func (s *CreateDetectionRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
