// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAgentSessionModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetAgentSessionModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetAgentSessionModelResponse
	GetStatusCode() *int32
	SetBody(v *SetAgentSessionModelResponseBody) *SetAgentSessionModelResponse
	GetBody() *SetAgentSessionModelResponseBody
}

type SetAgentSessionModelResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetAgentSessionModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetAgentSessionModelResponse) String() string {
	return dara.Prettify(s)
}

func (s SetAgentSessionModelResponse) GoString() string {
	return s.String()
}

func (s *SetAgentSessionModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetAgentSessionModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetAgentSessionModelResponse) GetBody() *SetAgentSessionModelResponseBody {
	return s.Body
}

func (s *SetAgentSessionModelResponse) SetHeaders(v map[string]*string) *SetAgentSessionModelResponse {
	s.Headers = v
	return s
}

func (s *SetAgentSessionModelResponse) SetStatusCode(v int32) *SetAgentSessionModelResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAgentSessionModelResponse) SetBody(v *SetAgentSessionModelResponseBody) *SetAgentSessionModelResponse {
	s.Body = v
	return s
}

func (s *SetAgentSessionModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
