// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolarAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPolarAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPolarAgentResponse
	GetStatusCode() *int32
	SetBody(v *GetPolarAgentResponseBody) *GetPolarAgentResponse
	GetBody() *GetPolarAgentResponseBody
}

type GetPolarAgentResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPolarAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPolarAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPolarAgentResponse) GoString() string {
	return s.String()
}

func (s *GetPolarAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPolarAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPolarAgentResponse) GetBody() *GetPolarAgentResponseBody {
	return s.Body
}

func (s *GetPolarAgentResponse) SetHeaders(v map[string]*string) *GetPolarAgentResponse {
	s.Headers = v
	return s
}

func (s *GetPolarAgentResponse) SetStatusCode(v int32) *GetPolarAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPolarAgentResponse) SetBody(v *GetPolarAgentResponseBody) *GetPolarAgentResponse {
	s.Body = v
	return s
}

func (s *GetPolarAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
