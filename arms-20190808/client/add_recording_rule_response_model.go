// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRecordingRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddRecordingRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddRecordingRuleResponse
	GetStatusCode() *int32
	SetBody(v *AddRecordingRuleResponseBody) *AddRecordingRuleResponse
	GetBody() *AddRecordingRuleResponseBody
}

type AddRecordingRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRecordingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRecordingRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s AddRecordingRuleResponse) GoString() string {
	return s.String()
}

func (s *AddRecordingRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddRecordingRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddRecordingRuleResponse) GetBody() *AddRecordingRuleResponseBody {
	return s.Body
}

func (s *AddRecordingRuleResponse) SetHeaders(v map[string]*string) *AddRecordingRuleResponse {
	s.Headers = v
	return s
}

func (s *AddRecordingRuleResponse) SetStatusCode(v int32) *AddRecordingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRecordingRuleResponse) SetBody(v *AddRecordingRuleResponseBody) *AddRecordingRuleResponse {
	s.Body = v
	return s
}

func (s *AddRecordingRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
