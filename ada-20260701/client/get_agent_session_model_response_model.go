// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSessionModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentSessionModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentSessionModelResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentSessionModelResponseBody) *GetAgentSessionModelResponse
	GetBody() *GetAgentSessionModelResponseBody
}

type GetAgentSessionModelResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentSessionModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentSessionModelResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSessionModelResponse) GoString() string {
	return s.String()
}

func (s *GetAgentSessionModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentSessionModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentSessionModelResponse) GetBody() *GetAgentSessionModelResponseBody {
	return s.Body
}

func (s *GetAgentSessionModelResponse) SetHeaders(v map[string]*string) *GetAgentSessionModelResponse {
	s.Headers = v
	return s
}

func (s *GetAgentSessionModelResponse) SetStatusCode(v int32) *GetAgentSessionModelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentSessionModelResponse) SetBody(v *GetAgentSessionModelResponseBody) *GetAgentSessionModelResponse {
	s.Body = v
	return s
}

func (s *GetAgentSessionModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
