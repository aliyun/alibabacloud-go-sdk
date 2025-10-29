// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveAIProduceRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLiveAIProduceRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLiveAIProduceRulesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLiveAIProduceRulesResponseBody) *UpdateLiveAIProduceRulesResponse
	GetBody() *UpdateLiveAIProduceRulesResponseBody
}

type UpdateLiveAIProduceRulesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLiveAIProduceRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLiveAIProduceRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveAIProduceRulesResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveAIProduceRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLiveAIProduceRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLiveAIProduceRulesResponse) GetBody() *UpdateLiveAIProduceRulesResponseBody {
	return s.Body
}

func (s *UpdateLiveAIProduceRulesResponse) SetHeaders(v map[string]*string) *UpdateLiveAIProduceRulesResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveAIProduceRulesResponse) SetStatusCode(v int32) *UpdateLiveAIProduceRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveAIProduceRulesResponse) SetBody(v *UpdateLiveAIProduceRulesResponseBody) *UpdateLiveAIProduceRulesResponse {
	s.Body = v
	return s
}

func (s *UpdateLiveAIProduceRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
