// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveAIProduceRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddLiveAIProduceRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddLiveAIProduceRulesResponse
	GetStatusCode() *int32
	SetBody(v *AddLiveAIProduceRulesResponseBody) *AddLiveAIProduceRulesResponse
	GetBody() *AddLiveAIProduceRulesResponseBody
}

type AddLiveAIProduceRulesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLiveAIProduceRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLiveAIProduceRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s AddLiveAIProduceRulesResponse) GoString() string {
	return s.String()
}

func (s *AddLiveAIProduceRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddLiveAIProduceRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddLiveAIProduceRulesResponse) GetBody() *AddLiveAIProduceRulesResponseBody {
	return s.Body
}

func (s *AddLiveAIProduceRulesResponse) SetHeaders(v map[string]*string) *AddLiveAIProduceRulesResponse {
	s.Headers = v
	return s
}

func (s *AddLiveAIProduceRulesResponse) SetStatusCode(v int32) *AddLiveAIProduceRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLiveAIProduceRulesResponse) SetBody(v *AddLiveAIProduceRulesResponseBody) *AddLiveAIProduceRulesResponse {
	s.Body = v
	return s
}

func (s *AddLiveAIProduceRulesResponse) Validate() error {
	return dara.Validate(s)
}
