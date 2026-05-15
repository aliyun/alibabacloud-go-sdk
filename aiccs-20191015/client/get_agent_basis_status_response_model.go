// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentBasisStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentBasisStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentBasisStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentBasisStatusResponseBody) *GetAgentBasisStatusResponse
	GetBody() *GetAgentBasisStatusResponseBody
}

type GetAgentBasisStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentBasisStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentBasisStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentBasisStatusResponse) GoString() string {
	return s.String()
}

func (s *GetAgentBasisStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentBasisStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentBasisStatusResponse) GetBody() *GetAgentBasisStatusResponseBody {
	return s.Body
}

func (s *GetAgentBasisStatusResponse) SetHeaders(v map[string]*string) *GetAgentBasisStatusResponse {
	s.Headers = v
	return s
}

func (s *GetAgentBasisStatusResponse) SetStatusCode(v int32) *GetAgentBasisStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentBasisStatusResponse) SetBody(v *GetAgentBasisStatusResponseBody) *GetAgentBasisStatusResponse {
	s.Body = v
	return s
}

func (s *GetAgentBasisStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
