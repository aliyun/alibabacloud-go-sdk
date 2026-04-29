// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolarClawAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePolarClawAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePolarClawAgentResponse
	GetStatusCode() *int32
	SetBody(v *DeletePolarClawAgentResponseBody) *DeletePolarClawAgentResponse
	GetBody() *DeletePolarClawAgentResponseBody
}

type DeletePolarClawAgentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePolarClawAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePolarClawAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarClawAgentResponse) GoString() string {
	return s.String()
}

func (s *DeletePolarClawAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePolarClawAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePolarClawAgentResponse) GetBody() *DeletePolarClawAgentResponseBody {
	return s.Body
}

func (s *DeletePolarClawAgentResponse) SetHeaders(v map[string]*string) *DeletePolarClawAgentResponse {
	s.Headers = v
	return s
}

func (s *DeletePolarClawAgentResponse) SetStatusCode(v int32) *DeletePolarClawAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePolarClawAgentResponse) SetBody(v *DeletePolarClawAgentResponseBody) *DeletePolarClawAgentResponse {
	s.Body = v
	return s
}

func (s *DeletePolarClawAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
