// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentWorkLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentWorkLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentWorkLogResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentWorkLogResponseBody) *GetAgentWorkLogResponse
	GetBody() *GetAgentWorkLogResponseBody
}

type GetAgentWorkLogResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentWorkLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentWorkLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentWorkLogResponse) GoString() string {
	return s.String()
}

func (s *GetAgentWorkLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentWorkLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentWorkLogResponse) GetBody() *GetAgentWorkLogResponseBody {
	return s.Body
}

func (s *GetAgentWorkLogResponse) SetHeaders(v map[string]*string) *GetAgentWorkLogResponse {
	s.Headers = v
	return s
}

func (s *GetAgentWorkLogResponse) SetStatusCode(v int32) *GetAgentWorkLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentWorkLogResponse) SetBody(v *GetAgentWorkLogResponseBody) *GetAgentWorkLogResponse {
	s.Body = v
	return s
}

func (s *GetAgentWorkLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
