// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAppAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddAppAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddAppAgentResponse
	GetStatusCode() *int32
	SetBody(v *AddAppAgentResponseBody) *AddAppAgentResponse
	GetBody() *AddAppAgentResponseBody
}

type AddAppAgentResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAppAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAppAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s AddAppAgentResponse) GoString() string {
	return s.String()
}

func (s *AddAppAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddAppAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddAppAgentResponse) GetBody() *AddAppAgentResponseBody {
	return s.Body
}

func (s *AddAppAgentResponse) SetHeaders(v map[string]*string) *AddAppAgentResponse {
	s.Headers = v
	return s
}

func (s *AddAppAgentResponse) SetStatusCode(v int32) *AddAppAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAppAgentResponse) SetBody(v *AddAppAgentResponseBody) *AddAppAgentResponse {
	s.Body = v
	return s
}

func (s *AddAppAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
