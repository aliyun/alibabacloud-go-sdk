// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveAIProduceRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveAIProduceRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveAIProduceRulesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveAIProduceRulesResponseBody) *DeleteLiveAIProduceRulesResponse
	GetBody() *DeleteLiveAIProduceRulesResponseBody
}

type DeleteLiveAIProduceRulesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveAIProduceRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveAIProduceRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveAIProduceRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveAIProduceRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveAIProduceRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveAIProduceRulesResponse) GetBody() *DeleteLiveAIProduceRulesResponseBody {
	return s.Body
}

func (s *DeleteLiveAIProduceRulesResponse) SetHeaders(v map[string]*string) *DeleteLiveAIProduceRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveAIProduceRulesResponse) SetStatusCode(v int32) *DeleteLiveAIProduceRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveAIProduceRulesResponse) SetBody(v *DeleteLiveAIProduceRulesResponseBody) *DeleteLiveAIProduceRulesResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveAIProduceRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
