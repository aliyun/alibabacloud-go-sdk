// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAgentSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelAgentSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelAgentSessionResponse
	GetStatusCode() *int32
	SetBody(v *CancelAgentSessionResponseBody) *CancelAgentSessionResponse
	GetBody() *CancelAgentSessionResponseBody
}

type CancelAgentSessionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelAgentSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelAgentSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelAgentSessionResponse) GoString() string {
	return s.String()
}

func (s *CancelAgentSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelAgentSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelAgentSessionResponse) GetBody() *CancelAgentSessionResponseBody {
	return s.Body
}

func (s *CancelAgentSessionResponse) SetHeaders(v map[string]*string) *CancelAgentSessionResponse {
	s.Headers = v
	return s
}

func (s *CancelAgentSessionResponse) SetStatusCode(v int32) *CancelAgentSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelAgentSessionResponse) SetBody(v *CancelAgentSessionResponseBody) *CancelAgentSessionResponse {
	s.Body = v
	return s
}

func (s *CancelAgentSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
