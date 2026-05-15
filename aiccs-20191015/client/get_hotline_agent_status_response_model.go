// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineAgentStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotlineAgentStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotlineAgentStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetHotlineAgentStatusResponseBody) *GetHotlineAgentStatusResponse
	GetBody() *GetHotlineAgentStatusResponseBody
}

type GetHotlineAgentStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotlineAgentStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotlineAgentStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineAgentStatusResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotlineAgentStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotlineAgentStatusResponse) GetBody() *GetHotlineAgentStatusResponseBody {
	return s.Body
}

func (s *GetHotlineAgentStatusResponse) SetHeaders(v map[string]*string) *GetHotlineAgentStatusResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineAgentStatusResponse) SetStatusCode(v int32) *GetHotlineAgentStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotlineAgentStatusResponse) SetBody(v *GetHotlineAgentStatusResponseBody) *GetHotlineAgentStatusResponse {
	s.Body = v
	return s
}

func (s *GetHotlineAgentStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
