// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentIndexRealTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentIndexRealTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentIndexRealTimeResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentIndexRealTimeResponseBody) *GetAgentIndexRealTimeResponse
	GetBody() *GetAgentIndexRealTimeResponseBody
}

type GetAgentIndexRealTimeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentIndexRealTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentIndexRealTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentIndexRealTimeResponse) GoString() string {
	return s.String()
}

func (s *GetAgentIndexRealTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentIndexRealTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentIndexRealTimeResponse) GetBody() *GetAgentIndexRealTimeResponseBody {
	return s.Body
}

func (s *GetAgentIndexRealTimeResponse) SetHeaders(v map[string]*string) *GetAgentIndexRealTimeResponse {
	s.Headers = v
	return s
}

func (s *GetAgentIndexRealTimeResponse) SetStatusCode(v int32) *GetAgentIndexRealTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentIndexRealTimeResponse) SetBody(v *GetAgentIndexRealTimeResponseBody) *GetAgentIndexRealTimeResponse {
	s.Body = v
	return s
}

func (s *GetAgentIndexRealTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
