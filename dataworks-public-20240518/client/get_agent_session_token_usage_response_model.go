// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSessionTokenUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentSessionTokenUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentSessionTokenUsageResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentSessionTokenUsageResponseBody) *GetAgentSessionTokenUsageResponse
	GetBody() *GetAgentSessionTokenUsageResponseBody
}

type GetAgentSessionTokenUsageResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentSessionTokenUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentSessionTokenUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSessionTokenUsageResponse) GoString() string {
	return s.String()
}

func (s *GetAgentSessionTokenUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentSessionTokenUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentSessionTokenUsageResponse) GetBody() *GetAgentSessionTokenUsageResponseBody {
	return s.Body
}

func (s *GetAgentSessionTokenUsageResponse) SetHeaders(v map[string]*string) *GetAgentSessionTokenUsageResponse {
	s.Headers = v
	return s
}

func (s *GetAgentSessionTokenUsageResponse) SetStatusCode(v int32) *GetAgentSessionTokenUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentSessionTokenUsageResponse) SetBody(v *GetAgentSessionTokenUsageResponseBody) *GetAgentSessionTokenUsageResponse {
	s.Body = v
	return s
}

func (s *GetAgentSessionTokenUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
