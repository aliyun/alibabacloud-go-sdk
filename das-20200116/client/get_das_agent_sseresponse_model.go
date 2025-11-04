// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDasAgentSSEResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDasAgentSSEResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDasAgentSSEResponse
	GetStatusCode() *int32
	SetBody(v *GetDasAgentSSEResponseBody) *GetDasAgentSSEResponse
	GetBody() *GetDasAgentSSEResponseBody
}

type GetDasAgentSSEResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDasAgentSSEResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDasAgentSSEResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDasAgentSSEResponse) GoString() string {
	return s.String()
}

func (s *GetDasAgentSSEResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDasAgentSSEResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDasAgentSSEResponse) GetBody() *GetDasAgentSSEResponseBody {
	return s.Body
}

func (s *GetDasAgentSSEResponse) SetHeaders(v map[string]*string) *GetDasAgentSSEResponse {
	s.Headers = v
	return s
}

func (s *GetDasAgentSSEResponse) SetStatusCode(v int32) *GetDasAgentSSEResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDasAgentSSEResponse) SetBody(v *GetDasAgentSSEResponseBody) *GetDasAgentSSEResponse {
	s.Body = v
	return s
}

func (s *GetDasAgentSSEResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
