// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecordingRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRecordingRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRecordingRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetRecordingRuleResponseBody) *GetRecordingRuleResponse
	GetBody() *GetRecordingRuleResponseBody
}

type GetRecordingRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRecordingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRecordingRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRecordingRuleResponse) GoString() string {
	return s.String()
}

func (s *GetRecordingRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRecordingRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRecordingRuleResponse) GetBody() *GetRecordingRuleResponseBody {
	return s.Body
}

func (s *GetRecordingRuleResponse) SetHeaders(v map[string]*string) *GetRecordingRuleResponse {
	s.Headers = v
	return s
}

func (s *GetRecordingRuleResponse) SetStatusCode(v int32) *GetRecordingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRecordingRuleResponse) SetBody(v *GetRecordingRuleResponseBody) *GetRecordingRuleResponse {
	s.Body = v
	return s
}

func (s *GetRecordingRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
